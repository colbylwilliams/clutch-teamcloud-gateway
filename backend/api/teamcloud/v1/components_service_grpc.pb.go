// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/components_service.proto

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

// ComponentsAPIClient is the client API for ComponentsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentsAPIClient interface {
	CreateComponent(ctx context.Context, in *CreateComponentRequest, opts ...grpc.CallOption) (*ComponentDataResult, error)
	DeleteComponent(ctx context.Context, in *DeleteComponentRequest, opts ...grpc.CallOption) (*StatusResult, error)
	GetComponent(ctx context.Context, in *GetComponentRequest, opts ...grpc.CallOption) (*ComponentDataResult, error)
	GetComponents(ctx context.Context, in *GetComponentsRequest, opts ...grpc.CallOption) (*ComponentListDataResult, error)
}

type componentsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentsAPIClient(cc grpc.ClientConnInterface) ComponentsAPIClient {
	return &componentsAPIClient{cc}
}

func (c *componentsAPIClient) CreateComponent(ctx context.Context, in *CreateComponentRequest, opts ...grpc.CallOption) (*ComponentDataResult, error) {
	out := new(ComponentDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ComponentsAPI/CreateComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsAPIClient) DeleteComponent(ctx context.Context, in *DeleteComponentRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ComponentsAPI/DeleteComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsAPIClient) GetComponent(ctx context.Context, in *GetComponentRequest, opts ...grpc.CallOption) (*ComponentDataResult, error) {
	out := new(ComponentDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ComponentsAPI/GetComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsAPIClient) GetComponents(ctx context.Context, in *GetComponentsRequest, opts ...grpc.CallOption) (*ComponentListDataResult, error) {
	out := new(ComponentListDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ComponentsAPI/GetComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentsAPIServer is the server API for ComponentsAPI service.
// All implementations should embed UnimplementedComponentsAPIServer
// for forward compatibility
type ComponentsAPIServer interface {
	CreateComponent(context.Context, *CreateComponentRequest) (*ComponentDataResult, error)
	DeleteComponent(context.Context, *DeleteComponentRequest) (*StatusResult, error)
	GetComponent(context.Context, *GetComponentRequest) (*ComponentDataResult, error)
	GetComponents(context.Context, *GetComponentsRequest) (*ComponentListDataResult, error)
}

// UnimplementedComponentsAPIServer should be embedded to have forward compatible implementations.
type UnimplementedComponentsAPIServer struct {
}

func (UnimplementedComponentsAPIServer) CreateComponent(context.Context, *CreateComponentRequest) (*ComponentDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComponent not implemented")
}
func (UnimplementedComponentsAPIServer) DeleteComponent(context.Context, *DeleteComponentRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComponent not implemented")
}
func (UnimplementedComponentsAPIServer) GetComponent(context.Context, *GetComponentRequest) (*ComponentDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponent not implemented")
}
func (UnimplementedComponentsAPIServer) GetComponents(context.Context, *GetComponentsRequest) (*ComponentListDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponents not implemented")
}

// UnsafeComponentsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentsAPIServer will
// result in compilation errors.
type UnsafeComponentsAPIServer interface {
	mustEmbedUnimplementedComponentsAPIServer()
}

func RegisterComponentsAPIServer(s grpc.ServiceRegistrar, srv ComponentsAPIServer) {
	s.RegisterService(&ComponentsAPI_ServiceDesc, srv)
}

func _ComponentsAPI_CreateComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsAPIServer).CreateComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ComponentsAPI/CreateComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsAPIServer).CreateComponent(ctx, req.(*CreateComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentsAPI_DeleteComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsAPIServer).DeleteComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ComponentsAPI/DeleteComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsAPIServer).DeleteComponent(ctx, req.(*DeleteComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentsAPI_GetComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsAPIServer).GetComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ComponentsAPI/GetComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsAPIServer).GetComponent(ctx, req.(*GetComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComponentsAPI_GetComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComponentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsAPIServer).GetComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ComponentsAPI/GetComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsAPIServer).GetComponents(ctx, req.(*GetComponentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComponentsAPI_ServiceDesc is the grpc.ServiceDesc for ComponentsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComponentsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.ComponentsAPI",
	HandlerType: (*ComponentsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComponent",
			Handler:    _ComponentsAPI_CreateComponent_Handler,
		},
		{
			MethodName: "DeleteComponent",
			Handler:    _ComponentsAPI_DeleteComponent_Handler,
		},
		{
			MethodName: "GetComponent",
			Handler:    _ComponentsAPI_GetComponent_Handler,
		},
		{
			MethodName: "GetComponents",
			Handler:    _ComponentsAPI_GetComponents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/components_service.proto",
}