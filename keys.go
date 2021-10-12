package spotify

import (
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "math/big"
   cryptoRand "crypto/rand"
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

type apList struct {
	ApListNoType []string `json:"ap_list"`
	ApList       []string `json:"accesspoint"`
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
