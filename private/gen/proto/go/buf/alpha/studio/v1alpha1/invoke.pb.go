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
// source: buf/alpha/studio/v1alpha1/invoke.proto

// The buf.alpha.studio.v1alpha1 package contains types used by the buf studio
// agent. Because browsers are not capable of sending gRPC requests directly,
// users can run a studio agent to that receives enveloped requests from the
// browser and forwards them as gRPC requests.
//
// Ideally the agent would be simple protocol translating HTTP proxy without
// requiring any custom envelope. Unfortunately, js in the browser cannot set
// per request proxy configuration and we cannot specify that we want to open a
// connection to the request agent while specifying a different server in the
// request's Host header. The studio agent and UI could communicate this through
// a custom header instead, but reading custom headers requires a CORS-preflight
// request.
//
// To facilitate easier deployment it in environments with complicated edge
// configuration, it is a goal for the agent and UI to communicate without the
// need for a CORS-preflight requests. This limits our ability to use custom
// headers and restricts allowed values for the Content-Type header. Due to this
// we cannot simply use gRPC-Web with an additional header, but instead rely on
// enveloping the request and responses in a base64 encoded binary proto message
// sent over a POST endpoint with text/plain as Content-Type.
//
// We may explore other transports such as WebSockets or WebTransport, at which
// point we should define proper proto services and methods here as well.

package studiov1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Headers encode HTTP headers.
type Headers struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key   string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	xxx_hidden_Value []string               `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Headers) Reset() {
	*x = Headers{}
	mi := &file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Headers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Headers) ProtoMessage() {}

func (x *Headers) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Headers) GetKey() string {
	if x != nil {
		return x.xxx_hidden_Key
	}
	return ""
}

func (x *Headers) GetValue() []string {
	if x != nil {
		return x.xxx_hidden_Value
	}
	return nil
}

func (x *Headers) SetKey(v string) {
	x.xxx_hidden_Key = v
}

func (x *Headers) SetValue(v []string) {
	x.xxx_hidden_Value = v
}

type Headers_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Key   string
	Value []string
}

func (b0 Headers_builder) Build() *Headers {
	m0 := &Headers{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Key = b.Key
	x.xxx_hidden_Value = b.Value
	return m0
}

// InvokeRequest encodes an enveloped RPC request. See the package documentation
// for more information.
type InvokeRequest struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Target  string                 `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	xxx_hidden_Headers *[]*Headers            `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	xxx_hidden_Body    []byte                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *InvokeRequest) Reset() {
	*x = InvokeRequest{}
	mi := &file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeRequest) ProtoMessage() {}

func (x *InvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *InvokeRequest) GetTarget() string {
	if x != nil {
		return x.xxx_hidden_Target
	}
	return ""
}

func (x *InvokeRequest) GetHeaders() []*Headers {
	if x != nil {
		if x.xxx_hidden_Headers != nil {
			return *x.xxx_hidden_Headers
		}
	}
	return nil
}

func (x *InvokeRequest) GetBody() []byte {
	if x != nil {
		return x.xxx_hidden_Body
	}
	return nil
}

func (x *InvokeRequest) SetTarget(v string) {
	x.xxx_hidden_Target = v
}

func (x *InvokeRequest) SetHeaders(v []*Headers) {
	x.xxx_hidden_Headers = &v
}

func (x *InvokeRequest) SetBody(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Body = v
}

type InvokeRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Target server the agent should forward this request to, e.g.
	// "https://api.acme.corp/pkg.Service/Method". Using the "http" scheme will
	// cause the request to be forwarded as h2c, whereas "https" forwards the
	// request with regular h2.
	Target string
	// Headers to send with the request. If body is set, a Content-Type header
	// must be specified.
	Headers []*Headers
	// The message to be sent in the request (without any protocol specific framing).
	Body []byte
}

func (b0 InvokeRequest_builder) Build() *InvokeRequest {
	m0 := &InvokeRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Target = b.Target
	x.xxx_hidden_Headers = &b.Headers
	x.xxx_hidden_Body = b.Body
	return m0
}

// InvokeResponse encodes an enveloped RPC response. See the package documentation
// for more information.
type InvokeResponse struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Headers  *[]*Headers            `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	xxx_hidden_Body     []byte                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	xxx_hidden_Trailers *[]*Headers            `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *InvokeResponse) Reset() {
	*x = InvokeResponse{}
	mi := &file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeResponse) ProtoMessage() {}

func (x *InvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *InvokeResponse) GetHeaders() []*Headers {
	if x != nil {
		if x.xxx_hidden_Headers != nil {
			return *x.xxx_hidden_Headers
		}
	}
	return nil
}

func (x *InvokeResponse) GetBody() []byte {
	if x != nil {
		return x.xxx_hidden_Body
	}
	return nil
}

func (x *InvokeResponse) GetTrailers() []*Headers {
	if x != nil {
		if x.xxx_hidden_Trailers != nil {
			return *x.xxx_hidden_Trailers
		}
	}
	return nil
}

func (x *InvokeResponse) SetHeaders(v []*Headers) {
	x.xxx_hidden_Headers = &v
}

func (x *InvokeResponse) SetBody(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Body = v
}

func (x *InvokeResponse) SetTrailers(v []*Headers) {
	x.xxx_hidden_Trailers = &v
}

type InvokeResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Headers received in the response.
	Headers []*Headers
	// The encoded message received in the response (without protocol specific framing).
	Body []byte
	// Trailers received in the response.
	Trailers []*Headers
}

func (b0 InvokeResponse_builder) Build() *InvokeResponse {
	m0 := &InvokeResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Headers = &b.Headers
	x.xxx_hidden_Body = b.Body
	x.xxx_hidden_Trailers = &b.Trailers
	return m0
}

var File_buf_alpha_studio_v1alpha1_invoke_proto protoreflect.FileDescriptor

var file_buf_alpha_studio_v1alpha1_invoke_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0x31, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x79, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x3c, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x08, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x42, 0x8a, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x42, 0x41, 0x53, 0xaa, 0x02, 0x19, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x19, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x53, 0x74, 0x75,
	0x64, 0x69, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x42,
	0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68,
	0x61, 0x3a, 0x3a, 0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_alpha_studio_v1alpha1_invoke_proto_goTypes = []any{
	(*Headers)(nil),        // 0: buf.alpha.studio.v1alpha1.Headers
	(*InvokeRequest)(nil),  // 1: buf.alpha.studio.v1alpha1.InvokeRequest
	(*InvokeResponse)(nil), // 2: buf.alpha.studio.v1alpha1.InvokeResponse
}
var file_buf_alpha_studio_v1alpha1_invoke_proto_depIdxs = []int32{
	0, // 0: buf.alpha.studio.v1alpha1.InvokeRequest.headers:type_name -> buf.alpha.studio.v1alpha1.Headers
	0, // 1: buf.alpha.studio.v1alpha1.InvokeResponse.headers:type_name -> buf.alpha.studio.v1alpha1.Headers
	0, // 2: buf.alpha.studio.v1alpha1.InvokeResponse.trailers:type_name -> buf.alpha.studio.v1alpha1.Headers
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buf_alpha_studio_v1alpha1_invoke_proto_init() }
func file_buf_alpha_studio_v1alpha1_invoke_proto_init() {
	if File_buf_alpha_studio_v1alpha1_invoke_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_studio_v1alpha1_invoke_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_studio_v1alpha1_invoke_proto_goTypes,
		DependencyIndexes: file_buf_alpha_studio_v1alpha1_invoke_proto_depIdxs,
		MessageInfos:      file_buf_alpha_studio_v1alpha1_invoke_proto_msgTypes,
	}.Build()
	File_buf_alpha_studio_v1alpha1_invoke_proto = out.File
	file_buf_alpha_studio_v1alpha1_invoke_proto_rawDesc = nil
	file_buf_alpha_studio_v1alpha1_invoke_proto_goTypes = nil
	file_buf_alpha_studio_v1alpha1_invoke_proto_depIdxs = nil
}
