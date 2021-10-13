package main

import (
   "flag"
   "github.com/89z/spotify"
)

func main() {
   var track, user, pass string
   flag.StringVar(&pass, "p", "", "password")
   flag.StringVar(&user, "u", "", "username")
   flag.StringVar(&track, "t", "", "track ID to download")
   flag.Parse()
   ses, err := spotify.Login(user, pass, "librespot")
   if err != nil {
      panic(err)
   }
   if err := ses.DownloadTrackID(track); err != nil {
      panic(err)
   }
}
