// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

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

// ServiceType indicates which service used to generate this message.
type ServiceType int32

const (
	ServiceType_CONSENSUS      ServiceType = 0
	ServiceType_STAKING        ServiceType = 1
	ServiceType_DRAND          ServiceType = 2
	ServiceType_CLIENT_SUPPORT ServiceType = 3
)

var ServiceType_name = map[int32]string{
	0: "CONSENSUS",
	1: "STAKING",
	2: "DRAND",
	3: "CLIENT_SUPPORT",
}

var ServiceType_value = map[string]int32{
	"CONSENSUS":      0,
	"STAKING":        1,
	"DRAND":          2,
	"CLIENT_SUPPORT": 3,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}

func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

// MessageType indicates what is the type of this message.
type MessageType int32

const (
	MessageType_NEWNODE_BEACON_STAKING MessageType = 0
	MessageType_ANNOUNCE               MessageType = 1
	MessageType_PREPARE                MessageType = 2
	MessageType_PREPARED               MessageType = 3
	MessageType_COMMIT                 MessageType = 4
	MessageType_COMMITTED              MessageType = 5
	MessageType_VIEWCHANGE             MessageType = 6
	MessageType_NEWVIEW                MessageType = 7
	MessageType_DRAND_INIT             MessageType = 10
	MessageType_DRAND_COMMIT           MessageType = 11
	MessageType_LOTTERY_REQUEST        MessageType = 12
)

var MessageType_name = map[int32]string{
	0:  "NEWNODE_BEACON_STAKING",
	1:  "ANNOUNCE",
	2:  "PREPARE",
	3:  "PREPARED",
	4:  "COMMIT",
	5:  "COMMITTED",
	6:  "VIEWCHANGE",
	7:  "NEWVIEW",
	10: "DRAND_INIT",
	11: "DRAND_COMMIT",
	12: "LOTTERY_REQUEST",
}

var MessageType_value = map[string]int32{
	"NEWNODE_BEACON_STAKING": 0,
	"ANNOUNCE":               1,
	"PREPARE":                2,
	"PREPARED":               3,
	"COMMIT":                 4,
	"COMMITTED":              5,
	"VIEWCHANGE":             6,
	"NEWVIEW":                7,
	"DRAND_INIT":             10,
	"DRAND_COMMIT":           11,
	"LOTTERY_REQUEST":        12,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

type LotteryRequest_Type int32

const (
	LotteryRequest_ENTER       LotteryRequest_Type = 0
	LotteryRequest_RESULT      LotteryRequest_Type = 1
	LotteryRequest_PICK_WINNER LotteryRequest_Type = 2
)

var LotteryRequest_Type_name = map[int32]string{
	0: "ENTER",
	1: "RESULT",
	2: "PICK_WINNER",
}

var LotteryRequest_Type_value = map[string]int32{
	"ENTER":       0,
	"RESULT":      1,
	"PICK_WINNER": 2,
}

func (x LotteryRequest_Type) String() string {
	return proto.EnumName(LotteryRequest_Type_name, int32(x))
}

func (LotteryRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3, 0}
}

// This is universal message for all communication protocols.
// There are different Requests for different message types.
// As we introduce a new type of message just add a new MessageType and new type of request in Message.
//
// The request field will be either one of the structure corresponding to the MessageType type.
type Message struct {
	ServiceType ServiceType `protobuf:"varint,1,opt,name=service_type,json=serviceType,proto3,enum=message.ServiceType" json:"service_type,omitempty"`
	Type        MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=message.MessageType" json:"type,omitempty"`
	Signature   []byte      `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*Message_Staking
	//	*Message_Consensus
	//	*Message_Drand
	//	*Message_Viewchange
	//	*Message_LotteryRequest
	Request              isMessage_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetServiceType() ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return ServiceType_CONSENSUS
}

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_NEWNODE_BEACON_STAKING
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type isMessage_Request interface {
	isMessage_Request()
}

type Message_Staking struct {
	Staking *StakingRequest `protobuf:"bytes,4,opt,name=staking,proto3,oneof"`
}

type Message_Consensus struct {
	Consensus *ConsensusRequest `protobuf:"bytes,5,opt,name=consensus,proto3,oneof"`
}

type Message_Drand struct {
	Drand *DrandRequest `protobuf:"bytes,6,opt,name=drand,proto3,oneof"`
}

type Message_Viewchange struct {
	Viewchange *ViewChangeRequest `protobuf:"bytes,7,opt,name=viewchange,proto3,oneof"`
}

type Message_LotteryRequest struct {
	LotteryRequest *LotteryRequest `protobuf:"bytes,8,opt,name=lottery_request,json=lotteryRequest,proto3,oneof"`
}

func (*Message_Staking) isMessage_Request() {}

func (*Message_Consensus) isMessage_Request() {}

func (*Message_Drand) isMessage_Request() {}

func (*Message_Viewchange) isMessage_Request() {}

func (*Message_LotteryRequest) isMessage_Request() {}

func (m *Message) GetRequest() isMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Message) GetStaking() *StakingRequest {
	if x, ok := m.GetRequest().(*Message_Staking); ok {
		return x.Staking
	}
	return nil
}

func (m *Message) GetConsensus() *ConsensusRequest {
	if x, ok := m.GetRequest().(*Message_Consensus); ok {
		return x.Consensus
	}
	return nil
}

func (m *Message) GetDrand() *DrandRequest {
	if x, ok := m.GetRequest().(*Message_Drand); ok {
		return x.Drand
	}
	return nil
}

func (m *Message) GetViewchange() *ViewChangeRequest {
	if x, ok := m.GetRequest().(*Message_Viewchange); ok {
		return x.Viewchange
	}
	return nil
}

func (m *Message) GetLotteryRequest() *LotteryRequest {
	if x, ok := m.GetRequest().(*Message_LotteryRequest); ok {
		return x.LotteryRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Staking)(nil),
		(*Message_Consensus)(nil),
		(*Message_Drand)(nil),
		(*Message_Viewchange)(nil),
		(*Message_LotteryRequest)(nil),
	}
}

type Response struct {
	ServiceType ServiceType `protobuf:"varint,1,opt,name=service_type,json=serviceType,proto3,enum=message.ServiceType" json:"service_type,omitempty"`
	Type        MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=message.MessageType" json:"type,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*Response_LotteryResponse
	Response             isResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetServiceType() ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return ServiceType_CONSENSUS
}

func (m *Response) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_NEWNODE_BEACON_STAKING
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_LotteryResponse struct {
	LotteryResponse *LotteryResponse `protobuf:"bytes,3,opt,name=lottery_response,json=lotteryResponse,proto3,oneof"`
}

func (*Response_LotteryResponse) isResponse_Response() {}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetLotteryResponse() *LotteryResponse {
	if x, ok := m.GetResponse().(*Response_LotteryResponse); ok {
		return x.LotteryResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_LotteryResponse)(nil),
	}
}

type LotteryResponse struct {
	Players              []string `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`
	Balances             []string `protobuf:"bytes,3,rep,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LotteryResponse) Reset()         { *m = LotteryResponse{} }
func (m *LotteryResponse) String() string { return proto.CompactTextString(m) }
func (*LotteryResponse) ProtoMessage()    {}
func (*LotteryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *LotteryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LotteryResponse.Unmarshal(m, b)
}
func (m *LotteryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LotteryResponse.Marshal(b, m, deterministic)
}
func (m *LotteryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LotteryResponse.Merge(m, src)
}
func (m *LotteryResponse) XXX_Size() int {
	return xxx_messageInfo_LotteryResponse.Size(m)
}
func (m *LotteryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LotteryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LotteryResponse proto.InternalMessageInfo

func (m *LotteryResponse) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *LotteryResponse) GetBalances() []string {
	if m != nil {
		return m.Balances
	}
	return nil
}

type LotteryRequest struct {
	Type                 LotteryRequest_Type `protobuf:"varint,1,opt,name=type,proto3,enum=message.LotteryRequest_Type" json:"type,omitempty"`
	PrivateKey           string              `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Amount               int64               `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LotteryRequest) Reset()         { *m = LotteryRequest{} }
func (m *LotteryRequest) String() string { return proto.CompactTextString(m) }
func (*LotteryRequest) ProtoMessage()    {}
func (*LotteryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *LotteryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LotteryRequest.Unmarshal(m, b)
}
func (m *LotteryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LotteryRequest.Marshal(b, m, deterministic)
}
func (m *LotteryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LotteryRequest.Merge(m, src)
}
func (m *LotteryRequest) XXX_Size() int {
	return xxx_messageInfo_LotteryRequest.Size(m)
}
func (m *LotteryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LotteryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LotteryRequest proto.InternalMessageInfo

func (m *LotteryRequest) GetType() LotteryRequest_Type {
	if m != nil {
		return m.Type
	}
	return LotteryRequest_ENTER
}

func (m *LotteryRequest) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *LotteryRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// Staking Request from new node to beacon node.
type StakingRequest struct {
	Transaction          []byte   `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakingRequest) Reset()         { *m = StakingRequest{} }
func (m *StakingRequest) String() string { return proto.CompactTextString(m) }
func (*StakingRequest) ProtoMessage()    {}
func (*StakingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *StakingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakingRequest.Unmarshal(m, b)
}
func (m *StakingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakingRequest.Marshal(b, m, deterministic)
}
func (m *StakingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakingRequest.Merge(m, src)
}
func (m *StakingRequest) XXX_Size() int {
	return xxx_messageInfo_StakingRequest.Size(m)
}
func (m *StakingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakingRequest proto.InternalMessageInfo

func (m *StakingRequest) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *StakingRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type ConsensusRequest struct {
	ViewId               uint32   `protobuf:"varint,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
	BlockNum             uint64   `protobuf:"varint,2,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	SenderPubkey         []byte   `protobuf:"bytes,4,opt,name=sender_pubkey,json=senderPubkey,proto3" json:"sender_pubkey,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusRequest) Reset()         { *m = ConsensusRequest{} }
func (m *ConsensusRequest) String() string { return proto.CompactTextString(m) }
func (*ConsensusRequest) ProtoMessage()    {}
func (*ConsensusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *ConsensusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusRequest.Unmarshal(m, b)
}
func (m *ConsensusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusRequest.Marshal(b, m, deterministic)
}
func (m *ConsensusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusRequest.Merge(m, src)
}
func (m *ConsensusRequest) XXX_Size() int {
	return xxx_messageInfo_ConsensusRequest.Size(m)
}
func (m *ConsensusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusRequest proto.InternalMessageInfo

func (m *ConsensusRequest) GetViewId() uint32 {
	if m != nil {
		return m.ViewId
	}
	return 0
}

func (m *ConsensusRequest) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *ConsensusRequest) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ConsensusRequest) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *ConsensusRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DrandRequest struct {
	SenderPubkey         []byte   `protobuf:"bytes,1,opt,name=sender_pubkey,json=senderPubkey,proto3" json:"sender_pubkey,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrandRequest) Reset()         { *m = DrandRequest{} }
func (m *DrandRequest) String() string { return proto.CompactTextString(m) }
func (*DrandRequest) ProtoMessage()    {}
func (*DrandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *DrandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrandRequest.Unmarshal(m, b)
}
func (m *DrandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrandRequest.Marshal(b, m, deterministic)
}
func (m *DrandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrandRequest.Merge(m, src)
}
func (m *DrandRequest) XXX_Size() int {
	return xxx_messageInfo_DrandRequest.Size(m)
}
func (m *DrandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DrandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DrandRequest proto.InternalMessageInfo

func (m *DrandRequest) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *DrandRequest) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *DrandRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ViewChangeRequest struct {
	ViewId        uint32 `protobuf:"varint,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
	BlockNum      uint64 `protobuf:"varint,2,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	SenderPubkey  []byte `protobuf:"bytes,3,opt,name=sender_pubkey,json=senderPubkey,proto3" json:"sender_pubkey,omitempty"`
	LeaderPubkey  []byte `protobuf:"bytes,4,opt,name=leader_pubkey,json=leaderPubkey,proto3" json:"leader_pubkey,omitempty"`
	Payload       []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	ViewchangeSig []byte `protobuf:"bytes,6,opt,name=viewchange_sig,json=viewchangeSig,proto3" json:"viewchange_sig,omitempty"`
	// below is for newview message only
	M1Aggsigs            []byte   `protobuf:"bytes,7,opt,name=m1_aggsigs,json=m1Aggsigs,proto3" json:"m1_aggsigs,omitempty"`
	M1Bitmap             []byte   `protobuf:"bytes,8,opt,name=m1_bitmap,json=m1Bitmap,proto3" json:"m1_bitmap,omitempty"`
	M2Aggsigs            []byte   `protobuf:"bytes,9,opt,name=m2_aggsigs,json=m2Aggsigs,proto3" json:"m2_aggsigs,omitempty"`
	M2Bitmap             []byte   `protobuf:"bytes,10,opt,name=m2_bitmap,json=m2Bitmap,proto3" json:"m2_bitmap,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewChangeRequest) Reset()         { *m = ViewChangeRequest{} }
func (m *ViewChangeRequest) String() string { return proto.CompactTextString(m) }
func (*ViewChangeRequest) ProtoMessage()    {}
func (*ViewChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *ViewChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewChangeRequest.Unmarshal(m, b)
}
func (m *ViewChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewChangeRequest.Marshal(b, m, deterministic)
}
func (m *ViewChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewChangeRequest.Merge(m, src)
}
func (m *ViewChangeRequest) XXX_Size() int {
	return xxx_messageInfo_ViewChangeRequest.Size(m)
}
func (m *ViewChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewChangeRequest proto.InternalMessageInfo

func (m *ViewChangeRequest) GetViewId() uint32 {
	if m != nil {
		return m.ViewId
	}
	return 0
}

func (m *ViewChangeRequest) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *ViewChangeRequest) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *ViewChangeRequest) GetLeaderPubkey() []byte {
	if m != nil {
		return m.LeaderPubkey
	}
	return nil
}

func (m *ViewChangeRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ViewChangeRequest) GetViewchangeSig() []byte {
	if m != nil {
		return m.ViewchangeSig
	}
	return nil
}

func (m *ViewChangeRequest) GetM1Aggsigs() []byte {
	if m != nil {
		return m.M1Aggsigs
	}
	return nil
}

func (m *ViewChangeRequest) GetM1Bitmap() []byte {
	if m != nil {
		return m.M1Bitmap
	}
	return nil
}

func (m *ViewChangeRequest) GetM2Aggsigs() []byte {
	if m != nil {
		return m.M2Aggsigs
	}
	return nil
}

func (m *ViewChangeRequest) GetM2Bitmap() []byte {
	if m != nil {
		return m.M2Bitmap
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.ServiceType", ServiceType_name, ServiceType_value)
	proto.RegisterEnum("message.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("message.LotteryRequest_Type", LotteryRequest_Type_name, LotteryRequest_Type_value)
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*Response)(nil), "message.Response")
	proto.RegisterType((*LotteryResponse)(nil), "message.LotteryResponse")
	proto.RegisterType((*LotteryRequest)(nil), "message.LotteryRequest")
	proto.RegisterType((*StakingRequest)(nil), "message.StakingRequest")
	proto.RegisterType((*ConsensusRequest)(nil), "message.ConsensusRequest")
	proto.RegisterType((*DrandRequest)(nil), "message.DrandRequest")
	proto.RegisterType((*ViewChangeRequest)(nil), "message.ViewChangeRequest")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 903 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcd, 0x8e, 0xe2, 0x46,
	0x10, 0xc7, 0x31, 0x30, 0x18, 0x97, 0x0d, 0xf4, 0x76, 0x92, 0x5d, 0x67, 0xb2, 0x51, 0x10, 0xab,
	0x48, 0x68, 0xa5, 0x8c, 0x76, 0xd8, 0x43, 0x14, 0x29, 0x17, 0x30, 0xad, 0xc1, 0x9a, 0x19, 0x43,
	0x1a, 0xb3, 0xa3, 0x9c, 0x2c, 0x03, 0x2d, 0xc6, 0x1a, 0x63, 0x13, 0xb7, 0x99, 0x15, 0x2f, 0x94,
	0x4b, 0xee, 0x39, 0x27, 0xf7, 0x3c, 0x54, 0xe4, 0xb6, 0x8d, 0xf9, 0xd8, 0x28, 0x52, 0x0e, 0xb9,
	0x51, 0xff, 0xaa, 0x5f, 0x75, 0x75, 0xb9, 0xab, 0x80, 0xc6, 0x9a, 0x71, 0xee, 0xae, 0xd8, 0xd5,
	0x26, 0x0a, 0xe3, 0x10, 0xcb, 0x99, 0xd9, 0xf9, 0xbd, 0x02, 0xf2, 0x7d, 0xfa, 0x1b, 0x7f, 0x0f,
	0x1a, 0x67, 0xd1, 0xb3, 0xb7, 0x60, 0x4e, 0xbc, 0xdb, 0x30, 0x5d, 0x6a, 0x4b, 0xdd, 0x66, 0xef,
	0xf3, 0xab, 0x1c, 0x9d, 0xa6, 0x4e, 0x7b, 0xb7, 0x61, 0x54, 0xe5, 0x85, 0x81, 0xbb, 0x50, 0x15,
	0x40, 0xf9, 0x04, 0xc8, 0x12, 0x0b, 0x40, 0x44, 0xe0, 0xd7, 0xa0, 0x70, 0x6f, 0x15, 0xb8, 0xf1,
	0x36, 0x62, 0x7a, 0xa5, 0x2d, 0x75, 0x35, 0x5a, 0x08, 0xf8, 0x3d, 0xc8, 0x3c, 0x76, 0x9f, 0xbc,
	0x60, 0xa5, 0x57, 0xdb, 0x52, 0x57, 0xed, 0xbd, 0x2a, 0xce, 0x4e, 0x75, 0xca, 0x7e, 0xd9, 0x32,
	0x1e, 0x8f, 0x4a, 0x34, 0x8f, 0xc4, 0x3f, 0x80, 0xb2, 0x08, 0x03, 0xce, 0x02, 0xbe, 0xe5, 0xfa,
	0x85, 0xc0, 0xbe, 0xdc, 0x63, 0x46, 0xee, 0x29, 0xc0, 0x22, 0x1a, 0x7f, 0x07, 0x17, 0xcb, 0xc8,
	0x0d, 0x96, 0x7a, 0x4d, 0x60, 0x5f, 0xec, 0xb1, 0x61, 0xa2, 0x16, 0x48, 0x1a, 0x85, 0x7f, 0x04,
	0x78, 0xf6, 0xd8, 0xc7, 0xc5, 0xa3, 0x1b, 0xac, 0x98, 0x2e, 0x0b, 0xe6, 0x72, 0xcf, 0x7c, 0xf0,
	0xd8, 0x47, 0x43, 0xb8, 0x0a, 0xf0, 0x20, 0x1e, 0x0f, 0xa0, 0xe5, 0x87, 0x71, 0xcc, 0xa2, 0x9d,
	0x13, 0xa5, 0x01, 0x7a, 0xfd, 0xe4, 0x92, 0x77, 0xa9, 0xbf, 0xe0, 0x9b, 0xfe, 0x91, 0x32, 0x50,
	0x40, 0xce, 0xd8, 0xce, 0x1f, 0x12, 0xd4, 0x29, 0xe3, 0x9b, 0xe4, 0x32, 0xff, 0xc7, 0x97, 0x23,
	0x80, 0x8a, 0xf2, 0xd3, 0x63, 0xc5, 0x07, 0x54, 0x7b, 0xfa, 0x79, 0xfd, 0xa9, 0x7f, 0x54, 0xa2,
	0x2d, 0xff, 0x58, 0x1a, 0x00, 0xd4, 0x73, 0xbc, 0x73, 0x03, 0xad, 0x13, 0x02, 0xeb, 0x20, 0x6f,
	0x7c, 0x77, 0xc7, 0x22, 0xae, 0x97, 0xdb, 0x95, 0xae, 0x42, 0x73, 0x13, 0x5f, 0x42, 0x7d, 0xee,
	0xfa, 0x6e, 0xb0, 0x60, 0x5c, 0xaf, 0x08, 0xd7, 0xde, 0xee, 0xfc, 0x26, 0x41, 0xf3, 0xb8, 0x77,
	0xf8, 0x5d, 0x76, 0xb1, 0xb4, 0x13, 0xaf, 0xff, 0xa1, 0xc5, 0x57, 0x07, 0x17, 0xfc, 0x06, 0xd4,
	0x4d, 0xe4, 0x3d, 0xbb, 0x31, 0x73, 0x9e, 0xd8, 0x4e, 0x74, 0x44, 0xa1, 0x90, 0x49, 0xb7, 0x6c,
	0x87, 0x5f, 0x42, 0xcd, 0x5d, 0x87, 0xdb, 0x20, 0x16, 0xf7, 0xae, 0xd0, 0xcc, 0xea, 0x5c, 0x41,
	0x55, 0xf4, 0x52, 0x81, 0x0b, 0x62, 0xd9, 0x84, 0xa2, 0x12, 0x06, 0xa8, 0x51, 0x32, 0x9d, 0xdd,
	0xd9, 0x48, 0xc2, 0x2d, 0x50, 0x27, 0xa6, 0x71, 0xeb, 0x3c, 0x98, 0x96, 0x45, 0x28, 0x2a, 0x77,
	0x6e, 0xa1, 0x79, 0xfc, 0x9a, 0x71, 0x1b, 0xd4, 0x38, 0x72, 0x03, 0xee, 0x2e, 0x62, 0x2f, 0x0c,
	0x44, 0xcd, 0x1a, 0x3d, 0x94, 0xf0, 0x2b, 0x90, 0x83, 0x70, 0xc9, 0x1c, 0x6f, 0x99, 0x15, 0x56,
	0x4b, 0x4c, 0x73, 0xd9, 0xf9, 0x55, 0x02, 0x74, 0xfa, 0xc8, 0x93, 0xe8, 0xe4, 0xe1, 0x25, 0xd1,
	0x49, 0xae, 0x06, 0xad, 0x25, 0xa6, 0xb9, 0xc4, 0x5f, 0x81, 0x32, 0xf7, 0xc3, 0xc5, 0x93, 0x13,
	0x6c, 0xd7, 0x22, 0x51, 0x95, 0xd6, 0x85, 0x60, 0x6d, 0xd7, 0xf8, 0x6b, 0x80, 0xd4, 0xf9, 0xe8,
	0xf2, 0xc7, 0x7c, 0x38, 0x85, 0x32, 0x72, 0xf9, 0x23, 0x7e, 0x03, 0x0d, 0xce, 0x82, 0x25, 0x8b,
	0x9c, 0xcd, 0x76, 0x9e, 0x74, 0xa8, 0x2a, 0x22, 0xb4, 0x54, 0x9c, 0x08, 0x4d, 0x7c, 0x3f, 0x77,
	0xe7, 0x87, 0xee, 0x52, 0x8c, 0xa2, 0x46, 0x73, 0xb3, 0xe3, 0x83, 0x76, 0x38, 0x55, 0xe7, 0xe9,
	0xa4, 0x4f, 0xa4, 0x3b, 0x2e, 0xa9, 0x7c, 0x5a, 0xd2, 0xc1, 0x69, 0x95, 0xe3, 0xd3, 0xfe, 0x2a,
	0xc3, 0x8b, 0xb3, 0x81, 0xfc, 0x8f, 0x7d, 0x39, 0xab, 0xb4, 0xf2, 0x89, 0x4a, 0xdf, 0x40, 0xc3,
	0x67, 0xee, 0x79, 0x77, 0x52, 0xf1, 0xdf, 0xba, 0x83, 0xbf, 0x85, 0x66, 0xb1, 0x2a, 0x1c, 0xee,
	0xad, 0xc4, 0x4a, 0xd2, 0x68, 0xa3, 0x50, 0xa7, 0xde, 0x2a, 0xe9, 0xc7, 0xfa, 0xda, 0x71, 0x57,
	0x2b, 0xee, 0xad, 0xb8, 0xd8, 0x40, 0x1a, 0x55, 0xd6, 0xd7, 0xfd, 0x54, 0x48, 0xae, 0xb1, 0xbe,
	0x76, 0xe6, 0x5e, 0xbc, 0x76, 0x37, 0x62, 0xb9, 0x68, 0xb4, 0xbe, 0xbe, 0x1e, 0x08, 0x5b, 0xb0,
	0xbd, 0x3d, 0xab, 0x64, 0x6c, 0xef, 0x90, 0xed, 0xe5, 0x2c, 0x64, 0x6c, 0x2f, 0x65, 0xdf, 0x8e,
	0x40, 0x3d, 0x58, 0x21, 0xb8, 0x01, 0x8a, 0x31, 0xb6, 0xa6, 0xc4, 0x9a, 0xce, 0xa6, 0xa8, 0x84,
	0x55, 0x90, 0xa7, 0x76, 0xff, 0xd6, 0xb4, 0x6e, 0x90, 0x94, 0x4c, 0xc1, 0x90, 0xf6, 0xad, 0x21,
	0x2a, 0x63, 0x0c, 0x4d, 0xe3, 0xce, 0x24, 0x96, 0xed, 0x4c, 0x67, 0x93, 0xc9, 0x98, 0xda, 0xa8,
	0xf2, 0xf6, 0x4f, 0x09, 0xd4, 0x83, 0xe5, 0x82, 0x2f, 0xe1, 0xa5, 0x45, 0x1e, 0xac, 0xf1, 0x90,
	0x38, 0x03, 0xd2, 0x37, 0xc6, 0x96, 0x93, 0xa7, 0x2a, 0x61, 0x0d, 0xea, 0x7d, 0xcb, 0x1a, 0xcf,
	0x2c, 0x83, 0x20, 0x29, 0x39, 0x65, 0x42, 0xc9, 0xa4, 0x4f, 0x09, 0x2a, 0x27, 0xae, 0xcc, 0x18,
	0xa2, 0x4a, 0x32, 0x6e, 0xc6, 0xf8, 0xfe, 0xde, 0xb4, 0x51, 0x35, 0xad, 0x2d, 0xf9, 0x6d, 0x93,
	0x21, 0xba, 0xc0, 0x4d, 0x80, 0x0f, 0x26, 0x79, 0x30, 0x46, 0x7d, 0xeb, 0x86, 0xa0, 0x5a, 0x92,
	0xc5, 0x22, 0x0f, 0x89, 0x84, 0xe4, 0xc4, 0x29, 0x6a, 0x75, 0x4c, 0xcb, 0xb4, 0x11, 0x60, 0x04,
	0x5a, 0x6a, 0x67, 0xd9, 0x54, 0xfc, 0x19, 0xb4, 0xee, 0xc6, 0xb6, 0x4d, 0xe8, 0xcf, 0x0e, 0x25,
	0x3f, 0xcd, 0xc8, 0xd4, 0x46, 0x5a, 0xaf, 0x0f, 0x0d, 0xc3, 0xf7, 0x58, 0x10, 0x67, 0x3d, 0xc1,
	0xef, 0x40, 0x9e, 0x44, 0xe1, 0x82, 0x71, 0x8e, 0xd1, 0xe9, 0x0a, 0xbd, 0x7c, 0xb1, 0x57, 0xf2,
	0x2d, 0xd7, 0x29, 0xcd, 0x6b, 0xe2, 0x6f, 0xf8, 0xfd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0x7b, 0x2c, 0xea, 0x97, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	Process(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Process(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.ClientService/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	Process(context.Context, *Message) (*Response, error)
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.ClientService/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Process(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _ClientService_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
