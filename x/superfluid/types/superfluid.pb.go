// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/superfluid.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// SuperfluidAsset is a unit to manage superfluid assets.
type SuperfluidAsset struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// TODO: do we need to keep superfluid_staked_amount in this object?
	SuperfluidStakedAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=superfluid_staked_amount,json=superfluidStakedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"superfluid_staked_amount" yaml:"superfluid_staked_amount"`
	Enabled                bool                                   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	RiskAdjustRatio        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=risk_adjust_ratio,json=riskAdjustRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"risk_adjust_ratio" yaml:"risk_adjust_ratio"`
}

func (m *SuperfluidAsset) Reset()         { *m = SuperfluidAsset{} }
func (m *SuperfluidAsset) String() string { return proto.CompactTextString(m) }
func (*SuperfluidAsset) ProtoMessage()    {}
func (*SuperfluidAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}
func (m *SuperfluidAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidAsset.Merge(m, src)
}
func (m *SuperfluidAsset) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidAsset proto.InternalMessageInfo

func (m *SuperfluidAsset) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *SuperfluidAsset) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*SuperfluidAsset)(nil), "osmosis.superfluid.SuperfluidAsset")
}

func init() {
	proto.RegisterFile("osmosis/superfluid/superfluid.proto", fileDescriptor_79d3c29d82dbb734)
}

var fileDescriptor_79d3c29d82dbb734 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xb3, 0x40,
	0x14, 0xc5, 0xa1, 0xdf, 0xe7, 0x3f, 0x36, 0x8d, 0xa4, 0x31, 0xa4, 0x0b, 0x68, 0x30, 0x31, 0xdd,
	0x94, 0x49, 0x75, 0xe7, 0xae, 0x8d, 0x1b, 0x8d, 0x1b, 0xe9, 0xce, 0x0d, 0x19, 0x60, 0x8a, 0x58,
	0x86, 0x21, 0xdc, 0x99, 0xc6, 0x3e, 0x83, 0x1b, 0xdf, 0xca, 0x2e, 0xbb, 0x34, 0x2e, 0x88, 0x69,
	0xdf, 0xa0, 0x4f, 0x60, 0x18, 0xa8, 0x6d, 0x34, 0x2e, 0x5c, 0x31, 0xe7, 0xfe, 0x6e, 0xce, 0x39,
	0x61, 0x46, 0x3b, 0x65, 0x40, 0x19, 0xc4, 0x80, 0x40, 0x64, 0x24, 0x1f, 0x27, 0x22, 0x0e, 0x77,
	0x8e, 0x4e, 0x96, 0x33, 0xce, 0x74, 0xbd, 0x5e, 0x72, 0xb6, 0xa4, 0xdd, 0x8a, 0x58, 0xc4, 0x24,
	0x46, 0xe5, 0xa9, 0xda, 0x6c, 0x9b, 0x11, 0x63, 0x51, 0x42, 0x90, 0x54, 0xbe, 0x18, 0xa3, 0x50,
	0xe4, 0x98, 0xc7, 0x2c, 0xad, 0xb9, 0xf5, 0x9d, 0xf3, 0x98, 0x12, 0xe0, 0x98, 0x66, 0x1b, 0x83,
	0x40, 0x66, 0x21, 0x1f, 0x03, 0x41, 0xd3, 0xbe, 0x4f, 0x38, 0xee, 0xa3, 0x80, 0xc5, 0xb5, 0x81,
	0xfd, 0xda, 0xd0, 0x9a, 0xa3, 0xaf, 0x16, 0x03, 0x00, 0xc2, 0xf5, 0x96, 0xb6, 0x17, 0x92, 0x94,
	0x51, 0x43, 0xed, 0xa8, 0xdd, 0x23, 0xb7, 0x12, 0xfa, 0xb3, 0xaa, 0x19, 0xdb, 0xbe, 0x1e, 0x70,
	0x3c, 0x21, 0xa1, 0x87, 0x29, 0x13, 0x29, 0x37, 0x1a, 0xe5, 0xe6, 0xf0, 0x6e, 0x5e, 0x58, 0xca,
	0x7b, 0x61, 0x9d, 0x45, 0x31, 0x7f, 0x10, 0xbe, 0x13, 0x30, 0x8a, 0xea, 0xfc, 0xea, 0xd3, 0x83,
	0x70, 0x82, 0xf8, 0x2c, 0x23, 0xe0, 0x5c, 0xa7, 0x7c, 0x5d, 0x58, 0xd6, 0x0c, 0xd3, 0xe4, 0xd2,
	0xfe, 0xcd, 0xd7, 0x76, 0x4f, 0xb6, 0x68, 0x24, 0xc9, 0x40, 0x02, 0xdd, 0xd0, 0x0e, 0x48, 0x8a,
	0xfd, 0x84, 0x84, 0xc6, 0xbf, 0x8e, 0xda, 0x3d, 0x74, 0x37, 0x52, 0x9f, 0x6a, 0xc7, 0x79, 0x0c,
	0x13, 0x0f, 0x87, 0x8f, 0x02, 0xb8, 0x27, 0x7f, 0x97, 0xf1, 0x5f, 0xf6, 0xbb, 0xf9, 0x43, 0xbf,
	0x2b, 0x12, 0xac, 0x0b, 0xcb, 0xa8, 0xfa, 0xfd, 0x30, 0xb4, 0xdd, 0x66, 0x39, 0x1b, 0xc8, 0x91,
	0x5b, 0x4e, 0x86, 0xb7, 0xf3, 0xa5, 0xa9, 0x2e, 0x96, 0xa6, 0xfa, 0xb1, 0x34, 0xd5, 0x97, 0x95,
	0xa9, 0x2c, 0x56, 0xa6, 0xf2, 0xb6, 0x32, 0x95, 0xfb, 0xf3, 0x9d, 0xb8, 0xfa, 0xe6, 0x7b, 0x09,
	0xf6, 0x61, 0x23, 0xd0, 0xd3, 0xee, 0x6b, 0x91, 0xf1, 0xfe, 0xbe, 0xbc, 0x9e, 0x8b, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0x1d, 0x2a, 0xc2, 0x50, 0x02, 0x00, 0x00,
}

func (m *SuperfluidAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RiskAdjustRatio.Size()
		i -= size
		if _, err := m.RiskAdjustRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSuperfluid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.SuperfluidStakedAmount.Size()
		i -= size
		if _, err := m.SuperfluidStakedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSuperfluid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSuperfluid(dAtA []byte, offset int, v uint64) int {
	offset -= sovSuperfluid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SuperfluidAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	l = m.SuperfluidStakedAmount.Size()
	n += 1 + l + sovSuperfluid(uint64(l))
	if m.Enabled {
		n += 2
	}
	l = m.RiskAdjustRatio.Size()
	n += 1 + l + sovSuperfluid(uint64(l))
	return n
}

func sovSuperfluid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSuperfluid(x uint64) (n int) {
	return sovSuperfluid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SuperfluidAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperfluidStakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SuperfluidStakedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RiskAdjustRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RiskAdjustRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func skipSuperfluid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
				return 0, ErrInvalidLengthSuperfluid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSuperfluid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSuperfluid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSuperfluid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSuperfluid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSuperfluid = fmt.Errorf("proto: unexpected end of group")
)
