package spotify

import (
   "encoding/binary"
   "fmt"
   "github.com/89z/spotify/pb"
   "github.com/golang/protobuf/proto"
)

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
      if file.GetFormat() == audioFile_OGG_VORBIS_160 {
         return file, nil
      }
   }
   msg := "could not find any files of the song in the specified formats"
   return nil, fmt.Errorf(msg)
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

