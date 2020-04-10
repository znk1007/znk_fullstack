// Code generated by protoc-gen-go. DO NOT EDIT.
// source: regist.proto

package userproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegistReq struct {
	//注册账号
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	//注册签名 (JSON字符串，参数密码：password，设备ID：deviceID，平台：platform[web,iOS,Android]，时间戳：timestamp，应用标识：appkey)
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistReq) Reset()         { *m = RegistReq{} }
func (m *RegistReq) String() string { return proto.CompactTextString(m) }
func (*RegistReq) ProtoMessage()    {}
func (*RegistReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33eb661bf4b357e0, []int{0}
}

func (m *RegistReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistReq.Unmarshal(m, b)
}
func (m *RegistReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistReq.Marshal(b, m, deterministic)
}
func (m *RegistReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistReq.Merge(m, src)
}
func (m *RegistReq) XXX_Size() int {
	return xxx_messageInfo_RegistReq.Size(m)
}
func (m *RegistReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistReq proto.InternalMessageInfo

func (m *RegistReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegistReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegistRes struct {
	//账号
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	//注册响应签名（JSON字符串，参数包括用户ID：userID，时间戳：timestamp，状态码：code，反馈消息：message）
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistRes) Reset()         { *m = RegistRes{} }
func (m *RegistRes) String() string { return proto.CompactTextString(m) }
func (*RegistRes) ProtoMessage()    {}
func (*RegistRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_33eb661bf4b357e0, []int{1}
}

func (m *RegistRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistRes.Unmarshal(m, b)
}
func (m *RegistRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistRes.Marshal(b, m, deterministic)
}
func (m *RegistRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistRes.Merge(m, src)
}
func (m *RegistRes) XXX_Size() int {
	return xxx_messageInfo_RegistRes.Size(m)
}
func (m *RegistRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistRes.DiscardUnknown(m)
}

var xxx_messageInfo_RegistRes proto.InternalMessageInfo

func (m *RegistRes) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegistRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistReq)(nil), "regist.RegistReq")
	proto.RegisterType((*RegistRes)(nil), "regist.RegistRes")
}

func init() {
	proto.RegisterFile("regist.proto", fileDescriptor_33eb661bf4b357e0)
}

var fileDescriptor_33eb661bf4b357e0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xac, 0xb9, 0x38,
	0x83, 0xc0, 0xac, 0xa0, 0xd4, 0x42, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc,
	0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0x3f,
	0x3b, 0x35, 0x4f, 0x82, 0x09, 0x2c, 0x0e, 0xe1, 0x20, 0x6b, 0x2e, 0x26, 0x55, 0xb3, 0x91, 0x23,
	0x4c, 0x73, 0x70, 0x51, 0x99, 0x90, 0x09, 0x17, 0x57, 0x69, 0x71, 0x6a, 0x51, 0x50, 0x6a, 0x66,
	0x7a, 0x71, 0x89, 0x90, 0xa0, 0x1e, 0xd4, 0xad, 0x70, 0xa7, 0x49, 0x61, 0x08, 0x15, 0x2b, 0x31,
	0x38, 0x71, 0x47, 0x71, 0x82, 0x74, 0x81, 0x7d, 0x94, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0xef, 0xf0, 0x48, 0xe8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegistSrvClient is the client API for RegistSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistSrvClient interface {
	//用户注册
	UserReigst(ctx context.Context, in *RegistReq, opts ...grpc.CallOption) (*RegistRes, error)
}

type registSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistSrvClient(cc grpc.ClientConnInterface) RegistSrvClient {
	return &registSrvClient{cc}
}

func (c *registSrvClient) UserReigst(ctx context.Context, in *RegistReq, opts ...grpc.CallOption) (*RegistRes, error) {
	out := new(RegistRes)
	err := c.cc.Invoke(ctx, "/regist.RegistSrv/userReigst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistSrvServer is the server API for RegistSrv service.
type RegistSrvServer interface {
	//用户注册
	UserReigst(context.Context, *RegistReq) (*RegistRes, error)
}

// UnimplementedRegistSrvServer can be embedded to have forward compatible implementations.
type UnimplementedRegistSrvServer struct {
}

func (*UnimplementedRegistSrvServer) UserReigst(ctx context.Context, req *RegistReq) (*RegistRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserReigst not implemented")
}

func RegisterRegistSrvServer(s *grpc.Server, srv RegistSrvServer) {
	s.RegisterService(&_RegistSrv_serviceDesc, srv)
}

func _RegistSrv_UserReigst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistSrvServer).UserReigst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regist.RegistSrv/UserReigst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistSrvServer).UserReigst(ctx, req.(*RegistReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "regist.RegistSrv",
	HandlerType: (*RegistSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "userReigst",
			Handler:    _RegistSrv_UserReigst_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regist.proto",
}
