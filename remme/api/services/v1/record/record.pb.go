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
	JobEvent
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

type JobEvent_Type int32

const (
	JobEvent_JOB_START   JobEvent_Type = 0
	JobEvent_TASK_START  JobEvent_Type = 1
	JobEvent_TASK_ERROR  JobEvent_Type = 2
	JobEvent_TASK_FINISH JobEvent_Type = 3
	JobEvent_JOB_FINISH  JobEvent_Type = 4
)

var JobEvent_Type_name = map[int32]string{
	0: "JOB_START",
	1: "TASK_START",
	2: "TASK_ERROR",
	3: "TASK_FINISH",
	4: "JOB_FINISH",
}
var JobEvent_Type_value = map[string]int32{
	"JOB_START":   0,
	"TASK_START":  1,
	"TASK_ERROR":  2,
	"TASK_FINISH": 3,
	"JOB_FINISH":  4,
}

func (x JobEvent_Type) String() string {
	return proto.EnumName(JobEvent_Type_name, int32(x))
}
func (JobEvent_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

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

// Status interface
// Start status
// Job start status (with subJob ID)
// Job Error status
// Job finish status
// Finish status
type JobEvent struct {
	Type      JobEvent_Type `protobuf:"varint,1,opt,name=type,enum=record.JobEvent_Type" json:"type,omitempty"`
	JobId     uint64        `protobuf:"varint,2,opt,name=jobId" json:"jobId,omitempty"`
	TaskId    uint64        `protobuf:"varint,3,opt,name=taskId" json:"taskId,omitempty"`
	Timestamp uint64        `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Hostname  string        `protobuf:"bytes,5,opt,name=hostname" json:"hostname,omitempty"`
	Email     string        `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
	Msg       string        `protobuf:"bytes,7,opt,name=msg" json:"msg,omitempty"`
	Version   string        `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`
}

func (m *JobEvent) Reset()                    { *m = JobEvent{} }
func (m *JobEvent) String() string            { return proto.CompactTextString(m) }
func (*JobEvent) ProtoMessage()               {}
func (*JobEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *JobEvent) GetType() JobEvent_Type {
	if m != nil {
		return m.Type
	}
	return JobEvent_JOB_START
}

func (m *JobEvent) GetJobId() uint64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *JobEvent) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *JobEvent) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *JobEvent) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *JobEvent) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *JobEvent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *JobEvent) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*RecordRequest)(nil), "record.RecordRequest")
	proto.RegisterType((*Failure)(nil), "record.Failure")
	proto.RegisterType((*LogRecord)(nil), "record.LogRecord")
	proto.RegisterType((*JobEvent)(nil), "record.JobEvent")
	proto.RegisterEnum("record.JobEvent_Type", JobEvent_Type_name, JobEvent_Type_value)
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
	TailEvents(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (Record_TailEventsClient, error)
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

func (c *recordClient) TailEvents(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (Record_TailEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Record_serviceDesc.Streams[0], c.cc, "/record.Record/TailEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordTailEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Record_TailEventsClient interface {
	Recv() (*JobEvent, error)
	grpc.ClientStream
}

type recordTailEventsClient struct {
	grpc.ClientStream
}

func (x *recordTailEventsClient) Recv() (*JobEvent, error) {
	m := new(JobEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Record service

type RecordServer interface {
	GetRecord(context.Context, *RecordRequest) (*LogRecord, error)
	TailEvents(*RecordRequest, Record_TailEventsServer) error
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

func _Record_TailEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecordRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordServer).TailEvents(m, &recordTailEventsServer{stream})
}

type Record_TailEventsServer interface {
	Send(*JobEvent) error
	grpc.ServerStream
}

type recordTailEventsServer struct {
	grpc.ServerStream
}

func (x *recordTailEventsServer) Send(m *JobEvent) error {
	return x.ServerStream.SendMsg(m)
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TailEvents",
			Handler:       _Record_TailEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "record.proto",
}

func init() { proto.RegisterFile("record.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x36, 0x4d, 0x9b, 0x57, 0xba, 0x05, 0x6b, 0x4c, 0xd1, 0x84, 0x50, 0x95, 0x03,
	0x2a, 0x42, 0x8a, 0x50, 0x39, 0x4c, 0x88, 0xd3, 0x2a, 0x6d, 0xd0, 0x81, 0xd8, 0xe4, 0x46, 0x5c,
	0x51, 0xd2, 0x3e, 0x5a, 0x43, 0x12, 0x07, 0xdb, 0x1d, 0x9a, 0xc4, 0xdf, 0xc4, 0x95, 0x7f, 0x0f,
	0xc5, 0x76, 0xb3, 0x76, 0x52, 0x6f, 0xfe, 0x7c, 0xfd, 0x5e, 0xde, 0x8f, 0xaf, 0x03, 0x4f, 0x04,
	0x2e, 0xb8, 0x58, 0xc6, 0x95, 0xe0, 0x8a, 0x13, 0xcf, 0x50, 0xf4, 0x1e, 0x86, 0x54, 0x9f, 0x28,
	0xfe, 0xda, 0xa0, 0x54, 0xe4, 0x04, 0xba, 0x3f, 0x78, 0x36, 0x5b, 0x86, 0xce, 0xc8, 0x19, 0xbb,
	0xd4, 0x40, 0xad, 0x4a, 0x56, 0x2e, 0x30, 0x6c, 0x8f, 0x9c, 0xf1, 0x90, 0x1a, 0x88, 0x0a, 0xe8,
	0x5d, 0xa5, 0x2c, 0xdf, 0x08, 0x24, 0x67, 0xd0, 0x5f, 0x73, 0xa9, 0xca, 0xb4, 0x40, 0x9d, 0xe9,
	0xd3, 0x86, 0xeb, 0x64, 0x2c, 0x52, 0x96, 0xeb, 0x64, 0x9f, 0x1a, 0x20, 0xa7, 0xe0, 0x09, 0x4c,
	0x25, 0x2f, 0xc3, 0x8e, 0x96, 0x2d, 0x91, 0x10, 0x7a, 0x77, 0x28, 0x24, 0xe3, 0x65, 0xe8, 0xea,
	0x8b, 0x2d, 0x46, 0x7f, 0x1d, 0xf0, 0x3f, 0xf3, 0x95, 0xe9, 0x97, 0x10, 0x70, 0x15, 0xb3, 0xd5,
	0x5c, 0xaa, 0xcf, 0x0f, 0xcd, 0xb7, 0x1f, 0x35, 0x6f, 0xea, 0x77, 0x76, 0xeb, 0x9f, 0x40, 0x57,
	0x09, 0x86, 0x52, 0x57, 0x71, 0xa9, 0x01, 0xf2, 0x02, 0x40, 0x71, 0x95, 0xe6, 0x73, 0xa6, 0x50,
	0x86, 0x5d, 0x7d, 0xb5, 0xa3, 0x90, 0xd7, 0xd0, 0xff, 0x6e, 0x46, 0x96, 0xa1, 0x37, 0xea, 0x8c,
	0x07, 0x93, 0xe3, 0xd8, 0x2e, 0xd6, 0xae, 0x82, 0x36, 0x01, 0xd1, 0xbf, 0x36, 0xf4, 0xaf, 0x79,
	0x76, 0x79, 0x87, 0xa5, 0x22, 0xaf, 0xc0, 0x55, 0xf7, 0x95, 0xe9, 0xf7, 0x68, 0xf2, 0x6c, 0x9b,
	0xb5, 0xbd, 0x8f, 0x93, 0xfb, 0x0a, 0xa9, 0x0e, 0x39, 0x30, 0xc6, 0x29, 0x78, 0x2a, 0x95, 0x3f,
	0x67, 0x4b, 0x3d, 0x87, 0x4b, 0x2d, 0x91, 0xe7, 0xe0, 0xd7, 0xc3, 0x4b, 0x95, 0x16, 0x95, 0x1d,
	0xe6, 0x41, 0xd8, 0x33, 0xa6, 0x7b, 0xc8, 0x18, 0x6f, 0x77, 0x31, 0x01, 0x74, 0x0a, 0xb9, 0x0a,
	0x7b, 0x5a, 0xab, 0x8f, 0xbb, 0x96, 0xf4, 0xf7, 0x2d, 0xf9, 0x0a, 0x6e, 0xdd, 0x37, 0x19, 0x82,
	0x7f, 0x7d, 0x33, 0xfd, 0x36, 0x4f, 0x2e, 0x68, 0x12, 0xb4, 0xc8, 0x11, 0x40, 0x72, 0x31, 0xff,
	0x64, 0xd9, 0x69, 0xf8, 0x92, 0xd2, 0x1b, 0x1a, 0xb4, 0xc9, 0x31, 0x0c, 0x34, 0x5f, 0xcd, 0xbe,
	0xcc, 0xe6, 0x1f, 0x83, 0x4e, 0x1d, 0x50, 0xe7, 0x5b, 0x76, 0x27, 0x7f, 0xc0, 0xb3, 0x36, 0x9f,
	0x83, 0xff, 0x01, 0x95, 0x85, 0x66, 0x6b, 0x7b, 0x6f, 0xf6, 0xec, 0xe9, 0x56, 0x6e, 0x5e, 0x47,
	0xd4, 0x22, 0xef, 0x00, 0x92, 0x94, 0xe5, 0x7a, 0xb9, 0xf2, 0x50, 0x66, 0xf0, 0xd8, 0x86, 0xa8,
	0xf5, 0xc6, 0x99, 0x9e, 0xc3, 0xcb, 0x05, 0x2f, 0xe2, 0x15, 0x53, 0xeb, 0x4d, 0x16, 0xff, 0x5e,
	0xb3, 0x1c, 0x73, 0xce, 0xab, 0x58, 0x60, 0x81, 0x05, 0x66, 0x28, 0x0a, 0x34, 0xbf, 0xd1, 0x74,
	0x60, 0x3e, 0x77, 0x5b, 0xc3, 0xad, 0x93, 0x79, 0x5a, 0x7d, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0xdf, 0xd1, 0xc5, 0x86, 0x6c, 0x03, 0x00, 0x00,
}
