// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: order_item_memo.proto

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

type OrderItemMemo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserUuid     string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	UserUsername string `protobuf:"bytes,3,opt,name=user_username,json=userUsername,proto3" json:"user_username,omitempty"`
	Body         string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DeletedAt    string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	OrderItem    int64  `protobuf:"varint,7,opt,name=order_item,json=orderItem,proto3" json:"order_item,omitempty"`
}

func (x *OrderItemMemo) Reset() {
	*x = OrderItemMemo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_item_memo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemMemo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemMemo) ProtoMessage() {}

func (x *OrderItemMemo) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_memo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemMemo.ProtoReflect.Descriptor instead.
func (*OrderItemMemo) Descriptor() ([]byte, []int) {
	return file_order_item_memo_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItemMemo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemMemo) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *OrderItemMemo) GetUserUsername() string {
	if x != nil {
		return x.UserUsername
	}
	return ""
}

func (x *OrderItemMemo) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *OrderItemMemo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderItemMemo) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *OrderItemMemo) GetOrderItem() int64 {
	if x != nil {
		return x.OrderItem
	}
	return 0
}

type OrderItemMemoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderItemMemoListRequest) Reset() {
	*x = OrderItemMemoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_item_memo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemMemoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemMemoListRequest) ProtoMessage() {}

func (x *OrderItemMemoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_memo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemMemoListRequest.ProtoReflect.Descriptor instead.
func (*OrderItemMemoListRequest) Descriptor() ([]byte, []int) {
	return file_order_item_memo_proto_rawDescGZIP(), []int{1}
}

type OrderItemMemoRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderItemMemoRetrieveRequest) Reset() {
	*x = OrderItemMemoRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_item_memo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemMemoRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemMemoRetrieveRequest) ProtoMessage() {}

func (x *OrderItemMemoRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_memo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemMemoRetrieveRequest.ProtoReflect.Descriptor instead.
func (*OrderItemMemoRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_order_item_memo_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItemMemoRetrieveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_order_item_memo_proto protoreflect.FileDescriptor

var file_order_item_memo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x6d,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74,
	0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x1a, 0x0a, 0x18, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x1c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x98, 0x03, 0x0a, 0x17, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x51, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d,
	0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x1c, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x08, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x2b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69,
	0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d,
	0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65,
	0x6d, 0x6f, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x1c, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07,
	0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69,
	0x74, 0x65, 0x6d, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x4d, 0x65, 0x6d, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_item_memo_proto_rawDescOnce sync.Once
	file_order_item_memo_proto_rawDescData = file_order_item_memo_proto_rawDesc
)

func file_order_item_memo_proto_rawDescGZIP() []byte {
	file_order_item_memo_proto_rawDescOnce.Do(func() {
		file_order_item_memo_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_item_memo_proto_rawDescData)
	})
	return file_order_item_memo_proto_rawDescData
}

var file_order_item_memo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_item_memo_proto_goTypes = []interface{}{
	(*OrderItemMemo)(nil),                // 0: orderitemmemo.OrderItemMemo
	(*OrderItemMemoListRequest)(nil),     // 1: orderitemmemo.OrderItemMemoListRequest
	(*OrderItemMemoRetrieveRequest)(nil), // 2: orderitemmemo.OrderItemMemoRetrieveRequest
	(*emptypb.Empty)(nil),                // 3: google.protobuf.Empty
}
var file_order_item_memo_proto_depIdxs = []int32{
	1, // 0: orderitemmemo.OrderItemMemoController.List:input_type -> orderitemmemo.OrderItemMemoListRequest
	0, // 1: orderitemmemo.OrderItemMemoController.Create:input_type -> orderitemmemo.OrderItemMemo
	2, // 2: orderitemmemo.OrderItemMemoController.Retrieve:input_type -> orderitemmemo.OrderItemMemoRetrieveRequest
	0, // 3: orderitemmemo.OrderItemMemoController.Update:input_type -> orderitemmemo.OrderItemMemo
	0, // 4: orderitemmemo.OrderItemMemoController.Destroy:input_type -> orderitemmemo.OrderItemMemo
	0, // 5: orderitemmemo.OrderItemMemoController.List:output_type -> orderitemmemo.OrderItemMemo
	0, // 6: orderitemmemo.OrderItemMemoController.Create:output_type -> orderitemmemo.OrderItemMemo
	0, // 7: orderitemmemo.OrderItemMemoController.Retrieve:output_type -> orderitemmemo.OrderItemMemo
	0, // 8: orderitemmemo.OrderItemMemoController.Update:output_type -> orderitemmemo.OrderItemMemo
	3, // 9: orderitemmemo.OrderItemMemoController.Destroy:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_item_memo_proto_init() }
func file_order_item_memo_proto_init() {
	if File_order_item_memo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_item_memo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemMemo); i {
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
		file_order_item_memo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemMemoListRequest); i {
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
		file_order_item_memo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemMemoRetrieveRequest); i {
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
			RawDescriptor: file_order_item_memo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_item_memo_proto_goTypes,
		DependencyIndexes: file_order_item_memo_proto_depIdxs,
		MessageInfos:      file_order_item_memo_proto_msgTypes,
	}.Build()
	File_order_item_memo_proto = out.File
	file_order_item_memo_proto_rawDesc = nil
	file_order_item_memo_proto_goTypes = nil
	file_order_item_memo_proto_depIdxs = nil
}
