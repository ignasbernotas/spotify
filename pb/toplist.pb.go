package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Toplist struct {
	Items            []string `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Toplist) Reset()                    { *m = Toplist{} }
func (m *Toplist) String() string            { return proto.CompactTextString(m) }
func (*Toplist) ProtoMessage()               {}
func (*Toplist) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *Toplist) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Toplist)(nil), "Spotify.Toplist")
}

func init() { proto.RegisterFile("toplist.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xc9, 0x2f, 0xc8,
	0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x0f, 0x2e, 0xc8, 0x2f, 0xc9,
	0x4c, 0xab, 0x54, 0x92, 0xe7, 0x62, 0x0f, 0x81, 0xc8, 0x08, 0x89, 0x70, 0xb1, 0x66, 0x96, 0xa4,
	0xe6, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x41, 0x38, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x91, 0x14, 0x78, 0x48, 0x39, 0x00, 0x00, 0x00,
}