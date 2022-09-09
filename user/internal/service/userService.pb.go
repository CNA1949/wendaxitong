// 指定当前proto语法的版本，有proto2和proto3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: userService.proto

//指定*.pb.go文件生成出来的package，一般与上面的path相同

package service

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

// user Model
type UserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName   string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Phone      string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email      string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	NumFans    uint64 `protobuf:"varint,6,opt,name=num_fans,json=numFans,proto3" json:"num_fans,omitempty"`
	NumIdols   uint64 `protobuf:"varint,7,opt,name=num_idols,json=numIdols,proto3" json:"num_idols,omitempty"`
	FansNames  string `protobuf:"bytes,8,opt,name=fans_names,json=fansNames,proto3" json:"fans_names,omitempty"`
	IdolsNames string `protobuf:"bytes,9,opt,name=idols_names,json=idolsNames,proto3" json:"idols_names,omitempty"`
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModel.ProtoReflect.Descriptor instead.
func (*UserModel) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{0}
}

func (x *UserModel) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserModel) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserModel) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserModel) GetNumFans() uint64 {
	if x != nil {
		return x.NumFans
	}
	return 0
}

func (x *UserModel) GetNumIdols() uint64 {
	if x != nil {
		return x.NumIdols
	}
	return 0
}

func (x *UserModel) GetFansNames() string {
	if x != nil {
		return x.FansNames
	}
	return ""
}

func (x *UserModel) GetIdolsNames() string {
	if x != nil {
		return x.IdolsNames
	}
	return ""
}

type UserInfoModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	NumFans  uint64 `protobuf:"varint,6,opt,name=num_fans,json=numFans,proto3" json:"num_fans,omitempty"`
	NumIdols uint64 `protobuf:"varint,7,opt,name=num_idols,json=numIdols,proto3" json:"num_idols,omitempty"`
}

func (x *UserInfoModel) Reset() {
	*x = UserInfoModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoModel) ProtoMessage() {}

func (x *UserInfoModel) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoModel.ProtoReflect.Descriptor instead.
func (*UserInfoModel) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoModel) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfoModel) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfoModel) GetNumFans() uint64 {
	if x != nil {
		return x.NumFans
	}
	return 0
}

func (x *UserInfoModel) GetNumIdols() uint64 {
	if x != nil {
		return x.NumIdols
	}
	return 0
}

// request model
// 消息 传输对象
type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserModel `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{2}
}

func (x *UserRequest) GetUserInfo() *UserModel {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserRequest2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserModel `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	Choose   uint64     `protobuf:"varint,2,opt,name=choose,proto3" json:"choose,omitempty"`
}

func (x *UserRequest2) Reset() {
	*x = UserRequest2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest2) ProtoMessage() {}

func (x *UserRequest2) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest2.ProtoReflect.Descriptor instead.
func (*UserRequest2) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{3}
}

func (x *UserRequest2) GetUserInfo() *UserModel {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *UserRequest2) GetChoose() uint64 {
	if x != nil {
		return x.Choose
	}
	return 0
}

// response model
type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    uint64     `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string     `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	UserInfo      *UserModel `protobuf:"bytes,3,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{4}
}

func (x *UserResponse) GetStatusCode() uint64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UserResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *UserResponse) GetUserInfo() *UserModel {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    uint64         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string         `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	UserInfo      *UserInfoModel `protobuf:"bytes,3,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_userService_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoResponse) GetStatusCode() uint64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UserInfoResponse) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *UserInfoResponse) GetUserInfo() *UserInfoModel {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_userService_proto protoreflect.FileDescriptor

var file_userService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xff, 0x01, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x61,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x46, 0x61, 0x6e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x6f, 0x6c, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x49, 0x64, 0x6f, 0x6c, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x61, 0x6e, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x61, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x64, 0x6f, 0x6c, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x6f, 0x6c, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x7b,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x61, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x46, 0x61, 0x6e, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x6f, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x49, 0x64, 0x6f, 0x6c, 0x73, 0x22, 0x3d, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x12, 0x2e, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68,
	0x6f, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x6f,
	0x73, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xc4, 0x03, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userService_proto_rawDescOnce sync.Once
	file_userService_proto_rawDescData = file_userService_proto_rawDesc
)

func file_userService_proto_rawDescGZIP() []byte {
	file_userService_proto_rawDescOnce.Do(func() {
		file_userService_proto_rawDescData = protoimpl.X.CompressGZIP(file_userService_proto_rawDescData)
	})
	return file_userService_proto_rawDescData
}

var file_userService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_userService_proto_goTypes = []interface{}{
	(*UserModel)(nil),        // 0: service.UserModel
	(*UserInfoModel)(nil),    // 1: service.UserInfoModel
	(*UserRequest)(nil),      // 2: service.UserRequest
	(*UserRequest2)(nil),     // 3: service.UserRequest2
	(*UserResponse)(nil),     // 4: service.UserResponse
	(*UserInfoResponse)(nil), // 5: service.UserInfoResponse
}
var file_userService_proto_depIdxs = []int32{
	0,  // 0: service.UserRequest.userInfo:type_name -> service.UserModel
	0,  // 1: service.UserRequest2.userInfo:type_name -> service.UserModel
	0,  // 2: service.UserResponse.userInfo:type_name -> service.UserModel
	1,  // 3: service.UserInfoResponse.userInfo:type_name -> service.UserInfoModel
	2,  // 4: service.UserService.UserLogin:input_type -> service.UserRequest
	2,  // 5: service.UserService.UserRegister:input_type -> service.UserRequest
	2,  // 6: service.UserService.DeleteUser:input_type -> service.UserRequest
	2,  // 7: service.UserService.ModifyUserInfo:input_type -> service.UserRequest
	2,  // 8: service.UserService.GetUserIdByUserName:input_type -> service.UserRequest
	2,  // 9: service.UserService.GetUserInfoByUserName:input_type -> service.UserRequest
	3,  // 10: service.UserService.FollowUser:input_type -> service.UserRequest2
	4,  // 11: service.UserService.UserLogin:output_type -> service.UserResponse
	4,  // 12: service.UserService.UserRegister:output_type -> service.UserResponse
	4,  // 13: service.UserService.DeleteUser:output_type -> service.UserResponse
	4,  // 14: service.UserService.ModifyUserInfo:output_type -> service.UserResponse
	4,  // 15: service.UserService.GetUserIdByUserName:output_type -> service.UserResponse
	4,  // 16: service.UserService.GetUserInfoByUserName:output_type -> service.UserResponse
	4,  // 17: service.UserService.FollowUser:output_type -> service.UserResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_userService_proto_init() }
func file_userService_proto_init() {
	if File_userService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModel); i {
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
		file_userService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoModel); i {
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
		file_userService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_userService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest2); i {
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
		file_userService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_userService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
			RawDescriptor: file_userService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userService_proto_goTypes,
		DependencyIndexes: file_userService_proto_depIdxs,
		MessageInfos:      file_userService_proto_msgTypes,
	}.Build()
	File_userService_proto = out.File
	file_userService_proto_rawDesc = nil
	file_userService_proto_goTypes = nil
	file_userService_proto_depIdxs = nil
}
