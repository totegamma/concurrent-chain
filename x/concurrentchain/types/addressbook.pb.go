// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: concurrentchain/concurrentchain/addressbook.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Addressbook struct {
	Index   string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Fqdn    string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Addressbook) Reset()         { *m = Addressbook{} }
func (m *Addressbook) String() string { return proto.CompactTextString(m) }
func (*Addressbook) ProtoMessage()    {}
func (*Addressbook) Descriptor() ([]byte, []int) {
	return fileDescriptor_af969b6b7c5a387e, []int{0}
}
func (m *Addressbook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Addressbook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Addressbook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Addressbook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addressbook.Merge(m, src)
}
func (m *Addressbook) XXX_Size() int {
	return m.Size()
}
func (m *Addressbook) XXX_DiscardUnknown() {
	xxx_messageInfo_Addressbook.DiscardUnknown(m)
}

var xxx_messageInfo_Addressbook proto.InternalMessageInfo

func (m *Addressbook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Addressbook) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Addressbook) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Addressbook)(nil), "concurrentchain.concurrentchain.Addressbook")
}

func init() {
	proto.RegisterFile("concurrentchain/concurrentchain/addressbook.proto", fileDescriptor_af969b6b7c5a387e)
}

var fileDescriptor_af969b6b7c5a387e = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xce, 0xcf, 0x4b,
	0x2e, 0x2d, 0x2a, 0x4a, 0xcd, 0x2b, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x47, 0xe7, 0x27, 0xa6,
	0xa4, 0x14, 0xa5, 0x16, 0x17, 0x27, 0xe5, 0xe7, 0x67, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0xc9, 0xa3, 0x29, 0xd1, 0x43, 0xe3, 0x2b, 0x05, 0x72, 0x71, 0x3b, 0x22, 0x74, 0x09, 0x89, 0x70,
	0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42,
	0x42, 0x5c, 0x2c, 0x69, 0x85, 0x29, 0x79, 0x12, 0x4c, 0x60, 0x41, 0x30, 0x5b, 0x48, 0x82, 0x8b,
	0x3d, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x19, 0x2c, 0x0c, 0xe3, 0x3a, 0x85, 0x9f,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6d, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x49, 0x7e, 0x49, 0x6a, 0x7a, 0x62, 0x6e, 0x6e, 0x22, 0x92,
	0x2f, 0x74, 0x21, 0xde, 0xa8, 0xc0, 0xf0, 0x58, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8,
	0x4f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x07, 0x58, 0x7e, 0x08, 0x01, 0x00, 0x00,
}

func (m *Addressbook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Addressbook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Addressbook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAddressbook(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Fqdn) > 0 {
		i -= len(m.Fqdn)
		copy(dAtA[i:], m.Fqdn)
		i = encodeVarintAddressbook(dAtA, i, uint64(len(m.Fqdn)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintAddressbook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAddressbook(dAtA []byte, offset int, v uint64) int {
	offset -= sovAddressbook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Addressbook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovAddressbook(uint64(l))
	}
	l = len(m.Fqdn)
	if l > 0 {
		n += 1 + l + sovAddressbook(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAddressbook(uint64(l))
	}
	return n
}

func sovAddressbook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAddressbook(x uint64) (n int) {
	return sovAddressbook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Addressbook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddressbook
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
			return fmt.Errorf("proto: Addressbook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Addressbook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressbook
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
				return ErrInvalidLengthAddressbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fqdn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressbook
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
				return ErrInvalidLengthAddressbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fqdn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddressbook
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
				return ErrInvalidLengthAddressbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddressbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAddressbook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddressbook
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
func skipAddressbook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddressbook
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
					return 0, ErrIntOverflowAddressbook
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
					return 0, ErrIntOverflowAddressbook
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
				return 0, ErrInvalidLengthAddressbook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAddressbook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAddressbook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAddressbook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddressbook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAddressbook = fmt.Errorf("proto: unexpected end of group")
)
