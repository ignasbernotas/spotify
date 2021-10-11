package spotify

import (
   "crypto/cipher"
   "github.com/89z/spotify/pb"
   "sync"
)

func (m *client) getTrack(id string) (*pb.Track, error) {
   result := new(pb.Track)
   err := m.mercuryGetProto("hm://metadata/4/track/" + id, result)
   if err != nil {
      return nil, err
   }
   return result, nil
}

func (p *player) loadTrack(file *pb.AudioFile, trackId []byte) (*audioFile, error) {
   audioFile := newAudioFileWithIdAndFormat(file.FileId, file.GetFormat(), p)
   // Start loading the audio key
   err := audioFile.loadKey(trackId)
   if err != nil {
      return nil, err
   }
   // Then start loading the audio itself
   audioFile.loadChunks()
   return audioFile, nil
}

type audioFile struct {
   chunkLoadOrder []int
   chunkLock      sync.RWMutex
   chunks         map[int]bool
   chunksLoading  bool
   cipher         cipher.Block
   cursor         int
   data           []byte
   decrypter      *audioFileDecrypter
   fileId         []byte
   format         pb.AudioFile_Format
   lock           sync.RWMutex
   player         *player
   responseChan   chan []byte
   size           uint32
}

func newAudioFileWithIdAndFormat(fileId []byte, format pb.AudioFile_Format, player *player) *audioFile {
   return &audioFile{
      chunkLock:     sync.RWMutex{},
      chunks:        map[int]bool{},
      chunksLoading: false,
      decrypter:     newAudioFileDecrypter(),
      fileId:        fileId,
      format:        format,
      player:        player,
      responseChan:  make(chan []byte),
      // Set an initial size to fetch the first chunk regardless of the actual size
      size: chunkSizeK,
   }
}

func (a *audioFile) headerOffset() int {
   switch a.format {
   case
   pb.AudioFile_OGG_VORBIS_160,
   pb.AudioFile_OGG_VORBIS_320,
   pb.AudioFile_OGG_VORBIS_96:
      return oggSkipBytesK
   }
   return 0
}
