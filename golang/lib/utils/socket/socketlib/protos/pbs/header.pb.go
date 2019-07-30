// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: header.proto

package socketproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// packet type for connect to server
type Type int32

const (
	Type_connect     Type = 0
	Type_disconnect  Type = 1
	Type_event       Type = 2
	Type_ack         Type = 3
	Type_error       Type = 4
	Type_binaryEvent Type = 6
	Type_binaryAck   Type = 7
	Type_typeMax     Type = 8
)

var Type_name = map[int32]string{
	0: "connect",
	1: "disconnect",
	2: "event",
	3: "ack",
	4: "error",
	6: "binaryEvent",
	7: "binaryAck",
	8: "typeMax",
}

var Type_value = map[string]int32{
	"connect":     0,
	"disconnect":  1,
	"event":       2,
	"ack":         3,
	"error":       4,
	"binaryEvent": 6,
	"binaryAck":   7,
	"typeMax":     8,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{0}
}

// header for packet
type Header struct {
	T         Type   `protobuf:"varint,1,opt,name=t,proto3,enum=socket.go.Type" json:"t,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ID        uint64 `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
	NeedAck   bool   `protobuf:"varint,4,opt,name=needAck,proto3" json:"needAck,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{0}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetT() Type {
	if m != nil {
		return m.T
	}
	return Type_connect
}

func (m *Header) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Header) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Header) GetNeedAck() bool {
	if m != nil {
		return m.NeedAck
	}
	return false
}

func init() {
	proto.RegisterEnum("socket.go.Type", Type_name, Type_value)
	proto.RegisterType((*Header)(nil), "socket.go.Header")
}

func init() { proto.RegisterFile("header.proto", fileDescriptor_6398613e36d6c2ce) }

var fileDescriptor_6398613e36d6c2ce = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xb3, 0xb9, 0x98, 0x5c, 0x26, 0x9a, 0x5b, 0xa6, 0x4a, 0xa1, 0x4b, 0x10, 0x84, 0x60,
	0x91, 0x42, 0x9f, 0xe0, 0xe4, 0x04, 0xaf, 0xb0, 0x09, 0x56, 0x76, 0x7b, 0x9b, 0x41, 0x8f, 0xe0,
	0x6e, 0xd8, 0x2c, 0x62, 0xde, 0xc2, 0xc7, 0xb2, 0xbc, 0xd2, 0x52, 0x92, 0x17, 0x91, 0xcb, 0x71,
	0x5a, 0x0d, 0xff, 0x37, 0x3f, 0x1f, 0xcc, 0xc0, 0xe9, 0x2b, 0xc9, 0x9a, 0x6c, 0xd9, 0x5a, 0xe3,
	0x0c, 0xc6, 0x9d, 0x51, 0x0d, 0xb9, 0xf2, 0xc5, 0x5c, 0x1a, 0x08, 0x1f, 0xa6, 0x15, 0x5e, 0x00,
	0x73, 0x19, 0xcb, 0x59, 0x91, 0xde, 0x2c, 0xca, 0xbf, 0x42, 0xf9, 0xd4, 0xb7, 0x54, 0x31, 0x87,
	0xe7, 0x10, 0x6b, 0xf9, 0x46, 0x5d, 0x2b, 0x15, 0x65, 0x7e, 0xce, 0x8a, 0xb8, 0xfa, 0x07, 0x98,
	0x82, 0xbf, 0x5e, 0x65, 0xb3, 0x9c, 0x15, 0x41, 0xe5, 0xaf, 0x57, 0x98, 0x41, 0xa4, 0x89, 0xea,
	0xa5, 0x6a, 0xb2, 0x20, 0x67, 0xc5, 0xbc, 0x3a, 0xc6, 0x6b, 0x03, 0xc1, 0x5e, 0x89, 0x09, 0x44,
	0xca, 0x68, 0x4d, 0xca, 0x71, 0x0f, 0x53, 0x80, 0x7a, 0xdb, 0x1d, 0x33, 0xc3, 0x18, 0x4e, 0xe8,
	0x9d, 0xb4, 0xe3, 0x3e, 0x46, 0x30, 0x93, 0xaa, 0xe1, 0xb3, 0x89, 0x59, 0x6b, 0x2c, 0x0f, 0x70,
	0x01, 0xc9, 0x66, 0xab, 0xa5, 0xed, 0xef, 0xa7, 0x52, 0x88, 0x67, 0x10, 0x1f, 0xc0, 0x52, 0x35,
	0x3c, 0xda, 0xbb, 0x5d, 0xdf, 0xd2, 0xa3, 0xfc, 0xe0, 0xf3, 0xbb, 0xab, 0xaf, 0x41, 0xb0, 0xdd,
	0x20, 0xd8, 0xcf, 0x20, 0xd8, 0xe7, 0x28, 0xbc, 0xdd, 0x28, 0xbc, 0xef, 0x51, 0x78, 0xcf, 0xc9,
	0xe1, 0xca, 0xe9, 0x27, 0x9b, 0x70, 0x1a, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x4f,
	0x98, 0xa5, 0x2a, 0x01, 0x00, 0x00,
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.T != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.T))
	}
	if len(m.Namespace) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHeader(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if m.ID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHeader(dAtA, i, uint64(m.ID))
	}
	if m.NeedAck {
		dAtA[i] = 0x20
		i++
		if m.NeedAck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintHeader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.T != 0 {
		n += 1 + sovHeader(uint64(m.T))
	}
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovHeader(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovHeader(uint64(m.ID))
	}
	if m.NeedAck {
		n += 2
	}
	return n
}

func sovHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeader(x uint64) (n int) {
	return sovHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeader
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			m.T = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.T |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
				return ErrInvalidLengthHeader
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedAck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeader
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
			m.NeedAck = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeader
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeader
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
func skipHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
					return 0, ErrIntOverflowHeader
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
				return 0, ErrInvalidLengthHeader
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHeader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeader
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
				next, err := skipHeader(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHeader
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
	ErrInvalidLengthHeader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeader   = fmt.Errorf("proto: integer overflow")
)
