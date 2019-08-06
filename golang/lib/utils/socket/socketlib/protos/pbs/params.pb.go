// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: params.proto

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
// connection parameters for server
type ConnParameters struct {
	PingInterval int64    `protobuf:"varint,1,opt,name=pingInterval,proto3" json:"pingInterval,omitempty"`
	PingTimeout  int64    `protobuf:"varint,2,opt,name=pingTimeout,proto3" json:"pingTimeout,omitempty"`
	SID          string   `protobuf:"bytes,3,opt,name=sID,proto3" json:"sID,omitempty"`
	Upgrades     []string `protobuf:"bytes,4,rep,name=upgrades,proto3" json:"upgrades,omitempty"`
}

func (m *ConnParameters) Reset()         { *m = ConnParameters{} }
func (m *ConnParameters) String() string { return proto.CompactTextString(m) }
func (*ConnParameters) ProtoMessage()    {}
func (*ConnParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_8679b07c520418a1, []int{0}
}
func (m *ConnParameters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnParameters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnParameters.Merge(m, src)
}
func (m *ConnParameters) XXX_Size() int {
	return m.Size()
}
func (m *ConnParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ConnParameters proto.InternalMessageInfo

func (m *ConnParameters) GetPingInterval() int64 {
	if m != nil {
		return m.PingInterval
	}
	return 0
}

func (m *ConnParameters) GetPingTimeout() int64 {
	if m != nil {
		return m.PingTimeout
	}
	return 0
}

func (m *ConnParameters) GetSID() string {
	if m != nil {
		return m.SID
	}
	return ""
}

func (m *ConnParameters) GetUpgrades() []string {
	if m != nil {
		return m.Upgrades
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnParameters)(nil), "core.go.ConnParameters")
}

func init() { proto.RegisterFile("params.proto", fileDescriptor_8679b07c520418a1) }

var fileDescriptor_8679b07c520418a1 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2c, 0x4a,
	0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4b,
	0xcf, 0x57, 0x6a, 0x61, 0xe4, 0xe2, 0x73, 0xce, 0xcf, 0xcb, 0x0b, 0x00, 0xc9, 0xa6, 0x96, 0xa4,
	0x16, 0x15, 0x0b, 0x29, 0x71, 0xf1, 0x14, 0x64, 0xe6, 0xa5, 0x7b, 0xe6, 0x95, 0xa4, 0x16, 0x95,
	0x25, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xa1, 0x88, 0x09, 0x29, 0x70, 0x71, 0x83,
	0xf8, 0x21, 0x99, 0xb9, 0xa9, 0xf9, 0xa5, 0x25, 0x12, 0x4c, 0x60, 0x25, 0xc8, 0x42, 0x42, 0x02,
	0x5c, 0xcc, 0xc5, 0x9e, 0x2e, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x14,
	0x17, 0x47, 0x69, 0x41, 0x7a, 0x51, 0x62, 0x4a, 0x6a, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0x67,
	0x10, 0x9c, 0xef, 0x24, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xcc,
	0x05, 0x49, 0xc5, 0x49, 0x6c, 0x60, 0x57, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x43,
	0xb7, 0x5d, 0xc5, 0x00, 0x00, 0x00,
}

func (m *ConnParameters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnParameters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PingInterval != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintParams(dAtA, i, uint64(m.PingInterval))
	}
	if m.PingTimeout != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintParams(dAtA, i, uint64(m.PingTimeout))
	}
	if len(m.SID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintParams(dAtA, i, uint64(len(m.SID)))
		i += copy(dAtA[i:], m.SID)
	}
	if len(m.Upgrades) > 0 {
		for _, s := range m.Upgrades {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConnParameters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PingInterval != 0 {
		n += 1 + sovParams(uint64(m.PingInterval))
	}
	if m.PingTimeout != 0 {
		n += 1 + sovParams(uint64(m.PingTimeout))
	}
	l = len(m.SID)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.Upgrades) > 0 {
		for _, s := range m.Upgrades {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConnParameters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ConnParameters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnParameters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingInterval", wireType)
			}
			m.PingInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PingInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PingTimeout", wireType)
			}
			m.PingTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PingTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upgrades", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Upgrades = append(m.Upgrades, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthParams
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowParams
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
				next, err := skipParams(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthParams
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
	ErrInvalidLengthParams = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams   = fmt.Errorf("proto: integer overflow")
)
