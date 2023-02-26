// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: estate.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_estate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_estate_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type Estate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Location  *Location              `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Estate) Reset() {
	*x = Estate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estate) ProtoMessage() {}

func (x *Estate) ProtoReflect() protoreflect.Message {
	mi := &file_estate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estate.ProtoReflect.Descriptor instead.
func (*Estate) Descriptor() ([]byte, []int) {
	return file_estate_proto_rawDescGZIP(), []int{1}
}

func (x *Estate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Estate) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Estate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Estate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Estate) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type CreateEstateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateEstateRequest) Reset() {
	*x = CreateEstateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEstateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEstateRequest) ProtoMessage() {}

func (x *CreateEstateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_estate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEstateRequest.ProtoReflect.Descriptor instead.
func (*CreateEstateRequest) Descriptor() ([]byte, []int) {
	return file_estate_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEstateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEstateRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type CreateEstateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estate *Estate `protobuf:"bytes,1,opt,name=estate,proto3" json:"estate,omitempty"`
}

func (x *CreateEstateResponse) Reset() {
	*x = CreateEstateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEstateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEstateResponse) ProtoMessage() {}

func (x *CreateEstateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_estate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEstateResponse.ProtoReflect.Descriptor instead.
func (*CreateEstateResponse) Descriptor() ([]byte, []int) {
	return file_estate_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEstateResponse) GetEstate() *Estate {
	if x != nil {
		return x.Estate
	}
	return nil
}

type GetEstatesByBoundBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopLeft     *Location `protobuf:"bytes,1,opt,name=top_left,json=topLeft,proto3" json:"top_left,omitempty"`
	BottomRight *Location `protobuf:"bytes,2,opt,name=bottom_right,json=bottomRight,proto3" json:"bottom_right,omitempty"`
}

func (x *GetEstatesByBoundBoxRequest) Reset() {
	*x = GetEstatesByBoundBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEstatesByBoundBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEstatesByBoundBoxRequest) ProtoMessage() {}

func (x *GetEstatesByBoundBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_estate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEstatesByBoundBoxRequest.ProtoReflect.Descriptor instead.
func (*GetEstatesByBoundBoxRequest) Descriptor() ([]byte, []int) {
	return file_estate_proto_rawDescGZIP(), []int{4}
}

func (x *GetEstatesByBoundBoxRequest) GetTopLeft() *Location {
	if x != nil {
		return x.TopLeft
	}
	return nil
}

func (x *GetEstatesByBoundBoxRequest) GetBottomRight() *Location {
	if x != nil {
		return x.BottomRight
	}
	return nil
}

type GetEstatesByBoundBoxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estates []*Estate `protobuf:"bytes,1,rep,name=estates,proto3" json:"estates,omitempty"`
	Hits    int32     `protobuf:"varint,2,opt,name=hits,proto3" json:"hits,omitempty"`
}

func (x *GetEstatesByBoundBoxResponse) Reset() {
	*x = GetEstatesByBoundBoxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_estate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEstatesByBoundBoxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEstatesByBoundBoxResponse) ProtoMessage() {}

func (x *GetEstatesByBoundBoxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_estate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEstatesByBoundBoxResponse.ProtoReflect.Descriptor instead.
func (*GetEstatesByBoundBoxResponse) Descriptor() ([]byte, []int) {
	return file_estate_proto_rawDescGZIP(), []int{5}
}

func (x *GetEstatesByBoundBoxResponse) GetEstates() []*Estate {
	if x != nil {
		return x.Estates
	}
	return nil
}

func (x *GetEstatesByBoundBoxResponse) GetHits() int32 {
	if x != nil {
		return x.Hits
	}
	return 0
}

var File_estate_proto protoreflect.FileDescriptor

var file_estate_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0xfc, 0x01, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x58, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3d, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x7d,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x42, 0x79, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x08, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x74, 0x6f, 0x70, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x32, 0x0a, 0x0c, 0x62, 0x6f, 0x74,
	0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x52, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5b, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x42, 0x79, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x07, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x32, 0xb5, 0x01, 0x0a, 0x0d, 0x45,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x42, 0x79, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x42, 0x79, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_estate_proto_rawDescOnce sync.Once
	file_estate_proto_rawDescData = file_estate_proto_rawDesc
)

func file_estate_proto_rawDescGZIP() []byte {
	file_estate_proto_rawDescOnce.Do(func() {
		file_estate_proto_rawDescData = protoimpl.X.CompressGZIP(file_estate_proto_rawDescData)
	})
	return file_estate_proto_rawDescData
}

var file_estate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_estate_proto_goTypes = []interface{}{
	(*Location)(nil),                     // 0: proto.Location
	(*Estate)(nil),                       // 1: proto.Estate
	(*CreateEstateRequest)(nil),          // 2: proto.CreateEstateRequest
	(*CreateEstateResponse)(nil),         // 3: proto.CreateEstateResponse
	(*GetEstatesByBoundBoxRequest)(nil),  // 4: proto.GetEstatesByBoundBoxRequest
	(*GetEstatesByBoundBoxResponse)(nil), // 5: proto.GetEstatesByBoundBoxResponse
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
}
var file_estate_proto_depIdxs = []int32{
	0,  // 0: proto.Estate.location:type_name -> proto.Location
	6,  // 1: proto.Estate.created_at:type_name -> google.protobuf.Timestamp
	6,  // 2: proto.Estate.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 3: proto.Estate.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 4: proto.CreateEstateRequest.location:type_name -> proto.Location
	1,  // 5: proto.CreateEstateResponse.estate:type_name -> proto.Estate
	0,  // 6: proto.GetEstatesByBoundBoxRequest.top_left:type_name -> proto.Location
	0,  // 7: proto.GetEstatesByBoundBoxRequest.bottom_right:type_name -> proto.Location
	1,  // 8: proto.GetEstatesByBoundBoxResponse.estates:type_name -> proto.Estate
	2,  // 9: proto.EstateService.CreateEstate:input_type -> proto.CreateEstateRequest
	4,  // 10: proto.EstateService.GetEstateByBound:input_type -> proto.GetEstatesByBoundBoxRequest
	3,  // 11: proto.EstateService.CreateEstate:output_type -> proto.CreateEstateResponse
	5,  // 12: proto.EstateService.GetEstateByBound:output_type -> proto.GetEstatesByBoundBoxResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_estate_proto_init() }
func file_estate_proto_init() {
	if File_estate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_estate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_estate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estate); i {
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
		file_estate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEstateRequest); i {
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
		file_estate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEstateResponse); i {
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
		file_estate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEstatesByBoundBoxRequest); i {
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
		file_estate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEstatesByBoundBoxResponse); i {
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
			RawDescriptor: file_estate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_estate_proto_goTypes,
		DependencyIndexes: file_estate_proto_depIdxs,
		MessageInfos:      file_estate_proto_msgTypes,
	}.Build()
	File_estate_proto = out.File
	file_estate_proto_rawDesc = nil
	file_estate_proto_goTypes = nil
	file_estate_proto_depIdxs = nil
}
