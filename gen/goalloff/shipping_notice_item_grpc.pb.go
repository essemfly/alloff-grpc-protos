// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: shipping_notice_item.proto

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

// ShippingNoticeItemControllerClient is the client API for ShippingNoticeItemController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingNoticeItemControllerClient interface {
	List(ctx context.Context, in *ShippingNoticeItemListRequest, opts ...grpc.CallOption) (ShippingNoticeItemController_ListClient, error)
	Create(ctx context.Context, in *ShippingNoticeItem, opts ...grpc.CallOption) (*ShippingNoticeItem, error)
	Retrieve(ctx context.Context, in *ShippingNoticeItemRetrieveRequest, opts ...grpc.CallOption) (*ShippingNoticeItem, error)
	Update(ctx context.Context, in *ShippingNoticeItem, opts ...grpc.CallOption) (*ShippingNoticeItem, error)
	Destroy(ctx context.Context, in *ShippingNoticeItem, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type shippingNoticeItemControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingNoticeItemControllerClient(cc grpc.ClientConnInterface) ShippingNoticeItemControllerClient {
	return &shippingNoticeItemControllerClient{cc}
}

func (c *shippingNoticeItemControllerClient) List(ctx context.Context, in *ShippingNoticeItemListRequest, opts ...grpc.CallOption) (ShippingNoticeItemController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShippingNoticeItemController_ServiceDesc.Streams[0], "/shipping_notice_item.ShippingNoticeItemController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &shippingNoticeItemControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShippingNoticeItemController_ListClient interface {
	Recv() (*ShippingNoticeItem, error)
	grpc.ClientStream
}

type shippingNoticeItemControllerListClient struct {
	grpc.ClientStream
}

func (x *shippingNoticeItemControllerListClient) Recv() (*ShippingNoticeItem, error) {
	m := new(ShippingNoticeItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shippingNoticeItemControllerClient) Create(ctx context.Context, in *ShippingNoticeItem, opts ...grpc.CallOption) (*ShippingNoticeItem, error) {
	out := new(ShippingNoticeItem)
	err := c.cc.Invoke(ctx, "/shipping_notice_item.ShippingNoticeItemController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingNoticeItemControllerClient) Retrieve(ctx context.Context, in *ShippingNoticeItemRetrieveRequest, opts ...grpc.CallOption) (*ShippingNoticeItem, error) {
	out := new(ShippingNoticeItem)
	err := c.cc.Invoke(ctx, "/shipping_notice_item.ShippingNoticeItemController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingNoticeItemControllerClient) Update(ctx context.Context, in *ShippingNoticeItem, opts ...grpc.CallOption) (*ShippingNoticeItem, error) {
	out := new(ShippingNoticeItem)
	err := c.cc.Invoke(ctx, "/shipping_notice_item.ShippingNoticeItemController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingNoticeItemControllerClient) Destroy(ctx context.Context, in *ShippingNoticeItem, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/shipping_notice_item.ShippingNoticeItemController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingNoticeItemControllerServer is the server API for ShippingNoticeItemController service.
// All implementations must embed UnimplementedShippingNoticeItemControllerServer
// for forward compatibility
type ShippingNoticeItemControllerServer interface {
	List(*ShippingNoticeItemListRequest, ShippingNoticeItemController_ListServer) error
	Create(context.Context, *ShippingNoticeItem) (*ShippingNoticeItem, error)
	Retrieve(context.Context, *ShippingNoticeItemRetrieveRequest) (*ShippingNoticeItem, error)
	Update(context.Context, *ShippingNoticeItem) (*ShippingNoticeItem, error)
	Destroy(context.Context, *ShippingNoticeItem) (*emptypb.Empty, error)
	mustEmbedUnimplementedShippingNoticeItemControllerServer()
}

// UnimplementedShippingNoticeItemControllerServer must be embedded to have forward compatible implementations.
type UnimplementedShippingNoticeItemControllerServer struct {
}

func (UnimplementedShippingNoticeItemControllerServer) List(*ShippingNoticeItemListRequest, ShippingNoticeItemController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedShippingNoticeItemControllerServer) Create(context.Context, *ShippingNoticeItem) (*ShippingNoticeItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShippingNoticeItemControllerServer) Retrieve(context.Context, *ShippingNoticeItemRetrieveRequest) (*ShippingNoticeItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedShippingNoticeItemControllerServer) Update(context.Context, *ShippingNoticeItem) (*ShippingNoticeItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedShippingNoticeItemControllerServer) Destroy(context.Context, *ShippingNoticeItem) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedShippingNoticeItemControllerServer) mustEmbedUnimplementedShippingNoticeItemControllerServer() {
}

// UnsafeShippingNoticeItemControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingNoticeItemControllerServer will
// result in compilation errors.
type UnsafeShippingNoticeItemControllerServer interface {
	mustEmbedUnimplementedShippingNoticeItemControllerServer()
}

func RegisterShippingNoticeItemControllerServer(s grpc.ServiceRegistrar, srv ShippingNoticeItemControllerServer) {
	s.RegisterService(&ShippingNoticeItemController_ServiceDesc, srv)
}

func _ShippingNoticeItemController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShippingNoticeItemListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShippingNoticeItemControllerServer).List(m, &shippingNoticeItemControllerListServer{stream})
}

type ShippingNoticeItemController_ListServer interface {
	Send(*ShippingNoticeItem) error
	grpc.ServerStream
}

type shippingNoticeItemControllerListServer struct {
	grpc.ServerStream
}

func (x *shippingNoticeItemControllerListServer) Send(m *ShippingNoticeItem) error {
	return x.ServerStream.SendMsg(m)
}

func _ShippingNoticeItemController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingNoticeItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingNoticeItemControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shipping_notice_item.ShippingNoticeItemController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingNoticeItemControllerServer).Create(ctx, req.(*ShippingNoticeItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingNoticeItemController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingNoticeItemRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingNoticeItemControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shipping_notice_item.ShippingNoticeItemController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingNoticeItemControllerServer).Retrieve(ctx, req.(*ShippingNoticeItemRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingNoticeItemController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingNoticeItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingNoticeItemControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shipping_notice_item.ShippingNoticeItemController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingNoticeItemControllerServer).Update(ctx, req.(*ShippingNoticeItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingNoticeItemController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShippingNoticeItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingNoticeItemControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shipping_notice_item.ShippingNoticeItemController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingNoticeItemControllerServer).Destroy(ctx, req.(*ShippingNoticeItem))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingNoticeItemController_ServiceDesc is the grpc.ServiceDesc for ShippingNoticeItemController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingNoticeItemController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shipping_notice_item.ShippingNoticeItemController",
	HandlerType: (*ShippingNoticeItemControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShippingNoticeItemController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _ShippingNoticeItemController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ShippingNoticeItemController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _ShippingNoticeItemController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ShippingNoticeItemController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shipping_notice_item.proto",
}
