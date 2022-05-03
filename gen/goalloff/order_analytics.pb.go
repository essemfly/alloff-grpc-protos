// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: order_analytics.proto

package goalloff

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DailyOrderStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DateFrom string `protobuf:"bytes,1,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo   string `protobuf:"bytes,2,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
}

func (x *DailyOrderStatRequest) Reset() {
	*x = DailyOrderStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyOrderStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyOrderStatRequest) ProtoMessage() {}

func (x *DailyOrderStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyOrderStatRequest.ProtoReflect.Descriptor instead.
func (*DailyOrderStatRequest) Descriptor() ([]byte, []int) {
	return file_order_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *DailyOrderStatRequest) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *DailyOrderStatRequest) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type OrderStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderStats []*OrderStat `protobuf:"bytes,1,rep,name=order_stats,json=orderStats,proto3" json:"order_stats,omitempty"`
	Summary    *OrderStat   `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *OrderStatResponse) Reset() {
	*x = OrderStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatResponse) ProtoMessage() {}

func (x *OrderStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatResponse.ProtoReflect.Descriptor instead.
func (*OrderStatResponse) Descriptor() ([]byte, []int) {
	return file_order_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *OrderStatResponse) GetOrderStats() []*OrderStat {
	if x != nil {
		return x.OrderStats
	}
	return nil
}

func (x *OrderStatResponse) GetSummary() *OrderStat {
	if x != nil {
		return x.Summary
	}
	return nil
}

type OrderStatFigure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Amount   int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Quantity int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderStatFigure) Reset() {
	*x = OrderStatFigure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatFigure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatFigure) ProtoMessage() {}

func (x *OrderStatFigure) ProtoReflect() protoreflect.Message {
	mi := &file_order_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatFigure.ProtoReflect.Descriptor instead.
func (*OrderStatFigure) Descriptor() ([]byte, []int) {
	return file_order_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *OrderStatFigure) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *OrderStatFigure) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderStatFigure) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OrderStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string           `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Total     *OrderStatFigure `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	Paid      *OrderStatFigure `protobuf:"bytes,3,opt,name=paid,proto3" json:"paid,omitempty"`
	Canceled  *OrderStatFigure `protobuf:"bytes,4,opt,name=canceled,proto3" json:"canceled,omitempty"`
}

func (x *OrderStat) Reset() {
	*x = OrderStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStat) ProtoMessage() {}

func (x *OrderStat) ProtoReflect() protoreflect.Message {
	mi := &file_order_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStat.ProtoReflect.Descriptor instead.
func (*OrderStat) Descriptor() ([]byte, []int) {
	return file_order_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *OrderStat) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *OrderStat) GetTotal() *OrderStatFigure {
	if x != nil {
		return x.Total
	}
	return nil
}

func (x *OrderStat) GetPaid() *OrderStatFigure {
	if x != nil {
		return x.Paid
	}
	return nil
}

func (x *OrderStat) GetCanceled() *OrderStatFigure {
	if x != nil {
		return x.Canceled
	}
	return nil
}

var File_order_analytics_proto protoreflect.FileDescriptor

var file_order_analytics_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x22, 0x4d, 0x0a, 0x15, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x5b, 0x0a,
	0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xd2, 0x01, 0x0a, 0x09, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a,
	0x04, 0x70, 0x61, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x04, 0x70, 0x61,
	0x69, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x46,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x32,
	0x78, 0x0a, 0x18, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x0e, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x25, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f,
	0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_analytics_proto_rawDescOnce sync.Once
	file_order_analytics_proto_rawDescData = file_order_analytics_proto_rawDesc
)

func file_order_analytics_proto_rawDescGZIP() []byte {
	file_order_analytics_proto_rawDescOnce.Do(func() {
		file_order_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_analytics_proto_rawDescData)
	})
	return file_order_analytics_proto_rawDescData
}

var file_order_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_analytics_proto_goTypes = []interface{}{
	(*DailyOrderStatRequest)(nil), // 0: orderanalytics.DailyOrderStatRequest
	(*OrderStatResponse)(nil),     // 1: orderanalytics.OrderStatResponse
	(*OrderStatFigure)(nil),       // 2: orderanalytics.OrderStatFigure
	(*OrderStat)(nil),             // 3: orderanalytics.OrderStat
}
var file_order_analytics_proto_depIdxs = []int32{
	3, // 0: orderanalytics.OrderStatResponse.order_stats:type_name -> orderanalytics.OrderStat
	3, // 1: orderanalytics.OrderStatResponse.summary:type_name -> orderanalytics.OrderStat
	2, // 2: orderanalytics.OrderStat.total:type_name -> orderanalytics.OrderStatFigure
	2, // 3: orderanalytics.OrderStat.paid:type_name -> orderanalytics.OrderStatFigure
	2, // 4: orderanalytics.OrderStat.canceled:type_name -> orderanalytics.OrderStatFigure
	0, // 5: orderanalytics.OrderAnalyticsController.DailyOrderStat:input_type -> orderanalytics.DailyOrderStatRequest
	1, // 6: orderanalytics.OrderAnalyticsController.DailyOrderStat:output_type -> orderanalytics.OrderStatResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_order_analytics_proto_init() }
func file_order_analytics_proto_init() {
	if File_order_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyOrderStatRequest); i {
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
		file_order_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatResponse); i {
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
		file_order_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatFigure); i {
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
		file_order_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStat); i {
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
			RawDescriptor: file_order_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_analytics_proto_goTypes,
		DependencyIndexes: file_order_analytics_proto_depIdxs,
		MessageInfos:      file_order_analytics_proto_msgTypes,
	}.Build()
	File_order_analytics_proto = out.File
	file_order_analytics_proto_rawDesc = nil
	file_order_analytics_proto_goTypes = nil
	file_order_analytics_proto_depIdxs = nil
}
