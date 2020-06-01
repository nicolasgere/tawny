// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Admin.proto

package tawny

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ChannelConfiguration struct {
	//* Define if a message can be push without admin api key. Common use case: Only allow server to push message
	AdminRequiredToPush  bool     `protobuf:"varint,1,opt,name=admin_required_to_push,json=adminRequiredToPush,proto3" json:"admin_required_to_push,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelConfiguration) Reset()         { *m = ChannelConfiguration{} }
func (m *ChannelConfiguration) String() string { return proto.CompactTextString(m) }
func (*ChannelConfiguration) ProtoMessage()    {}
func (*ChannelConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da8bee81a31f1f5, []int{0}
}

func (m *ChannelConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelConfiguration.Unmarshal(m, b)
}
func (m *ChannelConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelConfiguration.Marshal(b, m, deterministic)
}
func (m *ChannelConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelConfiguration.Merge(m, src)
}
func (m *ChannelConfiguration) XXX_Size() int {
	return xxx_messageInfo_ChannelConfiguration.Size(m)
}
func (m *ChannelConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelConfiguration proto.InternalMessageInfo

func (m *ChannelConfiguration) GetAdminRequiredToPush() bool {
	if m != nil {
		return m.AdminRequiredToPush
	}
	return false
}

type CreateOrUpdateChannelInput struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Configuration        *ChannelConfiguration `protobuf:"bytes,2,opt,name=configuration,proto3" json:"configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateOrUpdateChannelInput) Reset()         { *m = CreateOrUpdateChannelInput{} }
func (m *CreateOrUpdateChannelInput) String() string { return proto.CompactTextString(m) }
func (*CreateOrUpdateChannelInput) ProtoMessage()    {}
func (*CreateOrUpdateChannelInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da8bee81a31f1f5, []int{1}
}

func (m *CreateOrUpdateChannelInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrUpdateChannelInput.Unmarshal(m, b)
}
func (m *CreateOrUpdateChannelInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrUpdateChannelInput.Marshal(b, m, deterministic)
}
func (m *CreateOrUpdateChannelInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrUpdateChannelInput.Merge(m, src)
}
func (m *CreateOrUpdateChannelInput) XXX_Size() int {
	return xxx_messageInfo_CreateOrUpdateChannelInput.Size(m)
}
func (m *CreateOrUpdateChannelInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrUpdateChannelInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrUpdateChannelInput proto.InternalMessageInfo

func (m *CreateOrUpdateChannelInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrUpdateChannelInput) GetConfiguration() *ChannelConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type Channel struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Configuration        *ChannelConfiguration `protobuf:"bytes,2,opt,name=configuration,proto3" json:"configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da8bee81a31f1f5, []int{2}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Channel) GetConfiguration() *ChannelConfiguration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

type ListChannelOutput struct {
	Channels             []*Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListChannelOutput) Reset()         { *m = ListChannelOutput{} }
func (m *ListChannelOutput) String() string { return proto.CompactTextString(m) }
func (*ListChannelOutput) ProtoMessage()    {}
func (*ListChannelOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da8bee81a31f1f5, []int{3}
}

func (m *ListChannelOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListChannelOutput.Unmarshal(m, b)
}
func (m *ListChannelOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListChannelOutput.Marshal(b, m, deterministic)
}
func (m *ListChannelOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListChannelOutput.Merge(m, src)
}
func (m *ListChannelOutput) XXX_Size() int {
	return xxx_messageInfo_ListChannelOutput.Size(m)
}
func (m *ListChannelOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ListChannelOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ListChannelOutput proto.InternalMessageInfo

func (m *ListChannelOutput) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

type DeleteChannelInput struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteChannelInput) Reset()         { *m = DeleteChannelInput{} }
func (m *DeleteChannelInput) String() string { return proto.CompactTextString(m) }
func (*DeleteChannelInput) ProtoMessage()    {}
func (*DeleteChannelInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da8bee81a31f1f5, []int{4}
}

func (m *DeleteChannelInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteChannelInput.Unmarshal(m, b)
}
func (m *DeleteChannelInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteChannelInput.Marshal(b, m, deterministic)
}
func (m *DeleteChannelInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteChannelInput.Merge(m, src)
}
func (m *DeleteChannelInput) XXX_Size() int {
	return xxx_messageInfo_DeleteChannelInput.Size(m)
}
func (m *DeleteChannelInput) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteChannelInput.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteChannelInput proto.InternalMessageInfo

func (m *DeleteChannelInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ChannelConfiguration)(nil), "tawny.ChannelConfiguration")
	proto.RegisterType((*CreateOrUpdateChannelInput)(nil), "tawny.CreateOrUpdateChannelInput")
	proto.RegisterType((*Channel)(nil), "tawny.Channel")
	proto.RegisterType((*ListChannelOutput)(nil), "tawny.ListChannelOutput")
	proto.RegisterType((*DeleteChannelInput)(nil), "tawny.DeleteChannelInput")
}

func init() {
	proto.RegisterFile("Admin.proto", fileDescriptor_5da8bee81a31f1f5)
}

var fileDescriptor_5da8bee81a31f1f5 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x51, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0x7e, 0xe2, 0x54, 0x4c, 0x1c, 0x95, 0xd4, 0x72, 0xc1, 0x9e, 0x1a, 0x0f, 0x25, 0x81,
	0x1f, 0x60, 0x08, 0x78, 0x30, 0x9a, 0x60, 0x2a, 0x9e, 0x71, 0x81, 0x01, 0x9a, 0xc0, 0x6e, 0xdd,
	0xee, 0x6a, 0xf8, 0xd1, 0xfe, 0x07, 0xc3, 0x76, 0x31, 0x10, 0x6c, 0x3c, 0x79, 0xdb, 0xcc, 0xbc,
	0x7d, 0x6f, 0xde, 0x7b, 0xe0, 0xb6, 0xc7, 0x8b, 0x84, 0x47, 0xa9, 0x14, 0x4a, 0xe0, 0xa1, 0x62,
	0x9f, 0x7c, 0xe9, 0xd7, 0xa6, 0x42, 0x4c, 0xe7, 0xd4, 0x30, 0xc3, 0xa1, 0x9e, 0x34, 0x68, 0x91,
	0xaa, 0x65, 0x8e, 0x09, 0x1e, 0xe1, 0xb2, 0x33, 0x63, 0x9c, 0xd3, 0xbc, 0x23, 0xf8, 0x24, 0x99,
	0x6a, 0xc9, 0x54, 0x22, 0x38, 0xb6, 0xa0, 0xca, 0x56, 0x54, 0x03, 0x49, 0xef, 0x3a, 0x91, 0x34,
	0x1e, 0x28, 0x31, 0x48, 0x75, 0x36, 0xf3, 0x9c, 0xba, 0x13, 0x96, 0xe3, 0x0b, 0xb3, 0x8d, 0xed,
	0xb2, 0x2f, 0x9e, 0x75, 0x36, 0x0b, 0x32, 0xf0, 0x3b, 0x92, 0x98, 0xa2, 0x9e, 0x7c, 0x4d, 0xc7,
	0x4c, 0x91, 0xa5, 0x7e, 0xe0, 0xa9, 0x56, 0x88, 0x70, 0xc0, 0xd9, 0x82, 0x0c, 0xc1, 0x49, 0x6c,
	0xde, 0xd8, 0x86, 0xca, 0x68, 0x53, 0xd7, 0xdb, 0xab, 0x3b, 0xa1, 0xdb, 0xac, 0x45, 0xe6, 0xf4,
	0xe8, 0xb7, 0xd3, 0xe2, 0xed, 0x1f, 0xc1, 0x1b, 0x1c, 0x5b, 0xd8, 0x7f, 0x29, 0xdc, 0xc1, 0xf9,
	0x53, 0x92, 0x29, 0x0b, 0xed, 0x69, 0xb5, 0x72, 0x73, 0x0b, 0xe5, 0x51, 0x3e, 0xc8, 0x3c, 0xa7,
	0xbe, 0x1f, 0xba, 0xcd, 0xb3, 0x6d, 0xca, 0xf8, 0x67, 0x1f, 0x84, 0x80, 0x5d, 0x9a, 0xd3, 0xdf,
	0x79, 0x34, 0xbf, 0x1c, 0x38, 0x35, 0x15, 0xbe, 0x90, 0xfc, 0x48, 0x46, 0x84, 0x7d, 0xb8, 0xca,
	0x23, 0x5d, 0xab, 0xdb, 0x64, 0xf1, 0x66, 0xad, 0x56, 0x18, 0xb8, 0x5f, 0x8d, 0xf2, 0xe6, 0xa3,
	0x75, 0xf3, 0xd1, 0xfd, 0xaa, 0xf9, 0xa0, 0x84, 0x6d, 0x70, 0x37, 0x1c, 0x61, 0x01, 0xd0, 0xf7,
	0xac, 0xc6, 0x8e, 0xfb, 0xa0, 0x84, 0x5d, 0xa8, 0x6c, 0x79, 0xc2, 0x6b, 0x0b, 0xde, 0x75, 0x5a,
	0x7c, 0xc8, 0xf0, 0xc8, 0x4c, 0x5a, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0xef, 0x04, 0x43,
	0xb8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	//* Create a new channel
	CreateChannelOrUpdate(ctx context.Context, in *CreateOrUpdateChannelInput, opts ...grpc.CallOption) (*empty.Empty, error)
	//* List channels
	ListChannel(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListChannelOutput, error)
	//* Delete  channel
	DeleteChannel(ctx context.Context, in *DeleteChannelInput, opts ...grpc.CallOption) (*empty.Empty, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateChannelOrUpdate(ctx context.Context, in *CreateOrUpdateChannelInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tawny.AdminService/CreateChannelOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ListChannel(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListChannelOutput, error) {
	out := new(ListChannelOutput)
	err := c.cc.Invoke(ctx, "/tawny.AdminService/ListChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteChannel(ctx context.Context, in *DeleteChannelInput, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tawny.AdminService/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	//* Create a new channel
	CreateChannelOrUpdate(context.Context, *CreateOrUpdateChannelInput) (*empty.Empty, error)
	//* List channels
	ListChannel(context.Context, *empty.Empty) (*ListChannelOutput, error)
	//* Delete  channel
	DeleteChannel(context.Context, *DeleteChannelInput) (*empty.Empty, error)
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) CreateChannelOrUpdate(ctx context.Context, req *CreateOrUpdateChannelInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannelOrUpdate not implemented")
}
func (*UnimplementedAdminServiceServer) ListChannel(ctx context.Context, req *empty.Empty) (*ListChannelOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannel not implemented")
}
func (*UnimplementedAdminServiceServer) DeleteChannel(ctx context.Context, req *DeleteChannelInput) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_CreateChannelOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateChannelInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateChannelOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tawny.AdminService/CreateChannelOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateChannelOrUpdate(ctx, req.(*CreateOrUpdateChannelInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ListChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ListChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tawny.AdminService/ListChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ListChannel(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tawny.AdminService/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteChannel(ctx, req.(*DeleteChannelInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tawny.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannelOrUpdate",
			Handler:    _AdminService_CreateChannelOrUpdate_Handler,
		},
		{
			MethodName: "ListChannel",
			Handler:    _AdminService_ListChannel_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _AdminService_DeleteChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Admin.proto",
}