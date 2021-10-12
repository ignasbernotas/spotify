package pb
import "github.com/golang/protobuf/proto"

type Product int32

const (
	Product_PRODUCT_CLIENT              Product = 0
	Product_PRODUCT_LIBSPOTIFY          Product = 1
	Product_PRODUCT_MOBILE              Product = 2
	Product_PRODUCT_PARTNER             Product = 3
	Product_PRODUCT_LIBSPOTIFY_EMBEDDED Product = 5
)

var Product_name = map[int32]string{
	0: "PRODUCT_CLIENT",
	1: "PRODUCT_LIBSPOTIFY",
	2: "PRODUCT_MOBILE",
	3: "PRODUCT_PARTNER",
	5: "PRODUCT_LIBSPOTIFY_EMBEDDED",
}
var Product_value = map[string]int32{
	"PRODUCT_CLIENT":              0,
	"PRODUCT_LIBSPOTIFY":          1,
	"PRODUCT_MOBILE":              2,
	"PRODUCT_PARTNER":             3,
	"PRODUCT_LIBSPOTIFY_EMBEDDED": 5,
}

func (x Product) Enum() *Product {
	p := new(Product)
	*p = x
	return p
}
func (x Product) String() string {
	return proto.EnumName(Product_name, int32(x))
}
func (x *Product) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Product_value, data, "Product")
	if err != nil {
		return err
	}
	*x = Product(value)
	return nil
}
func (Product) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type ProductFlags int32

const (
	ProductFlags_PRODUCT_FLAG_NONE      ProductFlags = 0
	ProductFlags_PRODUCT_FLAG_DEV_BUILD ProductFlags = 1
)

var ProductFlags_name = map[int32]string{
	0: "PRODUCT_FLAG_NONE",
	1: "PRODUCT_FLAG_DEV_BUILD",
}
var ProductFlags_value = map[string]int32{
	"PRODUCT_FLAG_NONE":      0,
	"PRODUCT_FLAG_DEV_BUILD": 1,
}

func (x ProductFlags) Enum() *ProductFlags {
	p := new(ProductFlags)
	*p = x
	return p
}
func (x ProductFlags) String() string {
	return proto.EnumName(ProductFlags_name, int32(x))
}
func (x *ProductFlags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProductFlags_value, data, "ProductFlags")
	if err != nil {
		return err
	}
	*x = ProductFlags(value)
	return nil
}
func (ProductFlags) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type Platform int32

const (
	Platform_PLATFORM_WIN32_X86      Platform = 0
	Platform_PLATFORM_OSX_X86        Platform = 1
	Platform_PLATFORM_LINUX_X86      Platform = 2
	Platform_PLATFORM_IPHONE_ARM     Platform = 3
	Platform_PLATFORM_S60_ARM        Platform = 4
	Platform_PLATFORM_OSX_PPC        Platform = 5
	Platform_PLATFORM_ANDROID_ARM    Platform = 6
	Platform_PLATFORM_WINDOWS_CE_ARM Platform = 7
	Platform_PLATFORM_LINUX_X86_64   Platform = 8
	Platform_PLATFORM_OSX_X86_64     Platform = 9
	Platform_PLATFORM_PALM_ARM       Platform = 10
	Platform_PLATFORM_LINUX_SH       Platform = 11
	Platform_PLATFORM_FREEBSD_X86    Platform = 12
	Platform_PLATFORM_FREEBSD_X86_64 Platform = 13
	Platform_PLATFORM_BLACKBERRY_ARM Platform = 14
	Platform_PLATFORM_SONOS          Platform = 15
	Platform_PLATFORM_LINUX_MIPS     Platform = 16
	Platform_PLATFORM_LINUX_ARM      Platform = 17
	Platform_PLATFORM_LOGITECH_ARM   Platform = 18
	Platform_PLATFORM_LINUX_BLACKFIN Platform = 19
	Platform_PLATFORM_WP7_ARM        Platform = 20
	Platform_PLATFORM_ONKYO_ARM      Platform = 21
	Platform_PLATFORM_QNXNTO_ARM     Platform = 22
	Platform_PLATFORM_BCO_ARM        Platform = 23
)

var Platform_name = map[int32]string{
	0:  "PLATFORM_WIN32_X86",
	1:  "PLATFORM_OSX_X86",
	2:  "PLATFORM_LINUX_X86",
	3:  "PLATFORM_IPHONE_ARM",
	4:  "PLATFORM_S60_ARM",
	5:  "PLATFORM_OSX_PPC",
	6:  "PLATFORM_ANDROID_ARM",
	7:  "PLATFORM_WINDOWS_CE_ARM",
	8:  "PLATFORM_LINUX_X86_64",
	9:  "PLATFORM_OSX_X86_64",
	10: "PLATFORM_PALM_ARM",
	11: "PLATFORM_LINUX_SH",
	12: "PLATFORM_FREEBSD_X86",
	13: "PLATFORM_FREEBSD_X86_64",
	14: "PLATFORM_BLACKBERRY_ARM",
	15: "PLATFORM_SONOS",
	16: "PLATFORM_LINUX_MIPS",
	17: "PLATFORM_LINUX_ARM",
	18: "PLATFORM_LOGITECH_ARM",
	19: "PLATFORM_LINUX_BLACKFIN",
	20: "PLATFORM_WP7_ARM",
	21: "PLATFORM_ONKYO_ARM",
	22: "PLATFORM_QNXNTO_ARM",
	23: "PLATFORM_BCO_ARM",
}
var Platform_value = map[string]int32{
	"PLATFORM_WIN32_X86":      0,
	"PLATFORM_OSX_X86":        1,
	"PLATFORM_LINUX_X86":      2,
	"PLATFORM_IPHONE_ARM":     3,
	"PLATFORM_S60_ARM":        4,
	"PLATFORM_OSX_PPC":        5,
	"PLATFORM_ANDROID_ARM":    6,
	"PLATFORM_WINDOWS_CE_ARM": 7,
	"PLATFORM_LINUX_X86_64":   8,
	"PLATFORM_OSX_X86_64":     9,
	"PLATFORM_PALM_ARM":       10,
	"PLATFORM_LINUX_SH":       11,
	"PLATFORM_FREEBSD_X86":    12,
	"PLATFORM_FREEBSD_X86_64": 13,
	"PLATFORM_BLACKBERRY_ARM": 14,
	"PLATFORM_SONOS":          15,
	"PLATFORM_LINUX_MIPS":     16,
	"PLATFORM_LINUX_ARM":      17,
	"PLATFORM_LOGITECH_ARM":   18,
	"PLATFORM_LINUX_BLACKFIN": 19,
	"PLATFORM_WP7_ARM":        20,
	"PLATFORM_ONKYO_ARM":      21,
	"PLATFORM_QNXNTO_ARM":     22,
	"PLATFORM_BCO_ARM":        23,
}

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}
func (x Platform) String() string {
	return proto.EnumName(Platform_name, int32(x))
}
func (x *Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Platform_value, data, "Platform")
	if err != nil {
		return err
	}
	*x = Platform(value)
	return nil
}
func (Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type Fingerprint int32

const (
	Fingerprint_FINGERPRINT_GRAIN       Fingerprint = 0
	Fingerprint_FINGERPRINT_HMAC_RIPEMD Fingerprint = 1
)

var Fingerprint_name = map[int32]string{
	0: "FINGERPRINT_GRAIN",
	1: "FINGERPRINT_HMAC_RIPEMD",
}
var Fingerprint_value = map[string]int32{
	"FINGERPRINT_GRAIN":       0,
	"FINGERPRINT_HMAC_RIPEMD": 1,
}

func (x Fingerprint) Enum() *Fingerprint {
	p := new(Fingerprint)
	*p = x
	return p
}
func (x Fingerprint) String() string {
	return proto.EnumName(Fingerprint_name, int32(x))
}
func (x *Fingerprint) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Fingerprint_value, data, "Fingerprint")
	if err != nil {
		return err
	}
	*x = Fingerprint(value)
	return nil
}

func (Fingerprint) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type Cryptosuite int32

const (
   Cryptosuite_CRYPTO_SUITE_SHANNON       Cryptosuite = 0
   Cryptosuite_CRYPTO_SUITE_RC4_SHA1_HMAC Cryptosuite = 1
)

var Cryptosuite_name = map[int32]string{
   0: "CRYPTO_SUITE_SHANNON",
   1: "CRYPTO_SUITE_RC4_SHA1_HMAC",
}

var Cryptosuite_value = map[string]int32{
   "CRYPTO_SUITE_SHANNON":       0,
   "CRYPTO_SUITE_RC4_SHA1_HMAC": 1,
}

func (x Cryptosuite) Enum() *Cryptosuite {
	p := new(Cryptosuite)
	*p = x
	return p
}

func (x Cryptosuite) String() string {
	return proto.EnumName(Cryptosuite_name, int32(x))
}
func (x *Cryptosuite) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Cryptosuite_value, data, "Cryptosuite")
	if err != nil {
		return err
	}
	*x = Cryptosuite(value)
	return nil
}
func (Cryptosuite) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type Powscheme int32

const (
	Powscheme_POW_HASH_CASH Powscheme = 0
)

var Powscheme_name = map[int32]string{
	0: "POW_HASH_CASH",
}
var Powscheme_value = map[string]int32{
	"POW_HASH_CASH": 0,
}

func (x Powscheme) Enum() *Powscheme {
	p := new(Powscheme)
	*p = x
	return p
}
func (x Powscheme) String() string {
	return proto.EnumName(Powscheme_name, int32(x))
}
func (x *Powscheme) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Powscheme_value, data, "Powscheme")
	if err != nil {
		return err
	}
	*x = Powscheme(value)
	return nil
}
func (Powscheme) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type ErrorCode int32

const (
	ErrorCode_ProtocolError               ErrorCode = 0
	ErrorCode_TryAnotherAP                ErrorCode = 2
	ErrorCode_BadConnectionId             ErrorCode = 5
	ErrorCode_TravelRestriction           ErrorCode = 9
	ErrorCode_PremiumAccountRequired      ErrorCode = 11
	ErrorCode_BadCredentials              ErrorCode = 12
	ErrorCode_CouldNotValidateCredentials ErrorCode = 13
	ErrorCode_AccountExists               ErrorCode = 14
	ErrorCode_ExtraVerificationRequired   ErrorCode = 15
	ErrorCode_InvalidAppKey               ErrorCode = 16
	ErrorCode_ApplicationBanned           ErrorCode = 17
)

var ErrorCode_name = map[int32]string{
	0:  "ProtocolError",
	2:  "TryAnotherAP",
	5:  "BadConnectionId",
	9:  "TravelRestriction",
	11: "PremiumAccountRequired",
	12: "BadCredentials",
	13: "CouldNotValidateCredentials",
	14: "AccountExists",
	15: "ExtraVerificationRequired",
	16: "InvalidAppKey",
	17: "ApplicationBanned",
}
var ErrorCode_value = map[string]int32{
	"ProtocolError":               0,
	"TryAnotherAP":                2,
	"BadConnectionId":             5,
	"TravelRestriction":           9,
	"PremiumAccountRequired":      11,
	"BadCredentials":              12,
	"CouldNotValidateCredentials": 13,
	"AccountExists":               14,
	"ExtraVerificationRequired":   15,
	"InvalidAppKey":               16,
	"ApplicationBanned":           17,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type ClientHello struct {
	BuildInfo             *BuildInfo             `protobuf:"bytes,10,req,name=build_info,json=buildInfo" json:"build_info,omitempty"`
	FingerprintsSupported []Fingerprint          `protobuf:"varint,20,rep,name=fingerprints_supported,json=fingerprintsSupported,enum=Spotify.Fingerprint" json:"fingerprints_supported,omitempty"`
	CryptosuitesSupported []Cryptosuite          `protobuf:"varint,30,rep,name=cryptosuites_supported,json=cryptosuitesSupported,enum=Spotify.Cryptosuite" json:"cryptosuites_supported,omitempty"`
	PowschemesSupported   []Powscheme            `protobuf:"varint,40,rep,name=powschemes_supported,json=powschemesSupported,enum=Spotify.Powscheme" json:"powschemes_supported,omitempty"`
	LoginCryptoHello      *LoginCryptoHelloUnion `protobuf:"bytes,50,req,name=login_crypto_hello,json=loginCryptoHello" json:"login_crypto_hello,omitempty"`
	ClientNonce           []byte                 `protobuf:"bytes,60,req,name=client_nonce,json=clientNonce" json:"client_nonce,omitempty"`
	Padding               []byte                 `protobuf:"bytes,70,opt,name=padding" json:"padding,omitempty"`
	FeatureSet            *FeatureSet            `protobuf:"bytes,80,opt,name=feature_set,json=featureSet" json:"feature_set,omitempty"`
	XXX_unrecognized      []byte                 `json:"-"`
}

func (m *ClientHello) Reset()                    { *m = ClientHello{} }
func (m *ClientHello) String() string            { return proto.CompactTextString(m) }
func (*ClientHello) ProtoMessage()               {}
func (*ClientHello) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ClientHello) GetBuildInfo() *BuildInfo {
	if m != nil {
		return m.BuildInfo
	}
	return nil
}

func (m *ClientHello) GetFingerprintsSupported() []Fingerprint {
	if m != nil {
		return m.FingerprintsSupported
	}
	return nil
}

func (m *ClientHello) GetCryptosuitesSupported() []Cryptosuite {
	if m != nil {
		return m.CryptosuitesSupported
	}
	return nil
}

func (m *ClientHello) GetPowschemesSupported() []Powscheme {
	if m != nil {
		return m.PowschemesSupported
	}
	return nil
}

func (m *ClientHello) GetLoginCryptoHello() *LoginCryptoHelloUnion {
	if m != nil {
		return m.LoginCryptoHello
	}
	return nil
}

func (m *ClientHello) GetClientNonce() []byte {
	if m != nil {
		return m.ClientNonce
	}
	return nil
}

func (m *ClientHello) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

func (m *ClientHello) GetFeatureSet() *FeatureSet {
	if m != nil {
		return m.FeatureSet
	}
	return nil
}

type BuildInfo struct {
	Product          *Product       `protobuf:"varint,10,req,name=product,enum=Spotify.Product" json:"product,omitempty"`
	ProductFlags     []ProductFlags `protobuf:"varint,20,rep,name=product_flags,json=productFlags,enum=Spotify.ProductFlags" json:"product_flags,omitempty"`
	Platform         *Platform      `protobuf:"varint,30,req,name=platform,enum=Spotify.Platform" json:"platform,omitempty"`
	Version          *uint64        `protobuf:"varint,40,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *BuildInfo) Reset()                    { *m = BuildInfo{} }
func (m *BuildInfo) String() string            { return proto.CompactTextString(m) }
func (*BuildInfo) ProtoMessage()               {}
func (*BuildInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *BuildInfo) GetProduct() Product {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return Product_PRODUCT_CLIENT
}

func (m *BuildInfo) GetProductFlags() []ProductFlags {
	if m != nil {
		return m.ProductFlags
	}
	return nil
}

func (m *BuildInfo) GetPlatform() Platform {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return Platform_PLATFORM_WIN32_X86
}

func (m *BuildInfo) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

type LoginCryptoHelloUnion struct {
	DiffieHellman    *LoginCryptoDiffieHellmanHello `protobuf:"bytes,10,opt,name=diffie_hellman,json=diffieHellman" json:"diffie_hellman,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *LoginCryptoHelloUnion) Reset()                    { *m = LoginCryptoHelloUnion{} }
func (m *LoginCryptoHelloUnion) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoHelloUnion) ProtoMessage()               {}
func (*LoginCryptoHelloUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *LoginCryptoHelloUnion) GetDiffieHellman() *LoginCryptoDiffieHellmanHello {
	if m != nil {
		return m.DiffieHellman
	}
	return nil
}

type LoginCryptoDiffieHellmanHello struct {
	Gc               []byte  `protobuf:"bytes,10,req,name=gc" json:"gc,omitempty"`
	ServerKeysKnown  *uint32 `protobuf:"varint,20,req,name=server_keys_known,json=serverKeysKnown" json:"server_keys_known,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginCryptoDiffieHellmanHello) Reset()                    { *m = LoginCryptoDiffieHellmanHello{} }
func (m *LoginCryptoDiffieHellmanHello) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoDiffieHellmanHello) ProtoMessage()               {}
func (*LoginCryptoDiffieHellmanHello) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *LoginCryptoDiffieHellmanHello) GetGc() []byte {
	if m != nil {
		return m.Gc
	}
	return nil
}

func (m *LoginCryptoDiffieHellmanHello) GetServerKeysKnown() uint32 {
	if m != nil && m.ServerKeysKnown != nil {
		return *m.ServerKeysKnown
	}
	return 0
}

type FeatureSet struct {
	Autoupdate2      *bool  `protobuf:"varint,1,opt,name=autoupdate2" json:"autoupdate2,omitempty"`
	CurrentLocation  *bool  `protobuf:"varint,2,opt,name=current_location,json=currentLocation" json:"current_location,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FeatureSet) Reset()                    { *m = FeatureSet{} }
func (m *FeatureSet) String() string            { return proto.CompactTextString(m) }
func (*FeatureSet) ProtoMessage()               {}
func (*FeatureSet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *FeatureSet) GetAutoupdate2() bool {
	if m != nil && m.Autoupdate2 != nil {
		return *m.Autoupdate2
	}
	return false
}

func (m *FeatureSet) GetCurrentLocation() bool {
	if m != nil && m.CurrentLocation != nil {
		return *m.CurrentLocation
	}
	return false
}

type APResponseMessage struct {
	Challenge        *APChallenge            `protobuf:"bytes,10,opt,name=challenge" json:"challenge,omitempty"`
	Upgrade          *UpgradeRequiredMessage `protobuf:"bytes,20,opt,name=upgrade" json:"upgrade,omitempty"`
	LoginFailed      *APLoginFailed          `protobuf:"bytes,30,opt,name=login_failed,json=loginFailed" json:"login_failed,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *APResponseMessage) Reset()                    { *m = APResponseMessage{} }
func (m *APResponseMessage) String() string            { return proto.CompactTextString(m) }
func (*APResponseMessage) ProtoMessage()               {}
func (*APResponseMessage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *APResponseMessage) GetChallenge() *APChallenge {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *APResponseMessage) GetUpgrade() *UpgradeRequiredMessage {
	if m != nil {
		return m.Upgrade
	}
	return nil
}

func (m *APResponseMessage) GetLoginFailed() *APLoginFailed {
	if m != nil {
		return m.LoginFailed
	}
	return nil
}

type APChallenge struct {
	LoginCryptoChallenge *LoginCryptoChallengeUnion `protobuf:"bytes,10,req,name=login_crypto_challenge,json=loginCryptoChallenge" json:"login_crypto_challenge,omitempty"`
	FingerprintChallenge *FingerprintChallengeUnion `protobuf:"bytes,20,req,name=fingerprint_challenge,json=fingerprintChallenge" json:"fingerprint_challenge,omitempty"`
	PowChallenge         *PoWChallengeUnion         `protobuf:"bytes,30,req,name=pow_challenge,json=powChallenge" json:"pow_challenge,omitempty"`
	CryptoChallenge      *CryptoChallengeUnion      `protobuf:"bytes,40,req,name=crypto_challenge,json=cryptoChallenge" json:"crypto_challenge,omitempty"`
	ServerNonce          []byte                     `protobuf:"bytes,50,req,name=server_nonce,json=serverNonce" json:"server_nonce,omitempty"`
	Padding              []byte                     `protobuf:"bytes,60,opt,name=padding" json:"padding,omitempty"`
	XXX_unrecognized     []byte                     `json:"-"`
}

func (m *APChallenge) Reset()                    { *m = APChallenge{} }
func (m *APChallenge) String() string            { return proto.CompactTextString(m) }
func (*APChallenge) ProtoMessage()               {}
func (*APChallenge) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *APChallenge) GetLoginCryptoChallenge() *LoginCryptoChallengeUnion {
	if m != nil {
		return m.LoginCryptoChallenge
	}
	return nil
}

func (m *APChallenge) GetFingerprintChallenge() *FingerprintChallengeUnion {
	if m != nil {
		return m.FingerprintChallenge
	}
	return nil
}

func (m *APChallenge) GetPowChallenge() *PoWChallengeUnion {
	if m != nil {
		return m.PowChallenge
	}
	return nil
}

func (m *APChallenge) GetCryptoChallenge() *CryptoChallengeUnion {
	if m != nil {
		return m.CryptoChallenge
	}
	return nil
}

func (m *APChallenge) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *APChallenge) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

type LoginCryptoChallengeUnion struct {
	DiffieHellman    *LoginCryptoDiffieHellmanChallenge `protobuf:"bytes,10,opt,name=diffie_hellman,json=diffieHellman" json:"diffie_hellman,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *LoginCryptoChallengeUnion) Reset()                    { *m = LoginCryptoChallengeUnion{} }
func (m *LoginCryptoChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoChallengeUnion) ProtoMessage()               {}
func (*LoginCryptoChallengeUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *LoginCryptoChallengeUnion) GetDiffieHellman() *LoginCryptoDiffieHellmanChallenge {
	if m != nil {
		return m.DiffieHellman
	}
	return nil
}

type LoginCryptoDiffieHellmanChallenge struct {
	Gs                 []byte `protobuf:"bytes,10,req,name=gs" json:"gs,omitempty"`
	ServerSignatureKey *int32 `protobuf:"varint,20,req,name=server_signature_key,json=serverSignatureKey" json:"server_signature_key,omitempty"`
	GsSignature        []byte `protobuf:"bytes,30,req,name=gs_signature,json=gsSignature" json:"gs_signature,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *LoginCryptoDiffieHellmanChallenge) Reset()         { *m = LoginCryptoDiffieHellmanChallenge{} }
func (m *LoginCryptoDiffieHellmanChallenge) String() string { return proto.CompactTextString(m) }
func (*LoginCryptoDiffieHellmanChallenge) ProtoMessage()    {}
func (*LoginCryptoDiffieHellmanChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{8}
}

func (m *LoginCryptoDiffieHellmanChallenge) GetGs() []byte {
	if m != nil {
		return m.Gs
	}
	return nil
}

func (m *LoginCryptoDiffieHellmanChallenge) GetServerSignatureKey() int32 {
	if m != nil && m.ServerSignatureKey != nil {
		return *m.ServerSignatureKey
	}
	return 0
}

func (m *LoginCryptoDiffieHellmanChallenge) GetGsSignature() []byte {
	if m != nil {
		return m.GsSignature
	}
	return nil
}

type FingerprintChallengeUnion struct {
	Grain            *FingerprintGrainChallenge      `protobuf:"bytes,10,opt,name=grain" json:"grain,omitempty"`
	HmacRipemd       *FingerprintHmacRipemdChallenge `protobuf:"bytes,20,opt,name=hmac_ripemd,json=hmacRipemd" json:"hmac_ripemd,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *FingerprintChallengeUnion) Reset()                    { *m = FingerprintChallengeUnion{} }
func (m *FingerprintChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*FingerprintChallengeUnion) ProtoMessage()               {}
func (*FingerprintChallengeUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *FingerprintChallengeUnion) GetGrain() *FingerprintGrainChallenge {
	if m != nil {
		return m.Grain
	}
	return nil
}

func (m *FingerprintChallengeUnion) GetHmacRipemd() *FingerprintHmacRipemdChallenge {
	if m != nil {
		return m.HmacRipemd
	}
	return nil
}

type FingerprintGrainChallenge struct {
	Kek              []byte `protobuf:"bytes,10,req,name=kek" json:"kek,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintGrainChallenge) Reset()                    { *m = FingerprintGrainChallenge{} }
func (m *FingerprintGrainChallenge) String() string            { return proto.CompactTextString(m) }
func (*FingerprintGrainChallenge) ProtoMessage()               {}
func (*FingerprintGrainChallenge) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *FingerprintGrainChallenge) GetKek() []byte {
	if m != nil {
		return m.Kek
	}
	return nil
}

type FingerprintHmacRipemdChallenge struct {
	Challenge        []byte `protobuf:"bytes,10,req,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintHmacRipemdChallenge) Reset()         { *m = FingerprintHmacRipemdChallenge{} }
func (m *FingerprintHmacRipemdChallenge) String() string { return proto.CompactTextString(m) }
func (*FingerprintHmacRipemdChallenge) ProtoMessage()    {}
func (*FingerprintHmacRipemdChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{11}
}

func (m *FingerprintHmacRipemdChallenge) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

type PoWChallengeUnion struct {
	HashCash         *PoWHashCashChallenge `protobuf:"bytes,10,opt,name=hash_cash,json=hashCash" json:"hash_cash,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PoWChallengeUnion) Reset()                    { *m = PoWChallengeUnion{} }
func (m *PoWChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*PoWChallengeUnion) ProtoMessage()               {}
func (*PoWChallengeUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *PoWChallengeUnion) GetHashCash() *PoWHashCashChallenge {
	if m != nil {
		return m.HashCash
	}
	return nil
}

type PoWHashCashChallenge struct {
	Prefix           []byte `protobuf:"bytes,10,opt,name=prefix" json:"prefix,omitempty"`
	Length           *int32 `protobuf:"varint,20,opt,name=length" json:"length,omitempty"`
	Target           *int32 `protobuf:"varint,30,opt,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PoWHashCashChallenge) Reset()                    { *m = PoWHashCashChallenge{} }
func (m *PoWHashCashChallenge) String() string            { return proto.CompactTextString(m) }
func (*PoWHashCashChallenge) ProtoMessage()               {}
func (*PoWHashCashChallenge) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *PoWHashCashChallenge) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *PoWHashCashChallenge) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PoWHashCashChallenge) GetTarget() int32 {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return 0
}

type CryptoChallengeUnion struct {
	Shannon          *CryptoShannonChallenge     `protobuf:"bytes,10,opt,name=shannon" json:"shannon,omitempty"`
	Rc4Sha1Hmac      *CryptoRc4Sha1HmacChallenge `protobuf:"bytes,20,opt,name=rc4_sha1_hmac,json=rc4Sha1Hmac" json:"rc4_sha1_hmac,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *CryptoChallengeUnion) Reset()                    { *m = CryptoChallengeUnion{} }
func (m *CryptoChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*CryptoChallengeUnion) ProtoMessage()               {}
func (*CryptoChallengeUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *CryptoChallengeUnion) GetShannon() *CryptoShannonChallenge {
	if m != nil {
		return m.Shannon
	}
	return nil
}

func (m *CryptoChallengeUnion) GetRc4Sha1Hmac() *CryptoRc4Sha1HmacChallenge {
	if m != nil {
		return m.Rc4Sha1Hmac
	}
	return nil
}

type CryptoShannonChallenge struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoShannonChallenge) Reset()                    { *m = CryptoShannonChallenge{} }
func (m *CryptoShannonChallenge) String() string            { return proto.CompactTextString(m) }
func (*CryptoShannonChallenge) ProtoMessage()               {}
func (*CryptoShannonChallenge) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

type CryptoRc4Sha1HmacChallenge struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoRc4Sha1HmacChallenge) Reset()                    { *m = CryptoRc4Sha1HmacChallenge{} }
func (m *CryptoRc4Sha1HmacChallenge) String() string            { return proto.CompactTextString(m) }
func (*CryptoRc4Sha1HmacChallenge) ProtoMessage()               {}
func (*CryptoRc4Sha1HmacChallenge) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

type UpgradeRequiredMessage struct {
	UpgradeSignedPart []byte  `protobuf:"bytes,10,req,name=upgrade_signed_part,json=upgradeSignedPart" json:"upgrade_signed_part,omitempty"`
	Signature         []byte  `protobuf:"bytes,20,req,name=signature" json:"signature,omitempty"`
	HttpSuffix        *string `protobuf:"bytes,30,opt,name=http_suffix,json=httpSuffix" json:"http_suffix,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *UpgradeRequiredMessage) Reset()                    { *m = UpgradeRequiredMessage{} }
func (m *UpgradeRequiredMessage) String() string            { return proto.CompactTextString(m) }
func (*UpgradeRequiredMessage) ProtoMessage()               {}
func (*UpgradeRequiredMessage) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

func (m *UpgradeRequiredMessage) GetUpgradeSignedPart() []byte {
	if m != nil {
		return m.UpgradeSignedPart
	}
	return nil
}

func (m *UpgradeRequiredMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *UpgradeRequiredMessage) GetHttpSuffix() string {
	if m != nil && m.HttpSuffix != nil {
		return *m.HttpSuffix
	}
	return ""
}

type APLoginFailed struct {
	ErrorCode        *ErrorCode `protobuf:"varint,10,req,name=error_code,json=errorCode,enum=Spotify.ErrorCode" json:"error_code,omitempty"`
	RetryDelay       *int32     `protobuf:"varint,20,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
	Expiry           *int32     `protobuf:"varint,30,opt,name=expiry" json:"expiry,omitempty"`
	ErrorDescription *string    `protobuf:"bytes,40,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *APLoginFailed) Reset()                    { *m = APLoginFailed{} }
func (m *APLoginFailed) String() string            { return proto.CompactTextString(m) }
func (*APLoginFailed) ProtoMessage()               {}
func (*APLoginFailed) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

func (m *APLoginFailed) GetErrorCode() ErrorCode {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ErrorCode_ProtocolError
}

func (m *APLoginFailed) GetRetryDelay() int32 {
	if m != nil && m.RetryDelay != nil {
		return *m.RetryDelay
	}
	return 0
}

func (m *APLoginFailed) GetExpiry() int32 {
	if m != nil && m.Expiry != nil {
		return *m.Expiry
	}
	return 0
}

func (m *APLoginFailed) GetErrorDescription() string {
	if m != nil && m.ErrorDescription != nil {
		return *m.ErrorDescription
	}
	return ""
}

type ClientResponsePlaintext struct {
	LoginCryptoResponse *LoginCryptoResponseUnion `protobuf:"bytes,10,req,name=login_crypto_response,json=loginCryptoResponse" json:"login_crypto_response,omitempty"`
	PowResponse         *PoWResponseUnion         `protobuf:"bytes,20,req,name=pow_response,json=powResponse" json:"pow_response,omitempty"`
	CryptoResponse      *CryptoResponseUnion      `protobuf:"bytes,30,req,name=crypto_response,json=cryptoResponse" json:"crypto_response,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *ClientResponsePlaintext) Reset()                    { *m = ClientResponsePlaintext{} }
func (m *ClientResponsePlaintext) String() string            { return proto.CompactTextString(m) }
func (*ClientResponsePlaintext) ProtoMessage()               {}
func (*ClientResponsePlaintext) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

func (m *ClientResponsePlaintext) GetLoginCryptoResponse() *LoginCryptoResponseUnion {
	if m != nil {
		return m.LoginCryptoResponse
	}
	return nil
}

func (m *ClientResponsePlaintext) GetPowResponse() *PoWResponseUnion {
	if m != nil {
		return m.PowResponse
	}
	return nil
}

func (m *ClientResponsePlaintext) GetCryptoResponse() *CryptoResponseUnion {
	if m != nil {
		return m.CryptoResponse
	}
	return nil
}

type LoginCryptoResponseUnion struct {
	DiffieHellman    *LoginCryptoDiffieHellmanResponse `protobuf:"bytes,10,opt,name=diffie_hellman,json=diffieHellman" json:"diffie_hellman,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *LoginCryptoResponseUnion) Reset()                    { *m = LoginCryptoResponseUnion{} }
func (m *LoginCryptoResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoResponseUnion) ProtoMessage()               {}
func (*LoginCryptoResponseUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{20} }

func (m *LoginCryptoResponseUnion) GetDiffieHellman() *LoginCryptoDiffieHellmanResponse {
	if m != nil {
		return m.DiffieHellman
	}
	return nil
}

type LoginCryptoDiffieHellmanResponse struct {
	Hmac             []byte `protobuf:"bytes,10,req,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LoginCryptoDiffieHellmanResponse) Reset()         { *m = LoginCryptoDiffieHellmanResponse{} }
func (m *LoginCryptoDiffieHellmanResponse) String() string { return proto.CompactTextString(m) }
func (*LoginCryptoDiffieHellmanResponse) ProtoMessage()    {}
func (*LoginCryptoDiffieHellmanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{21}
}

func (m *LoginCryptoDiffieHellmanResponse) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type PoWResponseUnion struct {
	HashCash         *PoWHashCashResponse `protobuf:"bytes,10,opt,name=hash_cash,json=hashCash" json:"hash_cash,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PoWResponseUnion) Reset()                    { *m = PoWResponseUnion{} }
func (m *PoWResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*PoWResponseUnion) ProtoMessage()               {}
func (*PoWResponseUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{22} }

func (m *PoWResponseUnion) GetHashCash() *PoWHashCashResponse {
	if m != nil {
		return m.HashCash
	}
	return nil
}

type PoWHashCashResponse struct {
	HashSuffix       []byte `protobuf:"bytes,10,req,name=hash_suffix,json=hashSuffix" json:"hash_suffix,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PoWHashCashResponse) Reset()                    { *m = PoWHashCashResponse{} }
func (m *PoWHashCashResponse) String() string            { return proto.CompactTextString(m) }
func (*PoWHashCashResponse) ProtoMessage()               {}
func (*PoWHashCashResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{23} }

func (m *PoWHashCashResponse) GetHashSuffix() []byte {
	if m != nil {
		return m.HashSuffix
	}
	return nil
}

type CryptoResponseUnion struct {
	Shannon          *CryptoShannonResponse     `protobuf:"bytes,10,opt,name=shannon" json:"shannon,omitempty"`
	Rc4Sha1Hmac      *CryptoRc4Sha1HmacResponse `protobuf:"bytes,20,opt,name=rc4_sha1_hmac,json=rc4Sha1Hmac" json:"rc4_sha1_hmac,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CryptoResponseUnion) Reset()                    { *m = CryptoResponseUnion{} }
func (m *CryptoResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*CryptoResponseUnion) ProtoMessage()               {}
func (*CryptoResponseUnion) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{24} }

func (m *CryptoResponseUnion) GetShannon() *CryptoShannonResponse {
	if m != nil {
		return m.Shannon
	}
	return nil
}

func (m *CryptoResponseUnion) GetRc4Sha1Hmac() *CryptoRc4Sha1HmacResponse {
	if m != nil {
		return m.Rc4Sha1Hmac
	}
	return nil
}

type CryptoShannonResponse struct {
	Dummy            *int32 `protobuf:"varint,1,opt,name=dummy" json:"dummy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoShannonResponse) Reset()                    { *m = CryptoShannonResponse{} }
func (m *CryptoShannonResponse) String() string            { return proto.CompactTextString(m) }
func (*CryptoShannonResponse) ProtoMessage()               {}
func (*CryptoShannonResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{25} }

func (m *CryptoShannonResponse) GetDummy() int32 {
	if m != nil && m.Dummy != nil {
		return *m.Dummy
	}
	return 0
}

type CryptoRc4Sha1HmacResponse struct {
	Dummy            *int32 `protobuf:"varint,1,opt,name=dummy" json:"dummy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoRc4Sha1HmacResponse) Reset()                    { *m = CryptoRc4Sha1HmacResponse{} }
func (m *CryptoRc4Sha1HmacResponse) String() string            { return proto.CompactTextString(m) }
func (*CryptoRc4Sha1HmacResponse) ProtoMessage()               {}
func (*CryptoRc4Sha1HmacResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{26} }

func (m *CryptoRc4Sha1HmacResponse) GetDummy() int32 {
	if m != nil && m.Dummy != nil {
		return *m.Dummy
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientHello)(nil), "Spotify.ClientHello")
	proto.RegisterType((*BuildInfo)(nil), "Spotify.BuildInfo")
	proto.RegisterType((*LoginCryptoHelloUnion)(nil), "Spotify.LoginCryptoHelloUnion")
	proto.RegisterType((*LoginCryptoDiffieHellmanHello)(nil), "Spotify.LoginCryptoDiffieHellmanHello")
	proto.RegisterType((*FeatureSet)(nil), "Spotify.FeatureSet")
	proto.RegisterType((*APResponseMessage)(nil), "Spotify.APResponseMessage")
	proto.RegisterType((*APChallenge)(nil), "Spotify.APChallenge")
	proto.RegisterType((*LoginCryptoChallengeUnion)(nil), "Spotify.LoginCryptoChallengeUnion")
	proto.RegisterType((*LoginCryptoDiffieHellmanChallenge)(nil), "Spotify.LoginCryptoDiffieHellmanChallenge")
	proto.RegisterType((*FingerprintChallengeUnion)(nil), "Spotify.FingerprintChallengeUnion")
	proto.RegisterType((*FingerprintGrainChallenge)(nil), "Spotify.FingerprintGrainChallenge")
	proto.RegisterType((*FingerprintHmacRipemdChallenge)(nil), "Spotify.FingerprintHmacRipemdChallenge")
	proto.RegisterType((*PoWChallengeUnion)(nil), "Spotify.PoWChallengeUnion")
	proto.RegisterType((*PoWHashCashChallenge)(nil), "Spotify.PoWHashCashChallenge")
	proto.RegisterType((*CryptoChallengeUnion)(nil), "Spotify.CryptoChallengeUnion")
	proto.RegisterType((*CryptoShannonChallenge)(nil), "Spotify.CryptoShannonChallenge")
	proto.RegisterType((*CryptoRc4Sha1HmacChallenge)(nil), "Spotify.CryptoRc4Sha1HmacChallenge")
	proto.RegisterType((*UpgradeRequiredMessage)(nil), "Spotify.UpgradeRequiredMessage")
	proto.RegisterType((*APLoginFailed)(nil), "Spotify.APLoginFailed")
	proto.RegisterType((*ClientResponsePlaintext)(nil), "Spotify.ClientResponsePlaintext")
	proto.RegisterType((*LoginCryptoResponseUnion)(nil), "Spotify.LoginCryptoResponseUnion")
	proto.RegisterType((*LoginCryptoDiffieHellmanResponse)(nil), "Spotify.LoginCryptoDiffieHellmanResponse")
	proto.RegisterType((*PoWResponseUnion)(nil), "Spotify.PoWResponseUnion")
	proto.RegisterType((*PoWHashCashResponse)(nil), "Spotify.PoWHashCashResponse")
	proto.RegisterType((*CryptoResponseUnion)(nil), "Spotify.CryptoResponseUnion")
	proto.RegisterType((*CryptoShannonResponse)(nil), "Spotify.CryptoShannonResponse")
	proto.RegisterType((*CryptoRc4Sha1HmacResponse)(nil), "Spotify.CryptoRc4Sha1HmacResponse")
	proto.RegisterEnum("Spotify.Product", Product_name, Product_value)
	proto.RegisterEnum("Spotify.ProductFlags", ProductFlags_name, ProductFlags_value)
	proto.RegisterEnum("Spotify.Platform", Platform_name, Platform_value)
	proto.RegisterEnum("Spotify.Fingerprint", Fingerprint_name, Fingerprint_value)
	proto.RegisterEnum("Spotify.Cryptosuite", Cryptosuite_name, Cryptosuite_value)
	proto.RegisterEnum("Spotify.Powscheme", Powscheme_name, Powscheme_value)
	proto.RegisterEnum("Spotify.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("keyexchange.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x58, 0xdd, 0x6e, 0xdb, 0xc8,
	0x15, 0xb6, 0xec, 0x78, 0x1d, 0x1f, 0xc9, 0x36, 0x35, 0x96, 0x1d, 0x39, 0x9b, 0x38, 0x0e, 0x0b,
	0xb4, 0x5a, 0x17, 0x1b, 0x6c, 0xdc, 0x20, 0x4d, 0x16, 0x41, 0x0b, 0x4a, 0xa2, 0x2c, 0xc1, 0x12,
	0xc5, 0x1d, 0xca, 0x71, 0x82, 0x02, 0x25, 0xb8, 0xe4, 0x48, 0x22, 0x4c, 0x93, 0xec, 0x90, 0x4a,
	0xac, 0xab, 0x5e, 0x76, 0xfb, 0x04, 0x05, 0x7a, 0xd3, 0xdb, 0x3e, 0x44, 0xfb, 0x26, 0xbd, 0xea,
	0x3b, 0xf4, 0xbe, 0x98, 0xe1, 0xbf, 0x4c, 0x6f, 0x2e, 0x0c, 0x70, 0xbe, 0x73, 0xce, 0x37, 0x67,
	0xce, 0xdf, 0x8c, 0x0c, 0xf5, 0x6b, 0xb2, 0x24, 0xb7, 0xe6, 0xdc, 0x70, 0x67, 0xe4, 0x85, 0x4f,
	0xbd, 0xd0, 0x43, 0x5b, 0x9a, 0xef, 0x85, 0xf6, 0x74, 0x29, 0xfe, 0x77, 0x03, 0xaa, 0x1d, 0xc7,
	0x26, 0x6e, 0xd8, 0x27, 0x8e, 0xe3, 0xa1, 0x97, 0x00, 0x3f, 0x2e, 0x6c, 0xc7, 0xd2, 0x6d, 0x77,
	0xea, 0x35, 0xe1, 0x64, 0xbd, 0x55, 0x3d, 0x43, 0x2f, 0x62, 0xed, 0x17, 0x6d, 0x26, 0x1a, 0xb8,
	0x53, 0x0f, 0x6f, 0xff, 0x98, 0x7c, 0xa2, 0x0b, 0x38, 0x9c, 0xda, 0xee, 0x8c, 0x50, 0x9f, 0xda,
	0x6e, 0x18, 0xe8, 0xc1, 0xc2, 0xf7, 0x3d, 0x1a, 0x12, 0xab, 0xd9, 0x38, 0xd9, 0x68, 0xed, 0x9e,
	0x35, 0x52, 0xf3, 0x5e, 0xa6, 0x86, 0x0f, 0xf2, 0x36, 0x5a, 0x62, 0xc2, 0xc8, 0x4c, 0xba, 0xf4,
	0x43, 0x2f, 0x58, 0xd8, 0x21, 0xc9, 0x93, 0x1d, 0xaf, 0x90, 0x75, 0x32, 0x35, 0x7c, 0x90, 0xb7,
	0xc9, 0xc8, 0x64, 0x68, 0xf8, 0xde, 0xe7, 0xc0, 0x9c, 0x93, 0x9b, 0x02, 0x55, 0x8b, 0x53, 0x65,
	0xc7, 0x52, 0x13, 0x25, 0xbc, 0x9f, 0xe9, 0x67, 0x34, 0x43, 0x40, 0x8e, 0x37, 0xb3, 0x5d, 0x3d,
	0xda, 0x45, 0x9f, 0xb3, 0x48, 0x35, 0xcf, 0x78, 0x6c, 0x8e, 0x53, 0x92, 0x21, 0x53, 0x89, 0x9c,
	0xe2, 0xa1, 0xbc, 0x74, 0x6d, 0xcf, 0xc5, 0x82, 0xb3, 0x02, 0xa3, 0xe7, 0x50, 0x33, 0x79, 0xc0,
	0x75, 0xd7, 0x73, 0x4d, 0xd2, 0x7c, 0x77, 0xb2, 0xde, 0xaa, 0xe1, 0x6a, 0x84, 0x29, 0x0c, 0x42,
	0x4d, 0xd8, 0xf2, 0x0d, 0xcb, 0xb2, 0xdd, 0x59, 0xb3, 0x77, 0x52, 0x69, 0xd5, 0x70, 0xb2, 0x44,
	0xaf, 0xa0, 0x3a, 0x25, 0x46, 0xb8, 0xa0, 0x44, 0x0f, 0x48, 0xd8, 0x54, 0x4f, 0x2a, 0xad, 0xea,
	0xd9, 0x7e, 0x16, 0xe0, 0x48, 0xa6, 0x91, 0x10, 0xc3, 0x34, 0xfd, 0x16, 0xff, 0x55, 0x81, 0xed,
	0x34, 0x75, 0xe8, 0x14, 0xb6, 0x7c, 0xea, 0x59, 0x0b, 0x33, 0xe4, 0xf9, 0xdd, 0x3d, 0x13, 0xb2,
	0x40, 0x44, 0x38, 0x4e, 0x14, 0xd0, 0xf7, 0xb0, 0x13, 0x7f, 0xea, 0x53, 0xc7, 0x98, 0x05, 0x71,
	0x4a, 0x0f, 0x56, 0x2d, 0x7a, 0x4c, 0x88, 0x6b, 0x7e, 0x6e, 0x85, 0xbe, 0x85, 0x87, 0xbe, 0x63,
	0x84, 0x53, 0x8f, 0xde, 0x34, 0x8f, 0xf9, 0x46, 0xf5, 0xcc, 0x2c, 0x16, 0xe0, 0x54, 0x85, 0x1d,
	0xfa, 0x13, 0xa1, 0x81, 0xed, 0xb9, 0xcd, 0xd6, 0xc9, 0x7a, 0xeb, 0x01, 0x4e, 0x96, 0xe2, 0x14,
	0x0e, 0x4a, 0x83, 0x8b, 0x46, 0xb0, 0x6b, 0xd9, 0xd3, 0xa9, 0x4d, 0x78, 0x4a, 0x6e, 0x0c, 0xb7,
	0x09, 0x3c, 0x20, 0xbf, 0x2c, 0x4b, 0x4a, 0x97, 0x6b, 0xf6, 0x23, 0x45, 0x4e, 0x82, 0x77, 0xac,
	0x3c, 0x26, 0xfe, 0x01, 0x9e, 0xfe, 0xac, 0x3e, 0xda, 0x85, 0xf5, 0x99, 0xc9, 0x83, 0x56, 0xc3,
	0xeb, 0x33, 0x13, 0x9d, 0x42, 0x3d, 0x20, 0xf4, 0x13, 0xa1, 0xfa, 0x35, 0x59, 0x06, 0xfa, 0xb5,
	0xeb, 0x7d, 0x76, 0x9b, 0x8d, 0x93, 0xf5, 0xd6, 0x0e, 0xde, 0x8b, 0x04, 0x17, 0x64, 0x19, 0x5c,
	0x30, 0x58, 0xfc, 0x08, 0x90, 0x65, 0x07, 0x9d, 0x40, 0xd5, 0x58, 0x84, 0xde, 0xc2, 0xb7, 0x8c,
	0x90, 0x9c, 0x35, 0x2b, 0x27, 0x95, 0xd6, 0x43, 0x9c, 0x87, 0xd0, 0x37, 0x20, 0x98, 0x0b, 0x4a,
	0x59, 0x9d, 0x38, 0x9e, 0x69, 0x84, 0x2c, 0x2e, 0xeb, 0x5c, 0x6d, 0x2f, 0xc6, 0x87, 0x31, 0x2c,
	0xfe, 0xbb, 0x02, 0x75, 0x49, 0xc5, 0x24, 0xf0, 0x3d, 0x37, 0x20, 0x23, 0x12, 0x04, 0xc6, 0x8c,
	0xa0, 0x33, 0xd8, 0x36, 0xe7, 0x86, 0xe3, 0x10, 0x77, 0x46, 0xe2, 0xb8, 0x64, 0xcd, 0x23, 0xa9,
	0x9d, 0x44, 0x86, 0x33, 0x35, 0xf4, 0x16, 0xb6, 0x16, 0xfe, 0x8c, 0x1a, 0x16, 0x69, 0x36, 0xb8,
	0xc5, 0xb3, 0xd4, 0xe2, 0x32, 0xc2, 0x31, 0xf9, 0xd3, 0xc2, 0xa6, 0xc4, 0x8a, 0x77, 0xc1, 0x89,
	0x3e, 0x7a, 0x0b, 0xb5, 0xa8, 0x49, 0xa6, 0x86, 0xed, 0xf0, 0x76, 0x65, 0xf6, 0x87, 0xb9, 0x1d,
	0x79, 0x6c, 0x7b, 0x5c, 0x8a, 0xab, 0x4e, 0xb6, 0x10, 0x7f, 0xda, 0x80, 0x6a, 0xce, 0x21, 0xf4,
	0x01, 0x0e, 0x0b, 0xfd, 0x96, 0x3f, 0x06, 0xeb, 0x39, 0xb1, 0x2c, 0xbd, 0xa9, 0x79, 0xd4, 0x77,
	0x0d, 0xa7, 0x44, 0x84, 0xae, 0x20, 0x3f, 0x76, 0x72, 0xc4, 0x8d, 0x15, 0xe2, 0xdc, 0xa4, 0x5a,
	0x25, 0x9e, 0x96, 0x88, 0xd0, 0xef, 0x61, 0xc7, 0xf7, 0x3e, 0xe7, 0x08, 0x8f, 0x39, 0xe1, 0xe3,
	0xdc, 0x88, 0xb9, 0x5a, 0x21, 0xaa, 0xf9, 0xde, 0xe7, 0x8c, 0xa0, 0x0f, 0xc2, 0x9d, 0xd3, 0xb6,
	0x38, 0xc7, 0xd3, 0x95, 0x89, 0xb7, 0x42, 0xb3, 0x67, 0xae, 0x9c, 0xf1, 0x39, 0xd4, 0xe2, 0xa2,
	0x8c, 0xe6, 0xcb, 0x59, 0x34, 0x5f, 0x22, 0xec, 0xce, 0x7c, 0x79, 0x57, 0x98, 0x2f, 0xa2, 0x0b,
	0x47, 0xf7, 0xc6, 0x14, 0xfd, 0x70, 0x4f, 0xbb, 0x9d, 0x7e, 0xb1, 0xdd, 0xb2, 0x62, 0x5b, 0x69,
	0xb9, 0x9f, 0x2a, 0xf0, 0xfc, 0x8b, 0x46, 0xbc, 0xef, 0x82, 0xb4, 0xef, 0x02, 0xf4, 0x1d, 0x34,
	0xe2, 0x23, 0x06, 0xf6, 0xcc, 0x8d, 0xc6, 0xe1, 0x35, 0x59, 0xf2, 0x2c, 0x6e, 0x62, 0x14, 0xc9,
	0xb4, 0x44, 0x74, 0x41, 0x96, 0x2c, 0x28, 0xb3, 0x20, 0xd3, 0xe6, 0xe9, 0xa9, 0xe1, 0xea, 0x2c,
	0x48, 0xb5, 0xc4, 0x7f, 0x54, 0xe0, 0xe8, 0xde, 0xb4, 0xa3, 0x37, 0xb0, 0x39, 0xa3, 0x86, 0x9d,
	0x1c, 0xb9, 0xb4, 0x52, 0xce, 0x99, 0x42, 0x76, 0xd4, 0xc8, 0x00, 0xf5, 0xa1, 0x3a, 0xbf, 0x31,
	0x4c, 0x9d, 0xda, 0x3e, 0xb9, 0xb1, 0xe2, 0xbe, 0xfa, 0x55, 0x99, 0x7d, 0xff, 0xc6, 0x30, 0x31,
	0xd7, 0xca, 0x48, 0x60, 0x9e, 0x82, 0xe2, 0xb7, 0x05, 0x07, 0x8b, 0xbb, 0x21, 0x01, 0x36, 0xae,
	0xc9, 0x75, 0x1c, 0x24, 0xf6, 0x29, 0xfe, 0x0e, 0x8e, 0x7f, 0x9e, 0x1c, 0x3d, 0x29, 0x8e, 0x08,
	0x66, 0x99, 0x01, 0xe2, 0x18, 0xea, 0x77, 0xaa, 0x16, 0x7d, 0x0f, 0xdb, 0x73, 0x23, 0x98, 0xeb,
	0xa6, 0x11, 0xcc, 0xe3, 0x58, 0x3c, 0xcd, 0x17, 0x79, 0xdf, 0x08, 0xe6, 0x1d, 0xf6, 0x97, 0x9e,
	0xe0, 0xe1, 0x3c, 0x86, 0xc4, 0x3f, 0x42, 0xa3, 0x4c, 0x03, 0x1d, 0xc2, 0x57, 0x3e, 0x25, 0x53,
	0xfb, 0x96, 0x13, 0xd6, 0x70, 0xbc, 0x62, 0x38, 0x53, 0x08, 0xe7, 0x3c, 0x68, 0x9b, 0x38, 0x5e,
	0x31, 0x3c, 0x34, 0xe8, 0x8c, 0x84, 0x7c, 0xc8, 0x6c, 0xe2, 0x78, 0x25, 0xfe, 0xbd, 0x02, 0x8d,
	0xd2, 0xc2, 0x7d, 0x0b, 0x5b, 0xc1, 0xdc, 0x70, 0x5d, 0x2f, 0x49, 0xdf, 0xb3, 0x95, 0x9e, 0xd2,
	0x22, 0x69, 0xe6, 0x74, 0xa2, 0x8f, 0xce, 0x61, 0x87, 0x9a, 0xaf, 0xf4, 0x60, 0x6e, 0xbc, 0xd4,
	0x59, 0x2a, 0xe2, 0xfc, 0xfd, 0x62, 0x85, 0x00, 0x9b, 0xaf, 0xb4, 0xb9, 0xf1, 0x92, 0x05, 0x39,
	0x23, 0xa9, 0xd2, 0x0c, 0x15, 0x9b, 0x70, 0x58, 0xbe, 0x97, 0xf8, 0x04, 0x1e, 0xdf, 0x4f, 0x22,
	0xfe, 0xa5, 0x02, 0x87, 0xe5, 0xb3, 0x17, 0xbd, 0x80, 0xfd, 0x78, 0xfa, 0xf2, 0xca, 0x26, 0x96,
	0xee, 0x1b, 0x34, 0x8c, 0x13, 0x59, 0x8f, 0x45, 0x1a, 0x97, 0xa8, 0x06, 0x0d, 0x59, 0xba, 0xb3,
	0x0e, 0x68, 0x44, 0xe9, 0x4e, 0x01, 0xf4, 0x0c, 0xaa, 0xf3, 0x30, 0xf4, 0xf5, 0x60, 0x31, 0x65,
	0xa9, 0x60, 0xa1, 0xdd, 0xc6, 0xc0, 0x20, 0x8d, 0x23, 0xe2, 0x3f, 0x2b, 0xb0, 0x53, 0x98, 0xe2,
	0xec, 0xb1, 0x48, 0x28, 0xf5, 0xa8, 0x6e, 0x7a, 0x16, 0x89, 0x1f, 0x13, 0xd9, 0xab, 0x4a, 0x66,
	0xa2, 0x8e, 0x67, 0x11, 0xbc, 0x4d, 0x92, 0x4f, 0xb6, 0x0b, 0x25, 0x21, 0x5d, 0xea, 0x16, 0x71,
	0x8c, 0x65, 0x9c, 0x58, 0xe0, 0x50, 0x97, 0x21, 0x2c, 0xb9, 0xe4, 0xd6, 0xb7, 0xe9, 0x32, 0x49,
	0x6e, 0xb4, 0x42, 0xbf, 0x86, 0x7a, 0xb4, 0x97, 0x45, 0x02, 0x93, 0xda, 0x7e, 0x18, 0x3d, 0x14,
	0x98, 0x93, 0x02, 0x17, 0x74, 0x33, 0x5c, 0xfc, 0x5f, 0x05, 0x1e, 0x45, 0xaf, 0xda, 0xe4, 0x56,
	0x54, 0x1d, 0xc3, 0x76, 0x43, 0x72, 0x1b, 0xa2, 0x4b, 0x38, 0x28, 0xdc, 0x2e, 0x34, 0xd6, 0x88,
	0x2f, 0x97, 0xe7, 0x65, 0xc3, 0x2c, 0x61, 0x89, 0x46, 0xee, 0xbe, 0x73, 0x57, 0x82, 0xde, 0x01,
	0x1b, 0xe8, 0x19, 0x5b, 0x74, 0xa3, 0x1c, 0xe5, 0x7b, 0xa3, 0xc8, 0x52, 0xf5, 0xbd, 0xcf, 0xa9,
	0xb5, 0x0c, 0x7b, 0xab, 0xee, 0x44, 0x37, 0xc8, 0x93, 0xd5, 0x42, 0x2b, 0x70, 0xec, 0x9a, 0x05,
	0x50, 0x74, 0xa0, 0x79, 0x9f, 0xd7, 0x48, 0xbd, 0x67, 0x7a, 0x7f, 0xf3, 0xc5, 0xe9, 0x9d, 0xf0,
	0xac, 0x0e, 0xef, 0xd7, 0x70, 0xf2, 0x25, 0x13, 0x84, 0xe0, 0x01, 0x6f, 0x9b, 0xa8, 0x28, 0xf9,
	0xb7, 0x38, 0x02, 0x61, 0x35, 0x1a, 0xe8, 0xed, 0xdd, 0xb9, 0xf2, 0xa4, 0x6c, 0xae, 0xa4, 0xbe,
	0x64, 0x63, 0xe5, 0x35, 0xec, 0x97, 0x28, 0xf0, 0x7a, 0x66, 0x8c, 0x71, 0x3d, 0x47, 0x0e, 0x00,
	0x83, 0xe2, 0x7a, 0xfe, 0x5b, 0x05, 0xf6, 0xcb, 0x02, 0xf5, 0x66, 0x75, 0x5a, 0x1c, 0x97, 0x4f,
	0x8b, 0xd4, 0x95, 0x74, 0x58, 0xf4, 0xca, 0x87, 0x85, 0x78, 0xff, 0xb0, 0x48, 0x39, 0x0a, 0xb3,
	0xe2, 0x5b, 0x38, 0x28, 0xdd, 0x09, 0x35, 0x60, 0xd3, 0x5a, 0xdc, 0xdc, 0x2c, 0xf9, 0x83, 0x71,
	0x13, 0x47, 0x0b, 0xf1, 0x25, 0x1c, 0xdd, 0x4b, 0x5c, 0x6e, 0x72, 0xfa, 0x67, 0xd8, 0x8a, 0x5f,
	0xee, 0x08, 0xc1, 0xae, 0x8a, 0xc7, 0xdd, 0xcb, 0xce, 0x44, 0xef, 0x0c, 0x07, 0xb2, 0x32, 0x11,
	0xd6, 0xd0, 0x21, 0xa0, 0x04, 0x1b, 0x0e, 0xda, 0x9a, 0x3a, 0x9e, 0x0c, 0x7a, 0x1f, 0x85, 0x4a,
	0x5e, 0x77, 0x34, 0x6e, 0x0f, 0x86, 0xb2, 0xb0, 0x8e, 0xf6, 0x61, 0x2f, 0xc1, 0x54, 0x09, 0x4f,
	0x14, 0x19, 0x0b, 0x1b, 0xe8, 0x19, 0x7c, 0x7d, 0x97, 0x40, 0x97, 0x47, 0x6d, 0xb9, 0xdb, 0x95,
	0xbb, 0xc2, 0xe6, 0xa9, 0x04, 0xb5, 0xfc, 0x4f, 0x07, 0x74, 0x00, 0xf5, 0xc4, 0xa0, 0x37, 0x94,
	0xce, 0x75, 0x65, 0xac, 0xc8, 0xc2, 0x1a, 0x7a, 0x0c, 0x87, 0x05, 0xb8, 0x2b, 0xbf, 0xd7, 0xdb,
	0x97, 0x83, 0x61, 0x57, 0xa8, 0x9c, 0xfe, 0xe7, 0x01, 0x3c, 0x4c, 0x7e, 0x47, 0x70, 0x8f, 0x87,
	0xd2, 0xa4, 0x37, 0xc6, 0x23, 0xfd, 0x6a, 0xa0, 0xfc, 0xe6, 0x4c, 0xff, 0xf0, 0xe6, 0xb5, 0xb0,
	0x86, 0x1a, 0x20, 0xa4, 0xf8, 0x58, 0xfb, 0xc0, 0xd1, 0x4a, 0x41, 0x7b, 0x38, 0x50, 0x2e, 0x23,
	0x7c, 0x1d, 0x3d, 0x82, 0xfd, 0x14, 0x1f, 0xa8, 0xfd, 0xb1, 0x22, 0xeb, 0x12, 0x1e, 0x09, 0x1b,
	0x05, 0x1a, 0xed, 0xf5, 0x77, 0x1c, 0x7d, 0x70, 0x87, 0x5c, 0x55, 0x3b, 0xc2, 0x26, 0x6a, 0x42,
	0x23, 0x45, 0x25, 0xa5, 0x8b, 0xc7, 0x83, 0x2e, 0xd7, 0xff, 0x0a, 0x7d, 0x0d, 0x8f, 0xf2, 0x4e,
	0x76, 0xc7, 0x57, 0x9a, 0xde, 0x89, 0xb6, 0xd8, 0x42, 0x47, 0x70, 0x70, 0xd7, 0x27, 0xfd, 0xf5,
	0x2b, 0xe1, 0x61, 0xc1, 0xad, 0xf8, 0x10, 0x4c, 0xb0, 0xcd, 0xa3, 0x96, 0x08, 0x54, 0x69, 0x38,
	0xe2, 0x54, 0x50, 0x80, 0x23, 0x2a, 0xad, 0x2f, 0x54, 0x0b, 0x8e, 0xf5, 0xb0, 0x2c, 0xb7, 0xb5,
	0x2e, 0x3f, 0x77, 0xad, 0xe0, 0x58, 0x4e, 0xc2, 0x36, 0xd9, 0x29, 0x08, 0xdb, 0x43, 0xa9, 0x73,
	0xd1, 0x96, 0x31, 0xfe, 0xc8, 0xb7, 0xda, 0xe5, 0x15, 0x91, 0x06, 0x66, 0xac, 0x8c, 0x35, 0x61,
	0xaf, 0xe0, 0x6e, 0xb4, 0xfd, 0x68, 0xa0, 0x6a, 0x82, 0x50, 0x12, 0x76, 0x46, 0x52, 0x2f, 0x1e,
	0x7d, 0x7c, 0x3e, 0x98, 0xc8, 0x9d, 0x3e, 0x17, 0xa1, 0xc2, 0xe6, 0x91, 0x09, 0x77, 0xa1, 0x37,
	0x50, 0x84, 0xfd, 0x42, 0xfc, 0xaf, 0xd4, 0xdf, 0x72, 0x93, 0x46, 0x61, 0x97, 0xb1, 0x72, 0xf1,
	0x71, 0xcc, 0xf1, 0x83, 0x82, 0x5b, 0x3f, 0x28, 0x1f, 0x94, 0x49, 0x24, 0x38, 0x2c, 0xd0, 0xb4,
	0x3b, 0x11, 0xfa, 0xe8, 0x54, 0x82, 0x6a, 0xee, 0xf9, 0xc4, 0x62, 0xda, 0x1b, 0x28, 0xe7, 0x32,
	0x56, 0xf1, 0x40, 0x99, 0xe8, 0xe7, 0x58, 0x1a, 0x28, 0xc2, 0x1a, 0xf3, 0x2f, 0x0f, 0xf7, 0x47,
	0x52, 0x47, 0xc7, 0x03, 0x55, 0x1e, 0xb1, 0x0a, 0x3d, 0x87, 0x6a, 0xee, 0xbf, 0x14, 0x2c, 0xfe,
	0x1d, 0xfc, 0x51, 0x9d, 0x8c, 0x75, 0xed, 0x72, 0x30, 0x91, 0x75, 0xad, 0x2f, 0x29, 0xca, 0x98,
	0xb1, 0x1c, 0xc3, 0xe3, 0x82, 0x04, 0x77, 0x5e, 0x31, 0xe9, 0x4b, 0xce, 0x27, 0x54, 0x4e, 0x8f,
	0x61, 0x3b, 0xfd, 0x1f, 0x05, 0xaa, 0xc3, 0x8e, 0x3a, 0xbe, 0xd2, 0xfb, 0x92, 0xd6, 0xd7, 0x3b,
	0x92, 0xd6, 0x17, 0xd6, 0x4e, 0xff, 0xba, 0x0e, 0xdb, 0xe9, 0x75, 0xcb, 0x15, 0xa8, 0x17, 0x7a,
	0xa6, 0xe7, 0x70, 0x50, 0x58, 0x43, 0x02, 0xd4, 0x26, 0x74, 0x29, 0xb9, 0x5e, 0x38, 0x27, 0x54,
	0x52, 0xa3, 0xb6, 0x6d, 0x1b, 0x56, 0xc7, 0x73, 0x5d, 0x62, 0xb2, 0x3b, 0x73, 0x60, 0x09, 0x9b,
	0xec, 0x90, 0x13, 0x6a, 0x7c, 0x22, 0x0e, 0x26, 0x41, 0x48, 0x6d, 0x2e, 0x10, 0xb6, 0x79, 0x17,
	0x52, 0x72, 0x63, 0x2f, 0x6e, 0x24, 0xd3, 0xf4, 0x16, 0xec, 0x56, 0x8d, 0x5e, 0x22, 0x42, 0x95,
	0x15, 0x00, 0xe3, 0xa1, 0xc4, 0x22, 0x6e, 0x68, 0x1b, 0x4e, 0x20, 0xd4, 0x58, 0xf7, 0x77, 0xbc,
	0x85, 0x63, 0x29, 0x5e, 0xf8, 0xde, 0x70, 0x6c, 0xf6, 0x83, 0x36, 0xaf, 0xb0, 0xc3, 0x3c, 0x8c,
	0x99, 0xe4, 0x5b, 0x3b, 0x08, 0x03, 0x61, 0x17, 0x3d, 0x85, 0x23, 0xf9, 0x36, 0xa4, 0xc6, 0x7b,
	0x42, 0xed, 0xa9, 0x1d, 0xfd, 0xb2, 0x4d, 0xb7, 0xd9, 0x63, 0x16, 0x03, 0xf7, 0x13, 0x23, 0x93,
	0x7c, 0xff, 0x82, 0x2c, 0x05, 0x81, 0x39, 0x2b, 0xf9, 0xbe, 0x13, 0xeb, 0xb6, 0x0d, 0xd7, 0x25,
	0x96, 0x50, 0xff, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xb0, 0x64, 0xb7, 0xee, 0x12, 0x00,
	0x00,
}
