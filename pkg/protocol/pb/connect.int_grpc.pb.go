// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: connect.int.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ConnectInt_DeliverMessage_FullMethodName = "/pb.ConnectInt/DeliverMessage"
)

// ConnectIntClient is the client API for ConnectInt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectIntClient interface {
	// 消息投递
	DeliverMessage(ctx context.Context, in *DeliverMessageReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type connectIntClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectIntClient(cc grpc.ClientConnInterface) ConnectIntClient {
	return &connectIntClient{cc}
}

func (c *connectIntClient) DeliverMessage(ctx context.Context, in *DeliverMessageReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ConnectInt_DeliverMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectIntServer is the server API for ConnectInt service.
// All implementations must embed UnimplementedConnectIntServer
// for forward compatibility
type ConnectIntServer interface {
	// 消息投递
	DeliverMessage(context.Context, *DeliverMessageReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedConnectIntServer()
}

// UnimplementedConnectIntServer must be embedded to have forward compatible implementations.
type UnimplementedConnectIntServer struct {
}

func (UnimplementedConnectIntServer) DeliverMessage(context.Context, *DeliverMessageReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverMessage not implemented")
}
func (UnimplementedConnectIntServer) mustEmbedUnimplementedConnectIntServer() {}

// UnsafeConnectIntServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectIntServer will
// result in compilation errors.
type UnsafeConnectIntServer interface {
	mustEmbedUnimplementedConnectIntServer()
}

func RegisterConnectIntServer(s grpc.ServiceRegistrar, srv ConnectIntServer) {
	s.RegisterService(&ConnectInt_ServiceDesc, srv)
}

func _ConnectInt_DeliverMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectIntServer).DeliverMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConnectInt_DeliverMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectIntServer).DeliverMessage(ctx, req.(*DeliverMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectInt_ServiceDesc is the grpc.ServiceDesc for ConnectInt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectInt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ConnectInt",
	HandlerType: (*ConnectIntServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliverMessage",
			Handler:    _ConnectInt_DeliverMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connect.int.proto",
}
