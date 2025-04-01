//
// Teleport
// Copyright (C) 2024  Gravitational, Inc.
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
// source: teleport/notifications/v1/notifications_service.proto

package notificationsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// CreateUserNotificationRequest is the request for creating a user-specific notification.
type CreateUserNotificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// username is the username of the user the notification to create is for.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// notification is the notification to create.
	Notification  *Notification `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserNotificationRequest) Reset() {
	*x = CreateUserNotificationRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserNotificationRequest) ProtoMessage() {}

func (x *CreateUserNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateUserNotificationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserNotificationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserNotificationRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

// DeleteUserNotificationRequest is the request for deleting a user-specific notification.
type DeleteUserNotificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// username is the username of the user the notification to delete is for.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// notification_id is the ID of the notification to delete.
	NotificationId string `protobuf:"bytes,2,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeleteUserNotificationRequest) Reset() {
	*x = DeleteUserNotificationRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserNotificationRequest) ProtoMessage() {}

func (x *DeleteUserNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserNotificationRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserNotificationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteUserNotificationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeleteUserNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

// ListNotificationsRequest is the request for listing a user's notifications.
type ListNotificationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// page_size is the size of the page to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// page_token is the next_page_token value returned from a previous ListUserNotifications request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// filters specify search criteria to limit which notifications should be returned. If omitted, the default behavior will be to list all notifications.
	Filters       *NotificationFilters `protobuf:"bytes,3,opt,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNotificationsRequest) Reset() {
	*x = ListNotificationsRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsRequest) ProtoMessage() {}

func (x *ListNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsRequest.ProtoReflect.Descriptor instead.
func (*ListNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListNotificationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListNotificationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListNotificationsRequest) GetFilters() *NotificationFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

// NotificationFilters provide a mechanism to refine ListNotification results.
type NotificationFilters struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// username is the username of the user the notifications being listed are for.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// global_only is whether to only list global notifications (notifications capable of targetting multiple users).
	GlobalOnly bool `protobuf:"varint,2,opt,name=global_only,json=globalOnly,proto3" json:"global_only,omitempty"`
	// user_created_only is whether to only list user-created notifications (ie. notifications created by an admin via the tctl interface).
	UserCreatedOnly bool `protobuf:"varint,3,opt,name=user_created_only,json=userCreatedOnly,proto3" json:"user_created_only,omitempty"`
	// labels is used to request only notifications with specific labels.
	Labels        map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationFilters) Reset() {
	*x = NotificationFilters{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationFilters) ProtoMessage() {}

func (x *NotificationFilters) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationFilters.ProtoReflect.Descriptor instead.
func (*NotificationFilters) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationFilters) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NotificationFilters) GetGlobalOnly() bool {
	if x != nil {
		return x.GlobalOnly
	}
	return false
}

func (x *NotificationFilters) GetUserCreatedOnly() bool {
	if x != nil {
		return x.UserCreatedOnly
	}
	return false
}

func (x *NotificationFilters) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// ListNotificationsResponse is the response from listing a user's notifications.
type ListNotificationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// notifications is the notifications returned.
	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	// next_page_token contains the next page token to use as the start key for the next page of notifications.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// user_last_seen_notification_timestamp is the timestamp of the last notification the user has seen.
	UserLastSeenNotificationTimestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=user_last_seen_notification_timestamp,json=userLastSeenNotificationTimestamp,proto3" json:"user_last_seen_notification_timestamp,omitempty"`
	unknownFields                     protoimpl.UnknownFields
	sizeCache                         protoimpl.SizeCache
}

func (x *ListNotificationsResponse) Reset() {
	*x = ListNotificationsResponse{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotificationsResponse) ProtoMessage() {}

func (x *ListNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotificationsResponse.ProtoReflect.Descriptor instead.
func (*ListNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

func (x *ListNotificationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListNotificationsResponse) GetUserLastSeenNotificationTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.UserLastSeenNotificationTimestamp
	}
	return nil
}

// CreateGlobalNotificationRequest is the request for creating a global notification.
type CreateGlobalNotificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// global_notification is the global notification to create.
	GlobalNotification *GlobalNotification `protobuf:"bytes,1,opt,name=global_notification,json=globalNotification,proto3" json:"global_notification,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CreateGlobalNotificationRequest) Reset() {
	*x = CreateGlobalNotificationRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGlobalNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGlobalNotificationRequest) ProtoMessage() {}

func (x *CreateGlobalNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGlobalNotificationRequest.ProtoReflect.Descriptor instead.
func (*CreateGlobalNotificationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGlobalNotificationRequest) GetGlobalNotification() *GlobalNotification {
	if x != nil {
		return x.GlobalNotification
	}
	return nil
}

// DeleteGlobalNotificationRequest is the request for deleting a global notification.
type DeleteGlobalNotificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// notification_id is the ID of the notification to delete.
	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeleteGlobalNotificationRequest) Reset() {
	*x = DeleteGlobalNotificationRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGlobalNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGlobalNotificationRequest) ProtoMessage() {}

func (x *DeleteGlobalNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGlobalNotificationRequest.ProtoReflect.Descriptor instead.
func (*DeleteGlobalNotificationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteGlobalNotificationRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

// UpsertUserNotificationStateRequest is the request for creating or updating a user notification state.
type UpsertUserNotificationStateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// username is the username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// user_notification_state is the user notification state to create.
	UserNotificationState *UserNotificationState `protobuf:"bytes,2,opt,name=user_notification_state,json=userNotificationState,proto3" json:"user_notification_state,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *UpsertUserNotificationStateRequest) Reset() {
	*x = UpsertUserNotificationStateRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertUserNotificationStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserNotificationStateRequest) ProtoMessage() {}

func (x *UpsertUserNotificationStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserNotificationStateRequest.ProtoReflect.Descriptor instead.
func (*UpsertUserNotificationStateRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpsertUserNotificationStateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpsertUserNotificationStateRequest) GetUserNotificationState() *UserNotificationState {
	if x != nil {
		return x.UserNotificationState
	}
	return nil
}

// UpsertUserLastSeenNotificationRequest is the request for creating or updating a user's last seen notification.
type UpsertUserLastSeenNotificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// username is the username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// user_notification_state is the updated user last seen notification item.
	UserLastSeenNotification *UserLastSeenNotification `protobuf:"bytes,2,opt,name=user_last_seen_notification,json=userLastSeenNotification,proto3" json:"user_last_seen_notification,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *UpsertUserLastSeenNotificationRequest) Reset() {
	*x = UpsertUserLastSeenNotificationRequest{}
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertUserLastSeenNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserLastSeenNotificationRequest) ProtoMessage() {}

func (x *UpsertUserLastSeenNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_notifications_v1_notifications_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserLastSeenNotificationRequest.ProtoReflect.Descriptor instead.
func (*UpsertUserLastSeenNotificationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpsertUserLastSeenNotificationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpsertUserLastSeenNotificationRequest) GetUserLastSeenNotification() *UserLastSeenNotification {
	if x != nil {
		return x.UserLastSeenNotification
	}
	return nil
}

var File_teleport_notifications_v1_notifications_service_proto protoreflect.FileDescriptor

const file_teleport_notifications_v1_notifications_service_proto_rawDesc = "" +
	"\n" +
	"5teleport/notifications/v1/notifications_service.proto\x12\x19teleport.notifications.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a-teleport/notifications/v1/notifications.proto\"\x88\x01\n" +
	"\x1dCreateUserNotificationRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12K\n" +
	"\fnotification\x18\x02 \x01(\v2'.teleport.notifications.v1.NotificationR\fnotification\"d\n" +
	"\x1dDeleteUserNotificationRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12'\n" +
	"\x0fnotification_id\x18\x02 \x01(\tR\x0enotificationId\"\xa0\x01\n" +
	"\x18ListNotificationsRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\x12H\n" +
	"\afilters\x18\x03 \x01(\v2..teleport.notifications.v1.NotificationFiltersR\afilters\"\x8d\x02\n" +
	"\x13NotificationFilters\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1f\n" +
	"\vglobal_only\x18\x02 \x01(\bR\n" +
	"globalOnly\x12*\n" +
	"\x11user_created_only\x18\x03 \x01(\bR\x0fuserCreatedOnly\x12R\n" +
	"\x06labels\x18\x04 \x03(\v2:.teleport.notifications.v1.NotificationFilters.LabelsEntryR\x06labels\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x80\x02\n" +
	"\x19ListNotificationsResponse\x12M\n" +
	"\rnotifications\x18\x01 \x03(\v2'.teleport.notifications.v1.NotificationR\rnotifications\x12&\n" +
	"\x0fnext_page_token\x18\x03 \x01(\tR\rnextPageToken\x12l\n" +
	"%user_last_seen_notification_timestamp\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR!userLastSeenNotificationTimestamp\"\x81\x01\n" +
	"\x1fCreateGlobalNotificationRequest\x12^\n" +
	"\x13global_notification\x18\x01 \x01(\v2-.teleport.notifications.v1.GlobalNotificationR\x12globalNotification\"J\n" +
	"\x1fDeleteGlobalNotificationRequest\x12'\n" +
	"\x0fnotification_id\x18\x01 \x01(\tR\x0enotificationId\"\xaa\x01\n" +
	"\"UpsertUserNotificationStateRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12h\n" +
	"\x17user_notification_state\x18\x02 \x01(\v20.teleport.notifications.v1.UserNotificationStateR\x15userNotificationState\"\xb7\x01\n" +
	"%UpsertUserLastSeenNotificationRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12r\n" +
	"\x1buser_last_seen_notification\x18\x02 \x01(\v23.teleport.notifications.v1.UserLastSeenNotificationR\x18userLastSeenNotification2\xa1\a\n" +
	"\x13NotificationService\x12{\n" +
	"\x16CreateUserNotification\x128.teleport.notifications.v1.CreateUserNotificationRequest\x1a'.teleport.notifications.v1.Notification\x12j\n" +
	"\x16DeleteUserNotification\x128.teleport.notifications.v1.DeleteUserNotificationRequest\x1a\x16.google.protobuf.Empty\x12\x85\x01\n" +
	"\x18CreateGlobalNotification\x12:.teleport.notifications.v1.CreateGlobalNotificationRequest\x1a-.teleport.notifications.v1.GlobalNotification\x12n\n" +
	"\x18DeleteGlobalNotification\x12:.teleport.notifications.v1.DeleteGlobalNotificationRequest\x1a\x16.google.protobuf.Empty\x12~\n" +
	"\x11ListNotifications\x123.teleport.notifications.v1.ListNotificationsRequest\x1a4.teleport.notifications.v1.ListNotificationsResponse\x12\x8e\x01\n" +
	"\x1bUpsertUserNotificationState\x12=.teleport.notifications.v1.UpsertUserNotificationStateRequest\x1a0.teleport.notifications.v1.UserNotificationState\x12\x97\x01\n" +
	"\x1eUpsertUserLastSeenNotification\x12@.teleport.notifications.v1.UpsertUserLastSeenNotificationRequest\x1a3.teleport.notifications.v1.UserLastSeenNotificationB^Z\\github.com/gravitational/teleport/api/gen/proto/go/teleport/notifications/v1;notificationsv1b\x06proto3"

var (
	file_teleport_notifications_v1_notifications_service_proto_rawDescOnce sync.Once
	file_teleport_notifications_v1_notifications_service_proto_rawDescData []byte
)

func file_teleport_notifications_v1_notifications_service_proto_rawDescGZIP() []byte {
	file_teleport_notifications_v1_notifications_service_proto_rawDescOnce.Do(func() {
		file_teleport_notifications_v1_notifications_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_notifications_v1_notifications_service_proto_rawDesc), len(file_teleport_notifications_v1_notifications_service_proto_rawDesc)))
	})
	return file_teleport_notifications_v1_notifications_service_proto_rawDescData
}

var file_teleport_notifications_v1_notifications_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_teleport_notifications_v1_notifications_service_proto_goTypes = []any{
	(*CreateUserNotificationRequest)(nil),         // 0: teleport.notifications.v1.CreateUserNotificationRequest
	(*DeleteUserNotificationRequest)(nil),         // 1: teleport.notifications.v1.DeleteUserNotificationRequest
	(*ListNotificationsRequest)(nil),              // 2: teleport.notifications.v1.ListNotificationsRequest
	(*NotificationFilters)(nil),                   // 3: teleport.notifications.v1.NotificationFilters
	(*ListNotificationsResponse)(nil),             // 4: teleport.notifications.v1.ListNotificationsResponse
	(*CreateGlobalNotificationRequest)(nil),       // 5: teleport.notifications.v1.CreateGlobalNotificationRequest
	(*DeleteGlobalNotificationRequest)(nil),       // 6: teleport.notifications.v1.DeleteGlobalNotificationRequest
	(*UpsertUserNotificationStateRequest)(nil),    // 7: teleport.notifications.v1.UpsertUserNotificationStateRequest
	(*UpsertUserLastSeenNotificationRequest)(nil), // 8: teleport.notifications.v1.UpsertUserLastSeenNotificationRequest
	nil,                              // 9: teleport.notifications.v1.NotificationFilters.LabelsEntry
	(*Notification)(nil),             // 10: teleport.notifications.v1.Notification
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
	(*GlobalNotification)(nil),       // 12: teleport.notifications.v1.GlobalNotification
	(*UserNotificationState)(nil),    // 13: teleport.notifications.v1.UserNotificationState
	(*UserLastSeenNotification)(nil), // 14: teleport.notifications.v1.UserLastSeenNotification
	(*emptypb.Empty)(nil),            // 15: google.protobuf.Empty
}
var file_teleport_notifications_v1_notifications_service_proto_depIdxs = []int32{
	10, // 0: teleport.notifications.v1.CreateUserNotificationRequest.notification:type_name -> teleport.notifications.v1.Notification
	3,  // 1: teleport.notifications.v1.ListNotificationsRequest.filters:type_name -> teleport.notifications.v1.NotificationFilters
	9,  // 2: teleport.notifications.v1.NotificationFilters.labels:type_name -> teleport.notifications.v1.NotificationFilters.LabelsEntry
	10, // 3: teleport.notifications.v1.ListNotificationsResponse.notifications:type_name -> teleport.notifications.v1.Notification
	11, // 4: teleport.notifications.v1.ListNotificationsResponse.user_last_seen_notification_timestamp:type_name -> google.protobuf.Timestamp
	12, // 5: teleport.notifications.v1.CreateGlobalNotificationRequest.global_notification:type_name -> teleport.notifications.v1.GlobalNotification
	13, // 6: teleport.notifications.v1.UpsertUserNotificationStateRequest.user_notification_state:type_name -> teleport.notifications.v1.UserNotificationState
	14, // 7: teleport.notifications.v1.UpsertUserLastSeenNotificationRequest.user_last_seen_notification:type_name -> teleport.notifications.v1.UserLastSeenNotification
	0,  // 8: teleport.notifications.v1.NotificationService.CreateUserNotification:input_type -> teleport.notifications.v1.CreateUserNotificationRequest
	1,  // 9: teleport.notifications.v1.NotificationService.DeleteUserNotification:input_type -> teleport.notifications.v1.DeleteUserNotificationRequest
	5,  // 10: teleport.notifications.v1.NotificationService.CreateGlobalNotification:input_type -> teleport.notifications.v1.CreateGlobalNotificationRequest
	6,  // 11: teleport.notifications.v1.NotificationService.DeleteGlobalNotification:input_type -> teleport.notifications.v1.DeleteGlobalNotificationRequest
	2,  // 12: teleport.notifications.v1.NotificationService.ListNotifications:input_type -> teleport.notifications.v1.ListNotificationsRequest
	7,  // 13: teleport.notifications.v1.NotificationService.UpsertUserNotificationState:input_type -> teleport.notifications.v1.UpsertUserNotificationStateRequest
	8,  // 14: teleport.notifications.v1.NotificationService.UpsertUserLastSeenNotification:input_type -> teleport.notifications.v1.UpsertUserLastSeenNotificationRequest
	10, // 15: teleport.notifications.v1.NotificationService.CreateUserNotification:output_type -> teleport.notifications.v1.Notification
	15, // 16: teleport.notifications.v1.NotificationService.DeleteUserNotification:output_type -> google.protobuf.Empty
	12, // 17: teleport.notifications.v1.NotificationService.CreateGlobalNotification:output_type -> teleport.notifications.v1.GlobalNotification
	15, // 18: teleport.notifications.v1.NotificationService.DeleteGlobalNotification:output_type -> google.protobuf.Empty
	4,  // 19: teleport.notifications.v1.NotificationService.ListNotifications:output_type -> teleport.notifications.v1.ListNotificationsResponse
	13, // 20: teleport.notifications.v1.NotificationService.UpsertUserNotificationState:output_type -> teleport.notifications.v1.UserNotificationState
	14, // 21: teleport.notifications.v1.NotificationService.UpsertUserLastSeenNotification:output_type -> teleport.notifications.v1.UserLastSeenNotification
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_teleport_notifications_v1_notifications_service_proto_init() }
func file_teleport_notifications_v1_notifications_service_proto_init() {
	if File_teleport_notifications_v1_notifications_service_proto != nil {
		return
	}
	file_teleport_notifications_v1_notifications_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_notifications_v1_notifications_service_proto_rawDesc), len(file_teleport_notifications_v1_notifications_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_notifications_v1_notifications_service_proto_goTypes,
		DependencyIndexes: file_teleport_notifications_v1_notifications_service_proto_depIdxs,
		MessageInfos:      file_teleport_notifications_v1_notifications_service_proto_msgTypes,
	}.Build()
	File_teleport_notifications_v1_notifications_service_proto = out.File
	file_teleport_notifications_v1_notifications_service_proto_goTypes = nil
	file_teleport_notifications_v1_notifications_service_proto_depIdxs = nil
}
