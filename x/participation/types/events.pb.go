// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/participation/v1/events.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type EventAllocationsUsed struct {
	Participant    string                `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	AuctionId      uint64                `protobuf:"varint,2,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	NumAllocations cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=num_allocations,json=numAllocations,proto3,customtype=cosmossdk.io/math.Int" json:"num_allocations"`
}

func (m *EventAllocationsUsed) Reset()         { *m = EventAllocationsUsed{} }
func (m *EventAllocationsUsed) String() string { return proto.CompactTextString(m) }
func (*EventAllocationsUsed) ProtoMessage()    {}
func (*EventAllocationsUsed) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c252e1ef497c1ad, []int{0}
}
func (m *EventAllocationsUsed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAllocationsUsed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAllocationsUsed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAllocationsUsed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAllocationsUsed.Merge(m, src)
}
func (m *EventAllocationsUsed) XXX_Size() int {
	return m.Size()
}
func (m *EventAllocationsUsed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAllocationsUsed.DiscardUnknown(m)
}

var xxx_messageInfo_EventAllocationsUsed proto.InternalMessageInfo

func (m *EventAllocationsUsed) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *EventAllocationsUsed) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

type EventAllocationsWithdrawn struct {
	Participant string `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	AuctionId   uint64 `protobuf:"varint,2,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
}

func (m *EventAllocationsWithdrawn) Reset()         { *m = EventAllocationsWithdrawn{} }
func (m *EventAllocationsWithdrawn) String() string { return proto.CompactTextString(m) }
func (*EventAllocationsWithdrawn) ProtoMessage()    {}
func (*EventAllocationsWithdrawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c252e1ef497c1ad, []int{1}
}
func (m *EventAllocationsWithdrawn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAllocationsWithdrawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAllocationsWithdrawn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAllocationsWithdrawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAllocationsWithdrawn.Merge(m, src)
}
func (m *EventAllocationsWithdrawn) XXX_Size() int {
	return m.Size()
}
func (m *EventAllocationsWithdrawn) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAllocationsWithdrawn.DiscardUnknown(m)
}

var xxx_messageInfo_EventAllocationsWithdrawn proto.InternalMessageInfo

func (m *EventAllocationsWithdrawn) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *EventAllocationsWithdrawn) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func init() {
	proto.RegisterType((*EventAllocationsUsed)(nil), "network.participation.v1.EventAllocationsUsed")
	proto.RegisterType((*EventAllocationsWithdrawn)(nil), "network.participation.v1.EventAllocationsWithdrawn")
}

func init() {
	proto.RegisterFile("network/participation/v1/events.proto", fileDescriptor_2c252e1ef497c1ad)
}

var fileDescriptor_2c252e1ef497c1ad = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x1c, 0xc4, 0x37, 0x2a, 0x42, 0x23, 0x28, 0x2c, 0x15, 0xb6, 0x05, 0xb7, 0xa5, 0x20, 0x14, 0xa4,
	0x1b, 0x8a, 0x37, 0x6f, 0x2d, 0x78, 0xd8, 0x6b, 0x55, 0x04, 0x2f, 0x25, 0xdd, 0x84, 0x6d, 0x68,
	0x37, 0x59, 0x92, 0xff, 0x6e, 0xf5, 0x2d, 0x7c, 0x98, 0xbe, 0x83, 0x3d, 0x96, 0x9e, 0xc4, 0x43,
	0x91, 0xf6, 0x45, 0x64, 0x3f, 0xd4, 0xda, 0xbb, 0xb7, 0x64, 0x98, 0xfc, 0x26, 0xc3, 0xe0, 0x4b,
	0xc9, 0x61, 0xa6, 0xf4, 0x84, 0xc4, 0x54, 0x83, 0x08, 0x44, 0x4c, 0x41, 0x28, 0x49, 0xd2, 0x2e,
	0xe1, 0x29, 0x97, 0x60, 0xbc, 0x58, 0x2b, 0x50, 0xb6, 0x53, 0xda, 0xbc, 0x3f, 0x36, 0x2f, 0xed,
	0xd6, 0x6b, 0x81, 0x32, 0x91, 0x32, 0xc3, 0xdc, 0x47, 0x8a, 0x4b, 0xf1, 0xa8, 0x5e, 0x0d, 0x55,
	0xa8, 0x0a, 0x3d, 0x3b, 0x15, 0x6a, 0xeb, 0x0d, 0xe1, 0xea, 0x6d, 0xc6, 0xee, 0x4d, 0xa7, 0x2a,
	0xc8, 0x39, 0xe6, 0xc1, 0x70, 0x66, 0xdf, 0xe0, 0x93, 0x1f, 0xba, 0x04, 0x07, 0x35, 0x51, 0xbb,
	0xd2, 0x77, 0x56, 0xf3, 0x4e, 0xb5, 0xa4, 0xf6, 0x18, 0xd3, 0xdc, 0x98, 0x3b, 0xd0, 0x42, 0x86,
	0x83, 0x5d, 0xb3, 0x7d, 0x81, 0x31, 0x4d, 0x82, 0x8c, 0x35, 0x14, 0xcc, 0x39, 0x68, 0xa2, 0xf6,
	0xd1, 0xa0, 0x52, 0x2a, 0x3e, 0xb3, 0xef, 0xf1, 0x99, 0x4c, 0xa2, 0x21, 0xfd, 0x4d, 0x74, 0x0e,
	0x73, 0xfc, 0xd5, 0x62, 0xdd, 0xb0, 0x3e, 0xd6, 0x8d, 0xf3, 0x22, 0xc2, 0xb0, 0x89, 0x27, 0x14,
	0x89, 0x28, 0x8c, 0x3d, 0x5f, 0xc2, 0x6a, 0xde, 0xc1, 0x65, 0xb6, 0x2f, 0x61, 0x70, 0x2a, 0x93,
	0x68, 0xe7, 0xd3, 0xad, 0x14, 0xd7, 0xf6, 0x8b, 0x3c, 0x0a, 0x18, 0x33, 0x4d, 0x67, 0xf2, 0x1f,
	0xdb, 0xf4, 0xfd, 0xc5, 0xc6, 0x45, 0xcb, 0x8d, 0x8b, 0x3e, 0x37, 0x2e, 0x7a, 0xdd, 0xba, 0xd6,
	0x72, 0xeb, 0x5a, 0xef, 0x5b, 0xd7, 0x7a, 0x22, 0xa1, 0x80, 0x71, 0x32, 0xf2, 0x02, 0x15, 0x11,
	0x11, 0x4a, 0x01, 0x9c, 0x7c, 0xef, 0xfb, 0xbc, 0xb7, 0x30, 0xbc, 0xc4, 0xdc, 0x8c, 0x8e, 0xf3,
	0x4d, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x5a, 0x10, 0x2c, 0x07, 0x02, 0x00, 0x00,
}

func (m *EventAllocationsUsed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAllocationsUsed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAllocationsUsed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NumAllocations.Size()
		i -= size
		if _, err := m.NumAllocations.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AuctionId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAllocationsWithdrawn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAllocationsWithdrawn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAllocationsWithdrawn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuctionId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0xa
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
func (m *EventAllocationsUsed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.AuctionId != 0 {
		n += 1 + sovEvents(uint64(m.AuctionId))
	}
	l = m.NumAllocations.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventAllocationsWithdrawn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.AuctionId != 0 {
		n += 1 + sovEvents(uint64(m.AuctionId))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventAllocationsUsed) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventAllocationsUsed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAllocationsUsed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
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
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumAllocations", wireType)
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
			if err := m.NumAllocations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *EventAllocationsWithdrawn) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventAllocationsWithdrawn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAllocationsWithdrawn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
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
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
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
