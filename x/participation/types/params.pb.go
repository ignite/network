// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/participation/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	AllocationPrice       AllocationPrice `protobuf:"bytes,1,opt,name=allocation_price,json=allocationPrice,proto3" json:"allocation_price"`
	ParticipationTierList []Tier          `protobuf:"bytes,2,rep,name=participation_tier_list,json=participationTierList,proto3" json:"participation_tier_list"`
	// Time frame before auction starts where MsgParticipate can be called
	RegistrationPeriod time.Duration `protobuf:"bytes,3,opt,name=registration_period,json=registrationPeriod,proto3,stdduration" json:"registration_period"`
	// Delay after auction starts when allocations can be withdrawn
	WithdrawalDelay time.Duration `protobuf:"bytes,4,opt,name=withdrawal_delay,json=withdrawalDelay,proto3,stdduration" json:"withdrawal_delay"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b6c03ddd407c68b, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAllocationPrice() AllocationPrice {
	if m != nil {
		return m.AllocationPrice
	}
	return AllocationPrice{}
}

func (m *Params) GetParticipationTierList() []Tier {
	if m != nil {
		return m.ParticipationTierList
	}
	return nil
}

func (m *Params) GetRegistrationPeriod() time.Duration {
	if m != nil {
		return m.RegistrationPeriod
	}
	return 0
}

func (m *Params) GetWithdrawalDelay() time.Duration {
	if m != nil {
		return m.WithdrawalDelay
	}
	return 0
}

type AllocationPrice struct {
	// number of bonded tokens necessary to get one allocation
	Bonded cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=bonded,proto3,customtype=cosmossdk.io/math.Int" json:"bonded"`
}

func (m *AllocationPrice) Reset()         { *m = AllocationPrice{} }
func (m *AllocationPrice) String() string { return proto.CompactTextString(m) }
func (*AllocationPrice) ProtoMessage()    {}
func (*AllocationPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b6c03ddd407c68b, []int{1}
}
func (m *AllocationPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllocationPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllocationPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllocationPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationPrice.Merge(m, src)
}
func (m *AllocationPrice) XXX_Size() int {
	return m.Size()
}
func (m *AllocationPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationPrice.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationPrice proto.InternalMessageInfo

// Matches a number of required allocations with benefits
type Tier struct {
	TierId              uint64                `protobuf:"varint,1,opt,name=tier_id,json=tierId,proto3" json:"tier_id,omitempty"`
	RequiredAllocations cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=required_allocations,json=requiredAllocations,proto3,customtype=cosmossdk.io/math.Int" json:"required_allocations"`
	Benefits            TierBenefits          `protobuf:"bytes,3,opt,name=benefits,proto3" json:"benefits"`
}

func (m *Tier) Reset()         { *m = Tier{} }
func (m *Tier) String() string { return proto.CompactTextString(m) }
func (*Tier) ProtoMessage()    {}
func (*Tier) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b6c03ddd407c68b, []int{2}
}
func (m *Tier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tier.Merge(m, src)
}
func (m *Tier) XXX_Size() int {
	return m.Size()
}
func (m *Tier) XXX_DiscardUnknown() {
	xxx_messageInfo_Tier.DiscardUnknown(m)
}

var xxx_messageInfo_Tier proto.InternalMessageInfo

func (m *Tier) GetTierId() uint64 {
	if m != nil {
		return m.TierId
	}
	return 0
}

func (m *Tier) GetBenefits() TierBenefits {
	if m != nil {
		return m.Benefits
	}
	return TierBenefits{}
}

type TierBenefits struct {
	// max_bid_amount maximum amount an auction participant can bid
	MaxBidAmount cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=max_bid_amount,json=maxBidAmount,proto3,customtype=cosmossdk.io/math.Int" json:"max_bid_amount"`
}

func (m *TierBenefits) Reset()         { *m = TierBenefits{} }
func (m *TierBenefits) String() string { return proto.CompactTextString(m) }
func (*TierBenefits) ProtoMessage()    {}
func (*TierBenefits) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b6c03ddd407c68b, []int{3}
}
func (m *TierBenefits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TierBenefits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TierBenefits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TierBenefits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TierBenefits.Merge(m, src)
}
func (m *TierBenefits) XXX_Size() int {
	return m.Size()
}
func (m *TierBenefits) XXX_DiscardUnknown() {
	xxx_messageInfo_TierBenefits.DiscardUnknown(m)
}

var xxx_messageInfo_TierBenefits proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "network.participation.v1.Params")
	proto.RegisterType((*AllocationPrice)(nil), "network.participation.v1.AllocationPrice")
	proto.RegisterType((*Tier)(nil), "network.participation.v1.Tier")
	proto.RegisterType((*TierBenefits)(nil), "network.participation.v1.TierBenefits")
}

func init() {
	proto.RegisterFile("network/participation/v1/params.proto", fileDescriptor_9b6c03ddd407c68b)
}

var fileDescriptor_9b6c03ddd407c68b = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4f, 0x13, 0x41,
	0x18, 0xee, 0xd2, 0xa6, 0xe2, 0x40, 0x2c, 0x0e, 0x10, 0x16, 0x0e, 0x5b, 0xd2, 0x44, 0x45, 0x8d,
	0xb3, 0x01, 0x6f, 0xde, 0xa8, 0x1c, 0x6c, 0x62, 0x4c, 0x6d, 0x88, 0x07, 0x62, 0xdc, 0xcc, 0x76,
	0x86, 0xed, 0x1b, 0x76, 0x77, 0xea, 0xcc, 0x94, 0x96, 0xbf, 0xe0, 0x49, 0x6f, 0x1e, 0xfd, 0x09,
	0x1e, 0xfc, 0x11, 0x78, 0x23, 0x9e, 0x8c, 0x07, 0x34, 0xed, 0x41, 0x7f, 0x86, 0xd9, 0xd9, 0x29,
	0xa5, 0x24, 0x8d, 0x1f, 0x97, 0xcd, 0xbe, 0x1f, 0xcf, 0xf3, 0xbe, 0xf3, 0x3c, 0x2f, 0xba, 0x95,
	0x72, 0xdd, 0x17, 0xf2, 0xc8, 0xef, 0x52, 0xa9, 0xa1, 0x0d, 0x5d, 0xaa, 0x41, 0xa4, 0xfe, 0xf1,
	0x76, 0x96, 0xa0, 0x89, 0x22, 0x5d, 0x29, 0xb4, 0xc0, 0xae, 0x6d, 0x23, 0x53, 0x6d, 0xe4, 0x78,
	0x7b, 0xe3, 0x26, 0x4d, 0x20, 0x15, 0xbe, 0xf9, 0xe6, 0xcd, 0x1b, 0xeb, 0x6d, 0xa1, 0x12, 0xa1,
	0x02, 0x13, 0xf9, 0x79, 0x60, 0x4b, 0x2b, 0x91, 0x88, 0x44, 0x9e, 0xcf, 0xfe, 0x6c, 0xd6, 0x8b,
	0x84, 0x88, 0x62, 0xee, 0x9b, 0x28, 0xec, 0x1d, 0xfa, 0xac, 0x27, 0xf3, 0x01, 0x26, 0x53, 0x7b,
	0x57, 0x44, 0xe5, 0xa6, 0x59, 0x07, 0x1f, 0xa0, 0x25, 0x1a, 0xc7, 0xa2, 0x6d, 0xca, 0x41, 0x57,
	0x42, 0x9b, 0xbb, 0xce, 0xa6, 0xb3, 0xb5, 0xb0, 0x73, 0x97, 0xcc, 0xda, 0x91, 0xec, 0x5e, 0x20,
	0x9a, 0x19, 0xa0, 0x5e, 0x3a, 0x3d, 0xaf, 0x16, 0x5a, 0x15, 0x3a, 0x9d, 0xc6, 0x2f, 0xd1, 0xda,
	0x14, 0x34, 0xd0, 0xc0, 0x65, 0x10, 0x83, 0xd2, 0xee, 0xdc, 0x66, 0x71, 0x6b, 0x61, 0xc7, 0x9b,
	0x3d, 0x62, 0x1f, 0xb8, 0xb4, 0xbc, 0xab, 0x53, 0xc5, 0xac, 0xf0, 0x14, 0x94, 0xc6, 0xfb, 0x68,
	0x59, 0xf2, 0x08, 0x94, 0x96, 0x76, 0x77, 0x2e, 0x41, 0x30, 0xb7, 0x68, 0x96, 0x5f, 0x27, 0xb9,
	0x04, 0x64, 0x2c, 0x01, 0xd9, 0xb3, 0x12, 0xd4, 0xe7, 0x33, 0xd2, 0xf7, 0xdf, 0xab, 0x4e, 0x0b,
	0x5f, 0xc6, 0x37, 0x0d, 0x1c, 0x3f, 0x43, 0x4b, 0x7d, 0xd0, 0x1d, 0x26, 0x69, 0x9f, 0xc6, 0x01,
	0xe3, 0x31, 0x3d, 0x71, 0x4b, 0x7f, 0x4f, 0x59, 0x99, 0x80, 0xf7, 0x32, 0xec, 0xa3, 0x3b, 0xbf,
	0x3e, 0x54, 0x9d, 0x37, 0x3f, 0x3f, 0xde, 0xf3, 0xc6, 0x87, 0x31, 0xb8, 0x72, 0x1a, 0xb9, 0x11,
	0xb5, 0x17, 0xa8, 0x72, 0x45, 0x56, 0xfc, 0x18, 0x95, 0x43, 0x91, 0x32, 0xce, 0x8c, 0x23, 0xd7,
	0xeb, 0xf7, 0xb3, 0x31, 0xdf, 0xce, 0xab, 0xab, 0xf9, 0x09, 0x28, 0x76, 0x44, 0x40, 0xf8, 0x09,
	0xd5, 0x1d, 0xd2, 0x48, 0xf5, 0x97, 0x4f, 0x0f, 0x90, 0xbd, 0x8d, 0x46, 0xaa, 0x5b, 0x16, 0x5a,
	0xfb, 0xec, 0xa0, 0x52, 0xa6, 0x19, 0x5e, 0x43, 0xd7, 0x8c, 0xfe, 0x90, 0xd3, 0x95, 0x5a, 0xe5,
	0x2c, 0x6c, 0x30, 0xfc, 0x0a, 0xad, 0x48, 0xfe, 0xba, 0x07, 0x92, 0xb3, 0x60, 0x62, 0xa1, 0x72,
	0xe7, 0xfe, 0x7d, 0xe8, 0xf2, 0x98, 0x68, 0xf2, 0x14, 0x85, 0x9f, 0xa0, 0xf9, 0x90, 0xa7, 0xfc,
	0x10, 0xb4, 0xb2, 0xee, 0xdc, 0xfe, 0x83, 0xef, 0xb6, 0xdb, 0xfa, 0x7f, 0x81, 0xae, 0x51, 0xb4,
	0x78, 0xb9, 0x8e, 0x9f, 0xa3, 0x1b, 0x09, 0x1d, 0x04, 0x21, 0xb0, 0x80, 0x26, 0xa2, 0x97, 0xea,
	0xff, 0x11, 0x6a, 0x31, 0xa1, 0x83, 0x3a, 0xb0, 0x5d, 0x43, 0x50, 0x6f, 0x9c, 0x0e, 0x3d, 0xe7,
	0x6c, 0xe8, 0x39, 0x3f, 0x86, 0x9e, 0xf3, 0x76, 0xe4, 0x15, 0xce, 0x46, 0x5e, 0xe1, 0xeb, 0xc8,
	0x2b, 0x1c, 0xf8, 0x11, 0xe8, 0x4e, 0x2f, 0x24, 0x6d, 0x91, 0xf8, 0x10, 0xa5, 0xa0, 0xb9, 0x3f,
	0xcb, 0x52, 0x7d, 0xd2, 0xe5, 0x2a, 0x2c, 0x9b, 0x43, 0x79, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x0e, 0x89, 0x43, 0x13, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AllocationPrice.Equal(&that1.AllocationPrice) {
		return false
	}
	if len(this.ParticipationTierList) != len(that1.ParticipationTierList) {
		return false
	}
	for i := range this.ParticipationTierList {
		if !this.ParticipationTierList[i].Equal(&that1.ParticipationTierList[i]) {
			return false
		}
	}
	if this.RegistrationPeriod != that1.RegistrationPeriod {
		return false
	}
	if this.WithdrawalDelay != that1.WithdrawalDelay {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.WithdrawalDelay, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.WithdrawalDelay):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.RegistrationPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.RegistrationPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.ParticipationTierList) > 0 {
		for iNdEx := len(m.ParticipationTierList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ParticipationTierList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.AllocationPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AllocationPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllocationPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllocationPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Bonded.Size()
		i -= size
		if _, err := m.Bonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Tier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Benefits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.RequiredAllocations.Size()
		i -= size
		if _, err := m.RequiredAllocations.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.TierId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TierId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TierBenefits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TierBenefits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TierBenefits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxBidAmount.Size()
		i -= size
		if _, err := m.MaxBidAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AllocationPrice.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.ParticipationTierList) > 0 {
		for _, e := range m.ParticipationTierList {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.RegistrationPeriod)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.WithdrawalDelay)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *AllocationPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Bonded.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Tier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TierId != 0 {
		n += 1 + sovParams(uint64(m.TierId))
	}
	l = m.RequiredAllocations.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Benefits.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *TierBenefits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxBidAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocationPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AllocationPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationTierList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParticipationTierList = append(m.ParticipationTierList, Tier{})
			if err := m.ParticipationTierList[len(m.ParticipationTierList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistrationPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.RegistrationPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.WithdrawalDelay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *AllocationPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: AllocationPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllocationPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Tier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Tier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TierId", wireType)
			}
			m.TierId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TierId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequiredAllocations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RequiredAllocations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Benefits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Benefits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *TierBenefits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: TierBenefits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TierBenefits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBidAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
