// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto

/*
Package cluster is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto

It has these top-level messages:
	ContainersRequest
	ImagesRequest
	NodesRequest
	ContainersResponse
	ImagesResponse
	Node
	NodesResponse
	Container
*/
package cluster

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import stellar_services_node_v1 "github.com/ehazlett/stellar/api/services/node/v1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ContainersRequest struct {
	Filters []string `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (m *ContainersRequest) Reset()                    { *m = ContainersRequest{} }
func (m *ContainersRequest) String() string            { return proto.CompactTextString(m) }
func (*ContainersRequest) ProtoMessage()               {}
func (*ContainersRequest) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{0} }

func (m *ContainersRequest) GetFilters() []string {
	if m != nil {
		return m.Filters
	}
	return nil
}

type ImagesRequest struct {
}

func (m *ImagesRequest) Reset()                    { *m = ImagesRequest{} }
func (m *ImagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ImagesRequest) ProtoMessage()               {}
func (*ImagesRequest) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{1} }

type NodesRequest struct {
}

func (m *NodesRequest) Reset()                    { *m = NodesRequest{} }
func (m *NodesRequest) String() string            { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()               {}
func (*NodesRequest) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{2} }

type ContainersResponse struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
}

func (m *ContainersResponse) Reset()                    { *m = ContainersResponse{} }
func (m *ContainersResponse) String() string            { return proto.CompactTextString(m) }
func (*ContainersResponse) ProtoMessage()               {}
func (*ContainersResponse) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{3} }

func (m *ContainersResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type ImagesResponse struct {
	Images []*stellar_services_node_v1.Image `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
}

func (m *ImagesResponse) Reset()                    { *m = ImagesResponse{} }
func (m *ImagesResponse) String() string            { return proto.CompactTextString(m) }
func (*ImagesResponse) ProtoMessage()               {}
func (*ImagesResponse) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{4} }

func (m *ImagesResponse) GetImages() []*stellar_services_node_v1.Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Node struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{5} }

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type NodesResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *NodesResponse) Reset()                    { *m = NodesResponse{} }
func (m *NodesResponse) String() string            { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()               {}
func (*NodesResponse) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{6} }

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Container struct {
	Container *stellar_services_node_v1.Container `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Node      *Node                               `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptorCluster, []int{7} }

func (m *Container) GetContainer() *stellar_services_node_v1.Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Container) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainersRequest)(nil), "stellar.services.cluster.v1.ContainersRequest")
	proto.RegisterType((*ImagesRequest)(nil), "stellar.services.cluster.v1.ImagesRequest")
	proto.RegisterType((*NodesRequest)(nil), "stellar.services.cluster.v1.NodesRequest")
	proto.RegisterType((*ContainersResponse)(nil), "stellar.services.cluster.v1.ContainersResponse")
	proto.RegisterType((*ImagesResponse)(nil), "stellar.services.cluster.v1.ImagesResponse")
	proto.RegisterType((*Node)(nil), "stellar.services.cluster.v1.Node")
	proto.RegisterType((*NodesResponse)(nil), "stellar.services.cluster.v1.NodesResponse")
	proto.RegisterType((*Container)(nil), "stellar.services.cluster.v1.Container")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cluster service

type ClusterClient interface {
	Containers(ctx context.Context, in *ContainersRequest, opts ...grpc.CallOption) (*ContainersResponse, error)
	Images(ctx context.Context, in *ImagesRequest, opts ...grpc.CallOption) (*ImagesResponse, error)
	Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) Containers(ctx context.Context, in *ContainersRequest, opts ...grpc.CallOption) (*ContainersResponse, error) {
	out := new(ContainersResponse)
	err := grpc.Invoke(ctx, "/stellar.services.cluster.v1.Cluster/Containers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Images(ctx context.Context, in *ImagesRequest, opts ...grpc.CallOption) (*ImagesResponse, error) {
	out := new(ImagesResponse)
	err := grpc.Invoke(ctx, "/stellar.services.cluster.v1.Cluster/Images", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error) {
	out := new(NodesResponse)
	err := grpc.Invoke(ctx, "/stellar.services.cluster.v1.Cluster/Nodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cluster service

type ClusterServer interface {
	Containers(context.Context, *ContainersRequest) (*ContainersResponse, error)
	Images(context.Context, *ImagesRequest) (*ImagesResponse, error)
	Nodes(context.Context, *NodesRequest) (*NodesResponse, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_Containers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Containers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.cluster.v1.Cluster/Containers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Containers(ctx, req.(*ContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Images_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Images(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.cluster.v1.Cluster/Images",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Images(ctx, req.(*ImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.cluster.v1.Cluster/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Nodes(ctx, req.(*NodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.cluster.v1.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Containers",
			Handler:    _Cluster_Containers_Handler,
		},
		{
			MethodName: "Images",
			Handler:    _Cluster_Images_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _Cluster_Nodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto", fileDescriptorCluster)
}

var fileDescriptorCluster = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x25, 0x7d, 0xfd, 0x20, 0xb7, 0xaf, 0x7d, 0xbc, 0x59, 0x85, 0xba, 0x68, 0x8d, 0x20, 0xb5,
	0xe2, 0x84, 0x56, 0xc4, 0x45, 0x71, 0x51, 0x0b, 0x62, 0x37, 0x2e, 0xb2, 0x14, 0x11, 0xa6, 0xc9,
	0xd8, 0x06, 0xf2, 0x51, 0x33, 0xd3, 0x2c, 0xdc, 0xfb, 0xab, 0xfc, 0x73, 0x32, 0x93, 0x99, 0xa4,
	0x45, 0xac, 0xc1, 0x55, 0xee, 0xdc, 0x39, 0xe7, 0xdc, 0x7b, 0xee, 0xcd, 0xc0, 0x6c, 0x15, 0xf0,
	0xf5, 0x76, 0x89, 0xbd, 0x24, 0x72, 0xe8, 0x9a, 0xbc, 0x85, 0x94, 0x73, 0x87, 0x71, 0x1a, 0x86,
	0x24, 0x75, 0xc8, 0x26, 0x70, 0x18, 0x4d, 0xb3, 0xc0, 0xa3, 0xcc, 0xf1, 0xc2, 0x2d, 0xe3, 0x34,
	0x75, 0xb2, 0xb1, 0x0e, 0xf1, 0x26, 0x4d, 0x78, 0x82, 0x8e, 0x14, 0x1c, 0x6b, 0x28, 0xd6, 0xf7,
	0xd9, 0xb8, 0xd7, 0xdf, 0xd3, 0x88, 0x13, 0x9f, 0x0a, 0x01, 0xf1, 0xcd, 0xd9, 0xf6, 0x05, 0xfc,
	0x9f, 0x27, 0x31, 0x27, 0x41, 0x4c, 0x53, 0xe6, 0xd2, 0xd7, 0x2d, 0x65, 0x1c, 0x59, 0xd0, 0x7a,
	0x09, 0x42, 0x4e, 0x53, 0x66, 0x19, 0x83, 0x3f, 0x43, 0xd3, 0xd5, 0x47, 0xfb, 0x1f, 0x74, 0x16,
	0x11, 0x59, 0x51, 0x0d, 0xb5, 0xbb, 0xf0, 0xf7, 0x21, 0xf1, 0xcb, 0xf3, 0x13, 0xa0, 0x5d, 0x3d,
	0xb6, 0x49, 0x62, 0x46, 0xd1, 0x1d, 0x80, 0x57, 0x64, 0xa5, 0x66, 0x7b, 0x72, 0x8a, 0x0f, 0x34,
	0x8e, 0x0b, 0x11, 0x77, 0x87, 0x69, 0x2f, 0xa0, 0xab, 0xcb, 0x2b, 0xe5, 0x6b, 0x68, 0x06, 0x32,
	0xa3, 0x54, 0xfb, 0x5f, 0x55, 0xa5, 0xdb, 0x6c, 0x8c, 0x25, 0xd3, 0x55, 0x70, 0x1b, 0x43, 0x5d,
	0x34, 0x8e, 0x10, 0xd4, 0x63, 0x12, 0x51, 0xcb, 0x18, 0x18, 0x43, 0xd3, 0x95, 0xb1, 0xc8, 0x11,
	0xdf, 0x4f, 0xad, 0x5a, 0x9e, 0x13, 0xb1, 0x7d, 0x0f, 0x1d, 0x65, 0xb4, 0xa8, 0xdc, 0x10, 0xca,
	0xba, 0xf0, 0xf1, 0x41, 0x3b, 0x82, 0xea, 0xe6, 0x78, 0xfb, 0xdd, 0x00, 0xb3, 0xb0, 0x87, 0x66,
	0x60, 0x16, 0x06, 0x65, 0x13, 0xed, 0xc9, 0xc9, 0xf7, 0x1e, 0xca, 0xb1, 0x94, 0x2c, 0x74, 0x05,
	0x75, 0x71, 0x2f, 0xdb, 0xad, 0xd4, 0x88, 0x84, 0x4f, 0x3e, 0x6a, 0xd0, 0x9a, 0xe7, 0x37, 0x28,
	0x02, 0x28, 0xd7, 0x86, 0x70, 0xb5, 0xd5, 0xe8, 0xa5, 0xf7, 0x9c, 0xca, 0x78, 0x35, 0x3b, 0x02,
	0xcd, 0x7c, 0x8f, 0x68, 0x74, 0x90, 0xba, 0xf7, 0xaf, 0xf5, 0xce, 0x2b, 0x61, 0x55, 0x89, 0x67,
	0x68, 0xc8, 0x7d, 0xa1, 0xb3, 0x1f, 0xe7, 0x51, 0x14, 0x18, 0x55, 0x81, 0xe6, 0xfa, 0xb7, 0x37,
	0x8f, 0xd3, 0x5f, 0xbc, 0xdd, 0xa9, 0x0a, 0x97, 0x4d, 0xf9, 0xfc, 0x2e, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x62, 0xe7, 0x18, 0x35, 0x01, 0x04, 0x00, 0x00,
}
