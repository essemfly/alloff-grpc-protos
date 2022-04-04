// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: payment.proto

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

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ImpUid                string  `protobuf:"bytes,2,opt,name=imp_uid,json=impUid,proto3" json:"imp_uid,omitempty"`
	PaymentStatus         string  `protobuf:"bytes,3,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	Pg                    string  `protobuf:"bytes,4,opt,name=pg,proto3" json:"pg,omitempty"`
	PayMethod             string  `protobuf:"bytes,5,opt,name=pay_method,json=payMethod,proto3" json:"pay_method,omitempty"`
	Name                  string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	MerchantUid           string  `protobuf:"bytes,7,opt,name=merchant_uid,json=merchantUid,proto3" json:"merchant_uid,omitempty"`
	Amount                int32   `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
	BuyerName             string  `protobuf:"bytes,9,opt,name=buyer_name,json=buyerName,proto3" json:"buyer_name,omitempty"`
	BuyerMobile           string  `protobuf:"bytes,10,opt,name=buyer_mobile,json=buyerMobile,proto3" json:"buyer_mobile,omitempty"`
	BuyerAddress          string  `protobuf:"bytes,11,opt,name=buyer_address,json=buyerAddress,proto3" json:"buyer_address,omitempty"`
	BuyerPostCode         string  `protobuf:"bytes,12,opt,name=buyer_post_code,json=buyerPostCode,proto3" json:"buyer_post_code,omitempty"`
	Company               string  `protobuf:"bytes,13,opt,name=company,proto3" json:"company,omitempty"`
	AppScheme             string  `protobuf:"bytes,14,opt,name=app_scheme,json=appScheme,proto3" json:"app_scheme,omitempty"`
	CreatedAt             string  `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             string  `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PersonalCustomsNumber *string `protobuf:"bytes,17,opt,name=personal_customs_number,json=personalCustomsNumber,proto3,oneof" json:"personal_customs_number,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetImpUid() string {
	if x != nil {
		return x.ImpUid
	}
	return ""
}

func (x *Payment) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *Payment) GetPg() string {
	if x != nil {
		return x.Pg
	}
	return ""
}

func (x *Payment) GetPayMethod() string {
	if x != nil {
		return x.PayMethod
	}
	return ""
}

func (x *Payment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Payment) GetMerchantUid() string {
	if x != nil {
		return x.MerchantUid
	}
	return ""
}

func (x *Payment) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Payment) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *Payment) GetBuyerMobile() string {
	if x != nil {
		return x.BuyerMobile
	}
	return ""
}

func (x *Payment) GetBuyerAddress() string {
	if x != nil {
		return x.BuyerAddress
	}
	return ""
}

func (x *Payment) GetBuyerPostCode() string {
	if x != nil {
		return x.BuyerPostCode
	}
	return ""
}

func (x *Payment) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Payment) GetAppScheme() string {
	if x != nil {
		return x.AppScheme
	}
	return ""
}

func (x *Payment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Payment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Payment) GetPersonalCustomsNumber() string {
	if x != nil && x.PersonalCustomsNumber != nil {
		return *x.PersonalCustomsNumber
	}
	return ""
}

type PaymentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaymentListRequest) Reset() {
	*x = PaymentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentListRequest) ProtoMessage() {}

func (x *PaymentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentListRequest.ProtoReflect.Descriptor instead.
func (*PaymentListRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

type PaymentRetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PaymentRetrieveRequest) Reset() {
	*x = PaymentRetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRetrieveRequest) ProtoMessage() {}

func (x *PaymentRetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRetrieveRequest.ProtoReflect.Descriptor instead.
func (*PaymentRetrieveRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentRetrieveRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x04, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x70, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x70, 0x55, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70,
	0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x79, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x62, 0x75, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70,
	0x70, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x17, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x15, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x14,
	0x0a, 0x12, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa6,
	0x02, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x10, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_payment_proto_goTypes = []interface{}{
	(*Payment)(nil),                // 0: payment.Payment
	(*PaymentListRequest)(nil),     // 1: payment.PaymentListRequest
	(*PaymentRetrieveRequest)(nil), // 2: payment.PaymentRetrieveRequest
	(*emptypb.Empty)(nil),          // 3: google.protobuf.Empty
}
var file_payment_proto_depIdxs = []int32{
	1, // 0: payment.PaymentController.List:input_type -> payment.PaymentListRequest
	0, // 1: payment.PaymentController.Create:input_type -> payment.Payment
	2, // 2: payment.PaymentController.Retrieve:input_type -> payment.PaymentRetrieveRequest
	0, // 3: payment.PaymentController.Update:input_type -> payment.Payment
	0, // 4: payment.PaymentController.Destroy:input_type -> payment.Payment
	0, // 5: payment.PaymentController.List:output_type -> payment.Payment
	0, // 6: payment.PaymentController.Create:output_type -> payment.Payment
	0, // 7: payment.PaymentController.Retrieve:output_type -> payment.Payment
	0, // 8: payment.PaymentController.Update:output_type -> payment.Payment
	3, // 9: payment.PaymentController.Destroy:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentListRequest); i {
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
		file_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRetrieveRequest); i {
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
	file_payment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}
