// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.29.3
// source: helloworld/v1/greeter.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserModel) Reset() {
	*x = UserModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModel) ProtoMessage() {}

func (x *UserModel) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[0]
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
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *UserModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserModel) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type GetUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserReq) Reset() {
	*x = GetUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReq) ProtoMessage() {}

func (x *GetUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReq.ProtoReflect.Descriptor instead.
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserModel `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserReply) GetUser() *UserModel {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserReq) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserReq) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type DeleteUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteUserReply) Reset() {
	*x = DeleteUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReply) ProtoMessage() {}

func (x *DeleteUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReply.ProtoReflect.Descriptor instead.
func (*DeleteUserReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUserReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit   int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *ListUserReq) Reset() {
	*x = ListUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReq) ProtoMessage() {}

func (x *ListUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReq.ProtoReflect.Descriptor instead.
func (*ListUserReq) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{7}
}

func (x *ListUserReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListUserReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUserReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserModel `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_greeter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_greeter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_greeter_proto_rawDescGZIP(), []int{8}
}

func (x *ListUserReply) GetUsers() []*UserModel {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_helloworld_v1_greeter_proto protoreflect.FileDescriptor

var file_helloworld_v1_greeter_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x1c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x45, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x51, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xed, 0x03, 0x0a, 0x0a, 0x42, 0x66,
	0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1b,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x63, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x59,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0x4f, 0x0a, 0x1c, 0x64, 0x65, 0x76,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1a,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_helloworld_v1_greeter_proto_rawDescOnce sync.Once
	file_helloworld_v1_greeter_proto_rawDescData = file_helloworld_v1_greeter_proto_rawDesc
)

func file_helloworld_v1_greeter_proto_rawDescGZIP() []byte {
	file_helloworld_v1_greeter_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_greeter_proto_rawDescData)
	})
	return file_helloworld_v1_greeter_proto_rawDescData
}

var file_helloworld_v1_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_helloworld_v1_greeter_proto_goTypes = []interface{}{
	(*UserModel)(nil),       // 0: helloworld.v1.UserModel
	(*GetUserReq)(nil),      // 1: helloworld.v1.GetUserReq
	(*GetUserReply)(nil),    // 2: helloworld.v1.GetUserReply
	(*CreateUserReq)(nil),   // 3: helloworld.v1.CreateUserReq
	(*UpdateUserReq)(nil),   // 4: helloworld.v1.UpdateUserReq
	(*DeleteUserReq)(nil),   // 5: helloworld.v1.DeleteUserReq
	(*DeleteUserReply)(nil), // 6: helloworld.v1.DeleteUserReply
	(*ListUserReq)(nil),     // 7: helloworld.v1.ListUserReq
	(*ListUserReply)(nil),   // 8: helloworld.v1.ListUserReply
}
var file_helloworld_v1_greeter_proto_depIdxs = []int32{
	0, // 0: helloworld.v1.GetUserReply.user:type_name -> helloworld.v1.UserModel
	0, // 1: helloworld.v1.ListUserReply.users:type_name -> helloworld.v1.UserModel
	1, // 2: helloworld.v1.BffService.GetUser:input_type -> helloworld.v1.GetUserReq
	3, // 3: helloworld.v1.BffService.CreateUser:input_type -> helloworld.v1.CreateUserReq
	4, // 4: helloworld.v1.BffService.UpdateUser:input_type -> helloworld.v1.UpdateUserReq
	5, // 5: helloworld.v1.BffService.DeleteUser:input_type -> helloworld.v1.DeleteUserReq
	7, // 6: helloworld.v1.BffService.ListUsers:input_type -> helloworld.v1.ListUserReq
	2, // 7: helloworld.v1.BffService.GetUser:output_type -> helloworld.v1.GetUserReply
	2, // 8: helloworld.v1.BffService.CreateUser:output_type -> helloworld.v1.GetUserReply
	2, // 9: helloworld.v1.BffService.UpdateUser:output_type -> helloworld.v1.GetUserReply
	6, // 10: helloworld.v1.BffService.DeleteUser:output_type -> helloworld.v1.DeleteUserReply
	8, // 11: helloworld.v1.BffService.ListUsers:output_type -> helloworld.v1.ListUserReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_helloworld_v1_greeter_proto_init() }
func file_helloworld_v1_greeter_proto_init() {
	if File_helloworld_v1_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_v1_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_helloworld_v1_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReq); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReq); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReply); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReq); i {
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
		file_helloworld_v1_greeter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply); i {
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
			RawDescriptor: file_helloworld_v1_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_v1_greeter_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_greeter_proto_depIdxs,
		MessageInfos:      file_helloworld_v1_greeter_proto_msgTypes,
	}.Build()
	File_helloworld_v1_greeter_proto = out.File
	file_helloworld_v1_greeter_proto_rawDesc = nil
	file_helloworld_v1_greeter_proto_goTypes = nil
	file_helloworld_v1_greeter_proto_depIdxs = nil
}
