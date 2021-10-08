package discovery

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ConvertTo62(raw []byte) string {
	bi := big.Int{}
	bi.SetBytes(raw)
	rem := big.NewInt(0)
	base := big.NewInt(62)
	zero := big.NewInt(0)
	result := ""

	for bi.Cmp(zero) > 0 {
		_, rem = bi.DivMod(&bi, base, rem)
		result += string(alphabet[int(rem.Uint64())])
	}

	for len(result) < 22 {
		result += "0"
	}
	return reverse(result)
}

func Base62ToHex(b62 string) string {
	return fmt.Sprintf("%x", Convert62(b62))
}


var EMPTY_HTTP_RESPONSE = &http.Response{}

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


func GenerateDeviceId(name string) string {
	hash := sha1.Sum([]byte(name))
	hash64 := base64.StdEncoding.EncodeToString(hash[:])
	return hash64
}


const kAPEndpoint = "https://apresolve.spotify.com"

// APList is the JSON structure corresponding to the output of the AP endpoint resolve API
type APList struct {
	ApListNoType []string `json:"ap_list"`
	ApList       []string `json:"accesspoint"`
}

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
