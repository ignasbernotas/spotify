package spotify

import (
   "bytes"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/89z/spotify/crypto"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "net"
   "time"
)

type Session struct {
	/// Constructor references
	mercuryConstructor func(conn crypto.PacketStream) *crypto.Client
	shannonConstructor func(keys crypto.SharedKeys, conn crypto.PlainConnection) crypto.PacketStream

	/// Managers and helpers
	stream crypto.PacketStream
	mercury *crypto.Client
	discovery *crypto.Discovery
	player *crypto.Player
	tcpCon io.ReadWriter
	// keys are the encryption keys used to communicate with the server
	keys crypto.PrivateKeys

	/// State and variables
	// servers for this session
	deviceId string
	deviceName string
	// username is the currently authenticated canonical username
	username string
	reusableAuthBlob []byte
	country string
}

func (s *Session) Stream() crypto.PacketStream {
	return s.stream
}

func (s *Session) Discovery() *crypto.Discovery {
	return s.discovery
}

func (s *Session) Mercury() *crypto.Client {
	return s.mercury
}

func (s *Session) Player() *crypto.Player {
	return s.player
}

func (s *Session) Username() string {
	return s.username
}

func (s *Session) DeviceId() string {
	return s.deviceId
}

func (s *Session) ReusableAuthBlob() []byte {
	return s.reusableAuthBlob
}

func (s *Session) Country() string {
	return s.country
}

func (s *Session) startConnection() error {
	conn := crypto.MakePlainConnection(s.tcpCon, s.tcpCon)

	helloMessage := makeHelloMessage(s.keys.PubKey(), s.keys.ClientNonce())
	initClientPacket, err := conn.SendPrefixPacket([]byte{0, 4}, helloMessage)
	if err != nil {
		log.Fatal("Error writing client hello", err)
		return err
	}

	// Wait and read the hello reply
	initServerPacket, err := conn.RecvPacket()
	if err != nil {
		log.Fatal("Error receving packet for hello: ", err)
		return err
	}

	response := pb.APResponseMessage{}
	err = proto.Unmarshal(initServerPacket[4:], &response)
	if err != nil {
		log.Fatal("Failed to unmarshal server hello", err)
		return err
	}

	remoteKey := response.Challenge.LoginCryptoChallenge.DiffieHellman.Gs
	sharedKeys := s.keys.AddRemoteKey(remoteKey, initClientPacket, initServerPacket)

	plainResponse := &pb.ClientResponsePlaintext{
		LoginCryptoResponse: &pb.LoginCryptoResponseUnion{
			DiffieHellman: &pb.LoginCryptoDiffieHellmanResponse{
				Hmac: sharedKeys.Challenge(),
			},
		},
		PowResponse:    &pb.PoWResponseUnion{},
		CryptoResponse: &pb.CryptoResponseUnion{},
	}

	plainResponseMessage, err := proto.Marshal(plainResponse)
	if err != nil {
		log.Fatal("marshaling error: ", err)
		return err
	}

	_, err = conn.SendPrefixPacket([]byte{}, plainResponseMessage)
	if err != nil {
		log.Fatal("error writing client plain response ", err)
		return err
	}

	s.stream = s.shannonConstructor(sharedKeys, conn)
	s.mercury = s.mercuryConstructor(s.stream)

	s.player = crypto.CreatePlayer(s.stream, s.mercury)

	return nil
}

func setupSession() (*Session, error) {
	session := &Session{
		keys:               crypto.GenerateKeys(),
		mercuryConstructor: crypto.CreateMercury,
		shannonConstructor: crypto.CreateStream,
	}
	err := session.doConnect()

	return session, err
}

func sessionFromDiscovery(d *crypto.Discovery) (*Session, error) {
	s, err := setupSession()
	if err != nil {
		return nil, err
	}

	s.discovery = d
	s.deviceId = d.DeviceId()
	s.deviceName = d.DeviceName()

	err = s.startConnection()
	if err != nil {
		return s, err
	}

	loginPacket := s.getLoginBlobPacket(d.LoginBlob())
	return s, s.doLogin(loginPacket, d.LoginBlob().Username)
}

func (s *Session) doConnect() error {
	apUrl, err := APResolve()
	if err != nil {
		log.Println("Failed to get ap url", err)
		return err
	}

	s.tcpCon, err = net.Dial("tcp", apUrl)
	if err != nil {
		log.Println("Failed to connect:", err)
		return err
	}

	return err
}

func (s *Session) disconnect() {
	if s.tcpCon != nil {
		conn := s.tcpCon.(net.Conn)
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close tcp connection", err)
		}
		s.tcpCon = nil
	}
}

func (s *Session) doReconnect() error {
	s.disconnect()

	err := s.doConnect()
	if err != nil {
		return err
	}

	err = s.startConnection()
	if err != nil {
		return err
	}

	packet := makeLoginBlobPacket(s.username, s.reusableAuthBlob,
		pb.AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS.Enum(), s.deviceId)
	return s.doLogin(packet, s.username)
}

func (s *Session) planReconnect() {
	go func() {
		time.Sleep(1 * time.Second)

		if err := s.doReconnect(); err != nil {
			// Try to reconnect again in a second
			s.planReconnect()
		}
	}()
}

func (s *Session) runPollLoop() {
	for {
		cmd, data, err := s.stream.RecvPacket()
		if err != nil {
			log.Println("Error during RecvPacket: ", err)

			if err == io.EOF {
				// We've been disconnected, reconnect
				s.planReconnect()
				break
			}
		} else {
			s.handle(cmd, data)
		}
	}
}

func (s *Session) handle(cmd uint8, data []byte) {
	// log.Printf("handle, cmd=0x%x data=%x\n", cmd, data)
	// log.Printf("Handle: 0x%x", cmd)

	switch {
	case cmd == crypto.PacketPing:
		// Ping
		err := s.stream.SendPacket(crypto.PacketPong, data)
		if err != nil {
			log.Fatal("Error handling PacketPing", err)
		}

	case cmd == crypto.PacketPongAck:
		// Pong reply, ignore

	case cmd == crypto.PacketAesKey || cmd == crypto.PacketAesKeyError ||
		cmd == crypto.PacketStreamChunkRes:
		// Audio key and data responses
		s.player.HandleCmd(cmd, data)

	case cmd == crypto.PacketCountryCode:
		// Handle country code
		s.country = fmt.Sprintf("%s", data)

	case 0xb2 <= cmd && cmd <= 0xb6:
		err := s.mercury.Handle(cmd, bytes.NewReader(data))
		if err != nil {
			log.Fatal("Handle 0xbx", err)
		}

	case cmd == crypto.PacketSecretBlock:
		// Old RSA public key

	case cmd == crypto.PacketLegacyWelcome:
		// Empty welcome packet

	case cmd == crypto.PacketProductInfo:
		// Has some info about A/B testing status, product setup, etc... in an XML fashion.

	case cmd == 0x1f:
		// Unknown, data is zeroes only

	case cmd == crypto.PacketLicenseVersion:
	default:
		fmt.Printf("Unhandled cmd 0x%x\n", cmd)
	}
}

func (s *Session) poll() {
	cmd, data, err := s.stream.RecvPacket()
	if err != nil {
		log.Fatal("poll error", err)
	}
	s.handle(cmd, data)
}

func readInt(b *bytes.Buffer) uint32 {
	c, _ := b.ReadByte()
	lo := uint32(c)
	if lo&0x80 == 0 {
		return lo
	}

	c2, _ := b.ReadByte()
	hi := uint32(c2)
	return lo&0x7f | hi<<7
}

func readBytes(b *bytes.Buffer) []byte {
	length := readInt(b)
	data := make([]byte, length)
	b.Read(data)

	return data
}

func makeHelloMessage(publicKey []byte, nonce []byte) []byte {
	hello := &pb.ClientHello{
		BuildInfo: &pb.BuildInfo{
			Product:      pb.Product_PRODUCT_CLIENT.Enum(), // CHANGE THIS TO MAKE LIBRESPOT WORK WITH FREE ACCOUNTS
			ProductFlags: []pb.ProductFlags{pb.ProductFlags_PRODUCT_FLAG_NONE},
			Platform:     pb.Platform_PLATFORM_LINUX_X86_64.Enum(),
			Version:      proto.Uint64(0x10800000000),
		},
		FingerprintsSupported: []pb.Fingerprint{},
		CryptosuitesSupported: []pb.Cryptosuite{
			pb.Cryptosuite_CRYPTO_SUITE_SHANNON},
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
