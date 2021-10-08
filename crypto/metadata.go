package crypto

import (
   "encoding/json"
   "fmt"
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

type TopRecommendation struct {
}

type SearchResponse struct {
	Results         SearchResult `json:"results"`
	RequestId       string       `json:"requestId"`
	CategoriesOrder []string     `json:"categoriesOrder"`
}

type SearchResult struct {
	Tracks struct {
		Hits  []Track `json:"hits"`
		Total int     `json:"total"`
	} `json:"tracks"`

	Albums struct {
		Hits  []Album `json:"hits"`
		Total int     `json:"total"`
	} `json:"albums"`

	Artists struct {
		Hits  []Artist `json:"hits"`
		Total int      `json:"total"`
	} `json:"artists"`

	Playlists struct {
		Hits  []Playlist `json:"hits"`
		Total int        `json:"total"`
	} `json:"playlists"`

	Profiles struct {
		Hits  []Profile `json:"hits"`
		Total int       `json:"total"`
	} `json:"profiles"`

	Genres struct {
		Hits  []Genre `json:"hits"`
		Total int     `json:"total"`
	} `json:"genres"`

	TopHit struct {
		Hits  []TopHit `json:"hits"`
		Total int      `json:"total"`
	} `json:"topHit"`

	Shows struct {
		Hits  []Show `json:"hits"`
		Total int    `json:"total"`
	} `json:"shows"`

	VideoEpisodes struct {
		Hits  []VideoEpisode `json:"hits"`
		Total int            `json:"total"`
	} `json:"videoEpisodes"`

	TopRecommendations struct {
		Hits  []TopRecommendation `json:"hits"`
		Total int                 `json:"total"`
	} `json:"topRecommendations"`

	Error error
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

func (m *Client) GetTrack(id string) (*pb.Track, error) {
	uri := "hm://metadata/4/track/" + id
	result := &pb.Track{}
	err := m.mercuryGetProto(uri, result)
	return result, err
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
