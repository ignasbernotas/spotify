//package connection
package crypto

import (
	"encoding/binary"
	"io"
	"sync"
)

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
	PacketSecretBlock    = 0x02
	PacketPing           = 0x04
	PacketStreamChunk    = 0x08
	PacketStreamChunkRes = 0x09
	PacketChannelError   = 0x0a
	PacketChannelAbort   = 0x0b
	PacketRequestKey     = 0x0c
	PacketAesKey         = 0x0d
	PacketAesKeyError    = 0x0e

	PacketImage       = 0x19
	PacketCountryCode = 0x1b

	PacketPong    = 0x49
	PacketPongAck = 0x4a
	PacketPause   = 0x4b

	PacketProductInfo   = 0x50
	PacketLegacyWelcome = 0x69

	PacketLicenseVersion = 0x76

	PacketLogin       = 0xab
	PacketAPWelcome   = 0xac
	PacketAuthFailure = 0xad

	PacketMercuryReq   = 0xb2
	PacketMercurySub   = 0xb3
	PacketMercuryUnsub = 0xb4
)

type PacketStream interface {
	SendPacket(cmd uint8, data []byte) (err error)
	RecvPacket() (cmd uint8, buf []byte, err error)
}
