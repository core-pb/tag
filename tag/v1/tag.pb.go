// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: tag/v1/tag.proto

package tag

import (
	v1 "github.com/core-pb/dt/time/v1"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64           `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty" bun:",pk,autoincrement"`
	Key       string           `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" bun:",unique"`
	TypeId    uint64           `protobuf:"fixed64,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	ParentId  uint64           `protobuf:"fixed64,4,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Data      *structpb.Struct `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty" bun:"type:jsonb"`
	Info      *structpb.Struct `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty" bun:"type:jsonb"`
	CreatedAt *v1.Time         `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bun:"type:timestamptz"`
	UpdatedAt *v1.Time         `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bun:"type:timestamptz"`
	DeletedAt *v1.Time         `protobuf:"bytes,16,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty" bun:"type:timestamptz,soft_delete,nullzero"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tag_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_tag_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_tag_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Tag) GetTypeId() uint64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Tag) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Tag) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Tag) GetInfo() *structpb.Struct {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Tag) GetCreatedAt() *v1.Time {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Tag) GetUpdatedAt() *v1.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Tag) GetDeletedAt() *v1.Time {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_tag_v1_tag_proto protoreflect.FileDescriptor

var file_tag_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x74, 0x61, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x1c, 0x9a, 0x84, 0x9e, 0x03, 0x17, 0x62, 0x75, 0x6e, 0x3a,
	0x22, 0x2c, 0x70, 0x6b, 0x2c, 0x61, 0x75, 0x74, 0x6f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x62, 0x75, 0x6e, 0x3a, 0x22,
	0x2c, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x22, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06, 0x52, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03,
	0x10, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x6a, 0x73, 0x6f, 0x6e, 0x62,
	0x22, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x15,
	0x9a, 0x84, 0x9e, 0x03, 0x10, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x6a,
	0x73, 0x6f, 0x6e, 0x62, 0x22, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1b,
	0x9a, 0x84, 0x9e, 0x03, 0x16, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x74, 0x7a, 0x22, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x49, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x16,
	0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x74, 0x7a, 0x22, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x63, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x30, 0x9a, 0x84, 0x9e, 0x03, 0x2b, 0x62, 0x75, 0x6e, 0x3a, 0x22,
	0x74, 0x79, 0x70, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x74, 0x7a,
	0x2c, 0x73, 0x6f, 0x66, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2c, 0x6e, 0x75, 0x6c,
	0x6c, 0x7a, 0x65, 0x72, 0x6f, 0x22, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x67, 0x2f,
	0x74, 0x61, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tag_v1_tag_proto_rawDescOnce sync.Once
	file_tag_v1_tag_proto_rawDescData = file_tag_v1_tag_proto_rawDesc
)

func file_tag_v1_tag_proto_rawDescGZIP() []byte {
	file_tag_v1_tag_proto_rawDescOnce.Do(func() {
		file_tag_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_tag_v1_tag_proto_rawDescData)
	})
	return file_tag_v1_tag_proto_rawDescData
}

var file_tag_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tag_v1_tag_proto_goTypes = []any{
	(*Tag)(nil),             // 0: tag.v1.Tag
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
	(*v1.Time)(nil),         // 2: time.v1.Time
}
var file_tag_v1_tag_proto_depIdxs = []int32{
	1, // 0: tag.v1.Tag.data:type_name -> google.protobuf.Struct
	1, // 1: tag.v1.Tag.info:type_name -> google.protobuf.Struct
	2, // 2: tag.v1.Tag.created_at:type_name -> time.v1.Time
	2, // 3: tag.v1.Tag.updated_at:type_name -> time.v1.Time
	2, // 4: tag.v1.Tag.deleted_at:type_name -> time.v1.Time
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tag_v1_tag_proto_init() }
func file_tag_v1_tag_proto_init() {
	if File_tag_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tag_v1_tag_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Tag); i {
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
	file_tag_v1_tag_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tag_v1_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tag_v1_tag_proto_goTypes,
		DependencyIndexes: file_tag_v1_tag_proto_depIdxs,
		MessageInfos:      file_tag_v1_tag_proto_msgTypes,
	}.Build()
	File_tag_v1_tag_proto = out.File
	file_tag_v1_tag_proto_rawDesc = nil
	file_tag_v1_tag_proto_goTypes = nil
	file_tag_v1_tag_proto_depIdxs = nil
}
