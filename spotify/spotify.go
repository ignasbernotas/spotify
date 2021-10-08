package main

import (
   "flag"
   "github.com/89z/spotify"
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
	var session *spotify.Session
	var err error

	username := flag.String("username", EMPTY_STIRNG, "Username of the spotify account. Required.")
	password := flag.String("password", EMPTY_STIRNG, "Password of the spotify account. Required.")
	tracks := flag.String("tracks", EMPTY_STIRNG, "List of track IDs to download, separated by commas.")

	flag.Parse()
	if !isStringDefined(*username) || !isStringDefined(*password) {
		log.Fatal("Please specify a username and password")
	}

	session, err = spotify.Login(*username, *password, "librespot")
	if err != nil {
		log.Fatalf("Failed to login: %+v", err)
	}
	if isStringDefined(*tracks) {
		err = spotify.DownloadTrackList(session, stringListCleanup(strings.Split(*tracks, ",")))
		if err != nil {
			log.Fatalf("Failed to download tracks: %+v", err)
		}
	}
}