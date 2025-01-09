// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: gitpod/v1/environment_automation.proto

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
	EnvironmentAutomationService_CreateService_FullMethodName             = "/gitpod.v1.EnvironmentAutomationService/CreateService"
	EnvironmentAutomationService_GetService_FullMethodName                = "/gitpod.v1.EnvironmentAutomationService/GetService"
	EnvironmentAutomationService_ListServices_FullMethodName              = "/gitpod.v1.EnvironmentAutomationService/ListServices"
	EnvironmentAutomationService_UpdateService_FullMethodName             = "/gitpod.v1.EnvironmentAutomationService/UpdateService"
	EnvironmentAutomationService_DeleteService_FullMethodName             = "/gitpod.v1.EnvironmentAutomationService/DeleteService"
	EnvironmentAutomationService_StartService_FullMethodName              = "/gitpod.v1.EnvironmentAutomationService/StartService"
	EnvironmentAutomationService_StopService_FullMethodName               = "/gitpod.v1.EnvironmentAutomationService/StopService"
	EnvironmentAutomationService_UpsertAutomationsFile_FullMethodName     = "/gitpod.v1.EnvironmentAutomationService/UpsertAutomationsFile"
	EnvironmentAutomationService_CreateTask_FullMethodName                = "/gitpod.v1.EnvironmentAutomationService/CreateTask"
	EnvironmentAutomationService_GetTask_FullMethodName                   = "/gitpod.v1.EnvironmentAutomationService/GetTask"
	EnvironmentAutomationService_ListTasks_FullMethodName                 = "/gitpod.v1.EnvironmentAutomationService/ListTasks"
	EnvironmentAutomationService_UpdateTask_FullMethodName                = "/gitpod.v1.EnvironmentAutomationService/UpdateTask"
	EnvironmentAutomationService_DeleteTask_FullMethodName                = "/gitpod.v1.EnvironmentAutomationService/DeleteTask"
	EnvironmentAutomationService_StartTask_FullMethodName                 = "/gitpod.v1.EnvironmentAutomationService/StartTask"
	EnvironmentAutomationService_ListTaskExecutions_FullMethodName        = "/gitpod.v1.EnvironmentAutomationService/ListTaskExecutions"
	EnvironmentAutomationService_GetTaskExecution_FullMethodName          = "/gitpod.v1.EnvironmentAutomationService/GetTaskExecution"
	EnvironmentAutomationService_StopTaskExecution_FullMethodName         = "/gitpod.v1.EnvironmentAutomationService/StopTaskExecution"
	EnvironmentAutomationService_UpdateTaskExecutionStatus_FullMethodName = "/gitpod.v1.EnvironmentAutomationService/UpdateTaskExecutionStatus"
)

// EnvironmentAutomationServiceClient is the client API for EnvironmentAutomationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentAutomationServiceClient interface {
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error)
	// DeleteService deletes a service. This call does not block until the service is deleted.
	// If the service is not stopped it will be stopped before deletion.
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	// StartService starts a service. This call does not block until the service is started.
	// This call will not error if the service is already running or has been started.
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceResponse, error)
	// StopService stops a service. This call does not block until the service is stopped.
	// This call will not error if the service is already stopped or has been stopped.
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceResponse, error)
	// UpsertAutomationsFile upserts the automations file for the given environment.
	UpsertAutomationsFile(ctx context.Context, in *UpsertAutomationsFileRequest, opts ...grpc.CallOption) (*UpsertAutomationsFileResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error)
	ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
	// StartTask starts a task, i.e. creates a task execution.
	// This call does not block until the task is started; the task will be started asynchronously.
	StartTask(ctx context.Context, in *StartTaskRequest, opts ...grpc.CallOption) (*StartTaskResponse, error)
	ListTaskExecutions(ctx context.Context, in *ListTaskExecutionsRequest, opts ...grpc.CallOption) (*ListTaskExecutionsResponse, error)
	GetTaskExecution(ctx context.Context, in *GetTaskExecutionRequest, opts ...grpc.CallOption) (*GetTaskExecutionResponse, error)
	StopTaskExecution(ctx context.Context, in *StopTaskExecutionRequest, opts ...grpc.CallOption) (*StopTaskExecutionResponse, error)
	// UpdateTaskExecutionStatus updates the status of a task execution. Only the environment executing a task execution is expected to call this function.
	UpdateTaskExecutionStatus(ctx context.Context, in *UpdateTaskExecutionStatusRequest, opts ...grpc.CallOption) (*UpdateTaskExecutionStatusResponse, error)
}

type environmentAutomationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentAutomationServiceClient(cc grpc.ClientConnInterface) EnvironmentAutomationServiceClient {
	return &environmentAutomationServiceClient{cc}
}

func (c *environmentAutomationServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_CreateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_GetService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_ListServices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*UpdateServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateServiceResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_UpdateService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_DeleteService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartServiceResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_StartService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopServiceResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_StopService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) UpsertAutomationsFile(ctx context.Context, in *UpsertAutomationsFileRequest, opts ...grpc.CallOption) (*UpsertAutomationsFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertAutomationsFileResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_UpsertAutomationsFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_GetTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) ListTasks(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_ListTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) StartTask(ctx context.Context, in *StartTaskRequest, opts ...grpc.CallOption) (*StartTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartTaskResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_StartTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) ListTaskExecutions(ctx context.Context, in *ListTaskExecutionsRequest, opts ...grpc.CallOption) (*ListTaskExecutionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTaskExecutionsResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_ListTaskExecutions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) GetTaskExecution(ctx context.Context, in *GetTaskExecutionRequest, opts ...grpc.CallOption) (*GetTaskExecutionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskExecutionResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_GetTaskExecution_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) StopTaskExecution(ctx context.Context, in *StopTaskExecutionRequest, opts ...grpc.CallOption) (*StopTaskExecutionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopTaskExecutionResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_StopTaskExecution_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentAutomationServiceClient) UpdateTaskExecutionStatus(ctx context.Context, in *UpdateTaskExecutionStatusRequest, opts ...grpc.CallOption) (*UpdateTaskExecutionStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTaskExecutionStatusResponse)
	err := c.cc.Invoke(ctx, EnvironmentAutomationService_UpdateTaskExecutionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentAutomationServiceServer is the server API for EnvironmentAutomationService service.
// All implementations must embed UnimplementedEnvironmentAutomationServiceServer
// for forward compatibility.
type EnvironmentAutomationServiceServer interface {
	CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error)
	// DeleteService deletes a service. This call does not block until the service is deleted.
	// If the service is not stopped it will be stopped before deletion.
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	// StartService starts a service. This call does not block until the service is started.
	// This call will not error if the service is already running or has been started.
	StartService(context.Context, *StartServiceRequest) (*StartServiceResponse, error)
	// StopService stops a service. This call does not block until the service is stopped.
	// This call will not error if the service is already stopped or has been stopped.
	StopService(context.Context, *StopServiceRequest) (*StopServiceResponse, error)
	// UpsertAutomationsFile upserts the automations file for the given environment.
	UpsertAutomationsFile(context.Context, *UpsertAutomationsFileRequest) (*UpsertAutomationsFileResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error)
	ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	// StartTask starts a task, i.e. creates a task execution.
	// This call does not block until the task is started; the task will be started asynchronously.
	StartTask(context.Context, *StartTaskRequest) (*StartTaskResponse, error)
	ListTaskExecutions(context.Context, *ListTaskExecutionsRequest) (*ListTaskExecutionsResponse, error)
	GetTaskExecution(context.Context, *GetTaskExecutionRequest) (*GetTaskExecutionResponse, error)
	StopTaskExecution(context.Context, *StopTaskExecutionRequest) (*StopTaskExecutionResponse, error)
	// UpdateTaskExecutionStatus updates the status of a task execution. Only the environment executing a task execution is expected to call this function.
	UpdateTaskExecutionStatus(context.Context, *UpdateTaskExecutionStatusRequest) (*UpdateTaskExecutionStatusResponse, error)
	mustEmbedUnimplementedEnvironmentAutomationServiceServer()
}

// UnimplementedEnvironmentAutomationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEnvironmentAutomationServiceServer struct{}

func (UnimplementedEnvironmentAutomationServiceServer) CreateService(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) UpdateService(context.Context, *UpdateServiceRequest) (*UpdateServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) StartService(context.Context, *StartServiceRequest) (*StartServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) StopService(context.Context, *StopServiceRequest) (*StopServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopService not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) UpsertAutomationsFile(context.Context, *UpsertAutomationsFileRequest) (*UpsertAutomationsFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAutomationsFile not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) ListTasks(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasks not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) StartTask(context.Context, *StartTaskRequest) (*StartTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTask not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) ListTaskExecutions(context.Context, *ListTaskExecutionsRequest) (*ListTaskExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskExecutions not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) GetTaskExecution(context.Context, *GetTaskExecutionRequest) (*GetTaskExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskExecution not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) StopTaskExecution(context.Context, *StopTaskExecutionRequest) (*StopTaskExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTaskExecution not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) UpdateTaskExecutionStatus(context.Context, *UpdateTaskExecutionStatusRequest) (*UpdateTaskExecutionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskExecutionStatus not implemented")
}
func (UnimplementedEnvironmentAutomationServiceServer) mustEmbedUnimplementedEnvironmentAutomationServiceServer() {
}
func (UnimplementedEnvironmentAutomationServiceServer) testEmbeddedByValue() {}

// UnsafeEnvironmentAutomationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentAutomationServiceServer will
// result in compilation errors.
type UnsafeEnvironmentAutomationServiceServer interface {
	mustEmbedUnimplementedEnvironmentAutomationServiceServer()
}

func RegisterEnvironmentAutomationServiceServer(s grpc.ServiceRegistrar, srv EnvironmentAutomationServiceServer) {
	// If the following call pancis, it indicates UnimplementedEnvironmentAutomationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EnvironmentAutomationService_ServiceDesc, srv)
}

func _EnvironmentAutomationService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_UpdateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_StartService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_StopService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_UpsertAutomationsFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAutomationsFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).UpsertAutomationsFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_UpsertAutomationsFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).UpsertAutomationsFile(ctx, req.(*UpsertAutomationsFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_ListTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).ListTasks(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_StartTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).StartTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_StartTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).StartTask(ctx, req.(*StartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_ListTaskExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).ListTaskExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_ListTaskExecutions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).ListTaskExecutions(ctx, req.(*ListTaskExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_GetTaskExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).GetTaskExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_GetTaskExecution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).GetTaskExecution(ctx, req.(*GetTaskExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_StopTaskExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).StopTaskExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_StopTaskExecution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).StopTaskExecution(ctx, req.(*StopTaskExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentAutomationService_UpdateTaskExecutionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskExecutionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentAutomationServiceServer).UpdateTaskExecutionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EnvironmentAutomationService_UpdateTaskExecutionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentAutomationServiceServer).UpdateTaskExecutionStatus(ctx, req.(*UpdateTaskExecutionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentAutomationService_ServiceDesc is the grpc.ServiceDesc for EnvironmentAutomationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentAutomationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.v1.EnvironmentAutomationService",
	HandlerType: (*EnvironmentAutomationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateService",
			Handler:    _EnvironmentAutomationService_CreateService_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _EnvironmentAutomationService_GetService_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _EnvironmentAutomationService_ListServices_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _EnvironmentAutomationService_UpdateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _EnvironmentAutomationService_DeleteService_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _EnvironmentAutomationService_StartService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _EnvironmentAutomationService_StopService_Handler,
		},
		{
			MethodName: "UpsertAutomationsFile",
			Handler:    _EnvironmentAutomationService_UpsertAutomationsFile_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _EnvironmentAutomationService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _EnvironmentAutomationService_GetTask_Handler,
		},
		{
			MethodName: "ListTasks",
			Handler:    _EnvironmentAutomationService_ListTasks_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _EnvironmentAutomationService_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _EnvironmentAutomationService_DeleteTask_Handler,
		},
		{
			MethodName: "StartTask",
			Handler:    _EnvironmentAutomationService_StartTask_Handler,
		},
		{
			MethodName: "ListTaskExecutions",
			Handler:    _EnvironmentAutomationService_ListTaskExecutions_Handler,
		},
		{
			MethodName: "GetTaskExecution",
			Handler:    _EnvironmentAutomationService_GetTaskExecution_Handler,
		},
		{
			MethodName: "StopTaskExecution",
			Handler:    _EnvironmentAutomationService_StopTaskExecution_Handler,
		},
		{
			MethodName: "UpdateTaskExecutionStatus",
			Handler:    _EnvironmentAutomationService_UpdateTaskExecutionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/v1/environment_automation.proto",
}
