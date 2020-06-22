// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: proto/vocabulary/concept.proto

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

type ConceptType int32

const (
	ConceptType_Unknown  ConceptType = 0
	ConceptType_Personal ConceptType = 1
	ConceptType_Utensil  ConceptType = 2  // 器物
	ConceptType_Event    ConceptType = 3  //事件
	ConceptType_Organize ConceptType = 4  // 组织
	ConceptType_Idea     ConceptType = 5  //思想理论
	ConceptType_Book     ConceptType = 6  //经籍著作
	ConceptType_Culture  ConceptType = 7  //文化
	ConceptType_Faction  ConceptType = 8  //派别
	ConceptType_Nature   ConceptType = 9  //自然
	ConceptType_Honor    ConceptType = 10 //荣誉奖项
)

// Enum value maps for ConceptType.
var (
	ConceptType_name = map[int32]string{
		0:  "Unknown",
		1:  "Personal",
		2:  "Utensil",
		3:  "Event",
		4:  "Organize",
		5:  "Idea",
		6:  "Book",
		7:  "Culture",
		8:  "Faction",
		9:  "Nature",
		10: "Honor",
	}
	ConceptType_value = map[string]int32{
		"Unknown":  0,
		"Personal": 1,
		"Utensil":  2,
		"Event":    3,
		"Organize": 4,
		"Idea":     5,
		"Book":     6,
		"Culture":  7,
		"Faction":  8,
		"Nature":   9,
		"Honor":    10,
	}
)

func (x ConceptType) Enum() *ConceptType {
	p := new(ConceptType)
	*p = x
	return p
}

func (x ConceptType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConceptType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vocabulary_concept_proto_enumTypes[0].Descriptor()
}

func (ConceptType) Type() protoreflect.EnumType {
	return &file_proto_vocabulary_concept_proto_enumTypes[0]
}

func (x ConceptType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConceptType.Descriptor instead.
func (ConceptType) EnumDescriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{0}
}

type ConceptInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       ConceptType      `protobuf:"varint,1,opt,name=type,proto3,enum=omo.msp.vocabulary.ConceptType" json:"type,omitempty"`
	Name       string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Uid        string           `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cover      string           `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark     string           `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Table      string           `protobuf:"bytes,6,opt,name=table,proto3" json:"table,omitempty"`
	Children   []*ConceptInfo   `protobuf:"bytes,7,rep,name=children,proto3" json:"children,omitempty"`
	Attributes []*AttributeInfo `protobuf:"bytes,8,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *ConceptInfo) Reset() {
	*x = ConceptInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_concept_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptInfo) ProtoMessage() {}

func (x *ConceptInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_concept_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptInfo.ProtoReflect.Descriptor instead.
func (*ConceptInfo) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{0}
}

func (x *ConceptInfo) GetType() ConceptType {
	if x != nil {
		return x.Type
	}
	return ConceptType_Unknown
}

func (x *ConceptInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConceptInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ConceptInfo) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *ConceptInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConceptInfo) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *ConceptInfo) GetChildren() []*ConceptInfo {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *ConceptInfo) GetAttributes() []*AttributeInfo {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type ReqConceptAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   ConceptType `protobuf:"varint,1,opt,name=type,proto3,enum=omo.msp.vocabulary.ConceptType" json:"type,omitempty"`
	Name   string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Uid    string      `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cover  string      `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark string      `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Table  string      `protobuf:"bytes,6,opt,name=table,proto3" json:"table,omitempty"`
	Parent string      `protobuf:"bytes,7,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ReqConceptAdd) Reset() {
	*x = ReqConceptAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_concept_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqConceptAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqConceptAdd) ProtoMessage() {}

func (x *ReqConceptAdd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_concept_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqConceptAdd.ProtoReflect.Descriptor instead.
func (*ReqConceptAdd) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{1}
}

func (x *ReqConceptAdd) GetType() ConceptType {
	if x != nil {
		return x.Type
	}
	return ConceptType_Unknown
}

func (x *ReqConceptAdd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqConceptAdd) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReqConceptAdd) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *ReqConceptAdd) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ReqConceptAdd) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *ReqConceptAdd) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

type ReplyConceptInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *ConceptInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	ErrorCode ResultStatus `protobuf:"varint,2,opt,name=errorCode,proto3,enum=omo.msp.vocabulary.ResultStatus" json:"errorCode,omitempty"`
	ErrorMsg  string       `protobuf:"bytes,3,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *ReplyConceptInfo) Reset() {
	*x = ReplyConceptInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_concept_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyConceptInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyConceptInfo) ProtoMessage() {}

func (x *ReplyConceptInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_concept_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyConceptInfo.ProtoReflect.Descriptor instead.
func (*ReplyConceptInfo) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyConceptInfo) GetInfo() *ConceptInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReplyConceptInfo) GetErrorCode() ResultStatus {
	if x != nil {
		return x.ErrorCode
	}
	return ResultStatus_Success
}

func (x *ReplyConceptInfo) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type ReplyConceptList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag string         `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	List []*ConceptInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ReplyConceptList) Reset() {
	*x = ReplyConceptList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_concept_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyConceptList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyConceptList) ProtoMessage() {}

func (x *ReplyConceptList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_concept_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyConceptList.ProtoReflect.Descriptor instead.
func (*ReplyConceptList) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyConceptList) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *ReplyConceptList) GetList() []*ConceptInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type ReqConceptAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Concept string `protobuf:"bytes,2,opt,name=concept,proto3" json:"concept,omitempty"`
}

func (x *ReqConceptAttribute) Reset() {
	*x = ReqConceptAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_concept_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqConceptAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqConceptAttribute) ProtoMessage() {}

func (x *ReqConceptAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_concept_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqConceptAttribute.ProtoReflect.Descriptor instead.
func (*ReqConceptAttribute) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{4}
}

func (x *ReqConceptAttribute) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReqConceptAttribute) GetConcept() string {
	if x != nil {
		return x.Concept
	}
	return ""
}

type ReplyConceptAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Concept   string       `protobuf:"bytes,2,opt,name=concept,proto3" json:"concept,omitempty"`
	ErrorCode ResultStatus `protobuf:"varint,3,opt,name=errorCode,proto3,enum=omo.msp.vocabulary.ResultStatus" json:"errorCode,omitempty"`
}

func (x *ReplyConceptAttribute) Reset() {
	*x = ReplyConceptAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vocabulary_concept_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyConceptAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyConceptAttribute) ProtoMessage() {}

func (x *ReplyConceptAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vocabulary_concept_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyConceptAttribute.ProtoReflect.Descriptor instead.
func (*ReplyConceptAttribute) Descriptor() ([]byte, []int) {
	return file_proto_vocabulary_concept_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyConceptAttribute) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReplyConceptAttribute) GetConcept() string {
	if x != nil {
		return x.Concept
	}
	return ""
}

func (x *ReplyConceptAttribute) GetErrorCode() ResultStatus {
	if x != nil {
		return x.ErrorCode
	}
	return ResultStatus_Success
}

var File_proto_vocabulary_concept_proto protoreflect.FileDescriptor

var file_proto_vocabulary_concept_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61,
	0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61,
	0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76,
	0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12,
	0x41, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f,
	0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x41, 0x64, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63,
	0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x10,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x33, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c,
	0x61, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x22, 0x5b, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e,
	0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x41,
	0x0a, 0x13, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x93, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6c, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x64, 0x65, 0x61,
	0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x6f, 0x6e, 0x6f, 0x72, 0x10, 0x0a, 0x32, 0xac, 0x04,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x21, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e,
	0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x41, 0x64, 0x64, 0x1a, 0x24, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12,
	0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62,
	0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61,
	0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63,
	0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e,
	0x63, 0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0f, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x27, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x29, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a,
	0x29, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vocabulary_concept_proto_rawDescOnce sync.Once
	file_proto_vocabulary_concept_proto_rawDescData = file_proto_vocabulary_concept_proto_rawDesc
)

func file_proto_vocabulary_concept_proto_rawDescGZIP() []byte {
	file_proto_vocabulary_concept_proto_rawDescOnce.Do(func() {
		file_proto_vocabulary_concept_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vocabulary_concept_proto_rawDescData)
	})
	return file_proto_vocabulary_concept_proto_rawDescData
}

var file_proto_vocabulary_concept_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_vocabulary_concept_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_vocabulary_concept_proto_goTypes = []interface{}{
	(ConceptType)(0),              // 0: omo.msp.vocabulary.ConceptType
	(*ConceptInfo)(nil),           // 1: omo.msp.vocabulary.ConceptInfo
	(*ReqConceptAdd)(nil),         // 2: omo.msp.vocabulary.ReqConceptAdd
	(*ReplyConceptInfo)(nil),      // 3: omo.msp.vocabulary.ReplyConceptInfo
	(*ReplyConceptList)(nil),      // 4: omo.msp.vocabulary.ReplyConceptList
	(*ReqConceptAttribute)(nil),   // 5: omo.msp.vocabulary.ReqConceptAttribute
	(*ReplyConceptAttribute)(nil), // 6: omo.msp.vocabulary.ReplyConceptAttribute
	(*AttributeInfo)(nil),         // 7: omo.msp.vocabulary.AttributeInfo
	(ResultStatus)(0),             // 8: omo.msp.vocabulary.ResultStatus
	(*RequestInfo)(nil),           // 9: omo.msp.vocabulary.RequestInfo
	(*ReplyInfo)(nil),             // 10: omo.msp.vocabulary.ReplyInfo
}
var file_proto_vocabulary_concept_proto_depIdxs = []int32{
	0,  // 0: omo.msp.vocabulary.ConceptInfo.type:type_name -> omo.msp.vocabulary.ConceptType
	1,  // 1: omo.msp.vocabulary.ConceptInfo.children:type_name -> omo.msp.vocabulary.ConceptInfo
	7,  // 2: omo.msp.vocabulary.ConceptInfo.attributes:type_name -> omo.msp.vocabulary.AttributeInfo
	0,  // 3: omo.msp.vocabulary.ReqConceptAdd.type:type_name -> omo.msp.vocabulary.ConceptType
	1,  // 4: omo.msp.vocabulary.ReplyConceptInfo.info:type_name -> omo.msp.vocabulary.ConceptInfo
	8,  // 5: omo.msp.vocabulary.ReplyConceptInfo.errorCode:type_name -> omo.msp.vocabulary.ResultStatus
	1,  // 6: omo.msp.vocabulary.ReplyConceptList.list:type_name -> omo.msp.vocabulary.ConceptInfo
	8,  // 7: omo.msp.vocabulary.ReplyConceptAttribute.errorCode:type_name -> omo.msp.vocabulary.ResultStatus
	2,  // 8: omo.msp.vocabulary.ConceptService.AddOne:input_type -> omo.msp.vocabulary.ReqConceptAdd
	9,  // 9: omo.msp.vocabulary.ConceptService.GetOne:input_type -> omo.msp.vocabulary.RequestInfo
	9,  // 10: omo.msp.vocabulary.ConceptService.RemoveOne:input_type -> omo.msp.vocabulary.RequestInfo
	9,  // 11: omo.msp.vocabulary.ConceptService.GetAll:input_type -> omo.msp.vocabulary.RequestInfo
	5,  // 12: omo.msp.vocabulary.ConceptService.AppendAttribute:input_type -> omo.msp.vocabulary.ReqConceptAttribute
	5,  // 13: omo.msp.vocabulary.ConceptService.RemoveAttribute:input_type -> omo.msp.vocabulary.ReqConceptAttribute
	3,  // 14: omo.msp.vocabulary.ConceptService.AddOne:output_type -> omo.msp.vocabulary.ReplyConceptInfo
	3,  // 15: omo.msp.vocabulary.ConceptService.GetOne:output_type -> omo.msp.vocabulary.ReplyConceptInfo
	10, // 16: omo.msp.vocabulary.ConceptService.RemoveOne:output_type -> omo.msp.vocabulary.ReplyInfo
	4,  // 17: omo.msp.vocabulary.ConceptService.GetAll:output_type -> omo.msp.vocabulary.ReplyConceptList
	6,  // 18: omo.msp.vocabulary.ConceptService.AppendAttribute:output_type -> omo.msp.vocabulary.ReplyConceptAttribute
	6,  // 19: omo.msp.vocabulary.ConceptService.RemoveAttribute:output_type -> omo.msp.vocabulary.ReplyConceptAttribute
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_vocabulary_concept_proto_init() }
func file_proto_vocabulary_concept_proto_init() {
	if File_proto_vocabulary_concept_proto != nil {
		return
	}
	file_proto_vocabulary_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_vocabulary_concept_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptInfo); i {
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
		file_proto_vocabulary_concept_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqConceptAdd); i {
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
		file_proto_vocabulary_concept_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyConceptInfo); i {
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
		file_proto_vocabulary_concept_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyConceptList); i {
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
		file_proto_vocabulary_concept_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqConceptAttribute); i {
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
		file_proto_vocabulary_concept_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyConceptAttribute); i {
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
			RawDescriptor: file_proto_vocabulary_concept_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vocabulary_concept_proto_goTypes,
		DependencyIndexes: file_proto_vocabulary_concept_proto_depIdxs,
		EnumInfos:         file_proto_vocabulary_concept_proto_enumTypes,
		MessageInfos:      file_proto_vocabulary_concept_proto_msgTypes,
	}.Build()
	File_proto_vocabulary_concept_proto = out.File
	file_proto_vocabulary_concept_proto_rawDesc = nil
	file_proto_vocabulary_concept_proto_goTypes = nil
	file_proto_vocabulary_concept_proto_depIdxs = nil
}