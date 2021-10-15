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

func (x AuthenticationType) Enum() *AuthenticationType {
	p := new(AuthenticationType)
	*p = x
	return p
}

type AccountCreation int32

const (
	AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT AccountCreation = 1
	AccountCreation_ACCOUNT_CREATION_ALWAYS_CREATE AccountCreation = 3
)

func (x AccountCreation) Enum() *AccountCreation {
	p := new(AccountCreation)
	*p = x
	return p
}

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

func (x CpuFamily) Enum() *CpuFamily {
	p := new(CpuFamily)
	*p = x
	return p
}

type Brand int32

const (
	Brand_BRAND_UNBRANDED Brand = 0
	Brand_BRAND_INQ       Brand = 1
	Brand_BRAND_HTC       Brand = 2
	Brand_BRAND_NOKIA     Brand = 3
)

func (x Brand) Enum() *Brand {
	p := new(Brand)
	*p = x
	return p
}

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

func (x Os) Enum() *Os {
	p := new(Os)
	*p = x
	return p
}

type AccountType int32

const (
	AccountType_Spotify  AccountType = 0
	AccountType_Facebook AccountType = 1
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

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

type LoginCredentials struct {
	Username         *string             `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Typ              *AuthenticationType `protobuf:"varint,20,req,name=typ,enum=Spotify.AuthenticationType" json:"typ,omitempty"`
	AuthData         []byte              `protobuf:"bytes,30,opt,name=auth_data,json=authData" json:"auth_data,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *LoginCredentials) Reset()                    { *m = LoginCredentials{} }
func (m *LoginCredentials) String() string            { return proto.CompactTextString(m) }
func (*LoginCredentials) ProtoMessage()               {}

type FingerprintResponseUnion struct {
	Grain            *FingerprintGrainResponse      `protobuf:"bytes,10,opt,name=grain" json:"grain,omitempty"`
	HmacRipemd       *FingerprintHmacRipemdResponse `protobuf:"bytes,20,opt,name=hmac_ripemd,json=hmacRipemd" json:"hmac_ripemd,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *FingerprintResponseUnion) Reset()                    { *m = FingerprintResponseUnion{} }
func (m *FingerprintResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*FingerprintResponseUnion) ProtoMessage()               {}

type FingerprintGrainResponse struct {
	EncryptedKey     []byte `protobuf:"bytes,10,req,name=encrypted_key,json=encryptedKey" json:"encrypted_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintGrainResponse) Reset()                    { *m = FingerprintGrainResponse{} }
func (m *FingerprintGrainResponse) String() string            { return proto.CompactTextString(m) }
func (*FingerprintGrainResponse) ProtoMessage()               {}

type FingerprintHmacRipemdResponse struct {
	Hmac             []byte `protobuf:"bytes,10,req,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintHmacRipemdResponse) Reset()                    { *m = FingerprintHmacRipemdResponse{} }
func (m *FingerprintHmacRipemdResponse) String() string            { return proto.CompactTextString(m) }
func (*FingerprintHmacRipemdResponse) ProtoMessage()               {}

type PeerTicketUnion struct {
	PublicKey        *PeerTicketPublicKey `protobuf:"bytes,10,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	OldTicket        *PeerTicketOld       `protobuf:"bytes,20,opt,name=old_ticket,json=oldTicket" json:"old_ticket,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PeerTicketUnion) Reset()                    { *m = PeerTicketUnion{} }
func (m *PeerTicketUnion) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketUnion) ProtoMessage()               {}

type PeerTicketPublicKey struct {
	PublicKey        []byte `protobuf:"bytes,10,req,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PeerTicketPublicKey) Reset()                    { *m = PeerTicketPublicKey{} }
func (m *PeerTicketPublicKey) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketPublicKey) ProtoMessage()               {}

type PeerTicketOld struct {
	PeerTicket          []byte `protobuf:"bytes,10,req,name=peer_ticket,json=peerTicket" json:"peer_ticket,omitempty"`
	PeerTicketSignature []byte `protobuf:"bytes,20,req,name=peer_ticket_signature,json=peerTicketSignature" json:"peer_ticket_signature,omitempty"`
	XXX_unrecognized    []byte `json:"-"`
}

func (m *PeerTicketOld) Reset()                    { *m = PeerTicketOld{} }
func (m *PeerTicketOld) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketOld) ProtoMessage()               {}

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

type ClientInfo struct {
	Limited          *bool               `protobuf:"varint,1,opt,name=limited" json:"limited,omitempty"`
	Fb               *ClientInfoFacebook `protobuf:"bytes,2,opt,name=fb" json:"fb,omitempty"`
	Language         *string             `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}

type ClientInfoFacebook struct {
	MachineId        *string `protobuf:"bytes,1,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientInfoFacebook) Reset()                    { *m = ClientInfoFacebook{} }
func (m *ClientInfoFacebook) String() string            { return proto.CompactTextString(m) }
func (*ClientInfoFacebook) ProtoMessage()               {}

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

type AccountInfoFacebook struct {
	AccessToken      *string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	MachineId        *string `protobuf:"bytes,2,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AccountInfoFacebook) Reset()                    { *m = AccountInfoFacebook{} }
func (m *AccountInfoFacebook) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoFacebook) ProtoMessage()               {}
