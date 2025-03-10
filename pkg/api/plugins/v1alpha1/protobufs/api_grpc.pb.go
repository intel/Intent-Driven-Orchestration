// To regenerate api.pb.go run hack/generate_protobuf.sh or make api in the root of the repository

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: pkg/api/plugins/v1alpha1/protobufs/api.proto

package plugins

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
	Registration_Register_FullMethodName = "/plugins.Registration/Register"
)

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegistrationStatusResponse, error)
}

type registrationClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationClient(cc grpc.ClientConnInterface) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegistrationStatusResponse, error) {
	out := new(RegistrationStatusResponse)
	err := c.cc.Invoke(ctx, Registration_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
// All implementations must embed UnimplementedRegistrationServer
// for forward compatibility
type RegistrationServer interface {
	Register(context.Context, *RegisterRequest) (*RegistrationStatusResponse, error)
	mustEmbedUnimplementedRegistrationServer()
}

// UnimplementedRegistrationServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServer struct {
}

func (UnimplementedRegistrationServer) Register(context.Context, *RegisterRequest) (*RegistrationStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistrationServer) mustEmbedUnimplementedRegistrationServer() {}

// UnsafeRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServer will
// result in compilation errors.
type UnsafeRegistrationServer interface {
	mustEmbedUnimplementedRegistrationServer()
}

func RegisterRegistrationServer(s grpc.ServiceRegistrar, srv RegistrationServer) {
	s.RegisterService(&Registration_ServiceDesc, srv)
}

func _Registration_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registration_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registration_ServiceDesc is the grpc.ServiceDesc for Registration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registration_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/plugins/v1alpha1/protobufs/api.proto",
}

const (
	ActuatorPlugin_NextState_FullMethodName = "/plugins.ActuatorPlugin/NextState"
	ActuatorPlugin_Perform_FullMethodName   = "/plugins.ActuatorPlugin/Perform"
	ActuatorPlugin_Effect_FullMethodName    = "/plugins.ActuatorPlugin/Effect"
)

// ActuatorPluginClient is the client API for ActuatorPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActuatorPluginClient interface {
	// NextState should return a set of potential follow-up states for a given state if this actuator would potentially be used.
	NextState(ctx context.Context, opts ...grpc.CallOption) (ActuatorPlugin_NextStateClient, error)
	// Perform should perform those actions of the plan that it is in charge of
	Perform(ctx context.Context, in *PerformRequest, opts ...grpc.CallOption) (*Empty, error)
	// Effect should (optionally) recalculate the effect this actuator has for ALL objectives for this workload.
	Effect(ctx context.Context, in *EffectRequest, opts ...grpc.CallOption) (*Empty, error)
}

type actuatorPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewActuatorPluginClient(cc grpc.ClientConnInterface) ActuatorPluginClient {
	return &actuatorPluginClient{cc}
}

func (c *actuatorPluginClient) NextState(ctx context.Context, opts ...grpc.CallOption) (ActuatorPlugin_NextStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &ActuatorPlugin_ServiceDesc.Streams[0], ActuatorPlugin_NextState_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &actuatorPluginNextStateClient{stream}
	return x, nil
}

type ActuatorPlugin_NextStateClient interface {
	Send(*NextStateRequest) error
	Recv() (*NextStateResponse, error)
	grpc.ClientStream
}

type actuatorPluginNextStateClient struct {
	grpc.ClientStream
}

func (x *actuatorPluginNextStateClient) Send(m *NextStateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *actuatorPluginNextStateClient) Recv() (*NextStateResponse, error) {
	m := new(NextStateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *actuatorPluginClient) Perform(ctx context.Context, in *PerformRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ActuatorPlugin_Perform_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actuatorPluginClient) Effect(ctx context.Context, in *EffectRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ActuatorPlugin_Effect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActuatorPluginServer is the server API for ActuatorPlugin service.
// All implementations must embed UnimplementedActuatorPluginServer
// for forward compatibility
type ActuatorPluginServer interface {
	// NextState should return a set of potential follow-up states for a given state if this actuator would potentially be used.
	NextState(ActuatorPlugin_NextStateServer) error
	// Perform should perform those actions of the plan that it is in charge of
	Perform(context.Context, *PerformRequest) (*Empty, error)
	// Effect should (optionally) recalculate the effect this actuator has for ALL objectives for this workload.
	Effect(context.Context, *EffectRequest) (*Empty, error)
	mustEmbedUnimplementedActuatorPluginServer()
}

// UnimplementedActuatorPluginServer must be embedded to have forward compatible implementations.
type UnimplementedActuatorPluginServer struct {
}

func (UnimplementedActuatorPluginServer) NextState(ActuatorPlugin_NextStateServer) error {
	return status.Errorf(codes.Unimplemented, "method NextState not implemented")
}
func (UnimplementedActuatorPluginServer) Perform(context.Context, *PerformRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Perform not implemented")
}
func (UnimplementedActuatorPluginServer) Effect(context.Context, *EffectRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Effect not implemented")
}
func (UnimplementedActuatorPluginServer) mustEmbedUnimplementedActuatorPluginServer() {}

// UnsafeActuatorPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActuatorPluginServer will
// result in compilation errors.
type UnsafeActuatorPluginServer interface {
	mustEmbedUnimplementedActuatorPluginServer()
}

func RegisterActuatorPluginServer(s grpc.ServiceRegistrar, srv ActuatorPluginServer) {
	s.RegisterService(&ActuatorPlugin_ServiceDesc, srv)
}

func _ActuatorPlugin_NextState_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ActuatorPluginServer).NextState(&actuatorPluginNextStateServer{stream})
}

type ActuatorPlugin_NextStateServer interface {
	Send(*NextStateResponse) error
	Recv() (*NextStateRequest, error)
	grpc.ServerStream
}

type actuatorPluginNextStateServer struct {
	grpc.ServerStream
}

func (x *actuatorPluginNextStateServer) Send(m *NextStateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *actuatorPluginNextStateServer) Recv() (*NextStateRequest, error) {
	m := new(NextStateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ActuatorPlugin_Perform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActuatorPluginServer).Perform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActuatorPlugin_Perform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActuatorPluginServer).Perform(ctx, req.(*PerformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActuatorPlugin_Effect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EffectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActuatorPluginServer).Effect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActuatorPlugin_Effect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActuatorPluginServer).Effect(ctx, req.(*EffectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActuatorPlugin_ServiceDesc is the grpc.ServiceDesc for ActuatorPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActuatorPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.ActuatorPlugin",
	HandlerType: (*ActuatorPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Perform",
			Handler:    _ActuatorPlugin_Perform_Handler,
		},
		{
			MethodName: "Effect",
			Handler:    _ActuatorPlugin_Effect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NextState",
			Handler:       _ActuatorPlugin_NextState_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/api/plugins/v1alpha1/protobufs/api.proto",
}
