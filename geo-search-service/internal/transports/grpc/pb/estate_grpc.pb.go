// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: estate.proto

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

// EstateServiceClient is the client API for EstateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EstateServiceClient interface {
	CreateEstate(ctx context.Context, in *CreateEstateRequest, opts ...grpc.CallOption) (*CreateEstateResponse, error)
	GetEstateByBound(ctx context.Context, in *GetEstatesByBoundBoxRequest, opts ...grpc.CallOption) (*GetEstatesByBoundBoxResponse, error)
}

type estateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEstateServiceClient(cc grpc.ClientConnInterface) EstateServiceClient {
	return &estateServiceClient{cc}
}

func (c *estateServiceClient) CreateEstate(ctx context.Context, in *CreateEstateRequest, opts ...grpc.CallOption) (*CreateEstateResponse, error) {
	out := new(CreateEstateResponse)
	err := c.cc.Invoke(ctx, "/proto.EstateService/CreateEstate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *estateServiceClient) GetEstateByBound(ctx context.Context, in *GetEstatesByBoundBoxRequest, opts ...grpc.CallOption) (*GetEstatesByBoundBoxResponse, error) {
	out := new(GetEstatesByBoundBoxResponse)
	err := c.cc.Invoke(ctx, "/proto.EstateService/GetEstateByBound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EstateServiceServer is the server API for EstateService service.
// All implementations must embed UnimplementedEstateServiceServer
// for forward compatibility
type EstateServiceServer interface {
	CreateEstate(context.Context, *CreateEstateRequest) (*CreateEstateResponse, error)
	GetEstateByBound(context.Context, *GetEstatesByBoundBoxRequest) (*GetEstatesByBoundBoxResponse, error)
	mustEmbedUnimplementedEstateServiceServer()
}

// UnimplementedEstateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEstateServiceServer struct {
}

func (UnimplementedEstateServiceServer) CreateEstate(context.Context, *CreateEstateRequest) (*CreateEstateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEstate not implemented")
}
func (UnimplementedEstateServiceServer) GetEstateByBound(context.Context, *GetEstatesByBoundBoxRequest) (*GetEstatesByBoundBoxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEstateByBound not implemented")
}
func (UnimplementedEstateServiceServer) mustEmbedUnimplementedEstateServiceServer() {}

// UnsafeEstateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EstateServiceServer will
// result in compilation errors.
type UnsafeEstateServiceServer interface {
	mustEmbedUnimplementedEstateServiceServer()
}

func RegisterEstateServiceServer(s grpc.ServiceRegistrar, srv EstateServiceServer) {
	s.RegisterService(&EstateService_ServiceDesc, srv)
}

func _EstateService_CreateEstate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEstateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstateServiceServer).CreateEstate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EstateService/CreateEstate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstateServiceServer).CreateEstate(ctx, req.(*CreateEstateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EstateService_GetEstateByBound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEstatesByBoundBoxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstateServiceServer).GetEstateByBound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EstateService/GetEstateByBound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstateServiceServer).GetEstateByBound(ctx, req.(*GetEstatesByBoundBoxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EstateService_ServiceDesc is the grpc.ServiceDesc for EstateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EstateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EstateService",
	HandlerType: (*EstateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEstate",
			Handler:    _EstateService_CreateEstate_Handler,
		},
		{
			MethodName: "GetEstateByBound",
			Handler:    _EstateService_GetEstateByBound_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "estate.proto",
}
