package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/89z/spotify/Spotify"
	"github.com/89z/spotify/core"
	"github.com/89z/spotify/metadata"
	"github.com/89z/spotify/utils"
	"io"
	"os"
	"strings"
	"time"
)

func GetTrackFileAndInfo(session *core.Session, trackID string) (*SpotifyTrack, error) {
	// Get the track metadata: it holds information about which files and encodings are available
	track, err := session.Mercury().GetTrack(utils.Base62ToHex(trackID))
	if err != nil {
		return nil, fmt.Errorf("Failed to get track metadata: %s", err)
	}

	var selectedFile *Spotify.AudioFile = nil
	for _, file := range track.GetFile() {
		if file.GetFormat() == Spotify.AudioFile_OGG_VORBIS_160 {
			selectedFile = file
		}
	}
	if selectedFile == nil {
		return nil, fmt.Errorf("Could not find any files of the song in the specified formats")
	}

	// Synchronously load the track
	audioFile, err := session.Player().LoadTrack(selectedFile, track.GetGid())
	if err != nil {
		return nil, fmt.Errorf("Failed to download the track: %s", err)
	}

	return GetTrackInfo(audioFile, track), nil
}

func GetTrackInfo(audioFile io.Reader, track *Spotify.Track) *SpotifyTrack {
	serializedTrack := &SpotifyTrack{}
	serializedTrack.AudioFile = audioFile
	serializedTrack.TrackName = track.GetName()
	serializedTrack.TrackNumber = track.GetNumber()
	serializedTrack.TrackDuration = (track.GetDuration() / 1000) // convert ms to seconds
	serializedTrack.TrackDiscNumber = track.GetDiscNumber()

	album := track.GetAlbum()
	if album != nil {
		serializedTrack.Album.Name = album.GetName()
		serializedTrack.Album.Label = album.GetLabel()
		serializedTrack.Album.Genre = album.GetGenre()
		albumDate := album.GetDate()
		if albumDate != nil {
			serializedTrack.Album.Date = time.Date(int(albumDate.GetYear()), time.Month(int(albumDate.GetMonth())), int(albumDate.GetDay()), 0, 0, 0, 0, time.UTC)
		}

		albumArtists := album.GetArtist()
		if albumArtists != nil {
			for _, artist := range albumArtists {
				serializedTrack.Album.ArtistNames = append(serializedTrack.Album.ArtistNames, artist.GetName())
			}
		}
	}

	trackArtists := track.GetArtist()
	if trackArtists != nil {
		for _, artist := range trackArtists {
			serializedTrack.TrackArtistNames = append(serializedTrack.TrackArtistNames, artist.GetName())
		}
	}

	return serializedTrack
}


/* use these structs because they are much easier to work with than protobuf structs */
type SpotifyAlbum struct {
	Name        string
	Label       string
	Genre       []string
	Date        time.Time
	ArtistNames []string
}

type SpotifyTrack struct {
	AudioFile        io.Reader
	TrackName        string
	TrackNumber      int32
	TrackDuration    int32
	TrackDiscNumber  int32
	TrackArtistNames []string
	Album            SpotifyAlbum
}

const delimeter string = "; "


func saveReaderToNewFile(reader io.Reader, fileName string) error {
	newFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file with path %s: %s", fileName, err)
	}
	io.Copy(newFile, reader) // copy the reader to the writer

	newFile.Close() // don't defer since there's nothing in between and defer has a performance cost
	return nil
}

func getLastSplit(str string, delimiter string) string {
	str_split := strings.Split(str, delimiter)
	return str_split[len(str_split)-1]
}

func NiceJsonFormat(object interface{}) string {
	jsonBytes, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		return ""
	} else {
		return string(jsonBytes)
	}
}


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
