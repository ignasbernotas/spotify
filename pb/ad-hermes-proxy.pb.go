package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Rule struct {
	Type             *string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Times            *uint32 `protobuf:"varint,2,opt,name=times" json:"times,omitempty"`
	Interval         *uint64 `protobuf:"varint,3,opt,name=interval" json:"interval,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Rule) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Rule) GetTimes() uint32 {
	if m != nil && m.Times != nil {
		return *m.Times
	}
	return 0
}

func (m *Rule) GetInterval() uint64 {
	if m != nil && m.Interval != nil {
		return *m.Interval
	}
	return 0
}

type AdRequest struct {
	ClientLanguage   *string  `protobuf:"bytes,1,opt,name=client_language,json=clientLanguage" json:"client_language,omitempty"`
	Product          *string  `protobuf:"bytes,2,opt,name=product" json:"product,omitempty"`
	Version          *uint32  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Type             *string  `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	AvoidAds         []string `protobuf:"bytes,5,rep,name=avoidAds" json:"avoidAds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AdRequest) Reset()                    { *m = AdRequest{} }
func (m *AdRequest) String() string            { return proto.CompactTextString(m) }
func (*AdRequest) ProtoMessage()               {}
func (*AdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AdRequest) GetClientLanguage() string {
	if m != nil && m.ClientLanguage != nil {
		return *m.ClientLanguage
	}
	return ""
}

func (m *AdRequest) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *AdRequest) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *AdRequest) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *AdRequest) GetAvoidAds() []string {
	if m != nil {
		return m.AvoidAds
	}
	return nil
}

type AdQueueResponse struct {
	AdQueueEntry     []*AdQueueEntry `protobuf:"bytes,1,rep,name=adQueueEntry" json:"adQueueEntry,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AdQueueResponse) Reset()                    { *m = AdQueueResponse{} }
func (m *AdQueueResponse) String() string            { return proto.CompactTextString(m) }
func (*AdQueueResponse) ProtoMessage()               {}
func (*AdQueueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AdQueueResponse) GetAdQueueEntry() []*AdQueueEntry {
	if m != nil {
		return m.AdQueueEntry
	}
	return nil
}

type AdFile struct {
	Id               *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Format           *string `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdFile) Reset()                    { *m = AdFile{} }
func (m *AdFile) String() string            { return proto.CompactTextString(m) }
func (*AdFile) ProtoMessage()               {}
func (*AdFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AdFile) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AdFile) GetFormat() string {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return ""
}

type AdQueueEntry struct {
	StartTime        *uint64   `protobuf:"varint,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime          *uint64   `protobuf:"varint,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	Priority         *float64  `protobuf:"fixed64,3,opt,name=priority" json:"priority,omitempty"`
	Token            *string   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	AdVersion        *uint32   `protobuf:"varint,5,opt,name=ad_version,json=adVersion" json:"ad_version,omitempty"`
	Id               *string   `protobuf:"bytes,6,opt,name=id" json:"id,omitempty"`
	Type             *string   `protobuf:"bytes,7,opt,name=type" json:"type,omitempty"`
	Campaign         *string   `protobuf:"bytes,8,opt,name=campaign" json:"campaign,omitempty"`
	Advertiser       *string   `protobuf:"bytes,9,opt,name=advertiser" json:"advertiser,omitempty"`
	Url              *string   `protobuf:"bytes,10,opt,name=url" json:"url,omitempty"`
	Duration         *uint64   `protobuf:"varint,11,opt,name=duration" json:"duration,omitempty"`
	Expiry           *uint64   `protobuf:"varint,12,opt,name=expiry" json:"expiry,omitempty"`
	TrackingUrl      *string   `protobuf:"bytes,13,opt,name=tracking_url,json=trackingUrl" json:"tracking_url,omitempty"`
	BannerType       *string   `protobuf:"bytes,14,opt,name=banner_type,json=bannerType" json:"banner_type,omitempty"`
	Html             *string   `protobuf:"bytes,15,opt,name=html" json:"html,omitempty"`
	Image            *string   `protobuf:"bytes,16,opt,name=image" json:"image,omitempty"`
	BackgroundImage  *string   `protobuf:"bytes,17,opt,name=background_image,json=backgroundImage" json:"background_image,omitempty"`
	BackgroundUrl    *string   `protobuf:"bytes,18,opt,name=background_url,json=backgroundUrl" json:"background_url,omitempty"`
	BackgroundColor  *string   `protobuf:"bytes,19,opt,name=background_color,json=backgroundColor" json:"background_color,omitempty"`
	Title            *string   `protobuf:"bytes,20,opt,name=title" json:"title,omitempty"`
	Caption          *string   `protobuf:"bytes,21,opt,name=caption" json:"caption,omitempty"`
	File             []*AdFile `protobuf:"bytes,22,rep,name=file" json:"file,omitempty"`
	Rule             []*Rule   `protobuf:"bytes,23,rep,name=rule" json:"rule,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *AdQueueEntry) Reset()                    { *m = AdQueueEntry{} }
func (m *AdQueueEntry) String() string            { return proto.CompactTextString(m) }
func (*AdQueueEntry) ProtoMessage()               {}
func (*AdQueueEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AdQueueEntry) GetStartTime() uint64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *AdQueueEntry) GetEndTime() uint64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *AdQueueEntry) GetPriority() float64 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *AdQueueEntry) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *AdQueueEntry) GetAdVersion() uint32 {
	if m != nil && m.AdVersion != nil {
		return *m.AdVersion
	}
	return 0
}

func (m *AdQueueEntry) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AdQueueEntry) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *AdQueueEntry) GetCampaign() string {
	if m != nil && m.Campaign != nil {
		return *m.Campaign
	}
	return ""
}

func (m *AdQueueEntry) GetAdvertiser() string {
	if m != nil && m.Advertiser != nil {
		return *m.Advertiser
	}
	return ""
}

func (m *AdQueueEntry) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *AdQueueEntry) GetDuration() uint64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *AdQueueEntry) GetExpiry() uint64 {
	if m != nil && m.Expiry != nil {
		return *m.Expiry
	}
	return 0
}

func (m *AdQueueEntry) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *AdQueueEntry) GetBannerType() string {
	if m != nil && m.BannerType != nil {
		return *m.BannerType
	}
	return ""
}

func (m *AdQueueEntry) GetHtml() string {
	if m != nil && m.Html != nil {
		return *m.Html
	}
	return ""
}

func (m *AdQueueEntry) GetImage() string {
	if m != nil && m.Image != nil {
		return *m.Image
	}
	return ""
}

func (m *AdQueueEntry) GetBackgroundImage() string {
	if m != nil && m.BackgroundImage != nil {
		return *m.BackgroundImage
	}
	return ""
}

func (m *AdQueueEntry) GetBackgroundUrl() string {
	if m != nil && m.BackgroundUrl != nil {
		return *m.BackgroundUrl
	}
	return ""
}

func (m *AdQueueEntry) GetBackgroundColor() string {
	if m != nil && m.BackgroundColor != nil {
		return *m.BackgroundColor
	}
	return ""
}

func (m *AdQueueEntry) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *AdQueueEntry) GetCaption() string {
	if m != nil && m.Caption != nil {
		return *m.Caption
	}
	return ""
}

func (m *AdQueueEntry) GetFile() []*AdFile {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *AdQueueEntry) GetRule() []*Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "Spotify.Rule")
	proto.RegisterType((*AdRequest)(nil), "Spotify.AdRequest")
	proto.RegisterType((*AdQueueResponse)(nil), "Spotify.AdQueueResponse")
	proto.RegisterType((*AdFile)(nil), "Spotify.AdFile")
	proto.RegisterType((*AdQueueEntry)(nil), "Spotify.AdQueueEntry")
}

func init() { proto.RegisterFile("ad-hermes-proxy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0xda, 0x74, 0x6d, 0x6e, 0xff, 0xed, 0xe7, 0xdf, 0x36, 0xcc, 0x24, 0xa0, 0x2b, 0x42,
	0x94, 0x87, 0x4d, 0x88, 0x37, 0x1e, 0x2b, 0x04, 0x12, 0xd2, 0x5e, 0x30, 0x83, 0xd7, 0xca, 0x8b,
	0xef, 0x3a, 0x6b, 0x89, 0x1d, 0x1c, 0xa7, 0x5a, 0xbf, 0x09, 0xcf, 0x7c, 0x52, 0x64, 0x3b, 0xc9,
	0x52, 0xde, 0x7c, 0xce, 0xb9, 0x3a, 0xbe, 0xf7, 0xf8, 0x1a, 0x4e, 0xb9, 0xb8, 0xbc, 0x47, 0x93,
	0x63, 0x79, 0x59, 0x18, 0xfd, 0xb8, 0xbf, 0x2a, 0x8c, 0xb6, 0x9a, 0x0c, 0xbf, 0x17, 0xda, 0xca,
	0xbb, 0xfd, 0xf2, 0x1a, 0x62, 0x56, 0x65, 0x48, 0x08, 0xc4, 0x76, 0x5f, 0x20, 0x8d, 0x16, 0xd1,
	0x2a, 0x61, 0xfe, 0x4c, 0x4e, 0x60, 0x60, 0x65, 0x8e, 0x25, 0xed, 0x2d, 0xa2, 0xd5, 0x94, 0x05,
	0x40, 0xce, 0x61, 0x24, 0x95, 0x45, 0xb3, 0xe3, 0x19, 0xed, 0x2f, 0xa2, 0x55, 0xcc, 0x5a, 0xbc,
	0xfc, 0x1d, 0x41, 0xb2, 0x16, 0x0c, 0x7f, 0x55, 0x58, 0x5a, 0xf2, 0x16, 0xe6, 0x69, 0x26, 0x51,
	0xd9, 0x4d, 0xc6, 0xd5, 0xb6, 0xe2, 0xdb, 0xc6, 0x7e, 0x16, 0xe8, 0xeb, 0x9a, 0x25, 0x14, 0x86,
	0x85, 0xd1, 0xa2, 0x4a, 0xad, 0xbf, 0x2a, 0x61, 0x0d, 0x74, 0xca, 0x0e, 0x4d, 0x29, 0xb5, 0xf2,
	0x77, 0x4d, 0x59, 0x03, 0xdb, 0x86, 0xe3, 0x4e, 0xc3, 0xe7, 0x30, 0xe2, 0x3b, 0x2d, 0xc5, 0x5a,
	0x94, 0x74, 0xb0, 0xe8, 0xaf, 0x12, 0xd6, 0xe2, 0xe5, 0x35, 0xcc, 0xd7, 0xe2, 0x5b, 0x85, 0x15,
	0x32, 0x2c, 0x0b, 0xad, 0x4a, 0x24, 0x1f, 0x61, 0xc2, 0x03, 0xf5, 0x59, 0x59, 0xb3, 0xa7, 0xd1,
	0xa2, 0xbf, 0x1a, 0x7f, 0x38, 0xbd, 0xaa, 0xb3, 0xb9, 0x5a, 0x77, 0x44, 0x76, 0x50, 0xba, 0x7c,
	0x0f, 0x47, 0x6b, 0xf1, 0x45, 0x66, 0x48, 0x66, 0xd0, 0x93, 0xa2, 0x9e, 0xab, 0x27, 0x05, 0x39,
	0x83, 0xa3, 0x3b, 0x6d, 0x72, 0xde, 0x8c, 0x52, 0xa3, 0xe5, 0x9f, 0x01, 0x4c, 0xba, 0x86, 0xe4,
	0x05, 0x40, 0x69, 0xb9, 0xb1, 0x1b, 0x17, 0xab, 0x37, 0x88, 0x59, 0xe2, 0x99, 0x1b, 0x99, 0x23,
	0x79, 0x0e, 0x23, 0x54, 0x22, 0x88, 0x3d, 0x2f, 0x0e, 0x51, 0x09, 0x2f, 0x9d, 0xc3, 0xa8, 0x30,
	0x52, 0x1b, 0x69, 0xf7, 0x3e, 0x95, 0x88, 0xb5, 0xd8, 0xbf, 0x99, 0x7e, 0x40, 0x55, 0xe7, 0x12,
	0x80, 0xbb, 0x8b, 0x8b, 0x4d, 0x93, 0xe4, 0xc0, 0x27, 0x99, 0x70, 0xf1, 0xb3, 0xce, 0x32, 0xcc,
	0x70, 0xd4, 0xce, 0xd0, 0x64, 0x3b, 0x3c, 0xcc, 0x36, 0xe5, 0x79, 0xc1, 0xe5, 0x56, 0xd1, 0x91,
	0xe7, 0x5b, 0x4c, 0x5e, 0x3a, 0xfb, 0x1d, 0x1a, 0x2b, 0x4b, 0x34, 0x34, 0xf1, 0x6a, 0x87, 0x21,
	0xc7, 0xd0, 0xaf, 0x4c, 0x46, 0xc1, 0x0b, 0xee, 0xe8, 0xdc, 0x44, 0x65, 0xb8, 0x75, 0xed, 0x8c,
	0xc3, 0x12, 0x35, 0xd8, 0x25, 0x88, 0x8f, 0x85, 0x34, 0x7b, 0x3a, 0xf1, 0x4a, 0x8d, 0xc8, 0x05,
	0x4c, 0xac, 0xe1, 0xe9, 0x83, 0x54, 0xdb, 0x8d, 0xb3, 0x9b, 0x7a, 0xbb, 0x71, 0xc3, 0xfd, 0x30,
	0x19, 0x79, 0x05, 0xe3, 0x5b, 0xae, 0x14, 0x9a, 0x8d, 0xef, 0x7f, 0x16, 0x3a, 0x09, 0xd4, 0x8d,
	0x9b, 0x82, 0x40, 0x7c, 0x6f, 0xf3, 0x8c, 0xce, 0xc3, 0x64, 0xee, 0xec, 0x22, 0x93, 0xb9, 0x5b,
	0xce, 0xe3, 0x10, 0x99, 0x07, 0xe4, 0x1d, 0x1c, 0xdf, 0xf2, 0xf4, 0x61, 0x6b, 0x74, 0xa5, 0xc4,
	0x26, 0x14, 0xfc, 0xe7, 0x0b, 0xe6, 0x4f, 0xfc, 0x57, 0x5f, 0xfa, 0x06, 0x66, 0x9d, 0x52, 0xd7,
	0x1a, 0xf1, 0x85, 0xd3, 0x27, 0xd6, 0x35, 0x77, 0xe8, 0x98, 0xea, 0x4c, 0x1b, 0xfa, 0xff, 0xbf,
	0x8e, 0x9f, 0x1c, 0x1d, 0x7e, 0x9e, 0xcd, 0x90, 0x9e, 0xd4, 0xaf, 0xe8, 0x80, 0xfb, 0x0c, 0x29,
	0x2f, 0x7c, 0x66, 0xa7, 0xe1, 0x9b, 0xd4, 0x90, 0xbc, 0x86, 0xf8, 0x4e, 0x66, 0x48, 0xcf, 0xfc,
	0x06, 0xcf, 0x3b, 0x1b, 0xec, 0x76, 0x94, 0x79, 0x91, 0x5c, 0x40, 0x6c, 0xaa, 0x0c, 0xe9, 0x33,
	0x5f, 0x34, 0x6d, 0x8b, 0xdc, 0xff, 0x67, 0x5e, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x35,
	0xe8, 0x24, 0x2e, 0x04, 0x00, 0x00,
}