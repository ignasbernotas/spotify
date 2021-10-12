package pb
import "github.com/golang/protobuf/proto"

type MercuryReply_CachePolicy int32

const (
   MercuryReply_CACHE_NO      MercuryReply_CachePolicy = 1
   MercuryReply_CACHE_PRIVATE MercuryReply_CachePolicy = 2
   MercuryReply_CACHE_PUBLIC  MercuryReply_CachePolicy = 3
)

var MercuryReply_CachePolicy_name = map[int32]string{
   1: "CACHE_NO",
   2: "CACHE_PRIVATE",
   3: "CACHE_PUBLIC",
}

var MercuryReply_CachePolicy_value = map[string]int32{
   "CACHE_NO":      1,
   "CACHE_PRIVATE": 2,
   "CACHE_PUBLIC":  3,
}

func (x MercuryReply_CachePolicy) Enum() *MercuryReply_CachePolicy {
   p := new(MercuryReply_CachePolicy)
   *p = x
   return p
}

func (x MercuryReply_CachePolicy) String() string {
	return proto.EnumName(MercuryReply_CachePolicy_name, int32(x))
}

func (x *MercuryReply_CachePolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MercuryReply_CachePolicy_value, data, "MercuryReply_CachePolicy")
	if err != nil {
		return err
	}
	*x = MercuryReply_CachePolicy(value)
	return nil
}
func (MercuryReply_CachePolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{3, 0} }

type MercuryMultiGetRequest struct {
	Request          []*MercuryRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MercuryMultiGetRequest) Reset()                    { *m = MercuryMultiGetRequest{} }
func (m *MercuryMultiGetRequest) String() string            { return proto.CompactTextString(m) }
func (*MercuryMultiGetRequest) ProtoMessage()               {}
func (*MercuryMultiGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *MercuryMultiGetRequest) GetRequest() []*MercuryRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type MercuryMultiGetReply struct {
	Reply            []*MercuryReply `protobuf:"bytes,1,rep,name=reply" json:"reply,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MercuryMultiGetReply) Reset()                    { *m = MercuryMultiGetReply{} }
func (m *MercuryMultiGetReply) String() string            { return proto.CompactTextString(m) }
func (*MercuryMultiGetReply) ProtoMessage()               {}
func (*MercuryMultiGetReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *MercuryMultiGetReply) GetReply() []*MercuryReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type MercuryRequest struct {
	Uri              *string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType      *string `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body             []byte  `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Etag             []byte  `protobuf:"bytes,4,opt,name=etag" json:"etag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MercuryRequest) Reset()                    { *m = MercuryRequest{} }
func (m *MercuryRequest) String() string            { return proto.CompactTextString(m) }
func (*MercuryRequest) ProtoMessage()               {}
func (*MercuryRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *MercuryRequest) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *MercuryRequest) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *MercuryRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *MercuryRequest) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

type MercuryReply struct {
	StatusCode       *int32                    `protobuf:"zigzag32,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMessage    *string                   `protobuf:"bytes,2,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	CachePolicy      *MercuryReply_CachePolicy `protobuf:"varint,3,opt,name=cache_policy,json=cachePolicy,enum=Spotify.MercuryReply_CachePolicy" json:"cache_policy,omitempty"`
	Ttl              *int32                    `protobuf:"zigzag32,4,opt,name=ttl" json:"ttl,omitempty"`
	Etag             []byte                    `protobuf:"bytes,5,opt,name=etag" json:"etag,omitempty"`
	ContentType      *string                   `protobuf:"bytes,6,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body             []byte                    `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *MercuryReply) Reset()                    { *m = MercuryReply{} }
func (m *MercuryReply) String() string            { return proto.CompactTextString(m) }
func (*MercuryReply) ProtoMessage()               {}
func (*MercuryReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *MercuryReply) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *MercuryReply) GetStatusMessage() string {
	if m != nil && m.StatusMessage != nil {
		return *m.StatusMessage
	}
	return ""
}

func (m *MercuryReply) GetCachePolicy() MercuryReply_CachePolicy {
	if m != nil && m.CachePolicy != nil {
		return *m.CachePolicy
	}
	return MercuryReply_CACHE_NO
}

func (m *MercuryReply) GetTtl() int32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

func (m *MercuryReply) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

func (m *MercuryReply) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *MercuryReply) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	Uri              *string      `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	ContentType      *string      `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Method           *string      `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	StatusCode       *int32       `protobuf:"zigzag32,4,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	UserFields       []*UserField `protobuf:"bytes,6,rep,name=user_fields,json=userFields" json:"user_fields,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *Header) GetUri() string {
	if m != nil && m.Uri != nil {
		return *m.Uri
	}
	return ""
}

func (m *Header) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *Header) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Header) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

func (m *Header) GetUserFields() []*UserField {
	if m != nil {
		return m.UserFields
	}
	return nil
}

type UserField struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserField) Reset()                    { *m = UserField{} }
func (m *UserField) String() string            { return proto.CompactTextString(m) }
func (*UserField) ProtoMessage()               {}
func (*UserField) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *UserField) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *UserField) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MercuryMultiGetRequest)(nil), "Spotify.MercuryMultiGetRequest")
	proto.RegisterType((*MercuryMultiGetReply)(nil), "Spotify.MercuryMultiGetReply")
	proto.RegisterType((*MercuryRequest)(nil), "Spotify.MercuryRequest")
	proto.RegisterType((*MercuryReply)(nil), "Spotify.MercuryReply")
	proto.RegisterType((*Header)(nil), "Spotify.Header")
	proto.RegisterType((*UserField)(nil), "Spotify.UserField")
	proto.RegisterEnum("Spotify.MercuryReply_CachePolicy", MercuryReply_CachePolicy_name, MercuryReply_CachePolicy_value)
}

func init() { proto.RegisterFile("mercury.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x73, 0x25, 0x63, 0x27, 0x72, 0x56, 0xa5, 0xf8, 0x8d, 0xd4, 0x12, 0x52, 0x24, 0xa4,
	0x48, 0xb4, 0x3f, 0x40, 0x31, 0x85, 0x56, 0x10, 0xa8, 0x96, 0x96, 0x57, 0xcb, 0xd8, 0xd3, 0xd6,
	0xaa, 0xd3, 0x35, 0x7b, 0x41, 0xda, 0x6f, 0xe2, 0x13, 0xf8, 0x39, 0xb4, 0xde, 0x75, 0x12, 0x30,
	0xbc, 0xf4, 0xed, 0xcc, 0x99, 0xb3, 0x33, 0x67, 0x66, 0x16, 0xa6, 0x1b, 0xe4, 0xb9, 0xe2, 0x7a,
	0x55, 0x73, 0x26, 0x19, 0x19, 0x7f, 0xa9, 0x99, 0x2c, 0x6f, 0x74, 0xfc, 0x01, 0x0e, 0xd7, 0x36,
	0xb3, 0x56, 0x95, 0x2c, 0xdf, 0xa3, 0xa4, 0xf8, 0x5d, 0xa1, 0x90, 0xe4, 0x15, 0x8c, 0xb9, 0x85,
	0x91, 0xb7, 0xe8, 0x2f, 0xfd, 0xe3, 0x67, 0x2b, 0xf7, 0x68, 0xe5, 0x5e, 0x38, 0x25, 0x6d, 0x75,
	0x71, 0x02, 0x07, 0x9d, 0x62, 0x75, 0xa5, 0xc9, 0x4b, 0x18, 0x72, 0x03, 0x5c, 0xa1, 0xa7, 0xdd,
	0x42, 0x75, 0xa5, 0xa9, 0xd5, 0xc4, 0x1b, 0x98, 0xfd, 0x59, 0x9f, 0x84, 0xd0, 0x57, 0xbc, 0x8c,
	0xbc, 0x85, 0xb7, 0x9c, 0x50, 0x03, 0xc9, 0x11, 0x04, 0x39, 0x7b, 0x90, 0xf8, 0x20, 0x53, 0xa9,
	0x6b, 0x8c, 0x7a, 0x4d, 0xca, 0x77, 0xdc, 0x95, 0xae, 0x91, 0x10, 0x18, 0x7c, 0x63, 0x85, 0x8e,
	0xfa, 0x0b, 0x6f, 0x19, 0xd0, 0x06, 0x1b, 0x0e, 0x65, 0x76, 0x1b, 0x0d, 0x2c, 0x67, 0x70, 0xfc,
	0xab, 0x07, 0xc1, 0xbe, 0x0d, 0xf2, 0x1c, 0x7c, 0x21, 0x33, 0xa9, 0x44, 0x9a, 0xb3, 0x02, 0x9b,
	0xae, 0x73, 0x0a, 0x96, 0x4a, 0x58, 0x81, 0xe4, 0x05, 0xcc, 0x9c, 0x60, 0x83, 0x42, 0x64, 0xb7,
	0x6d, 0xfb, 0xa9, 0x65, 0xd7, 0x96, 0x24, 0x6f, 0x21, 0xc8, 0xb3, 0xfc, 0x0e, 0xd3, 0x9a, 0x55,
	0x65, 0x6e, 0x8d, 0xcc, 0x8e, 0x8f, 0xfe, 0x39, 0xfb, 0x2a, 0x31, 0xca, 0xcb, 0x46, 0x48, 0xfd,
	0x7c, 0x17, 0x98, 0xd9, 0xa5, 0xac, 0x1a, 0xc7, 0x73, 0x6a, 0xe0, 0x76, 0x88, 0xe1, 0x6e, 0x88,
	0xce, 0x3e, 0x46, 0xff, 0xdf, 0xc7, 0x78, 0xb7, 0x8f, 0xf8, 0x35, 0xf8, 0x7b, 0x8d, 0x49, 0x00,
	0x4f, 0x92, 0xd3, 0xe4, 0xfc, 0x2c, 0xfd, 0xf4, 0x39, 0xf4, 0xc8, 0x1c, 0xa6, 0x36, 0xba, 0xa4,
	0x17, 0x5f, 0x4f, 0xaf, 0xce, 0xc2, 0x1e, 0x09, 0x21, 0x70, 0xd4, 0xf5, 0x9b, 0x8f, 0x17, 0x49,
	0xd8, 0x8f, 0x7f, 0x7a, 0x30, 0x3a, 0xc7, 0xac, 0x40, 0xfe, 0xb8, 0x2b, 0x1d, 0xc2, 0x68, 0x83,
	0xf2, 0x8e, 0x15, 0xcd, 0x7a, 0x26, 0xd4, 0x45, 0x7f, 0x1f, 0x61, 0xd0, 0x39, 0xc2, 0x09, 0xf8,
	0x4a, 0x20, 0x4f, 0x6f, 0x4a, 0xac, 0x0a, 0x11, 0x8d, 0x9a, 0x8f, 0x45, 0xb6, 0xcb, 0xbd, 0x16,
	0xc8, 0xdf, 0x99, 0x14, 0x05, 0xd5, 0x42, 0x11, 0x9f, 0xc0, 0x64, 0x9b, 0x30, 0x7e, 0xef, 0x51,
	0xb7, 0x7e, 0xef, 0x51, 0x93, 0x03, 0x18, 0xfe, 0xc8, 0x2a, 0x65, 0x8d, 0x06, 0xd4, 0x06, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x4a, 0xca, 0xa1, 0x3a, 0x03, 0x00, 0x00,
}

type AudioFile struct {
   FileId           []byte            `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
   Format           *AudioFile_Format `protobuf:"varint,2,opt,name=format,enum=Spotify.AudioFile_Format" json:"format,omitempty"`
   XXX_unrecognized []byte            `json:"-"`
}

func (m *AudioFile) Reset()                    { *m = AudioFile{} }

func (m *AudioFile) String() string            { return proto.CompactTextString(m) }

func (*AudioFile) ProtoMessage()               {}

func (*AudioFile) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{15} }

func (m *AudioFile) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *AudioFile) GetFormat() AudioFile_Format {
   if m != nil && m.Format != nil {
      return *m.Format
   }
   return AudioFile_OGG_VORBIS_96
}
