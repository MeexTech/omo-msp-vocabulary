// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/edge.proto

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

type ReqVEdgeAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Center               string   `protobuf:"bytes,3,opt,name=center,proto3" json:"center,omitempty"`
	Relation             string   `protobuf:"bytes,4,opt,name=relation,proto3" json:"relation,omitempty"`
	Target               string   `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	Label                string   `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	Direction            uint32   `protobuf:"varint,7,opt,name=direction,proto3" json:"direction,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Weight               uint32   `protobuf:"varint,9,opt,name=weight,proto3" json:"weight,omitempty"`
	Thumb                string   `protobuf:"bytes,10,opt,name=thumb,proto3" json:"thumb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVEdgeAdd) Reset()         { *m = ReqVEdgeAdd{} }
func (m *ReqVEdgeAdd) String() string { return proto.CompactTextString(m) }
func (*ReqVEdgeAdd) ProtoMessage()    {}
func (*ReqVEdgeAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{0}
}

func (m *ReqVEdgeAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqVEdgeAdd.Unmarshal(m, b)
}
func (m *ReqVEdgeAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqVEdgeAdd.Marshal(b, m, deterministic)
}
func (m *ReqVEdgeAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqVEdgeAdd.Merge(m, src)
}
func (m *ReqVEdgeAdd) XXX_Size() int {
	return xxx_messageInfo_ReqVEdgeAdd.Size(m)
}
func (m *ReqVEdgeAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqVEdgeAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqVEdgeAdd proto.InternalMessageInfo

func (m *ReqVEdgeAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqVEdgeAdd) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ReqVEdgeAdd) GetCenter() string {
	if m != nil {
		return m.Center
	}
	return ""
}

func (m *ReqVEdgeAdd) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *ReqVEdgeAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqVEdgeAdd) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ReqVEdgeAdd) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqVEdgeAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqVEdgeAdd) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ReqVEdgeAdd) GetThumb() string {
	if m != nil {
		return m.Thumb
	}
	return ""
}

type ReqVEdgeUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Relation             string   `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	Target               string   `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Direction            uint32   `protobuf:"varint,6,opt,name=direction,proto3" json:"direction,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Thumb                string   `protobuf:"bytes,8,opt,name=thumb,proto3" json:"thumb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqVEdgeUpdate) Reset()         { *m = ReqVEdgeUpdate{} }
func (m *ReqVEdgeUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqVEdgeUpdate) ProtoMessage()    {}
func (*ReqVEdgeUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{1}
}

func (m *ReqVEdgeUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqVEdgeUpdate.Unmarshal(m, b)
}
func (m *ReqVEdgeUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqVEdgeUpdate.Marshal(b, m, deterministic)
}
func (m *ReqVEdgeUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqVEdgeUpdate.Merge(m, src)
}
func (m *ReqVEdgeUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqVEdgeUpdate.Size(m)
}
func (m *ReqVEdgeUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqVEdgeUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqVEdgeUpdate proto.InternalMessageInfo

func (m *ReqVEdgeUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetDirection() uint32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqVEdgeUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqVEdgeUpdate) GetThumb() string {
	if m != nil {
		return m.Thumb
	}
	return ""
}

type ReplyVEdgeInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *VEdgeInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyVEdgeInfo) Reset()         { *m = ReplyVEdgeInfo{} }
func (m *ReplyVEdgeInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyVEdgeInfo) ProtoMessage()    {}
func (*ReplyVEdgeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{2}
}

func (m *ReplyVEdgeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyVEdgeInfo.Unmarshal(m, b)
}
func (m *ReplyVEdgeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyVEdgeInfo.Marshal(b, m, deterministic)
}
func (m *ReplyVEdgeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyVEdgeInfo.Merge(m, src)
}
func (m *ReplyVEdgeInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyVEdgeInfo.Size(m)
}
func (m *ReplyVEdgeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyVEdgeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyVEdgeInfo proto.InternalMessageInfo

func (m *ReplyVEdgeInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyVEdgeInfo) GetInfo() *VEdgeInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyVEdgeList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	List                 []*VEdgeInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyVEdgeList) Reset()         { *m = ReplyVEdgeList{} }
func (m *ReplyVEdgeList) String() string { return proto.CompactTextString(m) }
func (*ReplyVEdgeList) ProtoMessage()    {}
func (*ReplyVEdgeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_06ad0988cd483830, []int{3}
}

func (m *ReplyVEdgeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyVEdgeList.Unmarshal(m, b)
}
func (m *ReplyVEdgeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyVEdgeList.Marshal(b, m, deterministic)
}
func (m *ReplyVEdgeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyVEdgeList.Merge(m, src)
}
func (m *ReplyVEdgeList) XXX_Size() int {
	return xxx_messageInfo_ReplyVEdgeList.Size(m)
}
func (m *ReplyVEdgeList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyVEdgeList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyVEdgeList proto.InternalMessageInfo

func (m *ReplyVEdgeList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyVEdgeList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyVEdgeList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyVEdgeList) GetList() []*VEdgeInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqVEdgeAdd)(nil), "omo.msp.vocabulary.ReqVEdgeAdd")
	proto.RegisterType((*ReqVEdgeUpdate)(nil), "omo.msp.vocabulary.ReqVEdgeUpdate")
	proto.RegisterType((*ReplyVEdgeInfo)(nil), "omo.msp.vocabulary.ReplyVEdgeInfo")
	proto.RegisterType((*ReplyVEdgeList)(nil), "omo.msp.vocabulary.ReplyVEdgeList")
}

func init() {
	proto.RegisterFile("proto/vocabulary/edge.proto", fileDescriptor_06ad0988cd483830)
}

var fileDescriptor_06ad0988cd483830 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x6d, 0x12, 0xc7, 0x4d, 0x36, 0x4d, 0x55, 0xad, 0x3e, 0x7d, 0x5a, 0x05, 0x22, 0x82, 0xb9,
	0xf4, 0xe4, 0x8a, 0x70, 0xe0, 0x9c, 0x4a, 0x50, 0x21, 0x81, 0x2a, 0x5c, 0x51, 0x24, 0x6e, 0x1b,
	0x7b, 0xea, 0xae, 0xb4, 0xf6, 0x9a, 0xf5, 0x3a, 0x28, 0x12, 0x37, 0xfe, 0x0c, 0x3f, 0x87, 0x3b,
	0x7f, 0x06, 0xed, 0xac, 0xeb, 0x46, 0x6d, 0x9d, 0x46, 0x70, 0xf3, 0x7b, 0x3b, 0xf3, 0xde, 0xbc,
	0xd1, 0x24, 0xe4, 0x49, 0xa1, 0x95, 0x51, 0x27, 0x2b, 0x15, 0xf3, 0x65, 0x25, 0xb9, 0x5e, 0x9f,
	0x40, 0x92, 0x42, 0x88, 0x2c, 0xa5, 0x2a, 0x53, 0x61, 0x56, 0x16, 0xe1, 0xed, 0xf3, 0x64, 0x7a,
	0xaf, 0x21, 0x56, 0x59, 0xa6, 0x72, 0xd7, 0x12, 0xfc, 0xe8, 0x92, 0x51, 0x04, 0x5f, 0x2f, 0xdf,
	0x24, 0x29, 0x2c, 0x92, 0x84, 0x52, 0xe2, 0xe5, 0x3c, 0x03, 0xd6, 0x99, 0x75, 0x8e, 0x87, 0x11,
	0x7e, 0xd3, 0xff, 0x89, 0x5f, 0xaa, 0x4a, 0xc7, 0xc0, 0xba, 0xc8, 0xd6, 0xc8, 0xf2, 0x31, 0xe4,
	0x06, 0x34, 0xeb, 0x39, 0xde, 0x21, 0x3a, 0x21, 0x03, 0x0d, 0x92, 0x1b, 0xa1, 0x72, 0xe6, 0xe1,
	0x4b, 0x83, 0x6d, 0x8f, 0xe1, 0x3a, 0x05, 0xc3, 0xfa, 0xae, 0xc7, 0x21, 0xfa, 0x1f, 0xe9, 0x4b,
	0xbe, 0x04, 0xc9, 0x7c, 0xa4, 0x1d, 0xa0, 0x4f, 0xc9, 0x30, 0x11, 0x1a, 0x62, 0x94, 0xda, 0x9f,
	0x75, 0x8e, 0xc7, 0xd1, 0x2d, 0x61, 0x7d, 0x54, 0x01, 0x9a, 0x1b, 0xa5, 0xd9, 0xc0, 0xf9, 0xdc,
	0x60, 0xeb, 0xf3, 0x0d, 0x44, 0x7a, 0x6d, 0xd8, 0x10, 0xdb, 0x6a, 0x64, 0x7d, 0xcc, 0x75, 0x95,
	0x2d, 0x19, 0x71, 0x3e, 0x08, 0x82, 0x5f, 0x1d, 0x72, 0x78, 0xb3, 0x85, 0x4f, 0x45, 0xc2, 0x0d,
	0xd0, 0x23, 0xd2, 0xab, 0x44, 0x52, 0xef, 0xc1, 0x7e, 0x36, 0xab, 0xe9, 0x6e, 0xac, 0x66, 0x33,
	0x6a, 0xaf, 0x35, 0xaa, 0xf7, 0x70, 0xd4, 0x7e, 0x6b, 0x54, 0x7f, 0x5b, 0xd4, 0xfd, 0x3b, 0x51,
	0x9b, 0x48, 0x83, 0xcd, 0x48, 0xdf, 0x6d, 0xa2, 0x42, 0xae, 0x31, 0xd3, 0xbb, 0xfc, 0x4a, 0xd1,
	0xd7, 0xc4, 0x2f, 0x0d, 0x37, 0x55, 0x89, 0xa1, 0x46, 0xf3, 0x67, 0xe1, 0xfd, 0x73, 0x09, 0xb1,
	0xe7, 0x02, 0xcb, 0xa2, 0xba, 0x9c, 0xbe, 0x24, 0x9e, 0xc8, 0xaf, 0x14, 0x06, 0x1f, 0xcd, 0xa7,
	0x0f, 0xb5, 0x35, 0x2e, 0x11, 0x96, 0x06, 0x3f, 0x3b, 0x9b, 0xf6, 0xef, 0x45, 0x69, 0xfe, 0xde,
	0xde, 0xe6, 0x53, 0x86, 0x4b, 0xf4, 0x1f, 0x47, 0x0e, 0x58, 0xb6, 0xe0, 0x29, 0x94, 0xb8, 0xf6,
	0x71, 0xe4, 0x80, 0x1d, 0x55, 0x8a, 0xd2, 0xb0, 0xd1, 0xac, 0xb7, 0xc3, 0xa8, 0xb6, 0x74, 0xfe,
	0xdb, 0x23, 0x07, 0xc8, 0x5d, 0x80, 0x5e, 0x89, 0x18, 0xe8, 0x39, 0xf1, 0x17, 0x49, 0x72, 0x9e,
	0x03, 0x6d, 0x19, 0xb1, 0xf9, 0xb5, 0x4c, 0x82, 0xd6, 0x0c, 0x8d, 0x4b, 0xb0, 0x67, 0x05, 0xcf,
	0xc0, 0x6c, 0x13, 0xac, 0xa0, 0x34, 0xb6, 0x78, 0x47, 0xc1, 0x0f, 0x64, 0x18, 0x41, 0xa6, 0x56,
	0xb0, 0x93, 0xe6, 0xb4, 0x55, 0xb3, 0x96, 0xbb, 0x24, 0xc4, 0x1d, 0x3d, 0x9e, 0x49, 0xb0, 0x2d,
	0xb4, 0xab, 0xdb, 0x71, 0xcc, 0x8f, 0x98, 0x7b, 0x21, 0x25, 0x7d, 0xbe, 0x65, 0xc6, 0xb7, 0x42,
	0x1a, 0xd0, 0x8f, 0x49, 0xda, 0x13, 0x0a, 0xf6, 0xe8, 0x67, 0x72, 0x70, 0x06, 0xc6, 0x1e, 0x88,
	0x28, 0x8d, 0x88, 0xff, 0x4d, 0xb8, 0x91, 0xc1, 0x1d, 0x1c, 0xba, 0x6c, 0xa7, 0x6b, 0xd7, 0x47,
	0x5f, 0xb4, 0x48, 0xbb, 0xb2, 0x5a, 0xfc, 0xb1, 0xdd, 0x9e, 0xd2, 0x2f, 0x47, 0x77, 0xff, 0x80,
	0x97, 0x3e, 0x32, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xe7, 0xdb, 0x02, 0xcc, 0x05,
	0x00, 0x00,
}
