package spotify

import (
   "bytes"
   "fmt"
   "io"
   "log"
   "net"
)

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
