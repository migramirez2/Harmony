// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus.proto

package consensus

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MessageType int32

const (
	MessageType_UNKNOWN         MessageType = 0
	MessageType_ANNOUNCE        MessageType = 1
	MessageType_COMMIT          MessageType = 2
	MessageType_CHALLENGE       MessageType = 3
	MessageType_RESPONSE        MessageType = 4
	MessageType_COLLECTIVE_SIG  MessageType = 5
	MessageType_FINAL_COMMIT    MessageType = 6
	MessageType_FINAL_CHALLENGE MessageType = 7
	MessageType_FINAL_RESPONSE  MessageType = 8
)

var MessageType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ANNOUNCE",
	2: "COMMIT",
	3: "CHALLENGE",
	4: "RESPONSE",
	5: "COLLECTIVE_SIG",
	6: "FINAL_COMMIT",
	7: "FINAL_CHALLENGE",
	8: "FINAL_RESPONSE",
}

var MessageType_value = map[string]int32{
	"UNKNOWN":         0,
	"ANNOUNCE":        1,
	"COMMIT":          2,
	"CHALLENGE":       3,
	"RESPONSE":        4,
	"COLLECTIVE_SIG":  5,
	"FINAL_COMMIT":    6,
	"FINAL_CHALLENGE": 7,
	"FINAL_RESPONSE":  8,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
}

type Message struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=consensus.MessageType" json:"type,omitempty"`
	ConsensusId          uint32      `protobuf:"varint,2,opt,name=consensus_id,json=consensusId,proto3" json:"consensus_id,omitempty"`
	SenderId             uint32      `protobuf:"varint,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	BlockHash            []byte      `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	Payload              []byte      `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            []byte      `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
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

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN
}

func (m *Message) GetConsensusId() uint32 {
	if m != nil {
		return m.ConsensusId
	}
	return 0
}

func (m *Message) GetSenderId() uint32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Message) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("consensus.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Message)(nil), "consensus.Message")
}

func init() { proto.RegisterFile("consensus.proto", fileDescriptor_56f0f2c53b3de771) }

var fileDescriptor_56f0f2c53b3de771 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdf, 0x4e, 0xc2, 0x30,
	0x14, 0xc6, 0x2d, 0x7f, 0x36, 0x76, 0x18, 0xd0, 0x1c, 0x13, 0xd3, 0x44, 0x4d, 0xd0, 0x2b, 0xc2,
	0x05, 0x17, 0xfa, 0x04, 0x64, 0xa9, 0xd0, 0x38, 0x3a, 0x33, 0x40, 0x2f, 0x97, 0xc2, 0x1a, 0x20,
	0x92, 0x6d, 0xa1, 0x70, 0xc1, 0xdb, 0xf8, 0x40, 0x3e, 0x94, 0x61, 0x60, 0xf5, 0xb2, 0xbf, 0xdf,
	0xf7, 0x7d, 0x69, 0x72, 0xa0, 0xb3, 0xcc, 0x33, 0xa3, 0x33, 0x73, 0x30, 0x83, 0x62, 0x97, 0xef,
	0x73, 0xf4, 0x2c, 0x78, 0xfc, 0x26, 0xe0, 0x4e, 0xb4, 0x31, 0x6a, 0xa5, 0xb1, 0x0f, 0xb5, 0xfd,
	0xb1, 0xd0, 0x8c, 0x74, 0x49, 0xaf, 0xfd, 0x74, 0x33, 0xf8, 0xab, 0x5d, 0x12, 0xb3, 0x63, 0xa1,
	0xe3, 0x32, 0x83, 0x0f, 0xe0, 0x5b, 0x9d, 0x6c, 0x52, 0x56, 0xe9, 0x92, 0x5e, 0x2b, 0x6e, 0x5a,
	0x26, 0x52, 0xbc, 0x05, 0xcf, 0xe8, 0x2c, 0xd5, 0xbb, 0x93, 0xaf, 0x96, 0xbe, 0x71, 0x06, 0x22,
	0xc5, 0x7b, 0x80, 0xc5, 0x36, 0x5f, 0x7e, 0x26, 0x6b, 0x65, 0xd6, 0xac, 0xd6, 0x25, 0x3d, 0x3f,
	0xf6, 0x4a, 0x32, 0x56, 0x66, 0x8d, 0x0c, 0xdc, 0x42, 0x1d, 0xb7, 0xb9, 0x4a, 0x59, 0xbd, 0x74,
	0xbf, 0x4f, 0xbc, 0x03, 0xcf, 0x6c, 0x56, 0x99, 0xda, 0x1f, 0x76, 0x9a, 0x39, 0xe7, 0x9e, 0x05,
	0xfd, 0x2f, 0x02, 0xcd, 0x7f, 0x9f, 0xc5, 0x26, 0xb8, 0x73, 0xf9, 0x2a, 0xa3, 0x0f, 0x49, 0xaf,
	0xd0, 0x87, 0xc6, 0x50, 0xca, 0x68, 0x2e, 0x03, 0x4e, 0x09, 0x02, 0x38, 0x41, 0x34, 0x99, 0x88,
	0x19, 0xad, 0x60, 0x0b, 0xbc, 0x60, 0x3c, 0x0c, 0x43, 0x2e, 0x47, 0x9c, 0x56, 0x4f, 0xc1, 0x98,
	0x4f, 0xdf, 0x22, 0x39, 0xe5, 0xb4, 0x86, 0x08, 0xed, 0x20, 0x0a, 0x43, 0x1e, 0xcc, 0xc4, 0x3b,
	0x4f, 0xa6, 0x62, 0x44, 0xeb, 0x48, 0xc1, 0x7f, 0x11, 0x72, 0x18, 0x26, 0x97, 0x09, 0x07, 0xaf,
	0xa1, 0x73, 0x21, 0x76, 0xc8, 0x3d, 0x55, 0xcf, 0xd0, 0xce, 0x35, 0x16, 0x4e, 0x79, 0x83, 0xe7,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xa2, 0x8a, 0x4f, 0x96, 0x01, 0x00, 0x00,
}
