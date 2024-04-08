// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: excersice.proto

package proto

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

// ExcersiceServiceClient is the client API for ExcersiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExcersiceServiceClient interface {
	UnarySum(ctx context.Context, in *UnarySumReqest, opts ...grpc.CallOption) (*UnarySumResponse, error)
}

type excersiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExcersiceServiceClient(cc grpc.ClientConnInterface) ExcersiceServiceClient {
	return &excersiceServiceClient{cc}
}

func (c *excersiceServiceClient) UnarySum(ctx context.Context, in *UnarySumReqest, opts ...grpc.CallOption) (*UnarySumResponse, error) {
	out := new(UnarySumResponse)
	err := c.cc.Invoke(ctx, "/excersice.ExcersiceService/UnarySum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExcersiceServiceServer is the server API for ExcersiceService service.
// All implementations must embed UnimplementedExcersiceServiceServer
// for forward compatibility
type ExcersiceServiceServer interface {
	UnarySum(context.Context, *UnarySumReqest) (*UnarySumResponse, error)
	mustEmbedUnimplementedExcersiceServiceServer()
}

// UnimplementedExcersiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExcersiceServiceServer struct {
}

func (UnimplementedExcersiceServiceServer) UnarySum(context.Context, *UnarySumReqest) (*UnarySumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnarySum not implemented")
}
func (UnimplementedExcersiceServiceServer) mustEmbedUnimplementedExcersiceServiceServer() {}

// UnsafeExcersiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExcersiceServiceServer will
// result in compilation errors.
type UnsafeExcersiceServiceServer interface {
	mustEmbedUnimplementedExcersiceServiceServer()
}

func RegisterExcersiceServiceServer(s grpc.ServiceRegistrar, srv ExcersiceServiceServer) {
	s.RegisterService(&ExcersiceService_ServiceDesc, srv)
}

func _ExcersiceService_UnarySum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnarySumReqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExcersiceServiceServer).UnarySum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/excersice.ExcersiceService/UnarySum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExcersiceServiceServer).UnarySum(ctx, req.(*UnarySumReqest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExcersiceService_ServiceDesc is the grpc.ServiceDesc for ExcersiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExcersiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "excersice.ExcersiceService",
	HandlerType: (*ExcersiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnarySum",
			Handler:    _ExcersiceService_UnarySum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "excersice.proto",
}
