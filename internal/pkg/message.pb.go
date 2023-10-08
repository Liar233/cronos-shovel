// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: pkg/message.proto

package pkg

import (
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

type DelayObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DateTime int64  `protobuf:"varint,2,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
}

func (x *DelayObject) Reset() {
	*x = DelayObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayObject) ProtoMessage() {}

func (x *DelayObject) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelayObject.ProtoReflect.Descriptor instead.
func (*DelayObject) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{0}
}

func (x *DelayObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DelayObject) GetDateTime() int64 {
	if x != nil {
		return x.DateTime
	}
	return 0
}

type MessageObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mask     string         `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
	Title    string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Channels []string       `protobuf:"bytes,4,rep,name=channels,proto3" json:"channels,omitempty"`
	Payload  []byte         `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Delays   []*DelayObject `protobuf:"bytes,6,rep,name=delays,proto3" json:"delays,omitempty"`
}

func (x *MessageObject) Reset() {
	*x = MessageObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageObject) ProtoMessage() {}

func (x *MessageObject) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageObject.ProtoReflect.Descriptor instead.
func (*MessageObject) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageObject) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *MessageObject) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MessageObject) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *MessageObject) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MessageObject) GetDelays() []*DelayObject {
	if x != nil {
		return x.Delays
	}
	return nil
}

type GetMessageListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMessageListRequest) Reset() {
	*x = GetMessageListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageListRequest) ProtoMessage() {}

func (x *GetMessageListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageListRequest.ProtoReflect.Descriptor instead.
func (*GetMessageListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{2}
}

type GetMessageListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    string           `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Messages []*MessageObject `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetMessageListResponse) Reset() {
	*x = GetMessageListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageListResponse) ProtoMessage() {}

func (x *GetMessageListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageListResponse.ProtoReflect.Descriptor instead.
func (*GetMessageListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageListResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetMessageListResponse) GetMessages() []*MessageObject {
	if x != nil {
		return x.Messages
	}
	return nil
}

type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mask     string   `protobuf:"bytes,1,opt,name=mask,proto3" json:"mask,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Channels []string `protobuf:"bytes,3,rep,name=channels,proto3" json:"channels,omitempty"`
	Payload  []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{4}
}

func (x *CreateMessageRequest) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *CreateMessageRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMessageRequest) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *CreateMessageRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Id       string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Mask     string   `protobuf:"bytes,3,opt,name=mask,proto3" json:"mask,omitempty"`
	Title    string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Channels []string `protobuf:"bytes,5,rep,name=channels,proto3" json:"channels,omitempty"`
	Payload  []byte   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateMessageResponse) Reset() {
	*x = CreateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageResponse) ProtoMessage() {}

func (x *CreateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{5}
}

func (x *CreateMessageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateMessageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateMessageResponse) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *CreateMessageResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateMessageResponse) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *CreateMessageResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type UpdateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mask     string   `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
	Title    string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Channels []string `protobuf:"bytes,4,rep,name=channels,proto3" json:"channels,omitempty"`
	Payload  []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UpdateMessageRequest) Reset() {
	*x = UpdateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRequest) ProtoMessage() {}

func (x *UpdateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMessageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMessageRequest) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *UpdateMessageRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMessageRequest) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *UpdateMessageRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type UpdateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Id       string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Mask     string   `protobuf:"bytes,3,opt,name=mask,proto3" json:"mask,omitempty"`
	Title    string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Channels []string `protobuf:"bytes,5,rep,name=channels,proto3" json:"channels,omitempty"`
	Payload  []byte   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UpdateMessageResponse) Reset() {
	*x = UpdateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageResponse) ProtoMessage() {}

func (x *UpdateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageResponse.ProtoReflect.Descriptor instead.
func (*UpdateMessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateMessageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *UpdateMessageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMessageResponse) GetMask() string {
	if x != nil {
		return x.Mask
	}
	return ""
}

func (x *UpdateMessageResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateMessageResponse) GetChannels() []string {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *UpdateMessageResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type DeleteMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMessageRequest) Reset() {
	*x = DeleteMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRequest) ProtoMessage() {}

func (x *DeleteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMessageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteMessageResponse) Reset() {
	*x = DeleteMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageResponse) ProtoMessage() {}

func (x *DeleteMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageResponse.ProtoReflect.Descriptor instead.
func (*DeleteMessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_message_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMessageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_message_proto protoreflect.FileDescriptor

var file_pkg_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x06, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x76, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x9d, 0x01, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x86, 0x01, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xde, 0x02, 0x0a,
	0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x69, 0x61, 0x72,
	0x32, 0x33, 0x33, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x6f, 0x73, 0x2d, 0x73, 0x68, 0x6f, 0x76, 0x65,
	0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_message_proto_rawDescOnce sync.Once
	file_pkg_message_proto_rawDescData = file_pkg_message_proto_rawDesc
)

func file_pkg_message_proto_rawDescGZIP() []byte {
	file_pkg_message_proto_rawDescOnce.Do(func() {
		file_pkg_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_message_proto_rawDescData)
	})
	return file_pkg_message_proto_rawDescData
}

var file_pkg_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_message_proto_goTypes = []interface{}{
	(*DelayObject)(nil),            // 0: message.DelayObject
	(*MessageObject)(nil),          // 1: message.MessageObject
	(*GetMessageListRequest)(nil),  // 2: message.GetMessageListRequest
	(*GetMessageListResponse)(nil), // 3: message.GetMessageListResponse
	(*CreateMessageRequest)(nil),   // 4: message.CreateMessageRequest
	(*CreateMessageResponse)(nil),  // 5: message.CreateMessageResponse
	(*UpdateMessageRequest)(nil),   // 6: message.UpdateMessageRequest
	(*UpdateMessageResponse)(nil),  // 7: message.UpdateMessageResponse
	(*DeleteMessageRequest)(nil),   // 8: message.DeleteMessageRequest
	(*DeleteMessageResponse)(nil),  // 9: message.DeleteMessageResponse
}
var file_pkg_message_proto_depIdxs = []int32{
	0, // 0: message.MessageObject.delays:type_name -> message.DelayObject
	1, // 1: message.GetMessageListResponse.messages:type_name -> message.MessageObject
	2, // 2: message.MessageController.GetMessageList:input_type -> message.GetMessageListRequest
	4, // 3: message.MessageController.CreateMessage:input_type -> message.CreateMessageRequest
	6, // 4: message.MessageController.UpdateMessage:input_type -> message.UpdateMessageRequest
	8, // 5: message.MessageController.DeleteMessage:input_type -> message.DeleteMessageRequest
	3, // 6: message.MessageController.GetMessageList:output_type -> message.GetMessageListResponse
	5, // 7: message.MessageController.CreateMessage:output_type -> message.CreateMessageResponse
	7, // 8: message.MessageController.UpdateMessage:output_type -> message.UpdateMessageResponse
	9, // 9: message.MessageController.DeleteMessage:output_type -> message.DeleteMessageResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_message_proto_init() }
func file_pkg_message_proto_init() {
	if File_pkg_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayObject); i {
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
		file_pkg_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageObject); i {
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
		file_pkg_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageListRequest); i {
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
		file_pkg_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageListResponse); i {
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
		file_pkg_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRequest); i {
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
		file_pkg_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageResponse); i {
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
		file_pkg_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRequest); i {
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
		file_pkg_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageResponse); i {
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
		file_pkg_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageRequest); i {
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
		file_pkg_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageResponse); i {
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
			RawDescriptor: file_pkg_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_message_proto_goTypes,
		DependencyIndexes: file_pkg_message_proto_depIdxs,
		MessageInfos:      file_pkg_message_proto_msgTypes,
	}.Build()
	File_pkg_message_proto = out.File
	file_pkg_message_proto_rawDesc = nil
	file_pkg_message_proto_goTypes = nil
	file_pkg_message_proto_depIdxs = nil
}
