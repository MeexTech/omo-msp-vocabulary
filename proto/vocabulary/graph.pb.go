// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vocabulary/graph.proto

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

type NodeInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Entity               string   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Labels               []string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{0}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *NodeInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *NodeInfo) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type LinkInfo struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Relation             string        `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	Direction            DirectionType `protobuf:"varint,4,opt,name=direction,proto3,enum=omo.msp.vocabulary.DirectionType" json:"direction,omitempty"`
	From                 string        `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	To                   string        `protobuf:"bytes,6,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LinkInfo) Reset()         { *m = LinkInfo{} }
func (m *LinkInfo) String() string { return proto.CompactTextString(m) }
func (*LinkInfo) ProtoMessage()    {}
func (*LinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{1}
}

func (m *LinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkInfo.Unmarshal(m, b)
}
func (m *LinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkInfo.Marshal(b, m, deterministic)
}
func (m *LinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkInfo.Merge(m, src)
}
func (m *LinkInfo) XXX_Size() int {
	return xxx_messageInfo_LinkInfo.Size(m)
}
func (m *LinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LinkInfo proto.InternalMessageInfo

func (m *LinkInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LinkInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LinkInfo) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

func (m *LinkInfo) GetDirection() DirectionType {
	if m != nil {
		return m.Direction
	}
	return DirectionType_Double
}

func (m *LinkInfo) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *LinkInfo) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type GraphInfo struct {
	Center               string      `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Nodes                []*NodeInfo `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Links                []*LinkInfo `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GraphInfo) Reset()         { *m = GraphInfo{} }
func (m *GraphInfo) String() string { return proto.CompactTextString(m) }
func (*GraphInfo) ProtoMessage()    {}
func (*GraphInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{2}
}

func (m *GraphInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphInfo.Unmarshal(m, b)
}
func (m *GraphInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphInfo.Marshal(b, m, deterministic)
}
func (m *GraphInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphInfo.Merge(m, src)
}
func (m *GraphInfo) XXX_Size() int {
	return xxx_messageInfo_GraphInfo.Size(m)
}
func (m *GraphInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GraphInfo proto.InternalMessageInfo

func (m *GraphInfo) GetCenter() string {
	if m != nil {
		return m.Center
	}
	return ""
}

func (m *GraphInfo) GetNodes() []*NodeInfo {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *GraphInfo) GetLinks() []*LinkInfo {
	if m != nil {
		return m.Links
	}
	return nil
}

type ReqNodeAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Entity               string   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Label                string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqNodeAdd) Reset()         { *m = ReqNodeAdd{} }
func (m *ReqNodeAdd) String() string { return proto.CompactTextString(m) }
func (*ReqNodeAdd) ProtoMessage()    {}
func (*ReqNodeAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{3}
}

func (m *ReqNodeAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqNodeAdd.Unmarshal(m, b)
}
func (m *ReqNodeAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqNodeAdd.Marshal(b, m, deterministic)
}
func (m *ReqNodeAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqNodeAdd.Merge(m, src)
}
func (m *ReqNodeAdd) XXX_Size() int {
	return xxx_messageInfo_ReqNodeAdd.Size(m)
}
func (m *ReqNodeAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqNodeAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqNodeAdd proto.InternalMessageInfo

func (m *ReqNodeAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqNodeAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqNodeAdd) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ReqNodeAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

type ReplyNodeInfo struct {
	Info                 *NodeInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyNodeInfo) Reset()         { *m = ReplyNodeInfo{} }
func (m *ReplyNodeInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyNodeInfo) ProtoMessage()    {}
func (*ReplyNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{4}
}

func (m *ReplyNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyNodeInfo.Unmarshal(m, b)
}
func (m *ReplyNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyNodeInfo.Marshal(b, m, deterministic)
}
func (m *ReplyNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyNodeInfo.Merge(m, src)
}
func (m *ReplyNodeInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyNodeInfo.Size(m)
}
func (m *ReplyNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyNodeInfo proto.InternalMessageInfo

func (m *ReplyNodeInfo) GetInfo() *NodeInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyNodeInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqLinkAdd struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Direction            DirectionType `protobuf:"varint,2,opt,name=direction,proto3,enum=omo.msp.vocabulary.DirectionType" json:"direction,omitempty"`
	Key                  string        `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	From                 string        `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   string        `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	Relation             string        `protobuf:"bytes,6,opt,name=relation,proto3" json:"relation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReqLinkAdd) Reset()         { *m = ReqLinkAdd{} }
func (m *ReqLinkAdd) String() string { return proto.CompactTextString(m) }
func (*ReqLinkAdd) ProtoMessage()    {}
func (*ReqLinkAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{5}
}

func (m *ReqLinkAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLinkAdd.Unmarshal(m, b)
}
func (m *ReqLinkAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLinkAdd.Marshal(b, m, deterministic)
}
func (m *ReqLinkAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLinkAdd.Merge(m, src)
}
func (m *ReqLinkAdd) XXX_Size() int {
	return xxx_messageInfo_ReqLinkAdd.Size(m)
}
func (m *ReqLinkAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLinkAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLinkAdd proto.InternalMessageInfo

func (m *ReqLinkAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqLinkAdd) GetDirection() DirectionType {
	if m != nil {
		return m.Direction
	}
	return DirectionType_Double
}

func (m *ReqLinkAdd) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqLinkAdd) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqLinkAdd) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqLinkAdd) GetRelation() string {
	if m != nil {
		return m.Relation
	}
	return ""
}

type ReplyLinkInfo struct {
	Info                 *LinkInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyLinkInfo) Reset()         { *m = ReplyLinkInfo{} }
func (m *ReplyLinkInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyLinkInfo) ProtoMessage()    {}
func (*ReplyLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{6}
}

func (m *ReplyLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyLinkInfo.Unmarshal(m, b)
}
func (m *ReplyLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyLinkInfo.Marshal(b, m, deterministic)
}
func (m *ReplyLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyLinkInfo.Merge(m, src)
}
func (m *ReplyLinkInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyLinkInfo.Size(m)
}
func (m *ReplyLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyLinkInfo proto.InternalMessageInfo

func (m *ReplyLinkInfo) GetInfo() *LinkInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyLinkInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqGraphPath struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqGraphPath) Reset()         { *m = ReqGraphPath{} }
func (m *ReqGraphPath) String() string { return proto.CompactTextString(m) }
func (*ReqGraphPath) ProtoMessage()    {}
func (*ReqGraphPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{7}
}

func (m *ReqGraphPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqGraphPath.Unmarshal(m, b)
}
func (m *ReqGraphPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqGraphPath.Marshal(b, m, deterministic)
}
func (m *ReqGraphPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqGraphPath.Merge(m, src)
}
func (m *ReqGraphPath) XXX_Size() int {
	return xxx_messageInfo_ReqGraphPath.Size(m)
}
func (m *ReqGraphPath) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqGraphPath.DiscardUnknown(m)
}

var xxx_messageInfo_ReqGraphPath proto.InternalMessageInfo

func (m *ReqGraphPath) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqGraphPath) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReplyGraphInfo struct {
	Graph                *GraphInfo   `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyGraphInfo) Reset()         { *m = ReplyGraphInfo{} }
func (m *ReplyGraphInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyGraphInfo) ProtoMessage()    {}
func (*ReplyGraphInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f4112fd2703b34, []int{8}
}

func (m *ReplyGraphInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyGraphInfo.Unmarshal(m, b)
}
func (m *ReplyGraphInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyGraphInfo.Marshal(b, m, deterministic)
}
func (m *ReplyGraphInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyGraphInfo.Merge(m, src)
}
func (m *ReplyGraphInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyGraphInfo.Size(m)
}
func (m *ReplyGraphInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyGraphInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyGraphInfo proto.InternalMessageInfo

func (m *ReplyGraphInfo) GetGraph() *GraphInfo {
	if m != nil {
		return m.Graph
	}
	return nil
}

func (m *ReplyGraphInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeInfo)(nil), "omo.msp.vocabulary.NodeInfo")
	proto.RegisterType((*LinkInfo)(nil), "omo.msp.vocabulary.LinkInfo")
	proto.RegisterType((*GraphInfo)(nil), "omo.msp.vocabulary.GraphInfo")
	proto.RegisterType((*ReqNodeAdd)(nil), "omo.msp.vocabulary.ReqNodeAdd")
	proto.RegisterType((*ReplyNodeInfo)(nil), "omo.msp.vocabulary.ReplyNodeInfo")
	proto.RegisterType((*ReqLinkAdd)(nil), "omo.msp.vocabulary.ReqLinkAdd")
	proto.RegisterType((*ReplyLinkInfo)(nil), "omo.msp.vocabulary.ReplyLinkInfo")
	proto.RegisterType((*ReqGraphPath)(nil), "omo.msp.vocabulary.ReqGraphPath")
	proto.RegisterType((*ReplyGraphInfo)(nil), "omo.msp.vocabulary.ReplyGraphInfo")
}

func init() {
	proto.RegisterFile("proto/vocabulary/graph.proto", fileDescriptor_32f4112fd2703b34)
}

var fileDescriptor_32f4112fd2703b34 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x76, 0xec, 0x26, 0xd3, 0x52, 0x55, 0xab, 0xaa, 0xb2, 0xa2, 0x06, 0x52, 0x9f, 0x72,
	0x4a, 0x91, 0x7b, 0xe0, 0x88, 0x8a, 0x10, 0x15, 0x12, 0x2a, 0x68, 0xcb, 0x89, 0x9b, 0x63, 0x6f,
	0xa8, 0x15, 0x7b, 0x37, 0xb1, 0x37, 0x91, 0xc2, 0x81, 0x1f, 0xe0, 0x63, 0x38, 0xf0, 0x5b, 0x7c,
	0x04, 0xda, 0xb1, 0xbd, 0x6e, 0x9a, 0x98, 0x5a, 0x2d, 0xb7, 0x9d, 0xf1, 0xcc, 0x9b, 0x37, 0x6f,
	0x5e, 0x1b, 0x38, 0x9d, 0x67, 0x42, 0x8a, 0xf3, 0x95, 0x08, 0x83, 0xc9, 0x32, 0x09, 0xb2, 0xf5,
	0xf9, 0xb7, 0x2c, 0x98, 0xdf, 0x8e, 0x31, 0x4d, 0x88, 0x48, 0xc5, 0x38, 0xcd, 0xe7, 0xe3, 0xfa,
	0x7b, 0x7f, 0xb0, 0xd5, 0x11, 0x8a, 0x34, 0x15, 0xbc, 0x68, 0xf1, 0x24, 0x74, 0xaf, 0x45, 0xc4,
	0x3e, 0xf0, 0xa9, 0x20, 0x87, 0x60, 0xc6, 0x91, 0x6b, 0x0c, 0x8d, 0x91, 0x45, 0xcd, 0x38, 0x22,
	0x04, 0x3a, 0x3c, 0x48, 0x99, 0x6b, 0x0e, 0x8d, 0x51, 0x8f, 0xe2, 0x9b, 0x1c, 0x83, 0x1d, 0x8a,
	0x15, 0xcb, 0x5c, 0x0b, 0x93, 0x45, 0x40, 0x4e, 0xc0, 0x61, 0x5c, 0xc6, 0x72, 0xed, 0x76, 0x30,
	0x5d, 0x46, 0x2a, 0x9f, 0x04, 0x13, 0x96, 0xe4, 0xae, 0x3d, 0xb4, 0x54, 0xbe, 0x88, 0xbc, 0x5f,
	0x06, 0x74, 0x3f, 0xc6, 0x7c, 0xd6, 0x7a, 0x6c, 0x1f, 0xba, 0x19, 0x4b, 0x02, 0x19, 0x0b, 0x5e,
	0x4e, 0xd6, 0x31, 0x79, 0x03, 0xbd, 0x28, 0xce, 0x58, 0x88, 0x1f, 0xd5, 0xfc, 0x43, 0xff, 0x6c,
	0xbc, 0xad, 0xc4, 0xf8, 0x5d, 0x55, 0xf4, 0x65, 0x3d, 0x67, 0xb4, 0xee, 0x51, 0x03, 0xa7, 0x99,
	0x48, 0x5d, 0xbb, 0x18, 0xa8, 0xde, 0x8a, 0x94, 0x14, 0xae, 0x83, 0x19, 0x53, 0x0a, 0xef, 0xa7,
	0x01, 0xbd, 0x2b, 0x25, 0x35, 0x52, 0x3e, 0x01, 0x27, 0x64, 0x5c, 0xb2, 0x0c, 0x69, 0xf7, 0x68,
	0x19, 0x11, 0x1f, 0x6c, 0x2e, 0x22, 0x96, 0xbb, 0xe6, 0xd0, 0x1a, 0xed, 0xfb, 0xa7, 0xbb, 0x68,
	0x54, 0x72, 0xd3, 0xa2, 0x54, 0xf5, 0x24, 0x31, 0x9f, 0xe5, 0xae, 0xd5, 0xdc, 0x53, 0x69, 0x45,
	0x8b, 0x52, 0x2f, 0x02, 0xa0, 0x6c, 0xa1, 0x90, 0x2e, 0xa3, 0x5a, 0x30, 0xe3, 0x8e, 0x60, 0xf5,
	0x45, 0xcc, 0x8d, 0x8b, 0x1c, 0x83, 0x8d, 0x37, 0xa8, 0xee, 0x87, 0x41, 0x7d, 0xd5, 0xce, 0x9d,
	0xab, 0x7a, 0xdf, 0xe1, 0x39, 0x65, 0xf3, 0x64, 0xad, 0x0d, 0xf2, 0x0a, 0x3a, 0x31, 0x9f, 0x0a,
	0x1c, 0xf4, 0xd0, 0x76, 0x58, 0x49, 0x5e, 0x83, 0x93, 0xcb, 0x40, 0x2e, 0x73, 0xa4, 0xb1, 0xef,
	0xbf, 0xdc, 0xd5, 0x83, 0x43, 0x6e, 0xb0, 0x8c, 0x96, 0xe5, 0xde, 0x6f, 0x03, 0x57, 0x54, 0x8b,
	0x37, 0xad, 0xb8, 0x71, 0x77, 0xf3, 0x11, 0x77, 0x3f, 0x02, 0x6b, 0xc6, 0xd6, 0xa5, 0x12, 0xea,
	0xa9, 0x9d, 0xd0, 0xd9, 0x72, 0x82, 0x5d, 0x39, 0x61, 0xc3, 0x8a, 0xce, 0xa6, 0x15, 0xb5, 0x62,
	0xda, 0xdb, 0x2d, 0x14, 0xd3, 0xb7, 0x7d, 0xa2, 0x62, 0x3e, 0x1c, 0x50, 0xb6, 0x40, 0x8f, 0x7e,
	0x0e, 0xe4, 0xad, 0xde, 0xc5, 0xd8, 0xda, 0xc5, 0xd4, 0xae, 0xfe, 0x01, 0x87, 0x08, 0x55, 0x3b,
	0xfb, 0x02, 0x6c, 0xfc, 0x8f, 0x52, 0x32, 0x1e, 0xec, 0x9a, 0xae, 0xab, 0x69, 0x51, 0xfb, 0x68,
	0xce, 0xfe, 0x1f, 0x1b, 0x0e, 0x10, 0xed, 0x86, 0x65, 0xab, 0x38, 0x64, 0xe4, 0x1a, 0xf6, 0x2e,
	0xa3, 0x48, 0x99, 0x88, 0xbc, 0xd8, 0x0d, 0x52, 0xb9, 0xbe, 0x7f, 0xd6, 0x38, 0xa4, 0xf2, 0xa0,
	0xf7, 0xac, 0xc4, 0x53, 0x12, 0x37, 0xe2, 0x95, 0x16, 0xfb, 0x07, 0x5e, 0x75, 0x21, 0xc4, 0x03,
	0xca, 0x52, 0xb1, 0x62, 0x48, 0xb1, 0x61, 0xcf, 0xc5, 0x92, 0xe5, 0x52, 0x35, 0xf4, 0x07, 0x8d,
	0x98, 0xf7, 0xf1, 0x90, 0xe2, 0xd3, 0xf1, 0x3e, 0xc1, 0xde, 0x15, 0x93, 0xed, 0xc8, 0xb5, 0x12,
	0xb0, 0x00, 0x6c, 0xc7, 0xae, 0x95, 0x82, 0x14, 0x7a, 0xef, 0x63, 0x8e, 0x27, 0xce, 0x1f, 0x86,
	0xf4, 0x1a, 0x21, 0xb5, 0x09, 0x11, 0xb3, 0xab, 0x30, 0xd1, 0xf6, 0xc3, 0x06, 0x48, 0xfd, 0x87,
	0xd1, 0x1a, 0x13, 0x79, 0x62, 0xea, 0x3f, 0xf1, 0x7c, 0x4b, 0xbe, 0x1e, 0xdd, 0xff, 0x35, 0x9e,
	0x38, 0x98, 0xb9, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x0a, 0xb2, 0xe0, 0xda, 0x07, 0x00,
	0x00,
}
