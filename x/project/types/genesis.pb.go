// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/project/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the project module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params             Params           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ProjectList        []Project        `protobuf:"bytes,2,rep,name=projectList,proto3" json:"projectList"`
	ProjectCount       uint64           `protobuf:"varint,3,opt,name=projectCount,proto3" json:"projectCount,omitempty"`
	ProjectChainsList  []ProjectChains  `protobuf:"bytes,4,rep,name=projectChainsList,proto3" json:"projectChainsList"`
	MainnetAccountList []MainnetAccount `protobuf:"bytes,5,rep,name=mainnetAccountList,proto3" json:"mainnetAccountList"`
	TotalShares        uint64           `protobuf:"varint,6,opt,name=totalShares,proto3" json:"totalShares,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f86dacbbe581600, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetProjectList() []Project {
	if m != nil {
		return m.ProjectList
	}
	return nil
}

func (m *GenesisState) GetProjectCount() uint64 {
	if m != nil {
		return m.ProjectCount
	}
	return 0
}

func (m *GenesisState) GetProjectChainsList() []ProjectChains {
	if m != nil {
		return m.ProjectChainsList
	}
	return nil
}

func (m *GenesisState) GetMainnetAccountList() []MainnetAccount {
	if m != nil {
		return m.MainnetAccountList
	}
	return nil
}

func (m *GenesisState) GetTotalShares() uint64 {
	if m != nil {
		return m.TotalShares
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "network.project.v1.GenesisState")
}

func init() { proto.RegisterFile("network/project/v1/genesis.proto", fileDescriptor_7f86dacbbe581600) }

var fileDescriptor_7f86dacbbe581600 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0xdb, 0xde, 0x4a, 0x38, 0x5d, 0x6a, 0x31, 0x44, 0x41, 0x4a, 0x43, 0x17, 0x02,
	0x43, 0xac, 0x96, 0x99, 0x81, 0x76, 0x60, 0x01, 0x09, 0xb5, 0x42, 0x42, 0x2c, 0x95, 0x1b, 0x59,
	0xa9, 0x81, 0xd8, 0x51, 0xe2, 0x16, 0x78, 0x04, 0x36, 0x1e, 0x83, 0x91, 0xc7, 0xe8, 0xd8, 0x91,
	0x09, 0xa1, 0x76, 0xe0, 0x35, 0x50, 0x6d, 0x07, 0xb5, 0x6a, 0x60, 0x89, 0xec, 0xa3, 0xcf, 0xdf,
	0x39, 0x27, 0x3f, 0xf4, 0x39, 0x95, 0x0f, 0x22, 0xbb, 0xc3, 0x69, 0x26, 0x6e, 0x69, 0x24, 0xf1,
	0xb4, 0x8d, 0x63, 0xca, 0x69, 0xce, 0xf2, 0x30, 0xcd, 0x84, 0x14, 0x08, 0x19, 0x22, 0x34, 0x44,
	0x38, 0x6d, 0xbb, 0x0d, 0x92, 0x30, 0x2e, 0xb0, 0xfa, 0x6a, 0xcc, 0xdd, 0x8d, 0x45, 0x2c, 0xd4,
	0x11, 0xaf, 0x4e, 0xa6, 0x1a, 0x94, 0xe8, 0x13, 0xc2, 0x38, 0xa7, 0x72, 0x48, 0xa2, 0x48, 0x4c,
	0xb8, 0x34, 0x64, 0xb3, 0x84, 0x4c, 0x49, 0x46, 0x12, 0x33, 0x87, 0x5b, 0x36, 0x69, 0x31, 0x92,
	0x26, 0x0e, 0x7e, 0x27, 0x86, 0xd1, 0x98, 0x30, 0x6e, 0x54, 0xad, 0xe7, 0x0a, 0xac, 0x9f, 0xe9,
	0x25, 0x07, 0x92, 0x48, 0x8a, 0x4e, 0x60, 0x4d, 0xf7, 0x72, 0x80, 0x0f, 0x02, 0xbb, 0xe3, 0x86,
	0xdb, 0x4b, 0x87, 0x97, 0x8a, 0xe8, 0xee, 0xcc, 0x3e, 0x9a, 0xd6, 0xeb, 0xd7, 0xdb, 0x11, 0xe8,
	0x9b, 0x47, 0xa8, 0x07, 0x6d, 0xc3, 0x9d, 0xb3, 0x5c, 0x3a, 0xff, 0xfc, 0x4a, 0x60, 0x77, 0xf6,
	0x4a, 0x1d, 0xfa, 0xd8, 0xad, 0xae, 0x24, 0xfd, 0xf5, 0x57, 0xa8, 0x05, 0xeb, 0xe6, 0xda, 0x5b,
	0xfd, 0x16, 0xa7, 0xe2, 0x83, 0xa0, 0xda, 0xdf, 0xa8, 0xa1, 0x2b, 0xd8, 0x28, 0xee, 0x6a, 0x1f,
	0xd5, 0xae, 0xaa, 0xda, 0xed, 0xff, 0xd1, 0x4e, 0xc3, 0xa6, 0xe9, 0xb6, 0x01, 0x5d, 0x43, 0x64,
	0x42, 0x39, 0xd5, 0x99, 0x28, 0xef, 0x7f, 0xe5, 0x6d, 0x95, 0x79, 0x2f, 0x36, 0x68, 0x23, 0x2e,
	0x71, 0x20, 0x1f, 0xda, 0x52, 0x48, 0x72, 0x3f, 0x18, 0x93, 0x8c, 0xe6, 0x4e, 0x4d, 0xed, 0xb4,
	0x5e, 0xea, 0xf6, 0x66, 0x0b, 0x0f, 0xcc, 0x17, 0x1e, 0xf8, 0x5c, 0x78, 0xe0, 0x65, 0xe9, 0x59,
	0xf3, 0xa5, 0x67, 0xbd, 0x2f, 0x3d, 0xeb, 0xe6, 0x30, 0x66, 0x72, 0x3c, 0x19, 0x85, 0x91, 0x48,
	0x30, 0x8b, 0x39, 0x93, 0x14, 0x17, 0x01, 0x3f, 0xfe, 0x44, 0x2c, 0x9f, 0x52, 0x9a, 0x8f, 0x6a,
	0x2a, 0xd7, 0xe3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x6f, 0xcf, 0x96, 0xce, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalShares != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TotalShares))
		i--
		dAtA[i] = 0x30
	}
	if len(m.MainnetAccountList) > 0 {
		for iNdEx := len(m.MainnetAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MainnetAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ProjectChainsList) > 0 {
		for iNdEx := len(m.ProjectChainsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectChainsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ProjectCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ProjectCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProjectList) > 0 {
		for iNdEx := len(m.ProjectList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ProjectList) > 0 {
		for _, e := range m.ProjectList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ProjectCount != 0 {
		n += 1 + sovGenesis(uint64(m.ProjectCount))
	}
	if len(m.ProjectChainsList) > 0 {
		for _, e := range m.ProjectChainsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MainnetAccountList) > 0 {
		for _, e := range m.MainnetAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TotalShares != 0 {
		n += 1 + sovGenesis(uint64(m.TotalShares))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectList = append(m.ProjectList, Project{})
			if err := m.ProjectList[len(m.ProjectList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectCount", wireType)
			}
			m.ProjectCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectChainsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectChainsList = append(m.ProjectChainsList, ProjectChains{})
			if err := m.ProjectChainsList[len(m.ProjectChainsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainnetAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MainnetAccountList = append(m.MainnetAccountList, MainnetAccount{})
			if err := m.MainnetAccountList[len(m.MainnetAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			m.TotalShares = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalShares |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
