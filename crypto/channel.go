package crypto

import (
   "bytes"
   "crypto/aes"
   "crypto/cipher"
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "encoding/binary"
   "errors"
   "log"
   "math"
   "math/big"
)

type headerFunc func(channel *Channel, id byte, data *bytes.Reader) uint16
type dataFunc func(channel *Channel, data []byte) uint16
type releaseFunc func(channel *Channel)

type Channel struct {
	num       uint16
	dataMode  bool
	onHeader  headerFunc
	onData    dataFunc
	onRelease releaseFunc
}

func NewChannel(num uint16, release releaseFunc) *Channel {
	return &Channel{
		num:       num,
		dataMode:  false,
		onRelease: release,
	}
}

func (c *Channel) handlePacket(data []byte) {
	dataReader := bytes.NewReader(data)

	if !c.dataMode {
		// Read the header
		// fmt.Printf("[channel] Reading in header mode, size=%d\n", dataReader.Len())

		length := uint16(0)
		var err error = nil
		for err == nil {
			err = binary.Read(dataReader, binary.BigEndian, &length)

			if err != nil {
				break
			}

			// fmt.Printf("[channel] Header part length: %d\n", length)

			if length > 0 {
				var headerId uint8
				binary.Read(dataReader, binary.BigEndian, &headerId)

				// fmt.Printf("[channel] Header ID: 0x%x\n", headerId)

				read := uint16(0)
				if c.onHeader != nil {
					read = c.onHeader(c, headerId, dataReader)
				}

				// Consume the remaining un-read data
				dataReader.Read(make([]byte, length-read))
			}
		}

		if c.onData != nil {
			// fmt.Printf("[channel] Switching channel to dataMode\n")
			c.dataMode = true
		} else {
			c.onRelease(c)
		}
	} else {
		// fmt.Printf("[channel] Reading in dataMode\n")

		if len(data) == 0 {
			if c.onData != nil {
				c.onData(c, nil)
			}

			c.onRelease(c)
		} else {
			if c.onData != nil {
				c.onData(c, data)
			}
		}
	}

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

func buildKeyRequest(seq []byte, trackId []byte, fileId []byte) []byte {
	buf := new(bytes.Buffer)

	buf.Write(fileId)
	buf.Write(trackId)
	buf.Write(seq)
	binary.Write(buf, binary.BigEndian, uint16(0x0000))

	return buf.Bytes()
}


var AUDIO_AESIV = []byte{0x72, 0xe0, 0x67, 0xfb, 0xdd, 0xcb, 0xcf, 0x77, 0xeb, 0xe8, 0xbc, 0x64, 0x3f, 0x63, 0x0d, 0x93}

type AudioFileDecrypter struct {
	ivDiff *big.Int
	ivInt  *big.Int
}

func CreateCipher(key []byte) cipher.Block {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	return block
}

func NewAudioFileDecrypter() *AudioFileDecrypter {
	return &AudioFileDecrypter{
		ivDiff: new(big.Int),
		ivInt:  new(big.Int),
	}
}

func (afd *AudioFileDecrypter) DecryptAudioWithBlock(index int, block cipher.Block, ciphertext []byte, plaintext []byte) []byte {
	length := len(ciphertext)
	// plaintext := bufferPool.Get().([]byte) // make([]byte, length)
	byteBaseOffset := index * kChunkSize * 4

	// The actual IV is the base IV + index*0x100, where index is the chunk index sized 1024 words (so each 4096 bytes
	// block has its own IV). As we are retrieving 32768 words (131072 bytes) to speed up network operations, we need
	// to process the data by 4096 bytes blocks to decrypt with the correct key.

	// We pre-calculate the base IV for the first chunk we are processing, then just proceed to add 0x100 at
	// every iteration.
	afd.ivInt.SetBytes(AUDIO_AESIV)
	afd.ivDiff.SetInt64(int64((byteBaseOffset / 4096) * 0x100))
	afd.ivInt.Add(afd.ivInt, afd.ivDiff)

	afd.ivDiff.SetInt64(int64(0x100))

	for i := 0; i < length; i += 4096 {
		// fmt.Printf("IV (chunk index %d): %x\n", chunkIndex, ivBytes)
		endBytes := int(math.Min(float64(i+4096), float64(length)))

		stream := cipher.NewCTR(block, afd.ivInt.Bytes())
		stream.XORKeyStream(plaintext[i:endBytes], ciphertext[i:endBytes])

		afd.ivInt.Add(afd.ivInt, afd.ivDiff)
	}

	return plaintext[0:length]
}


// BlobInfo is the structure holding authentication blob data. The blob is an
// encoded/encrypted byte array (encoded as base64), holding the encryption
// keys, the deviceId, and the username.
type BlobInfo struct {
	Username    string
	DecodedBlob string
}

func decodeBlob(blob64 string, client64 string, keys PrivateKeys) (string, error) {

	clientKey, err := base64.StdEncoding.DecodeString(client64)
	if err != nil {
		return "", err
	}

	blobBytes, err := base64.StdEncoding.DecodeString(blob64)
	if err != nil {
		return "", err
	}

	clientKey_be := new(big.Int)
	clientKey_be.SetBytes(clientKey)

	sharedKey := Powm(clientKey_be, keys.PrivateKey(), keys.Prime())
	iv := blobBytes[0:16]
	encryptedPart := blobBytes[16 : len(blobBytes)-20]
	ckSum := blobBytes[len(blobBytes)-20:]
	key := sha1.Sum(sharedKey.Bytes())
	base_key := key[:16]
	hash := hmac.New(sha1.New, base_key)

	hash.Write([]byte("checksum"))
	checksum_key := hash.Sum(nil)
	hash.Reset()

	hash.Write([]byte("encryption"))
	encryption_key := hash.Sum(nil)
	hash.Reset()

	macHash := hmac.New(sha1.New, checksum_key)
	macHash.Write(encryptedPart)
	mac := macHash.Sum(nil)

	if !bytes.Equal(mac, ckSum) {
		log.Println("add user error, mac doesn't match")
		return "", errors.New("mac mismatch")
	}

	block, _ := aes.NewCipher(encryption_key[0:16])
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(encryptedPart, encryptedPart)

	return string(encryptedPart), nil
}

func decryptBlob(blob []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	bs := block.BlockSize()
	if len(blob)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	plaintext := make([]byte, len(blob))

	plain := plaintext
	for len(blob) > 0 {
		block.Decrypt(plaintext, blob)
		plaintext = plaintext[bs:]
		blob = blob[bs:]
	}

	l := len(plain)
	for i := 0; i < l-0x10; i++ {
		plain[l-i-1] = plain[l-i-1] ^ plain[l-i-0x11]
	}

	return plain
}
