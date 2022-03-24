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
// source: teamcloud/v1/component_template.proto

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

type ComponentTemplate_TypeEnum int32

const (
	ComponentTemplate_ENVIRONMENT ComponentTemplate_TypeEnum = 0
	ComponentTemplate_REPOSITORY  ComponentTemplate_TypeEnum = 1
	ComponentTemplate_NAMESPACE   ComponentTemplate_TypeEnum = 2
)

// Enum value maps for ComponentTemplate_TypeEnum.
var (
	ComponentTemplate_TypeEnum_name = map[int32]string{
		0: "ENVIRONMENT",
		1: "REPOSITORY",
		2: "NAMESPACE",
	}
	ComponentTemplate_TypeEnum_value = map[string]int32{
		"ENVIRONMENT": 0,
		"REPOSITORY":  1,
		"NAMESPACE":   2,
	}
)

func (x ComponentTemplate_TypeEnum) Enum() *ComponentTemplate_TypeEnum {
	p := new(ComponentTemplate_TypeEnum)
	*p = x
	return p
}

func (x ComponentTemplate_TypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComponentTemplate_TypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_teamcloud_v1_component_template_proto_enumTypes[0].Descriptor()
}

func (ComponentTemplate_TypeEnum) Type() protoreflect.EnumType {
	return &file_teamcloud_v1_component_template_proto_enumTypes[0]
}

func (x ComponentTemplate_TypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComponentTemplate_TypeEnum.Descriptor instead.
func (ComponentTemplate_TypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_teamcloud_v1_component_template_proto_rawDescGZIP(), []int{0, 0}
}

type ComponentTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization     string                        `protobuf:"bytes,105180469,opt,name=organization,proto3" json:"organization,omitempty"`
	OrganizationName string                        `protobuf:"bytes,218494373,opt,name=organizationName,proto3" json:"organizationName,omitempty"`
	ParentId         string                        `protobuf:"bytes,101420903,opt,name=parentId,proto3" json:"parentId,omitempty"`
	DisplayName      *string                       `protobuf:"bytes,103536240,opt,name=displayName,proto3,oneof" json:"displayName,omitempty"`
	Description      *string                       `protobuf:"bytes,113933319,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Repository       *RepositoryReference          `protobuf:"bytes,340187981,opt,name=repository,proto3" json:"repository,omitempty"`
	Permissions      *ComponentTemplatePermissions `protobuf:"bytes,59962502,opt,name=permissions,proto3" json:"permissions,omitempty"`
	InputJsonSchema  *string                       `protobuf:"bytes,297910254,opt,name=inputJsonSchema,proto3,oneof" json:"inputJsonSchema,omitempty"`
	Tasks            []*ComponentTaskTemplate      `protobuf:"bytes,110132110,rep,name=tasks,proto3" json:"tasks,omitempty"`
	TaskRunner       *ComponentTaskRunner          `protobuf:"bytes,405342486,opt,name=taskRunner,proto3" json:"taskRunner,omitempty"`
	Type             ComponentTemplate_TypeEnum    `protobuf:"varint,3575610,opt,name=type,proto3,enum=colbylwilliams.teamcloud.v1.ComponentTemplate_TypeEnum" json:"type,omitempty"`
	Folder           *string                       `protobuf:"bytes,195224468,opt,name=folder,proto3,oneof" json:"folder,omitempty"`
	Configuration    *string                       `protobuf:"bytes,322139385,opt,name=configuration,proto3,oneof" json:"configuration,omitempty"`
	Id               string                        `protobuf:"bytes,3355,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ComponentTemplate) Reset() {
	*x = ComponentTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_component_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentTemplate) ProtoMessage() {}

func (x *ComponentTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_component_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentTemplate.ProtoReflect.Descriptor instead.
func (*ComponentTemplate) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_component_template_proto_rawDescGZIP(), []int{0}
}

func (x *ComponentTemplate) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *ComponentTemplate) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *ComponentTemplate) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ComponentTemplate) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *ComponentTemplate) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ComponentTemplate) GetRepository() *RepositoryReference {
	if x != nil {
		return x.Repository
	}
	return nil
}

func (x *ComponentTemplate) GetPermissions() *ComponentTemplatePermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *ComponentTemplate) GetInputJsonSchema() string {
	if x != nil && x.InputJsonSchema != nil {
		return *x.InputJsonSchema
	}
	return ""
}

func (x *ComponentTemplate) GetTasks() []*ComponentTaskTemplate {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *ComponentTemplate) GetTaskRunner() *ComponentTaskRunner {
	if x != nil {
		return x.TaskRunner
	}
	return nil
}

func (x *ComponentTemplate) GetType() ComponentTemplate_TypeEnum {
	if x != nil {
		return x.Type
	}
	return ComponentTemplate_ENVIRONMENT
}

func (x *ComponentTemplate) GetFolder() string {
	if x != nil && x.Folder != nil {
		return *x.Folder
	}
	return ""
}

func (x *ComponentTemplate) GetConfiguration() string {
	if x != nil && x.Configuration != nil {
		return *x.Configuration
	}
	return ""
}

func (x *ComponentTemplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_teamcloud_v1_component_template_proto protoreflect.FileDescriptor

var file_teamcloud_v1_component_template_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77,
	0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x07, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xb5, 0xda, 0x93,
	0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0xa5, 0xeb, 0x97, 0x68, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0xe7,
	0x9e, 0xae, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0xf0, 0xac, 0xaf, 0x31, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x87, 0xf8, 0xa9, 0x36, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0xcd, 0xb6, 0x9b, 0xa2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x5e, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x86, 0xe9, 0xcb, 0x1c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c,
	0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x0f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0xee,
	0xff, 0x86, 0x8e, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x4b,
	0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x8e, 0xf7, 0xc1, 0x34, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d,
	0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x54, 0x0a, 0x0a, 0x74,
	0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x96, 0x92, 0xa4, 0xc1, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c,
	0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x4e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0xba, 0x9e, 0xda, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69,
	0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x94, 0xc7, 0x8b, 0x5d,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0xf9, 0xe9, 0xcd, 0x99, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x9b, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f, 0x0a,
	0x0b, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x59,
	0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c,
	0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65,
	0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x76, 0x31, 0x50, 0x00, 0x50, 0x01, 0x50, 0x02, 0x50,
	0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teamcloud_v1_component_template_proto_rawDescOnce sync.Once
	file_teamcloud_v1_component_template_proto_rawDescData = file_teamcloud_v1_component_template_proto_rawDesc
)

func file_teamcloud_v1_component_template_proto_rawDescGZIP() []byte {
	file_teamcloud_v1_component_template_proto_rawDescOnce.Do(func() {
		file_teamcloud_v1_component_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_teamcloud_v1_component_template_proto_rawDescData)
	})
	return file_teamcloud_v1_component_template_proto_rawDescData
}

var file_teamcloud_v1_component_template_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_teamcloud_v1_component_template_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teamcloud_v1_component_template_proto_goTypes = []interface{}{
	(ComponentTemplate_TypeEnum)(0),      // 0: colbylwilliams.teamcloud.v1.ComponentTemplate.TypeEnum
	(*ComponentTemplate)(nil),            // 1: colbylwilliams.teamcloud.v1.ComponentTemplate
	(*RepositoryReference)(nil),          // 2: colbylwilliams.teamcloud.v1.RepositoryReference
	(*ComponentTemplatePermissions)(nil), // 3: colbylwilliams.teamcloud.v1.ComponentTemplatePermissions
	(*ComponentTaskTemplate)(nil),        // 4: colbylwilliams.teamcloud.v1.ComponentTaskTemplate
	(*ComponentTaskRunner)(nil),          // 5: colbylwilliams.teamcloud.v1.ComponentTaskRunner
}
var file_teamcloud_v1_component_template_proto_depIdxs = []int32{
	2, // 0: colbylwilliams.teamcloud.v1.ComponentTemplate.repository:type_name -> colbylwilliams.teamcloud.v1.RepositoryReference
	3, // 1: colbylwilliams.teamcloud.v1.ComponentTemplate.permissions:type_name -> colbylwilliams.teamcloud.v1.ComponentTemplatePermissions
	4, // 2: colbylwilliams.teamcloud.v1.ComponentTemplate.tasks:type_name -> colbylwilliams.teamcloud.v1.ComponentTaskTemplate
	5, // 3: colbylwilliams.teamcloud.v1.ComponentTemplate.taskRunner:type_name -> colbylwilliams.teamcloud.v1.ComponentTaskRunner
	0, // 4: colbylwilliams.teamcloud.v1.ComponentTemplate.type:type_name -> colbylwilliams.teamcloud.v1.ComponentTemplate.TypeEnum
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_teamcloud_v1_component_template_proto_init() }
func file_teamcloud_v1_component_template_proto_init() {
	if File_teamcloud_v1_component_template_proto != nil {
		return
	}
	file_teamcloud_v1_component_task_runner_proto_init()
	file_teamcloud_v1_component_task_template_proto_init()
	file_teamcloud_v1_component_template_permissions_proto_init()
	file_teamcloud_v1_repository_reference_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teamcloud_v1_component_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentTemplate); i {
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
	file_teamcloud_v1_component_template_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teamcloud_v1_component_template_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teamcloud_v1_component_template_proto_goTypes,
		DependencyIndexes: file_teamcloud_v1_component_template_proto_depIdxs,
		EnumInfos:         file_teamcloud_v1_component_template_proto_enumTypes,
		MessageInfos:      file_teamcloud_v1_component_template_proto_msgTypes,
	}.Build()
	File_teamcloud_v1_component_template_proto = out.File
	file_teamcloud_v1_component_template_proto_rawDesc = nil
	file_teamcloud_v1_component_template_proto_goTypes = nil
	file_teamcloud_v1_component_template_proto_depIdxs = nil
}