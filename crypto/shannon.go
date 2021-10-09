package crypto

import (
   "bytes"
   "encoding/binary"
   "io"
   "log"
   "sync"
)

type ShannonStream struct {
	sendNonce  uint32
	SendCipher shn_ctx
	RecvCipher shn_ctx

	recvNonce uint32
	Reader    io.Reader
	Writer    io.Writer

	Mutex *sync.Mutex
}

func SetKey(ctx *shn_ctx, key []uint8) {
	shn_key(ctx, key, len(key))

	nonce := make([]byte, 4)
	binary.BigEndian.PutUint32(nonce, 0)
	shn_nonce(ctx, nonce, len(nonce))
}

func (s *ShannonStream) SendPacket(cmd uint8, data []byte) (err error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	_, err = s.Write(cipherPacket(cmd, data))
	if err != nil {
		return
	}
	err = s.FinishSend()
	return
}

func cipherPacket(cmd uint8, data []byte) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, cmd)
	binary.Write(buf, binary.BigEndian, uint16(len(data)))
	buf.Write(data)
	return buf.Bytes()
}

func (s *ShannonStream) Encrypt(message string) []byte {
	messageBytes := []byte(message)
	return s.EncryptBytes(messageBytes)
}

func (s *ShannonStream) EncryptBytes(messageBytes []byte) []byte {
	shn_encrypt(&s.SendCipher, messageBytes, len(messageBytes))
	return messageBytes
}

func (s *ShannonStream) Decrypt(messageBytes []byte) []byte {
	shn_decrypt(&s.RecvCipher, messageBytes, len(messageBytes))
	return messageBytes
}

func (s *ShannonStream) WrapReader(reader io.Reader) {
	s.Reader = reader
}

func (s *ShannonStream) WrapWriter(writer io.Writer) {
	s.Writer = writer
}

func (s *ShannonStream) Read(p []byte) (n int, err error) {
   n, err = s.Reader.Read(p)
   //p = s.Decrypt(p)
   s.Decrypt(p)
   return n, err
}

func (s *ShannonStream) Write(p []byte) (n int, err error) {
	p = s.EncryptBytes(p)
	return s.Writer.Write(p)
}

func (s *ShannonStream) FinishSend() (err error) {
	count := 4
	mac := make([]byte, count)
	shn_finish(&s.SendCipher, mac, count)

	s.sendNonce += 1
	nonce := make([]uint8, 4)
	binary.BigEndian.PutUint32(nonce, s.sendNonce)
	shn_nonce(&s.SendCipher, nonce, len(nonce))

	_, err = s.Writer.Write(mac)
	return
}

func (s *ShannonStream) finishRecv() {
	count := 4

	mac := make([]byte, count)
	io.ReadFull(s.Reader, mac)

	mac2 := make([]byte, count)
	shn_finish(&s.RecvCipher, mac2, count)

	if !bytes.Equal(mac, mac2) {
		log.Println("received mac doesn't match")
	}

	s.recvNonce += 1
	nonce := make([]uint8, 4)
	binary.BigEndian.PutUint32(nonce, s.recvNonce)
	shn_nonce(&s.RecvCipher, nonce, len(nonce))
}

func (s *ShannonStream) RecvPacket() (cmd uint8, buf []byte, err error) {
	err = binary.Read(s, binary.BigEndian, &cmd)
	if err != nil {
		return
	}

	var size uint16
	err = binary.Read(s, binary.BigEndian, &size)
	if err != nil {
		return
	}

	if size > 0 {
		buf = make([]byte, size)
		_, err = io.ReadFull(s.Reader, buf)
		if err != nil {
			return
		}
		buf = s.Decrypt(buf)

	}
	s.finishRecv()

	return cmd, buf, err
}
