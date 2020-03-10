// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nullstring.proto

/*
Package nullstring is a generated protocol buffer package.

It is generated from these files:
	nullstring.proto

It has these top-level messages:
	NullString
*/
package nullstring

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// NullString ...
type NullString struct {
	// Represents the actual string
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// is set to true if the string is null
	IsNotNull bool `protobuf:"varint,2,opt,name=is_not_null,json=isNotNull" json:"is_not_null,omitempty"`
}

func (m *NullString) Reset()                    { *m = NullString{} }
func (m *NullString) String() string            { return proto.CompactTextString(m) }
func (*NullString) ProtoMessage()               {}
func (*NullString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NullString) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *NullString) GetIsNotNull() bool {
	if m != nil {
		return m.IsNotNull
	}
	return false
}

func init() {
	proto.RegisterType((*NullString)(nil), "weiwolves.protobuf.NullString")
}

func init() { proto.RegisterFile("nullstring.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x2b, 0xcd, 0xc9,
	0x29, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xc9,
	0x4e, 0x4b, 0x4a, 0x2c, 0x4e, 0xcd, 0x81, 0xf0, 0x93, 0x4a, 0xd3, 0x94, 0x1c, 0xb8, 0xb8, 0xfc,
	0x4a, 0x73, 0x72, 0x82, 0xc1, 0xca, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x39, 0x2e, 0xee, 0xcc, 0xe2, 0xf8, 0xbc, 0xfc,
	0x92, 0x78, 0x90, 0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x9c, 0x99, 0xc5, 0x7e, 0xf9,
	0x25, 0x20, 0x9d, 0x4e, 0xfa, 0x51, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0x30, 0x1b, 0xf4, 0x61, 0x36, 0xe8, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x23, 0xdc,
	0x92, 0xc4, 0x06, 0x96, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xa8, 0x8d, 0xeb, 0xa0,
	0x00, 0x00, 0x00,
}
