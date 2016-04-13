// Code generated by protoc-gen-go.
// source: src/protos/system/system.proto
// DO NOT EDIT!

/*
Package system is a generated protocol buffer package.

It is generated from these files:
	src/protos/system/system.proto

It has these top-level messages:
	System_LogC2S
	System_LogS2C
*/
package system

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

type System_LogC2S struct {
	Dir              *string `protobuf:"bytes,1,req,name=Dir,json=dir" json:"Dir,omitempty"`
	Type             *uint32 `protobuf:"varint,2,req,name=Type,json=type" json:"Type,omitempty"`
	Content          *string `protobuf:"bytes,3,req,name=Content,json=content" json:"Content,omitempty"`
	Time             *int64  `protobuf:"varint,4,req,name=Time,json=time" json:"Time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *System_LogC2S) Reset()                    { *m = System_LogC2S{} }
func (m *System_LogC2S) String() string            { return proto.CompactTextString(m) }
func (*System_LogC2S) ProtoMessage()               {}
func (*System_LogC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *System_LogC2S) GetDir() string {
	if m != nil && m.Dir != nil {
		return *m.Dir
	}
	return ""
}

func (m *System_LogC2S) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *System_LogC2S) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *System_LogC2S) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type System_LogS2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_LogS2C) Reset()                    { *m = System_LogS2C{} }
func (m *System_LogS2C) String() string            { return proto.CompactTextString(m) }
func (*System_LogS2C) ProtoMessage()               {}
func (*System_LogS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*System_LogC2S)(nil), "System_LogC2S")
	proto.RegisterType((*System_LogS2C)(nil), "System_LogS2C")
}

var fileDescriptor0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0x85, 0x52, 0x7a,
	0x60, 0x41, 0xa5, 0x64, 0x2e, 0xde, 0x60, 0x30, 0x3f, 0xde, 0x27, 0x3f, 0xdd, 0xd9, 0x28, 0x58,
	0x48, 0x80, 0x8b, 0xd9, 0x25, 0xb3, 0x48, 0x82, 0x51, 0x81, 0x49, 0x83, 0x33, 0x88, 0x39, 0x25,
	0xb3, 0x48, 0x48, 0x88, 0x8b, 0x25, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x09, 0x28, 0xc4, 0x1b, 0xc4,
	0x52, 0x02, 0x64, 0x0b, 0x49, 0x70, 0xb1, 0x3b, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0x30,
	0x83, 0x55, 0xb2, 0x27, 0x43, 0xb8, 0x60, 0xd5, 0x99, 0xb9, 0xa9, 0x12, 0x2c, 0x40, 0x61, 0x66,
	0xa0, 0x6a, 0x20, 0x5b, 0x89, 0x1f, 0xd9, 0x92, 0x60, 0x23, 0x67, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf1, 0x91, 0x50, 0xd2, 0x96, 0x00, 0x00, 0x00,
}
