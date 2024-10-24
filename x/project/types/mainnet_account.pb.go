// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/project/v1/mainnet_account.proto

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

type MainnetAccount struct {
	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Shares    Shares `protobuf:"bytes,3,rep,name=shares,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"shares"`
}

func (m *MainnetAccount) Reset()         { *m = MainnetAccount{} }
func (m *MainnetAccount) String() string { return proto.CompactTextString(m) }
func (*MainnetAccount) ProtoMessage()    {}
func (*MainnetAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_071f89fc0af8d597, []int{0}
}
func (m *MainnetAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MainnetAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MainnetAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MainnetAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainnetAccount.Merge(m, src)
}
func (m *MainnetAccount) XXX_Size() int {
	return m.Size()
}
func (m *MainnetAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MainnetAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MainnetAccount proto.InternalMessageInfo

func (m *MainnetAccount) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *MainnetAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MainnetAccount) GetShares() Shares {
	if m != nil {
		return m.Shares
	}
	return nil
}

type MainnetAccountBalance struct {
	ProjectId uint64                                   `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Address   string                                   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Coins     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *MainnetAccountBalance) Reset()         { *m = MainnetAccountBalance{} }
func (m *MainnetAccountBalance) String() string { return proto.CompactTextString(m) }
func (*MainnetAccountBalance) ProtoMessage()    {}
func (*MainnetAccountBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_071f89fc0af8d597, []int{1}
}
func (m *MainnetAccountBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MainnetAccountBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MainnetAccountBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MainnetAccountBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainnetAccountBalance.Merge(m, src)
}
func (m *MainnetAccountBalance) XXX_Size() int {
	return m.Size()
}
func (m *MainnetAccountBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_MainnetAccountBalance.DiscardUnknown(m)
}

var xxx_messageInfo_MainnetAccountBalance proto.InternalMessageInfo

func (m *MainnetAccountBalance) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *MainnetAccountBalance) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MainnetAccountBalance) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*MainnetAccount)(nil), "network.project.v1.MainnetAccount")
	proto.RegisterType((*MainnetAccountBalance)(nil), "network.project.v1.MainnetAccountBalance")
}

func init() {
	proto.RegisterFile("network/project/v1/mainnet_account.proto", fileDescriptor_071f89fc0af8d597)
}

var fileDescriptor_071f89fc0af8d597 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0xd1, 0x2f, 0x33, 0xd4, 0xcf,
	0x4d, 0xcc, 0xcc, 0xcb, 0x4b, 0x2d, 0x89, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xaa, 0xd4, 0x83, 0xaa, 0xd4, 0x2b, 0x33, 0x94, 0x12,
	0x4c, 0xcc, 0xcd, 0xcc, 0xcb, 0xd7, 0x07, 0x93, 0x10, 0x65, 0x52, 0x72, 0xc9, 0xf9, 0xc5, 0xb9,
	0xf9, 0xc5, 0xfa, 0x49, 0x89, 0xc5, 0xa9, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89, 0x86, 0xfa,
	0xc9, 0xf9, 0x99, 0x79, 0x50, 0x79, 0x49, 0x88, 0x7c, 0x3c, 0x98, 0xa7, 0x0f, 0xe1, 0x40, 0xa5,
	0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0x21, 0xe2, 0x20, 0x16, 0x44, 0x54, 0xe9, 0x05, 0x23, 0x17, 0x9f,
	0x2f, 0xc4, 0x45, 0x8e, 0x10, 0x07, 0x09, 0xc9, 0x72, 0x71, 0x41, 0x1d, 0x11, 0x9f, 0x99, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x09, 0x15, 0xf1, 0x4c, 0x11, 0x32, 0xe2, 0x62, 0x4f,
	0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x74, 0x92, 0xb8, 0xb4,
	0x45, 0x57, 0x04, 0x6a, 0x95, 0x23, 0x44, 0x26, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xa6,
	0x50, 0xa8, 0x89, 0x91, 0x8b, 0xad, 0x38, 0x23, 0xb1, 0x28, 0xb5, 0x58, 0x82, 0x59, 0x81, 0x59,
	0x83, 0xdb, 0x48, 0x52, 0x0f, 0xaa, 0x01, 0xe4, 0x11, 0x3d, 0xa8, 0x47, 0xf4, 0x9c, 0xf3, 0x33,
	0xf3, 0x9c, 0xfc, 0x4f, 0xdc, 0x93, 0x67, 0xf8, 0x75, 0x4f, 0x5e, 0x3d, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17, 0xea, 0x11, 0x28, 0xa5, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x52,
	0x59, 0x90, 0x5a, 0x0c, 0xd6, 0xb0, 0xea, 0xbe, 0x3c, 0x5b, 0x30, 0xd8, 0xec, 0x59, 0xcf, 0x37,
	0x68, 0xf1, 0xe4, 0xa4, 0xa6, 0x27, 0x26, 0x57, 0xc6, 0x83, 0xc2, 0xa4, 0x78, 0xc5, 0xf3, 0x0d,
	0x5a, 0x8c, 0x41, 0x50, 0x9b, 0x95, 0xfa, 0x99, 0xb8, 0x44, 0x51, 0xbd, 0xea, 0x94, 0x98, 0x93,
	0x98, 0x97, 0x9c, 0x4a, 0x0b, 0x1f, 0x2f, 0x62, 0xe4, 0x62, 0x05, 0xbb, 0x81, 0xb0, 0x87, 0x0b,
	0x49, 0xf7, 0xb0, 0x06, 0x91, 0x4a, 0x71, 0x05, 0x09, 0xc4, 0x69, 0x4e, 0xce, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x89, 0x64, 0x41, 0x66, 0x7a, 0x5e, 0x66, 0x49,
	0xaa, 0x3e, 0x2c, 0x29, 0x57, 0xc0, 0x13, 0x33, 0xd8, 0x9e, 0x24, 0x36, 0x70, 0x42, 0x32, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x88, 0x92, 0x79, 0x39, 0xec, 0x02, 0x00, 0x00,
}

func (m *MainnetAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MainnetAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MainnetAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMainnetAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMainnetAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProjectId != 0 {
		i = encodeVarintMainnetAccount(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MainnetAccountBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MainnetAccountBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MainnetAccountBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMainnetAccount(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMainnetAccount(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProjectId != 0 {
		i = encodeVarintMainnetAccount(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMainnetAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovMainnetAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MainnetAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectId != 0 {
		n += 1 + sovMainnetAccount(uint64(m.ProjectId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMainnetAccount(uint64(l))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovMainnetAccount(uint64(l))
		}
	}
	return n
}

func (m *MainnetAccountBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectId != 0 {
		n += 1 + sovMainnetAccount(uint64(m.ProjectId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMainnetAccount(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovMainnetAccount(uint64(l))
		}
	}
	return n
}

func sovMainnetAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMainnetAccount(x uint64) (n int) {
	return sovMainnetAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MainnetAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMainnetAccount
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
			return fmt.Errorf("proto: MainnetAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MainnetAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainnetAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainnetAccount
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
				return ErrInvalidLengthMainnetAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMainnetAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainnetAccount
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
				return ErrInvalidLengthMainnetAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMainnetAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMainnetAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMainnetAccount
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
func (m *MainnetAccountBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMainnetAccount
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
			return fmt.Errorf("proto: MainnetAccountBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MainnetAccountBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainnetAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainnetAccount
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
				return ErrInvalidLengthMainnetAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMainnetAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainnetAccount
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
				return ErrInvalidLengthMainnetAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMainnetAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMainnetAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMainnetAccount
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
func skipMainnetAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMainnetAccount
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
					return 0, ErrIntOverflowMainnetAccount
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
					return 0, ErrIntOverflowMainnetAccount
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
				return 0, ErrInvalidLengthMainnetAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMainnetAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMainnetAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMainnetAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMainnetAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMainnetAccount = fmt.Errorf("proto: unexpected end of group")
)
