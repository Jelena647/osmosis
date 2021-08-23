// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AssetInfoRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *AssetInfoRequest) Reset()         { *m = AssetInfoRequest{} }
func (m *AssetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*AssetInfoRequest) ProtoMessage()    {}
func (*AssetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d9448e4ed3943f, []int{0}
}
func (m *AssetInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetInfoRequest.Merge(m, src)
}
func (m *AssetInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *AssetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssetInfoRequest proto.InternalMessageInfo

func (m *AssetInfoRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type AssetInfoResponse struct {
	Asset *SuperfluidAsset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (m *AssetInfoResponse) Reset()         { *m = AssetInfoResponse{} }
func (m *AssetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*AssetInfoResponse) ProtoMessage()    {}
func (*AssetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d9448e4ed3943f, []int{1}
}
func (m *AssetInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetInfoResponse.Merge(m, src)
}
func (m *AssetInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *AssetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AssetInfoResponse proto.InternalMessageInfo

func (m *AssetInfoResponse) GetAsset() *SuperfluidAsset {
	if m != nil {
		return m.Asset
	}
	return nil
}

type AllAssetsRequest struct {
}

func (m *AllAssetsRequest) Reset()         { *m = AllAssetsRequest{} }
func (m *AllAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*AllAssetsRequest) ProtoMessage()    {}
func (*AllAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d9448e4ed3943f, []int{2}
}
func (m *AllAssetsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllAssetsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllAssetsRequest.Merge(m, src)
}
func (m *AllAssetsRequest) XXX_Size() int {
	return m.Size()
}
func (m *AllAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllAssetsRequest proto.InternalMessageInfo

type AllAssetsResponse struct {
	Assets []*SuperfluidAsset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
}

func (m *AllAssetsResponse) Reset()         { *m = AllAssetsResponse{} }
func (m *AllAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*AllAssetsResponse) ProtoMessage()    {}
func (*AllAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d9448e4ed3943f, []int{3}
}
func (m *AllAssetsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllAssetsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllAssetsResponse.Merge(m, src)
}
func (m *AllAssetsResponse) XXX_Size() int {
	return m.Size()
}
func (m *AllAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllAssetsResponse proto.InternalMessageInfo

func (m *AllAssetsResponse) GetAssets() []*SuperfluidAsset {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*AssetInfoRequest)(nil), "osmosis.superfluid.AssetInfoRequest")
	proto.RegisterType((*AssetInfoResponse)(nil), "osmosis.superfluid.AssetInfoResponse")
	proto.RegisterType((*AllAssetsRequest)(nil), "osmosis.superfluid.AllAssetsRequest")
	proto.RegisterType((*AllAssetsResponse)(nil), "osmosis.superfluid.AllAssetsResponse")
}

func init() { proto.RegisterFile("osmosis/superfluid/query.proto", fileDescriptor_e3d9448e4ed3943f) }

var fileDescriptor_e3d9448e4ed3943f = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x6d, 0x9e, 0xf4, 0xc1, 0x1b, 0x37, 0xef, 0x0d, 0x6f, 0x51, 0x8a, 0x8c, 0x25, 0x55, 0xe9,
	0xc6, 0x8c, 0xad, 0x20, 0x88, 0xab, 0xba, 0x13, 0x44, 0xb4, 0xee, 0xdc, 0x94, 0x49, 0x3b, 0x89,
	0x03, 0xc9, 0xdc, 0x34, 0x77, 0x22, 0x16, 0x71, 0xe3, 0xca, 0x65, 0xc1, 0x6f, 0xf0, 0x5f, 0x5c,
	0x16, 0xdc, 0xb8, 0x94, 0xd6, 0x0f, 0x91, 0x4c, 0x26, 0x6d, 0x89, 0xad, 0xba, 0xbb, 0xb9, 0xe7,
	0xdc, 0x73, 0xcf, 0x3d, 0x13, 0xc2, 0x00, 0x53, 0x40, 0x85, 0x1c, 0x8b, 0x4c, 0xe6, 0x51, 0x52,
	0xa8, 0x39, 0x5f, 0x14, 0x32, 0x5f, 0x06, 0x59, 0x0e, 0x06, 0x28, 0x75, 0x78, 0xb0, 0xc7, 0xbb,
	0xd7, 0x31, 0xc4, 0x60, 0x61, 0x5e, 0x56, 0x15, 0xb3, 0xcb, 0x66, 0x96, 0xca, 0x43, 0x81, 0x92,
	0xbf, 0x1b, 0x86, 0xd2, 0x88, 0x21, 0x9f, 0x81, 0xd2, 0x0e, 0xbf, 0x15, 0x03, 0xc4, 0x89, 0xe4,
	0x22, 0x53, 0x5c, 0x68, 0x0d, 0x46, 0x18, 0x05, 0x1a, 0x1d, 0x7a, 0xdb, 0xa1, 0xf6, 0x2b, 0x2c,
	0x22, 0x6e, 0x54, 0x2a, 0xd1, 0x88, 0x34, 0xab, 0xe5, 0x9b, 0x84, 0x79, 0x91, 0x5b, 0x05, 0x87,
	0xf7, 0x8f, 0x1c, 0xb2, 0x2f, 0x2b, 0x92, 0x3f, 0x20, 0x97, 0x63, 0x44, 0x69, 0x9e, 0xe9, 0x08,
	0x26, 0x72, 0x51, 0x48, 0x34, 0xf4, 0x9a, 0xb4, 0xe7, 0x52, 0x43, 0xda, 0xf1, 0x7a, 0xde, 0xe0,
	0x62, 0x52, 0x7d, 0xf8, 0x2f, 0xc8, 0xd5, 0x01, 0x13, 0x33, 0xd0, 0x28, 0xe9, 0x63, 0xd2, 0x16,
	0x65, 0xd3, 0x52, 0x6f, 0x8e, 0xfa, 0xc1, 0x9f, 0xe1, 0x04, 0xaf, 0x77, 0xa5, 0x9d, 0x9f, 0x54,
	0x13, 0x3e, 0x25, 0x97, 0xe3, 0x24, 0xb1, 0x2d, 0x74, 0x9b, 0xfd, 0x97, 0xe4, 0xea, 0xa0, 0xe7,
	0x76, 0x3c, 0x21, 0xe7, 0x76, 0x02, 0x3b, 0x5e, 0xef, 0xc6, 0xff, 0x2e, 0x71, 0x23, 0xa3, 0xaf,
	0x67, 0xa4, 0xfd, 0xaa, 0x7c, 0x3d, 0xba, 0xf2, 0xc8, 0xc5, 0xee, 0x00, 0x7a, 0xe7, 0x98, 0x48,
	0x33, 0x89, 0xee, 0xdd, 0x7f, 0xb0, 0x2a, 0x87, 0xfe, 0xa3, 0x4f, 0xdf, 0x7f, 0x7d, 0x39, 0x7b,
	0x40, 0x03, 0x7e, 0x24, 0xf2, 0xfa, 0xe1, 0xad, 0xa1, 0xa9, 0xd2, 0x11, 0xf0, 0x0f, 0x36, 0xd1,
	0x8f, 0xf4, 0x73, 0x69, 0xa9, 0xbe, 0xf7, 0x84, 0xa5, 0x46, 0x44, 0x27, 0x2c, 0x35, 0x43, 0xf3,
	0x03, 0x6b, 0x69, 0x40, 0xef, 0xfd, 0xd5, 0x52, 0x92, 0x4c, 0xab, 0x9c, 0x9e, 0x3e, 0xff, 0xb6,
	0x61, 0xde, 0x7a, 0xc3, 0xbc, 0x9f, 0x1b, 0xe6, 0xad, 0xb6, 0xac, 0xb5, 0xde, 0xb2, 0xd6, 0x8f,
	0x2d, 0x6b, 0xbd, 0x19, 0xc5, 0xca, 0xbc, 0x2d, 0xc2, 0x60, 0x06, 0x69, 0xad, 0x75, 0x3f, 0x11,
	0x21, 0xee, 0x84, 0xdf, 0x1f, 0x4a, 0x9b, 0x65, 0x26, 0x31, 0x3c, 0xb7, 0x3f, 0xd7, 0xc3, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xbe, 0x1a, 0x0e, 0x4c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns superfluid asset info
	AssetInfo(ctx context.Context, in *AssetInfoRequest, opts ...grpc.CallOption) (*AssetInfoResponse, error)
	// Returns all superfluid assets info
	AllAssets(ctx context.Context, in *AllAssetsRequest, opts ...grpc.CallOption) (*AllAssetsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AssetInfo(ctx context.Context, in *AssetInfoRequest, opts ...grpc.CallOption) (*AssetInfoResponse, error) {
	out := new(AssetInfoResponse)
	err := c.cc.Invoke(ctx, "/osmosis.superfluid.Query/AssetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllAssets(ctx context.Context, in *AllAssetsRequest, opts ...grpc.CallOption) (*AllAssetsResponse, error) {
	out := new(AllAssetsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.superfluid.Query/AllAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns superfluid asset info
	AssetInfo(context.Context, *AssetInfoRequest) (*AssetInfoResponse, error)
	// Returns all superfluid assets info
	AllAssets(context.Context, *AllAssetsRequest) (*AllAssetsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AssetInfo(ctx context.Context, req *AssetInfoRequest) (*AssetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssetInfo not implemented")
}
func (*UnimplementedQueryServer) AllAssets(ctx context.Context, req *AllAssetsRequest) (*AllAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllAssets not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AssetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AssetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.superfluid.Query/AssetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AssetInfo(ctx, req.(*AssetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.superfluid.Query/AllAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllAssets(ctx, req.(*AllAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.superfluid.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssetInfo",
			Handler:    _Query_AssetInfo_Handler,
		},
		{
			MethodName: "AllAssets",
			Handler:    _Query_AllAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/superfluid/query.proto",
}

func (m *AssetInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AssetInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Asset != nil {
		{
			size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllAssetsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllAssetsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllAssetsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *AllAssetsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllAssetsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllAssetsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AssetInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AssetInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Asset != nil {
		l = m.Asset.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AllAssetsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *AllAssetsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssetInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: AssetInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *AssetInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: AssetInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Asset == nil {
				m.Asset = &SuperfluidAsset{}
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *AllAssetsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: AllAssetsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllAssetsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *AllAssetsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: AllAssetsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllAssetsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, &SuperfluidAsset{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
