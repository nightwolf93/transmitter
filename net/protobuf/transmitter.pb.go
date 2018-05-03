// Code generated by protoc-gen-go. DO NOT EDIT.
// source: net/protobuf/transmitter.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The client driver
type ClientDriver struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProtocolVersion      int32    `protobuf:"varint,2,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientDriver) Reset()         { *m = ClientDriver{} }
func (m *ClientDriver) String() string { return proto.CompactTextString(m) }
func (*ClientDriver) ProtoMessage()    {}
func (*ClientDriver) Descriptor() ([]byte, []int) {
	return fileDescriptor_transmitter_1f0abee039cd5034, []int{0}
}
func (m *ClientDriver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientDriver.Unmarshal(m, b)
}
func (m *ClientDriver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientDriver.Marshal(b, m, deterministic)
}
func (dst *ClientDriver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientDriver.Merge(dst, src)
}
func (m *ClientDriver) XXX_Size() int {
	return xxx_messageInfo_ClientDriver.Size(m)
}
func (m *ClientDriver) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientDriver.DiscardUnknown(m)
}

var xxx_messageInfo_ClientDriver proto.InternalMessageInfo

func (m *ClientDriver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClientDriver) GetProtocolVersion() int32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

// Server settings and informations
type ServerInformations struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProtocolVersion      int32    `protobuf:"varint,3,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInformations) Reset()         { *m = ServerInformations{} }
func (m *ServerInformations) String() string { return proto.CompactTextString(m) }
func (*ServerInformations) ProtoMessage()    {}
func (*ServerInformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_transmitter_1f0abee039cd5034, []int{1}
}
func (m *ServerInformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInformations.Unmarshal(m, b)
}
func (m *ServerInformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInformations.Marshal(b, m, deterministic)
}
func (dst *ServerInformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInformations.Merge(dst, src)
}
func (m *ServerInformations) XXX_Size() int {
	return xxx_messageInfo_ServerInformations.Size(m)
}
func (m *ServerInformations) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInformations.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInformations proto.InternalMessageInfo

func (m *ServerInformations) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerInformations) GetProtocolVersion() int32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

// The server give to the client all the informations about him on the server
type HandshakeRequest struct {
	Uid                  string              `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	ServerInformations   *ServerInformations `protobuf:"bytes,2,opt,name=serverInformations" json:"serverInformations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HandshakeRequest) Reset()         { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()    {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transmitter_1f0abee039cd5034, []int{2}
}
func (m *HandshakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeRequest.Unmarshal(m, b)
}
func (m *HandshakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeRequest.Marshal(b, m, deterministic)
}
func (dst *HandshakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeRequest.Merge(dst, src)
}
func (m *HandshakeRequest) XXX_Size() int {
	return xxx_messageInfo_HandshakeRequest.Size(m)
}
func (m *HandshakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeRequest proto.InternalMessageInfo

func (m *HandshakeRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *HandshakeRequest) GetServerInformations() *ServerInformations {
	if m != nil {
		return m.ServerInformations
	}
	return nil
}

// The client respond with the informations about him, like the driver he use
type HandshakeResponse struct {
	ClientDriver         *ClientDriver `protobuf:"bytes,1,opt,name=clientDriver" json:"clientDriver,omitempty"`
	Token                string        `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HandshakeResponse) Reset()         { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()    {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transmitter_1f0abee039cd5034, []int{3}
}
func (m *HandshakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandshakeResponse.Unmarshal(m, b)
}
func (m *HandshakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandshakeResponse.Marshal(b, m, deterministic)
}
func (dst *HandshakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeResponse.Merge(dst, src)
}
func (m *HandshakeResponse) XXX_Size() int {
	return xxx_messageInfo_HandshakeResponse.Size(m)
}
func (m *HandshakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeResponse proto.InternalMessageInfo

func (m *HandshakeResponse) GetClientDriver() *ClientDriver {
	if m != nil {
		return m.ClientDriver
	}
	return nil
}

func (m *HandshakeResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientDriver)(nil), "transmitter.ClientDriver")
	proto.RegisterType((*ServerInformations)(nil), "transmitter.ServerInformations")
	proto.RegisterType((*HandshakeRequest)(nil), "transmitter.HandshakeRequest")
	proto.RegisterType((*HandshakeResponse)(nil), "transmitter.HandshakeResponse")
}

func init() {
	proto.RegisterFile("net/protobuf/transmitter.proto", fileDescriptor_transmitter_1f0abee039cd5034)
}

var fileDescriptor_transmitter_1f0abee039cd5034 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0x42, 0x41,
	0x10, 0xc6, 0x79, 0x99, 0x91, 0xf3, 0x84, 0x6c, 0xe8, 0x60, 0x97, 0x92, 0x77, 0x7a, 0x27, 0x05,
	0x3b, 0x77, 0xa9, 0x0e, 0x05, 0x41, 0xb0, 0x41, 0x87, 0x6e, 0xab, 0x8e, 0xb8, 0xe8, 0x9b, 0xb5,
	0xd9, 0x59, 0xff, 0xfe, 0x68, 0x2d, 0x58, 0xd3, 0x83, 0xb7, 0x6f, 0x3f, 0x76, 0xbe, 0xf9, 0xf8,
	0x0d, 0xdc, 0x30, 0xe9, 0x68, 0x2d, 0x5e, 0xfd, 0x24, 0xce, 0x47, 0x2a, 0x96, 0x43, 0xe3, 0x54,
	0x49, 0x86, 0xc9, 0xc4, 0x32, 0xb3, 0xaa, 0x57, 0xe8, 0x3e, 0xae, 0x1c, 0xb1, 0x3e, 0x89, 0xdb,
	0x90, 0x20, 0xc2, 0x29, 0xdb, 0x86, 0xfa, 0xc5, 0xa0, 0xa8, 0x3b, 0x26, 0x69, 0xac, 0xe1, 0x22,
	0x4d, 0x4e, 0xfd, 0xea, 0x83, 0x24, 0x38, 0xcf, 0xfd, 0x93, 0x41, 0x51, 0xb7, 0xcd, 0x7f, 0xbb,
	0x32, 0x80, 0xef, 0x24, 0x1b, 0x92, 0x17, 0x9e, 0x7b, 0x69, 0xac, 0x3a, 0xcf, 0xe1, 0xd8, 0xcc,
	0xd6, 0xe1, 0xcc, 0x08, 0xbd, 0x67, 0xcb, 0xb3, 0xb0, 0xb0, 0x4b, 0x32, 0xf4, 0x15, 0x29, 0x28,
	0xf6, 0xa0, 0x15, 0xdd, 0xec, 0x37, 0xf0, 0x47, 0xe2, 0x1b, 0x60, 0xd8, 0xdb, 0x9c, 0x6a, 0x96,
	0xe3, 0xdb, 0x61, 0x0e, 0x61, 0xbf, 0xa0, 0x39, 0x30, 0x5a, 0x2d, 0xe0, 0x32, 0x5b, 0x1b, 0xd6,
	0x9e, 0x03, 0xe1, 0x3d, 0x74, 0xa7, 0x19, 0xad, 0x54, 0xa0, 0x1c, 0x5f, 0xef, 0xe4, 0xe7, 0x38,
	0xcd, 0xce, 0x77, 0xbc, 0x82, 0xb6, 0xfa, 0x25, 0x6d, 0xf1, 0x75, 0xcc, 0xf6, 0xf1, 0x00, 0x9f,
	0xe7, 0x7f, 0xd7, 0x9a, 0x9c, 0x25, 0x75, 0xf7, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x4b, 0xf1,
	0x35, 0xc4, 0x01, 0x00, 0x00,
}