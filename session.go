package spotify

import (
   "bytes"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "net"
   "time"
)

type Session struct {
	/// Constructor references
	mercuryConstructor func(conn packetStream) *client
	shannonConstructor func(keys sharedKeys, conn plainConnection) packetStream
	/// Managers and helpers
	stream packetStream
	mercury *client
	discovery *discovery
	player *player
	tcpCon io.ReadWriter
	// keys are the encryption keys used to communicate with the server
	keys privateKeys

	/// State and variables
	// servers for this session
	deviceId string
	deviceName string
	// username is the currently authenticated canonical username
	username string
	reusableAuthBlob []byte
	country string
}

func (s *Session) startConnection() error {
   conn := makePlainConnection(s.tcpCon, s.tcpCon)
   helloMessage := makeHelloMessage(s.keys.pubKey(), s.keys.clientNonce)
   initClientPacket, err := conn.sendPrefixPacket([]byte{0, 4}, helloMessage)
   if err != nil {
      log.Fatal("Error writing client hello", err)
      return err
   }
   // Wait and read the hello reply
   initServerPacket, err := conn.recvPacket()
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
   sharedKeys := s.keys.addRemoteKey(
      remoteKey, initClientPacket, initServerPacket,
   )
   plainResponse := &pb.ClientResponsePlaintext{
      CryptoResponse: &pb.CryptoResponseUnion{},
      LoginCryptoResponse: &pb.LoginCryptoResponseUnion{
         DiffieHellman: &pb.LoginCryptoDiffieHellmanResponse{
            Hmac: sharedKeys.challenge,
         },
      },
      PowResponse:    &pb.PoWResponseUnion{},
   }
   plainResponseMessage, err := proto.Marshal(plainResponse)
   if err != nil {
   log.Fatal("marshaling error: ", err)
   return err
   }
   _, err = conn.sendPrefixPacket([]byte{}, plainResponseMessage)
   if err != nil {
   log.Fatal("error writing client plain response ", err)
   return err
   }
   s.stream = s.shannonConstructor(sharedKeys, conn)
   s.mercury = s.mercuryConstructor(s.stream)
   s.player = createPlayer(s.stream, s.mercury)
   return nil
}

func (s *Session) doConnect() error {
   apUrl, err := apResolve()
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
   switch {
   case cmd == packetPing:
      // Ping
      err := s.stream.SendPacket(packetPong, data)
      if err != nil {
      log.Fatal("Error handling PacketPing", err)
      }
   case cmd == packetPongAck:
      // Pong reply, ignore
   case cmd == packetAesKey, cmd == packetAesKeyError, cmd == packetStreamChunkRes:
      // Audio key and data responses
      s.player.handleCmd(cmd, data)
   case cmd == packetCountryCode:
      s.country = fmt.Sprintf("%s", data)
   case 0xb2 <= cmd && cmd <= 0xb6:
      err := s.mercury.handle(cmd, bytes.NewReader(data))
      if err != nil {
      log.Fatal("Handle 0xbx", err)
      }
   case cmd == packetSecretBlock:
      // Old RSA public key
   case cmd == packetLegacyWelcome:
      // Empty welcome packet
   case cmd == packetProductInfo:
      // Has some info about A/B testing status, product setup, etc... in an
      // XML fashion.
   case cmd == 0x1f:
      // Unknown, data is zeroes only
   case cmd == packetLicenseVersion:
   default:
      fmt.Printf("Unhandled cmd 0x%x\n", cmd)
   }
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
