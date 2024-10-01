// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/monitoringc/v1/genesis.proto

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

// GenesisState defines the monitoringc module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params                           Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortID                           string                         `protobuf:"bytes,2,opt,name=portID,proto3" json:"portID,omitempty"`
	LaunchIDFromChannelIDList        []LaunchIDFromChannelID        `protobuf:"bytes,3,rep,name=launchIDFromChannelIDList,proto3" json:"launchIDFromChannelIDList"`
	LaunchIDFromVerifiedClientIDList []LaunchIDFromVerifiedClientID `protobuf:"bytes,4,rep,name=launchIDFromVerifiedClientIDList,proto3" json:"launchIDFromVerifiedClientIDList"`
	MonitoringHistoryList            []MonitoringHistory            `protobuf:"bytes,5,rep,name=monitoringHistoryList,proto3" json:"monitoringHistoryList"`
	VerifiedClientIDList             []VerifiedClientID             `protobuf:"bytes,6,rep,name=verifiedClientIDList,proto3" json:"verifiedClientIDList"`
	ProviderClientIDList             []ProviderClientID             `protobuf:"bytes,7,rep,name=providerClientIDList,proto3" json:"providerClientIDList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab576c05c1adbe3d, []int{0}
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

func (m *GenesisState) GetPortID() string {
	if m != nil {
		return m.PortID
	}
	return ""
}

func (m *GenesisState) GetLaunchIDFromChannelIDList() []LaunchIDFromChannelID {
	if m != nil {
		return m.LaunchIDFromChannelIDList
	}
	return nil
}

func (m *GenesisState) GetLaunchIDFromVerifiedClientIDList() []LaunchIDFromVerifiedClientID {
	if m != nil {
		return m.LaunchIDFromVerifiedClientIDList
	}
	return nil
}

func (m *GenesisState) GetMonitoringHistoryList() []MonitoringHistory {
	if m != nil {
		return m.MonitoringHistoryList
	}
	return nil
}

func (m *GenesisState) GetVerifiedClientIDList() []VerifiedClientID {
	if m != nil {
		return m.VerifiedClientIDList
	}
	return nil
}

func (m *GenesisState) GetProviderClientIDList() []ProviderClientID {
	if m != nil {
		return m.ProviderClientIDList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "network.monitoringc.v1.GenesisState")
}

func init() {
	proto.RegisterFile("network/monitoringc/v1/genesis.proto", fileDescriptor_ab576c05c1adbe3d)
}

var fileDescriptor_ab576c05c1adbe3d = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0x87, 0x3b, 0xee, 0x6e, 0xcd, 0xce, 0x7a, 0xb1, 0x59, 0x09, 0x72, 0xa8, 0x8d, 0x7a, 0xa8,
	0x26, 0xb4, 0x01, 0x8d, 0x47, 0x13, 0x81, 0x88, 0x24, 0x98, 0x10, 0x4c, 0x3c, 0x78, 0x69, 0x4a,
	0x3b, 0xb4, 0x13, 0xe9, 0x4c, 0x9d, 0x0e, 0x55, 0xbe, 0x80, 0x67, 0x0f, 0x7e, 0x08, 0x8f, 0x7e,
	0x0c, 0x8e, 0x1c, 0x3d, 0x19, 0x03, 0x07, 0xbf, 0x86, 0x61, 0x3a, 0xc8, 0x9f, 0x74, 0x16, 0x2e,
	0x64, 0x86, 0xbe, 0xcf, 0xef, 0x79, 0x67, 0xe6, 0x85, 0x8f, 0x09, 0xe2, 0x9f, 0x29, 0xfb, 0xe8,
	0x26, 0x94, 0x60, 0x4e, 0x19, 0x26, 0x51, 0xe0, 0xe6, 0x0d, 0x37, 0x42, 0x04, 0x65, 0x38, 0x73,
	0x52, 0x46, 0x39, 0x35, 0x2a, 0xb2, 0xca, 0xd9, 0xa9, 0x72, 0xf2, 0x46, 0xed, 0xae, 0x9f, 0x60,
	0x42, 0x5d, 0xf1, 0x5b, 0x94, 0xd6, 0xae, 0x23, 0x1a, 0x51, 0xb1, 0x74, 0xd7, 0x2b, 0xf9, 0xef,
	0x0b, 0x85, 0x66, 0xe2, 0x4f, 0x49, 0x10, 0x7b, 0x38, 0xf4, 0xc6, 0x8c, 0x26, 0x5e, 0x10, 0xfb,
	0x84, 0xa0, 0x89, 0x87, 0x43, 0xc9, 0xbd, 0x3c, 0x8d, 0xcb, 0x11, 0xc3, 0x63, 0x8c, 0x42, 0x2f,
	0x98, 0x60, 0x44, 0xf8, 0x96, 0x77, 0x15, 0xfc, 0x76, 0xeb, 0xc5, 0x38, 0xe3, 0x94, 0xcd, 0x24,
	0xf0, 0x48, 0x01, 0xa4, 0x3e, 0xf3, 0x93, 0xec, 0x48, 0x6a, 0xca, 0x68, 0x8e, 0x43, 0xc4, 0x4e,
	0x6e, 0x43, 0xd5, 0xf7, 0xc3, 0xef, 0x17, 0xf0, 0x4e, 0xb7, 0x78, 0x82, 0x77, 0xdc, 0xe7, 0xc8,
	0x78, 0x05, 0xf5, 0xa2, 0x85, 0x2a, 0xb0, 0x80, 0x7d, 0xd5, 0x34, 0x9d, 0xf2, 0x27, 0x71, 0x06,
	0xa2, 0xaa, 0x75, 0x39, 0xff, 0xfd, 0x40, 0xfb, 0xf1, 0xf7, 0xe7, 0x53, 0x30, 0x94, 0xa0, 0x51,
	0x81, 0x7a, 0x4a, 0x19, 0xef, 0x75, 0xaa, 0xb7, 0x2c, 0x60, 0x5f, 0x0e, 0xe5, 0xce, 0xf8, 0x04,
	0xef, 0x17, 0xd7, 0xd9, 0xeb, 0xbc, 0x66, 0x34, 0x69, 0x17, 0x6f, 0xd0, 0xeb, 0xf4, 0x71, 0xc6,
	0xab, 0x67, 0xd6, 0x99, 0x7d, 0xd5, 0xac, 0xab, 0x6c, 0xfd, 0x32, 0xb0, 0x75, 0xbe, 0x96, 0x0f,
	0xd5, 0xa9, 0xc6, 0x57, 0x00, 0xad, 0xdd, 0xaf, 0xef, 0xe5, 0x3d, 0xb4, 0xc5, 0x35, 0x48, 0xf5,
	0xb9, 0x50, 0x3f, 0x3f, 0x45, 0x7d, 0xc8, 0xcb, 0x0e, 0x8e, 0x3a, 0x0c, 0x04, 0xef, 0x6d, 0x63,
	0xdf, 0x14, 0x93, 0x20, 0xe4, 0x17, 0x42, 0xfe, 0x44, 0x25, 0x7f, 0x7b, 0x08, 0x49, 0x63, 0x79,
	0x9a, 0x31, 0x82, 0xd7, 0x79, 0xd9, 0x11, 0x75, 0x61, 0xb1, 0x55, 0x16, 0xc5, 0xb1, 0x4a, 0xb3,
	0xd6, 0x8e, 0xcd, 0xfc, 0xed, 0x39, 0x6e, 0xdf, 0xec, 0x18, 0x1c, 0x30, 0x1b, 0x47, 0x59, 0x56,
	0xab, 0x3b, 0x5f, 0x9a, 0x60, 0xb1, 0x34, 0xc1, 0x9f, 0xa5, 0x09, 0xbe, 0xad, 0x4c, 0x6d, 0xb1,
	0x32, 0xb5, 0x5f, 0x2b, 0x53, 0xfb, 0x50, 0x8f, 0x30, 0x8f, 0xa7, 0x23, 0x27, 0xa0, 0x89, 0x8b,
	0x23, 0x82, 0x39, 0xfa, 0x3f, 0xf3, 0x5f, 0xf6, 0xa6, 0x9e, 0xcf, 0x52, 0x94, 0x8d, 0x74, 0x31,
	0xe6, 0xcf, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x85, 0x95, 0x19, 0x8f, 0x7f, 0x04, 0x00, 0x00,
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
	if len(m.ProviderClientIDList) > 0 {
		for iNdEx := len(m.ProviderClientIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderClientIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.VerifiedClientIDList) > 0 {
		for iNdEx := len(m.VerifiedClientIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerifiedClientIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.MonitoringHistoryList) > 0 {
		for iNdEx := len(m.MonitoringHistoryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MonitoringHistoryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.LaunchIDFromVerifiedClientIDList) > 0 {
		for iNdEx := len(m.LaunchIDFromVerifiedClientIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDFromVerifiedClientIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.LaunchIDFromChannelIDList) > 0 {
		for iNdEx := len(m.LaunchIDFromChannelIDList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDFromChannelIDList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortID) > 0 {
		i -= len(m.PortID)
		copy(dAtA[i:], m.PortID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortID)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.PortID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.LaunchIDFromChannelIDList) > 0 {
		for _, e := range m.LaunchIDFromChannelIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDFromVerifiedClientIDList) > 0 {
		for _, e := range m.LaunchIDFromVerifiedClientIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MonitoringHistoryList) > 0 {
		for _, e := range m.MonitoringHistoryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VerifiedClientIDList) > 0 {
		for _, e := range m.VerifiedClientIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProviderClientIDList) > 0 {
		for _, e := range m.ProviderClientIDList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field PortID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDFromChannelIDList", wireType)
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
			m.LaunchIDFromChannelIDList = append(m.LaunchIDFromChannelIDList, LaunchIDFromChannelID{})
			if err := m.LaunchIDFromChannelIDList[len(m.LaunchIDFromChannelIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDFromVerifiedClientIDList", wireType)
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
			m.LaunchIDFromVerifiedClientIDList = append(m.LaunchIDFromVerifiedClientIDList, LaunchIDFromVerifiedClientID{})
			if err := m.LaunchIDFromVerifiedClientIDList[len(m.LaunchIDFromVerifiedClientIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoringHistoryList", wireType)
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
			m.MonitoringHistoryList = append(m.MonitoringHistoryList, MonitoringHistory{})
			if err := m.MonitoringHistoryList[len(m.MonitoringHistoryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiedClientIDList", wireType)
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
			m.VerifiedClientIDList = append(m.VerifiedClientIDList, VerifiedClientID{})
			if err := m.VerifiedClientIDList[len(m.VerifiedClientIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderClientIDList", wireType)
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
			m.ProviderClientIDList = append(m.ProviderClientIDList, ProviderClientID{})
			if err := m.ProviderClientIDList[len(m.ProviderClientIDList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
