// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/user_service.proto

package pb

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

// UserserviceClient is the client API for Userservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserserviceClient interface {
	GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserserviceClient(cc grpc.ClientConnInterface) UserserviceClient {
	return &userserviceClient{cc}
}

func (c *userserviceClient) GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.Userservice/GetMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserserviceServer is the server API for Userservice service.
// All implementations must embed UnimplementedUserserviceServer
// for forward compatibility
type UserserviceServer interface {
	GetMe(context.Context, *GetMeRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserserviceServer()
}

// UnimplementedUserserviceServer must be embedded to have forward compatible implementations.
type UnimplementedUserserviceServer struct {
}

func (UnimplementedUserserviceServer) GetMe(context.Context, *GetMeRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedUserserviceServer) mustEmbedUnimplementedUserserviceServer() {}

// UnsafeUserserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserserviceServer will
// result in compilation errors.
type UnsafeUserserviceServer interface {
	mustEmbedUnimplementedUserserviceServer()
}

func RegisterUserserviceServer(s grpc.ServiceRegistrar, srv UserserviceServer) {
	s.RegisterService(&Userservice_ServiceDesc, srv)
}

func _Userservice_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserserviceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Userservice/GetMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserserviceServer).GetMe(ctx, req.(*GetMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Userservice_ServiceDesc is the grpc.ServiceDesc for Userservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Userservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Userservice",
	HandlerType: (*UserserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMe",
			Handler:    _Userservice_GetMe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user_service.proto",
}
