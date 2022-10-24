// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: agentpb/agent.proto

package agentpb

import (
	context "context"
	taskpb "github.com/devilsec/btedr/proto/taskpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentRPCClient is the client API for AgentRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentRPCClient interface {
	Register(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error)
	// Retrieve a task from the server (beacon mode)
	GetTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*taskpb.Task, error)
	TaskResult(ctx context.Context, in *Result, opts ...grpc.CallOption) (*Empty, error)
}

type agentRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentRPCClient(cc grpc.ClientConnInterface) AgentRPCClient {
	return &agentRPCClient{cc}
}

func (c *agentRPCClient) Register(ctx context.Context, in *Registration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/btedrpb.AgentRPC/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRPCClient) GetTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*taskpb.Task, error) {
	out := new(taskpb.Task)
	err := c.cc.Invoke(ctx, "/btedrpb.AgentRPC/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRPCClient) TaskResult(ctx context.Context, in *Result, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/btedrpb.AgentRPC/TaskResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentRPCServer is the server API for AgentRPC service.
// All implementations must embed UnimplementedAgentRPCServer
// for forward compatibility
type AgentRPCServer interface {
	Register(context.Context, *Registration) (*Empty, error)
	// Retrieve a task from the server (beacon mode)
	GetTask(context.Context, *Request) (*taskpb.Task, error)
	TaskResult(context.Context, *Result) (*Empty, error)
	mustEmbedUnimplementedAgentRPCServer()
}

// UnimplementedAgentRPCServer must be embedded to have forward compatible implementations.
type UnimplementedAgentRPCServer struct {
}

func (UnimplementedAgentRPCServer) Register(context.Context, *Registration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAgentRPCServer) GetTask(context.Context, *Request) (*taskpb.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedAgentRPCServer) TaskResult(context.Context, *Result) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskResult not implemented")
}
func (UnimplementedAgentRPCServer) mustEmbedUnimplementedAgentRPCServer() {}

// UnsafeAgentRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentRPCServer will
// result in compilation errors.
type UnsafeAgentRPCServer interface {
	mustEmbedUnimplementedAgentRPCServer()
}

func RegisterAgentRPCServer(s grpc.ServiceRegistrar, srv AgentRPCServer) {
	s.RegisterService(&AgentRPC_ServiceDesc, srv)
}

func _AgentRPC_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRPCServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btedrpb.AgentRPC/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRPCServer).Register(ctx, req.(*Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRPC_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRPCServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btedrpb.AgentRPC/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRPCServer).GetTask(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRPC_TaskResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Result)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRPCServer).TaskResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/btedrpb.AgentRPC/TaskResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRPCServer).TaskResult(ctx, req.(*Result))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentRPC_ServiceDesc is the grpc.ServiceDesc for AgentRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "btedrpb.AgentRPC",
	HandlerType: (*AgentRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AgentRPC_Register_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _AgentRPC_GetTask_Handler,
		},
		{
			MethodName: "TaskResult",
			Handler:    _AgentRPC_TaskResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agentpb/agent.proto",
}
