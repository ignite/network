// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: network/project/v1/tx.proto

package projectv1

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
	Msg_UpdateParams_FullMethodName             = "/network.project.v1.Msg/UpdateParams"
	Msg_CreateProject_FullMethodName            = "/network.project.v1.Msg/CreateProject"
	Msg_EditProject_FullMethodName              = "/network.project.v1.Msg/EditProject"
	Msg_UpdateTotalSupply_FullMethodName        = "/network.project.v1.Msg/UpdateTotalSupply"
	Msg_UpdateSpecialAllocations_FullMethodName = "/network.project.v1.Msg/UpdateSpecialAllocations"
	Msg_InitializeMainnet_FullMethodName        = "/network.project.v1.Msg/InitializeMainnet"
	Msg_MintVouchers_FullMethodName             = "/network.project.v1.Msg/MintVouchers"
	Msg_BurnVouchers_FullMethodName             = "/network.project.v1.Msg/BurnVouchers"
	Msg_RedeemVouchers_FullMethodName           = "/network.project.v1.Msg/RedeemVouchers"
	Msg_UnredeemVouchers_FullMethodName         = "/network.project.v1.Msg/UnredeemVouchers"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateProject(ctx context.Context, in *MsgCreateProject, opts ...grpc.CallOption) (*MsgCreateProjectResponse, error)
	EditProject(ctx context.Context, in *MsgEditProject, opts ...grpc.CallOption) (*MsgEditProjectResponse, error)
	UpdateTotalSupply(ctx context.Context, in *MsgUpdateTotalSupply, opts ...grpc.CallOption) (*MsgUpdateTotalSupplyResponse, error)
	UpdateSpecialAllocations(ctx context.Context, in *MsgUpdateSpecialAllocations, opts ...grpc.CallOption) (*MsgUpdateSpecialAllocationsResponse, error)
	InitializeMainnet(ctx context.Context, in *MsgInitializeMainnet, opts ...grpc.CallOption) (*MsgInitializeMainnetResponse, error)
	MintVouchers(ctx context.Context, in *MsgMintVouchers, opts ...grpc.CallOption) (*MsgMintVouchersResponse, error)
	BurnVouchers(ctx context.Context, in *MsgBurnVouchers, opts ...grpc.CallOption) (*MsgBurnVouchersResponse, error)
	RedeemVouchers(ctx context.Context, in *MsgRedeemVouchers, opts ...grpc.CallOption) (*MsgRedeemVouchersResponse, error)
	UnredeemVouchers(ctx context.Context, in *MsgUnredeemVouchers, opts ...grpc.CallOption) (*MsgUnredeemVouchersResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateProject(ctx context.Context, in *MsgCreateProject, opts ...grpc.CallOption) (*MsgCreateProjectResponse, error) {
	out := new(MsgCreateProjectResponse)
	err := c.cc.Invoke(ctx, Msg_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditProject(ctx context.Context, in *MsgEditProject, opts ...grpc.CallOption) (*MsgEditProjectResponse, error) {
	out := new(MsgEditProjectResponse)
	err := c.cc.Invoke(ctx, Msg_EditProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateTotalSupply(ctx context.Context, in *MsgUpdateTotalSupply, opts ...grpc.CallOption) (*MsgUpdateTotalSupplyResponse, error) {
	out := new(MsgUpdateTotalSupplyResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateTotalSupply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSpecialAllocations(ctx context.Context, in *MsgUpdateSpecialAllocations, opts ...grpc.CallOption) (*MsgUpdateSpecialAllocationsResponse, error) {
	out := new(MsgUpdateSpecialAllocationsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateSpecialAllocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InitializeMainnet(ctx context.Context, in *MsgInitializeMainnet, opts ...grpc.CallOption) (*MsgInitializeMainnetResponse, error) {
	out := new(MsgInitializeMainnetResponse)
	err := c.cc.Invoke(ctx, Msg_InitializeMainnet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintVouchers(ctx context.Context, in *MsgMintVouchers, opts ...grpc.CallOption) (*MsgMintVouchersResponse, error) {
	out := new(MsgMintVouchersResponse)
	err := c.cc.Invoke(ctx, Msg_MintVouchers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnVouchers(ctx context.Context, in *MsgBurnVouchers, opts ...grpc.CallOption) (*MsgBurnVouchersResponse, error) {
	out := new(MsgBurnVouchersResponse)
	err := c.cc.Invoke(ctx, Msg_BurnVouchers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RedeemVouchers(ctx context.Context, in *MsgRedeemVouchers, opts ...grpc.CallOption) (*MsgRedeemVouchersResponse, error) {
	out := new(MsgRedeemVouchersResponse)
	err := c.cc.Invoke(ctx, Msg_RedeemVouchers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnredeemVouchers(ctx context.Context, in *MsgUnredeemVouchers, opts ...grpc.CallOption) (*MsgUnredeemVouchersResponse, error) {
	out := new(MsgUnredeemVouchersResponse)
	err := c.cc.Invoke(ctx, Msg_UnredeemVouchers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateProject(context.Context, *MsgCreateProject) (*MsgCreateProjectResponse, error)
	EditProject(context.Context, *MsgEditProject) (*MsgEditProjectResponse, error)
	UpdateTotalSupply(context.Context, *MsgUpdateTotalSupply) (*MsgUpdateTotalSupplyResponse, error)
	UpdateSpecialAllocations(context.Context, *MsgUpdateSpecialAllocations) (*MsgUpdateSpecialAllocationsResponse, error)
	InitializeMainnet(context.Context, *MsgInitializeMainnet) (*MsgInitializeMainnetResponse, error)
	MintVouchers(context.Context, *MsgMintVouchers) (*MsgMintVouchersResponse, error)
	BurnVouchers(context.Context, *MsgBurnVouchers) (*MsgBurnVouchersResponse, error)
	RedeemVouchers(context.Context, *MsgRedeemVouchers) (*MsgRedeemVouchersResponse, error)
	UnredeemVouchers(context.Context, *MsgUnredeemVouchers) (*MsgUnredeemVouchersResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateProject(context.Context, *MsgCreateProject) (*MsgCreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedMsgServer) EditProject(context.Context, *MsgEditProject) (*MsgEditProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProject not implemented")
}
func (UnimplementedMsgServer) UpdateTotalSupply(context.Context, *MsgUpdateTotalSupply) (*MsgUpdateTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTotalSupply not implemented")
}
func (UnimplementedMsgServer) UpdateSpecialAllocations(context.Context, *MsgUpdateSpecialAllocations) (*MsgUpdateSpecialAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSpecialAllocations not implemented")
}
func (UnimplementedMsgServer) InitializeMainnet(context.Context, *MsgInitializeMainnet) (*MsgInitializeMainnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeMainnet not implemented")
}
func (UnimplementedMsgServer) MintVouchers(context.Context, *MsgMintVouchers) (*MsgMintVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintVouchers not implemented")
}
func (UnimplementedMsgServer) BurnVouchers(context.Context, *MsgBurnVouchers) (*MsgBurnVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnVouchers not implemented")
}
func (UnimplementedMsgServer) RedeemVouchers(context.Context, *MsgRedeemVouchers) (*MsgRedeemVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemVouchers not implemented")
}
func (UnimplementedMsgServer) UnredeemVouchers(context.Context, *MsgUnredeemVouchers) (*MsgUnredeemVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnredeemVouchers not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProject(ctx, req.(*MsgCreateProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditProject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditProject(ctx, req.(*MsgEditProject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateTotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateTotalSupply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateTotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateTotalSupply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateTotalSupply(ctx, req.(*MsgUpdateTotalSupply))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSpecialAllocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateSpecialAllocations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSpecialAllocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateSpecialAllocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSpecialAllocations(ctx, req.(*MsgUpdateSpecialAllocations))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InitializeMainnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitializeMainnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitializeMainnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InitializeMainnet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitializeMainnet(ctx, req.(*MsgInitializeMainnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintVouchers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintVouchers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintVouchers(ctx, req.(*MsgMintVouchers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnVouchers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnVouchers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnVouchers(ctx, req.(*MsgBurnVouchers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RedeemVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRedeemVouchers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RedeemVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RedeemVouchers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RedeemVouchers(ctx, req.(*MsgRedeemVouchers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnredeemVouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnredeemVouchers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnredeemVouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UnredeemVouchers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnredeemVouchers(ctx, req.(*MsgUnredeemVouchers))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network.project.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Msg_CreateProject_Handler,
		},
		{
			MethodName: "EditProject",
			Handler:    _Msg_EditProject_Handler,
		},
		{
			MethodName: "UpdateTotalSupply",
			Handler:    _Msg_UpdateTotalSupply_Handler,
		},
		{
			MethodName: "UpdateSpecialAllocations",
			Handler:    _Msg_UpdateSpecialAllocations_Handler,
		},
		{
			MethodName: "InitializeMainnet",
			Handler:    _Msg_InitializeMainnet_Handler,
		},
		{
			MethodName: "MintVouchers",
			Handler:    _Msg_MintVouchers_Handler,
		},
		{
			MethodName: "BurnVouchers",
			Handler:    _Msg_BurnVouchers_Handler,
		},
		{
			MethodName: "RedeemVouchers",
			Handler:    _Msg_RedeemVouchers_Handler,
		},
		{
			MethodName: "UnredeemVouchers",
			Handler:    _Msg_UnredeemVouchers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network/project/v1/tx.proto",
}
