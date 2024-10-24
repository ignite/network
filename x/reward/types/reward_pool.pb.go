// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/reward/v1/reward_pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type RewardPool struct {
	LaunchId            uint64                                   `protobuf:"varint,1,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	Provider            string                                   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	InitialCoins        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_coins,json=initialCoins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_coins"`
	RemainingCoins      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=remaining_coins,json=remainingCoins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"remaining_coins"`
	LastRewardHeight    int64                                    `protobuf:"varint,5,opt,name=last_reward_height,json=lastRewardHeight,proto3" json:"last_reward_height,omitempty"`
	CurrentRewardHeight int64                                    `protobuf:"varint,6,opt,name=current_reward_height,json=currentRewardHeight,proto3" json:"current_reward_height,omitempty"`
	Closed              bool                                     `protobuf:"varint,7,opt,name=closed,proto3" json:"closed,omitempty"`
}

func (m *RewardPool) Reset()         { *m = RewardPool{} }
func (m *RewardPool) String() string { return proto.CompactTextString(m) }
func (*RewardPool) ProtoMessage()    {}
func (*RewardPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec0cc30cd2290ef1, []int{0}
}
func (m *RewardPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardPool.Merge(m, src)
}
func (m *RewardPool) XXX_Size() int {
	return m.Size()
}
func (m *RewardPool) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardPool.DiscardUnknown(m)
}

var xxx_messageInfo_RewardPool proto.InternalMessageInfo

func (m *RewardPool) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

func (m *RewardPool) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *RewardPool) GetInitialCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialCoins
	}
	return nil
}

func (m *RewardPool) GetRemainingCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.RemainingCoins
	}
	return nil
}

func (m *RewardPool) GetLastRewardHeight() int64 {
	if m != nil {
		return m.LastRewardHeight
	}
	return 0
}

func (m *RewardPool) GetCurrentRewardHeight() int64 {
	if m != nil {
		return m.CurrentRewardHeight
	}
	return 0
}

func (m *RewardPool) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func init() {
	proto.RegisterType((*RewardPool)(nil), "network.reward.v1.RewardPool")
}

func init() {
	proto.RegisterFile("network/reward/v1/reward_pool.proto", fileDescriptor_ec0cc30cd2290ef1)
}

var fileDescriptor_ec0cc30cd2290ef1 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xbf, 0x6e, 0x13, 0x31,
	0x1c, 0x8e, 0x49, 0x08, 0xa9, 0x29, 0x7f, 0x7a, 0x14, 0x74, 0x2d, 0xd2, 0xe5, 0x04, 0x03, 0xa7,
	0x8a, 0x9e, 0x95, 0xc2, 0x0b, 0x10, 0x16, 0xd8, 0xd0, 0xb1, 0xb1, 0x9c, 0x9c, 0xb3, 0xe5, 0x58,
	0xbd, 0xf8, 0x17, 0x6c, 0xe7, 0x4a, 0xdf, 0x82, 0x99, 0x27, 0x40, 0x4c, 0x45, 0x82, 0x47, 0x40,
	0xea, 0x58, 0x31, 0x31, 0x15, 0x94, 0x0c, 0x7d, 0x07, 0x26, 0x74, 0xb6, 0x5b, 0x01, 0x13, 0x23,
	0xcb, 0x9d, 0x7f, 0xdf, 0x1f, 0xf9, 0xfb, 0x64, 0x1b, 0xdf, 0x57, 0xdc, 0x1e, 0x80, 0xde, 0x27,
	0x9a, 0x1f, 0x50, 0xcd, 0x48, 0x33, 0x0a, 0xab, 0x72, 0x0e, 0x50, 0xe7, 0x73, 0x0d, 0x16, 0xa2,
	0x8d, 0x20, 0xca, 0x3d, 0x95, 0x37, 0xa3, 0xed, 0x0d, 0x3a, 0x93, 0x0a, 0x88, 0xfb, 0x7a, 0xd5,
	0x76, 0x52, 0x81, 0x99, 0x81, 0x21, 0x13, 0x6a, 0x38, 0x69, 0x46, 0x13, 0x6e, 0xe9, 0x88, 0x54,
	0x20, 0x55, 0xe0, 0xb7, 0x3c, 0x5f, 0xba, 0x89, 0xf8, 0x21, 0x50, 0x9b, 0x02, 0x04, 0x78, 0xbc,
	0x5d, 0x79, 0xf4, 0xde, 0x97, 0x1e, 0xc6, 0x85, 0xdb, 0xf1, 0x05, 0x40, 0x1d, 0xdd, 0xc5, 0x6b,
	0x35, 0x5d, 0xa8, 0x6a, 0x5a, 0x4a, 0x16, 0xa3, 0x14, 0x65, 0xbd, 0x62, 0xe0, 0x81, 0xe7, 0x2c,
	0x7a, 0x8c, 0x07, 0x73, 0x0d, 0x8d, 0x64, 0x5c, 0xc7, 0x97, 0x52, 0x94, 0xad, 0x8d, 0xe3, 0xaf,
	0x9f, 0x76, 0x37, 0xc3, 0x2e, 0x4f, 0x18, 0xd3, 0xdc, 0x98, 0x97, 0x56, 0x4b, 0x25, 0x8a, 0x0b,
	0x65, 0xf4, 0x11, 0xe1, 0x6b, 0x52, 0x49, 0x2b, 0x69, 0x5d, 0xb6, 0x49, 0x4d, 0xdc, 0x4d, 0xbb,
	0xd9, 0xd5, 0xbd, 0xad, 0x3c, 0x18, 0xdb, 0x2e, 0x79, 0xe8, 0x92, 0x3f, 0x05, 0xa9, 0xc6, 0xaf,
	0x8f, 0x4f, 0x87, 0x9d, 0x9f, 0xa7, 0xc3, 0x07, 0x42, 0xda, 0xe9, 0x62, 0x92, 0x57, 0x30, 0x0b,
	0x5d, 0xc2, 0x6f, 0xd7, 0xb0, 0x7d, 0x62, 0x0f, 0xe7, 0xdc, 0x38, 0xc3, 0x87, 0xef, 0xc3, 0xec,
	0x1f, 0xa5, 0xe6, 0xdd, 0xd9, 0xd1, 0xce, 0x7a, 0xcd, 0x05, 0xad, 0x0e, 0x7d, 0x9c, 0xf7, 0x67,
	0x47, 0x3b, 0xa8, 0x58, 0x0f, 0x11, 0x9d, 0x26, 0xfa, 0x8c, 0xf0, 0x0d, 0xcd, 0x67, 0x54, 0x2a,
	0xa9, 0x44, 0x48, 0xdd, 0xfb, 0x0f, 0x53, 0x5f, 0xbf, 0x08, 0xe9, 0x73, 0x3f, 0xc4, 0x51, 0x4d,
	0x8d, 0x2d, 0xc3, 0xf5, 0x9a, 0x72, 0x29, 0xa6, 0x36, 0xbe, 0x9c, 0xa2, 0xac, 0x5b, 0xdc, 0x6c,
	0x19, 0x7f, 0xd4, 0xcf, 0x1c, 0x1e, 0xed, 0xe1, 0xdb, 0xd5, 0x42, 0x6b, 0xae, 0xfe, 0x36, 0xf4,
	0x9d, 0xe1, 0x56, 0x20, 0xff, 0xf0, 0xdc, 0xc1, 0xfd, 0xaa, 0x06, 0xc3, 0x59, 0x7c, 0x25, 0x45,
	0xd9, 0xa0, 0x08, 0xd3, 0x78, 0x7c, 0xbc, 0x4c, 0xd0, 0xc9, 0x32, 0x41, 0x3f, 0x96, 0x09, 0x7a,
	0xbb, 0x4a, 0x3a, 0x27, 0xab, 0xa4, 0xf3, 0x6d, 0x95, 0x74, 0x5e, 0xfd, 0xde, 0x51, 0x0a, 0x25,
	0x2d, 0x27, 0xe7, 0xef, 0xe1, 0xcd, 0xf9, 0x8b, 0x70, 0x4d, 0x27, 0x7d, 0x77, 0x25, 0x1f, 0xfd,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x67, 0xf3, 0xd7, 0x13, 0x30, 0x03, 0x00, 0x00,
}

func (m *RewardPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Closed {
		i--
		if m.Closed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.CurrentRewardHeight != 0 {
		i = encodeVarintRewardPool(dAtA, i, uint64(m.CurrentRewardHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.LastRewardHeight != 0 {
		i = encodeVarintRewardPool(dAtA, i, uint64(m.LastRewardHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RemainingCoins) > 0 {
		for iNdEx := len(m.RemainingCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RemainingCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRewardPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InitialCoins) > 0 {
		for iNdEx := len(m.InitialCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRewardPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintRewardPool(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchId != 0 {
		i = encodeVarintRewardPool(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRewardPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewardPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchId != 0 {
		n += 1 + sovRewardPool(uint64(m.LaunchId))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovRewardPool(uint64(l))
	}
	if len(m.InitialCoins) > 0 {
		for _, e := range m.InitialCoins {
			l = e.Size()
			n += 1 + l + sovRewardPool(uint64(l))
		}
	}
	if len(m.RemainingCoins) > 0 {
		for _, e := range m.RemainingCoins {
			l = e.Size()
			n += 1 + l + sovRewardPool(uint64(l))
		}
	}
	if m.LastRewardHeight != 0 {
		n += 1 + sovRewardPool(uint64(m.LastRewardHeight))
	}
	if m.CurrentRewardHeight != 0 {
		n += 1 + sovRewardPool(uint64(m.CurrentRewardHeight))
	}
	if m.Closed {
		n += 2
	}
	return n
}

func sovRewardPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewardPool(x uint64) (n int) {
	return sovRewardPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewardPool
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
			return fmt.Errorf("proto: RewardPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
				return ErrInvalidLengthRewardPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
				return ErrInvalidLengthRewardPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewardPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialCoins = append(m.InitialCoins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.InitialCoins[len(m.InitialCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
				return ErrInvalidLengthRewardPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewardPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemainingCoins = append(m.RemainingCoins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.RemainingCoins[len(m.RemainingCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastRewardHeight", wireType)
			}
			m.LastRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentRewardHeight", wireType)
			}
			m.CurrentRewardHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentRewardHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Closed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Closed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRewardPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewardPool
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
func skipRewardPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewardPool
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
					return 0, ErrIntOverflowRewardPool
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
					return 0, ErrIntOverflowRewardPool
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
				return 0, ErrInvalidLengthRewardPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewardPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewardPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewardPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewardPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewardPool = fmt.Errorf("proto: unexpected end of group")
)
