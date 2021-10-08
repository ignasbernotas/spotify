package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

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
