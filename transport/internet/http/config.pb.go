package http

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

type Config struct {
	Host                 []string `protobuf:"bytes,1,rep,name=host" json:"host,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_93ae84938c3897ac, []int{0}
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

func (m *Config) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Config) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.http.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/http/config.proto", fileDescriptor_config_93ae84938c3897ac)
}

var fileDescriptor_config_93ae84938c3897ac = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0xcf, 0x28,
	0x29, 0x29, 0xd0, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x82, 0x69, 0x2a, 0x4a, 0xd5, 0x83, 0x6b, 0xd0, 0x83, 0x69, 0xd0, 0x03, 0x69, 0x50, 0x32,
	0xe0, 0x62, 0x73, 0x06, 0xeb, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0xc8, 0x2f, 0x2e, 0x91, 0x60, 0x54,
	0x60, 0xd6, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0x05, 0x89, 0x25, 0x19, 0x12, 0x4c, 0x0a, 0x8c,
	0x20, 0x31, 0x10, 0xdb, 0x29, 0x94, 0x4b, 0x2d, 0x39, 0x3f, 0x57, 0x8f, 0xb0, 0xd9, 0x01, 0x8c,
	0x51, 0x2c, 0x20, 0x7a, 0x15, 0x93, 0x52, 0x98, 0x51, 0x50, 0x62, 0xa5, 0x9e, 0x33, 0x48, 0x71,
	0x08, 0x5c, 0xb1, 0x27, 0x4c, 0xb1, 0x47, 0x49, 0x49, 0x41, 0x12, 0x1b, 0xd8, 0xcd, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xf2, 0x95, 0x63, 0xea, 0x00, 0x00, 0x00,
}
