// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: errorpb/errorpb.proto

package errorpb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	metapb "tinykv/proto/metapb"
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

type NotLeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId uint64       `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Leader   *metapb.Peer `protobuf:"bytes,2,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *NotLeader) Reset() {
	*x = NotLeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotLeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotLeader) ProtoMessage() {}

func (x *NotLeader) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotLeader.ProtoReflect.Descriptor instead.
func (*NotLeader) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{0}
}

func (x *NotLeader) GetRegionId() uint64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *NotLeader) GetLeader() *metapb.Peer {
	if x != nil {
		return x.Leader
	}
	return nil
}

type StoreNotMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestStoreId uint64 `protobuf:"varint,1,opt,name=request_store_id,json=requestStoreId,proto3" json:"request_store_id,omitempty"`
	ActualStoreId  uint64 `protobuf:"varint,2,opt,name=actual_store_id,json=actualStoreId,proto3" json:"actual_store_id,omitempty"`
}

func (x *StoreNotMatch) Reset() {
	*x = StoreNotMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreNotMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreNotMatch) ProtoMessage() {}

func (x *StoreNotMatch) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreNotMatch.ProtoReflect.Descriptor instead.
func (*StoreNotMatch) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{1}
}

func (x *StoreNotMatch) GetRequestStoreId() uint64 {
	if x != nil {
		return x.RequestStoreId
	}
	return 0
}

func (x *StoreNotMatch) GetActualStoreId() uint64 {
	if x != nil {
		return x.ActualStoreId
	}
	return 0
}

type RegionNotFound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId uint64 `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *RegionNotFound) Reset() {
	*x = RegionNotFound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionNotFound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionNotFound) ProtoMessage() {}

func (x *RegionNotFound) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionNotFound.ProtoReflect.Descriptor instead.
func (*RegionNotFound) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{2}
}

func (x *RegionNotFound) GetRegionId() uint64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

type KeyNotInRegion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RegionId uint64 `protobuf:"varint,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	StartKey []byte `protobuf:"bytes,3,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey   []byte `protobuf:"bytes,4,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
}

func (x *KeyNotInRegion) Reset() {
	*x = KeyNotInRegion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyNotInRegion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyNotInRegion) ProtoMessage() {}

func (x *KeyNotInRegion) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyNotInRegion.ProtoReflect.Descriptor instead.
func (*KeyNotInRegion) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{3}
}

func (x *KeyNotInRegion) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyNotInRegion) GetRegionId() uint64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *KeyNotInRegion) GetStartKey() []byte {
	if x != nil {
		return x.StartKey
	}
	return nil
}

func (x *KeyNotInRegion) GetEndKey() []byte {
	if x != nil {
		return x.EndKey
	}
	return nil
}

type EpochNotMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentRegions []*metapb.Region `protobuf:"bytes,1,rep,name=current_regions,json=currentRegions,proto3" json:"current_regions,omitempty"`
}

func (x *EpochNotMatch) Reset() {
	*x = EpochNotMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochNotMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochNotMatch) ProtoMessage() {}

func (x *EpochNotMatch) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochNotMatch.ProtoReflect.Descriptor instead.
func (*EpochNotMatch) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{4}
}

func (x *EpochNotMatch) GetCurrentRegions() []*metapb.Region {
	if x != nil {
		return x.CurrentRegions
	}
	return nil
}

type StaleCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StaleCommand) Reset() {
	*x = StaleCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaleCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaleCommand) ProtoMessage() {}

func (x *StaleCommand) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaleCommand.ProtoReflect.Descriptor instead.
func (*StaleCommand) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{5}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	NotLeader      *NotLeader      `protobuf:"bytes,2,opt,name=not_leader,json=notLeader,proto3" json:"not_leader,omitempty"`
	RegionNotFound *RegionNotFound `protobuf:"bytes,3,opt,name=region_not_found,json=regionNotFound,proto3" json:"region_not_found,omitempty"`
	KeyNotInRegion *KeyNotInRegion `protobuf:"bytes,4,opt,name=key_not_in_region,json=keyNotInRegion,proto3" json:"key_not_in_region,omitempty"`
	EpochNotMatch  *EpochNotMatch  `protobuf:"bytes,5,opt,name=epoch_not_match,json=epochNotMatch,proto3" json:"epoch_not_match,omitempty"`
	StaleCommand   *StaleCommand   `protobuf:"bytes,7,opt,name=stale_command,json=staleCommand,proto3" json:"stale_command,omitempty"`
	StoreNotMatch  *StoreNotMatch  `protobuf:"bytes,8,opt,name=store_not_match,json=storeNotMatch,proto3" json:"store_not_match,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorpb_errorpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errorpb_errorpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errorpb_errorpb_proto_rawDescGZIP(), []int{6}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetNotLeader() *NotLeader {
	if x != nil {
		return x.NotLeader
	}
	return nil
}

func (x *Error) GetRegionNotFound() *RegionNotFound {
	if x != nil {
		return x.RegionNotFound
	}
	return nil
}

func (x *Error) GetKeyNotInRegion() *KeyNotInRegion {
	if x != nil {
		return x.KeyNotInRegion
	}
	return nil
}

func (x *Error) GetEpochNotMatch() *EpochNotMatch {
	if x != nil {
		return x.EpochNotMatch
	}
	return nil
}

func (x *Error) GetStaleCommand() *StaleCommand {
	if x != nil {
		return x.StaleCommand
	}
	return nil
}

func (x *Error) GetStoreNotMatch() *StoreNotMatch {
	if x != nil {
		return x.StoreNotMatch
	}
	return nil
}

var File_errorpb_errorpb_proto protoreflect.FileDescriptor

var file_errorpb_errorpb_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62,
	0x1a, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x06, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x61, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x75, 0x61,
	0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x4e, 0x6f,
	0x74, 0x49, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x22, 0x48,
	0x0a, 0x0d, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x37, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xa4, 0x03, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x0a,
	0x6e, 0x6f, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x41, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x42, 0x0a, 0x11, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x4e, 0x6f, 0x74, 0x49, 0x6e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6b, 0x65, 0x79, 0x4e, 0x6f, 0x74, 0x49, 0x6e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x4e,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0d, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x6f,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3a, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x42,
	0x16, 0x5a, 0x14, 0x74, 0x69, 0x6e, 0x79, 0x6b, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errorpb_errorpb_proto_rawDescOnce sync.Once
	file_errorpb_errorpb_proto_rawDescData = file_errorpb_errorpb_proto_rawDesc
)

func file_errorpb_errorpb_proto_rawDescGZIP() []byte {
	file_errorpb_errorpb_proto_rawDescOnce.Do(func() {
		file_errorpb_errorpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorpb_errorpb_proto_rawDescData)
	})
	return file_errorpb_errorpb_proto_rawDescData
}

var file_errorpb_errorpb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_errorpb_errorpb_proto_goTypes = []interface{}{
	(*NotLeader)(nil),      // 0: errorpb.NotLeader
	(*StoreNotMatch)(nil),  // 1: errorpb.StoreNotMatch
	(*RegionNotFound)(nil), // 2: errorpb.RegionNotFound
	(*KeyNotInRegion)(nil), // 3: errorpb.KeyNotInRegion
	(*EpochNotMatch)(nil),  // 4: errorpb.EpochNotMatch
	(*StaleCommand)(nil),   // 5: errorpb.StaleCommand
	(*Error)(nil),          // 6: errorpb.Error
	(*metapb.Peer)(nil),    // 7: metapb.Peer
	(*metapb.Region)(nil),  // 8: metapb.Region
}
var file_errorpb_errorpb_proto_depIdxs = []int32{
	7, // 0: errorpb.NotLeader.leader:type_name -> metapb.Peer
	8, // 1: errorpb.EpochNotMatch.current_regions:type_name -> metapb.Region
	0, // 2: errorpb.Error.not_leader:type_name -> errorpb.NotLeader
	2, // 3: errorpb.Error.region_not_found:type_name -> errorpb.RegionNotFound
	3, // 4: errorpb.Error.key_not_in_region:type_name -> errorpb.KeyNotInRegion
	4, // 5: errorpb.Error.epoch_not_match:type_name -> errorpb.EpochNotMatch
	5, // 6: errorpb.Error.stale_command:type_name -> errorpb.StaleCommand
	1, // 7: errorpb.Error.store_not_match:type_name -> errorpb.StoreNotMatch
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_errorpb_errorpb_proto_init() }
func file_errorpb_errorpb_proto_init() {
	if File_errorpb_errorpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errorpb_errorpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotLeader); i {
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
		file_errorpb_errorpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreNotMatch); i {
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
		file_errorpb_errorpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionNotFound); i {
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
		file_errorpb_errorpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyNotInRegion); i {
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
		file_errorpb_errorpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochNotMatch); i {
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
		file_errorpb_errorpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaleCommand); i {
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
		file_errorpb_errorpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_errorpb_errorpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errorpb_errorpb_proto_goTypes,
		DependencyIndexes: file_errorpb_errorpb_proto_depIdxs,
		MessageInfos:      file_errorpb_errorpb_proto_msgTypes,
	}.Build()
	File_errorpb_errorpb_proto = out.File
	file_errorpb_errorpb_proto_rawDesc = nil
	file_errorpb_errorpb_proto_goTypes = nil
	file_errorpb_errorpb_proto_depIdxs = nil
}
