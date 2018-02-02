// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcservice.proto

/*
Package grpcservice is a generated protocol buffer package.

It is generated from these files:
	grpcservice.proto

It has these top-level messages:
	WrinkledItem
	SmoothItem
	Item
*/
package grpcservice

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

// WrinkledItem is used for items with multiple wrinkes
type WrinkledItem struct {
	// customer will contain the name of the customer
	Customer string `protobuf:"bytes,1,opt,name=Customer" json:"Customer,omitempty"`
	// nested messages can be embedded directly using the tag compose:"embed"
	Item *Item `protobuf:"bytes,2,opt,name=Item" json:"Item,omitempty"`
}

func (m *WrinkledItem) Reset()                    { *m = WrinkledItem{} }
func (m *WrinkledItem) String() string            { return proto.CompactTextString(m) }
func (*WrinkledItem) ProtoMessage()               {}
func (*WrinkledItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WrinkledItem) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *WrinkledItem) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

// SmoothItem is used as response with associated costs
type SmoothItem struct {
	Item *Item `protobuf:"bytes,1,opt,name=Item" json:"Item,omitempty"`
	Cost int32 `protobuf:"varint,2,opt,name=Cost" json:"Cost,omitempty"`
}

func (m *SmoothItem) Reset()                    { *m = SmoothItem{} }
func (m *SmoothItem) String() string            { return proto.CompactTextString(m) }
func (*SmoothItem) ProtoMessage()               {}
func (*SmoothItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SmoothItem) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *SmoothItem) GetCost() int32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

// Item contains the information about a specific item
type Item struct {
	Name     string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Wrinkels int32  `protobuf:"varint,2,opt,name=Wrinkels" json:"Wrinkels,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetWrinkels() int32 {
	if m != nil {
		return m.Wrinkels
	}
	return 0
}

func init() {
	proto.RegisterType((*WrinkledItem)(nil), "grpcservice.WrinkledItem")
	proto.RegisterType((*SmoothItem)(nil), "grpcservice.SmoothItem")
	proto.RegisterType((*Item)(nil), "grpcservice.Item")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Butler service

type ButlerClient interface {
	// Iron will take a wrinkled item and remove all wrinkels
	Iron(ctx context.Context, in *WrinkledItem, opts ...grpc.CallOption) (*SmoothItem, error)
}

type butlerClient struct {
	cc *grpc.ClientConn
}

func NewButlerClient(cc *grpc.ClientConn) ButlerClient {
	return &butlerClient{cc}
}

func (c *butlerClient) Iron(ctx context.Context, in *WrinkledItem, opts ...grpc.CallOption) (*SmoothItem, error) {
	out := new(SmoothItem)
	err := grpc.Invoke(ctx, "/grpcservice.Butler/Iron", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Butler service

type ButlerServer interface {
	// Iron will take a wrinkled item and remove all wrinkels
	Iron(context.Context, *WrinkledItem) (*SmoothItem, error)
}

func RegisterButlerServer(s *grpc.Server, srv ButlerServer) {
	s.RegisterService(&_Butler_serviceDesc, srv)
}

func _Butler_Iron_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrinkledItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ButlerServer).Iron(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcservice.Butler/Iron",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ButlerServer).Iron(ctx, req.(*WrinkledItem))
	}
	return interceptor(ctx, in, info, handler)
}

var _Butler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcservice.Butler",
	HandlerType: (*ButlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Iron",
			Handler:    _Butler_Iron_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcservice.proto",
}

func init() { proto.RegisterFile("grpcservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2f, 0x2a, 0x48,
	0x2e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x0a, 0xe4, 0xe2, 0x09, 0x2f, 0xca, 0xcc, 0xcb, 0xce, 0x49, 0x4d, 0xf1, 0x2c, 0x49,
	0xcd, 0x15, 0x92, 0xe2, 0xe2, 0x70, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2d, 0x92, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x85, 0x54, 0xb9, 0x58, 0x40, 0x6a, 0x24, 0x98, 0x14, 0x18,
	0x35, 0xb8, 0x8d, 0x04, 0xf5, 0x90, 0x8d, 0x06, 0x49, 0x04, 0x81, 0xa5, 0x95, 0xdc, 0xb9, 0xb8,
	0x82, 0x73, 0xf3, 0xf3, 0x4b, 0x32, 0xc0, 0x06, 0xc2, 0x34, 0x31, 0xe2, 0xd5, 0x24, 0x24, 0xc4,
	0xc5, 0xe2, 0x9c, 0x5f, 0x5c, 0x02, 0x36, 0x9b, 0x35, 0x08, 0xcc, 0x56, 0x32, 0xe3, 0x82, 0xcb,
	0xf9, 0x25, 0xe6, 0xa6, 0x42, 0xdd, 0x03, 0x66, 0x83, 0xdc, 0x09, 0x76, 0x77, 0x6a, 0x4e, 0x31,
	0x54, 0x0f, 0x9c, 0x6f, 0xe4, 0xc6, 0xc5, 0xe6, 0x54, 0x5a, 0x92, 0x93, 0x5a, 0x24, 0x64, 0xc3,
	0xc5, 0xe2, 0x59, 0x94, 0x9f, 0x27, 0x24, 0x89, 0x62, 0x2d, 0xb2, 0x87, 0xa5, 0xc4, 0x51, 0xa4,
	0x10, 0x0e, 0x57, 0x62, 0x48, 0x62, 0x03, 0x87, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0xed, 0x60, 0x87, 0x44, 0x01, 0x00, 0x00,
}
