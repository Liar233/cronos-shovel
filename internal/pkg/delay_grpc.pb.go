// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pkg/delay.proto

package pkg

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

const (
	DelayController_CreateDelay_FullMethodName = "/delay.DelayController/CreateDelay"
	DelayController_DeleteDelay_FullMethodName = "/delay.DelayController/DeleteDelay"
)

// DelayControllerClient is the client API for DelayController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DelayControllerClient interface {
	CreateDelay(ctx context.Context, in *CreateDelayRequest, opts ...grpc.CallOption) (*CreateDelayResponse, error)
	DeleteDelay(ctx context.Context, in *DeleteDelayRequest, opts ...grpc.CallOption) (*DeleteDelayResponse, error)
}

type delayControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewDelayControllerClient(cc grpc.ClientConnInterface) DelayControllerClient {
	return &delayControllerClient{cc}
}

func (c *delayControllerClient) CreateDelay(ctx context.Context, in *CreateDelayRequest, opts ...grpc.CallOption) (*CreateDelayResponse, error) {
	out := new(CreateDelayResponse)
	err := c.cc.Invoke(ctx, DelayController_CreateDelay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *delayControllerClient) DeleteDelay(ctx context.Context, in *DeleteDelayRequest, opts ...grpc.CallOption) (*DeleteDelayResponse, error) {
	out := new(DeleteDelayResponse)
	err := c.cc.Invoke(ctx, DelayController_DeleteDelay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelayControllerServer is the server API for DelayController service.
// All implementations must embed UnimplementedDelayControllerServer
// for forward compatibility
type DelayControllerServer interface {
	CreateDelay(context.Context, *CreateDelayRequest) (*CreateDelayResponse, error)
	DeleteDelay(context.Context, *DeleteDelayRequest) (*DeleteDelayResponse, error)
	mustEmbedUnimplementedDelayControllerServer()
}

// UnimplementedDelayControllerServer must be embedded to have forward compatible implementations.
type UnimplementedDelayControllerServer struct {
}

func (UnimplementedDelayControllerServer) CreateDelay(context.Context, *CreateDelayRequest) (*CreateDelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDelay not implemented")
}
func (UnimplementedDelayControllerServer) DeleteDelay(context.Context, *DeleteDelayRequest) (*DeleteDelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDelay not implemented")
}
func (UnimplementedDelayControllerServer) mustEmbedUnimplementedDelayControllerServer() {}

// UnsafeDelayControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DelayControllerServer will
// result in compilation errors.
type UnsafeDelayControllerServer interface {
	mustEmbedUnimplementedDelayControllerServer()
}

func RegisterDelayControllerServer(s grpc.ServiceRegistrar, srv DelayControllerServer) {
	s.RegisterService(&DelayController_ServiceDesc, srv)
}

func _DelayController_CreateDelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelayControllerServer).CreateDelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DelayController_CreateDelay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelayControllerServer).CreateDelay(ctx, req.(*CreateDelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DelayController_DeleteDelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelayControllerServer).DeleteDelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DelayController_DeleteDelay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelayControllerServer).DeleteDelay(ctx, req.(*DeleteDelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DelayController_ServiceDesc is the grpc.ServiceDesc for DelayController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DelayController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delay.DelayController",
	HandlerType: (*DelayControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDelay",
			Handler:    _DelayController_CreateDelay_Handler,
		},
		{
			MethodName: "DeleteDelay",
			Handler:    _DelayController_DeleteDelay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/delay.proto",
}