package spotify

import (
   "encoding/binary"
   "encoding/hex"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "io"
   "math/big"
   "os"
)

func (m *internal) parseResponse(cmd uint8, reader io.Reader) (*response, error) {
   seq, flags, count, err := handleHead(reader)
   if err != nil {
      return nil, err
   }
   seqKey := string(seq)
   pend := m.Pending[seqKey]
   for i := uint16(0); i < count; i++ {
      part, err := parsePart(reader)
      if err != nil {
         return nil, err
      }
      if pend.partial != nil {
         part = append(pend.partial, part...)
         pend.partial = nil
      }
      if i == count-1 && (flags == 2) {
         pend.partial = part
      } else {
         pend.parts = append(pend.parts, part)
      }
   }
   if flags == 1 {
      delete(m.Pending, seqKey)
      hData := pend.parts[0]
      var head pb.Header
      fmt.Printf("%q\n", hData)
      err := proto.Unmarshal(hData, &head)
      if err != nil {
         return nil, err
      }
      return &response{
         headerData: hData,
         payload: pend.parts[1:],
         seqKey: seqKey,
         statusCode: head.GetStatusCode(),
         uri: *head.Uri,
      }, nil
   } else {
      m.Pending[seqKey] = pend
   }
   return nil, nil
}
func encodeRequest(seq []byte, req request) ([]byte, error) {
   buf, err := encodeMercuryHead(seq, uint16(1+len(req.payload)), uint8(1))
   if err != nil {
      return nil, err
   }
   header := &pb.Header{
      Method: proto.String(req.method),
      Uri:    proto.String(req.uri),
   }
   if req.contentType != "" {
      header.ContentType = proto.String(req.contentType)
   }
   hData, err := proto.Marshal(header)
   if err != nil {
      return nil, err
   }
   err = binary.Write(buf, binary.BigEndian, uint16(len(hData)))
   if err != nil {
      return nil, err
   }
   _, err = buf.Write(hData)
   if err != nil {
      return nil, err
   }
   for _, p := range req.payload {
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

func makeLoginBlobPacket(username string, authData []byte, authType *pb.AuthenticationType, deviceId string) ([]byte, error) {
   packet := &pb.ClientResponseEncrypted{
      LoginCredentials: &pb.LoginCredentials{
         Username: proto.String(username),
         Typ:      authType,
         AuthData: authData,
      },
      AccountCreation: pb.AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT.Enum(),
      SystemInfo: &pb.SystemInfo{
         Brand:                   pb.Brand_BRAND_UNBRANDED.Enum(),
         BrandFlags:              proto.Uint32(0),
         CpuFamily:               pb.CpuFamily_CPU_X86_64.Enum(),
         CpuSubtype:              proto.Uint32(0),
         DeviceId:                proto.String("libspotify"),
         Os:                      pb.Os_OS_LINUX.Enum(),
         OsExt:                   proto.Uint32(0),
         OsVersion:               proto.Uint32(0),
         SystemInformationString: proto.String("Linux [x86-64 0]"),
      },
      PlatformModel: proto.String("PC desktop"),
      VersionString: proto.String("1.1.10.546.ge08ef575"),
      ClientInfo: &pb.ClientInfo{
         Language: proto.String("en"),
         Limited:  proto.Bool(false),
      },
   }
   return proto.Marshal(packet)
}

func getFormat(track pb.Track) (*pb.AudioFile, error) {
   for _, file := range track.GetFile() {
      if file.GetFormat() == AudioFile_OGG_VORBIS_160 {
         return file, nil
      }
   }
   msg := "could not find any files of the song in the specified formats"
   return nil, fmt.Errorf(msg)
}

func (ses *session) DownloadTrackID(id string) error {
   b62 := new(big.Int)
   b62.SetString(id, 62)
   id = hex.EncodeToString(b62.Bytes())
   addr := "hm://metadata/4/track/" + id
   fmt.Println("GET", addr)
   data := ses.mercury.mercuryGet(addr)
   var trk pb.Track
   err := proto.Unmarshal(data, &trk)
   if err != nil {
      return err
   }
   fSelect, err := getFormat(trk)
   if err != nil {
      return err
   }
   trackID := trk.GetGid()
   aFile := newAudioFileWithIdAndFormat(
      fSelect.FileId,
      AudioFile_OGG_VORBIS_160,
      ses.player,
   )
   // Start loading the audio key
   if err := aFile.loadKey(trackID); err != nil {
      return err
   }
   // Then start loading the audio itself
   aFile.loadChunks()
   file, err := os.Create("file.ogg")
   if err != nil {
      return err
   }
   defer file.Close()
   if _, err := file.ReadFrom(aFile); err != nil {
      return err
   }
   return nil
}

func (s *session) loginSession(username string, password string, deviceName string) error {
   s.deviceId = generateDeviceId(deviceName)
   s.deviceName = deviceName
   err := s.startConnection()
   if err != nil {
      return err
   }
   loginPacket, err := makeLoginBlobPacket(
      username,
      []byte(password),
      pb.AuthenticationType_AUTHENTICATION_UNKNOWN.Enum(),
      s.deviceId,
   )
   if err != nil {
      return err
   }
   return s.doLogin(loginPacket, username)
}

func (s *session) doLogin(packet []byte, username string) error {
   err := s.stream.sendPacket(packetLogin, packet)
   if err != nil {
      return err
   }
   // Pll once for authentication response
   cmd, data, err := s.stream.recvPacket()
   if err != nil {
      return err
   }
   switch cmd {
   case packetAuthFailure:
      failure := &pb.APLoginFailed{}
      err := proto.Unmarshal(data, failure)
      if err != nil {
         return err
      }
      return fmt.Errorf("errorCode %v", failure.ErrorCode)
   case packetAPWelcome:
      welcome := new(pb.APWelcome)
      err := proto.Unmarshal(data, welcome)
      if err != nil {
         return fmt.Errorf("authentication failed: %v", err)
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
   return fmt.Errorf("authentication failed: unexpected cmd %v", cmd)
}

func makeHelloMessage(publicKey []byte, nonce []byte) ([]byte, error) {
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
   return proto.Marshal(hello)
}

func (s *session) startConnection() error {
   conn := makePlainConnection(s.tcpCon, s.tcpCon)
   helloMessage, err := makeHelloMessage(
      s.keys.publicKey.Bytes(),
      s.keys.clientNonce,
   )
   if err != nil {
      return err
   }
   initClientPacket, err := conn.sendPrefixPacket([]byte{0, 4}, helloMessage)
   if err != nil {
      return err
   }
   // Wait and read the hello reply
   initServerPacket, err := conn.recvPacket()
   if err != nil {
      return err
   }
   response := pb.APResponseMessage{}
   err = proto.Unmarshal(initServerPacket[4:], &response)
   if err != nil {
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
      return err
   }
   _, err = conn.sendPrefixPacket([]byte{}, plainResponseMessage)
   if err != nil {
      return err
   }
   s.stream = s.shannonConstructor(sharedKeys, conn)
   s.mercury = s.mercuryConstructor(s.stream)
   s.player = createPlayer(s.stream, s.mercury)
   return nil
}
