// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/server/service.proto

package server

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	pps "github.com/pachyderm/pachyderm/v2/src/pps"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CancelRequest struct {
	JobID                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	DataFilters          []string `protobuf:"bytes,1,rep,name=data_filters,json=dataFilters,proto3" json:"data_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelRequest) Reset()         { *m = CancelRequest{} }
func (m *CancelRequest) String() string { return proto.CompactTextString(m) }
func (*CancelRequest) ProtoMessage()    {}
func (*CancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4407c0c45dc0204, []int{0}
}
func (m *CancelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelRequest.Merge(m, src)
}
func (m *CancelRequest) XXX_Size() int {
	return m.Size()
}
func (m *CancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelRequest proto.InternalMessageInfo

func (m *CancelRequest) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *CancelRequest) GetDataFilters() []string {
	if m != nil {
		return m.DataFilters
	}
	return nil
}

type CancelResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelResponse) Reset()         { *m = CancelResponse{} }
func (m *CancelResponse) String() string { return proto.CompactTextString(m) }
func (*CancelResponse) ProtoMessage()    {}
func (*CancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4407c0c45dc0204, []int{1}
}
func (m *CancelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CancelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResponse.Merge(m, src)
}
func (m *CancelResponse) XXX_Size() int {
	return m.Size()
}
func (m *CancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResponse proto.InternalMessageInfo

func (m *CancelResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*CancelRequest)(nil), "server.CancelRequest")
	proto.RegisterType((*CancelResponse)(nil), "server.CancelResponse")
}

func init() {
	proto.RegisterFile("server/worker/server/service.proto", fileDescriptor_c4407c0c45dc0204)
}

var fileDescriptor_c4407c0c45dc0204 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x5f, 0x4b, 0xf3, 0x30,
	0x18, 0xc5, 0xd7, 0xf7, 0x65, 0xd5, 0x45, 0x27, 0x18, 0x74, 0x8c, 0x0a, 0x73, 0xf6, 0x6a, 0x78,
	0x91, 0xc0, 0xc4, 0x0b, 0xbd, 0x9c, 0x7f, 0x70, 0x5e, 0x56, 0x41, 0xf0, 0x66, 0xb4, 0xe9, 0xb3,
	0xae, 0x73, 0xdb, 0x13, 0x93, 0x74, 0x32, 0x3f, 0xa1, 0x97, 0x7e, 0x02, 0x91, 0x7e, 0x12, 0x69,
	0xb3, 0x82, 0x8a, 0x57, 0x79, 0xce, 0x2f, 0x87, 0xe4, 0x9c, 0x84, 0xf8, 0x1a, 0xd4, 0x12, 0x14,
	0x7f, 0x41, 0xf5, 0x04, 0x8a, 0xaf, 0x55, 0xb1, 0xa4, 0x02, 0x98, 0x54, 0x68, 0x90, 0xba, 0x96,
	0x7a, 0x4d, 0x29, 0x35, 0x97, 0x52, 0x5b, 0xec, 0xed, 0x25, 0x98, 0x60, 0x39, 0xf2, 0x62, 0x5a,
	0xd3, 0x83, 0x04, 0x31, 0x99, 0x01, 0x2f, 0x55, 0x94, 0x8d, 0x39, 0xcc, 0xa5, 0x59, 0xd9, 0x4d,
	0xff, 0x9e, 0x34, 0x2f, 0xc2, 0x85, 0x80, 0x59, 0x00, 0xcf, 0x19, 0x68, 0x43, 0xbb, 0xc4, 0x9d,
	0x62, 0x34, 0x4a, 0xe3, 0xf6, 0xbf, 0xae, 0xd3, 0x6b, 0x0c, 0x1a, 0xf9, 0xc7, 0x61, 0xfd, 0x16,
	0xa3, 0xe1, 0x65, 0x50, 0x9f, 0x62, 0x34, 0x8c, 0xe9, 0x11, 0xd9, 0x8e, 0x43, 0x13, 0x8e, 0xc6,
	0xe9, 0xcc, 0x80, 0xd2, 0x6d, 0xa7, 0xfb, 0xbf, 0xd7, 0x08, 0xb6, 0x0a, 0x76, 0x6d, 0x91, 0x7f,
	0x4c, 0x76, 0xaa, 0x53, 0xb5, 0xc4, 0x85, 0x06, 0xda, 0x26, 0x1b, 0x3a, 0x13, 0x02, 0x74, 0xe1,
	0x77, 0x7a, 0x9b, 0x41, 0x25, 0xfb, 0xaf, 0xc4, 0x7d, 0x28, 0xab, 0xd2, 0x53, 0xe2, 0xde, 0x99,
	0xd0, 0x64, 0x9a, 0xb6, 0x98, 0xcd, 0xcc, 0xaa, 0xcc, 0xec, 0xaa, 0xc8, 0xec, 0xed, 0xb2, 0xa2,
	0xac, 0xb5, 0x5b, 0xab, 0x5f, 0xa3, 0x67, 0xc4, 0xb5, 0x97, 0xd1, 0x7d, 0x66, 0xdf, 0x85, 0xfd,
	0xa8, 0xe4, 0xb5, 0x7e, 0x63, 0x9b, 0xc9, 0xaf, 0x0d, 0x6e, 0xde, 0xf2, 0x8e, 0xf3, 0x9e, 0x77,
	0x9c, 0xcf, 0xbc, 0xe3, 0x3c, 0x9e, 0x27, 0xa9, 0x99, 0x64, 0x11, 0x13, 0x38, 0xe7, 0x32, 0x14,
	0x93, 0x55, 0x0c, 0xea, 0xfb, 0xb4, 0xec, 0x73, 0xad, 0x04, 0xff, 0xeb, 0x7f, 0x22, 0xb7, 0x4c,
	0x7a, 0xf2, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xf3, 0xda, 0x1a, 0xbe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	Status(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error)
	Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Status(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*pps.WorkerStatus, error) {
	out := new(pps.WorkerStatus)
	err := c.cc.Invoke(ctx, "/server.Worker/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Cancel(ctx context.Context, in *CancelRequest, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, "/server.Worker/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	Status(context.Context, *types.Empty) (*pps.WorkerStatus, error)
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)
}

// UnimplementedWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (*UnimplementedWorkerServer) Status(ctx context.Context, req *types.Empty) (*pps.WorkerStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedWorkerServer) Cancel(ctx context.Context, req *CancelRequest) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Worker/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Status(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Worker/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Cancel(ctx, req.(*CancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Worker_Status_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Worker_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/worker/server/service.proto",
}

func (m *CancelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintService(dAtA, i, uint64(len(m.JobID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DataFilters) > 0 {
		for iNdEx := len(m.DataFilters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DataFilters[iNdEx])
			copy(dAtA[i:], m.DataFilters[iNdEx])
			i = encodeVarintService(dAtA, i, uint64(len(m.DataFilters[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CancelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CancelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CancelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DataFilters) > 0 {
		for _, s := range m.DataFilters {
			l = len(s)
			n += 1 + l + sovService(uint64(l))
		}
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CancelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CancelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: CancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataFilters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataFilters = append(m.DataFilters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *CancelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: CancelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)