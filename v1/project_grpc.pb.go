// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: gitpod/v1/project.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProjectService_CreateProject_FullMethodName                = "/gitpod.v1.ProjectService/CreateProject"
	ProjectService_CreateProjectFromEnvironment_FullMethodName = "/gitpod.v1.ProjectService/CreateProjectFromEnvironment"
	ProjectService_GetProject_FullMethodName                   = "/gitpod.v1.ProjectService/GetProject"
	ProjectService_UpdateProject_FullMethodName                = "/gitpod.v1.ProjectService/UpdateProject"
	ProjectService_ListProjects_FullMethodName                 = "/gitpod.v1.ProjectService/ListProjects"
	ProjectService_DeleteProject_FullMethodName                = "/gitpod.v1.ProjectService/DeleteProject"
	ProjectService_ListProjectPolicies_FullMethodName          = "/gitpod.v1.ProjectService/ListProjectPolicies"
	ProjectService_CreateProjectPolicy_FullMethodName          = "/gitpod.v1.ProjectService/CreateProjectPolicy"
	ProjectService_DeleteProjectPolicy_FullMethodName          = "/gitpod.v1.ProjectService/DeleteProjectPolicy"
	ProjectService_UpdateProjectPolicy_FullMethodName          = "/gitpod.v1.ProjectService/UpdateProjectPolicy"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	// CreateProject creates a new Project.
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	// CreateProject creates a new Project using an environment as template.
	CreateProjectFromEnvironment(ctx context.Context, in *CreateProjectFromEnvironmentRequest, opts ...grpc.CallOption) (*CreateProjectFromEnvironmentResponse, error)
	// GetProject retrieves a single Project.
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	// UpdateProject updates the properties of a Project.
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
	// ListProjects lists all projects the caller has access to.
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// DeleteProject deletes the specified project.
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
	// ListProjectPolicies lists policies for a project.
	ListProjectPolicies(ctx context.Context, in *ListProjectPoliciesRequest, opts ...grpc.CallOption) (*ListProjectPoliciesResponse, error)
	// CreateProjectPolicy creates a Project Policy.
	CreateProjectPolicy(ctx context.Context, in *CreateProjectPolicyRequest, opts ...grpc.CallOption) (*CreateProjectPolicyResponse, error)
	// DeleteProjectPolicy deletes a Project Policy.
	DeleteProjectPolicy(ctx context.Context, in *DeleteProjectPolicyRequest, opts ...grpc.CallOption) (*DeleteProjectPolicyResponse, error)
	// UpdateProjectPolicy updates a Project Policy.
	UpdateProjectPolicy(ctx context.Context, in *UpdateProjectPolicyRequest, opts ...grpc.CallOption) (*UpdateProjectPolicyResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) CreateProjectFromEnvironment(ctx context.Context, in *CreateProjectFromEnvironmentRequest, opts ...grpc.CallOption) (*CreateProjectFromEnvironmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectFromEnvironmentResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProjectFromEnvironment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_ListProjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) ListProjectPolicies(ctx context.Context, in *ListProjectPoliciesRequest, opts ...grpc.CallOption) (*ListProjectPoliciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectPoliciesResponse)
	err := c.cc.Invoke(ctx, ProjectService_ListProjectPolicies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) CreateProjectPolicy(ctx context.Context, in *CreateProjectPolicyRequest, opts ...grpc.CallOption) (*CreateProjectPolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectPolicyResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProjectPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProjectPolicy(ctx context.Context, in *DeleteProjectPolicyRequest, opts ...grpc.CallOption) (*DeleteProjectPolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProjectPolicyResponse)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProjectPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProjectPolicy(ctx context.Context, in *UpdateProjectPolicyRequest, opts ...grpc.CallOption) (*UpdateProjectPolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProjectPolicyResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProjectPolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility.
type ProjectServiceServer interface {
	// CreateProject creates a new Project.
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	// CreateProject creates a new Project using an environment as template.
	CreateProjectFromEnvironment(context.Context, *CreateProjectFromEnvironmentRequest) (*CreateProjectFromEnvironmentResponse, error)
	// GetProject retrieves a single Project.
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	// UpdateProject updates the properties of a Project.
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error)
	// ListProjects lists all projects the caller has access to.
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// DeleteProject deletes the specified project.
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	// ListProjectPolicies lists policies for a project.
	ListProjectPolicies(context.Context, *ListProjectPoliciesRequest) (*ListProjectPoliciesResponse, error)
	// CreateProjectPolicy creates a Project Policy.
	CreateProjectPolicy(context.Context, *CreateProjectPolicyRequest) (*CreateProjectPolicyResponse, error)
	// DeleteProjectPolicy deletes a Project Policy.
	DeleteProjectPolicy(context.Context, *DeleteProjectPolicyRequest) (*DeleteProjectPolicyResponse, error)
	// UpdateProjectPolicy updates a Project Policy.
	UpdateProjectPolicy(context.Context, *UpdateProjectPolicyRequest) (*UpdateProjectPolicyResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectServiceServer struct{}

func (UnimplementedProjectServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) CreateProjectFromEnvironment(context.Context, *CreateProjectFromEnvironmentRequest) (*CreateProjectFromEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProjectFromEnvironment not implemented")
}
func (UnimplementedProjectServiceServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectServiceServer) ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectServiceServer) ListProjectPolicies(context.Context, *ListProjectPoliciesRequest) (*ListProjectPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjectPolicies not implemented")
}
func (UnimplementedProjectServiceServer) CreateProjectPolicy(context.Context, *CreateProjectPolicyRequest) (*CreateProjectPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProjectPolicy not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProjectPolicy(context.Context, *DeleteProjectPolicyRequest) (*DeleteProjectPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProjectPolicy not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProjectPolicy(context.Context, *UpdateProjectPolicyRequest) (*UpdateProjectPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectPolicy not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}
func (UnimplementedProjectServiceServer) testEmbeddedByValue()                        {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	// If the following call pancis, it indicates UnimplementedProjectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_CreateProjectFromEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectFromEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProjectFromEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProjectFromEnvironment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProjectFromEnvironment(ctx, req.(*CreateProjectFromEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_ListProjectPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).ListProjectPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_ListProjectPolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).ListProjectPolicies(ctx, req.(*ListProjectPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_CreateProjectPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProjectPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProjectPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProjectPolicy(ctx, req.(*CreateProjectPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProjectPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProjectPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProjectPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProjectPolicy(ctx, req.(*DeleteProjectPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProjectPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProjectPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProjectPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProjectPolicy(ctx, req.(*UpdateProjectPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "CreateProjectFromEnvironment",
			Handler:    _ProjectService_CreateProjectFromEnvironment_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _ProjectService_UpdateProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _ProjectService_ListProjects_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
		{
			MethodName: "ListProjectPolicies",
			Handler:    _ProjectService_ListProjectPolicies_Handler,
		},
		{
			MethodName: "CreateProjectPolicy",
			Handler:    _ProjectService_CreateProjectPolicy_Handler,
		},
		{
			MethodName: "DeleteProjectPolicy",
			Handler:    _ProjectService_DeleteProjectPolicy_Handler,
		},
		{
			MethodName: "UpdateProjectPolicy",
			Handler:    _ProjectService_UpdateProjectPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/v1/project.proto",
}
