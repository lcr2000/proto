// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
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

type ReportReq struct {
	// 平台
	PlatformId *string `protobuf:"bytes,1,req,name=platform_id,json=platformId" json:"platform_id,omitempty"`
	// 活动
	ActId *string `protobuf:"bytes,2,req,name=act_id,json=actId" json:"act_id,omitempty"`
	// 统计项
	ItemId *string `protobuf:"bytes,3,req,name=item_id,json=itemId" json:"item_id,omitempty"`
	// 数值
	Amount               *float32 `protobuf:"fixed32,4,req,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReq) Reset()         { *m = ReportReq{} }
func (m *ReportReq) String() string { return proto.CompactTextString(m) }
func (*ReportReq) ProtoMessage()    {}
func (*ReportReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *ReportReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReportReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReportReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReportReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReq.Merge(m, src)
}
func (m *ReportReq) XXX_Size() int {
	return m.Size()
}
func (m *ReportReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReq proto.InternalMessageInfo

func (m *ReportReq) GetPlatformId() string {
	if m != nil && m.PlatformId != nil {
		return *m.PlatformId
	}
	return ""
}

func (m *ReportReq) GetActId() string {
	if m != nil && m.ActId != nil {
		return *m.ActId
	}
	return ""
}

func (m *ReportReq) GetItemId() string {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return ""
}

func (m *ReportReq) GetAmount() float32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

type ReportRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRsp) Reset()         { *m = ReportRsp{} }
func (m *ReportRsp) String() string { return proto.CompactTextString(m) }
func (*ReportRsp) ProtoMessage()    {}
func (*ReportRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *ReportRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReportRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReportRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReportRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRsp.Merge(m, src)
}
func (m *ReportRsp) XXX_Size() int {
	return m.Size()
}
func (m *ReportRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRsp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReportReq)(nil), "api.ReportReq")
	proto.RegisterType((*ReportRsp)(nil), "api.ReportRsp")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0x2a, 0xe1, 0xe2, 0x0c, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x09, 0x4a,
	0x2d, 0x14, 0x92, 0xe7, 0xe2, 0x2e, 0xc8, 0x49, 0x2c, 0x49, 0xcb, 0x2f, 0xca, 0x8d, 0xcf, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0xe2, 0x82, 0x09, 0x79, 0xa6, 0x08, 0x89, 0x72, 0xb1,
	0x25, 0x26, 0x97, 0x80, 0xe4, 0x98, 0xc0, 0x72, 0xac, 0x89, 0xc9, 0x25, 0x9e, 0x29, 0x42, 0xe2,
	0x5c, 0xec, 0x99, 0x25, 0xa9, 0x60, 0x3d, 0xcc, 0x60, 0x71, 0x36, 0x10, 0xd7, 0x33, 0x45, 0x48,
	0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x45, 0x81, 0x49, 0x83, 0x29, 0x08,
	0xca, 0x53, 0xe2, 0x86, 0xdb, 0x5a, 0x5c, 0x60, 0x14, 0xc8, 0xc5, 0x15, 0x0c, 0x72, 0x54, 0x71,
	0x49, 0x66, 0x72, 0xb1, 0x90, 0x33, 0x17, 0x1b, 0x44, 0x4a, 0x88, 0x4f, 0x0f, 0xe4, 0x13, 0xb8,
	0xeb, 0xa4, 0x50, 0xf8, 0xc5, 0x05, 0x4a, 0x92, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12, 0x56, 0x12,
	0x2c, 0x86, 0xeb, 0xd6, 0x2f, 0x02, 0xcb, 0x3a, 0xf1, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x9a, 0xa6, 0xee,
	0x12, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatisticsClient is the client API for Statistics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatisticsClient interface {
	// 上报
	Report(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*ReportRsp, error)
}

type statisticsClient struct {
	cc *grpc.ClientConn
}

func NewStatisticsClient(cc *grpc.ClientConn) StatisticsClient {
	return &statisticsClient{cc}
}

func (c *statisticsClient) Report(ctx context.Context, in *ReportReq, opts ...grpc.CallOption) (*ReportRsp, error) {
	out := new(ReportRsp)
	err := c.cc.Invoke(ctx, "/api.Statistics/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsServer is the server API for Statistics service.
type StatisticsServer interface {
	// 上报
	Report(context.Context, *ReportReq) (*ReportRsp, error)
}

// UnimplementedStatisticsServer can be embedded to have forward compatible implementations.
type UnimplementedStatisticsServer struct {
}

func (*UnimplementedStatisticsServer) Report(ctx context.Context, req *ReportReq) (*ReportRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterStatisticsServer(s *grpc.Server, srv StatisticsServer) {
	s.RegisterService(&_Statistics_serviceDesc, srv)
}

func _Statistics_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticsServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Statistics/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticsServer).Report(ctx, req.(*ReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Statistics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Statistics",
	HandlerType: (*StatisticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _Statistics_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *ReportReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReportReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReportReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Amount == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("amount")
	} else {
		i -= 4
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(*m.Amount))))
		i--
		dAtA[i] = 0x25
	}
	if m.ItemId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("item_id")
	} else {
		i -= len(*m.ItemId)
		copy(dAtA[i:], *m.ItemId)
		i = encodeVarintApi(dAtA, i, uint64(len(*m.ItemId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ActId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("act_id")
	} else {
		i -= len(*m.ActId)
		copy(dAtA[i:], *m.ActId)
		i = encodeVarintApi(dAtA, i, uint64(len(*m.ActId)))
		i--
		dAtA[i] = 0x12
	}
	if m.PlatformId == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("platform_id")
	} else {
		i -= len(*m.PlatformId)
		copy(dAtA[i:], *m.PlatformId)
		i = encodeVarintApi(dAtA, i, uint64(len(*m.PlatformId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReportRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReportRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReportRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReportReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlatformId != nil {
		l = len(*m.PlatformId)
		n += 1 + l + sovApi(uint64(l))
	}
	if m.ActId != nil {
		l = len(*m.ActId)
		n += 1 + l + sovApi(uint64(l))
	}
	if m.ItemId != nil {
		l = len(*m.ItemId)
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Amount != nil {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReportRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReportReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ReportReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReportReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlatformId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.PlatformId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ActId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ItemId = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			v2 := float32(math.Float32frombits(v))
			m.Amount = &v2
			hasFields[0] |= uint64(0x00000008)
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("platform_id")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("act_id")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("item_id")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("amount")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReportRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ReportRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReportRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
