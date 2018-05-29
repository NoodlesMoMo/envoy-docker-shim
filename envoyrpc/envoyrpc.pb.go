// Code generated by protoc-gen-go.
// source: envoyrpc.proto
// DO NOT EDIT!

/*
Package envoyrpc is a generated protocol buffer package.

It is generated from these files:
	envoyrpc.proto

It has these top-level messages:
	RegistrarRequest
	RegistrarReply
*/
package envoyrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegistrarRequest_Action int32

const (
	RegistrarRequest_REGISTER   RegistrarRequest_Action = 0
	RegistrarRequest_DEREGISTER RegistrarRequest_Action = 1
)

var RegistrarRequest_Action_name = map[int32]string{
	0: "REGISTER",
	1: "DEREGISTER",
}
var RegistrarRequest_Action_value = map[string]int32{
	"REGISTER":   0,
	"DEREGISTER": 1,
}

func (x RegistrarRequest_Action) String() string {
	return proto.EnumName(RegistrarRequest_Action_name, int32(x))
}
func (RegistrarRequest_Action) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// The requested listener and cluster member
type RegistrarRequest struct {
	FrontendAddr string                  `protobuf:"bytes,1,opt,name=frontend_addr,json=frontendAddr" json:"frontend_addr,omitempty"`
	FrontendPort int32                   `protobuf:"varint,2,opt,name=frontend_port,json=frontendPort" json:"frontend_port,omitempty"`
	BackendAddr  string                  `protobuf:"bytes,3,opt,name=backend_addr,json=backendAddr" json:"backend_addr,omitempty"`
	BackendPort  int32                   `protobuf:"varint,4,opt,name=backend_port,json=backendPort" json:"backend_port,omitempty"`
	Action       RegistrarRequest_Action `protobuf:"varint,5,opt,name=action,enum=envoyrpc.RegistrarRequest_Action" json:"action,omitempty"`
}

func (m *RegistrarRequest) Reset()                    { *m = RegistrarRequest{} }
func (m *RegistrarRequest) String() string            { return proto.CompactTextString(m) }
func (*RegistrarRequest) ProtoMessage()               {}
func (*RegistrarRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegistrarRequest) GetFrontendAddr() string {
	if m != nil {
		return m.FrontendAddr
	}
	return ""
}

func (m *RegistrarRequest) GetFrontendPort() int32 {
	if m != nil {
		return m.FrontendPort
	}
	return 0
}

func (m *RegistrarRequest) GetBackendAddr() string {
	if m != nil {
		return m.BackendAddr
	}
	return ""
}

func (m *RegistrarRequest) GetBackendPort() int32 {
	if m != nil {
		return m.BackendPort
	}
	return 0
}

func (m *RegistrarRequest) GetAction() RegistrarRequest_Action {
	if m != nil {
		return m.Action
	}
	return RegistrarRequest_REGISTER
}

// The response message containing the status
type RegistrarReply struct {
	StatusCode int32 `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
}

func (m *RegistrarReply) Reset()                    { *m = RegistrarReply{} }
func (m *RegistrarReply) String() string            { return proto.CompactTextString(m) }
func (*RegistrarReply) ProtoMessage()               {}
func (*RegistrarReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegistrarReply) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func init() {
	proto.RegisterType((*RegistrarRequest)(nil), "envoyrpc.RegistrarRequest")
	proto.RegisterType((*RegistrarReply)(nil), "envoyrpc.RegistrarReply")
	proto.RegisterEnum("envoyrpc.RegistrarRequest_Action", RegistrarRequest_Action_name, RegistrarRequest_Action_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Registrar service

type RegistrarClient interface {
	Register(ctx context.Context, in *RegistrarRequest, opts ...grpc.CallOption) (*RegistrarReply, error)
}

type registrarClient struct {
	cc *grpc.ClientConn
}

func NewRegistrarClient(cc *grpc.ClientConn) RegistrarClient {
	return &registrarClient{cc}
}

func (c *registrarClient) Register(ctx context.Context, in *RegistrarRequest, opts ...grpc.CallOption) (*RegistrarReply, error) {
	out := new(RegistrarReply)
	err := grpc.Invoke(ctx, "/envoyrpc.Registrar/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registrar service

type RegistrarServer interface {
	Register(context.Context, *RegistrarRequest) (*RegistrarReply, error)
}

func RegisterRegistrarServer(s *grpc.Server, srv RegistrarServer) {
	s.RegisterService(&_Registrar_serviceDesc, srv)
}

func _Registrar_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrarServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoyrpc.Registrar/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrarServer).Register(ctx, req.(*RegistrarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registrar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoyrpc.Registrar",
	HandlerType: (*RegistrarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registrar_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoyrpc.proto",
}

func init() { proto.RegisterFile("envoyrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xeb, 0x7e, 0x5f, 0xab, 0xf4, 0x9a, 0x46, 0x95, 0xc5, 0x10, 0x75, 0x21, 0x0d, 0x12,
	0x64, 0x8a, 0x44, 0x99, 0x18, 0x1b, 0x88, 0x10, 0x13, 0x91, 0x61, 0xaf, 0x52, 0xdb, 0x54, 0x11,
	0x25, 0x67, 0x1c, 0x17, 0x29, 0x3f, 0x80, 0xff, 0x8d, 0xea, 0x34, 0xa1, 0x20, 0xd8, 0x7c, 0xaf,
	0x9f, 0x7b, 0x74, 0xf6, 0x81, 0x27, 0xcb, 0x77, 0xac, 0xb5, 0xe2, 0xb1, 0xd2, 0x68, 0x90, 0x3a,
	0x6d, 0x1d, 0x7e, 0xf4, 0x61, 0xca, 0xe4, 0xa6, 0xa8, 0x8c, 0xce, 0x35, 0x93, 0x6f, 0x3b, 0x59,
	0x19, 0x7a, 0x06, 0x93, 0x67, 0x8d, 0xa5, 0x91, 0xa5, 0x58, 0xe5, 0x42, 0x68, 0x9f, 0x04, 0x24,
	0x1a, 0x31, 0xb7, 0x0d, 0x97, 0x42, 0xe8, 0x6f, 0x90, 0x42, 0x6d, 0xfc, 0x7e, 0x40, 0xa2, 0xc1,
	0x17, 0x94, 0xa1, 0x36, 0x74, 0x0e, 0xee, 0x3a, 0xe7, 0x2f, 0x9d, 0xe8, 0x9f, 0x15, 0x8d, 0x0f,
	0x99, 0xf5, 0x1c, 0x21, 0x56, 0xf3, 0xdf, 0x6a, 0x5a, 0xc4, 0x5a, 0xae, 0x61, 0x98, 0x73, 0x53,
	0x60, 0xe9, 0x0f, 0x02, 0x12, 0x79, 0x8b, 0x79, 0xdc, 0xbd, 0xe7, 0xe7, 0xec, 0xf1, 0xd2, 0x82,
	0xec, 0xd0, 0x10, 0x9e, 0xc3, 0xb0, 0x49, 0xa8, 0x0b, 0x0e, 0x4b, 0xef, 0xee, 0x1f, 0x9f, 0x52,
	0x36, 0xed, 0x51, 0x0f, 0xe0, 0x36, 0xed, 0x6a, 0x12, 0x5e, 0x82, 0x77, 0xa4, 0x52, 0xdb, 0x9a,
	0x9e, 0xc2, 0xb8, 0x32, 0xb9, 0xd9, 0x55, 0x2b, 0x8e, 0x42, 0xda, 0x2f, 0x18, 0x30, 0x68, 0xa2,
	0x1b, 0x14, 0x72, 0xf1, 0x00, 0xa3, 0xae, 0x85, 0x26, 0xe0, 0x34, 0x85, 0xd4, 0x74, 0xf6, 0xf7,
	0x78, 0x33, 0xff, 0xd7, 0x3b, 0xb5, 0xad, 0xc3, 0x5e, 0x72, 0x01, 0x27, 0x1c, 0x5f, 0xe3, 0x0d,
	0x96, 0x85, 0xd1, 0xd8, 0x81, 0xc9, 0x24, 0xdd, 0x9f, 0x98, 0xe2, 0xd9, 0x7e, 0x79, 0x19, 0x59,
	0x0f, 0xed, 0x16, 0xaf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xc4, 0x0d, 0xbe, 0xd7, 0x01,
	0x00, 0x00,
}
