package spotify

import (
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "math/big"
   "os"
   "time"
)

func Login(username string, password string, deviceName string) (*Session, error) {
   private := new(big.Int)
   private.SetBytes(RandomVec(95))
   DH_GENERATOR := big.NewInt(0x2)
   DH_PRIME := new(big.Int)
   // datatracker.ietf.org/doc/html/rfc2412#appendix-E.1
   DH_PRIME.SetBytes([]byte{
      0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc9, 0x0f, 0xda, 0xa2,
      0x21, 0x68, 0xc2, 0x34, 0xc4, 0xc6, 0x62, 0x8b, 0x80, 0xdc, 0x1c, 0xd1,
      0x29, 0x02, 0x4e, 0x08, 0x8a, 0x67, 0xcc, 0x74, 0x02, 0x0b, 0xbe, 0xa6,
      0x3b, 0x13, 0x9b, 0x22, 0x51, 0x4a, 0x08, 0x79, 0x8e, 0x34, 0x04, 0xdd,
      0xef, 0x95, 0x19, 0xb3, 0xcd, 0x3a, 0x43, 0x1b, 0x30, 0x2b, 0x0a, 0x6d,
      0xf2, 0x5f, 0x14, 0x37, 0x4f, 0xe1, 0x35, 0x6d, 0x6d, 0x51, 0xc2, 0x45,
      0xe4, 0x85, 0xb5, 0x76, 0x62, 0x5e, 0x7e, 0xc6, 0xf4, 0x4c, 0x42, 0xe9,
      0xa6, 0x3a, 0x36, 0x20, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
   })
   ses := &Session{
      keys: PrivateKeys{
         clientNonce: RandomVec(0x10),
         generator:   DH_GENERATOR,
         prime:       DH_PRIME,
         privateKey: private,
         publicKey:  Powm(DH_GENERATOR, private, DH_PRIME),
      },
      mercuryConstructor: createMercury,
      shannonConstructor: CreateStream,
   }
   err := ses.doConnect()
   if err != nil {
      return nil, err
   }
   if err := ses.loginSession(username, password, deviceName); err != nil {
      return nil, err
   }
   return ses, nil
}

func DownloadTrackID(ses *Session, id string) error {
   hex := fmt.Sprintf("%x", Convert62(id))
   tra, err := ses.Mercury().getTrack(hex)
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
   audioFile, err := ses.Player().loadTrack(selectedFile, tra.GetGid())
   if err != nil {
      return fmt.Errorf("failed to download the track %v", err)
   }
   track := GetTrackInfo(tra)
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

func (s *Session) loginSession(username string, password string, deviceName string) error {
	s.deviceId = GenerateDeviceId(deviceName)
	s.deviceName = deviceName

	err := s.startConnection()
	if err != nil {
		return err
	}
	loginPacket := makeLoginPasswordPacket(username, password, s.deviceId)
	return s.doLogin(loginPacket, username)
}

func (s *Session) doLogin(packet []byte, username string) error {
	err := s.stream.SendPacket(packetLogin, packet)
	if err != nil {
		log.Fatal("bad shannon write", err)
	}

	// Pll once for authentication response
	welcome, err := s.handleLogin()
	if err != nil {
		return err
	}

	// Store the few interesting values
	s.username = welcome.GetCanonicalUsername()
	if s.username == "" {
		s.username = s.discovery.LoginBlob().Username
	}
	s.reusableAuthBlob = welcome.GetReusableAuthCredentials()

	// Poll for acknowledge before loading - needed for gopherjs
	go s.runPollLoop()

	return nil
}

func (s *Session) handleLogin() (*pb.APWelcome, error) {
	cmd, data, err := s.stream.RecvPacket()
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	if cmd == packetAuthFailure {
		failure := &pb.APLoginFailed{}
		err := proto.Unmarshal(data, failure)
		if err != nil {
			return nil, fmt.Errorf("authenticated failed: %v", err)
		}
		return nil, fmt.Errorf("authentication failed: %s", failure.ErrorCode)
	} else if cmd == packetAPWelcome {
		welcome := &pb.APWelcome{}
		err := proto.Unmarshal(data, welcome)
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %v", err)
		}
		return welcome, nil
	} else {
		return nil, fmt.Errorf("authentication failed: unexpected cmd %v", cmd)
	}
}

func makeLoginPasswordPacket(username string, password string, deviceId string) []byte {
	return makeLoginBlobPacket(username, []byte(password),
		pb.AuthenticationType_AUTHENTICATION_UNKNOWN.Enum(), deviceId)
}

func makeLoginBlobPacket(username string, authData []byte,
	authType *pb.AuthenticationType, deviceId string) []byte {
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

func GetTrackInfo(track *pb.Track) *SpotifyTrack {
   enc := new(SpotifyTrack)
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

