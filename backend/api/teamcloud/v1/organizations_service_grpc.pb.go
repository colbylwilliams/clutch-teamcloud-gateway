// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/organizations_service.proto

package teamcloudv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrganizationsAPIClient is the client API for OrganizationsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationsAPIClient interface {
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*OrganizationDataResult, error)
	DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*StatusResult, error)
	GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*OrganizationDataResult, error)
	GetOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OrganizationListDataResult, error)
}

type organizationsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationsAPIClient(cc grpc.ClientConnInterface) OrganizationsAPIClient {
	return &organizationsAPIClient{cc}
}

func (c *organizationsAPIClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*OrganizationDataResult, error) {
	out := new(OrganizationDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationsAPI/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsAPIClient) DeleteOrganization(ctx context.Context, in *DeleteOrganizationRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationsAPI/DeleteOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsAPIClient) GetOrganization(ctx context.Context, in *GetOrganizationRequest, opts ...grpc.CallOption) (*OrganizationDataResult, error) {
	out := new(OrganizationDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationsAPI/GetOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsAPIClient) GetOrganizations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OrganizationListDataResult, error) {
	out := new(OrganizationListDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.OrganizationsAPI/GetOrganizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationsAPIServer is the server API for OrganizationsAPI service.
// All implementations should embed UnimplementedOrganizationsAPIServer
// for forward compatibility
type OrganizationsAPIServer interface {
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*OrganizationDataResult, error)
	DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*StatusResult, error)
	GetOrganization(context.Context, *GetOrganizationRequest) (*OrganizationDataResult, error)
	GetOrganizations(context.Context, *emptypb.Empty) (*OrganizationListDataResult, error)
}

// UnimplementedOrganizationsAPIServer should be embedded to have forward compatible implementations.
type UnimplementedOrganizationsAPIServer struct {
}

func (UnimplementedOrganizationsAPIServer) CreateOrganization(context.Context, *CreateOrganizationRequest) (*OrganizationDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizationsAPIServer) DeleteOrganization(context.Context, *DeleteOrganizationRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganization not implemented")
}
func (UnimplementedOrganizationsAPIServer) GetOrganization(context.Context, *GetOrganizationRequest) (*OrganizationDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganization not implemented")
}
func (UnimplementedOrganizationsAPIServer) GetOrganizations(context.Context, *emptypb.Empty) (*OrganizationListDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizations not implemented")
}

// UnsafeOrganizationsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationsAPIServer will
// result in compilation errors.
type UnsafeOrganizationsAPIServer interface {
	mustEmbedUnimplementedOrganizationsAPIServer()
}

func RegisterOrganizationsAPIServer(s grpc.ServiceRegistrar, srv OrganizationsAPIServer) {
	s.RegisterService(&OrganizationsAPI_ServiceDesc, srv)
}

func _OrganizationsAPI_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsAPIServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationsAPI/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsAPIServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationsAPI_DeleteOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsAPIServer).DeleteOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationsAPI/DeleteOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsAPIServer).DeleteOrganization(ctx, req.(*DeleteOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationsAPI_GetOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsAPIServer).GetOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationsAPI/GetOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsAPIServer).GetOrganization(ctx, req.(*GetOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationsAPI_GetOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsAPIServer).GetOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.OrganizationsAPI/GetOrganizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsAPIServer).GetOrganizations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationsAPI_ServiceDesc is the grpc.ServiceDesc for OrganizationsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.OrganizationsAPI",
	HandlerType: (*OrganizationsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrganization",
			Handler:    _OrganizationsAPI_CreateOrganization_Handler,
		},
		{
			MethodName: "DeleteOrganization",
			Handler:    _OrganizationsAPI_DeleteOrganization_Handler,
		},
		{
			MethodName: "GetOrganization",
			Handler:    _OrganizationsAPI_GetOrganization_Handler,
		},
		{
			MethodName: "GetOrganizations",
			Handler:    _OrganizationsAPI_GetOrganizations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/organizations_service.proto",
}
