// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

package paymentpb

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

// UserStatus - user status
type UserStatus int32

const (
	// NORMALUSER - normal user
	UserStatus_NORMALUSER UserStatus = 0
	// CANPAY - can pay
	UserStatus_CANPAY UserStatus = 1
	// CANCOLLECT - can collect
	UserStatus_CANCOLLECT UserStatus = 2
	// frozen
	UserStatus_FROZEN UserStatus = 3
)

var UserStatus_name = map[int32]string{
	0: "NORMALUSER",
	1: "CANPAY",
	2: "CANCOLLECT",
	3: "FROZEN",
}
var UserStatus_value = map[string]int32{
	"NORMALUSER": 0,
	"CANPAY":     1,
	"CANCOLLECT": 2,
	"FROZEN":     3,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{0}
}

// Currency - currency
type Currency int32

const (
	// NONECURRENCY - none
	Currency_NONECURRENCY Currency = 0
	// USD - usd
	Currency_USD Currency = 1
	// EUR - eur
	Currency_EUR Currency = 2
)

var Currency_name = map[int32]string{
	0: "NONECURRENCY",
	1: "USD",
	2: "EUR",
}
var Currency_value = map[string]int32{
	"NONECURRENCY": 0,
	"USD":          1,
	"EUR":          2,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}
func (Currency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{1}
}

type PaymentStatus int32

const (
	// CREATED - The transaction was successfully created
	PaymentStatus_CREATED PaymentStatus = 0
	// APPROVED - The transaction was approved
	PaymentStatus_APPROVED PaymentStatus = 1
	// FAILED - The transaction was failed
	PaymentStatus_FAILED PaymentStatus = 2
)

var PaymentStatus_name = map[int32]string{
	0: "CREATED",
	1: "APPROVED",
	2: "FAILED",
}
var PaymentStatus_value = map[string]int32{
	"CREATED":  0,
	"APPROVED": 1,
	"FAILED":   2,
}

func (x PaymentStatus) String() string {
	return proto.EnumName(PaymentStatus_name, int32(x))
}
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{2}
}

// UserCurrency - user currency
type UserCurrency struct {
	// currencyString - currency string
	CurrencyString string `protobuf:"bytes,1,opt,name=currencyString,proto3" json:"currencyString,omitempty"`
	// balance - balance
	Balance int64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	// currency - currency
	Currency             Currency `protobuf:"varint,3,opt,name=currency,proto3,enum=paymentpb.Currency" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCurrency) Reset()         { *m = UserCurrency{} }
func (m *UserCurrency) String() string { return proto.CompactTextString(m) }
func (*UserCurrency) ProtoMessage()    {}
func (*UserCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{0}
}
func (m *UserCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCurrency.Unmarshal(m, b)
}
func (m *UserCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCurrency.Marshal(b, m, deterministic)
}
func (dst *UserCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCurrency.Merge(dst, src)
}
func (m *UserCurrency) XXX_Size() int {
	return xxx_messageInfo_UserCurrency.Size(m)
}
func (m *UserCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_UserCurrency proto.InternalMessageInfo

func (m *UserCurrency) GetCurrencyString() string {
	if m != nil {
		return m.CurrencyString
	}
	return ""
}

func (m *UserCurrency) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *UserCurrency) GetCurrency() Currency {
	if m != nil {
		return m.Currency
	}
	return Currency_NONECURRENCY
}

// UserCurrencies - user currencies
type UserCurrencies struct {
	// currencies - currencies
	Currencies           map[string]*UserCurrency `protobuf:"bytes,1,rep,name=currencies,proto3" json:"currencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UserCurrencies) Reset()         { *m = UserCurrencies{} }
func (m *UserCurrencies) String() string { return proto.CompactTextString(m) }
func (*UserCurrencies) ProtoMessage()    {}
func (*UserCurrencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{1}
}
func (m *UserCurrencies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCurrencies.Unmarshal(m, b)
}
func (m *UserCurrencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCurrencies.Marshal(b, m, deterministic)
}
func (dst *UserCurrencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCurrencies.Merge(dst, src)
}
func (m *UserCurrencies) XXX_Size() int {
	return xxx_messageInfo_UserCurrencies.Size(m)
}
func (m *UserCurrencies) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCurrencies.DiscardUnknown(m)
}

var xxx_messageInfo_UserCurrencies proto.InternalMessageInfo

func (m *UserCurrencies) GetCurrencies() map[string]*UserCurrency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

// User - user info
type User struct {
	// userID - User ID
	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - User name
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// status - status
	Status UserStatus `protobuf:"varint,3,opt,name=status,proto3,enum=paymentpb.UserStatus" json:"status,omitempty"`
	// registerTime - time of the register
	RegisterTime int64 `protobuf:"varint,4,opt,name=registerTime,proto3" json:"registerTime,omitempty"`
	// currencies - currencies
	UserCurrencies       *UserCurrencies `protobuf:"bytes,100,opt,name=userCurrencies,proto3" json:"userCurrencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{2}
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

func (m *User) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_NORMALUSER
}

func (m *User) GetRegisterTime() int64 {
	if m != nil {
		return m.RegisterTime
	}
	return 0
}

func (m *User) GetUserCurrencies() *UserCurrencies {
	if m != nil {
		return m.UserCurrencies
	}
	return nil
}

// UserPayment - user payment
type UserPayment struct {
	// paymentID - Payment unique identifier
	PaymentID int64 `protobuf:"varint,1,opt,name=paymentID,proto3" json:"paymentID,omitempty"`
	// payer - payer's userID
	Payer int64 `protobuf:"varint,2,opt,name=payer,proto3" json:"payer,omitempty"`
	// payee - payee's userID
	Payee int64 `protobuf:"varint,3,opt,name=payee,proto3" json:"payee,omitempty"`
	// currency - currency
	Currency Currency `protobuf:"varint,4,opt,name=currency,proto3,enum=paymentpb.Currency" json:"currency,omitempty"`
	// amount - amount
	Amount int64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	// status - status
	Status PaymentStatus `protobuf:"varint,6,opt,name=status,proto3,enum=paymentpb.PaymentStatus" json:"status,omitempty"`
	// note - payment note
	Note string `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
	// startBalancePayer - payer's start balance
	StartBalancePayer int64 `protobuf:"varint,8,opt,name=startBalancePayer,proto3" json:"startBalancePayer,omitempty"`
	// endBalancePayer - payer's end balance
	EndBalancePayer int64 `protobuf:"varint,9,opt,name=endBalancePayer,proto3" json:"endBalancePayer,omitempty"`
	// startBalancePayee - payee's start balance
	StartBalancePayee int64 `protobuf:"varint,10,opt,name=startBalancePayee,proto3" json:"startBalancePayee,omitempty"`
	// endBalancePayee - payee's end balance
	EndBalancePayee      int64    `protobuf:"varint,11,opt,name=endBalancePayee,proto3" json:"endBalancePayee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserPayment) Reset()         { *m = UserPayment{} }
func (m *UserPayment) String() string { return proto.CompactTextString(m) }
func (*UserPayment) ProtoMessage()    {}
func (*UserPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{3}
}
func (m *UserPayment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPayment.Unmarshal(m, b)
}
func (m *UserPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPayment.Marshal(b, m, deterministic)
}
func (dst *UserPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPayment.Merge(dst, src)
}
func (m *UserPayment) XXX_Size() int {
	return xxx_messageInfo_UserPayment.Size(m)
}
func (m *UserPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPayment.DiscardUnknown(m)
}

var xxx_messageInfo_UserPayment proto.InternalMessageInfo

func (m *UserPayment) GetPaymentID() int64 {
	if m != nil {
		return m.PaymentID
	}
	return 0
}

func (m *UserPayment) GetPayer() int64 {
	if m != nil {
		return m.Payer
	}
	return 0
}

func (m *UserPayment) GetPayee() int64 {
	if m != nil {
		return m.Payee
	}
	return 0
}

func (m *UserPayment) GetCurrency() Currency {
	if m != nil {
		return m.Currency
	}
	return Currency_NONECURRENCY
}

func (m *UserPayment) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UserPayment) GetStatus() PaymentStatus {
	if m != nil {
		return m.Status
	}
	return PaymentStatus_CREATED
}

func (m *UserPayment) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *UserPayment) GetStartBalancePayer() int64 {
	if m != nil {
		return m.StartBalancePayer
	}
	return 0
}

func (m *UserPayment) GetEndBalancePayer() int64 {
	if m != nil {
		return m.EndBalancePayer
	}
	return 0
}

func (m *UserPayment) GetStartBalancePayee() int64 {
	if m != nil {
		return m.StartBalancePayee
	}
	return 0
}

func (m *UserPayment) GetEndBalancePayee() int64 {
	if m != nil {
		return m.EndBalancePayee
	}
	return 0
}

// UserPayments
type UserPayments struct {
	// totalNums - total nums
	TotalNums int32 `protobuf:"varint,1,opt,name=totalNums,proto3" json:"totalNums,omitempty"`
	// startIndex - current start index
	StartIndex int32 `protobuf:"varint,2,opt,name=startIndex,proto3" json:"startIndex,omitempty"`
	// pageNums - page nums
	PageNums int32 `protobuf:"varint,3,opt,name=pageNums,proto3" json:"pageNums,omitempty"`
	// payments - payments
	Payments             []*UserPayment `protobuf:"bytes,4,rep,name=payments,proto3" json:"payments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserPayments) Reset()         { *m = UserPayments{} }
func (m *UserPayments) String() string { return proto.CompactTextString(m) }
func (*UserPayments) ProtoMessage()    {}
func (*UserPayments) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{4}
}
func (m *UserPayments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPayments.Unmarshal(m, b)
}
func (m *UserPayments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPayments.Marshal(b, m, deterministic)
}
func (dst *UserPayments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPayments.Merge(dst, src)
}
func (m *UserPayments) XXX_Size() int {
	return xxx_messageInfo_UserPayments.Size(m)
}
func (m *UserPayments) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPayments.DiscardUnknown(m)
}

var xxx_messageInfo_UserPayments proto.InternalMessageInfo

func (m *UserPayments) GetTotalNums() int32 {
	if m != nil {
		return m.TotalNums
	}
	return 0
}

func (m *UserPayments) GetStartIndex() int32 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *UserPayments) GetPageNums() int32 {
	if m != nil {
		return m.PageNums
	}
	return 0
}

func (m *UserPayments) GetPayments() []*UserPayment {
	if m != nil {
		return m.Payments
	}
	return nil
}

// UserList
type UserList struct {
	// totalNums - total nums
	TotalNums int32 `protobuf:"varint,1,opt,name=totalNums,proto3" json:"totalNums,omitempty"`
	// startIndex - current start index
	StartIndex int32 `protobuf:"varint,2,opt,name=startIndex,proto3" json:"startIndex,omitempty"`
	// pageNums - page nums
	PageNums int32 `protobuf:"varint,3,opt,name=pageNums,proto3" json:"pageNums,omitempty"`
	// payments - payments
	Users                []*User  `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8fe6b4518a1b0d77, []int{5}
}
func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (dst *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(dst, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetTotalNums() int32 {
	if m != nil {
		return m.TotalNums
	}
	return 0
}

func (m *UserList) GetStartIndex() int32 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *UserList) GetPageNums() int32 {
	if m != nil {
		return m.PageNums
	}
	return 0
}

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UserCurrency)(nil), "paymentpb.UserCurrency")
	proto.RegisterType((*UserCurrencies)(nil), "paymentpb.UserCurrencies")
	proto.RegisterMapType((map[string]*UserCurrency)(nil), "paymentpb.UserCurrencies.CurrenciesEntry")
	proto.RegisterType((*User)(nil), "paymentpb.User")
	proto.RegisterType((*UserPayment)(nil), "paymentpb.UserPayment")
	proto.RegisterType((*UserPayments)(nil), "paymentpb.UserPayments")
	proto.RegisterType((*UserList)(nil), "paymentpb.UserList")
	proto.RegisterEnum("paymentpb.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterEnum("paymentpb.Currency", Currency_name, Currency_value)
	proto.RegisterEnum("paymentpb.PaymentStatus", PaymentStatus_name, PaymentStatus_value)
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor_payment_8fe6b4518a1b0d77) }

var fileDescriptor_payment_8fe6b4518a1b0d77 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xee, 0xc6, 0xf9, 0x9c, 0xa4, 0xa9, 0xdf, 0x79, 0xa1, 0x2c, 0x15, 0x42, 0x91, 0x25, 0x50,
	0xa8, 0x68, 0x40, 0x41, 0x42, 0x88, 0x9b, 0x71, 0x8c, 0x14, 0x29, 0x38, 0xd1, 0xa6, 0xa9, 0x54,
	0x6e, 0x6e, 0xba, 0xaa, 0x22, 0x1a, 0x27, 0xb2, 0xd7, 0x88, 0x1c, 0x39, 0x73, 0xe7, 0xc4, 0xbf,
	0xe0, 0x8f, 0xf0, 0x93, 0xd0, 0xae, 0xd7, 0x4e, 0xec, 0x06, 0x89, 0x0b, 0xb7, 0x99, 0x79, 0x66,
	0xe6, 0x99, 0xaf, 0x5d, 0x38, 0x5c, 0xfb, 0x9b, 0x25, 0x0f, 0x44, 0x6f, 0x1d, 0xae, 0xc4, 0x0a,
	0x1b, 0x5a, 0x5d, 0x5f, 0x59, 0x5f, 0x09, 0xb4, 0x66, 0x11, 0x0f, 0x9d, 0x38, 0x0c, 0x79, 0x30,
	0xdf, 0xe0, 0x53, 0x68, 0xcf, 0xb5, 0x3c, 0x15, 0xe1, 0x22, 0xb8, 0xa1, 0xa4, 0x43, 0xba, 0x0d,
	0x56, 0xb0, 0x22, 0x85, 0xda, 0x95, 0x7f, 0xeb, 0x07, 0x73, 0x4e, 0x4b, 0x1d, 0xd2, 0x35, 0x58,
	0xaa, 0xe2, 0x0b, 0xa8, 0xa7, 0xbe, 0xd4, 0xe8, 0x90, 0x6e, 0xbb, 0xff, 0x7f, 0x2f, 0x23, 0xec,
	0xa5, 0x44, 0x2c, 0x73, 0xb2, 0x7e, 0x12, 0x68, 0xef, 0xd4, 0xb0, 0xe0, 0x11, 0x0e, 0x01, 0xe6,
	0x99, 0x46, 0x49, 0xc7, 0xe8, 0x36, 0xfb, 0xcf, 0x76, 0xb2, 0xe4, 0xdd, 0x7b, 0x5b, 0xd1, 0x0d,
	0x44, 0xb8, 0x61, 0x3b, 0xc1, 0x27, 0x17, 0x70, 0x54, 0x80, 0xd1, 0x04, 0xe3, 0x13, 0xdf, 0xe8,
	0xc6, 0xa4, 0x88, 0x67, 0x50, 0xf9, 0xec, 0xdf, 0xc6, 0x49, 0x2f, 0xcd, 0xfe, 0x83, 0xfd, 0x54,
	0x1b, 0x96, 0x78, 0xbd, 0x2d, 0xbd, 0x21, 0xd6, 0x2f, 0x02, 0x65, 0x89, 0xe1, 0x31, 0x54, 0xe3,
	0x88, 0x87, 0xc3, 0x81, 0x4a, 0x68, 0x30, 0xad, 0xe1, 0x09, 0xd4, 0xa5, 0xe4, 0xf9, 0xcb, 0x24,
	0x6d, 0x83, 0x65, 0x3a, 0x9e, 0x41, 0x35, 0x12, 0xbe, 0x88, 0x23, 0x3d, 0xa1, 0xfb, 0x05, 0xc2,
	0xa9, 0x02, 0x99, 0x76, 0x42, 0x0b, 0x5a, 0x21, 0xbf, 0x59, 0x44, 0x82, 0x87, 0xe7, 0x8b, 0x25,
	0xa7, 0x65, 0x45, 0x94, 0xb3, 0xa1, 0x0d, 0xed, 0x38, 0x37, 0x15, 0x7a, 0xad, 0x7a, 0x79, 0xf8,
	0xc7, 0xb1, 0xb1, 0x42, 0x80, 0xf5, 0xdd, 0x80, 0xa6, 0x74, 0x99, 0x24, 0x01, 0xf8, 0x08, 0xd2,
	0x4b, 0xc9, 0x9a, 0xdb, 0x1a, 0xf0, 0x1e, 0x54, 0xd6, 0xfe, 0x86, 0x87, 0x7a, 0xff, 0x89, 0x92,
	0x5a, 0xb9, 0x6a, 0x4c, 0x5b, 0xf3, 0x37, 0x51, 0xfe, 0x8b, 0x9b, 0x90, 0x43, 0xf5, 0x97, 0xab,
	0x38, 0x10, 0xb4, 0x92, 0x0c, 0x35, 0xd1, 0xf0, 0x65, 0x36, 0xb8, 0xaa, 0x4a, 0x43, 0x77, 0xd2,
	0xe8, 0xb2, 0x0b, 0xb3, 0x43, 0x28, 0x07, 0x2b, 0xc1, 0x69, 0x4d, 0xad, 0x40, 0xc9, 0xf8, 0x1c,
	0xfe, 0x8b, 0x84, 0x1f, 0x8a, 0x77, 0xc9, 0xc9, 0x4e, 0x54, 0x1b, 0x75, 0x45, 0x74, 0x17, 0xc0,
	0x2e, 0x1c, 0xf1, 0xe0, 0x3a, 0xe7, 0xdb, 0x50, 0xbe, 0x45, 0xf3, 0xbe, 0xbc, 0x9c, 0xc2, 0xfe,
	0xbc, 0xfc, 0x6e, 0x5e, 0x4e, 0x9b, 0xfb, 0xf2, 0x72, 0xeb, 0x87, 0x7e, 0xa5, 0xba, 0xc3, 0x48,
	0x6e, 0x46, 0xac, 0x84, 0x7f, 0xeb, 0xc5, 0xcb, 0x48, 0x6d, 0xa6, 0xc2, 0xb6, 0x06, 0x7c, 0x0c,
	0xa0, 0xd8, 0x86, 0xc1, 0x35, 0xff, 0xa2, 0xd6, 0x53, 0x61, 0x3b, 0x16, 0x79, 0x99, 0x6b, 0xff,
	0x86, 0xab, 0x60, 0x43, 0xa1, 0x99, 0x8e, 0x7d, 0x89, 0x25, 0x2c, 0xb4, 0xac, 0xde, 0xdd, 0x71,
	0xe1, 0x80, 0x74, 0x11, 0x2c, 0xf3, 0xb3, 0xbe, 0x11, 0xa8, 0x4b, 0x64, 0xb4, 0x88, 0xc4, 0x3f,
	0x2c, 0xed, 0x09, 0x54, 0xe4, 0xc1, 0xa6, 0x75, 0x1d, 0x15, 0xea, 0x62, 0x09, 0x7a, 0x3a, 0x00,
	0xd8, 0x3e, 0x21, 0x6c, 0x03, 0x78, 0x63, 0xf6, 0xc1, 0x1e, 0xcd, 0xa6, 0x2e, 0x33, 0x0f, 0x10,
	0xa0, 0xea, 0xd8, 0xde, 0xc4, 0xbe, 0x34, 0x89, 0xc4, 0x1c, 0xdb, 0x73, 0xc6, 0xa3, 0x91, 0xeb,
	0x9c, 0x9b, 0x25, 0x89, 0xbd, 0x67, 0xe3, 0x8f, 0xae, 0x67, 0x1a, 0xa7, 0x3d, 0xa8, 0x67, 0x7f,
	0xa2, 0x09, 0x2d, 0x6f, 0xec, 0xb9, 0xce, 0x8c, 0x31, 0xd7, 0x73, 0x2e, 0xcd, 0x03, 0xac, 0x81,
	0x31, 0x9b, 0x0e, 0x4c, 0x22, 0x05, 0x77, 0xc6, 0xcc, 0xd2, 0xe9, 0x6b, 0x38, 0xcc, 0xdd, 0x1f,
	0x36, 0xa1, 0xe6, 0x30, 0xd7, 0x3e, 0x77, 0x07, 0xe6, 0x01, 0xb6, 0xa0, 0x6e, 0x4f, 0x26, 0x6c,
	0x7c, 0xe1, 0xca, 0x20, 0xc9, 0x63, 0x0f, 0x47, 0xee, 0xc0, 0x2c, 0x5d, 0x55, 0xd5, 0x97, 0xfc,
	0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0xa9, 0xc9, 0x46, 0xa3, 0x05, 0x00, 0x00,
}
