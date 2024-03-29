// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
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

// 通知类型 业务+通知平台
type AlarmType int32

const (
	// coupon钉钉通知
	AlarmType_CouponDingTalk AlarmType = 1
	// coupon企业微信通知
	AlarmType_CouponQyWeChat AlarmType = 2
)

var AlarmType_name = map[int32]string{
	1: "CouponDingTalk",
	2: "CouponQyWeChat",
}

var AlarmType_value = map[string]int32{
	"CouponDingTalk": 1,
	"CouponQyWeChat": 2,
}

func (x AlarmType) Enum() *AlarmType {
	p := new(AlarmType)
	*p = x
	return p
}

func (x AlarmType) String() string {
	return proto.EnumName(AlarmType_name, int32(x))
}

func (x *AlarmType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AlarmType_value, data, "AlarmType")
	if err != nil {
		return err
	}
	*x = AlarmType(value)
	return nil
}

func (AlarmType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type NotifyReq struct {
	// 类型
	AlarmType *AlarmType `protobuf:"varint,1,req,name=alarm_type,json=alarmType,enum=api.AlarmType" json:"alarm_type,omitempty"`
	// 标题
	Title *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// 内容
	Content              *string  `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyReq) Reset()         { *m = NotifyReq{} }
func (m *NotifyReq) String() string { return proto.CompactTextString(m) }
func (*NotifyReq) ProtoMessage()    {}
func (*NotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *NotifyReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReq.Merge(m, src)
}
func (m *NotifyReq) XXX_Size() int {
	return m.Size()
}
func (m *NotifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReq proto.InternalMessageInfo

func (m *NotifyReq) GetAlarmType() AlarmType {
	if m != nil && m.AlarmType != nil {
		return *m.AlarmType
	}
	return AlarmType_CouponDingTalk
}

func (m *NotifyReq) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *NotifyReq) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type NotifyRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyRsp) Reset()         { *m = NotifyRsp{} }
func (m *NotifyRsp) String() string { return proto.CompactTextString(m) }
func (*NotifyRsp) ProtoMessage()    {}
func (*NotifyRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *NotifyRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRsp.Merge(m, src)
}
func (m *NotifyRsp) XXX_Size() int {
	return m.Size()
}
func (m *NotifyRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyRsp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("api.AlarmType", AlarmType_name, AlarmType_value)
	proto.RegisterType((*NotifyReq)(nil), "api.NotifyReq")
	proto.RegisterType((*NotifyRsp)(nil), "api.NotifyRsp")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0xca, 0xe2, 0xe2, 0xf4, 0xcb, 0x2f, 0xc9, 0x4c, 0xab, 0x0c, 0x4a,
	0x2d, 0x14, 0xd2, 0xe5, 0xe2, 0x4a, 0xcc, 0x49, 0x2c, 0xca, 0x8d, 0x2f, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x54, 0x60, 0xd2, 0xe0, 0x33, 0xe2, 0xd3, 0x03, 0x99, 0xe7, 0x08, 0x12, 0x0e, 0xa9, 0x2c,
	0x48, 0x0d, 0xe2, 0x4c, 0x84, 0x31, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xfc, 0xbc, 0x92,
	0xd4, 0xbc, 0x12, 0x09, 0x66, 0x05, 0x26, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x89, 0x1b, 0x6e, 0x57,
	0x71, 0x81, 0x96, 0x31, 0x17, 0x27, 0xdc, 0x50, 0x21, 0x21, 0x2e, 0x3e, 0xe7, 0xfc, 0xd2, 0x82,
	0xfc, 0x3c, 0x97, 0xcc, 0xbc, 0xf4, 0x90, 0xc4, 0x9c, 0x6c, 0x01, 0x46, 0x84, 0x58, 0x60, 0x65,
	0x78, 0xaa, 0x73, 0x46, 0x62, 0x89, 0x00, 0x93, 0x91, 0x3b, 0x17, 0x2b, 0x58, 0x93, 0x90, 0x1d,
	0x17, 0x1b, 0xc4, 0x28, 0x21, 0x88, 0xfb, 0xe0, 0x7e, 0x90, 0x42, 0xe1, 0x17, 0x17, 0x28, 0x89,
	0x34, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x4f, 0x89, 0x07, 0xec, 0x6e, 0xfd, 0x3c, 0xb0, 0x84, 0x13,
	0xcf, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x08, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0x86, 0xc0, 0x77, 0x33, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlarmClient is the client API for Alarm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlarmClient interface {
	// 通知
	Notify(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyRsp, error)
}

type alarmClient struct {
	cc *grpc.ClientConn
}

func NewAlarmClient(cc *grpc.ClientConn) AlarmClient {
	return &alarmClient{cc}
}

func (c *alarmClient) Notify(ctx context.Context, in *NotifyReq, opts ...grpc.CallOption) (*NotifyRsp, error) {
	out := new(NotifyRsp)
	err := c.cc.Invoke(ctx, "/api.Alarm/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlarmServer is the server API for Alarm service.
type AlarmServer interface {
	// 通知
	Notify(context.Context, *NotifyReq) (*NotifyRsp, error)
}

// UnimplementedAlarmServer can be embedded to have forward compatible implementations.
type UnimplementedAlarmServer struct {
}

func (*UnimplementedAlarmServer) Notify(ctx context.Context, req *NotifyReq) (*NotifyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}

func RegisterAlarmServer(s *grpc.Server, srv AlarmServer) {
	s.RegisterService(&_Alarm_serviceDesc, srv)
}

func _Alarm_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Alarm/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServer).Notify(ctx, req.(*NotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alarm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Alarm",
	HandlerType: (*AlarmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Alarm_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *NotifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Content == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("content")
	} else {
		i -= len(*m.Content)
		copy(dAtA[i:], *m.Content)
		i = encodeVarintApi(dAtA, i, uint64(len(*m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Title != nil {
		i -= len(*m.Title)
		copy(dAtA[i:], *m.Title)
		i = encodeVarintApi(dAtA, i, uint64(len(*m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.AlarmType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("alarm_type")
	} else {
		i = encodeVarintApi(dAtA, i, uint64(*m.AlarmType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NotifyRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *NotifyReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AlarmType != nil {
		n += 1 + sovApi(uint64(*m.AlarmType))
	}
	if m.Title != nil {
		l = len(*m.Title)
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Content != nil {
		l = len(*m.Content)
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NotifyRsp) Size() (n int) {
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
func (m *NotifyReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NotifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlarmType", wireType)
			}
			var v AlarmType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= AlarmType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlarmType = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
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
			m.Content = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
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
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("alarm_type")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("content")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyRsp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NotifyRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyRsp: illegal tag %d (wire type %d)", fieldNum, wire)
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
