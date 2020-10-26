// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: models/relation.proto

package models

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GRPCRelationsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32  `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
}

func (x *GRPCRelationsReadRequest) Reset() {
	*x = GRPCRelationsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCRelationsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCRelationsReadRequest) ProtoMessage() {}

func (x *GRPCRelationsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCRelationsReadRequest.ProtoReflect.Descriptor instead.
func (*GRPCRelationsReadRequest) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{0}
}

func (x *GRPCRelationsReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GRPCRelationsReadRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GRPCRelationsReadRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type GRPCRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectID string `protobuf:"bytes,1,opt,name=subjectID,proto3" json:"subjectID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ObjectID  string `protobuf:"bytes,3,opt,name=objectID,proto3" json:"objectID,omitempty"`
}

func (x *GRPCRelation) Reset() {
	*x = GRPCRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCRelation) ProtoMessage() {}

func (x *GRPCRelation) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCRelation.ProtoReflect.Descriptor instead.
func (*GRPCRelation) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{1}
}

func (x *GRPCRelation) GetSubjectID() string {
	if x != nil {
		return x.SubjectID
	}
	return ""
}

func (x *GRPCRelation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GRPCRelation) GetObjectID() string {
	if x != nil {
		return x.ObjectID
	}
	return ""
}

type GRPCRelationsReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations []*GRPCRelation `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (x *GRPCRelationsReadResponse) Reset() {
	*x = GRPCRelationsReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCRelationsReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCRelationsReadResponse) ProtoMessage() {}

func (x *GRPCRelationsReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCRelationsReadResponse.ProtoReflect.Descriptor instead.
func (*GRPCRelationsReadResponse) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GRPCRelationsReadResponse) GetRelations() []*GRPCRelation {
	if x != nil {
		return x.Relations
	}
	return nil
}

type GRPCRelationsWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GRPCRelationsWriteResponse) Reset() {
	*x = GRPCRelationsWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GRPCRelationsWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GRPCRelationsWriteResponse) ProtoMessage() {}

func (x *GRPCRelationsWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GRPCRelationsWriteResponse.ProtoReflect.Descriptor instead.
func (*GRPCRelationsWriteResponse) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{3}
}

var File_models_relation_proto protoreflect.FileDescriptor

var file_models_relation_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22,
	0x58, 0x0a, 0x18, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x0c, 0x47, 0x52, 0x50,
	0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x19, 0x47, 0x52, 0x50, 0x43, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x52, 0x50, 0x43,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc6, 0x01, 0x0a, 0x12, 0x47, 0x52, 0x50, 0x43, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x56, 0x0a,
	0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52, 0x50, 0x43,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x5f, 0x0a, 0x12, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_relation_proto_rawDescOnce sync.Once
	file_models_relation_proto_rawDescData = file_models_relation_proto_rawDesc
)

func file_models_relation_proto_rawDescGZIP() []byte {
	file_models_relation_proto_rawDescOnce.Do(func() {
		file_models_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_relation_proto_rawDescData)
	})
	return file_models_relation_proto_rawDescData
}

var file_models_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_models_relation_proto_goTypes = []interface{}{
	(*GRPCRelationsReadRequest)(nil),   // 0: models.GRPCRelationsReadRequest
	(*GRPCRelation)(nil),               // 1: models.GRPCRelation
	(*GRPCRelationsReadResponse)(nil),  // 2: models.GRPCRelationsReadResponse
	(*GRPCRelationsWriteResponse)(nil), // 3: models.GRPCRelationsWriteResponse
}
var file_models_relation_proto_depIdxs = []int32{
	1, // 0: models.GRPCRelationsReadResponse.relations:type_name -> models.GRPCRelation
	0, // 1: models.GRPCRelationReader.RelationsByUser:input_type -> models.GRPCRelationsReadRequest
	0, // 2: models.GRPCRelationReader.RelationsByObject:input_type -> models.GRPCRelationsReadRequest
	1, // 3: models.GRPCRelationWriter.WriteRelation:input_type -> models.GRPCRelation
	2, // 4: models.GRPCRelationReader.RelationsByUser:output_type -> models.GRPCRelationsReadResponse
	2, // 5: models.GRPCRelationReader.RelationsByObject:output_type -> models.GRPCRelationsReadResponse
	3, // 6: models.GRPCRelationWriter.WriteRelation:output_type -> models.GRPCRelationsWriteResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_relation_proto_init() }
func file_models_relation_proto_init() {
	if File_models_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCRelationsReadRequest); i {
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
		file_models_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCRelation); i {
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
		file_models_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCRelationsReadResponse); i {
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
		file_models_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GRPCRelationsWriteResponse); i {
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
			RawDescriptor: file_models_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_models_relation_proto_goTypes,
		DependencyIndexes: file_models_relation_proto_depIdxs,
		MessageInfos:      file_models_relation_proto_msgTypes,
	}.Build()
	File_models_relation_proto = out.File
	file_models_relation_proto_rawDesc = nil
	file_models_relation_proto_goTypes = nil
	file_models_relation_proto_depIdxs = nil
}
