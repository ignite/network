// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: network/participation/v1/query.proto

package participationv1

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
	Query_Params_FullMethodName                     = "/network.participation.v1.Query/Params"
	Query_GetAuctionUsedAllocations_FullMethodName  = "/network.participation.v1.Query/GetAuctionUsedAllocations"
	Query_ListAuctionUsedAllocations_FullMethodName = "/network.participation.v1.Query/ListAuctionUsedAllocations"
	Query_GetUsedAllocations_FullMethodName         = "/network.participation.v1.Query/GetUsedAllocations"
	Query_ListUsedAllocations_FullMethodName        = "/network.participation.v1.Query/ListUsedAllocations"
	Query_TotalAllocations_FullMethodName           = "/network.participation.v1.Query/TotalAllocations"
	Query_AvailableAllocations_FullMethodName       = "/network.participation.v1.Query/AvailableAllocations"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of AuctionUsedAllocations items.
	GetAuctionUsedAllocations(ctx context.Context, in *QueryGetAuctionUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryGetAuctionUsedAllocationsResponse, error)
	ListAuctionUsedAllocations(ctx context.Context, in *QueryAllAuctionUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryAllAuctionUsedAllocationsResponse, error)
	// Queries a list of UsedAllocations items.
	GetUsedAllocations(ctx context.Context, in *QueryGetUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryGetUsedAllocationsResponse, error)
	ListUsedAllocations(ctx context.Context, in *QueryAllUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryAllUsedAllocationsResponse, error)
	// Queries a list of TotalAllocations items.
	TotalAllocations(ctx context.Context, in *QueryTotalAllocationsRequest, opts ...grpc.CallOption) (*QueryTotalAllocationsResponse, error)
	// Queries a list of AvailableAllocations items.
	AvailableAllocations(ctx context.Context, in *QueryAvailableAllocationsRequest, opts ...grpc.CallOption) (*QueryAvailableAllocationsResponse, error)
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

func (c *queryClient) GetAuctionUsedAllocations(ctx context.Context, in *QueryGetAuctionUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryGetAuctionUsedAllocationsResponse, error) {
	out := new(QueryGetAuctionUsedAllocationsResponse)
	err := c.cc.Invoke(ctx, Query_GetAuctionUsedAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListAuctionUsedAllocations(ctx context.Context, in *QueryAllAuctionUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryAllAuctionUsedAllocationsResponse, error) {
	out := new(QueryAllAuctionUsedAllocationsResponse)
	err := c.cc.Invoke(ctx, Query_ListAuctionUsedAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetUsedAllocations(ctx context.Context, in *QueryGetUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryGetUsedAllocationsResponse, error) {
	out := new(QueryGetUsedAllocationsResponse)
	err := c.cc.Invoke(ctx, Query_GetUsedAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListUsedAllocations(ctx context.Context, in *QueryAllUsedAllocationsRequest, opts ...grpc.CallOption) (*QueryAllUsedAllocationsResponse, error) {
	out := new(QueryAllUsedAllocationsResponse)
	err := c.cc.Invoke(ctx, Query_ListUsedAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalAllocations(ctx context.Context, in *QueryTotalAllocationsRequest, opts ...grpc.CallOption) (*QueryTotalAllocationsResponse, error) {
	out := new(QueryTotalAllocationsResponse)
	err := c.cc.Invoke(ctx, Query_TotalAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AvailableAllocations(ctx context.Context, in *QueryAvailableAllocationsRequest, opts ...grpc.CallOption) (*QueryAvailableAllocationsResponse, error) {
	out := new(QueryAvailableAllocationsResponse)
	err := c.cc.Invoke(ctx, Query_AvailableAllocations_FullMethodName, in, out, opts...)
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
	// Queries a list of AuctionUsedAllocations items.
	GetAuctionUsedAllocations(context.Context, *QueryGetAuctionUsedAllocationsRequest) (*QueryGetAuctionUsedAllocationsResponse, error)
	ListAuctionUsedAllocations(context.Context, *QueryAllAuctionUsedAllocationsRequest) (*QueryAllAuctionUsedAllocationsResponse, error)
	// Queries a list of UsedAllocations items.
	GetUsedAllocations(context.Context, *QueryGetUsedAllocationsRequest) (*QueryGetUsedAllocationsResponse, error)
	ListUsedAllocations(context.Context, *QueryAllUsedAllocationsRequest) (*QueryAllUsedAllocationsResponse, error)
	// Queries a list of TotalAllocations items.
	TotalAllocations(context.Context, *QueryTotalAllocationsRequest) (*QueryTotalAllocationsResponse, error)
	// Queries a list of AvailableAllocations items.
	AvailableAllocations(context.Context, *QueryAvailableAllocationsRequest) (*QueryAvailableAllocationsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetAuctionUsedAllocations(context.Context, *QueryGetAuctionUsedAllocationsRequest) (*QueryGetAuctionUsedAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuctionUsedAllocations not implemented")
}
func (UnimplementedQueryServer) ListAuctionUsedAllocations(context.Context, *QueryAllAuctionUsedAllocationsRequest) (*QueryAllAuctionUsedAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuctionUsedAllocations not implemented")
}
func (UnimplementedQueryServer) GetUsedAllocations(context.Context, *QueryGetUsedAllocationsRequest) (*QueryGetUsedAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsedAllocations not implemented")
}
func (UnimplementedQueryServer) ListUsedAllocations(context.Context, *QueryAllUsedAllocationsRequest) (*QueryAllUsedAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsedAllocations not implemented")
}
func (UnimplementedQueryServer) TotalAllocations(context.Context, *QueryTotalAllocationsRequest) (*QueryTotalAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalAllocations not implemented")
}
func (UnimplementedQueryServer) AvailableAllocations(context.Context, *QueryAvailableAllocationsRequest) (*QueryAvailableAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailableAllocations not implemented")
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

func _Query_GetAuctionUsedAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAuctionUsedAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAuctionUsedAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetAuctionUsedAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAuctionUsedAllocations(ctx, req.(*QueryGetAuctionUsedAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListAuctionUsedAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAuctionUsedAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListAuctionUsedAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListAuctionUsedAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListAuctionUsedAllocations(ctx, req.(*QueryAllAuctionUsedAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetUsedAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetUsedAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetUsedAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetUsedAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetUsedAllocations(ctx, req.(*QueryGetUsedAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListUsedAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllUsedAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListUsedAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListUsedAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListUsedAllocations(ctx, req.(*QueryAllUsedAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TotalAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalAllocations(ctx, req.(*QueryTotalAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AvailableAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAvailableAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AvailableAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AvailableAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AvailableAllocations(ctx, req.(*QueryAvailableAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network.participation.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetAuctionUsedAllocations",
			Handler:    _Query_GetAuctionUsedAllocations_Handler,
		},
		{
			MethodName: "ListAuctionUsedAllocations",
			Handler:    _Query_ListAuctionUsedAllocations_Handler,
		},
		{
			MethodName: "GetUsedAllocations",
			Handler:    _Query_GetUsedAllocations_Handler,
		},
		{
			MethodName: "ListUsedAllocations",
			Handler:    _Query_ListUsedAllocations_Handler,
		},
		{
			MethodName: "TotalAllocations",
			Handler:    _Query_TotalAllocations_Handler,
		},
		{
			MethodName: "AvailableAllocations",
			Handler:    _Query_AvailableAllocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network/participation/v1/query.proto",
}
