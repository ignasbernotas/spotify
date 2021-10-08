package spotify

import (
   "fmt"
   "os"
   "path"
   "strconv"
   "strings"
)

const baseOutputDirectory string = "output"

func createTrackDirectory(track *SpotifyTrack) (string, error) {
	var mainArtistName string
	var albumName string

	albumName = track.Album.Name

	if len(track.TrackArtistNames) == 0 {
		mainArtistName = "Unknown"
	} else {
		mainArtistName = track.TrackArtistNames[0]
	}

	// if theres another disc, make a folder for it
	if track.TrackDiscNumber > 1 {
		albumName = path.Join(albumName, "Disc "+strconv.Itoa(int(track.TrackDiscNumber)))
	}

	newPath := path.Join(baseOutputDirectory, mainArtistName, mainArtistName+" - "+albumName)
	err := os.MkdirAll(newPath, os.ModePerm)

	return newPath, err
}

func trackOutputFilename(track *SpotifyTrack, outputDirectory string) string {
	return path.Join(outputDirectory, strconv.Itoa(int(track.TrackNumber))+" - "+track.TrackName+".ogg")
}

// NEED THIS
func downloadTrackId(session *Session, id string) error {
	track, err := GetTrackFileAndInfo(session, id)
	if err != nil {
		return err
	}

	outputDirectory, err := createTrackDirectory(track)
	if err != nil {
		return err
	}
	outputPath := trackOutputFilename(track, outputDirectory)

	fmt.Printf("Downloading: %s - %s (album track #%d) [%s] to %s\n", strings.Join(track.TrackArtistNames, ", "), track.TrackName, track.TrackNumber, id, outputPath)

	err = saveReaderToNewFile(track.AudioFile, outputPath)
	if err != nil {
		return err
	}
	return nil
}

// NEED THIS
func DownloadTrackList(session *Session, idList []string) error {
   for _, id := range idList {
      err := downloadTrackId(session, id)
      if err != nil {
         return fmt.Errorf("failed to download track %q %+v", id, err)
      }
   }
   return nil
}
func GetArtistTracks(session *Session, id string) (*[]string, error) {
	var albumUris []string
	var trackIds []string

	response, err := session.Mercury().GetArtistInfo(id, session.Username())
	if err != nil {
		return nil, err
	}

	// get all albums, singles, and compilations of an artist (spotify treats all of these as "albums")
	for _, album := range response.Releases.Albums.Releases {
		albumUris = append(albumUris, album.Uri)
	}
	for _, compilation := range response.Releases.Compilations.Releases {
		albumUris = append(albumUris, compilation.Uri)
	}
	for _, single := range response.Releases.Singles.Releases {
		albumUris = append(albumUris, single.Uri)
	}

	// get all track ids
	for _, albumId := range albumUris {
		albumTracks, err := GetAlbumTracks(session, removeSpotifyUriPrefix(albumId))
		if err != nil {
			break
		}
		trackIds = append(trackIds, *albumTracks...)
	}
	return &trackIds, err
}

func GetAlbumTracks(session *Session, id string) (*[]string, error) {
	var trackUris []string

	response, err := session.Mercury().GetAlbumInfo(id, session.Username())
	if err != nil {
		return nil, err
	}

	for _, disc := range response.Discs {
		for _, track := range disc.Tracks {
			trackUris = append(trackUris, removeSpotifyUriPrefix(track.Uri))
		}
	}

	return &trackUris, nil
}

func Login(username string, password string, deviceName string) (*Session, error) {
	return CoreLogin(username, password, deviceName)
}

func LoginSaved(username string, authData []byte, deviceName string) (*Session, error) {
	return CoreLoginSaved(username, authData, deviceName)
}

func LoginOAuth(deviceName string, clientId string, clientSecret string) (*Session, error) {
	return CoreLoginOAuth(deviceName, clientId, clientSecret)
}
