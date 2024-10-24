// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: network/profile/v1/query.proto

package profilev1

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
	Query_Params_FullMethodName                        = "/network.profile.v1.Query/Params"
	Query_GetCoordinator_FullMethodName                = "/network.profile.v1.Query/GetCoordinator"
	Query_ListCoordinator_FullMethodName               = "/network.profile.v1.Query/ListCoordinator"
	Query_GetValidator_FullMethodName                  = "/network.profile.v1.Query/GetValidator"
	Query_ListValidator_FullMethodName                 = "/network.profile.v1.Query/ListValidator"
	Query_GetCoordinatorByAddress_FullMethodName       = "/network.profile.v1.Query/GetCoordinatorByAddress"
	Query_GetValidatorByOperatorAddress_FullMethodName = "/network.profile.v1.Query/GetValidatorByOperatorAddress"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Coordinator items.
	GetCoordinator(ctx context.Context, in *QueryGetCoordinatorRequest, opts ...grpc.CallOption) (*QueryGetCoordinatorResponse, error)
	ListCoordinator(ctx context.Context, in *QueryAllCoordinatorRequest, opts ...grpc.CallOption) (*QueryAllCoordinatorResponse, error)
	// Queries a list of Validator items.
	GetValidator(ctx context.Context, in *QueryGetValidatorRequest, opts ...grpc.CallOption) (*QueryGetValidatorResponse, error)
	ListValidator(ctx context.Context, in *QueryAllValidatorRequest, opts ...grpc.CallOption) (*QueryAllValidatorResponse, error)
	// Queries a list of GetCoordinatorByAddress items.
	GetCoordinatorByAddress(ctx context.Context, in *QueryGetCoordinatorByAddressRequest, opts ...grpc.CallOption) (*QueryGetCoordinatorByAddressResponse, error)
	// Queries a list of GetValidatorByOperatorAddress items.
	GetValidatorByOperatorAddress(ctx context.Context, in *QueryGetValidatorByOperatorAddressRequest, opts ...grpc.CallOption) (*QueryGetValidatorByOperatorAddressResponse, error)
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

func (c *queryClient) GetCoordinator(ctx context.Context, in *QueryGetCoordinatorRequest, opts ...grpc.CallOption) (*QueryGetCoordinatorResponse, error) {
	out := new(QueryGetCoordinatorResponse)
	err := c.cc.Invoke(ctx, Query_GetCoordinator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListCoordinator(ctx context.Context, in *QueryAllCoordinatorRequest, opts ...grpc.CallOption) (*QueryAllCoordinatorResponse, error) {
	out := new(QueryAllCoordinatorResponse)
	err := c.cc.Invoke(ctx, Query_ListCoordinator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetValidator(ctx context.Context, in *QueryGetValidatorRequest, opts ...grpc.CallOption) (*QueryGetValidatorResponse, error) {
	out := new(QueryGetValidatorResponse)
	err := c.cc.Invoke(ctx, Query_GetValidator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListValidator(ctx context.Context, in *QueryAllValidatorRequest, opts ...grpc.CallOption) (*QueryAllValidatorResponse, error) {
	out := new(QueryAllValidatorResponse)
	err := c.cc.Invoke(ctx, Query_ListValidator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetCoordinatorByAddress(ctx context.Context, in *QueryGetCoordinatorByAddressRequest, opts ...grpc.CallOption) (*QueryGetCoordinatorByAddressResponse, error) {
	out := new(QueryGetCoordinatorByAddressResponse)
	err := c.cc.Invoke(ctx, Query_GetCoordinatorByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetValidatorByOperatorAddress(ctx context.Context, in *QueryGetValidatorByOperatorAddressRequest, opts ...grpc.CallOption) (*QueryGetValidatorByOperatorAddressResponse, error) {
	out := new(QueryGetValidatorByOperatorAddressResponse)
	err := c.cc.Invoke(ctx, Query_GetValidatorByOperatorAddress_FullMethodName, in, out, opts...)
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
	// Queries a list of Coordinator items.
	GetCoordinator(context.Context, *QueryGetCoordinatorRequest) (*QueryGetCoordinatorResponse, error)
	ListCoordinator(context.Context, *QueryAllCoordinatorRequest) (*QueryAllCoordinatorResponse, error)
	// Queries a list of Validator items.
	GetValidator(context.Context, *QueryGetValidatorRequest) (*QueryGetValidatorResponse, error)
	ListValidator(context.Context, *QueryAllValidatorRequest) (*QueryAllValidatorResponse, error)
	// Queries a list of GetCoordinatorByAddress items.
	GetCoordinatorByAddress(context.Context, *QueryGetCoordinatorByAddressRequest) (*QueryGetCoordinatorByAddressResponse, error)
	// Queries a list of GetValidatorByOperatorAddress items.
	GetValidatorByOperatorAddress(context.Context, *QueryGetValidatorByOperatorAddressRequest) (*QueryGetValidatorByOperatorAddressResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetCoordinator(context.Context, *QueryGetCoordinatorRequest) (*QueryGetCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoordinator not implemented")
}
func (UnimplementedQueryServer) ListCoordinator(context.Context, *QueryAllCoordinatorRequest) (*QueryAllCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCoordinator not implemented")
}
func (UnimplementedQueryServer) GetValidator(context.Context, *QueryGetValidatorRequest) (*QueryGetValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidator not implemented")
}
func (UnimplementedQueryServer) ListValidator(context.Context, *QueryAllValidatorRequest) (*QueryAllValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListValidator not implemented")
}
func (UnimplementedQueryServer) GetCoordinatorByAddress(context.Context, *QueryGetCoordinatorByAddressRequest) (*QueryGetCoordinatorByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoordinatorByAddress not implemented")
}
func (UnimplementedQueryServer) GetValidatorByOperatorAddress(context.Context, *QueryGetValidatorByOperatorAddressRequest) (*QueryGetValidatorByOperatorAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorByOperatorAddress not implemented")
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

func _Query_GetCoordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCoordinatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetCoordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetCoordinator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetCoordinator(ctx, req.(*QueryGetCoordinatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListCoordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCoordinatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListCoordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListCoordinator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListCoordinator(ctx, req.(*QueryAllCoordinatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetValidator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetValidator(ctx, req.(*QueryGetValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListValidator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListValidator(ctx, req.(*QueryAllValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetCoordinatorByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCoordinatorByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetCoordinatorByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetCoordinatorByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetCoordinatorByAddress(ctx, req.(*QueryGetCoordinatorByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetValidatorByOperatorAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetValidatorByOperatorAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetValidatorByOperatorAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetValidatorByOperatorAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetValidatorByOperatorAddress(ctx, req.(*QueryGetValidatorByOperatorAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network.profile.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetCoordinator",
			Handler:    _Query_GetCoordinator_Handler,
		},
		{
			MethodName: "ListCoordinator",
			Handler:    _Query_ListCoordinator_Handler,
		},
		{
			MethodName: "GetValidator",
			Handler:    _Query_GetValidator_Handler,
		},
		{
			MethodName: "ListValidator",
			Handler:    _Query_ListValidator_Handler,
		},
		{
			MethodName: "GetCoordinatorByAddress",
			Handler:    _Query_GetCoordinatorByAddress_Handler,
		},
		{
			MethodName: "GetValidatorByOperatorAddress",
			Handler:    _Query_GetValidatorByOperatorAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network/profile/v1/query.proto",
}
