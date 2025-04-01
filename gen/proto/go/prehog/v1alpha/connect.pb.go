//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: prehog/v1alpha/connect.proto

package prehogv1alpha

import (
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

type ConnectClusterLoginEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// local/github/saml/oidc/passwordless
	ConnectorType string `protobuf:"bytes,3,opt,name=connector_type,json=connectorType,proto3" json:"connector_type,omitempty"`
	Arch          string `protobuf:"bytes,4,opt,name=arch,proto3" json:"arch,omitempty"`
	Os            string `protobuf:"bytes,5,opt,name=os,proto3" json:"os,omitempty"`
	OsVersion     string `protobuf:"bytes,6,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	AppVersion    string `protobuf:"bytes,7,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectClusterLoginEvent) Reset() {
	*x = ConnectClusterLoginEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectClusterLoginEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectClusterLoginEvent) ProtoMessage() {}

func (x *ConnectClusterLoginEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectClusterLoginEvent.ProtoReflect.Descriptor instead.
func (*ConnectClusterLoginEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectClusterLoginEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectClusterLoginEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ConnectClusterLoginEvent) GetConnectorType() string {
	if x != nil {
		return x.ConnectorType
	}
	return ""
}

func (x *ConnectClusterLoginEvent) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *ConnectClusterLoginEvent) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *ConnectClusterLoginEvent) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *ConnectClusterLoginEvent) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

type ConnectProtocolUseEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// one of ssh/db/kube
	Protocol string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// origin denotes which part of Connect UI was used to access a resource.
	// One of resource_table/search_bar/connection_list/reopened_session/vnet (optional for backwards
	// compatibility).
	Origin string `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	// access_through describes whether a resource was accessed by speaking to the proxy service
	// directly, through a local proxy or through VNet.
	// One of proxy_service/local_proxy/vnet (optional for backwards compatibility).
	AccessThrough string `protobuf:"bytes,5,opt,name=access_through,json=accessThrough,proto3" json:"access_through,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectProtocolUseEvent) Reset() {
	*x = ConnectProtocolUseEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectProtocolUseEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectProtocolUseEvent) ProtoMessage() {}

func (x *ConnectProtocolUseEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectProtocolUseEvent.ProtoReflect.Descriptor instead.
func (*ConnectProtocolUseEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectProtocolUseEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectProtocolUseEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ConnectProtocolUseEvent) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ConnectProtocolUseEvent) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *ConnectProtocolUseEvent) GetAccessThrough() string {
	if x != nil {
		return x.AccessThrough
	}
	return ""
}

type ConnectAccessRequestCreateEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// one of role/resource
	Kind          string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectAccessRequestCreateEvent) Reset() {
	*x = ConnectAccessRequestCreateEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectAccessRequestCreateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectAccessRequestCreateEvent) ProtoMessage() {}

func (x *ConnectAccessRequestCreateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectAccessRequestCreateEvent.ProtoReflect.Descriptor instead.
func (*ConnectAccessRequestCreateEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectAccessRequestCreateEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectAccessRequestCreateEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ConnectAccessRequestCreateEvent) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type ConnectAccessRequestReviewEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectAccessRequestReviewEvent) Reset() {
	*x = ConnectAccessRequestReviewEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectAccessRequestReviewEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectAccessRequestReviewEvent) ProtoMessage() {}

func (x *ConnectAccessRequestReviewEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectAccessRequestReviewEvent.ProtoReflect.Descriptor instead.
func (*ConnectAccessRequestReviewEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectAccessRequestReviewEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectAccessRequestReviewEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type ConnectAccessRequestAssumeRoleEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectAccessRequestAssumeRoleEvent) Reset() {
	*x = ConnectAccessRequestAssumeRoleEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectAccessRequestAssumeRoleEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectAccessRequestAssumeRoleEvent) ProtoMessage() {}

func (x *ConnectAccessRequestAssumeRoleEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectAccessRequestAssumeRoleEvent.ProtoReflect.Descriptor instead.
func (*ConnectAccessRequestAssumeRoleEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectAccessRequestAssumeRoleEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectAccessRequestAssumeRoleEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type ConnectFileTransferRunEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	IsUpload      bool   `protobuf:"varint,3,opt,name=is_upload,json=isUpload,proto3" json:"is_upload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectFileTransferRunEvent) Reset() {
	*x = ConnectFileTransferRunEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectFileTransferRunEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectFileTransferRunEvent) ProtoMessage() {}

func (x *ConnectFileTransferRunEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectFileTransferRunEvent.ProtoReflect.Descriptor instead.
func (*ConnectFileTransferRunEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectFileTransferRunEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectFileTransferRunEvent) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ConnectFileTransferRunEvent) GetIsUpload() bool {
	if x != nil {
		return x.IsUpload
	}
	return false
}

type ConnectUserJobRoleUpdateEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Software Engineer, DevOps Engineer etc.
	JobRole       string `protobuf:"bytes,1,opt,name=job_role,json=jobRole,proto3" json:"job_role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectUserJobRoleUpdateEvent) Reset() {
	*x = ConnectUserJobRoleUpdateEvent{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectUserJobRoleUpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectUserJobRoleUpdateEvent) ProtoMessage() {}

func (x *ConnectUserJobRoleUpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectUserJobRoleUpdateEvent.ProtoReflect.Descriptor instead.
func (*ConnectUserJobRoleUpdateEvent) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{6}
}

func (x *ConnectUserJobRoleUpdateEvent) GetJobRole() string {
	if x != nil {
		return x.JobRole
	}
	return ""
}

type ConnectConnectMyComputerSetup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Success  bool   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	// empty when success is true
	FailedStep    string `protobuf:"bytes,4,opt,name=failed_step,json=failedStep,proto3" json:"failed_step,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectConnectMyComputerSetup) Reset() {
	*x = ConnectConnectMyComputerSetup{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectConnectMyComputerSetup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectConnectMyComputerSetup) ProtoMessage() {}

func (x *ConnectConnectMyComputerSetup) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectConnectMyComputerSetup.ProtoReflect.Descriptor instead.
func (*ConnectConnectMyComputerSetup) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{7}
}

func (x *ConnectConnectMyComputerSetup) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectConnectMyComputerSetup) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ConnectConnectMyComputerSetup) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ConnectConnectMyComputerSetup) GetFailedStep() string {
	if x != nil {
		return x.FailedStep
	}
	return ""
}

type ConnectConnectMyComputerAgentStart struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// anonymized
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// anonymized
	UserName      string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectConnectMyComputerAgentStart) Reset() {
	*x = ConnectConnectMyComputerAgentStart{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectConnectMyComputerAgentStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectConnectMyComputerAgentStart) ProtoMessage() {}

func (x *ConnectConnectMyComputerAgentStart) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectConnectMyComputerAgentStart.ProtoReflect.Descriptor instead.
func (*ConnectConnectMyComputerAgentStart) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{8}
}

func (x *ConnectConnectMyComputerAgentStart) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ConnectConnectMyComputerAgentStart) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type SubmitConnectEventRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// UUID
	DistinctId string `protobuf:"bytes,1,opt,name=distinct_id,json=distinctId,proto3" json:"distinct_id,omitempty"`
	// optional, will default to the ingest time if unset
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//
	//	*SubmitConnectEventRequest_ClusterLogin
	//	*SubmitConnectEventRequest_ProtocolUse
	//	*SubmitConnectEventRequest_AccessRequestCreate
	//	*SubmitConnectEventRequest_AccessRequestReview
	//	*SubmitConnectEventRequest_AccessRequestAssumeRole
	//	*SubmitConnectEventRequest_FileTransferRun
	//	*SubmitConnectEventRequest_UserJobRoleUpdate
	//	*SubmitConnectEventRequest_ConnectMyComputerSetup
	//	*SubmitConnectEventRequest_ConnectMyComputerAgentStart
	Event         isSubmitConnectEventRequest_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitConnectEventRequest) Reset() {
	*x = SubmitConnectEventRequest{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitConnectEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitConnectEventRequest) ProtoMessage() {}

func (x *SubmitConnectEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitConnectEventRequest.ProtoReflect.Descriptor instead.
func (*SubmitConnectEventRequest) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{9}
}

func (x *SubmitConnectEventRequest) GetDistinctId() string {
	if x != nil {
		return x.DistinctId
	}
	return ""
}

func (x *SubmitConnectEventRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetEvent() isSubmitConnectEventRequest_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetClusterLogin() *ConnectClusterLoginEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_ClusterLogin); ok {
			return x.ClusterLogin
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetProtocolUse() *ConnectProtocolUseEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_ProtocolUse); ok {
			return x.ProtocolUse
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetAccessRequestCreate() *ConnectAccessRequestCreateEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_AccessRequestCreate); ok {
			return x.AccessRequestCreate
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetAccessRequestReview() *ConnectAccessRequestReviewEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_AccessRequestReview); ok {
			return x.AccessRequestReview
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetAccessRequestAssumeRole() *ConnectAccessRequestAssumeRoleEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_AccessRequestAssumeRole); ok {
			return x.AccessRequestAssumeRole
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetFileTransferRun() *ConnectFileTransferRunEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_FileTransferRun); ok {
			return x.FileTransferRun
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetUserJobRoleUpdate() *ConnectUserJobRoleUpdateEvent {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_UserJobRoleUpdate); ok {
			return x.UserJobRoleUpdate
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetConnectMyComputerSetup() *ConnectConnectMyComputerSetup {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_ConnectMyComputerSetup); ok {
			return x.ConnectMyComputerSetup
		}
	}
	return nil
}

func (x *SubmitConnectEventRequest) GetConnectMyComputerAgentStart() *ConnectConnectMyComputerAgentStart {
	if x != nil {
		if x, ok := x.Event.(*SubmitConnectEventRequest_ConnectMyComputerAgentStart); ok {
			return x.ConnectMyComputerAgentStart
		}
	}
	return nil
}

type isSubmitConnectEventRequest_Event interface {
	isSubmitConnectEventRequest_Event()
}

type SubmitConnectEventRequest_ClusterLogin struct {
	ClusterLogin *ConnectClusterLoginEvent `protobuf:"bytes,3,opt,name=cluster_login,json=clusterLogin,proto3,oneof"`
}

type SubmitConnectEventRequest_ProtocolUse struct {
	ProtocolUse *ConnectProtocolUseEvent `protobuf:"bytes,4,opt,name=protocol_use,json=protocolUse,proto3,oneof"`
}

type SubmitConnectEventRequest_AccessRequestCreate struct {
	AccessRequestCreate *ConnectAccessRequestCreateEvent `protobuf:"bytes,5,opt,name=access_request_create,json=accessRequestCreate,proto3,oneof"`
}

type SubmitConnectEventRequest_AccessRequestReview struct {
	AccessRequestReview *ConnectAccessRequestReviewEvent `protobuf:"bytes,6,opt,name=access_request_review,json=accessRequestReview,proto3,oneof"`
}

type SubmitConnectEventRequest_AccessRequestAssumeRole struct {
	AccessRequestAssumeRole *ConnectAccessRequestAssumeRoleEvent `protobuf:"bytes,7,opt,name=access_request_assume_role,json=accessRequestAssumeRole,proto3,oneof"`
}

type SubmitConnectEventRequest_FileTransferRun struct {
	FileTransferRun *ConnectFileTransferRunEvent `protobuf:"bytes,8,opt,name=file_transfer_run,json=fileTransferRun,proto3,oneof"`
}

type SubmitConnectEventRequest_UserJobRoleUpdate struct {
	UserJobRoleUpdate *ConnectUserJobRoleUpdateEvent `protobuf:"bytes,9,opt,name=user_job_role_update,json=userJobRoleUpdate,proto3,oneof"`
}

type SubmitConnectEventRequest_ConnectMyComputerSetup struct {
	ConnectMyComputerSetup *ConnectConnectMyComputerSetup `protobuf:"bytes,10,opt,name=connect_my_computer_setup,json=connectMyComputerSetup,proto3,oneof"`
}

type SubmitConnectEventRequest_ConnectMyComputerAgentStart struct {
	ConnectMyComputerAgentStart *ConnectConnectMyComputerAgentStart `protobuf:"bytes,11,opt,name=connect_my_computer_agent_start,json=connectMyComputerAgentStart,proto3,oneof"`
}

func (*SubmitConnectEventRequest_ClusterLogin) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_ProtocolUse) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_AccessRequestCreate) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_AccessRequestReview) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_AccessRequestAssumeRole) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_FileTransferRun) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_UserJobRoleUpdate) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_ConnectMyComputerSetup) isSubmitConnectEventRequest_Event() {}

func (*SubmitConnectEventRequest_ConnectMyComputerAgentStart) isSubmitConnectEventRequest_Event() {}

type SubmitConnectEventResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitConnectEventResponse) Reset() {
	*x = SubmitConnectEventResponse{}
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitConnectEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitConnectEventResponse) ProtoMessage() {}

func (x *SubmitConnectEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prehog_v1alpha_connect_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitConnectEventResponse.ProtoReflect.Descriptor instead.
func (*SubmitConnectEventResponse) Descriptor() ([]byte, []int) {
	return file_prehog_v1alpha_connect_proto_rawDescGZIP(), []int{10}
}

var File_prehog_v1alpha_connect_proto protoreflect.FileDescriptor

const file_prehog_v1alpha_connect_proto_rawDesc = "" +
	"\n" +
	"\x1cprehog/v1alpha/connect.proto\x12\x0eprehog.v1alpha\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe5\x01\n" +
	"\x18ConnectClusterLoginEvent\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12%\n" +
	"\x0econnector_type\x18\x03 \x01(\tR\rconnectorType\x12\x12\n" +
	"\x04arch\x18\x04 \x01(\tR\x04arch\x12\x0e\n" +
	"\x02os\x18\x05 \x01(\tR\x02os\x12\x1d\n" +
	"\n" +
	"os_version\x18\x06 \x01(\tR\tosVersion\x12\x1f\n" +
	"\vapp_version\x18\a \x01(\tR\n" +
	"appVersion\"\xb4\x01\n" +
	"\x17ConnectProtocolUseEvent\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12\x1a\n" +
	"\bprotocol\x18\x03 \x01(\tR\bprotocol\x12\x16\n" +
	"\x06origin\x18\x04 \x01(\tR\x06origin\x12%\n" +
	"\x0eaccess_through\x18\x05 \x01(\tR\raccessThrough\"u\n" +
	"\x1fConnectAccessRequestCreateEvent\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12\x12\n" +
	"\x04kind\x18\x03 \x01(\tR\x04kind\"a\n" +
	"\x1fConnectAccessRequestReviewEvent\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\"e\n" +
	"#ConnectAccessRequestAssumeRoleEvent\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\"z\n" +
	"\x1bConnectFileTransferRunEvent\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12\x1b\n" +
	"\tis_upload\x18\x03 \x01(\bR\bisUpload\":\n" +
	"\x1dConnectUserJobRoleUpdateEvent\x12\x19\n" +
	"\bjob_role\x18\x01 \x01(\tR\ajobRole\"\x9a\x01\n" +
	"\x1dConnectConnectMyComputerSetup\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\x12\x18\n" +
	"\asuccess\x18\x03 \x01(\bR\asuccess\x12\x1f\n" +
	"\vfailed_step\x18\x04 \x01(\tR\n" +
	"failedStep\"d\n" +
	"\"ConnectConnectMyComputerAgentStart\x12!\n" +
	"\fcluster_name\x18\x01 \x01(\tR\vclusterName\x12\x1b\n" +
	"\tuser_name\x18\x02 \x01(\tR\buserName\"\x85\b\n" +
	"\x19SubmitConnectEventRequest\x12\x1f\n" +
	"\vdistinct_id\x18\x01 \x01(\tR\n" +
	"distinctId\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12O\n" +
	"\rcluster_login\x18\x03 \x01(\v2(.prehog.v1alpha.ConnectClusterLoginEventH\x00R\fclusterLogin\x12L\n" +
	"\fprotocol_use\x18\x04 \x01(\v2'.prehog.v1alpha.ConnectProtocolUseEventH\x00R\vprotocolUse\x12e\n" +
	"\x15access_request_create\x18\x05 \x01(\v2/.prehog.v1alpha.ConnectAccessRequestCreateEventH\x00R\x13accessRequestCreate\x12e\n" +
	"\x15access_request_review\x18\x06 \x01(\v2/.prehog.v1alpha.ConnectAccessRequestReviewEventH\x00R\x13accessRequestReview\x12r\n" +
	"\x1aaccess_request_assume_role\x18\a \x01(\v23.prehog.v1alpha.ConnectAccessRequestAssumeRoleEventH\x00R\x17accessRequestAssumeRole\x12Y\n" +
	"\x11file_transfer_run\x18\b \x01(\v2+.prehog.v1alpha.ConnectFileTransferRunEventH\x00R\x0ffileTransferRun\x12`\n" +
	"\x14user_job_role_update\x18\t \x01(\v2-.prehog.v1alpha.ConnectUserJobRoleUpdateEventH\x00R\x11userJobRoleUpdate\x12j\n" +
	"\x19connect_my_computer_setup\x18\n" +
	" \x01(\v2-.prehog.v1alpha.ConnectConnectMyComputerSetupH\x00R\x16connectMyComputerSetup\x12z\n" +
	"\x1fconnect_my_computer_agent_start\x18\v \x01(\v22.prehog.v1alpha.ConnectConnectMyComputerAgentStartH\x00R\x1bconnectMyComputerAgentStartB\a\n" +
	"\x05event\"\x1c\n" +
	"\x1aSubmitConnectEventResponse2\x88\x01\n" +
	"\x17ConnectReportingService\x12m\n" +
	"\x12SubmitConnectEvent\x12).prehog.v1alpha.SubmitConnectEventRequest\x1a*.prehog.v1alpha.SubmitConnectEventResponse\"\x00B\xc8\x01\n" +
	"\x12com.prehog.v1alphaB\fConnectProtoP\x01ZKgithub.com/gravitational/teleport/gen/proto/go/prehog/v1alpha;prehogv1alpha\xa2\x02\x03PXX\xaa\x02\x0ePrehog.V1alpha\xca\x02\x0ePrehog\\V1alpha\xe2\x02\x1aPrehog\\V1alpha\\GPBMetadata\xea\x02\x0fPrehog::V1alphab\x06proto3"

var (
	file_prehog_v1alpha_connect_proto_rawDescOnce sync.Once
	file_prehog_v1alpha_connect_proto_rawDescData []byte
)

func file_prehog_v1alpha_connect_proto_rawDescGZIP() []byte {
	file_prehog_v1alpha_connect_proto_rawDescOnce.Do(func() {
		file_prehog_v1alpha_connect_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_prehog_v1alpha_connect_proto_rawDesc), len(file_prehog_v1alpha_connect_proto_rawDesc)))
	})
	return file_prehog_v1alpha_connect_proto_rawDescData
}

var file_prehog_v1alpha_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_prehog_v1alpha_connect_proto_goTypes = []any{
	(*ConnectClusterLoginEvent)(nil),            // 0: prehog.v1alpha.ConnectClusterLoginEvent
	(*ConnectProtocolUseEvent)(nil),             // 1: prehog.v1alpha.ConnectProtocolUseEvent
	(*ConnectAccessRequestCreateEvent)(nil),     // 2: prehog.v1alpha.ConnectAccessRequestCreateEvent
	(*ConnectAccessRequestReviewEvent)(nil),     // 3: prehog.v1alpha.ConnectAccessRequestReviewEvent
	(*ConnectAccessRequestAssumeRoleEvent)(nil), // 4: prehog.v1alpha.ConnectAccessRequestAssumeRoleEvent
	(*ConnectFileTransferRunEvent)(nil),         // 5: prehog.v1alpha.ConnectFileTransferRunEvent
	(*ConnectUserJobRoleUpdateEvent)(nil),       // 6: prehog.v1alpha.ConnectUserJobRoleUpdateEvent
	(*ConnectConnectMyComputerSetup)(nil),       // 7: prehog.v1alpha.ConnectConnectMyComputerSetup
	(*ConnectConnectMyComputerAgentStart)(nil),  // 8: prehog.v1alpha.ConnectConnectMyComputerAgentStart
	(*SubmitConnectEventRequest)(nil),           // 9: prehog.v1alpha.SubmitConnectEventRequest
	(*SubmitConnectEventResponse)(nil),          // 10: prehog.v1alpha.SubmitConnectEventResponse
	(*timestamppb.Timestamp)(nil),               // 11: google.protobuf.Timestamp
}
var file_prehog_v1alpha_connect_proto_depIdxs = []int32{
	11, // 0: prehog.v1alpha.SubmitConnectEventRequest.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 1: prehog.v1alpha.SubmitConnectEventRequest.cluster_login:type_name -> prehog.v1alpha.ConnectClusterLoginEvent
	1,  // 2: prehog.v1alpha.SubmitConnectEventRequest.protocol_use:type_name -> prehog.v1alpha.ConnectProtocolUseEvent
	2,  // 3: prehog.v1alpha.SubmitConnectEventRequest.access_request_create:type_name -> prehog.v1alpha.ConnectAccessRequestCreateEvent
	3,  // 4: prehog.v1alpha.SubmitConnectEventRequest.access_request_review:type_name -> prehog.v1alpha.ConnectAccessRequestReviewEvent
	4,  // 5: prehog.v1alpha.SubmitConnectEventRequest.access_request_assume_role:type_name -> prehog.v1alpha.ConnectAccessRequestAssumeRoleEvent
	5,  // 6: prehog.v1alpha.SubmitConnectEventRequest.file_transfer_run:type_name -> prehog.v1alpha.ConnectFileTransferRunEvent
	6,  // 7: prehog.v1alpha.SubmitConnectEventRequest.user_job_role_update:type_name -> prehog.v1alpha.ConnectUserJobRoleUpdateEvent
	7,  // 8: prehog.v1alpha.SubmitConnectEventRequest.connect_my_computer_setup:type_name -> prehog.v1alpha.ConnectConnectMyComputerSetup
	8,  // 9: prehog.v1alpha.SubmitConnectEventRequest.connect_my_computer_agent_start:type_name -> prehog.v1alpha.ConnectConnectMyComputerAgentStart
	9,  // 10: prehog.v1alpha.ConnectReportingService.SubmitConnectEvent:input_type -> prehog.v1alpha.SubmitConnectEventRequest
	10, // 11: prehog.v1alpha.ConnectReportingService.SubmitConnectEvent:output_type -> prehog.v1alpha.SubmitConnectEventResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_prehog_v1alpha_connect_proto_init() }
func file_prehog_v1alpha_connect_proto_init() {
	if File_prehog_v1alpha_connect_proto != nil {
		return
	}
	file_prehog_v1alpha_connect_proto_msgTypes[9].OneofWrappers = []any{
		(*SubmitConnectEventRequest_ClusterLogin)(nil),
		(*SubmitConnectEventRequest_ProtocolUse)(nil),
		(*SubmitConnectEventRequest_AccessRequestCreate)(nil),
		(*SubmitConnectEventRequest_AccessRequestReview)(nil),
		(*SubmitConnectEventRequest_AccessRequestAssumeRole)(nil),
		(*SubmitConnectEventRequest_FileTransferRun)(nil),
		(*SubmitConnectEventRequest_UserJobRoleUpdate)(nil),
		(*SubmitConnectEventRequest_ConnectMyComputerSetup)(nil),
		(*SubmitConnectEventRequest_ConnectMyComputerAgentStart)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_prehog_v1alpha_connect_proto_rawDesc), len(file_prehog_v1alpha_connect_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prehog_v1alpha_connect_proto_goTypes,
		DependencyIndexes: file_prehog_v1alpha_connect_proto_depIdxs,
		MessageInfos:      file_prehog_v1alpha_connect_proto_msgTypes,
	}.Build()
	File_prehog_v1alpha_connect_proto = out.File
	file_prehog_v1alpha_connect_proto_goTypes = nil
	file_prehog_v1alpha_connect_proto_depIdxs = nil
}
