//
//TeamCloud
//
//API for working with a TeamCloud instance.
//
//The version of the OpenAPI document: v1
//
//Contact: colbyw@microsoft.com
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: teamcloud/v1/project_template_definition.proto

package teamcloudv1

import (
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

type ProjectTemplateDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string                `protobuf:"bytes,103536240,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Repository  *RepositoryDefinition `protobuf:"bytes,340187981,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *ProjectTemplateDefinition) Reset() {
	*x = ProjectTemplateDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_project_template_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTemplateDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTemplateDefinition) ProtoMessage() {}

func (x *ProjectTemplateDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_project_template_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTemplateDefinition.ProtoReflect.Descriptor instead.
func (*ProjectTemplateDefinition) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_project_template_definition_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectTemplateDefinition) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ProjectTemplateDefinition) GetRepository() *RepositoryDefinition {
	if x != nil {
		return x.Repository
	}
	return nil
}

var File_teamcloud_v1_project_template_definition_proto protoreflect.FileDescriptor

var file_teamcloud_v1_project_template_definition_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0xf0, 0xac, 0xaf, 0x31, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0xcd, 0xb6, 0x9b, 0xa2, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69,
	0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2f, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x76, 0x31, 0x50, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teamcloud_v1_project_template_definition_proto_rawDescOnce sync.Once
	file_teamcloud_v1_project_template_definition_proto_rawDescData = file_teamcloud_v1_project_template_definition_proto_rawDesc
)

func file_teamcloud_v1_project_template_definition_proto_rawDescGZIP() []byte {
	file_teamcloud_v1_project_template_definition_proto_rawDescOnce.Do(func() {
		file_teamcloud_v1_project_template_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_teamcloud_v1_project_template_definition_proto_rawDescData)
	})
	return file_teamcloud_v1_project_template_definition_proto_rawDescData
}

var file_teamcloud_v1_project_template_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teamcloud_v1_project_template_definition_proto_goTypes = []interface{}{
	(*ProjectTemplateDefinition)(nil), // 0: colbylwilliams.teamcloud.v1.ProjectTemplateDefinition
	(*RepositoryDefinition)(nil),      // 1: colbylwilliams.teamcloud.v1.RepositoryDefinition
}
var file_teamcloud_v1_project_template_definition_proto_depIdxs = []int32{
	1, // 0: colbylwilliams.teamcloud.v1.ProjectTemplateDefinition.repository:type_name -> colbylwilliams.teamcloud.v1.RepositoryDefinition
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teamcloud_v1_project_template_definition_proto_init() }
func file_teamcloud_v1_project_template_definition_proto_init() {
	if File_teamcloud_v1_project_template_definition_proto != nil {
		return
	}
	file_teamcloud_v1_repository_definition_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teamcloud_v1_project_template_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTemplateDefinition); i {
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
			RawDescriptor: file_teamcloud_v1_project_template_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teamcloud_v1_project_template_definition_proto_goTypes,
		DependencyIndexes: file_teamcloud_v1_project_template_definition_proto_depIdxs,
		MessageInfos:      file_teamcloud_v1_project_template_definition_proto_msgTypes,
	}.Build()
	File_teamcloud_v1_project_template_definition_proto = out.File
	file_teamcloud_v1_project_template_definition_proto_rawDesc = nil
	file_teamcloud_v1_project_template_definition_proto_goTypes = nil
	file_teamcloud_v1_project_template_definition_proto_depIdxs = nil
}
