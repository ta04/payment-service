// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

package com_ta04_srv_payment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Payment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int32    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentMethodId      int32    `protobuf:"varint,3,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{0}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Payment) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Payment) GetPaymentMethodId() int32 {
	if m != nil {
		return m.PaymentMethodId
	}
	return 0
}

func (m *Payment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Order struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId            int32    `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	UserId               int32    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *Order) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetAllPaymentsRequest struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPaymentsRequest) Reset()         { *m = GetAllPaymentsRequest{} }
func (m *GetAllPaymentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllPaymentsRequest) ProtoMessage()    {}
func (*GetAllPaymentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{2}
}

func (m *GetAllPaymentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPaymentsRequest.Unmarshal(m, b)
}
func (m *GetAllPaymentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPaymentsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllPaymentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPaymentsRequest.Merge(m, src)
}
func (m *GetAllPaymentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllPaymentsRequest.Size(m)
}
func (m *GetAllPaymentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPaymentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPaymentsRequest proto.InternalMessageInfo

func (m *GetAllPaymentsRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetOnePaymentRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId              int32    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOnePaymentRequest) Reset()         { *m = GetOnePaymentRequest{} }
func (m *GetOnePaymentRequest) String() string { return proto.CompactTextString(m) }
func (*GetOnePaymentRequest) ProtoMessage()    {}
func (*GetOnePaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{3}
}

func (m *GetOnePaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOnePaymentRequest.Unmarshal(m, b)
}
func (m *GetOnePaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOnePaymentRequest.Marshal(b, m, deterministic)
}
func (m *GetOnePaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOnePaymentRequest.Merge(m, src)
}
func (m *GetOnePaymentRequest) XXX_Size() int {
	return xxx_messageInfo_GetOnePaymentRequest.Size(m)
}
func (m *GetOnePaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOnePaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOnePaymentRequest proto.InternalMessageInfo

func (m *GetOnePaymentRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetOnePaymentRequest) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type Response struct {
	Payment              *Payment   `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	Error                *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Payments             []*Payment `protobuf:"bytes,3,rep,name=payments,proto3" json:"payments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetPayment() *Payment {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Payment)(nil), "com.ta04.srv.payment.Payment")
	proto.RegisterType((*Order)(nil), "com.ta04.srv.payment.Order")
	proto.RegisterType((*GetAllPaymentsRequest)(nil), "com.ta04.srv.payment.GetAllPaymentsRequest")
	proto.RegisterType((*GetOnePaymentRequest)(nil), "com.ta04.srv.payment.GetOnePaymentRequest")
	proto.RegisterType((*Response)(nil), "com.ta04.srv.payment.Response")
	proto.RegisterType((*Error)(nil), "com.ta04.srv.payment.Error")
}

func init() {
	proto.RegisterFile("payment.proto", fileDescriptor_6362648dfa63d410)
}

var fileDescriptor_6362648dfa63d410 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0xaf, 0xd2, 0x40,
	0x10, 0xb7, 0x85, 0x52, 0x18, 0x02, 0xea, 0x04, 0xb5, 0x6a, 0x30, 0xa4, 0x27, 0x82, 0x49, 0x55,
	0xd4, 0x18, 0x8f, 0xc4, 0x18, 0xc2, 0xc1, 0x60, 0x4a, 0x3c, 0x1a, 0x52, 0xbb, 0x13, 0x6c, 0x42,
	0xd9, 0xba, 0xbb, 0x25, 0xfa, 0xb1, 0xfc, 0x7a, 0x9e, 0x4c, 0xb7, 0xdb, 0xf7, 0x20, 0xf4, 0xf1,
	0x38, 0xbc, 0xdb, 0xce, 0xec, 0xef, 0xcf, 0xec, 0x6f, 0xb2, 0xd0, 0xcb, 0xa2, 0x3f, 0x29, 0xed,
	0x54, 0x90, 0x09, 0xae, 0x38, 0x0e, 0x62, 0x9e, 0x06, 0x2a, 0x7a, 0xfd, 0x2e, 0x90, 0x62, 0x1f,
	0x98, 0x3b, 0xff, 0x37, 0xb8, 0x5f, 0xcb, 0x23, 0xf6, 0xc1, 0x4e, 0x98, 0x67, 0x8d, 0xac, 0xb1,
	0x13, 0xda, 0x09, 0xc3, 0xa7, 0xd0, 0xe6, 0x82, 0x91, 0x58, 0x27, 0xcc, 0xb3, 0x75, 0xd7, 0xd5,
	0xf5, 0x82, 0xe1, 0x04, 0x1e, 0x1a, 0x81, 0x75, 0x4a, 0xea, 0x27, 0x67, 0x05, 0xa6, 0xa1, 0x31,
	0xf7, 0xcd, 0xc5, 0x17, 0xdd, 0x5f, 0x30, 0x7c, 0x0c, 0x2d, 0xa9, 0x22, 0x95, 0x4b, 0xaf, 0x39,
	0xb2, 0xc6, 0x9d, 0xd0, 0x54, 0xfe, 0x06, 0x9c, 0x65, 0x21, 0x77, 0xe2, 0x3b, 0x04, 0xc8, 0x04,
	0x67, 0x79, 0xac, 0xae, 0x9d, 0x3b, 0xa6, 0xb3, 0x60, 0xf8, 0x04, 0xdc, 0x5c, 0x96, 0x53, 0x95,
	0x8e, 0xad, 0xa2, 0x3c, 0x63, 0xf4, 0x0a, 0x1e, 0xcd, 0x49, 0xcd, 0xb6, 0x5b, 0xf3, 0x50, 0x19,
	0xd2, 0xaf, 0x9c, 0xa4, 0x3a, 0x20, 0x58, 0x47, 0x84, 0x19, 0x0c, 0xe6, 0xa4, 0x96, 0x3b, 0x32,
	0x84, 0x0a, 0x7f, 0x79, 0x40, 0xfe, 0x5f, 0x0b, 0xda, 0x21, 0xc9, 0x8c, 0xef, 0x24, 0xe1, 0x07,
	0x70, 0x4d, 0x28, 0x9a, 0xdc, 0x9d, 0x0e, 0x83, 0xba, 0x5d, 0x04, 0x95, 0x5d, 0x85, 0xc6, 0x37,
	0xe0, 0x90, 0x10, 0x5c, 0x68, 0xf5, 0xee, 0xf4, 0x79, 0x3d, 0xed, 0x73, 0x01, 0x09, 0x4b, 0x24,
	0x7e, 0x84, 0xb6, 0xe9, 0x4b, 0xaf, 0x31, 0x6a, 0xdc, 0x6e, 0x76, 0x05, 0xf7, 0xdf, 0x83, 0xa3,
	0xa5, 0x10, 0xa1, 0x19, 0x73, 0x46, 0xe6, 0xa5, 0xfa, 0x8c, 0x1e, 0xb8, 0x29, 0x49, 0x19, 0x6d,
	0x48, 0x0f, 0xd3, 0x09, 0xab, 0x72, 0xfa, 0xcf, 0x86, 0xbe, 0x11, 0x5b, 0x91, 0xd8, 0x27, 0x31,
	0xe1, 0x1a, 0xfa, 0xc7, 0x89, 0xe3, 0xcb, 0xfa, 0x21, 0x6a, 0xf7, 0xf2, 0xec, 0x45, 0x3d, 0xb8,
	0xca, 0xd3, 0xbf, 0x87, 0xdf, 0xa1, 0x77, 0xb4, 0x21, 0x9c, 0xdc, 0xa8, 0x7f, 0xb2, 0xc6, 0x0b,
	0xe4, 0x57, 0xf0, 0xe0, 0x93, 0xa0, 0x48, 0xd1, 0x81, 0xc3, 0xf9, 0x18, 0x2f, 0x13, 0xfd, 0x96,
	0xb1, 0xbb, 0x15, 0xfd, 0xd1, 0xd2, 0x7f, 0xfb, 0xed, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f,
	0xf4, 0x32, 0xc2, 0xec, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PaymentService service

type PaymentServiceClient interface {
	GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest, opts ...client.CallOption) (*Response, error)
	GetOnePayment(ctx context.Context, in *GetOnePaymentRequest, opts ...client.CallOption) (*Response, error)
	CreateOnePayment(ctx context.Context, in *Payment, opts ...client.CallOption) (*Response, error)
	UpdateOnePayment(ctx context.Context, in *Payment, opts ...client.CallOption) (*Response, error)
}

type paymentServiceClient struct {
	c           client.Client
	serviceName string
}

func NewPaymentServiceClient(serviceName string, c client.Client) PaymentServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.ta04.srv.payment"
	}
	return &paymentServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *paymentServiceClient) GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PaymentService.GetAllPayments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetOnePayment(ctx context.Context, in *GetOnePaymentRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PaymentService.GetOnePayment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CreateOnePayment(ctx context.Context, in *Payment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PaymentService.CreateOnePayment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) UpdateOnePayment(ctx context.Context, in *Payment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "PaymentService.UpdateOnePayment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentService service

type PaymentServiceHandler interface {
	GetAllPayments(context.Context, *GetAllPaymentsRequest, *Response) error
	GetOnePayment(context.Context, *GetOnePaymentRequest, *Response) error
	CreateOnePayment(context.Context, *Payment, *Response) error
	UpdateOnePayment(context.Context, *Payment, *Response) error
}

func RegisterPaymentServiceHandler(s server.Server, hdlr PaymentServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PaymentService{hdlr}, opts...))
}

type PaymentService struct {
	PaymentServiceHandler
}

func (h *PaymentService) GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest, out *Response) error {
	return h.PaymentServiceHandler.GetAllPayments(ctx, in, out)
}

func (h *PaymentService) GetOnePayment(ctx context.Context, in *GetOnePaymentRequest, out *Response) error {
	return h.PaymentServiceHandler.GetOnePayment(ctx, in, out)
}

func (h *PaymentService) CreateOnePayment(ctx context.Context, in *Payment, out *Response) error {
	return h.PaymentServiceHandler.CreateOnePayment(ctx, in, out)
}

func (h *PaymentService) UpdateOnePayment(ctx context.Context, in *Payment, out *Response) error {
	return h.PaymentServiceHandler.UpdateOnePayment(ctx, in, out)
}