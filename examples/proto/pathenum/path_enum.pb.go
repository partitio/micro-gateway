// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/pathenum/path_enum.proto

package pathenum // import "github.com/partitio/micro-gateway/examples/proto/pathenum"

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

type PathEnum int32

const (
	PathEnum_ABC PathEnum = 0
	PathEnum_DEF PathEnum = 1
)

var PathEnum_name = map[int32]string{
	0: "ABC",
	1: "DEF",
}
var PathEnum_value = map[string]int32{
	"ABC": 0,
	"DEF": 1,
}

func (x PathEnum) String() string {
	return proto.EnumName(PathEnum_name, int32(x))
}
func (PathEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_path_enum_8ae4c6b975612f7a, []int{0}
}

type MessagePathEnum_NestedPathEnum int32

const (
	MessagePathEnum_GHI MessagePathEnum_NestedPathEnum = 0
	MessagePathEnum_JKL MessagePathEnum_NestedPathEnum = 1
)

var MessagePathEnum_NestedPathEnum_name = map[int32]string{
	0: "GHI",
	1: "JKL",
}
var MessagePathEnum_NestedPathEnum_value = map[string]int32{
	"GHI": 0,
	"JKL": 1,
}

func (x MessagePathEnum_NestedPathEnum) String() string {
	return proto.EnumName(MessagePathEnum_NestedPathEnum_name, int32(x))
}
func (MessagePathEnum_NestedPathEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_path_enum_8ae4c6b975612f7a, []int{0, 0}
}

type MessagePathEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagePathEnum) Reset()         { *m = MessagePathEnum{} }
func (m *MessagePathEnum) String() string { return proto.CompactTextString(m) }
func (*MessagePathEnum) ProtoMessage()    {}
func (*MessagePathEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_path_enum_8ae4c6b975612f7a, []int{0}
}
func (m *MessagePathEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagePathEnum.Unmarshal(m, b)
}
func (m *MessagePathEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagePathEnum.Marshal(b, m, deterministic)
}
func (dst *MessagePathEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagePathEnum.Merge(dst, src)
}
func (m *MessagePathEnum) XXX_Size() int {
	return xxx_messageInfo_MessagePathEnum.Size(m)
}
func (m *MessagePathEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagePathEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessagePathEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MessagePathEnum)(nil), "grpc.gateway.examples.pathenum.MessagePathEnum")
	proto.RegisterEnum("grpc.gateway.examples.pathenum.PathEnum", PathEnum_name, PathEnum_value)
	proto.RegisterEnum("grpc.gateway.examples.pathenum.MessagePathEnum_NestedPathEnum", MessagePathEnum_NestedPathEnum_name, MessagePathEnum_NestedPathEnum_value)
}

func init() {
	proto.RegisterFile("examples/proto/pathenum/path_enum.proto", fileDescriptor_path_enum_8ae4c6b975612f7a)
}

var fileDescriptor_path_enum_8ae4c6b975612f7a = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x48, 0x2c, 0xc9, 0x48,
	0xcd, 0x2b, 0xcd, 0x05, 0x33, 0xe2, 0x41, 0x2c, 0x3d, 0xb0, 0x84, 0x90, 0x5c, 0x7a, 0x51, 0x41,
	0xb2, 0x5e, 0x7a, 0x62, 0x49, 0x6a, 0x79, 0x62, 0xa5, 0x1e, 0x4c, 0x97, 0x1e, 0x4c, 0xbd, 0x92,
	0x29, 0x17, 0xbf, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0x6a, 0x40, 0x62, 0x49, 0x86, 0x2b, 0x48,
	0x48, 0x89, 0x8b, 0xcf, 0x2f, 0xb5, 0xb8, 0x24, 0x35, 0x05, 0x26, 0x22, 0xc4, 0xce, 0xc5, 0xec,
	0xee, 0xe1, 0x29, 0xc0, 0x00, 0x62, 0x78, 0x79, 0xfb, 0x08, 0x30, 0x6a, 0xc9, 0x70, 0x71, 0x20,
	0xcb, 0x3a, 0x3a, 0x39, 0x43, 0x64, 0x5d, 0x5c, 0xdd, 0x04, 0x18, 0x9d, 0xac, 0xa3, 0x2c, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x41, 0x4e, 0x2a, 0x2a, 0xc9, 0x2c, 0xc9, 0xcc,
	0xd7, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x85, 0xba, 0x45, 0x1f, 0x87, 0x0f, 0x92, 0xd8, 0xc0,
	0x7c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xc0, 0x6b, 0x75, 0xe3, 0x00, 0x00, 0x00,
}
