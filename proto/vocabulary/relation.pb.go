// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: proto/vocabulary/relation.proto

package vocabulary

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RelationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type     string          `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name     string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Remark   string          `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Time     uint64          `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	Custom   bool            `protobuf:"varint,6,opt,name=custom,proto3" json:"custom,omitempty"`
	Children []*RelationInfo `protobuf:"bytes,7,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *RelationInfo) Reset() {
	*x = RelationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationInfo) ProtoMessage() {}

func (x *RelationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationInfo.ProtoReflect.Descriptor instead.
func (*RelationInfo) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RelationInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RelationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RelationInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *RelationInfo) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RelationInfo) GetCustom() bool {
	if x != nil {
		return x.Custom
	}
	return false
}

func (x *RelationInfo) GetChildren() []*RelationInfo {
	if x != nil {
		return x.Children
	}
	return nil
}

type ReqRelationAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Parent string `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Custom bool   `protobuf:"varint,4,opt,name=custom,proto3" json:"custom,omitempty"`
}

func (x *ReqRelationAdd) Reset() {
	*x = ReqRelationAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRelationAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRelationAdd) ProtoMessage() {}

func (x *ReqRelationAdd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRelationAdd.ProtoReflect.Descriptor instead.
func (*ReqRelationAdd) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_relation_proto_rawDescGZIP(), []int{1}
}

func (x *ReqRelationAdd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqRelationAdd) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ReqRelationAdd) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ReqRelationAdd) GetCustom() bool {
	if x != nil {
		return x.Custom
	}
	return false
}

type ReplyRelationOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *RelationInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	ErrorCode ResultStatus  `protobuf:"varint,2,opt,name=errorCode,proto3,enum=omo.msp.vocabulary.ResultStatus" json:"errorCode,omitempty"`
}

func (x *ReplyRelationOne) Reset() {
	*x = ReplyRelationOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRelationOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRelationOne) ProtoMessage() {}

func (x *ReplyRelationOne) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRelationOne.ProtoReflect.Descriptor instead.
func (*ReplyRelationOne) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_relation_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyRelationOne) GetInfo() *RelationInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReplyRelationOne) GetErrorCode() ResultStatus {
	if x != nil {
		return x.ErrorCode
	}
	return ResultStatus_Success
}

type ReplyRelationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*RelationInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ReplyRelationList) Reset() {
	*x = ReplyRelationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRelationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRelationList) ProtoMessage() {}

func (x *ReplyRelationList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRelationList.ProtoReflect.Descriptor instead.
func (*ReplyRelationList) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_relation_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyRelationList) GetList() []*RelationInfo {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_vocabulary_relation_proto protoreflect.FileDescriptor

var file_proto_vocabulary_relation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62,
	0x75, 0x6c, 0x61, 0x72, 0x79, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63,
	0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x22, 0x6c, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x22,
	0x88, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63,
	0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x49, 0x0a, 0x11, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xdd, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x41, 0x64, 0x64,
	0x4f, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f,
	0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x6e, 0x65, 0x12,
	0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62,
	0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x2e, 0x6f, 0x6d,
	0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x25, 0x2e, 0x6f,
	0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72,
	0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_vocabulary_relation_proto_rawDescOnce sync.Once
	file_proto_vocabulary_relation_proto_rawDescData = file_proto_vocabulary_relation_proto_rawDesc
)

func file_proto_vocabulary_relation_proto_rawDescGZIP() []byte {
	file_proto_vocabulary_relation_proto_rawDescOnce.Do(func() {
		file_proto_vocabulary_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vocabulary_relation_proto_rawDescData)
	})
	return file_proto_vocabulary_relation_proto_rawDescData
}

var file_proto_vocabulary_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_vocabulary_relation_proto_goTypes = []interface{}{
	(*RelationInfo)(nil),      // 0: omo.msp.vocabulary.RelationInfo
	(*ReqRelationAdd)(nil),    // 1: omo.msp.vocabulary.ReqRelationAdd
	(*ReplyRelationOne)(nil),  // 2: omo.msp.vocabulary.ReplyRelationOne
	(*ReplyRelationList)(nil), // 3: omo.msp.vocabulary.ReplyRelationList
	(ResultStatus)(0),         // 4: omo.msp.vocabulary.ResultStatus
	(*RequestInfo)(nil),       // 5: omo.msp.vocabulary.RequestInfo
	(*ReplyInfo)(nil),         // 6: omo.msp.vocabulary.ReplyInfo
}
var file_proto_vocabulary_relation_proto_depIdxs = []int32{
	0, // 0: omo.msp.vocabulary.RelationInfo.children:type_name -> omo.msp.vocabulary.RelationInfo
	0, // 1: omo.msp.vocabulary.ReplyRelationOne.info:type_name -> omo.msp.vocabulary.RelationInfo
	4, // 2: omo.msp.vocabulary.ReplyRelationOne.errorCode:type_name -> omo.msp.vocabulary.ResultStatus
	0, // 3: omo.msp.vocabulary.ReplyRelationList.list:type_name -> omo.msp.vocabulary.RelationInfo
	1, // 4: omo.msp.vocabulary.RelationService.AddOne:input_type -> omo.msp.vocabulary.ReqRelationAdd
	5, // 5: omo.msp.vocabulary.RelationService.GetOne:input_type -> omo.msp.vocabulary.RequestInfo
	5, // 6: omo.msp.vocabulary.RelationService.RemoveOne:input_type -> omo.msp.vocabulary.RequestInfo
	5, // 7: omo.msp.vocabulary.RelationService.GetAll:input_type -> omo.msp.vocabulary.RequestInfo
	2, // 8: omo.msp.vocabulary.RelationService.AddOne:output_type -> omo.msp.vocabulary.ReplyRelationOne
	2, // 9: omo.msp.vocabulary.RelationService.GetOne:output_type -> omo.msp.vocabulary.ReplyRelationOne
	6, // 10: omo.msp.vocabulary.RelationService.RemoveOne:output_type -> omo.msp.vocabulary.ReplyInfo
	3, // 11: omo.msp.vocabulary.RelationService.GetAll:output_type -> omo.msp.vocabulary.ReplyRelationList
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_vocabulary_relation_proto_init() }
func file_proto_vocabulary_relation_proto_init() {
	if File_proto_vocabulary_relation_proto != nil {
		return
	}
	file_proto_vocabulary_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_vocabulary_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_vocabulary_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRelationAdd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_vocabulary_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRelationOne); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_vocabulary_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRelationList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_vocabulary_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vocabulary_relation_proto_goTypes,
		DependencyIndexes: file_proto_vocabulary_relation_proto_depIdxs,
		MessageInfos:      file_proto_vocabulary_relation_proto_msgTypes,
	}.Build()
	File_proto_vocabulary_relation_proto = out.File
	file_proto_vocabulary_relation_proto_rawDesc = nil
	file_proto_vocabulary_relation_proto_goTypes = nil
	file_proto_vocabulary_relation_proto_depIdxs = nil
}
