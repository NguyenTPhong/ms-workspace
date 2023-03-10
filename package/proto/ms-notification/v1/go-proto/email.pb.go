// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/ms-notification/v1/email.proto

package go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendActiveEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Code   string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url    string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	UserId int64  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SendActiveEmailRequest) Reset() {
	*x = SendActiveEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ms_notification_v1_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendActiveEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendActiveEmailRequest) ProtoMessage() {}

func (x *SendActiveEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ms_notification_v1_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendActiveEmailRequest.ProtoReflect.Descriptor instead.
func (*SendActiveEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_ms_notification_v1_email_proto_rawDescGZIP(), []int{0}
}

func (x *SendActiveEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendActiveEmailRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SendActiveEmailRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SendActiveEmailRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SendActiveEmailRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SendActiveEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *SendActiveEmailResponse) Reset() {
	*x = SendActiveEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ms_notification_v1_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendActiveEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendActiveEmailResponse) ProtoMessage() {}

func (x *SendActiveEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ms_notification_v1_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendActiveEmailResponse.ProtoReflect.Descriptor instead.
func (*SendActiveEmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_ms_notification_v1_email_proto_rawDescGZIP(), []int{1}
}

func (x *SendActiveEmailResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

var File_proto_ms_notification_v1_email_proto protoreflect.FileDescriptor

var file_proto_ms_notification_v1_email_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6d, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x16,
	0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x38, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x32, 0x7c, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x53, 0x65, 0x6e,
	0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x2e, 0x6d,
	0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x73, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ms_notification_v1_email_proto_rawDescOnce sync.Once
	file_proto_ms_notification_v1_email_proto_rawDescData = file_proto_ms_notification_v1_email_proto_rawDesc
)

func file_proto_ms_notification_v1_email_proto_rawDescGZIP() []byte {
	file_proto_ms_notification_v1_email_proto_rawDescOnce.Do(func() {
		file_proto_ms_notification_v1_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ms_notification_v1_email_proto_rawDescData)
	})
	return file_proto_ms_notification_v1_email_proto_rawDescData
}

var file_proto_ms_notification_v1_email_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ms_notification_v1_email_proto_goTypes = []interface{}{
	(*SendActiveEmailRequest)(nil),  // 0: ms_notification.v1.SendActiveEmailRequest
	(*SendActiveEmailResponse)(nil), // 1: ms_notification.v1.SendActiveEmailResponse
}
var file_proto_ms_notification_v1_email_proto_depIdxs = []int32{
	0, // 0: ms_notification.v1.EmailService.SendActiveEmail:input_type -> ms_notification.v1.SendActiveEmailRequest
	1, // 1: ms_notification.v1.EmailService.SendActiveEmail:output_type -> ms_notification.v1.SendActiveEmailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ms_notification_v1_email_proto_init() }
func file_proto_ms_notification_v1_email_proto_init() {
	if File_proto_ms_notification_v1_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ms_notification_v1_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendActiveEmailRequest); i {
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
		file_proto_ms_notification_v1_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendActiveEmailResponse); i {
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
			RawDescriptor: file_proto_ms_notification_v1_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ms_notification_v1_email_proto_goTypes,
		DependencyIndexes: file_proto_ms_notification_v1_email_proto_depIdxs,
		MessageInfos:      file_proto_ms_notification_v1_email_proto_msgTypes,
	}.Build()
	File_proto_ms_notification_v1_email_proto = out.File
	file_proto_ms_notification_v1_email_proto_rawDesc = nil
	file_proto_ms_notification_v1_email_proto_goTypes = nil
	file_proto_ms_notification_v1_email_proto_depIdxs = nil
}
