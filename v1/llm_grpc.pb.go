// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: gitpod/v1/llm.proto

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
	LLMService_InvokeModel_FullMethodName = "/gitpod.v1.LLMService/InvokeModel"
)

// LLMServiceClient is the client API for LLMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Experimental LLM service. This will go away once we stabilize that API.
// Do not rely on this service at all.
type LLMServiceClient interface {
	InvokeModel(ctx context.Context, in *InvokeModelRequest, opts ...grpc.CallOption) (*InvokeModelResponse, error)
}

type lLMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLLMServiceClient(cc grpc.ClientConnInterface) LLMServiceClient {
	return &lLMServiceClient{cc}
}

func (c *lLMServiceClient) InvokeModel(ctx context.Context, in *InvokeModelRequest, opts ...grpc.CallOption) (*InvokeModelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InvokeModelResponse)
	err := c.cc.Invoke(ctx, LLMService_InvokeModel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LLMServiceServer is the server API for LLMService service.
// All implementations must embed UnimplementedLLMServiceServer
// for forward compatibility.
//
// Experimental LLM service. This will go away once we stabilize that API.
// Do not rely on this service at all.
type LLMServiceServer interface {
	InvokeModel(context.Context, *InvokeModelRequest) (*InvokeModelResponse, error)
	mustEmbedUnimplementedLLMServiceServer()
}

// UnimplementedLLMServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLLMServiceServer struct{}

func (UnimplementedLLMServiceServer) InvokeModel(context.Context, *InvokeModelRequest) (*InvokeModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeModel not implemented")
}
func (UnimplementedLLMServiceServer) mustEmbedUnimplementedLLMServiceServer() {}
func (UnimplementedLLMServiceServer) testEmbeddedByValue()                    {}

// UnsafeLLMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LLMServiceServer will
// result in compilation errors.
type UnsafeLLMServiceServer interface {
	mustEmbedUnimplementedLLMServiceServer()
}

func RegisterLLMServiceServer(s grpc.ServiceRegistrar, srv LLMServiceServer) {
	// If the following call pancis, it indicates UnimplementedLLMServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LLMService_ServiceDesc, srv)
}

func _LLMService_InvokeModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServiceServer).InvokeModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LLMService_InvokeModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServiceServer).InvokeModel(ctx, req.(*InvokeModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LLMService_ServiceDesc is the grpc.ServiceDesc for LLMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LLMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.v1.LLMService",
	HandlerType: (*LLMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InvokeModel",
			Handler:    _LLMService_InvokeModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/v1/llm.proto",
}
