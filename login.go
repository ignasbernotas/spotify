package spotify

import (
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "os"
   "time"
)

var Version = "master"
var BuildID = "dev"

// NEED THIS
func CoreLogin(username string, password string, deviceName string) (*Session, error) {
	s, err := setupSession()
	if err != nil {
		return s, err
	}

	return s, s.loginSession(username, password, deviceName)
}

// NEED THIS
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

func CoreLoginSaved(username string, authData []byte, deviceName string) (*Session, error) {
	s, err := setupSession()
	if err != nil {
		return s, err
	}
	s.deviceId = GenerateDeviceId(deviceName)
	s.deviceName = deviceName

	err = s.startConnection()
	if err != nil {
		return s, err
	}

	packet := makeLoginBlobPacket(username, authData,
		pb.AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS.Enum(), s.deviceId)
	return s, s.doLogin(packet, username)
}

// NEED THIS
func (s *Session) doLogin(packet []byte, username string) error {
	err := s.stream.SendPacket(PacketLogin, packet)
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
	// s.poll()
	go s.runPollLoop()

	return nil
}

func (s *Session) handleLogin() (*pb.APWelcome, error) {
	cmd, data, err := s.stream.RecvPacket()
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	if cmd == PacketAuthFailure {
		failure := &pb.APLoginFailed{}
		err := proto.Unmarshal(data, failure)
		if err != nil {
			return nil, fmt.Errorf("authenticated failed: %v", err)
		}
		return nil, fmt.Errorf("authentication failed: %s", failure.ErrorCode)
	} else if cmd == PacketAPWelcome {
		welcome := &pb.APWelcome{}
		err := proto.Unmarshal(data, welcome)
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %v", err)
		}
		// fmt.Println("Authentication succeeded: Welcome,", welcome.GetCanonicalUsername())
		// fmt.Println("Blob type:", welcome.GetReusableAuthCredentialsType())
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

	// TODO: Fix PremiumAccountRequired
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
