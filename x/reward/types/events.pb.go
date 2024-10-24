// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/reward/v1/events.proto

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

type EventRewardPoolCreated struct {
	LaunchId uint64 `protobuf:"varint,1,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (m *EventRewardPoolCreated) Reset()         { *m = EventRewardPoolCreated{} }
func (m *EventRewardPoolCreated) String() string { return proto.CompactTextString(m) }
func (*EventRewardPoolCreated) ProtoMessage()    {}
func (*EventRewardPoolCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8864807f722abba, []int{0}
}
func (m *EventRewardPoolCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardPoolCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardPoolCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardPoolCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardPoolCreated.Merge(m, src)
}
func (m *EventRewardPoolCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardPoolCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardPoolCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardPoolCreated proto.InternalMessageInfo

func (m *EventRewardPoolCreated) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

func (m *EventRewardPoolCreated) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type EventRewardPoolRemoved struct {
	LaunchId uint64 `protobuf:"varint,1,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
}

func (m *EventRewardPoolRemoved) Reset()         { *m = EventRewardPoolRemoved{} }
func (m *EventRewardPoolRemoved) String() string { return proto.CompactTextString(m) }
func (*EventRewardPoolRemoved) ProtoMessage()    {}
func (*EventRewardPoolRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8864807f722abba, []int{1}
}
func (m *EventRewardPoolRemoved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardPoolRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardPoolRemoved.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardPoolRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardPoolRemoved.Merge(m, src)
}
func (m *EventRewardPoolRemoved) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardPoolRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardPoolRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardPoolRemoved proto.InternalMessageInfo

func (m *EventRewardPoolRemoved) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

type EventRewardsDistributed struct {
	LaunchId uint64                                   `protobuf:"varint,1,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	Receiver string                                   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Rewards  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=rewards,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"rewards"`
}

func (m *EventRewardsDistributed) Reset()         { *m = EventRewardsDistributed{} }
func (m *EventRewardsDistributed) String() string { return proto.CompactTextString(m) }
func (*EventRewardsDistributed) ProtoMessage()    {}
func (*EventRewardsDistributed) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8864807f722abba, []int{2}
}
func (m *EventRewardsDistributed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardsDistributed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardsDistributed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardsDistributed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardsDistributed.Merge(m, src)
}
func (m *EventRewardsDistributed) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardsDistributed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardsDistributed.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardsDistributed proto.InternalMessageInfo

func (m *EventRewardsDistributed) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

func (m *EventRewardsDistributed) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *EventRewardsDistributed) GetRewards() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*EventRewardPoolCreated)(nil), "network.reward.v1.EventRewardPoolCreated")
	proto.RegisterType((*EventRewardPoolRemoved)(nil), "network.reward.v1.EventRewardPoolRemoved")
	proto.RegisterType((*EventRewardsDistributed)(nil), "network.reward.v1.EventRewardsDistributed")
}

func init() { proto.RegisterFile("network/reward/v1/events.proto", fileDescriptor_c8864807f722abba) }

var fileDescriptor_c8864807f722abba = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xbe, 0x4d, 0x10, 0x24, 0x86, 0x26, 0xa7, 0x08, 0x9c, 0x20, 0xed, 0x9d, 0x8e, 0x02, 0x2b,
	0x52, 0xbc, 0x32, 0x3f, 0x0f, 0x80, 0x03, 0x05, 0x1d, 0x32, 0x1d, 0x8d, 0x65, 0x7b, 0x47, 0xce,
	0xea, 0xec, 0x1d, 0xb3, 0xbb, 0xe7, 0x90, 0xb7, 0xa0, 0xa1, 0xe1, 0x01, 0x10, 0xa2, 0x4a, 0xc1,
	0x43, 0xa4, 0x8c, 0xa8, 0xa8, 0x02, 0xba, 0x2b, 0xf2, 0x0e, 0x54, 0xc8, 0xde, 0x35, 0x8a, 0x94,
	0x82, 0x6b, 0xec, 0x99, 0xf9, 0xbe, 0xd1, 0xf7, 0x69, 0xe7, 0xf3, 0xa8, 0x04, 0x73, 0x82, 0x6a,
	0xce, 0x14, 0x9c, 0x64, 0x8a, 0xb3, 0x36, 0x62, 0xd0, 0x82, 0x34, 0x3a, 0x6c, 0x14, 0x1a, 0x1c,
	0xef, 0x38, 0x3c, 0xb4, 0x78, 0xd8, 0x46, 0xfb, 0x3b, 0x59, 0x2d, 0x24, 0xb2, 0xfe, 0x6b, 0x59,
	0xfb, 0xb4, 0x40, 0x5d, 0xa3, 0x66, 0x79, 0xa6, 0x81, 0xb5, 0x51, 0x0e, 0x26, 0x8b, 0x58, 0x81,
	0x42, 0x3a, 0x7c, 0xcf, 0xe2, 0x69, 0xdf, 0x31, 0xdb, 0x38, 0x68, 0xb7, 0xc4, 0x12, 0xed, 0xbc,
	0xab, 0xdc, 0xf4, 0xd1, 0x4d, 0x5b, 0xb6, 0x4a, 0x1b, 0xc4, 0xca, 0x92, 0x66, 0x73, 0xef, 0xfe,
	0xab, 0xce, 0x6b, 0xd2, 0x23, 0x6f, 0x10, 0xab, 0x23, 0x05, 0x99, 0x01, 0x3e, 0x7e, 0xe8, 0x6d,
	0x57, 0xd9, 0x42, 0x16, 0xc7, 0xa9, 0xe0, 0x3e, 0x99, 0x92, 0xe0, 0x56, 0xb2, 0x65, 0x07, 0xaf,
	0xf9, 0xf8, 0x99, 0xb7, 0xd5, 0x28, 0x6c, 0x05, 0x07, 0xe5, 0x6f, 0x4c, 0x49, 0xb0, 0x1d, 0xfb,
	0x3f, 0xbe, 0x1f, 0xee, 0x3a, 0x57, 0x2f, 0x38, 0x57, 0xa0, 0xf5, 0x5b, 0xa3, 0x84, 0x2c, 0x93,
	0x7f, 0xcc, 0xd9, 0xf3, 0x1b, 0x62, 0x09, 0xd4, 0xd8, 0xfe, 0x47, 0x6c, 0xf6, 0x69, 0xc3, 0x7b,
	0x70, 0x6d, 0x4f, 0xbf, 0x14, 0xda, 0x28, 0x91, 0x2f, 0xd6, 0x71, 0xa9, 0xa0, 0x00, 0xd1, 0xae,
	0xe3, 0x72, 0x60, 0x8e, 0xbf, 0x10, 0xef, 0x8e, 0x7d, 0x28, 0xed, 0x6f, 0x4e, 0x37, 0x83, 0xbb,
	0x4f, 0xf6, 0x42, 0xb7, 0xd2, 0xdd, 0x26, 0x74, 0xb7, 0x09, 0x8f, 0x50, 0xc8, 0xf8, 0xfd, 0xf9,
	0xe5, 0x64, 0xf4, 0xe7, 0x72, 0xf2, 0xb8, 0x14, 0xe6, 0x78, 0x91, 0x87, 0x05, 0xd6, 0xee, 0x36,
	0xee, 0x77, 0xa8, 0xf9, 0x9c, 0x99, 0xd3, 0x06, 0x74, 0xbf, 0xf0, 0xed, 0xd7, 0x24, 0x58, 0x93,
	0xaa, 0x3f, 0x5f, 0x9d, 0x1d, 0xdc, 0xab, 0xa0, 0xcc, 0x8a, 0xd3, 0xb4, 0x0b, 0x82, 0xfe, 0x7a,
	0x75, 0x76, 0x40, 0x92, 0xc1, 0x5c, 0x1c, 0x9f, 0x2f, 0x29, 0xb9, 0x58, 0x52, 0xf2, 0x7b, 0x49,
	0xc9, 0xc7, 0x15, 0x1d, 0x5d, 0xac, 0xe8, 0xe8, 0xe7, 0x8a, 0x8e, 0xde, 0x5d, 0x97, 0x10, 0xa5,
	0x14, 0x06, 0xd8, 0x10, 0x86, 0x0f, 0x43, 0x1c, 0x7a, 0xa1, 0xfc, 0x76, 0x1f, 0x83, 0xa7, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x39, 0x4e, 0x0e, 0xc4, 0x02, 0x00, 0x00,
}

func (m *EventRewardPoolCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardPoolCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardPoolCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRewardPoolRemoved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardPoolRemoved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardPoolRemoved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LaunchId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRewardsDistributed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardsDistributed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardsDistributed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRewardPoolCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchId != 0 {
		n += 1 + sovEvents(uint64(m.LaunchId))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRewardPoolRemoved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchId != 0 {
		n += 1 + sovEvents(uint64(m.LaunchId))
	}
	return n
}

func (m *EventRewardsDistributed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchId != 0 {
		n += 1 + sovEvents(uint64(m.LaunchId))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRewardPoolCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardPoolCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardPoolCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRewardPoolRemoved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardPoolRemoved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardPoolRemoved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRewardsDistributed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardsDistributed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardsDistributed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
