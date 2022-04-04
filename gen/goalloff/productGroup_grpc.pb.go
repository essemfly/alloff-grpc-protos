// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package goalloff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductGroupClient is the client API for ProductGroup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductGroupClient interface {
	GetProductGroup(ctx context.Context, in *GetProductGroupRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error)
	CreateProductGroup(ctx context.Context, in *CreateProductGroupRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error)
	ListProductGroups(ctx context.Context, in *ListProductGroupsRequest, opts ...grpc.CallOption) (*ListProductGroupsResponse, error)
	EditProductGroup(ctx context.Context, in *EditProductGroupRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error)
	PushProductsInProductGroup(ctx context.Context, in *ProductsInPgRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error)
	UpdateProductsInProductGroup(ctx context.Context, in *ProductsInPgRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error)
	RemoveProductInProductGroup(ctx context.Context, in *RemoveProductInPgRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error)
}

type productGroupClient struct {
	cc grpc.ClientConnInterface
}

func NewProductGroupClient(cc grpc.ClientConnInterface) ProductGroupClient {
	return &productGroupClient{cc}
}

func (c *productGroupClient) GetProductGroup(ctx context.Context, in *GetProductGroupRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error) {
	out := new(ProductGroupMessage)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/GetProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGroupClient) CreateProductGroup(ctx context.Context, in *CreateProductGroupRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error) {
	out := new(ProductGroupMessage)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/CreateProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGroupClient) ListProductGroups(ctx context.Context, in *ListProductGroupsRequest, opts ...grpc.CallOption) (*ListProductGroupsResponse, error) {
	out := new(ListProductGroupsResponse)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/ListProductGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGroupClient) EditProductGroup(ctx context.Context, in *EditProductGroupRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error) {
	out := new(ProductGroupMessage)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/EditProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGroupClient) PushProductsInProductGroup(ctx context.Context, in *ProductsInPgRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error) {
	out := new(ProductGroupMessage)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/PushProductsInProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGroupClient) UpdateProductsInProductGroup(ctx context.Context, in *ProductsInPgRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error) {
	out := new(ProductGroupMessage)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/UpdateProductsInProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productGroupClient) RemoveProductInProductGroup(ctx context.Context, in *RemoveProductInPgRequest, opts ...grpc.CallOption) (*ProductGroupMessage, error) {
	out := new(ProductGroupMessage)
	err := c.cc.Invoke(ctx, "/goalloff.ProductGroup/RemoveProductInProductGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductGroupServer is the server API for ProductGroup service.
// All implementations must embed UnimplementedProductGroupServer
// for forward compatibility
type ProductGroupServer interface {
	GetProductGroup(context.Context, *GetProductGroupRequest) (*ProductGroupMessage, error)
	CreateProductGroup(context.Context, *CreateProductGroupRequest) (*ProductGroupMessage, error)
	ListProductGroups(context.Context, *ListProductGroupsRequest) (*ListProductGroupsResponse, error)
	EditProductGroup(context.Context, *EditProductGroupRequest) (*ProductGroupMessage, error)
	PushProductsInProductGroup(context.Context, *ProductsInPgRequest) (*ProductGroupMessage, error)
	UpdateProductsInProductGroup(context.Context, *ProductsInPgRequest) (*ProductGroupMessage, error)
	RemoveProductInProductGroup(context.Context, *RemoveProductInPgRequest) (*ProductGroupMessage, error)
	mustEmbedUnimplementedProductGroupServer()
}

// UnimplementedProductGroupServer must be embedded to have forward compatible implementations.
type UnimplementedProductGroupServer struct {
}

func (UnimplementedProductGroupServer) GetProductGroup(context.Context, *GetProductGroupRequest) (*ProductGroupMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductGroup not implemented")
}
func (UnimplementedProductGroupServer) CreateProductGroup(context.Context, *CreateProductGroupRequest) (*ProductGroupMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductGroup not implemented")
}
func (UnimplementedProductGroupServer) ListProductGroups(context.Context, *ListProductGroupsRequest) (*ListProductGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductGroups not implemented")
}
func (UnimplementedProductGroupServer) EditProductGroup(context.Context, *EditProductGroupRequest) (*ProductGroupMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProductGroup not implemented")
}
func (UnimplementedProductGroupServer) PushProductsInProductGroup(context.Context, *ProductsInPgRequest) (*ProductGroupMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushProductsInProductGroup not implemented")
}
func (UnimplementedProductGroupServer) UpdateProductsInProductGroup(context.Context, *ProductsInPgRequest) (*ProductGroupMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductsInProductGroup not implemented")
}
func (UnimplementedProductGroupServer) RemoveProductInProductGroup(context.Context, *RemoveProductInPgRequest) (*ProductGroupMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductInProductGroup not implemented")
}
func (UnimplementedProductGroupServer) mustEmbedUnimplementedProductGroupServer() {}

// UnsafeProductGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductGroupServer will
// result in compilation errors.
type UnsafeProductGroupServer interface {
	mustEmbedUnimplementedProductGroupServer()
}

func RegisterProductGroupServer(s grpc.ServiceRegistrar, srv ProductGroupServer) {
	s.RegisterService(&ProductGroup_ServiceDesc, srv)
}

func _ProductGroup_GetProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).GetProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/GetProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).GetProductGroup(ctx, req.(*GetProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGroup_CreateProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).CreateProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/CreateProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).CreateProductGroup(ctx, req.(*CreateProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGroup_ListProductGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).ListProductGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/ListProductGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).ListProductGroups(ctx, req.(*ListProductGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGroup_EditProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProductGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).EditProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/EditProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).EditProductGroup(ctx, req.(*EditProductGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGroup_PushProductsInProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsInPgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).PushProductsInProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/PushProductsInProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).PushProductsInProductGroup(ctx, req.(*ProductsInPgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGroup_UpdateProductsInProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsInPgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).UpdateProductsInProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/UpdateProductsInProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).UpdateProductsInProductGroup(ctx, req.(*ProductsInPgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductGroup_RemoveProductInProductGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductInPgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductGroupServer).RemoveProductInProductGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goalloff.ProductGroup/RemoveProductInProductGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductGroupServer).RemoveProductInProductGroup(ctx, req.(*RemoveProductInPgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductGroup_ServiceDesc is the grpc.ServiceDesc for ProductGroup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductGroup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goalloff.ProductGroup",
	HandlerType: (*ProductGroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductGroup",
			Handler:    _ProductGroup_GetProductGroup_Handler,
		},
		{
			MethodName: "CreateProductGroup",
			Handler:    _ProductGroup_CreateProductGroup_Handler,
		},
		{
			MethodName: "ListProductGroups",
			Handler:    _ProductGroup_ListProductGroups_Handler,
		},
		{
			MethodName: "EditProductGroup",
			Handler:    _ProductGroup_EditProductGroup_Handler,
		},
		{
			MethodName: "PushProductsInProductGroup",
			Handler:    _ProductGroup_PushProductsInProductGroup_Handler,
		},
		{
			MethodName: "UpdateProductsInProductGroup",
			Handler:    _ProductGroup_UpdateProductsInProductGroup_Handler,
		},
		{
			MethodName: "RemoveProductInProductGroup",
			Handler:    _ProductGroup_RemoveProductInProductGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productGroup.proto",
}
