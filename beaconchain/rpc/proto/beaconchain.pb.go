// Code generated by protoc-gen-go. DO NOT EDIT.
// source: beaconchain.proto

package beaconchain

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

// FetchLeadersRequest is the request to fetch the current leaders.
type FetchLeadersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchLeadersRequest) Reset()         { *m = FetchLeadersRequest{} }
func (m *FetchLeadersRequest) String() string { return proto.CompactTextString(m) }
func (*FetchLeadersRequest) ProtoMessage()    {}
func (*FetchLeadersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_474fd8061d1037cf, []int{0}
}

func (m *FetchLeadersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchLeadersRequest.Unmarshal(m, b)
}
func (m *FetchLeadersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchLeadersRequest.Marshal(b, m, deterministic)
}
func (m *FetchLeadersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchLeadersRequest.Merge(m, src)
}
func (m *FetchLeadersRequest) XXX_Size() int {
	return xxx_messageInfo_FetchLeadersRequest.Size(m)
}
func (m *FetchLeadersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchLeadersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchLeadersRequest proto.InternalMessageInfo

// FetchLeadersResponse is the response of FetchLeadersRequest.
type FetchLeadersResponse struct {
	Leaders              []*FetchLeadersResponse_Leader `protobuf:"bytes,1,rep,name=leaders,proto3" json:"leaders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *FetchLeadersResponse) Reset()         { *m = FetchLeadersResponse{} }
func (m *FetchLeadersResponse) String() string { return proto.CompactTextString(m) }
func (*FetchLeadersResponse) ProtoMessage()    {}
func (*FetchLeadersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_474fd8061d1037cf, []int{1}
}

func (m *FetchLeadersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchLeadersResponse.Unmarshal(m, b)
}
func (m *FetchLeadersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchLeadersResponse.Marshal(b, m, deterministic)
}
func (m *FetchLeadersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchLeadersResponse.Merge(m, src)
}
func (m *FetchLeadersResponse) XXX_Size() int {
	return xxx_messageInfo_FetchLeadersResponse.Size(m)
}
func (m *FetchLeadersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchLeadersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchLeadersResponse proto.InternalMessageInfo

func (m *FetchLeadersResponse) GetLeaders() []*FetchLeadersResponse_Leader {
	if m != nil {
		return m.Leaders
	}
	return nil
}

type FetchLeadersResponse_Leader struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ShardId              uint32   `protobuf:"varint,3,opt,name=shardId,proto3" json:"shardId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchLeadersResponse_Leader) Reset()         { *m = FetchLeadersResponse_Leader{} }
func (m *FetchLeadersResponse_Leader) String() string { return proto.CompactTextString(m) }
func (*FetchLeadersResponse_Leader) ProtoMessage()    {}
func (*FetchLeadersResponse_Leader) Descriptor() ([]byte, []int) {
	return fileDescriptor_474fd8061d1037cf, []int{1, 0}
}

func (m *FetchLeadersResponse_Leader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchLeadersResponse_Leader.Unmarshal(m, b)
}
func (m *FetchLeadersResponse_Leader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchLeadersResponse_Leader.Marshal(b, m, deterministic)
}
func (m *FetchLeadersResponse_Leader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchLeadersResponse_Leader.Merge(m, src)
}
func (m *FetchLeadersResponse_Leader) XXX_Size() int {
	return xxx_messageInfo_FetchLeadersResponse_Leader.Size(m)
}
func (m *FetchLeadersResponse_Leader) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchLeadersResponse_Leader.DiscardUnknown(m)
}

var xxx_messageInfo_FetchLeadersResponse_Leader proto.InternalMessageInfo

func (m *FetchLeadersResponse_Leader) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *FetchLeadersResponse_Leader) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *FetchLeadersResponse_Leader) GetShardId() uint32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func init() {
	proto.RegisterType((*FetchLeadersRequest)(nil), "beaconchain.FetchLeadersRequest")
	proto.RegisterType((*FetchLeadersResponse)(nil), "beaconchain.FetchLeadersResponse")
	proto.RegisterType((*FetchLeadersResponse_Leader)(nil), "beaconchain.FetchLeadersResponse.Leader")
}

func init() { proto.RegisterFile("beaconchain.proto", fileDescriptor_474fd8061d1037cf) }

var fileDescriptor_474fd8061d1037cf = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4a, 0x4d, 0x4c,
	0xce, 0xcf, 0x4b, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x12, 0xe5, 0x12, 0x76, 0x4b, 0x2d, 0x49, 0xce, 0xf0, 0x49, 0x4d, 0x4c, 0x49, 0x2d,
	0x2a, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x5a, 0xc4, 0xc8, 0x25, 0x82, 0x2a, 0x5e,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xe4, 0xc4, 0xc5, 0x9e, 0x03, 0x11, 0x92, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x36, 0xd2, 0xd0, 0x43, 0xb6, 0x01, 0x9b, 0x1e, 0x3d, 0x08, 0x3f, 0x08, 0xa6, 0x51,
	0xca, 0x8d, 0x8b, 0x0d, 0x22, 0x24, 0xc4, 0xc7, 0xc5, 0x94, 0x59, 0x20, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x94, 0x59, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x90, 0x5f, 0x54, 0x22, 0xc1, 0x04,
	0x16, 0x01, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0x33, 0x12, 0x8b, 0x52, 0x3c, 0x53, 0x24, 0x98,
	0x15, 0x18, 0x35, 0x78, 0x83, 0x60, 0x5c, 0xa3, 0x6c, 0x2e, 0x21, 0x27, 0xb0, 0xdd, 0xce, 0x20,
	0xbb, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x42, 0xb9, 0x78, 0x90, 0x5d, 0x21, 0xa4,
	0x80, 0xc7, 0x81, 0x60, 0xcf, 0x4a, 0x29, 0x12, 0xf4, 0x82, 0x12, 0x43, 0x12, 0x1b, 0x38, 0xf0,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0x0b, 0x9f, 0xda, 0x51, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconChainServiceClient is the client API for BeaconChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconChainServiceClient interface {
	FetchLeaders(ctx context.Context, in *FetchLeadersRequest, opts ...grpc.CallOption) (*FetchLeadersResponse, error)
}

type beaconChainServiceClient struct {
	cc *grpc.ClientConn
}

func NewBeaconChainServiceClient(cc *grpc.ClientConn) BeaconChainServiceClient {
	return &beaconChainServiceClient{cc}
}

func (c *beaconChainServiceClient) FetchLeaders(ctx context.Context, in *FetchLeadersRequest, opts ...grpc.CallOption) (*FetchLeadersResponse, error) {
	out := new(FetchLeadersResponse)
	err := c.cc.Invoke(ctx, "/beaconchain.BeaconChainService/FetchLeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconChainServiceServer is the server API for BeaconChainService service.
type BeaconChainServiceServer interface {
	FetchLeaders(context.Context, *FetchLeadersRequest) (*FetchLeadersResponse, error)
}

func RegisterBeaconChainServiceServer(s *grpc.Server, srv BeaconChainServiceServer) {
	s.RegisterService(&_BeaconChainService_serviceDesc, srv)
}

func _BeaconChainService_FetchLeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconChainServiceServer).FetchLeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beaconchain.BeaconChainService/FetchLeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconChainServiceServer).FetchLeaders(ctx, req.(*FetchLeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconChainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beaconchain.BeaconChainService",
	HandlerType: (*BeaconChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchLeaders",
			Handler:    _BeaconChainService_FetchLeaders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beaconchain.proto",
}
