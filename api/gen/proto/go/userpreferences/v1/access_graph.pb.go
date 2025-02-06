// Copyright 2023 Gravitational, Inc
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
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: teleport/userpreferences/v1/access_graph.proto

package userpreferencesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// AccessGraphUserPreferences is the user preferences for Access Graph.
type AccessGraphUserPreferences struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// has_been_redirected is true if the user has already been redirected to the Access Graph
	// on login, after having signed up for a trial from the Teleport Policy page.
	HasBeenRedirected bool `protobuf:"varint,1,opt,name=has_been_redirected,json=hasBeenRedirected,proto3" json:"has_been_redirected,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *AccessGraphUserPreferences) Reset() {
	*x = AccessGraphUserPreferences{}
	mi := &file_teleport_userpreferences_v1_access_graph_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessGraphUserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGraphUserPreferences) ProtoMessage() {}

func (x *AccessGraphUserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_access_graph_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGraphUserPreferences.ProtoReflect.Descriptor instead.
func (*AccessGraphUserPreferences) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_access_graph_proto_rawDescGZIP(), []int{0}
}

func (x *AccessGraphUserPreferences) GetHasBeenRedirected() bool {
	if x != nil {
		return x.HasBeenRedirected
	}
	return false
}

var File_teleport_userpreferences_v1_access_graph_proto protoreflect.FileDescriptor

var file_teleport_userpreferences_v1_access_graph_proto_rawDesc = string([]byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x4c, 0x0a,
	0x1a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x68,
	0x61, 0x73, 0x5f, 0x62, 0x65, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x42, 0x65, 0x65,
	0x6e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x59, 0x5a, 0x57, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teleport_userpreferences_v1_access_graph_proto_rawDescOnce sync.Once
	file_teleport_userpreferences_v1_access_graph_proto_rawDescData []byte
)

func file_teleport_userpreferences_v1_access_graph_proto_rawDescGZIP() []byte {
	file_teleport_userpreferences_v1_access_graph_proto_rawDescOnce.Do(func() {
		file_teleport_userpreferences_v1_access_graph_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_userpreferences_v1_access_graph_proto_rawDesc), len(file_teleport_userpreferences_v1_access_graph_proto_rawDesc)))
	})
	return file_teleport_userpreferences_v1_access_graph_proto_rawDescData
}

var file_teleport_userpreferences_v1_access_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_userpreferences_v1_access_graph_proto_goTypes = []any{
	(*AccessGraphUserPreferences)(nil), // 0: teleport.userpreferences.v1.AccessGraphUserPreferences
}
var file_teleport_userpreferences_v1_access_graph_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleport_userpreferences_v1_access_graph_proto_init() }
func file_teleport_userpreferences_v1_access_graph_proto_init() {
	if File_teleport_userpreferences_v1_access_graph_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_userpreferences_v1_access_graph_proto_rawDesc), len(file_teleport_userpreferences_v1_access_graph_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userpreferences_v1_access_graph_proto_goTypes,
		DependencyIndexes: file_teleport_userpreferences_v1_access_graph_proto_depIdxs,
		MessageInfos:      file_teleport_userpreferences_v1_access_graph_proto_msgTypes,
	}.Build()
	File_teleport_userpreferences_v1_access_graph_proto = out.File
	file_teleport_userpreferences_v1_access_graph_proto_goTypes = nil
	file_teleport_userpreferences_v1_access_graph_proto_depIdxs = nil
}
