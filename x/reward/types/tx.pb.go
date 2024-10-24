// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/reward/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3137090ae719fd02, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3137090ae719fd02, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgSetRewards struct {
	Provider         string                                   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	LaunchId         uint64                                   `protobuf:"varint,2,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	Coins            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	LastRewardHeight int64                                    `protobuf:"varint,4,opt,name=last_reward_height,json=lastRewardHeight,proto3" json:"last_reward_height,omitempty"`
}

func (m *MsgSetRewards) Reset()         { *m = MsgSetRewards{} }
func (m *MsgSetRewards) String() string { return proto.CompactTextString(m) }
func (*MsgSetRewards) ProtoMessage()    {}
func (*MsgSetRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_3137090ae719fd02, []int{2}
}
func (m *MsgSetRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetRewards.Merge(m, src)
}
func (m *MsgSetRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetRewards proto.InternalMessageInfo

func (m *MsgSetRewards) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *MsgSetRewards) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

func (m *MsgSetRewards) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *MsgSetRewards) GetLastRewardHeight() int64 {
	if m != nil {
		return m.LastRewardHeight
	}
	return 0
}

type MsgSetRewardsResponse struct {
	PreviousCoins            github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=previous_coins,json=previousCoins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"previous_coins"`
	PreviousLastRewardHeight int64                                    `protobuf:"varint,2,opt,name=previous_last_reward_height,json=previousLastRewardHeight,proto3" json:"previous_last_reward_height,omitempty"`
	NewCoins                 github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=new_coins,json=newCoins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"new_coins"`
	NewLastRewardHeight      int64                                    `protobuf:"varint,4,opt,name=new_last_reward_height,json=newLastRewardHeight,proto3" json:"new_last_reward_height,omitempty"`
}

func (m *MsgSetRewardsResponse) Reset()         { *m = MsgSetRewardsResponse{} }
func (m *MsgSetRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetRewardsResponse) ProtoMessage()    {}
func (*MsgSetRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3137090ae719fd02, []int{3}
}
func (m *MsgSetRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetRewardsResponse.Merge(m, src)
}
func (m *MsgSetRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetRewardsResponse proto.InternalMessageInfo

func (m *MsgSetRewardsResponse) GetPreviousCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PreviousCoins
	}
	return nil
}

func (m *MsgSetRewardsResponse) GetPreviousLastRewardHeight() int64 {
	if m != nil {
		return m.PreviousLastRewardHeight
	}
	return 0
}

func (m *MsgSetRewardsResponse) GetNewCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.NewCoins
	}
	return nil
}

func (m *MsgSetRewardsResponse) GetNewLastRewardHeight() int64 {
	if m != nil {
		return m.NewLastRewardHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "network.reward.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "network.reward.v1.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgSetRewards)(nil), "network.reward.v1.MsgSetRewards")
	proto.RegisterType((*MsgSetRewardsResponse)(nil), "network.reward.v1.MsgSetRewardsResponse")
}

func init() { proto.RegisterFile("network/reward/v1/tx.proto", fileDescriptor_3137090ae719fd02) }

var fileDescriptor_3137090ae719fd02 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xce, 0x35, 0x6d, 0xd5, 0xdc, 0xaf, 0xfd, 0x41, 0x8f, 0x42, 0x5d, 0x57, 0x72, 0xad, 0x2c,
	0x44, 0x11, 0xb5, 0x95, 0x16, 0x31, 0x54, 0x30, 0x10, 0x16, 0x90, 0xa8, 0x84, 0x5c, 0x21, 0x21,
	0x06, 0x2c, 0x27, 0x3e, 0x5d, 0x4e, 0xad, 0x7d, 0xc6, 0x77, 0x49, 0xda, 0x0d, 0x31, 0x32, 0x31,
	0x33, 0x32, 0x01, 0x53, 0x24, 0x18, 0x19, 0x11, 0xea, 0x58, 0x31, 0x31, 0x15, 0xd4, 0x0e, 0x99,
	0xf8, 0x07, 0x98, 0x90, 0xef, 0xce, 0x69, 0x9b, 0x06, 0x85, 0xb1, 0x4b, 0x92, 0x7b, 0xdf, 0xf7,
	0xde, 0xfb, 0xde, 0xf7, 0x72, 0x07, 0xcd, 0x18, 0x8b, 0x2e, 0x4b, 0xb7, 0xdd, 0x14, 0x77, 0x83,
	0x34, 0x74, 0x3b, 0x35, 0x57, 0xec, 0x3a, 0x49, 0xca, 0x04, 0x43, 0xf3, 0x1a, 0x73, 0x14, 0xe6,
	0x74, 0x6a, 0xe6, 0x7c, 0x10, 0xd1, 0x98, 0xb9, 0xf2, 0x53, 0xb1, 0x4c, 0xab, 0xc9, 0x78, 0xc4,
	0xb8, 0xdb, 0x08, 0x38, 0x76, 0x3b, 0xb5, 0x06, 0x16, 0x41, 0xcd, 0x6d, 0x32, 0x1a, 0x6b, 0x7c,
	0x51, 0xe3, 0x11, 0x27, 0x59, 0xf5, 0x88, 0x13, 0x0d, 0x2c, 0x29, 0xc0, 0x97, 0x27, 0x57, 0x1d,
	0x34, 0xb4, 0x40, 0x18, 0x61, 0x2a, 0x9e, 0xfd, 0xca, 0x3b, 0x9d, 0xd7, 0x9a, 0x04, 0x69, 0x10,
	0xe9, 0xac, 0xf2, 0x67, 0x00, 0x2f, 0x6d, 0x72, 0xf2, 0x38, 0x09, 0x03, 0x81, 0x1f, 0x49, 0x04,
	0xdd, 0x82, 0xa5, 0xa0, 0x2d, 0x5a, 0x2c, 0xa5, 0x62, 0xcf, 0x00, 0x36, 0xa8, 0x94, 0xea, 0xc6,
	0xb7, 0x4f, 0xab, 0x0b, 0xba, 0xdd, 0xdd, 0x30, 0x4c, 0x31, 0xe7, 0x5b, 0x22, 0xa5, 0x31, 0xf1,
	0x4e, 0xa8, 0xe8, 0x36, 0x9c, 0x56, 0xb5, 0x8d, 0x09, 0x1b, 0x54, 0xfe, 0x5b, 0x5b, 0x72, 0xce,
	0x99, 0xe1, 0xa8, 0x16, 0xf5, 0xd2, 0xfe, 0xe1, 0x4a, 0xe1, 0x5d, 0xbf, 0x57, 0x05, 0x9e, 0xce,
	0xd9, 0x58, 0x7f, 0xd9, 0xef, 0x55, 0x4f, 0xaa, 0xbd, 0xea, 0xf7, 0xaa, 0x76, 0x2e, 0x7e, 0x37,
	0x97, 0x3f, 0x24, 0xb5, 0xbc, 0x04, 0x17, 0x87, 0x42, 0x1e, 0xe6, 0x09, 0x8b, 0x39, 0x2e, 0x7f,
	0x99, 0x80, 0x73, 0x9b, 0x9c, 0x6c, 0x61, 0xe1, 0xc9, 0x5c, 0x8e, 0x6e, 0xc2, 0x99, 0x24, 0x65,
	0x1d, 0x1a, 0xe2, 0x74, 0xec, 0x58, 0x03, 0x26, 0x5a, 0x86, 0xa5, 0x9d, 0xa0, 0x1d, 0x37, 0x5b,
	0x3e, 0x0d, 0xe5, 0x60, 0x93, 0xde, 0x8c, 0x0a, 0x3c, 0x08, 0xd1, 0x5b, 0x00, 0xa7, 0xb2, 0xbd,
	0x71, 0xa3, 0x68, 0x17, 0xe5, 0xc8, 0xba, 0x5a, 0xb6, 0x59, 0x47, 0x6f, 0xd6, 0xb9, 0xc7, 0x68,
	0x5c, 0x7f, 0x9e, 0x8d, 0xfc, 0xfb, 0x70, 0xe5, 0x3a, 0xa1, 0xa2, 0xd5, 0x6e, 0x38, 0x4d, 0x16,
	0xe9, 0x05, 0xea, 0xaf, 0x55, 0x1e, 0x6e, 0xbb, 0x62, 0x2f, 0xc1, 0x5c, 0x26, 0x7c, 0xf8, 0xb1,
	0x52, 0xf9, 0x47, 0x2a, 0x7f, 0xd3, 0xef, 0x55, 0x67, 0x77, 0x30, 0x09, 0x9a, 0x7b, 0xbe, 0x94,
	0xa3, 0xac, 0x55, 0xd2, 0xd0, 0x0d, 0x88, 0x76, 0x02, 0x2e, 0x7c, 0xe5, 0xa1, 0xdf, 0xc2, 0x94,
	0xb4, 0x84, 0x31, 0x69, 0x83, 0x4a, 0xd1, 0xbb, 0x9c, 0x21, 0xca, 0xa0, 0xfb, 0x32, 0xbe, 0x31,
	0x97, 0xed, 0x61, 0x30, 0x7e, 0xf9, 0x57, 0x11, 0x5e, 0x3d, 0x63, 0x63, 0x6e, 0x30, 0xfa, 0x08,
	0xe0, 0xff, 0x49, 0x8a, 0x3b, 0x94, 0xb5, 0xb9, 0xea, 0x6a, 0x80, 0x0b, 0x68, 0xc2, 0x5c, 0xae,
	0x51, 0x92, 0xd0, 0x1d, 0xb8, 0x3c, 0x10, 0x3d, 0xc2, 0x95, 0x09, 0xe9, 0x8a, 0x91, 0x53, 0x1e,
	0x0e, 0xb9, 0x83, 0xde, 0x03, 0x58, 0x8a, 0x71, 0xd7, 0xbf, 0xb8, 0x4b, 0x9f, 0x89, 0x71, 0x57,
	0x8d, 0xba, 0x0e, 0xaf, 0x65, 0x52, 0xff, 0xba, 0xfb, 0x2b, 0x31, 0xee, 0x0e, 0x0f, 0xb8, 0xf6,
	0x15, 0xc0, 0xe2, 0x26, 0x27, 0xe8, 0x19, 0x9c, 0x3d, 0xf3, 0x28, 0x94, 0x47, 0x5c, 0xe6, 0xa1,
	0xab, 0x67, 0x56, 0xc7, 0x73, 0x06, 0xff, 0x9e, 0x27, 0x10, 0x9e, 0xba, 0x9a, 0xf6, 0xe8, 0xcc,
	0x13, 0x86, 0x59, 0x19, 0xc7, 0xc8, 0x2b, 0x9b, 0x53, 0x2f, 0x32, 0x1f, 0xea, 0xf5, 0xfd, 0x23,
	0x0b, 0x1c, 0x1c, 0x59, 0xe0, 0xe7, 0x91, 0x05, 0x5e, 0x1f, 0x5b, 0x85, 0x83, 0x63, 0xab, 0xf0,
	0xfd, 0xd8, 0x2a, 0x3c, 0x3d, 0xed, 0x30, 0x25, 0x31, 0x15, 0xd8, 0x3d, 0xf7, 0xd0, 0x48, 0x9f,
	0x1b, 0xd3, 0xf2, 0x91, 0x5c, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xf6, 0x1b, 0x57, 0xf2,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	SetRewards(ctx context.Context, in *MsgSetRewards, opts ...grpc.CallOption) (*MsgSetRewardsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/network.reward.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetRewards(ctx context.Context, in *MsgSetRewards, opts ...grpc.CallOption) (*MsgSetRewardsResponse, error) {
	out := new(MsgSetRewardsResponse)
	err := c.cc.Invoke(ctx, "/network.reward.v1.Msg/SetRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	SetRewards(context.Context, *MsgSetRewards) (*MsgSetRewardsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) SetRewards(ctx context.Context, req *MsgSetRewards) (*MsgSetRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRewards not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
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
		FullMethod: "/network.reward.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/network.reward.v1.Msg/SetRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetRewards(ctx, req.(*MsgSetRewards))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "network.reward.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SetRewards",
			Handler:    _Msg_SetRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network/reward/v1/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastRewardHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.LastRewardHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.LaunchId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewLastRewardHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NewLastRewardHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NewCoins) > 0 {
		for iNdEx := len(m.NewCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NewCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PreviousLastRewardHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PreviousLastRewardHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PreviousCoins) > 0 {
		for iNdEx := len(m.PreviousCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PreviousCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.LaunchId != 0 {
		n += 1 + sovTx(uint64(m.LaunchId))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.LastRewardHeight != 0 {
		n += 1 + sovTx(uint64(m.LastRewardHeight))
	}
	return n
}

func (m *MsgSetRewardsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PreviousCoins) > 0 {
		for _, e := range m.PreviousCoins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.PreviousLastRewardHeight != 0 {
		n += 1 + sovTx(uint64(m.PreviousLastRewardHeight))
	}
	if len(m.NewCoins) > 0 {
		for _, e := range m.NewCoins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.NewLastRewardHeight != 0 {
		n += 1 + sovTx(uint64(m.NewLastRewardHeight))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRewardHeight", wireType)
			}
			m.LastRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastRewardHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetRewardsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousCoins = append(m.PreviousCoins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.PreviousCoins[len(m.PreviousCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousLastRewardHeight", wireType)
			}
			m.PreviousLastRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreviousLastRewardHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewCoins = append(m.NewCoins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.NewCoins[len(m.NewCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewLastRewardHeight", wireType)
			}
			m.NewLastRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewLastRewardHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
