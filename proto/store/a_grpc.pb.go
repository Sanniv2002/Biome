// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/store/a.proto

package store

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Store_AddParentContainer_FullMethodName   = "/store.Store/AddParentContainer"
	Store_AddChildContainer_FullMethodName    = "/store.Store/AddChildContainer"
	Store_RemoveContainer_FullMethodName      = "/store.Store/RemoveContainer"
	Store_GetAllContainers_FullMethodName     = "/store.Store/GetAllContainers"
	Store_UpdateContainerStats_FullMethodName = "/store.Store/UpdateContainerStats"
	Store_GetContainerStats_FullMethodName    = "/store.Store/GetContainerStats"
	Store_GetContainerLimits_FullMethodName   = "/store.Store/GetContainerLimits"
	Store_UpdatePortMappings_FullMethodName   = "/store.Store/UpdatePortMappings"
	Store_GetAllConfigs_FullMethodName        = "/store.Store/GetAllConfigs"
)

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreClient interface {
	AddParentContainer(ctx context.Context, in *AddParentContainerRequest, opts ...grpc.CallOption) (*AddParentContainerResponse, error)
	AddChildContainer(ctx context.Context, in *AddChildContainerRequest, opts ...grpc.CallOption) (*AddChildContainerResponse, error)
	RemoveContainer(ctx context.Context, in *RemoveContainerRequest, opts ...grpc.CallOption) (*RemoveContainerResponse, error)
	GetAllContainers(ctx context.Context, in *GetAllContainersRequest, opts ...grpc.CallOption) (*GetAllContainersResponse, error)
	UpdateContainerStats(ctx context.Context, in *UpdateContainerStatsRequest, opts ...grpc.CallOption) (*UpdateContainerStatsResponse, error)
	GetContainerStats(ctx context.Context, in *GetContainerStatsRequest, opts ...grpc.CallOption) (*GetContainerStatsResponse, error)
	GetContainerLimits(ctx context.Context, in *GetContainerLimitsRequest, opts ...grpc.CallOption) (*GetContainerLimitsResponse, error)
	UpdatePortMappings(ctx context.Context, in *UpdatePortMappingsRequest, opts ...grpc.CallOption) (*UpdatePortMappingsResponse, error)
	GetAllConfigs(ctx context.Context, in *GetAllConfigsRequest, opts ...grpc.CallOption) (*GetAllConfigsResponse, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) AddParentContainer(ctx context.Context, in *AddParentContainerRequest, opts ...grpc.CallOption) (*AddParentContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddParentContainerResponse)
	err := c.cc.Invoke(ctx, Store_AddParentContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) AddChildContainer(ctx context.Context, in *AddChildContainerRequest, opts ...grpc.CallOption) (*AddChildContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddChildContainerResponse)
	err := c.cc.Invoke(ctx, Store_AddChildContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) RemoveContainer(ctx context.Context, in *RemoveContainerRequest, opts ...grpc.CallOption) (*RemoveContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveContainerResponse)
	err := c.cc.Invoke(ctx, Store_RemoveContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetAllContainers(ctx context.Context, in *GetAllContainersRequest, opts ...grpc.CallOption) (*GetAllContainersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllContainersResponse)
	err := c.cc.Invoke(ctx, Store_GetAllContainers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdateContainerStats(ctx context.Context, in *UpdateContainerStatsRequest, opts ...grpc.CallOption) (*UpdateContainerStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateContainerStatsResponse)
	err := c.cc.Invoke(ctx, Store_UpdateContainerStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetContainerStats(ctx context.Context, in *GetContainerStatsRequest, opts ...grpc.CallOption) (*GetContainerStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContainerStatsResponse)
	err := c.cc.Invoke(ctx, Store_GetContainerStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetContainerLimits(ctx context.Context, in *GetContainerLimitsRequest, opts ...grpc.CallOption) (*GetContainerLimitsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContainerLimitsResponse)
	err := c.cc.Invoke(ctx, Store_GetContainerLimits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) UpdatePortMappings(ctx context.Context, in *UpdatePortMappingsRequest, opts ...grpc.CallOption) (*UpdatePortMappingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePortMappingsResponse)
	err := c.cc.Invoke(ctx, Store_UpdatePortMappings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) GetAllConfigs(ctx context.Context, in *GetAllConfigsRequest, opts ...grpc.CallOption) (*GetAllConfigsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllConfigsResponse)
	err := c.cc.Invoke(ctx, Store_GetAllConfigs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
// All implementations must embed UnimplementedStoreServer
// for forward compatibility
type StoreServer interface {
	AddParentContainer(context.Context, *AddParentContainerRequest) (*AddParentContainerResponse, error)
	AddChildContainer(context.Context, *AddChildContainerRequest) (*AddChildContainerResponse, error)
	RemoveContainer(context.Context, *RemoveContainerRequest) (*RemoveContainerResponse, error)
	GetAllContainers(context.Context, *GetAllContainersRequest) (*GetAllContainersResponse, error)
	UpdateContainerStats(context.Context, *UpdateContainerStatsRequest) (*UpdateContainerStatsResponse, error)
	GetContainerStats(context.Context, *GetContainerStatsRequest) (*GetContainerStatsResponse, error)
	GetContainerLimits(context.Context, *GetContainerLimitsRequest) (*GetContainerLimitsResponse, error)
	UpdatePortMappings(context.Context, *UpdatePortMappingsRequest) (*UpdatePortMappingsResponse, error)
	GetAllConfigs(context.Context, *GetAllConfigsRequest) (*GetAllConfigsResponse, error)
	mustEmbedUnimplementedStoreServer()
}

// UnimplementedStoreServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (UnimplementedStoreServer) AddParentContainer(context.Context, *AddParentContainerRequest) (*AddParentContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParentContainer not implemented")
}
func (UnimplementedStoreServer) AddChildContainer(context.Context, *AddChildContainerRequest) (*AddChildContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChildContainer not implemented")
}
func (UnimplementedStoreServer) RemoveContainer(context.Context, *RemoveContainerRequest) (*RemoveContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveContainer not implemented")
}
func (UnimplementedStoreServer) GetAllContainers(context.Context, *GetAllContainersRequest) (*GetAllContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllContainers not implemented")
}
func (UnimplementedStoreServer) UpdateContainerStats(context.Context, *UpdateContainerStatsRequest) (*UpdateContainerStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContainerStats not implemented")
}
func (UnimplementedStoreServer) GetContainerStats(context.Context, *GetContainerStatsRequest) (*GetContainerStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerStats not implemented")
}
func (UnimplementedStoreServer) GetContainerLimits(context.Context, *GetContainerLimitsRequest) (*GetContainerLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerLimits not implemented")
}
func (UnimplementedStoreServer) UpdatePortMappings(context.Context, *UpdatePortMappingsRequest) (*UpdatePortMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePortMappings not implemented")
}
func (UnimplementedStoreServer) GetAllConfigs(context.Context, *GetAllConfigsRequest) (*GetAllConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConfigs not implemented")
}
func (UnimplementedStoreServer) mustEmbedUnimplementedStoreServer() {}

// UnsafeStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServer will
// result in compilation errors.
type UnsafeStoreServer interface {
	mustEmbedUnimplementedStoreServer()
}

func RegisterStoreServer(s grpc.ServiceRegistrar, srv StoreServer) {
	s.RegisterService(&Store_ServiceDesc, srv)
}

func _Store_AddParentContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParentContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).AddParentContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_AddParentContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).AddParentContainer(ctx, req.(*AddParentContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_AddChildContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChildContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).AddChildContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_AddChildContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).AddChildContainer(ctx, req.(*AddChildContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_RemoveContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).RemoveContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_RemoveContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).RemoveContainer(ctx, req.(*RemoveContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetAllContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetAllContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_GetAllContainers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetAllContainers(ctx, req.(*GetAllContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdateContainerStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContainerStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdateContainerStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_UpdateContainerStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdateContainerStats(ctx, req.(*UpdateContainerStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetContainerStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetContainerStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_GetContainerStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetContainerStats(ctx, req.(*GetContainerStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetContainerLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetContainerLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_GetContainerLimits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetContainerLimits(ctx, req.(*GetContainerLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_UpdatePortMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePortMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).UpdatePortMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_UpdatePortMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).UpdatePortMappings(ctx, req.(*UpdatePortMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_GetAllConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).GetAllConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Store_GetAllConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).GetAllConfigs(ctx, req.(*GetAllConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Store_ServiceDesc is the grpc.ServiceDesc for Store service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Store_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddParentContainer",
			Handler:    _Store_AddParentContainer_Handler,
		},
		{
			MethodName: "AddChildContainer",
			Handler:    _Store_AddChildContainer_Handler,
		},
		{
			MethodName: "RemoveContainer",
			Handler:    _Store_RemoveContainer_Handler,
		},
		{
			MethodName: "GetAllContainers",
			Handler:    _Store_GetAllContainers_Handler,
		},
		{
			MethodName: "UpdateContainerStats",
			Handler:    _Store_UpdateContainerStats_Handler,
		},
		{
			MethodName: "GetContainerStats",
			Handler:    _Store_GetContainerStats_Handler,
		},
		{
			MethodName: "GetContainerLimits",
			Handler:    _Store_GetContainerLimits_Handler,
		},
		{
			MethodName: "UpdatePortMappings",
			Handler:    _Store_UpdatePortMappings_Handler,
		},
		{
			MethodName: "GetAllConfigs",
			Handler:    _Store_GetAllConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/store/a.proto",
}