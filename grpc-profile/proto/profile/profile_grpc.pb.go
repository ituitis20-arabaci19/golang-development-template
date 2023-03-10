// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/profile/profile.proto

package profile

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

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	// create a profile
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateMulti(ctx context.Context, opts ...grpc.CallOption) (ProfileService_CreateMultiClient, error)
	ReadUser(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (ProfileService_ReadUserClient, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) CreateMulti(ctx context.Context, opts ...grpc.CallOption) (ProfileService_CreateMultiClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProfileService_ServiceDesc.Streams[0], "/profile.ProfileService/CreateMulti", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileServiceCreateMultiClient{stream}
	return x, nil
}

type ProfileService_CreateMultiClient interface {
	Send(*CreateUserRequest) error
	CloseAndRecv() (*CreateUserResponse, error)
	grpc.ClientStream
}

type profileServiceCreateMultiClient struct {
	grpc.ClientStream
}

func (x *profileServiceCreateMultiClient) Send(m *CreateUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profileServiceCreateMultiClient) CloseAndRecv() (*CreateUserResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileServiceClient) ReadUser(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (ProfileService_ReadUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProfileService_ServiceDesc.Streams[1], "/profile.ProfileService/ReadUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileServiceReadUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProfileService_ReadUserClient interface {
	Recv() (*ReadUserResponse, error)
	grpc.ClientStream
}

type profileServiceReadUserClient struct {
	grpc.ClientStream
}

func (x *profileServiceReadUserClient) Recv() (*ReadUserResponse, error) {
	m := new(ReadUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	// create a profile
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateMulti(ProfileService_CreateMultiServer) error
	ReadUser(*ReadUserRequest, ProfileService_ReadUserServer) error
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProfileServiceServer) CreateMulti(ProfileService_CreateMultiServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateMulti not implemented")
}
func (UnimplementedProfileServiceServer) ReadUser(*ReadUserRequest, ProfileService_ReadUserServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadUser not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_CreateMulti_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfileServiceServer).CreateMulti(&profileServiceCreateMultiServer{stream})
}

type ProfileService_CreateMultiServer interface {
	SendAndClose(*CreateUserResponse) error
	Recv() (*CreateUserRequest, error)
	grpc.ServerStream
}

type profileServiceCreateMultiServer struct {
	grpc.ServerStream
}

func (x *profileServiceCreateMultiServer) SendAndClose(m *CreateUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profileServiceCreateMultiServer) Recv() (*CreateUserRequest, error) {
	m := new(CreateUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfileService_ReadUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadUserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProfileServiceServer).ReadUser(m, &profileServiceReadUserServer{stream})
}

type ProfileService_ReadUserServer interface {
	Send(*ReadUserResponse) error
	grpc.ServerStream
}

type profileServiceReadUserServer struct {
	grpc.ServerStream
}

func (x *profileServiceReadUserServer) Send(m *ReadUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProfileService_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateMulti",
			Handler:       _ProfileService_CreateMulti_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ReadUser",
			Handler:       _ProfileService_ReadUser_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/profile/profile.proto",
}
