package spotify

import (
   "bytes"
   "crypto/aes"
   "crypto/cipher"
   "encoding/binary"
   "encoding/json"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "math"
   "sync"
)

type client struct {
	subscriptions map[string][]chan Response
	callbacks     map[string]Callback
	inter      *Internal
	cbMu          sync.Mutex
}

func createMercury(stream PacketStream) *client {
	client := &client{
		callbacks:     make(map[string]Callback),
		subscriptions: make(map[string][]chan Response),
		inter: &Internal{
			Pending: make(map[string]Pending),
			Stream:  stream,
		},
	}
	return client
}

func (m *client) handle(cmd uint8, reader io.Reader) (err error) {
	resp, err := m.inter.ParseResponse(cmd, reader)
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
      Request{
         Method:  "GET", Payload: [][]byte{}, Uri: url,
      },
      func(res Response) {
         done <- res.CombinePayload()
      },
   )
   result := <-done
   return result
}

func (m *client) mercuryGetJson(url string, result interface{}) error {
   data := m.mercuryGet(url)
   return json.Unmarshal(data, result)
}

func (m *client) mercuryGetProto(url string, result proto.Message) error {
   data := m.mercuryGet(url)
   return proto.Unmarshal(data, result)
}

func (m *client) getTrack(id string) (*pb.Track, error) {
   result := new(pb.Track)
   err := m.mercuryGetProto("hm://metadata/4/track/" + id, result)
   if err != nil {
      return nil, err
   }
   return result, nil
}

func (m *client) request(req Request, cb Callback) (err error) {
   seq, err := m.inter.Request(req)
   if err != nil {
      if cb != nil {
         cb(Response{StatusCode: 500})
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
   channels    map[uint16]*Channel
   mercury  *client
   nextChan    uint16
   seqChans    sync.Map
   stream   PacketStream
}

func CreatePlayer(conn PacketStream, client *client) *player {
	return &player{
		stream:   conn,
		mercury:  client,
		channels: map[uint16]*Channel{},
		seqChans: sync.Map{},
		chanLock: sync.Mutex{},
		nextChan: 0,
	}
}

func (p *player) LoadTrack(file *pb.AudioFile, trackId []byte) (*AudioFile, error) {
   audioFile := newAudioFileWithIdAndFormat(file.FileId, file.GetFormat(), p)
   // Start loading the audio key
   err := audioFile.loadKey(trackId)
   if err != nil {
      return nil, err
   }
   // Then start loading the audio itself
   audioFile.loadChunks()
   return audioFile, nil
}

func (p *player) loadTrackKey(trackId []byte, fileId []byte) ([]byte, error) {
   seqInt, seq := p.mercury.inter.NextSeq()
   p.seqChans.Store(seqInt, make(chan []byte))
   req := BuildKeyRequest(seq, trackId, fileId)
   err := p.stream.SendPacket(PacketRequestKey, req)
   if err != nil {
   log.Println("Error while sending packet", err)
      return nil, err
   }
   channel, _ := p.seqChans.Load(seqInt)
   key := <-channel.(chan []byte)
   p.seqChans.Delete(seqInt)
   return key, nil
}

func (p *player) AllocateChannel() *Channel {
	p.chanLock.Lock()
	channel := NewChannel(p.nextChan, p.releaseChannel)
	p.nextChan++

	p.channels[channel.Num] = channel
	p.chanLock.Unlock()
	return channel
}

func (p *player) HandleCmd(cmd byte, data []byte) {
	switch {
	case cmd == PacketAesKey:
		dataReader := bytes.NewReader(data)
		var seqNum uint32
		binary.Read(dataReader, binary.BigEndian, &seqNum)

		if channel, ok := p.seqChans.Load(seqNum); ok {
			channel.(chan []byte) <- data[4:20]
		} else {
			fmt.Printf("[player] Unknown channel for audio key seqNum %d\n", seqNum)
		}

	case cmd == PacketAesKeyError:
		// Audio key error
		fmt.Println("[player] Audio key error!")
		fmt.Printf("%x\n", data)

	case cmd == PacketStreamChunkRes:
		var channel uint16
		dataReader := bytes.NewReader(data)
		binary.Read(dataReader, binary.BigEndian, &channel)

		if val, ok := p.channels[channel]; ok {
			val.HandlePacket(data[2:])
		} else {
			fmt.Printf("Unknown channel!\n")
		}
	}
}

func (p *player) releaseChannel(channel *Channel) {
	p.chanLock.Lock()
	delete(p.channels, channel.Num)
	p.chanLock.Unlock()
}

type AudioFile struct {
	size           uint32
	lock           sync.RWMutex
	format         pb.AudioFile_Format
	fileId         []byte
	player         *player
	cipher         cipher.Block
	decrypter      *AudioFileDecrypter
	responseChan   chan []byte
	chunkLock      sync.RWMutex
	chunkLoadOrder []int
	data           []byte
	cursor         int
	chunks         map[int]bool
	chunksLoading  bool
}

func newAudioFileWithIdAndFormat(fileId []byte, format pb.AudioFile_Format, player *player) *AudioFile {
	return &AudioFile{
		player:        player,
		fileId:        fileId,
		format:        format,
		decrypter:     NewAudioFileDecrypter(),
		size:          ChunkSizeK, // Set an initial size to fetch the first chunk regardless of the actual size
		responseChan:  make(chan []byte),
		chunks:        map[int]bool{},
		chunkLock:     sync.RWMutex{},
		chunksLoading: false,
	}
}

func (a *AudioFile) Read(buf []byte) (int, error) {
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
		// fmt.Printf("[audiofile] Cursor: %d, len: %d, matching chunk %d\n", a.cursor, length, chunkIdx)

		if chunkIdx >= a.totalChunks() {
			// We've reached the last chunk, so we can signal EOF
			eof = true
			break
		} else if !a.hasChunk(chunkIdx) {
			a.requestChunk(chunkIdx)
			// fmt.Printf("[audiofile] Doesn't have chunk %d yet, queuing\n", chunkIdx)
			break
		} else {
			// cursorEnd is the ending position in the output buffer. It is either the current outBufCursor + the size
			// of a chunk, in bytes, or the length of the buffer, whichever is smallest.
			cursorEnd := Min(outBufCursor+ChunkByteSizeK, length)
			writtenLen := cursorEnd - outBufCursor

			// Calculate where our data cursor will end: either at the boundary of the current chunk, or the end
			// of the song itself
			dataCursorEnd := Min(a.cursor+writtenLen, (chunkIdx+1)*ChunkByteSizeK)
			dataCursorEnd = Min(dataCursorEnd, int(a.size))

			writtenLen = dataCursorEnd - a.cursor

			if writtenLen <= 0 {
				// No more space in the output buffer, bail out
				break
			}

			// Copy into the output buffer, from the current outBufCursor, up to the cursorEnd. The source is the
			// current cursor inside the audio file, up to the dataCursorEnd.
			copy(buf[outBufCursor:cursorEnd], a.data[a.cursor:dataCursorEnd])
			outBufCursor += writtenLen
			a.cursor += writtenLen
			totalWritten += writtenLen

			// Update our current chunk, if we need to
			chunkIdx = a.chunkIndexAtByte(a.cursor)
		}
	}

	// The only error we can return here, is if we reach the end of the stream
	var err error
	if eof {
		err = io.EOF
	}

	return totalWritten, err
}

func (a *AudioFile) headerOffset() int {
	switch {
	case a.format == pb.AudioFile_OGG_VORBIS_96 || a.format == pb.AudioFile_OGG_VORBIS_160 ||
		a.format == pb.AudioFile_OGG_VORBIS_320:
		return OggSkipBytesK

	default:
		return 0
	}
}

func (a *AudioFile) chunkIndexAtByte(byteIndex int) int {
	return int(math.Floor(float64(byteIndex) / float64(ChunkSizeK) / 4.0))
}

func (a *AudioFile) hasChunk(index int) bool {
	a.chunkLock.RLock()
	has, ok := a.chunks[index]
	a.chunkLock.RUnlock()

	return has && ok
}

func (a *AudioFile) loadKey(trackId []byte) error {
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

func (a *AudioFile) totalChunks() int {
	a.lock.RLock()
	size := a.size
	a.lock.RUnlock()
	return int(math.Ceil(float64(size) / float64(ChunkSizeK) / 4.0))
}

func (a *AudioFile) loadChunks() {
	// By default, we will load the track in the normal order. If we need to skip to a specific piece of audio,
	// we will prepend the chunks needed so that we load them as soon as possible. Since loadNextChunk will check
	// if a chunk is already loaded (using hasChunk), we won't be downloading the same chunk multiple times.

	// We can however only download the first chunk for now, as we have no idea how many chunks this track has. The
	// remaining chunks will be added once we get the headers with the file size.
	a.chunkLoadOrder = append(a.chunkLoadOrder, 0)

	go a.loadNextChunk()
}

func (a *AudioFile) requestChunk(chunkIndex int) {
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

func (a *AudioFile) loadChunk(chunkIndex int) error {
   chunkData := make([]byte, ChunkByteSizeK)
   channel := a.player.AllocateChannel()
   channel.OnHeader = a.onChannelHeader
   channel.OnData = a.onChannelData
   chunkOffsetStart := uint32(chunkIndex * ChunkSizeK)
   chunkOffsetEnd := uint32((chunkIndex + 1) * ChunkSizeK)
   err := a.player.stream.SendPacket(
      PacketStreamChunk,
      BuildAudioChunkRequest(
         channel.Num, a.fileId, chunkOffsetStart, chunkOffsetEnd,
      ),
   )
   if err != nil {
   return err
   }
   chunkSz := 0
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

func (a *AudioFile) loadNextChunk() {
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

func (a *AudioFile) putEncryptedChunk(index int, data []byte) {
	byteIndex := index * ChunkByteSizeK
	a.decrypter.DecryptAudioWithBlock(index, a.cipher, data, a.data[byteIndex:byteIndex+len(data)])

	a.chunkLock.Lock()
	a.chunks[index] = true
	a.chunkLock.Unlock()
}

func (a *AudioFile) onChannelHeader(channel *Channel, id byte, data *bytes.Reader) uint16 {
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

func (a *AudioFile) onChannelData(channel *Channel, data []byte) uint16 {
	if data != nil {
		a.responseChan <- data

		return 0 // uint16(len(data))
	} else {
		a.responseChan <- []byte{}
		return 0
	}

}
