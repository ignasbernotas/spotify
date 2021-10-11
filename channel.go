package spotify

import (
   "bytes"
   "crypto/cipher"
   "encoding/binary"
   "io"
   "math"
   "math/big"
   "sync"
)

const (
   packetStreamChunk    = 0x08
   packetRequestKey     = 0x0c
   packetAesKey         = 0x0d
   packetAesKeyError    = 0x0e
   packetStreamChunkRes = 0x09
   packetLogin       = 0xab
   packetAuthFailure = 0xad
   packetAPWelcome   = 0xac
   packetPing           = 0x04
   packetPong    = 0x49
   packetPongAck = 0x4a
   packetCountryCode = 0x1b
   packetSecretBlock    = 0x02
   packetLegacyWelcome = 0x69
   packetProductInfo   = 0x50
   packetLicenseVersion = 0x76
)

var audio_AESIV = []byte{
   0x72, 0xe0, 0x67, 0xfb, 0xdd, 0xcb, 0xcf, 0x77,
   0xeb, 0xe8, 0xbc, 0x64, 0x3f, 0x63, 0x0d, 0x93,
}

func buildAudioChunkRequest(channel uint16, fileId []byte, start uint32, end uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, channel)
	binary.Write(buf, binary.BigEndian, uint8(0x0))
	binary.Write(buf, binary.BigEndian, uint8(0x1))
	binary.Write(buf, binary.BigEndian, uint16(0x0000))
	binary.Write(buf, binary.BigEndian, uint32(0x00000000))
	binary.Write(buf, binary.BigEndian, uint32(0x00009C40))
	binary.Write(buf, binary.BigEndian, uint32(0x00020000))
	buf.Write(fileId)
	binary.Write(buf, binary.BigEndian, start)
	binary.Write(buf, binary.BigEndian, end)

	return buf.Bytes()
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

func buildKeyRequest(seq []byte, trackId []byte, fileId []byte) []byte {
	buf := new(bytes.Buffer)

	buf.Write(fileId)
	buf.Write(trackId)
	buf.Write(seq)
	binary.Write(buf, binary.BigEndian, uint16(0x0000))

	return buf.Bytes()
}

type audioFileDecrypter struct {
	ivDiff *big.Int
	ivInt  *big.Int
}

func newAudioFileDecrypter() *audioFileDecrypter {
   return &audioFileDecrypter{
      ivDiff: new(big.Int), ivInt: new(big.Int),
   }
}

func (afd *audioFileDecrypter) decryptAudioWithBlock(index int, block cipher.Block, ciphertext []byte, plaintext []byte) []byte {
   length := len(ciphertext)
   byteBaseOffset := index * chunkSizeK * 4
   // The actual IV is the base IV + index*0x100, where index is the chunk
   // index sized 1024 words (so each 4096 bytes block has its own IV). As we
   // are retrieving 32768 words (131072 bytes) to speed up network operations,
   // we need to process the data by 4096 bytes blocks to decrypt with the
   // correct key. We pre-calculate the base IV for the first chunk we are
   // processing, then just proceed to add 0x100 at every iteration.
   afd.ivInt.SetBytes(audio_AESIV)
   afd.ivDiff.SetInt64(int64((byteBaseOffset / 4096) * 0x100))
   afd.ivInt.Add(afd.ivInt, afd.ivDiff)
   afd.ivDiff.SetInt64(int64(0x100))
   for i := 0; i < length; i += 4096 {
      endBytes := int(math.Min(float64(i+4096), float64(length)))
      stream := cipher.NewCTR(block, afd.ivInt.Bytes())
      stream.XORKeyStream(plaintext[i:endBytes], ciphertext[i:endBytes])
      afd.ivInt.Add(afd.ivInt, afd.ivDiff)
   }
   return plaintext[0:length]
}
type channel struct {
	Num       uint16
	dataMode  bool
	OnHeader  headerFunc
	OnData    dataFunc
	onRelease releaseFunc
}

func newChannel(num uint16, release releaseFunc) *channel {
	return &channel{
		Num:       num,
		dataMode:  false,
		onRelease: release,
	}
}

func (c *channel) handlePacket(data []byte) {
	dataReader := bytes.NewReader(data)

	if !c.dataMode {
		// Read the header
		length := uint16(0)
		var err error = nil
		for err == nil {
			err = binary.Read(dataReader, binary.BigEndian, &length)

			if err != nil {
				break
			}
			if length > 0 {
				var headerId uint8
				binary.Read(dataReader, binary.BigEndian, &headerId)
				read := uint16(0)
				if c.OnHeader != nil {
					read = c.OnHeader(c, headerId, dataReader)
				}

				// Consume the remaining un-read data
				dataReader.Read(make([]byte, length-read))
			}
		}

		if c.OnData != nil {
			c.dataMode = true
		} else {
			c.onRelease(c)
		}
	} else {
		if len(data) == 0 {
			if c.OnData != nil {
				c.OnData(c, nil)
			}

			c.onRelease(c)
		} else {
			if c.OnData != nil {
				c.OnData(c, data)
			}
		}
	}

}

type dataFunc func(channel *channel, data []byte) uint16

type headerFunc func(channel *channel, id byte, data *bytes.Reader) uint16

type packetStream interface {
   recvPacket() (cmd uint8, buf []byte, err error)
   sendPacket(cmd uint8, data []byte) error
}

func createStream(keys sharedKeys, conn plainConnection) packetStream {
	s := &shannonStream{
		Reader: conn.Reader,
		Writer: conn.Writer,
		Mutex:  &sync.Mutex{},
	}
	setKey(&s.RecvCipher, keys.recvKey)
	setKey(&s.SendCipher, keys.sendKey)
	return s
}

type plainConnection struct {
	Writer io.Writer
	Reader io.Reader
	mutex  *sync.Mutex
}

func makePlainConnection(reader io.Reader, writer io.Writer) plainConnection {
	return plainConnection{
		Reader: reader,
		Writer: writer,
		mutex:  &sync.Mutex{},
	}
}

func (p *plainConnection) sendPrefixPacket(prefix []byte, data []byte) ([]byte, error) {
   packet := makePacketPrefix(prefix, data)
   p.mutex.Lock()
   _, err := p.Writer.Write(packet)
   if err != nil {
      return nil, err
   }
   p.mutex.Unlock()
   return packet, nil
}

func (p *plainConnection) recvPacket() (buf []byte, err error) {
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

type releaseFunc func(channel *channel)
