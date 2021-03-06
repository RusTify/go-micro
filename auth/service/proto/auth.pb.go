// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/service/proto/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Token struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64             `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expiry               int64             `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Subject              string            `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Roles                []string          `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Namespace            string            `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Token) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Token) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *Token) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Token) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Token) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Token) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type Account struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret               *Token            `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Roles                []string          `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Namespace            string            `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{1}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetSecret() *Token {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *Account) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Account) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Account) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type Resource struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Endpoint             string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type GenerateRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Roles                []string          `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SecretExpiry         int64             `protobuf:"varint,4,opt,name=secret_expiry,json=secretExpiry,proto3" json:"secret_expiry,omitempty"`
	Namespace            string            `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{3}
}

func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(m, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GenerateRequest) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *GenerateRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GenerateRequest) GetSecretExpiry() int64 {
	if m != nil {
		return m.SecretExpiry
	}
	return 0
}

func (m *GenerateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type GenerateResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{4}
}

func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (m *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(m, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

func (m *GenerateResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type GrantRequest struct {
	Role                 string    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GrantRequest) Reset()         { *m = GrantRequest{} }
func (m *GrantRequest) String() string { return proto.CompactTextString(m) }
func (*GrantRequest) ProtoMessage()    {}
func (*GrantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{5}
}

func (m *GrantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrantRequest.Unmarshal(m, b)
}
func (m *GrantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrantRequest.Marshal(b, m, deterministic)
}
func (m *GrantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantRequest.Merge(m, src)
}
func (m *GrantRequest) XXX_Size() int {
	return xxx_messageInfo_GrantRequest.Size(m)
}
func (m *GrantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrantRequest proto.InternalMessageInfo

func (m *GrantRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GrantRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type GrantResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrantResponse) Reset()         { *m = GrantResponse{} }
func (m *GrantResponse) String() string { return proto.CompactTextString(m) }
func (*GrantResponse) ProtoMessage()    {}
func (*GrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{6}
}

func (m *GrantResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrantResponse.Unmarshal(m, b)
}
func (m *GrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrantResponse.Marshal(b, m, deterministic)
}
func (m *GrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantResponse.Merge(m, src)
}
func (m *GrantResponse) XXX_Size() int {
	return xxx_messageInfo_GrantResponse.Size(m)
}
func (m *GrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrantResponse proto.InternalMessageInfo

type RevokeRequest struct {
	Role                 string    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RevokeRequest) Reset()         { *m = RevokeRequest{} }
func (m *RevokeRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeRequest) ProtoMessage()    {}
func (*RevokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{7}
}

func (m *RevokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeRequest.Unmarshal(m, b)
}
func (m *RevokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeRequest.Marshal(b, m, deterministic)
}
func (m *RevokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeRequest.Merge(m, src)
}
func (m *RevokeRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeRequest.Size(m)
}
func (m *RevokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeRequest proto.InternalMessageInfo

func (m *RevokeRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RevokeRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type RevokeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeResponse) Reset()         { *m = RevokeResponse{} }
func (m *RevokeResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeResponse) ProtoMessage()    {}
func (*RevokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{8}
}

func (m *RevokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeResponse.Unmarshal(m, b)
}
func (m *RevokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeResponse.Marshal(b, m, deterministic)
}
func (m *RevokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeResponse.Merge(m, src)
}
func (m *RevokeResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeResponse.Size(m)
}
func (m *RevokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeResponse proto.InternalMessageInfo

type InspectRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectRequest) Reset()         { *m = InspectRequest{} }
func (m *InspectRequest) String() string { return proto.CompactTextString(m) }
func (*InspectRequest) ProtoMessage()    {}
func (*InspectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{9}
}

func (m *InspectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectRequest.Unmarshal(m, b)
}
func (m *InspectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectRequest.Marshal(b, m, deterministic)
}
func (m *InspectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectRequest.Merge(m, src)
}
func (m *InspectRequest) XXX_Size() int {
	return xxx_messageInfo_InspectRequest.Size(m)
}
func (m *InspectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InspectRequest proto.InternalMessageInfo

func (m *InspectRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type InspectResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectResponse) Reset()         { *m = InspectResponse{} }
func (m *InspectResponse) String() string { return proto.CompactTextString(m) }
func (*InspectResponse) ProtoMessage()    {}
func (*InspectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{10}
}

func (m *InspectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectResponse.Unmarshal(m, b)
}
func (m *InspectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectResponse.Marshal(b, m, deterministic)
}
func (m *InspectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectResponse.Merge(m, src)
}
func (m *InspectResponse) XXX_Size() int {
	return xxx_messageInfo_InspectResponse.Size(m)
}
func (m *InspectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InspectResponse proto.InternalMessageInfo

func (m *InspectResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type RefreshRequest struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	TokenExpiry          int64    `protobuf:"varint,2,opt,name=token_expiry,json=tokenExpiry,proto3" json:"token_expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{11}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *RefreshRequest) GetTokenExpiry() int64 {
	if m != nil {
		return m.TokenExpiry
	}
	return 0
}

type RefreshResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshResponse) Reset()         { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()    {}
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{12}
}

func (m *RefreshResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshResponse.Unmarshal(m, b)
}
func (m *RefreshResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshResponse.Marshal(b, m, deterministic)
}
func (m *RefreshResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshResponse.Merge(m, src)
}
func (m *RefreshResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshResponse.Size(m)
}
func (m *RefreshResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshResponse proto.InternalMessageInfo

func (m *RefreshResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "go.micro.auth.Token")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.Token.MetadataEntry")
	proto.RegisterType((*Account)(nil), "go.micro.auth.Account")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.Account.MetadataEntry")
	proto.RegisterType((*Resource)(nil), "go.micro.auth.Resource")
	proto.RegisterType((*GenerateRequest)(nil), "go.micro.auth.GenerateRequest")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.GenerateRequest.MetadataEntry")
	proto.RegisterType((*GenerateResponse)(nil), "go.micro.auth.GenerateResponse")
	proto.RegisterType((*GrantRequest)(nil), "go.micro.auth.GrantRequest")
	proto.RegisterType((*GrantResponse)(nil), "go.micro.auth.GrantResponse")
	proto.RegisterType((*RevokeRequest)(nil), "go.micro.auth.RevokeRequest")
	proto.RegisterType((*RevokeResponse)(nil), "go.micro.auth.RevokeResponse")
	proto.RegisterType((*InspectRequest)(nil), "go.micro.auth.InspectRequest")
	proto.RegisterType((*InspectResponse)(nil), "go.micro.auth.InspectResponse")
	proto.RegisterType((*RefreshRequest)(nil), "go.micro.auth.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "go.micro.auth.RefreshResponse")
}

func init() { proto.RegisterFile("auth/service/proto/auth.proto", fileDescriptor_21300bfacc51fc2a) }

var fileDescriptor_21300bfacc51fc2a = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xae, 0xed, 0x7c, 0x38, 0x93, 0xa6, 0x89, 0x56, 0x55, 0x5e, 0x2b, 0x7a, 0x5b, 0x82, 0x41,
	0x28, 0x42, 0x95, 0x83, 0xd2, 0x0b, 0xa2, 0x02, 0x51, 0xa0, 0x2a, 0x1f, 0x2a, 0x07, 0x0b, 0x09,
	0xc4, 0x05, 0xb9, 0xce, 0x40, 0x4c, 0x1a, 0xdb, 0xac, 0xd7, 0x11, 0x39, 0xf2, 0x87, 0xf8, 0x67,
	0xfc, 0x01, 0x4e, 0x68, 0xd7, 0xbb, 0x8e, 0xe3, 0x24, 0x08, 0xa1, 0x72, 0x89, 0x66, 0x66, 0x67,
	0x9e, 0x99, 0x79, 0xf6, 0xf1, 0x06, 0x0e, 0xbc, 0x94, 0x4d, 0x86, 0x09, 0xd2, 0x79, 0xe0, 0xe3,
	0x30, 0xa6, 0x11, 0x8b, 0x86, 0x3c, 0xe4, 0x08, 0x93, 0xb4, 0x3e, 0x45, 0xce, 0x2c, 0xf0, 0x69,
	0xe4, 0xf0, 0xa0, 0xfd, 0x5d, 0x87, 0xea, 0x9b, 0x68, 0x8a, 0x21, 0xd9, 0x87, 0x2a, 0xe3, 0x86,
	0xa5, 0xf5, 0xb5, 0x41, 0xc3, 0xcd, 0x1c, 0x42, 0xa0, 0xc2, 0x16, 0x31, 0x5a, 0xba, 0x08, 0x0a,
	0x9b, 0x58, 0x50, 0xf7, 0x29, 0x7a, 0x0c, 0xc7, 0x96, 0xd1, 0xd7, 0x06, 0x86, 0xab, 0x5c, 0xd2,
	0x85, 0x1a, 0x7e, 0x8d, 0x03, 0xba, 0xb0, 0x2a, 0xe2, 0x40, 0x7a, 0xbc, 0x22, 0x49, 0x2f, 0x3f,
	0xa3, 0xcf, 0xac, 0xaa, 0x00, 0x52, 0x2e, 0xef, 0x4a, 0xa3, 0x2b, 0x4c, 0xac, 0x5a, 0xdf, 0xe0,
	0x5d, 0x85, 0x43, 0x1e, 0x81, 0x39, 0x43, 0xe6, 0x8d, 0x3d, 0xe6, 0x59, 0xf5, 0xbe, 0x31, 0x68,
	0x8e, 0x6c, 0x67, 0x65, 0x6e, 0x47, 0xcc, 0xec, 0x5c, 0xc8, 0xa4, 0xb3, 0x90, 0xd1, 0x85, 0x9b,
	0xd7, 0x90, 0xff, 0xa1, 0x11, 0x7a, 0x33, 0x4c, 0x62, 0xcf, 0x47, 0xcb, 0x14, 0x1d, 0x97, 0x81,
	0xde, 0x09, 0xb4, 0x56, 0x0a, 0x49, 0x07, 0x8c, 0x29, 0x2e, 0xe4, 0xe2, 0xdc, 0xe4, 0x63, 0xcd,
	0xbd, 0xab, 0x54, 0xed, 0x9d, 0x39, 0x0f, 0xf4, 0xfb, 0x9a, 0xfd, 0x53, 0x83, 0xfa, 0xa9, 0xef,
	0x47, 0x69, 0xc8, 0xc8, 0x1e, 0xe8, 0xc1, 0x58, 0x96, 0xe9, 0xc1, 0x98, 0x1c, 0x41, 0x2d, 0x41,
	0x9f, 0x22, 0x13, 0x65, 0xcd, 0xd1, 0xfe, 0xa6, 0xa1, 0x5d, 0x99, 0xb3, 0x5c, 0xdd, 0x28, 0xae,
	0xfe, 0xb8, 0xb0, 0x7a, 0x45, 0xac, 0x7e, 0xbb, 0x84, 0x22, 0xbb, 0xff, 0xd9, 0xf2, 0xd5, 0x6b,
	0x5d, 0xfe, 0x35, 0x98, 0x2e, 0x26, 0x51, 0x4a, 0x7d, 0xe4, 0xca, 0xe0, 0xa8, 0xb2, 0x50, 0xd8,
	0x1b, 0xd5, 0xd2, 0x03, 0x13, 0xc3, 0x71, 0x1c, 0x05, 0x21, 0x13, 0x72, 0x69, 0xb8, 0xb9, 0x6f,
	0x7f, 0xd3, 0xa1, 0x7d, 0x8e, 0x21, 0x52, 0x8f, 0xa1, 0x8b, 0x5f, 0x52, 0x4c, 0xd6, 0x49, 0xcd,
	0x69, 0xd2, 0x8b, 0x34, 0x3d, 0x2f, 0xd0, 0x64, 0x08, 0x9a, 0x8e, 0x4a, 0x34, 0x95, 0x70, 0xb7,
	0xd2, 0x75, 0x0b, 0x5a, 0xd9, 0x85, 0x7c, 0x58, 0x91, 0xee, 0x6e, 0x16, 0x3c, 0xcb, 0x04, 0xfc,
	0x0f, 0x39, 0x7d, 0x06, 0x9d, 0xe5, 0xa8, 0x49, 0x1c, 0x85, 0x09, 0x92, 0x7b, 0x50, 0xf7, 0xb2,
	0x5b, 0x16, 0x18, 0xcd, 0x51, 0x77, 0xb3, 0x06, 0x5c, 0x95, 0x66, 0xbf, 0x85, 0xdd, 0x73, 0xea,
	0x85, 0x4c, 0xb1, 0x48, 0xa0, 0xc2, 0x89, 0x52, 0xb7, 0xc3, 0x6d, 0x72, 0x0c, 0x26, 0x95, 0xb7,
	0x27, 0x05, 0xfa, 0x5f, 0x09, 0x56, 0x5d, 0xae, 0x9b, 0x27, 0xda, 0x6d, 0x68, 0x49, 0xe0, 0x6c,
	0x36, 0xfb, 0x1d, 0xb4, 0x5c, 0x9c, 0x47, 0x53, 0xbc, 0xf6, 0x56, 0x1d, 0xd8, 0x53, 0xc8, 0xb2,
	0xd7, 0x1d, 0xd8, 0x7b, 0x11, 0x26, 0x31, 0xfa, 0xf9, 0x5e, 0x1b, 0x5f, 0x29, 0xfb, 0x29, 0xb4,
	0xf3, 0xbc, 0xbf, 0xa6, 0xf0, 0x15, 0x6f, 0xff, 0x91, 0x62, 0x32, 0x51, 0xcd, 0xba, 0xf9, 0xf7,
	0x9c, 0x75, 0x53, 0x5f, 0xee, 0x4d, 0xd8, 0x15, 0x7d, 0x95, 0x62, 0x74, 0xa1, 0x98, 0xa6, 0x88,
	0x65, 0x82, 0xb1, 0x1f, 0x42, 0x3b, 0x07, 0x93, 0x13, 0xdd, 0x2d, 0x8e, 0xbe, 0xed, 0x71, 0xc8,
	0x52, 0x46, 0x3f, 0x34, 0xa8, 0x9c, 0xa6, 0x6c, 0x42, 0x2e, 0xc0, 0x54, 0xea, 0x20, 0x87, 0xbf,
	0x57, 0x78, 0xef, 0xc6, 0xd6, 0x73, 0x49, 0xe7, 0x0e, 0x79, 0x09, 0x75, 0x49, 0x14, 0x39, 0x28,
	0x65, 0xaf, 0x12, 0xdd, 0x3b, 0xdc, 0x76, 0x5c, 0xc4, 0x92, 0x2b, 0xae, 0x61, 0xad, 0xf2, 0xb8,
	0x86, 0x55, 0x62, 0xc6, 0xde, 0x79, 0xd2, 0x7c, 0xdf, 0xe0, 0x27, 0x27, 0xfc, 0xe7, 0xb2, 0x26,
	0xfe, 0xa9, 0x8e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x52, 0x2c, 0xfc, 0x9c, 0xca, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	Inspect(ctx context.Context, in *InspectRequest, opts ...grpc.CallOption) (*InspectResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/go.micro.auth.Auth/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Inspect(ctx context.Context, in *InspectRequest, opts ...grpc.CallOption) (*InspectResponse, error) {
	out := new(InspectResponse)
	err := c.cc.Invoke(ctx, "/go.micro.auth.Auth/Inspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/go.micro.auth.Auth/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	Inspect(context.Context, *InspectRequest) (*InspectResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) Generate(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (*UnimplementedAuthServer) Inspect(ctx context.Context, req *InspectRequest) (*InspectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inspect not implemented")
}
func (*UnimplementedAuthServer) Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.auth.Auth/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Inspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InspectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Inspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.auth.Auth/Inspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Inspect(ctx, req.(*InspectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.auth.Auth/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Auth_Generate_Handler,
		},
		{
			MethodName: "Inspect",
			Handler:    _Auth_Inspect_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Auth_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/service/proto/auth.proto",
}
