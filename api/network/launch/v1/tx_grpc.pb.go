// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: network/launch/v1/tx.proto

package launchv1

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
	Msg_UpdateParams_FullMethodName            = "/network.launch.v1.Msg/UpdateParams"
	Msg_CreateChain_FullMethodName             = "/network.launch.v1.Msg/CreateChain"
	Msg_EditChain_FullMethodName               = "/network.launch.v1.Msg/EditChain"
	Msg_UpdateLaunchInformation_FullMethodName = "/network.launch.v1.Msg/UpdateLaunchInformation"
	Msg_SendRequest_FullMethodName             = "/network.launch.v1.Msg/SendRequest"
	Msg_SettleRequest_FullMethodName           = "/network.launch.v1.Msg/SettleRequest"
	Msg_TriggerLaunch_FullMethodName           = "/network.launch.v1.Msg/TriggerLaunch"
	Msg_RevertLaunch_FullMethodName            = "/network.launch.v1.Msg/RevertLaunch"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateChain(ctx context.Context, in *MsgCreateChain, opts ...grpc.CallOption) (*MsgCreateChainResponse, error)
	EditChain(ctx context.Context, in *MsgEditChain, opts ...grpc.CallOption) (*MsgEditChainResponse, error)
	UpdateLaunchInformation(ctx context.Context, in *MsgUpdateLaunchInformation, opts ...grpc.CallOption) (*MsgUpdateLaunchInformationResponse, error)
	SendRequest(ctx context.Context, in *MsgSendRequest, opts ...grpc.CallOption) (*MsgSendRequestResponse, error)
	SettleRequest(ctx context.Context, in *MsgSettleRequest, opts ...grpc.CallOption) (*MsgSettleRequestResponse, error)
	TriggerLaunch(ctx context.Context, in *MsgTriggerLaunch, opts ...grpc.CallOption) (*MsgTriggerLaunchResponse, error)
	RevertLaunch(ctx context.Context, in *MsgRevertLaunch, opts ...grpc.CallOption) (*MsgRevertLaunchResponse, error)
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

func (c *msgClient) CreateChain(ctx context.Context, in *MsgCreateChain, opts ...grpc.CallOption) (*MsgCreateChainResponse, error) {
	out := new(MsgCreateChainResponse)
	err := c.cc.Invoke(ctx, Msg_CreateChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditChain(ctx context.Context, in *MsgEditChain, opts ...grpc.CallOption) (*MsgEditChainResponse, error) {
	out := new(MsgEditChainResponse)
	err := c.cc.Invoke(ctx, Msg_EditChain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateLaunchInformation(ctx context.Context, in *MsgUpdateLaunchInformation, opts ...grpc.CallOption) (*MsgUpdateLaunchInformationResponse, error) {
	out := new(MsgUpdateLaunchInformationResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateLaunchInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendRequest(ctx context.Context, in *MsgSendRequest, opts ...grpc.CallOption) (*MsgSendRequestResponse, error) {
	out := new(MsgSendRequestResponse)
	err := c.cc.Invoke(ctx, Msg_SendRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SettleRequest(ctx context.Context, in *MsgSettleRequest, opts ...grpc.CallOption) (*MsgSettleRequestResponse, error) {
	out := new(MsgSettleRequestResponse)
	err := c.cc.Invoke(ctx, Msg_SettleRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TriggerLaunch(ctx context.Context, in *MsgTriggerLaunch, opts ...grpc.CallOption) (*MsgTriggerLaunchResponse, error) {
	out := new(MsgTriggerLaunchResponse)
	err := c.cc.Invoke(ctx, Msg_TriggerLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevertLaunch(ctx context.Context, in *MsgRevertLaunch, opts ...grpc.CallOption) (*MsgRevertLaunchResponse, error) {
	out := new(MsgRevertLaunchResponse)
	err := c.cc.Invoke(ctx, Msg_RevertLaunch_FullMethodName, in, out, opts...)
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
	CreateChain(context.Context, *MsgCreateChain) (*MsgCreateChainResponse, error)
	EditChain(context.Context, *MsgEditChain) (*MsgEditChainResponse, error)
	UpdateLaunchInformation(context.Context, *MsgUpdateLaunchInformation) (*MsgUpdateLaunchInformationResponse, error)
	SendRequest(context.Context, *MsgSendRequest) (*MsgSendRequestResponse, error)
	SettleRequest(context.Context, *MsgSettleRequest) (*MsgSettleRequestResponse, error)
	TriggerLaunch(context.Context, *MsgTriggerLaunch) (*MsgTriggerLaunchResponse, error)
	RevertLaunch(context.Context, *MsgRevertLaunch) (*MsgRevertLaunchResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateChain(context.Context, *MsgCreateChain) (*MsgCreateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChain not implemented")
}
func (UnimplementedMsgServer) EditChain(context.Context, *MsgEditChain) (*MsgEditChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditChain not implemented")
}
func (UnimplementedMsgServer) UpdateLaunchInformation(context.Context, *MsgUpdateLaunchInformation) (*MsgUpdateLaunchInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLaunchInformation not implemented")
}
func (UnimplementedMsgServer) SendRequest(context.Context, *MsgSendRequest) (*MsgSendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedMsgServer) SettleRequest(context.Context, *MsgSettleRequest) (*MsgSettleRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SettleRequest not implemented")
}
func (UnimplementedMsgServer) TriggerLaunch(context.Context, *MsgTriggerLaunch) (*MsgTriggerLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerLaunch not implemented")
}
func (UnimplementedMsgServer) RevertLaunch(context.Context, *MsgRevertLaunch) (*MsgRevertLaunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevertLaunch not implemented")
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

func _Msg_CreateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateChain(ctx, req.(*MsgCreateChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditChain(ctx, req.(*MsgEditChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateLaunchInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateLaunchInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateLaunchInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateLaunchInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateLaunchInformation(ctx, req.(*MsgUpdateLaunchInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SendRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendRequest(ctx, req.(*MsgSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SettleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSettleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SettleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SettleRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SettleRequest(ctx, req.(*MsgSettleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TriggerLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTriggerLaunch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TriggerLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TriggerLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TriggerLaunch(ctx, req.(*MsgTriggerLaunch))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevertLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevertLaunch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevertLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RevertLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevertLaunch(ctx, req.(*MsgRevertLaunch))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "network.launch.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateChain",
			Handler:    _Msg_CreateChain_Handler,
		},
		{
			MethodName: "EditChain",
			Handler:    _Msg_EditChain_Handler,
		},
		{
			MethodName: "UpdateLaunchInformation",
			Handler:    _Msg_UpdateLaunchInformation_Handler,
		},
		{
			MethodName: "SendRequest",
			Handler:    _Msg_SendRequest_Handler,
		},
		{
			MethodName: "SettleRequest",
			Handler:    _Msg_SettleRequest_Handler,
		},
		{
			MethodName: "TriggerLaunch",
			Handler:    _Msg_TriggerLaunch_Handler,
		},
		{
			MethodName: "RevertLaunch",
			Handler:    _Msg_RevertLaunch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network/launch/v1/tx.proto",
}
