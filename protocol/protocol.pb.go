// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/protocol.proto

package protocol

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MESSAGE_REQUEST   MessageType = 0
	MESSAGE_RESPONSE  MessageType = 1
	MESSAGE_HEARTBEAT MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "MESSAGE_REQUEST",
	1: "MESSAGE_RESPONSE",
	2: "MESSAGE_HEARTBEAT",
}

var MessageType_value = map[string]int32{
	"MESSAGE_REQUEST":   0,
	"MESSAGE_RESPONSE":  1,
	"MESSAGE_HEARTBEAT": 2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{0}
}

type ErrorCode int32

const (
	ERROR_NO                ErrorCode = 0
	ERROR_TOO_MANY_REQUESTS ErrorCode = 1
	ERROR_NOT_FOUND         ErrorCode = 2
	ERROR_BAD_REQUEST       ErrorCode = 3
	ERROR_NOT_IMPLEMENTED   ErrorCode = 4
	ERROR_INTERNAL_SERVER   ErrorCode = 5
	ERROR_USER_DEFINED      ErrorCode = 256
)

var ErrorCode_name = map[int32]string{
	0:   "ERROR_NO",
	1:   "ERROR_TOO_MANY_REQUESTS",
	2:   "ERROR_NOT_FOUND",
	3:   "ERROR_BAD_REQUEST",
	4:   "ERROR_NOT_IMPLEMENTED",
	5:   "ERROR_INTERNAL_SERVER",
	256: "ERROR_USER_DEFINED",
}

var ErrorCode_value = map[string]int32{
	"ERROR_NO":                0,
	"ERROR_TOO_MANY_REQUESTS": 1,
	"ERROR_NOT_FOUND":         2,
	"ERROR_BAD_REQUEST":       3,
	"ERROR_NOT_IMPLEMENTED":   4,
	"ERROR_INTERNAL_SERVER":   5,
	"ERROR_USER_DEFINED":      256,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{1}
}

type Greeting struct {
	Channel   Greeting_Channel `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel"`
	Handshake []byte           `protobuf:"bytes,2,opt,name=handshake,proto3" json:"handshake,omitempty"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{0}
}
func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return m.Size()
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetChannel() Greeting_Channel {
	if m != nil {
		return m.Channel
	}
	return Greeting_Channel{}
}

func (m *Greeting) GetHandshake() []byte {
	if m != nil {
		return m.Handshake
	}
	return nil
}

type Greeting_Channel struct {
	Id                 []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timeout            int32  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	IncomingWindowSize int32  `protobuf:"varint,3,opt,name=incoming_window_size,json=incomingWindowSize,proto3" json:"incoming_window_size,omitempty"`
	OutgoingWindowSize int32  `protobuf:"varint,4,opt,name=outgoing_window_size,json=outgoingWindowSize,proto3" json:"outgoing_window_size,omitempty"`
}

func (m *Greeting_Channel) Reset()         { *m = Greeting_Channel{} }
func (m *Greeting_Channel) String() string { return proto.CompactTextString(m) }
func (*Greeting_Channel) ProtoMessage()    {}
func (*Greeting_Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{0, 0}
}
func (m *Greeting_Channel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Greeting_Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Greeting_Channel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Greeting_Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting_Channel.Merge(m, src)
}
func (m *Greeting_Channel) XXX_Size() int {
	return m.Size()
}
func (m *Greeting_Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting_Channel proto.InternalMessageInfo

func (m *Greeting_Channel) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Greeting_Channel) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Greeting_Channel) GetIncomingWindowSize() int32 {
	if m != nil {
		return m.IncomingWindowSize
	}
	return 0
}

func (m *Greeting_Channel) GetOutgoingWindowSize() int32 {
	if m != nil {
		return m.OutgoingWindowSize
	}
	return 0
}

type RequestHeader struct {
	SequenceNumber int32  `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	ServiceName    string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	MethodName     string `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	FifoKey        string `protobuf:"bytes,4,opt,name=fifo_key,json=fifoKey,proto3" json:"fifo_key,omitempty"`
	ExtraData      []byte `protobuf:"bytes,5,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	TraceId        []byte `protobuf:"bytes,6,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId         int32  `protobuf:"varint,7,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{1}
}
func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return m.Size()
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *RequestHeader) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RequestHeader) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *RequestHeader) GetFifoKey() string {
	if m != nil {
		return m.FifoKey
	}
	return ""
}

func (m *RequestHeader) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *RequestHeader) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *RequestHeader) GetSpanId() int32 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

type ResponseHeader struct {
	SequenceNumber int32 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	NextSpanId     int32 `protobuf:"varint,2,opt,name=next_span_id,json=nextSpanId,proto3" json:"next_span_id,omitempty"`
	ErrorCode      int32 `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{2}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *ResponseHeader) GetNextSpanId() int32 {
	if m != nil {
		return m.NextSpanId
	}
	return 0
}

func (m *ResponseHeader) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type Heartbeat struct {
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_87968d26f3046c60, []int{3}
}
func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return m.Size()
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*Greeting)(nil), "Greeting")
	proto.RegisterType((*Greeting_Channel)(nil), "Greeting.Channel")
	proto.RegisterType((*RequestHeader)(nil), "RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "ResponseHeader")
	proto.RegisterType((*Heartbeat)(nil), "Heartbeat")
}

func init() { proto.RegisterFile("protocol/protocol.proto", fileDescriptor_87968d26f3046c60) }

var fileDescriptor_87968d26f3046c60 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbf, 0x52, 0xdb, 0x4c,
	0x10, 0xb7, 0x0c, 0xc6, 0x78, 0xed, 0x0f, 0xc4, 0x7d, 0x30, 0x36, 0x7c, 0x5f, 0x0c, 0x71, 0x13,
	0x86, 0x19, 0x4c, 0xfe, 0xbc, 0x40, 0x6c, 0x7c, 0x80, 0x27, 0x58, 0x26, 0x27, 0x91, 0x4c, 0xd2,
	0xdc, 0x9c, 0xa5, 0xc5, 0xd6, 0x80, 0x75, 0x8e, 0x74, 0x0e, 0xe0, 0x2a, 0x8f, 0x90, 0x2e, 0xcf,
	0x91, 0xb7, 0xa0, 0xa4, 0x4c, 0x95, 0x49, 0xa0, 0x4b, 0x9f, 0x3e, 0xa3, 0x13, 0xc2, 0x99, 0x74,
	0xe9, 0xf6, 0xf7, 0x6f, 0x77, 0xa5, 0x9b, 0x85, 0xf2, 0x28, 0x94, 0x4a, 0xba, 0xf2, 0x6c, 0x27,
	0x2d, 0xea, 0xba, 0x58, 0xdb, 0xee, 0xfb, 0x6a, 0x30, 0xee, 0xd5, 0x5d, 0x39, 0xdc, 0xe9, 0xcb,
	0xbe, 0x4c, 0xf4, 0xde, 0xf8, 0x44, 0x23, 0x0d, 0x74, 0x95, 0xd8, 0x6b, 0x3f, 0x0d, 0x98, 0xdf,
	0x0f, 0x11, 0x95, 0x1f, 0xf4, 0xc9, 0x13, 0xc8, 0xbb, 0x03, 0x11, 0x04, 0x78, 0x56, 0x31, 0x36,
	0x8c, 0xcd, 0xe2, 0xd3, 0xa5, 0x7a, 0xaa, 0xd5, 0x77, 0x13, 0xa1, 0x39, 0x7b, 0xf5, 0x75, 0x3d,
	0xc3, 0x52, 0x1f, 0xf9, 0x1f, 0x0a, 0x03, 0x11, 0x78, 0xd1, 0x40, 0x9c, 0x62, 0x25, 0xbb, 0x61,
	0x6c, 0x96, 0xd8, 0x94, 0x58, 0xfb, 0x64, 0x40, 0xfe, 0x2e, 0x48, 0x16, 0x20, 0xeb, 0x7b, 0xba,
	0x6f, 0x89, 0x65, 0x7d, 0x8f, 0x54, 0x20, 0xaf, 0xfc, 0x21, 0xca, 0xb1, 0xd2, 0xb9, 0x1c, 0x4b,
	0x21, 0x79, 0x0c, 0xcb, 0x7e, 0xe0, 0xca, 0xa1, 0x1f, 0xf4, 0xf9, 0xb9, 0x1f, 0x78, 0xf2, 0x9c,
	0x47, 0xfe, 0x04, 0x2b, 0x33, 0xda, 0x46, 0x52, 0xed, 0xb5, 0x96, 0x6c, 0x7f, 0x82, 0x71, 0x42,
	0x8e, 0x55, 0x5f, 0xfe, 0x99, 0x98, 0x4d, 0x12, 0xa9, 0x36, 0x4d, 0xd4, 0x7e, 0x18, 0xf0, 0x0f,
	0xc3, 0x77, 0x63, 0x8c, 0xd4, 0x01, 0x0a, 0x0f, 0x43, 0xf2, 0x08, 0x16, 0xa3, 0x98, 0x08, 0x5c,
	0xe4, 0xc1, 0x78, 0xd8, 0xc3, 0x50, 0x2f, 0x9b, 0x63, 0x0b, 0x29, 0x6d, 0x69, 0x96, 0x3c, 0x84,
	0x52, 0x84, 0xe1, 0x7b, 0x3f, 0xf6, 0x89, 0x61, 0xf2, 0xd5, 0x05, 0x56, 0xbc, 0xe3, 0x2c, 0x31,
	0x44, 0xb2, 0x0e, 0xc5, 0x21, 0xaa, 0x81, 0xf4, 0x12, 0xc7, 0x8c, 0x76, 0x40, 0x42, 0x69, 0xc3,
	0x2a, 0xcc, 0x9f, 0xf8, 0x27, 0x92, 0x9f, 0xe2, 0xa5, 0x5e, 0xb2, 0xc0, 0xf2, 0x31, 0x7e, 0x81,
	0x97, 0xe4, 0x01, 0x00, 0x5e, 0xa8, 0x50, 0x70, 0x4f, 0x28, 0x51, 0xc9, 0x25, 0xbf, 0x54, 0x33,
	0x2d, 0xa1, 0x44, 0x9c, 0x54, 0xa1, 0x70, 0x91, 0xfb, 0x5e, 0x65, 0x4e, 0x8b, 0x79, 0x8d, 0xdb,
	0x1e, 0x29, 0x43, 0x3e, 0x1a, 0x89, 0x20, 0x56, 0xf2, 0x7a, 0xf3, 0xb9, 0x18, 0xb6, 0xbd, 0xda,
	0x04, 0x16, 0x18, 0x46, 0x23, 0x19, 0x44, 0xf8, 0xb7, 0x1f, 0xbb, 0x01, 0xa5, 0x00, 0x2f, 0x14,
	0x4f, 0x1b, 0x27, 0x4f, 0x05, 0x31, 0x67, 0xeb, 0xe6, 0x7a, 0xdf, 0x30, 0x94, 0x21, 0x77, 0xa5,
	0x97, 0xbe, 0x51, 0x41, 0x33, 0xbb, 0xd2, 0xc3, 0x5a, 0x11, 0x0a, 0x07, 0x28, 0x42, 0xd5, 0x43,
	0xa1, 0xb6, 0xba, 0x50, 0xec, 0x60, 0x14, 0x89, 0x3e, 0x3a, 0x97, 0x23, 0x24, 0xff, 0xc2, 0x62,
	0x87, 0xda, 0x76, 0x63, 0x9f, 0x72, 0x46, 0x5f, 0x1e, 0x53, 0xdb, 0x31, 0x33, 0x64, 0x19, 0xcc,
	0x29, 0x69, 0x1f, 0x75, 0x2d, 0x9b, 0x9a, 0x06, 0x59, 0x81, 0xa5, 0x94, 0x3d, 0xa0, 0x0d, 0xe6,
	0x34, 0x69, 0xc3, 0x31, 0xb3, 0x5b, 0x9f, 0x0d, 0x28, 0xd0, 0x74, 0x16, 0x29, 0xc1, 0x3c, 0x65,
	0xac, 0xcb, 0xb8, 0xd5, 0x35, 0x33, 0xe4, 0x3f, 0x28, 0x27, 0xc8, 0xe9, 0x76, 0x79, 0xa7, 0x61,
	0xbd, 0x49, 0x87, 0xd8, 0xa6, 0x11, 0x8f, 0x4e, 0xad, 0x0e, 0xdf, 0xeb, 0x1e, 0x5b, 0x2d, 0x33,
	0x1b, 0x0f, 0x49, 0xc8, 0x66, 0xa3, 0x75, 0xbf, 0xd1, 0x0c, 0x59, 0x85, 0x95, 0xa9, 0xb7, 0xdd,
	0x39, 0x3a, 0xa4, 0x1d, 0x6a, 0x39, 0xb4, 0x65, 0xce, 0x4e, 0xa5, 0xb6, 0xe5, 0x50, 0x66, 0x35,
	0x0e, 0xb9, 0x4d, 0xd9, 0x2b, 0xca, 0xcc, 0x1c, 0x29, 0x03, 0x49, 0xa4, 0x63, 0x9b, 0x32, 0xde,
	0xa2, 0x7b, 0x6d, 0x8b, 0xb6, 0xcc, 0x0f, 0xd9, 0xe6, 0xf3, 0xeb, 0xef, 0xd5, 0xcc, 0xd5, 0x4d,
	0xd5, 0xb8, 0xbe, 0xa9, 0x1a, 0xdf, 0x6e, 0xaa, 0xc6, 0xc7, 0xdb, 0x6a, 0xe6, 0xfa, 0xb6, 0x9a,
	0xf9, 0x72, 0x5b, 0xcd, 0xbc, 0xad, 0xfd, 0x76, 0xbf, 0x67, 0xa8, 0xb6, 0x27, 0xdb, 0xf1, 0x0d,
	0xf7, 0xc2, 0x91, 0x7b, 0x7f, 0xe9, 0xbd, 0x39, 0x5d, 0x3d, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0xc2, 0x93, 0x41, 0x05, 0x04, 0x00, 0x00,
}

func (m *Greeting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Greeting) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintProtocol(dAtA, i, uint64(m.Channel.Size()))
	n1, err := m.Channel.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Handshake) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Handshake)))
		i += copy(dAtA[i:], m.Handshake)
	}
	return i, nil
}

func (m *Greeting_Channel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Greeting_Channel) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.Timeout))
	}
	if m.IncomingWindowSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.IncomingWindowSize))
	}
	if m.OutgoingWindowSize != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.OutgoingWindowSize))
	}
	return i, nil
}

func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SequenceNumber))
	}
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.MethodName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.MethodName)))
		i += copy(dAtA[i:], m.MethodName)
	}
	if len(m.FifoKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.FifoKey)))
		i += copy(dAtA[i:], m.FifoKey)
	}
	if len(m.ExtraData) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.ExtraData)))
		i += copy(dAtA[i:], m.ExtraData)
	}
	if len(m.TraceId) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.TraceId)))
		i += copy(dAtA[i:], m.TraceId)
	}
	if m.SpanId != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SpanId))
	}
	return i, nil
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SequenceNumber))
	}
	if m.NextSpanId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.NextSpanId))
	}
	if m.ErrorCode != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.ErrorCode))
	}
	return i, nil
}

func (m *Heartbeat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Heartbeat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintProtocol(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Greeting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Channel.Size()
	n += 1 + l + sovProtocol(uint64(l))
	l = len(m.Handshake)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	return n
}

func (m *Greeting_Channel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovProtocol(uint64(m.Timeout))
	}
	if m.IncomingWindowSize != 0 {
		n += 1 + sovProtocol(uint64(m.IncomingWindowSize))
	}
	if m.OutgoingWindowSize != 0 {
		n += 1 + sovProtocol(uint64(m.OutgoingWindowSize))
	}
	return n
}

func (m *RequestHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		n += 1 + sovProtocol(uint64(m.SequenceNumber))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.MethodName)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.FifoKey)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.ExtraData)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.TraceId)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SpanId != 0 {
		n += 1 + sovProtocol(uint64(m.SpanId))
	}
	return n
}

func (m *ResponseHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		n += 1 + sovProtocol(uint64(m.SequenceNumber))
	}
	if m.NextSpanId != 0 {
		n += 1 + sovProtocol(uint64(m.NextSpanId))
	}
	if m.ErrorCode != 0 {
		n += 1 + sovProtocol(uint64(m.ErrorCode))
	}
	return n
}

func (m *Heartbeat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovProtocol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Greeting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: Greeting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Greeting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Channel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handshake", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handshake = append(m.Handshake[:0], dAtA[iNdEx:postIndex]...)
			if m.Handshake == nil {
				m.Handshake = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
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
func (m *Greeting_Channel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: Channel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Channel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
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
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncomingWindowSize", wireType)
			}
			m.IncomingWindowSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IncomingWindowSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutgoingWindowSize", wireType)
			}
			m.OutgoingWindowSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutgoingWindowSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
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
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraData = append(m.ExtraData[:0], dAtA[iNdEx:postIndex]...)
			if m.ExtraData == nil {
				m.ExtraData = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
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
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProtocol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = append(m.TraceId[:0], dAtA[iNdEx:postIndex]...)
			if m.TraceId == nil {
				m.TraceId = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpanId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
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
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextSpanId", wireType)
			}
			m.NextSpanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextSpanId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
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
func (m *Heartbeat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
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
			return fmt.Errorf("proto: Heartbeat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Heartbeat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProtocol
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
func skipProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
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
					return 0, ErrIntOverflowProtocol
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
					return 0, ErrIntOverflowProtocol
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
				return 0, ErrInvalidLengthProtocol
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
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
				next, err := skipProtocol(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthProtocol
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
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)
