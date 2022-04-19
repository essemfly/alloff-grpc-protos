// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: topbanner.proto

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

type GetTopBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId string `protobuf:"bytes,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
}

func (x *GetTopBannerRequest) Reset() {
	*x = GetTopBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopBannerRequest) ProtoMessage() {}

func (x *GetTopBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopBannerRequest.ProtoReflect.Descriptor instead.
func (*GetTopBannerRequest) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{0}
}

func (x *GetTopBannerRequest) GetBannerId() string {
	if x != nil {
		return x.BannerId
	}
	return ""
}

type ListTopBannersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListTopBannersRequest) Reset() {
	*x = ListTopBannersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopBannersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopBannersRequest) ProtoMessage() {}

func (x *ListTopBannersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopBannersRequest.ProtoReflect.Descriptor instead.
func (*ListTopBannersRequest) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{1}
}

func (x *ListTopBannersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTopBannersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type EditTopBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId     string  `protobuf:"bytes,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	BannerImage  *string `protobuf:"bytes,2,opt,name=banner_image,json=bannerImage,proto3,oneof" json:"banner_image,omitempty"`
	ExhibitionId *string `protobuf:"bytes,3,opt,name=exhibition_id,json=exhibitionId,proto3,oneof" json:"exhibition_id,omitempty"`
	Title        *string `protobuf:"bytes,4,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Subtitle     *string `protobuf:"bytes,5,opt,name=subtitle,proto3,oneof" json:"subtitle,omitempty"`
	IsLive       *bool   `protobuf:"varint,6,opt,name=is_live,json=isLive,proto3,oneof" json:"is_live,omitempty"`
	Weight       *int32  `protobuf:"varint,7,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *EditTopBannerRequest) Reset() {
	*x = EditTopBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditTopBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditTopBannerRequest) ProtoMessage() {}

func (x *EditTopBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditTopBannerRequest.ProtoReflect.Descriptor instead.
func (*EditTopBannerRequest) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{2}
}

func (x *EditTopBannerRequest) GetBannerId() string {
	if x != nil {
		return x.BannerId
	}
	return ""
}

func (x *EditTopBannerRequest) GetBannerImage() string {
	if x != nil && x.BannerImage != nil {
		return *x.BannerImage
	}
	return ""
}

func (x *EditTopBannerRequest) GetExhibitionId() string {
	if x != nil && x.ExhibitionId != nil {
		return *x.ExhibitionId
	}
	return ""
}

func (x *EditTopBannerRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *EditTopBannerRequest) GetSubtitle() string {
	if x != nil && x.Subtitle != nil {
		return *x.Subtitle
	}
	return ""
}

func (x *EditTopBannerRequest) GetIsLive() bool {
	if x != nil && x.IsLive != nil {
		return *x.IsLive
	}
	return false
}

func (x *EditTopBannerRequest) GetWeight() int32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type CreateTopBannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerImage  string `protobuf:"bytes,1,opt,name=banner_image,json=bannerImage,proto3" json:"banner_image,omitempty"`
	ExhibitionId string `protobuf:"bytes,2,opt,name=exhibition_id,json=exhibitionId,proto3" json:"exhibition_id,omitempty"`
	Title        string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle     string `protobuf:"bytes,4,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	IsLive       bool   `protobuf:"varint,5,opt,name=is_live,json=isLive,proto3" json:"is_live,omitempty"`
	Weight       int32  `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *CreateTopBannerRequest) Reset() {
	*x = CreateTopBannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopBannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopBannerRequest) ProtoMessage() {}

func (x *CreateTopBannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopBannerRequest.ProtoReflect.Descriptor instead.
func (*CreateTopBannerRequest) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTopBannerRequest) GetBannerImage() string {
	if x != nil {
		return x.BannerImage
	}
	return ""
}

func (x *CreateTopBannerRequest) GetExhibitionId() string {
	if x != nil {
		return x.ExhibitionId
	}
	return ""
}

func (x *CreateTopBannerRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTopBannerRequest) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *CreateTopBannerRequest) GetIsLive() bool {
	if x != nil {
		return x.IsLive
	}
	return false
}

func (x *CreateTopBannerRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type GetTopBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banner *TopBannerMessage `protobuf:"bytes,1,opt,name=banner,proto3" json:"banner,omitempty"`
}

func (x *GetTopBannerResponse) Reset() {
	*x = GetTopBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopBannerResponse) ProtoMessage() {}

func (x *GetTopBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopBannerResponse.ProtoReflect.Descriptor instead.
func (*GetTopBannerResponse) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopBannerResponse) GetBanner() *TopBannerMessage {
	if x != nil {
		return x.Banner
	}
	return nil
}

type ListTopBannersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banners     []*TopBannerMessage `protobuf:"bytes,1,rep,name=banners,proto3" json:"banners,omitempty"`
	Offset      int32               `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit       int32               `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	TotalCounts int32               `protobuf:"varint,4,opt,name=total_counts,json=totalCounts,proto3" json:"total_counts,omitempty"`
}

func (x *ListTopBannersResponse) Reset() {
	*x = ListTopBannersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopBannersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopBannersResponse) ProtoMessage() {}

func (x *ListTopBannersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopBannersResponse.ProtoReflect.Descriptor instead.
func (*ListTopBannersResponse) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{5}
}

func (x *ListTopBannersResponse) GetBanners() []*TopBannerMessage {
	if x != nil {
		return x.Banners
	}
	return nil
}

func (x *ListTopBannersResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTopBannersResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTopBannersResponse) GetTotalCounts() int32 {
	if x != nil {
		return x.TotalCounts
	}
	return 0
}

type EditTopBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banner *TopBannerMessage `protobuf:"bytes,1,opt,name=banner,proto3" json:"banner,omitempty"`
}

func (x *EditTopBannerResponse) Reset() {
	*x = EditTopBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditTopBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditTopBannerResponse) ProtoMessage() {}

func (x *EditTopBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditTopBannerResponse.ProtoReflect.Descriptor instead.
func (*EditTopBannerResponse) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{6}
}

func (x *EditTopBannerResponse) GetBanner() *TopBannerMessage {
	if x != nil {
		return x.Banner
	}
	return nil
}

type CreateTopBannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banner *TopBannerMessage `protobuf:"bytes,1,opt,name=banner,proto3" json:"banner,omitempty"`
}

func (x *CreateTopBannerResponse) Reset() {
	*x = CreateTopBannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopBannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopBannerResponse) ProtoMessage() {}

func (x *CreateTopBannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopBannerResponse.ProtoReflect.Descriptor instead.
func (*CreateTopBannerResponse) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTopBannerResponse) GetBanner() *TopBannerMessage {
	if x != nil {
		return x.Banner
	}
	return nil
}

type TopBannerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BannerId     string `protobuf:"bytes,1,opt,name=banner_id,json=bannerId,proto3" json:"banner_id,omitempty"`
	BannerImage  string `protobuf:"bytes,2,opt,name=banner_image,json=bannerImage,proto3" json:"banner_image,omitempty"`
	ExhibitionId string `protobuf:"bytes,3,opt,name=exhibition_id,json=exhibitionId,proto3" json:"exhibition_id,omitempty"`
	Title        string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle     string `protobuf:"bytes,5,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	IsLive       bool   `protobuf:"varint,6,opt,name=is_live,json=isLive,proto3" json:"is_live,omitempty"`
	Weight       int32  `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *TopBannerMessage) Reset() {
	*x = TopBannerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topbanner_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopBannerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopBannerMessage) ProtoMessage() {}

func (x *TopBannerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_topbanner_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopBannerMessage.ProtoReflect.Descriptor instead.
func (*TopBannerMessage) Descriptor() ([]byte, []int) {
	return file_topbanner_proto_rawDescGZIP(), []int{8}
}

func (x *TopBannerMessage) GetBannerId() string {
	if x != nil {
		return x.BannerId
	}
	return ""
}

func (x *TopBannerMessage) GetBannerImage() string {
	if x != nil {
		return x.BannerImage
	}
	return ""
}

func (x *TopBannerMessage) GetExhibitionId() string {
	if x != nil {
		return x.ExhibitionId
	}
	return ""
}

func (x *TopBannerMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TopBannerMessage) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *TopBannerMessage) GetIsLive() bool {
	if x != nil {
		return x.IsLive
	}
	return false
}

func (x *TopBannerMessage) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_topbanner_proto protoreflect.FileDescriptor

var file_topbanner_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x70, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x22, 0x32, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x45, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xcd, 0x02, 0x0a, 0x14, 0x45, 0x64, 0x69, 0x74, 0x54,
	0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0c,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x65,
	0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x06, 0x69,
	0x73, 0x4c, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x78, 0x68, 0x69, 0x62,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x68,
	0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x4c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x4a, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e,
	0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x9f, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e,
	0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x45, 0x64,
	0x69, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x54,
	0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x54, 0x6f,
	0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0xda, 0x01, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x32, 0xd9, 0x02, 0x0a, 0x09, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x54, 0x6f, 0x70,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66,
	0x66, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66,
	0x66, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x67, 0x6f, 0x61,
	0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67,
	0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x70, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_topbanner_proto_rawDescOnce sync.Once
	file_topbanner_proto_rawDescData = file_topbanner_proto_rawDesc
)

func file_topbanner_proto_rawDescGZIP() []byte {
	file_topbanner_proto_rawDescOnce.Do(func() {
		file_topbanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_topbanner_proto_rawDescData)
	})
	return file_topbanner_proto_rawDescData
}

var file_topbanner_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_topbanner_proto_goTypes = []interface{}{
	(*GetTopBannerRequest)(nil),     // 0: goalloff.GetTopBannerRequest
	(*ListTopBannersRequest)(nil),   // 1: goalloff.ListTopBannersRequest
	(*EditTopBannerRequest)(nil),    // 2: goalloff.EditTopBannerRequest
	(*CreateTopBannerRequest)(nil),  // 3: goalloff.CreateTopBannerRequest
	(*GetTopBannerResponse)(nil),    // 4: goalloff.GetTopBannerResponse
	(*ListTopBannersResponse)(nil),  // 5: goalloff.ListTopBannersResponse
	(*EditTopBannerResponse)(nil),   // 6: goalloff.EditTopBannerResponse
	(*CreateTopBannerResponse)(nil), // 7: goalloff.CreateTopBannerResponse
	(*TopBannerMessage)(nil),        // 8: goalloff.TopBannerMessage
}
var file_topbanner_proto_depIdxs = []int32{
	8, // 0: goalloff.GetTopBannerResponse.banner:type_name -> goalloff.TopBannerMessage
	8, // 1: goalloff.ListTopBannersResponse.banners:type_name -> goalloff.TopBannerMessage
	8, // 2: goalloff.EditTopBannerResponse.banner:type_name -> goalloff.TopBannerMessage
	8, // 3: goalloff.CreateTopBannerResponse.banner:type_name -> goalloff.TopBannerMessage
	0, // 4: goalloff.TopBanner.GetTopBanner:input_type -> goalloff.GetTopBannerRequest
	1, // 5: goalloff.TopBanner.ListTopBanners:input_type -> goalloff.ListTopBannersRequest
	2, // 6: goalloff.TopBanner.EditTopBanner:input_type -> goalloff.EditTopBannerRequest
	3, // 7: goalloff.TopBanner.CreateTopBanner:input_type -> goalloff.CreateTopBannerRequest
	4, // 8: goalloff.TopBanner.GetTopBanner:output_type -> goalloff.GetTopBannerResponse
	5, // 9: goalloff.TopBanner.ListTopBanners:output_type -> goalloff.ListTopBannersResponse
	6, // 10: goalloff.TopBanner.EditTopBanner:output_type -> goalloff.EditTopBannerResponse
	7, // 11: goalloff.TopBanner.CreateTopBanner:output_type -> goalloff.CreateTopBannerResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_topbanner_proto_init() }
func file_topbanner_proto_init() {
	if File_topbanner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_topbanner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopBannerRequest); i {
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
		file_topbanner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopBannersRequest); i {
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
		file_topbanner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditTopBannerRequest); i {
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
		file_topbanner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopBannerRequest); i {
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
		file_topbanner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopBannerResponse); i {
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
		file_topbanner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopBannersResponse); i {
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
		file_topbanner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditTopBannerResponse); i {
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
		file_topbanner_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTopBannerResponse); i {
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
		file_topbanner_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopBannerMessage); i {
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
	file_topbanner_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_topbanner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_topbanner_proto_goTypes,
		DependencyIndexes: file_topbanner_proto_depIdxs,
		MessageInfos:      file_topbanner_proto_msgTypes,
	}.Build()
	File_topbanner_proto = out.File
	file_topbanner_proto_rawDesc = nil
	file_topbanner_proto_goTypes = nil
	file_topbanner_proto_depIdxs = nil
}
