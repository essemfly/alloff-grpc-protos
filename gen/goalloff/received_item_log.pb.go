// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: received_item_log.proto

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

type ReceivedItemLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserUuid     string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	UserUsername string `protobuf:"bytes,3,opt,name=user_username,json=userUsername,proto3" json:"user_username,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	FieldName    string `protobuf:"bytes,6,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Before       string `protobuf:"bytes,7,opt,name=before,proto3" json:"before,omitempty"`
	After        string `protobuf:"bytes,8,opt,name=after,proto3" json:"after,omitempty"`
	ReceivedItem int64  `protobuf:"varint,9,opt,name=received_item,json=receivedItem,proto3" json:"received_item,omitempty"`
}

func (x *ReceivedItemLog) Reset() {
	*x = ReceivedItemLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_received_item_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivedItemLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedItemLog) ProtoMessage() {}

func (x *ReceivedItemLog) ProtoReflect() protoreflect.Message {
	mi := &file_received_item_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedItemLog.ProtoReflect.Descriptor instead.
func (*ReceivedItemLog) Descriptor() ([]byte, []int) {
	return file_received_item_log_proto_rawDescGZIP(), []int{0}
}

func (x *ReceivedItemLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReceivedItemLog) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *ReceivedItemLog) GetUserUsername() string {
	if x != nil {
		return x.UserUsername
	}
	return ""
}

func (x *ReceivedItemLog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReceivedItemLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ReceivedItemLog) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *ReceivedItemLog) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *ReceivedItemLog) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *ReceivedItemLog) GetReceivedItem() int64 {
	if x != nil {
		return x.ReceivedItem
	}
	return 0
}

type ReceivedItemLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReceivedItemLogListRequest) Reset() {
	*x = ReceivedItemLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_received_item_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivedItemLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedItemLogListRequest) ProtoMessage() {}

func (x *ReceivedItemLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_received_item_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedItemLogListRequest.ProtoReflect.Descriptor instead.
func (*ReceivedItemLogListRequest) Descriptor() ([]byte, []int) {
	return file_received_item_log_proto_rawDescGZIP(), []int{1}
}

type ReceivedItemLogRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReceivedItemLogRetrieveRequest) Reset() {
	*x = ReceivedItemLogRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_received_item_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivedItemLogRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedItemLogRetrieveRequest) ProtoMessage() {}

func (x *ReceivedItemLogRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_received_item_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedItemLogRetrieveRequest.ProtoReflect.Descriptor instead.
func (*ReceivedItemLogRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_received_item_log_proto_rawDescGZIP(), []int{2}
}

func (x *ReceivedItemLogRetrieveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_received_item_log_proto protoreflect.FileDescriptor

var file_received_item_log_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x22, 0x1c, 0x0a, 0x1a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30,
	0x0a, 0x1e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xbe, 0x03, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x59,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x69, 0x74,
	0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x69, 0x74,
	0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x08, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x2f, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x69,
	0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x07, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x69, 0x74, 0x65, 0x6d, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x65, 0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66,
	0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_received_item_log_proto_rawDescOnce sync.Once
	file_received_item_log_proto_rawDescData = file_received_item_log_proto_rawDesc
)

func file_received_item_log_proto_rawDescGZIP() []byte {
	file_received_item_log_proto_rawDescOnce.Do(func() {
		file_received_item_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_received_item_log_proto_rawDescData)
	})
	return file_received_item_log_proto_rawDescData
}

var file_received_item_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_received_item_log_proto_goTypes = []interface{}{
	(*ReceivedItemLog)(nil),                // 0: receiveditemlog.ReceivedItemLog
	(*ReceivedItemLogListRequest)(nil),     // 1: receiveditemlog.ReceivedItemLogListRequest
	(*ReceivedItemLogRetrieveRequest)(nil), // 2: receiveditemlog.ReceivedItemLogRetrieveRequest
	(*emptypb.Empty)(nil),                  // 3: google.protobuf.Empty
}
var file_received_item_log_proto_depIdxs = []int32{
	1, // 0: receiveditemlog.ReceivedItemLogController.List:input_type -> receiveditemlog.ReceivedItemLogListRequest
	0, // 1: receiveditemlog.ReceivedItemLogController.Create:input_type -> receiveditemlog.ReceivedItemLog
	2, // 2: receiveditemlog.ReceivedItemLogController.Retrieve:input_type -> receiveditemlog.ReceivedItemLogRetrieveRequest
	0, // 3: receiveditemlog.ReceivedItemLogController.Update:input_type -> receiveditemlog.ReceivedItemLog
	0, // 4: receiveditemlog.ReceivedItemLogController.Destroy:input_type -> receiveditemlog.ReceivedItemLog
	0, // 5: receiveditemlog.ReceivedItemLogController.List:output_type -> receiveditemlog.ReceivedItemLog
	0, // 6: receiveditemlog.ReceivedItemLogController.Create:output_type -> receiveditemlog.ReceivedItemLog
	0, // 7: receiveditemlog.ReceivedItemLogController.Retrieve:output_type -> receiveditemlog.ReceivedItemLog
	0, // 8: receiveditemlog.ReceivedItemLogController.Update:output_type -> receiveditemlog.ReceivedItemLog
	3, // 9: receiveditemlog.ReceivedItemLogController.Destroy:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_received_item_log_proto_init() }
func file_received_item_log_proto_init() {
	if File_received_item_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_received_item_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivedItemLog); i {
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
		file_received_item_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivedItemLogListRequest); i {
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
		file_received_item_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivedItemLogRetrieveRequest); i {
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
			RawDescriptor: file_received_item_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_received_item_log_proto_goTypes,
		DependencyIndexes: file_received_item_log_proto_depIdxs,
		MessageInfos:      file_received_item_log_proto_msgTypes,
	}.Build()
	File_received_item_log_proto = out.File
	file_received_item_log_proto_rawDesc = nil
	file_received_item_log_proto_goTypes = nil
	file_received_item_log_proto_depIdxs = nil
}