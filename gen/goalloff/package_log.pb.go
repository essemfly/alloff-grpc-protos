// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: package_log.proto

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

type PackageLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	FieldName   string `protobuf:"bytes,4,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Before      string `protobuf:"bytes,5,opt,name=before,proto3" json:"before,omitempty"`
	After       string `protobuf:"bytes,6,opt,name=after,proto3" json:"after,omitempty"`
	CreatedBy   int32  `protobuf:"varint,7,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Package     int64  `protobuf:"varint,8,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *PackageLog) Reset() {
	*x = PackageLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageLog) ProtoMessage() {}

func (x *PackageLog) ProtoReflect() protoreflect.Message {
	mi := &file_package_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageLog.ProtoReflect.Descriptor instead.
func (*PackageLog) Descriptor() ([]byte, []int) {
	return file_package_log_proto_rawDescGZIP(), []int{0}
}

func (x *PackageLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PackageLog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PackageLog) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PackageLog) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *PackageLog) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *PackageLog) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *PackageLog) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *PackageLog) GetPackage() int64 {
	if x != nil {
		return x.Package
	}
	return 0
}

type PackageLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PackageLogListRequest) Reset() {
	*x = PackageLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageLogListRequest) ProtoMessage() {}

func (x *PackageLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageLogListRequest.ProtoReflect.Descriptor instead.
func (*PackageLogListRequest) Descriptor() ([]byte, []int) {
	return file_package_log_proto_rawDescGZIP(), []int{1}
}

type PackageLogRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PackageLogRetrieveRequest) Reset() {
	*x = PackageLogRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageLogRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageLogRetrieveRequest) ProtoMessage() {}

func (x *PackageLogRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageLogRetrieveRequest.ProtoReflect.Descriptor instead.
func (*PackageLogRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_package_log_proto_rawDescGZIP(), []int{2}
}

func (x *PackageLogRetrieveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_package_log_proto protoreflect.FileDescriptor

var file_package_log_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01,
	0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x19,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe8, 0x02, 0x0a, 0x14, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x1a, 0x17,
	0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x1a, 0x17, 0x2e, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c,
	0x6c, 0x6f, 0x66, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_log_proto_rawDescOnce sync.Once
	file_package_log_proto_rawDescData = file_package_log_proto_rawDesc
)

func file_package_log_proto_rawDescGZIP() []byte {
	file_package_log_proto_rawDescOnce.Do(func() {
		file_package_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_log_proto_rawDescData)
	})
	return file_package_log_proto_rawDescData
}

var file_package_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_package_log_proto_goTypes = []interface{}{
	(*PackageLog)(nil),                // 0: package_log.PackageLog
	(*PackageLogListRequest)(nil),     // 1: package_log.PackageLogListRequest
	(*PackageLogRetrieveRequest)(nil), // 2: package_log.PackageLogRetrieveRequest
	(*emptypb.Empty)(nil),             // 3: google.protobuf.Empty
}
var file_package_log_proto_depIdxs = []int32{
	1, // 0: package_log.PackageLogController.List:input_type -> package_log.PackageLogListRequest
	0, // 1: package_log.PackageLogController.Create:input_type -> package_log.PackageLog
	2, // 2: package_log.PackageLogController.Retrieve:input_type -> package_log.PackageLogRetrieveRequest
	0, // 3: package_log.PackageLogController.Update:input_type -> package_log.PackageLog
	0, // 4: package_log.PackageLogController.Destroy:input_type -> package_log.PackageLog
	0, // 5: package_log.PackageLogController.List:output_type -> package_log.PackageLog
	0, // 6: package_log.PackageLogController.Create:output_type -> package_log.PackageLog
	0, // 7: package_log.PackageLogController.Retrieve:output_type -> package_log.PackageLog
	0, // 8: package_log.PackageLogController.Update:output_type -> package_log.PackageLog
	3, // 9: package_log.PackageLogController.Destroy:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_package_log_proto_init() }
func file_package_log_proto_init() {
	if File_package_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_package_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageLog); i {
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
		file_package_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageLogListRequest); i {
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
		file_package_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageLogRetrieveRequest); i {
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
			RawDescriptor: file_package_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_log_proto_goTypes,
		DependencyIndexes: file_package_log_proto_depIdxs,
		MessageInfos:      file_package_log_proto_msgTypes,
	}.Build()
	File_package_log_proto = out.File
	file_package_log_proto_rawDesc = nil
	file_package_log_proto_goTypes = nil
	file_package_log_proto_depIdxs = nil
}
