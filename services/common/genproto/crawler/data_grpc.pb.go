// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: data.proto

package crawler

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
	CrawlerService_CreateCrawler_FullMethodName = "/CrawlerService/CreateCrawler"
	CrawlerService_GetCrawler_FullMethodName    = "/CrawlerService/GetCrawler"
)

// CrawlerServiceClient is the client API for CrawlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlerServiceClient interface {
	CreateCrawler(ctx context.Context, in *CreateCrawlerRequest, opts ...grpc.CallOption) (*CreateCrawlerResponse, error)
	GetCrawler(ctx context.Context, in *GetCrawlerRequest, opts ...grpc.CallOption) (*GetCrawlerResponse, error)
}

type crawlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlerServiceClient(cc grpc.ClientConnInterface) CrawlerServiceClient {
	return &crawlerServiceClient{cc}
}

func (c *crawlerServiceClient) CreateCrawler(ctx context.Context, in *CreateCrawlerRequest, opts ...grpc.CallOption) (*CreateCrawlerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCrawlerResponse)
	err := c.cc.Invoke(ctx, CrawlerService_CreateCrawler_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerServiceClient) GetCrawler(ctx context.Context, in *GetCrawlerRequest, opts ...grpc.CallOption) (*GetCrawlerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCrawlerResponse)
	err := c.cc.Invoke(ctx, CrawlerService_GetCrawler_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrawlerServiceServer is the server API for CrawlerService service.
// All implementations must embed UnimplementedCrawlerServiceServer
// for forward compatibility.
type CrawlerServiceServer interface {
	CreateCrawler(context.Context, *CreateCrawlerRequest) (*CreateCrawlerResponse, error)
	GetCrawler(context.Context, *GetCrawlerRequest) (*GetCrawlerResponse, error)
	mustEmbedUnimplementedCrawlerServiceServer()
}

// UnimplementedCrawlerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCrawlerServiceServer struct{}

func (UnimplementedCrawlerServiceServer) CreateCrawler(context.Context, *CreateCrawlerRequest) (*CreateCrawlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCrawler not implemented")
}
func (UnimplementedCrawlerServiceServer) GetCrawler(context.Context, *GetCrawlerRequest) (*GetCrawlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrawler not implemented")
}
func (UnimplementedCrawlerServiceServer) mustEmbedUnimplementedCrawlerServiceServer() {}
func (UnimplementedCrawlerServiceServer) testEmbeddedByValue()                        {}

// UnsafeCrawlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlerServiceServer will
// result in compilation errors.
type UnsafeCrawlerServiceServer interface {
	mustEmbedUnimplementedCrawlerServiceServer()
}

func RegisterCrawlerServiceServer(s grpc.ServiceRegistrar, srv CrawlerServiceServer) {
	// If the following call pancis, it indicates UnimplementedCrawlerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CrawlerService_ServiceDesc, srv)
}

func _CrawlerService_CreateCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCrawlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).CreateCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlerService_CreateCrawler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).CreateCrawler(ctx, req.(*CreateCrawlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlerService_GetCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrawlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).GetCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlerService_GetCrawler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).GetCrawler(ctx, req.(*GetCrawlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CrawlerService_ServiceDesc is the grpc.ServiceDesc for CrawlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CrawlerService",
	HandlerType: (*CrawlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCrawler",
			Handler:    _CrawlerService_CreateCrawler_Handler,
		},
		{
			MethodName: "GetCrawler",
			Handler:    _CrawlerService_GetCrawler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
