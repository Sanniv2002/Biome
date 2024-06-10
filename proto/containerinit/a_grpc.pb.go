// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: proto/containerinit/a.proto

package containerinit

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
	ContainerInit_StartContainer_FullMethodName = "/containerinit.ContainerInit/StartContainer"
	ContainerInit_ScaleContainer_FullMethodName = "/containerinit.ContainerInit/ScaleContainer"
	ContainerInit_KillContainer_FullMethodName  = "/containerinit.ContainerInit/KillContainer"
)

// ContainerInitClient is the client API for ContainerInit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerInitClient interface {
	StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*StartContainerResponse, error)
	ScaleContainer(ctx context.Context, in *ScaleContainerRequest, opts ...grpc.CallOption) (*ScaleContainerResponse, error)
	KillContainer(ctx context.Context, in *KillContainerRequest, opts ...grpc.CallOption) (*KillContainerResponse, error)
}

type containerInitClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerInitClient(cc grpc.ClientConnInterface) ContainerInitClient {
	return &containerInitClient{cc}
}

func (c *containerInitClient) StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*StartContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartContainerResponse)
	err := c.cc.Invoke(ctx, ContainerInit_StartContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerInitClient) ScaleContainer(ctx context.Context, in *ScaleContainerRequest, opts ...grpc.CallOption) (*ScaleContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScaleContainerResponse)
	err := c.cc.Invoke(ctx, ContainerInit_ScaleContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerInitClient) KillContainer(ctx context.Context, in *KillContainerRequest, opts ...grpc.CallOption) (*KillContainerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KillContainerResponse)
	err := c.cc.Invoke(ctx, ContainerInit_KillContainer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerInitServer is the server API for ContainerInit service.
// All implementations must embed UnimplementedContainerInitServer
// for forward compatibility
type ContainerInitServer interface {
	StartContainer(context.Context, *StartContainerRequest) (*StartContainerResponse, error)
	ScaleContainer(context.Context, *ScaleContainerRequest) (*ScaleContainerResponse, error)
	KillContainer(context.Context, *KillContainerRequest) (*KillContainerResponse, error)
	mustEmbedUnimplementedContainerInitServer()
}

// UnimplementedContainerInitServer must be embedded to have forward compatible implementations.
type UnimplementedContainerInitServer struct {
}

func (UnimplementedContainerInitServer) StartContainer(context.Context, *StartContainerRequest) (*StartContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartContainer not implemented")
}
func (UnimplementedContainerInitServer) ScaleContainer(context.Context, *ScaleContainerRequest) (*ScaleContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleContainer not implemented")
}
func (UnimplementedContainerInitServer) KillContainer(context.Context, *KillContainerRequest) (*KillContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillContainer not implemented")
}
func (UnimplementedContainerInitServer) mustEmbedUnimplementedContainerInitServer() {}

// UnsafeContainerInitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerInitServer will
// result in compilation errors.
type UnsafeContainerInitServer interface {
	mustEmbedUnimplementedContainerInitServer()
}

func RegisterContainerInitServer(s grpc.ServiceRegistrar, srv ContainerInitServer) {
	s.RegisterService(&ContainerInit_ServiceDesc, srv)
}

func _ContainerInit_StartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerInitServer).StartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContainerInit_StartContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerInitServer).StartContainer(ctx, req.(*StartContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerInit_ScaleContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerInitServer).ScaleContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContainerInit_ScaleContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerInitServer).ScaleContainer(ctx, req.(*ScaleContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerInit_KillContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerInitServer).KillContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContainerInit_KillContainer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerInitServer).KillContainer(ctx, req.(*KillContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContainerInit_ServiceDesc is the grpc.ServiceDesc for ContainerInit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerInit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerinit.ContainerInit",
	HandlerType: (*ContainerInitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartContainer",
			Handler:    _ContainerInit_StartContainer_Handler,
		},
		{
			MethodName: "ScaleContainer",
			Handler:    _ContainerInit_ScaleContainer_Handler,
		},
		{
			MethodName: "KillContainer",
			Handler:    _ContainerInit_KillContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/containerinit/a.proto",
}
