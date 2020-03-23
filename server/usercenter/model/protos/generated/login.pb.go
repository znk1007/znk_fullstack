// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package userproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LoginReq struct {
	//账号
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	//登录签名jwt（JSON字符串，参数-用户ID：userID，密码：password[CBCEncrypt]，时间戳：timestamp，设备ID：deviceID，平台：platform[web,iOS,Android]，应用标识：appkey）
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginRes struct {
	//账号
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	//登录签名jwt（JSON字符串，参数-状态码：code，反馈消息：message，时间戳：timestamp，用户信息：user）
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "login.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "login.LoginRes")
}

func init() {
	proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f)
}

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xac, 0xb8, 0x38, 0x7c,
	0x40, 0x8c, 0xa0, 0xd4, 0x42, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0x3f, 0x3b,
	0x35, 0x4f, 0x82, 0x09, 0x2c, 0x0e, 0xe1, 0x20, 0xe9, 0x2d, 0x26, 0x55, 0xaf, 0x91, 0x29, 0x54,
	0x6f, 0x70, 0x51, 0x99, 0x90, 0x26, 0x17, 0xc4, 0x31, 0x42, 0xfc, 0x7a, 0x10, 0x17, 0xc2, 0x5c,
	0x24, 0x85, 0x2e, 0xe0, 0xc4, 0x1d, 0xc5, 0x59, 0x5a, 0x9c, 0x5a, 0x04, 0xf6, 0x42, 0x12, 0x1b,
	0x98, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x8c, 0xfb, 0xca, 0xd8, 0x00, 0x00, 0x00,
}
