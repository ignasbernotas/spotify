package spotify

import (
   "bytes"
   "encoding/binary"
   "io"
   "log"
   "sync"
)

type shannonStream struct {
   mutex *sync.Mutex
   reader    io.Reader
   recvCipher shn_ctx
   recvNonce uint32
   sendCipher shn_ctx
   sendNonce  uint32
   writer    io.Writer
}

func setKey(ctx *shn_ctx, key []uint8) {
	shn_key(ctx, key, len(key))
	nonce := make([]byte, 4)
	binary.BigEndian.PutUint32(nonce, 0)
	shn_nonce(ctx, nonce, len(nonce))
}

func (s *shannonStream) sendPacket(cmd uint8, data []byte) error {
   s.mutex.Lock()
   defer s.mutex.Unlock()
   _, err := s.Write(cipherPacket(cmd, data))
   if err != nil {
      return err
   }
   return s.finishSend()
}

func cipherPacket(cmd uint8, data []byte) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, cmd)
	binary.Write(buf, binary.BigEndian, uint16(len(data)))
	buf.Write(data)
	return buf.Bytes()
}

func (s *shannonStream) encryptBytes(messageBytes []byte) []byte {
   shn_encrypt(&s.sendCipher, messageBytes, len(messageBytes))
   return messageBytes
}

func (s *shannonStream) decrypt(messageBytes []byte) []byte {
	shn_decrypt(&s.recvCipher, messageBytes, len(messageBytes))
	return messageBytes
}

func (s *shannonStream) Read(p []byte) (int, error) {
   n, err := s.reader.Read(p)
   if err != nil {
      return 0, err
   }
   s.decrypt(p)
   return n, nil
}

func (s *shannonStream) Write(p []byte) (int, error) {
   p = s.encryptBytes(p)
   return s.writer.Write(p)
}

func (s *shannonStream) finishSend() (err error) {
	count := 4
	mac := make([]byte, count)
	shn_finish(&s.sendCipher, mac, count)

	s.sendNonce += 1
	nonce := make([]uint8, 4)
	binary.BigEndian.PutUint32(nonce, s.sendNonce)
	shn_nonce(&s.sendCipher, nonce, len(nonce))

	_, err = s.writer.Write(mac)
	return
}

func (s *shannonStream) finishRecv() {
	count := 4

	mac := make([]byte, count)
	io.ReadFull(s.reader, mac)

	mac2 := make([]byte, count)
	shn_finish(&s.recvCipher, mac2, count)

	if !bytes.Equal(mac, mac2) {
		log.Println("received mac doesn't match")
	}

	s.recvNonce += 1
	nonce := make([]uint8, 4)
	binary.BigEndian.PutUint32(nonce, s.recvNonce)
	shn_nonce(&s.recvCipher, nonce, len(nonce))
}

func (s *shannonStream) recvPacket() (cmd uint8, buf []byte, err error) {
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
		_, err = io.ReadFull(s.reader, buf)
		if err != nil {
			return
		}
		buf = s.decrypt(buf)

	}
	s.finishRecv()

	return cmd, buf, err
}

type shn_ctx struct {
   R     [num]uint32
   crc   [num]uint32
   initR [num]uint32
   konst uint32
   mbuf  uint32
   nbuf  int
   sbuf  uint32
}

const (
   initkonst uint32 = 0x6996c53a
   keyP int = 13
   num int = 16
)

// some useful macros -- machine independent little-endian

func toByte(x uint32, i int) uint8 {
	return uint8((x >> uint(8*i)) & 0xFF)
}

func rotl(w uint32, x int) uint32 {
	return w<<uint(x) | (w&0xffffffff)>>uint(32-x)
}

func byte2word(b []byte) uint32 {
	return (uint32(b[3])&0xFF)<<24 | (uint32(b[2])&0xFF)<<16 | (uint32(b[1])&0xFF)<<8 | (uint32(b[0]) & 0xFF)
}

func word2byte(w uint32, b []byte) {
	b[3] = byte(toByte(w, 3))
	b[2] = byte(toByte(w, 2))
	b[1] = byte(toByte(w, 1))
	b[0] = byte(toByte(w, 0))
}

// Nonlinear transform (sbox) of a word. There are two slightly different
// combinations.
func sbox1(w uint32) uint32 {
	w ^= rotl(w, 5) | rotl(w, 7)
	w ^= rotl(w, 19) | rotl(w, 22)
	return w
}

func sbox2(w uint32) uint32 {
	w ^= rotl(w, 7) | rotl(w, 22)
	w ^= rotl(w, 5) | rotl(w, 19)
	return w
}

// cycle the contents of the register and calculate output word in c->sbuf
func cycle(c *shn_ctx) {
	var t uint32
	var i int

	/* nonlinear feedback function */
	t = c.R[12] ^ c.R[13] ^ c.konst

	t = sbox1(t) ^ rotl(c.R[0], 1)

	/* shift register */
	for i = 1; i < num; i++ {
		c.R[i-1] = c.R[i]
	}
	c.R[num-1] = t
	t = sbox2(c.R[2] ^ c.R[15])
	c.R[0] ^= t
	c.sbuf = t ^ c.R[8] ^ c.R[12]
}

func crcfunc(c *shn_ctx, i uint32) {
   var t uint32
   var j int
   t = c.crc[0] ^ c.crc[2] ^ c.crc[15] ^ i
   for j = 1; j < num; j++ {
      c.crc[j-1] = c.crc[j]
   }
   c.crc[num-1] = t
}

func macfunc(c *shn_ctx, i uint32) {
	crcfunc(c, i)
	c.R[keyP] ^= i
}

// initialise to known state
func shn_initstate(c *shn_ctx) {
	var i int

	/* Register initialised to Fibonacci numbers; Counter zeroed. */
	c.R[0] = 1

	c.R[1] = 1
	for i = 2; i < num; i++ {
		c.R[i] = c.R[i-1] + c.R[i-2]
	}
	c.konst = initkonst
}

// Save the current register state
func shn_savestate(c *shn_ctx) {
	var i int

	for i = 0; i < num; i++ {
		c.initR[i] = c.R[i]
	}
}

// initialise to previously saved register state
func shn_reloadstate(c *shn_ctx) {
	var i int

	for i = 0; i < num; i++ {
		c.R[i] = c.initR[i]
	}
}

// Initialise "konst"
func shn_genkonst(c *shn_ctx) {
	c.konst = c.R[0]
}

// Load key material into the register
func addkey(c *shn_ctx, k uint32) {
	c.R[keyP] ^= k
}

// extra nonlinear diffusion of register for key and MAC
func shn_diffuse(c *shn_ctx) {
   for i := 0; i < 16; i++ {
      cycle(c)
   }
}

func shn_loadkey(c *shn_ctx, key []byte, keylen int) {
	var i int
	var j int
	var k uint32
	var xtra [4]uint8

	/* start folding in key */
	for i = 0; i < keylen&^0x3; i += 4 {
		k = byte2word(key[i:])
		addkey(c, k)
		cycle(c)
	}

	/* if there were any extra key bytes, zero pad to a word */
	if i < keylen {
		for j = 0; i < keylen; i++ { /* i unchanged */
			xtra[j] = uint8(key[i])
			j++ /* j unchanged */
		}
		for ; j < 4; j++ {
			xtra[j] = 0
		}
		k = byte2word(xtra[:])
		addkey(c, k)
		cycle(c)
	}

	addkey(c, uint32(keylen))

	cycle(c)

	/* save a copy of the register */
	for i = 0; i < num; i++ {
		c.crc[i] = c.R[i]
	}

	/* now diffuse */
	shn_diffuse(c)

	/* now xor the copy back -- makes key loading irreversible */
	for i = 0; i < num; i++ {
		c.R[i] ^= c.crc[i]
	}
}

// Published "key" interface
func shn_key(c *shn_ctx, key []byte, keylen int) {
	shn_initstate(c)
	shn_loadkey(c, key, keylen)
	shn_genkonst(c) /* in case we proceed to stream generation */
	shn_savestate(c)
	c.nbuf = 0
}

// Published "IV" interface
func shn_nonce(c *shn_ctx, nonce []byte, noncelen int) {
	shn_reloadstate(c)
	c.konst = initkonst
	shn_loadkey(c, nonce, noncelen)
	shn_genkonst(c)
	c.nbuf = 0
}

// Combined MAC and encryption. Note that plaintext is accumulated for MAC.
func shn_encrypt(c *shn_ctx, buf []byte, nbytes int) {
	var endbuf []byte
	var t uint32 = 0

	/* Handle any previously buffered bytes */
	if c.nbuf != 0 {
		for c.nbuf != 0 && nbytes != 0 {
			c.mbuf ^= uint32(buf[0]) << uint(32-c.nbuf)
			buf[0] ^= byte((c.sbuf >> uint(32-c.nbuf)) & 0xFF)
			buf = buf[1:]
			c.nbuf -= 8
			nbytes--
		}

		if c.nbuf != 0 { /* not a whole word yet */
			return
		}

		/* LFSR already cycled */
		macfunc(c, c.mbuf)
	}

	/* Handle whole words */
	endbuf = buf[uint32(nbytes)&^(uint32(0x03)):]

	for -cap(buf) < -cap(endbuf) {
		cycle(c)
		t = byte2word(buf)
		macfunc(c, t)
		t ^= c.sbuf
		word2byte(t, buf)
		buf = buf[4:]
	}

	/* Handle any trailing bytes */
	nbytes &= 0x03

	if nbytes != 0 {
		cycle(c)
		c.mbuf = 0
		c.nbuf = 32
		for c.nbuf != 0 && nbytes != 0 {
			c.mbuf ^= uint32(buf[0]) << uint(32-c.nbuf)
			buf[0] ^= byte((c.sbuf >> uint(32-c.nbuf)) & 0xFF)
			buf = buf[1:]
			c.nbuf -= 8
			nbytes--
		}
	}
}

// Combined MAC and decryption. Note that plaintext is accumulated for MAC.
func shn_decrypt(c *shn_ctx, buf []byte, nbytes int) {
	var endbuf []byte
	var t uint32 = 0

	/* Handle any previously buffered bytes */
	if c.nbuf != 0 {
		for c.nbuf != 0 && nbytes != 0 {
			buf[0] ^= byte((c.sbuf >> uint(32-c.nbuf)) & 0xFF)
			c.mbuf ^= uint32(buf[0]) << uint(32-c.nbuf)
			buf = buf[1:]
			c.nbuf -= 8
			nbytes--
		}

		if c.nbuf != 0 { /* not a whole word yet */
			return
		}

		/* LFSR already cycled */
		macfunc(c, c.mbuf)
	}

	/* Handle whole words */
	endbuf = buf[uint32(nbytes)&^(uint32(0x03)):]

	for -cap(buf) < -cap(endbuf) {
		cycle(c)
		t = byte2word(buf) ^ c.sbuf
		macfunc(c, t)
		word2byte(t, buf)
		buf = buf[4:]
	}

	/* Handle any trailing bytes */
	nbytes &= 0x03

	if nbytes != 0 {
		cycle(c)
		c.mbuf = 0
		c.nbuf = 32
		for c.nbuf != 0 && nbytes != 0 {
			buf[0] ^= byte((c.sbuf >> uint(32-c.nbuf)) & 0xFF)
			c.mbuf ^= uint32(buf[0]) << uint(32-c.nbuf)
			buf = buf[1:]
			c.nbuf -= 8
			nbytes--
		}
	}
}

// Having accumulated a MAC, finish processing and return it. Note that any
// unprocessed bytes are treated as if they were encrypted zero bytes, so
// plaintext (zero) is accumulated.
func shn_finish(c *shn_ctx, buf []byte, nbytes int) {
   var i int
   /* Handle any previously buffered bytes */
   if c.nbuf != 0 {
      /* LFSR already cycled */
      macfunc(c, c.mbuf)
   }
   cycle(c)
   addkey(c, initkonst^(uint32(c.nbuf)<<3))
   c.nbuf = 0
   for i = 0; i < num; i++ {
      c.R[i] ^= c.crc[i]
   }
   shn_diffuse(c)
   /* produce output from the stream buffer */
   for nbytes > 0 {
      cycle(c)
      if nbytes >= 4 {
         word2byte(c.sbuf, buf)
         nbytes -= 4
         buf = buf[4:]
      } else {
         for i = 0; i < nbytes; i++ {
            buf[i] = byte(toByte(c.sbuf, i))
         }
         break
      }
   }
}

func createStream(keys sharedKeys, conn plainConnection) packetStream {
   s := &shannonStream{
      mutex:  &sync.Mutex{},
      reader: conn.Reader,
      writer: conn.Writer,
   }
   setKey(&s.recvCipher, keys.recvKey)
   setKey(&s.sendCipher, keys.sendKey)
   return s
}

