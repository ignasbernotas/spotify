package spotify

import (
   "crypto/cipher"
   "encoding/binary"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "log"
   "sync"
   "time"
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

func makeHelloMessage(publicKey []byte, nonce []byte) []byte {
   hello := &pb.ClientHello{
      BuildInfo: &pb.BuildInfo{
         Platform:     pb.Platform_PLATFORM_LINUX_X86_64.Enum(),
         // authentication failed: PremiumAccountRequired
         // Product: pb.Product_PRODUCT_PARTNER.Enum(),
         // CHANGE THIS TO MAKE LIBRESPOT WORK WITH FREE ACCOUNTS
         Product: pb.Product_PRODUCT_CLIENT.Enum(),
         ProductFlags: []pb.ProductFlags{pb.ProductFlags_PRODUCT_FLAG_NONE},
         Version:      proto.Uint64(0x10800000000),
      },
      FingerprintsSupported: []pb.Fingerprint{},
      CryptosuitesSupported: []pb.Cryptosuite{
         pb.Cryptosuite_CRYPTO_SUITE_SHANNON,
      },
      LoginCryptoHello: &pb.LoginCryptoHelloUnion{
         DiffieHellman: &pb.LoginCryptoDiffieHellmanHello{
            Gc:              publicKey,
            ServerKeysKnown: proto.Uint32(1),
         },
      },
      ClientNonce: nonce,
      FeatureSet: &pb.FeatureSet{
         Autoupdate2: proto.Bool(true),
      },
      Padding: []byte{0x1e},
   }
   packetData, err := proto.Marshal(hello)
   if err != nil {
      log.Fatal("login marshaling error: ", err)
   }
   return packetData
}

func encodeRequest(seq []byte, req request) ([]byte, error) {
	buf, err := encodeMercuryHead(seq, uint16(1+len(req.Payload)), uint8(1))
	if err != nil {
		return nil, err
	}

	header := &pb.Header{
		Uri:    proto.String(req.Uri),
		Method: proto.String(req.Method),
	}

	if req.ContentType != "" {
		header.ContentType = proto.String(req.ContentType)
	}

	headerData, err := proto.Marshal(header)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint16(len(headerData)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(headerData)
	if err != nil {
		return nil, err
	}

	for _, p := range req.Payload {
		err = binary.Write(buf, binary.BigEndian, uint16(len(p)))
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

func (m *internal) completeRequest(cmd uint8, pending pending, seqKey string) (*response, error) {
	headerData := pending.parts[0]
	header := &pb.Header{}
	err := proto.Unmarshal(headerData, header)
	if err != nil {
		return nil, err
	}

	return &response{
		HeaderData: headerData,
		Uri:        *header.Uri,
		Payload:    pending.parts[1:],
		StatusCode: header.GetStatusCode(),
		SeqKey:     seqKey,
	}, nil

}

func makeLoginBlobPacket(username string, authData []byte, authType *pb.AuthenticationType, deviceId string) []byte {
	packet := &pb.ClientResponseEncrypted{
		LoginCredentials: &pb.LoginCredentials{
			Username: proto.String(username),
			Typ:      authType,
			AuthData: authData,
		},
		AccountCreation: pb.AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT.Enum(),
		SystemInfo: &pb.SystemInfo{
			CpuFamily:               pb.CpuFamily_CPU_X86_64.Enum(),
			CpuSubtype:              proto.Uint32(0),
			Brand:                   pb.Brand_BRAND_UNBRANDED.Enum(),
			BrandFlags:              proto.Uint32(0),
			Os:                      pb.Os_OS_LINUX.Enum(),
			OsVersion:               proto.Uint32(0),
			OsExt:                   proto.Uint32(0),
			SystemInformationString: proto.String("Linux [x86-64 0]"),
			DeviceId:                proto.String("libspotify"),
		},
		PlatformModel: proto.String("PC desktop"),
		VersionString: proto.String("1.1.10.546.ge08ef575"),
		ClientInfo: &pb.ClientInfo{
			Limited:  proto.Bool(false),
			Language: proto.String("en"),
		},
	}
	packetData, err := proto.Marshal(packet)
	if err != nil {
		log.Fatal("login marshaling error: ", err)
	}
	return packetData
}

func getTrackInfo(track *pb.Track) *spotifyTrack {
   enc := new(spotifyTrack)
   enc.TrackName = track.GetName()
   enc.TrackNumber = track.GetNumber()
   // convert ms to seconds
   enc.TrackDuration = (track.GetDuration() / 1000)
   enc.TrackDiscNumber = track.GetDiscNumber()
   album := track.GetAlbum()
   if album != nil {
      enc.Album.Name = album.GetName()
      enc.Album.Label = album.GetLabel()
      enc.Album.Genre = album.GetGenre()
      albumDate := album.GetDate()
      if albumDate != nil {
         enc.Album.Date = time.Date(
            int(albumDate.GetYear()),
            time.Month(int(albumDate.GetMonth())),
            int(albumDate.GetDay()), 0, 0, 0, 0, time.UTC,
         )
      }
      albumArtists := album.GetArtist()
      for _, artist := range albumArtists {
         enc.Album.ArtistNames = append(
            enc.Album.ArtistNames, artist.GetName(),
         )
      }
   }
   trackArtists := track.GetArtist()
   for _, artist := range trackArtists {
      enc.TrackArtistNames = append(
         enc.TrackArtistNames, artist.GetName(),
      )
   }
   return enc
}
