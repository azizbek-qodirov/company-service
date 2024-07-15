// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: staffer-protos/position.proto

package genprotos

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

type PositionCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DepartmentId string `protobuf:"bytes,3,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
}

func (x *PositionCreateReq) Reset() {
	*x = PositionCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCreateReq) ProtoMessage() {}

func (x *PositionCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionCreateReq.ProtoReflect.Descriptor instead.
func (*PositionCreateReq) Descriptor() ([]byte, []int) {
	return file_staffer_protos_position_proto_rawDescGZIP(), []int{0}
}

func (x *PositionCreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PositionCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PositionCreateReq) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

type PositionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DepartmentId string `protobuf:"bytes,4,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
}

func (x *PositionRes) Reset() {
	*x = PositionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionRes) ProtoMessage() {}

func (x *PositionRes) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionRes.ProtoReflect.Descriptor instead.
func (*PositionRes) Descriptor() ([]byte, []int) {
	return file_staffer_protos_position_proto_rawDescGZIP(), []int{1}
}

func (x *PositionRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PositionRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PositionRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PositionRes) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

type PositionUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PositionUpdateReq) Reset() {
	*x = PositionUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionUpdateReq) ProtoMessage() {}

func (x *PositionUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionUpdateReq.ProtoReflect.Descriptor instead.
func (*PositionUpdateReq) Descriptor() ([]byte, []int) {
	return file_staffer_protos_position_proto_rawDescGZIP(), []int{2}
}

func (x *PositionUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PositionUpdateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PositionUpdateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PositionGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepartmentId string      `protobuf:"bytes,1,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	Filter       *Pagination `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *PositionGetAllReq) Reset() {
	*x = PositionGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionGetAllReq) ProtoMessage() {}

func (x *PositionGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionGetAllReq.ProtoReflect.Descriptor instead.
func (*PositionGetAllReq) Descriptor() ([]byte, []int) {
	return file_staffer_protos_position_proto_rawDescGZIP(), []int{3}
}

func (x *PositionGetAllReq) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *PositionGetAllReq) GetFilter() *Pagination {
	if x != nil {
		return x.Filter
	}
	return nil
}

type PositionGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *PositionRes `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *PositionGetByIdRes) Reset() {
	*x = PositionGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionGetByIdRes) ProtoMessage() {}

func (x *PositionGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_position_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionGetByIdRes.ProtoReflect.Descriptor instead.
func (*PositionGetByIdRes) Descriptor() ([]byte, []int) {
	return file_staffer_protos_position_proto_rawDescGZIP(), []int{4}
}

func (x *PositionGetByIdRes) GetPosition() *PositionRes {
	if x != nil {
		return x.Position
	}
	return nil
}

type PositionGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*PositionRes `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *PositionGetAllRes) Reset() {
	*x = PositionGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_position_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionGetAllRes) ProtoMessage() {}

func (x *PositionGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_position_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionGetAllRes.ProtoReflect.Descriptor instead.
func (*PositionGetAllRes) Descriptor() ([]byte, []int) {
	return file_staffer_protos_position_proto_rawDescGZIP(), []int{5}
}

func (x *PositionGetAllRes) GetPositions() []*PositionRes {
	if x != nil {
		return x.Positions
	}
	return nil
}

var File_staffer_protos_position_proto protoreflect.FileDescriptor

var file_staffer_protos_position_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x61, 0x66, 0x66, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x19, 0x73, 0x74, 0x61, 0x66, 0x66, 0x65, 0x72, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x70, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x5b, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x63, 0x0a, 0x11,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x44, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x96,
	0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x42, 0x79,
	0x69, 0x64, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staffer_protos_position_proto_rawDescOnce sync.Once
	file_staffer_protos_position_proto_rawDescData = file_staffer_protos_position_proto_rawDesc
)

func file_staffer_protos_position_proto_rawDescGZIP() []byte {
	file_staffer_protos_position_proto_rawDescOnce.Do(func() {
		file_staffer_protos_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_staffer_protos_position_proto_rawDescData)
	})
	return file_staffer_protos_position_proto_rawDescData
}

var file_staffer_protos_position_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_staffer_protos_position_proto_goTypes = []any{
	(*PositionCreateReq)(nil),  // 0: staff.PositionCreateReq
	(*PositionRes)(nil),        // 1: staff.PositionRes
	(*PositionUpdateReq)(nil),  // 2: staff.PositionUpdateReq
	(*PositionGetAllReq)(nil),  // 3: staff.PositionGetAllReq
	(*PositionGetByIdRes)(nil), // 4: staff.PositionGetByIdRes
	(*PositionGetAllRes)(nil),  // 5: staff.PositionGetAllRes
	(*Pagination)(nil),         // 6: staff.Pagination
	(*Byid)(nil),               // 7: staff.Byid
	(*Void)(nil),               // 8: staff.Void
}
var file_staffer_protos_position_proto_depIdxs = []int32{
	6, // 0: staff.PositionGetAllReq.filter:type_name -> staff.Pagination
	1, // 1: staff.PositionGetByIdRes.position:type_name -> staff.PositionRes
	1, // 2: staff.PositionGetAllRes.positions:type_name -> staff.PositionRes
	0, // 3: staff.PositionService.Create:input_type -> staff.PositionCreateReq
	7, // 4: staff.PositionService.GetById:input_type -> staff.Byid
	2, // 5: staff.PositionService.Update:input_type -> staff.PositionUpdateReq
	7, // 6: staff.PositionService.Delete:input_type -> staff.Byid
	3, // 7: staff.PositionService.GetAll:input_type -> staff.PositionGetAllReq
	1, // 8: staff.PositionService.Create:output_type -> staff.PositionRes
	4, // 9: staff.PositionService.GetById:output_type -> staff.PositionGetByIdRes
	1, // 10: staff.PositionService.Update:output_type -> staff.PositionRes
	8, // 11: staff.PositionService.Delete:output_type -> staff.Void
	5, // 12: staff.PositionService.GetAll:output_type -> staff.PositionGetAllRes
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_staffer_protos_position_proto_init() }
func file_staffer_protos_position_proto_init() {
	if File_staffer_protos_position_proto != nil {
		return
	}
	file_staffer_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staffer_protos_position_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PositionCreateReq); i {
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
		file_staffer_protos_position_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PositionRes); i {
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
		file_staffer_protos_position_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PositionUpdateReq); i {
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
		file_staffer_protos_position_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PositionGetAllReq); i {
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
		file_staffer_protos_position_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PositionGetByIdRes); i {
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
		file_staffer_protos_position_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PositionGetAllRes); i {
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
			RawDescriptor: file_staffer_protos_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staffer_protos_position_proto_goTypes,
		DependencyIndexes: file_staffer_protos_position_proto_depIdxs,
		MessageInfos:      file_staffer_protos_position_proto_msgTypes,
	}.Build()
	File_staffer_protos_position_proto = out.File
	file_staffer_protos_position_proto_rawDesc = nil
	file_staffer_protos_position_proto_goTypes = nil
	file_staffer_protos_position_proto_depIdxs = nil
}
