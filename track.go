package spotify

import (
   "fmt"
   "github.com/89z/spotify/pb"
   "io"
   "os"
   "time"
)

// use these structs because they are much easier to work with than protobuf
// structs
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

func DownloadTrackID(ses *Session, id string) error {
   tra, err := ses.Mercury().GetTrack(Base62ToHex(id))
   if err != nil {
      return fmt.Errorf("failed to get track metadata %v", err)
   }
   var selectedFile *pb.AudioFile = nil
   for _, file := range tra.GetFile() {
      if file.GetFormat() == pb.AudioFile_OGG_VORBIS_160 {
         selectedFile = file
      }
   }
   if selectedFile == nil {
      msg := "could not find any files of the song in the specified formats"
      return fmt.Errorf(msg)
   }
   audioFile, err := ses.Player().LoadTrack(selectedFile, tra.GetGid())
   if err != nil {
      return fmt.Errorf("failed to download the track %v", err)
   }
   track := GetTrackInfo(audioFile, tra)
   fmt.Printf("%+v\n", track)
   file, err := os.Create(track.TrackName + ".ogg")
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(track.AudioFile); err != nil {
      return err
   }
   return nil
}

func GetTrackInfo(audioFile io.Reader, track *pb.Track) *SpotifyTrack {
   serializedTrack := &SpotifyTrack{}
   serializedTrack.AudioFile = audioFile
   serializedTrack.TrackName = track.GetName()
   serializedTrack.TrackNumber = track.GetNumber()
   // convert ms to seconds
   serializedTrack.TrackDuration = (track.GetDuration() / 1000)
   serializedTrack.TrackDiscNumber = track.GetDiscNumber()
   album := track.GetAlbum()
   if album != nil {
      serializedTrack.Album.Name = album.GetName()
      serializedTrack.Album.Label = album.GetLabel()
      serializedTrack.Album.Genre = album.GetGenre()
      albumDate := album.GetDate()
      if albumDate != nil {
         serializedTrack.Album.Date = time.Date(
            int(albumDate.GetYear()),
            time.Month(int(albumDate.GetMonth())),
            int(albumDate.GetDay()), 0, 0, 0, 0, time.UTC,
         )
      }
      albumArtists := album.GetArtist()
      for _, artist := range albumArtists {
         serializedTrack.Album.ArtistNames = append(
            serializedTrack.Album.ArtistNames, artist.GetName(),
         )
      }
   }
   trackArtists := track.GetArtist()
   for _, artist := range trackArtists {
      serializedTrack.TrackArtistNames = append(
         serializedTrack.TrackArtistNames, artist.GetName(),
      )
   }
   return serializedTrack
}
