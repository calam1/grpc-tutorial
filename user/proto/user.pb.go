// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{1}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{2}
}
func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (dst *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(dst, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserRequest) Reset()         { *m = ReadUserRequest{} }
func (m *ReadUserRequest) String() string { return proto.CompactTextString(m) }
func (*ReadUserRequest) ProtoMessage()    {}
func (*ReadUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{3}
}
func (m *ReadUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserRequest.Unmarshal(m, b)
}
func (m *ReadUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserRequest.Marshal(b, m, deterministic)
}
func (dst *ReadUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserRequest.Merge(dst, src)
}
func (m *ReadUserRequest) XXX_Size() int {
	return xxx_messageInfo_ReadUserRequest.Size(m)
}
func (m *ReadUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserRequest proto.InternalMessageInfo

func (m *ReadUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadUserResponse) Reset()         { *m = ReadUserResponse{} }
func (m *ReadUserResponse) String() string { return proto.CompactTextString(m) }
func (*ReadUserResponse) ProtoMessage()    {}
func (*ReadUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{4}
}
func (m *ReadUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadUserResponse.Unmarshal(m, b)
}
func (m *ReadUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadUserResponse.Marshal(b, m, deterministic)
}
func (dst *ReadUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadUserResponse.Merge(dst, src)
}
func (m *ReadUserResponse) XXX_Size() int {
	return xxx_messageInfo_ReadUserResponse.Size(m)
}
func (m *ReadUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadUserResponse proto.InternalMessageInfo

func (m *ReadUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{5}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(dst, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUserResponse struct {
	RowsDeleted          int64    `protobuf:"varint,1,opt,name=rowsDeleted,proto3" json:"rowsDeleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_780cb17b605244eb, []int{6}
}
func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(dst, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetRowsDeleted() int64 {
	if m != nil {
		return m.RowsDeleted
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "CreateUserResponse")
	proto.RegisterType((*ReadUserRequest)(nil), "ReadUserRequest")
	proto.RegisterType((*ReadUserResponse)(nil), "ReadUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "DeleteUserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Read(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (*ReadUserResponse, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *ReadUserRequest, opts ...grpc.CallOption) (*ReadUserResponse, error) {
	out := new(ReadUserResponse)
	err := c.cc.Invoke(ctx, "/UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Read(context.Context, *ReadUserRequest) (*ReadUserResponse, error)
	Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*ReadUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_780cb17b605244eb) }

var fileDescriptor_user_780cb17b605244eb = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0x0e, 0x2d, 0xa2, 0xfd, 0x48, 0xb4, 0x4c, 0x2f, 0x48, 0x3c, 0x20, 0x7a, 0x68, 0x62, 0xdc,
	0xa4, 0x35, 0xf1, 0x05, 0xf4, 0xec, 0x01, 0xe3, 0x03, 0xac, 0xee, 0x98, 0x60, 0xaa, 0xd4, 0x1d,
	0xaa, 0xaf, 0xe3, 0xa3, 0x1a, 0x16, 0x6a, 0x11, 0x8c, 0xc7, 0x99, 0xef, 0xdb, 0xef, 0x67, 0xb2,
	0xc0, 0x46, 0xd8, 0xaa, 0xb5, 0x2d, 0xab, 0x32, 0x7b, 0x81, 0xff, 0x20, 0x6c, 0xe9, 0x10, 0xa3,
	0xc2, 0xc4, 0x5e, 0xea, 0xcd, 0xc7, 0xf9, 0xa8, 0x30, 0x74, 0x82, 0xc9, 0x73, 0x61, 0xa5, 0xba,
	0xd3, 0xaf, 0x1c, 0x8f, 0x52, 0x6f, 0x3e, 0xc9, 0x77, 0x0b, 0x4a, 0x70, 0xb0, 0xd2, 0x2d, 0x38,
	0x76, 0xe0, 0xcf, 0x4c, 0x31, 0xf6, 0xb5, 0x31, 0x96, 0x45, 0x62, 0xdf, 0x41, 0xdb, 0x31, 0x53,
	0x88, 0x6e, 0x2c, 0xeb, 0x8a, 0x6b, 0xc7, 0x9c, 0xdf, 0x37, 0x2c, 0x15, 0x1d, 0xc3, 0xaf, 0xe3,
	0x38, 0xeb, 0x70, 0xb9, 0xa7, 0x1c, 0xe6, 0x56, 0xd9, 0x39, 0xa8, 0xcb, 0x97, 0x75, 0xf9, 0x26,
	0xdc, 0x4f, 0x9a, 0x9d, 0xe2, 0x28, 0x67, 0x6d, 0xba, 0x9a, 0x7d, 0xca, 0x25, 0xa6, 0x3b, 0x4a,
	0x2b, 0xf3, 0x8f, 0xef, 0x19, 0xa2, 0x5b, 0x5e, 0xf1, 0xef, 0x9c, 0x7d, 0xcd, 0x6b, 0x50, 0x97,
	0xd4, 0xaa, 0xa6, 0x08, 0x6d, 0xf9, 0x29, 0x0d, 0xb2, 0xa5, 0x77, 0x57, 0xcb, 0x2f, 0x0f, 0x61,
	0xfd, 0xe4, 0x9e, 0xed, 0x47, 0xf1, 0xc4, 0xb4, 0x40, 0xd0, 0x94, 0x24, 0x52, 0x83, 0xeb, 0x24,
	0x33, 0xf5, 0xc7, 0x05, 0x2e, 0xe0, 0xd7, 0x75, 0x68, 0xaa, 0x7a, 0xc5, 0x93, 0x48, 0x0d, 0x7a,
	0x2e, 0x10, 0x34, 0xd6, 0x44, 0x6a, 0xd0, 0x2a, 0x99, 0xa9, 0x61, 0x89, 0xc7, 0xc0, 0x7d, 0x8d,
	0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x32, 0x5c, 0xac, 0x28, 0x02, 0x00, 0x00,
}
