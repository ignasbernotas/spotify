package spotify

import (
   "crypto/hmac"
   "crypto/sha1"
   "encoding/base64"
   "encoding/json"
   "errors"
   "log"
   "math/big"
   "math/rand"
   "net/http"
   "strconv"
   "strings"
   "sync"
   "time"
   cryptoRand "crypto/rand"
)

func CreateStream(keys SharedKeys, conn plainConnection) packetStream {
	s := &ShannonStream{
		Reader: conn.Reader,
		Writer: conn.Writer,
		Mutex:  &sync.Mutex{},
	}
	SetKey(&s.RecvCipher, keys.recvKey)
	SetKey(&s.SendCipher, keys.sendKey)
	return s
}

type PrivateKeys struct {
	privateKey *big.Int
	publicKey  *big.Int

	generator   *big.Int
	prime       *big.Int
	clientNonce []byte
}

type SharedKeys struct {
	challenge []byte
	sendKey   []byte
	recvKey   []byte
}

func RandomVec(count int) []byte {
	c := count
	b := make([]byte, c)
	_, err := cryptoRand.Read(b)
	if err != nil {
		log.Fatal("error:", err)
	}
	return b
}

func Powm(base, exp, modulus *big.Int) *big.Int {
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


// NEED THIS
func (p *PrivateKeys) AddRemoteKey(remote []byte, clientPacket []byte, serverPacket []byte) SharedKeys {
	remote_be := new(big.Int)
	remote_be.SetBytes(remote)
	shared_key := Powm(remote_be, p.privateKey, p.prime)

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

	return SharedKeys{
		challenge: mac.Sum(nil),
		sendKey:   data[0x14:0x34],
		recvKey:   data[0x34:0x54],
	}
}

// NEED THIS
func (p *PrivateKeys) PubKey() []byte {
	return p.publicKey.Bytes()
}

// NEED THIS
func (p *PrivateKeys) ClientNonce() []byte {
	return p.clientNonce
}

// NEED THIS
func (s *SharedKeys) Challenge() []byte {
	return s.challenge
}

// BlobInfo is the structure holding authentication blob data. The blob is an
// encoded/encrypted byte array (encoded as base64), holding the encryption
// keys, the deviceId, and the username.
type BlobInfo struct {
   Username    string
   DecodedBlob string
}

// Discovery stores the information about Spotify Connect Discovery Request
type Discovery struct {
   loginBlob  BlobInfo
}

// NEED THIS
func (d *Discovery) LoginBlob() BlobInfo {
	return d.loginBlob
}

type Artist struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	Uri   string `json:"uri"`
}

type Album struct {
	Artists []Artist `json:"artists"`
	Image   string   `json:"image"`
	Name    string   `json:"name"`
	Uri     string   `json:"uri"`
}

type Track struct {
	Album      Album    `json:"album"`
	Artists    []Artist `json:"artists"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Uri        string   `json:"uri"`
	Duration   int      `json:"duration"`
	Popularity float32  `json:"popularity"`
}

const (
   alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
   kAPEndpoint = "https://apresolve.spotify.com"
)

var EMPTY_HTTP_RESPONSE = &http.Response{}

var STANDARD_APRESOLVE_HEADERS = map[string]string{
   "User-Agent":                        "Spotify/111000546 (8; 0; 5)",
   "x-spotify-ap-resolve-pod-override": "0",
}

// APResolve fetches the available Spotify servers (AP) and picks a random one
func APResolve() (string, error) {
	var endpoints APList

	var unixTimestamp = strconv.Itoa(int(time.Now().Unix()))
	r, err := HttpGetHeaders(kAPEndpoint+"/?time="+unixTimestamp+"&type=accesspoint", STANDARD_APRESOLVE_HEADERS)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&endpoints)
	if err != nil {
		return "", err
	}
	if len(endpoints.ApList) > 0 {
		return endpoints.ApList[rand.Intn(len(endpoints.ApList))], nil
	} else if len(endpoints.ApListNoType) > 0 {
		return endpoints.ApListNoType[rand.Intn(len(endpoints.ApListNoType))], nil
	} else {
		return "", errors.New("AP endpoint list is empty")
	}

}

func Convert62(id string) []byte {
	base := big.NewInt(62)

	n := &big.Int{}
	for _, c := range []byte(id) {
		d := big.NewInt(int64(strings.IndexByte(alphabet, c)))
		n = n.Mul(n, base)
		n = n.Add(n, d)
	}

	nBytes := n.Bytes()
	if len(nBytes) < 16 {
		paddingBytes := make([]byte, 16-len(nBytes))
		nBytes = append(paddingBytes, nBytes...)
	}
	return nBytes
}

func GenerateDeviceId(name string) string {
   hash := sha1.Sum([]byte(name))
   return base64.StdEncoding.EncodeToString(hash[:])
}

func HttpGetHeaders(link string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return EMPTY_HTTP_RESPONSE, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return EMPTY_HTTP_RESPONSE, err
	}

	return response, nil
}

// APList is the JSON structure corresponding to the output of the AP endpoint resolve API
type APList struct {
	ApListNoType []string `json:"ap_list"`
	ApList       []string `json:"accesspoint"`
}
