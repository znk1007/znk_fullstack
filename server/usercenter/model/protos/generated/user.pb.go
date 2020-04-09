// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type Permission int32

const (
	Permission_super   Permission = 0
	Permission_admin   Permission = 1
	Permission_user    Permission = 2
	Permission_visitor Permission = 3
)

var Permission_name = map[int32]string{
	0: "super",
	1: "admin",
	2: "user",
	3: "visitor",
}

var Permission_value = map[string]int32{
	"super":   0,
	"admin":   1,
	"user":    2,
	"visitor": 3,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}

func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

type User struct {
	UserID               string     `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Account              string     `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Nickname             string     `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Phone                string     `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Photo                string     `protobuf:"bytes,6,opt,name=photo,proto3" json:"photo,omitempty"`
	CreatedAt            string     `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string     `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Active               int32      `protobuf:"varint,9,opt,name=active,proto3" json:"active,omitempty"`
	Permission           Permission `protobuf:"varint,11,opt,name=permission,proto3,enum=user.Permission" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *User) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *User) GetActive() int32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *User) GetPermission() Permission {
	if m != nil {
		return m.Permission
	}
	return Permission_super
}

func init() {
	proto.RegisterEnum("user.Permission", Permission_name, Permission_value)
	proto.RegisterType((*User)(nil), "user.User")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4b, 0xc5, 0x30,
	0x10, 0x80, 0x6d, 0x5f, 0x7f, 0x5e, 0x41, 0xca, 0x21, 0x72, 0x88, 0x43, 0x71, 0x2a, 0x0e, 0x0f,
	0xd1, 0xc9, 0x51, 0x71, 0x71, 0x93, 0x82, 0x8b, 0x5b, 0x6c, 0x0f, 0x0c, 0xda, 0xa4, 0x24, 0xe9,
	0xfb, 0x6b, 0xfc, 0x63, 0x25, 0x69, 0xed, 0x7b, 0x53, 0xee, 0xfb, 0xbe, 0x5b, 0x72, 0x00, 0xb3,
	0x65, 0xb3, 0x9f, 0x8c, 0x76, 0x1a, 0x13, 0x3f, 0xdf, 0xfc, 0xc6, 0x90, 0xbc, 0x5b, 0x36, 0x78,
	0x09, 0x99, 0x17, 0xaf, 0x2f, 0x14, 0x35, 0x51, 0x5b, 0x76, 0x2b, 0x21, 0x41, 0x2e, 0xfa, 0x5e,
	0xcf, 0xca, 0x51, 0x1c, 0xc2, 0x3f, 0xe2, 0x15, 0x14, 0x4a, 0xf6, 0xdf, 0x4a, 0x8c, 0x4c, 0xbb,
	0x90, 0x36, 0xc6, 0x0b, 0x48, 0xa7, 0x2f, 0xad, 0x98, 0x92, 0x10, 0x16, 0xf0, 0x96, 0x47, 0x21,
	0x7f, 0x28, 0x5d, 0x6c, 0x80, 0x75, 0xd7, 0x69, 0xca, 0xb6, 0x5d, 0xa7, 0xf1, 0x1a, 0xca, 0xde,
	0xb0, 0x70, 0x3c, 0x3c, 0x39, 0xca, 0x43, 0x39, 0x0a, 0x5f, 0xe7, 0x69, 0x58, 0x6b, 0xb1, 0xd4,
	0x4d, 0xf8, 0xbf, 0x88, 0xde, 0xc9, 0x03, 0x53, 0xd9, 0x44, 0x6d, 0xda, 0xad, 0x84, 0x77, 0x00,
	0x13, 0x9b, 0x51, 0x5a, 0x2b, 0xb5, 0xa2, 0xaa, 0x89, 0xda, 0xf3, 0xfb, 0x7a, 0x1f, 0x6e, 0xf2,
	0xb6, 0xf9, 0xee, 0x64, 0xe7, 0xf6, 0x11, 0xe0, 0x58, 0xb0, 0x84, 0xd4, 0xce, 0x13, 0x9b, 0xfa,
	0xcc, 0x8f, 0x62, 0x18, 0xa5, 0xaa, 0x23, 0x2c, 0x20, 0x9c, 0xb2, 0x8e, 0xb1, 0x82, 0xfc, 0x20,
	0xad, 0x74, 0xda, 0xd4, 0xbb, 0xe7, 0xea, 0xa3, 0xf4, 0x3a, 0x1c, 0xfb, 0x33, 0x0b, 0xcf, 0xc3,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x2c, 0x5e, 0xc9, 0x81, 0x01, 0x00, 0x00,
}
