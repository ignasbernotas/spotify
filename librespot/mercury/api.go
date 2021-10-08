package mercury

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/89z/spotify/Spotify"
	"github.com/89z/spotify/librespot/metadata"
	"net/url"
)

func (m *Client) mercuryGet(url string) []byte {
	done := make(chan []byte)
	go m.Request(Request{
		Method:  "GET",
		Uri:     url,
		Payload: [][]byte{},
	}, func(res Response) {
		done <- res.CombinePayload()
	})

	result := <-done
	return result
}

func (m *Client) mercuryGetJson(url string, result interface{}) (err error) {
	data := m.mercuryGet(url)
	// fmt.Printf("Received:\n%s\n\n\n", data)
	err = json.Unmarshal(data, result)
	return
}

func (m *Client) mercuryGetProto(url string, result proto.Message) (err error) {
	data := m.mercuryGet(url)
	// ioutil.WriteFile("/tmp/proto.blob", data, 0644)
	err = proto.Unmarshal(data, result)
	return
}

func (m *Client) GetToken(clientId string, scopes string) (*metadata.Token, error) {
	uri := fmt.Sprintf("hm://keymaster/token/authenticated?client_id=%s&scope=%s", url.QueryEscape(clientId),
		url.QueryEscape(scopes))

	token := &metadata.Token{}
	err := m.mercuryGetJson(uri, token)
	return token, err
}

func (m *Client) Search(search string, limit int, country string, username string) (*metadata.SearchResponse, error) {
	v := url.Values{}
	v.Set("entityVersion", "2")
	v.Set("limit", fmt.Sprintf("%d", limit))
	v.Set("imageSize", "large")
	v.Set("catalogue", "")
	v.Set("country", country)
	v.Set("platform", "zelda")
	v.Set("username", username)

	uri := fmt.Sprintf("hm://searchview/km/v4/search/%s?%s", url.QueryEscape(search), v.Encode())

	result := &metadata.SearchResponse{}
	err := m.mercuryGetJson(uri, result)
	return result, err
}

func (m *Client) Suggest(search string) (*metadata.SuggestResult, error) {
	uri := "hm://searchview/km/v3/suggest/" + url.QueryEscape(search) + "?limit=3&intent=2516516747764520149&sequence=0&catalogue=&country=&locale=&platform=zelda&username="
	data := m.mercuryGet(uri)

	return parseSuggest(data)
}

func (m *Client) GetTrack(id string) (*Spotify.Track, error) {
	uri := "hm://metadata/4/track/" + id
	result := &Spotify.Track{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetArtist(id string) (*Spotify.Artist, error) {
	uri := "hm://metadata/4/artist/" + id
	result := &Spotify.Artist{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetAlbum(id string) (*Spotify.Album, error) {
	uri := "hm://metadata/4/album/" + id
	result := &Spotify.Album{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

// new functions
func (m *Client) GetArtistInfo(id string, username string) (*metadata.DetailPageArtist, error) {
	uri := fmt.Sprintf("hm://artist/v1/%s/desktop?format=json&catalogue=free&locale=en&username=%scat=1", id, username)

	result := &metadata.DetailPageArtist{}
	err := m.mercuryGetJson(uri, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (m *Client) GetAlbumInfo(id string, username string) (*metadata.DetailPageGenericRelease, error) {
	uri := fmt.Sprintf("hm://album/v1/album-app/album/spotify:album:%s/desktop?catalogue=free&locale=en&username=%s", id, username)

	result := &metadata.DetailPageGenericRelease{}
	err := m.mercuryGetJson(uri, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseSuggest(body []byte) (*metadata.SuggestResult, error) {
	result := &metadata.SuggestResult{}
	err := json.Unmarshal(body, result)
	if err != nil {
		fmt.Println("err", err)
	}

	for _, s := range result.Sections {
		switch s.Typ {
		case "top-results":
			err = json.Unmarshal(s.RawItems, &result.TopHits)
		case "album-results":
			err = json.Unmarshal(s.RawItems, &result.Albums)
		case "artist-results":
			err = json.Unmarshal(s.RawItems, &result.Artists)
		case "track-results":
			err = json.Unmarshal(s.RawItems, &result.Tracks)
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
