// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// SupplierControllerClient is the client API for SupplierController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierControllerClient interface {
	List(ctx context.Context, in *SupplierListRequest, opts ...grpc.CallOption) (SupplierController_ListClient, error)
	Create(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error)
	Retrieve(ctx context.Context, in *SupplierRetrieveRequest, opts ...grpc.CallOption) (*Supplier, error)
	Update(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error)
	Destroy(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type supplierControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierControllerClient(cc grpc.ClientConnInterface) SupplierControllerClient {
	return &supplierControllerClient{cc}
}

func (c *supplierControllerClient) List(ctx context.Context, in *SupplierListRequest, opts ...grpc.CallOption) (SupplierController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &SupplierController_ServiceDesc.Streams[0], "/supplier.SupplierController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &supplierControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SupplierController_ListClient interface {
	Recv() (*Supplier, error)
	grpc.ClientStream
}

type supplierControllerListClient struct {
	grpc.ClientStream
}

func (x *supplierControllerListClient) Recv() (*Supplier, error) {
	m := new(Supplier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *supplierControllerClient) Create(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/supplier.SupplierController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierControllerClient) Retrieve(ctx context.Context, in *SupplierRetrieveRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/supplier.SupplierController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierControllerClient) Update(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/supplier.SupplierController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierControllerClient) Destroy(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/supplier.SupplierController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupplierControllerServer is the server API for SupplierController service.
// All implementations must embed UnimplementedSupplierControllerServer
// for forward compatibility
type SupplierControllerServer interface {
	List(*SupplierListRequest, SupplierController_ListServer) error
	Create(context.Context, *Supplier) (*Supplier, error)
	Retrieve(context.Context, *SupplierRetrieveRequest) (*Supplier, error)
	Update(context.Context, *Supplier) (*Supplier, error)
	Destroy(context.Context, *Supplier) (*emptypb.Empty, error)
	mustEmbedUnimplementedSupplierControllerServer()
}

// UnimplementedSupplierControllerServer must be embedded to have forward compatible implementations.
type UnimplementedSupplierControllerServer struct {
}

func (UnimplementedSupplierControllerServer) List(*SupplierListRequest, SupplierController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSupplierControllerServer) Create(context.Context, *Supplier) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSupplierControllerServer) Retrieve(context.Context, *SupplierRetrieveRequest) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedSupplierControllerServer) Update(context.Context, *Supplier) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSupplierControllerServer) Destroy(context.Context, *Supplier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedSupplierControllerServer) mustEmbedUnimplementedSupplierControllerServer() {}

// UnsafeSupplierControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierControllerServer will
// result in compilation errors.
type UnsafeSupplierControllerServer interface {
	mustEmbedUnimplementedSupplierControllerServer()
}

func RegisterSupplierControllerServer(s grpc.ServiceRegistrar, srv SupplierControllerServer) {
	s.RegisterService(&SupplierController_ServiceDesc, srv)
}

func _SupplierController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SupplierListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SupplierControllerServer).List(m, &supplierControllerListServer{stream})
}

type SupplierController_ListServer interface {
	Send(*Supplier) error
	grpc.ServerStream
}

type supplierControllerListServer struct {
	grpc.ServerStream
}

func (x *supplierControllerListServer) Send(m *Supplier) error {
	return x.ServerStream.SendMsg(m)
}

func _SupplierController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierControllerServer).Create(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplierRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierControllerServer).Retrieve(ctx, req.(*SupplierRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierControllerServer).Update(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierControllerServer).Destroy(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

// SupplierController_ServiceDesc is the grpc.ServiceDesc for SupplierController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supplier.SupplierController",
	HandlerType: (*SupplierControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SupplierController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _SupplierController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SupplierController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _SupplierController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _SupplierController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "supplier.proto",
}
