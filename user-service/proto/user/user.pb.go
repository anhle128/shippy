// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	User
	Request
	Response
	Token
	Error
*/
package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Company  string `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	User   *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Token  *Token   `protobuf:"bytes,4,opt,name=Token" json:"Token,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Token struct {
	Token  string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid  bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0xaa, 0xd3, 0x40,
	0x10, 0xc6, 0xdb, 0xfc, 0x6b, 0x3b, 0x45, 0xc1, 0x41, 0x71, 0xa9, 0x37, 0x25, 0x57, 0x82, 0x18,
	0xa5, 0x5e, 0x4a, 0xc1, 0x52, 0x4a, 0xef, 0xd7, 0x3f, 0xf7, 0x31, 0x19, 0x74, 0x31, 0xc9, 0xc6,
	0xdd, 0x6d, 0xc4, 0x37, 0xf4, 0x0d, 0x7c, 0x1d, 0xd9, 0x49, 0x2b, 0x07, 0x4e, 0xd2, 0xc3, 0x39,
	0x37, 0xc9, 0xce, 0x37, 0xbf, 0xdd, 0x7c, 0x99, 0x2f, 0x81, 0x67, 0xad, 0xd1, 0x4e, 0xbf, 0x39,
	0x59, 0x32, 0x7c, 0xc9, 0xb8, 0xc6, 0x27, 0xdf, 0x74, 0x56, 0xab, 0xc2, 0xe8, 0xcc, 0x9a, 0x2e,
	0xf3, 0x8d, 0xb4, 0x83, 0xe8, 0xb3, 0x25, 0x83, 0x8f, 0x21, 0x50, 0xa5, 0x98, 0xae, 0xa7, 0x2f,
	0x17, 0x32, 0x50, 0x25, 0x22, 0x44, 0x4d, 0x5e, 0x93, 0x08, 0x58, 0xe1, 0x35, 0x0a, 0x98, 0x15,
	0xba, 0x6e, 0xf3, 0xe6, 0xb7, 0x08, 0x59, 0xbe, 0x94, 0xf8, 0x14, 0x62, 0xaa, 0x73, 0x55, 0x89,
	0x88, 0xf5, 0xbe, 0xc0, 0x15, 0xcc, 0xdb, 0xdc, 0xda, 0x5f, 0xda, 0x94, 0x22, 0xe6, 0xc6, 0xff,
	0x3a, 0x5d, 0xc0, 0x4c, 0xd2, 0xcf, 0x13, 0x59, 0x97, 0xfe, 0x99, 0xc2, 0x5c, 0x92, 0x6d, 0x75,
	0x63, 0x09, 0x5f, 0x41, 0xe4, 0x7d, 0xb1, 0x93, 0xe5, 0xe6, 0x79, 0x76, 0xcb, 0x71, 0xe6, 0xed,
	0x4a, 0x86, 0xf0, 0x35, 0xc4, 0xfe, 0x6e, 0x45, 0xb0, 0x0e, 0xaf, 0xd1, 0x3d, 0x85, 0x6f, 0x21,
	0x21, 0x63, 0xb4, 0xb1, 0x22, 0x64, 0x5e, 0x0c, 0xf0, 0x07, 0x0f, 0xc8, 0x33, 0x87, 0x19, 0xc4,
	0x9f, 0xf4, 0x0f, 0x6a, 0xf8, 0xbd, 0x86, 0x37, 0x70, 0x5f, 0xf6, 0x58, 0x4a, 0x67, 0xde, 0x0f,
	0xc4, 0xf1, 0xc6, 0x7e, 0xa2, 0x7d, 0xe1, 0xd5, 0x2e, 0xaf, 0x54, 0xc9, 0x53, 0x9d, 0xcb, 0xbe,
	0xb8, 0xbf, 0xad, 0x74, 0x0b, 0x31, 0x0b, 0x3e, 0xa5, 0x42, 0x97, 0xc4, 0x4f, 0x89, 0x25, 0xaf,
	0x71, 0x0d, 0xcb, 0x92, 0x6c, 0x61, 0x54, 0xeb, 0x94, 0x6e, 0xce, 0x01, 0xde, 0x94, 0x36, 0x7f,
	0x03, 0x58, 0xfa, 0xb9, 0x7c, 0x24, 0xd3, 0xa9, 0x82, 0xf0, 0x03, 0x24, 0x7b, 0x43, 0xb9, 0x23,
	0x1c, 0x9b, 0xe0, 0xea, 0xc5, 0x40, 0xe3, 0x92, 0x59, 0x3a, 0xc1, 0x2d, 0x84, 0x47, 0x72, 0x0f,
	0xde, 0xbe, 0x87, 0xe4, 0x48, 0x6e, 0x57, 0x55, 0xb8, 0x1a, 0x04, 0xf9, 0x3b, 0xb9, 0xeb, 0x90,
	0xf7, 0x10, 0xed, 0x4e, 0xee, 0xfb, 0xb8, 0x89, 0xd1, 0xf4, 0xd2, 0x09, 0x1e, 0xe0, 0xd1, 0x17,
	0x1f, 0x46, 0xee, 0xa8, 0x0f, 0x70, 0x14, 0xbe, 0x76, 0xcc, 0xd7, 0x84, 0xff, 0xb3, 0x77, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0xbb, 0x04, 0x25, 0x80, 0x03, 0x00, 0x00,
}