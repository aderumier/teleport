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
// source: teleport/provisioning/v1/provisioning.proto

package provisioningv1

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// ProvisioningState indicates the current stage of the provisioning pipeline a
// resource is in.
type ProvisioningState int32

const (
	// PROVISIONING_STATE_UNSPECIFIED is the `unset` default value. This should
	// never be used in practice.
	ProvisioningState_PROVISIONING_STATE_UNSPECIFIED ProvisioningState = 0
	// PROVISIONING_STATE_STALE indicates that Teleport has detected a change in
	// the principal resource, but that has not yet been reflected in the
	// downstream system.
	ProvisioningState_PROVISIONING_STATE_STALE ProvisioningState = 1
	// PROVISIONING_STATE_PROVISIONED indicates that the principal has been synced
	// with the downstram system and no further work is necessary.
	ProvisioningState_PROVISIONING_STATE_PROVISIONED ProvisioningState = 2
	// PROVISIONING_STATE_DELETED indicates that the principal has been deleted in
	// Teleport, and must be de-provisioned in the downstream system and the
	// provisioning state record deleted.
	ProvisioningState_PROVISIONING_STATE_DELETED ProvisioningState = 3
)

// Enum value maps for ProvisioningState.
var (
	ProvisioningState_name = map[int32]string{
		0: "PROVISIONING_STATE_UNSPECIFIED",
		1: "PROVISIONING_STATE_STALE",
		2: "PROVISIONING_STATE_PROVISIONED",
		3: "PROVISIONING_STATE_DELETED",
	}
	ProvisioningState_value = map[string]int32{
		"PROVISIONING_STATE_UNSPECIFIED": 0,
		"PROVISIONING_STATE_STALE":       1,
		"PROVISIONING_STATE_PROVISIONED": 2,
		"PROVISIONING_STATE_DELETED":     3,
	}
)

func (x ProvisioningState) Enum() *ProvisioningState {
	p := new(ProvisioningState)
	*p = x
	return p
}

func (x ProvisioningState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProvisioningState) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_provisioning_v1_provisioning_proto_enumTypes[0].Descriptor()
}

func (ProvisioningState) Type() protoreflect.EnumType {
	return &file_teleport_provisioning_v1_provisioning_proto_enumTypes[0]
}

func (x ProvisioningState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProvisioningState.Descriptor instead.
func (ProvisioningState) EnumDescriptor() ([]byte, []int) {
	return file_teleport_provisioning_v1_provisioning_proto_rawDescGZIP(), []int{0}
}

// PrincipalType indicates the type of principal represented by a PrincipalState
type PrincipalType int32

const (
	// PRINCIPAL_TYPE_UNSPECIFIED is the `unset` default value. This should
	// never be used in practice.
	PrincipalType_PRINCIPAL_TYPE_UNSPECIFIED PrincipalType = 0
	// PRINCIPAL_TYPE_USER indicates that the target principal is a Teleport user
	PrincipalType_PRINCIPAL_TYPE_USER PrincipalType = 1
	// PRINCIPAL_TYPE_USER indicates that the target principal is a Teleport
	// Access List
	PrincipalType_PRINCIPAL_TYPE_ACCESS_LIST PrincipalType = 2
)

// Enum value maps for PrincipalType.
var (
	PrincipalType_name = map[int32]string{
		0: "PRINCIPAL_TYPE_UNSPECIFIED",
		1: "PRINCIPAL_TYPE_USER",
		2: "PRINCIPAL_TYPE_ACCESS_LIST",
	}
	PrincipalType_value = map[string]int32{
		"PRINCIPAL_TYPE_UNSPECIFIED": 0,
		"PRINCIPAL_TYPE_USER":        1,
		"PRINCIPAL_TYPE_ACCESS_LIST": 2,
	}
)

func (x PrincipalType) Enum() *PrincipalType {
	p := new(PrincipalType)
	*p = x
	return p
}

func (x PrincipalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrincipalType) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_provisioning_v1_provisioning_proto_enumTypes[1].Descriptor()
}

func (PrincipalType) Type() protoreflect.EnumType {
	return &file_teleport_provisioning_v1_provisioning_proto_enumTypes[1]
}

func (x PrincipalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrincipalType.Descriptor instead.
func (PrincipalType) EnumDescriptor() ([]byte, []int) {
	return file_teleport_provisioning_v1_provisioning_proto_rawDescGZIP(), []int{1}
}

// PrincipalState describes the provisioning state of a Teleport user in a
// downstream system
type PrincipalState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	SubKind       string                 `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	Version       string                 `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Metadata      *v1.Metadata           `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *PrincipalStateSpec    `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *PrincipalStateStatus  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrincipalState) Reset() {
	*x = PrincipalState{}
	mi := &file_teleport_provisioning_v1_provisioning_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrincipalState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalState) ProtoMessage() {}

func (x *PrincipalState) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_provisioning_v1_provisioning_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalState.ProtoReflect.Descriptor instead.
func (*PrincipalState) Descriptor() ([]byte, []int) {
	return file_teleport_provisioning_v1_provisioning_proto_rawDescGZIP(), []int{0}
}

func (x *PrincipalState) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *PrincipalState) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *PrincipalState) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PrincipalState) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PrincipalState) GetSpec() *PrincipalStateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *PrincipalState) GetStatus() *PrincipalStateStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// PrincipalStateSpec describes the current state of a provisioning operation. It
// serves as a Teleport-local record of the downstream state.
type PrincipalStateSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// DownstreamId identifies the downstream service that this state applies to.
	DownstreamId string `protobuf:"bytes,1,opt,name=downstream_id,json=downstreamId,proto3" json:"downstream_id,omitempty"`
	// PrincipalType identifies what kind of principal this state applies to, either
	// a User or a Group (i.e. AccessList)
	PrincipalType PrincipalType `protobuf:"varint,2,opt,name=principal_type,json=principalType,proto3,enum=teleport.provisioning.v1.PrincipalType" json:"principal_type,omitempty"`
	// PrincipalId identifies the Teleport User or Access List that this state
	// applies to
	PrincipalId   string `protobuf:"bytes,3,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrincipalStateSpec) Reset() {
	*x = PrincipalStateSpec{}
	mi := &file_teleport_provisioning_v1_provisioning_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrincipalStateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalStateSpec) ProtoMessage() {}

func (x *PrincipalStateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_provisioning_v1_provisioning_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalStateSpec.ProtoReflect.Descriptor instead.
func (*PrincipalStateSpec) Descriptor() ([]byte, []int) {
	return file_teleport_provisioning_v1_provisioning_proto_rawDescGZIP(), []int{1}
}

func (x *PrincipalStateSpec) GetDownstreamId() string {
	if x != nil {
		return x.DownstreamId
	}
	return ""
}

func (x *PrincipalStateSpec) GetPrincipalType() PrincipalType {
	if x != nil {
		return x.PrincipalType
	}
	return PrincipalType_PRINCIPAL_TYPE_UNSPECIFIED
}

func (x *PrincipalStateSpec) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

// PrincipalStateStatus contains the runtime-writable status block for the
// PrincipalState resource
type PrincipalStateStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ProvisioningState indicates the resource's current state in the
	// provisioning process state machine.
	ProvisioningState ProvisioningState `protobuf:"varint,5,opt,name=provisioning_state,json=provisioningState,proto3,enum=teleport.provisioning.v1.ProvisioningState" json:"provisioning_state,omitempty"`
	// ExternalID holds the ID used by the downstream system to represent this
	// principal
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// LastProvisioned records the last time this record was provisioined into
	// the downstream system.
	LastProvisioned *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_provisioned,json=lastProvisioned,proto3" json:"last_provisioned,omitempty"`
	// Error holds a description of the last provisioing error, if any.
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	// Revision holds the revision of the principal record provisioned into the
	// downstream system. Used to assert that the latest revision of the principal
	// is provisioned downstream and detect changes in the principal that require
	// re-provisoning.
	ProvisionedPrincipalRevision string `protobuf:"bytes,6,opt,name=provisioned_principal_revision,json=provisionedPrincipalRevision,proto3" json:"provisioned_principal_revision,omitempty"`
	// ActiveLocks holds the list of known active locks on the principal. Used to
	// store the lock state across restarts of Teleport in order to detect state
	// changes that may happen while Teleport is not running (e.g. a storage
	// backend deleting an expired lock record while Teleport is offline)
	ActiveLocks   []string `protobuf:"bytes,7,rep,name=active_locks,json=activeLocks,proto3" json:"active_locks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrincipalStateStatus) Reset() {
	*x = PrincipalStateStatus{}
	mi := &file_teleport_provisioning_v1_provisioning_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrincipalStateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalStateStatus) ProtoMessage() {}

func (x *PrincipalStateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_provisioning_v1_provisioning_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalStateStatus.ProtoReflect.Descriptor instead.
func (*PrincipalStateStatus) Descriptor() ([]byte, []int) {
	return file_teleport_provisioning_v1_provisioning_proto_rawDescGZIP(), []int{2}
}

func (x *PrincipalStateStatus) GetProvisioningState() ProvisioningState {
	if x != nil {
		return x.ProvisioningState
	}
	return ProvisioningState_PROVISIONING_STATE_UNSPECIFIED
}

func (x *PrincipalStateStatus) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *PrincipalStateStatus) GetLastProvisioned() *timestamppb.Timestamp {
	if x != nil {
		return x.LastProvisioned
	}
	return nil
}

func (x *PrincipalStateStatus) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PrincipalStateStatus) GetProvisionedPrincipalRevision() string {
	if x != nil {
		return x.ProvisionedPrincipalRevision
	}
	return ""
}

func (x *PrincipalStateStatus) GetActiveLocks() []string {
	if x != nil {
		return x.ActiveLocks
	}
	return nil
}

var File_teleport_provisioning_v1_provisioning_proto protoreflect.FileDescriptor

const file_teleport_provisioning_v1_provisioning_proto_rawDesc = "" +
	"\n" +
	"+teleport/provisioning/v1/provisioning.proto\x12\x18teleport.provisioning.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a!teleport/header/v1/metadata.proto\"\x9d\x02\n" +
	"\x0ePrincipalState\x12\x12\n" +
	"\x04kind\x18\x01 \x01(\tR\x04kind\x12\x19\n" +
	"\bsub_kind\x18\x02 \x01(\tR\asubKind\x12\x18\n" +
	"\aversion\x18\x03 \x01(\tR\aversion\x128\n" +
	"\bmetadata\x18\x04 \x01(\v2\x1c.teleport.header.v1.MetadataR\bmetadata\x12@\n" +
	"\x04spec\x18\x05 \x01(\v2,.teleport.provisioning.v1.PrincipalStateSpecR\x04spec\x12F\n" +
	"\x06status\x18\x06 \x01(\v2..teleport.provisioning.v1.PrincipalStateStatusR\x06status\"\xac\x01\n" +
	"\x12PrincipalStateSpec\x12#\n" +
	"\rdownstream_id\x18\x01 \x01(\tR\fdownstreamId\x12N\n" +
	"\x0eprincipal_type\x18\x02 \x01(\x0e2'.teleport.provisioning.v1.PrincipalTypeR\rprincipalType\x12!\n" +
	"\fprincipal_id\x18\x03 \x01(\tR\vprincipalId\"\xe7\x02\n" +
	"\x14PrincipalStateStatus\x12Z\n" +
	"\x12provisioning_state\x18\x05 \x01(\x0e2+.teleport.provisioning.v1.ProvisioningStateR\x11provisioningState\x12\x1f\n" +
	"\vexternal_id\x18\x02 \x01(\tR\n" +
	"externalId\x12E\n" +
	"\x10last_provisioned\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x0flastProvisioned\x12\x14\n" +
	"\x05error\x18\x04 \x01(\tR\x05error\x12D\n" +
	"\x1eprovisioned_principal_revision\x18\x06 \x01(\tR\x1cprovisionedPrincipalRevision\x12!\n" +
	"\factive_locks\x18\a \x03(\tR\vactiveLocksJ\x04\b\x01\x10\x02R\x06status*\x99\x01\n" +
	"\x11ProvisioningState\x12\"\n" +
	"\x1ePROVISIONING_STATE_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18PROVISIONING_STATE_STALE\x10\x01\x12\"\n" +
	"\x1ePROVISIONING_STATE_PROVISIONED\x10\x02\x12\x1e\n" +
	"\x1aPROVISIONING_STATE_DELETED\x10\x03*h\n" +
	"\rPrincipalType\x12\x1e\n" +
	"\x1aPRINCIPAL_TYPE_UNSPECIFIED\x10\x00\x12\x17\n" +
	"\x13PRINCIPAL_TYPE_USER\x10\x01\x12\x1e\n" +
	"\x1aPRINCIPAL_TYPE_ACCESS_LIST\x10\x02B\\ZZgithub.com/gravitational/teleport/api/gen/proto/go/teleport/provisioning/v1;provisioningv1b\x06proto3"

var (
	file_teleport_provisioning_v1_provisioning_proto_rawDescOnce sync.Once
	file_teleport_provisioning_v1_provisioning_proto_rawDescData []byte
)

func file_teleport_provisioning_v1_provisioning_proto_rawDescGZIP() []byte {
	file_teleport_provisioning_v1_provisioning_proto_rawDescOnce.Do(func() {
		file_teleport_provisioning_v1_provisioning_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_provisioning_v1_provisioning_proto_rawDesc), len(file_teleport_provisioning_v1_provisioning_proto_rawDesc)))
	})
	return file_teleport_provisioning_v1_provisioning_proto_rawDescData
}

var file_teleport_provisioning_v1_provisioning_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_teleport_provisioning_v1_provisioning_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_teleport_provisioning_v1_provisioning_proto_goTypes = []any{
	(ProvisioningState)(0),        // 0: teleport.provisioning.v1.ProvisioningState
	(PrincipalType)(0),            // 1: teleport.provisioning.v1.PrincipalType
	(*PrincipalState)(nil),        // 2: teleport.provisioning.v1.PrincipalState
	(*PrincipalStateSpec)(nil),    // 3: teleport.provisioning.v1.PrincipalStateSpec
	(*PrincipalStateStatus)(nil),  // 4: teleport.provisioning.v1.PrincipalStateStatus
	(*v1.Metadata)(nil),           // 5: teleport.header.v1.Metadata
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_teleport_provisioning_v1_provisioning_proto_depIdxs = []int32{
	5, // 0: teleport.provisioning.v1.PrincipalState.metadata:type_name -> teleport.header.v1.Metadata
	3, // 1: teleport.provisioning.v1.PrincipalState.spec:type_name -> teleport.provisioning.v1.PrincipalStateSpec
	4, // 2: teleport.provisioning.v1.PrincipalState.status:type_name -> teleport.provisioning.v1.PrincipalStateStatus
	1, // 3: teleport.provisioning.v1.PrincipalStateSpec.principal_type:type_name -> teleport.provisioning.v1.PrincipalType
	0, // 4: teleport.provisioning.v1.PrincipalStateStatus.provisioning_state:type_name -> teleport.provisioning.v1.ProvisioningState
	6, // 5: teleport.provisioning.v1.PrincipalStateStatus.last_provisioned:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_teleport_provisioning_v1_provisioning_proto_init() }
func file_teleport_provisioning_v1_provisioning_proto_init() {
	if File_teleport_provisioning_v1_provisioning_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_provisioning_v1_provisioning_proto_rawDesc), len(file_teleport_provisioning_v1_provisioning_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_provisioning_v1_provisioning_proto_goTypes,
		DependencyIndexes: file_teleport_provisioning_v1_provisioning_proto_depIdxs,
		EnumInfos:         file_teleport_provisioning_v1_provisioning_proto_enumTypes,
		MessageInfos:      file_teleport_provisioning_v1_provisioning_proto_msgTypes,
	}.Build()
	File_teleport_provisioning_v1_provisioning_proto = out.File
	file_teleport_provisioning_v1_provisioning_proto_goTypes = nil
	file_teleport_provisioning_v1_provisioning_proto_depIdxs = nil
}
