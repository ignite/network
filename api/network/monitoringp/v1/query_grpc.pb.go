// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: network/monitoringp/v1/query.proto

package monitoringpv1

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
	Query_Params_FullMethodName                 = "/network.monitoringp.v1.Query/Params"
	Query_GetMonitoringInfo_FullMethodName      = "/network.monitoringp.v1.Query/GetMonitoringInfo"
	Query_GetConnectionChannelID_FullMethodName = "/network.monitoringp.v1.Query/GetConnectionChannelID"
	Query_GetConsumerClientID_FullMethodName    = "/network.monitoringp.v1.Query/GetConsumerClientID"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a MonitoringInfo by index.
	GetMonitoringInfo(ctx context.Context, in *QueryGetMonitoringInfoRequest, opts ...grpc.CallOption) (*QueryGetMonitoringInfoResponse, error)
	// Queries a ConnectionChannelID by index.
	GetConnectionChannelID(ctx context.Context, in *QueryGetConnectionChannelIDRequest, opts ...grpc.CallOption) (*QueryGetConnectionChannelIDResponse, error)
	// Queries a ConsumerClientID by index.
	GetConsumerClientID(ctx context.Context, in *QueryGetConsumerClientIDRequest, opts ...grpc.CallOption) (*QueryGetConsumerClientIDResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetMonitoringInfo(ctx context.Context, in *QueryGetMonitoringInfoRequest, opts ...grpc.CallOption) (*QueryGetMonitoringInfoResponse, error) {
	out := new(QueryGetMonitoringInfoResponse)
	err := c.cc.Invoke(ctx, Query_GetMonitoringInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetConnectionChannelID(ctx context.Context, in *QueryGetConnectionChannelIDRequest, opts ...grpc.CallOption) (*QueryGetConnectionChannelIDResponse, error) {
	out := new(QueryGetConnectionChannelIDResponse)
	err := c.cc.Invoke(ctx, Query_GetConnectionChannelID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetConsumerClientID(ctx context.Context, in *QueryGetConsumerClientIDRequest, opts ...grpc.CallOption) (*QueryGetConsumerClientIDResponse, error) {
	out := new(QueryGetConsumerClientIDResponse)
	err := c.cc.Invoke(ctx, Query_GetConsumerClientID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a MonitoringInfo by index.
	GetMonitoringInfo(context.Context, *QueryGetMonitoringInfoRequest) (*QueryGetMonitoringInfoResponse, error)
	// Queries a ConnectionChannelID by index.
	GetConnectionChannelID(context.Context, *QueryGetConnectionChannelIDRequest) (*QueryGetConnectionChannelIDResponse, error)
	// Queries a ConsumerClientID by index.
	GetConsumerClientID(context.Context, *QueryGetConsumerClientIDRequest) (*QueryGetConsumerClientIDResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetMonitoringInfo(context.Context, *QueryGetMonitoringInfoRequest) (*QueryGetMonitoringInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonitoringInfo not implemented")
}
func (UnimplementedQueryServer) GetConnectionChannelID(context.Context, *QueryGetConnectionChannelIDRequest) (*QueryGetConnectionChannelIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionChannelID not implemented")
}
func (UnimplementedQueryServer) GetConsumerClientID(context.Context, *QueryGetConsumerClientIDRequest) (*QueryGetConsumerClientIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerClientID not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetMonitoringInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMonitoringInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetMonitoringInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetMonitoringInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetMonitoringInfo(ctx, req.(*QueryGetMonitoringInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetConnectionChannelID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetConnectionChannelIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetConnectionChannelID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetConnectionChannelID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetConnectionChannelID(ctx, req.(*QueryGetConnectionChannelIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetConsumerClientID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetConsumerClientIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetConsumerClientID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetConsumerClientID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetConsumerClientID(ctx, req.(*QueryGetConsumerClientIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network.monitoringp.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetMonitoringInfo",
			Handler:    _Query_GetMonitoringInfo_Handler,
		},
		{
			MethodName: "GetConnectionChannelID",
			Handler:    _Query_GetConnectionChannelID_Handler,
		},
		{
			MethodName: "GetConsumerClientID",
			Handler:    _Query_GetConsumerClientID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network/monitoringp/v1/query.proto",
}