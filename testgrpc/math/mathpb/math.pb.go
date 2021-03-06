// Code generated by protoc-gen-go. DO NOT EDIT.
// source: math.proto

package mathpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Numbers struct {
	Num1                 int32    `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2                 int32    `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()    {}
func (*Numbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_math_128713b35d8af79c, []int{0}
}
func (m *Numbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numbers.Unmarshal(m, b)
}
func (m *Numbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numbers.Marshal(b, m, deterministic)
}
func (dst *Numbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numbers.Merge(dst, src)
}
func (m *Numbers) XXX_Size() int {
	return xxx_messageInfo_Numbers.Size(m)
}
func (m *Numbers) XXX_DiscardUnknown() {
	xxx_messageInfo_Numbers.DiscardUnknown(m)
}

var xxx_messageInfo_Numbers proto.InternalMessageInfo

func (m *Numbers) GetNum1() int32 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *Numbers) GetNum2() int32 {
	if m != nil {
		return m.Num2
	}
	return 0
}

type NumbersRequest struct {
	Numbers              *Numbers `protobuf:"bytes,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumbersRequest) Reset()         { *m = NumbersRequest{} }
func (m *NumbersRequest) String() string { return proto.CompactTextString(m) }
func (*NumbersRequest) ProtoMessage()    {}
func (*NumbersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_math_128713b35d8af79c, []int{1}
}
func (m *NumbersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumbersRequest.Unmarshal(m, b)
}
func (m *NumbersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumbersRequest.Marshal(b, m, deterministic)
}
func (dst *NumbersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumbersRequest.Merge(dst, src)
}
func (m *NumbersRequest) XXX_Size() int {
	return xxx_messageInfo_NumbersRequest.Size(m)
}
func (m *NumbersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumbersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumbersRequest proto.InternalMessageInfo

func (m *NumbersRequest) GetNumbers() *Numbers {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type NumbersResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumbersResponse) Reset()         { *m = NumbersResponse{} }
func (m *NumbersResponse) String() string { return proto.CompactTextString(m) }
func (*NumbersResponse) ProtoMessage()    {}
func (*NumbersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_math_128713b35d8af79c, []int{2}
}
func (m *NumbersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumbersResponse.Unmarshal(m, b)
}
func (m *NumbersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumbersResponse.Marshal(b, m, deterministic)
}
func (dst *NumbersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumbersResponse.Merge(dst, src)
}
func (m *NumbersResponse) XXX_Size() int {
	return xxx_messageInfo_NumbersResponse.Size(m)
}
func (m *NumbersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NumbersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NumbersResponse proto.InternalMessageInfo

func (m *NumbersResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "math.Numbers")
	proto.RegisterType((*NumbersRequest)(nil), "math.NumbersRequest")
	proto.RegisterType((*NumbersResponse)(nil), "math.NumbersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NumbersServiceClient is the client API for NumbersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NumbersServiceClient interface {
	// Unary
	Compute(ctx context.Context, in *NumbersRequest, opts ...grpc.CallOption) (*NumbersResponse, error)
}

type numbersServiceClient struct {
	cc *grpc.ClientConn
}

func NewNumbersServiceClient(cc *grpc.ClientConn) NumbersServiceClient {
	return &numbersServiceClient{cc}
}

func (c *numbersServiceClient) Compute(ctx context.Context, in *NumbersRequest, opts ...grpc.CallOption) (*NumbersResponse, error) {
	out := new(NumbersResponse)
	err := c.cc.Invoke(ctx, "/math.NumbersService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumbersServiceServer is the server API for NumbersService service.
type NumbersServiceServer interface {
	// Unary
	Compute(context.Context, *NumbersRequest) (*NumbersResponse, error)
}

func RegisterNumbersServiceServer(s *grpc.Server, srv NumbersServiceServer) {
	s.RegisterService(&_NumbersService_serviceDesc, srv)
}

func _NumbersService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumbersServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/math.NumbersService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumbersServiceServer).Compute(ctx, req.(*NumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NumbersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "math.NumbersService",
	HandlerType: (*NumbersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _NumbersService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "math.proto",
}

func init() { proto.RegisterFile("math.proto", fileDescriptor_math_128713b35d8af79c) }

var fileDescriptor_math_128713b35d8af79c = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x0c, 0xb9, 0xd8, 0xfd, 0x4a,
	0x73, 0x93, 0x52, 0x8b, 0x8a, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x4a, 0x73, 0x0d, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0xa8, 0x98, 0x91, 0x04, 0x13, 0x5c, 0xcc, 0x48, 0xc9, 0x92,
	0x8b, 0x0f, 0xaa, 0x25, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x9d, 0x8b, 0x3d, 0x0f,
	0x22, 0x02, 0xd6, 0xcc, 0x6d, 0xc4, 0xab, 0x07, 0xb6, 0x08, 0xa6, 0x0c, 0x26, 0xab, 0xa4, 0xc9,
	0xc5, 0x0f, 0xd7, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a,
	0x5c, 0x9a, 0x53, 0x02, 0xb5, 0x17, 0xca, 0x33, 0xf2, 0x82, 0xdb, 0x12, 0x9c, 0x5a, 0x54, 0x96,
	0x99, 0x9c, 0x2a, 0x64, 0xc1, 0xc5, 0xee, 0x9c, 0x9f, 0x5b, 0x50, 0x5a, 0x92, 0x2a, 0x24, 0x82,
	0x6a, 0x3e, 0xc4, 0x19, 0x52, 0xa2, 0x68, 0xa2, 0x10, 0x1b, 0x94, 0x18, 0x9c, 0x38, 0xa2, 0xd8,
	0x40, 0x32, 0x05, 0x49, 0x49, 0x6c, 0x60, 0xbf, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x08,
	0x1d, 0x88, 0x6d, 0x09, 0x01, 0x00, 0x00,
}
