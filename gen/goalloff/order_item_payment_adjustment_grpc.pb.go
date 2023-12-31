// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: order_item_payment_adjustment.proto

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

// OrderItemPaymentAdjustmentControllerClient is the client API for OrderItemPaymentAdjustmentController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderItemPaymentAdjustmentControllerClient interface {
	List(ctx context.Context, in *OrderItemPaymentAdjustmentListRequest, opts ...grpc.CallOption) (OrderItemPaymentAdjustmentController_ListClient, error)
	Create(ctx context.Context, in *OrderItemPaymentAdjustment, opts ...grpc.CallOption) (*OrderItemPaymentAdjustment, error)
	Retrieve(ctx context.Context, in *OrderItemPaymentAdjustmentRetrieveRequest, opts ...grpc.CallOption) (*OrderItemPaymentAdjustment, error)
	Update(ctx context.Context, in *OrderItemPaymentAdjustment, opts ...grpc.CallOption) (*OrderItemPaymentAdjustment, error)
	Destroy(ctx context.Context, in *OrderItemPaymentAdjustment, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderItemPaymentAdjustmentControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderItemPaymentAdjustmentControllerClient(cc grpc.ClientConnInterface) OrderItemPaymentAdjustmentControllerClient {
	return &orderItemPaymentAdjustmentControllerClient{cc}
}

func (c *orderItemPaymentAdjustmentControllerClient) List(ctx context.Context, in *OrderItemPaymentAdjustmentListRequest, opts ...grpc.CallOption) (OrderItemPaymentAdjustmentController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderItemPaymentAdjustmentController_ServiceDesc.Streams[0], "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderItemPaymentAdjustmentControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderItemPaymentAdjustmentController_ListClient interface {
	Recv() (*OrderItemPaymentAdjustment, error)
	grpc.ClientStream
}

type orderItemPaymentAdjustmentControllerListClient struct {
	grpc.ClientStream
}

func (x *orderItemPaymentAdjustmentControllerListClient) Recv() (*OrderItemPaymentAdjustment, error) {
	m := new(OrderItemPaymentAdjustment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderItemPaymentAdjustmentControllerClient) Create(ctx context.Context, in *OrderItemPaymentAdjustment, opts ...grpc.CallOption) (*OrderItemPaymentAdjustment, error) {
	out := new(OrderItemPaymentAdjustment)
	err := c.cc.Invoke(ctx, "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemPaymentAdjustmentControllerClient) Retrieve(ctx context.Context, in *OrderItemPaymentAdjustmentRetrieveRequest, opts ...grpc.CallOption) (*OrderItemPaymentAdjustment, error) {
	out := new(OrderItemPaymentAdjustment)
	err := c.cc.Invoke(ctx, "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemPaymentAdjustmentControllerClient) Update(ctx context.Context, in *OrderItemPaymentAdjustment, opts ...grpc.CallOption) (*OrderItemPaymentAdjustment, error) {
	out := new(OrderItemPaymentAdjustment)
	err := c.cc.Invoke(ctx, "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderItemPaymentAdjustmentControllerClient) Destroy(ctx context.Context, in *OrderItemPaymentAdjustment, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderItemPaymentAdjustmentControllerServer is the server API for OrderItemPaymentAdjustmentController service.
// All implementations must embed UnimplementedOrderItemPaymentAdjustmentControllerServer
// for forward compatibility
type OrderItemPaymentAdjustmentControllerServer interface {
	List(*OrderItemPaymentAdjustmentListRequest, OrderItemPaymentAdjustmentController_ListServer) error
	Create(context.Context, *OrderItemPaymentAdjustment) (*OrderItemPaymentAdjustment, error)
	Retrieve(context.Context, *OrderItemPaymentAdjustmentRetrieveRequest) (*OrderItemPaymentAdjustment, error)
	Update(context.Context, *OrderItemPaymentAdjustment) (*OrderItemPaymentAdjustment, error)
	Destroy(context.Context, *OrderItemPaymentAdjustment) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderItemPaymentAdjustmentControllerServer()
}

// UnimplementedOrderItemPaymentAdjustmentControllerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderItemPaymentAdjustmentControllerServer struct {
}

func (UnimplementedOrderItemPaymentAdjustmentControllerServer) List(*OrderItemPaymentAdjustmentListRequest, OrderItemPaymentAdjustmentController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderItemPaymentAdjustmentControllerServer) Create(context.Context, *OrderItemPaymentAdjustment) (*OrderItemPaymentAdjustment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderItemPaymentAdjustmentControllerServer) Retrieve(context.Context, *OrderItemPaymentAdjustmentRetrieveRequest) (*OrderItemPaymentAdjustment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedOrderItemPaymentAdjustmentControllerServer) Update(context.Context, *OrderItemPaymentAdjustment) (*OrderItemPaymentAdjustment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderItemPaymentAdjustmentControllerServer) Destroy(context.Context, *OrderItemPaymentAdjustment) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedOrderItemPaymentAdjustmentControllerServer) mustEmbedUnimplementedOrderItemPaymentAdjustmentControllerServer() {
}

// UnsafeOrderItemPaymentAdjustmentControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderItemPaymentAdjustmentControllerServer will
// result in compilation errors.
type UnsafeOrderItemPaymentAdjustmentControllerServer interface {
	mustEmbedUnimplementedOrderItemPaymentAdjustmentControllerServer()
}

func RegisterOrderItemPaymentAdjustmentControllerServer(s grpc.ServiceRegistrar, srv OrderItemPaymentAdjustmentControllerServer) {
	s.RegisterService(&OrderItemPaymentAdjustmentController_ServiceDesc, srv)
}

func _OrderItemPaymentAdjustmentController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderItemPaymentAdjustmentListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderItemPaymentAdjustmentControllerServer).List(m, &orderItemPaymentAdjustmentControllerListServer{stream})
}

type OrderItemPaymentAdjustmentController_ListServer interface {
	Send(*OrderItemPaymentAdjustment) error
	grpc.ServerStream
}

type orderItemPaymentAdjustmentControllerListServer struct {
	grpc.ServerStream
}

func (x *orderItemPaymentAdjustmentControllerListServer) Send(m *OrderItemPaymentAdjustment) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderItemPaymentAdjustmentController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemPaymentAdjustment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Create(ctx, req.(*OrderItemPaymentAdjustment))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemPaymentAdjustmentController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemPaymentAdjustmentRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Retrieve(ctx, req.(*OrderItemPaymentAdjustmentRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemPaymentAdjustmentController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemPaymentAdjustment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Update(ctx, req.(*OrderItemPaymentAdjustment))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderItemPaymentAdjustmentController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemPaymentAdjustment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderitempaymentadjustment.OrderItemPaymentAdjustmentController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderItemPaymentAdjustmentControllerServer).Destroy(ctx, req.(*OrderItemPaymentAdjustment))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderItemPaymentAdjustmentController_ServiceDesc is the grpc.ServiceDesc for OrderItemPaymentAdjustmentController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderItemPaymentAdjustmentController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderitempaymentadjustment.OrderItemPaymentAdjustmentController",
	HandlerType: (*OrderItemPaymentAdjustmentControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderItemPaymentAdjustmentController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _OrderItemPaymentAdjustmentController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderItemPaymentAdjustmentController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _OrderItemPaymentAdjustmentController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _OrderItemPaymentAdjustmentController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order_item_payment_adjustment.proto",
}
