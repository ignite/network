// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/project/v1/project.proto

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

type Project struct {
	ProjectId          uint64                                   `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName        string                                   `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	CoordinatorId      uint64                                   `protobuf:"varint,3,opt,name=coordinator_id,json=coordinatorId,proto3" json:"coordinator_id,omitempty"`
	CreatedAt          int64                                    `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	MainnetId          uint64                                   `protobuf:"varint,5,opt,name=mainnet_id,json=mainnetId,proto3" json:"mainnet_id,omitempty"`
	MainnetInitialized bool                                     `protobuf:"varint,6,opt,name=mainnet_initialized,json=mainnetInitialized,proto3" json:"mainnet_initialized,omitempty"`
	TotalSupply        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=total_supply,json=totalSupply,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_supply"`
	AllocatedShares    Shares                                   `protobuf:"bytes,8,rep,name=allocated_shares,json=allocatedShares,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"allocated_shares"`
	SpecialAllocations SpecialAllocations                       `protobuf:"bytes,9,opt,name=special_allocations,json=specialAllocations,proto3" json:"special_allocations"`
	Metadata           []byte                                   `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_bce081b9714122e9, []int{0}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Project.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return m.Size()
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *Project) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *Project) GetCoordinatorId() uint64 {
	if m != nil {
		return m.CoordinatorId
	}
	return 0
}

func (m *Project) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Project) GetMainnetId() uint64 {
	if m != nil {
		return m.MainnetId
	}
	return 0
}

func (m *Project) GetMainnetInitialized() bool {
	if m != nil {
		return m.MainnetInitialized
	}
	return false
}

func (m *Project) GetTotalSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

func (m *Project) GetAllocatedShares() Shares {
	if m != nil {
		return m.AllocatedShares
	}
	return nil
}

func (m *Project) GetSpecialAllocations() SpecialAllocations {
	if m != nil {
		return m.SpecialAllocations
	}
	return SpecialAllocations{}
}

func (m *Project) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Project)(nil), "network.project.v1.Project")
}

func init() { proto.RegisterFile("network/project/v1/project.proto", fileDescriptor_bce081b9714122e9) }

var fileDescriptor_bce081b9714122e9 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x6b, 0xda, 0x75, 0xad, 0x5b, 0xfe, 0x79, 0x1c, 0xb2, 0x4a, 0xa4, 0x01, 0x09, 0x08,
	0x13, 0xc4, 0xea, 0x78, 0x82, 0x75, 0xa7, 0x5e, 0x00, 0xa5, 0x37, 0x24, 0x14, 0xb9, 0xb1, 0xd5,
	0x99, 0x25, 0x76, 0x88, 0xbd, 0x42, 0x79, 0x0a, 0x2e, 0x5c, 0x78, 0x00, 0x84, 0xb8, 0xb0, 0xc7,
	0xd8, 0x71, 0x47, 0x4e, 0x03, 0xb5, 0x87, 0xbd, 0x03, 0x27, 0x14, 0xc7, 0xe9, 0x26, 0x6d, 0x08,
	0x76, 0x69, 0xed, 0xef, 0xf7, 0xeb, 0x5f, 0x3e, 0xfe, 0xd9, 0x86, 0x9e, 0x60, 0xfa, 0x9d, 0xcc,
	0xf7, 0x71, 0x96, 0xcb, 0x37, 0x2c, 0xd6, 0x78, 0x36, 0xa8, 0x86, 0x41, 0x96, 0x4b, 0x2d, 0x11,
	0xb2, 0x89, 0xa0, 0x92, 0x67, 0x83, 0xde, 0x6d, 0x92, 0x72, 0x21, 0xb1, 0xf9, 0x2d, 0x63, 0x3d,
	0x37, 0x96, 0x2a, 0x95, 0x0a, 0x4f, 0x88, 0x62, 0x78, 0x36, 0x98, 0x30, 0x4d, 0x06, 0x38, 0x96,
	0x5c, 0x58, 0x7f, 0xb3, 0xf4, 0x23, 0x33, 0xc3, 0xe5, 0xc4, 0x5a, 0x77, 0xa6, 0x72, 0x2a, 0x4b,
	0xbd, 0x18, 0x59, 0xf5, 0xc9, 0x25, 0x64, 0x2a, 0x63, 0x31, 0x27, 0x49, 0x44, 0x92, 0x44, 0xc6,
	0x44, 0x73, 0x29, 0x6c, 0x8d, 0xfb, 0x5f, 0xd6, 0xe0, 0xfa, 0xcb, 0x32, 0x88, 0xee, 0x42, 0x68,
	0xd7, 0x44, 0x9c, 0x3a, 0xc0, 0x03, 0x7e, 0x23, 0x6c, 0x5b, 0x65, 0x44, 0xd1, 0x3d, 0xd8, 0xad,
	0x6c, 0x41, 0x52, 0xe6, 0x5c, 0xf3, 0x80, 0xdf, 0x0e, 0x3b, 0x56, 0x7b, 0x4e, 0x52, 0x86, 0x1e,
	0xc0, 0x1b, 0xb1, 0x94, 0x39, 0xe5, 0x82, 0x68, 0x99, 0x17, 0x55, 0xea, 0xa6, 0xca, 0xf5, 0x73,
	0xea, 0x88, 0x16, 0x1f, 0x8a, 0x73, 0x46, 0x34, 0xa3, 0x11, 0xd1, 0x4e, 0xc3, 0x03, 0x7e, 0x3d,
	0x6c, 0x5b, 0x65, 0xc7, 0x70, 0xa4, 0x84, 0x0b, 0xc1, 0x0c, 0xc7, 0x5a, 0xc9, 0x61, 0x95, 0x11,
	0x45, 0x18, 0x6e, 0xac, 0x6c, 0xc1, 0x35, 0x27, 0x09, 0xff, 0xc0, 0xa8, 0xd3, 0xf4, 0x80, 0xdf,
	0x0a, 0x51, 0x95, 0x3b, 0x73, 0xd0, 0x77, 0x00, 0xbb, 0x5a, 0x6a, 0x92, 0x44, 0xea, 0x20, 0xcb,
	0x92, 0xb9, 0xb3, 0xee, 0xd5, 0xfd, 0xce, 0xf6, 0x66, 0x60, 0xbb, 0x59, 0xb4, 0x3e, 0xb0, 0xad,
	0x0f, 0x76, 0x25, 0x17, 0xc3, 0xb7, 0x47, 0x27, 0xfd, 0xda, 0xef, 0x93, 0xfe, 0xa3, 0x29, 0xd7,
	0x7b, 0x07, 0x93, 0x20, 0x96, 0xa9, 0x6d, 0xbd, 0xfd, 0x7b, 0xaa, 0xe8, 0x3e, 0xd6, 0xf3, 0x8c,
	0x29, 0xb3, 0xe0, 0xdb, 0xcf, 0xbe, 0xff, 0x9f, 0x51, 0xf5, 0xf9, 0xf4, 0x70, 0xab, 0x9b, 0xb0,
	0x29, 0x89, 0xe7, 0x51, 0x71, 0xce, 0xea, 0xeb, 0xe9, 0xe1, 0x16, 0x08, 0x3b, 0x86, 0x70, 0x6c,
	0x00, 0xd1, 0x27, 0x00, 0x6f, 0xd9, 0xb3, 0x62, 0x34, 0x52, 0x7b, 0x24, 0x67, 0xca, 0x69, 0xfd,
	0x8b, 0xfa, 0xc5, 0xd5, 0xa9, 0x9b, 0x63, 0x53, 0xfb, 0x2f, 0x4c, 0x37, 0x57, 0x0c, 0x65, 0x0c,
	0xbd, 0x86, 0x1b, 0x97, 0x5c, 0x25, 0xa7, 0xed, 0x01, 0xbf, 0xb3, 0xfd, 0x30, 0xb8, 0x78, 0xe3,
	0x83, 0x71, 0x19, 0xdf, 0x39, 0x4b, 0x0f, 0x1b, 0x05, 0x66, 0x88, 0xd4, 0x05, 0x07, 0xf5, 0x60,
	0x2b, 0x65, 0x9a, 0x50, 0xa2, 0x89, 0x03, 0x3d, 0xe0, 0x77, 0xc3, 0xd5, 0x7c, 0xb8, 0x7b, 0xb4,
	0x70, 0xc1, 0xf1, 0xc2, 0x05, 0xbf, 0x16, 0x2e, 0xf8, 0xb8, 0x74, 0x6b, 0xc7, 0x4b, 0xb7, 0xf6,
	0x63, 0xe9, 0xd6, 0x5e, 0x3d, 0x3e, 0xb7, 0x5d, 0x3e, 0x15, 0x5c, 0x33, 0x5c, 0x3d, 0x81, 0xf7,
	0xab, 0x47, 0x60, 0x76, 0x3d, 0x69, 0x9a, 0x4b, 0xff, 0xec, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0x31, 0xdd, 0x8c, 0xbe, 0x03, 0x00, 0x00,
}

func (m *Project) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Project) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Project) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x52
	}
	{
		size, err := m.SpecialAllocations.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProject(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.AllocatedShares) > 0 {
		for iNdEx := len(m.AllocatedShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllocatedShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.TotalSupply) > 0 {
		for iNdEx := len(m.TotalSupply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalSupply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.MainnetInitialized {
		i--
		if m.MainnetInitialized {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.MainnetId != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.MainnetId))
		i--
		dAtA[i] = 0x28
	}
	if m.CreatedAt != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if m.CoordinatorId != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.CoordinatorId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProjectName) > 0 {
		i -= len(m.ProjectName)
		copy(dAtA[i:], m.ProjectName)
		i = encodeVarintProject(dAtA, i, uint64(len(m.ProjectName)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProjectId != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProject(dAtA []byte, offset int, v uint64) int {
	offset -= sovProject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Project) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectId != 0 {
		n += 1 + sovProject(uint64(m.ProjectId))
	}
	l = len(m.ProjectName)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.CoordinatorId != 0 {
		n += 1 + sovProject(uint64(m.CoordinatorId))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovProject(uint64(m.CreatedAt))
	}
	if m.MainnetId != 0 {
		n += 1 + sovProject(uint64(m.MainnetId))
	}
	if m.MainnetInitialized {
		n += 2
	}
	if len(m.TotalSupply) > 0 {
		for _, e := range m.TotalSupply {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if len(m.AllocatedShares) > 0 {
		for _, e := range m.AllocatedShares {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	l = m.SpecialAllocations.Size()
	n += 1 + l + sovProject(uint64(l))
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	return n
}

func sovProject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProject(x uint64) (n int) {
	return sovProject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Project) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
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
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorId", wireType)
			}
			m.CoordinatorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainnetId", wireType)
			}
			m.MainnetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainnetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainnetInitialized", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
			m.MainnetInitialized = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalSupply = append(m.TotalSupply, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.TotalSupply[len(m.TotalSupply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllocatedShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllocatedShares = append(m.AllocatedShares, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.AllocatedShares[len(m.AllocatedShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpecialAllocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpecialAllocations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata[:0], dAtA[iNdEx:postIndex]...)
			if m.Metadata == nil {
				m.Metadata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
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
func skipProject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProject
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
					return 0, ErrIntOverflowProject
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
					return 0, ErrIntOverflowProject
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
				return 0, ErrInvalidLengthProject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProject = fmt.Errorf("proto: unexpected end of group")
)
