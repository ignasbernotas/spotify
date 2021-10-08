package crypto

import (
   "encoding/json"
   "github.com/golang/protobuf/proto"
   "github.com/89z/spotify/pb"
)

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

type Playlist struct {
	Name           string `json:"name"`
	Uri            string `json:"uri"`
	Image          string `json:"image"`
	FollowersCount int    `json:"followersCount"`
	Author         string `json:"author"`
}

type Profile struct {
	Name           string `json:"name"`
	Uri            string `json:"uri"`
	Image          string `json:"image"`
	FollowersCount int    `json:"followersCount"`
}

type Genre struct {
	Name  string `json:"name"`
	Uri   string `json:"uri"`
	Image string `json:"image"`
}

type TopHit struct {
	Uri            string `json:"uri"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	Verified       bool   `json:"verified"`
	Following      bool   `json:"following"`
	FollowersCount int    `json:"followersCount"`
	Author         string `json:"author"`
	Log            struct {
		Origin string `json:"origin"`
		TopHit string `json:"top_hit"`
	} `json:"log"`
	Artists []Artist `json:"artists"`
	Album   Album    `json:"album"`
}

type Show struct {
	Name     string `json:"name"`
	Uri      string `json:"uri"`
	Image    string `json:"image"`
	ShowType string `json:"showType"`
}

type VideoEpisode struct {
	Name  string `json:"name"`
	Uri   string `json:"uri"`
	Image string `json:"image"`
}

type SuggestResult struct {
	Sections []struct {
		RawItems json.RawMessage `json:"items"`
		Typ      string          `json:"type"`
	} `json:"sections"`
	Albums  []Artist
	Artists []Album
	Tracks  []Track
	TopHits []TopHit
	Error   error
}

type Token struct {
	AccessToken string   `json:"accessToken"`
	ExpiresIn   int      `json:"expiresIn"`
	TokenType   string   `json:"tokenType"`
	Scope       []string `json:"scope"`
}


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

func (m *Client) mercuryGetJson(url string, result interface{}) error {
   data := m.mercuryGet(url)
   return json.Unmarshal(data, result)
}

func (m *Client) mercuryGetProto(url string, result proto.Message) error {
   data := m.mercuryGet(url)
   return proto.Unmarshal(data, result)
}

func (m *Client) GetTrack(id string) (*pb.Track, error) {
   uri := "hm://metadata/4/track/" + id
   result := &pb.Track{}
   err := m.mercuryGetProto(uri, result)
   return result, err
}
