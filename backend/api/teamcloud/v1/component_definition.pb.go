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
// source: teamcloud/v1/component_definition.proto

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

type ComponentDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId        string  `protobuf:"bytes,230268727,opt,name=templateId,proto3" json:"templateId,omitempty"`
	DisplayName       string  `protobuf:"bytes,103536240,opt,name=displayName,proto3" json:"displayName,omitempty"`
	InputJson         *string `protobuf:"bytes,96060373,opt,name=inputJson,proto3,oneof" json:"inputJson,omitempty"`
	DeploymentScopeId *string `protobuf:"bytes,336779257,opt,name=deploymentScopeId,proto3,oneof" json:"deploymentScopeId,omitempty"`
}

func (x *ComponentDefinition) Reset() {
	*x = ComponentDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_component_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentDefinition) ProtoMessage() {}

func (x *ComponentDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_component_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentDefinition.ProtoReflect.Descriptor instead.
func (*ComponentDefinition) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_component_definition_proto_rawDescGZIP(), []int{0}
}

func (x *ComponentDefinition) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *ComponentDefinition) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ComponentDefinition) GetInputJson() string {
	if x != nil && x.InputJson != nil {
		return *x.InputJson
	}
	return ""
}

func (x *ComponentDefinition) GetDeploymentScopeId() string {
	if x != nil && x.DeploymentScopeId != nil {
		return *x.DeploymentScopeId
	}
	return ""
}

var File_teamcloud_v1_component_definition_proto protoreflect.FileDescriptor

var file_teamcloud_v1_component_definition_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6c, 0x62, 0x79,
	0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x22, 0xde, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0xb7, 0xbe, 0xe6,
	0x6d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0xf0, 0xac, 0xaf, 0x31, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a,
	0x73, 0x6f, 0x6e, 0x18, 0xd5, 0x87, 0xe7, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x11,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49,
	0x64, 0x18, 0xf9, 0xaf, 0xcb, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x11, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a, 0x73, 0x6f,
	0x6e, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c,
	0x69, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2d, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teamcloud_v1_component_definition_proto_rawDescOnce sync.Once
	file_teamcloud_v1_component_definition_proto_rawDescData = file_teamcloud_v1_component_definition_proto_rawDesc
)

func file_teamcloud_v1_component_definition_proto_rawDescGZIP() []byte {
	file_teamcloud_v1_component_definition_proto_rawDescOnce.Do(func() {
		file_teamcloud_v1_component_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_teamcloud_v1_component_definition_proto_rawDescData)
	})
	return file_teamcloud_v1_component_definition_proto_rawDescData
}

var file_teamcloud_v1_component_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teamcloud_v1_component_definition_proto_goTypes = []interface{}{
	(*ComponentDefinition)(nil), // 0: colbylwilliams.teamcloud.v1.ComponentDefinition
}
var file_teamcloud_v1_component_definition_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teamcloud_v1_component_definition_proto_init() }
func file_teamcloud_v1_component_definition_proto_init() {
	if File_teamcloud_v1_component_definition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teamcloud_v1_component_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentDefinition); i {
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
	file_teamcloud_v1_component_definition_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teamcloud_v1_component_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teamcloud_v1_component_definition_proto_goTypes,
		DependencyIndexes: file_teamcloud_v1_component_definition_proto_depIdxs,
		MessageInfos:      file_teamcloud_v1_component_definition_proto_msgTypes,
	}.Build()
	File_teamcloud_v1_component_definition_proto = out.File
	file_teamcloud_v1_component_definition_proto_rawDesc = nil
	file_teamcloud_v1_component_definition_proto_goTypes = nil
	file_teamcloud_v1_component_definition_proto_depIdxs = nil
}
