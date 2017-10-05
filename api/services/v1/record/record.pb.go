// Code generated by protoc-gen-go. DO NOT EDIT.
// source: record.proto

/*
Package record is a generated protocol buffer package.

It is generated from these files:
	record.proto

It has these top-level messages:
	RecordRequest
	Failure
	LogRecord
*/
package record

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

type RecordRequest struct {
	JobId uint64 `protobuf:"varint,1,opt,name=jobId" json:"jobId,omitempty"`
	Since uint32 `protobuf:"varint,2,opt,name=since" json:"since,omitempty"`
}

func (m *RecordRequest) Reset()                    { *m = RecordRequest{} }
func (m *RecordRequest) String() string            { return proto.CompactTextString(m) }
func (*RecordRequest) ProtoMessage()               {}
func (*RecordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecordRequest) GetJobId() uint64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *RecordRequest) GetSince() uint32 {
	if m != nil {
		return m.Since
	}
	return 0
}

type Failure struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Reason   string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	Version  string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
}

func (m *Failure) Reset()                    { *m = Failure{} }
func (m *Failure) String() string            { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()               {}
func (*Failure) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Failure) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Failure) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Failure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Failure) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type LogRecord struct {
	Time       uint64     `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	JobId      uint64     `protobuf:"varint,2,opt,name=jobId" json:"jobId,omitempty"`
	Email      string     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Tries      uint64     `protobuf:"varint,4,opt,name=tries" json:"tries,omitempty"`
	TotalSites uint64     `protobuf:"varint,5,opt,name=totalSites" json:"totalSites,omitempty"`
	Failures   []*Failure `protobuf:"bytes,6,rep,name=failures" json:"failures,omitempty"`
}

func (m *LogRecord) Reset()                    { *m = LogRecord{} }
func (m *LogRecord) String() string            { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()               {}
func (*LogRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogRecord) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogRecord) GetJobId() uint64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *LogRecord) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LogRecord) GetTries() uint64 {
	if m != nil {
		return m.Tries
	}
	return 0
}

func (m *LogRecord) GetTotalSites() uint64 {
	if m != nil {
		return m.TotalSites
	}
	return 0
}

func (m *LogRecord) GetFailures() []*Failure {
	if m != nil {
		return m.Failures
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordRequest)(nil), "record.RecordRequest")
	proto.RegisterType((*Failure)(nil), "record.Failure")
	proto.RegisterType((*LogRecord)(nil), "record.LogRecord")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Record service

type RecordClient interface {
	GetRecord(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*LogRecord, error)
}

type recordClient struct {
	cc *grpc.ClientConn
}

func NewRecordClient(cc *grpc.ClientConn) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) GetRecord(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*LogRecord, error) {
	out := new(LogRecord)
	err := grpc.Invoke(ctx, "/record.Record/GetRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Record service

type RecordServer interface {
	GetRecord(context.Context, *RecordRequest) (*LogRecord, error)
}

func RegisterRecordServer(s *grpc.Server, srv RecordServer) {
	s.RegisterService(&_Record_serviceDesc, srv)
}

func _Record_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/record.Record/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).GetRecord(ctx, req.(*RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Record_serviceDesc = grpc.ServiceDesc{
	ServiceName: "record.Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecord",
			Handler:    _Record_GetRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record.proto",
}

func init() { proto.RegisterFile("record.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0xff, 0x85, 0x52, 0xe8, 0xf0, 0x27, 0xc6, 0x0d, 0x9a, 0x0d, 0x07, 0x43, 0x7a, 0x30,
	0x24, 0x26, 0x3d, 0xe0, 0x81, 0x83, 0x27, 0x39, 0x68, 0x4c, 0x3c, 0x90, 0xf5, 0x13, 0xb4, 0x30,
	0xc2, 0x9a, 0x6e, 0x07, 0x77, 0x17, 0xfd, 0x56, 0x7e, 0x46, 0xc3, 0xee, 0xb6, 0xa2, 0xb7, 0xfd,
	0xbd, 0xe9, 0xcb, 0x7b, 0x33, 0x85, 0xff, 0x1a, 0xd7, 0xa4, 0x37, 0xf9, 0x5e, 0x93, 0x25, 0x96,
	0x78, 0xca, 0xee, 0x60, 0x24, 0xdc, 0x4b, 0xe0, 0xfb, 0x01, 0x8d, 0x65, 0x63, 0xe8, 0xbd, 0x51,
	0xf9, 0xb4, 0xe1, 0xd1, 0x34, 0x9a, 0xc5, 0xc2, 0xc3, 0x51, 0x35, 0xb2, 0x5e, 0x23, 0xef, 0x4c,
	0xa3, 0xd9, 0x48, 0x78, 0xc8, 0x14, 0xf4, 0x1f, 0x0a, 0x59, 0x1d, 0x34, 0xb2, 0x09, 0x0c, 0x76,
	0x64, 0x6c, 0x5d, 0x28, 0x74, 0xce, 0x54, 0xb4, 0x7c, 0x34, 0xa3, 0x2a, 0x64, 0xe5, 0xcc, 0xa9,
	0xf0, 0xc0, 0x2e, 0x21, 0xd1, 0x58, 0x18, 0xaa, 0x79, 0xd7, 0xc9, 0x81, 0x18, 0x87, 0xfe, 0x07,
	0x6a, 0x23, 0xa9, 0xe6, 0xb1, 0x1b, 0x34, 0x98, 0x7d, 0x45, 0x90, 0x3e, 0xd3, 0xd6, 0xf7, 0x65,
	0x0c, 0x62, 0x2b, 0x43, 0x5a, 0x2c, 0xdc, 0xfb, 0xa7, 0x7c, 0xe7, 0x4f, 0x79, 0x9f, 0xdf, 0x3d,
	0xcd, 0x1f, 0x43, 0xcf, 0x6a, 0x89, 0xc6, 0xa5, 0xc4, 0xc2, 0x03, 0xbb, 0x02, 0xb0, 0x64, 0x8b,
	0xea, 0x45, 0x5a, 0x34, 0xbc, 0xe7, 0x46, 0x27, 0x0a, 0xbb, 0x81, 0xc1, 0xab, 0x5f, 0xd9, 0xf0,
	0x64, 0xda, 0x9d, 0x0d, 0xe7, 0x67, 0x79, 0x38, 0x6c, 0x38, 0x85, 0x68, 0x3f, 0x98, 0xdf, 0x43,
	0x12, 0xca, 0x2e, 0x20, 0x7d, 0x44, 0x1b, 0xe0, 0xa2, 0x71, 0xfc, 0xba, 0xfc, 0xe4, 0xbc, 0x91,
	0xdb, 0x1d, 0xb3, 0x7f, 0xcb, 0x05, 0x5c, 0xaf, 0x49, 0xe5, 0x5b, 0x69, 0x77, 0x87, 0x32, 0xff,
	0xdc, 0xc9, 0x0a, 0x2b, 0xa2, 0x7d, 0xae, 0x51, 0xa1, 0xc2, 0x12, 0xb5, 0x42, 0xff, 0x47, 0x97,
	0x43, 0xef, 0x59, 0x1d, 0x61, 0x15, 0x95, 0x89, 0x53, 0x6f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x30, 0x61, 0xf4, 0x6d, 0xf7, 0x01, 0x00, 0x00,
}
