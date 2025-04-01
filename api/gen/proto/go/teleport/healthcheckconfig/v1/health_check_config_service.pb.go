// Copyright 2025 Gravitational, Inc.
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
// source: teleport/healthcheckconfig/v1/health_check_config_service.proto

package healthcheckconfigv1

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

// Request for CreateHealthCheckConfig.
type CreateHealthCheckConfigRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Config is the resource to create.
	Config        *HealthCheckConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateHealthCheckConfigRequest) Reset() {
	*x = CreateHealthCheckConfigRequest{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHealthCheckConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHealthCheckConfigRequest) ProtoMessage() {}

func (x *CreateHealthCheckConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHealthCheckConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateHealthCheckConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHealthCheckConfigRequest) GetConfig() *HealthCheckConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// GetHealthCheckConfigRequest is a request for GetHealthCheckConfig.
type GetHealthCheckConfigRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the name of the HealthCheckConfig to retrieve.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHealthCheckConfigRequest) Reset() {
	*x = GetHealthCheckConfigRequest{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHealthCheckConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthCheckConfigRequest) ProtoMessage() {}

func (x *GetHealthCheckConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthCheckConfigRequest.ProtoReflect.Descriptor instead.
func (*GetHealthCheckConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHealthCheckConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ListHealthCheckConfigsRequest is the request for ListHealthCheckConfigs.
type ListHealthCheckConfigsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PageSize is the maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// PageToken is the page token value returned from a prior list request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHealthCheckConfigsRequest) Reset() {
	*x = ListHealthCheckConfigsRequest{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHealthCheckConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHealthCheckConfigsRequest) ProtoMessage() {}

func (x *ListHealthCheckConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHealthCheckConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListHealthCheckConfigsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListHealthCheckConfigsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListHealthCheckConfigsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// ListHealthCheckConfigsResponse is the response from ListHealthCheckConfigs.
type ListHealthCheckConfigsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Configs is a page of health check configs.
	Configs []*HealthCheckConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	// NextPageToken is the token that can be used to retrieve the next page of
	// results or empty if there are no more pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHealthCheckConfigsResponse) Reset() {
	*x = ListHealthCheckConfigsResponse{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHealthCheckConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHealthCheckConfigsResponse) ProtoMessage() {}

func (x *ListHealthCheckConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHealthCheckConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListHealthCheckConfigsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListHealthCheckConfigsResponse) GetConfigs() []*HealthCheckConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *ListHealthCheckConfigsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request for UpdateHealthCheckConfig.
type UpdateHealthCheckConfigRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Config is the resource to update.
	Config        *HealthCheckConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateHealthCheckConfigRequest) Reset() {
	*x = UpdateHealthCheckConfigRequest{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateHealthCheckConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHealthCheckConfigRequest) ProtoMessage() {}

func (x *UpdateHealthCheckConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHealthCheckConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateHealthCheckConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateHealthCheckConfigRequest) GetConfig() *HealthCheckConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Request for UpsertHealthCheckConfig.
type UpsertHealthCheckConfigRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Config is the resource to upsert.
	Config        *HealthCheckConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertHealthCheckConfigRequest) Reset() {
	*x = UpsertHealthCheckConfigRequest{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertHealthCheckConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertHealthCheckConfigRequest) ProtoMessage() {}

func (x *UpsertHealthCheckConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertHealthCheckConfigRequest.ProtoReflect.Descriptor instead.
func (*UpsertHealthCheckConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertHealthCheckConfigRequest) GetConfig() *HealthCheckConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// Request for DeleteHealthCheckConfig.
type DeleteHealthCheckConfigRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the name of the HealthCheckConfig to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteHealthCheckConfigRequest) Reset() {
	*x = DeleteHealthCheckConfigRequest{}
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteHealthCheckConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHealthCheckConfigRequest) ProtoMessage() {}

func (x *DeleteHealthCheckConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHealthCheckConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteHealthCheckConfigRequest) Descriptor() ([]byte, []int) {
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteHealthCheckConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_teleport_healthcheckconfig_v1_health_check_config_service_proto protoreflect.FileDescriptor

const file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDesc = "" +
	"\n" +
	"?teleport/healthcheckconfig/v1/health_check_config_service.proto\x12\x1dteleport.healthcheckconfig.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a7teleport/healthcheckconfig/v1/health_check_config.proto\"j\n" +
	"\x1eCreateHealthCheckConfigRequest\x12H\n" +
	"\x06config\x18\x01 \x01(\v20.teleport.healthcheckconfig.v1.HealthCheckConfigR\x06config\"1\n" +
	"\x1bGetHealthCheckConfigRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"[\n" +
	"\x1dListHealthCheckConfigsRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"\x94\x01\n" +
	"\x1eListHealthCheckConfigsResponse\x12J\n" +
	"\aconfigs\x18\x01 \x03(\v20.teleport.healthcheckconfig.v1.HealthCheckConfigR\aconfigs\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"j\n" +
	"\x1eUpdateHealthCheckConfigRequest\x12H\n" +
	"\x06config\x18\x01 \x01(\v20.teleport.healthcheckconfig.v1.HealthCheckConfigR\x06config\"j\n" +
	"\x1eUpsertHealthCheckConfigRequest\x12H\n" +
	"\x06config\x18\x01 \x01(\v20.teleport.healthcheckconfig.v1.HealthCheckConfigR\x06config\"4\n" +
	"\x1eDeleteHealthCheckConfigRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name2\xd2\x06\n" +
	"\x18HealthCheckConfigService\x12\x8a\x01\n" +
	"\x17CreateHealthCheckConfig\x12=.teleport.healthcheckconfig.v1.CreateHealthCheckConfigRequest\x1a0.teleport.healthcheckconfig.v1.HealthCheckConfig\x12\x84\x01\n" +
	"\x14GetHealthCheckConfig\x12:.teleport.healthcheckconfig.v1.GetHealthCheckConfigRequest\x1a0.teleport.healthcheckconfig.v1.HealthCheckConfig\x12\x95\x01\n" +
	"\x16ListHealthCheckConfigs\x12<.teleport.healthcheckconfig.v1.ListHealthCheckConfigsRequest\x1a=.teleport.healthcheckconfig.v1.ListHealthCheckConfigsResponse\x12\x8a\x01\n" +
	"\x17UpdateHealthCheckConfig\x12=.teleport.healthcheckconfig.v1.UpdateHealthCheckConfigRequest\x1a0.teleport.healthcheckconfig.v1.HealthCheckConfig\x12\x8a\x01\n" +
	"\x17UpsertHealthCheckConfig\x12=.teleport.healthcheckconfig.v1.UpsertHealthCheckConfigRequest\x1a0.teleport.healthcheckconfig.v1.HealthCheckConfig\x12p\n" +
	"\x17DeleteHealthCheckConfig\x12=.teleport.healthcheckconfig.v1.DeleteHealthCheckConfigRequest\x1a\x16.google.protobuf.EmptyBfZdgithub.com/gravitational/teleport/api/gen/proto/go/teleport/healthcheckconfig/v1;healthcheckconfigv1b\x06proto3"

var (
	file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescOnce sync.Once
	file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescData []byte
)

func file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescGZIP() []byte {
	file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescOnce.Do(func() {
		file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDesc), len(file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDesc)))
	})
	return file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDescData
}

var file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_teleport_healthcheckconfig_v1_health_check_config_service_proto_goTypes = []any{
	(*CreateHealthCheckConfigRequest)(nil), // 0: teleport.healthcheckconfig.v1.CreateHealthCheckConfigRequest
	(*GetHealthCheckConfigRequest)(nil),    // 1: teleport.healthcheckconfig.v1.GetHealthCheckConfigRequest
	(*ListHealthCheckConfigsRequest)(nil),  // 2: teleport.healthcheckconfig.v1.ListHealthCheckConfigsRequest
	(*ListHealthCheckConfigsResponse)(nil), // 3: teleport.healthcheckconfig.v1.ListHealthCheckConfigsResponse
	(*UpdateHealthCheckConfigRequest)(nil), // 4: teleport.healthcheckconfig.v1.UpdateHealthCheckConfigRequest
	(*UpsertHealthCheckConfigRequest)(nil), // 5: teleport.healthcheckconfig.v1.UpsertHealthCheckConfigRequest
	(*DeleteHealthCheckConfigRequest)(nil), // 6: teleport.healthcheckconfig.v1.DeleteHealthCheckConfigRequest
	(*HealthCheckConfig)(nil),              // 7: teleport.healthcheckconfig.v1.HealthCheckConfig
	(*emptypb.Empty)(nil),                  // 8: google.protobuf.Empty
}
var file_teleport_healthcheckconfig_v1_health_check_config_service_proto_depIdxs = []int32{
	7,  // 0: teleport.healthcheckconfig.v1.CreateHealthCheckConfigRequest.config:type_name -> teleport.healthcheckconfig.v1.HealthCheckConfig
	7,  // 1: teleport.healthcheckconfig.v1.ListHealthCheckConfigsResponse.configs:type_name -> teleport.healthcheckconfig.v1.HealthCheckConfig
	7,  // 2: teleport.healthcheckconfig.v1.UpdateHealthCheckConfigRequest.config:type_name -> teleport.healthcheckconfig.v1.HealthCheckConfig
	7,  // 3: teleport.healthcheckconfig.v1.UpsertHealthCheckConfigRequest.config:type_name -> teleport.healthcheckconfig.v1.HealthCheckConfig
	0,  // 4: teleport.healthcheckconfig.v1.HealthCheckConfigService.CreateHealthCheckConfig:input_type -> teleport.healthcheckconfig.v1.CreateHealthCheckConfigRequest
	1,  // 5: teleport.healthcheckconfig.v1.HealthCheckConfigService.GetHealthCheckConfig:input_type -> teleport.healthcheckconfig.v1.GetHealthCheckConfigRequest
	2,  // 6: teleport.healthcheckconfig.v1.HealthCheckConfigService.ListHealthCheckConfigs:input_type -> teleport.healthcheckconfig.v1.ListHealthCheckConfigsRequest
	4,  // 7: teleport.healthcheckconfig.v1.HealthCheckConfigService.UpdateHealthCheckConfig:input_type -> teleport.healthcheckconfig.v1.UpdateHealthCheckConfigRequest
	5,  // 8: teleport.healthcheckconfig.v1.HealthCheckConfigService.UpsertHealthCheckConfig:input_type -> teleport.healthcheckconfig.v1.UpsertHealthCheckConfigRequest
	6,  // 9: teleport.healthcheckconfig.v1.HealthCheckConfigService.DeleteHealthCheckConfig:input_type -> teleport.healthcheckconfig.v1.DeleteHealthCheckConfigRequest
	7,  // 10: teleport.healthcheckconfig.v1.HealthCheckConfigService.CreateHealthCheckConfig:output_type -> teleport.healthcheckconfig.v1.HealthCheckConfig
	7,  // 11: teleport.healthcheckconfig.v1.HealthCheckConfigService.GetHealthCheckConfig:output_type -> teleport.healthcheckconfig.v1.HealthCheckConfig
	3,  // 12: teleport.healthcheckconfig.v1.HealthCheckConfigService.ListHealthCheckConfigs:output_type -> teleport.healthcheckconfig.v1.ListHealthCheckConfigsResponse
	7,  // 13: teleport.healthcheckconfig.v1.HealthCheckConfigService.UpdateHealthCheckConfig:output_type -> teleport.healthcheckconfig.v1.HealthCheckConfig
	7,  // 14: teleport.healthcheckconfig.v1.HealthCheckConfigService.UpsertHealthCheckConfig:output_type -> teleport.healthcheckconfig.v1.HealthCheckConfig
	8,  // 15: teleport.healthcheckconfig.v1.HealthCheckConfigService.DeleteHealthCheckConfig:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_healthcheckconfig_v1_health_check_config_service_proto_init() }
func file_teleport_healthcheckconfig_v1_health_check_config_service_proto_init() {
	if File_teleport_healthcheckconfig_v1_health_check_config_service_proto != nil {
		return
	}
	file_teleport_healthcheckconfig_v1_health_check_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDesc), len(file_teleport_healthcheckconfig_v1_health_check_config_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_healthcheckconfig_v1_health_check_config_service_proto_goTypes,
		DependencyIndexes: file_teleport_healthcheckconfig_v1_health_check_config_service_proto_depIdxs,
		MessageInfos:      file_teleport_healthcheckconfig_v1_health_check_config_service_proto_msgTypes,
	}.Build()
	File_teleport_healthcheckconfig_v1_health_check_config_service_proto = out.File
	file_teleport_healthcheckconfig_v1_health_check_config_service_proto_goTypes = nil
	file_teleport_healthcheckconfig_v1_health_check_config_service_proto_depIdxs = nil
}
