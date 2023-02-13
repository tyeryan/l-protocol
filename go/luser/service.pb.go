// Code generated by protoc-gen-go. DO NOT EDIT.
// source: luser/service.proto

package luser // import "github.com/tyeryan/l-protocol/go/luser"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type UserCreationReq struct {
	MCC                  string   `protobuf:"bytes,1,opt,name=MCC,proto3" json:"MCC,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Lang                 string   `protobuf:"bytes,3,opt,name=Lang,proto3" json:"Lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreationReq) Reset()         { *m = UserCreationReq{} }
func (m *UserCreationReq) String() string { return proto.CompactTextString(m) }
func (*UserCreationReq) ProtoMessage()    {}
func (*UserCreationReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ee490807268de434, []int{0}
}
func (m *UserCreationReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreationReq.Unmarshal(m, b)
}
func (m *UserCreationReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreationReq.Marshal(b, m, deterministic)
}
func (dst *UserCreationReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreationReq.Merge(dst, src)
}
func (m *UserCreationReq) XXX_Size() int {
	return xxx_messageInfo_UserCreationReq.Size(m)
}
func (m *UserCreationReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreationReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreationReq proto.InternalMessageInfo

func (m *UserCreationReq) GetMCC() string {
	if m != nil {
		return m.MCC
	}
	return ""
}

func (m *UserCreationReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserCreationReq) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type UserCreationRsp struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IsFound              bool     `protobuf:"varint,2,opt,name=IsFound,proto3" json:"IsFound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreationRsp) Reset()         { *m = UserCreationRsp{} }
func (m *UserCreationRsp) String() string { return proto.CompactTextString(m) }
func (*UserCreationRsp) ProtoMessage()    {}
func (*UserCreationRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ee490807268de434, []int{1}
}
func (m *UserCreationRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreationRsp.Unmarshal(m, b)
}
func (m *UserCreationRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreationRsp.Marshal(b, m, deterministic)
}
func (dst *UserCreationRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreationRsp.Merge(dst, src)
}
func (m *UserCreationRsp) XXX_Size() int {
	return xxx_messageInfo_UserCreationRsp.Size(m)
}
func (m *UserCreationRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreationRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreationRsp proto.InternalMessageInfo

func (m *UserCreationRsp) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UserCreationRsp) GetIsFound() bool {
	if m != nil {
		return m.IsFound
	}
	return false
}

type User struct {
	MCC                  string               `protobuf:"bytes,2,opt,name=MCC,proto3" json:"MCC,omitempty"`
	PhoneNumber          string               `protobuf:"bytes,3,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email                string               `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	EmailVerified        bool                 `protobuf:"varint,5,opt,name=EmailVerified,proto3" json:"EmailVerified,omitempty"`
	FirstName            string               `protobuf:"bytes,6,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string               `protobuf:"bytes,7,opt,name=LastName,proto3" json:"LastName,omitempty"`
	DOB                  string               `protobuf:"bytes,8,opt,name=DOB,proto3" json:"DOB,omitempty"`
	Nationality          string               `protobuf:"bytes,9,opt,name=Nationality,proto3" json:"Nationality,omitempty"`
	Language             string               `protobuf:"bytes,10,opt,name=Language,proto3" json:"Language,omitempty"`
	Gender               string               `protobuf:"bytes,11,opt,name=Gender,proto3" json:"Gender,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,18,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,19,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ee490807268de434, []int{2}
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

func (m *User) GetMCC() string {
	if m != nil {
		return m.MCC
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
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

func (m *User) GetDOB() string {
	if m != nil {
		return m.DOB
	}
	return ""
}

func (m *User) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *User) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *User) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*UserCreationReq)(nil), "luser.UserCreationReq")
	proto.RegisterType((*UserCreationRsp)(nil), "luser.UserCreationRsp")
	proto.RegisterType((*User)(nil), "luser.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LUserClient is the client API for LUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LUserClient interface {
	GetOrCreateNewUser(ctx context.Context, in *UserCreationReq, opts ...grpc.CallOption) (*UserCreationRsp, error)
}

type lUserClient struct {
	cc *grpc.ClientConn
}

func NewLUserClient(cc *grpc.ClientConn) LUserClient {
	return &lUserClient{cc}
}

func (c *lUserClient) GetOrCreateNewUser(ctx context.Context, in *UserCreationReq, opts ...grpc.CallOption) (*UserCreationRsp, error) {
	out := new(UserCreationRsp)
	err := c.cc.Invoke(ctx, "/luser.LUser/GetOrCreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LUserServer is the server API for LUser service.
type LUserServer interface {
	GetOrCreateNewUser(context.Context, *UserCreationReq) (*UserCreationRsp, error)
}

func RegisterLUserServer(s *grpc.Server, srv LUserServer) {
	s.RegisterService(&_LUser_serviceDesc, srv)
}

func _LUser_GetOrCreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LUserServer).GetOrCreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luser.LUser/GetOrCreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LUserServer).GetOrCreateNewUser(ctx, req.(*UserCreationReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luser.LUser",
	HandlerType: (*LUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrCreateNewUser",
			Handler:    _LUser_GetOrCreateNewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "luser/service.proto",
}

func init() { proto.RegisterFile("luser/service.proto", fileDescriptor_service_ee490807268de434) }

var fileDescriptor_service_ee490807268de434 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xb5, 0x9f, 0xdb, 0xde, 0xe2, 0x07, 0xb3, 0xb2, 0x0c, 0x45, 0xb0, 0x14, 0x91, 0xbe, 0x98,
	0xc0, 0xfa, 0x22, 0xf8, 0x64, 0x1b, 0xbb, 0x14, 0x6a, 0x22, 0xc1, 0x15, 0xf4, 0x6d, 0xd2, 0xdc,
	0xcd, 0x0e, 0x24, 0x99, 0x38, 0x33, 0x51, 0xfa, 0x1f, 0xfc, 0xd1, 0x92, 0x3b, 0xc9, 0xee, 0xba,
	0x2a, 0xbe, 0xdd, 0x7b, 0xce, 0x3d, 0xa7, 0xa7, 0x39, 0x03, 0xa7, 0x79, 0x6d, 0x50, 0xfb, 0x06,
	0xf5, 0x77, 0x79, 0x40, 0xaf, 0xd2, 0xca, 0x2a, 0x36, 0x22, 0x70, 0xfe, 0x3c, 0x53, 0x2a, 0xcb,
	0xd1, 0x27, 0x30, 0xa9, 0xaf, 0x7c, 0x2b, 0x0b, 0x34, 0x56, 0x14, 0x95, 0xbb, 0x5b, 0x7e, 0x81,
	0xc7, 0x97, 0x06, 0xf5, 0x46, 0xa3, 0xb0, 0x52, 0x95, 0x31, 0x7e, 0x63, 0x4f, 0x60, 0xf0, 0x61,
	0xb3, 0xe1, 0xbd, 0x45, 0x6f, 0x35, 0x8d, 0x9b, 0x91, 0x2d, 0x60, 0xf6, 0xf1, 0x5a, 0x95, 0x18,
	0xd6, 0x45, 0x82, 0x9a, 0xf7, 0x89, 0xb9, 0x0b, 0x31, 0x06, 0xc3, 0xbd, 0x28, 0x33, 0x3e, 0x20,
	0x8a, 0xe6, 0xe5, 0xdb, 0x7b, 0xd6, 0xa6, 0x62, 0x8f, 0xa0, 0xbf, 0x0b, 0x5a, 0xe7, 0xfe, 0x2e,
	0x60, 0x1c, 0x4e, 0x76, 0x66, 0xab, 0xea, 0x32, 0x25, 0xd3, 0x49, 0xdc, 0xad, 0xcb, 0x9f, 0x03,
	0x18, 0x36, 0xea, 0x2e, 0x4d, 0xff, 0x9f, 0x69, 0x06, 0x7f, 0xa6, 0x79, 0x0a, 0xa3, 0xf7, 0x85,
	0x90, 0x39, 0x1f, 0x12, 0xe7, 0x16, 0xf6, 0x02, 0x1e, 0xd2, 0xf0, 0x19, 0xb5, 0xbc, 0x92, 0x98,
	0xf2, 0x11, 0xfd, 0xe4, 0xef, 0x20, 0x7b, 0x06, 0xd3, 0xad, 0xd4, 0xc6, 0x86, 0xa2, 0x40, 0x3e,
	0x26, 0xfd, 0x2d, 0xc0, 0xe6, 0x30, 0xd9, 0x8b, 0x96, 0x3c, 0x21, 0xf2, 0x66, 0x6f, 0x92, 0x06,
	0xd1, 0x9a, 0x4f, 0x5c, 0xd2, 0x20, 0x5a, 0x37, 0x49, 0x43, 0xfa, 0xef, 0x22, 0x97, 0xf6, 0xc8,
	0xa7, 0x2e, 0xe9, 0x1d, 0xc8, 0xf9, 0x95, 0x59, 0x2d, 0x32, 0xe4, 0xd0, 0xf9, 0xb9, 0x9d, 0x9d,
	0xc1, 0xf8, 0x02, 0xcb, 0x14, 0x35, 0x9f, 0x11, 0xd3, 0x6e, 0xec, 0x0d, 0x4c, 0xe9, 0x9b, 0x62,
	0xfa, 0xce, 0x72, 0xb6, 0xe8, 0xad, 0x66, 0xe7, 0x73, 0xcf, 0xf5, 0xec, 0x75, 0x3d, 0x7b, 0x9f,
	0xba, 0x9e, 0xe3, 0xdb, 0xe3, 0x46, 0x79, 0x59, 0xa5, 0xad, 0xf2, 0xf4, 0xff, 0xca, 0x9b, 0xe3,
	0xf3, 0x08, 0x46, 0x7b, 0xaa, 0x63, 0x0b, 0xec, 0x02, 0x6d, 0xe4, 0x5a, 0xc5, 0x10, 0x7f, 0x10,
	0x7a, 0xe6, 0xd1, 0x73, 0xf3, 0xee, 0x3d, 0xa5, 0xf9, 0x5f, 0x71, 0x53, 0x2d, 0x1f, 0xac, 0x57,
	0x5f, 0x5f, 0x66, 0xd2, 0x5e, 0xd7, 0x89, 0x77, 0x50, 0x85, 0x6f, 0x8f, 0xa8, 0x8f, 0xa2, 0xf4,
	0xf3, 0x57, 0x14, 0xe3, 0xa0, 0x72, 0x3f, 0x53, 0x3e, 0x69, 0x93, 0x31, 0x41, 0xaf, 0x7f, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x49, 0x5a, 0xcb, 0xe7, 0x02, 0x00, 0x00,
}
