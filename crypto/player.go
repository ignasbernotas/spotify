package crypto

import (
   "bytes"
   "encoding/binary"
   "fmt"
   "github.com/89z/spotify/pb"
   "io"
   "log"
   "sync"
)

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


// PlainConnection represents an unencrypted connection to a Spotify AP
type PlainConnection struct {
	Writer io.Writer
	Reader io.Reader
	mutex  *sync.Mutex
}

func makePacketPrefix(prefix []byte, data []byte) []byte {
	size := len(prefix) + 4 + len(data)
	buf := make([]byte, 0, size)
	buf = append(buf, prefix...)
	sizeBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeBuf, uint32(size))
	buf = append(buf, sizeBuf...)
	return append(buf, data...)
}

func MakePlainConnection(reader io.Reader, writer io.Writer) PlainConnection {
	return PlainConnection{
		Reader: reader,
		Writer: writer,
		mutex:  &sync.Mutex{},
	}
}

func (p *PlainConnection) SendPrefixPacket(prefix []byte, data []byte) (packet []byte, err error) {
	packet = makePacketPrefix(prefix, data)

	p.mutex.Lock()
	_, err = p.Writer.Write(packet)
	p.mutex.Unlock()

	return
}

func (p *PlainConnection) RecvPacket() (buf []byte, err error) {
	var size uint32
	err = binary.Read(p.Reader, binary.BigEndian, &size)
	if err != nil {
		return
	}
	buf = make([]byte, size)
	binary.BigEndian.PutUint32(buf, size)
	_, err = io.ReadFull(p.Reader, buf[4:])
	if err != nil {
		return
	}
	return buf, nil
}

const (
   PacketStreamChunk    = 0x08
   PacketRequestKey     = 0x0c
   PacketAesKey         = 0x0d
   PacketAesKeyError    = 0x0e
   PacketStreamChunkRes = 0x09
   PacketLogin       = 0xab
   PacketAuthFailure = 0xad
   PacketAPWelcome   = 0xac
   PacketPing           = 0x04
   PacketPong    = 0x49
   PacketPongAck = 0x4a
   PacketCountryCode = 0x1b
   PacketSecretBlock    = 0x02
   PacketLegacyWelcome = 0x69
   PacketProductInfo   = 0x50
   PacketLicenseVersion = 0x76
)

type PacketStream interface {
	SendPacket(cmd uint8, data []byte) (err error)
	RecvPacket() (cmd uint8, buf []byte, err error)
}
