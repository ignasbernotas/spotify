package spotify

import (
   "fmt"
   "github.com/89z/spotify/pb"
   "math/big"
   "os"
   "strings"
)

func convert62(id string) []byte {
   base := big.NewInt(62)
   n := &big.Int{}
   for _, c := range []byte(id) {
      d := big.NewInt(int64(strings.IndexByte(alphabet, c)))
      n = n.Mul(n, base)
      n = n.Add(n, d)
   }
   nBytes := n.Bytes()
   if len(nBytes) < 16 {
      paddingBytes := make([]byte, 16-len(nBytes))
      nBytes = append(paddingBytes, nBytes...)
   }
   return nBytes
}


func (ses *session) DownloadTrackID(id string) error {
   hex := fmt.Sprintf("%x", convert62(id))
   tra, err := ses.mercury.getTrack(hex)
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
   audioFile, err := ses.player.loadTrack(selectedFile, tra.GetGid())
   if err != nil {
      return fmt.Errorf("failed to download the track %v", err)
   }
   track := getTrackInfo(tra)
   fmt.Printf("%+v\n", track)
   file, err := os.Create(track.TrackName + ".ogg")
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(audioFile); err != nil {
      return err
   }
   return nil
}
