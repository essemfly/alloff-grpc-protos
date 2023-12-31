// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: refund_item_history.proto

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

// RefundItemHistoryControllerClient is the client API for RefundItemHistoryController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RefundItemHistoryControllerClient interface {
	List(ctx context.Context, in *RefundItemHistoryListRequest, opts ...grpc.CallOption) (RefundItemHistoryController_ListClient, error)
	Create(ctx context.Context, in *RefundItemHistory, opts ...grpc.CallOption) (*RefundItemHistory, error)
	Retrieve(ctx context.Context, in *RefundItemHistoryRetrieveRequest, opts ...grpc.CallOption) (*RefundItemHistory, error)
	Update(ctx context.Context, in *RefundItemHistory, opts ...grpc.CallOption) (*RefundItemHistory, error)
	Destroy(ctx context.Context, in *RefundItemHistory, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type refundItemHistoryControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewRefundItemHistoryControllerClient(cc grpc.ClientConnInterface) RefundItemHistoryControllerClient {
	return &refundItemHistoryControllerClient{cc}
}

func (c *refundItemHistoryControllerClient) List(ctx context.Context, in *RefundItemHistoryListRequest, opts ...grpc.CallOption) (RefundItemHistoryController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &RefundItemHistoryController_ServiceDesc.Streams[0], "/refunditemhistory.RefundItemHistoryController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &refundItemHistoryControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RefundItemHistoryController_ListClient interface {
	Recv() (*RefundItemHistory, error)
	grpc.ClientStream
}

type refundItemHistoryControllerListClient struct {
	grpc.ClientStream
}

func (x *refundItemHistoryControllerListClient) Recv() (*RefundItemHistory, error) {
	m := new(RefundItemHistory)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *refundItemHistoryControllerClient) Create(ctx context.Context, in *RefundItemHistory, opts ...grpc.CallOption) (*RefundItemHistory, error) {
	out := new(RefundItemHistory)
	err := c.cc.Invoke(ctx, "/refunditemhistory.RefundItemHistoryController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundItemHistoryControllerClient) Retrieve(ctx context.Context, in *RefundItemHistoryRetrieveRequest, opts ...grpc.CallOption) (*RefundItemHistory, error) {
	out := new(RefundItemHistory)
	err := c.cc.Invoke(ctx, "/refunditemhistory.RefundItemHistoryController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundItemHistoryControllerClient) Update(ctx context.Context, in *RefundItemHistory, opts ...grpc.CallOption) (*RefundItemHistory, error) {
	out := new(RefundItemHistory)
	err := c.cc.Invoke(ctx, "/refunditemhistory.RefundItemHistoryController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *refundItemHistoryControllerClient) Destroy(ctx context.Context, in *RefundItemHistory, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/refunditemhistory.RefundItemHistoryController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RefundItemHistoryControllerServer is the server API for RefundItemHistoryController service.
// All implementations must embed UnimplementedRefundItemHistoryControllerServer
// for forward compatibility
type RefundItemHistoryControllerServer interface {
	List(*RefundItemHistoryListRequest, RefundItemHistoryController_ListServer) error
	Create(context.Context, *RefundItemHistory) (*RefundItemHistory, error)
	Retrieve(context.Context, *RefundItemHistoryRetrieveRequest) (*RefundItemHistory, error)
	Update(context.Context, *RefundItemHistory) (*RefundItemHistory, error)
	Destroy(context.Context, *RefundItemHistory) (*emptypb.Empty, error)
	mustEmbedUnimplementedRefundItemHistoryControllerServer()
}

// UnimplementedRefundItemHistoryControllerServer must be embedded to have forward compatible implementations.
type UnimplementedRefundItemHistoryControllerServer struct {
}

func (UnimplementedRefundItemHistoryControllerServer) List(*RefundItemHistoryListRequest, RefundItemHistoryController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRefundItemHistoryControllerServer) Create(context.Context, *RefundItemHistory) (*RefundItemHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRefundItemHistoryControllerServer) Retrieve(context.Context, *RefundItemHistoryRetrieveRequest) (*RefundItemHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedRefundItemHistoryControllerServer) Update(context.Context, *RefundItemHistory) (*RefundItemHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRefundItemHistoryControllerServer) Destroy(context.Context, *RefundItemHistory) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedRefundItemHistoryControllerServer) mustEmbedUnimplementedRefundItemHistoryControllerServer() {
}

// UnsafeRefundItemHistoryControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RefundItemHistoryControllerServer will
// result in compilation errors.
type UnsafeRefundItemHistoryControllerServer interface {
	mustEmbedUnimplementedRefundItemHistoryControllerServer()
}

func RegisterRefundItemHistoryControllerServer(s grpc.ServiceRegistrar, srv RefundItemHistoryControllerServer) {
	s.RegisterService(&RefundItemHistoryController_ServiceDesc, srv)
}

func _RefundItemHistoryController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RefundItemHistoryListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RefundItemHistoryControllerServer).List(m, &refundItemHistoryControllerListServer{stream})
}

type RefundItemHistoryController_ListServer interface {
	Send(*RefundItemHistory) error
	grpc.ServerStream
}

type refundItemHistoryControllerListServer struct {
	grpc.ServerStream
}

func (x *refundItemHistoryControllerListServer) Send(m *RefundItemHistory) error {
	return x.ServerStream.SendMsg(m)
}

func _RefundItemHistoryController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundItemHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefundItemHistoryControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/refunditemhistory.RefundItemHistoryController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefundItemHistoryControllerServer).Create(ctx, req.(*RefundItemHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RefundItemHistoryController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundItemHistoryRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefundItemHistoryControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/refunditemhistory.RefundItemHistoryController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefundItemHistoryControllerServer).Retrieve(ctx, req.(*RefundItemHistoryRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RefundItemHistoryController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundItemHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefundItemHistoryControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/refunditemhistory.RefundItemHistoryController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefundItemHistoryControllerServer).Update(ctx, req.(*RefundItemHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _RefundItemHistoryController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundItemHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefundItemHistoryControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/refunditemhistory.RefundItemHistoryController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefundItemHistoryControllerServer).Destroy(ctx, req.(*RefundItemHistory))
	}
	return interceptor(ctx, in, info, handler)
}

// RefundItemHistoryController_ServiceDesc is the grpc.ServiceDesc for RefundItemHistoryController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RefundItemHistoryController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "refunditemhistory.RefundItemHistoryController",
	HandlerType: (*RefundItemHistoryControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RefundItemHistoryController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _RefundItemHistoryController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RefundItemHistoryController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _RefundItemHistoryController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _RefundItemHistoryController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "refund_item_history.proto",
}
