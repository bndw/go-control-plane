// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/mongo_proxy.proto

package envoy_api_v2_filter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MongoProxy struct {
	// The human readable prefix to use when emitting statistics for the
	// MongoDB proxy filter. See the statistics documentation for more information.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated by runtime.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog" json:"access_log,omitempty"`
}

func (m *MongoProxy) Reset()                    { *m = MongoProxy{} }
func (m *MongoProxy) String() string            { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()               {}
func (*MongoProxy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.api.v2.filter.MongoProxy")
}

func init() { proto.RegisterFile("api/filter/mongo_proxy.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xcf, 0xcd, 0xcf, 0x4b, 0xcf, 0x8f, 0x2f, 0x28, 0xca,
	0xaf, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4,
	0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0xd2, 0x83, 0x28, 0x53, 0xf2, 0xe1, 0xe2, 0xf2, 0x05, 0xa9,
	0x0c, 0x00, 0x29, 0x14, 0x92, 0xe7, 0xe2, 0x2e, 0x2e, 0x49, 0x2c, 0x89, 0x2f, 0x28, 0x4a, 0x4d,
	0xcb, 0xac, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x02, 0x09, 0x05, 0x80, 0x45, 0x84,
	0x64, 0xb9, 0xb8, 0x12, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0xe3, 0x73, 0xf2, 0xd3, 0x25, 0x98, 0xc0,
	0xf2, 0x9c, 0x10, 0x11, 0x9f, 0xfc, 0xf4, 0x24, 0x36, 0xb0, 0x4d, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x49, 0x22, 0x2e, 0x21, 0x89, 0x00, 0x00, 0x00,
}
