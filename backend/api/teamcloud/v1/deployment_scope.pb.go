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
// source: teamcloud/v1/deployment_scope.proto

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

type DeploymentScope_TypeEnum int32

const (
	DeploymentScope_AZURERESOURCEMANAGER DeploymentScope_TypeEnum = 0
	DeploymentScope_AZUREDEVOPS          DeploymentScope_TypeEnum = 1
	DeploymentScope_GITHUB               DeploymentScope_TypeEnum = 2
	DeploymentScope_KUBERNETES           DeploymentScope_TypeEnum = 3
)

// Enum value maps for DeploymentScope_TypeEnum.
var (
	DeploymentScope_TypeEnum_name = map[int32]string{
		0: "AZURERESOURCEMANAGER",
		1: "AZUREDEVOPS",
		2: "GITHUB",
		3: "KUBERNETES",
	}
	DeploymentScope_TypeEnum_value = map[string]int32{
		"AZURERESOURCEMANAGER": 0,
		"AZUREDEVOPS":          1,
		"GITHUB":               2,
		"KUBERNETES":           3,
	}
)

func (x DeploymentScope_TypeEnum) Enum() *DeploymentScope_TypeEnum {
	p := new(DeploymentScope_TypeEnum)
	*p = x
	return p
}

func (x DeploymentScope_TypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentScope_TypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_teamcloud_v1_deployment_scope_proto_enumTypes[0].Descriptor()
}

func (DeploymentScope_TypeEnum) Type() protoreflect.EnumType {
	return &file_teamcloud_v1_deployment_scope_proto_enumTypes[0]
}

func (x DeploymentScope_TypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentScope_TypeEnum.Descriptor instead.
func (DeploymentScope_TypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_teamcloud_v1_deployment_scope_proto_rawDescGZIP(), []int{0, 0}
}

type DeploymentScope_ComponentTypesEnum int32

const (
	DeploymentScope_ENVIRONMENT DeploymentScope_ComponentTypesEnum = 0
	DeploymentScope_REPOSITORY  DeploymentScope_ComponentTypesEnum = 1
	DeploymentScope_NAMESPACE   DeploymentScope_ComponentTypesEnum = 2
)

// Enum value maps for DeploymentScope_ComponentTypesEnum.
var (
	DeploymentScope_ComponentTypesEnum_name = map[int32]string{
		0: "ENVIRONMENT",
		1: "REPOSITORY",
		2: "NAMESPACE",
	}
	DeploymentScope_ComponentTypesEnum_value = map[string]int32{
		"ENVIRONMENT": 0,
		"REPOSITORY":  1,
		"NAMESPACE":   2,
	}
)

func (x DeploymentScope_ComponentTypesEnum) Enum() *DeploymentScope_ComponentTypesEnum {
	p := new(DeploymentScope_ComponentTypesEnum)
	*p = x
	return p
}

func (x DeploymentScope_ComponentTypesEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentScope_ComponentTypesEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_teamcloud_v1_deployment_scope_proto_enumTypes[1].Descriptor()
}

func (DeploymentScope_ComponentTypesEnum) Type() protoreflect.EnumType {
	return &file_teamcloud_v1_deployment_scope_proto_enumTypes[1]
}

func (x DeploymentScope_ComponentTypesEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentScope_ComponentTypesEnum.Descriptor instead.
func (DeploymentScope_ComponentTypesEnum) EnumDescriptor() ([]byte, []int) {
	return file_teamcloud_v1_deployment_scope_proto_rawDescGZIP(), []int{0, 1}
}

type DeploymentScope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization      string                             `protobuf:"bytes,105180469,opt,name=organization,proto3" json:"organization,omitempty"`
	OrganizationName  string                             `protobuf:"bytes,218494373,opt,name=organizationName,proto3" json:"organizationName,omitempty"`
	DisplayName       string                             `protobuf:"bytes,103536240,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Slug              string                             `protobuf:"bytes,3533483,opt,name=slug,proto3" json:"slug,omitempty"`
	IsDefault         bool                               `protobuf:"varint,428154296,opt,name=isDefault,proto3" json:"isDefault,omitempty"`
	Type              DeploymentScope_TypeEnum           `protobuf:"varint,3575610,opt,name=type,proto3,enum=colbylwilliams.teamcloud.v1.DeploymentScope_TypeEnum" json:"type,omitempty"`
	InputDataSchema   *string                            `protobuf:"bytes,9787051,opt,name=inputDataSchema,proto3,oneof" json:"inputDataSchema,omitempty"`
	InputData         *string                            `protobuf:"bytes,95864471,opt,name=inputData,proto3,oneof" json:"inputData,omitempty"`
	ManagementGroupId *string                            `protobuf:"bytes,70728247,opt,name=managementGroupId,proto3,oneof" json:"managementGroupId,omitempty"`
	SubscriptionIds   []string                           `protobuf:"bytes,328379303,rep,name=subscriptionIds,proto3" json:"subscriptionIds,omitempty"`
	Authorizable      bool                               `protobuf:"varint,388247383,opt,name=authorizable,proto3" json:"authorizable,omitempty"`
	Authorized        bool                               `protobuf:"varint,426969703,opt,name=authorized,proto3" json:"authorized,omitempty"`
	AuthorizeUrl      *string                            `protobuf:"bytes,388354247,opt,name=authorizeUrl,proto3,oneof" json:"authorizeUrl,omitempty"`
	ComponentTypes    DeploymentScope_ComponentTypesEnum `protobuf:"varint,286581118,opt,name=componentTypes,proto3,enum=colbylwilliams.teamcloud.v1.DeploymentScope_ComponentTypesEnum" json:"componentTypes,omitempty"`
	Id                string                             `protobuf:"bytes,3355,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeploymentScope) Reset() {
	*x = DeploymentScope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamcloud_v1_deployment_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentScope) ProtoMessage() {}

func (x *DeploymentScope) ProtoReflect() protoreflect.Message {
	mi := &file_teamcloud_v1_deployment_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentScope.ProtoReflect.Descriptor instead.
func (*DeploymentScope) Descriptor() ([]byte, []int) {
	return file_teamcloud_v1_deployment_scope_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentScope) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *DeploymentScope) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *DeploymentScope) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DeploymentScope) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *DeploymentScope) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

func (x *DeploymentScope) GetType() DeploymentScope_TypeEnum {
	if x != nil {
		return x.Type
	}
	return DeploymentScope_AZURERESOURCEMANAGER
}

func (x *DeploymentScope) GetInputDataSchema() string {
	if x != nil && x.InputDataSchema != nil {
		return *x.InputDataSchema
	}
	return ""
}

func (x *DeploymentScope) GetInputData() string {
	if x != nil && x.InputData != nil {
		return *x.InputData
	}
	return ""
}

func (x *DeploymentScope) GetManagementGroupId() string {
	if x != nil && x.ManagementGroupId != nil {
		return *x.ManagementGroupId
	}
	return ""
}

func (x *DeploymentScope) GetSubscriptionIds() []string {
	if x != nil {
		return x.SubscriptionIds
	}
	return nil
}

func (x *DeploymentScope) GetAuthorizable() bool {
	if x != nil {
		return x.Authorizable
	}
	return false
}

func (x *DeploymentScope) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

func (x *DeploymentScope) GetAuthorizeUrl() string {
	if x != nil && x.AuthorizeUrl != nil {
		return *x.AuthorizeUrl
	}
	return ""
}

func (x *DeploymentScope) GetComponentTypes() DeploymentScope_ComponentTypesEnum {
	if x != nil {
		return x.ComponentTypes
	}
	return DeploymentScope_ENVIRONMENT
}

func (x *DeploymentScope) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_teamcloud_v1_deployment_scope_proto protoreflect.FileDescriptor

var file_teamcloud_v1_deployment_scope_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c,
	0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x31, 0x22, 0xa8, 0x07, 0x0a, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xb5, 0xda, 0x93, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0xa5, 0xeb, 0x97, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0xf0, 0xac, 0xaf, 0x31,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0xab, 0xd5, 0xd7, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0xb8, 0xbb, 0x94, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0xba, 0x9e, 0xda, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f,
	0x6c, 0x62, 0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x0f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0xab, 0xad, 0xd5, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x09, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x97, 0x8d, 0xdb, 0x2d, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x34, 0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0xb7, 0xf4, 0xdc, 0x21, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0xa7, 0xd7, 0xca, 0x9c, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0xd7, 0xde, 0x90, 0xb9, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0xe7, 0x94, 0xcc, 0xcb, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x12, 0x2b, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c,
	0x18, 0xc7, 0xa1, 0x97, 0xb9, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0c, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x6b, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18,
	0xfe, 0xc2, 0xd3, 0x88, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x6c, 0x62,
	0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x9b, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x08, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x5a, 0x55, 0x52, 0x45,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x5a, 0x55, 0x52, 0x45, 0x44, 0x45, 0x56, 0x4f, 0x50, 0x53,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x10, 0x03, 0x22, 0x44,
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x50, 0x4f, 0x53, 0x49, 0x54,
	0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41,
	0x43, 0x45, 0x10, 0x02, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x59, 0x5a,
	0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x62,
	0x79, 0x6c, 0x77, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63,
	0x68, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x61,
	0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teamcloud_v1_deployment_scope_proto_rawDescOnce sync.Once
	file_teamcloud_v1_deployment_scope_proto_rawDescData = file_teamcloud_v1_deployment_scope_proto_rawDesc
)

func file_teamcloud_v1_deployment_scope_proto_rawDescGZIP() []byte {
	file_teamcloud_v1_deployment_scope_proto_rawDescOnce.Do(func() {
		file_teamcloud_v1_deployment_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_teamcloud_v1_deployment_scope_proto_rawDescData)
	})
	return file_teamcloud_v1_deployment_scope_proto_rawDescData
}

var file_teamcloud_v1_deployment_scope_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_teamcloud_v1_deployment_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teamcloud_v1_deployment_scope_proto_goTypes = []interface{}{
	(DeploymentScope_TypeEnum)(0),           // 0: colbylwilliams.teamcloud.v1.DeploymentScope.TypeEnum
	(DeploymentScope_ComponentTypesEnum)(0), // 1: colbylwilliams.teamcloud.v1.DeploymentScope.ComponentTypesEnum
	(*DeploymentScope)(nil),                 // 2: colbylwilliams.teamcloud.v1.DeploymentScope
}
var file_teamcloud_v1_deployment_scope_proto_depIdxs = []int32{
	0, // 0: colbylwilliams.teamcloud.v1.DeploymentScope.type:type_name -> colbylwilliams.teamcloud.v1.DeploymentScope.TypeEnum
	1, // 1: colbylwilliams.teamcloud.v1.DeploymentScope.componentTypes:type_name -> colbylwilliams.teamcloud.v1.DeploymentScope.ComponentTypesEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_teamcloud_v1_deployment_scope_proto_init() }
func file_teamcloud_v1_deployment_scope_proto_init() {
	if File_teamcloud_v1_deployment_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teamcloud_v1_deployment_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentScope); i {
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
	file_teamcloud_v1_deployment_scope_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teamcloud_v1_deployment_scope_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teamcloud_v1_deployment_scope_proto_goTypes,
		DependencyIndexes: file_teamcloud_v1_deployment_scope_proto_depIdxs,
		EnumInfos:         file_teamcloud_v1_deployment_scope_proto_enumTypes,
		MessageInfos:      file_teamcloud_v1_deployment_scope_proto_msgTypes,
	}.Build()
	File_teamcloud_v1_deployment_scope_proto = out.File
	file_teamcloud_v1_deployment_scope_proto_rawDesc = nil
	file_teamcloud_v1_deployment_scope_proto_goTypes = nil
	file_teamcloud_v1_deployment_scope_proto_depIdxs = nil
}
