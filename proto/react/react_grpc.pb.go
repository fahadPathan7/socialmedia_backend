// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: proto/react/react.proto

package react

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReactService_CreateAReact_FullMethodName           = "/react.ReactService/CreateAReact"
	ReactService_ReadAReact_FullMethodName             = "/react.ReactService/ReadAReact"
	ReactService_ReadAllReactsOfAPost_FullMethodName   = "/react.ReactService/ReadAllReactsOfAPost"
	ReactService_UpdateAReact_FullMethodName           = "/react.ReactService/UpdateAReact"
	ReactService_DeleteAReact_FullMethodName           = "/react.ReactService/DeleteAReact"
	ReactService_DeleteAllReactsOfAPost_FullMethodName = "/react.ReactService/DeleteAllReactsOfAPost"
)

// ReactServiceClient is the client API for ReactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ReactService definition
type ReactServiceClient interface {
	CreateAReact(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	ReadAReact(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	ReadAllReactsOfAPost(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
	UpdateAReact(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	DeleteAReact(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteAllReactsOfAPost(ctx context.Context, in *DeleteAllReactsOfAPostRequest, opts ...grpc.CallOption) (*DeleteAllReactsOfAPostResponse, error)
}

type reactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReactServiceClient(cc grpc.ClientConnInterface) ReactServiceClient {
	return &reactServiceClient{cc}
}

func (c *reactServiceClient) CreateAReact(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ReactService_CreateAReact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactServiceClient) ReadAReact(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, ReactService_ReadAReact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactServiceClient) ReadAllReactsOfAPost(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, ReactService_ReadAllReactsOfAPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactServiceClient) UpdateAReact(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, ReactService_UpdateAReact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactServiceClient) DeleteAReact(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ReactService_DeleteAReact_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reactServiceClient) DeleteAllReactsOfAPost(ctx context.Context, in *DeleteAllReactsOfAPostRequest, opts ...grpc.CallOption) (*DeleteAllReactsOfAPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAllReactsOfAPostResponse)
	err := c.cc.Invoke(ctx, ReactService_DeleteAllReactsOfAPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReactServiceServer is the server API for ReactService service.
// All implementations must embed UnimplementedReactServiceServer
// for forward compatibility.
//
// ReactService definition
type ReactServiceServer interface {
	CreateAReact(context.Context, *CreateRequest) (*CreateResponse, error)
	ReadAReact(context.Context, *ReadRequest) (*ReadResponse, error)
	ReadAllReactsOfAPost(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	UpdateAReact(context.Context, *UpdateRequest) (*UpdateResponse, error)
	DeleteAReact(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteAllReactsOfAPost(context.Context, *DeleteAllReactsOfAPostRequest) (*DeleteAllReactsOfAPostResponse, error)
	mustEmbedUnimplementedReactServiceServer()
}

// UnimplementedReactServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReactServiceServer struct{}

func (UnimplementedReactServiceServer) CreateAReact(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAReact not implemented")
}
func (UnimplementedReactServiceServer) ReadAReact(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAReact not implemented")
}
func (UnimplementedReactServiceServer) ReadAllReactsOfAPost(context.Context, *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllReactsOfAPost not implemented")
}
func (UnimplementedReactServiceServer) UpdateAReact(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAReact not implemented")
}
func (UnimplementedReactServiceServer) DeleteAReact(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAReact not implemented")
}
func (UnimplementedReactServiceServer) DeleteAllReactsOfAPost(context.Context, *DeleteAllReactsOfAPostRequest) (*DeleteAllReactsOfAPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllReactsOfAPost not implemented")
}
func (UnimplementedReactServiceServer) mustEmbedUnimplementedReactServiceServer() {}
func (UnimplementedReactServiceServer) testEmbeddedByValue()                      {}

// UnsafeReactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReactServiceServer will
// result in compilation errors.
type UnsafeReactServiceServer interface {
	mustEmbedUnimplementedReactServiceServer()
}

func RegisterReactServiceServer(s grpc.ServiceRegistrar, srv ReactServiceServer) {
	// If the following call pancis, it indicates UnimplementedReactServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReactService_ServiceDesc, srv)
}

func _ReactService_CreateAReact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactServiceServer).CreateAReact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactService_CreateAReact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactServiceServer).CreateAReact(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactService_ReadAReact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactServiceServer).ReadAReact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactService_ReadAReact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactServiceServer).ReadAReact(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactService_ReadAllReactsOfAPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactServiceServer).ReadAllReactsOfAPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactService_ReadAllReactsOfAPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactServiceServer).ReadAllReactsOfAPost(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactService_UpdateAReact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactServiceServer).UpdateAReact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactService_UpdateAReact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactServiceServer).UpdateAReact(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactService_DeleteAReact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactServiceServer).DeleteAReact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactService_DeleteAReact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactServiceServer).DeleteAReact(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReactService_DeleteAllReactsOfAPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllReactsOfAPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReactServiceServer).DeleteAllReactsOfAPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReactService_DeleteAllReactsOfAPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReactServiceServer).DeleteAllReactsOfAPost(ctx, req.(*DeleteAllReactsOfAPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReactService_ServiceDesc is the grpc.ServiceDesc for ReactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "react.ReactService",
	HandlerType: (*ReactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAReact",
			Handler:    _ReactService_CreateAReact_Handler,
		},
		{
			MethodName: "ReadAReact",
			Handler:    _ReactService_ReadAReact_Handler,
		},
		{
			MethodName: "ReadAllReactsOfAPost",
			Handler:    _ReactService_ReadAllReactsOfAPost_Handler,
		},
		{
			MethodName: "UpdateAReact",
			Handler:    _ReactService_UpdateAReact_Handler,
		},
		{
			MethodName: "DeleteAReact",
			Handler:    _ReactService_DeleteAReact_Handler,
		},
		{
			MethodName: "DeleteAllReactsOfAPost",
			Handler:    _ReactService_DeleteAllReactsOfAPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/react/react.proto",
}
