// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gitlab.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_gitlab_d4a072aee4b3acfb, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Project struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
	PathWithNamespace    string   `protobuf:"bytes,2,opt,name=PathWithNamespace" json:"PathWithNamespace,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_gitlab_d4a072aee4b3acfb, []int{1}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (dst *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(dst, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Project) GetPathWithNamespace() string {
	if m != nil {
		return m.PathWithNamespace
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Projects struct {
	Projects             []*Project `protobuf:"bytes,1,rep,name=Projects" json:"Projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Projects) Reset()         { *m = Projects{} }
func (m *Projects) String() string { return proto.CompactTextString(m) }
func (*Projects) ProtoMessage()    {}
func (*Projects) Descriptor() ([]byte, []int) {
	return fileDescriptor_gitlab_d4a072aee4b3acfb, []int{2}
}
func (m *Projects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Projects.Unmarshal(m, b)
}
func (m *Projects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Projects.Marshal(b, m, deterministic)
}
func (dst *Projects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Projects.Merge(dst, src)
}
func (m *Projects) XXX_Size() int {
	return xxx_messageInfo_Projects.Size(m)
}
func (m *Projects) XXX_DiscardUnknown() {
	xxx_messageInfo_Projects.DiscardUnknown(m)
}

var xxx_messageInfo_Projects proto.InternalMessageInfo

func (m *Projects) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ProjectPredicate struct {
	// can be an int or project path
	Id                   string   `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectPredicate) Reset()         { *m = ProjectPredicate{} }
func (m *ProjectPredicate) String() string { return proto.CompactTextString(m) }
func (*ProjectPredicate) ProtoMessage()    {}
func (*ProjectPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_gitlab_d4a072aee4b3acfb, []int{3}
}
func (m *ProjectPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectPredicate.Unmarshal(m, b)
}
func (m *ProjectPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectPredicate.Marshal(b, m, deterministic)
}
func (dst *ProjectPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectPredicate.Merge(dst, src)
}
func (m *ProjectPredicate) XXX_Size() int {
	return xxx_messageInfo_ProjectPredicate.Size(m)
}
func (m *ProjectPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectPredicate proto.InternalMessageInfo

func (m *ProjectPredicate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Environment struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=Slug" json:"Slug,omitempty"`
	ExternalURL          string   `protobuf:"bytes,3,opt,name=ExternalURL" json:"ExternalURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_gitlab_d4a072aee4b3acfb, []int{4}
}
func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (dst *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(dst, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Environment) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Environment) GetExternalURL() string {
	if m != nil {
		return m.ExternalURL
	}
	return ""
}

type Environments struct {
	Environments         []*Environment `protobuf:"bytes,1,rep,name=Environments" json:"Environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Environments) Reset()         { *m = Environments{} }
func (m *Environments) String() string { return proto.CompactTextString(m) }
func (*Environments) ProtoMessage()    {}
func (*Environments) Descriptor() ([]byte, []int) {
	return fileDescriptor_gitlab_d4a072aee4b3acfb, []int{5}
}
func (m *Environments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environments.Unmarshal(m, b)
}
func (m *Environments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environments.Marshal(b, m, deterministic)
}
func (dst *Environments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environments.Merge(dst, src)
}
func (m *Environments) XXX_Size() int {
	return xxx_messageInfo_Environments.Size(m)
}
func (m *Environments) XXX_DiscardUnknown() {
	xxx_messageInfo_Environments.DiscardUnknown(m)
}

var xxx_messageInfo_Environments proto.InternalMessageInfo

func (m *Environments) GetEnvironments() []*Environment {
	if m != nil {
		return m.Environments
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "rpc.User")
	proto.RegisterType((*Project)(nil), "rpc.Project")
	proto.RegisterType((*Projects)(nil), "rpc.Projects")
	proto.RegisterType((*ProjectPredicate)(nil), "rpc.ProjectPredicate")
	proto.RegisterType((*Environment)(nil), "rpc.Environment")
	proto.RegisterType((*Environments)(nil), "rpc.Environments")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GitlabClient is the client API for Gitlab service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GitlabClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	MyUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error)
	MyProjects(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Projects, error)
	MyEnvironments(ctx context.Context, in *ProjectPredicate, opts ...grpc.CallOption) (*Environments, error)
}

type gitlabClient struct {
	cc *grpc.ClientConn
}

func NewGitlabClient(cc *grpc.ClientConn) GitlabClient {
	return &gitlabClient{cc}
}

func (c *gitlabClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpc.Gitlab/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabClient) MyUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/rpc.Gitlab/MyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabClient) MyProjects(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Projects, error) {
	out := new(Projects)
	err := c.cc.Invoke(ctx, "/rpc.Gitlab/MyProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitlabClient) MyEnvironments(ctx context.Context, in *ProjectPredicate, opts ...grpc.CallOption) (*Environments, error) {
	out := new(Environments)
	err := c.cc.Invoke(ctx, "/rpc.Gitlab/MyEnvironments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitlabServer is the server API for Gitlab service.
type GitlabServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	MyUser(context.Context, *empty.Empty) (*User, error)
	MyProjects(context.Context, *empty.Empty) (*Projects, error)
	MyEnvironments(context.Context, *ProjectPredicate) (*Environments, error)
}

func RegisterGitlabServer(s *grpc.Server, srv GitlabServer) {
	s.RegisterService(&_Gitlab_serviceDesc, srv)
}

func _Gitlab_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Gitlab/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gitlab_MyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServer).MyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Gitlab/MyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServer).MyUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gitlab_MyProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServer).MyProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Gitlab/MyProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServer).MyProjects(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gitlab_MyEnvironments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectPredicate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitlabServer).MyEnvironments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Gitlab/MyEnvironments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitlabServer).MyEnvironments(ctx, req.(*ProjectPredicate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gitlab_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Gitlab",
	HandlerType: (*GitlabServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Gitlab_Ping_Handler,
		},
		{
			MethodName: "MyUser",
			Handler:    _Gitlab_MyUser_Handler,
		},
		{
			MethodName: "MyProjects",
			Handler:    _Gitlab_MyProjects_Handler,
		},
		{
			MethodName: "MyEnvironments",
			Handler:    _Gitlab_MyEnvironments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitlab.proto",
}

func init() { proto.RegisterFile("gitlab.proto", fileDescriptor_gitlab_d4a072aee4b3acfb) }

var fileDescriptor_gitlab_d4a072aee4b3acfb = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0x25, 0x2a, 0x56, 0x47, 0x2b, 0x3a, 0xd0, 0x12, 0xd2, 0x8b, 0xec, 0x49, 0x68, 0x89, 0xa0,
	0xd2, 0x43, 0xcf, 0x95, 0x22, 0xd4, 0x12, 0x52, 0xc4, 0x63, 0x89, 0x71, 0x1b, 0x53, 0x92, 0x6c,
	0xd8, 0x5d, 0x4b, 0xf3, 0xb3, 0xfd, 0x96, 0xb2, 0xeb, 0x36, 0xc6, 0x16, 0x4f, 0x3b, 0xf3, 0xe6,
	0xed, 0xcc, 0x9b, 0x79, 0xd0, 0x8d, 0x62, 0x99, 0x04, 0x1b, 0x37, 0xe7, 0x4c, 0x32, 0xac, 0xf3,
	0x3c, 0x74, 0x6e, 0x22, 0xc6, 0xa2, 0x84, 0x8e, 0x35, 0xb4, 0xd9, 0xbf, 0x8f, 0x69, 0x9a, 0xcb,
	0xe2, 0xc0, 0x20, 0x04, 0x1a, 0x2b, 0x41, 0x39, 0x3a, 0xd0, 0x52, 0x6f, 0x16, 0xa4, 0xd4, 0xb6,
	0x86, 0xd6, 0xa8, 0xed, 0x97, 0x39, 0x79, 0x83, 0x0b, 0x8f, 0xb3, 0x0f, 0x1a, 0x4a, 0x44, 0x68,
	0x78, 0x81, 0xdc, 0x19, 0x8a, 0x8e, 0xf1, 0x0e, 0x06, 0xea, 0x5d, 0xc7, 0x72, 0xf7, 0x12, 0xa4,
	0x54, 0xe4, 0x41, 0x48, 0xed, 0x9a, 0x26, 0xfc, 0x2f, 0xa8, 0x0e, 0x2a, 0xb1, 0xeb, 0x87, 0x0e,
	0x2a, 0x26, 0x33, 0x68, 0x99, 0x01, 0x02, 0x47, 0xc7, 0xd8, 0xb6, 0x86, 0xf5, 0x51, 0x67, 0xd2,
	0x75, 0x79, 0x1e, 0xba, 0x06, 0xf4, 0xcb, 0x2a, 0x21, 0xd0, 0x37, 0xb1, 0xc7, 0xe9, 0x36, 0x0e,
	0x03, 0x49, 0xb1, 0x07, 0xb5, 0xc5, 0xd6, 0xa8, 0xab, 0x2d, 0xb6, 0x64, 0x0d, 0x9d, 0x79, 0xf6,
	0x19, 0x73, 0x96, 0xa5, 0x34, 0x93, 0xe5, 0x70, 0xeb, 0x38, 0x5c, 0x61, 0xaf, 0xc9, 0x3e, 0x32,
	0x8a, 0x75, 0x8c, 0x43, 0xe8, 0xcc, 0xbf, 0xa4, 0x5a, 0x3f, 0x59, 0xf9, 0xcf, 0x46, 0x6b, 0x15,
	0x22, 0x8f, 0xd0, 0xad, 0x34, 0x16, 0x38, 0x3b, 0xcd, 0x8d, 0xf4, 0xbe, 0x96, 0x5e, 0x29, 0xf8,
	0x27, 0xac, 0xc9, 0xb7, 0x05, 0xcd, 0x27, 0x6d, 0x18, 0xde, 0x43, 0xc3, 0x8b, 0xb3, 0x08, 0xaf,
	0xdd, 0x83, 0x5d, 0xee, 0xaf, 0x5d, 0xee, 0x5c, 0xd9, 0xe5, 0x9c, 0xc1, 0xf1, 0x16, 0x9a, 0xcb,
	0x42, 0x5b, 0x78, 0xee, 0x67, 0x5b, 0x8b, 0xd0, 0x94, 0x29, 0xc0, 0xb2, 0x28, 0x4f, 0x7d, 0xee,
	0xc3, 0x65, 0xf5, 0xe0, 0x02, 0x1f, 0xa0, 0xb7, 0x2c, 0x4e, 0x96, 0xbd, 0xaa, 0x12, 0xca, 0xe3,
	0x3b, 0x83, 0xbf, 0xdb, 0x8a, 0x4d, 0x53, 0xb7, 0x9e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6d,
	0x7e, 0x6f, 0x96, 0x97, 0x02, 0x00, 0x00,
}