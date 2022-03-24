// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/status_service.proto

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

// StatusAPIClient is the client API for StatusAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatusAPIClient interface {
	GetProjectStatus(ctx context.Context, in *GetProjectStatusRequest, opts ...grpc.CallOption) (*StatusResult, error)
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*StatusResult, error)
}

type statusAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusAPIClient(cc grpc.ClientConnInterface) StatusAPIClient {
	return &statusAPIClient{cc}
}

func (c *statusAPIClient) GetProjectStatus(ctx context.Context, in *GetProjectStatusRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.StatusAPI/GetProjectStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusAPIClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.StatusAPI/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusAPIServer is the server API for StatusAPI service.
// All implementations should embed UnimplementedStatusAPIServer
// for forward compatibility
type StatusAPIServer interface {
	GetProjectStatus(context.Context, *GetProjectStatusRequest) (*StatusResult, error)
	GetStatus(context.Context, *GetStatusRequest) (*StatusResult, error)
}

// UnimplementedStatusAPIServer should be embedded to have forward compatible implementations.
type UnimplementedStatusAPIServer struct {
}

func (UnimplementedStatusAPIServer) GetProjectStatus(context.Context, *GetProjectStatusRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectStatus not implemented")
}
func (UnimplementedStatusAPIServer) GetStatus(context.Context, *GetStatusRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}

// UnsafeStatusAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatusAPIServer will
// result in compilation errors.
type UnsafeStatusAPIServer interface {
	mustEmbedUnimplementedStatusAPIServer()
}

func RegisterStatusAPIServer(s grpc.ServiceRegistrar, srv StatusAPIServer) {
	s.RegisterService(&StatusAPI_ServiceDesc, srv)
}

func _StatusAPI_GetProjectStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusAPIServer).GetProjectStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.StatusAPI/GetProjectStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusAPIServer).GetProjectStatus(ctx, req.(*GetProjectStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusAPI_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusAPIServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.StatusAPI/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusAPIServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatusAPI_ServiceDesc is the grpc.ServiceDesc for StatusAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatusAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.StatusAPI",
	HandlerType: (*StatusAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProjectStatus",
			Handler:    _StatusAPI_GetProjectStatus_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _StatusAPI_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/status_service.proto",
}