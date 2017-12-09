// Code generated by protoc-gen-go. DO NOT EDIT.
// source: changer.proto

/*
Package changer is a generated protocol buffer package.

It is generated from these files:
	changer.proto

It has these top-level messages:
	ManagersRequest
	ManagersReply
	PasswdConfig
	ChangeRequest
	Status
*/
package changer

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

type ChangeRequest_Manager int32

const (
	ChangeRequest_LASTPASS ChangeRequest_Manager = 0
)

var ChangeRequest_Manager_name = map[int32]string{
	0: "LASTPASS",
}
var ChangeRequest_Manager_value = map[string]int32{
	"LASTPASS": 0,
}

func (x ChangeRequest_Manager) String() string {
	return proto.EnumName(ChangeRequest_Manager_name, int32(x))
}
func (ChangeRequest_Manager) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Status_Type int32

const (
	Status_JOB_START   Status_Type = 0
	Status_TASK_START  Status_Type = 1
	Status_TASK_ERROR  Status_Type = 2
	Status_TASK_FINISH Status_Type = 3
	Status_JOB_FINISH  Status_Type = 4
)

var Status_Type_name = map[int32]string{
	0: "JOB_START",
	1: "TASK_START",
	2: "TASK_ERROR",
	3: "TASK_FINISH",
	4: "JOB_FINISH",
}
var Status_Type_value = map[string]int32{
	"JOB_START":   0,
	"TASK_START":  1,
	"TASK_ERROR":  2,
	"TASK_FINISH": 3,
	"JOB_FINISH":  4,
}

func (x Status_Type) String() string {
	return proto.EnumName(Status_Type_name, int32(x))
}
func (Status_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type ManagersRequest struct {
}

func (m *ManagersRequest) Reset()                    { *m = ManagersRequest{} }
func (m *ManagersRequest) String() string            { return proto.CompactTextString(m) }
func (*ManagersRequest) ProtoMessage()               {}
func (*ManagersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ManagersReply struct {
	Managers []string `protobuf:"bytes,1,rep,name=managers" json:"managers,omitempty"`
}

func (m *ManagersReply) Reset()                    { *m = ManagersReply{} }
func (m *ManagersReply) String() string            { return proto.CompactTextString(m) }
func (*ManagersReply) ProtoMessage()               {}
func (*ManagersReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ManagersReply) GetManagers() []string {
	if m != nil {
		return m.Managers
	}
	return nil
}

type PasswdConfig struct {
	Length       uint32 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
	Numbers      bool   `protobuf:"varint,2,opt,name=numbers" json:"numbers,omitempty"`
	SpecialChars bool   `protobuf:"varint,3,opt,name=specialChars" json:"specialChars,omitempty"`
}

func (m *PasswdConfig) Reset()                    { *m = PasswdConfig{} }
func (m *PasswdConfig) String() string            { return proto.CompactTextString(m) }
func (*PasswdConfig) ProtoMessage()               {}
func (*PasswdConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PasswdConfig) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *PasswdConfig) GetNumbers() bool {
	if m != nil {
		return m.Numbers
	}
	return false
}

func (m *PasswdConfig) GetSpecialChars() bool {
	if m != nil {
		return m.SpecialChars
	}
	return false
}

// The request message containing the user's name.
type ChangeRequest struct {
	Manager      ChangeRequest_Manager `protobuf:"varint,1,opt,name=manager,enum=changer.ChangeRequest_Manager" json:"manager,omitempty"`
	Email        string                `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password     string                `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	PasswdConfig *PasswdConfig         `protobuf:"bytes,4,opt,name=passwdConfig" json:"passwdConfig,omitempty"`
}

func (m *ChangeRequest) Reset()                    { *m = ChangeRequest{} }
func (m *ChangeRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangeRequest) ProtoMessage()               {}
func (*ChangeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ChangeRequest) GetManager() ChangeRequest_Manager {
	if m != nil {
		return m.Manager
	}
	return ChangeRequest_LASTPASS
}

func (m *ChangeRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ChangeRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ChangeRequest) GetPasswdConfig() *PasswdConfig {
	if m != nil {
		return m.PasswdConfig
	}
	return nil
}

// Status interface
// Start status
// Job start status (with subJob ID)
// Job Error status
// Job finish status
// Finish status
type Status struct {
	Type     Status_Type `protobuf:"varint,1,opt,name=type,enum=changer.Status_Type" json:"type,omitempty"`
	JobId    uint64      `protobuf:"varint,2,opt,name=jobId" json:"jobId,omitempty"`
	TaskId   uint64      `protobuf:"varint,3,opt,name=taskId" json:"taskId,omitempty"`
	Hostname string      `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	Email    string      `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Msg      string      `protobuf:"bytes,6,opt,name=msg" json:"msg,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Status) GetType() Status_Type {
	if m != nil {
		return m.Type
	}
	return Status_JOB_START
}

func (m *Status) GetJobId() uint64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *Status) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *Status) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Status) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ManagersRequest)(nil), "changer.ManagersRequest")
	proto.RegisterType((*ManagersReply)(nil), "changer.ManagersReply")
	proto.RegisterType((*PasswdConfig)(nil), "changer.PasswdConfig")
	proto.RegisterType((*ChangeRequest)(nil), "changer.ChangeRequest")
	proto.RegisterType((*Status)(nil), "changer.Status")
	proto.RegisterEnum("changer.ChangeRequest_Manager", ChangeRequest_Manager_name, ChangeRequest_Manager_value)
	proto.RegisterEnum("changer.Status_Type", Status_Type_name, Status_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Changer service

type ChangerClient interface {
	// Sends a greeting
	ChangePassword(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (Changer_ChangePasswordClient, error)
	GetManagers(ctx context.Context, in *ManagersRequest, opts ...grpc.CallOption) (*ManagersReply, error)
}

type changerClient struct {
	cc *grpc.ClientConn
}

func NewChangerClient(cc *grpc.ClientConn) ChangerClient {
	return &changerClient{cc}
}

func (c *changerClient) ChangePassword(ctx context.Context, in *ChangeRequest, opts ...grpc.CallOption) (Changer_ChangePasswordClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Changer_serviceDesc.Streams[0], c.cc, "/changer.Changer/ChangePassword", opts...)
	if err != nil {
		return nil, err
	}
	x := &changerChangePasswordClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Changer_ChangePasswordClient interface {
	Recv() (*Status, error)
	grpc.ClientStream
}

type changerChangePasswordClient struct {
	grpc.ClientStream
}

func (x *changerChangePasswordClient) Recv() (*Status, error) {
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *changerClient) GetManagers(ctx context.Context, in *ManagersRequest, opts ...grpc.CallOption) (*ManagersReply, error) {
	out := new(ManagersReply)
	err := grpc.Invoke(ctx, "/changer.Changer/GetManagers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Changer service

type ChangerServer interface {
	// Sends a greeting
	ChangePassword(*ChangeRequest, Changer_ChangePasswordServer) error
	GetManagers(context.Context, *ManagersRequest) (*ManagersReply, error)
}

func RegisterChangerServer(s *grpc.Server, srv ChangerServer) {
	s.RegisterService(&_Changer_serviceDesc, srv)
}

func _Changer_ChangePassword_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChangerServer).ChangePassword(m, &changerChangePasswordServer{stream})
}

type Changer_ChangePasswordServer interface {
	Send(*Status) error
	grpc.ServerStream
}

type changerChangePasswordServer struct {
	grpc.ServerStream
}

func (x *changerChangePasswordServer) Send(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func _Changer_GetManagers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangerServer).GetManagers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/changer.Changer/GetManagers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangerServer).GetManagers(ctx, req.(*ManagersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Changer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "changer.Changer",
	HandlerType: (*ChangerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagers",
			Handler:    _Changer_GetManagers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChangePassword",
			Handler:       _Changer_ChangePassword_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "changer.proto",
}

func init() { proto.RegisterFile("changer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xeb, 0x69, 0xa6, 0x97, 0xd3, 0x2b, 0xd6, 0x50, 0xa2, 0x2e, 0x50, 0xe5, 0x05, 0xaa,
	0x84, 0x14, 0xa1, 0xb2, 0x19, 0x16, 0x2c, 0xda, 0x8a, 0x4b, 0xb9, 0x4d, 0xe5, 0x44, 0x6c, 0x91,
	0xd3, 0x9a, 0x24, 0x90, 0x1b, 0xb1, 0xab, 0x51, 0x9f, 0x81, 0x87, 0xe3, 0x85, 0x58, 0x8c, 0xec,
	0x38, 0x99, 0x76, 0x34, 0x3b, 0x7f, 0xc7, 0x27, 0x3e, 0xff, 0xff, 0xdb, 0x81, 0xc1, 0x2e, 0x64,
	0x69, 0xc0, 0x0b, 0x27, 0x2f, 0x32, 0x99, 0xe1, 0xb6, 0x41, 0xf2, 0x04, 0x46, 0x5f, 0x59, 0xca,
	0x02, 0x5e, 0x08, 0xca, 0xff, 0x1c, 0xb8, 0x90, 0xe4, 0x25, 0x0c, 0xee, 0x4b, 0x79, 0x7c, 0xc4,
	0x53, 0xe8, 0x24, 0xa6, 0x60, 0xa3, 0x59, 0x73, 0xde, 0xa5, 0x35, 0x93, 0x3d, 0xf4, 0xb7, 0x4c,
	0x88, 0xdb, 0xfd, 0x3a, 0x4b, 0x7f, 0x46, 0x01, 0x9e, 0x40, 0x2b, 0xe6, 0x69, 0x20, 0x43, 0x1b,
	0xcd, 0xd0, 0x7c, 0x40, 0x0d, 0x61, 0x1b, 0xda, 0xe9, 0x21, 0xf1, 0xd5, 0x11, 0x17, 0x33, 0x34,
	0xef, 0xd0, 0x0a, 0x31, 0x81, 0xbe, 0xc8, 0xf9, 0x2e, 0x62, 0xf1, 0x3a, 0x64, 0x85, 0xb0, 0x9b,
	0x7a, 0xfb, 0xac, 0x46, 0xfe, 0x21, 0x18, 0xac, 0xb5, 0x62, 0x23, 0x12, 0x5f, 0x43, 0xdb, 0x68,
	0xd0, 0x83, 0x86, 0x8b, 0xe7, 0x4e, 0xe5, 0xf0, 0xac, 0xd1, 0x31, 0x56, 0x68, 0xd5, 0x8e, 0xaf,
	0xe0, 0x92, 0x27, 0x2c, 0x8a, 0xb5, 0x8e, 0x2e, 0x2d, 0x41, 0x79, 0xcc, 0x95, 0x8f, 0xac, 0xd8,
	0x6b, 0x05, 0x5d, 0x5a, 0x33, 0x7e, 0x03, 0xfd, 0xfc, 0xc4, 0xa3, 0x6d, 0xcd, 0xd0, 0xbc, 0xb7,
	0x78, 0x5a, 0x0f, 0x3c, 0x0d, 0x80, 0x9e, 0xb5, 0x92, 0x67, 0xd0, 0x36, 0x02, 0x70, 0x1f, 0x3a,
	0x5f, 0x96, 0xae, 0xb7, 0x5d, 0xba, 0xee, 0xb8, 0x41, 0xfe, 0x23, 0x68, 0xb9, 0x92, 0xc9, 0x83,
	0xc0, 0x73, 0xb0, 0xe4, 0x31, 0xe7, 0xc6, 0xc7, 0x55, 0x7d, 0x6c, 0xb9, 0xed, 0x78, 0xc7, 0x9c,
	0x53, 0xdd, 0xa1, 0xa4, 0xff, 0xca, 0xfc, 0xcd, 0x5e, 0x4b, 0xb7, 0x68, 0x09, 0x2a, 0x72, 0xc9,
	0xc4, 0xef, 0x4d, 0x29, 0xdc, 0xa2, 0x86, 0x94, 0xa5, 0x30, 0x13, 0x32, 0x65, 0x09, 0xd7, 0x92,
	0xbb, 0xb4, 0xe6, 0xfb, 0x10, 0x2e, 0x4f, 0x43, 0x18, 0x43, 0x33, 0x11, 0x81, 0xdd, 0xd2, 0x35,
	0xb5, 0x24, 0xdf, 0xc1, 0x52, 0xf3, 0xf1, 0x00, 0xba, 0x9f, 0x6e, 0x56, 0x3f, 0x5c, 0x6f, 0x49,
	0xbd, 0x71, 0x03, 0x0f, 0x01, 0xbc, 0xa5, 0xfb, 0xd9, 0x30, 0xaa, 0xf9, 0x1d, 0xa5, 0x37, 0x74,
	0x7c, 0x81, 0x47, 0xd0, 0xd3, 0xfc, 0x7e, 0xf3, 0x6d, 0xe3, 0x7e, 0x1c, 0x37, 0x55, 0x83, 0xfa,
	0xde, 0xb0, 0xb5, 0xf8, 0x8b, 0xa0, 0x5d, 0xde, 0x53, 0x81, 0xdf, 0xc2, 0xb0, 0x5c, 0x6e, 0xab,
	0xc0, 0x27, 0x8f, 0xdf, 0xe5, 0x74, 0xf4, 0x20, 0x1b, 0xd2, 0x78, 0x85, 0xf0, 0x12, 0x7a, 0x1f,
	0xb8, 0xac, 0x5e, 0x2c, 0xb6, 0xeb, 0x9e, 0x07, 0xef, 0x7a, 0x3a, 0x79, 0x64, 0x27, 0x8f, 0x8f,
	0xa4, 0xb1, 0xba, 0x86, 0x17, 0xbb, 0x2c, 0x71, 0x82, 0x48, 0x86, 0x07, 0xdf, 0xb9, 0x0d, 0xa3,
	0x98, 0xc7, 0x59, 0x96, 0x3b, 0x05, 0x4f, 0x78, 0xc2, 0x7d, 0x5e, 0x24, 0xbc, 0xfc, 0x6f, 0x56,
	0x7d, 0x23, 0x7a, 0xab, 0x68, 0x8b, 0xfc, 0x96, 0x2e, 0xbf, 0xbe, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x85, 0xf6, 0x10, 0x3f, 0x5f, 0x03, 0x00, 0x00,
}