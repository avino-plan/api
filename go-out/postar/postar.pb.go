// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: postar.proto

package postar

import (
	base "github.com/avino-plan/api/go-out/base"
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

// Email wraps all information of using smtp service.
type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receivers []string `protobuf:"bytes,1,rep,name=receivers,proto3" json:"receivers,omitempty"`               // Receivers.
	Subject   string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`                   // Subject.
	BodyType  string   `protobuf:"bytes,3,opt,name=body_type,json=bodyType,proto3" json:"body_type,omitempty"` // Body type.
	Body      string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`                         // Body.
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_postar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_postar_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetReceivers() []string {
	if x != nil {
		return x.Receivers
	}
	return nil
}

func (x *Email) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Email) GetBodyType() string {
	if x != nil {
		return x.BodyType
	}
	return ""
}

func (x *Email) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// SendEmailOptions is the options of sending emails.
type SendEmailOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async         bool  `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`                                      // Sending this email asynchronously.
	TimeoutMillis int32 `protobuf:"varint,2,opt,name=timeout_millis,json=timeoutMillis,proto3" json:"timeout_millis,omitempty"` // Sending timeout in ms.
}

func (x *SendEmailOptions) Reset() {
	*x = SendEmailOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailOptions) ProtoMessage() {}

func (x *SendEmailOptions) ProtoReflect() protoreflect.Message {
	mi := &file_postar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailOptions.ProtoReflect.Descriptor instead.
func (*SendEmailOptions) Descriptor() ([]byte, []int) {
	return file_postar_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailOptions) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendEmailOptions) GetTimeoutMillis() int32 {
	if x != nil {
		return x.TimeoutMillis
	}
	return 0
}

// SendEmailRequest is the request of SendEmail.
type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email   *Email            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`     // Sending email.
	Options *SendEmailOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"` // Sending options.
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_postar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_postar_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailRequest) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *SendEmailRequest) GetOptions() *SendEmailOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

// SendEmailResponse is the response of SendEmail.
type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    base.ServerCode `protobuf:"varint,1,opt,name=code,proto3,enum=github.com.avinoplan.api.base.ServerCode" json:"code,omitempty"` // 0 is ok.
	Msg     string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`                                                  // For messaging.
	TraceId string          `protobuf:"bytes,3,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`                           // For tracing.
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_postar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_postar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_postar_proto_rawDescGZIP(), []int{3}
}

func (x *SendEmailResponse) GetCode() base.ServerCode {
	if x != nil {
		return x.Code
	}
	return base.ServerCode(0)
}

func (x *SendEmailResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SendEmailResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

var File_postar_proto protoreflect.FileDescriptor

var file_postar_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e, 0x6f,
	0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x1a,
	0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x4f, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x9d,
	0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x4b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x61, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7f,
	0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x32,
	0x83, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x72, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x31,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x76, 0x69, 0x6e,
	0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x76, 0x69, 0x6e, 0x6f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x69, 0x6e, 0x6f, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x6f, 0x75, 0x74, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_postar_proto_rawDescOnce sync.Once
	file_postar_proto_rawDescData = file_postar_proto_rawDesc
)

func file_postar_proto_rawDescGZIP() []byte {
	file_postar_proto_rawDescOnce.Do(func() {
		file_postar_proto_rawDescData = protoimpl.X.CompressGZIP(file_postar_proto_rawDescData)
	})
	return file_postar_proto_rawDescData
}

var file_postar_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_postar_proto_goTypes = []interface{}{
	(*Email)(nil),             // 0: github.com.avinoplan.api.postar.Email
	(*SendEmailOptions)(nil),  // 1: github.com.avinoplan.api.postar.SendEmailOptions
	(*SendEmailRequest)(nil),  // 2: github.com.avinoplan.api.postar.SendEmailRequest
	(*SendEmailResponse)(nil), // 3: github.com.avinoplan.api.postar.SendEmailResponse
	(base.ServerCode)(0),      // 4: github.com.avinoplan.api.base.ServerCode
}
var file_postar_proto_depIdxs = []int32{
	0, // 0: github.com.avinoplan.api.postar.SendEmailRequest.email:type_name -> github.com.avinoplan.api.postar.Email
	1, // 1: github.com.avinoplan.api.postar.SendEmailRequest.options:type_name -> github.com.avinoplan.api.postar.SendEmailOptions
	4, // 2: github.com.avinoplan.api.postar.SendEmailResponse.code:type_name -> github.com.avinoplan.api.base.ServerCode
	2, // 3: github.com.avinoplan.api.postar.PostarService.SendEmail:input_type -> github.com.avinoplan.api.postar.SendEmailRequest
	3, // 4: github.com.avinoplan.api.postar.PostarService.SendEmail:output_type -> github.com.avinoplan.api.postar.SendEmailResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_postar_proto_init() }
func file_postar_proto_init() {
	if File_postar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_postar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_postar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailOptions); i {
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
		file_postar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_postar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailResponse); i {
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
			RawDescriptor: file_postar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_postar_proto_goTypes,
		DependencyIndexes: file_postar_proto_depIdxs,
		MessageInfos:      file_postar_proto_msgTypes,
	}.Build()
	File_postar_proto = out.File
	file_postar_proto_rawDesc = nil
	file_postar_proto_goTypes = nil
	file_postar_proto_depIdxs = nil
}
