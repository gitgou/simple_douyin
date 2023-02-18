// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: redis.proto

package redisdemo

import (
	context "context"
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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg   string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	ServiceTime int64  `protobuf:"varint,3,opt,name=service_time,json=serviceTime,proto3" json:"service_time,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *BaseResp) GetServiceTime() int64 {
	if x != nil {
		return x.ServiceTime
	}
	return 0
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=expire,proto3" json:"expire,omitempty"` //key expire time, seconds, 0 : persist
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{1}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SetRequest) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{2}
}

func (x *SetResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type GetIncreIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetIncreIdRequest) Reset() {
	*x = GetIncreIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIncreIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIncreIdRequest) ProtoMessage() {}

func (x *GetIncreIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIncreIdRequest.ProtoReflect.Descriptor instead.
func (*GetIncreIdRequest) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{3}
}

func (x *GetIncreIdRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetIncreIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
	Id       int64     `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetIncreIdResponse) Reset() {
	*x = GetIncreIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIncreIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIncreIdResponse) ProtoMessage() {}

func (x *GetIncreIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIncreIdResponse.ProtoReflect.Descriptor instead.
func (*GetIncreIdResponse) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{4}
}

func (x *GetIncreIdResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetIncreIdResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ZSET Increase 调用 API
type ZSETIncreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`               //zset key
	Member    string  `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`         // member key
	Increment float32 `protobuf:"fixed32,3,opt,name=increment,proto3" json:"increment,omitempty"` // added value
}

func (x *ZSETIncreRequest) Reset() {
	*x = ZSETIncreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZSETIncreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZSETIncreRequest) ProtoMessage() {}

func (x *ZSETIncreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZSETIncreRequest.ProtoReflect.Descriptor instead.
func (*ZSETIncreRequest) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{5}
}

func (x *ZSETIncreRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ZSETIncreRequest) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *ZSETIncreRequest) GetIncrement() float32 {
	if x != nil {
		return x.Increment
	}
	return 0
}

type ZSETIncreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
}

func (x *ZSETIncreResponse) Reset() {
	*x = ZSETIncreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZSETIncreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZSETIncreResponse) ProtoMessage() {}

func (x *ZSETIncreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZSETIncreResponse.ProtoReflect.Descriptor instead.
func (*ZSETIncreResponse) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{6}
}

func (x *ZSETIncreResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type ZSETGetMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Member string `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *ZSETGetMemberRequest) Reset() {
	*x = ZSETGetMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZSETGetMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZSETGetMemberRequest) ProtoMessage() {}

func (x *ZSETGetMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZSETGetMemberRequest.ProtoReflect.Descriptor instead.
func (*ZSETGetMemberRequest) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{7}
}

func (x *ZSETGetMemberRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ZSETGetMemberRequest) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

type ZSETGetMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float32 `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ZSETGetMemberResponse) Reset() {
	*x = ZSETGetMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZSETGetMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZSETGetMemberResponse) ProtoMessage() {}

func (x *ZSETGetMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZSETGetMemberResponse.ProtoReflect.Descriptor instead.
func (*ZSETGetMemberResponse) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{8}
}

func (x *ZSETGetMemberResponse) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GetUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserInfoRequest) Reset() {
	*x = GetUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRequest) ProtoMessage() {}

func (x *GetUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`       //key 标记信息
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"` //值
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{10}
}

func (x *UserInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UserInfo) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GetUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo []*UserInfo `protobuf:"bytes,1,rep,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *GetUserInfoResponse) Reset() {
	*x = GetUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redis_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResponse) ProtoMessage() {}

func (x *GetUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redis_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_redis_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserInfoResponse) GetUserInfo() []*UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_redis_proto protoreflect.FileDescriptor

var file_redis_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x22, 0x6d, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x10, 0x5a, 0x53,
	0x45, 0x54, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x11, 0x5a, 0x53, 0x45, 0x54, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x40, 0x0a, 0x14, 0x5a, 0x53, 0x45,
	0x54, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x15, 0x5a,
	0x53, 0x45, 0x54, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x44, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x32, 0xe5, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x69, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x09, 0x5a, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x5a, 0x53, 0x45, 0x54, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x5a, 0x53,
	0x45, 0x54, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x5a, 0x53, 0x65, 0x74, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x5a, 0x53, 0x45, 0x54,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x5a, 0x53, 0x45, 0x54, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x67, 0x6f, 0x75,
	0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x64, 0x65,
	0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redis_proto_rawDescOnce sync.Once
	file_redis_proto_rawDescData = file_redis_proto_rawDesc
)

func file_redis_proto_rawDescGZIP() []byte {
	file_redis_proto_rawDescOnce.Do(func() {
		file_redis_proto_rawDescData = protoimpl.X.CompressGZIP(file_redis_proto_rawDescData)
	})
	return file_redis_proto_rawDescData
}

var file_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_redis_proto_goTypes = []interface{}{
	(*BaseResp)(nil),              // 0: douyin.BaseResp
	(*SetRequest)(nil),            // 1: douyin.SetRequest
	(*SetResponse)(nil),           // 2: douyin.SetResponse
	(*GetIncreIdRequest)(nil),     // 3: douyin.GetIncreIdRequest
	(*GetIncreIdResponse)(nil),    // 4: douyin.GetIncreIdResponse
	(*ZSETIncreRequest)(nil),      // 5: douyin.ZSETIncreRequest
	(*ZSETIncreResponse)(nil),     // 6: douyin.ZSETIncreResponse
	(*ZSETGetMemberRequest)(nil),  // 7: douyin.ZSETGetMemberRequest
	(*ZSETGetMemberResponse)(nil), // 8: douyin.ZSETGetMemberResponse
	(*GetUserInfoRequest)(nil),    // 9: douyin.GetUserInfoRequest
	(*UserInfo)(nil),              // 10: douyin.UserInfo
	(*GetUserInfoResponse)(nil),   // 11: douyin.GetUserInfoResponse
}
var file_redis_proto_depIdxs = []int32{
	0,  // 0: douyin.SetResponse.baseResp:type_name -> douyin.BaseResp
	0,  // 1: douyin.GetIncreIdResponse.baseResp:type_name -> douyin.BaseResp
	0,  // 2: douyin.ZSETIncreResponse.baseResp:type_name -> douyin.BaseResp
	10, // 3: douyin.GetUserInfoResponse.user_info:type_name -> douyin.UserInfo
	1,  // 4: douyin.RedisService.Set:input_type -> douyin.SetRequest
	3,  // 5: douyin.RedisService.GetIncreId:input_type -> douyin.GetIncreIdRequest
	5,  // 6: douyin.RedisService.ZSetIncre:input_type -> douyin.ZSETIncreRequest
	7,  // 7: douyin.RedisService.ZSetGetMember:input_type -> douyin.ZSETGetMemberRequest
	9,  // 8: douyin.RedisService.GetUserInfo:input_type -> douyin.GetUserInfoRequest
	2,  // 9: douyin.RedisService.Set:output_type -> douyin.SetResponse
	4,  // 10: douyin.RedisService.GetIncreId:output_type -> douyin.GetIncreIdResponse
	6,  // 11: douyin.RedisService.ZSetIncre:output_type -> douyin.ZSETIncreResponse
	8,  // 12: douyin.RedisService.ZSetGetMember:output_type -> douyin.ZSETGetMemberResponse
	11, // 13: douyin.RedisService.GetUserInfo:output_type -> douyin.GetUserInfoResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_redis_proto_init() }
func file_redis_proto_init() {
	if File_redis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_redis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_redis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_redis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIncreIdRequest); i {
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
		file_redis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIncreIdResponse); i {
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
		file_redis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZSETIncreRequest); i {
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
		file_redis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZSETIncreResponse); i {
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
		file_redis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZSETGetMemberRequest); i {
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
		file_redis_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZSETGetMemberResponse); i {
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
		file_redis_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRequest); i {
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
		file_redis_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_redis_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResponse); i {
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
			RawDescriptor: file_redis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redis_proto_goTypes,
		DependencyIndexes: file_redis_proto_depIdxs,
		MessageInfos:      file_redis_proto_msgTypes,
	}.Build()
	File_redis_proto = out.File
	file_redis_proto_rawDesc = nil
	file_redis_proto_goTypes = nil
	file_redis_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type RedisService interface {
	Set(ctx context.Context, req *SetRequest) (res *SetResponse, err error)
	GetIncreId(ctx context.Context, req *GetIncreIdRequest) (res *GetIncreIdResponse, err error)
	ZSetIncre(ctx context.Context, req *ZSETIncreRequest) (res *ZSETIncreResponse, err error)
	ZSetGetMember(ctx context.Context, req *ZSETGetMemberRequest) (res *ZSETGetMemberResponse, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (res *GetUserInfoResponse, err error)
}
