// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: some-ms.proto

package __

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
	SomeMsService_Create_FullMethodName = "/protobuf.SomeMsService/Create"
	SomeMsService_Update_FullMethodName = "/protobuf.SomeMsService/Update"
	SomeMsService_Delete_FullMethodName = "/protobuf.SomeMsService/Delete"
	SomeMsService_Get_FullMethodName    = "/protobuf.SomeMsService/Get"
)

// SomeMsServiceClient is the client API for SomeMsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SomeMsServiceClient interface {
	Create(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	Update(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	Delete(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
	//Возвращает список записей
	Get(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
}

type someMsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSomeMsServiceClient(cc grpc.ClientConnInterface) SomeMsServiceClient {
	return &someMsServiceClient{cc}
}

func (c *someMsServiceClient) Create(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, SomeMsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *someMsServiceClient) Update(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, SomeMsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *someMsServiceClient) Delete(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, SomeMsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *someMsServiceClient) Get(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, SomeMsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SomeMsServiceServer is the server API for SomeMsService service.
// All implementations must embed UnimplementedSomeMsServiceServer
// for forward compatibility
type SomeMsServiceServer interface {
	Create(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	Update(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	Delete(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	//Возвращает список записей
	Get(context.Context, *GetItemRequest) (*GetItemResponse, error)
	mustEmbedUnimplementedSomeMsServiceServer()
}

// UnimplementedSomeMsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSomeMsServiceServer struct {
}

func (UnimplementedSomeMsServiceServer) Create(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSomeMsServiceServer) Update(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSomeMsServiceServer) Delete(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSomeMsServiceServer) Get(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSomeMsServiceServer) mustEmbedUnimplementedSomeMsServiceServer() {}

// UnsafeSomeMsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SomeMsServiceServer will
// result in compilation errors.
type UnsafeSomeMsServiceServer interface {
	mustEmbedUnimplementedSomeMsServiceServer()
}

func RegisterSomeMsServiceServer(s grpc.ServiceRegistrar, srv SomeMsServiceServer) {
	s.RegisterService(&SomeMsService_ServiceDesc, srv)
}

func _SomeMsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeMsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeMsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeMsServiceServer).Create(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SomeMsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeMsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeMsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeMsServiceServer).Update(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SomeMsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeMsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeMsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeMsServiceServer).Delete(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SomeMsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeMsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeMsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeMsServiceServer).Get(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SomeMsService_ServiceDesc is the grpc.ServiceDesc for SomeMsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SomeMsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.SomeMsService",
	HandlerType: (*SomeMsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SomeMsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SomeMsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SomeMsService_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SomeMsService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "some-ms.proto",
}
