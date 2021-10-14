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

type LoginCryptoHelloUnion struct {
	DiffieHellman    *LoginCryptoDiffieHellmanHello `protobuf:"bytes,10,opt,name=diffie_hellman,json=diffieHellman" json:"diffie_hellman,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *LoginCryptoHelloUnion) Reset()                    { *m = LoginCryptoHelloUnion{} }
func (m *LoginCryptoHelloUnion) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoHelloUnion) ProtoMessage()               {}

type LoginCryptoDiffieHellmanHello struct {
	Gc               []byte  `protobuf:"bytes,10,req,name=gc" json:"gc,omitempty"`
	ServerKeysKnown  *uint32 `protobuf:"varint,20,req,name=server_keys_known,json=serverKeysKnown" json:"server_keys_known,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LoginCryptoDiffieHellmanHello) Reset()                    { *m = LoginCryptoDiffieHellmanHello{} }
func (m *LoginCryptoDiffieHellmanHello) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoDiffieHellmanHello) ProtoMessage()               {}

type FeatureSet struct {
	Autoupdate2      *bool  `protobuf:"varint,1,opt,name=autoupdate2" json:"autoupdate2,omitempty"`
	CurrentLocation  *bool  `protobuf:"varint,2,opt,name=current_location,json=currentLocation" json:"current_location,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FeatureSet) Reset()                    { *m = FeatureSet{} }
func (m *FeatureSet) String() string            { return proto.CompactTextString(m) }
func (*FeatureSet) ProtoMessage()               {}

type APResponseMessage struct {
	Challenge        *APChallenge            `protobuf:"bytes,10,opt,name=challenge" json:"challenge,omitempty"`
	Upgrade          *UpgradeRequiredMessage `protobuf:"bytes,20,opt,name=upgrade" json:"upgrade,omitempty"`
	LoginFailed      *APLoginFailed          `protobuf:"bytes,30,opt,name=login_failed,json=loginFailed" json:"login_failed,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *APResponseMessage) Reset()                    { *m = APResponseMessage{} }
func (m *APResponseMessage) String() string            { return proto.CompactTextString(m) }
func (*APResponseMessage) ProtoMessage()               {}

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

type LoginCryptoChallengeUnion struct {
	DiffieHellman    *LoginCryptoDiffieHellmanChallenge `protobuf:"bytes,10,opt,name=diffie_hellman,json=diffieHellman" json:"diffie_hellman,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *LoginCryptoChallengeUnion) Reset()                    { *m = LoginCryptoChallengeUnion{} }
func (m *LoginCryptoChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoChallengeUnion) ProtoMessage()               {}

type LoginCryptoDiffieHellmanChallenge struct {
	Gs                 []byte `protobuf:"bytes,10,req,name=gs" json:"gs,omitempty"`
	ServerSignatureKey *int32 `protobuf:"varint,20,req,name=server_signature_key,json=serverSignatureKey" json:"server_signature_key,omitempty"`
	GsSignature        []byte `protobuf:"bytes,30,req,name=gs_signature,json=gsSignature" json:"gs_signature,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *LoginCryptoDiffieHellmanChallenge) Reset()         { *m = LoginCryptoDiffieHellmanChallenge{} }
func (m *LoginCryptoDiffieHellmanChallenge) String() string { return proto.CompactTextString(m) }
func (*LoginCryptoDiffieHellmanChallenge) ProtoMessage()    {}

type FingerprintChallengeUnion struct {
	Grain            *FingerprintGrainChallenge      `protobuf:"bytes,10,opt,name=grain" json:"grain,omitempty"`
	HmacRipemd       *FingerprintHmacRipemdChallenge `protobuf:"bytes,20,opt,name=hmac_ripemd,json=hmacRipemd" json:"hmac_ripemd,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *FingerprintChallengeUnion) Reset()                    { *m = FingerprintChallengeUnion{} }
func (m *FingerprintChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*FingerprintChallengeUnion) ProtoMessage()               {}

type FingerprintGrainChallenge struct {
	Kek              []byte `protobuf:"bytes,10,req,name=kek" json:"kek,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintGrainChallenge) Reset()                    { *m = FingerprintGrainChallenge{} }
func (m *FingerprintGrainChallenge) String() string            { return proto.CompactTextString(m) }
func (*FingerprintGrainChallenge) ProtoMessage()               {}

type FingerprintHmacRipemdChallenge struct {
	Challenge        []byte `protobuf:"bytes,10,req,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintHmacRipemdChallenge) Reset()         { *m = FingerprintHmacRipemdChallenge{} }
func (m *FingerprintHmacRipemdChallenge) String() string { return proto.CompactTextString(m) }
func (*FingerprintHmacRipemdChallenge) ProtoMessage()    {}

type PoWChallengeUnion struct {
	HashCash         *PoWHashCashChallenge `protobuf:"bytes,10,opt,name=hash_cash,json=hashCash" json:"hash_cash,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PoWChallengeUnion) Reset()                    { *m = PoWChallengeUnion{} }
func (m *PoWChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*PoWChallengeUnion) ProtoMessage()               {}

type PoWHashCashChallenge struct {
	Prefix           []byte `protobuf:"bytes,10,opt,name=prefix" json:"prefix,omitempty"`
	Length           *int32 `protobuf:"varint,20,opt,name=length" json:"length,omitempty"`
	Target           *int32 `protobuf:"varint,30,opt,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PoWHashCashChallenge) Reset()                    { *m = PoWHashCashChallenge{} }
func (m *PoWHashCashChallenge) String() string            { return proto.CompactTextString(m) }
func (*PoWHashCashChallenge) ProtoMessage()               {}

type CryptoChallengeUnion struct {
	Shannon          *CryptoShannonChallenge     `protobuf:"bytes,10,opt,name=shannon" json:"shannon,omitempty"`
	Rc4Sha1Hmac      *CryptoRc4Sha1HmacChallenge `protobuf:"bytes,20,opt,name=rc4_sha1_hmac,json=rc4Sha1Hmac" json:"rc4_sha1_hmac,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *CryptoChallengeUnion) Reset()                    { *m = CryptoChallengeUnion{} }
func (m *CryptoChallengeUnion) String() string            { return proto.CompactTextString(m) }
func (*CryptoChallengeUnion) ProtoMessage()               {}

type CryptoShannonChallenge struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoShannonChallenge) Reset()                    { *m = CryptoShannonChallenge{} }
func (m *CryptoShannonChallenge) String() string            { return proto.CompactTextString(m) }
func (*CryptoShannonChallenge) ProtoMessage()               {}

type CryptoRc4Sha1HmacChallenge struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoRc4Sha1HmacChallenge) Reset()                    { *m = CryptoRc4Sha1HmacChallenge{} }
func (m *CryptoRc4Sha1HmacChallenge) String() string            { return proto.CompactTextString(m) }
func (*CryptoRc4Sha1HmacChallenge) ProtoMessage()               {}

type UpgradeRequiredMessage struct {
	UpgradeSignedPart []byte  `protobuf:"bytes,10,req,name=upgrade_signed_part,json=upgradeSignedPart" json:"upgrade_signed_part,omitempty"`
	Signature         []byte  `protobuf:"bytes,20,req,name=signature" json:"signature,omitempty"`
	HttpSuffix        *string `protobuf:"bytes,30,opt,name=http_suffix,json=httpSuffix" json:"http_suffix,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *UpgradeRequiredMessage) Reset()                    { *m = UpgradeRequiredMessage{} }
func (m *UpgradeRequiredMessage) String() string            { return proto.CompactTextString(m) }
func (*UpgradeRequiredMessage) ProtoMessage()               {}

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

type ClientResponsePlaintext struct {
	LoginCryptoResponse *LoginCryptoResponseUnion `protobuf:"bytes,10,req,name=login_crypto_response,json=loginCryptoResponse" json:"login_crypto_response,omitempty"`
	PowResponse         *PoWResponseUnion         `protobuf:"bytes,20,req,name=pow_response,json=powResponse" json:"pow_response,omitempty"`
	CryptoResponse      *CryptoResponseUnion      `protobuf:"bytes,30,req,name=crypto_response,json=cryptoResponse" json:"crypto_response,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *ClientResponsePlaintext) Reset()                    { *m = ClientResponsePlaintext{} }
func (m *ClientResponsePlaintext) String() string            { return proto.CompactTextString(m) }
func (*ClientResponsePlaintext) ProtoMessage()               {}

type LoginCryptoResponseUnion struct {
	DiffieHellman    *LoginCryptoDiffieHellmanResponse `protobuf:"bytes,10,opt,name=diffie_hellman,json=diffieHellman" json:"diffie_hellman,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *LoginCryptoResponseUnion) Reset()                    { *m = LoginCryptoResponseUnion{} }
func (m *LoginCryptoResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*LoginCryptoResponseUnion) ProtoMessage()               {}

type LoginCryptoDiffieHellmanResponse struct {
	Hmac             []byte `protobuf:"bytes,10,req,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LoginCryptoDiffieHellmanResponse) Reset()         { *m = LoginCryptoDiffieHellmanResponse{} }
func (m *LoginCryptoDiffieHellmanResponse) String() string { return proto.CompactTextString(m) }
func (*LoginCryptoDiffieHellmanResponse) ProtoMessage()    {}

type PoWResponseUnion struct {
	HashCash         *PoWHashCashResponse `protobuf:"bytes,10,opt,name=hash_cash,json=hashCash" json:"hash_cash,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PoWResponseUnion) Reset()                    { *m = PoWResponseUnion{} }
func (m *PoWResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*PoWResponseUnion) ProtoMessage()               {}

type PoWHashCashResponse struct {
	HashSuffix       []byte `protobuf:"bytes,10,req,name=hash_suffix,json=hashSuffix" json:"hash_suffix,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PoWHashCashResponse) Reset()                    { *m = PoWHashCashResponse{} }
func (m *PoWHashCashResponse) String() string            { return proto.CompactTextString(m) }
func (*PoWHashCashResponse) ProtoMessage()               {}

type CryptoResponseUnion struct {
	Shannon          *CryptoShannonResponse     `protobuf:"bytes,10,opt,name=shannon" json:"shannon,omitempty"`
	Rc4Sha1Hmac      *CryptoRc4Sha1HmacResponse `protobuf:"bytes,20,opt,name=rc4_sha1_hmac,json=rc4Sha1Hmac" json:"rc4_sha1_hmac,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CryptoResponseUnion) Reset()                    { *m = CryptoResponseUnion{} }
func (m *CryptoResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*CryptoResponseUnion) ProtoMessage()               {}

type CryptoShannonResponse struct {
	Dummy            *int32 `protobuf:"varint,1,opt,name=dummy" json:"dummy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoShannonResponse) Reset()                    { *m = CryptoShannonResponse{} }
func (m *CryptoShannonResponse) String() string            { return proto.CompactTextString(m) }
func (*CryptoShannonResponse) ProtoMessage()               {}

type CryptoRc4Sha1HmacResponse struct {
	Dummy            *int32 `protobuf:"varint,1,opt,name=dummy" json:"dummy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CryptoRc4Sha1HmacResponse) Reset()                    { *m = CryptoRc4Sha1HmacResponse{} }
func (m *CryptoRc4Sha1HmacResponse) String() string            { return proto.CompactTextString(m) }
func (*CryptoRc4Sha1HmacResponse) ProtoMessage()               {}


