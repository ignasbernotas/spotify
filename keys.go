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
   "time"
   cryptoRand "crypto/rand"
)

const (
   alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
   kAPEndpoint = "https://apresolve.spotify.com"
)

var empty_HTTP_RESPONSE = &http.Response{}

var standard_APRESOLVE_HEADERS = map[string]string{
   "User-Agent":                        "Spotify/111000546 (8; 0; 5)",
   "x-spotify-ap-resolve-pod-override": "0",
}

func apResolve() (string, error) {
   var unixTimestamp = strconv.Itoa(int(time.Now().Unix()))
   r, err := httpGetHeaders(
      kAPEndpoint+"/?time="+unixTimestamp+"&type=accesspoint",
      standard_APRESOLVE_HEADERS,
   )
   if err != nil {
   return "", err
   }
   defer r.Body.Close()
   var endpoints apList
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

func convert62(id string) []byte {
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

func generateDeviceId(name string) string {
   hash := sha1.Sum([]byte(name))
   return base64.StdEncoding.EncodeToString(hash[:])
}

func httpGetHeaders(link string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return empty_HTTP_RESPONSE, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return empty_HTTP_RESPONSE, err
	}

	return response, nil
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

func randomVec(count int) []byte {
	c := count
	b := make([]byte, c)
	_, err := cryptoRand.Read(b)
	if err != nil {
		log.Fatal("error:", err)
	}
	return b
}

type album struct {
	Artists []artist `json:"artists"`
	Image   string   `json:"image"`
	Name    string   `json:"name"`
	Uri     string   `json:"uri"`
}

type apList struct {
	ApListNoType []string `json:"ap_list"`
	ApList       []string `json:"accesspoint"`
}

type artist struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	Uri   string `json:"uri"`
}

type blobInfo struct {
   Username    string
   DecodedBlob string
}

// Discovery stores the information about Spotify Connect Discovery Request
type discovery struct {
   loginBlob blobInfo
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

func (p *privateKeys) pubKey() []byte {
	return p.publicKey.Bytes()
}

type sharedKeys struct {
	challenge []byte
	sendKey   []byte
	recvKey   []byte
}

type track struct {
	Album      album    `json:"album"`
	Artists    []artist `json:"artists"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Uri        string   `json:"uri"`
	Duration   int      `json:"duration"`
	Popularity float32  `json:"popularity"`
}
