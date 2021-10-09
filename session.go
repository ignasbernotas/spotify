package spotify

import (
   "bytes"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "log"
   "math/big"
   "net"
   "os"
   "time"
)

type session struct {
   /// Constructor references
   mercuryConstructor func(conn packetStream) *client
   shannonConstructor func(keys sharedKeys, conn plainConnection) packetStream
   /// Managers and helpers
   stream packetStream
   mercury *client
   player *player
   tcpCon io.ReadWriter
   // keys are the encryption keys used to communicate with the server
   keys privateKeys
   /// State and variables
   deviceId string
   deviceName string
   // username is the currently authenticated canonical username
   username string
   reusableAuthBlob []byte
   country string
   discovery *blobInfo
}

func Login(username string, password string, deviceName string) (*session, error) {
   private := new(big.Int)
   private.SetBytes(randomVec(95))
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
   ses := &session{
      keys: privateKeys{
         clientNonce: randomVec(0x10),
         generator:   DH_GENERATOR,
         prime:       DH_PRIME,
         privateKey: private,
         publicKey:  powm(DH_GENERATOR, private, DH_PRIME),
      },
      mercuryConstructor: createMercury,
      shannonConstructor: createStream,
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

func (s *session) disconnect() {
	if s.tcpCon != nil {
		conn := s.tcpCon.(net.Conn)
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close tcp connection", err)
		}
		s.tcpCon = nil
	}
}

func (s *session) doConnect() error {
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

func (s *session) doReconnect() error {
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

func (s *session) handle(cmd uint8, data []byte) {
   switch {
   case cmd == packetPing:
      // Ping
      err := s.stream.sendPacket(packetPong, data)
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

func (s *session) planReconnect() {
	go func() {
		time.Sleep(1 * time.Second)

		if err := s.doReconnect(); err != nil {
			// Try to reconnect again in a second
			s.planReconnect()
		}
	}()
}

func (s *session) runPollLoop() {
	for {
		cmd, data, err := s.stream.recvPacket()
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

func (s *session) startConnection() error {
   conn := makePlainConnection(s.tcpCon, s.tcpCon)
   helloMessage := makeHelloMessage(
      s.keys.publicKey.Bytes(),
      s.keys.clientNonce,
   )
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

func (s *session) loginSession(username string, password string, deviceName string) error {
   s.deviceId = generateDeviceId(deviceName)
   s.deviceName = deviceName
   err := s.startConnection()
   if err != nil {
      return err
   }
   loginPacket := makeLoginBlobPacket(
      username, []byte(password),
      pb.AuthenticationType_AUTHENTICATION_UNKNOWN.Enum(), s.deviceId,
   )
   return s.doLogin(loginPacket, username)
}

func (s *session) doLogin(packet []byte, username string) error {
   err := s.stream.sendPacket(packetLogin, packet)
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
      s.username = s.discovery.Username
   }
   s.reusableAuthBlob = welcome.GetReusableAuthCredentials()
   // Poll for acknowledge before loading - needed for gopherjs
   go s.runPollLoop()
   return nil
}

func (s *session) handleLogin() (*pb.APWelcome, error) {
	cmd, data, err := s.stream.recvPacket()
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
