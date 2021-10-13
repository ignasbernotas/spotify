package spotify

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "encoding/binary"
   "fmt"
   "io"
   "math/big"
   "sync"
   cryptoRand "crypto/rand"
)

type request struct {
   contentType string
   method      string
   payload     [][]byte
   uri         string
}

type response struct {
   headerData []byte
   payload    [][]byte
   seqKey     string
   statusCode int32
   uri        string
}

const chunkByteSizeK = chunkSizeK * 4

// Number of bytes to skip at the beginning of the file
const oggSkipBytesK = 167

func encodeMercuryHead(seq []byte, partsLength uint16, flags uint8) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   err := binary.Write(buf, binary.BigEndian, uint16(len(seq)))
   if err != nil {
      return nil, err
   }
   if _, err := buf.Write(seq); err != nil {
      return nil, err
   }
   if err := binary.Write(buf, binary.BigEndian, uint8(flags)); err != nil {
      return nil, err
   }
   if err := binary.Write(buf, binary.BigEndian, partsLength); err != nil {
      return nil, err
   }
   return buf, nil
}

func handleHead(reader io.Reader) ([]byte, uint8, uint16, error) {
   var seqLength uint16
   err := binary.Read(reader, binary.BigEndian, &seqLength)
   if err != nil {
      return nil, 0, 0, err
   }
   seq := make([]byte, seqLength)
   if _, err := io.ReadFull(reader, seq); err != nil {
      return nil, 0, 0, fmt.Errorf("read seq %v", err)
   }
   var flags uint8
   if err := binary.Read(reader, binary.BigEndian, &flags); err != nil {
      return nil, 0, 0, fmt.Errorf("read flags %v", err)
   }
   var count uint16
   if err := binary.Read(reader, binary.BigEndian, &count); err != nil {
      return nil, 0, 0, fmt.Errorf("read count %v", err)
   }
   return seq, flags, count, nil
}

func min(a, b int) int {
   if a < b {
      return a
   }
   return b
}

func parsePart(reader io.Reader) ([]byte, error) {
   var size uint16
   binary.Read(reader, binary.BigEndian, &size)
   buf := make([]byte, size)
   _, err := io.ReadFull(reader, buf)
   if err != nil {
      return nil, err
   }
   return buf, nil
}

type callback func(response)

type internal struct {
   Pending map[string]pending
   Stream  packetStream
   nextSequence uint32
   seqLock sync.Mutex
}

func (m *internal) nextSeq() (uint32, []byte) {
   m.seqLock.Lock()
   seq := make([]byte, 4)
   seqInt := m.nextSequence
   binary.BigEndian.PutUint32(seq, seqInt)
   m.nextSequence += 1
   m.seqLock.Unlock()
   return seqInt, seq
}

func (m *internal) request(req request) (seqKey string, err error) {
   _, seq := m.nextSeq()
   data, err := encodeRequest(seq, req)
   if err != nil {
      return "", err
   }
   var cmd uint8
   switch {
   case req.method == "SUB":
      cmd = 0xb3
   case req.method == "UNSUB":
      cmd = 0xb4
   default:
      cmd = 0xb2
   }
   err = m.Stream.sendPacket(cmd, data)
   if err != nil {
      return "", err
   }
   return string(seq), nil
}

type pending struct {
	parts   [][]byte
	partial []byte
}

func (res *response) combinePayload() []byte {
	body := make([]byte, 0)
	for _, p := range res.payload {
		body = append(body, p...)
	}
	return body
}

const (
   AudioFile_OGG_VORBIS_96   = 0
   AudioFile_OGG_VORBIS_160  = 1
   AudioFile_OGG_VORBIS_320  = 2
)

func generateDeviceId(name string) string {
   hash := sha1.Sum([]byte(name))
   return base64.StdEncoding.EncodeToString(hash[:])
}

func powm(base, exp, modulus *big.Int) *big.Int {
	exp2 := big.NewInt(0).SetBytes(exp.Bytes())
	base2 := big.NewInt(0).SetBytes(base.Bytes())
	modulus2 := big.NewInt(0).SetBytes(modulus.Bytes())
	zero := big.NewInt(0)
	result := big.NewInt(1)
	temp := new(big.Int)

	for zero.Cmp(exp2) != 0 {
		if temp.Rem(exp2, big.NewInt(2)).Cmp(zero) != 0 {
			result = result.Mul(result, base2)
			result = result.Rem(result, modulus2)
		}
		exp2 = exp2.Rsh(exp2, 1)
		base2 = base2.Mul(base2, base2)
		base2 = base2.Rem(base2, modulus2)
	}
	return result
}

func randomVec(count int) ([]byte, error) {
   b := make([]byte, count)
   _, err := cryptoRand.Read(b)
   if err != nil {
      return nil, err
   }
   return b, nil
}

type blobInfo struct {
   Username    string
   DecodedBlob string
}

type privateKeys struct {
   clientNonce []byte
   generator   *big.Int
   prime       *big.Int
   privateKey *big.Int
   publicKey  *big.Int
}

func (p *privateKeys) addRemoteKey(remote []byte, clientPacket []byte, serverPacket []byte) sharedKeys {
	remote_be := new(big.Int)
	remote_be.SetBytes(remote)
	shared_key := powm(remote_be, p.privateKey, p.prime)
	data := make([]byte, 0, 100)
	mac := hmac.New(sha1.New, shared_key.Bytes())

	for i := 1; i < 6; i++ {
		mac.Write(clientPacket)
		mac.Write(serverPacket)
		mac.Write([]byte{uint8(i)})
		data = append(data, mac.Sum(nil)...)
		mac.Reset()
	}

	mac = hmac.New(sha1.New, data[0:0x14])
	mac.Write(clientPacket)
	mac.Write(serverPacket)

	return sharedKeys{
		challenge: mac.Sum(nil),
		sendKey:   data[0x14:0x34],
		recvKey:   data[0x34:0x54],
	}
}

type sharedKeys struct {
	challenge []byte
	sendKey   []byte
	recvKey   []byte
}
