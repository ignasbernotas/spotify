package spotify

import (
   "encoding/binary"
   "encoding/hex"
   "fmt"
   "github.com/segmentio/encoding/proto"
   "io"
   "math/big"
   "os"
)

func encodeRequest(seq []byte, req request) ([]byte, error) {
   buf, err := encodeMercuryHead(
      seq, 1+uint16(len(req.payload)), 1,
   )
   if err != nil {
      return nil, err
   }
   header := Header{
      ContentType: req.contentType,
      Method: req.method,
      URI: req.uri,
   }
   hData, err := proto.Marshal(header)
   if err != nil {
      return nil, err
   }
   // must use uint16
   err = binary.Write(buf, binary.BigEndian, uint16(len(hData)))
   if err != nil {
      return nil, err
   }
   _, err = buf.Write(hData)
   if err != nil {
      return nil, err
   }
   for _, p := range req.payload {
      err = binary.Write(buf, binary.BigEndian, len(p))
      if err != nil {
         return nil, err
      }
      _, err = buf.Write(p)
      if err != nil {
         return nil, err
      }
   }
   return buf.Bytes(), nil
}

type Header struct {
   URI string `protobuf:"bytes,1"`
   ContentType string `protobuf:"bytes,2"`
   Method string `protobuf:"bytes,3"`
   StatusCode int32 `protobuf:"zigzag32,4"`
   UserFields []struct {
      Key string `protobuf:"bytes,1"`
      Value string `protobuf:"bytes,2"`
   } `protobuf:"bytes,6"`
}

func (m *internal) parseResponse(cmd uint8, reader io.Reader) (*response, error) {
   seq, flags, count, err := handleHead(reader)
   if err != nil {
      return nil, err
   }
   seqKey := string(seq)
   pend := m.Pending[seqKey]
   for i := uint16(0); i < count; i++ {
      part, err := parsePart(reader)
      if err != nil {
         return nil, err
      }
      if pend.partial != nil {
         part = append(pend.partial, part...)
         pend.partial = nil
      }
      if i == count-1 && (flags == 2) {
         pend.partial = part
      } else {
         pend.parts = append(pend.parts, part)
      }
   }
   if flags == 1 {
      delete(m.Pending, seqKey)
      hData := pend.parts[0]
      var head Header
      err := proto.Unmarshal(hData, &head)
      if err != nil {
         return nil, err
      }
      return &response{
         headerData: hData,
         payload: pend.parts[1:],
         seqKey: seqKey,
         statusCode: head.StatusCode,
         uri: head.URI,
      }, nil
   } else {
      m.Pending[seqKey] = pend
   }
   return nil, nil
}

type AudioFile_Format int32

const (
   AudioFile_AAC_160        AudioFile_Format = 10
   AudioFile_AAC_320        AudioFile_Format = 11
   AudioFile_MP3_160        AudioFile_Format = 5
   AudioFile_MP3_160_ENC    AudioFile_Format = 7
   AudioFile_MP3_256        AudioFile_Format = 3
   AudioFile_MP3_320        AudioFile_Format = 4
   AudioFile_MP3_96         AudioFile_Format = 6
   AudioFile_OGG_VORBIS_160 AudioFile_Format = 1
   AudioFile_OGG_VORBIS_320 AudioFile_Format = 2
   AudioFile_OGG_VORBIS_96  AudioFile_Format = 0
   AudioFile_OTHER2         AudioFile_Format = 8
   AudioFile_OTHER3         AudioFile_Format = 9
   AudioFile_OTHER4         AudioFile_Format = 12
   AudioFile_OTHER5         AudioFile_Format = 13
)

type AudioFile struct {
   FileId []byte            `protobuf:"bytes,1"`
   Format AudioFile_Format `protobuf:"varint,2"`
}

type Track struct {
   GID []byte         `protobuf:"bytes,1"`
   Name             *string        `protobuf:"bytes,2"`
   Number           *int32         `protobuf:"zigzag32,5"`
   DiscNumber       *int32         `protobuf:"zigzag32,6"`
   Duration         *int32         `protobuf:"zigzag32,7"`
   Popularity       *int32         `protobuf:"zigzag32,8"`
   Explicit         *bool          `protobuf:"varint,9"`
   File             []*AudioFile   `protobuf:"bytes,12"`
}

func getFormat(track Track) (*AudioFile, error) {
   for _, file := range track.File {
      if file.Format == audioFile_OGG_VORBIS_160 {
         return file, nil
      }
   }
   msg := "could not find any files of the song in the specified formats"
   return nil, fmt.Errorf(msg)
}

func (ses *Session) DownloadTrackID(id string) error {
   b62 := new(big.Int)
   b62.SetString(id, 62)
   id = hex.EncodeToString(b62.Bytes())
   addr := "hm://metadata/4/track/" + id
   fmt.Println("GET", addr)
   data := ses.mercury.mercuryGet(addr)
   var trk Track
   err := proto.Unmarshal(data, &trk)
   if err != nil {
      return err
   }
   fSelect, err := getFormat(trk)
   if err != nil {
      return err
   }
   aFile := newAudioFileWithIdAndFormat(
      fSelect.FileId,
      audioFile_OGG_VORBIS_160,
      ses.player,
   )
   // Start loading the audio key
   if err := aFile.loadKey(trk.GID); err != nil {
      return err
   }
   // Then start loading the audio itself
   aFile.loadChunks()
   file, err := os.Create("file.ogg")
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(aFile); err != nil {
      return err
   }
   return nil
}
