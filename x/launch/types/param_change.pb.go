// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/launch/v1/param_change.proto

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

type ParamChange struct {
	LaunchId uint64 `protobuf:"varint,1,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	Module   string `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	Param    string `protobuf:"bytes,3,opt,name=param,proto3" json:"param,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ParamChange) Reset()         { *m = ParamChange{} }
func (m *ParamChange) String() string { return proto.CompactTextString(m) }
func (*ParamChange) ProtoMessage()    {}
func (*ParamChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dc6682ab9eafba3, []int{0}
}
func (m *ParamChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamChange.Merge(m, src)
}
func (m *ParamChange) XXX_Size() int {
	return m.Size()
}
func (m *ParamChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamChange.DiscardUnknown(m)
}

var xxx_messageInfo_ParamChange proto.InternalMessageInfo

func (m *ParamChange) GetLaunchId() uint64 {
	if m != nil {
		return m.LaunchId
	}
	return 0
}

func (m *ParamChange) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ParamChange) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *ParamChange) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ParamChange)(nil), "network.launch.v1.ParamChange")
}

func init() {
	proto.RegisterFile("network/launch/v1/param_change.proto", fileDescriptor_2dc6682ab9eafba3)
}

var fileDescriptor_2dc6682ab9eafba3 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc9, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0x49, 0x2c, 0xcd, 0x4b, 0xce, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x48,
	0x2c, 0x4a, 0xcc, 0x8d, 0x4f, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x84, 0xaa, 0xd2, 0x83, 0xa8, 0xd2, 0x2b, 0x33, 0x54, 0xca, 0xe3, 0xe2, 0x0e, 0x00,
	0x29, 0x74, 0x06, 0xab, 0x13, 0x92, 0xe6, 0xe2, 0x84, 0xc8, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xb0, 0x04, 0x71, 0x40, 0x04, 0x3c, 0x53, 0x84, 0xc4, 0xb8, 0xd8, 0x72, 0xf3, 0x53,
	0x4a, 0x73, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x11, 0x2e, 0x56,
	0xb0, 0x65, 0x12, 0xcc, 0x60, 0x61, 0x08, 0x07, 0x24, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1,
	0xa2, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe1, 0x38, 0x39, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x46, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x66, 0x7a, 0x5e, 0x66, 0x49, 0xaa, 0x3e, 0xcc, 0x53, 0x15, 0x30, 0x6f, 0x95, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x7d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x06, 0xea, 0xca, 0x65,
	0xf5, 0x00, 0x00, 0x00,
}

func (m *ParamChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintParamChange(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Param) > 0 {
		i -= len(m.Param)
		copy(dAtA[i:], m.Param)
		i = encodeVarintParamChange(dAtA, i, uint64(len(m.Param)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintParamChange(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchId != 0 {
		i = encodeVarintParamChange(dAtA, i, uint64(m.LaunchId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParamChange(dAtA []byte, offset int, v uint64) int {
	offset -= sovParamChange(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParamChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchId != 0 {
		n += 1 + sovParamChange(uint64(m.LaunchId))
	}
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovParamChange(uint64(l))
	}
	l = len(m.Param)
	if l > 0 {
		n += 1 + l + sovParamChange(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovParamChange(uint64(l))
	}
	return n
}

func sovParamChange(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParamChange(x uint64) (n int) {
	return sovParamChange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParamChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParamChange
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
			return fmt.Errorf("proto: ParamChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchId", wireType)
			}
			m.LaunchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParamChange
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
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParamChange
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
				return ErrInvalidLengthParamChange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParamChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParamChange
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
				return ErrInvalidLengthParamChange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParamChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Param = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParamChange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParamChange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParamChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParamChange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParamChange
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
func skipParamChange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParamChange
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
					return 0, ErrIntOverflowParamChange
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
					return 0, ErrIntOverflowParamChange
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
				return 0, ErrInvalidLengthParamChange
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParamChange
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParamChange
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParamChange        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParamChange          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParamChange = fmt.Errorf("proto: unexpected end of group")
)
