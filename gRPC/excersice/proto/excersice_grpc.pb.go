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
	ServerStreamPrimeNumber(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (ExcersiceService_ServerStreamPrimeNumberClient, error)
	ClientStreamAvg(ctx context.Context, opts ...grpc.CallOption) (ExcersiceService_ClientStreamAvgClient, error)
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

func (c *excersiceServiceClient) ServerStreamPrimeNumber(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (ExcersiceService_ServerStreamPrimeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExcersiceService_ServiceDesc.Streams[0], "/excersice.ExcersiceService/ServerStreamPrimeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &excersiceServiceServerStreamPrimeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExcersiceService_ServerStreamPrimeNumberClient interface {
	Recv() (*ServerStreamResponse, error)
	grpc.ClientStream
}

type excersiceServiceServerStreamPrimeNumberClient struct {
	grpc.ClientStream
}

func (x *excersiceServiceServerStreamPrimeNumberClient) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *excersiceServiceClient) ClientStreamAvg(ctx context.Context, opts ...grpc.CallOption) (ExcersiceService_ClientStreamAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExcersiceService_ServiceDesc.Streams[1], "/excersice.ExcersiceService/ClientStreamAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &excersiceServiceClientStreamAvgClient{stream}
	return x, nil
}

type ExcersiceService_ClientStreamAvgClient interface {
	Send(*ClientStreamRequest) error
	CloseAndRecv() (*ClientStreamResponse, error)
	grpc.ClientStream
}

type excersiceServiceClientStreamAvgClient struct {
	grpc.ClientStream
}

func (x *excersiceServiceClientStreamAvgClient) Send(m *ClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *excersiceServiceClientStreamAvgClient) CloseAndRecv() (*ClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExcersiceServiceServer is the server API for ExcersiceService service.
// All implementations must embed UnimplementedExcersiceServiceServer
// for forward compatibility
type ExcersiceServiceServer interface {
	UnarySum(context.Context, *UnarySumReqest) (*UnarySumResponse, error)
	ServerStreamPrimeNumber(*ServerStreamRequest, ExcersiceService_ServerStreamPrimeNumberServer) error
	ClientStreamAvg(ExcersiceService_ClientStreamAvgServer) error
	mustEmbedUnimplementedExcersiceServiceServer()
}

// UnimplementedExcersiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExcersiceServiceServer struct {
}

func (UnimplementedExcersiceServiceServer) UnarySum(context.Context, *UnarySumReqest) (*UnarySumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnarySum not implemented")
}
func (UnimplementedExcersiceServiceServer) ServerStreamPrimeNumber(*ServerStreamRequest, ExcersiceService_ServerStreamPrimeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamPrimeNumber not implemented")
}
func (UnimplementedExcersiceServiceServer) ClientStreamAvg(ExcersiceService_ClientStreamAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamAvg not implemented")
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

func _ExcersiceService_ServerStreamPrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExcersiceServiceServer).ServerStreamPrimeNumber(m, &excersiceServiceServerStreamPrimeNumberServer{stream})
}

type ExcersiceService_ServerStreamPrimeNumberServer interface {
	Send(*ServerStreamResponse) error
	grpc.ServerStream
}

type excersiceServiceServerStreamPrimeNumberServer struct {
	grpc.ServerStream
}

func (x *excersiceServiceServerStreamPrimeNumberServer) Send(m *ServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ExcersiceService_ClientStreamAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExcersiceServiceServer).ClientStreamAvg(&excersiceServiceClientStreamAvgServer{stream})
}

type ExcersiceService_ClientStreamAvgServer interface {
	SendAndClose(*ClientStreamResponse) error
	Recv() (*ClientStreamRequest, error)
	grpc.ServerStream
}

type excersiceServiceClientStreamAvgServer struct {
	grpc.ServerStream
}

func (x *excersiceServiceClientStreamAvgServer) SendAndClose(m *ClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *excersiceServiceClientStreamAvgServer) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamPrimeNumber",
			Handler:       _ExcersiceService_ServerStreamPrimeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamAvg",
			Handler:       _ExcersiceService_ClientStreamAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "excersice.proto",
}
