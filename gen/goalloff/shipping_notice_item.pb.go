// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: shipping_notice_item.proto

package goalloff

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ShippingNoticeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NoticeId  int64      `protobuf:"varint,2,opt,name=notice_id,json=noticeId,proto3" json:"notice_id,omitempty"`
	Inventory *Inventory `protobuf:"bytes,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	OrderItem *OrderItem `protobuf:"bytes,4,opt,name=order_item,json=orderItem,proto3" json:"order_item,omitempty"`
	Package   *Package   `protobuf:"bytes,5,opt,name=package,proto3,oneof" json:"package,omitempty"`
}

func (x *ShippingNoticeItem) Reset() {
	*x = ShippingNoticeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_notice_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingNoticeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingNoticeItem) ProtoMessage() {}

func (x *ShippingNoticeItem) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_notice_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingNoticeItem.ProtoReflect.Descriptor instead.
func (*ShippingNoticeItem) Descriptor() ([]byte, []int) {
	return file_shipping_notice_item_proto_rawDescGZIP(), []int{0}
}

func (x *ShippingNoticeItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShippingNoticeItem) GetNoticeId() int64 {
	if x != nil {
		return x.NoticeId
	}
	return 0
}

func (x *ShippingNoticeItem) GetInventory() *Inventory {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *ShippingNoticeItem) GetOrderItem() *OrderItem {
	if x != nil {
		return x.OrderItem
	}
	return nil
}

func (x *ShippingNoticeItem) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type ShippingNoticeItemListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShippingNoticeItemListRequest) Reset() {
	*x = ShippingNoticeItemListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_notice_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingNoticeItemListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingNoticeItemListRequest) ProtoMessage() {}

func (x *ShippingNoticeItemListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_notice_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingNoticeItemListRequest.ProtoReflect.Descriptor instead.
func (*ShippingNoticeItemListRequest) Descriptor() ([]byte, []int) {
	return file_shipping_notice_item_proto_rawDescGZIP(), []int{1}
}

type ShippingNoticeItemRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShippingNoticeItemRetrieveRequest) Reset() {
	*x = ShippingNoticeItemRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_notice_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingNoticeItemRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingNoticeItemRetrieveRequest) ProtoMessage() {}

func (x *ShippingNoticeItemRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_notice_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingNoticeItemRetrieveRequest.ProtoReflect.Descriptor instead.
func (*ShippingNoticeItemRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_shipping_notice_item_proto_rawDescGZIP(), []int{2}
}

func (x *ShippingNoticeItemRetrieveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_shipping_notice_item_proto protoreflect.FileDescriptor

var file_shipping_notice_item_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x12, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x09,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2f,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x21,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x89, 0x04, 0x0a, 0x1c, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x69, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5e, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x1a, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x6f, 0x0a,
	0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x37, 0x2e, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x5e,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x1a, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x28, 0x2e, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x73,
	0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipping_notice_item_proto_rawDescOnce sync.Once
	file_shipping_notice_item_proto_rawDescData = file_shipping_notice_item_proto_rawDesc
)

func file_shipping_notice_item_proto_rawDescGZIP() []byte {
	file_shipping_notice_item_proto_rawDescOnce.Do(func() {
		file_shipping_notice_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipping_notice_item_proto_rawDescData)
	})
	return file_shipping_notice_item_proto_rawDescData
}

var file_shipping_notice_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shipping_notice_item_proto_goTypes = []interface{}{
	(*ShippingNoticeItem)(nil),                // 0: shipping_notice_item.ShippingNoticeItem
	(*ShippingNoticeItemListRequest)(nil),     // 1: shipping_notice_item.ShippingNoticeItemListRequest
	(*ShippingNoticeItemRetrieveRequest)(nil), // 2: shipping_notice_item.ShippingNoticeItemRetrieveRequest
	(*Inventory)(nil),                         // 3: inventory.Inventory
	(*OrderItem)(nil),                         // 4: orderitem.OrderItem
	(*Package)(nil),                           // 5: package.Package
	(*emptypb.Empty)(nil),                     // 6: google.protobuf.Empty
}
var file_shipping_notice_item_proto_depIdxs = []int32{
	3, // 0: shipping_notice_item.ShippingNoticeItem.inventory:type_name -> inventory.Inventory
	4, // 1: shipping_notice_item.ShippingNoticeItem.order_item:type_name -> orderitem.OrderItem
	5, // 2: shipping_notice_item.ShippingNoticeItem.package:type_name -> package.Package
	1, // 3: shipping_notice_item.ShippingNoticeItemController.List:input_type -> shipping_notice_item.ShippingNoticeItemListRequest
	0, // 4: shipping_notice_item.ShippingNoticeItemController.Create:input_type -> shipping_notice_item.ShippingNoticeItem
	2, // 5: shipping_notice_item.ShippingNoticeItemController.Retrieve:input_type -> shipping_notice_item.ShippingNoticeItemRetrieveRequest
	0, // 6: shipping_notice_item.ShippingNoticeItemController.Update:input_type -> shipping_notice_item.ShippingNoticeItem
	0, // 7: shipping_notice_item.ShippingNoticeItemController.Destroy:input_type -> shipping_notice_item.ShippingNoticeItem
	0, // 8: shipping_notice_item.ShippingNoticeItemController.List:output_type -> shipping_notice_item.ShippingNoticeItem
	0, // 9: shipping_notice_item.ShippingNoticeItemController.Create:output_type -> shipping_notice_item.ShippingNoticeItem
	0, // 10: shipping_notice_item.ShippingNoticeItemController.Retrieve:output_type -> shipping_notice_item.ShippingNoticeItem
	0, // 11: shipping_notice_item.ShippingNoticeItemController.Update:output_type -> shipping_notice_item.ShippingNoticeItem
	6, // 12: shipping_notice_item.ShippingNoticeItemController.Destroy:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shipping_notice_item_proto_init() }
func file_shipping_notice_item_proto_init() {
	if File_shipping_notice_item_proto != nil {
		return
	}
	file_inventory_proto_init()
	file_package_proto_init()
	file_order_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shipping_notice_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingNoticeItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shipping_notice_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingNoticeItemListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shipping_notice_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingNoticeItemRetrieveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_shipping_notice_item_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shipping_notice_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipping_notice_item_proto_goTypes,
		DependencyIndexes: file_shipping_notice_item_proto_depIdxs,
		MessageInfos:      file_shipping_notice_item_proto_msgTypes,
	}.Build()
	File_shipping_notice_item_proto = out.File
	file_shipping_notice_item_proto_rawDesc = nil
	file_shipping_notice_item_proto_goTypes = nil
	file_shipping_notice_item_proto_depIdxs = nil
}
