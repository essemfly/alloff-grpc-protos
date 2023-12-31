// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: order_item_memo.proto

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

// OrderItemMemoControllerClient is the client API for OrderItemMemoController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderItemMemoControllerClient interface {
	List(ctx context.Context, in *OrderItemMemoListRequest, opts ...grpc.CallOption) (OrderItemMemoController_ListClient, error)
	Create(ctx context.Context, in *OrderItemMemo, opts ...grpc.CallOption) (*OrderItemMemo, error)
	Retrieve(ctx context.Context, in *OrderItemMemoRetrieveRequest, opts ...grpc.CallOption) (*OrderItemMemo, error)
	Update(ctx context.Context, in *OrderItemMemo, opts ...grpc.CallOption) (*OrderItemMemo, error)
	Destroy(ctx context.Context, in *OrderItemMemo, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderItemMemoControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderItemMemoControllerClient(cc grpc.ClientConnInterface) OrderItemMemoControllerClient {
	return &orderItemMemoControllerClient{cc}
}

func (c *orderItemMemoControllerClient) List(ctx context.Context, in *OrderItemMemoListRequest, opts ...grpc.CallOption) (OrderItemMemoController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderItemMemoController_ServiceDesc.Streams[0], "/orderitemmemo.OrderItemMemoController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderItemMemoControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderItemMemoController_ListClient interface {
	Recv() (*OrderItemMemo, error)
	grpc.ClientStream
}

type orderItemMemoControllerListClient struct {
	grpc.ClientStream
}

func (x *orderItemMemoControllerListClient) Recv() (*OrderItemMemo, error) {
	m := new(OrderItemMemo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderItemMemoControllerClient) Create(ctx context.Context, in *OrderItemMemo, opts ...grpc.CallOption) (*OrderItemMemo, error) {
	out := new(OrderItemMemo)
	err := c.cc.Invoke(ctx, "/orderitemmemo.OrderItemMemoController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemMemoControllerClient) Retrieve(ctx context.Context, in *OrderItemMemoRetrieveRequest, opts ...grpc.CallOption) (*OrderItemMemo, error) {
	out := new(OrderItemMemo)
	err := c.cc.Invoke(ctx, "/orderitemmemo.OrderItemMemoController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemMemoControllerClient) Update(ctx context.Context, in *OrderItemMemo, opts ...grpc.CallOption) (*OrderItemMemo, error) {
	out := new(OrderItemMemo)
	err := c.cc.Invoke(ctx, "/orderitemmemo.OrderItemMemoController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemMemoControllerClient) Destroy(ctx context.Context, in *OrderItemMemo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/orderitemmemo.OrderItemMemoController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderItemMemoControllerServer is the server API for OrderItemMemoController service.
// All implementations must embed UnimplementedOrderItemMemoControllerServer
// for forward compatibility
type OrderItemMemoControllerServer interface {
	List(*OrderItemMemoListRequest, OrderItemMemoController_ListServer) error
	Create(context.Context, *OrderItemMemo) (*OrderItemMemo, error)
	Retrieve(context.Context, *OrderItemMemoRetrieveRequest) (*OrderItemMemo, error)
	Update(context.Context, *OrderItemMemo) (*OrderItemMemo, error)
	Destroy(context.Context, *OrderItemMemo) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderItemMemoControllerServer()
}

// UnimplementedOrderItemMemoControllerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderItemMemoControllerServer struct {
}

func (UnimplementedOrderItemMemoControllerServer) List(*OrderItemMemoListRequest, OrderItemMemoController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderItemMemoControllerServer) Create(context.Context, *OrderItemMemo) (*OrderItemMemo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderItemMemoControllerServer) Retrieve(context.Context, *OrderItemMemoRetrieveRequest) (*OrderItemMemo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedOrderItemMemoControllerServer) Update(context.Context, *OrderItemMemo) (*OrderItemMemo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderItemMemoControllerServer) Destroy(context.Context, *OrderItemMemo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedOrderItemMemoControllerServer) mustEmbedUnimplementedOrderItemMemoControllerServer() {
}

// UnsafeOrderItemMemoControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderItemMemoControllerServer will
// result in compilation errors.
type UnsafeOrderItemMemoControllerServer interface {
	mustEmbedUnimplementedOrderItemMemoControllerServer()
}

func RegisterOrderItemMemoControllerServer(s grpc.ServiceRegistrar, srv OrderItemMemoControllerServer) {
	s.RegisterService(&OrderItemMemoController_ServiceDesc, srv)
}

func _OrderItemMemoController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderItemMemoListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderItemMemoControllerServer).List(m, &orderItemMemoControllerListServer{stream})
}

type OrderItemMemoController_ListServer interface {
	Send(*OrderItemMemo) error
	grpc.ServerStream
}

type orderItemMemoControllerListServer struct {
	grpc.ServerStream
}

func (x *orderItemMemoControllerListServer) Send(m *OrderItemMemo) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderItemMemoController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemMemo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemMemoControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemmemo.OrderItemMemoController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemMemoControllerServer).Create(ctx, req.(*OrderItemMemo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemMemoController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemMemoRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemMemoControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemmemo.OrderItemMemoController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemMemoControllerServer).Retrieve(ctx, req.(*OrderItemMemoRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemMemoController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemMemo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemMemoControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemmemo.OrderItemMemoController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemMemoControllerServer).Update(ctx, req.(*OrderItemMemo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemMemoController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemMemo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemMemoControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemmemo.OrderItemMemoController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemMemoControllerServer).Destroy(ctx, req.(*OrderItemMemo))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderItemMemoController_ServiceDesc is the grpc.ServiceDesc for OrderItemMemoController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderItemMemoController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderitemmemo.OrderItemMemoController",
	HandlerType: (*OrderItemMemoControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderItemMemoController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _OrderItemMemoController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderItemMemoController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _OrderItemMemoController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _OrderItemMemoController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order_item_memo.proto",
}
