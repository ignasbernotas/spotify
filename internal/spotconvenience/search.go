package spotconvenience

import (
	"fmt"
	"github.com/librespot-org/librespot-golang/librespot/core"
	"github.com/librespot-org/librespot-golang/librespot/metadata"
	"strings"
)

func removeSpotifyUriPrefix(uri string) string {
	return getLastSplit(uri, ":")
}

func Search(session *core.Session, query string) (*metadata.SearchResult, error) {
	response, err := session.Mercury().Search(query, 12, session.Country(), session.Username())

	if err != nil {
		return nil, fmt.Errorf("Failed to search:", err)
	}

	results := response.Results
	if results.Error != nil {
		return nil, fmt.Errorf("Search result error:", results.Error)
	}

	for _, result := range results.Artists.Hits {
		fmt.Printf("Artist: %s (%s)\n", result.Name, removeSpotifyUriPrefix(result.Uri))
	}
	fmt.Printf("\n")
	for _, result := range results.Albums.Hits {
		artistList := []string{}
		for _, artist := range result.Artists {
			artistList = append(artistList, artist.Name)
		}

		fmt.Printf("Album: %s - %s (%s)\n", strings.Join(artistList, ", "), result.Name, removeSpotifyUriPrefix(result.Uri))
	}
	fmt.Printf("\n")
	for _, result := range results.Tracks.Hits {
		artistList := []string{}
		for _, artist := range result.Artists {
			artistList = append(artistList, artist.Name)
		}

		fmt.Printf("Track: %s - %s (%s)\n", strings.Join(artistList, ", "), result.Name, removeSpotifyUriPrefix(result.Uri))
	}

	return &results, nil
}
