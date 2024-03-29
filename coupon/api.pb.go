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

type ReceiveStatus int32

const (
	// 成功
	ReceiveStatus_Success ReceiveStatus = 0
	// 活动未打开
	ReceiveStatus_ActNotOpen ReceiveStatus = 1
	// 未知用户
	ReceiveStatus_UserNotConfig ReceiveStatus = 2
	// 用户配置未打开
	ReceiveStatus_UserNotOpen ReceiveStatus = 3
	// 领取目标不存在
	ReceiveStatus_CouponNotExist ReceiveStatus = 4
	// 重复领取
	ReceiveStatus_ReceiveRepeat ReceiveStatus = 5
	// 未知错误
	ReceiveStatus_UnknownError ReceiveStatus = 6
)

var ReceiveStatus_name = map[int32]string{
	0: "Success",
	1: "ActNotOpen",
	2: "UserNotConfig",
	3: "UserNotOpen",
	4: "CouponNotExist",
	5: "ReceiveRepeat",
	6: "UnknownError",
}

var ReceiveStatus_value = map[string]int32{
	"Success":        0,
	"ActNotOpen":     1,
	"UserNotConfig":  2,
	"UserNotOpen":    3,
	"CouponNotExist": 4,
	"ReceiveRepeat":  5,
	"UnknownError":   6,
}

func (x ReceiveStatus) Enum() *ReceiveStatus {
	p := new(ReceiveStatus)
	*p = x
	return p
}

func (x ReceiveStatus) String() string {
	return proto.EnumName(ReceiveStatus_name, int32(x))
}

func (x *ReceiveStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReceiveStatus_value, data, "ReceiveStatus")
	if err != nil {
		return err
	}
	*x = ReceiveStatus(value)
	return nil
}

func (ReceiveStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type ReceiveAllReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveAllReq) Reset()         { *m = ReceiveAllReq{} }
func (m *ReceiveAllReq) String() string { return proto.CompactTextString(m) }
func (*ReceiveAllReq) ProtoMessage()    {}
func (*ReceiveAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *ReceiveAllReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiveAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiveAllReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceiveAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveAllReq.Merge(m, src)
}
func (m *ReceiveAllReq) XXX_Size() int {
	return m.Size()
}
func (m *ReceiveAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveAllReq proto.InternalMessageInfo

type ReceiveInfo struct {
	// id
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// 结果
	Status               *ReceiveStatus `protobuf:"varint,2,opt,name=status,enum=api.ReceiveStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReceiveInfo) Reset()         { *m = ReceiveInfo{} }
func (m *ReceiveInfo) String() string { return proto.CompactTextString(m) }
func (*ReceiveInfo) ProtoMessage()    {}
func (*ReceiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *ReceiveInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiveInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveInfo.Merge(m, src)
}
func (m *ReceiveInfo) XXX_Size() int {
	return m.Size()
}
func (m *ReceiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveInfo proto.InternalMessageInfo

func (m *ReceiveInfo) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ReceiveInfo) GetStatus() ReceiveStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ReceiveStatus_Success
}

type ReceiveRsp struct {
	// 领取结果
	ReceiveList          []*ReceiveInfo `protobuf:"bytes,1,rep,name=receiveList" json:"receiveList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReceiveRsp) Reset()         { *m = ReceiveRsp{} }
func (m *ReceiveRsp) String() string { return proto.CompactTextString(m) }
func (*ReceiveRsp) ProtoMessage()    {}
func (*ReceiveRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}
func (m *ReceiveRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiveRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiveRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceiveRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveRsp.Merge(m, src)
}
func (m *ReceiveRsp) XXX_Size() int {
	return m.Size()
}
func (m *ReceiveRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveRsp proto.InternalMessageInfo

func (m *ReceiveRsp) GetReceiveList() []*ReceiveInfo {
	if m != nil {
		return m.ReceiveList
	}
	return nil
}

type ReceiveByIDReq struct {
	// id
	Id                   []string `protobuf:"bytes,1,rep,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveByIDReq) Reset()         { *m = ReceiveByIDReq{} }
func (m *ReceiveByIDReq) String() string { return proto.CompactTextString(m) }
func (*ReceiveByIDReq) ProtoMessage()    {}
func (*ReceiveByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}
func (m *ReceiveByIDReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiveByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiveByIDReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReceiveByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveByIDReq.Merge(m, src)
}
func (m *ReceiveByIDReq) XXX_Size() int {
	return m.Size()
}
func (m *ReceiveByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveByIDReq proto.InternalMessageInfo

func (m *ReceiveByIDReq) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.ReceiveStatus", ReceiveStatus_name, ReceiveStatus_value)
	proto.RegisterType((*ReceiveAllReq)(nil), "api.ReceiveAllReq")
	proto.RegisterType((*ReceiveInfo)(nil), "api.ReceiveInfo")
	proto.RegisterType((*ReceiveRsp)(nil), "api.ReceiveRsp")
	proto.RegisterType((*ReceiveByIDReq)(nil), "api.ReceiveByIDReq")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4d, 0x8e, 0xd3, 0x30,
	0x14, 0xc6, 0x09, 0x14, 0xcd, 0xcb, 0x4c, 0x6a, 0x1e, 0x23, 0x54, 0x55, 0xa3, 0x2a, 0xca, 0xaa,
	0x9a, 0xc5, 0x54, 0xca, 0x09, 0xe8, 0x94, 0x2e, 0x2a, 0x41, 0x11, 0xae, 0xba, 0xae, 0x4c, 0xea,
	0x56, 0x16, 0x91, 0xed, 0xc6, 0x2e, 0xd0, 0x2d, 0x3b, 0xd6, 0xdc, 0x82, 0x93, 0xb0, 0x44, 0xe2,
	0x02, 0xa8, 0xe2, 0x20, 0x28, 0x89, 0x89, 0x02, 0x2c, 0xfd, 0xde, 0xf7, 0xf3, 0xbe, 0xcf, 0x70,
	0xc1, 0x8d, 0xbc, 0x33, 0xa5, 0x76, 0x1a, 0x43, 0x6e, 0xe4, 0xf0, 0x66, 0xaf, 0xf5, 0xbe, 0x10,
	0x13, 0x6e, 0xe4, 0x84, 0x2b, 0xa5, 0x1d, 0x77, 0x52, 0x2b, 0xdb, 0x40, 0xd2, 0x3e, 0x5c, 0x31,
	0x91, 0x0b, 0xf9, 0x5e, 0x4c, 0x8b, 0x82, 0x89, 0x43, 0xba, 0x80, 0xc8, 0x0f, 0x16, 0x6a, 0xa7,
	0x31, 0x86, 0x40, 0x6e, 0x07, 0x24, 0x21, 0xe3, 0x0b, 0x16, 0xc8, 0x2d, 0xde, 0x42, 0xcf, 0x3a,
	0xee, 0x8e, 0x76, 0x10, 0x24, 0x64, 0x1c, 0x67, 0x78, 0x57, 0xd9, 0x79, 0xc6, 0xaa, 0xde, 0x30,
	0x8f, 0x48, 0x9f, 0x03, 0xf8, 0x05, 0xb3, 0x06, 0x33, 0x88, 0xca, 0xe6, 0xf5, 0x52, 0x5a, 0x37,
	0x20, 0x49, 0x38, 0x8e, 0x32, 0xda, 0xa5, 0x57, 0x86, 0xac, 0x0b, 0x4a, 0x13, 0x88, 0xfd, 0xee,
	0xfe, 0xb4, 0x78, 0xc1, 0xc4, 0xa1, 0xbd, 0x27, 0x6c, 0xee, 0xb9, 0xfd, 0x4c, 0xda, 0x00, 0x8d,
	0x3b, 0x46, 0xf0, 0x78, 0x75, 0xcc, 0x73, 0x61, 0x2d, 0x7d, 0x80, 0x31, 0xc0, 0x34, 0x77, 0x4b,
	0xed, 0x5e, 0x1b, 0xa1, 0x28, 0xc1, 0x27, 0x70, 0xb5, 0xb6, 0xa2, 0x5c, 0x6a, 0x37, 0xd3, 0x6a,
	0x27, 0xf7, 0x34, 0xc0, 0x3e, 0x44, 0x7e, 0x54, 0x63, 0x42, 0x44, 0x88, 0x67, 0xfa, 0x68, 0xb4,
	0x5a, 0x6a, 0x37, 0xff, 0x28, 0xad, 0xa3, 0x0f, 0x2b, 0xde, 0x9f, 0x28, 0xc2, 0x08, 0xee, 0xe8,
	0x23, 0xa4, 0x70, 0xb9, 0x56, 0xef, 0x94, 0xfe, 0xa0, 0xe6, 0x65, 0xa9, 0x4b, 0xda, 0xcb, 0xbe,
	0x12, 0xe8, 0x35, 0x4c, 0x7c, 0xd5, 0x46, 0x9f, 0x16, 0x05, 0xfe, 0x55, 0x52, 0xd3, 0xf3, 0xb0,
	0xdf, 0x9d, 0x31, 0x6b, 0xd2, 0xe1, 0xa7, 0x1f, 0xbf, 0xbe, 0x04, 0xd7, 0x29, 0xe6, 0xb5, 0xc4,
	0xc4, 0xf7, 0xb0, 0xe1, 0x45, 0x81, 0x6f, 0xda, 0x4f, 0xa9, 0x7a, 0xc0, 0xa7, 0x5d, 0xae, 0x6f,
	0xe6, 0x7f, 0xc1, 0x9b, 0x5a, 0xf0, 0x59, 0x7a, 0xfd, 0x8f, 0xe0, 0xdb, 0xd3, 0x46, 0x6e, 0xef,
	0x2f, 0xbf, 0x9d, 0x47, 0xe4, 0xfb, 0x79, 0x44, 0x7e, 0x9e, 0x47, 0xe4, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x15, 0x75, 0xf3, 0xb7, 0x35, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CouponClient is the client API for Coupon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CouponClient interface {
	// 领取所有
	ReceiveAll(ctx context.Context, in *ReceiveAllReq, opts ...grpc.CallOption) (*ReceiveRsp, error)
	// 根据id领取
	ReceiveByID(ctx context.Context, in *ReceiveByIDReq, opts ...grpc.CallOption) (*ReceiveRsp, error)
}

type couponClient struct {
	cc *grpc.ClientConn
}

func NewCouponClient(cc *grpc.ClientConn) CouponClient {
	return &couponClient{cc}
}

func (c *couponClient) ReceiveAll(ctx context.Context, in *ReceiveAllReq, opts ...grpc.CallOption) (*ReceiveRsp, error) {
	out := new(ReceiveRsp)
	err := c.cc.Invoke(ctx, "/api.Coupon/ReceiveAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) ReceiveByID(ctx context.Context, in *ReceiveByIDReq, opts ...grpc.CallOption) (*ReceiveRsp, error) {
	out := new(ReceiveRsp)
	err := c.cc.Invoke(ctx, "/api.Coupon/ReceiveByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServer is the server API for Coupon service.
type CouponServer interface {
	// 领取所有
	ReceiveAll(context.Context, *ReceiveAllReq) (*ReceiveRsp, error)
	// 根据id领取
	ReceiveByID(context.Context, *ReceiveByIDReq) (*ReceiveRsp, error)
}

// UnimplementedCouponServer can be embedded to have forward compatible implementations.
type UnimplementedCouponServer struct {
}

func (*UnimplementedCouponServer) ReceiveAll(ctx context.Context, req *ReceiveAllReq) (*ReceiveRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveAll not implemented")
}
func (*UnimplementedCouponServer) ReceiveByID(ctx context.Context, req *ReceiveByIDReq) (*ReceiveRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveByID not implemented")
}

func RegisterCouponServer(s *grpc.Server, srv CouponServer) {
	s.RegisterService(&_Coupon_serviceDesc, srv)
}

func _Coupon_ReceiveAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).ReceiveAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Coupon/ReceiveAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).ReceiveAll(ctx, req.(*ReceiveAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_ReceiveByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).ReceiveByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Coupon/ReceiveByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).ReceiveByID(ctx, req.(*ReceiveByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coupon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Coupon",
	HandlerType: (*CouponServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveAll",
			Handler:    _Coupon_ReceiveAll_Handler,
		},
		{
			MethodName: "ReceiveByID",
			Handler:    _Coupon_ReceiveByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *ReceiveAllReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiveAllReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceiveAllReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *ReceiveInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiveInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceiveInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != nil {
		i = encodeVarintApi(dAtA, i, uint64(*m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != nil {
		i -= len(*m.Id)
		copy(dAtA[i:], *m.Id)
		i = encodeVarintApi(dAtA, i, uint64(len(*m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReceiveRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiveRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceiveRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ReceiveList) > 0 {
		for iNdEx := len(m.ReceiveList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReceiveList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReceiveByIDReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiveByIDReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReceiveByIDReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Id) > 0 {
		for iNdEx := len(m.Id) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Id[iNdEx])
			copy(dAtA[i:], m.Id[iNdEx])
			i = encodeVarintApi(dAtA, i, uint64(len(m.Id[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
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
func (m *ReceiveAllReq) Size() (n int) {
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

func (m *ReceiveInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = len(*m.Id)
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Status != nil {
		n += 1 + sovApi(uint64(*m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReceiveRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ReceiveList) > 0 {
		for _, e := range m.ReceiveList {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReceiveByIDReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Id) > 0 {
		for _, s := range m.Id {
			l = len(s)
			n += 1 + l + sovApi(uint64(l))
		}
	}
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
func (m *ReceiveAllReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReceiveAllReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiveAllReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ReceiveInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReceiveInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiveInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = &s
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v ReceiveStatus
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= ReceiveStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = &v
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
func (m *ReceiveRsp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReceiveRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiveRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiveList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiveList = append(m.ReceiveList, &ReceiveInfo{})
			if err := m.ReceiveList[len(m.ReceiveList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *ReceiveByIDReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ReceiveByIDReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiveByIDReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = append(m.Id, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
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
