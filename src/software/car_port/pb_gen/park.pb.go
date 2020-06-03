// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: park.proto

package pb_gen

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Park struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                   //停车场id
	Name       string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                //停车场名字
	Longitude  float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`                    //停车场经度
	Latitude   float64 `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`                      //停车场维度
	Charge     int32   `protobuf:"varint,5,opt,name=charge,proto3" json:"charge,omitempty"`                           //停车场费用
	EmptyPorts int32   `protobuf:"varint,6,opt,name=empty_ports,json=emptyPorts,proto3" json:"empty_ports,omitempty"` //空余车位
	TotalPorts int32   `protobuf:"varint,7,opt,name=total_ports,json=totalPorts,proto3" json:"total_ports,omitempty"` //总车位
}

func (x *Park) Reset() {
	*x = Park{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Park) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Park) ProtoMessage() {}

func (x *Park) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Park.ProtoReflect.Descriptor instead.
func (*Park) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{0}
}

func (x *Park) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Park) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Park) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Park) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Park) GetCharge() int32 {
	if x != nil {
		return x.Charge
	}
	return 0
}

func (x *Park) GetEmptyPorts() int32 {
	if x != nil {
		return x.EmptyPorts
	}
	return 0
}

func (x *Park) GetTotalPorts() int32 {
	if x != nil {
		return x.TotalPorts
	}
	return 0
}

type ReqAddPark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Park *Park `protobuf:"bytes,1,opt,name=park,proto3" json:"park,omitempty"`
}

func (x *ReqAddPark) Reset() {
	*x = ReqAddPark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqAddPark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqAddPark) ProtoMessage() {}

func (x *ReqAddPark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqAddPark.ProtoReflect.Descriptor instead.
func (*ReqAddPark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{1}
}

func (x *ReqAddPark) GetPark() *Park {
	if x != nil {
		return x.Park
	}
	return nil
}

type RespAddPark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *RespAddPark) Reset() {
	*x = RespAddPark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespAddPark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespAddPark) ProtoMessage() {}

func (x *RespAddPark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespAddPark.ProtoReflect.Descriptor instead.
func (*RespAddPark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{2}
}

func (x *RespAddPark) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *RespAddPark) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

type ReqUpdatePark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Park *Park `protobuf:"bytes,1,opt,name=park,proto3" json:"park,omitempty"`
}

func (x *ReqUpdatePark) Reset() {
	*x = ReqUpdatePark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdatePark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdatePark) ProtoMessage() {}

func (x *ReqUpdatePark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdatePark.ProtoReflect.Descriptor instead.
func (*ReqUpdatePark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{3}
}

func (x *ReqUpdatePark) GetPark() *Park {
	if x != nil {
		return x.Park
	}
	return nil
}

type RespUpdatePark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
}

func (x *RespUpdatePark) Reset() {
	*x = RespUpdatePark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespUpdatePark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespUpdatePark) ProtoMessage() {}

func (x *RespUpdatePark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespUpdatePark.ProtoReflect.Descriptor instead.
func (*RespUpdatePark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{4}
}

func (x *RespUpdatePark) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *RespUpdatePark) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

type ReqGetPark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParkId int64 `protobuf:"varint,1,opt,name=park_id,json=parkId,proto3" json:"park_id,omitempty"`
}

func (x *ReqGetPark) Reset() {
	*x = ReqGetPark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetPark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetPark) ProtoMessage() {}

func (x *ReqGetPark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetPark.ProtoReflect.Descriptor instead.
func (*ReqGetPark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{5}
}

func (x *ReqGetPark) GetParkId() int64 {
	if x != nil {
		return x.ParkId
	}
	return 0
}

type RespGetPark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
	Park    *Park  `protobuf:"bytes,3,opt,name=park,proto3" json:"park,omitempty"`
}

func (x *RespGetPark) Reset() {
	*x = RespGetPark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespGetPark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespGetPark) ProtoMessage() {}

func (x *RespGetPark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespGetPark.ProtoReflect.Descriptor instead.
func (*RespGetPark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{6}
}

func (x *RespGetPark) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *RespGetPark) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

func (x *RespGetPark) GetPark() *Park {
	if x != nil {
		return x.Park
	}
	return nil
}

type ReqMGetPark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Num   int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ReqMGetPark) Reset() {
	*x = ReqMGetPark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMGetPark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMGetPark) ProtoMessage() {}

func (x *ReqMGetPark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMGetPark.ProtoReflect.Descriptor instead.
func (*ReqMGetPark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{7}
}

func (x *ReqMGetPark) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ReqMGetPark) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type RespMGetPark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo   int32   `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrTips string  `protobuf:"bytes,2,opt,name=err_tips,json=errTips,proto3" json:"err_tips,omitempty"`
	Parks   []*Park `protobuf:"bytes,3,rep,name=Parks,proto3" json:"Parks,omitempty"`
	Count   int32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RespMGetPark) Reset() {
	*x = RespMGetPark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_park_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespMGetPark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespMGetPark) ProtoMessage() {}

func (x *RespMGetPark) ProtoReflect() protoreflect.Message {
	mi := &file_park_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespMGetPark.ProtoReflect.Descriptor instead.
func (*RespMGetPark) Descriptor() ([]byte, []int) {
	return file_park_proto_rawDescGZIP(), []int{8}
}

func (x *RespMGetPark) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *RespMGetPark) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

func (x *RespMGetPark) GetParks() []*Park {
	if x != nil {
		return x.Parks
	}
	return nil
}

func (x *RespMGetPark) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_park_proto protoreflect.FileDescriptor

var file_park_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61,
	0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x41, 0x64,
	0x64, 0x50, 0x61, 0x72, 0x6b, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x6b, 0x52, 0x04, 0x70, 0x61, 0x72, 0x6b, 0x22, 0x3f, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x70, 0x41, 0x64, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f,
	0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x22, 0x33, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x6b, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x52, 0x04, 0x70, 0x61, 0x72, 0x6b, 0x22,
	0x42, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x6b, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f,
	0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54,
	0x69, 0x70, 0x73, 0x22, 0x25, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x6b, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x72, 0x6b, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x0b, 0x52, 0x65,
	0x73, 0x70, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72,
	0x5f, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x61, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x52, 0x04, 0x70, 0x61, 0x72, 0x6b, 0x22,
	0x35, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x4d, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x7c, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x4d, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x72, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x72, 0x72, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x52, 0x05, 0x50, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65,
	0x2f, 0x63, 0x61, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x5f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_park_proto_rawDescOnce sync.Once
	file_park_proto_rawDescData = file_park_proto_rawDesc
)

func file_park_proto_rawDescGZIP() []byte {
	file_park_proto_rawDescOnce.Do(func() {
		file_park_proto_rawDescData = protoimpl.X.CompressGZIP(file_park_proto_rawDescData)
	})
	return file_park_proto_rawDescData
}

var file_park_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_park_proto_goTypes = []interface{}{
	(*Park)(nil),           // 0: car_port.Park
	(*ReqAddPark)(nil),     // 1: car_port.ReqAddPark
	(*RespAddPark)(nil),    // 2: car_port.RespAddPark
	(*ReqUpdatePark)(nil),  // 3: car_port.ReqUpdatePark
	(*RespUpdatePark)(nil), // 4: car_port.RespUpdatePark
	(*ReqGetPark)(nil),     // 5: car_port.ReqGetPark
	(*RespGetPark)(nil),    // 6: car_port.RespGetPark
	(*ReqMGetPark)(nil),    // 7: car_port.ReqMGetPark
	(*RespMGetPark)(nil),   // 8: car_port.RespMGetPark
}
var file_park_proto_depIdxs = []int32{
	0, // 0: car_port.ReqAddPark.park:type_name -> car_port.Park
	0, // 1: car_port.ReqUpdatePark.park:type_name -> car_port.Park
	0, // 2: car_port.RespGetPark.park:type_name -> car_port.Park
	0, // 3: car_port.RespMGetPark.Parks:type_name -> car_port.Park
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_park_proto_init() }
func file_park_proto_init() {
	if File_park_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_park_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Park); i {
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
		file_park_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqAddPark); i {
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
		file_park_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespAddPark); i {
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
		file_park_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdatePark); i {
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
		file_park_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespUpdatePark); i {
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
		file_park_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetPark); i {
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
		file_park_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespGetPark); i {
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
		file_park_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMGetPark); i {
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
		file_park_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespMGetPark); i {
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
			RawDescriptor: file_park_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_park_proto_goTypes,
		DependencyIndexes: file_park_proto_depIdxs,
		MessageInfos:      file_park_proto_msgTypes,
	}.Build()
	File_park_proto = out.File
	file_park_proto_rawDesc = nil
	file_park_proto_goTypes = nil
	file_park_proto_depIdxs = nil
}
