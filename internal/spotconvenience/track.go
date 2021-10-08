package spotconvenience

import (
	"fmt"
	"github.com/librespot-org/librespot-golang/Spotify"
	"github.com/librespot-org/librespot-golang/librespot/core"
	"github.com/librespot-org/librespot-golang/librespot/utils"
	"io"
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
