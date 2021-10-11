package spotify

import (
   "bytes"
   "crypto/aes"
   "encoding/binary"
   "fmt"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "math"
   "sync"
)

const chunkSizeK = 32768

type client struct {
	subscriptions map[string][]chan response
	callbacks     map[string]callback
	inter      *internal
	cbMu          sync.Mutex
}

func createMercury(stream packetStream) *client {
	client := &client{
		callbacks:     make(map[string]callback),
		subscriptions: make(map[string][]chan response),
		inter: &internal{
			Pending: make(map[string]pending),
			Stream:  stream,
		},
	}
	return client
}

func (m *client) handle(cmd uint8, reader io.Reader) (err error) {
	resp, err := m.inter.parseResponse(cmd, reader)
	if err != nil {
		return
	}
	if resp != nil {
		if cmd == 0xb5 {
			chList, ok := m.subscriptions[resp.Uri]
			if ok {
				for _, ch := range chList {
					ch <- *resp
				}
			}
		} else {
			m.cbMu.Lock()
			cb, ok := m.callbacks[resp.SeqKey]
			delete(m.callbacks, resp.SeqKey) // no-op if element does not exist
			m.cbMu.Unlock()
			if ok {
				cb(*resp)
			}
		}
	}
	return

}

func (m *client) mercuryGet(url string) []byte {
   done := make(chan []byte)
   go m.request(
      request{
         Method:  "GET", Payload: [][]byte{}, Uri: url,
      },
      func(res response) {
         done <- res.combinePayload()
      },
   )
   result := <-done
   return result
}

func (m *client) mercuryGetProto(url string, result proto.Message) error {
   data := m.mercuryGet(url)
   return proto.Unmarshal(data, result)
}

func (m *client) request(req request, cb callback) (err error) {
   seq, err := m.inter.request(req)
   if err != nil {
      if cb != nil {
         cb(response{StatusCode: 500})
      }
      return err
   }
   m.cbMu.Lock()
   m.callbacks[string(seq)] = cb
   m.cbMu.Unlock()
   return nil
}

type player struct {
   chanLock    sync.Mutex
   channels    map[uint16]*channel
   mercury  *client
   nextChan    uint16
   seqChans    sync.Map
   stream   packetStream
}

func createPlayer(conn packetStream, client *client) *player {
	return &player{
		stream:   conn,
		mercury:  client,
		channels: map[uint16]*channel{},
		seqChans: sync.Map{},
		chanLock: sync.Mutex{},
		nextChan: 0,
	}
}

func (p *player) loadTrackKey(trackId []byte, fileId []byte) ([]byte, error) {
   seqInt, seq := p.mercury.inter.nextSeq()
   p.seqChans.Store(seqInt, make(chan []byte))
   req := buildKeyRequest(seq, trackId, fileId)
   err := p.stream.sendPacket(packetRequestKey, req)
   if err != nil {
   log.Println("Error while sending packet", err)
      return nil, err
   }
   channel, _ := p.seqChans.Load(seqInt)
   key := <-channel.(chan []byte)
   p.seqChans.Delete(seqInt)
   return key, nil
}

func (p *player) handleCmd(cmd byte, data []byte) {
	switch {
	case cmd == packetAesKey:
		dataReader := bytes.NewReader(data)
		var seqNum uint32
		binary.Read(dataReader, binary.BigEndian, &seqNum)

		if channel, ok := p.seqChans.Load(seqNum); ok {
			channel.(chan []byte) <- data[4:20]
		} else {
			fmt.Printf("[player] Unknown channel for audio key seqNum %d\n", seqNum)
		}

	case cmd == packetAesKeyError:
		// Audio key error
		fmt.Println("[player] Audio key error!")
		fmt.Printf("%x\n", data)

	case cmd == packetStreamChunkRes:
		var channel uint16
		dataReader := bytes.NewReader(data)
		binary.Read(dataReader, binary.BigEndian, &channel)

		if val, ok := p.channels[channel]; ok {
			val.handlePacket(data[2:])
		} else {
			fmt.Printf("Unknown channel!\n")
		}
	}
}

func (p *player) releaseChannel(channel *channel) {
	p.chanLock.Lock()
	delete(p.channels, channel.Num)
	p.chanLock.Unlock()
}

func (a *audioFile) Read(buf []byte) (int, error) {
   length := len(buf)
   outBufCursor := 0
   totalWritten := 0
   eof := false
   a.lock.RLock()
   size := a.size
   a.lock.RUnlock()
   // Offset the data start by the header, if needed
   if a.cursor == 0 {
      a.cursor += a.headerOffset()
   } else if uint32(a.cursor) >= size {
      // We're at the end
      return 0, io.EOF
   }
   chunkIdx := a.chunkIndexAtByte(a.cursor)
   for totalWritten < length {
      if chunkIdx >= a.totalChunks() {
         // We've reached the last chunk, so we can signal EOF
         eof = true
         break
      } else if !a.hasChunk(chunkIdx) {
         a.requestChunk(chunkIdx)
         break
      } else {
         // cursorEnd is the ending position in the output buffer. It is either
         // the current outBufCursor + the size of a chunk, in bytes, or the
         // length of the buffer, whichever is smallest.
         cursorEnd := min(outBufCursor+chunkByteSizeK, length)
         writtenLen := cursorEnd - outBufCursor
         // Calculate where our data cursor will end: either at the boundary of
         // the current chunk, or the end of the song itself
         dataCursorEnd := min(a.cursor+writtenLen, (chunkIdx+1)*chunkByteSizeK)
         dataCursorEnd = min(dataCursorEnd, int(a.size))
         writtenLen = dataCursorEnd - a.cursor
         if writtenLen <= 0 {
            // No more space in the output buffer, bail out
            break
         }
         // Copy into the output buffer, from the current outBufCursor, up to the
         // cursorEnd. The source is the current cursor inside the audio file,
         // up to the dataCursorEnd.
         copy(buf[outBufCursor:cursorEnd], a.data[a.cursor:dataCursorEnd])
         outBufCursor += writtenLen
         a.cursor += writtenLen
         totalWritten += writtenLen
         // Update our current chunk, if we need to
         chunkIdx = a.chunkIndexAtByte(a.cursor)
      }
   }
   // The only error we can return here, is if we reach the end of the stream
   if eof {
      return 0, io.EOF
   }
   return totalWritten, nil
}

func (a *audioFile) chunkIndexAtByte(byteIndex int) int {
   return int(math.Floor(float64(byteIndex) / float64(chunkSizeK) / 4.0))
}

func (a *audioFile) hasChunk(index int) bool {
	a.chunkLock.RLock()
	has, ok := a.chunks[index]
	a.chunkLock.RUnlock()

	return has && ok
}

func (a *audioFile) loadKey(trackId []byte) error {
	key, err := a.player.loadTrackKey(trackId, a.fileId)
	if err != nil {
		fmt.Printf("[audiofile] Unable to load key: %s\n", err)
		return err
	}

	a.cipher, err = aes.NewCipher(key)
	if err != nil {
		return err
	}

	return nil
}

func (a *audioFile) totalChunks() int {
	a.lock.RLock()
	size := a.size
	a.lock.RUnlock()
	return int(math.Ceil(float64(size) / float64(chunkSizeK) / 4.0))
}

func (a *audioFile) loadChunks() {
	// By default, we will load the track in the normal order. If we need to skip to a specific piece of audio,
	// we will prepend the chunks needed so that we load them as soon as possible. Since loadNextChunk will check
	// if a chunk is already loaded (using hasChunk), we won't be downloading the same chunk multiple times.

	// We can however only download the first chunk for now, as we have no idea how many chunks this track has. The
	// remaining chunks will be added once we get the headers with the file size.
	a.chunkLoadOrder = append(a.chunkLoadOrder, 0)

	go a.loadNextChunk()
}

func (a *audioFile) requestChunk(chunkIndex int) {
	a.chunkLock.RLock()

	// Check if we don't already have this chunk in the 2 next chunks requested
	if len(a.chunkLoadOrder) >= 1 && a.chunkLoadOrder[0] == chunkIndex ||
		len(a.chunkLoadOrder) >= 2 && a.chunkLoadOrder[1] == chunkIndex {
		a.chunkLock.RUnlock()
		return
	}

	a.chunkLock.RUnlock()

	// Set an artificial limit of 500 chunks to prevent overflows and buggy readers/seekers
	a.chunkLock.Lock()

	if len(a.chunkLoadOrder) < 500 {
		a.chunkLoadOrder = append([]int{chunkIndex}, a.chunkLoadOrder...)
	}

	a.chunkLock.Unlock()
}

func (a *audioFile) loadChunk(chunkIndex int) error {
   a.player.chanLock.Lock()
   channel := newChannel(a.player.nextChan, a.player.releaseChannel)
   a.player.nextChan++
   a.player.channels[channel.Num] = channel
   a.player.chanLock.Unlock()
   channel.OnHeader = a.onChannelHeader
   channel.OnData = a.onChannelData
   chunkOffsetStart := uint32(chunkIndex * chunkSizeK)
   chunkOffsetEnd := uint32((chunkIndex + 1) * chunkSizeK)
   err := a.player.stream.sendPacket(
      packetStreamChunk,
      buildAudioChunkRequest(
         channel.Num, a.fileId, chunkOffsetStart, chunkOffsetEnd,
      ),
   )
   if err != nil {
      return err
   }
   chunkSz := 0
   chunkData := make([]byte, chunkByteSizeK)
   for {
      chunk := <-a.responseChan
      chunkLen := len(chunk)
      if chunkLen > 0 {
         copy(chunkData[chunkSz:chunkSz+chunkLen], chunk)
         chunkSz += chunkLen
      } else {
         break
      }
   }
   a.putEncryptedChunk(chunkIndex, chunkData[0:chunkSz])
   return nil
}

func (a *audioFile) loadNextChunk() {
   a.chunkLock.Lock()
   if a.chunksLoading {
      // We are already loading a chunk, don't need to start another goroutine
      a.chunkLock.Unlock()
      return
   }
   a.chunksLoading = true
   chunkIndex := a.chunkLoadOrder[0]
   a.chunkLoadOrder = a.chunkLoadOrder[1:]
   a.chunkLock.Unlock()
   if !a.hasChunk(chunkIndex) {
      a.loadChunk(chunkIndex)
   }
   a.chunkLock.Lock()
   a.chunksLoading = false
   if len(a.chunkLoadOrder) > 0 {
      a.chunkLock.Unlock()
      a.loadNextChunk()
   } else {
      a.chunkLock.Unlock()
   }
}

func (a *audioFile) putEncryptedChunk(index int, data []byte) {
   byteIndex := index * chunkByteSizeK
   a.decrypter.decryptAudioWithBlock(
      index, a.cipher, data, a.data[byteIndex:byteIndex+len(data)],
   )
   a.chunkLock.Lock()
   a.chunks[index] = true
   a.chunkLock.Unlock()
}

func (a *audioFile) onChannelHeader(channel *channel, id byte, data *bytes.Reader) uint16 {
   read := uint16(0)
   if id == 0x3 {
      var size uint32
      binary.Read(data, binary.BigEndian, &size)
      size *= 4
      if a.size != size {
         a.lock.Lock()
         a.size = size
         a.lock.Unlock()
         if a.data == nil {
            a.data = make([]byte, size)
         }
         a.chunkLock.Lock()
         for i := 0; i < a.totalChunks(); i++ {
            a.chunkLoadOrder = append(a.chunkLoadOrder, i)
         }
         a.chunkLock.Unlock()
         // Re-launch the chunk loading system. It will check itself if another
         // goroutine is already loading chunks.
         go a.loadNextChunk()
      }
      // Return 4 bytes read
      read = 4
   }
   return read
}

func (a *audioFile) onChannelData(channel *channel, data []byte) uint16 {
	if data != nil {
		a.responseChan <- data

		return 0 // uint16(len(data))
	} else {
		a.responseChan <- []byte{}
		return 0
	}

}
