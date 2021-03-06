// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package userinterface

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type QueryUserInfoReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserInfoReq) Reset()         { *m = QueryUserInfoReq{} }
func (m *QueryUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoReq) ProtoMessage()    {}
func (*QueryUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *QueryUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoReq.Unmarshal(m, b)
}
func (m *QueryUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoReq.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoReq.Merge(m, src)
}
func (m *QueryUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoReq.Size(m)
}
func (m *QueryUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoReq proto.InternalMessageInfo

func (m *QueryUserInfoReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryUserInfoReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type QueryUserInfoResp struct {
	Code                 int64                   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *QueryUserInfoResp_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QueryUserInfoResp) Reset()         { *m = QueryUserInfoResp{} }
func (m *QueryUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoResp) ProtoMessage()    {}
func (*QueryUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *QueryUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoResp.Unmarshal(m, b)
}
func (m *QueryUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoResp.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoResp.Merge(m, src)
}
func (m *QueryUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoResp.Size(m)
}
func (m *QueryUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoResp proto.InternalMessageInfo

func (m *QueryUserInfoResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryUserInfoResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *QueryUserInfoResp) GetData() *QueryUserInfoResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type QueryUserInfoResp_Data struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Profile              string   `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
	LastActive           string   `protobuf:"bytes,5,opt,name=lastActive,proto3" json:"lastActive,omitempty"`
	SignupAt             string   `protobuf:"bytes,6,opt,name=signupAt,proto3" json:"signupAt,omitempty"`
	EmailValidated       bool     `protobuf:"varint,7,opt,name=emailValidated,proto3" json:"emailValidated,omitempty"`
	PhoneValidated       bool     `protobuf:"varint,8,opt,name=phoneValidated,proto3" json:"phoneValidated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserInfoResp_Data) Reset()         { *m = QueryUserInfoResp_Data{} }
func (m *QueryUserInfoResp_Data) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoResp_Data) ProtoMessage()    {}
func (*QueryUserInfoResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3, 0}
}

func (m *QueryUserInfoResp_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoResp_Data.Unmarshal(m, b)
}
func (m *QueryUserInfoResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoResp_Data.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoResp_Data.Merge(m, src)
}
func (m *QueryUserInfoResp_Data) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoResp_Data.Size(m)
}
func (m *QueryUserInfoResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoResp_Data proto.InternalMessageInfo

func (m *QueryUserInfoResp_Data) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetLastActive() string {
	if m != nil {
		return m.LastActive
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *QueryUserInfoResp_Data) GetEmailValidated() bool {
	if m != nil {
		return m.EmailValidated
	}
	return false
}

func (m *QueryUserInfoResp_Data) GetPhoneValidated() bool {
	if m != nil {
		return m.PhoneValidated
	}
	return false
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "userinterface.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "userinterface.RegisterResp")
	proto.RegisterType((*QueryUserInfoReq)(nil), "userinterface.QueryUserInfoReq")
	proto.RegisterType((*QueryUserInfoResp)(nil), "userinterface.QueryUserInfoResp")
	proto.RegisterType((*QueryUserInfoResp_Data)(nil), "userinterface.QueryUserInfoResp.Data")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcb, 0x4e, 0xe3, 0x30,
	0x14, 0x55, 0xda, 0xf4, 0x75, 0xd3, 0x8e, 0x66, 0xac, 0x59, 0x58, 0x19, 0x69, 0x88, 0x22, 0x81,
	0xba, 0xca, 0xa2, 0xac, 0x90, 0xd8, 0x54, 0x02, 0x21, 0x76, 0x60, 0x1e, 0x7b, 0x93, 0xdc, 0x16,
	0x8b, 0x34, 0x4e, 0x6d, 0xb7, 0x88, 0x5f, 0xe2, 0xf7, 0xe0, 0x03, 0x90, 0xf3, 0xa0, 0x0f, 0x44,
	0x2b, 0x76, 0x3e, 0xe7, 0x5c, 0x9f, 0x7b, 0x72, 0x7d, 0x03, 0xb0, 0xd0, 0xa8, 0xa2, 0x5c, 0x49,
	0x23, 0xc9, 0xc0, 0x9e, 0x45, 0x66, 0x50, 0x4d, 0x78, 0x8c, 0xe1, 0x39, 0x78, 0x0c, 0xa7, 0x42,
	0x1b, 0x54, 0x0c, 0xe7, 0xc4, 0x87, 0xae, 0xd5, 0x33, 0x3e, 0x43, 0xea, 0x04, 0xce, 0xb0, 0xc7,
	0x3e, 0xb1, 0xd5, 0x72, 0xae, 0xf5, 0xb3, 0x54, 0x09, 0x6d, 0x94, 0x5a, 0x8d, 0xc3, 0x53, 0xe8,
	0xaf, 0x6c, 0x74, 0x4e, 0x08, 0xb8, 0xb1, 0x4c, 0x4a, 0x8f, 0x26, 0x2b, 0xce, 0x84, 0x42, 0x67,
	0x86, 0x5a, 0xf3, 0x29, 0x56, 0xd7, 0x6b, 0x18, 0x5e, 0xc1, 0xef, 0xeb, 0x05, 0xaa, 0x97, 0x3b,
	0x8d, 0xea, 0x32, 0x9b, 0xc8, 0x7d, 0x49, 0x02, 0xf0, 0x78, 0x1c, 0xa3, 0xd6, 0xb7, 0xf2, 0x09,
	0xb3, 0xca, 0x6d, 0x9d, 0x0a, 0xdf, 0x1a, 0xf0, 0x67, 0xcb, 0xf2, 0xa7, 0xa9, 0xc8, 0x09, 0xb8,
	0x09, 0x37, 0x9c, 0x36, 0x03, 0x67, 0xe8, 0x8d, 0x0e, 0xa3, 0x8d, 0xc1, 0x45, 0x5f, 0xdc, 0xa3,
	0x33, 0x6e, 0x38, 0x2b, 0xae, 0xf8, 0xef, 0x0e, 0xb8, 0x16, 0xee, 0xfc, 0x8a, 0xbf, 0xd0, 0xc2,
	0x19, 0x17, 0x69, 0xd5, 0xb7, 0x04, 0x96, 0xcd, 0x1f, 0x65, 0x86, 0x45, 0xdb, 0x1e, 0x2b, 0x81,
	0x4d, 0x99, 0x2b, 0x39, 0x11, 0x29, 0x52, 0xb7, 0x4c, 0x59, 0x41, 0xf2, 0x1f, 0x20, 0xe5, 0xda,
	0x8c, 0x63, 0x23, 0x96, 0x48, 0x5b, 0x85, 0xb8, 0xc6, 0xd8, 0x04, 0x5a, 0x4c, 0xb3, 0x45, 0x3e,
	0x36, 0xb4, 0x5d, 0x26, 0xa8, 0x31, 0x39, 0x82, 0x5f, 0x45, 0xd3, 0x7b, 0x9e, 0x8a, 0x84, 0x1b,
	0x4c, 0x68, 0x27, 0x70, 0x86, 0x5d, 0xb6, 0xc5, 0xda, 0xba, 0x22, 0xc6, 0xaa, 0xae, 0x5b, 0xd6,
	0x6d, 0xb2, 0xa3, 0x57, 0x07, 0x3c, 0x3b, 0x92, 0x1b, 0x54, 0x4b, 0x11, 0x23, 0xb9, 0x80, 0xbe,
	0x85, 0xf5, 0x66, 0x10, 0x7f, 0x6b, 0x86, 0x6b, 0x9b, 0xe7, 0xff, 0xfb, 0x56, 0xd3, 0x39, 0x61,
	0x30, 0xd8, 0x98, 0x37, 0x39, 0xd8, 0xfd, 0x1a, 0x73, 0x3f, 0xd8, 0xf7, 0x5c, 0x0f, 0xed, 0xe2,
	0x7f, 0x38, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa6, 0x85, 0x71, 0x1d, 0x03, 0x00, 0x00,
}
