package pb

import (
   "github.com/golang/protobuf/proto"
)

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
