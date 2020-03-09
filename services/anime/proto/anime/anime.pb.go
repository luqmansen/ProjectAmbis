// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/anime/anime.proto

package hanako_srv_anime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	AnimeId              string   `protobuf:"bytes,1,opt,name=animeId,proto3" json:"animeId,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c518ae8dab0f146, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAnimeId() string {
	if m != nil {
		return m.AnimeId
	}
	return ""
}

func (m *Request) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Anime struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Sources              []string `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Episodes             int32    `protobuf:"varint,5,opt,name=episodes,proto3" json:"episodes,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Picture              string   `protobuf:"bytes,7,opt,name=picture,proto3" json:"picture,omitempty"`
	Thumbnail            string   `protobuf:"bytes,8,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Synonyms             []string `protobuf:"bytes,9,rep,name=synonyms,proto3" json:"synonyms,omitempty"`
	Relations            []string `protobuf:"bytes,10,rep,name=relations,proto3" json:"relations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Anime) Reset()         { *m = Anime{} }
func (m *Anime) String() string { return proto.CompactTextString(m) }
func (*Anime) ProtoMessage()    {}
func (*Anime) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c518ae8dab0f146, []int{1}
}

func (m *Anime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Anime.Unmarshal(m, b)
}
func (m *Anime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Anime.Marshal(b, m, deterministic)
}
func (m *Anime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Anime.Merge(m, src)
}
func (m *Anime) XXX_Size() int {
	return xxx_messageInfo_Anime.Size(m)
}
func (m *Anime) XXX_DiscardUnknown() {
	xxx_messageInfo_Anime.DiscardUnknown(m)
}

var xxx_messageInfo_Anime proto.InternalMessageInfo

func (m *Anime) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Anime) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Anime) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Anime) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Anime) GetEpisodes() int32 {
	if m != nil {
		return m.Episodes
	}
	return 0
}

func (m *Anime) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Anime) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *Anime) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *Anime) GetSynonyms() []string {
	if m != nil {
		return m.Synonyms
	}
	return nil
}

func (m *Anime) GetRelations() []string {
	if m != nil {
		return m.Relations
	}
	return nil
}

type Results struct {
	Animes               []*Anime `protobuf:"bytes,1,rep,name=animes,proto3" json:"animes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Results) Reset()         { *m = Results{} }
func (m *Results) String() string { return proto.CompactTextString(m) }
func (*Results) ProtoMessage()    {}
func (*Results) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c518ae8dab0f146, []int{2}
}

func (m *Results) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Results.Unmarshal(m, b)
}
func (m *Results) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Results.Marshal(b, m, deterministic)
}
func (m *Results) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Results.Merge(m, src)
}
func (m *Results) XXX_Size() int {
	return xxx_messageInfo_Results.Size(m)
}
func (m *Results) XXX_DiscardUnknown() {
	xxx_messageInfo_Results.DiscardUnknown(m)
}

var xxx_messageInfo_Results proto.InternalMessageInfo

func (m *Results) GetAnimes() []*Anime {
	if m != nil {
		return m.Animes
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "hanako.srv.anime.Request")
	proto.RegisterType((*Anime)(nil), "hanako.srv.anime.Anime")
	proto.RegisterType((*Results)(nil), "hanako.srv.anime.Results")
}

func init() { proto.RegisterFile("proto/anime/anime.proto", fileDescriptor_2c518ae8dab0f146) }

var fileDescriptor_2c518ae8dab0f146 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x05, 0x0a, 0x2d, 0xa3, 0x31, 0x66, 0x62, 0xec, 0xda, 0x78, 0x20, 0x9c, 0x7a, 0xa2,
	0x49, 0xbd, 0x79, 0xd2, 0x78, 0x30, 0x3d, 0x78, 0xc1, 0x27, 0xa0, 0xed, 0x24, 0xdd, 0x48, 0x59,
	0x64, 0x96, 0x26, 0x3d, 0xfa, 0x10, 0xbe, 0xaf, 0xd9, 0x29, 0x96, 0x44, 0x7b, 0xf2, 0x42, 0xf8,
	0xbf, 0x61, 0xbe, 0x5d, 0x66, 0x17, 0xc6, 0x75, 0x63, 0xac, 0x99, 0x15, 0x95, 0xde, 0xd2, 0xe1,
	0x99, 0x09, 0xc1, 0xab, 0x4d, 0x51, 0x15, 0xef, 0x26, 0xe3, 0x66, 0x97, 0x09, 0x4f, 0x5f, 0x61,
	0x98, 0xd3, 0x47, 0x4b, 0x6c, 0x51, 0xc1, 0x50, 0xd8, 0x62, 0xad, 0xbc, 0xc4, 0x9b, 0xc6, 0xf9,
	0x4f, 0xc4, 0x6b, 0x08, 0xad, 0xb6, 0x25, 0x29, 0x5f, 0xf8, 0x21, 0x20, 0xc2, 0xc0, 0xee, 0x6b,
	0x52, 0x81, 0x40, 0x79, 0x4f, 0x3f, 0x7d, 0x08, 0x9f, 0x5c, 0x17, 0x5e, 0x82, 0x7f, 0x14, 0xf9,
	0x8b, 0xb5, 0xb3, 0xb3, 0x69, 0x9b, 0x15, 0xb1, 0xf2, 0x93, 0xc0, 0xd9, 0xbb, 0xd8, 0xdb, 0x83,
	0x53, 0xf6, 0x41, 0x6f, 0xc7, 0x09, 0x8c, 0xa8, 0xd6, 0x6c, 0xd6, 0xc4, 0x2a, 0x4c, 0xbc, 0x69,
	0x98, 0x1f, 0x33, 0xde, 0x40, 0xc4, 0xb6, 0xb0, 0x2d, 0xab, 0x48, 0x3a, 0xba, 0xe4, 0xd6, 0xad,
	0xf5, 0xca, 0xb6, 0x0d, 0xa9, 0xe1, 0xe1, 0xaf, 0xba, 0x88, 0x77, 0x10, 0xdb, 0x4d, 0xbb, 0x5d,
	0x56, 0x85, 0x2e, 0xd5, 0x48, 0x6a, 0x3d, 0x70, 0x6b, 0xf1, 0xbe, 0x32, 0xd5, 0x7e, 0xcb, 0x2a,
	0x96, 0x0d, 0x1f, 0xb3, 0xeb, 0x6c, 0xa8, 0x2c, 0xac, 0x36, 0x15, 0x2b, 0x90, 0x62, 0x0f, 0xd2,
	0x07, 0x37, 0x52, 0x6e, 0x4b, 0xcb, 0x38, 0x83, 0x48, 0x66, 0xc8, 0xca, 0x4b, 0x82, 0xe9, 0xf9,
	0x7c, 0x9c, 0xfd, 0x3e, 0x80, 0x4c, 0xa6, 0x95, 0x77, 0x9f, 0xcd, 0xbf, 0x3c, 0xb8, 0x10, 0xf2,
	0x46, 0xcd, 0x4e, 0xaf, 0x08, 0x9f, 0x21, 0x7e, 0x21, 0x2b, 0x88, 0xf1, 0xf6, 0x6f, 0x7b, 0x77,
	0x78, 0x93, 0x93, 0x25, 0xd9, 0x44, 0x7a, 0x86, 0x8f, 0x10, 0x39, 0x49, 0x59, 0xfe, 0xd7, 0xb0,
	0x8c, 0xe4, 0xfe, 0xdc, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xb1, 0xe3, 0x28, 0x5a, 0x02,
	0x00, 0x00,
}