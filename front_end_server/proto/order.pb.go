// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Side int32

const (
	Side_SELL Side = 0
	Side_BUY  Side = 1
)

var Side_name = map[int32]string{
	0: "SELL",
	1: "BUY",
}

var Side_value = map[string]int32{
	"SELL": 0,
	"BUY":  1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}

func (Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

type Type int32

const (
	Type_LIMIT  Type = 0
	Type_MARKET Type = 1
	Type_STOP   Type = 2
)

var Type_name = map[int32]string{
	0: "LIMIT",
	1: "MARKET",
	2: "STOP",
}

var Type_value = map[string]int32{
	"LIMIT":  0,
	"MARKET": 1,
	"STOP":   2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

type OrderRequest struct {
	UserId               string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId              string               `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Amount               uint64               `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Price                uint64               `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Side                 Side                 `protobuf:"varint,5,opt,name=Side,proto3,enum=proto.Side" json:"Side,omitempty"`
	Type                 Type                 `protobuf:"varint,6,opt,name=type,proto3,enum=proto.Type" json:"type,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OrderRequest) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderRequest) GetSide() Side {
	if m != nil {
		return m.Side
	}
	return Side_SELL
}

func (m *OrderRequest) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_LIMIT
}

func (m *OrderRequest) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.Side", Side_name, Side_value)
	proto.RegisterEnum("proto.Type", Type_name, Type_value)
	proto.RegisterType((*OrderRequest)(nil), "proto.OrderRequest")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x6f, 0xf2, 0x30,
	0x18, 0xc4, 0x31, 0xe4, 0x0f, 0x79, 0xf2, 0xea, 0x55, 0x64, 0x55, 0xad, 0x61, 0x21, 0xea, 0xd2,
	0x88, 0x21, 0x48, 0x74, 0xea, 0x48, 0x25, 0x86, 0xa8, 0x41, 0x54, 0x26, 0x1d, 0x3a, 0xa1, 0x80,
	0x9f, 0xa2, 0x48, 0x4d, 0x9d, 0x3a, 0xce, 0xc0, 0x87, 0xee, 0x77, 0xa8, 0xec, 0xc0, 0xd0, 0xc9,
	0xfa, 0x3d, 0x77, 0xa7, 0xf3, 0x41, 0x28, 0x95, 0x40, 0x95, 0x36, 0x4a, 0x6a, 0x49, 0x5d, 0xfb,
	0x4c, 0x67, 0x27, 0x29, 0x4f, 0x9f, 0xb8, 0xb0, 0x74, 0xe8, 0x3e, 0x16, 0xba, 0xaa, 0xb1, 0xd5,
	0x65, 0xdd, 0xf4, 0xbe, 0xfb, 0x1f, 0x02, 0xff, 0xb6, 0x26, 0xc7, 0xf1, 0xbb, 0xc3, 0x56, 0xd3,
	0x3b, 0xf0, 0xbb, 0x16, 0xd5, 0xbe, 0x12, 0x8c, 0xc4, 0x24, 0x09, 0xb8, 0x67, 0x30, 0x13, 0x74,
	0x02, 0x63, 0x5b, 0x60, 0x94, 0xa1, 0x55, 0x7c, 0xcb, 0x99, 0xa0, 0xb7, 0xe0, 0x95, 0xb5, 0xec,
	0xbe, 0x34, 0x1b, 0xc5, 0x24, 0x71, 0xf8, 0x85, 0xe8, 0x0d, 0xb8, 0x8d, 0xaa, 0x8e, 0xc8, 0x1c,
	0x7b, 0xee, 0x81, 0xce, 0xc0, 0xd9, 0x55, 0x02, 0x99, 0x1b, 0x93, 0xe4, 0xff, 0x32, 0xec, 0x3f,
	0x92, 0x9a, 0x13, 0xb7, 0x82, 0x31, 0xe8, 0x73, 0x83, 0xcc, 0xfb, 0x63, 0x28, 0xce, 0x0d, 0x72,
	0x2b, 0xd0, 0x27, 0x80, 0xa3, 0xc2, 0x52, 0xa3, 0xd8, 0x97, 0x9a, 0xf9, 0x31, 0x49, 0xc2, 0xe5,
	0x34, 0xed, 0xa7, 0xa6, 0xd7, 0xa9, 0x69, 0x71, 0x9d, 0xca, 0x83, 0x8b, 0x7b, 0xa5, 0xe7, 0x93,
	0xbe, 0x9c, 0x8e, 0xc1, 0xd9, 0xad, 0xf3, 0x3c, 0x1a, 0x50, 0x1f, 0x46, 0xcf, 0x6f, 0xef, 0x11,
	0x99, 0x3f, 0x80, 0x63, 0x3a, 0x68, 0x00, 0x6e, 0x9e, 0x6d, 0xb2, 0x22, 0x1a, 0x50, 0x00, 0x6f,
	0xb3, 0xe2, 0x2f, 0xeb, 0x22, 0x22, 0x36, 0x51, 0x6c, 0x5f, 0xa3, 0xe1, 0xc1, 0xb3, 0x15, 0x8f,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x49, 0x84, 0x6c, 0x71, 0x01, 0x00, 0x00,
}
