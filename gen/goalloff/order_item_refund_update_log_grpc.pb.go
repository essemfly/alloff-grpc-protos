// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: order_item_refund_update_log.proto

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

// OrderItemRefundUpdateLogControllerClient is the client API for OrderItemRefundUpdateLogController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderItemRefundUpdateLogControllerClient interface {
	List(ctx context.Context, in *OrderItemRefundUpdateLogListRequest, opts ...grpc.CallOption) (OrderItemRefundUpdateLogController_ListClient, error)
	Create(ctx context.Context, in *OrderItemRefundUpdateLog, opts ...grpc.CallOption) (*OrderItemRefundUpdateLog, error)
	Retrieve(ctx context.Context, in *OrderItemRefundUpdateLogRetrieveRequest, opts ...grpc.CallOption) (*OrderItemRefundUpdateLog, error)
	Update(ctx context.Context, in *OrderItemRefundUpdateLog, opts ...grpc.CallOption) (*OrderItemRefundUpdateLog, error)
	Destroy(ctx context.Context, in *OrderItemRefundUpdateLog, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderItemRefundUpdateLogControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderItemRefundUpdateLogControllerClient(cc grpc.ClientConnInterface) OrderItemRefundUpdateLogControllerClient {
	return &orderItemRefundUpdateLogControllerClient{cc}
}

func (c *orderItemRefundUpdateLogControllerClient) List(ctx context.Context, in *OrderItemRefundUpdateLogListRequest, opts ...grpc.CallOption) (OrderItemRefundUpdateLogController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderItemRefundUpdateLogController_ServiceDesc.Streams[0], "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderItemRefundUpdateLogControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderItemRefundUpdateLogController_ListClient interface {
	Recv() (*OrderItemRefundUpdateLog, error)
	grpc.ClientStream
}

type orderItemRefundUpdateLogControllerListClient struct {
	grpc.ClientStream
}

func (x *orderItemRefundUpdateLogControllerListClient) Recv() (*OrderItemRefundUpdateLog, error) {
	m := new(OrderItemRefundUpdateLog)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderItemRefundUpdateLogControllerClient) Create(ctx context.Context, in *OrderItemRefundUpdateLog, opts ...grpc.CallOption) (*OrderItemRefundUpdateLog, error) {
	out := new(OrderItemRefundUpdateLog)
	err := c.cc.Invoke(ctx, "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemRefundUpdateLogControllerClient) Retrieve(ctx context.Context, in *OrderItemRefundUpdateLogRetrieveRequest, opts ...grpc.CallOption) (*OrderItemRefundUpdateLog, error) {
	out := new(OrderItemRefundUpdateLog)
	err := c.cc.Invoke(ctx, "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemRefundUpdateLogControllerClient) Update(ctx context.Context, in *OrderItemRefundUpdateLog, opts ...grpc.CallOption) (*OrderItemRefundUpdateLog, error) {
	out := new(OrderItemRefundUpdateLog)
	err := c.cc.Invoke(ctx, "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemRefundUpdateLogControllerClient) Destroy(ctx context.Context, in *OrderItemRefundUpdateLog, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderItemRefundUpdateLogControllerServer is the server API for OrderItemRefundUpdateLogController service.
// All implementations must embed UnimplementedOrderItemRefundUpdateLogControllerServer
// for forward compatibility
type OrderItemRefundUpdateLogControllerServer interface {
	List(*OrderItemRefundUpdateLogListRequest, OrderItemRefundUpdateLogController_ListServer) error
	Create(context.Context, *OrderItemRefundUpdateLog) (*OrderItemRefundUpdateLog, error)
	Retrieve(context.Context, *OrderItemRefundUpdateLogRetrieveRequest) (*OrderItemRefundUpdateLog, error)
	Update(context.Context, *OrderItemRefundUpdateLog) (*OrderItemRefundUpdateLog, error)
	Destroy(context.Context, *OrderItemRefundUpdateLog) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderItemRefundUpdateLogControllerServer()
}

// UnimplementedOrderItemRefundUpdateLogControllerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderItemRefundUpdateLogControllerServer struct {
}

func (UnimplementedOrderItemRefundUpdateLogControllerServer) List(*OrderItemRefundUpdateLogListRequest, OrderItemRefundUpdateLogController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderItemRefundUpdateLogControllerServer) Create(context.Context, *OrderItemRefundUpdateLog) (*OrderItemRefundUpdateLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderItemRefundUpdateLogControllerServer) Retrieve(context.Context, *OrderItemRefundUpdateLogRetrieveRequest) (*OrderItemRefundUpdateLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedOrderItemRefundUpdateLogControllerServer) Update(context.Context, *OrderItemRefundUpdateLog) (*OrderItemRefundUpdateLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderItemRefundUpdateLogControllerServer) Destroy(context.Context, *OrderItemRefundUpdateLog) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedOrderItemRefundUpdateLogControllerServer) mustEmbedUnimplementedOrderItemRefundUpdateLogControllerServer() {
}

// UnsafeOrderItemRefundUpdateLogControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderItemRefundUpdateLogControllerServer will
// result in compilation errors.
type UnsafeOrderItemRefundUpdateLogControllerServer interface {
	mustEmbedUnimplementedOrderItemRefundUpdateLogControllerServer()
}

func RegisterOrderItemRefundUpdateLogControllerServer(s grpc.ServiceRegistrar, srv OrderItemRefundUpdateLogControllerServer) {
	s.RegisterService(&OrderItemRefundUpdateLogController_ServiceDesc, srv)
}

func _OrderItemRefundUpdateLogController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderItemRefundUpdateLogListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderItemRefundUpdateLogControllerServer).List(m, &orderItemRefundUpdateLogControllerListServer{stream})
}

type OrderItemRefundUpdateLogController_ListServer interface {
	Send(*OrderItemRefundUpdateLog) error
	grpc.ServerStream
}

type orderItemRefundUpdateLogControllerListServer struct {
	grpc.ServerStream
}

func (x *orderItemRefundUpdateLogControllerListServer) Send(m *OrderItemRefundUpdateLog) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderItemRefundUpdateLogController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemRefundUpdateLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemRefundUpdateLogControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemRefundUpdateLogControllerServer).Create(ctx, req.(*OrderItemRefundUpdateLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemRefundUpdateLogController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemRefundUpdateLogRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemRefundUpdateLogControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemRefundUpdateLogControllerServer).Retrieve(ctx, req.(*OrderItemRefundUpdateLogRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemRefundUpdateLogController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemRefundUpdateLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemRefundUpdateLogControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemRefundUpdateLogControllerServer).Update(ctx, req.(*OrderItemRefundUpdateLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemRefundUpdateLogController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemRefundUpdateLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemRefundUpdateLogControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitemrefundupdatelog.OrderItemRefundUpdateLogController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemRefundUpdateLogControllerServer).Destroy(ctx, req.(*OrderItemRefundUpdateLog))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderItemRefundUpdateLogController_ServiceDesc is the grpc.ServiceDesc for OrderItemRefundUpdateLogController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderItemRefundUpdateLogController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderitemrefundupdatelog.OrderItemRefundUpdateLogController",
	HandlerType: (*OrderItemRefundUpdateLogControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderItemRefundUpdateLogController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _OrderItemRefundUpdateLogController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderItemRefundUpdateLogController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _OrderItemRefundUpdateLogController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _OrderItemRefundUpdateLogController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order_item_refund_update_log.proto",
}
