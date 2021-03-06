// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/deployment_scopes_authorization_service.proto

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

// DeploymentScopesAuthorizationAPIClient is the client API for DeploymentScopesAuthorizationAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentScopesAuthorizationAPIClient interface {
	InitializeAuthorization(ctx context.Context, in *InitializeAuthorizationRequest, opts ...grpc.CallOption) (*DeploymentScopeDataResult, error)
}

type deploymentScopesAuthorizationAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentScopesAuthorizationAPIClient(cc grpc.ClientConnInterface) DeploymentScopesAuthorizationAPIClient {
	return &deploymentScopesAuthorizationAPIClient{cc}
}

func (c *deploymentScopesAuthorizationAPIClient) InitializeAuthorization(ctx context.Context, in *InitializeAuthorizationRequest, opts ...grpc.CallOption) (*DeploymentScopeDataResult, error) {
	out := new(DeploymentScopeDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.DeploymentScopesAuthorizationAPI/InitializeAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentScopesAuthorizationAPIServer is the server API for DeploymentScopesAuthorizationAPI service.
// All implementations should embed UnimplementedDeploymentScopesAuthorizationAPIServer
// for forward compatibility
type DeploymentScopesAuthorizationAPIServer interface {
	InitializeAuthorization(context.Context, *InitializeAuthorizationRequest) (*DeploymentScopeDataResult, error)
}

// UnimplementedDeploymentScopesAuthorizationAPIServer should be embedded to have forward compatible implementations.
type UnimplementedDeploymentScopesAuthorizationAPIServer struct {
}

func (UnimplementedDeploymentScopesAuthorizationAPIServer) InitializeAuthorization(context.Context, *InitializeAuthorizationRequest) (*DeploymentScopeDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeAuthorization not implemented")
}

// UnsafeDeploymentScopesAuthorizationAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentScopesAuthorizationAPIServer will
// result in compilation errors.
type UnsafeDeploymentScopesAuthorizationAPIServer interface {
	mustEmbedUnimplementedDeploymentScopesAuthorizationAPIServer()
}

func RegisterDeploymentScopesAuthorizationAPIServer(s grpc.ServiceRegistrar, srv DeploymentScopesAuthorizationAPIServer) {
	s.RegisterService(&DeploymentScopesAuthorizationAPI_ServiceDesc, srv)
}

func _DeploymentScopesAuthorizationAPI_InitializeAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentScopesAuthorizationAPIServer).InitializeAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.DeploymentScopesAuthorizationAPI/InitializeAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentScopesAuthorizationAPIServer).InitializeAuthorization(ctx, req.(*InitializeAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentScopesAuthorizationAPI_ServiceDesc is the grpc.ServiceDesc for DeploymentScopesAuthorizationAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentScopesAuthorizationAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.DeploymentScopesAuthorizationAPI",
	HandlerType: (*DeploymentScopesAuthorizationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeAuthorization",
			Handler:    _DeploymentScopesAuthorizationAPI_InitializeAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/deployment_scopes_authorization_service.proto",
}
