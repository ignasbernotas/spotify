package pb

import (
   "github.com/golang/protobuf/proto"
)

type AudioFile_Format int32

const (
	AudioFile_OGG_VORBIS_96  AudioFile_Format = 0
	AudioFile_OGG_VORBIS_160 AudioFile_Format = 1
	AudioFile_OGG_VORBIS_320 AudioFile_Format = 2
	AudioFile_MP3_256        AudioFile_Format = 3
	AudioFile_MP3_320        AudioFile_Format = 4
	AudioFile_MP3_160        AudioFile_Format = 5
	AudioFile_MP3_96         AudioFile_Format = 6
	AudioFile_MP3_160_ENC    AudioFile_Format = 7
	AudioFile_OTHER2         AudioFile_Format = 8
	AudioFile_OTHER3         AudioFile_Format = 9
	AudioFile_AAC_160        AudioFile_Format = 10
	AudioFile_AAC_320        AudioFile_Format = 11
	AudioFile_OTHER4         AudioFile_Format = 12
	AudioFile_OTHER5         AudioFile_Format = 13
)

type Track struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Number           *int32         `protobuf:"zigzag32,5,opt,name=number" json:"number,omitempty"`
	DiscNumber       *int32         `protobuf:"zigzag32,6,opt,name=disc_number,json=discNumber" json:"disc_number,omitempty"`
	Duration         *int32         `protobuf:"zigzag32,7,opt,name=duration" json:"duration,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,8,opt,name=popularity" json:"popularity,omitempty"`
	Explicit         *bool          `protobuf:"varint,9,opt,name=explicit" json:"explicit,omitempty"`
	File             []*AudioFile   `protobuf:"bytes,12,rep,name=file" json:"file,omitempty"`
	Alternative      []*Track       `protobuf:"bytes,13,rep,name=alternative" json:"alternative,omitempty"`
	Preview          []*AudioFile   `protobuf:"bytes,15,rep,name=preview" json:"preview,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Track) Reset()                    { *m = Track{} }
func (m *Track) String() string            { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()               {}

func (m *Track) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Track) GetFile() []*AudioFile {
	if m != nil {
		return m.File
	}
	return nil
}

type MercuryMultiGetRequest struct {
	Request          []*MercuryRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MercuryMultiGetRequest) Reset()                    { *m = MercuryMultiGetRequest{} }
func (m *MercuryMultiGetRequest) String() string            { return proto.CompactTextString(m) }
func (*MercuryMultiGetRequest) ProtoMessage()               {}

type MercuryMultiGetReply struct {
	Reply            []*MercuryReply `protobuf:"bytes,1,rep,name=reply" json:"reply,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MercuryMultiGetReply) Reset()                    { *m = MercuryMultiGetReply{} }
func (m *MercuryMultiGetReply) String() string            { return proto.CompactTextString(m) }
func (*MercuryMultiGetReply) ProtoMessage()               {}

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

type MercuryReply struct {
	StatusCode       *int32                    `protobuf:"zigzag32,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMessage    *string                   `protobuf:"bytes,2,opt,name=status_message,json=statusMessage" json:"status_message,omitempty"`
	Ttl              *int32                    `protobuf:"zigzag32,4,opt,name=ttl" json:"ttl,omitempty"`
	Etag             []byte                    `protobuf:"bytes,5,opt,name=etag" json:"etag,omitempty"`
	ContentType      *string                   `protobuf:"bytes,6,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Body             []byte                    `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *MercuryReply) Reset()                    { *m = MercuryReply{} }
func (m *MercuryReply) String() string            { return proto.CompactTextString(m) }
func (*MercuryReply) ProtoMessage()               {}

type AudioFile struct {
   FileId           []byte            `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
   Format           *AudioFile_Format `protobuf:"varint,2,opt,name=format,enum=Spotify.AudioFile_Format" json:"format,omitempty"`
   XXX_unrecognized []byte            `json:"-"`
}

func (m *AudioFile) Reset()                    { *m = AudioFile{} }
func (m *AudioFile) String() string            { return proto.CompactTextString(m) }
func (*AudioFile) ProtoMessage()               {}

func (m *AudioFile) GetFormat() AudioFile_Format {
   if m != nil && m.Format != nil {
      return *m.Format
   }
   return AudioFile_OGG_VORBIS_96
}
