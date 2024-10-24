// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/monitoringp/v1/monitoring_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/ignite/network/pkg/types"
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

type MonitoringInfo struct {
	Transmitted     bool                  `protobuf:"varint,1,opt,name=transmitted,proto3" json:"transmitted,omitempty"`
	SignatureCounts types.SignatureCounts `protobuf:"bytes,2,opt,name=signature_counts,json=signatureCounts,proto3" json:"signature_counts"`
}

func (m *MonitoringInfo) Reset()         { *m = MonitoringInfo{} }
func (m *MonitoringInfo) String() string { return proto.CompactTextString(m) }
func (*MonitoringInfo) ProtoMessage()    {}
func (*MonitoringInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e6e0ccfd1879c9c, []int{0}
}
func (m *MonitoringInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonitoringInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonitoringInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonitoringInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringInfo.Merge(m, src)
}
func (m *MonitoringInfo) XXX_Size() int {
	return m.Size()
}
func (m *MonitoringInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringInfo proto.InternalMessageInfo

func (m *MonitoringInfo) GetTransmitted() bool {
	if m != nil {
		return m.Transmitted
	}
	return false
}

func (m *MonitoringInfo) GetSignatureCounts() types.SignatureCounts {
	if m != nil {
		return m.SignatureCounts
	}
	return types.SignatureCounts{}
}

func init() {
	proto.RegisterType((*MonitoringInfo)(nil), "network.monitoringp.v1.MonitoringInfo")
}

func init() {
	proto.RegisterFile("network/monitoringp/v1/monitoring_info.proto", fileDescriptor_1e6e0ccfd1879c9c)
}

var fileDescriptor_1e6e0ccfd1879c9c = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc9, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x2f, 0xd0,
	0x2f, 0x33, 0x44, 0xe2, 0xc6, 0x67, 0xe6, 0xa5, 0xe5, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0x89, 0x41, 0x55, 0xeb, 0x21, 0xa9, 0xd6, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x2b, 0xd1, 0x07, 0xb1, 0x20, 0xaa, 0xa5, 0xe4, 0x60, 0x66, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x23,
	0x19, 0x09, 0x91, 0x57, 0x6a, 0x66, 0xe4, 0xe2, 0xf3, 0x85, 0x0b, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b,
	0x29, 0x70, 0x71, 0x97, 0x14, 0x25, 0xe6, 0x15, 0xe7, 0x66, 0x96, 0x94, 0xa4, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x04, 0x21, 0x0b, 0x09, 0xf9, 0x73, 0x09, 0x14, 0x67, 0xa6, 0xe7, 0x25,
	0x96, 0x94, 0x16, 0xa5, 0xc6, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x1b, 0xc9, 0xe9, 0xc1, 0x5c, 0x07, 0xb6, 0x4f, 0x2f, 0x18, 0xa6, 0xcc, 0x19, 0xac, 0xca,
	0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0xfe, 0x62, 0x34, 0x61, 0xf7, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0xcf, 0x4c, 0xcf, 0xcb, 0x2c, 0x49, 0xd5, 0x87, 0xf9, 0xa8, 0x02, 0x25, 0xbc, 0xc0,
	0xf6, 0x25, 0xb1, 0x81, 0x7d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x55, 0x6d, 0x48, 0xa4,
	0x53, 0x01, 0x00, 0x00,
}

func (m *MonitoringInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoringInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonitoringInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SignatureCounts.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMonitoringInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Transmitted {
		i--
		if m.Transmitted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMonitoringInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovMonitoringInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MonitoringInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Transmitted {
		n += 2
	}
	l = m.SignatureCounts.Size()
	n += 1 + l + sovMonitoringInfo(uint64(l))
	return n
}

func sovMonitoringInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMonitoringInfo(x uint64) (n int) {
	return sovMonitoringInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MonitoringInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoringInfo
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
			return fmt.Errorf("proto: MonitoringInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoringInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transmitted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoringInfo
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
			m.Transmitted = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureCounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoringInfo
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
				return ErrInvalidLengthMonitoringInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoringInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SignatureCounts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoringInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMonitoringInfo
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
func skipMonitoringInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoringInfo
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
					return 0, ErrIntOverflowMonitoringInfo
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
					return 0, ErrIntOverflowMonitoringInfo
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
				return 0, ErrInvalidLengthMonitoringInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMonitoringInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMonitoringInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMonitoringInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoringInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMonitoringInfo = fmt.Errorf("proto: unexpected end of group")
)
