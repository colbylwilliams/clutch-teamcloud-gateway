// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: teamcloud/v1/project_identities_service.proto

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

// ProjectIdentitiesAPIClient is the client API for ProjectIdentitiesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectIdentitiesAPIClient interface {
	CreateProjectIdentity(ctx context.Context, in *CreateProjectIdentityRequest, opts ...grpc.CallOption) (*ProjectIdentityDataResult, error)
	DeleteProjectIdentity(ctx context.Context, in *DeleteProjectIdentityRequest, opts ...grpc.CallOption) (*ProjectIdentityDataResult, error)
	GetProjectIdentities(ctx context.Context, in *GetProjectIdentitiesRequest, opts ...grpc.CallOption) (*ProjectIdentityListDataResult, error)
	GetProjectIdentity(ctx context.Context, in *GetProjectIdentityRequest, opts ...grpc.CallOption) (*ProjectIdentityDataResult, error)
	UpdateProjectIdentity(ctx context.Context, in *UpdateProjectIdentityRequest, opts ...grpc.CallOption) (*StatusResult, error)
}

type projectIdentitiesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectIdentitiesAPIClient(cc grpc.ClientConnInterface) ProjectIdentitiesAPIClient {
	return &projectIdentitiesAPIClient{cc}
}

func (c *projectIdentitiesAPIClient) CreateProjectIdentity(ctx context.Context, in *CreateProjectIdentityRequest, opts ...grpc.CallOption) (*ProjectIdentityDataResult, error) {
	out := new(ProjectIdentityDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/CreateProjectIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectIdentitiesAPIClient) DeleteProjectIdentity(ctx context.Context, in *DeleteProjectIdentityRequest, opts ...grpc.CallOption) (*ProjectIdentityDataResult, error) {
	out := new(ProjectIdentityDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/DeleteProjectIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectIdentitiesAPIClient) GetProjectIdentities(ctx context.Context, in *GetProjectIdentitiesRequest, opts ...grpc.CallOption) (*ProjectIdentityListDataResult, error) {
	out := new(ProjectIdentityListDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/GetProjectIdentities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectIdentitiesAPIClient) GetProjectIdentity(ctx context.Context, in *GetProjectIdentityRequest, opts ...grpc.CallOption) (*ProjectIdentityDataResult, error) {
	out := new(ProjectIdentityDataResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/GetProjectIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectIdentitiesAPIClient) UpdateProjectIdentity(ctx context.Context, in *UpdateProjectIdentityRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/UpdateProjectIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectIdentitiesAPIServer is the server API for ProjectIdentitiesAPI service.
// All implementations should embed UnimplementedProjectIdentitiesAPIServer
// for forward compatibility
type ProjectIdentitiesAPIServer interface {
	CreateProjectIdentity(context.Context, *CreateProjectIdentityRequest) (*ProjectIdentityDataResult, error)
	DeleteProjectIdentity(context.Context, *DeleteProjectIdentityRequest) (*ProjectIdentityDataResult, error)
	GetProjectIdentities(context.Context, *GetProjectIdentitiesRequest) (*ProjectIdentityListDataResult, error)
	GetProjectIdentity(context.Context, *GetProjectIdentityRequest) (*ProjectIdentityDataResult, error)
	UpdateProjectIdentity(context.Context, *UpdateProjectIdentityRequest) (*StatusResult, error)
}

// UnimplementedProjectIdentitiesAPIServer should be embedded to have forward compatible implementations.
type UnimplementedProjectIdentitiesAPIServer struct {
}

func (UnimplementedProjectIdentitiesAPIServer) CreateProjectIdentity(context.Context, *CreateProjectIdentityRequest) (*ProjectIdentityDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProjectIdentity not implemented")
}
func (UnimplementedProjectIdentitiesAPIServer) DeleteProjectIdentity(context.Context, *DeleteProjectIdentityRequest) (*ProjectIdentityDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectIdentity not implemented")
}
func (UnimplementedProjectIdentitiesAPIServer) GetProjectIdentities(context.Context, *GetProjectIdentitiesRequest) (*ProjectIdentityListDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectIdentities not implemented")
}
func (UnimplementedProjectIdentitiesAPIServer) GetProjectIdentity(context.Context, *GetProjectIdentityRequest) (*ProjectIdentityDataResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectIdentity not implemented")
}
func (UnimplementedProjectIdentitiesAPIServer) UpdateProjectIdentity(context.Context, *UpdateProjectIdentityRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectIdentity not implemented")
}

// UnsafeProjectIdentitiesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectIdentitiesAPIServer will
// result in compilation errors.
type UnsafeProjectIdentitiesAPIServer interface {
	mustEmbedUnimplementedProjectIdentitiesAPIServer()
}

func RegisterProjectIdentitiesAPIServer(s grpc.ServiceRegistrar, srv ProjectIdentitiesAPIServer) {
	s.RegisterService(&ProjectIdentitiesAPI_ServiceDesc, srv)
}

func _ProjectIdentitiesAPI_CreateProjectIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectIdentitiesAPIServer).CreateProjectIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/CreateProjectIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectIdentitiesAPIServer).CreateProjectIdentity(ctx, req.(*CreateProjectIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectIdentitiesAPI_DeleteProjectIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectIdentitiesAPIServer).DeleteProjectIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/DeleteProjectIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectIdentitiesAPIServer).DeleteProjectIdentity(ctx, req.(*DeleteProjectIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectIdentitiesAPI_GetProjectIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectIdentitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectIdentitiesAPIServer).GetProjectIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/GetProjectIdentities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectIdentitiesAPIServer).GetProjectIdentities(ctx, req.(*GetProjectIdentitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectIdentitiesAPI_GetProjectIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectIdentitiesAPIServer).GetProjectIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/GetProjectIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectIdentitiesAPIServer).GetProjectIdentity(ctx, req.(*GetProjectIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectIdentitiesAPI_UpdateProjectIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectIdentitiesAPIServer).UpdateProjectIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI/UpdateProjectIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectIdentitiesAPIServer).UpdateProjectIdentity(ctx, req.(*UpdateProjectIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectIdentitiesAPI_ServiceDesc is the grpc.ServiceDesc for ProjectIdentitiesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectIdentitiesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "colbylwilliams.teamcloud.v1.ProjectIdentitiesAPI",
	HandlerType: (*ProjectIdentitiesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProjectIdentity",
			Handler:    _ProjectIdentitiesAPI_CreateProjectIdentity_Handler,
		},
		{
			MethodName: "DeleteProjectIdentity",
			Handler:    _ProjectIdentitiesAPI_DeleteProjectIdentity_Handler,
		},
		{
			MethodName: "GetProjectIdentities",
			Handler:    _ProjectIdentitiesAPI_GetProjectIdentities_Handler,
		},
		{
			MethodName: "GetProjectIdentity",
			Handler:    _ProjectIdentitiesAPI_GetProjectIdentity_Handler,
		},
		{
			MethodName: "UpdateProjectIdentity",
			Handler:    _ProjectIdentitiesAPI_UpdateProjectIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teamcloud/v1/project_identities_service.proto",
}