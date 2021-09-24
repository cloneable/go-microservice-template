// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package healthz

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthzClient is the client API for Healthz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthzClient interface {
	Check(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type healthzClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthzClient(cc grpc.ClientConnInterface) HealthzClient {
	return &healthzClient{cc}
}

func (c *healthzClient) Check(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/go_microservice_template.healthz.Healthz/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthzServer is the server API for Healthz service.
// All implementations should embed UnimplementedHealthzServer
// for forward compatibility
type HealthzServer interface {
	Check(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error)
}

// UnimplementedHealthzServer should be embedded to have forward compatible implementations.
type UnimplementedHealthzServer struct {
}

func (UnimplementedHealthzServer) Check(context.Context, *emptypb.Empty) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

// UnsafeHealthzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthzServer will
// result in compilation errors.
type UnsafeHealthzServer interface {
	mustEmbedUnimplementedHealthzServer()
}

func RegisterHealthzServer(s grpc.ServiceRegistrar, srv HealthzServer) {
	s.RegisterService(&Healthz_ServiceDesc, srv)
}

func _Healthz_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_microservice_template.healthz.Healthz/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).Check(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Healthz_ServiceDesc is the grpc.ServiceDesc for Healthz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Healthz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_microservice_template.healthz.Healthz",
	HandlerType: (*HealthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Healthz_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/healthz/healthz.proto",
}
