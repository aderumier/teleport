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
// source: teleport/identitycenter/v1/identitycenter_service.proto

package identitycenterv1

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

// DeleteAllIdentityCenterAccountsRequest is a request to delete all Identity Center imported accounts.
type DeleteAllIdentityCenterAccountsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllIdentityCenterAccountsRequest) Reset() {
	*x = DeleteAllIdentityCenterAccountsRequest{}
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllIdentityCenterAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllIdentityCenterAccountsRequest) ProtoMessage() {}

func (x *DeleteAllIdentityCenterAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllIdentityCenterAccountsRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllIdentityCenterAccountsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescGZIP(), []int{0}
}

// DeleteAllAccountAssignmentsRequest is a request to delete all Identity Center account assignments.
type DeleteAllAccountAssignmentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllAccountAssignmentsRequest) Reset() {
	*x = DeleteAllAccountAssignmentsRequest{}
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllAccountAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllAccountAssignmentsRequest) ProtoMessage() {}

func (x *DeleteAllAccountAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllAccountAssignmentsRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllAccountAssignmentsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescGZIP(), []int{1}
}

// DeleteAllPrincipalAssignmentsRequest is a request to delete all Identity Center principal assignments.
type DeleteAllPrincipalAssignmentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllPrincipalAssignmentsRequest) Reset() {
	*x = DeleteAllPrincipalAssignmentsRequest{}
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllPrincipalAssignmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllPrincipalAssignmentsRequest) ProtoMessage() {}

func (x *DeleteAllPrincipalAssignmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllPrincipalAssignmentsRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllPrincipalAssignmentsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescGZIP(), []int{2}
}

// DeleteAllPermissionSetsRequest is a request to delete all Identity Center permission sets.
type DeleteAllPermissionSetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllPermissionSetsRequest) Reset() {
	*x = DeleteAllPermissionSetsRequest{}
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllPermissionSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllPermissionSetsRequest) ProtoMessage() {}

func (x *DeleteAllPermissionSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllPermissionSetsRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllPermissionSetsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescGZIP(), []int{3}
}

var File_teleport_identitycenter_v1_identitycenter_service_proto protoreflect.FileDescriptor

const file_teleport_identitycenter_v1_identitycenter_service_proto_rawDesc = "" +
	"\n" +
	"7teleport/identitycenter/v1/identitycenter_service.proto\x12\x1ateleport.identitycenter.v1\x1a\x1bgoogle/protobuf/empty.proto\"(\n" +
	"&DeleteAllIdentityCenterAccountsRequest\"$\n" +
	"\"DeleteAllAccountAssignmentsRequest\"&\n" +
	"$DeleteAllPrincipalAssignmentsRequest\" \n" +
	"\x1eDeleteAllPermissionSetsRequest2\xf7\x03\n" +
	"\x15IdentityCenterService\x12}\n" +
	"\x1fDeleteAllIdentityCenterAccounts\x12B.teleport.identitycenter.v1.DeleteAllIdentityCenterAccountsRequest\x1a\x16.google.protobuf.Empty\x12u\n" +
	"\x1bDeleteAllAccountAssignments\x12>.teleport.identitycenter.v1.DeleteAllAccountAssignmentsRequest\x1a\x16.google.protobuf.Empty\x12y\n" +
	"\x1dDeleteAllPrincipalAssignments\x12@.teleport.identitycenter.v1.DeleteAllPrincipalAssignmentsRequest\x1a\x16.google.protobuf.Empty\x12m\n" +
	"\x17DeleteAllPermissionSets\x12:.teleport.identitycenter.v1.DeleteAllPermissionSetsRequest\x1a\x16.google.protobuf.EmptyB`Z^github.com/gravitational/teleport/api/gen/proto/go/teleport/identitycenter/v1;identitycenterv1b\x06proto3"

var (
	file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescOnce sync.Once
	file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescData []byte
)

func file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescGZIP() []byte {
	file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescOnce.Do(func() {
		file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_identitycenter_v1_identitycenter_service_proto_rawDesc), len(file_teleport_identitycenter_v1_identitycenter_service_proto_rawDesc)))
	})
	return file_teleport_identitycenter_v1_identitycenter_service_proto_rawDescData
}

var file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_identitycenter_v1_identitycenter_service_proto_goTypes = []any{
	(*DeleteAllIdentityCenterAccountsRequest)(nil), // 0: teleport.identitycenter.v1.DeleteAllIdentityCenterAccountsRequest
	(*DeleteAllAccountAssignmentsRequest)(nil),     // 1: teleport.identitycenter.v1.DeleteAllAccountAssignmentsRequest
	(*DeleteAllPrincipalAssignmentsRequest)(nil),   // 2: teleport.identitycenter.v1.DeleteAllPrincipalAssignmentsRequest
	(*DeleteAllPermissionSetsRequest)(nil),         // 3: teleport.identitycenter.v1.DeleteAllPermissionSetsRequest
	(*emptypb.Empty)(nil),                          // 4: google.protobuf.Empty
}
var file_teleport_identitycenter_v1_identitycenter_service_proto_depIdxs = []int32{
	0, // 0: teleport.identitycenter.v1.IdentityCenterService.DeleteAllIdentityCenterAccounts:input_type -> teleport.identitycenter.v1.DeleteAllIdentityCenterAccountsRequest
	1, // 1: teleport.identitycenter.v1.IdentityCenterService.DeleteAllAccountAssignments:input_type -> teleport.identitycenter.v1.DeleteAllAccountAssignmentsRequest
	2, // 2: teleport.identitycenter.v1.IdentityCenterService.DeleteAllPrincipalAssignments:input_type -> teleport.identitycenter.v1.DeleteAllPrincipalAssignmentsRequest
	3, // 3: teleport.identitycenter.v1.IdentityCenterService.DeleteAllPermissionSets:input_type -> teleport.identitycenter.v1.DeleteAllPermissionSetsRequest
	4, // 4: teleport.identitycenter.v1.IdentityCenterService.DeleteAllIdentityCenterAccounts:output_type -> google.protobuf.Empty
	4, // 5: teleport.identitycenter.v1.IdentityCenterService.DeleteAllAccountAssignments:output_type -> google.protobuf.Empty
	4, // 6: teleport.identitycenter.v1.IdentityCenterService.DeleteAllPrincipalAssignments:output_type -> google.protobuf.Empty
	4, // 7: teleport.identitycenter.v1.IdentityCenterService.DeleteAllPermissionSets:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_identitycenter_v1_identitycenter_service_proto_init() }
func file_teleport_identitycenter_v1_identitycenter_service_proto_init() {
	if File_teleport_identitycenter_v1_identitycenter_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_identitycenter_v1_identitycenter_service_proto_rawDesc), len(file_teleport_identitycenter_v1_identitycenter_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_identitycenter_v1_identitycenter_service_proto_goTypes,
		DependencyIndexes: file_teleport_identitycenter_v1_identitycenter_service_proto_depIdxs,
		MessageInfos:      file_teleport_identitycenter_v1_identitycenter_service_proto_msgTypes,
	}.Build()
	File_teleport_identitycenter_v1_identitycenter_service_proto = out.File
	file_teleport_identitycenter_v1_identitycenter_service_proto_goTypes = nil
	file_teleport_identitycenter_v1_identitycenter_service_proto_depIdxs = nil
}
