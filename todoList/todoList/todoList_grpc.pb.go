// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: todoList.proto

package todoList

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

const (
	Adder_GetTodoList_FullMethodName = "/todoList.adder/getTodoList"
)

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderClient interface {
	GetTodoList(ctx context.Context, in *GetTodoListReq, opts ...grpc.CallOption) (*GetTodoListResp, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) GetTodoList(ctx context.Context, in *GetTodoListReq, opts ...grpc.CallOption) (*GetTodoListResp, error) {
	out := new(GetTodoListResp)
	err := c.cc.Invoke(ctx, Adder_GetTodoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServer is the server API for Adder service.
// All implementations must embed UnimplementedAdderServer
// for forward compatibility
type AdderServer interface {
	GetTodoList(context.Context, *GetTodoListReq) (*GetTodoListResp, error)
	mustEmbedUnimplementedAdderServer()
}

// UnimplementedAdderServer must be embedded to have forward compatible implementations.
type UnimplementedAdderServer struct {
}

func (UnimplementedAdderServer) GetTodoList(context.Context, *GetTodoListReq) (*GetTodoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoList not implemented")
}
func (UnimplementedAdderServer) mustEmbedUnimplementedAdderServer() {}

// UnsafeAdderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdderServer will
// result in compilation errors.
type UnsafeAdderServer interface {
	mustEmbedUnimplementedAdderServer()
}

func RegisterAdderServer(s grpc.ServiceRegistrar, srv AdderServer) {
	s.RegisterService(&Adder_ServiceDesc, srv)
}

func _Adder_GetTodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).GetTodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Adder_GetTodoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).GetTodoList(ctx, req.(*GetTodoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Adder_ServiceDesc is the grpc.ServiceDesc for Adder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Adder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todoList.adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getTodoList",
			Handler:    _Adder_GetTodoList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todoList.proto",
}