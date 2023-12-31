// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: inventory_receipt_log.proto

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

// InventoryReceiptLogControllerClient is the client API for InventoryReceiptLogController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryReceiptLogControllerClient interface {
	List(ctx context.Context, in *InventoryReceiptLogListRequest, opts ...grpc.CallOption) (InventoryReceiptLogController_ListClient, error)
	Create(ctx context.Context, in *InventoryReceiptLog, opts ...grpc.CallOption) (*InventoryReceiptLog, error)
	Retrieve(ctx context.Context, in *InventoryReceiptLogRetrieveRequest, opts ...grpc.CallOption) (*InventoryReceiptLog, error)
	Update(ctx context.Context, in *InventoryReceiptLog, opts ...grpc.CallOption) (*InventoryReceiptLog, error)
	Destroy(ctx context.Context, in *InventoryReceiptLog, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type inventoryReceiptLogControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryReceiptLogControllerClient(cc grpc.ClientConnInterface) InventoryReceiptLogControllerClient {
	return &inventoryReceiptLogControllerClient{cc}
}

func (c *inventoryReceiptLogControllerClient) List(ctx context.Context, in *InventoryReceiptLogListRequest, opts ...grpc.CallOption) (InventoryReceiptLogController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &InventoryReceiptLogController_ServiceDesc.Streams[0], "/inventoryreceiptlog.InventoryReceiptLogController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &inventoryReceiptLogControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InventoryReceiptLogController_ListClient interface {
	Recv() (*InventoryReceiptLog, error)
	grpc.ClientStream
}

type inventoryReceiptLogControllerListClient struct {
	grpc.ClientStream
}

func (x *inventoryReceiptLogControllerListClient) Recv() (*InventoryReceiptLog, error) {
	m := new(InventoryReceiptLog)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *inventoryReceiptLogControllerClient) Create(ctx context.Context, in *InventoryReceiptLog, opts ...grpc.CallOption) (*InventoryReceiptLog, error) {
	out := new(InventoryReceiptLog)
	err := c.cc.Invoke(ctx, "/inventoryreceiptlog.InventoryReceiptLogController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryReceiptLogControllerClient) Retrieve(ctx context.Context, in *InventoryReceiptLogRetrieveRequest, opts ...grpc.CallOption) (*InventoryReceiptLog, error) {
	out := new(InventoryReceiptLog)
	err := c.cc.Invoke(ctx, "/inventoryreceiptlog.InventoryReceiptLogController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryReceiptLogControllerClient) Update(ctx context.Context, in *InventoryReceiptLog, opts ...grpc.CallOption) (*InventoryReceiptLog, error) {
	out := new(InventoryReceiptLog)
	err := c.cc.Invoke(ctx, "/inventoryreceiptlog.InventoryReceiptLogController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryReceiptLogControllerClient) Destroy(ctx context.Context, in *InventoryReceiptLog, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/inventoryreceiptlog.InventoryReceiptLogController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryReceiptLogControllerServer is the server API for InventoryReceiptLogController service.
// All implementations must embed UnimplementedInventoryReceiptLogControllerServer
// for forward compatibility
type InventoryReceiptLogControllerServer interface {
	List(*InventoryReceiptLogListRequest, InventoryReceiptLogController_ListServer) error
	Create(context.Context, *InventoryReceiptLog) (*InventoryReceiptLog, error)
	Retrieve(context.Context, *InventoryReceiptLogRetrieveRequest) (*InventoryReceiptLog, error)
	Update(context.Context, *InventoryReceiptLog) (*InventoryReceiptLog, error)
	Destroy(context.Context, *InventoryReceiptLog) (*emptypb.Empty, error)
	mustEmbedUnimplementedInventoryReceiptLogControllerServer()
}

// UnimplementedInventoryReceiptLogControllerServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryReceiptLogControllerServer struct {
}

func (UnimplementedInventoryReceiptLogControllerServer) List(*InventoryReceiptLogListRequest, InventoryReceiptLogController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedInventoryReceiptLogControllerServer) Create(context.Context, *InventoryReceiptLog) (*InventoryReceiptLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInventoryReceiptLogControllerServer) Retrieve(context.Context, *InventoryReceiptLogRetrieveRequest) (*InventoryReceiptLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedInventoryReceiptLogControllerServer) Update(context.Context, *InventoryReceiptLog) (*InventoryReceiptLog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedInventoryReceiptLogControllerServer) Destroy(context.Context, *InventoryReceiptLog) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedInventoryReceiptLogControllerServer) mustEmbedUnimplementedInventoryReceiptLogControllerServer() {
}

// UnsafeInventoryReceiptLogControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryReceiptLogControllerServer will
// result in compilation errors.
type UnsafeInventoryReceiptLogControllerServer interface {
	mustEmbedUnimplementedInventoryReceiptLogControllerServer()
}

func RegisterInventoryReceiptLogControllerServer(s grpc.ServiceRegistrar, srv InventoryReceiptLogControllerServer) {
	s.RegisterService(&InventoryReceiptLogController_ServiceDesc, srv)
}

func _InventoryReceiptLogController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InventoryReceiptLogListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InventoryReceiptLogControllerServer).List(m, &inventoryReceiptLogControllerListServer{stream})
}

type InventoryReceiptLogController_ListServer interface {
	Send(*InventoryReceiptLog) error
	grpc.ServerStream
}

type inventoryReceiptLogControllerListServer struct {
	grpc.ServerStream
}

func (x *inventoryReceiptLogControllerListServer) Send(m *InventoryReceiptLog) error {
	return x.ServerStream.SendMsg(m)
}

func _InventoryReceiptLogController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryReceiptLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryReceiptLogControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryreceiptlog.InventoryReceiptLogController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryReceiptLogControllerServer).Create(ctx, req.(*InventoryReceiptLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryReceiptLogController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryReceiptLogRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryReceiptLogControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryreceiptlog.InventoryReceiptLogController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryReceiptLogControllerServer).Retrieve(ctx, req.(*InventoryReceiptLogRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryReceiptLogController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryReceiptLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryReceiptLogControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryreceiptlog.InventoryReceiptLogController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryReceiptLogControllerServer).Update(ctx, req.(*InventoryReceiptLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryReceiptLogController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryReceiptLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryReceiptLogControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventoryreceiptlog.InventoryReceiptLogController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryReceiptLogControllerServer).Destroy(ctx, req.(*InventoryReceiptLog))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryReceiptLogController_ServiceDesc is the grpc.ServiceDesc for InventoryReceiptLogController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryReceiptLogController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventoryreceiptlog.InventoryReceiptLogController",
	HandlerType: (*InventoryReceiptLogControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _InventoryReceiptLogController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _InventoryReceiptLogController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _InventoryReceiptLogController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _InventoryReceiptLogController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _InventoryReceiptLogController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "inventory_receipt_log.proto",
}
