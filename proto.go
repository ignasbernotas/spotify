package spotify

import (
   "encoding/binary"
   "encoding/hex"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
   "math/big"
   "os"
)

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
      username, []byte(password),
      pb.AuthenticationType_AUTHENTICATION_UNKNOWN.Enum(), s.deviceId,
   )
   if err != nil {
      return err
   }
   return s.doLogin(loginPacket, username)
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
