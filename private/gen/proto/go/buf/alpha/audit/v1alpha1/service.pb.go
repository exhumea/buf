// Copyright 2020-2025 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: buf/alpha/audit/v1alpha1/service.proto

package auditv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListAuditedEventsRequest struct {
	state                protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_PageSize  uint32                 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	xxx_hidden_PageToken string                 `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	xxx_hidden_Start     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	xxx_hidden_End       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ListAuditedEventsRequest) Reset() {
	*x = ListAuditedEventsRequest{}
	mi := &file_buf_alpha_audit_v1alpha1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuditedEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditedEventsRequest) ProtoMessage() {}

func (x *ListAuditedEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListAuditedEventsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.xxx_hidden_PageSize
	}
	return 0
}

func (x *ListAuditedEventsRequest) GetPageToken() string {
	if x != nil {
		return x.xxx_hidden_PageToken
	}
	return ""
}

func (x *ListAuditedEventsRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_Start
	}
	return nil
}

func (x *ListAuditedEventsRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_End
	}
	return nil
}

func (x *ListAuditedEventsRequest) SetPageSize(v uint32) {
	x.xxx_hidden_PageSize = v
}

func (x *ListAuditedEventsRequest) SetPageToken(v string) {
	x.xxx_hidden_PageToken = v
}

func (x *ListAuditedEventsRequest) SetStart(v *timestamppb.Timestamp) {
	x.xxx_hidden_Start = v
}

func (x *ListAuditedEventsRequest) SetEnd(v *timestamppb.Timestamp) {
	x.xxx_hidden_End = v
}

func (x *ListAuditedEventsRequest) HasStart() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Start != nil
}

func (x *ListAuditedEventsRequest) HasEnd() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_End != nil
}

func (x *ListAuditedEventsRequest) ClearStart() {
	x.xxx_hidden_Start = nil
}

func (x *ListAuditedEventsRequest) ClearEnd() {
	x.xxx_hidden_End = nil
}

type ListAuditedEventsRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The page size for listing audited events, values between 1-250.
	PageSize uint32
	// The page token for paginating. The first page is returned if this is empty.
	PageToken string
	// The start timestamp to filter events from.
	Start *timestamppb.Timestamp
	// The end timestamp to filter events to.
	End *timestamppb.Timestamp
}

func (b0 ListAuditedEventsRequest_builder) Build() *ListAuditedEventsRequest {
	m0 := &ListAuditedEventsRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_PageSize = b.PageSize
	x.xxx_hidden_PageToken = b.PageToken
	x.xxx_hidden_Start = b.Start
	x.xxx_hidden_End = b.End
	return m0
}

type ListAuditedEventsResponse struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Events        *[]*Event              `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	xxx_hidden_NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ListAuditedEventsResponse) Reset() {
	*x = ListAuditedEventsResponse{}
	mi := &file_buf_alpha_audit_v1alpha1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuditedEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditedEventsResponse) ProtoMessage() {}

func (x *ListAuditedEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_audit_v1alpha1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListAuditedEventsResponse) GetEvents() []*Event {
	if x != nil {
		if x.xxx_hidden_Events != nil {
			return *x.xxx_hidden_Events
		}
	}
	return nil
}

func (x *ListAuditedEventsResponse) GetNextPageToken() string {
	if x != nil {
		return x.xxx_hidden_NextPageToken
	}
	return ""
}

func (x *ListAuditedEventsResponse) SetEvents(v []*Event) {
	x.xxx_hidden_Events = &v
}

func (x *ListAuditedEventsResponse) SetNextPageToken(v string) {
	x.xxx_hidden_NextPageToken = v
}

type ListAuditedEventsResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The list of audited events in the current page.
	Events []*Event
	// The token for the next page of audited events. There are no more pages if
	// this is empty.
	NextPageToken string
}

func (b0 ListAuditedEventsResponse_builder) Build() *ListAuditedEventsResponse {
	m0 := &ListAuditedEventsResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Events = &b.Events
	x.xxx_hidden_NextPageToken = b.NextPageToken
	return m0
}

var File_buf_alpha_audit_v1alpha1_service_proto protoreflect.FileDescriptor

var file_buf_alpha_audit_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x24, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x7c, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x32, 0x92, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x81, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0x84, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41,
	0x41, 0xaa, 0x02, 0x18, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18, 0x42,
	0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1b, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_buf_alpha_audit_v1alpha1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_alpha_audit_v1alpha1_service_proto_goTypes = []any{
	(*ListAuditedEventsRequest)(nil),  // 0: buf.alpha.audit.v1alpha1.ListAuditedEventsRequest
	(*ListAuditedEventsResponse)(nil), // 1: buf.alpha.audit.v1alpha1.ListAuditedEventsResponse
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
	(*Event)(nil),                     // 3: buf.alpha.audit.v1alpha1.Event
}
var file_buf_alpha_audit_v1alpha1_service_proto_depIdxs = []int32{
	2, // 0: buf.alpha.audit.v1alpha1.ListAuditedEventsRequest.start:type_name -> google.protobuf.Timestamp
	2, // 1: buf.alpha.audit.v1alpha1.ListAuditedEventsRequest.end:type_name -> google.protobuf.Timestamp
	3, // 2: buf.alpha.audit.v1alpha1.ListAuditedEventsResponse.events:type_name -> buf.alpha.audit.v1alpha1.Event
	0, // 3: buf.alpha.audit.v1alpha1.AuditService.ListAuditedEvents:input_type -> buf.alpha.audit.v1alpha1.ListAuditedEventsRequest
	1, // 4: buf.alpha.audit.v1alpha1.AuditService.ListAuditedEvents:output_type -> buf.alpha.audit.v1alpha1.ListAuditedEventsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buf_alpha_audit_v1alpha1_service_proto_init() }
func file_buf_alpha_audit_v1alpha1_service_proto_init() {
	if File_buf_alpha_audit_v1alpha1_service_proto != nil {
		return
	}
	file_buf_alpha_audit_v1alpha1_event_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_audit_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_alpha_audit_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_buf_alpha_audit_v1alpha1_service_proto_depIdxs,
		MessageInfos:      file_buf_alpha_audit_v1alpha1_service_proto_msgTypes,
	}.Build()
	File_buf_alpha_audit_v1alpha1_service_proto = out.File
	file_buf_alpha_audit_v1alpha1_service_proto_rawDesc = nil
	file_buf_alpha_audit_v1alpha1_service_proto_goTypes = nil
	file_buf_alpha_audit_v1alpha1_service_proto_depIdxs = nil
}
