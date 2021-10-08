package main

import (
	"flag"
	// "fmt"
	"fmt"
	"github.com/chris124567/spotrip/internal/spotconvenience"
	"github.com/librespot-org/librespot-golang/librespot"
	"github.com/librespot-org/librespot-golang/librespot/core"
	"log"
	"strings"
)

const EMPTY_STIRNG string = ""

func isStringDefined(str string) bool {
	return str != EMPTY_STIRNG
}

func stringListCleanup(strList []string) []string {
	length := len(strList)
	for i := 0; i < length; i++ {
		strList[i] = strings.TrimSpace(strList[i])
	}
	return strList
}

func main() {
	var session *core.Session
	var err error

	username := flag.String("username", EMPTY_STIRNG, "Username of the spotify account. Required.")
	password := flag.String("password", EMPTY_STIRNG, "Password of the spotify account. Required.")
	searchQuery := flag.String("search", EMPTY_STIRNG, "Search query.")
	artists := flag.String("artists", EMPTY_STIRNG, "List of artist IDs to download (downloads all of their albums, singles, and compilations), separated by commas.")
	albums := flag.String("albums", EMPTY_STIRNG, "List of album IDs to download, separated by commas.")
	tracks := flag.String("tracks", EMPTY_STIRNG, "List of track IDs to download, separated by commas.")
	artistInfo := flag.String("artist_info", EMPTY_STIRNG, "Specify an artist ID in this field to get JSON formatted information about that artist.")

	flag.Parse()
	if !isStringDefined(*username) || !isStringDefined(*password) {
		log.Fatal("Please specify a username and password")
	}

	session, err = librespot.Login(*username, *password, "librespot")
	if err != nil {
		log.Fatalf("Failed to login: %+v", err)
	}

	if isStringDefined(*searchQuery) {
		spotconvenience.Search(session, *searchQuery)
	}

	if isStringDefined(*artists) {
		for _, artistId := range stringListCleanup(strings.Split(*artists, ",")) {
			uris, err := spotconvenience.GetArtistTracks(session, artistId)
			if err != nil {
				log.Fatalf("Failed to get track list for artist %s: %+v", artistId, err)
			}
			err = spotconvenience.DownloadTrackList(session, *uris)
			if err != nil {
				log.Fatalf("Failed to download tracks for artist %s: %+v", artistId, err)
			}
		}
	}
	if isStringDefined(*albums) {
		for _, albumId := range stringListCleanup(strings.Split(*albums, ",")) {
			uris, err := spotconvenience.GetAlbumTracks(session, albumId)
			if err != nil {
				log.Fatalf("Failed to get track list for album %s: %+v", albumId, err)
			}
			err = spotconvenience.DownloadTrackList(session, *uris)
			if err != nil {
				log.Fatalf("Failed to download tracks for album %s: %+v", albumId, err)
			}
		}
	}
	if isStringDefined(*tracks) {
		err = spotconvenience.DownloadTrackList(session, stringListCleanup(strings.Split(*tracks, ",")))
		if err != nil {
			log.Fatalf("Failed to download tracks: %+v", err)
		}
	}
	if isStringDefined(*artistInfo) {
		response, err := session.Mercury().GetArtistInfo(*artistInfo, session.Username())
		if err != nil {
			log.Fatalf("Failed to get artist info for artist %s: %+v", *artistInfo, err)
		}
		fmt.Print(spotconvenience.NiceJsonFormat(response))
	}
}
