// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.0
// source: auth/auth.proto

package auth

import (
	context "context"
	sdkws "github.com/SXTSOFT/protocol/sdkws"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserBindingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32           `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID,omitempty"`
	User       *sdkws.UserInfo `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserBindingReq) Reset() {
	*x = UserBindingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBindingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBindingReq) ProtoMessage() {}

func (x *UserBindingReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBindingReq.ProtoReflect.Descriptor instead.
func (*UserBindingReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserBindingReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *UserBindingReq) GetUser() *sdkws.UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type UserBindingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserBindingResp) Reset() {
	*x = UserBindingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBindingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBindingResp) ProtoMessage() {}

func (x *UserBindingResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBindingResp.ProtoReflect.Descriptor instead.
func (*UserBindingResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserBindingResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserBindingResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret     string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	PlatformID int32  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserTokenReq) Reset() {
	*x = UserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenReq) ProtoMessage() {}

func (x *UserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenReq.ProtoReflect.Descriptor instead.
func (*UserTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserTokenReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UserTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *UserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ExpireTimeSeconds int64  `protobuf:"varint,3,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *UserTokenResp) Reset() {
	*x = UserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenResp) ProtoMessage() {}

func (x *UserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenResp.ProtoReflect.Descriptor instead.
func (*UserTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type ForceLogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ForceLogoutReq) Reset() {
	*x = ForceLogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutReq) ProtoMessage() {}

func (x *ForceLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutReq.ProtoReflect.Descriptor instead.
func (*ForceLogoutReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ForceLogoutReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ForceLogoutReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ForceLogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForceLogoutResp) Reset() {
	*x = ForceLogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutResp) ProtoMessage() {}

func (x *ForceLogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutResp.ProtoReflect.Descriptor instead.
func (*ForceLogoutResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

type ParseTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ParseTokenReq) Reset() {
	*x = ParseTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenReq) ProtoMessage() {}

func (x *ParseTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenReq.ProtoReflect.Descriptor instead.
func (*ParseTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *ParseTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ParseTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID            string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Platform          string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	ExpireTimeSeconds int64  `protobuf:"varint,4,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *ParseTokenResp) Reset() {
	*x = ParseTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenResp) ProtoMessage() {}

func (x *ParseTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenResp.ProtoReflect.Descriptor instead.
func (*ParseTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *ParseTokenResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ParseTokenResp) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ParseTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x1a, 0x11, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x77,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x22, 0x48, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x11, 0x0a, 0x0f, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25,
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x72, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2c, 0x0a, 0x11, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x32, 0xd5, 0x02, 0x0a, 0x04, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x54, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x21, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x21, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x4f, 0x70, 0x65,
	0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51,
	0x0a, 0x0a, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x21,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4f, 0x70, 0x65, 0x6e, 0x49, 0x4d, 0x53, 0x44, 0x4b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData = file_auth_auth_proto_rawDesc
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_proto_rawDescData)
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_auth_proto_goTypes = []interface{}{
	(*UserBindingReq)(nil),  // 0: OpenIMServer.auth.userBindingReq
	(*UserBindingResp)(nil), // 1: OpenIMServer.auth.userBindingResp
	(*UserTokenReq)(nil),    // 2: OpenIMServer.auth.userTokenReq
	(*UserTokenResp)(nil),   // 3: OpenIMServer.auth.userTokenResp
	(*ForceLogoutReq)(nil),  // 4: OpenIMServer.auth.forceLogoutReq
	(*ForceLogoutResp)(nil), // 5: OpenIMServer.auth.forceLogoutResp
	(*ParseTokenReq)(nil),   // 6: OpenIMServer.auth.parseTokenReq
	(*ParseTokenResp)(nil),  // 7: OpenIMServer.auth.parseTokenResp
	(*sdkws.UserInfo)(nil),  // 8: OpenIMServer.sdkws.UserInfo
}
var file_auth_auth_proto_depIdxs = []int32{
	8, // 0: OpenIMServer.auth.userBindingReq.user:type_name -> OpenIMServer.sdkws.UserInfo
	0, // 1: OpenIMServer.auth.Auth.userBinding:input_type -> OpenIMServer.auth.userBindingReq
	2, // 2: OpenIMServer.auth.Auth.userToken:input_type -> OpenIMServer.auth.userTokenReq
	4, // 3: OpenIMServer.auth.Auth.forceLogout:input_type -> OpenIMServer.auth.forceLogoutReq
	6, // 4: OpenIMServer.auth.Auth.parseToken:input_type -> OpenIMServer.auth.parseTokenReq
	1, // 5: OpenIMServer.auth.Auth.userBinding:output_type -> OpenIMServer.auth.userBindingResp
	3, // 6: OpenIMServer.auth.Auth.userToken:output_type -> OpenIMServer.auth.userTokenResp
	5, // 7: OpenIMServer.auth.Auth.forceLogout:output_type -> OpenIMServer.auth.forceLogoutResp
	7, // 8: OpenIMServer.auth.Auth.parseToken:output_type -> OpenIMServer.auth.parseTokenResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBindingReq); i {
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
		file_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBindingResp); i {
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
		file_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenReq); i {
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
		file_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTokenResp); i {
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
		file_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutReq); i {
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
		file_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutResp); i {
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
		file_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenReq); i {
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
		file_auth_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenResp); i {
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
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// user binding
	UserBinding(ctx context.Context, in *UserBindingReq, opts ...grpc.CallOption) (*UserBindingResp, error)
	// 生成token
	UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error)
	// 强制退出登录
	ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...grpc.CallOption) (*ForceLogoutResp, error)
	// 解析token
	ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) UserBinding(ctx context.Context, in *UserBindingReq, opts ...grpc.CallOption) (*UserBindingResp, error) {
	out := new(UserBindingResp)
	err := c.cc.Invoke(ctx, "/OpenIMServer.auth.Auth/userBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error) {
	out := new(UserTokenResp)
	err := c.cc.Invoke(ctx, "/OpenIMServer.auth.Auth/userToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ForceLogout(ctx context.Context, in *ForceLogoutReq, opts ...grpc.CallOption) (*ForceLogoutResp, error) {
	out := new(ForceLogoutResp)
	err := c.cc.Invoke(ctx, "/OpenIMServer.auth.Auth/forceLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error) {
	out := new(ParseTokenResp)
	err := c.cc.Invoke(ctx, "/OpenIMServer.auth.Auth/parseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// user binding
	UserBinding(context.Context, *UserBindingReq) (*UserBindingResp, error)
	// 生成token
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
	// 强制退出登录
	ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error)
	// 解析token
	ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) UserBinding(context.Context, *UserBindingReq) (*UserBindingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBinding not implemented")
}
func (*UnimplementedAuthServer) UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken not implemented")
}
func (*UnimplementedAuthServer) ForceLogout(context.Context, *ForceLogoutReq) (*ForceLogoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceLogout not implemented")
}
func (*UnimplementedAuthServer) ParseToken(context.Context, *ParseTokenReq) (*ParseTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_UserBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBindingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OpenIMServer.auth.Auth/UserBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserBinding(ctx, req.(*UserBindingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OpenIMServer.auth.Auth/UserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserToken(ctx, req.(*UserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ForceLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceLogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ForceLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OpenIMServer.auth.Auth/ForceLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ForceLogout(ctx, req.(*ForceLogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ParseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ParseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OpenIMServer.auth.Auth/ParseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ParseToken(ctx, req.(*ParseTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OpenIMServer.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "userBinding",
			Handler:    _Auth_UserBinding_Handler,
		},
		{
			MethodName: "userToken",
			Handler:    _Auth_UserToken_Handler,
		},
		{
			MethodName: "forceLogout",
			Handler:    _Auth_ForceLogout_Handler,
		},
		{
			MethodName: "parseToken",
			Handler:    _Auth_ParseToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
