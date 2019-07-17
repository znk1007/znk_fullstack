// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: updatephoto.proto

package protos

import (
	context "context"
	fmt "fmt"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UpdatePhotoRequest struct {
	Account   string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	SessionId string `protobuf:"bytes,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Photo     string `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo,omitempty"`
	Device    string `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
}

func (m *UpdatePhotoRequest) Reset()         { *m = UpdatePhotoRequest{} }
func (m *UpdatePhotoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePhotoRequest) ProtoMessage()    {}
func (*UpdatePhotoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e81b693890f6e8, []int{0}
}
func (m *UpdatePhotoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePhotoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePhotoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePhotoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePhotoRequest.Merge(m, src)
}
func (m *UpdatePhotoRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePhotoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePhotoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePhotoRequest proto.InternalMessageInfo

func (m *UpdatePhotoRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UpdatePhotoRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdatePhotoRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *UpdatePhotoRequest) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

func (m *UpdatePhotoRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type UpdatePhotoResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Status  int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *UpdatePhotoResponse) Reset()         { *m = UpdatePhotoResponse{} }
func (m *UpdatePhotoResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePhotoResponse) ProtoMessage()    {}
func (*UpdatePhotoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96e81b693890f6e8, []int{1}
}
func (m *UpdatePhotoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePhotoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePhotoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePhotoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePhotoResponse.Merge(m, src)
}
func (m *UpdatePhotoResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePhotoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePhotoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePhotoResponse proto.InternalMessageInfo

func (m *UpdatePhotoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UpdatePhotoResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdatePhotoResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*UpdatePhotoRequest)(nil), "protos.updatephoto.UpdatePhotoRequest")
	proto.RegisterType((*UpdatePhotoResponse)(nil), "protos.updatephoto.UpdatePhotoResponse")
}

func init() { proto.RegisterFile("updatephoto.proto", fileDescriptor_96e81b693890f6e8) }

var fileDescriptor_96e81b693890f6e8 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x68, 0x82, 0x7a, 0x4c, 0x1c, 0x08, 0x59, 0x08, 0x59, 0x55, 0x07, 0x60, 0xca,
	0x00, 0x6f, 0xc0, 0xd6, 0x0d, 0x45, 0x62, 0x00, 0xa6, 0x10, 0x9f, 0x68, 0x07, 0xe2, 0xd0, 0xb3,
	0x79, 0x0e, 0xc4, 0x53, 0x31, 0x76, 0x64, 0x44, 0xc9, 0x8b, 0xa0, 0x9c, 0x53, 0x51, 0xd4, 0xa1,
	0xd3, 0xf9, 0xf3, 0x6f, 0xfd, 0xfe, 0x74, 0x70, 0x14, 0x1a, 0x5b, 0x7a, 0x6a, 0xe6, 0xce, 0xbb,
	0xbc, 0x59, 0x3a, 0xef, 0x10, 0x65, 0x70, 0xbe, 0x91, 0x4c, 0x3f, 0x15, 0xe0, 0xbd, 0xf0, 0x5d,
	0xcf, 0x05, 0xbd, 0x05, 0x62, 0x8f, 0x1a, 0x0e, 0xca, 0xaa, 0x72, 0xa1, 0xf6, 0x5a, 0x4d, 0xd4,
	0xd5, 0xb8, 0x58, 0x23, 0x9e, 0x42, 0x16, 0x98, 0x96, 0x33, 0xab, 0xf7, 0x24, 0x18, 0x08, 0xcf,
	0x61, 0xcc, 0xc4, 0xbc, 0x70, 0xf5, 0xcc, 0xea, 0x7d, 0x89, 0xfe, 0x2e, 0xf0, 0x04, 0x52, 0xf9,
	0x4f, 0x8f, 0x24, 0x89, 0xd0, 0x77, 0x59, 0x7a, 0x5f, 0x54, 0xa4, 0xd3, 0xd8, 0x15, 0x69, 0xfa,
	0x04, 0xc7, 0xff, 0x9c, 0xb8, 0x71, 0x35, 0x53, 0x2f, 0xf5, 0x4a, 0xcc, 0xe5, 0x0b, 0xad, 0xa5,
	0x06, 0x44, 0x84, 0x51, 0xe5, 0x2c, 0x89, 0x52, 0x5a, 0xc8, 0xb9, 0x2f, 0x67, 0x5f, 0xfa, 0xc0,
	0x62, 0x93, 0x16, 0x03, 0x5d, 0xcf, 0xe1, 0x70, 0xa3, 0x1c, 0x1f, 0x20, 0x8b, 0xfb, 0xc0, 0x8b,
	0x7c, 0x7b, 0x3f, 0xf9, 0xf6, 0x6e, 0xce, 0x2e, 0x77, 0xbe, 0x8b, 0xbe, 0xb7, 0x93, 0xaf, 0xd6,
	0xa8, 0x55, 0x6b, 0xd4, 0x4f, 0x6b, 0xd4, 0x47, 0x67, 0x92, 0x55, 0x67, 0x92, 0xef, 0xce, 0x24,
	0x8f, 0x59, 0x6c, 0x78, 0x8e, 0xf3, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0x08, 0x21, 0x78, 0xfa,
	0xad, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdatePhotoClient is the client API for UpdatePhoto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdatePhotoClient interface {
	Update(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*UpdatePhotoResponse, error)
}

type updatePhotoClient struct {
	cc *grpc.ClientConn
}

func NewUpdatePhotoClient(cc *grpc.ClientConn) UpdatePhotoClient {
	return &updatePhotoClient{cc}
}

func (c *updatePhotoClient) Update(ctx context.Context, in *UpdatePhotoRequest, opts ...grpc.CallOption) (*UpdatePhotoResponse, error) {
	out := new(UpdatePhotoResponse)
	err := c.cc.Invoke(ctx, "/protos.updatephoto.UpdatePhoto/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdatePhotoServer is the server API for UpdatePhoto service.
type UpdatePhotoServer interface {
	Update(context.Context, *UpdatePhotoRequest) (*UpdatePhotoResponse, error)
}

// UnimplementedUpdatePhotoServer can be embedded to have forward compatible implementations.
type UnimplementedUpdatePhotoServer struct {
}

func (*UnimplementedUpdatePhotoServer) Update(ctx context.Context, req *UpdatePhotoRequest) (*UpdatePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterUpdatePhotoServer(s *grpc.Server, srv UpdatePhotoServer) {
	s.RegisterService(&_UpdatePhoto_serviceDesc, srv)
}

func _UpdatePhoto_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdatePhotoServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.updatephoto.UpdatePhoto/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdatePhotoServer).Update(ctx, req.(*UpdatePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpdatePhoto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.updatephoto.UpdatePhoto",
	HandlerType: (*UpdatePhotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "update",
			Handler:    _UpdatePhoto_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updatephoto.proto",
}

func (m *UpdatePhotoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePhotoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.SessionId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(len(m.SessionId)))
		i += copy(dAtA[i:], m.SessionId)
	}
	if len(m.Photo) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(len(m.Photo)))
		i += copy(dAtA[i:], m.Photo)
	}
	if len(m.Device) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	return i, nil
}

func (m *UpdatePhotoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePhotoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Code != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(m.Code))
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUpdatephoto(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeVarintUpdatephoto(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UpdatePhotoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovUpdatephoto(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovUpdatephoto(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovUpdatephoto(uint64(l))
	}
	l = len(m.Photo)
	if l > 0 {
		n += 1 + l + sovUpdatephoto(uint64(l))
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovUpdatephoto(uint64(l))
	}
	return n
}

func (m *UpdatePhotoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovUpdatephoto(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovUpdatephoto(uint64(m.Code))
	}
	if m.Status != 0 {
		n += 1 + sovUpdatephoto(uint64(m.Status))
	}
	return n
}

func sovUpdatephoto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpdatephoto(x uint64) (n int) {
	return sovUpdatephoto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdatePhotoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdatephoto
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
			return fmt.Errorf("proto: UpdatePhotoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePhotoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
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
				return ErrInvalidLengthUpdatephoto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
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
				return ErrInvalidLengthUpdatephoto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
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
				return ErrInvalidLengthUpdatephoto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Photo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
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
				return ErrInvalidLengthUpdatephoto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Photo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
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
				return ErrInvalidLengthUpdatephoto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpdatephoto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUpdatephoto
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
func (m *UpdatePhotoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdatephoto
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
			return fmt.Errorf("proto: UpdatePhotoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePhotoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
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
				return ErrInvalidLengthUpdatephoto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdatephoto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUpdatephoto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpdatephoto
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUpdatephoto
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
func skipUpdatephoto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpdatephoto
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
					return 0, ErrIntOverflowUpdatephoto
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUpdatephoto
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
				return 0, ErrInvalidLengthUpdatephoto
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthUpdatephoto
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUpdatephoto
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUpdatephoto(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthUpdatephoto
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUpdatephoto = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpdatephoto   = fmt.Errorf("proto: integer overflow")
)
