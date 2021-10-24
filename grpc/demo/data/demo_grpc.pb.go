// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data

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

// DemoClient is the client API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoClient interface {
	// 简单模式。一个请求，一个响应。
	Add(ctx context.Context, in *TwoNum, opts ...grpc.CallOption) (*Response, error)
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	//服务端流模式，客户端发送一个请求，服务端返回多次。
	GetStream(ctx context.Context, in *TwoNum, opts ...grpc.CallOption) (Demo_GetStreamClient, error)
	//客户端流模式，客户端发送多次请求，服务端响应一次。
	PutStream(ctx context.Context, opts ...grpc.CallOption) (Demo_PutStreamClient, error)
	//双向流，发送和接收同时进行，互不干扰
	DoubleStream(ctx context.Context, opts ...grpc.CallOption) (Demo_DoubleStreamClient, error)
}

type demoClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoClient(cc grpc.ClientConnInterface) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) Add(ctx context.Context, in *TwoNum, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/demo.Demo/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/demo.Demo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) GetStream(ctx context.Context, in *TwoNum, opts ...grpc.CallOption) (Demo_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Demo_ServiceDesc.Streams[0], "/demo.Demo/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Demo_GetStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type demoGetStreamClient struct {
	grpc.ClientStream
}

func (x *demoGetStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoClient) PutStream(ctx context.Context, opts ...grpc.CallOption) (Demo_PutStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Demo_ServiceDesc.Streams[1], "/demo.Demo/PutStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoPutStreamClient{stream}
	return x, nil
}

type Demo_PutStreamClient interface {
	Send(*OneNum) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type demoPutStreamClient struct {
	grpc.ClientStream
}

func (x *demoPutStreamClient) Send(m *OneNum) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoPutStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoClient) DoubleStream(ctx context.Context, opts ...grpc.CallOption) (Demo_DoubleStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Demo_ServiceDesc.Streams[2], "/demo.Demo/DoubleStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoDoubleStreamClient{stream}
	return x, nil
}

type Demo_DoubleStreamClient interface {
	Send(*TwoNum) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type demoDoubleStreamClient struct {
	grpc.ClientStream
}

func (x *demoDoubleStreamClient) Send(m *TwoNum) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoDoubleStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoServer is the server API for Demo service.
// All implementations must embed UnimplementedDemoServer
// for forward compatibility
type DemoServer interface {
	// 简单模式。一个请求，一个响应。
	Add(context.Context, *TwoNum) (*Response, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	//服务端流模式，客户端发送一个请求，服务端返回多次。
	GetStream(*TwoNum, Demo_GetStreamServer) error
	//客户端流模式，客户端发送多次请求，服务端响应一次。
	PutStream(Demo_PutStreamServer) error
	//双向流，发送和接收同时进行，互不干扰
	DoubleStream(Demo_DoubleStreamServer) error
	mustEmbedUnimplementedDemoServer()
}

// UnimplementedDemoServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServer struct {
}

func (UnimplementedDemoServer) Add(context.Context, *TwoNum) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDemoServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDemoServer) GetStream(*TwoNum, Demo_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedDemoServer) PutStream(Demo_PutStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutStream not implemented")
}
func (UnimplementedDemoServer) DoubleStream(Demo_DoubleStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoubleStream not implemented")
}
func (UnimplementedDemoServer) mustEmbedUnimplementedDemoServer() {}

// UnsafeDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServer will
// result in compilation errors.
type UnsafeDemoServer interface {
	mustEmbedUnimplementedDemoServer()
}

func RegisterDemoServer(s grpc.ServiceRegistrar, srv DemoServer) {
	s.RegisterService(&Demo_ServiceDesc, srv)
}

func _Demo_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwoNum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Add(ctx, req.(*TwoNum))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.Demo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TwoNum)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServer).GetStream(m, &demoGetStreamServer{stream})
}

type Demo_GetStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type demoGetStreamServer struct {
	grpc.ServerStream
}

func (x *demoGetStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Demo_PutStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).PutStream(&demoPutStreamServer{stream})
}

type Demo_PutStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*OneNum, error)
	grpc.ServerStream
}

type demoPutStreamServer struct {
	grpc.ServerStream
}

func (x *demoPutStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoPutStreamServer) Recv() (*OneNum, error) {
	m := new(OneNum)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Demo_DoubleStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).DoubleStream(&demoDoubleStreamServer{stream})
}

type Demo_DoubleStreamServer interface {
	Send(*Response) error
	Recv() (*TwoNum, error)
	grpc.ServerStream
}

type demoDoubleStreamServer struct {
	grpc.ServerStream
}

func (x *demoDoubleStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoDoubleStreamServer) Recv() (*TwoNum, error) {
	m := new(TwoNum)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Demo_ServiceDesc is the grpc.ServiceDesc for Demo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Demo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Demo_Add_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Demo_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Demo_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutStream",
			Handler:       _Demo_PutStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoubleStream",
			Handler:       _Demo_DoubleStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demo.proto",
}
