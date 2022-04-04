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

// PackageControllerClient is the client API for PackageController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageControllerClient interface {
	List(ctx context.Context, in *PackageListRequest, opts ...grpc.CallOption) (PackageController_ListClient, error)
	Create(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error)
	Retrieve(ctx context.Context, in *PackageRetrieveRequest, opts ...grpc.CallOption) (*Package, error)
	Update(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error)
	Destroy(ctx context.Context, in *Package, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type packageControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageControllerClient(cc grpc.ClientConnInterface) PackageControllerClient {
	return &packageControllerClient{cc}
}

func (c *packageControllerClient) List(ctx context.Context, in *PackageListRequest, opts ...grpc.CallOption) (PackageController_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &PackageController_ServiceDesc.Streams[0], "/package.PackageController/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &packageControllerListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PackageController_ListClient interface {
	Recv() (*Package, error)
	grpc.ClientStream
}

type packageControllerListClient struct {
	grpc.ClientStream
}

func (x *packageControllerListClient) Recv() (*Package, error) {
	m := new(Package)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *packageControllerClient) Create(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/package.PackageController/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageControllerClient) Retrieve(ctx context.Context, in *PackageRetrieveRequest, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/package.PackageController/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageControllerClient) Update(ctx context.Context, in *Package, opts ...grpc.CallOption) (*Package, error) {
	out := new(Package)
	err := c.cc.Invoke(ctx, "/package.PackageController/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageControllerClient) Destroy(ctx context.Context, in *Package, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/package.PackageController/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageControllerServer is the server API for PackageController service.
// All implementations must embed UnimplementedPackageControllerServer
// for forward compatibility
type PackageControllerServer interface {
	List(*PackageListRequest, PackageController_ListServer) error
	Create(context.Context, *Package) (*Package, error)
	Retrieve(context.Context, *PackageRetrieveRequest) (*Package, error)
	Update(context.Context, *Package) (*Package, error)
	Destroy(context.Context, *Package) (*emptypb.Empty, error)
	mustEmbedUnimplementedPackageControllerServer()
}

// UnimplementedPackageControllerServer must be embedded to have forward compatible implementations.
type UnimplementedPackageControllerServer struct {
}

func (UnimplementedPackageControllerServer) List(*PackageListRequest, PackageController_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPackageControllerServer) Create(context.Context, *Package) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPackageControllerServer) Retrieve(context.Context, *PackageRetrieveRequest) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedPackageControllerServer) Update(context.Context, *Package) (*Package, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPackageControllerServer) Destroy(context.Context, *Package) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedPackageControllerServer) mustEmbedUnimplementedPackageControllerServer() {}

// UnsafePackageControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageControllerServer will
// result in compilation errors.
type UnsafePackageControllerServer interface {
	mustEmbedUnimplementedPackageControllerServer()
}

func RegisterPackageControllerServer(s grpc.ServiceRegistrar, srv PackageControllerServer) {
	s.RegisterService(&PackageController_ServiceDesc, srv)
}

func _PackageController_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PackageListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PackageControllerServer).List(m, &packageControllerListServer{stream})
}

type PackageController_ListServer interface {
	Send(*Package) error
	grpc.ServerStream
}

type packageControllerListServer struct {
	grpc.ServerStream
}

func (x *packageControllerListServer) Send(m *Package) error {
	return x.ServerStream.SendMsg(m)
}

func _PackageController_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageControllerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/package.PackageController/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageControllerServer).Create(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageController_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackageRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageControllerServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/package.PackageController/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageControllerServer).Retrieve(ctx, req.(*PackageRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageController_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageControllerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/package.PackageController/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageControllerServer).Update(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageController_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageControllerServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/package.PackageController/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageControllerServer).Destroy(ctx, req.(*Package))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageController_ServiceDesc is the grpc.ServiceDesc for PackageController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "package.PackageController",
	HandlerType: (*PackageControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PackageController_Create_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _PackageController_Retrieve_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PackageController_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _PackageController_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _PackageController_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "package.proto",
}
