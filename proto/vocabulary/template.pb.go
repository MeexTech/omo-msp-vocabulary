// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/template.proto

package vocabulary

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

type TemplateInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	Limit                uint32   `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	Name                 string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,11,opt,name=cover,proto3" json:"cover,omitempty"`
	Names                []string `protobuf:"bytes,21,rep,name=names,proto3" json:"names,omitempty"`
	Covers               []string `protobuf:"bytes,22,rep,name=covers,proto3" json:"covers,omitempty"`
	Scores               []string `protobuf:"bytes,23,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateInfo) Reset()         { *m = TemplateInfo{} }
func (m *TemplateInfo) String() string { return proto.CompactTextString(m) }
func (*TemplateInfo) ProtoMessage()    {}
func (*TemplateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b79888774522469, []int{0}
}

func (m *TemplateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateInfo.Unmarshal(m, b)
}
func (m *TemplateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateInfo.Marshal(b, m, deterministic)
}
func (m *TemplateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateInfo.Merge(m, src)
}
func (m *TemplateInfo) XXX_Size() int {
	return xxx_messageInfo_TemplateInfo.Size(m)
}
func (m *TemplateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateInfo proto.InternalMessageInfo

func (m *TemplateInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TemplateInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TemplateInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TemplateInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TemplateInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TemplateInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TemplateInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TemplateInfo) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TemplateInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TemplateInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *TemplateInfo) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *TemplateInfo) GetCovers() []string {
	if m != nil {
		return m.Covers
	}
	return nil
}

func (m *TemplateInfo) GetScores() []string {
	if m != nil {
		return m.Scores
	}
	return nil
}

type ReqTemplateAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Names                []string `protobuf:"bytes,11,rep,name=names,proto3" json:"names,omitempty"`
	Covers               []string `protobuf:"bytes,12,rep,name=covers,proto3" json:"covers,omitempty"`
	Scores               []string `protobuf:"bytes,13,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTemplateAdd) Reset()         { *m = ReqTemplateAdd{} }
func (m *ReqTemplateAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTemplateAdd) ProtoMessage()    {}
func (*ReqTemplateAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b79888774522469, []int{1}
}

func (m *ReqTemplateAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTemplateAdd.Unmarshal(m, b)
}
func (m *ReqTemplateAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTemplateAdd.Marshal(b, m, deterministic)
}
func (m *ReqTemplateAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTemplateAdd.Merge(m, src)
}
func (m *ReqTemplateAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTemplateAdd.Size(m)
}
func (m *ReqTemplateAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTemplateAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTemplateAdd proto.InternalMessageInfo

func (m *ReqTemplateAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTemplateAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqTemplateAdd) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqTemplateAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqTemplateAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTemplateAdd) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ReqTemplateAdd) GetCovers() []string {
	if m != nil {
		return m.Covers
	}
	return nil
}

func (m *ReqTemplateAdd) GetScores() []string {
	if m != nil {
		return m.Scores
	}
	return nil
}

type ReqTemplateUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Limit                uint32   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Type                 uint32   `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Names                []string `protobuf:"bytes,11,rep,name=names,proto3" json:"names,omitempty"`
	Covers               []string `protobuf:"bytes,12,rep,name=covers,proto3" json:"covers,omitempty"`
	Scores               []string `protobuf:"bytes,13,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTemplateUpdate) Reset()         { *m = ReqTemplateUpdate{} }
func (m *ReqTemplateUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqTemplateUpdate) ProtoMessage()    {}
func (*ReqTemplateUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b79888774522469, []int{2}
}

func (m *ReqTemplateUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTemplateUpdate.Unmarshal(m, b)
}
func (m *ReqTemplateUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTemplateUpdate.Marshal(b, m, deterministic)
}
func (m *ReqTemplateUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTemplateUpdate.Merge(m, src)
}
func (m *ReqTemplateUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqTemplateUpdate.Size(m)
}
func (m *ReqTemplateUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTemplateUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTemplateUpdate proto.InternalMessageInfo

func (m *ReqTemplateUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqTemplateUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTemplateUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqTemplateUpdate) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqTemplateUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqTemplateUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTemplateUpdate) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ReqTemplateUpdate) GetCovers() []string {
	if m != nil {
		return m.Covers
	}
	return nil
}

func (m *ReqTemplateUpdate) GetScores() []string {
	if m != nil {
		return m.Scores
	}
	return nil
}

type ReplyTemplateInfo struct {
	Info                 *TemplateInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyTemplateInfo) Reset()         { *m = ReplyTemplateInfo{} }
func (m *ReplyTemplateInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyTemplateInfo) ProtoMessage()    {}
func (*ReplyTemplateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b79888774522469, []int{3}
}

func (m *ReplyTemplateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTemplateInfo.Unmarshal(m, b)
}
func (m *ReplyTemplateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTemplateInfo.Marshal(b, m, deterministic)
}
func (m *ReplyTemplateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTemplateInfo.Merge(m, src)
}
func (m *ReplyTemplateInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyTemplateInfo.Size(m)
}
func (m *ReplyTemplateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTemplateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTemplateInfo proto.InternalMessageInfo

func (m *ReplyTemplateInfo) GetInfo() *TemplateInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyTemplateInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyTemplateList struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32          `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	List                 []*TemplateInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyTemplateList) Reset()         { *m = ReplyTemplateList{} }
func (m *ReplyTemplateList) String() string { return proto.CompactTextString(m) }
func (*ReplyTemplateList) ProtoMessage()    {}
func (*ReplyTemplateList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b79888774522469, []int{4}
}

func (m *ReplyTemplateList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTemplateList.Unmarshal(m, b)
}
func (m *ReplyTemplateList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTemplateList.Marshal(b, m, deterministic)
}
func (m *ReplyTemplateList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTemplateList.Merge(m, src)
}
func (m *ReplyTemplateList) XXX_Size() int {
	return xxx_messageInfo_ReplyTemplateList.Size(m)
}
func (m *ReplyTemplateList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTemplateList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTemplateList proto.InternalMessageInfo

func (m *ReplyTemplateList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTemplateList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTemplateList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyTemplateList) GetList() []*TemplateInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*TemplateInfo)(nil), "omo.msp.vocabulary.TemplateInfo")
	proto.RegisterType((*ReqTemplateAdd)(nil), "omo.msp.vocabulary.ReqTemplateAdd")
	proto.RegisterType((*ReqTemplateUpdate)(nil), "omo.msp.vocabulary.ReqTemplateUpdate")
	proto.RegisterType((*ReplyTemplateInfo)(nil), "omo.msp.vocabulary.ReplyTemplateInfo")
	proto.RegisterType((*ReplyTemplateList)(nil), "omo.msp.vocabulary.ReplyTemplateList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/template.proto", fileDescriptor_5b79888774522469)
}

var fileDescriptor_5b79888774522469 = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xa3, 0x3f, 0x56, 0x9a, 0x71, 0xec, 0xa4, 0x4b, 0x9b, 0x2e, 0x86, 0x10, 0x55, 0x25,
	0xe0, 0x93, 0x03, 0x6e, 0xa1, 0xe7, 0xe4, 0x50, 0x53, 0x68, 0x29, 0x28, 0x09, 0x85, 0x42, 0x0b,
	0x8a, 0xb4, 0x29, 0x0b, 0x92, 0x57, 0xd1, 0xae, 0x0d, 0xbe, 0xf6, 0x81, 0x7a, 0xea, 0x33, 0xf4,
	0x25, 0xfa, 0x24, 0xbd, 0x95, 0x9d, 0x95, 0x5c, 0xc5, 0x91, 0x12, 0xf5, 0xd0, 0xdb, 0xce, 0xcc,
	0xb7, 0x9f, 0xe6, 0x37, 0x9e, 0xc5, 0x70, 0x94, 0x17, 0x42, 0x89, 0x93, 0xa5, 0x88, 0xa3, 0xab,
	0x45, 0x1a, 0x15, 0xab, 0x13, 0xc5, 0xb2, 0x3c, 0x8d, 0x14, 0x9b, 0x60, 0x85, 0x10, 0x91, 0x89,
	0x49, 0x26, 0xf3, 0xc9, 0x5f, 0xc9, 0xe8, 0xf0, 0xce, 0xa5, 0x58, 0x64, 0x99, 0x98, 0x9b, 0x2b,
	0xc1, 0x0f, 0x1b, 0x76, 0x2f, 0x4a, 0x97, 0xb7, 0xf3, 0x6b, 0x41, 0xf6, 0xc1, 0x59, 0xf0, 0x84,
	0x5a, 0xbe, 0x35, 0xde, 0x09, 0xf5, 0x91, 0x0c, 0xc1, 0xe6, 0x09, 0xb5, 0x7d, 0x6b, 0xec, 0x86,
	0x36, 0x4f, 0x08, 0x85, 0xed, 0xb8, 0x60, 0x91, 0x62, 0x09, 0x75, 0x7c, 0x6b, 0xec, 0x84, 0x55,
	0xa8, 0x2b, 0x8b, 0x3c, 0xc1, 0x8a, 0x6b, 0x2a, 0x65, 0xb8, 0xbe, 0x23, 0x0a, 0xda, 0x43, 0xe7,
	0x2a, 0x24, 0x23, 0x78, 0x24, 0x72, 0x56, 0x60, 0xc9, 0xc3, 0xd2, 0x3a, 0x26, 0x04, 0x5c, 0xb5,
	0xca, 0x19, 0xdd, 0xf6, 0xad, 0xf1, 0x20, 0xc4, 0x33, 0x79, 0x02, 0xbd, 0x94, 0x67, 0x5c, 0xd1,
	0x1d, 0x4c, 0x9a, 0x40, 0x2b, 0xe7, 0x51, 0xc6, 0x28, 0xa0, 0x03, 0x9e, 0xb5, 0x32, 0x16, 0x4b,
	0x56, 0xd0, 0x3e, 0x26, 0x4d, 0xa0, 0xb3, 0xba, 0x2a, 0xe9, 0x53, 0xdf, 0xd1, 0x59, 0x0c, 0xc8,
	0x01, 0x78, 0x58, 0x96, 0xf4, 0x00, 0xd3, 0x65, 0xa4, 0xf3, 0x32, 0x16, 0x05, 0x93, 0xf4, 0x99,
	0xc9, 0x9b, 0x28, 0xf8, 0x69, 0xc1, 0x30, 0x64, 0x37, 0xd5, 0xe4, 0x4e, 0x93, 0x64, 0xdd, 0x82,
	0xd5, 0xd4, 0x82, 0xbd, 0xd1, 0x82, 0x41, 0x70, 0x36, 0x10, 0x10, 0xd6, 0xad, 0xc1, 0xd6, 0x87,
	0xd3, 0xdb, 0x18, 0xce, 0x1a, 0xa4, 0xdf, 0x0c, 0xb2, 0xdb, 0x02, 0x32, 0xb8, 0x05, 0xf2, 0xcb,
	0x82, 0xc7, 0x35, 0x90, 0x4b, 0xfc, 0xbd, 0x1a, 0x96, 0xa0, 0xa2, 0xb3, 0x9b, 0xe8, 0x9c, 0x46,
	0x3a, 0xb7, 0x89, 0xae, 0xd7, 0x42, 0xe7, 0xfd, 0x17, 0xba, 0x6f, 0x48, 0x97, 0xa7, 0xab, 0x5b,
	0x2b, 0xfe, 0x0a, 0x5c, 0x3e, 0xbf, 0x16, 0x88, 0xd7, 0x9f, 0xfa, 0x93, 0xbb, 0xaf, 0x66, 0x52,
	0xd7, 0x87, 0xa8, 0x26, 0xaf, 0xc1, 0x93, 0x2a, 0x52, 0x0b, 0x89, 0x33, 0xe8, 0x4f, 0x8f, 0x9a,
	0xee, 0xe1, 0xc7, 0xce, 0x51, 0x16, 0x96, 0xf2, 0xe0, 0xfb, 0x66, 0x13, 0xef, 0xb8, 0x54, 0x35,
	0x3b, 0xeb, 0x9f, 0xec, 0xf4, 0x64, 0x94, 0x50, 0x51, 0x8a, 0x6d, 0x0c, 0x42, 0x13, 0xe8, 0x6c,
	0x1e, 0x7d, 0x65, 0xb2, 0xda, 0x29, 0x0c, 0x34, 0x69, 0xca, 0xa5, 0xc2, 0x21, 0x76, 0x22, 0xd5,
	0xea, 0xe9, 0x6f, 0x17, 0xf6, 0xaa, 0xf4, 0x39, 0x2b, 0x96, 0x3c, 0x66, 0xe4, 0x12, 0xbc, 0xd3,
	0x24, 0xf9, 0x30, 0x67, 0x24, 0x68, 0x6e, 0xb4, 0xfe, 0x16, 0x46, 0xc7, 0xad, 0x30, 0xf5, 0xcf,
	0x05, 0x5b, 0x24, 0x04, 0x6f, 0xc6, 0x94, 0xb6, 0x6d, 0xe1, 0xbf, 0x59, 0x30, 0xa9, 0xb4, 0xb8,
	0xbb, 0xe7, 0x7b, 0xd8, 0x09, 0x59, 0x26, 0x96, 0xac, 0x93, 0xed, 0x61, 0xab, 0x6d, 0x69, 0xf7,
	0x05, 0x86, 0xe6, 0x55, 0x9c, 0xad, 0xde, 0xf0, 0x54, 0xb1, 0x82, 0xbc, 0x68, 0xf1, 0x34, 0x32,
	0x23, 0xea, 0xde, 0xee, 0x05, 0x40, 0xe9, 0x1f, 0x49, 0x46, 0x8e, 0x1f, 0x98, 0xae, 0x91, 0x3e,
	0xdc, 0xf5, 0x67, 0xd8, 0x9b, 0x31, 0xa5, 0x37, 0x6d, 0xdd, 0xf6, 0xf3, 0x7b, 0x46, 0xd1, 0xb9,
	0x69, 0xed, 0x18, 0x6c, 0x91, 0x8f, 0xb0, 0x3b, 0x63, 0x4a, 0x6f, 0x26, 0x97, 0x8a, 0xc7, 0x5d,
	0xbc, 0x83, 0x7b, 0x17, 0x1c, 0x6d, 0x82, 0xad, 0x33, 0xf2, 0x69, 0x7f, 0xf3, 0x0f, 0xeb, 0xca,
	0xc3, 0xcc, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x99, 0x79, 0xd8, 0x00, 0x07, 0x00,
	0x00,
}
