// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/tcp_proxy.proto

package envoy_api_v2_filter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [V2-API-DIFF] The route match now takes place in the FilterChainMatch table.
type TcpProxy struct {
	// The human readable prefix to use when emitting statistics for the
	// TCP proxy filter. See the statistics documentation for more information.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	// The upstream cluster to connect to.
	// TODO(htuch): look into shared WeightedCluster support with RDS.
	Cluster string `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	// [V2-API-DIFF] The idle timeout for connections managed by the TCP proxy
	// filter. The idle timeout is defined as the period in which there is no
	// active traffic. If not set, there is no idle timeout. When the idle timeout
	// is reached the connection will be closed. The distinction between
	// downstream_idle_timeout/upstream_idle_timeout provides a means to set
	// timeout based on the last byte sent on the downstream/upstream connection.
	DownstreamIdleTimeout *google_protobuf.Duration `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout" json:"downstream_idle_timeout,omitempty"`
	UpstreamIdleTimeout   *google_protobuf.Duration `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout" json:"upstream_idle_timeout,omitempty"`
}

func (m *TcpProxy) Reset()                    { *m = TcpProxy{} }
func (m *TcpProxy) String() string            { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()               {}
func (*TcpProxy) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *TcpProxy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *google_protobuf.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.api.v2.filter.TcpProxy")
}

func init() { proto.RegisterFile("api/filter/tcp_proxy.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x59, 0x15, 0xff, 0xa4, 0xb7, 0x94, 0xe2, 0xda, 0x83, 0x16, 0x4f, 0x3d, 0x25, 0x50,
	0x5f, 0xc1, 0x8b, 0x07, 0xa1, 0x96, 0xde, 0x43, 0xba, 0x99, 0x2d, 0x81, 0x74, 0x67, 0xc8, 0x4e,
	0x6a, 0xfb, 0xc4, 0xbe, 0x86, 0x34, 0x71, 0xd1, 0x83, 0xd0, 0xe3, 0x7c, 0xf3, 0xfd, 0x7e, 0x7c,
	0x62, 0x6a, 0xc9, 0xeb, 0xd6, 0x07, 0x86, 0xa8, 0xb9, 0x21, 0x43, 0x11, 0x0f, 0x47, 0x45, 0x11,
	0x19, 0xe5, 0x18, 0xba, 0x3d, 0x1e, 0x95, 0x25, 0xaf, 0xf6, 0x0b, 0x55, 0x4a, 0xd3, 0xc7, 0x2d,
	0xe2, 0x36, 0x80, 0xce, 0x95, 0x4d, 0x6a, 0xb5, 0x4b, 0xd1, 0xb2, 0xc7, 0xae, 0x40, 0xcf, 0x5f,
	0x95, 0xb8, 0x5d, 0x37, 0xb4, 0x3c, 0x79, 0xe4, 0x93, 0x18, 0xf5, 0x6c, 0xd9, 0x50, 0x84, 0xd6,
	0x1f, 0xea, 0x6a, 0x56, 0xcd, 0xef, 0x56, 0xe2, 0x14, 0x2d, 0x73, 0x22, 0x6b, 0x71, 0xd3, 0x84,
	0xd4, 0x33, 0xc4, 0xfa, 0x22, 0x3f, 0x87, 0x53, 0x7e, 0x88, 0x7b, 0x87, 0x9f, 0x5d, 0xcf, 0x11,
	0xec, 0xce, 0x78, 0x17, 0xc0, 0xb0, 0xdf, 0x01, 0x26, 0xae, 0x2f, 0x67, 0xd5, 0x7c, 0xb4, 0x78,
	0x50, 0x65, 0x89, 0x1a, 0x96, 0xa8, 0xd7, 0x9f, 0x25, 0xab, 0xc9, 0x2f, 0xf9, 0xe6, 0x02, 0xac,
	0x0b, 0x27, 0xdf, 0xc5, 0x24, 0xd1, 0x7f, 0xc2, 0xab, 0x73, 0xc2, 0xf1, 0xc0, 0xfd, 0xd1, 0x6d,
	0xae, 0x73, 0xef, 0xe5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xae, 0x41, 0x3b, 0xc5, 0x43, 0x01, 0x00,
	0x00,
}
