// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/teamcloud.proto

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

// TeamCloudAPIClient is the client API for TeamCloudAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamCloudAPIClient interface {
	GetTeamCloudInfo(ctx context.Context, in *GetTeamCloudInfoRequest, opts ...grpc.CallOption) (*GetTeamCloudInfoResponse, error)
}

type teamCloudAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamCloudAPIClient(cc grpc.ClientConnInterface) TeamCloudAPIClient {
	return &teamCloudAPIClient{cc}
}

func (c *teamCloudAPIClient) GetTeamCloudInfo(ctx context.Context, in *GetTeamCloudInfoRequest, opts ...grpc.CallOption) (*GetTeamCloudInfoResponse, error) {
	out := new(GetTeamCloudInfoResponse)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.TeamCloudAPI/GetTeamCloudInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamCloudAPIServer is the server API for TeamCloudAPI service.
// All implementations should embed UnimplementedTeamCloudAPIServer
// for forward compatibility
type TeamCloudAPIServer interface {
	GetTeamCloudInfo(context.Context, *GetTeamCloudInfoRequest) (*GetTeamCloudInfoResponse, error)
}

// UnimplementedTeamCloudAPIServer should be embedded to have forward compatible implementations.
type UnimplementedTeamCloudAPIServer struct {
}

func (UnimplementedTeamCloudAPIServer) GetTeamCloudInfo(context.Context, *GetTeamCloudInfoRequest) (*GetTeamCloudInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamCloudInfo not implemented")
}

// UnsafeTeamCloudAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamCloudAPIServer will
// result in compilation errors.
type UnsafeTeamCloudAPIServer interface {
	mustEmbedUnimplementedTeamCloudAPIServer()
}

func RegisterTeamCloudAPIServer(s grpc.ServiceRegistrar, srv TeamCloudAPIServer) {
	s.RegisterService(&TeamCloudAPI_ServiceDesc, srv)
}

func _TeamCloudAPI_GetTeamCloudInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamCloudInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamCloudAPIServer).GetTeamCloudInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.TeamCloudAPI/GetTeamCloudInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamCloudAPIServer).GetTeamCloudInfo(ctx, req.(*GetTeamCloudInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamCloudAPI_ServiceDesc is the grpc.ServiceDesc for TeamCloudAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamCloudAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.TeamCloudAPI",
	HandlerType: (*TeamCloudAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeamCloudInfo",
			Handler:    _TeamCloudAPI_GetTeamCloudInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/teamcloud.proto",
}
