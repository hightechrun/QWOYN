// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aquifer/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgDeposit struct {
	Sender string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgDeposit) Reset()         { *m = MsgDeposit{} }
func (m *MsgDeposit) String() string { return proto.CompactTextString(m) }
func (*MsgDeposit) ProtoMessage()    {}
func (*MsgDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_29346906793f3475, []int{0}
}
func (m *MsgDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeposit.Merge(m, src)
}
func (m *MsgDeposit) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeposit proto.InternalMessageInfo

func (m *MsgDeposit) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgDeposit) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgDepositDepositResponse struct {
}

func (m *MsgDepositDepositResponse) Reset()         { *m = MsgDepositDepositResponse{} }
func (m *MsgDepositDepositResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositDepositResponse) ProtoMessage()    {}
func (*MsgDepositDepositResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_29346906793f3475, []int{1}
}
func (m *MsgDepositDepositResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositDepositResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositDepositResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositDepositResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositDepositResponse.Merge(m, src)
}
func (m *MsgDepositDepositResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositDepositResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositDepositResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositDepositResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDeposit)(nil), "cosmichorizon.qwoyn.aquifer.MsgDeposit")
	proto.RegisterType((*MsgDepositDepositResponse)(nil), "cosmichorizon.qwoyn.aquifer.MsgDepositDepositResponse")
}

func init() { proto.RegisterFile("aquifer/tx.proto", fileDescriptor_29346906793f3475) }

var fileDescriptor_29346906793f3475 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x1b, 0x95, 0x8a, 0xf1, 0x22, 0x45, 0x64, 0xeb, 0x20, 0x8e, 0x5d, 0xdc, 0xc5, 0x84,
	0x4d, 0xd0, 0xfb, 0xf4, 0xe2, 0x61, 0x97, 0x1e, 0x05, 0x0f, 0x6d, 0x8d, 0x5d, 0x84, 0xe6, 0x6b,
	0xfb, 0xa5, 0xba, 0xf9, 0x2b, 0xfc, 0x59, 0x3b, 0xee, 0xe8, 0x49, 0xa4, 0xfd, 0x23, 0xd2, 0x36,
	0x3a, 0xf0, 0x20, 0x9e, 0xf2, 0x26, 0x79, 0xf3, 0x3c, 0x24, 0xa1, 0x47, 0x61, 0x5e, 0xaa, 0x47,
	0x59, 0x08, 0xb3, 0xe4, 0x59, 0x01, 0x06, 0xbc, 0x41, 0x0c, 0x98, 0xaa, 0x78, 0x01, 0x85, 0x7a,
	0x05, 0xcd, 0xf3, 0x17, 0x58, 0x69, 0x6e, 0x5b, 0x3e, 0x6b, 0x36, 0x01, 0x45, 0x14, 0xa2, 0x14,
	0xcf, 0x93, 0x48, 0x9a, 0x70, 0x22, 0x62, 0x50, 0xba, 0x3b, 0xec, 0x1f, 0x27, 0x90, 0x40, 0x1b,
	0x45, 0x93, 0xba, 0xd5, 0xd1, 0x3d, 0xa5, 0x73, 0x4c, 0x6e, 0x64, 0x06, 0xa8, 0x8c, 0x77, 0x42,
	0x5d, 0x94, 0xfa, 0x41, 0x16, 0x3d, 0x32, 0x24, 0xe3, 0x83, 0xc0, 0xce, 0xbc, 0x2b, 0xea, 0x86,
	0x29, 0x94, 0xda, 0xf4, 0x76, 0x86, 0x64, 0x7c, 0x38, 0xed, 0xf3, 0x4e, 0xc6, 0x1b, 0x19, 0xb7,
	0x32, 0x7e, 0x0d, 0x4a, 0xcf, 0xf6, 0xd6, 0x1f, 0xa7, 0x4e, 0x60, 0xeb, 0xa3, 0x01, 0xed, 0x6f,
	0xf1, 0x76, 0x08, 0x24, 0x66, 0xa0, 0x51, 0x4e, 0x73, 0xba, 0x3b, 0xc7, 0xc4, 0x7b, 0xa2, 0xfb,
	0xdf, 0xfe, 0x33, 0xfe, 0xc7, 0x0d, 0xf9, 0x96, 0xe4, 0x5f, 0xfe, 0xb3, 0xf8, 0x4b, 0x39, 0xbb,
	0x5d, 0x57, 0x8c, 0x6c, 0x2a, 0x46, 0x3e, 0x2b, 0x46, 0xde, 0x6a, 0xe6, 0x6c, 0x6a, 0xe6, 0xbc,
	0xd7, 0xcc, 0xb9, 0x13, 0x89, 0x32, 0x8b, 0x32, 0xe2, 0x31, 0xa4, 0xa2, 0x63, 0x9f, 0x5b, 0xb8,
	0x68, 0xe1, 0x62, 0x29, 0x7e, 0xfe, 0x63, 0x95, 0x49, 0x8c, 0xdc, 0xf6, 0x01, 0x2f, 0xbe, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x69, 0x3e, 0x03, 0x59, 0xa7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositDepositResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositDepositResponse, error) {
	out := new(MsgDepositDepositResponse)
	err := c.cc.Invoke(ctx, "/cosmichorizon.qwoyn.aquifer.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Deposit(context.Context, *MsgDeposit) (*MsgDepositDepositResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Deposit(ctx context.Context, req *MsgDeposit) (*MsgDepositDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmichorizon.qwoyn.aquifer.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmichorizon.qwoyn.aquifer.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aquifer/tx.proto",
}

func (m *MsgDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositDepositResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositDepositResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositDepositResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgDepositDepositResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDepositDepositResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDepositDepositResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositDepositResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
