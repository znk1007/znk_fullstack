// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: buffer.proto

package pbs

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

// generate command-line: protoc -I=protos --gogofaster_out=plugins=grpc:protos/pbs protos/*.proto
// buffer: binary buffer for emit args, used for transport
type Buffer struct {
	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	IsBinary bool   `protobuf:"varint,2,opt,name=isBinary,proto3" json:"isBinary,omitempty"`
	Num      uint64 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e21993b51421452, []int{0}
}
func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return m.Size()
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Buffer) GetIsBinary() bool {
	if m != nil {
		return m.IsBinary
	}
	return false
}

func (m *Buffer) GetNum() uint64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*Buffer)(nil), "socket.go.Buffer")
}

func init() { proto.RegisterFile("buffer.proto", fileDescriptor_1e21993b51421452) }

var fileDescriptor_1e21993b51421452 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0x4d, 0x4b,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0xce, 0x4f, 0xce, 0x4e, 0x2d,
	0xd1, 0x4b, 0xcf, 0x57, 0xf2, 0xe2, 0x62, 0x73, 0x02, 0x4b, 0x09, 0x09, 0x71, 0xb1, 0xa4, 0x24,
	0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x42, 0x52, 0x5c, 0x1c, 0x99,
	0xc5, 0x4e, 0x99, 0x79, 0x89, 0x45, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x70, 0xbe,
	0x90, 0x00, 0x17, 0x73, 0x5e, 0x69, 0xae, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x88, 0xe9,
	0x24, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xcc, 0x05, 0x49, 0xc5,
	0x49, 0x6c, 0x60, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xf2, 0x81, 0x30, 0x8c,
	0x00, 0x00, 0x00,
}

func (m *Buffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Buffer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.IsBinary {
		dAtA[i] = 0x10
		i++
		if m.IsBinary {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Num != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.Num))
	}
	return i, nil
}

func encodeVarintBuffer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Buffer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovBuffer(uint64(l))
	}
	if m.IsBinary {
		n += 2
	}
	if m.Num != 0 {
		n += 1 + sovBuffer(uint64(m.Num))
	}
	return n
}

func sovBuffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBuffer(x uint64) (n int) {
	return sovBuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Buffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: Buffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Buffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBuffer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBinary", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
			m.IsBinary = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBuffer
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
func skipBuffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
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
				return 0, ErrInvalidLengthBuffer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBuffer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBuffer
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
				next, err := skipBuffer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBuffer
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
	ErrInvalidLengthBuffer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuffer   = fmt.Errorf("proto: integer overflow")
)