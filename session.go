package spotify

import (
   "bytes"
   "fmt"
   "io"
   "log"
   "math/big"
   "net"
)

const (
   AudioFile_OGG_VORBIS_96   = 0
   AudioFile_OGG_VORBIS_160  = 1
   AudioFile_OGG_VORBIS_320  = 2
)

func (s *session) doConnect() error {
   con, err := net.Dial("tcp", "ap.spotify.com:80")
   if err != nil {
      return err
   }
   s.tcpCon = con
   return nil
}

func Login(username string, password string, deviceName string) (*session, error) {
   private := new(big.Int)
   ran, err := randomVec(95)
   if err != nil {
      return nil, err
   }
   private.SetBytes(ran)
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
   non, err := randomVec(0x10)
   if err != nil {
      return nil, err
   }
   ses := &session{
      keys: privateKeys{
         clientNonce: non,
         generator:   DH_GENERATOR,
         prime:       DH_PRIME,
         privateKey: private,
         publicKey:  powm(DH_GENERATOR, private, DH_PRIME),
      },
      mercuryConstructor: createMercury,
      shannonConstructor: createStream,
   }
   if err := ses.doConnect(); err != nil {
      return nil, err
   }
   if err := ses.loginSession(username, password, deviceName); err != nil {
      return nil, err
   }
   return ses, nil
}


func (s *session) handle(cmd uint8, data []byte) {
   switch cmd {
   case 0x1f:
      // Unknown, data is zeroes only
   case packetAesKey, packetAesKeyError, packetStreamChunkRes:
      // Audio key and data responses
      s.player.handleCmd(cmd, data)
   case packetCountryCode:
      s.country = string(data)
   case packetLegacyWelcome:
      // Empty welcome packet
   case packetLicenseVersion:
   case packetPing:
      // Ping
      err := s.stream.sendPacket(packetPong, data)
      if err != nil {
         log.Fatal("Error handling PacketPing", err)
      }
   case packetPongAck:
      // Pong reply, ignore
   case packetProductInfo:
      // Has some info about A/B testing status, product setup, etc... in an
      // XML fashion.
   case packetSecretBlock:
      // Old RSA public key
   default:
      if 0xb2 <= cmd && cmd <= 0xb6 {
         err := s.mercury.handle(cmd, bytes.NewReader(data))
         if err != nil {
            log.Fatal("Handle 0xbx", err)
         }
      } else {
         fmt.Printf("Unhandled cmd 0x%x\n", cmd)
      }
   }
}

func (s *session) runPollLoop() {
   for {
      cmd, data, err := s.stream.recvPacket()
      if err != nil {
         log.Println("Error during RecvPacket: ", err)
         if err == io.EOF {
            break
         }
      } else {
         s.handle(cmd, data)
      }
   }
}

func (s *session) doLogin(packet []byte, username string) error {
   err := s.stream.sendPacket(packetLogin, packet)
   if err != nil {
      return err
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
