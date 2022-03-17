// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_grpc

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

// ExGoGrpcClient is the client API for ExGoGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExGoGrpcClient interface {
	HelloGrpc(ctx context.Context, in *HelloGrpcRequest, opts ...grpc.CallOption) (*HelloGrpcResponse, error)
}

type exGoGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewExGoGrpcClient(cc grpc.ClientConnInterface) ExGoGrpcClient {
	return &exGoGrpcClient{cc}
}

func (c *exGoGrpcClient) HelloGrpc(ctx context.Context, in *HelloGrpcRequest, opts ...grpc.CallOption) (*HelloGrpcResponse, error) {
	out := new(HelloGrpcResponse)
	err := c.cc.Invoke(ctx, "/exgogrpc.ExGoGrpc/HelloGrpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExGoGrpcServer is the server API for ExGoGrpc service.
// All implementations must embed UnimplementedExGoGrpcServer
// for forward compatibility
type ExGoGrpcServer interface {
	HelloGrpc(context.Context, *HelloGrpcRequest) (*HelloGrpcResponse, error)
	mustEmbedUnimplementedExGoGrpcServer()
}

// UnimplementedExGoGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedExGoGrpcServer struct {
}

func (UnimplementedExGoGrpcServer) HelloGrpc(context.Context, *HelloGrpcRequest) (*HelloGrpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloGrpc not implemented")
}
func (UnimplementedExGoGrpcServer) mustEmbedUnimplementedExGoGrpcServer() {}

// UnsafeExGoGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExGoGrpcServer will
// result in compilation errors.
type UnsafeExGoGrpcServer interface {
	mustEmbedUnimplementedExGoGrpcServer()
}

func RegisterExGoGrpcServer(s grpc.ServiceRegistrar, srv ExGoGrpcServer) {
	s.RegisterService(&ExGoGrpc_ServiceDesc, srv)
}

func _ExGoGrpc_HelloGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloGrpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExGoGrpcServer).HelloGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exgogrpc.ExGoGrpc/HelloGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExGoGrpcServer).HelloGrpc(ctx, req.(*HelloGrpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExGoGrpc_ServiceDesc is the grpc.ServiceDesc for ExGoGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExGoGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exgogrpc.ExGoGrpc",
	HandlerType: (*ExGoGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloGrpc",
			Handler:    _ExGoGrpc_HelloGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/go_grpc/go_grpc.proto",
}