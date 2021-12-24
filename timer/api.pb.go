// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
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

type ErrCode int32

const (
	// 成功
	ErrCode_Success ErrCode = 0
	// 参数为空
	ErrCode_ParameterNull ErrCode = 1
	// 逻辑错误
	ErrCode_LogicError ErrCode = 2
)

var ErrCode_name = map[int32]string{
	0: "Success",
	1: "ParameterNull",
	2: "LogicError",
}

var ErrCode_value = map[string]int32{
	"Success":       0,
	"ParameterNull": 1,
	"LogicError":    2,
}

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (x *ErrCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrCode_value, data, "ErrCode")
	if err != nil {
		return err
	}
	*x = ErrCode(value)
	return nil
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type GetAllTimerReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllTimerReq) Reset()         { *m = GetAllTimerReq{} }
func (m *GetAllTimerReq) String() string { return proto.CompactTextString(m) }
func (*GetAllTimerReq) ProtoMessage()    {}
func (*GetAllTimerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *GetAllTimerReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllTimerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllTimerReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllTimerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTimerReq.Merge(m, src)
}
func (m *GetAllTimerReq) XXX_Size() int {
	return m.Size()
}
func (m *GetAllTimerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTimerReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTimerReq proto.InternalMessageInfo

type GetAllTimerRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllTimerRsp) Reset()         { *m = GetAllTimerRsp{} }
func (m *GetAllTimerRsp) String() string { return proto.CompactTextString(m) }
func (*GetAllTimerRsp) ProtoMessage()    {}
func (*GetAllTimerRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *GetAllTimerRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAllTimerRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAllTimerRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAllTimerRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTimerRsp.Merge(m, src)
}
func (m *GetAllTimerRsp) XXX_Size() int {
	return m.Size()
}
func (m *GetAllTimerRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTimerRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTimerRsp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("api.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*GetAllTimerReq)(nil), "api.GetAllTimerReq")
	proto.RegisterType((*GetAllTimerRsp)(nil), "api.GetAllTimerRsp")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0x12, 0xe0, 0xe2, 0x73, 0x4f, 0x2d, 0x71, 0xcc, 0xc9, 0x09, 0xc9,
	0xcc, 0x4d, 0x2d, 0x0a, 0x4a, 0x2d, 0x44, 0x17, 0x29, 0x2e, 0xd0, 0xb2, 0xe4, 0x62, 0x77, 0x2d,
	0x2a, 0x72, 0xce, 0x4f, 0x49, 0x15, 0xe2, 0xe6, 0x62, 0x0f, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e,
	0x16, 0x60, 0x10, 0x12, 0xe4, 0xe2, 0x0d, 0x48, 0x2c, 0x4a, 0xcc, 0x4d, 0x2d, 0x49, 0x2d, 0xf2,
	0x2b, 0xcd, 0xc9, 0x11, 0x60, 0x14, 0xe2, 0xe3, 0xe2, 0xf2, 0xc9, 0x4f, 0xcf, 0x4c, 0x76, 0x2d,
	0x2a, 0xca, 0x2f, 0x12, 0x60, 0x32, 0x8a, 0xe6, 0x62, 0x05, 0x1b, 0x23, 0x14, 0xc4, 0xc5, 0x8d,
	0x64, 0xaa, 0x90, 0xb0, 0x1e, 0xc8, 0x95, 0xa8, 0x36, 0x4b, 0x61, 0x0a, 0x16, 0x17, 0x28, 0x49,
	0x36, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x58, 0x49, 0xb0, 0x04, 0x24, 0xa2, 0x8f, 0x24, 0xeb, 0xc4,
	0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x02, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x93, 0x28, 0x11, 0x06, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimerClient is the client API for Timer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimerClient interface {
	// 获取所有定时器
	GetAllTimer(ctx context.Context, in *GetAllTimerReq, opts ...grpc.CallOption) (*GetAllTimerRsp, error)
}

type timerClient struct {
	cc *grpc.ClientConn
}

func NewTimerClient(cc *grpc.ClientConn) TimerClient {
	return &timerClient{cc}
}

func (c *timerClient) GetAllTimer(ctx context.Context, in *GetAllTimerReq, opts ...grpc.CallOption) (*GetAllTimerRsp, error) {
	out := new(GetAllTimerRsp)
	err := c.cc.Invoke(ctx, "/api.Timer/GetAllTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimerServer is the server API for Timer service.
type TimerServer interface {
	// 获取所有定时器
	GetAllTimer(context.Context, *GetAllTimerReq) (*GetAllTimerRsp, error)
}

// UnimplementedTimerServer can be embedded to have forward compatible implementations.
type UnimplementedTimerServer struct {
}

func (*UnimplementedTimerServer) GetAllTimer(ctx context.Context, req *GetAllTimerReq) (*GetAllTimerRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTimer not implemented")
}

func RegisterTimerServer(s *grpc.Server, srv TimerServer) {
	s.RegisterService(&_Timer_serviceDesc, srv)
}

func _Timer_GetAllTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTimerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).GetAllTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Timer/GetAllTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).GetAllTimer(ctx, req.(*GetAllTimerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Timer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Timer",
	HandlerType: (*TimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTimer",
			Handler:    _Timer_GetAllTimer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *GetAllTimerReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllTimerReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllTimerReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *GetAllTimerRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAllTimerRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAllTimerRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *GetAllTimerReq) Size() (n int) {
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

func (m *GetAllTimerRsp) Size() (n int) {
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
func (m *GetAllTimerReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAllTimerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllTimerReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *GetAllTimerRsp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAllTimerRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAllTimerRsp: illegal tag %d (wire type %d)", fieldNum, wire)
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
