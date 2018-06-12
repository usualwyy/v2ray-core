package command

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	"context"

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

type GetStatsRequest struct {
	// Name of the stat counter.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Whether or not to reset the counter to fetching its value.
	Reset_               bool     `protobuf:"varint,2,opt,name=reset" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_a956deb495595f76, []int{0}
}
func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (dst *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(dst, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type Stat struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_a956deb495595f76, []int{1}
}
func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (dst *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(dst, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stat) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetStatsResponse struct {
	Stat                 *Stat    `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_a956deb495595f76, []int{2}
}
func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (dst *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(dst, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetStat() *Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_command_a956deb495595f76, []int{3}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetStatsRequest)(nil), "v2ray.core.app.stats.command.GetStatsRequest")
	proto.RegisterType((*Stat)(nil), "v2ray.core.app.stats.command.Stat")
	proto.RegisterType((*GetStatsResponse)(nil), "v2ray.core.app.stats.command.GetStatsResponse")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.stats.command.Config")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
}

type statsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatsServiceClient(cc *grpc.ClientConn) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
type StatsServiceServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
}

func RegisterStatsServiceServer(s *grpc.Server, srv StatsServiceServer) {
	s.RegisterService(&_StatsService_serviceDesc, srv)
}

func _StatsService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2ray.core.app.stats.command.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _StatsService_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2ray.com/core/app/stats/command/command.proto",
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/stats/command/command.proto", fileDescriptor_command_a956deb495595f76)
}

var fileDescriptor_command_a956deb495595f76 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4b, 0x03, 0x31,
	0x14, 0xc0, 0xbd, 0x5a, 0xeb, 0xf9, 0x14, 0x94, 0xe0, 0x50, 0xa4, 0xc3, 0x91, 0xa9, 0x8b, 0xef,
	0xe4, 0x04, 0x17, 0x27, 0xbd, 0x41, 0x10, 0x07, 0x49, 0xc1, 0xc1, 0x2d, 0xc6, 0xa7, 0x14, 0xcd,
	0x25, 0x26, 0xe9, 0x41, 0xf1, 0x1b, 0xf9, 0x29, 0x25, 0xb9, 0x1e, 0x82, 0xe0, 0xe1, 0x94, 0xf7,
	0x92, 0xdf, 0xef, 0xfd, 0x21, 0x80, 0x6d, 0xe5, 0xe4, 0x1a, 0x95, 0xd1, 0xa5, 0x32, 0x8e, 0x4a,
	0x69, 0x6d, 0xe9, 0x83, 0x0c, 0xbe, 0x54, 0x46, 0x6b, 0xd9, 0x3c, 0xf7, 0x27, 0x5a, 0x67, 0x82,
	0x61, 0xb3, 0x9e, 0x77, 0x84, 0xd2, 0x5a, 0x4c, 0x2c, 0x6e, 0x18, 0x7e, 0x09, 0x87, 0x37, 0x14,
	0x16, 0xf1, 0x4e, 0xd0, 0xc7, 0x8a, 0x7c, 0x60, 0x0c, 0xc6, 0x8d, 0xd4, 0x34, 0xcd, 0x8a, 0x6c,
	0xbe, 0x27, 0x52, 0xcc, 0x8e, 0x61, 0xc7, 0x91, 0xa7, 0x30, 0x1d, 0x15, 0xd9, 0x3c, 0x17, 0x5d,
	0xc2, 0xcf, 0x60, 0x1c, 0xcd, 0xbf, 0x8c, 0x56, 0xbe, 0xaf, 0x28, 0x19, 0xdb, 0xa2, 0x4b, 0xf8,
	0x2d, 0x1c, 0xfd, 0xb4, 0xf3, 0xd6, 0x34, 0x9e, 0xd8, 0x05, 0x8c, 0xe3, 0x4c, 0xc9, 0xde, 0xaf,
	0x38, 0x0e, 0xcd, 0x8b, 0x51, 0x15, 0x89, 0xe7, 0x39, 0x4c, 0x6a, 0xd3, 0xbc, 0x2c, 0x5f, 0xab,
	0x4f, 0x38, 0x48, 0x25, 0x17, 0xe4, 0xda, 0xa5, 0x22, 0xf6, 0x06, 0x79, 0xdf, 0x85, 0x9d, 0x0e,
	0xd7, 0xfb, 0xb5, 0xfc, 0x09, 0xfe, 0x17, 0xef, 0x86, 0xe7, 0x5b, 0xd7, 0x77, 0x50, 0x28, 0xa3,
	0x07, 0xb5, 0xfb, 0xec, 0x71, 0x77, 0x13, 0x7e, 0x8d, 0x66, 0x0f, 0x95, 0x90, 0x6b, 0xac, 0x23,
	0x79, 0x65, 0x6d, 0xda, 0xc8, 0x63, 0xdd, 0x3d, 0x3f, 0x4d, 0xd2, 0xa7, 0x9d, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x3a, 0x8a, 0xf3, 0xe6, 0x01, 0x00, 0x00,
}
