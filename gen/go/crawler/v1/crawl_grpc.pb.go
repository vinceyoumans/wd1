// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: crawler/v1/crawl.proto

package crawlerpb

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

// CrawlerServiceClient is the client API for CrawlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlerServiceClient interface {
	AddCrawl(ctx context.Context, in *AddCrawlRequest, opts ...grpc.CallOption) (*AddCrawlResponse, error)
	ListCrawl(ctx context.Context, in *ListCrawlRequest, opts ...grpc.CallOption) (*ListCrawlResponse, error)
	LongListCrawl(ctx context.Context, in *LongListCrawlRequest, opts ...grpc.CallOption) (*LongListCrawlResponse, error)
	MonitorCrawl(ctx context.Context, in *MonitorCrawlRequest, opts ...grpc.CallOption) (CrawlerService_MonitorCrawlClient, error)
	MonitorCrawl01(ctx context.Context, in *MonitorCrawl01Request, opts ...grpc.CallOption) (CrawlerService_MonitorCrawl01Client, error)
}

type crawlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlerServiceClient(cc grpc.ClientConnInterface) CrawlerServiceClient {
	return &crawlerServiceClient{cc}
}

func (c *crawlerServiceClient) AddCrawl(ctx context.Context, in *AddCrawlRequest, opts ...grpc.CallOption) (*AddCrawlResponse, error) {
	out := new(AddCrawlResponse)
	err := c.cc.Invoke(ctx, "/crawler.v1.CrawlerService/AddCrawl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerServiceClient) ListCrawl(ctx context.Context, in *ListCrawlRequest, opts ...grpc.CallOption) (*ListCrawlResponse, error) {
	out := new(ListCrawlResponse)
	err := c.cc.Invoke(ctx, "/crawler.v1.CrawlerService/ListCrawl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerServiceClient) LongListCrawl(ctx context.Context, in *LongListCrawlRequest, opts ...grpc.CallOption) (*LongListCrawlResponse, error) {
	out := new(LongListCrawlResponse)
	err := c.cc.Invoke(ctx, "/crawler.v1.CrawlerService/LongListCrawl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerServiceClient) MonitorCrawl(ctx context.Context, in *MonitorCrawlRequest, opts ...grpc.CallOption) (CrawlerService_MonitorCrawlClient, error) {
	stream, err := c.cc.NewStream(ctx, &CrawlerService_ServiceDesc.Streams[0], "/crawler.v1.CrawlerService/MonitorCrawl", opts...)
	if err != nil {
		return nil, err
	}
	x := &crawlerServiceMonitorCrawlClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrawlerService_MonitorCrawlClient interface {
	Recv() (*MonitorCrawlResponse, error)
	grpc.ClientStream
}

type crawlerServiceMonitorCrawlClient struct {
	grpc.ClientStream
}

func (x *crawlerServiceMonitorCrawlClient) Recv() (*MonitorCrawlResponse, error) {
	m := new(MonitorCrawlResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crawlerServiceClient) MonitorCrawl01(ctx context.Context, in *MonitorCrawl01Request, opts ...grpc.CallOption) (CrawlerService_MonitorCrawl01Client, error) {
	stream, err := c.cc.NewStream(ctx, &CrawlerService_ServiceDesc.Streams[1], "/crawler.v1.CrawlerService/MonitorCrawl01", opts...)
	if err != nil {
		return nil, err
	}
	x := &crawlerServiceMonitorCrawl01Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrawlerService_MonitorCrawl01Client interface {
	Recv() (*MonitorCrawl01Response, error)
	grpc.ClientStream
}

type crawlerServiceMonitorCrawl01Client struct {
	grpc.ClientStream
}

func (x *crawlerServiceMonitorCrawl01Client) Recv() (*MonitorCrawl01Response, error) {
	m := new(MonitorCrawl01Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CrawlerServiceServer is the server API for CrawlerService service.
// All implementations must embed UnimplementedCrawlerServiceServer
// for forward compatibility
type CrawlerServiceServer interface {
	AddCrawl(context.Context, *AddCrawlRequest) (*AddCrawlResponse, error)
	// rpc CancelCrawl (CancelCrawlRequest) returns (CancelCrawlResponse) {} // cancels a specific job
	ListCrawl(context.Context, *ListCrawlRequest) (*ListCrawlResponse, error)
	LongListCrawl(context.Context, *LongListCrawlRequest) (*LongListCrawlResponse, error)
	MonitorCrawl(*MonitorCrawlRequest, CrawlerService_MonitorCrawlServer) error
	MonitorCrawl01(*MonitorCrawl01Request, CrawlerService_MonitorCrawl01Server) error
	mustEmbedUnimplementedCrawlerServiceServer()
}

// UnimplementedCrawlerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrawlerServiceServer struct {
}

func (UnimplementedCrawlerServiceServer) AddCrawl(context.Context, *AddCrawlRequest) (*AddCrawlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCrawl not implemented")
}
func (UnimplementedCrawlerServiceServer) ListCrawl(context.Context, *ListCrawlRequest) (*ListCrawlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCrawl not implemented")
}
func (UnimplementedCrawlerServiceServer) LongListCrawl(context.Context, *LongListCrawlRequest) (*LongListCrawlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LongListCrawl not implemented")
}
func (UnimplementedCrawlerServiceServer) MonitorCrawl(*MonitorCrawlRequest, CrawlerService_MonitorCrawlServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorCrawl not implemented")
}
func (UnimplementedCrawlerServiceServer) MonitorCrawl01(*MonitorCrawl01Request, CrawlerService_MonitorCrawl01Server) error {
	return status.Errorf(codes.Unimplemented, "method MonitorCrawl01 not implemented")
}
func (UnimplementedCrawlerServiceServer) mustEmbedUnimplementedCrawlerServiceServer() {}

// UnsafeCrawlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlerServiceServer will
// result in compilation errors.
type UnsafeCrawlerServiceServer interface {
	mustEmbedUnimplementedCrawlerServiceServer()
}

func RegisterCrawlerServiceServer(s grpc.ServiceRegistrar, srv CrawlerServiceServer) {
	s.RegisterService(&CrawlerService_ServiceDesc, srv)
}

func _CrawlerService_AddCrawl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCrawlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).AddCrawl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crawler.v1.CrawlerService/AddCrawl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).AddCrawl(ctx, req.(*AddCrawlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlerService_ListCrawl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCrawlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).ListCrawl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crawler.v1.CrawlerService/ListCrawl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).ListCrawl(ctx, req.(*ListCrawlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlerService_LongListCrawl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongListCrawlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).LongListCrawl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crawler.v1.CrawlerService/LongListCrawl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).LongListCrawl(ctx, req.(*LongListCrawlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlerService_MonitorCrawl_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorCrawlRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlerServiceServer).MonitorCrawl(m, &crawlerServiceMonitorCrawlServer{stream})
}

type CrawlerService_MonitorCrawlServer interface {
	Send(*MonitorCrawlResponse) error
	grpc.ServerStream
}

type crawlerServiceMonitorCrawlServer struct {
	grpc.ServerStream
}

func (x *crawlerServiceMonitorCrawlServer) Send(m *MonitorCrawlResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrawlerService_MonitorCrawl01_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorCrawl01Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrawlerServiceServer).MonitorCrawl01(m, &crawlerServiceMonitorCrawl01Server{stream})
}

type CrawlerService_MonitorCrawl01Server interface {
	Send(*MonitorCrawl01Response) error
	grpc.ServerStream
}

type crawlerServiceMonitorCrawl01Server struct {
	grpc.ServerStream
}

func (x *crawlerServiceMonitorCrawl01Server) Send(m *MonitorCrawl01Response) error {
	return x.ServerStream.SendMsg(m)
}

// CrawlerService_ServiceDesc is the grpc.ServiceDesc for CrawlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crawler.v1.CrawlerService",
	HandlerType: (*CrawlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCrawl",
			Handler:    _CrawlerService_AddCrawl_Handler,
		},
		{
			MethodName: "ListCrawl",
			Handler:    _CrawlerService_ListCrawl_Handler,
		},
		{
			MethodName: "LongListCrawl",
			Handler:    _CrawlerService_LongListCrawl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorCrawl",
			Handler:       _CrawlerService_MonitorCrawl_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MonitorCrawl01",
			Handler:       _CrawlerService_MonitorCrawl01_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "crawler/v1/crawl.proto",
}
