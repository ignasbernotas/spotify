package pb
import "github.com/golang/protobuf/proto"

type Album_Type int32

const (
	Album_ALBUM       Album_Type = 1
	Album_SINGLE      Album_Type = 2
	Album_COMPILATION Album_Type = 3
	Album_EP          Album_Type = 4
)

var Album_Type_name = map[int32]string{
	1: "ALBUM",
	2: "SINGLE",
	3: "COMPILATION",
	4: "EP",
}
var Album_Type_value = map[string]int32{
	"ALBUM":       1,
	"SINGLE":      2,
	"COMPILATION": 3,
	"EP":          4,
}

func (x Album_Type) Enum() *Album_Type {
	p := new(Album_Type)
	*p = x
	return p
}
func (x Album_Type) String() string {
	return proto.EnumName(Album_Type_name, int32(x))
}
func (x *Album_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Album_Type_value, data, "Album_Type")
	if err != nil {
		return err
	}
	*x = Album_Type(value)
	return nil
}

type Image_Size int32

const (
	Image_DEFAULT Image_Size = 0
	Image_SMALL   Image_Size = 1
	Image_LARGE   Image_Size = 2
	Image_XLARGE  Image_Size = 3
)

var Image_Size_name = map[int32]string{
	0: "DEFAULT",
	1: "SMALL",
	2: "LARGE",
	3: "XLARGE",
}
var Image_Size_value = map[string]int32{
	"DEFAULT": 0,
	"SMALL":   1,
	"LARGE":   2,
	"XLARGE":  3,
}

func (x Image_Size) Enum() *Image_Size {
	p := new(Image_Size)
	*p = x
	return p
}
func (x Image_Size) String() string {
	return proto.EnumName(Image_Size_name, int32(x))
}
func (x *Image_Size) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Image_Size_value, data, "Image_Size")
	if err != nil {
		return err
	}
	*x = Image_Size(value)
	return nil
}

type Copyright_Type int32

const (
	Copyright_P Copyright_Type = 0
	Copyright_C Copyright_Type = 1
)

var Copyright_Type_name = map[int32]string{
	0: "P",
	1: "C",
}
var Copyright_Type_value = map[string]int32{
	"P": 0,
	"C": 1,
}

func (x Copyright_Type) Enum() *Copyright_Type {
	p := new(Copyright_Type)
	*p = x
	return p
}
func (x Copyright_Type) String() string {
	return proto.EnumName(Copyright_Type_name, int32(x))
}
func (x *Copyright_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Copyright_Type_value, data, "Copyright_Type")
	if err != nil {
		return err
	}
	*x = Copyright_Type(value)
	return nil
}

type Restriction_Type int32

const (
	Restriction_STREAMING Restriction_Type = 0
)

var Restriction_Type_name = map[int32]string{
	0: "STREAMING",
}
var Restriction_Type_value = map[string]int32{
	"STREAMING": 0,
}

func (x Restriction_Type) Enum() *Restriction_Type {
	p := new(Restriction_Type)
	*p = x
	return p
}
func (x Restriction_Type) String() string {
	return proto.EnumName(Restriction_Type_name, int32(x))
}
func (x *Restriction_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Restriction_Type_value, data, "Restriction_Type")
	if err != nil {
		return err
	}
	*x = Restriction_Type(value)
	return nil
}

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

var AudioFile_Format_name = map[int32]string{
	0:  "OGG_VORBIS_96",
	1:  "OGG_VORBIS_160",
	2:  "OGG_VORBIS_320",
	3:  "MP3_256",
	4:  "MP3_320",
	5:  "MP3_160",
	6:  "MP3_96",
	7:  "MP3_160_ENC",
	8:  "OTHER2",
	9:  "OTHER3",
	10: "AAC_160",
	11: "AAC_320",
	12: "OTHER4",
	13: "OTHER5",
}
var AudioFile_Format_value = map[string]int32{
	"OGG_VORBIS_96":  0,
	"OGG_VORBIS_160": 1,
	"OGG_VORBIS_320": 2,
	"MP3_256":        3,
	"MP3_320":        4,
	"MP3_160":        5,
	"MP3_96":         6,
	"MP3_160_ENC":    7,
	"OTHER2":         8,
	"OTHER3":         9,
	"AAC_160":        10,
	"AAC_320":        11,
	"OTHER4":         12,
	"OTHER5":         13,
}

func (x AudioFile_Format) Enum() *AudioFile_Format {
	p := new(AudioFile_Format)
	*p = x
	return p
}
func (x AudioFile_Format) String() string {
	return proto.EnumName(AudioFile_Format_name, int32(x))
}
func (x *AudioFile_Format) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AudioFile_Format_value, data, "AudioFile_Format")
	if err != nil {
		return err
	}
	*x = AudioFile_Format(value)
	return nil
}

type TopTracks struct {
	Country          *string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	Track            []*Track `protobuf:"bytes,2,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TopTracks) Reset()                    { *m = TopTracks{} }
func (m *TopTracks) String() string            { return proto.CompactTextString(m) }
func (*TopTracks) ProtoMessage()               {}

type ActivityPeriod struct {
	StartYear        *int32 `protobuf:"zigzag32,1,opt,name=start_year,json=startYear" json:"start_year,omitempty"`
	EndYear          *int32 `protobuf:"zigzag32,2,opt,name=end_year,json=endYear" json:"end_year,omitempty"`
	Decade           *int32 `protobuf:"zigzag32,3,opt,name=decade" json:"decade,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActivityPeriod) Reset()                    { *m = ActivityPeriod{} }
func (m *ActivityPeriod) String() string            { return proto.CompactTextString(m) }
func (*ActivityPeriod) ProtoMessage()               {}

type Artist struct {
	Gid                  []byte            `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name                 *string           `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Popularity           *int32            `protobuf:"zigzag32,3,opt,name=popularity" json:"popularity,omitempty"`
	TopTrack             []*TopTracks      `protobuf:"bytes,4,rep,name=top_track,json=topTrack" json:"top_track,omitempty"`
	AlbumGroup           []*AlbumGroup     `protobuf:"bytes,5,rep,name=album_group,json=albumGroup" json:"album_group,omitempty"`
	SingleGroup          []*AlbumGroup     `protobuf:"bytes,6,rep,name=single_group,json=singleGroup" json:"single_group,omitempty"`
	CompilationGroup     []*AlbumGroup     `protobuf:"bytes,7,rep,name=compilation_group,json=compilationGroup" json:"compilation_group,omitempty"`
	AppearsOnGroup       []*AlbumGroup     `protobuf:"bytes,8,rep,name=appears_on_group,json=appearsOnGroup" json:"appears_on_group,omitempty"`
	Genre                []string          `protobuf:"bytes,9,rep,name=genre" json:"genre,omitempty"`
	ExternalId           []*ExternalId     `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Portrait             []*Image          `protobuf:"bytes,11,rep,name=portrait" json:"portrait,omitempty"`
	Biography            []*Biography      `protobuf:"bytes,12,rep,name=biography" json:"biography,omitempty"`
	ActivityPeriod       []*ActivityPeriod `protobuf:"bytes,13,rep,name=activity_period,json=activityPeriod" json:"activity_period,omitempty"`
	Restriction          []*Restriction    `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related              []*Artist         `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	IsPortraitAlbumCover *bool             `protobuf:"varint,16,opt,name=is_portrait_album_cover,json=isPortraitAlbumCover" json:"is_portrait_album_cover,omitempty"`
	PortraitGroup        *ImageGroup       `protobuf:"bytes,17,opt,name=portrait_group,json=portraitGroup" json:"portrait_group,omitempty"`
	XXX_unrecognized     []byte            `json:"-"`
}

func (*Artist) ProtoMessage()               {}
func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }

type AlbumGroup struct {
	Album            []*Album `protobuf:"bytes,1,rep,name=album" json:"album,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AlbumGroup) Reset()                    { *m = AlbumGroup{} }
func (m *AlbumGroup) String() string            { return proto.CompactTextString(m) }
func (*AlbumGroup) ProtoMessage()               {}

type Date struct {
	Year             *int32 `protobuf:"zigzag32,1,opt,name=year" json:"year,omitempty"`
	Month            *int32 `protobuf:"zigzag32,2,opt,name=month" json:"month,omitempty"`
	Day              *int32 `protobuf:"zigzag32,3,opt,name=day" json:"day,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}

type Album struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,3,rep,name=artist" json:"artist,omitempty"`
	Typ              *Album_Type    `protobuf:"varint,4,opt,name=typ,enum=Spotify.Album_Type" json:"typ,omitempty"`
	Label            *string        `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	Date             *Date          `protobuf:"bytes,6,opt,name=date" json:"date,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,7,opt,name=popularity" json:"popularity,omitempty"`
	Genre            []string       `protobuf:"bytes,8,rep,name=genre" json:"genre,omitempty"`
	Cover            []*Image       `protobuf:"bytes,9,rep,name=cover" json:"cover,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Disc             []*Disc        `protobuf:"bytes,11,rep,name=disc" json:"disc,omitempty"`
	Review           []string       `protobuf:"bytes,12,rep,name=review" json:"review,omitempty"`
	Copyright        []*Copyright   `protobuf:"bytes,13,rep,name=copyright" json:"copyright,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related          []*Album       `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,16,rep,name=sale_period,json=salePeriod" json:"sale_period,omitempty"`
	CoverGroup       *ImageGroup    `protobuf:"bytes,17,opt,name=cover_group,json=coverGroup" json:"cover_group,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Album) Reset()                    { *m = Album{} }
func (m *Album) String() string            { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()               {}

type Track struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Album            *Album         `protobuf:"bytes,3,opt,name=album" json:"album,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,4,rep,name=artist" json:"artist,omitempty"`
	Number           *int32         `protobuf:"zigzag32,5,opt,name=number" json:"number,omitempty"`
	DiscNumber       *int32         `protobuf:"zigzag32,6,opt,name=disc_number,json=discNumber" json:"disc_number,omitempty"`
	Duration         *int32         `protobuf:"zigzag32,7,opt,name=duration" json:"duration,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,8,opt,name=popularity" json:"popularity,omitempty"`
	Explicit         *bool          `protobuf:"varint,9,opt,name=explicit" json:"explicit,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,11,rep,name=restriction" json:"restriction,omitempty"`
	File             []*AudioFile   `protobuf:"bytes,12,rep,name=file" json:"file,omitempty"`
	Alternative      []*Track       `protobuf:"bytes,13,rep,name=alternative" json:"alternative,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,14,rep,name=sale_period,json=salePeriod" json:"sale_period,omitempty"`
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

type Image struct {
	FileId           []byte      `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Size             *Image_Size `protobuf:"varint,2,opt,name=size,enum=Spotify.Image_Size" json:"size,omitempty"`
	Width            *int32      `protobuf:"zigzag32,3,opt,name=width" json:"width,omitempty"`
	Height           *int32      `protobuf:"zigzag32,4,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}

type ImageGroup struct {
	Image            []*Image `protobuf:"bytes,1,rep,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ImageGroup) Reset()                    { *m = ImageGroup{} }
func (m *ImageGroup) String() string            { return proto.CompactTextString(m) }
func (*ImageGroup) ProtoMessage()               {}

type Biography struct {
	Text             *string       `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Portrait         []*Image      `protobuf:"bytes,2,rep,name=portrait" json:"portrait,omitempty"`
	PortraitGroup    []*ImageGroup `protobuf:"bytes,3,rep,name=portrait_group,json=portraitGroup" json:"portrait_group,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Biography) Reset()                    { *m = Biography{} }
func (m *Biography) String() string            { return proto.CompactTextString(m) }
func (*Biography) ProtoMessage()               {}

type Disc struct {
	Number           *int32   `protobuf:"zigzag32,1,opt,name=number" json:"number,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Track            []*Track `protobuf:"bytes,3,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Disc) Reset()                    { *m = Disc{} }
func (m *Disc) String() string            { return proto.CompactTextString(m) }
func (*Disc) ProtoMessage()               {}

type Copyright struct {
	Typ              *Copyright_Type `protobuf:"varint,1,opt,name=typ,enum=Spotify.Copyright_Type" json:"typ,omitempty"`
	Text             *string         `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Copyright) Reset()                    { *m = Copyright{} }
func (m *Copyright) String() string            { return proto.CompactTextString(m) }
func (*Copyright) ProtoMessage()               {}

type Restriction struct {
	CountriesAllowed   *string           `protobuf:"bytes,2,opt,name=countries_allowed,json=countriesAllowed" json:"countries_allowed,omitempty"`
	CountriesForbidden *string           `protobuf:"bytes,3,opt,name=countries_forbidden,json=countriesForbidden" json:"countries_forbidden,omitempty"`
	Typ                *Restriction_Type `protobuf:"varint,4,opt,name=typ,enum=Spotify.Restriction_Type" json:"typ,omitempty"`
	CatalogueStr       []string          `protobuf:"bytes,5,rep,name=catalogue_str,json=catalogueStr" json:"catalogue_str,omitempty"`
	XXX_unrecognized   []byte            `json:"-"`
}

func (m *Restriction) Reset()                    { *m = Restriction{} }
func (m *Restriction) String() string            { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()               {}

type SalePeriod struct {
	Restriction      []*Restriction `protobuf:"bytes,1,rep,name=restriction" json:"restriction,omitempty"`
	Start            *Date          `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End              *Date          `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SalePeriod) Reset()                    { *m = SalePeriod{} }
func (m *SalePeriod) String() string            { return proto.CompactTextString(m) }
func (*SalePeriod) ProtoMessage()               {}

type ExternalId struct {
	Typ              *string `protobuf:"bytes,1,opt,name=typ" json:"typ,omitempty"`
	Id               *string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalId) Reset()                    { *m = ExternalId{} }
func (m *ExternalId) String() string            { return proto.CompactTextString(m) }
func (*ExternalId) ProtoMessage()               {}

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

func init() {
   proto.RegisterType((*MercuryMultiGetRequest)(nil), "Spotify.MercuryMultiGetRequest")
   proto.RegisterType((*MercuryMultiGetReply)(nil), "Spotify.MercuryMultiGetReply")
   proto.RegisterType((*MercuryRequest)(nil), "Spotify.MercuryRequest")
   proto.RegisterType((*MercuryReply)(nil), "Spotify.MercuryReply")
   proto.RegisterEnum("Spotify.MercuryReply_CachePolicy", MercuryReply_CachePolicy_name, MercuryReply_CachePolicy_value)
}

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
