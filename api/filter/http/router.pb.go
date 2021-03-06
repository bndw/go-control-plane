// Code generated by protoc-gen-go.
// source: api/filter/http/router.proto
// DO NOT EDIT!

package http

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Router struct {
	DynamicStats   *google_protobuf1.BoolValue `protobuf:"bytes,1,opt,name=dynamic_stats,json=dynamicStats" json:"dynamic_stats,omitempty"`
	StartChildSpan bool                        `protobuf:"varint,2,opt,name=start_child_span,json=startChildSpan" json:"start_child_span,omitempty"`
}

func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
func (*Router) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Router) GetDynamicStats() *google_protobuf1.BoolValue {
	if m != nil {
		return m.DynamicStats
	}
	return nil
}

func (m *Router) GetStartChildSpan() bool {
	if m != nil {
		return m.StartChildSpan
	}
	return false
}

func init() {
	proto.RegisterType((*Router)(nil), "envoy.api.v2.filter.http.Router")
}

func init() { proto.RegisterFile("api/filter/http/router.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xcc, 0xcd, 0x6a, 0x85, 0x30,
	0x10, 0xc5, 0x71, 0xec, 0xe2, 0x52, 0xd2, 0x0f, 0x8a, 0x2b, 0xb9, 0x94, 0x22, 0x5d, 0xb9, 0x9a,
	0x80, 0x7d, 0x80, 0x42, 0xfb, 0x06, 0x0a, 0xdd, 0xca, 0xa8, 0x51, 0x03, 0x69, 0x66, 0x48, 0x46,
	0x8b, 0x6f, 0x5f, 0x8c, 0xed, 0xf6, 0xf0, 0x3f, 0x3f, 0xf5, 0x8c, 0x6c, 0xf5, 0x64, 0x9d, 0x98,
	0xa0, 0x17, 0x11, 0xd6, 0x81, 0x56, 0x31, 0x01, 0x38, 0x90, 0x50, 0x5e, 0x18, 0xbf, 0xd1, 0x0e,
	0xc8, 0x16, 0xb6, 0x1a, 0xce, 0x0c, 0x8e, 0xec, 0xfa, 0x32, 0x13, 0xcd, 0xce, 0xe8, 0xd4, 0xf5,
	0xeb, 0xa4, 0x7f, 0x02, 0x32, 0x9b, 0x10, 0xcf, 0xe7, 0x6b, 0x54, 0x97, 0x26, 0x49, 0xf9, 0xbb,
	0x7a, 0x18, 0x77, 0x8f, 0xdf, 0x76, 0xe8, 0xa2, 0xa0, 0xc4, 0x22, 0x2b, 0xb3, 0xea, 0xae, 0xbe,
	0xc2, 0x29, 0xc0, 0xbf, 0x00, 0x1f, 0x44, 0xee, 0x0b, 0xdd, 0x6a, 0x9a, 0xfb, 0xbf, 0x43, 0x7b,
	0xf4, 0x79, 0xa5, 0x9e, 0xa2, 0x60, 0x90, 0x6e, 0x58, 0xac, 0x1b, 0xbb, 0xc8, 0xe8, 0x8b, 0x9b,
	0x32, 0xab, 0x6e, 0x9b, 0xc7, 0xb4, 0x7f, 0x1e, 0x73, 0xcb, 0xe8, 0xfb, 0x4b, 0xb2, 0xde, 0x7e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x47, 0xfb, 0x02, 0xa8, 0xd5, 0x00, 0x00, 0x00,
}
