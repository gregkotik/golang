// Code generated by protoc-gen-go. DO NOT EDIT.
// source: record.proto

package storerecords

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The request message containing the record
type StoreRecordRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	MobileNumber         string   `protobuf:"bytes,4,opt,name=mobileNumber,proto3" json:"mobileNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRecordRequest) Reset()         { *m = StoreRecordRequest{} }
func (m *StoreRecordRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRecordRequest) ProtoMessage()    {}
func (*StoreRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0}
}

func (m *StoreRecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRecordRequest.Unmarshal(m, b)
}
func (m *StoreRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRecordRequest.Marshal(b, m, deterministic)
}
func (m *StoreRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRecordRequest.Merge(m, src)
}
func (m *StoreRecordRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRecordRequest.Size(m)
}
func (m *StoreRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRecordRequest proto.InternalMessageInfo

func (m *StoreRecordRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StoreRecordRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoreRecordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *StoreRecordRequest) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

// The response message containing the ack
type StoreRecordReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRecordReply) Reset()         { *m = StoreRecordReply{} }
func (m *StoreRecordReply) String() string { return proto.CompactTextString(m) }
func (*StoreRecordReply) ProtoMessage()    {}
func (*StoreRecordReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{1}
}

func (m *StoreRecordReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRecordReply.Unmarshal(m, b)
}
func (m *StoreRecordReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRecordReply.Marshal(b, m, deterministic)
}
func (m *StoreRecordReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRecordReply.Merge(m, src)
}
func (m *StoreRecordReply) XXX_Size() int {
	return xxx_messageInfo_StoreRecordReply.Size(m)
}
func (m *StoreRecordReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRecordReply.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRecordReply proto.InternalMessageInfo

func (m *StoreRecordReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreRecordRequest)(nil), "storerecords.StoreRecordRequest")
	proto.RegisterType((*StoreRecordReply)(nil), "storerecords.StoreRecordReply")
}

func init() { proto.RegisterFile("record.proto", fileDescriptor_bf94fd919e302a1d) }

var fileDescriptor_bf94fd919e302a1d = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x8a, 0xc2, 0x30,
	0x14, 0x45, 0xa7, 0x9d, 0xce, 0x48, 0x9f, 0x41, 0xe4, 0xe1, 0x22, 0xb8, 0x90, 0x92, 0x55, 0x17,
	0xd2, 0x85, 0xfe, 0x87, 0x60, 0xfc, 0x82, 0xd4, 0x3e, 0x24, 0x90, 0x98, 0x9a, 0xb4, 0x8b, 0xfe,
	0xbd, 0x98, 0x22, 0xb4, 0x08, 0xee, 0x72, 0x4e, 0x2e, 0x5c, 0xee, 0x03, 0xe6, 0xe9, 0xea, 0x7c,
	0x53, 0xb5, 0xde, 0x75, 0x0e, 0x59, 0xe8, 0x9c, 0xa7, 0x51, 0x05, 0xe1, 0x01, 0x2f, 0x2f, 0x96,
	0x91, 0x25, 0x3d, 0x7a, 0x0a, 0x1d, 0xae, 0x20, 0xd5, 0x0d, 0x4f, 0x8a, 0xa4, 0xcc, 0x64, 0xaa,
	0x1b, 0x44, 0xc8, 0xee, 0xca, 0x12, 0x4f, 0x8b, 0xa4, 0xcc, 0x65, 0x7c, 0xe3, 0x06, 0xfe, 0xc8,
	0x2a, 0x6d, 0xf8, 0x6f, 0x94, 0x23, 0xa0, 0x00, 0x66, 0x5d, 0xad, 0x0d, 0x9d, 0x7a, 0x5b, 0x93,
	0xe7, 0x59, 0xfc, 0x9c, 0x39, 0xb1, 0x87, 0xf5, 0xac, 0xb3, 0x35, 0x03, 0x72, 0x58, 0x58, 0x0a,
	0x41, 0xdd, 0x28, 0xd6, 0xe6, 0xf2, 0x8d, 0x07, 0x05, 0x6c, 0x92, 0x0e, 0x78, 0x86, 0xe5, 0x84,
	0xb1, 0xa8, 0xa6, 0x7b, 0xaa, 0xcf, 0x31, 0xdb, 0xdd, 0x97, 0x44, 0x6b, 0x06, 0xf1, 0x53, 0xff,
	0xc7, 0xcb, 0x1c, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0xeb, 0x1b, 0x47, 0x29, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreRecordsClient is the client API for StoreRecords service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreRecordsClient interface {
	// Sends a record
	StoreRecord(ctx context.Context, in *StoreRecordRequest, opts ...grpc.CallOption) (*StoreRecordReply, error)
}

type storeRecordsClient struct {
	cc *grpc.ClientConn
}

func NewStoreRecordsClient(cc *grpc.ClientConn) StoreRecordsClient {
	return &storeRecordsClient{cc}
}

func (c *storeRecordsClient) StoreRecord(ctx context.Context, in *StoreRecordRequest, opts ...grpc.CallOption) (*StoreRecordReply, error) {
	out := new(StoreRecordReply)
	err := c.cc.Invoke(ctx, "/storerecords.StoreRecords/StoreRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreRecordsServer is the server API for StoreRecords service.
type StoreRecordsServer interface {
	// Sends a record
	StoreRecord(context.Context, *StoreRecordRequest) (*StoreRecordReply, error)
}

func RegisterStoreRecordsServer(s *grpc.Server, srv StoreRecordsServer) {
	s.RegisterService(&_StoreRecords_serviceDesc, srv)
}

func _StoreRecords_StoreRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreRecordsServer).StoreRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storerecords.StoreRecords/StoreRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreRecordsServer).StoreRecord(ctx, req.(*StoreRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreRecords_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storerecords.StoreRecords",
	HandlerType: (*StoreRecordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreRecord",
			Handler:    _StoreRecords_StoreRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record.proto",
}
