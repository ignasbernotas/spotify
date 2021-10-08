package crypto

import (
   "encoding/json"
   "fmt"
   "github.com/golang/protobuf/proto"
   "github.com/89z/spotify/pb"
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

func (m *Client) mercuryGetProto(url string, result proto.Message) error {
   data := m.mercuryGet(url)
   return proto.Unmarshal(data, result)
}

func (m *Client) Suggest(search string) (*SuggestResult, error) {
	uri := "hm://searchview/km/v3/suggest/" + url.QueryEscape(search) + "?limit=3&intent=2516516747764520149&sequence=0&catalogue=&country=&locale=&platform=zelda&username="
	data := m.mercuryGet(uri)

	return parseSuggest(data)
}

func (m *Client) GetTrack(id string) (*pb.Track, error) {
	uri := "hm://metadata/4/track/" + id
	result := &pb.Track{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetArtist(id string) (*pb.Artist, error) {
	uri := "hm://metadata/4/artist/" + id
	result := &pb.Artist{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetAlbum(id string) (*pb.Album, error) {
	uri := "hm://metadata/4/album/" + id
	result := &pb.Album{}
	err := m.mercuryGetProto(uri, result)
	return result, err
}

func (m *Client) GetArtistInfo(id string, username string) (*DetailPageArtist, error) {
	uri := fmt.Sprintf("hm://artist/v1/%s/desktop?format=json&catalogue=free&locale=en&username=%scat=1", id, username)

	result := &DetailPageArtist{}
	err := m.mercuryGetJson(uri, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *Client) GetAlbumInfo(id string, username string) (*DetailPageGenericRelease, error) {
	uri := fmt.Sprintf("hm://album/v1/album-app/album/spotify:album:%s/desktop?catalogue=free&locale=en&username=%s", id, username)

	result := &DetailPageGenericRelease{}
	err := m.mercuryGetJson(uri, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseSuggest(body []byte) (*SuggestResult, error) {
	result := &SuggestResult{}
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
