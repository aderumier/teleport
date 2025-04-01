// Copyright 2024 Gravitational, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: teleport/userprovisioning/v2/statichostuser_service.proto

package userprovisioningv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for GetStaticHostUser.
type GetStaticHostUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the user to retrieve, this take priority over current_user.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStaticHostUserRequest) Reset() {
	*x = GetStaticHostUserRequest{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticHostUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticHostUserRequest) ProtoMessage() {}

func (x *GetStaticHostUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticHostUserRequest.ProtoReflect.Descriptor instead.
func (*GetStaticHostUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetStaticHostUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for ListStaticHostUsers.
//
// Follows the pagination semantics of
// https://cloud.google.com/apis/design/standard_methods#list.
type ListStaticHostUsersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStaticHostUsersRequest) Reset() {
	*x = ListStaticHostUsersRequest{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStaticHostUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaticHostUsersRequest) ProtoMessage() {}

func (x *ListStaticHostUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaticHostUsersRequest.ProtoReflect.Descriptor instead.
func (*ListStaticHostUsersRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListStaticHostUsersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListStaticHostUsersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response for ListStaticHostUsers.
type ListStaticHostUsersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Static host users that matched the search.
	Users []*StaticHostUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStaticHostUsersResponse) Reset() {
	*x = ListStaticHostUsersResponse{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStaticHostUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaticHostUsersResponse) ProtoMessage() {}

func (x *ListStaticHostUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaticHostUsersResponse.ProtoReflect.Descriptor instead.
func (*ListStaticHostUsersResponse) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListStaticHostUsersResponse) GetUsers() []*StaticHostUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListStaticHostUsersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request for CreateStaticHostUser.
type CreateStaticHostUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The static host user resource to create.
	User          *StaticHostUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStaticHostUserRequest) Reset() {
	*x = CreateStaticHostUserRequest{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStaticHostUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaticHostUserRequest) ProtoMessage() {}

func (x *CreateStaticHostUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaticHostUserRequest.ProtoReflect.Descriptor instead.
func (*CreateStaticHostUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateStaticHostUserRequest) GetUser() *StaticHostUser {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for UpdateStaticHostUser.
type UpdateStaticHostUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The static host user resource to update.
	User          *StaticHostUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStaticHostUserRequest) Reset() {
	*x = UpdateStaticHostUserRequest{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStaticHostUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaticHostUserRequest) ProtoMessage() {}

func (x *UpdateStaticHostUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaticHostUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateStaticHostUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStaticHostUserRequest) GetUser() *StaticHostUser {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for UpsertStaticHostUser.
type UpsertStaticHostUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The static host user resource to upsert.
	User          *StaticHostUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertStaticHostUserRequest) Reset() {
	*x = UpsertStaticHostUserRequest{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertStaticHostUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertStaticHostUserRequest) ProtoMessage() {}

func (x *UpsertStaticHostUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertStaticHostUserRequest.ProtoReflect.Descriptor instead.
func (*UpsertStaticHostUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertStaticHostUserRequest) GetUser() *StaticHostUser {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for DeleteStaticHostUser.
type DeleteStaticHostUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the user resource to remove.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStaticHostUserRequest) Reset() {
	*x = DeleteStaticHostUserRequest{}
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStaticHostUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaticHostUserRequest) ProtoMessage() {}

func (x *DeleteStaticHostUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaticHostUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteStaticHostUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteStaticHostUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_teleport_userprovisioning_v2_statichostuser_service_proto protoreflect.FileDescriptor

const file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDesc = "" +
	"\n" +
	"9teleport/userprovisioning/v2/statichostuser_service.proto\x12\x1cteleport.userprovisioning.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a1teleport/userprovisioning/v2/statichostuser.proto\".\n" +
	"\x18GetStaticHostUserRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"X\n" +
	"\x1aListStaticHostUsersRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"\x89\x01\n" +
	"\x1bListStaticHostUsersResponse\x12B\n" +
	"\x05users\x18\x01 \x03(\v2,.teleport.userprovisioning.v2.StaticHostUserR\x05users\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"_\n" +
	"\x1bCreateStaticHostUserRequest\x12@\n" +
	"\x04user\x18\x01 \x01(\v2,.teleport.userprovisioning.v2.StaticHostUserR\x04user\"_\n" +
	"\x1bUpdateStaticHostUserRequest\x12@\n" +
	"\x04user\x18\x02 \x01(\v2,.teleport.userprovisioning.v2.StaticHostUserR\x04user\"_\n" +
	"\x1bUpsertStaticHostUserRequest\x12@\n" +
	"\x04user\x18\x02 \x01(\v2,.teleport.userprovisioning.v2.StaticHostUserR\x04user\"1\n" +
	"\x1bDeleteStaticHostUserRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name2\x8e\x06\n" +
	"\x16StaticHostUsersService\x12y\n" +
	"\x11GetStaticHostUser\x126.teleport.userprovisioning.v2.GetStaticHostUserRequest\x1a,.teleport.userprovisioning.v2.StaticHostUser\x12\x8a\x01\n" +
	"\x13ListStaticHostUsers\x128.teleport.userprovisioning.v2.ListStaticHostUsersRequest\x1a9.teleport.userprovisioning.v2.ListStaticHostUsersResponse\x12\x7f\n" +
	"\x14CreateStaticHostUser\x129.teleport.userprovisioning.v2.CreateStaticHostUserRequest\x1a,.teleport.userprovisioning.v2.StaticHostUser\x12\x7f\n" +
	"\x14UpdateStaticHostUser\x129.teleport.userprovisioning.v2.UpdateStaticHostUserRequest\x1a,.teleport.userprovisioning.v2.StaticHostUser\x12\x7f\n" +
	"\x14UpsertStaticHostUser\x129.teleport.userprovisioning.v2.UpsertStaticHostUserRequest\x1a,.teleport.userprovisioning.v2.StaticHostUser\x12i\n" +
	"\x14DeleteStaticHostUser\x129.teleport.userprovisioning.v2.DeleteStaticHostUserRequest\x1a\x16.google.protobuf.EmptyBdZbgithub.com/gravitational/teleport/api/gen/proto/go/teleport/userprovisioning/v2;userprovisioningv2b\x06proto3"

var (
	file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescOnce sync.Once
	file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescData []byte
)

func file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescGZIP() []byte {
	file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescOnce.Do(func() {
		file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDesc), len(file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDesc)))
	})
	return file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDescData
}

var file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_teleport_userprovisioning_v2_statichostuser_service_proto_goTypes = []any{
	(*GetStaticHostUserRequest)(nil),    // 0: teleport.userprovisioning.v2.GetStaticHostUserRequest
	(*ListStaticHostUsersRequest)(nil),  // 1: teleport.userprovisioning.v2.ListStaticHostUsersRequest
	(*ListStaticHostUsersResponse)(nil), // 2: teleport.userprovisioning.v2.ListStaticHostUsersResponse
	(*CreateStaticHostUserRequest)(nil), // 3: teleport.userprovisioning.v2.CreateStaticHostUserRequest
	(*UpdateStaticHostUserRequest)(nil), // 4: teleport.userprovisioning.v2.UpdateStaticHostUserRequest
	(*UpsertStaticHostUserRequest)(nil), // 5: teleport.userprovisioning.v2.UpsertStaticHostUserRequest
	(*DeleteStaticHostUserRequest)(nil), // 6: teleport.userprovisioning.v2.DeleteStaticHostUserRequest
	(*StaticHostUser)(nil),              // 7: teleport.userprovisioning.v2.StaticHostUser
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_teleport_userprovisioning_v2_statichostuser_service_proto_depIdxs = []int32{
	7,  // 0: teleport.userprovisioning.v2.ListStaticHostUsersResponse.users:type_name -> teleport.userprovisioning.v2.StaticHostUser
	7,  // 1: teleport.userprovisioning.v2.CreateStaticHostUserRequest.user:type_name -> teleport.userprovisioning.v2.StaticHostUser
	7,  // 2: teleport.userprovisioning.v2.UpdateStaticHostUserRequest.user:type_name -> teleport.userprovisioning.v2.StaticHostUser
	7,  // 3: teleport.userprovisioning.v2.UpsertStaticHostUserRequest.user:type_name -> teleport.userprovisioning.v2.StaticHostUser
	0,  // 4: teleport.userprovisioning.v2.StaticHostUsersService.GetStaticHostUser:input_type -> teleport.userprovisioning.v2.GetStaticHostUserRequest
	1,  // 5: teleport.userprovisioning.v2.StaticHostUsersService.ListStaticHostUsers:input_type -> teleport.userprovisioning.v2.ListStaticHostUsersRequest
	3,  // 6: teleport.userprovisioning.v2.StaticHostUsersService.CreateStaticHostUser:input_type -> teleport.userprovisioning.v2.CreateStaticHostUserRequest
	4,  // 7: teleport.userprovisioning.v2.StaticHostUsersService.UpdateStaticHostUser:input_type -> teleport.userprovisioning.v2.UpdateStaticHostUserRequest
	5,  // 8: teleport.userprovisioning.v2.StaticHostUsersService.UpsertStaticHostUser:input_type -> teleport.userprovisioning.v2.UpsertStaticHostUserRequest
	6,  // 9: teleport.userprovisioning.v2.StaticHostUsersService.DeleteStaticHostUser:input_type -> teleport.userprovisioning.v2.DeleteStaticHostUserRequest
	7,  // 10: teleport.userprovisioning.v2.StaticHostUsersService.GetStaticHostUser:output_type -> teleport.userprovisioning.v2.StaticHostUser
	2,  // 11: teleport.userprovisioning.v2.StaticHostUsersService.ListStaticHostUsers:output_type -> teleport.userprovisioning.v2.ListStaticHostUsersResponse
	7,  // 12: teleport.userprovisioning.v2.StaticHostUsersService.CreateStaticHostUser:output_type -> teleport.userprovisioning.v2.StaticHostUser
	7,  // 13: teleport.userprovisioning.v2.StaticHostUsersService.UpdateStaticHostUser:output_type -> teleport.userprovisioning.v2.StaticHostUser
	7,  // 14: teleport.userprovisioning.v2.StaticHostUsersService.UpsertStaticHostUser:output_type -> teleport.userprovisioning.v2.StaticHostUser
	8,  // 15: teleport.userprovisioning.v2.StaticHostUsersService.DeleteStaticHostUser:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_userprovisioning_v2_statichostuser_service_proto_init() }
func file_teleport_userprovisioning_v2_statichostuser_service_proto_init() {
	if File_teleport_userprovisioning_v2_statichostuser_service_proto != nil {
		return
	}
	file_teleport_userprovisioning_v2_statichostuser_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDesc), len(file_teleport_userprovisioning_v2_statichostuser_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_userprovisioning_v2_statichostuser_service_proto_goTypes,
		DependencyIndexes: file_teleport_userprovisioning_v2_statichostuser_service_proto_depIdxs,
		MessageInfos:      file_teleport_userprovisioning_v2_statichostuser_service_proto_msgTypes,
	}.Build()
	File_teleport_userprovisioning_v2_statichostuser_service_proto = out.File
	file_teleport_userprovisioning_v2_statichostuser_service_proto_goTypes = nil
	file_teleport_userprovisioning_v2_statichostuser_service_proto_depIdxs = nil
}
