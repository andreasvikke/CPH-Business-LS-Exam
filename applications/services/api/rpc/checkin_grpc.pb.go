// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CheckInProtoClient is the client API for CheckInProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckInProtoClient interface {
	GetCheckInById(ctx context.Context, in *wrappers.Int64Value, opts ...grpc.CallOption) (*CheckIn, error)
	GetAllCheckIns(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckIns, error)
}

type checkInProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckInProtoClient(cc grpc.ClientConnInterface) CheckInProtoClient {
	return &checkInProtoClient{cc}
}

func (c *checkInProtoClient) GetCheckInById(ctx context.Context, in *wrappers.Int64Value, opts ...grpc.CallOption) (*CheckIn, error) {
	out := new(CheckIn)
	err := c.cc.Invoke(ctx, "/rpc.CheckInProto/GetCheckInById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInProtoClient) GetAllCheckIns(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckIns, error) {
	out := new(CheckIns)
	err := c.cc.Invoke(ctx, "/rpc.CheckInProto/GetAllCheckIns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckInProtoServer is the server API for CheckInProto service.
// All implementations must embed UnimplementedCheckInProtoServer
// for forward compatibility
type CheckInProtoServer interface {
	GetCheckInById(context.Context, *wrappers.Int64Value) (*CheckIn, error)
	GetAllCheckIns(context.Context, *empty.Empty) (*CheckIns, error)
	mustEmbedUnimplementedCheckInProtoServer()
}

// UnimplementedCheckInProtoServer must be embedded to have forward compatible implementations.
type UnimplementedCheckInProtoServer struct {
}

func (UnimplementedCheckInProtoServer) GetCheckInById(context.Context, *wrappers.Int64Value) (*CheckIn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckInById not implemented")
}
func (UnimplementedCheckInProtoServer) GetAllCheckIns(context.Context, *empty.Empty) (*CheckIns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCheckIns not implemented")
}
func (UnimplementedCheckInProtoServer) mustEmbedUnimplementedCheckInProtoServer() {}

// UnsafeCheckInProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckInProtoServer will
// result in compilation errors.
type UnsafeCheckInProtoServer interface {
	mustEmbedUnimplementedCheckInProtoServer()
}

func RegisterCheckInProtoServer(s grpc.ServiceRegistrar, srv CheckInProtoServer) {
	s.RegisterService(&CheckInProto_ServiceDesc, srv)
}

func _CheckInProto_GetCheckInById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInProtoServer).GetCheckInById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInProto/GetCheckInById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInProtoServer).GetCheckInById(ctx, req.(*wrappers.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInProto_GetAllCheckIns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInProtoServer).GetAllCheckIns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInProto/GetAllCheckIns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInProtoServer).GetAllCheckIns(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckInProto_ServiceDesc is the grpc.ServiceDesc for CheckInProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckInProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.CheckInProto",
	HandlerType: (*CheckInProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCheckInById",
			Handler:    _CheckInProto_GetCheckInById_Handler,
		},
		{
			MethodName: "GetAllCheckIns",
			Handler:    _CheckInProto_GetAllCheckIns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkin.proto",
}