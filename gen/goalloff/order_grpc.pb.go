// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: order.proto

package goalloff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderControllerClient is the client API for OrderController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderControllerClient interface {
	List(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (OrderController_ListClient, error)
	Create(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	Retrieve(ctx context.Context, in *OrderRetrieveRequest, opts ...grpc.CallOption) (*Order, error)
	Update(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	Destroy(ctx context.Context, in *Order, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderControllerClient(cc grpc.ClientConnInterface) OrderControllerClient {
	return &orderControllerClient{cc}
}

func (c *orderControllerClient) List(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (OrderController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderController_ServiceDesc.Streams[0], "/order.OrderController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderController_ListClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderControllerListClient struct {
	grpc.ClientStream
}

func (x *orderControllerListClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderControllerClient) Create(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderControllerClient) Retrieve(ctx context.Context, in *OrderRetrieveRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderControllerClient) Update(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/order.OrderController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderControllerClient) Destroy(ctx context.Context, in *Order, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order.OrderController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderControllerServer is the server API for OrderController service.
// All implementations must embed UnimplementedOrderControllerServer
// for forward compatibility
type OrderControllerServer interface {
	List(*OrderListRequest, OrderController_ListServer) error
	Create(context.Context, *Order) (*Order, error)
	Retrieve(context.Context, *OrderRetrieveRequest) (*Order, error)
	Update(context.Context, *Order) (*Order, error)
	Destroy(context.Context, *Order) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderControllerServer()
}

// UnimplementedOrderControllerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderControllerServer struct {
}

func (UnimplementedOrderControllerServer) List(*OrderListRequest, OrderController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderControllerServer) Create(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderControllerServer) Retrieve(context.Context, *OrderRetrieveRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedOrderControllerServer) Update(context.Context, *Order) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderControllerServer) Destroy(context.Context, *Order) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedOrderControllerServer) mustEmbedUnimplementedOrderControllerServer() {}

// UnsafeOrderControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderControllerServer will
// result in compilation errors.
type UnsafeOrderControllerServer interface {
	mustEmbedUnimplementedOrderControllerServer()
}

func RegisterOrderControllerServer(s grpc.ServiceRegistrar, srv OrderControllerServer) {
	s.RegisterService(&OrderController_ServiceDesc, srv)
}

func _OrderController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderControllerServer).List(m, &orderControllerListServer{stream})
}

type OrderController_ListServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderControllerListServer struct {
	grpc.ServerStream
}

func (x *orderControllerListServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderControllerServer).Create(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderControllerServer).Retrieve(ctx, req.(*OrderRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderControllerServer).Update(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderControllerServer).Destroy(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderController_ServiceDesc is the grpc.ServiceDesc for OrderController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderController",
	HandlerType: (*OrderControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _OrderController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _OrderController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _OrderController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order.proto",
}
