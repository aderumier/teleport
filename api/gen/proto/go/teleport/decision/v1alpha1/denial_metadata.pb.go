// Copyright 2024 Gravitational, Inc
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
// source: teleport/decision/v1alpha1/denial_metadata.proto

package decisionpb

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

// Metadata for access denials.
type DenialMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// FeatureAssertions is a list of EnforcementFeature that the PEP (Policy
	// Enforcement Point) *must* implement in order to correctly enforce the
	// decision. Note that denials rarely need feature assertions since they
	// typically "fail safe" anyway.
	FeatureAssertions []EnforcementFeature `protobuf:"varint,1,rep,packed,name=feature_assertions,json=featureAssertions,proto3,enum=teleport.decision.v1alpha1.EnforcementFeature" json:"feature_assertions,omitempty"`
	// PdpVersion is the version of the PDP (Policy Decision Point) that evaluated
	// the decision request.
	PdpVersion string `protobuf:"bytes,2,opt,name=pdp_version,json=pdpVersion,proto3" json:"pdp_version,omitempty"`
	// UserMessage is a sanitized message safe for return to the subject identity
	// of the decision request.
	UserMessage   string `protobuf:"bytes,3,opt,name=user_message,json=userMessage,proto3" json:"user_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DenialMetadata) Reset() {
	*x = DenialMetadata{}
	mi := &file_teleport_decision_v1alpha1_denial_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DenialMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenialMetadata) ProtoMessage() {}

func (x *DenialMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_decision_v1alpha1_denial_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenialMetadata.ProtoReflect.Descriptor instead.
func (*DenialMetadata) Descriptor() ([]byte, []int) {
	return file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *DenialMetadata) GetFeatureAssertions() []EnforcementFeature {
	if x != nil {
		return x.FeatureAssertions
	}
	return nil
}

func (x *DenialMetadata) GetPdpVersion() string {
	if x != nil {
		return x.PdpVersion
	}
	return ""
}

func (x *DenialMetadata) GetUserMessage() string {
	if x != nil {
		return x.UserMessage
	}
	return ""
}

var File_teleport_decision_v1alpha1_denial_metadata_proto protoreflect.FileDescriptor

const file_teleport_decision_v1alpha1_denial_metadata_proto_rawDesc = "" +
	"\n" +
	"0teleport/decision/v1alpha1/denial_metadata.proto\x12\x1ateleport.decision.v1alpha1\x1a4teleport/decision/v1alpha1/enforcement_feature.proto\"\xb3\x01\n" +
	"\x0eDenialMetadata\x12]\n" +
	"\x12feature_assertions\x18\x01 \x03(\x0e2..teleport.decision.v1alpha1.EnforcementFeatureR\x11featureAssertions\x12\x1f\n" +
	"\vpdp_version\x18\x02 \x01(\tR\n" +
	"pdpVersion\x12!\n" +
	"\fuser_message\x18\x03 \x01(\tR\vuserMessageBZZXgithub.com/gravitational/teleport/api/gen/proto/go/teleport/decision/v1alpha1;decisionpbb\x06proto3"

var (
	file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescOnce sync.Once
	file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescData []byte
)

func file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescGZIP() []byte {
	file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescOnce.Do(func() {
		file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_decision_v1alpha1_denial_metadata_proto_rawDesc), len(file_teleport_decision_v1alpha1_denial_metadata_proto_rawDesc)))
	})
	return file_teleport_decision_v1alpha1_denial_metadata_proto_rawDescData
}

var file_teleport_decision_v1alpha1_denial_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_teleport_decision_v1alpha1_denial_metadata_proto_goTypes = []any{
	(*DenialMetadata)(nil),  // 0: teleport.decision.v1alpha1.DenialMetadata
	(EnforcementFeature)(0), // 1: teleport.decision.v1alpha1.EnforcementFeature
}
var file_teleport_decision_v1alpha1_denial_metadata_proto_depIdxs = []int32{
	1, // 0: teleport.decision.v1alpha1.DenialMetadata.feature_assertions:type_name -> teleport.decision.v1alpha1.EnforcementFeature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_decision_v1alpha1_denial_metadata_proto_init() }
func file_teleport_decision_v1alpha1_denial_metadata_proto_init() {
	if File_teleport_decision_v1alpha1_denial_metadata_proto != nil {
		return
	}
	file_teleport_decision_v1alpha1_enforcement_feature_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_decision_v1alpha1_denial_metadata_proto_rawDesc), len(file_teleport_decision_v1alpha1_denial_metadata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_decision_v1alpha1_denial_metadata_proto_goTypes,
		DependencyIndexes: file_teleport_decision_v1alpha1_denial_metadata_proto_depIdxs,
		MessageInfos:      file_teleport_decision_v1alpha1_denial_metadata_proto_msgTypes,
	}.Build()
	File_teleport_decision_v1alpha1_denial_metadata_proto = out.File
	file_teleport_decision_v1alpha1_denial_metadata_proto_goTypes = nil
	file_teleport_decision_v1alpha1_denial_metadata_proto_depIdxs = nil
}
