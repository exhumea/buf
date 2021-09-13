// Copyright 2020-2021 Buf Technologies, Inc.
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: buf/alpha/registry/v1alpha1/repository_tag.proto

package registryv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RepositoryTagServiceClient is the client API for RepositoryTagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryTagServiceClient interface {
	// CreateRepositoryTag creates a new repository tag.
	CreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest, opts ...grpc.CallOption) (*CreateRepositoryTagResponse, error)
	// ListRepositoryTags lists the repository tags associated with a Repository.
	ListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest, opts ...grpc.CallOption) (*ListRepositoryTagsResponse, error)
}

type repositoryTagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryTagServiceClient(cc grpc.ClientConnInterface) RepositoryTagServiceClient {
	return &repositoryTagServiceClient{cc}
}

func (c *repositoryTagServiceClient) CreateRepositoryTag(ctx context.Context, in *CreateRepositoryTagRequest, opts ...grpc.CallOption) (*CreateRepositoryTagResponse, error) {
	out := new(CreateRepositoryTagResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.RepositoryTagService/CreateRepositoryTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryTagServiceClient) ListRepositoryTags(ctx context.Context, in *ListRepositoryTagsRequest, opts ...grpc.CallOption) (*ListRepositoryTagsResponse, error) {
	out := new(ListRepositoryTagsResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.RepositoryTagService/ListRepositoryTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryTagServiceServer is the server API for RepositoryTagService service.
// All implementations should embed UnimplementedRepositoryTagServiceServer
// for forward compatibility
type RepositoryTagServiceServer interface {
	// CreateRepositoryTag creates a new repository tag.
	CreateRepositoryTag(context.Context, *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error)
	// ListRepositoryTags lists the repository tags associated with a Repository.
	ListRepositoryTags(context.Context, *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error)
}

// UnimplementedRepositoryTagServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRepositoryTagServiceServer struct {
}

func (UnimplementedRepositoryTagServiceServer) CreateRepositoryTag(context.Context, *CreateRepositoryTagRequest) (*CreateRepositoryTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepositoryTag not implemented")
}
func (UnimplementedRepositoryTagServiceServer) ListRepositoryTags(context.Context, *ListRepositoryTagsRequest) (*ListRepositoryTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositoryTags not implemented")
}

// UnsafeRepositoryTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryTagServiceServer will
// result in compilation errors.
type UnsafeRepositoryTagServiceServer interface {
	mustEmbedUnimplementedRepositoryTagServiceServer()
}

func RegisterRepositoryTagServiceServer(s grpc.ServiceRegistrar, srv RepositoryTagServiceServer) {
	s.RegisterService(&RepositoryTagService_ServiceDesc, srv)
}

func _RepositoryTagService_CreateRepositoryTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryTagServiceServer).CreateRepositoryTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.RepositoryTagService/CreateRepositoryTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryTagServiceServer).CreateRepositoryTag(ctx, req.(*CreateRepositoryTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryTagService_ListRepositoryTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoryTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryTagServiceServer).ListRepositoryTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.RepositoryTagService/ListRepositoryTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryTagServiceServer).ListRepositoryTags(ctx, req.(*ListRepositoryTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepositoryTagService_ServiceDesc is the grpc.ServiceDesc for RepositoryTagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepositoryTagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.RepositoryTagService",
	HandlerType: (*RepositoryTagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepositoryTag",
			Handler:    _RepositoryTagService_CreateRepositoryTag_Handler,
		},
		{
			MethodName: "ListRepositoryTags",
			Handler:    _RepositoryTagService_ListRepositoryTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/repository_tag.proto",
}