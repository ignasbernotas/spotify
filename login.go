package spotify

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/89z/spotify/crypto"
   "github.com/golang/protobuf/proto"
   "io/ioutil"
   "log"
   "net/http"
   "net/url"
)

var Version = "master"
var BuildID = "dev"

func CoreLogin(username string, password string, deviceName string) (*Session, error) {
	s, err := setupSession()
	if err != nil {
		return s, err
	}

	return s, s.loginSession(username, password, deviceName)
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

func CoreLoginOAuth(deviceName string, clientId string, clientSecret string) (*Session, error) {
	token := getOAuthToken(clientId, clientSecret)
	return loginOAuthToken(token.AccessToken, deviceName)
}

func loginOAuthToken(accessToken string, deviceName string) (*Session, error) {
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

	packet := makeLoginBlobPacket("", []byte(accessToken),
		pb.AuthenticationType_AUTHENTICATION_SPOTIFY_TOKEN.Enum(), s.deviceId)
	return s, s.doLogin(packet, "")
}

func (s *Session) doLogin(packet []byte, username string) error {
	err := s.stream.SendPacket(crypto.PacketLogin, packet)
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

	if cmd == crypto.PacketAuthFailure {
		failure := &pb.APLoginFailed{}
		err := proto.Unmarshal(data, failure)
		if err != nil {
			return nil, fmt.Errorf("authenticated failed: %v", err)
		}
		return nil, fmt.Errorf("authentication failed: %s", failure.ErrorCode)
	} else if cmd == crypto.PacketAPWelcome {
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

func (s *Session) getLoginBlobPacket(blob crypto.BlobInfo) []byte {
	data, _ := base64.StdEncoding.DecodeString(blob.DecodedBlob)

	buffer := bytes.NewBuffer(data)
	buffer.ReadByte()
	readBytes(buffer)
	buffer.ReadByte()
	authNum := readInt(buffer)
	authType := pb.AuthenticationType(authNum)
	buffer.ReadByte()
	authData := readBytes(buffer)

	return makeLoginBlobPacket(blob.Username, authData, &authType, s.deviceId)
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


type OAuth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	Error        string
}

func GetOauthAccessToken(code string, redirectUri string, clientId string, clientSecret string) (*OAuth, error) {
	val := url.Values{}
	val.Set("grant_type", "authorization_code")
	val.Set("code", code)
	val.Set("redirect_uri", redirectUri)
	val.Set("client_id", clientId)
	val.Set("client_secret", clientSecret)

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", val)
	if err != nil {
		// Retry since there is an nginx bug that causes http2 streams to get
		// an initial REFUSED_STREAM response
		// https://github.com/curl/curl/issues/804
		resp, err = http.PostForm("https://accounts.spotify.com/api/token", val)
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()
	auth := OAuth{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return nil, err
	}
	if auth.Error != "" {
		return nil, fmt.Errorf("error getting token %v", auth.Error)
	}
	return &auth, nil
}

func getOAuthToken(clientId string, clientSecret string) OAuth {
	ch := make(chan OAuth)

	fmt.Println("go to this url")
	urlPath := "https://accounts.spotify.com/authorize?" +
		"client_id=" + clientId +
		"&response_type=code" +
		"&redirect_uri=http://localhost:8888/callback" +
		"&scope=streaming"
	fmt.Println(urlPath)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		auth, err := GetOauthAccessToken(params.Get("code"), "http://localhost:8888/callback", clientId, clientSecret)
		if err != nil {
			fmt.Fprintf(w, "Error getting token %q", err)
			return
		}
		fmt.Fprintf(w, "Got token, loggin in")
		ch <- *auth
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8888", nil))
	}()

	return <-ch
}
