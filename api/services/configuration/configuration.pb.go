// Code generated by protoc-gen-go.
// source: github.com/ehazlett/interlock/api/services/configuration/configuration.proto
// DO NOT EDIT!

/*
Package configuration is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/interlock/api/services/configuration/configuration.proto

It has these top-level messages:
	Config
	Backend
	ContextRoot
	ConfigRequest
	ConfigResponse
*/
package configuration

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import interlock_v1_types "github.com/ehazlett/interlock/api/types"

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

type Config struct {
	Version  string     `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Backends []*Backend `protobuf:"bytes,2,rep,name=backends" json:"backends,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Config) GetBackends() []*Backend {
	if m != nil {
		return m.Backends
	}
	return nil
}

type Backend struct {
	Name               string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Hosts              []string                `protobuf:"bytes,2,rep,name=hosts" json:"hosts,omitempty"`
	ContextRoots       map[string]*ContextRoot `protobuf:"bytes,3,rep,name=context_roots,json=contextRoots" json:"context_roots,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Port               uint32                  `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	Ssl                bool                    `protobuf:"varint,5,opt,name=ssl" json:"ssl,omitempty"`
	SslPort            uint32                  `protobuf:"varint,6,opt,name=ssl_port,json=sslPort" json:"ssl_port,omitempty"`
	SslCert            string                  `protobuf:"bytes,7,opt,name=ssl_cert,json=sslCert" json:"ssl_cert,omitempty"`
	SslCertKey         string                  `protobuf:"bytes,8,opt,name=ssl_cert_key,json=sslCertKey" json:"ssl_cert_key,omitempty"`
	SslOnly            bool                    `protobuf:"varint,9,opt,name=ssl_only,json=sslOnly" json:"ssl_only,omitempty"`
	SslBackend         bool                    `protobuf:"varint,10,opt,name=ssl_backend,json=sslBackend" json:"ssl_backend,omitempty"`
	WebsocketEndpoints []string                `protobuf:"bytes,11,rep,name=websocket_endpoints,json=websocketEndpoints" json:"websocket_endpoints,omitempty"`
	IpHash             bool                    `protobuf:"varint,12,opt,name=ip_hash,json=ipHash" json:"ip_hash,omitempty"`
	Targets            []string                `protobuf:"bytes,13,rep,name=targets" json:"targets,omitempty"`
}

func (m *Backend) Reset()                    { *m = Backend{} }
func (m *Backend) String() string            { return proto.CompactTextString(m) }
func (*Backend) ProtoMessage()               {}
func (*Backend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Backend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Backend) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Backend) GetContextRoots() map[string]*ContextRoot {
	if m != nil {
		return m.ContextRoots
	}
	return nil
}

func (m *Backend) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Backend) GetSsl() bool {
	if m != nil {
		return m.Ssl
	}
	return false
}

func (m *Backend) GetSslPort() uint32 {
	if m != nil {
		return m.SslPort
	}
	return 0
}

func (m *Backend) GetSslCert() string {
	if m != nil {
		return m.SslCert
	}
	return ""
}

func (m *Backend) GetSslCertKey() string {
	if m != nil {
		return m.SslCertKey
	}
	return ""
}

func (m *Backend) GetSslOnly() bool {
	if m != nil {
		return m.SslOnly
	}
	return false
}

func (m *Backend) GetSslBackend() bool {
	if m != nil {
		return m.SslBackend
	}
	return false
}

func (m *Backend) GetWebsocketEndpoints() []string {
	if m != nil {
		return m.WebsocketEndpoints
	}
	return nil
}

func (m *Backend) GetIpHash() bool {
	if m != nil {
		return m.IpHash
	}
	return false
}

func (m *Backend) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ContextRoot struct {
	Name    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path    string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Rewrite bool     `protobuf:"varint,3,opt,name=rewrite" json:"rewrite,omitempty"`
	Targets []string `protobuf:"bytes,4,rep,name=targets" json:"targets,omitempty"`
}

func (m *ContextRoot) Reset()                    { *m = ContextRoot{} }
func (m *ContextRoot) String() string            { return proto.CompactTextString(m) }
func (*ContextRoot) ProtoMessage()               {}
func (*ContextRoot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContextRoot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContextRoot) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ContextRoot) GetRewrite() bool {
	if m != nil {
		return m.Rewrite
	}
	return false
}

func (m *ContextRoot) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ConfigRequest struct {
	ServiceCluster string `protobuf:"bytes,1,opt,name=service_cluster,json=serviceCluster" json:"service_cluster,omitempty"`
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ConfigRequest) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

type ConfigResponse struct {
	Config       *Config                          `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	PluginConfig *interlock_v1_types.PluginConfig `protobuf:"bytes,2,opt,name=plugin_config,json=pluginConfig" json:"plugin_config,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConfigResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigResponse) GetPluginConfig() *interlock_v1_types.PluginConfig {
	if m != nil {
		return m.PluginConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "interlock.v1.configuration.Config")
	proto.RegisterType((*Backend)(nil), "interlock.v1.configuration.Backend")
	proto.RegisterType((*ContextRoot)(nil), "interlock.v1.configuration.ContextRoot")
	proto.RegisterType((*ConfigRequest)(nil), "interlock.v1.configuration.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "interlock.v1.configuration.ConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Configuration service

type ConfigurationClient interface {
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type configurationClient struct {
	cc *grpc.ClientConn
}

func NewConfigurationClient(cc *grpc.ClientConn) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/interlock.v1.configuration.Configuration/Config", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Configuration service

type ConfigurationServer interface {
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interlock.v1.configuration.Configuration/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interlock.v1.configuration.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Config",
			Handler:    _Configuration_Config_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/interlock/api/services/configuration/configuration.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/interlock/api/services/configuration/configuration.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xb9, 0x49, 0xf3, 0xb1, 0x49, 0xfa, 0x83, 0x05, 0x89, 0x25, 0x17, 0x2c, 0x73, 0x68,
	0xe0, 0x60, 0x8b, 0x20, 0xa4, 0xaa, 0x08, 0x21, 0x11, 0x55, 0x42, 0x02, 0x89, 0xca, 0x47, 0x2e,
	0x96, 0xe3, 0x0e, 0xf1, 0x2a, 0xee, 0xae, 0xd9, 0xd9, 0xa4, 0x98, 0x7f, 0x83, 0x03, 0xff, 0x2e,
	0xda, 0x0f, 0x47, 0x89, 0xf8, 0x68, 0x6f, 0x33, 0xfb, 0xe6, 0xcd, 0xdb, 0x79, 0xeb, 0x31, 0xf9,
	0xb8, 0xe2, 0xba, 0xdc, 0x2c, 0xe3, 0x42, 0x5e, 0x27, 0x50, 0xe6, 0xdf, 0x2b, 0xd0, 0x3a, 0xe1,
	0x42, 0x83, 0xaa, 0x64, 0xb1, 0x4e, 0xf2, 0x9a, 0x27, 0x08, 0x6a, 0xcb, 0x0b, 0xc0, 0xa4, 0x90,
	0xe2, 0x0b, 0x5f, 0x6d, 0x54, 0xae, 0xb9, 0x14, 0x87, 0x59, 0x5c, 0x2b, 0xa9, 0x25, 0x9d, 0xee,
	0x98, 0xf1, 0xf6, 0x45, 0x7c, 0x50, 0x31, 0x7d, 0x7d, 0xbb, 0x92, 0x6e, 0x6a, 0xc0, 0xa4, 0xae,
	0x36, 0x2b, 0x2e, 0x32, 0xc7, 0x76, 0x8d, 0xa3, 0x82, 0xf4, 0x16, 0x36, 0xa7, 0x8c, 0xf4, 0xb7,
	0xa0, 0x90, 0x4b, 0xc1, 0x82, 0x30, 0x98, 0x0d, 0xd3, 0x36, 0xa5, 0x6f, 0xc9, 0x60, 0x99, 0x17,
	0x6b, 0x10, 0x57, 0xc8, 0x8e, 0xc2, 0xce, 0x6c, 0x34, 0x7f, 0x1a, 0xff, 0xfd, 0x3e, 0xf1, 0x3b,
	0x57, 0x9b, 0xee, 0x48, 0xd1, 0xcf, 0x2e, 0xe9, 0xfb, 0x53, 0x4a, 0x49, 0x57, 0xe4, 0xd7, 0xe0,
	0x35, 0x6c, 0x4c, 0x1f, 0x92, 0xe3, 0x52, 0xa2, 0x76, 0xdd, 0x87, 0xa9, 0x4b, 0xe8, 0x67, 0x32,
	0x29, 0xa4, 0xd0, 0xf0, 0x4d, 0x67, 0x4a, 0x4a, 0x8d, 0xac, 0x63, 0xb5, 0x5f, 0xdd, 0x41, 0x3b,
	0x5e, 0x38, 0x62, 0x6a, 0x78, 0x17, 0x42, 0xab, 0x26, 0x1d, 0x17, 0x7b, 0x47, 0xe6, 0x16, 0xb5,
	0x54, 0x9a, 0x75, 0xc3, 0x60, 0x36, 0x49, 0x6d, 0x4c, 0xef, 0x91, 0x0e, 0x62, 0xc5, 0x8e, 0xc3,
	0x60, 0x36, 0x48, 0x4d, 0x48, 0x1f, 0x93, 0x01, 0x62, 0x95, 0xd9, 0xca, 0x9e, 0xad, 0xec, 0x23,
	0x56, 0x97, 0xa6, 0xd8, 0x43, 0x05, 0x28, 0xcd, 0xfa, 0xce, 0x2e, 0xc4, 0x6a, 0x01, 0x4a, 0xd3,
	0x90, 0x8c, 0x5b, 0x28, 0x5b, 0x43, 0xc3, 0x06, 0x16, 0x26, 0x1e, 0xfe, 0x00, 0x4d, 0x4b, 0x96,
	0xa2, 0x6a, 0xd8, 0xd0, 0xca, 0x19, 0xf2, 0x27, 0x51, 0x35, 0xf4, 0x09, 0x19, 0x19, 0xc8, 0x5b,
	0xc7, 0x88, 0x45, 0x0d, 0xb7, 0xf5, 0x2f, 0x21, 0x0f, 0x6e, 0x60, 0x89, 0xb2, 0x58, 0x83, 0xce,
	0x40, 0x5c, 0xd5, 0x92, 0x0b, 0x8d, 0x6c, 0x64, 0x9d, 0xa3, 0x3b, 0xe8, 0xa2, 0x45, 0xe8, 0x23,
	0xd2, 0xe7, 0x75, 0x56, 0xe6, 0x58, 0xb2, 0xb1, 0xed, 0xd6, 0xe3, 0xf5, 0xfb, 0x1c, 0x4b, 0xf3,
	0xe0, 0x3a, 0x57, 0x2b, 0xd0, 0xc8, 0x26, 0x96, 0xdd, 0xa6, 0xd3, 0x92, 0xdc, 0xff, 0xcd, 0x40,
	0x63, 0x8f, 0x99, 0xc6, 0xbd, 0x9b, 0x09, 0xe9, 0x1b, 0x72, 0xbc, 0xcd, 0xab, 0x0d, 0xb0, 0xa3,
	0x30, 0x98, 0x8d, 0xe6, 0xa7, 0xff, 0x7a, 0x98, 0xbd, 0x7e, 0xa9, 0x63, 0x9d, 0x1f, 0x9d, 0x05,
	0x11, 0x27, 0xa3, 0x3d, 0xe4, 0x8f, 0x1f, 0x87, 0x79, 0xaa, 0x5c, 0x97, 0x56, 0x64, 0x98, 0xda,
	0xd8, 0x5c, 0x5d, 0xc1, 0x8d, 0xe2, 0x1a, 0x58, 0xc7, 0xf9, 0xe7, 0xd3, 0xfd, 0xa1, 0xba, 0x07,
	0x43, 0x45, 0x67, 0x64, 0xe2, 0xbe, 0xf4, 0x14, 0xbe, 0x6e, 0x00, 0x35, 0x3d, 0x25, 0xff, 0xfb,
	0x0d, 0xcc, 0x8a, 0x6a, 0x83, 0x1a, 0x94, 0xd7, 0x3d, 0xf1, 0xc7, 0x0b, 0x77, 0x1a, 0xfd, 0x08,
	0xc8, 0x49, 0x4b, 0xc5, 0x5a, 0x0a, 0x04, 0x7a, 0x4e, 0x7a, 0x6e, 0x3e, 0x4b, 0x19, 0xcd, 0xa3,
	0x5b, 0x66, 0x37, 0x5c, 0xcf, 0xa0, 0x17, 0x64, 0x72, 0xb0, 0x89, 0xde, 0xbe, 0xf0, 0xb0, 0x85,
	0x5d, 0xd9, 0xf8, 0xd2, 0x16, 0xfa, 0x06, 0xe3, 0x7a, 0x2f, 0x9b, 0xab, 0x76, 0x1e, 0x2f, 0x43,
	0xf3, 0xdd, 0x2a, 0x3f, 0xbb, 0xc3, 0x6d, 0x9c, 0x09, 0xd3, 0xe7, 0x77, 0x29, 0x75, 0x43, 0x47,
	0xff, 0x2d, 0x7b, 0xf6, 0xa7, 0xf1, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x63, 0xba,
	0xb8, 0xdd, 0x04, 0x00, 0x00,
}