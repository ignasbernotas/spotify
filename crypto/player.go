package crypto

import (
   "bytes"
   "encoding/binary"
   "encoding/json"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "sync"
)

type Client struct {
	subscriptions map[string][]chan Response
	callbacks     map[string]Callback
	internal      *Internal
	cbMu          sync.Mutex
}

func CreateMercury(stream PacketStream) *Client {
	client := &Client{
		callbacks:     make(map[string]Callback),
		subscriptions: make(map[string][]chan Response),
		internal: &Internal{
			pending: make(map[string]Pending),
			stream:  stream,
		},
	}
	return client
}

func (m *Client) Handle(cmd uint8, reader io.Reader) (err error) {
	response, err := m.internal.parseResponse(cmd, reader)
	if err != nil {
		return
	}
	if response != nil {
		if cmd == 0xb5 {
			chList, ok := m.subscriptions[response.Uri]
			if ok {
				for _, ch := range chList {
					ch <- *response
				}
			}
		} else {
			m.cbMu.Lock()
			cb, ok := m.callbacks[response.SeqKey]
			delete(m.callbacks, response.SeqKey) // no-op if element does not exist
			m.cbMu.Unlock()
			if ok {
				cb(*response)
			}
		}
	}
	return

}

func (m *Client) mercuryGet(url string) []byte {
	done := make(chan []byte)
	go m.Request(Request{
		Method:  "GET",
		Uri:     url,
		Payload: [][]byte{},
	}, func(res Response) {
		done <- res.CombinePayload()
	})

	result := <-done
	return result
}

func (m *Client) mercuryGetJson(url string, result interface{}) error {
   data := m.mercuryGet(url)
   return json.Unmarshal(data, result)
}

func (m *Client) mercuryGetProto(url string, result proto.Message) error {
   data := m.mercuryGet(url)
   return proto.Unmarshal(data, result)
}

func (m *Client) GetTrack(id string) (*pb.Track, error) {
   result := new(pb.Track)
   err := m.mercuryGetProto("hm://metadata/4/track/" + id, result)
   if err != nil {
      return nil, err
   }
   return result, nil
}

func (m *Client) Request(req Request, cb Callback) (err error) {
   seq, err := m.internal.request(req)
   if err != nil {
      // Call the callback with a 500 error-code so that the request doesn't
      // remain pending in case of error
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

func (m *Client) NextSeqWithInt() (uint32, []byte) {
	return m.internal.NextSeq()
}

type Player struct {
   chanLock    sync.Mutex
   channels    map[uint16]*Channel
   mercury  *Client
   nextChan    uint16
   seqChans    sync.Map
   stream   PacketStream
}

func CreatePlayer(conn PacketStream, client *Client) *Player {
	return &Player{
		stream:   conn,
		mercury:  client,
		channels: map[uint16]*Channel{},
		seqChans: sync.Map{},
		chanLock: sync.Mutex{},
		nextChan: 0,
	}
}

func (p *Player) LoadTrack(file *pb.AudioFile, trackId []byte) (*AudioFile, error) {
   // Allocate an AudioFile and a channel
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

func (p *Player) loadTrackKey(trackId []byte, fileId []byte) ([]byte, error) {
	seqInt, seq := p.mercury.NextSeqWithInt()

	p.seqChans.Store(seqInt, make(chan []byte))

	req := buildKeyRequest(seq, trackId, fileId)
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

func (p *Player) AllocateChannel() *Channel {
	p.chanLock.Lock()
	channel := NewChannel(p.nextChan, p.releaseChannel)
	p.nextChan++

	p.channels[channel.num] = channel
	p.chanLock.Unlock()

	return channel
}

func (p *Player) HandleCmd(cmd byte, data []byte) {
	switch {
	case cmd == PacketAesKey:
		// Audio key response
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
		// Audio data response
		var channel uint16
		dataReader := bytes.NewReader(data)
		binary.Read(dataReader, binary.BigEndian, &channel)

		// fmt.Printf("[player] Data on channel %d: %d bytes\n", channel, len(data[2:]))

		if val, ok := p.channels[channel]; ok {
			val.handlePacket(data[2:])
		} else {
			fmt.Printf("Unknown channel!\n")
		}
	}
}

func (p *Player) releaseChannel(channel *Channel) {
	p.chanLock.Lock()
	delete(p.channels, channel.num)
	p.chanLock.Unlock()
	// fmt.Printf("[player] Released channel %d\n", channel.num)
}
