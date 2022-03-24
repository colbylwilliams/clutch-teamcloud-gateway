// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/organization_audit_service.proto

package teamcloudv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationAuditAPIClient is the client API for OrganizationAuditAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationAuditAPIClient interface {
	GetAuditCommands(ctx context.Context, in *GetAuditCommandsRequest, opts ...grpc.CallOption) (*StringListDataResult, error)
	GetAuditEntries(ctx context.Context, in *GetAuditEntriesRequest, opts ...grpc.CallOption) (*CommandAuditEntityListDataResult, error)
	GetAuditEntry(ctx context.Context, in *GetAuditEntryRequest, opts ...grpc.CallOption) (*CommandAuditEntityDataResult, error)
}

type organizationAuditAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationAuditAPIClient(cc grpc.ClientConnInterface) OrganizationAuditAPIClient {
	return &organizationAuditAPIClient{cc}
}

func (c *organizationAuditAPIClient) GetAuditCommands(ctx context.Context, in *GetAuditCommandsRequest, opts ...grpc.CallOption) (*StringListDataResult, error) {
	out := new(StringListDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationAuditAPI/GetAuditCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAuditAPIClient) GetAuditEntries(ctx context.Context, in *GetAuditEntriesRequest, opts ...grpc.CallOption) (*CommandAuditEntityListDataResult, error) {
	out := new(CommandAuditEntityListDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationAuditAPI/GetAuditEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationAuditAPIClient) GetAuditEntry(ctx context.Context, in *GetAuditEntryRequest, opts ...grpc.CallOption) (*CommandAuditEntityDataResult, error) {
	out := new(CommandAuditEntityDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationAuditAPI/GetAuditEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationAuditAPIServer is the server API for OrganizationAuditAPI service.
// All implementations should embed UnimplementedOrganizationAuditAPIServer
// for forward compatibility
type OrganizationAuditAPIServer interface {
	GetAuditCommands(context.Context, *GetAuditCommandsRequest) (*StringListDataResult, error)
	GetAuditEntries(context.Context, *GetAuditEntriesRequest) (*CommandAuditEntityListDataResult, error)
	GetAuditEntry(context.Context, *GetAuditEntryRequest) (*CommandAuditEntityDataResult, error)
}

// UnimplementedOrganizationAuditAPIServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationAuditAPIServer struct {
}

func (UnimplementedOrganizationAuditAPIServer) GetAuditCommands(context.Context, *GetAuditCommandsRequest) (*StringListDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditCommands not implemented")
}
func (UnimplementedOrganizationAuditAPIServer) GetAuditEntries(context.Context, *GetAuditEntriesRequest) (*CommandAuditEntityListDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditEntries not implemented")
}
func (UnimplementedOrganizationAuditAPIServer) GetAuditEntry(context.Context, *GetAuditEntryRequest) (*CommandAuditEntityDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditEntry not implemented")
}

// UnsafeOrganizationAuditAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationAuditAPIServer will
// result in compilation errors.
type UnsafeOrganizationAuditAPIServer interface {
	mustEmbedUnimplementedOrganizationAuditAPIServer()
}

func RegisterOrganizationAuditAPIServer(s grpc.ServiceRegistrar, srv OrganizationAuditAPIServer) {
	s.RegisterService(&OrganizationAuditAPI_ServiceDesc, srv)
}

func _OrganizationAuditAPI_GetAuditCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditCommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAuditAPIServer).GetAuditCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationAuditAPI/GetAuditCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAuditAPIServer).GetAuditCommands(ctx, req.(*GetAuditCommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAuditAPI_GetAuditEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAuditAPIServer).GetAuditEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationAuditAPI/GetAuditEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAuditAPIServer).GetAuditEntries(ctx, req.(*GetAuditEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationAuditAPI_GetAuditEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuditEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationAuditAPIServer).GetAuditEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationAuditAPI/GetAuditEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationAuditAPIServer).GetAuditEntry(ctx, req.(*GetAuditEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationAuditAPI_ServiceDesc is the grpc.ServiceDesc for OrganizationAuditAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationAuditAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.OrganizationAuditAPI",
	HandlerType: (*OrganizationAuditAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuditCommands",
			Handler:    _OrganizationAuditAPI_GetAuditCommands_Handler,
		},
		{
			MethodName: "GetAuditEntries",
			Handler:    _OrganizationAuditAPI_GetAuditEntries_Handler,
		},
		{
			MethodName: "GetAuditEntry",
			Handler:    _OrganizationAuditAPI_GetAuditEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/organization_audit_service.proto",
}