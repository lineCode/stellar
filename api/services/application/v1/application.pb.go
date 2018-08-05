// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/application/v1/application.proto

/*
Package application is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/application/v1/application.proto

It has these top-level messages:
	Process
	Mount
	Service
	CreateRequest
	DeleteRequest
	ListRequest
	App
	ListResponse
	GetRequest
	GetResponse
*/
package application

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

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

type Process struct {
	Uid  uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid  uint32   `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Args []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	Env  []string `protobuf:"bytes,4,rep,name=env" json:"env,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{0} }

func (m *Process) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Process) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *Process) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Process) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

type Mount struct {
	Type        string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Source      string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Destination string   `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Options     []string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (m *Mount) String() string            { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{1} }

func (m *Mount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Mount) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Mount) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

type Service struct {
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image       string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Runtime     string   `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Snapshotter string   `protobuf:"bytes,4,opt,name=snapshotter,proto3" json:"snapshotter,omitempty"`
	Process     *Process `protobuf:"bytes,5,opt,name=process" json:"process,omitempty"`
	Labels      []string `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty"`
	Network     bool     `protobuf:"varint,7,opt,name=network,proto3" json:"network,omitempty"`
	Mounts      []*Mount `protobuf:"bytes,8,rep,name=mounts" json:"mounts,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{2} }

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Service) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *Service) GetSnapshotter() string {
	if m != nil {
		return m.Snapshotter
	}
	return ""
}

func (m *Service) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *Service) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Service) GetNetwork() bool {
	if m != nil {
		return m.Network
	}
	return false
}

func (m *Service) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

type CreateRequest struct {
	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels   []string   `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	Services []*Service `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{3} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateRequest) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type DeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{4} }

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{5} }

type App struct {
	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Services []*Service `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{6} }

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListResponse struct {
	Applications []*App `protobuf:"bytes,1,rep,name=applications" json:"applications,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{7} }

func (m *ListResponse) GetApplications() []*App {
	if m != nil {
		return m.Applications
	}
	return nil
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{8} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Application *App `protobuf:"bytes,1,opt,name=application" json:"application,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{9} }

func (m *GetResponse) GetApplication() *App {
	if m != nil {
		return m.Application
	}
	return nil
}

func init() {
	proto.RegisterType((*Process)(nil), "stellar.services.application.v1.Process")
	proto.RegisterType((*Mount)(nil), "stellar.services.application.v1.Mount")
	proto.RegisterType((*Service)(nil), "stellar.services.application.v1.Service")
	proto.RegisterType((*CreateRequest)(nil), "stellar.services.application.v1.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "stellar.services.application.v1.DeleteRequest")
	proto.RegisterType((*ListRequest)(nil), "stellar.services.application.v1.ListRequest")
	proto.RegisterType((*App)(nil), "stellar.services.application.v1.App")
	proto.RegisterType((*ListResponse)(nil), "stellar.services.application.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "stellar.services.application.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "stellar.services.application.v1.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Application service

type ApplicationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type applicationClient struct {
	cc *grpc.ClientConn
}

func NewApplicationClient(cc *grpc.ClientConn) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationServer interface {
	Create(context.Context, *CreateRequest) (*google_protobuf.Empty, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterApplicationServer(s *grpc.Server, srv ApplicationServer) {
	s.RegisterService(&_Application_serviceDesc, srv)
}

func _Application_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Application_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.application.v1.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Application_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Application_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/application/v1/application.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/application/v1/application.proto", fileDescriptorApplication)
}

var fileDescriptorApplication = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6b, 0xd4, 0x40,
	0x10, 0x27, 0x97, 0xfb, 0x68, 0x27, 0x3d, 0x90, 0x45, 0xca, 0x72, 0x3e, 0x18, 0xa2, 0x48, 0xc0,
	0x9a, 0xd0, 0xfa, 0x28, 0x08, 0xfd, 0xb2, 0x0a, 0x0a, 0x25, 0x45, 0x10, 0x5f, 0x34, 0x97, 0x1b,
	0x73, 0x8b, 0x49, 0x76, 0xcd, 0x6e, 0x4e, 0xea, 0x9b, 0x7f, 0x8c, 0xff, 0xa3, 0x8f, 0xb2, 0x9b,
	0xa4, 0x97, 0x83, 0xda, 0x1c, 0x7d, 0x9b, 0x99, 0xcc, 0xef, 0x63, 0x67, 0xc2, 0xc0, 0xbb, 0x94,
	0xa9, 0x65, 0x35, 0x0f, 0x12, 0x9e, 0x87, 0xb8, 0x8c, 0x7f, 0x65, 0xa8, 0x54, 0x28, 0x15, 0x66,
	0x59, 0x5c, 0x86, 0xb1, 0x60, 0xa1, 0xc4, 0x72, 0xc5, 0x12, 0x94, 0x61, 0x2c, 0x44, 0xc6, 0x92,
	0x58, 0x31, 0x5e, 0x84, 0xab, 0xc3, 0x6e, 0x1a, 0x88, 0x92, 0x2b, 0x4e, 0x1e, 0x37, 0xb0, 0xa0,
	0x85, 0x04, 0xdd, 0x9e, 0xd5, 0xe1, 0xec, 0x51, 0xca, 0x79, 0x9a, 0x61, 0x68, 0xda, 0xe7, 0xd5,
	0xb7, 0x10, 0x73, 0xa1, 0xae, 0x6b, 0xb4, 0x77, 0x05, 0x93, 0xcb, 0x92, 0x27, 0x28, 0x25, 0x79,
	0x00, 0x76, 0xc5, 0x16, 0xd4, 0x72, 0x2d, 0x7f, 0x1a, 0xe9, 0x50, 0x57, 0x52, 0xb6, 0xa0, 0x83,
	0xba, 0x92, 0xb2, 0x05, 0x21, 0x30, 0x8c, 0xcb, 0x54, 0x52, 0xdb, 0xb5, 0xfd, 0xdd, 0xc8, 0xc4,
	0xba, 0x0b, 0x8b, 0x15, 0x1d, 0x9a, 0x92, 0x0e, 0x3d, 0x0e, 0xa3, 0x0f, 0xbc, 0x2a, 0x94, 0x6e,
	0x57, 0xd7, 0x02, 0x0d, 0xe7, 0x6e, 0x64, 0x62, 0xb2, 0x0f, 0x63, 0xc9, 0xab, 0x32, 0x41, 0xc3,
	0xbb, 0x1b, 0x35, 0x19, 0x71, 0xc1, 0x59, 0xa0, 0x54, 0xac, 0x30, 0xc6, 0xe9, 0xd0, 0x7c, 0xec,
	0x96, 0x08, 0x85, 0x09, 0x17, 0x3a, 0x92, 0x74, 0x64, 0xc4, 0xda, 0xd4, 0xfb, 0x33, 0x80, 0xc9,
	0x55, 0xfd, 0x7c, 0xad, 0x59, 0xc4, 0xf9, 0x8d, 0xa6, 0x8e, 0xc9, 0x43, 0x18, 0xb1, 0x3c, 0x4e,
	0x5b, 0xc9, 0x3a, 0xd1, 0x7c, 0x65, 0x55, 0x28, 0x96, 0x23, 0xb5, 0x4d, 0xbd, 0x4d, 0xb5, 0x17,
	0x59, 0xc4, 0x42, 0x2e, 0xb9, 0x52, 0x58, 0xb6, 0x5e, 0x3a, 0x25, 0x72, 0x02, 0x13, 0x51, 0xcf,
	0x8d, 0x8e, 0x5c, 0xcb, 0x77, 0x8e, 0xfc, 0xa0, 0x67, 0x0f, 0x41, 0x33, 0xe7, 0xa8, 0x05, 0xea,
	0x49, 0x64, 0xf1, 0x1c, 0x33, 0x49, 0xc7, 0xe6, 0x39, 0x4d, 0xa6, 0x7d, 0x15, 0xa8, 0x7e, 0xf2,
	0xf2, 0x3b, 0x9d, 0xb8, 0x96, 0xbf, 0x13, 0xb5, 0x29, 0x79, 0x0d, 0xe3, 0x5c, 0x0f, 0x56, 0xd2,
	0x1d, 0xd7, 0xf6, 0x9d, 0xa3, 0x67, 0xbd, 0xa2, 0x66, 0x0f, 0x51, 0x83, 0xf2, 0x7e, 0x5b, 0x30,
	0x3d, 0x2d, 0x31, 0x56, 0x18, 0xe1, 0x8f, 0x0a, 0xa5, 0xba, 0x75, 0x5a, 0x6b, 0x5f, 0x83, 0x0d,
	0x5f, 0x67, 0xb0, 0xd3, 0xca, 0x98, 0x1f, 0x60, 0x9b, 0x47, 0x37, 0x5b, 0x89, 0x6e, 0x90, 0xde,
	0x13, 0x98, 0x9e, 0x61, 0x86, 0x77, 0x5a, 0xf0, 0xa6, 0xe0, 0xbc, 0x67, 0x52, 0x35, 0x2d, 0xde,
	0x17, 0xb0, 0x8f, 0x85, 0xb8, 0xd5, 0x6c, 0xd7, 0xd4, 0xe0, 0xde, 0xa6, 0x3e, 0xc1, 0x5e, 0xad,
	0x27, 0x05, 0x2f, 0x24, 0x92, 0xb7, 0xb0, 0xd7, 0xc1, 0x48, 0x6a, 0x19, 0xe6, 0xa7, 0xbd, 0xcc,
	0xc7, 0x42, 0x44, 0x1b, 0x48, 0xcf, 0x05, 0xb8, 0x40, 0x75, 0xd7, 0x5b, 0x3f, 0x82, 0x63, 0x3a,
	0x1a, 0xe9, 0x37, 0xe0, 0x74, 0x08, 0x4c, 0xe7, 0xb6, 0xca, 0x5d, 0xe0, 0xd1, 0xdf, 0x01, 0x38,
	0xc7, 0xeb, 0x9c, 0x5c, 0xc2, 0xb8, 0x5e, 0x3d, 0x09, 0x7a, 0xc9, 0x36, 0xfe, 0x91, 0xd9, 0x7e,
	0x50, 0x5f, 0x90, 0xa0, 0xbd, 0x20, 0xc1, 0xb9, 0xbe, 0x20, 0x9a, 0xb1, 0xde, 0xe4, 0x16, 0x8c,
	0x1b, 0x2b, 0xff, 0x2f, 0x63, 0x02, 0x43, 0xbd, 0x06, 0x72, 0xd0, 0xcb, 0xd7, 0xf9, 0x3b, 0x66,
	0x2f, 0xb6, 0xec, 0x6e, 0x06, 0xfc, 0x15, 0xec, 0x0b, 0x54, 0xe4, 0x79, 0x2f, 0x6a, 0xbd, 0xb7,
	0xd9, 0xc1, 0x76, 0xcd, 0xb5, 0xc2, 0xc9, 0xf9, 0xe7, 0xd3, 0x7b, 0xde, 0xf7, 0x57, 0x9d, 0x74,
	0x3e, 0x36, 0xd3, 0x79, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x54, 0x38, 0x5d, 0xd6, 0x2d, 0x06,
	0x00, 0x00,
}
