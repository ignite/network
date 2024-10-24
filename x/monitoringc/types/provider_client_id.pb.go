// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/monitoringc/v1/provider_client_id.proto

package types

import (
	fmt "fmt"
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

type ProviderClientID struct {
	LaunchId uint64 `protobuf:"varint,1,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *ProviderClientID) Reset()         { *m = ProviderClientID{} }
func (m *ProviderClientID) String() string { return proto.CompactTextString(m) }
func (*ProviderClientID) ProtoMessage()    {}
func (*ProviderClientID) Descriptor() ([]byte, []int) {
	return fileDescriptor_20cbe9ae26cd49bf, []int{0}
}
func (m *ProviderClientID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderClientID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderClientID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderClientID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderClientID.Merge(m, src)
}
func (m *ProviderClientID) XXX_Size() int {
	return m.Size()
}
func (m *ProviderClientID) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderClientID.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderClientID proto.InternalMessageInfo

func (m *ProviderClientID) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

func (m *ProviderClientID) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*ProviderClientID)(nil), "network.monitoringc.v1.ProviderClientID")
}

func init() {
	proto.RegisterFile("network/monitoringc/v1/provider_client_id.proto", fileDescriptor_20cbe9ae26cd49bf)
}

var fileDescriptor_20cbe9ae26cd49bf = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcf, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x4f, 0xd6,
	0x2f, 0x33, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0x8a, 0x4f, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x89, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x83, 0x6a, 0xd0,
	0x43, 0xd2, 0xa0, 0x57, 0x66, 0xa8, 0xe4, 0xc3, 0x25, 0x10, 0x00, 0xd5, 0xe3, 0x0c, 0xd6, 0xe2,
	0xe9, 0x22, 0x24, 0xcd, 0xc5, 0x99, 0x93, 0x58, 0x9a, 0x97, 0x9c, 0x11, 0x9f, 0x99, 0x22, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x01, 0x11, 0xf0, 0x4c, 0x01, 0x49, 0xc2, 0xcd, 0x96, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x38, 0xb9, 0x9f, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x66, 0x7a, 0x5e, 0x66, 0x49, 0x2a, 0xdc, 0x0b, 0x15, 0x28, 0x9e, 0x28, 0xa9,
	0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xbb, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x51, 0xa5,
	0x37, 0xa8, 0xe8, 0x00, 0x00, 0x00,
}

func (m *ProviderClientID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderClientID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderClientID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintProviderClientId(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchId != 0 {
		i = encodeVarintProviderClientId(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProviderClientId(dAtA []byte, offset int, v uint64) int {
	offset -= sovProviderClientId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProviderClientID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchId != 0 {
		n += 1 + sovProviderClientId(uint64(m.LaunchId))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovProviderClientId(uint64(l))
	}
	return n
}

func sovProviderClientId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProviderClientId(x uint64) (n int) {
	return sovProviderClientId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderClientID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProviderClientId
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
			return fmt.Errorf("proto: ProviderClientID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderClientID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderClientId
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderClientId
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
				return ErrInvalidLengthProviderClientId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderClientId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProviderClientId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProviderClientId
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
func skipProviderClientId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProviderClientId
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
					return 0, ErrIntOverflowProviderClientId
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
					return 0, ErrIntOverflowProviderClientId
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
				return 0, ErrInvalidLengthProviderClientId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProviderClientId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProviderClientId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProviderClientId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProviderClientId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProviderClientId = fmt.Errorf("proto: unexpected end of group")
)
