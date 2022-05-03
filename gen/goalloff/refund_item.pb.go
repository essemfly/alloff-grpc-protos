// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: refund_item.proto

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

type RefundItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RefundFee    int32  `protobuf:"varint,2,opt,name=refund_fee,json=refundFee,proto3" json:"refund_fee,omitempty"`
	RefundAmount int32  `protobuf:"varint,3,opt,name=refund_amount,json=refundAmount,proto3" json:"refund_amount,omitempty"`
	CreatedAt    string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Order        int64  `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
	OrderItem    int64  `protobuf:"varint,7,opt,name=order_item,json=orderItem,proto3" json:"order_item,omitempty"`
}

func (x *RefundItem) Reset() {
	*x = RefundItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refund_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundItem) ProtoMessage() {}

func (x *RefundItem) ProtoReflect() protoreflect.Message {
	mi := &file_refund_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundItem.ProtoReflect.Descriptor instead.
func (*RefundItem) Descriptor() ([]byte, []int) {
	return file_refund_item_proto_rawDescGZIP(), []int{0}
}

func (x *RefundItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RefundItem) GetRefundFee() int32 {
	if x != nil {
		return x.RefundFee
	}
	return 0
}

func (x *RefundItem) GetRefundAmount() int32 {
	if x != nil {
		return x.RefundAmount
	}
	return 0
}

func (x *RefundItem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RefundItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RefundItem) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *RefundItem) GetOrderItem() int64 {
	if x != nil {
		return x.OrderItem
	}
	return 0
}

type RefundItemListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefundItemListRequest) Reset() {
	*x = RefundItemListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refund_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundItemListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundItemListRequest) ProtoMessage() {}

func (x *RefundItemListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_refund_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundItemListRequest.ProtoReflect.Descriptor instead.
func (*RefundItemListRequest) Descriptor() ([]byte, []int) {
	return file_refund_item_proto_rawDescGZIP(), []int{1}
}

type RefundItemRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RefundItemRetrieveRequest) Reset() {
	*x = RefundItemRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_refund_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundItemRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundItemRetrieveRequest) ProtoMessage() {}

func (x *RefundItemRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_refund_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundItemRetrieveRequest.ProtoReflect.Descriptor instead.
func (*RefundItemRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_refund_item_proto_rawDescGZIP(), []int{2}
}

func (x *RefundItemRetrieveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_refund_item_proto protoreflect.FileDescriptor

var file_refund_item_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a,
	0x0a, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x46, 0x65, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x19, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xdf, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x12, 0x25, 0x2e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x72, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x75, 0x74,
	0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x6c,
	0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_refund_item_proto_rawDescOnce sync.Once
	file_refund_item_proto_rawDescData = file_refund_item_proto_rawDesc
)

func file_refund_item_proto_rawDescGZIP() []byte {
	file_refund_item_proto_rawDescOnce.Do(func() {
		file_refund_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_refund_item_proto_rawDescData)
	})
	return file_refund_item_proto_rawDescData
}

var file_refund_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_refund_item_proto_goTypes = []interface{}{
	(*RefundItem)(nil),                // 0: refunditem.RefundItem
	(*RefundItemListRequest)(nil),     // 1: refunditem.RefundItemListRequest
	(*RefundItemRetrieveRequest)(nil), // 2: refunditem.RefundItemRetrieveRequest
	(*emptypb.Empty)(nil),             // 3: google.protobuf.Empty
}
var file_refund_item_proto_depIdxs = []int32{
	1, // 0: refunditem.RefundItemController.List:input_type -> refunditem.RefundItemListRequest
	0, // 1: refunditem.RefundItemController.Create:input_type -> refunditem.RefundItem
	2, // 2: refunditem.RefundItemController.Retrieve:input_type -> refunditem.RefundItemRetrieveRequest
	0, // 3: refunditem.RefundItemController.Update:input_type -> refunditem.RefundItem
	0, // 4: refunditem.RefundItemController.Destroy:input_type -> refunditem.RefundItem
	0, // 5: refunditem.RefundItemController.List:output_type -> refunditem.RefundItem
	0, // 6: refunditem.RefundItemController.Create:output_type -> refunditem.RefundItem
	0, // 7: refunditem.RefundItemController.Retrieve:output_type -> refunditem.RefundItem
	0, // 8: refunditem.RefundItemController.Update:output_type -> refunditem.RefundItem
	3, // 9: refunditem.RefundItemController.Destroy:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_refund_item_proto_init() }
func file_refund_item_proto_init() {
	if File_refund_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_refund_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundItem); i {
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
		file_refund_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundItemListRequest); i {
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
		file_refund_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundItemRetrieveRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_refund_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_refund_item_proto_goTypes,
		DependencyIndexes: file_refund_item_proto_depIdxs,
		MessageInfos:      file_refund_item_proto_msgTypes,
	}.Build()
	File_refund_item_proto = out.File
	file_refund_item_proto_rawDesc = nil
	file_refund_item_proto_goTypes = nil
	file_refund_item_proto_depIdxs = nil
}
