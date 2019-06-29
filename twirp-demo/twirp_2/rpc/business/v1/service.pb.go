// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package business_v1

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

type DelReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelReq) Reset()         { *m = DelReq{} }
func (m *DelReq) String() string { return proto.CompactTextString(m) }
func (*DelReq) ProtoMessage()    {}
func (*DelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_90af7b75fdb7eab7, []int{0}
}
func (m *DelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelReq.Unmarshal(m, b)
}
func (m *DelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelReq.Marshal(b, m, deterministic)
}
func (dst *DelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelReq.Merge(dst, src)
}
func (m *DelReq) XXX_Size() int {
	return xxx_messageInfo_DelReq.Size(m)
}
func (m *DelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelReq proto.InternalMessageInfo

func (m *DelReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SetReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReq) Reset()         { *m = SetReq{} }
func (m *SetReq) String() string { return proto.CompactTextString(m) }
func (*SetReq) ProtoMessage()    {}
func (*SetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_90af7b75fdb7eab7, []int{1}
}
func (m *SetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReq.Unmarshal(m, b)
}
func (m *SetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReq.Marshal(b, m, deterministic)
}
func (dst *SetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReq.Merge(dst, src)
}
func (m *SetReq) XXX_Size() int {
	return xxx_messageInfo_SetReq.Size(m)
}
func (m *SetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetReq proto.InternalMessageInfo

func (m *SetReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetResp struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *DefaultResp `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SetResp) Reset()         { *m = SetResp{} }
func (m *SetResp) String() string { return proto.CompactTextString(m) }
func (*SetResp) ProtoMessage()    {}
func (*SetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_90af7b75fdb7eab7, []int{2}
}
func (m *SetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResp.Unmarshal(m, b)
}
func (m *SetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResp.Marshal(b, m, deterministic)
}
func (dst *SetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResp.Merge(dst, src)
}
func (m *SetResp) XXX_Size() int {
	return xxx_messageInfo_SetResp.Size(m)
}
func (m *SetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResp.DiscardUnknown(m)
}

var xxx_messageInfo_SetResp proto.InternalMessageInfo

func (m *SetResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SetResp) GetData() *DefaultResp {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_90af7b75fdb7eab7, []int{3}
}
func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (dst *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(dst, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetResp struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Time                 string       `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Data                 *DefaultResp `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_90af7b75fdb7eab7, []int{4}
}
func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (dst *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(dst, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetResp) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *GetResp) GetData() *DefaultResp {
	if m != nil {
		return m.Data
	}
	return nil
}

type DefaultResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultResp) Reset()         { *m = DefaultResp{} }
func (m *DefaultResp) String() string { return proto.CompactTextString(m) }
func (*DefaultResp) ProtoMessage()    {}
func (*DefaultResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_90af7b75fdb7eab7, []int{5}
}
func (m *DefaultResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefaultResp.Unmarshal(m, b)
}
func (m *DefaultResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefaultResp.Marshal(b, m, deterministic)
}
func (dst *DefaultResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultResp.Merge(dst, src)
}
func (m *DefaultResp) XXX_Size() int {
	return xxx_messageInfo_DefaultResp.Size(m)
}
func (m *DefaultResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultResp.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultResp proto.InternalMessageInfo

func (m *DefaultResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DefaultResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*DelReq)(nil), "business.v1.DelReq")
	proto.RegisterType((*SetReq)(nil), "business.v1.SetReq")
	proto.RegisterType((*SetResp)(nil), "business.v1.SetResp")
	proto.RegisterType((*GetReq)(nil), "business.v1.GetReq")
	proto.RegisterType((*GetResp)(nil), "business.v1.GetResp")
	proto.RegisterType((*DefaultResp)(nil), "business.v1.DefaultResp")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_90af7b75fdb7eab7) }

var fileDescriptor_service_90af7b75fdb7eab7 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x1e, 0xa4, 0xb0, 0x11, 0x08, 0x19, 0x0e, 0x56, 0xc5, 0xa1, 0xca, 0xa9, 0x07, 0x64,
	0x41, 0xcb, 0x17, 0xa0, 0x48, 0xbe, 0x3b, 0x5f, 0xe0, 0xd6, 0x0b, 0xb2, 0x94, 0xd4, 0xa1, 0x76,
	0xfb, 0x45, 0x7c, 0x28, 0xb2, 0x9d, 0x42, 0x49, 0xa9, 0xc4, 0x6d, 0x34, 0xe3, 0xd9, 0x99, 0xf5,
	0xc2, 0xb5, 0xc5, 0xed, 0x5e, 0xaf, 0x91, 0xf5, 0x5b, 0xe3, 0x0c, 0x29, 0x57, 0x3b, 0xab, 0x37,
	0x68, 0x2d, 0xdb, 0x3f, 0x57, 0x14, 0x8a, 0x1a, 0x5b, 0x81, 0x1f, 0xe4, 0x06, 0x52, 0xad, 0x68,
	0x32, 0x4b, 0xe6, 0x17, 0x22, 0xd5, 0xaa, 0x7a, 0x80, 0xa2, 0x41, 0xe7, 0x15, 0x02, 0xf9, 0x46,
	0x76, 0x18, 0xb4, 0x2b, 0x11, 0x70, 0xc5, 0x61, 0x12, 0x54, 0xdb, 0x8f, 0x8d, 0xe4, 0x11, 0x72,
	0x25, 0x9d, 0xa4, 0xe9, 0x2c, 0x99, 0x97, 0x0b, 0xca, 0x8e, 0xe2, 0x58, 0x8d, 0x6f, 0x72, 0xd7,
	0x06, 0x9f, 0x08, 0xaf, 0x7c, 0x01, 0x1e, 0x63, 0xc6, 0x05, 0x0c, 0x4c, 0xf8, 0x99, 0x88, 0x43,
	0xa3, 0xf4, 0xa7, 0x91, 0xe7, 0x9c, 0xee, 0x90, 0x66, 0x91, 0xf3, 0xf8, 0xbb, 0x4a, 0xfe, 0xaf,
	0x2a, 0x4b, 0x28, 0x8f, 0x48, 0x3f, 0x70, 0x6d, 0x14, 0x0e, 0xb1, 0x01, 0x93, 0x5b, 0xc8, 0x3a,
	0xfb, 0x3e, 0xe4, 0x7a, 0xb8, 0xf8, 0x4c, 0xe0, 0xf2, 0x75, 0x18, 0x4b, 0x9e, 0x20, 0xe3, 0xe8,
	0xc8, 0xdd, 0xaf, 0xa0, 0xb8, 0xde, 0xf4, 0xfe, 0x94, 0xb4, 0xbd, 0x77, 0x34, 0x27, 0x8e, 0xe6,
	0x2f, 0xc7, 0xe1, 0xbb, 0x5f, 0x20, 0xab, 0xb1, 0x1d, 0x39, 0xe2, 0x0d, 0xa7, 0x67, 0x37, 0x5c,
	0x15, 0xe1, 0xf6, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xf5, 0x28, 0xda, 0x0c, 0x02,
	0x00, 0x00,
}
