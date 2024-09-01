// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: tag/v1/service_internal.proto

package tag

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTagIDTreeizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to From:
	//
	//	*GetTagIDTreeizeRequest_TagId
	//	*GetTagIDTreeizeRequest_TypeId
	From isGetTagIDTreeizeRequest_From `protobuf_oneof:"from"`
}

func (x *GetTagIDTreeizeRequest) Reset() {
	*x = GetTagIDTreeizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagIDTreeizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagIDTreeizeRequest) ProtoMessage() {}

func (x *GetTagIDTreeizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagIDTreeizeRequest.ProtoReflect.Descriptor instead.
func (*GetTagIDTreeizeRequest) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{0}
}

func (m *GetTagIDTreeizeRequest) GetFrom() isGetTagIDTreeizeRequest_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (x *GetTagIDTreeizeRequest) GetTagId() uint64 {
	if x, ok := x.GetFrom().(*GetTagIDTreeizeRequest_TagId); ok {
		return x.TagId
	}
	return 0
}

func (x *GetTagIDTreeizeRequest) GetTypeId() uint64 {
	if x, ok := x.GetFrom().(*GetTagIDTreeizeRequest_TypeId); ok {
		return x.TypeId
	}
	return 0
}

type isGetTagIDTreeizeRequest_From interface {
	isGetTagIDTreeizeRequest_From()
}

type GetTagIDTreeizeRequest_TagId struct {
	TagId uint64 `protobuf:"fixed64,1,opt,name=tag_id,json=tagId,proto3,oneof"`
}

type GetTagIDTreeizeRequest_TypeId struct {
	TypeId uint64 `protobuf:"fixed64,2,opt,name=type_id,json=typeId,proto3,oneof"`
}

func (*GetTagIDTreeizeRequest_TagId) isGetTagIDTreeizeRequest_From() {}

func (*GetTagIDTreeizeRequest_TypeId) isGetTagIDTreeizeRequest_From() {}

type GetTagIDTreeizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *TagIDTreeize `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type *Type         `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetTagIDTreeizeResponse) Reset() {
	*x = GetTagIDTreeizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagIDTreeizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagIDTreeizeResponse) ProtoMessage() {}

func (x *GetTagIDTreeizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagIDTreeizeResponse.ProtoReflect.Descriptor instead.
func (*GetTagIDTreeizeResponse) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{1}
}

func (x *GetTagIDTreeizeResponse) GetData() *TagIDTreeize {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetTagIDTreeizeResponse) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type TagIDTreeize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64          `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	Child []*TagIDTreeize `protobuf:"bytes,2,rep,name=child,proto3" json:"child,omitempty"`
}

func (x *TagIDTreeize) Reset() {
	*x = TagIDTreeize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagIDTreeize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagIDTreeize) ProtoMessage() {}

func (x *TagIDTreeize) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagIDTreeize.ProtoReflect.Descriptor instead.
func (*TagIDTreeize) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{2}
}

func (x *TagIDTreeize) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TagIDTreeize) GetChild() []*TagIDTreeize {
	if x != nil {
		return x.Child
	}
	return nil
}

type BindRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleId   uint64           `protobuf:"fixed64,1,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	ExternalId uint64           `protobuf:"fixed64,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	TagId      uint64           `protobuf:"fixed64,3,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	Data       *structpb.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BindRelationRequest) Reset() {
	*x = BindRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRelationRequest) ProtoMessage() {}

func (x *BindRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRelationRequest.ProtoReflect.Descriptor instead.
func (*BindRelationRequest) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{3}
}

func (x *BindRelationRequest) GetModuleId() uint64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *BindRelationRequest) GetExternalId() uint64 {
	if x != nil {
		return x.ExternalId
	}
	return 0
}

func (x *BindRelationRequest) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *BindRelationRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type BindRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CleanTagId   []uint64 `protobuf:"fixed64,1,rep,packed,name=clean_tag_id,json=cleanTagId,proto3" json:"clean_tag_id,omitempty"`
	InheritTagId []uint64 `protobuf:"fixed64,2,rep,packed,name=inherit_tag_id,json=inheritTagId,proto3" json:"inherit_tag_id,omitempty"`
}

func (x *BindRelationResponse) Reset() {
	*x = BindRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRelationResponse) ProtoMessage() {}

func (x *BindRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRelationResponse.ProtoReflect.Descriptor instead.
func (*BindRelationResponse) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{4}
}

func (x *BindRelationResponse) GetCleanTagId() []uint64 {
	if x != nil {
		return x.CleanTagId
	}
	return nil
}

func (x *BindRelationResponse) GetInheritTagId() []uint64 {
	if x != nil {
		return x.InheritTagId
	}
	return nil
}

type UnbindRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleId   uint64 `protobuf:"fixed64,1,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	ExternalId uint64 `protobuf:"fixed64,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	TagId      uint64 `protobuf:"fixed64,3,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
}

func (x *UnbindRelationRequest) Reset() {
	*x = UnbindRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRelationRequest) ProtoMessage() {}

func (x *UnbindRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRelationRequest.ProtoReflect.Descriptor instead.
func (*UnbindRelationRequest) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{5}
}

func (x *UnbindRelationRequest) GetModuleId() uint64 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *UnbindRelationRequest) GetExternalId() uint64 {
	if x != nil {
		return x.ExternalId
	}
	return 0
}

func (x *UnbindRelationRequest) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

type UnbindRelationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CleanTagId []uint64 `protobuf:"fixed64,1,rep,packed,name=clean_tag_id,json=cleanTagId,proto3" json:"clean_tag_id,omitempty"`
}

func (x *UnbindRelationResponse) Reset() {
	*x = UnbindRelationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_service_internal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRelationResponse) ProtoMessage() {}

func (x *UnbindRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_service_internal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRelationResponse.ProtoReflect.Descriptor instead.
func (*UnbindRelationResponse) Descriptor() ([]byte, []int) {
	return file_tag_v1_service_internal_proto_rawDescGZIP(), []int{6}
}

func (x *UnbindRelationResponse) GetCleanTagId() []uint64 {
	if x != nil {
		return x.CleanTagId
	}
	return nil
}

var File_tag_v1_service_internal_proto protoreflect.FileDescriptor

var file_tag_v1_service_internal_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x49, 0x44, 0x54, 0x72, 0x65, 0x65, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x06, 0x48, 0x00, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x48, 0x00, 0x52, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x65,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x54, 0x72, 0x65, 0x65, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x67, 0x49, 0x44, 0x54, 0x72, 0x65, 0x65, 0x69, 0x7a, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x0c, 0x54, 0x61, 0x67, 0x49, 0x44, 0x54, 0x72,
	0x65, 0x65, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x67, 0x49, 0x44, 0x54, 0x72, 0x65, 0x65, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x14, 0x42,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e,
	0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74,
	0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0c, 0x69,
	0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x54, 0x61, 0x67, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x15, 0x55,
	0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x16, 0x55, 0x6e, 0x62,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0a, 0x63, 0x6c, 0x65, 0x61, 0x6e,
	0x54, 0x61, 0x67, 0x49, 0x64, 0x32, 0x80, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x12, 0x54, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x54, 0x72,
	0x65, 0x65, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x54, 0x72, 0x65, 0x65, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x49, 0x44, 0x54, 0x72, 0x65, 0x65, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x70, 0x62, 0x2f, 0x74,
	0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tag_v1_service_internal_proto_rawDescOnce sync.Once
	file_tag_v1_service_internal_proto_rawDescData = file_tag_v1_service_internal_proto_rawDesc
)

func file_tag_v1_service_internal_proto_rawDescGZIP() []byte {
	file_tag_v1_service_internal_proto_rawDescOnce.Do(func() {
		file_tag_v1_service_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_v1_service_internal_proto_rawDescData)
	})
	return file_tag_v1_service_internal_proto_rawDescData
}

var file_tag_v1_service_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tag_v1_service_internal_proto_goTypes = []any{
	(*GetTagIDTreeizeRequest)(nil),  // 0: tag.v1.GetTagIDTreeizeRequest
	(*GetTagIDTreeizeResponse)(nil), // 1: tag.v1.GetTagIDTreeizeResponse
	(*TagIDTreeize)(nil),            // 2: tag.v1.TagIDTreeize
	(*BindRelationRequest)(nil),     // 3: tag.v1.BindRelationRequest
	(*BindRelationResponse)(nil),    // 4: tag.v1.BindRelationResponse
	(*UnbindRelationRequest)(nil),   // 5: tag.v1.UnbindRelationRequest
	(*UnbindRelationResponse)(nil),  // 6: tag.v1.UnbindRelationResponse
	(*Type)(nil),                    // 7: tag.v1.Type
	(*structpb.Struct)(nil),         // 8: google.protobuf.Struct
}
var file_tag_v1_service_internal_proto_depIdxs = []int32{
	2, // 0: tag.v1.GetTagIDTreeizeResponse.data:type_name -> tag.v1.TagIDTreeize
	7, // 1: tag.v1.GetTagIDTreeizeResponse.type:type_name -> tag.v1.Type
	2, // 2: tag.v1.TagIDTreeize.child:type_name -> tag.v1.TagIDTreeize
	8, // 3: tag.v1.BindRelationRequest.data:type_name -> google.protobuf.Struct
	0, // 4: tag.v1.Internal.GetTagIDTreeize:input_type -> tag.v1.GetTagIDTreeizeRequest
	3, // 5: tag.v1.Internal.BindRelation:input_type -> tag.v1.BindRelationRequest
	5, // 6: tag.v1.Internal.UnbindRelation:input_type -> tag.v1.UnbindRelationRequest
	1, // 7: tag.v1.Internal.GetTagIDTreeize:output_type -> tag.v1.GetTagIDTreeizeResponse
	4, // 8: tag.v1.Internal.BindRelation:output_type -> tag.v1.BindRelationResponse
	6, // 9: tag.v1.Internal.UnbindRelation:output_type -> tag.v1.UnbindRelationResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tag_v1_service_internal_proto_init() }
func file_tag_v1_service_internal_proto_init() {
	if File_tag_v1_service_internal_proto != nil {
		return
	}
	file_tag_v1_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tag_v1_service_internal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTagIDTreeizeRequest); i {
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
		file_tag_v1_service_internal_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTagIDTreeizeResponse); i {
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
		file_tag_v1_service_internal_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TagIDTreeize); i {
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
		file_tag_v1_service_internal_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BindRelationRequest); i {
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
		file_tag_v1_service_internal_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BindRelationResponse); i {
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
		file_tag_v1_service_internal_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UnbindRelationRequest); i {
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
		file_tag_v1_service_internal_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UnbindRelationResponse); i {
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
	file_tag_v1_service_internal_proto_msgTypes[0].OneofWrappers = []any{
		(*GetTagIDTreeizeRequest_TagId)(nil),
		(*GetTagIDTreeizeRequest_TypeId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tag_v1_service_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tag_v1_service_internal_proto_goTypes,
		DependencyIndexes: file_tag_v1_service_internal_proto_depIdxs,
		MessageInfos:      file_tag_v1_service_internal_proto_msgTypes,
	}.Build()
	File_tag_v1_service_internal_proto = out.File
	file_tag_v1_service_internal_proto_rawDesc = nil
	file_tag_v1_service_internal_proto_goTypes = nil
	file_tag_v1_service_internal_proto_depIdxs = nil
}
