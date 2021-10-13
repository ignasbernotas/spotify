package pb

import (
   "github.com/golang/protobuf/proto"
)

type Header struct {
   ContentType      *string      `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
   Method           *string      `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
   StatusCode       *int32       `protobuf:"zigzag32,4,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
   Uri              *string      `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
   UserFields       []*UserField `protobuf:"bytes,6,rep,name=user_fields,json=userFields" json:"user_fields,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}

func (m *Header) GetStatusCode() int32 {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return 0
}

type UserField struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UserField) Reset()                    { *m = UserField{} }
func (m *UserField) String() string            { return proto.CompactTextString(m) }
func (*UserField) ProtoMessage()               {}
