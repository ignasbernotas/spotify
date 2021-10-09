package pb

import "github.com/golang/protobuf/proto"

type AuthenticationType int32

const (
	AuthenticationType_AUTHENTICATION_USER_PASS                   AuthenticationType = 0
	AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS  AuthenticationType = 1
	AuthenticationType_AUTHENTICATION_STORED_FACEBOOK_CREDENTIALS AuthenticationType = 2
	AuthenticationType_AUTHENTICATION_SPOTIFY_TOKEN               AuthenticationType = 3
	AuthenticationType_AUTHENTICATION_FACEBOOK_TOKEN              AuthenticationType = 4
	AuthenticationType_AUTHENTICATION_UNKNOWN                     AuthenticationType = 8
)

var AuthenticationType_name = map[int32]string{
	0: "AUTHENTICATION_USER_PASS",
	1: "AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS",
	2: "AUTHENTICATION_STORED_FACEBOOK_CREDENTIALS",
	3: "AUTHENTICATION_SPOTIFY_TOKEN",
	4: "AUTHENTICATION_FACEBOOK_TOKEN",
	8: "AUTHENTICATION_UNKNOWN",
}
var AuthenticationType_value = map[string]int32{
	"AUTHENTICATION_USER_PASS":                   0,
	"AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS":  1,
	"AUTHENTICATION_STORED_FACEBOOK_CREDENTIALS": 2,
	"AUTHENTICATION_SPOTIFY_TOKEN":               3,
	"AUTHENTICATION_FACEBOOK_TOKEN":              4,
	"AUTHENTICATION_UNKNOWN":                     8,
}

func (x AuthenticationType) Enum() *AuthenticationType {
	p := new(AuthenticationType)
	*p = x
	return p
}
func (x AuthenticationType) String() string {
	return proto.EnumName(AuthenticationType_name, int32(x))
}
func (x *AuthenticationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthenticationType_value, data, "AuthenticationType")
	if err != nil {
		return err
	}
	*x = AuthenticationType(value)
	return nil
}
func (AuthenticationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type AccountCreation int32

const (
	AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT AccountCreation = 1
	AccountCreation_ACCOUNT_CREATION_ALWAYS_CREATE AccountCreation = 3
)

var AccountCreation_name = map[int32]string{
	1: "ACCOUNT_CREATION_ALWAYS_PROMPT",
	3: "ACCOUNT_CREATION_ALWAYS_CREATE",
}
var AccountCreation_value = map[string]int32{
	"ACCOUNT_CREATION_ALWAYS_PROMPT": 1,
	"ACCOUNT_CREATION_ALWAYS_CREATE": 3,
}

func (x AccountCreation) Enum() *AccountCreation {
	p := new(AccountCreation)
	*p = x
	return p
}
func (x AccountCreation) String() string {
	return proto.EnumName(AccountCreation_name, int32(x))
}
func (x *AccountCreation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AccountCreation_value, data, "AccountCreation")
	if err != nil {
		return err
	}
	*x = AccountCreation(value)
	return nil
}
func (AccountCreation) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type CpuFamily int32

const (
	CpuFamily_CPU_UNKNOWN  CpuFamily = 0
	CpuFamily_CPU_X86      CpuFamily = 1
	CpuFamily_CPU_X86_64   CpuFamily = 2
	CpuFamily_CPU_PPC      CpuFamily = 3
	CpuFamily_CPU_PPC_64   CpuFamily = 4
	CpuFamily_CPU_ARM      CpuFamily = 5
	CpuFamily_CPU_IA64     CpuFamily = 6
	CpuFamily_CPU_SH       CpuFamily = 7
	CpuFamily_CPU_MIPS     CpuFamily = 8
	CpuFamily_CPU_BLACKFIN CpuFamily = 9
)

var CpuFamily_name = map[int32]string{
	0: "CPU_UNKNOWN",
	1: "CPU_X86",
	2: "CPU_X86_64",
	3: "CPU_PPC",
	4: "CPU_PPC_64",
	5: "CPU_ARM",
	6: "CPU_IA64",
	7: "CPU_SH",
	8: "CPU_MIPS",
	9: "CPU_BLACKFIN",
}
var CpuFamily_value = map[string]int32{
	"CPU_UNKNOWN":  0,
	"CPU_X86":      1,
	"CPU_X86_64":   2,
	"CPU_PPC":      3,
	"CPU_PPC_64":   4,
	"CPU_ARM":      5,
	"CPU_IA64":     6,
	"CPU_SH":       7,
	"CPU_MIPS":     8,
	"CPU_BLACKFIN": 9,
}

func (x CpuFamily) Enum() *CpuFamily {
	p := new(CpuFamily)
	*p = x
	return p
}
func (x CpuFamily) String() string {
	return proto.EnumName(CpuFamily_name, int32(x))
}
func (x *CpuFamily) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CpuFamily_value, data, "CpuFamily")
	if err != nil {
		return err
	}
	*x = CpuFamily(value)
	return nil
}
func (CpuFamily) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type Brand int32

const (
	Brand_BRAND_UNBRANDED Brand = 0
	Brand_BRAND_INQ       Brand = 1
	Brand_BRAND_HTC       Brand = 2
	Brand_BRAND_NOKIA     Brand = 3
)

var Brand_name = map[int32]string{
	0: "BRAND_UNBRANDED",
	1: "BRAND_INQ",
	2: "BRAND_HTC",
	3: "BRAND_NOKIA",
}
var Brand_value = map[string]int32{
	"BRAND_UNBRANDED": 0,
	"BRAND_INQ":       1,
	"BRAND_HTC":       2,
	"BRAND_NOKIA":     3,
}

func (x Brand) Enum() *Brand {
	p := new(Brand)
	*p = x
	return p
}
func (x Brand) String() string {
	return proto.EnumName(Brand_name, int32(x))
}
func (x *Brand) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Brand_value, data, "Brand")
	if err != nil {
		return err
	}
	*x = Brand(value)
	return nil
}
func (Brand) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type Os int32

const (
	Os_OS_UNKNOWN    Os = 0
	Os_OS_WINDOWS    Os = 1
	Os_OS_OSX        Os = 2
	Os_OS_IPHONE     Os = 3
	Os_OS_S60        Os = 4
	Os_OS_LINUX      Os = 5
	Os_OS_WINDOWS_CE Os = 6
	Os_OS_ANDROID    Os = 7
	Os_OS_PALM       Os = 8
	Os_OS_FREEBSD    Os = 9
	Os_OS_BLACKBERRY Os = 10
	Os_OS_SONOS      Os = 11
	Os_OS_LOGITECH   Os = 12
	Os_OS_WP7        Os = 13
	Os_OS_ONKYO      Os = 14
	Os_OS_PHILIPS    Os = 15
	Os_OS_WD         Os = 16
	Os_OS_VOLVO      Os = 17
	Os_OS_TIVO       Os = 18
	Os_OS_AWOX       Os = 19
	Os_OS_MEEGO      Os = 20
	Os_OS_QNXNTO     Os = 21
	Os_OS_BCO        Os = 22
)

var Os_name = map[int32]string{
	0:  "OS_UNKNOWN",
	1:  "OS_WINDOWS",
	2:  "OS_OSX",
	3:  "OS_IPHONE",
	4:  "OS_S60",
	5:  "OS_LINUX",
	6:  "OS_WINDOWS_CE",
	7:  "OS_ANDROID",
	8:  "OS_PALM",
	9:  "OS_FREEBSD",
	10: "OS_BLACKBERRY",
	11: "OS_SONOS",
	12: "OS_LOGITECH",
	13: "OS_WP7",
	14: "OS_ONKYO",
	15: "OS_PHILIPS",
	16: "OS_WD",
	17: "OS_VOLVO",
	18: "OS_TIVO",
	19: "OS_AWOX",
	20: "OS_MEEGO",
	21: "OS_QNXNTO",
	22: "OS_BCO",
}
var Os_value = map[string]int32{
	"OS_UNKNOWN":    0,
	"OS_WINDOWS":    1,
	"OS_OSX":        2,
	"OS_IPHONE":     3,
	"OS_S60":        4,
	"OS_LINUX":      5,
	"OS_WINDOWS_CE": 6,
	"OS_ANDROID":    7,
	"OS_PALM":       8,
	"OS_FREEBSD":    9,
	"OS_BLACKBERRY": 10,
	"OS_SONOS":      11,
	"OS_LOGITECH":   12,
	"OS_WP7":        13,
	"OS_ONKYO":      14,
	"OS_PHILIPS":    15,
	"OS_WD":         16,
	"OS_VOLVO":      17,
	"OS_TIVO":       18,
	"OS_AWOX":       19,
	"OS_MEEGO":      20,
	"OS_QNXNTO":     21,
	"OS_BCO":        22,
}

func (x Os) Enum() *Os {
	p := new(Os)
	*p = x
	return p
}
func (x Os) String() string {
	return proto.EnumName(Os_name, int32(x))
}
func (x *Os) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Os_value, data, "Os")
	if err != nil {
		return err
	}
	*x = Os(value)
	return nil
}
func (Os) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type AccountType int32

const (
	AccountType_Spotify  AccountType = 0
	AccountType_Facebook AccountType = 1
)

var AccountType_name = map[int32]string{
	0: "Spotify",
	1: "Facebook",
}
var AccountType_value = map[string]int32{
	"Spotify":  0,
	"Facebook": 1,
}

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}
func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}
func (x *AccountType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AccountType_value, data, "AccountType")
	if err != nil {
		return err
	}
	*x = AccountType(value)
	return nil
}
func (AccountType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type ClientResponseEncrypted struct {
	LoginCredentials    *LoginCredentials         `protobuf:"bytes,10,req,name=login_credentials,json=loginCredentials" json:"login_credentials,omitempty"`
	AccountCreation     *AccountCreation          `protobuf:"varint,20,opt,name=account_creation,json=accountCreation,enum=Spotify.AccountCreation" json:"account_creation,omitempty"`
	FingerprintResponse *FingerprintResponseUnion `protobuf:"bytes,30,opt,name=fingerprint_response,json=fingerprintResponse" json:"fingerprint_response,omitempty"`
	PeerTicket          *PeerTicketUnion          `protobuf:"bytes,40,opt,name=peer_ticket,json=peerTicket" json:"peer_ticket,omitempty"`
	SystemInfo          *SystemInfo               `protobuf:"bytes,50,req,name=system_info,json=systemInfo" json:"system_info,omitempty"`
	PlatformModel       *string                   `protobuf:"bytes,60,opt,name=platform_model,json=platformModel" json:"platform_model,omitempty"`
	VersionString       *string                   `protobuf:"bytes,70,opt,name=version_string,json=versionString" json:"version_string,omitempty"`
	Appkey              *LibspotifyAppKey         `protobuf:"bytes,80,opt,name=appkey" json:"appkey,omitempty"`
	ClientInfo          *ClientInfo               `protobuf:"bytes,90,opt,name=client_info,json=clientInfo" json:"client_info,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *ClientResponseEncrypted) Reset()                    { *m = ClientResponseEncrypted{} }
func (m *ClientResponseEncrypted) String() string            { return proto.CompactTextString(m) }
func (*ClientResponseEncrypted) ProtoMessage()               {}
func (*ClientResponseEncrypted) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ClientResponseEncrypted) GetLoginCredentials() *LoginCredentials {
	if m != nil {
		return m.LoginCredentials
	}
	return nil
}

func (m *ClientResponseEncrypted) GetAccountCreation() AccountCreation {
	if m != nil && m.AccountCreation != nil {
		return *m.AccountCreation
	}
	return AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT
}

func (m *ClientResponseEncrypted) GetFingerprintResponse() *FingerprintResponseUnion {
	if m != nil {
		return m.FingerprintResponse
	}
	return nil
}

func (m *ClientResponseEncrypted) GetPeerTicket() *PeerTicketUnion {
	if m != nil {
		return m.PeerTicket
	}
	return nil
}

func (m *ClientResponseEncrypted) GetSystemInfo() *SystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
}

func (m *ClientResponseEncrypted) GetPlatformModel() string {
	if m != nil && m.PlatformModel != nil {
		return *m.PlatformModel
	}
	return ""
}

func (m *ClientResponseEncrypted) GetVersionString() string {
	if m != nil && m.VersionString != nil {
		return *m.VersionString
	}
	return ""
}

func (m *ClientResponseEncrypted) GetAppkey() *LibspotifyAppKey {
	if m != nil {
		return m.Appkey
	}
	return nil
}

func (m *ClientResponseEncrypted) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

type LoginCredentials struct {
	Username         *string             `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Typ              *AuthenticationType `protobuf:"varint,20,req,name=typ,enum=Spotify.AuthenticationType" json:"typ,omitempty"`
	AuthData         []byte              `protobuf:"bytes,30,opt,name=auth_data,json=authData" json:"auth_data,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *LoginCredentials) Reset()                    { *m = LoginCredentials{} }
func (m *LoginCredentials) String() string            { return proto.CompactTextString(m) }
func (*LoginCredentials) ProtoMessage()               {}
func (*LoginCredentials) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *LoginCredentials) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *LoginCredentials) GetTyp() AuthenticationType {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return AuthenticationType_AUTHENTICATION_USER_PASS
}

func (m *LoginCredentials) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

type FingerprintResponseUnion struct {
	Grain            *FingerprintGrainResponse      `protobuf:"bytes,10,opt,name=grain" json:"grain,omitempty"`
	HmacRipemd       *FingerprintHmacRipemdResponse `protobuf:"bytes,20,opt,name=hmac_ripemd,json=hmacRipemd" json:"hmac_ripemd,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *FingerprintResponseUnion) Reset()                    { *m = FingerprintResponseUnion{} }
func (m *FingerprintResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*FingerprintResponseUnion) ProtoMessage()               {}
func (*FingerprintResponseUnion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *FingerprintResponseUnion) GetGrain() *FingerprintGrainResponse {
	if m != nil {
		return m.Grain
	}
	return nil
}

func (m *FingerprintResponseUnion) GetHmacRipemd() *FingerprintHmacRipemdResponse {
	if m != nil {
		return m.HmacRipemd
	}
	return nil
}

type FingerprintGrainResponse struct {
	EncryptedKey     []byte `protobuf:"bytes,10,req,name=encrypted_key,json=encryptedKey" json:"encrypted_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintGrainResponse) Reset()                    { *m = FingerprintGrainResponse{} }
func (m *FingerprintGrainResponse) String() string            { return proto.CompactTextString(m) }
func (*FingerprintGrainResponse) ProtoMessage()               {}
func (*FingerprintGrainResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *FingerprintGrainResponse) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

type FingerprintHmacRipemdResponse struct {
	Hmac             []byte `protobuf:"bytes,10,req,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintHmacRipemdResponse) Reset()                    { *m = FingerprintHmacRipemdResponse{} }
func (m *FingerprintHmacRipemdResponse) String() string            { return proto.CompactTextString(m) }
func (*FingerprintHmacRipemdResponse) ProtoMessage()               {}
func (*FingerprintHmacRipemdResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *FingerprintHmacRipemdResponse) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type PeerTicketUnion struct {
	PublicKey        *PeerTicketPublicKey `protobuf:"bytes,10,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	OldTicket        *PeerTicketOld       `protobuf:"bytes,20,opt,name=old_ticket,json=oldTicket" json:"old_ticket,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PeerTicketUnion) Reset()                    { *m = PeerTicketUnion{} }
func (m *PeerTicketUnion) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketUnion) ProtoMessage()               {}
func (*PeerTicketUnion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *PeerTicketUnion) GetPublicKey() *PeerTicketPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PeerTicketUnion) GetOldTicket() *PeerTicketOld {
	if m != nil {
		return m.OldTicket
	}
	return nil
}

type PeerTicketPublicKey struct {
	PublicKey        []byte `protobuf:"bytes,10,req,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PeerTicketPublicKey) Reset()                    { *m = PeerTicketPublicKey{} }
func (m *PeerTicketPublicKey) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketPublicKey) ProtoMessage()               {}
func (*PeerTicketPublicKey) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *PeerTicketPublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type PeerTicketOld struct {
	PeerTicket          []byte `protobuf:"bytes,10,req,name=peer_ticket,json=peerTicket" json:"peer_ticket,omitempty"`
	PeerTicketSignature []byte `protobuf:"bytes,20,req,name=peer_ticket_signature,json=peerTicketSignature" json:"peer_ticket_signature,omitempty"`
	XXX_unrecognized    []byte `json:"-"`
}

func (m *PeerTicketOld) Reset()                    { *m = PeerTicketOld{} }
func (m *PeerTicketOld) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketOld) ProtoMessage()               {}
func (*PeerTicketOld) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *PeerTicketOld) GetPeerTicket() []byte {
	if m != nil {
		return m.PeerTicket
	}
	return nil
}

func (m *PeerTicketOld) GetPeerTicketSignature() []byte {
	if m != nil {
		return m.PeerTicketSignature
	}
	return nil
}

type SystemInfo struct {
	CpuFamily               *CpuFamily `protobuf:"varint,10,req,name=cpu_family,json=cpuFamily,enum=Spotify.CpuFamily" json:"cpu_family,omitempty"`
	CpuSubtype              *uint32    `protobuf:"varint,20,opt,name=cpu_subtype,json=cpuSubtype" json:"cpu_subtype,omitempty"`
	CpuExt                  *uint32    `protobuf:"varint,30,opt,name=cpu_ext,json=cpuExt" json:"cpu_ext,omitempty"`
	Brand                   *Brand     `protobuf:"varint,40,opt,name=brand,enum=Spotify.Brand" json:"brand,omitempty"`
	BrandFlags              *uint32    `protobuf:"varint,50,opt,name=brand_flags,json=brandFlags" json:"brand_flags,omitempty"`
	Os                      *Os        `protobuf:"varint,60,req,name=os,enum=Spotify.Os" json:"os,omitempty"`
	OsVersion               *uint32    `protobuf:"varint,70,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	OsExt                   *uint32    `protobuf:"varint,80,opt,name=os_ext,json=osExt" json:"os_ext,omitempty"`
	SystemInformationString *string    `protobuf:"bytes,90,opt,name=system_information_string,json=systemInformationString" json:"system_information_string,omitempty"`
	DeviceId                *string    `protobuf:"bytes,100,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	XXX_unrecognized        []byte     `json:"-"`
}

func (m *SystemInfo) Reset()                    { *m = SystemInfo{} }
func (m *SystemInfo) String() string            { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()               {}
func (*SystemInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *SystemInfo) GetCpuFamily() CpuFamily {
	if m != nil && m.CpuFamily != nil {
		return *m.CpuFamily
	}
	return CpuFamily_CPU_UNKNOWN
}

func (m *SystemInfo) GetCpuSubtype() uint32 {
	if m != nil && m.CpuSubtype != nil {
		return *m.CpuSubtype
	}
	return 0
}

func (m *SystemInfo) GetCpuExt() uint32 {
	if m != nil && m.CpuExt != nil {
		return *m.CpuExt
	}
	return 0
}

func (m *SystemInfo) GetBrand() Brand {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return Brand_BRAND_UNBRANDED
}

func (m *SystemInfo) GetBrandFlags() uint32 {
	if m != nil && m.BrandFlags != nil {
		return *m.BrandFlags
	}
	return 0
}

func (m *SystemInfo) GetOs() Os {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return Os_OS_UNKNOWN
}

func (m *SystemInfo) GetOsVersion() uint32 {
	if m != nil && m.OsVersion != nil {
		return *m.OsVersion
	}
	return 0
}

func (m *SystemInfo) GetOsExt() uint32 {
	if m != nil && m.OsExt != nil {
		return *m.OsExt
	}
	return 0
}

func (m *SystemInfo) GetSystemInformationString() string {
	if m != nil && m.SystemInformationString != nil {
		return *m.SystemInformationString
	}
	return ""
}

func (m *SystemInfo) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

type LibspotifyAppKey struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Devkey           []byte  `protobuf:"bytes,2,req,name=devkey" json:"devkey,omitempty"`
	Signature        []byte  `protobuf:"bytes,3,req,name=signature" json:"signature,omitempty"`
	Useragent        *string `protobuf:"bytes,4,req,name=useragent" json:"useragent,omitempty"`
	CallbackHash     []byte  `protobuf:"bytes,5,req,name=callback_hash,json=callbackHash" json:"callback_hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LibspotifyAppKey) Reset()                    { *m = LibspotifyAppKey{} }
func (m *LibspotifyAppKey) String() string            { return proto.CompactTextString(m) }
func (*LibspotifyAppKey) ProtoMessage()               {}
func (*LibspotifyAppKey) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *LibspotifyAppKey) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *LibspotifyAppKey) GetDevkey() []byte {
	if m != nil {
		return m.Devkey
	}
	return nil
}

func (m *LibspotifyAppKey) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *LibspotifyAppKey) GetUseragent() string {
	if m != nil && m.Useragent != nil {
		return *m.Useragent
	}
	return ""
}

func (m *LibspotifyAppKey) GetCallbackHash() []byte {
	if m != nil {
		return m.CallbackHash
	}
	return nil
}

type ClientInfo struct {
	Limited          *bool               `protobuf:"varint,1,opt,name=limited" json:"limited,omitempty"`
	Fb               *ClientInfoFacebook `protobuf:"bytes,2,opt,name=fb" json:"fb,omitempty"`
	Language         *string             `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *ClientInfo) GetLimited() bool {
	if m != nil && m.Limited != nil {
		return *m.Limited
	}
	return false
}

func (m *ClientInfo) GetFb() *ClientInfoFacebook {
	if m != nil {
		return m.Fb
	}
	return nil
}

func (m *ClientInfo) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type ClientInfoFacebook struct {
	MachineId        *string `protobuf:"bytes,1,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientInfoFacebook) Reset()                    { *m = ClientInfoFacebook{} }
func (m *ClientInfoFacebook) String() string            { return proto.CompactTextString(m) }
func (*ClientInfoFacebook) ProtoMessage()               {}
func (*ClientInfoFacebook) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ClientInfoFacebook) GetMachineId() string {
	if m != nil && m.MachineId != nil {
		return *m.MachineId
	}
	return ""
}

type APWelcome struct {
	CanonicalUsername           *string              `protobuf:"bytes,10,req,name=canonical_username,json=canonicalUsername" json:"canonical_username,omitempty"`
	AccountTypeLoggedIn         *AccountType         `protobuf:"varint,20,req,name=account_type_logged_in,json=accountTypeLoggedIn,enum=Spotify.AccountType" json:"account_type_logged_in,omitempty"`
	CredentialsTypeLoggedIn     *AccountType         `protobuf:"varint,25,req,name=credentials_type_logged_in,json=credentialsTypeLoggedIn,enum=Spotify.AccountType" json:"credentials_type_logged_in,omitempty"`
	ReusableAuthCredentialsType *AuthenticationType  `protobuf:"varint,30,req,name=reusable_auth_credentials_type,json=reusableAuthCredentialsType,enum=Spotify.AuthenticationType" json:"reusable_auth_credentials_type,omitempty"`
	ReusableAuthCredentials     []byte               `protobuf:"bytes,40,req,name=reusable_auth_credentials,json=reusableAuthCredentials" json:"reusable_auth_credentials,omitempty"`
	LfsSecret                   []byte               `protobuf:"bytes,50,opt,name=lfs_secret,json=lfsSecret" json:"lfs_secret,omitempty"`
	AccountInfo                 *AccountInfo         `protobuf:"bytes,60,opt,name=account_info,json=accountInfo" json:"account_info,omitempty"`
	Fb                          *AccountInfoFacebook `protobuf:"bytes,70,opt,name=fb" json:"fb,omitempty"`
	XXX_unrecognized            []byte               `json:"-"`
}

func (m *APWelcome) Reset()                    { *m = APWelcome{} }
func (m *APWelcome) String() string            { return proto.CompactTextString(m) }
func (*APWelcome) ProtoMessage()               {}
func (*APWelcome) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *APWelcome) GetCanonicalUsername() string {
	if m != nil && m.CanonicalUsername != nil {
		return *m.CanonicalUsername
	}
	return ""
}

func (m *APWelcome) GetAccountTypeLoggedIn() AccountType {
	if m != nil && m.AccountTypeLoggedIn != nil {
		return *m.AccountTypeLoggedIn
	}
	return AccountType_Spotify
}

func (m *APWelcome) GetCredentialsTypeLoggedIn() AccountType {
	if m != nil && m.CredentialsTypeLoggedIn != nil {
		return *m.CredentialsTypeLoggedIn
	}
	return AccountType_Spotify
}

func (m *APWelcome) GetReusableAuthCredentialsType() AuthenticationType {
	if m != nil && m.ReusableAuthCredentialsType != nil {
		return *m.ReusableAuthCredentialsType
	}
	return AuthenticationType_AUTHENTICATION_USER_PASS
}

func (m *APWelcome) GetReusableAuthCredentials() []byte {
	if m != nil {
		return m.ReusableAuthCredentials
	}
	return nil
}

func (m *APWelcome) GetLfsSecret() []byte {
	if m != nil {
		return m.LfsSecret
	}
	return nil
}

func (m *APWelcome) GetAccountInfo() *AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

func (m *APWelcome) GetFb() *AccountInfoFacebook {
	if m != nil {
		return m.Fb
	}
	return nil
}

type AccountInfo struct {
	Spotify          *AccountInfoSpotify  `protobuf:"bytes,1,opt,name=spotify" json:"spotify,omitempty"`
	Facebook         *AccountInfoFacebook `protobuf:"bytes,2,opt,name=facebook" json:"facebook,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AccountInfo) Reset()                    { *m = AccountInfo{} }
func (m *AccountInfo) String() string            { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()               {}
func (*AccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *AccountInfo) GetSpotify() *AccountInfoSpotify {
	if m != nil {
		return m.Spotify
	}
	return nil
}

func (m *AccountInfo) GetFacebook() *AccountInfoFacebook {
	if m != nil {
		return m.Facebook
	}
	return nil
}

type AccountInfoSpotify struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AccountInfoSpotify) Reset()                    { *m = AccountInfoSpotify{} }
func (m *AccountInfoSpotify) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoSpotify) ProtoMessage()               {}
func (*AccountInfoSpotify) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

type AccountInfoFacebook struct {
	AccessToken      *string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	MachineId        *string `protobuf:"bytes,2,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AccountInfoFacebook) Reset()                    { *m = AccountInfoFacebook{} }
func (m *AccountInfoFacebook) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoFacebook) ProtoMessage()               {}
func (*AccountInfoFacebook) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *AccountInfoFacebook) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *AccountInfoFacebook) GetMachineId() string {
	if m != nil && m.MachineId != nil {
		return *m.MachineId
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientResponseEncrypted)(nil), "Spotify.ClientResponseEncrypted")
	proto.RegisterType((*LoginCredentials)(nil), "Spotify.LoginCredentials")
	proto.RegisterType((*FingerprintResponseUnion)(nil), "Spotify.FingerprintResponseUnion")
	proto.RegisterType((*FingerprintGrainResponse)(nil), "Spotify.FingerprintGrainResponse")
	proto.RegisterType((*FingerprintHmacRipemdResponse)(nil), "Spotify.FingerprintHmacRipemdResponse")
	proto.RegisterType((*PeerTicketUnion)(nil), "Spotify.PeerTicketUnion")
	proto.RegisterType((*PeerTicketPublicKey)(nil), "Spotify.PeerTicketPublicKey")
	proto.RegisterType((*PeerTicketOld)(nil), "Spotify.PeerTicketOld")
	proto.RegisterType((*SystemInfo)(nil), "Spotify.SystemInfo")
	proto.RegisterType((*LibspotifyAppKey)(nil), "Spotify.LibspotifyAppKey")
	proto.RegisterType((*ClientInfo)(nil), "Spotify.ClientInfo")
	proto.RegisterType((*ClientInfoFacebook)(nil), "Spotify.ClientInfoFacebook")
	proto.RegisterType((*APWelcome)(nil), "Spotify.APWelcome")
	proto.RegisterType((*AccountInfo)(nil), "Spotify.AccountInfo")
	proto.RegisterType((*AccountInfoSpotify)(nil), "Spotify.AccountInfoSpotify")
	proto.RegisterType((*AccountInfoFacebook)(nil), "Spotify.AccountInfoFacebook")
	proto.RegisterEnum("Spotify.AuthenticationType", AuthenticationType_name, AuthenticationType_value)
	proto.RegisterEnum("Spotify.AccountCreation", AccountCreation_name, AccountCreation_value)
	proto.RegisterEnum("Spotify.CpuFamily", CpuFamily_name, CpuFamily_value)
	proto.RegisterEnum("Spotify.Brand", Brand_name, Brand_value)
	proto.RegisterEnum("Spotify.Os", Os_name, Os_value)
	proto.RegisterEnum("Spotify.AccountType", AccountType_name, AccountType_value)
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x57, 0x4b, 0x73, 0xdb, 0xc8,
	0x11, 0x36, 0xa1, 0x27, 0x9b, 0xa4, 0x34, 0x1e, 0xc9, 0x36, 0xbc, 0xb2, 0x1d, 0x9a, 0x79, 0x14,
	0xa3, 0x64, 0x5d, 0x59, 0xae, 0xd7, 0xde, 0x24, 0xae, 0x4a, 0x41, 0x20, 0x68, 0xa1, 0x48, 0x01,
	0xf0, 0x80, 0xb4, 0xa4, 0xbd, 0x20, 0x20, 0x30, 0xa4, 0x50, 0x02, 0x01, 0x14, 0x00, 0xba, 0x96,
	0x39, 0xe4, 0x96, 0x7f, 0x91, 0xdc, 0x92, 0x1f, 0x94, 0x4b, 0x7e, 0x45, 0x7e, 0x42, 0x0e, 0xa9,
	0x19, 0x3c, 0xf8, 0x90, 0xec, 0x3d, 0x09, 0xdd, 0xfd, 0xf5, 0xd7, 0x3d, 0xd3, 0x8f, 0x11, 0xe1,
	0xd8, 0x9e, 0xa7, 0x37, 0x34, 0x48, 0x3d, 0xc7, 0x4e, 0xbd, 0x30, 0x78, 0x15, 0xc5, 0x61, 0x1a,
	0xe2, 0x3d, 0x33, 0x0a, 0x53, 0x6f, 0xb2, 0x68, 0xfd, 0x73, 0x1b, 0x9e, 0xc8, 0xbe, 0x47, 0x83,
	0x94, 0xd0, 0x24, 0x0a, 0x83, 0x84, 0x2a, 0x81, 0x13, 0x2f, 0xa2, 0x94, 0xba, 0xb8, 0x07, 0x0f,
	0xfd, 0x70, 0xea, 0x05, 0x96, 0x13, 0x53, 0x97, 0x51, 0xd8, 0x7e, 0x22, 0x42, 0x53, 0x68, 0xd7,
	0x3a, 0x4f, 0x5f, 0xe5, 0x04, 0xaf, 0x06, 0x0c, 0x21, 0x2f, 0x01, 0x04, 0xf9, 0x1b, 0x1a, 0x2c,
	0x03, 0xb2, 0x1d, 0x27, 0x9c, 0x07, 0x29, 0x63, 0xe2, 0x69, 0x88, 0xc7, 0xcd, 0x4a, 0xfb, 0xa0,
	0x23, 0x96, 0x34, 0x52, 0x06, 0x90, 0x73, 0x3b, 0x39, 0xb4, 0xd7, 0x15, 0x78, 0x08, 0xc7, 0x13,
	0x2f, 0x98, 0xd2, 0x38, 0x8a, 0xbd, 0x20, 0xb5, 0xe2, 0x3c, 0x5b, 0xf1, 0x45, 0xb3, 0xd2, 0xae,
	0x75, 0x5e, 0x96, 0x44, 0xbd, 0x25, 0xa8, 0x38, 0xd1, 0x28, 0x60, 0x8c, 0x47, 0x93, 0xbb, 0x16,
	0xfc, 0x7b, 0xa8, 0x45, 0x94, 0xc6, 0x56, 0xea, 0x39, 0xb7, 0x34, 0x15, 0xdb, 0x9c, 0x6c, 0x99,
	0x95, 0x41, 0x69, 0x3c, 0xe4, 0xa6, 0x8c, 0x03, 0xa2, 0x52, 0x81, 0x5f, 0x43, 0x2d, 0x59, 0x24,
	0x29, 0x9d, 0x59, 0x5e, 0x30, 0x09, 0xc5, 0x0e, 0xbf, 0x97, 0xa3, 0xd2, 0xd5, 0xe4, 0x36, 0x35,
	0x98, 0x84, 0x04, 0x92, 0xf2, 0x1b, 0xff, 0x12, 0x0e, 0x22, 0xdf, 0x4e, 0x27, 0x61, 0x3c, 0xb3,
	0x66, 0xa1, 0x4b, 0x7d, 0xf1, 0x5d, 0xb3, 0xd2, 0xae, 0x92, 0x46, 0xa1, 0xbd, 0x60, 0x4a, 0x06,
	0xfb, 0x44, 0xe3, 0xc4, 0x0b, 0x03, 0x2b, 0x49, 0x63, 0x2f, 0x98, 0x8a, 0xbd, 0x0c, 0x96, 0x6b,
	0x4d, 0xae, 0xc4, 0xdf, 0xc0, 0xae, 0x1d, 0x45, 0xb7, 0x74, 0x21, 0x1a, 0x3c, 0xf3, 0x95, 0xb2,
	0x78, 0xe3, 0x24, 0xfb, 0x94, 0xa2, 0xa8, 0x4f, 0x17, 0x24, 0x07, 0xb2, 0xb4, 0x1d, 0x5e, 0xef,
	0x2c, 0xed, 0x1f, 0xb8, 0xdf, 0x32, 0xed, 0xac, 0x17, 0xb2, 0xb4, 0x9d, 0xf2, 0xbb, 0xf5, 0x17,
	0x40, 0x9b, 0x85, 0xc6, 0x5f, 0xc1, 0xfe, 0x3c, 0xa1, 0x71, 0x60, 0xcf, 0xa8, 0x08, 0x3c, 0xbb,
	0x52, 0xc6, 0x5f, 0xc3, 0x56, 0xba, 0x88, 0xc4, 0xe3, 0xa6, 0xd0, 0x3e, 0xe8, 0x9c, 0x2c, 0xab,
	0xbc, 0xd6, 0x8b, 0xc3, 0x45, 0x44, 0x09, 0xc3, 0xe1, 0x13, 0xa8, 0xb2, 0x36, 0xb5, 0x5c, 0x3b,
	0xb5, 0x79, 0x45, 0xeb, 0x64, 0x9f, 0x29, 0xba, 0x76, 0x6a, 0xb7, 0xfe, 0x5e, 0x01, 0xf1, 0x73,
	0x55, 0xc5, 0x6f, 0x61, 0x67, 0x1a, 0xdb, 0x5e, 0xc0, 0x33, 0xf8, 0x4c, 0x1f, 0xbc, 0x67, 0x80,
	0xc2, 0x8d, 0x64, 0x78, 0xfc, 0x1e, 0x6a, 0x37, 0x33, 0xdb, 0xb1, 0x62, 0x2f, 0xa2, 0x33, 0x97,
	0xf7, 0x63, 0xad, 0xf3, 0xab, 0xfb, 0xdc, 0xcf, 0x67, 0xb6, 0x43, 0x38, 0xaa, 0xe4, 0x80, 0x9b,
	0x52, 0xd7, 0xfa, 0xd3, 0x5a, 0x76, 0x6b, 0xb1, 0xf0, 0xcf, 0xa1, 0x41, 0x8b, 0x71, 0xb2, 0x58,
	0x99, 0xd8, 0xf4, 0xd4, 0x49, 0xbd, 0x54, 0xf6, 0xe9, 0xa2, 0xf5, 0x2d, 0x3c, 0xff, 0x62, 0x34,
	0x8c, 0x61, 0x9b, 0xc5, 0xcb, 0x9d, 0xf9, 0x77, 0xeb, 0x6f, 0x15, 0x38, 0xdc, 0xe8, 0x4e, 0xfc,
	0x47, 0x80, 0x68, 0x3e, 0xf6, 0x3d, 0x27, 0x0f, 0xc5, 0x4e, 0xf4, 0xec, 0x9e, 0x5e, 0x36, 0x38,
	0x88, 0x35, 0x45, 0x35, 0x2a, 0x3e, 0xf1, 0x77, 0x00, 0xa1, 0xef, 0x16, 0x83, 0x90, 0x5d, 0xc7,
	0xe3, 0x7b, 0x9c, 0x75, 0xdf, 0x25, 0xd5, 0xd0, 0x77, 0x33, 0xa9, 0xf5, 0x1a, 0x8e, 0xee, 0x21,
	0xc6, 0xcf, 0x37, 0x52, 0x61, 0x89, 0x2f, 0x83, 0xb5, 0x5c, 0x68, 0xac, 0x31, 0xe2, 0x9f, 0xad,
	0xcf, 0x61, 0xe6, 0xb0, 0x3a, 0x6d, 0x1d, 0x78, 0xb4, 0x02, 0xb0, 0x12, 0x6f, 0x1a, 0xd8, 0xe9,
	0x3c, 0xa6, 0xbc, 0xc5, 0xea, 0xe4, 0x68, 0x09, 0x35, 0x0b, 0x53, 0xeb, 0xbf, 0x02, 0xc0, 0x72,
	0x0c, 0xf1, 0x37, 0x00, 0x4e, 0x34, 0xb7, 0x26, 0xf6, 0xcc, 0xf3, 0xb3, 0x9c, 0x0e, 0x3a, 0x78,
	0xd9, 0xf8, 0xd1, 0xbc, 0xc7, 0x2d, 0xa4, 0xea, 0x14, 0x9f, 0x2c, 0x2d, 0xe6, 0x92, 0xcc, 0xc7,
	0xe9, 0x22, 0xa2, 0xfc, 0x56, 0x1a, 0x84, 0xb1, 0x98, 0x99, 0x06, 0x3f, 0x81, 0x3d, 0x06, 0xa0,
	0x3f, 0xa6, 0xbc, 0x6d, 0x1b, 0x64, 0xd7, 0x89, 0xe6, 0xca, 0x8f, 0x29, 0xfe, 0x05, 0xec, 0x8c,
	0x63, 0x3b, 0x70, 0xf9, 0x4a, 0x39, 0xe8, 0x1c, 0x94, 0x71, 0xce, 0x98, 0x96, 0x64, 0x46, 0xc6,
	0xcf, 0x3f, 0xac, 0x89, 0x6f, 0x4f, 0x13, 0xb1, 0x93, 0xf1, 0x73, 0x55, 0x8f, 0x69, 0xf0, 0x09,
	0x08, 0x61, 0x22, 0xbe, 0xe3, 0xb9, 0xd6, 0x4a, 0x0e, 0x3d, 0x21, 0x42, 0x98, 0xb0, 0x4b, 0x0e,
	0x13, 0x2b, 0xdf, 0x08, 0x7c, 0x41, 0x34, 0x48, 0x35, 0x4c, 0x3e, 0x66, 0x0a, 0xfc, 0x08, 0x76,
	0xc3, 0x84, 0xa7, 0x66, 0x70, 0xd3, 0x4e, 0x98, 0xb0, 0xcc, 0xfe, 0x00, 0x4f, 0x57, 0xf6, 0x56,
	0x3c, 0xe3, 0xa3, 0x58, 0x6c, 0x99, 0x1f, 0xf8, 0x1c, 0x3f, 0x59, 0x2e, 0xac, 0xdc, 0x9e, 0xef,
	0x9b, 0x13, 0xa8, 0xba, 0xf4, 0x93, 0xe7, 0x50, 0xcb, 0x73, 0x45, 0x37, 0x9b, 0xf9, 0x4c, 0xa1,
	0xba, 0xad, 0x7f, 0x55, 0x00, 0x6d, 0xae, 0x1d, 0x2c, 0xc2, 0x5e, 0x91, 0x60, 0xa5, 0x29, 0xb4,
	0x1b, 0xa4, 0x10, 0xf1, 0x63, 0xd8, 0x75, 0xe9, 0x27, 0xd6, 0x1e, 0x02, 0x2f, 0x61, 0x2e, 0xe1,
	0x67, 0x50, 0x5d, 0x56, 0x77, 0x2b, 0xeb, 0x9c, 0x52, 0xc1, 0xac, 0x6c, 0xc9, 0xd8, 0x53, 0x1a,
	0xa4, 0xe2, 0x76, 0x53, 0x68, 0x57, 0xc9, 0x52, 0xc1, 0xe6, 0xcd, 0xb1, 0x7d, 0x7f, 0x6c, 0x3b,
	0xb7, 0xd6, 0x8d, 0x9d, 0xdc, 0x88, 0x3b, 0xd9, 0xbc, 0x15, 0xca, 0x73, 0x3b, 0xb9, 0x69, 0x85,
	0x00, 0xcb, 0x2d, 0xc7, 0x12, 0xf4, 0xbd, 0x99, 0x97, 0x52, 0x57, 0xac, 0x34, 0x2b, 0xed, 0x7d,
	0x52, 0x88, 0xf8, 0x37, 0x20, 0x4c, 0xc6, 0xa2, 0xc0, 0x27, 0xe1, 0xe4, 0x9e, 0x05, 0xd9, 0xb3,
	0x1d, 0x3a, 0x0e, 0xc3, 0x5b, 0x22, 0x4c, 0xc6, 0x6c, 0x19, 0xfa, 0x76, 0x30, 0x9d, 0xdb, 0x53,
	0x96, 0x34, 0xbf, 0x98, 0x42, 0x6e, 0x7d, 0x0b, 0xf8, 0xae, 0x17, 0xab, 0xde, 0xcc, 0x76, 0x6e,
	0xbc, 0x80, 0x5f, 0x66, 0x85, 0xfb, 0x54, 0x73, 0x8d, 0xea, 0xb6, 0xfe, 0xb7, 0x05, 0x55, 0xc9,
	0xb8, 0xa4, 0xbe, 0x13, 0xf2, 0x7d, 0x8a, 0x1d, 0x3b, 0x08, 0x03, 0xcf, 0xb1, 0x7d, 0x6b, 0x65,
	0xeb, 0xb2, 0xf3, 0x3f, 0x2c, 0x2d, 0xa3, 0x62, 0xfd, 0xaa, 0xf0, 0xb8, 0x78, 0x71, 0x59, 0x9b,
	0x5a, 0x7e, 0x38, 0x9d, 0x52, 0xd7, 0xf2, 0x82, 0x7c, 0x23, 0x1f, 0x6f, 0xbe, 0xbb, 0x7c, 0x15,
	0x1f, 0xd9, 0x4b, 0x61, 0xc0, 0x3d, 0xd4, 0x00, 0x7f, 0x80, 0xaf, 0x56, 0x9e, 0xff, 0x4d, 0xba,
	0xa7, 0x5f, 0xa0, 0x7b, 0xb2, 0xe2, 0xb7, 0x46, 0xf9, 0x67, 0x78, 0x11, 0xd3, 0x79, 0x62, 0x8f,
	0x7d, 0x6a, 0xf1, 0xb5, 0xbf, 0x19, 0x40, 0x7c, 0xf1, 0xd3, 0xef, 0xc6, 0x49, 0x41, 0xc1, 0x6c,
	0xf2, 0x7a, 0x24, 0xd6, 0xe3, 0x9f, 0x8d, 0x20, 0xb6, 0x79, 0x4f, 0x3c, 0xf9, 0x8c, 0x3f, 0xab,
	0x8b, 0x3f, 0x49, 0xac, 0x84, 0x3a, 0x31, 0x4d, 0xf9, 0x48, 0xd6, 0x49, 0xd5, 0x9f, 0x24, 0x26,
	0x57, 0xe0, 0xb7, 0x50, 0x2f, 0xae, 0x96, 0x3f, 0xa0, 0xef, 0x78, 0x7f, 0xdc, 0xb9, 0x01, 0xfe,
	0x82, 0xd6, 0xec, 0xa5, 0x80, 0x7f, 0xcb, 0xdb, 0xa9, 0xb7, 0xb1, 0x95, 0x57, 0xe0, 0xab, 0xfd,
	0xd4, 0xfa, 0x2b, 0xd4, 0x56, 0x4c, 0xf8, 0x3b, 0xd8, 0xcb, 0xe7, 0x8a, 0x77, 0xca, 0x6a, 0x43,
	0xae, 0xc0, 0x72, 0x15, 0x29, 0xb0, 0xf8, 0x7b, 0xd8, 0x9f, 0xe4, 0xac, 0x79, 0x23, 0x7f, 0x39,
	0x72, 0x89, 0x6e, 0x1d, 0x03, 0xbe, 0x4b, 0xdc, 0xba, 0x84, 0xa3, 0x7b, 0xdc, 0xf0, 0x4b, 0x7e,
	0x27, 0x34, 0x49, 0xac, 0x34, 0xbc, 0xa5, 0x41, 0xde, 0xcc, 0xb5, 0x4c, 0x37, 0x64, 0xaa, 0x8d,
	0x6e, 0x17, 0x36, 0xba, 0xfd, 0xf4, 0x3f, 0x15, 0xc0, 0x77, 0x8b, 0x8c, 0x9f, 0x81, 0x28, 0x8d,
	0x86, 0xe7, 0x8a, 0x36, 0x54, 0x65, 0x69, 0xa8, 0xea, 0x9a, 0x35, 0x32, 0x15, 0x62, 0x19, 0x92,
	0x69, 0xa2, 0x07, 0xf8, 0x6b, 0xf8, 0xf5, 0x86, 0xd5, 0x1c, 0xea, 0x44, 0xe9, 0x5a, 0xa6, 0xa1,
	0x0f, 0xd5, 0xde, 0xb5, 0x25, 0x13, 0xa5, 0xcb, 0xac, 0xd2, 0xc0, 0x44, 0x15, 0xfc, 0x0a, 0x4e,
	0xef, 0x87, 0xf7, 0x24, 0x59, 0x39, 0xd3, 0xf5, 0xfe, 0x1a, 0x5e, 0xc0, 0x4d, 0x78, 0xb6, 0x89,
	0xcf, 0x79, 0x87, 0x7a, 0x5f, 0xd1, 0xd0, 0x16, 0x7e, 0x09, 0xcf, 0x37, 0x10, 0x25, 0x55, 0x06,
	0xd9, 0x3e, 0xbd, 0x86, 0xc3, 0x8d, 0x7f, 0x6d, 0x71, 0x0b, 0x5e, 0x48, 0xb2, 0xac, 0x8f, 0xb4,
	0x21, 0x0b, 0x98, 0xf9, 0x49, 0x83, 0x4b, 0xe9, 0xda, 0xb4, 0x0c, 0xa2, 0x5f, 0x18, 0x43, 0x54,
	0xf9, 0x12, 0x86, 0xcb, 0x0a, 0xda, 0x3a, 0xfd, 0x47, 0x05, 0xaa, 0xe5, 0xab, 0x85, 0x0f, 0xa1,
	0x26, 0x1b, 0x23, 0x6b, 0xa4, 0xf5, 0x35, 0xfd, 0x52, 0x43, 0x0f, 0x70, 0x0d, 0xf6, 0x98, 0xe2,
	0xea, 0xfb, 0x37, 0xa8, 0x82, 0x0f, 0x00, 0x72, 0xc1, 0x7a, 0xf3, 0x1a, 0x09, 0x85, 0xd1, 0x30,
	0x64, 0xb4, 0x55, 0x18, 0x0d, 0x43, 0x66, 0xc6, 0xed, 0xc2, 0x28, 0x91, 0x0b, 0xb4, 0x83, 0xeb,
	0xb0, 0xcf, 0x04, 0x55, 0x7a, 0xf3, 0x1a, 0xed, 0x62, 0x80, 0x5d, 0x26, 0x99, 0xe7, 0x68, 0xaf,
	0xb0, 0x5c, 0xa8, 0x86, 0x89, 0xf6, 0x31, 0x82, 0x3a, 0x93, 0xce, 0x06, 0x92, 0xdc, 0xef, 0xa9,
	0x1a, 0xaa, 0x9e, 0xf6, 0x61, 0x87, 0x3f, 0x76, 0xf8, 0x08, 0x0e, 0xcf, 0x88, 0xa4, 0x75, 0xad,
	0x91, 0xc6, 0xff, 0x2a, 0x5d, 0xf4, 0x00, 0x37, 0xa0, 0x9a, 0x29, 0x55, 0xed, 0x03, 0xaa, 0x2c,
	0xc5, 0xf3, 0xa1, 0x8c, 0x04, 0x76, 0x9a, 0x4c, 0xd4, 0xf4, 0xbe, 0x2a, 0xa1, 0xad, 0xd3, 0x7f,
	0x0b, 0x20, 0xe8, 0x09, 0x4b, 0x55, 0x37, 0x57, 0x0e, 0x99, 0xc9, 0x97, 0xaa, 0xd6, 0xd5, 0x2f,
	0x59, 0x8d, 0x01, 0x76, 0x75, 0xd3, 0xd2, 0xcd, 0x2b, 0x24, 0x30, 0x4a, 0xdd, 0xb4, 0x54, 0xe3,
	0x5c, 0xd7, 0x14, 0xb4, 0x95, 0x9b, 0xcc, 0x37, 0xbf, 0x43, 0xdb, 0x2c, 0x75, 0xdd, 0xb4, 0x06,
	0xaa, 0x36, 0xba, 0x42, 0x3b, 0xf8, 0x21, 0x34, 0x96, 0x24, 0x96, 0xac, 0xa0, 0xdd, 0x9c, 0x57,
	0xd2, 0xba, 0x44, 0x57, 0xbb, 0x68, 0x8f, 0x5d, 0x89, 0x6e, 0x5a, 0x86, 0x34, 0xb8, 0x40, 0xfb,
	0xb9, 0xb1, 0x47, 0x14, 0xe5, 0xcc, 0xec, 0xa2, 0x6a, 0xee, 0xcf, 0x4f, 0x7e, 0xa6, 0x10, 0x72,
	0x8d, 0x20, 0x0f, 0x60, 0xea, 0x9a, 0x6e, 0xa2, 0x1a, 0x3b, 0x0d, 0x0b, 0xa7, 0xbf, 0x57, 0x87,
	0x8a, 0x7c, 0x8e, 0xea, 0x79, 0x2e, 0x97, 0xc6, 0x5b, 0xd4, 0xc8, 0xa1, 0xba, 0xd6, 0xbf, 0xd6,
	0xd1, 0x41, 0xce, 0x6d, 0x9c, 0xab, 0x03, 0x76, 0xad, 0x87, 0xb8, 0x0a, 0x3b, 0x0c, 0xd9, 0x45,
	0x28, 0x07, 0x7e, 0xd4, 0x07, 0x1f, 0x75, 0xf4, 0x30, 0xcf, 0x68, 0xa8, 0x7e, 0xd4, 0x11, 0xce,
	0x05, 0xe9, 0x52, 0xbf, 0x42, 0x47, 0x39, 0xee, 0x42, 0x51, 0xde, 0xeb, 0xe8, 0x38, 0xbf, 0x85,
	0x0f, 0xda, 0x95, 0x36, 0xd4, 0xd1, 0xa3, 0x3c, 0xf2, 0x99, 0xac, 0xa3, 0xc7, 0xa7, 0xed, 0x72,
	0xc7, 0xf0, 0x61, 0xab, 0x41, 0xf1, 0xab, 0x10, 0x3d, 0x60, 0x24, 0xc5, 0x78, 0xa3, 0xca, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x5f, 0x12, 0x81, 0x45, 0x0e, 0x00, 0x00,
}
