// cartcast/proto/cartcast.proto
// This file defines the CartCast service and the associated message types for handling
// upselling cart data in a gRPC context.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: cartcast.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CartCastService_CreateCartCast_FullMethodName       = "/upselling.CartCastService/CreateCartCast"
	CartCastService_ReadCartCast_FullMethodName         = "/upselling.CartCastService/ReadCartCast"
	CartCastService_ReadByCoordsCartCast_FullMethodName = "/upselling.CartCastService/ReadByCoordsCartCast"
	CartCastService_UpdateCartCast_FullMethodName       = "/upselling.CartCastService/UpdateCartCast"
	CartCastService_DeleteCartCast_FullMethodName       = "/upselling.CartCastService/DeleteCartCast"
	CartCastService_ListCartCast_FullMethodName         = "/upselling.CartCastService/ListCartCast"
)

// CartCastServiceClient is the client API for CartCastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the CartCastService service with RPC methods for managing upselling items.
type CartCastServiceClient interface {
	CreateCartCast(ctx context.Context, in *CartCast, opts ...grpc.CallOption) (*CartCastId, error)
	ReadCartCast(ctx context.Context, in *CartCastId, opts ...grpc.CallOption) (*CartCast, error)
	ReadByCoordsCartCast(ctx context.Context, in *CartCastCoords, opts ...grpc.CallOption) (*CartCast, error)
	UpdateCartCast(ctx context.Context, in *CartCast, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCartCast(ctx context.Context, in *CartCastId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListCartCast(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CartCast], error)
}

type cartCastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartCastServiceClient(cc grpc.ClientConnInterface) CartCastServiceClient {
	return &cartCastServiceClient{cc}
}

func (c *cartCastServiceClient) CreateCartCast(ctx context.Context, in *CartCast, opts ...grpc.CallOption) (*CartCastId, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartCastId)
	err := c.cc.Invoke(ctx, CartCastService_CreateCartCast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartCastServiceClient) ReadCartCast(ctx context.Context, in *CartCastId, opts ...grpc.CallOption) (*CartCast, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartCast)
	err := c.cc.Invoke(ctx, CartCastService_ReadCartCast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartCastServiceClient) ReadByCoordsCartCast(ctx context.Context, in *CartCastCoords, opts ...grpc.CallOption) (*CartCast, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartCast)
	err := c.cc.Invoke(ctx, CartCastService_ReadByCoordsCartCast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartCastServiceClient) UpdateCartCast(ctx context.Context, in *CartCast, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CartCastService_UpdateCartCast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartCastServiceClient) DeleteCartCast(ctx context.Context, in *CartCastId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CartCastService_DeleteCartCast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartCastServiceClient) ListCartCast(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CartCast], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CartCastService_ServiceDesc.Streams[0], CartCastService_ListCartCast_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, CartCast]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CartCastService_ListCartCastClient = grpc.ServerStreamingClient[CartCast]

// CartCastServiceServer is the server API for CartCastService service.
// All implementations must embed UnimplementedCartCastServiceServer
// for forward compatibility.
//
// Define the CartCastService service with RPC methods for managing upselling items.
type CartCastServiceServer interface {
	CreateCartCast(context.Context, *CartCast) (*CartCastId, error)
	ReadCartCast(context.Context, *CartCastId) (*CartCast, error)
	ReadByCoordsCartCast(context.Context, *CartCastCoords) (*CartCast, error)
	UpdateCartCast(context.Context, *CartCast) (*emptypb.Empty, error)
	DeleteCartCast(context.Context, *CartCastId) (*emptypb.Empty, error)
	ListCartCast(*emptypb.Empty, grpc.ServerStreamingServer[CartCast]) error
	mustEmbedUnimplementedCartCastServiceServer()
}

// UnimplementedCartCastServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCartCastServiceServer struct{}

func (UnimplementedCartCastServiceServer) CreateCartCast(context.Context, *CartCast) (*CartCastId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCartCast not implemented")
}
func (UnimplementedCartCastServiceServer) ReadCartCast(context.Context, *CartCastId) (*CartCast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCartCast not implemented")
}
func (UnimplementedCartCastServiceServer) ReadByCoordsCartCast(context.Context, *CartCastCoords) (*CartCast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadByCoordsCartCast not implemented")
}
func (UnimplementedCartCastServiceServer) UpdateCartCast(context.Context, *CartCast) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartCast not implemented")
}
func (UnimplementedCartCastServiceServer) DeleteCartCast(context.Context, *CartCastId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartCast not implemented")
}
func (UnimplementedCartCastServiceServer) ListCartCast(*emptypb.Empty, grpc.ServerStreamingServer[CartCast]) error {
	return status.Errorf(codes.Unimplemented, "method ListCartCast not implemented")
}
func (UnimplementedCartCastServiceServer) mustEmbedUnimplementedCartCastServiceServer() {}
func (UnimplementedCartCastServiceServer) testEmbeddedByValue()                         {}

// UnsafeCartCastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartCastServiceServer will
// result in compilation errors.
type UnsafeCartCastServiceServer interface {
	mustEmbedUnimplementedCartCastServiceServer()
}

func RegisterCartCastServiceServer(s grpc.ServiceRegistrar, srv CartCastServiceServer) {
	// If the following call pancis, it indicates UnimplementedCartCastServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CartCastService_ServiceDesc, srv)
}

func _CartCastService_CreateCartCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCast)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartCastServiceServer).CreateCartCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartCastService_CreateCartCast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartCastServiceServer).CreateCartCast(ctx, req.(*CartCast))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartCastService_ReadCartCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCastId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartCastServiceServer).ReadCartCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartCastService_ReadCartCast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartCastServiceServer).ReadCartCast(ctx, req.(*CartCastId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartCastService_ReadByCoordsCartCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCastCoords)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartCastServiceServer).ReadByCoordsCartCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartCastService_ReadByCoordsCartCast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartCastServiceServer).ReadByCoordsCartCast(ctx, req.(*CartCastCoords))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartCastService_UpdateCartCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCast)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartCastServiceServer).UpdateCartCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartCastService_UpdateCartCast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartCastServiceServer).UpdateCartCast(ctx, req.(*CartCast))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartCastService_DeleteCartCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCastId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartCastServiceServer).DeleteCartCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartCastService_DeleteCartCast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartCastServiceServer).DeleteCartCast(ctx, req.(*CartCastId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartCastService_ListCartCast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CartCastServiceServer).ListCartCast(m, &grpc.GenericServerStream[emptypb.Empty, CartCast]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CartCastService_ListCartCastServer = grpc.ServerStreamingServer[CartCast]

// CartCastService_ServiceDesc is the grpc.ServiceDesc for CartCastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartCastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upselling.CartCastService",
	HandlerType: (*CartCastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCartCast",
			Handler:    _CartCastService_CreateCartCast_Handler,
		},
		{
			MethodName: "ReadCartCast",
			Handler:    _CartCastService_ReadCartCast_Handler,
		},
		{
			MethodName: "ReadByCoordsCartCast",
			Handler:    _CartCastService_ReadByCoordsCartCast_Handler,
		},
		{
			MethodName: "UpdateCartCast",
			Handler:    _CartCastService_UpdateCartCast_Handler,
		},
		{
			MethodName: "DeleteCartCast",
			Handler:    _CartCastService_DeleteCartCast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCartCast",
			Handler:       _CartCastService_ListCartCast_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cartcast.proto",
}