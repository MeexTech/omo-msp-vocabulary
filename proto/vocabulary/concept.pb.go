// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/concept.proto

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

type ConceptInfo struct {
	Type                 uint32         `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string         `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cover                string         `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string         `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Table                string         `protobuf:"bytes,6,opt,name=table,proto3" json:"table,omitempty"`
	Time                 int64          `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
	Updated              int64          `protobuf:"varint,8,opt,name=updated,proto3" json:"updated,omitempty"`
	Created              int64          `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string         `protobuf:"bytes,10,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,11,opt,name=operator,proto3" json:"operator,omitempty"`
	Parent               string         `protobuf:"bytes,12,opt,name=parent,proto3" json:"parent,omitempty"`
	Scene                uint32         `protobuf:"varint,13,opt,name=scene,proto3" json:"scene,omitempty"`
	Attributes           []string       `protobuf:"bytes,14,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Children             []*ConceptInfo `protobuf:"bytes,15,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConceptInfo) Reset()         { *m = ConceptInfo{} }
func (m *ConceptInfo) String() string { return proto.CompactTextString(m) }
func (*ConceptInfo) ProtoMessage()    {}
func (*ConceptInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4eb41a88802c3b, []int{0}
}

func (m *ConceptInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConceptInfo.Unmarshal(m, b)
}
func (m *ConceptInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConceptInfo.Marshal(b, m, deterministic)
}
func (m *ConceptInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConceptInfo.Merge(m, src)
}
func (m *ConceptInfo) XXX_Size() int {
	return xxx_messageInfo_ConceptInfo.Size(m)
}
func (m *ConceptInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConceptInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConceptInfo proto.InternalMessageInfo

func (m *ConceptInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ConceptInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConceptInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ConceptInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ConceptInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ConceptInfo) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ConceptInfo) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ConceptInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ConceptInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ConceptInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ConceptInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ConceptInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ConceptInfo) GetScene() uint32 {
	if m != nil {
		return m.Scene
	}
	return 0
}

func (m *ConceptInfo) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ConceptInfo) GetChildren() []*ConceptInfo {
	if m != nil {
		return m.Children
	}
	return nil
}

type ReqConceptAdd struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Table                string   `protobuf:"bytes,6,opt,name=table,proto3" json:"table,omitempty"`
	Parent               string   `protobuf:"bytes,7,opt,name=parent,proto3" json:"parent,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Scene                uint32   `protobuf:"varint,9,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqConceptAdd) Reset()         { *m = ReqConceptAdd{} }
func (m *ReqConceptAdd) String() string { return proto.CompactTextString(m) }
func (*ReqConceptAdd) ProtoMessage()    {}
func (*ReqConceptAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4eb41a88802c3b, []int{1}
}

func (m *ReqConceptAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqConceptAdd.Unmarshal(m, b)
}
func (m *ReqConceptAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqConceptAdd.Marshal(b, m, deterministic)
}
func (m *ReqConceptAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqConceptAdd.Merge(m, src)
}
func (m *ReqConceptAdd) XXX_Size() int {
	return xxx_messageInfo_ReqConceptAdd.Size(m)
}
func (m *ReqConceptAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqConceptAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqConceptAdd proto.InternalMessageInfo

func (m *ReqConceptAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqConceptAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqConceptAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqConceptAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqConceptAdd) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ReqConceptAdd) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqConceptAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqConceptAdd) GetScene() uint32 {
	if m != nil {
		return m.Scene
	}
	return 0
}

type ReqConceptUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Scene                uint32   `protobuf:"varint,7,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqConceptUpdate) Reset()         { *m = ReqConceptUpdate{} }
func (m *ReqConceptUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqConceptUpdate) ProtoMessage()    {}
func (*ReqConceptUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4eb41a88802c3b, []int{2}
}

func (m *ReqConceptUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqConceptUpdate.Unmarshal(m, b)
}
func (m *ReqConceptUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqConceptUpdate.Marshal(b, m, deterministic)
}
func (m *ReqConceptUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqConceptUpdate.Merge(m, src)
}
func (m *ReqConceptUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqConceptUpdate.Size(m)
}
func (m *ReqConceptUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqConceptUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqConceptUpdate proto.InternalMessageInfo

func (m *ReqConceptUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqConceptUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqConceptUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqConceptUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqConceptUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqConceptUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqConceptUpdate) GetScene() uint32 {
	if m != nil {
		return m.Scene
	}
	return 0
}

type ReplyConceptInfo struct {
	Info                 *ConceptInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyConceptInfo) Reset()         { *m = ReplyConceptInfo{} }
func (m *ReplyConceptInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyConceptInfo) ProtoMessage()    {}
func (*ReplyConceptInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4eb41a88802c3b, []int{3}
}

func (m *ReplyConceptInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyConceptInfo.Unmarshal(m, b)
}
func (m *ReplyConceptInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyConceptInfo.Marshal(b, m, deterministic)
}
func (m *ReplyConceptInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyConceptInfo.Merge(m, src)
}
func (m *ReplyConceptInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyConceptInfo.Size(m)
}
func (m *ReplyConceptInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyConceptInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyConceptInfo proto.InternalMessageInfo

func (m *ReplyConceptInfo) GetInfo() *ConceptInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyConceptInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyConceptList struct {
	Flag                 string         `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	List                 []*ConceptInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyConceptList) Reset()         { *m = ReplyConceptList{} }
func (m *ReplyConceptList) String() string { return proto.CompactTextString(m) }
func (*ReplyConceptList) ProtoMessage()    {}
func (*ReplyConceptList) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4eb41a88802c3b, []int{4}
}

func (m *ReplyConceptList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyConceptList.Unmarshal(m, b)
}
func (m *ReplyConceptList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyConceptList.Marshal(b, m, deterministic)
}
func (m *ReplyConceptList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyConceptList.Merge(m, src)
}
func (m *ReplyConceptList) XXX_Size() int {
	return xxx_messageInfo_ReplyConceptList.Size(m)
}
func (m *ReplyConceptList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyConceptList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyConceptList proto.InternalMessageInfo

func (m *ReplyConceptList) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *ReplyConceptList) GetList() []*ConceptInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyConceptList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyConceptAttrs struct {
	Concept              string       `protobuf:"bytes,1,opt,name=concept,proto3" json:"concept,omitempty"`
	Attributes           []string     `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyConceptAttrs) Reset()         { *m = ReplyConceptAttrs{} }
func (m *ReplyConceptAttrs) String() string { return proto.CompactTextString(m) }
func (*ReplyConceptAttrs) ProtoMessage()    {}
func (*ReplyConceptAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f4eb41a88802c3b, []int{5}
}

func (m *ReplyConceptAttrs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyConceptAttrs.Unmarshal(m, b)
}
func (m *ReplyConceptAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyConceptAttrs.Marshal(b, m, deterministic)
}
func (m *ReplyConceptAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyConceptAttrs.Merge(m, src)
}
func (m *ReplyConceptAttrs) XXX_Size() int {
	return xxx_messageInfo_ReplyConceptAttrs.Size(m)
}
func (m *ReplyConceptAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyConceptAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyConceptAttrs proto.InternalMessageInfo

func (m *ReplyConceptAttrs) GetConcept() string {
	if m != nil {
		return m.Concept
	}
	return ""
}

func (m *ReplyConceptAttrs) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ReplyConceptAttrs) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ConceptInfo)(nil), "omo.msp.vocabulary.ConceptInfo")
	proto.RegisterType((*ReqConceptAdd)(nil), "omo.msp.vocabulary.ReqConceptAdd")
	proto.RegisterType((*ReqConceptUpdate)(nil), "omo.msp.vocabulary.ReqConceptUpdate")
	proto.RegisterType((*ReplyConceptInfo)(nil), "omo.msp.vocabulary.ReplyConceptInfo")
	proto.RegisterType((*ReplyConceptList)(nil), "omo.msp.vocabulary.ReplyConceptList")
	proto.RegisterType((*ReplyConceptAttrs)(nil), "omo.msp.vocabulary.ReplyConceptAttrs")
}

func init() {
	proto.RegisterFile("proto/vocabulary/concept.proto", fileDescriptor_7f4eb41a88802c3b)
}

var fileDescriptor_7f4eb41a88802c3b = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x8e, 0xe3, 0xd4, 0x49, 0xa6, 0x7f, 0x7e, 0xf9, 0xad, 0x10, 0x5a, 0x45, 0x6a, 0x09, 0x56,
	0x91, 0x72, 0x4a, 0xa5, 0xf6, 0xc0, 0x81, 0x53, 0x40, 0xa2, 0x42, 0x02, 0x21, 0x1c, 0x01, 0x12,
	0xe2, 0xe2, 0xd8, 0x13, 0xb0, 0xb0, 0xbd, 0xee, 0x7a, 0x1d, 0x29, 0x37, 0x4e, 0x3c, 0x41, 0x1f,
	0x83, 0xd7, 0xe0, 0x61, 0x78, 0x0b, 0xb4, 0x63, 0x27, 0x71, 0xd3, 0x38, 0x98, 0xaa, 0xb7, 0xfd,
	0xe6, 0x9f, 0xbf, 0xf9, 0x66, 0x46, 0x86, 0x93, 0x44, 0x0a, 0x25, 0xce, 0xe6, 0xc2, 0x73, 0xa7,
	0x59, 0xe8, 0xca, 0xc5, 0x99, 0x27, 0x62, 0x0f, 0x13, 0x35, 0x22, 0x07, 0x63, 0x22, 0x12, 0xa3,
	0x28, 0x4d, 0x46, 0xeb, 0x88, 0xfe, 0xf1, 0x96, 0x9c, 0x28, 0x12, 0x71, 0x9e, 0x62, 0x5f, 0x9b,
	0xb0, 0xff, 0x22, 0x2f, 0xf2, 0x2a, 0x9e, 0x09, 0xc6, 0xa0, 0xa5, 0x16, 0x09, 0x72, 0x63, 0x60,
	0x0c, 0x0f, 0x1d, 0x7a, 0x6b, 0x5b, 0xec, 0x46, 0xc8, 0x9b, 0x03, 0x63, 0xd8, 0x75, 0xe8, 0xcd,
	0x7a, 0x60, 0x66, 0x81, 0xcf, 0x4d, 0x32, 0xe9, 0x27, 0x7b, 0x00, 0x7b, 0x9e, 0x98, 0xa3, 0xe4,
	0x2d, 0xb2, 0xe5, 0x80, 0x3d, 0x04, 0x4b, 0x62, 0xe4, 0xca, 0x6f, 0x7c, 0x8f, 0xcc, 0x05, 0xd2,
	0xd1, 0xca, 0x9d, 0x86, 0xc8, 0xad, 0x3c, 0x9a, 0x00, 0x7d, 0x3d, 0x88, 0x90, 0xb7, 0x07, 0xc6,
	0xd0, 0x74, 0xe8, 0xcd, 0x38, 0xb4, 0xb3, 0xc4, 0x77, 0x15, 0xfa, 0xbc, 0x43, 0xe6, 0x25, 0xd4,
	0x1e, 0x4f, 0x22, 0x79, 0xba, 0xb9, 0xa7, 0x80, 0x2b, 0x8f, 0x90, 0x1c, 0xa8, 0xfe, 0x12, 0xb2,
	0x3e, 0x74, 0x44, 0x82, 0x92, 0x5c, 0xfb, 0xe4, 0x5a, 0x61, 0xcd, 0x35, 0x71, 0x25, 0xc6, 0x8a,
	0x1f, 0xe4, 0x5c, 0x73, 0xa4, 0xb9, 0xa6, 0x1e, 0xc6, 0xc8, 0x0f, 0x49, 0x94, 0x1c, 0xb0, 0x13,
	0x00, 0x57, 0x29, 0x19, 0x4c, 0x33, 0x85, 0x29, 0x3f, 0x1a, 0x98, 0xc3, 0xae, 0x53, 0xb2, 0xb0,
	0x67, 0xd0, 0xf1, 0xbe, 0x06, 0xa1, 0x2f, 0x31, 0xe6, 0xff, 0x0d, 0xcc, 0xe1, 0xfe, 0xf9, 0xa3,
	0xd1, 0xed, 0xf9, 0x8c, 0x4a, 0xe2, 0x3b, 0xab, 0x04, 0xfb, 0x97, 0x01, 0x87, 0x0e, 0x5e, 0x15,
	0xce, 0xb1, 0xef, 0xd7, 0x1e, 0xcc, 0x7d, 0x8c, 0x61, 0x2d, 0x44, 0xfb, 0x86, 0x10, 0x65, 0xf1,
	0x3a, 0x1b, 0xe2, 0xad, 0x44, 0xea, 0x96, 0x44, 0xb2, 0x7f, 0x1a, 0xd0, 0x5b, 0xf7, 0xf1, 0x9e,
	0x06, 0xb7, 0xdc, 0x1d, 0x63, 0xbd, 0x3b, 0xcb, 0xe6, 0x9a, 0x5b, 0x9a, 0x33, 0xef, 0xdc, 0x5c,
	0x99, 0xae, 0x55, 0x45, 0xb7, 0x5d, 0xa6, 0xfb, 0x9d, 0xe8, 0x26, 0xe1, 0xa2, 0x7c, 0x12, 0x17,
	0xd0, 0x0a, 0xe2, 0x99, 0x20, 0xbe, 0x35, 0x86, 0x48, 0xc1, 0xec, 0x29, 0x58, 0xa9, 0x72, 0x55,
	0x96, 0x52, 0x4f, 0x15, 0x69, 0xf4, 0xa9, 0x09, 0x85, 0x39, 0x45, 0xb8, 0x7d, 0xbd, 0x41, 0xe1,
	0x75, 0x90, 0x2a, 0xad, 0xc5, 0x2c, 0x74, 0xbf, 0x14, 0x92, 0xd1, 0x5b, 0xd3, 0x0a, 0x83, 0x54,
	0xf1, 0x66, 0xbd, 0xdd, 0xa2, 0xe0, 0x12, 0x2d, 0xf3, 0xdf, 0x68, 0xfd, 0x30, 0xe0, 0xff, 0x32,
	0xad, 0xb1, 0x52, 0x32, 0xa5, 0x3b, 0xcb, 0x71, 0x41, 0x6d, 0x09, 0x37, 0xae, 0xa3, 0x79, 0xeb,
	0x3a, 0xee, 0x4a, 0xe4, 0xfc, 0x77, 0x0b, 0x8e, 0x0a, 0x0e, 0x13, 0x94, 0xf3, 0xc0, 0x43, 0x36,
	0x01, 0x6b, 0xec, 0xfb, 0x6f, 0x63, 0x64, 0x8f, 0xb7, 0x57, 0x29, 0xdd, 0x51, 0xff, 0xb4, 0xf2,
	0x43, 0x25, 0xb5, 0xec, 0x06, 0x7b, 0x07, 0xd6, 0x25, 0x2a, 0x5d, 0xb4, 0x82, 0xda, 0x55, 0x86,
	0x29, 0x05, 0xd7, 0x2e, 0xf9, 0x06, 0xba, 0x0e, 0x46, 0x62, 0x8e, 0xb5, 0xaa, 0x1e, 0x57, 0x56,
	0xbd, 0xc1, 0x70, 0x1c, 0x86, 0xf7, 0xc0, 0x50, 0x6f, 0x99, 0xdd, 0x60, 0x1f, 0xe1, 0xe0, 0x12,
	0x95, 0x56, 0x3c, 0x48, 0x55, 0xe0, 0x55, 0xea, 0xa9, 0x0b, 0xbf, 0x0c, 0x42, 0x85, 0xb2, 0x6f,
	0xef, 0x1c, 0x1c, 0x95, 0xb1, 0x1b, 0xec, 0x03, 0x58, 0xc5, 0xf1, 0x9f, 0xee, 0x1e, 0x51, 0x1e,
	0x55, 0x5b, 0xd2, 0xcf, 0xd0, 0xcb, 0x33, 0xc6, 0xeb, 0xd5, 0xda, 0xa5, 0x86, 0xee, 0xb3, 0xff,
	0xe4, 0x6f, 0xc5, 0x69, 0xb9, 0xed, 0xc6, 0x73, 0xf6, 0xa9, 0xb7, 0xf9, 0xf7, 0x9c, 0x5a, 0x64,
	0xb9, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x31, 0x94, 0x7a, 0x37, 0x8c, 0x07, 0x00, 0x00,
}
