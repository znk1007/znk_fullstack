// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: updateonline.proto

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

type UpdateOnlineRequest struct {
	Account   string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	SessionId string `protobuf:"bytes,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	Online    bool   `protobuf:"varint,4,opt,name=online,proto3" json:"online,omitempty"`
	Device    string `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`
}

func (m *UpdateOnlineRequest) Reset()         { *m = UpdateOnlineRequest{} }
func (m *UpdateOnlineRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateOnlineRequest) ProtoMessage()    {}
func (*UpdateOnlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad52548456099ed4, []int{0}
}
func (m *UpdateOnlineRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateOnlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateOnlineRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateOnlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOnlineRequest.Merge(m, src)
}
func (m *UpdateOnlineRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdateOnlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOnlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOnlineRequest proto.InternalMessageInfo

func (m *UpdateOnlineRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UpdateOnlineRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateOnlineRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *UpdateOnlineRequest) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *UpdateOnlineRequest) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

type UpdateOnlineResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Status  int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *UpdateOnlineResponse) Reset()         { *m = UpdateOnlineResponse{} }
func (m *UpdateOnlineResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateOnlineResponse) ProtoMessage()    {}
func (*UpdateOnlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad52548456099ed4, []int{1}
}
func (m *UpdateOnlineResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateOnlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateOnlineResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateOnlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOnlineResponse.Merge(m, src)
}
func (m *UpdateOnlineResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdateOnlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOnlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOnlineResponse proto.InternalMessageInfo

func (m *UpdateOnlineResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UpdateOnlineResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateOnlineResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*UpdateOnlineRequest)(nil), "protos.updateonline.UpdateOnlineRequest")
	proto.RegisterType((*UpdateOnlineResponse)(nil), "protos.updateonline.UpdateOnlineResponse")
}

func init() { proto.RegisterFile("updateonline.proto", fileDescriptor_ad52548456099ed4) }

var fileDescriptor_ad52548456099ed4 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x68, 0x02, 0xb5, 0x98, 0xae, 0x08, 0x59, 0x08, 0x59, 0x51, 0xa7, 0xb0, 0x64,
	0x80, 0x37, 0x60, 0xeb, 0x84, 0x64, 0x89, 0x05, 0xc1, 0x10, 0xe2, 0x13, 0x8a, 0x44, 0xed, 0xd0,
	0xb3, 0x79, 0x0e, 0x16, 0xde, 0x89, 0xb1, 0x23, 0x23, 0x4a, 0x5e, 0x04, 0xc5, 0x4e, 0x45, 0x91,
	0x32, 0x30, 0x9d, 0xff, 0xff, 0xec, 0xdf, 0x9f, 0x7e, 0x0e, 0xbe, 0xd5, 0x95, 0x43, 0x6b, 0x5e,
	0x1a, 0x83, 0x65, 0xbb, 0xb1, 0xce, 0xc2, 0x22, 0x0c, 0x2a, 0xf7, 0x57, 0xcb, 0x0f, 0xc6, 0x17,
	0x77, 0xc1, 0xb8, 0x0d, 0x86, 0xc2, 0x57, 0x8f, 0xe4, 0x40, 0xf0, 0xa3, 0xaa, 0xae, 0xad, 0x37,
	0x4e, 0xb0, 0x9c, 0x15, 0x73, 0xb5, 0x93, 0x70, 0xc6, 0x33, 0x4f, 0xb8, 0x59, 0x69, 0x71, 0x10,
	0x16, 0xa3, 0x82, 0x0b, 0x3e, 0x27, 0x24, 0x6a, 0xac, 0x59, 0x69, 0x71, 0x18, 0x56, 0xbf, 0xc6,
	0xf0, 0x2a, 0xfe, 0x28, 0x66, 0x39, 0x2b, 0x8e, 0xd5, 0xa8, 0x06, 0x5f, 0xe3, 0x5b, 0x53, 0xa3,
	0x48, 0x63, 0x5a, 0x54, 0xcb, 0x07, 0x7e, 0xfa, 0x17, 0x8b, 0x5a, 0x6b, 0x08, 0x07, 0xae, 0x35,
	0x12, 0x55, 0xcf, 0xb8, 0xe3, 0x1a, 0x25, 0x00, 0x9f, 0xd5, 0x56, 0x63, 0xa0, 0x4a, 0x55, 0x38,
	0x0f, 0xe9, 0xe4, 0x2a, 0xe7, 0x29, 0x00, 0xa5, 0x6a, 0x54, 0x57, 0x6b, 0x7e, 0xb2, 0x9f, 0x0e,
	0x8f, 0x3c, 0x8b, 0xad, 0x40, 0x51, 0x4e, 0xb4, 0x54, 0x4e, 0x34, 0x74, 0x7e, 0xf9, 0x8f, 0x9b,
	0x11, 0xfa, 0x26, 0xff, 0xec, 0x24, 0xdb, 0x76, 0x92, 0x7d, 0x77, 0x92, 0xbd, 0xf7, 0x32, 0xd9,
	0xf6, 0x32, 0xf9, 0xea, 0x65, 0x72, 0x9f, 0xc5, 0x8c, 0xa7, 0x38, 0xaf, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x10, 0x82, 0x4e, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdateOnlineClient is the client API for UpdateOnline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateOnlineClient interface {
	Update(ctx context.Context, in *UpdateOnlineRequest, opts ...grpc.CallOption) (*UpdateOnlineResponse, error)
}

type updateOnlineClient struct {
	cc *grpc.ClientConn
}

func NewUpdateOnlineClient(cc *grpc.ClientConn) UpdateOnlineClient {
	return &updateOnlineClient{cc}
}

func (c *updateOnlineClient) Update(ctx context.Context, in *UpdateOnlineRequest, opts ...grpc.CallOption) (*UpdateOnlineResponse, error) {
	out := new(UpdateOnlineResponse)
	err := c.cc.Invoke(ctx, "/protos.updateonline.UpdateOnline/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateOnlineServer is the server API for UpdateOnline service.
type UpdateOnlineServer interface {
	Update(context.Context, *UpdateOnlineRequest) (*UpdateOnlineResponse, error)
}

// UnimplementedUpdateOnlineServer can be embedded to have forward compatible implementations.
type UnimplementedUpdateOnlineServer struct {
}

func (*UnimplementedUpdateOnlineServer) Update(ctx context.Context, req *UpdateOnlineRequest) (*UpdateOnlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterUpdateOnlineServer(s *grpc.Server, srv UpdateOnlineServer) {
	s.RegisterService(&_UpdateOnline_serviceDesc, srv)
}

func _UpdateOnline_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateOnlineServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.updateonline.UpdateOnline/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateOnlineServer).Update(ctx, req.(*UpdateOnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpdateOnline_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.updateonline.UpdateOnline",
	HandlerType: (*UpdateOnlineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "update",
			Handler:    _UpdateOnline_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updateonline.proto",
}

func (m *UpdateOnlineRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateOnlineRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(len(m.Account)))
		i += copy(dAtA[i:], m.Account)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.SessionId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(len(m.SessionId)))
		i += copy(dAtA[i:], m.SessionId)
	}
	if m.Online {
		dAtA[i] = 0x20
		i++
		if m.Online {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Device) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	return i, nil
}

func (m *UpdateOnlineResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateOnlineResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Code != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(m.Code))
	}
	if m.Status != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUpdateonline(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeVarintUpdateonline(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UpdateOnlineRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovUpdateonline(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovUpdateonline(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovUpdateonline(uint64(l))
	}
	if m.Online {
		n += 2
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovUpdateonline(uint64(l))
	}
	return n
}

func (m *UpdateOnlineResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovUpdateonline(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovUpdateonline(uint64(m.Code))
	}
	if m.Status != 0 {
		n += 1 + sovUpdateonline(uint64(m.Status))
	}
	return n
}

func sovUpdateonline(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUpdateonline(x uint64) (n int) {
	return sovUpdateonline(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateOnlineRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdateonline
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
			return fmt.Errorf("proto: UpdateOnlineRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateOnlineRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdateonline
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
				return ErrInvalidLengthUpdateonline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateonline
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
					return ErrIntOverflowUpdateonline
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
				return ErrInvalidLengthUpdateonline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateonline
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
					return ErrIntOverflowUpdateonline
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
				return ErrInvalidLengthUpdateonline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateonline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Online", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdateonline
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Online = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdateonline
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
				return ErrInvalidLengthUpdateonline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateonline
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpdateonline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpdateonline
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUpdateonline
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
func (m *UpdateOnlineResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpdateonline
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
			return fmt.Errorf("proto: UpdateOnlineResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateOnlineResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpdateonline
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
				return ErrInvalidLengthUpdateonline
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUpdateonline
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
					return ErrIntOverflowUpdateonline
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
					return ErrIntOverflowUpdateonline
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
			skippy, err := skipUpdateonline(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpdateonline
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUpdateonline
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
func skipUpdateonline(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpdateonline
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
					return 0, ErrIntOverflowUpdateonline
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
					return 0, ErrIntOverflowUpdateonline
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
				return 0, ErrInvalidLengthUpdateonline
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthUpdateonline
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUpdateonline
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
				next, err := skipUpdateonline(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthUpdateonline
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
	ErrInvalidLengthUpdateonline = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpdateonline   = fmt.Errorf("proto: integer overflow")
)
