// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/storage/chunk/chunk.proto

package chunk

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CompressionAlgo int32

const (
	CompressionAlgo_NONE            CompressionAlgo = 0
	CompressionAlgo_GZIP_BEST_SPEED CompressionAlgo = 1
)

var CompressionAlgo_name = map[int32]string{
	0: "NONE",
	1: "GZIP_BEST_SPEED",
}

var CompressionAlgo_value = map[string]int32{
	"NONE":            0,
	"GZIP_BEST_SPEED": 1,
}

func (x CompressionAlgo) String() string {
	return proto.EnumName(CompressionAlgo_name, int32(x))
}

func (CompressionAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b743b4a788792d7, []int{0}
}

type EncryptionAlgo int32

const (
	EncryptionAlgo_CHACHA20 EncryptionAlgo = 0
)

var EncryptionAlgo_name = map[int32]string{
	0: "CHACHA20",
}

var EncryptionAlgo_value = map[string]int32{
	"CHACHA20": 0,
}

func (x EncryptionAlgo) String() string {
	return proto.EnumName(EncryptionAlgo_name, int32(x))
}

func (EncryptionAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4b743b4a788792d7, []int{1}
}

// DataRef is a reference to data within a chunk.
type DataRef struct {
	// The chunk the referenced data is located in.
	Ref *Ref `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// The hash of the data being referenced.
	// This field is empty when it is equal to the chunk hash (the ref is the whole chunk).
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The offset and size used for accessing the data within the chunk.
	OffsetBytes          int64    `protobuf:"varint,3,opt,name=offset_bytes,json=offsetBytes,proto3" json:"offset_bytes,omitempty"`
	SizeBytes            int64    `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRef) Reset()         { *m = DataRef{} }
func (m *DataRef) String() string { return proto.CompactTextString(m) }
func (*DataRef) ProtoMessage()    {}
func (*DataRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b743b4a788792d7, []int{0}
}
func (m *DataRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRef.Merge(m, src)
}
func (m *DataRef) XXX_Size() int {
	return m.Size()
}
func (m *DataRef) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRef.DiscardUnknown(m)
}

var xxx_messageInfo_DataRef proto.InternalMessageInfo

func (m *DataRef) GetRef() *Ref {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *DataRef) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *DataRef) GetOffsetBytes() int64 {
	if m != nil {
		return m.OffsetBytes
	}
	return 0
}

func (m *DataRef) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

type Ref struct {
	Id                   []byte          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SizeBytes            int64           `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	Edge                 bool            `protobuf:"varint,3,opt,name=edge,proto3" json:"edge,omitempty"`
	Dek                  []byte          `protobuf:"bytes,4,opt,name=dek,proto3" json:"dek,omitempty"`
	EncryptionAlgo       EncryptionAlgo  `protobuf:"varint,5,opt,name=encryption_algo,json=encryptionAlgo,proto3,enum=chunk.EncryptionAlgo" json:"encryption_algo,omitempty"`
	CompressionAlgo      CompressionAlgo `protobuf:"varint,6,opt,name=compression_algo,json=compressionAlgo,proto3,enum=chunk.CompressionAlgo" json:"compression_algo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Ref) Reset()         { *m = Ref{} }
func (m *Ref) String() string { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()    {}
func (*Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b743b4a788792d7, []int{1}
}
func (m *Ref) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ref.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ref.Merge(m, src)
}
func (m *Ref) XXX_Size() int {
	return m.Size()
}
func (m *Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_Ref proto.InternalMessageInfo

func (m *Ref) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ref) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func (m *Ref) GetEdge() bool {
	if m != nil {
		return m.Edge
	}
	return false
}

func (m *Ref) GetDek() []byte {
	if m != nil {
		return m.Dek
	}
	return nil
}

func (m *Ref) GetEncryptionAlgo() EncryptionAlgo {
	if m != nil {
		return m.EncryptionAlgo
	}
	return EncryptionAlgo_CHACHA20
}

func (m *Ref) GetCompressionAlgo() CompressionAlgo {
	if m != nil {
		return m.CompressionAlgo
	}
	return CompressionAlgo_NONE
}

func init() {
	proto.RegisterEnum("chunk.CompressionAlgo", CompressionAlgo_name, CompressionAlgo_value)
	proto.RegisterEnum("chunk.EncryptionAlgo", EncryptionAlgo_name, EncryptionAlgo_value)
	proto.RegisterType((*DataRef)(nil), "chunk.DataRef")
	proto.RegisterType((*Ref)(nil), "chunk.Ref")
}

func init() {
	proto.RegisterFile("internal/storage/chunk/chunk.proto", fileDescriptor_4b743b4a788792d7)
}

var fileDescriptor_4b743b4a788792d7 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x7d, 0x93, 0xf4, 0x3d, 0xfb, 0xee, 0x0b, 0x49, 0x18, 0x51, 0xb2, 0xd0, 0x10, 0xbb, 0x0a,
	0x6f, 0xd1, 0x3c, 0xe2, 0x56, 0x84, 0xb4, 0x0d, 0x56, 0x17, 0xb5, 0x4c, 0x5d, 0x75, 0x13, 0xd2,
	0x64, 0xf2, 0x41, 0xdb, 0x4c, 0x98, 0x99, 0x0a, 0x15, 0xfc, 0x7f, 0x2e, 0xfd, 0x03, 0x82, 0xf4,
	0x97, 0x48, 0xa6, 0xa5, 0xda, 0xe2, 0x66, 0x38, 0xf7, 0xdc, 0x73, 0xcf, 0xb9, 0x03, 0x17, 0x06,
	0x75, 0x23, 0x29, 0x6f, 0xd2, 0x4d, 0x20, 0x24, 0xe3, 0x69, 0x49, 0x83, 0xac, 0xda, 0x35, 0xeb,
	0xe3, 0x3b, 0x6c, 0x39, 0x93, 0x0c, 0xdf, 0xaa, 0x62, 0xf0, 0x1d, 0x9e, 0x4d, 0x52, 0x99, 0x12,
	0x5a, 0xe0, 0x57, 0xa0, 0x73, 0x5a, 0x38, 0xc8, 0x43, 0xfe, 0x43, 0x08, 0xc3, 0xa3, 0x98, 0xd0,
	0x82, 0x74, 0x34, 0xc6, 0xd0, 0xab, 0x52, 0x51, 0x39, 0x9a, 0x87, 0xfc, 0x7b, 0xa2, 0x30, 0x7e,
	0x03, 0x06, 0x2b, 0x0a, 0x41, 0x65, 0xb2, 0xda, 0x4b, 0x2a, 0x1c, 0xdd, 0x43, 0xbe, 0x4e, 0x1e,
	0x8e, 0xdc, 0xa8, 0xa3, 0xf0, 0x6b, 0x00, 0x51, 0x7f, 0xa3, 0x27, 0x41, 0x4f, 0x09, 0xee, 0x3b,
	0x46, 0xb5, 0x07, 0xbf, 0x10, 0xe8, 0x5d, 0xb6, 0x09, 0x5a, 0x9d, 0xab, 0x68, 0x83, 0x68, 0x75,
	0x7e, 0x35, 0xa6, 0x5d, 0x8d, 0x75, 0xcb, 0xd0, 0xbc, 0xa4, 0x2a, 0xb0, 0x4f, 0x14, 0xc6, 0x36,
	0xe8, 0x39, 0x5d, 0xab, 0x08, 0x83, 0x74, 0x10, 0xbf, 0x07, 0x8b, 0x36, 0x19, 0xdf, 0xb7, 0xb2,
	0x66, 0x4d, 0x92, 0x6e, 0x4a, 0xe6, 0xdc, 0x7a, 0xc8, 0x37, 0xc3, 0x17, 0xa7, 0xcf, 0xc5, 0xe7,
	0x6e, 0xb4, 0x29, 0x19, 0x31, 0xe9, 0x45, 0x8d, 0x23, 0xb0, 0x33, 0xb6, 0x6d, 0x39, 0x15, 0xe2,
	0x6c, 0x70, 0xa7, 0x0c, 0x5e, 0x9e, 0x0c, 0xc6, 0x7f, 0xdb, 0xca, 0xc1, 0xca, 0x2e, 0x89, 0xc7,
	0x27, 0xb0, 0xae, 0x34, 0xb8, 0x0f, 0xbd, 0xd9, 0xe7, 0x59, 0x6c, 0xdf, 0xe0, 0xe7, 0x60, 0x7d,
	0x58, 0x7e, 0x9c, 0x27, 0xa3, 0x78, 0xf1, 0x25, 0x59, 0xcc, 0xe3, 0x78, 0x62, 0xa3, 0x47, 0x17,
	0xcc, 0xcb, 0xb5, 0xb0, 0x01, 0xfd, 0xf1, 0x34, 0x1a, 0x4f, 0xa3, 0xf0, 0xc9, 0xbe, 0x19, 0x7d,
	0xfa, 0x71, 0x70, 0xd1, 0xcf, 0x83, 0x8b, 0x7e, 0x1f, 0x5c, 0xb4, 0x7c, 0x57, 0xd6, 0xb2, 0xda,
	0xad, 0x86, 0x19, 0xdb, 0x06, 0x6d, 0x9a, 0x55, 0xfb, 0x9c, 0xf2, 0x7f, 0xd1, 0xd7, 0x30, 0x10,
	0x3c, 0x0b, 0xfe, 0x7f, 0x0f, 0xab, 0x3b, 0x75, 0x0a, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x44, 0x31, 0x38, 0xf9, 0x30, 0x02, 0x00, 0x00,
}

func (m *DataRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataRef) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataRef) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SizeBytes != 0 {
		i = encodeVarintChunk(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x20
	}
	if m.OffsetBytes != 0 {
		i = encodeVarintChunk(dAtA, i, uint64(m.OffsetBytes))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintChunk(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Ref != nil {
		{
			size, err := m.Ref.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintChunk(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Ref) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ref) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ref) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CompressionAlgo != 0 {
		i = encodeVarintChunk(dAtA, i, uint64(m.CompressionAlgo))
		i--
		dAtA[i] = 0x30
	}
	if m.EncryptionAlgo != 0 {
		i = encodeVarintChunk(dAtA, i, uint64(m.EncryptionAlgo))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Dek) > 0 {
		i -= len(m.Dek)
		copy(dAtA[i:], m.Dek)
		i = encodeVarintChunk(dAtA, i, uint64(len(m.Dek)))
		i--
		dAtA[i] = 0x22
	}
	if m.Edge {
		i--
		if m.Edge {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.SizeBytes != 0 {
		i = encodeVarintChunk(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintChunk(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChunk(dAtA []byte, offset int, v uint64) int {
	offset -= sovChunk(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ref != nil {
		l = m.Ref.Size()
		n += 1 + l + sovChunk(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovChunk(uint64(l))
	}
	if m.OffsetBytes != 0 {
		n += 1 + sovChunk(uint64(m.OffsetBytes))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovChunk(uint64(m.SizeBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Ref) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovChunk(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovChunk(uint64(m.SizeBytes))
	}
	if m.Edge {
		n += 2
	}
	l = len(m.Dek)
	if l > 0 {
		n += 1 + l + sovChunk(uint64(l))
	}
	if m.EncryptionAlgo != 0 {
		n += 1 + sovChunk(uint64(m.EncryptionAlgo))
	}
	if m.CompressionAlgo != 0 {
		n += 1 + sovChunk(uint64(m.CompressionAlgo))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovChunk(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChunk(x uint64) (n int) {
	return sovChunk(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChunk
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
			return fmt.Errorf("proto: DataRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthChunk
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChunk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ref == nil {
				m.Ref = &Ref{}
			}
			if err := m.Ref.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
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
				return ErrInvalidLengthChunk
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChunk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OffsetBytes", wireType)
			}
			m.OffsetBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OffsetBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChunk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChunk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ref) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChunk
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
			return fmt.Errorf("proto: Ref: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ref: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
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
				return ErrInvalidLengthChunk
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChunk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Edge", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
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
			m.Edge = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dek", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
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
				return ErrInvalidLengthChunk
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChunk
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dek = append(m.Dek[:0], dAtA[iNdEx:postIndex]...)
			if m.Dek == nil {
				m.Dek = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionAlgo", wireType)
			}
			m.EncryptionAlgo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EncryptionAlgo |= EncryptionAlgo(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompressionAlgo", wireType)
			}
			m.CompressionAlgo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChunk
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompressionAlgo |= CompressionAlgo(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChunk(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChunk
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChunk(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChunk
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
					return 0, ErrIntOverflowChunk
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
					return 0, ErrIntOverflowChunk
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
				return 0, ErrInvalidLengthChunk
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChunk
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChunk
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChunk        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChunk          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChunk = fmt.Errorf("proto: unexpected end of group")
)