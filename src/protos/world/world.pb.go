// Code generated by protoc-gen-go.
// source: src/protos/world/world.proto
// DO NOT EDIT!

/*
Package world is a generated protocol buffer package.

It is generated from these files:
	src/protos/world/world.proto

It has these top-level messages:
	World_PingC2S
	World_PingS2C
	World_RegisterRoleIDC2S
	World_RegisterRoleIDS2C
	World_SendChatC2S
	World_SendChatS2C
	World_Receive_ChatS2C
*/
package world

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type World_PingC2S struct {
	Content          *string `protobuf:"bytes,1,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *World_PingC2S) Reset()                    { *m = World_PingC2S{} }
func (m *World_PingC2S) String() string            { return proto.CompactTextString(m) }
func (*World_PingC2S) ProtoMessage()               {}
func (*World_PingC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *World_PingC2S) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type World_PingS2C struct {
	Content          *string `protobuf:"bytes,1,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *World_PingS2C) Reset()                    { *m = World_PingS2C{} }
func (m *World_PingS2C) String() string            { return proto.CompactTextString(m) }
func (*World_PingS2C) ProtoMessage()               {}
func (*World_PingS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *World_PingS2C) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type World_RegisterRoleIDC2S struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *World_RegisterRoleIDC2S) Reset()                    { *m = World_RegisterRoleIDC2S{} }
func (m *World_RegisterRoleIDC2S) String() string            { return proto.CompactTextString(m) }
func (*World_RegisterRoleIDC2S) ProtoMessage()               {}
func (*World_RegisterRoleIDC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *World_RegisterRoleIDC2S) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type World_RegisterRoleIDS2C struct {
	Result           *bool   `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	RoleID           *uint64 `protobuf:"varint,2,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *World_RegisterRoleIDS2C) Reset()                    { *m = World_RegisterRoleIDS2C{} }
func (m *World_RegisterRoleIDS2C) String() string            { return proto.CompactTextString(m) }
func (*World_RegisterRoleIDS2C) ProtoMessage()               {}
func (*World_RegisterRoleIDS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *World_RegisterRoleIDS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func (m *World_RegisterRoleIDS2C) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type World_SendChatC2S struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	Content          *string `protobuf:"bytes,2,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *World_SendChatC2S) Reset()                    { *m = World_SendChatC2S{} }
func (m *World_SendChatC2S) String() string            { return proto.CompactTextString(m) }
func (*World_SendChatC2S) ProtoMessage()               {}
func (*World_SendChatC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *World_SendChatC2S) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *World_SendChatC2S) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type World_SendChatS2C struct {
	Result           *bool  `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *World_SendChatS2C) Reset()                    { *m = World_SendChatS2C{} }
func (m *World_SendChatS2C) String() string            { return proto.CompactTextString(m) }
func (*World_SendChatS2C) ProtoMessage()               {}
func (*World_SendChatS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *World_SendChatS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

// 接收聊天信息
type World_Receive_ChatS2C struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	Content          *string `protobuf:"bytes,3,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *World_Receive_ChatS2C) Reset()                    { *m = World_Receive_ChatS2C{} }
func (m *World_Receive_ChatS2C) String() string            { return proto.CompactTextString(m) }
func (*World_Receive_ChatS2C) ProtoMessage()               {}
func (*World_Receive_ChatS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *World_Receive_ChatS2C) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *World_Receive_ChatS2C) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*World_PingC2S)(nil), "World_PingC2S")
	proto.RegisterType((*World_PingS2C)(nil), "World_PingS2C")
	proto.RegisterType((*World_RegisterRoleIDC2S)(nil), "World_RegisterRoleIDC2S")
	proto.RegisterType((*World_RegisterRoleIDS2C)(nil), "World_RegisterRoleIDS2C")
	proto.RegisterType((*World_SendChatC2S)(nil), "World_SendChatC2S")
	proto.RegisterType((*World_SendChatS2C)(nil), "World_SendChatS2C")
	proto.RegisterType((*World_Receive_ChatS2C)(nil), "World_Receive_ChatS2C")
}

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x81, 0x90, 0x7a, 0x60,
	0x21, 0x25, 0x4d, 0x2e, 0xde, 0x70, 0x10, 0x37, 0x3e, 0x20, 0x33, 0x2f, 0xdd, 0xd9, 0x28, 0x58,
	0x48, 0x82, 0x8b, 0xdd, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf, 0x44, 0x82, 0x51, 0x81, 0x49, 0x83,
	0x33, 0x88, 0x3d, 0x19, 0xc2, 0x45, 0x55, 0x1a, 0x6c, 0xe4, 0x8c, 0x47, 0xa9, 0x21, 0x97, 0x38,
	0x44, 0x69, 0x50, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x51, 0x50, 0x7e, 0x4e, 0xaa, 0xa7, 0x0b,
	0xc8, 0x7c, 0x31, 0x2e, 0x36, 0x08, 0x07, 0xac, 0x87, 0x25, 0x88, 0xad, 0x08, 0xcc, 0x53, 0xf2,
	0xc4, 0xae, 0x05, 0x64, 0x0f, 0x48, 0x4b, 0x6a, 0x71, 0x69, 0x0e, 0xc4, 0x1a, 0x0e, 0xa0, 0x16,
	0x30, 0x0f, 0xc9, 0x28, 0x26, 0x14, 0xa3, 0x5c, 0xb9, 0x04, 0x21, 0x46, 0x05, 0xa7, 0xe6, 0xa5,
	0x38, 0x67, 0x24, 0x96, 0xe0, 0xb1, 0x17, 0xd9, 0x13, 0x4c, 0xa8, 0x9e, 0xd0, 0x46, 0x37, 0x06,
	0x8f, 0x5b, 0x80, 0xce, 0x17, 0x85, 0x39, 0x3f, 0x39, 0x35, 0xb3, 0x2c, 0x35, 0x1e, 0x59, 0x03,
	0x01, 0x7b, 0x99, 0x51, 0xec, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0x35, 0x0e, 0xed, 0xa5, 0xb1,
	0x01, 0x00, 0x00,
}
