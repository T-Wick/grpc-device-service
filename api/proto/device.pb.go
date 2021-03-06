// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Device_DeviceType int32

const (
	Device_onOff  Device_DeviceType = 0
	Device_dimmer Device_DeviceType = 1
	Device_sensor Device_DeviceType = 2
)

var Device_DeviceType_name = map[int32]string{
	0: "onOff",
	1: "dimmer",
	2: "sensor",
}

var Device_DeviceType_value = map[string]int32{
	"onOff":  0,
	"dimmer": 1,
	"sensor": 2,
}

func (x Device_DeviceType) String() string {
	return proto.EnumName(Device_DeviceType_name, int32(x))
}

func (Device_DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{4, 0}
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type User struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string    `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string    `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Devices              []*Device `protobuf:"bytes,4,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type ID struct {
	Account              *User    `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{2}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetAccount() *User {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *ID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateDevice struct {
	Account              *User    `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Value                int32    `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDevice) Reset()         { *m = UpdateDevice{} }
func (m *UpdateDevice) String() string { return proto.CompactTextString(m) }
func (*UpdateDevice) ProtoMessage()    {}
func (*UpdateDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{3}
}

func (m *UpdateDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDevice.Unmarshal(m, b)
}
func (m *UpdateDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDevice.Marshal(b, m, deterministic)
}
func (m *UpdateDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDevice.Merge(m, src)
}
func (m *UpdateDevice) XXX_Size() int {
	return xxx_messageInfo_UpdateDevice.Size(m)
}
func (m *UpdateDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDevice.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDevice proto.InternalMessageInfo

func (m *UpdateDevice) GetAccount() *User {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *UpdateDevice) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateDevice) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Device struct {
	Id                   int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hardware             string            `protobuf:"bytes,2,opt,name=hardware,proto3" json:"hardware,omitempty"`
	Name                 string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Location             string            `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Type                 Device_DeviceType `protobuf:"varint,5,opt,name=type,proto3,enum=pb.Device_DeviceType" json:"type,omitempty"`
	Unit                 string            `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	State                int32             `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	Account              *User             `protobuf:"bytes,8,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{4}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Device) GetHardware() string {
	if m != nil {
		return m.Hardware
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Device) GetType() Device_DeviceType {
	if m != nil {
		return m.Type
	}
	return Device_onOff
}

func (m *Device) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Device) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Device) GetAccount() *User {
	if m != nil {
		return m.Account
	}
	return nil
}

type Devices struct {
	Device               []*Device `protobuf:"bytes,1,rep,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Devices) Reset()         { *m = Devices{} }
func (m *Devices) String() string { return proto.CompactTextString(m) }
func (*Devices) ProtoMessage()    {}
func (*Devices) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{5}
}

func (m *Devices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Devices.Unmarshal(m, b)
}
func (m *Devices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Devices.Marshal(b, m, deterministic)
}
func (m *Devices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Devices.Merge(m, src)
}
func (m *Devices) XXX_Size() int {
	return xxx_messageInfo_Devices.Size(m)
}
func (m *Devices) XXX_DiscardUnknown() {
	xxx_messageInfo_Devices.DiscardUnknown(m)
}

var xxx_messageInfo_Devices proto.InternalMessageInfo

func (m *Devices) GetDevice() []*Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type Empty struct {
	Account              *User    `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_870276a56ac00da5, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetAccount() *User {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.Device_DeviceType", Device_DeviceType_name, Device_DeviceType_value)
	proto.RegisterType((*CreateUserRequest)(nil), "pb.CreateUserRequest")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*ID)(nil), "pb.ID")
	proto.RegisterType((*UpdateDevice)(nil), "pb.UpdateDevice")
	proto.RegisterType((*Device)(nil), "pb.Device")
	proto.RegisterType((*Devices)(nil), "pb.Devices")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

func init() { proto.RegisterFile("device.proto", fileDescriptor_870276a56ac00da5) }

var fileDescriptor_870276a56ac00da5 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x6d, 0xd2, 0x24, 0x4d, 0xa7, 0xdd, 0xaa, 0x58, 0x20, 0x59, 0x39, 0x55, 0x16, 0x48, 0x5d,
	0x10, 0x45, 0x2a, 0x17, 0x0e, 0x5c, 0x80, 0x22, 0xd4, 0x13, 0x52, 0x96, 0x95, 0xe0, 0xe8, 0x26,
	0xb3, 0xac, 0xa5, 0x7c, 0x61, 0xbb, 0x5b, 0xf5, 0x4f, 0xf0, 0x87, 0xb9, 0x20, 0xdb, 0x69, 0x1b,
	0x28, 0x48, 0x08, 0xed, 0x29, 0x7e, 0x9e, 0x37, 0xcf, 0xcf, 0xcf, 0x13, 0x18, 0xe7, 0x78, 0x27,
	0x32, 0x5c, 0x34, 0xb2, 0xd6, 0x35, 0xf1, 0x9b, 0x0d, 0xfb, 0x02, 0x0f, 0xde, 0x49, 0xe4, 0x1a,
	0xaf, 0x15, 0xca, 0x14, 0xbf, 0x6d, 0x51, 0x69, 0x42, 0x20, 0xa8, 0x78, 0x89, 0xd4, 0x9b, 0x79,
	0xf3, 0x61, 0x6a, 0xd7, 0xe4, 0x21, 0x84, 0x58, 0x72, 0x51, 0x50, 0xdf, 0x6e, 0x3a, 0x40, 0x12,
	0x88, 0x1b, 0xae, 0xd4, 0xae, 0x96, 0x39, 0xed, 0xdb, 0xc2, 0x11, 0x33, 0x09, 0x81, 0x11, 0xbd,
	0x1f, 0x35, 0xf2, 0x18, 0x06, 0xce, 0xbc, 0xa2, 0xc1, 0xac, 0x3f, 0x1f, 0x2d, 0x61, 0xd1, 0x6c,
	0x16, 0x2b, 0xbb, 0x95, 0x1e, 0x4a, 0xec, 0x15, 0xf8, 0xeb, 0x15, 0x61, 0x30, 0xe0, 0x59, 0x56,
	0x6f, 0x2b, 0x6d, 0x0f, 0x1d, 0x2d, 0x63, 0xc3, 0xb5, 0x37, 0x3c, 0x14, 0xc8, 0x04, 0x7c, 0x91,
	0xdb, 0xe3, 0xc3, 0xd4, 0x17, 0x39, 0xfb, 0x0c, 0xe3, 0xeb, 0x26, 0xe7, 0x1a, 0x9d, 0xe4, 0xff,
	0x68, 0x98, 0x5b, 0xdd, 0xf1, 0x62, 0x8b, 0xd6, 0x7c, 0x98, 0x3a, 0xc0, 0xbe, 0xfb, 0x10, 0xb5,
	0xa2, 0xae, 0xc1, 0x3b, 0x36, 0x24, 0x10, 0xdf, 0x72, 0x99, 0xef, 0xb8, 0xc4, 0x36, 0x89, 0x23,
	0x3e, 0xc6, 0xd6, 0xef, 0xc4, 0x96, 0x40, 0x5c, 0xd4, 0x19, 0xd7, 0xa2, 0xae, 0x68, 0xe0, 0xf8,
	0x07, 0x4c, 0x2e, 0x21, 0xd0, 0xfb, 0x06, 0x69, 0x38, 0xf3, 0xe6, 0x93, 0xe5, 0xa3, 0x53, 0x3a,
	0xed, 0xe7, 0xd3, 0xbe, 0xc1, 0xd4, 0x52, 0x8c, 0xf4, 0xb6, 0x12, 0x9a, 0x46, 0x4e, 0xda, 0xac,
	0x8d, 0x77, 0xa5, 0xb9, 0x46, 0x3a, 0x70, 0xde, 0x2d, 0xe8, 0xa6, 0x10, 0xff, 0x25, 0x05, 0xf6,
	0x02, 0xe0, 0x74, 0x02, 0x19, 0x42, 0x58, 0x57, 0x1f, 0x6f, 0x6e, 0xa6, 0x3d, 0x02, 0x10, 0xe5,
	0xa2, 0x2c, 0x51, 0x4e, 0x3d, 0xb3, 0x56, 0x58, 0xa9, 0x5a, 0x4e, 0x7d, 0xf6, 0x1c, 0x06, 0xae,
	0x41, 0x11, 0x06, 0x91, 0x7b, 0x3a, 0xea, 0x9d, 0x3d, 0x6a, 0x5b, 0x61, 0xcf, 0x20, 0x7c, 0x5f,
	0x36, 0x7a, 0xff, 0x2f, 0x4f, 0xb2, 0xfc, 0xe1, 0xc1, 0x85, 0xeb, 0xbf, 0x42, 0x69, 0x33, 0x7f,
	0x0d, 0x70, 0x9a, 0x70, 0x62, 0x73, 0x39, 0x9b, 0xf8, 0xe4, 0xcf, 0xdb, 0xac, 0x47, 0x2e, 0xe1,
	0xe2, 0x03, 0xea, 0x37, 0x45, 0x71, 0x70, 0x3c, 0x34, 0x4c, 0xeb, 0x27, 0x19, 0x9d, 0xcc, 0x2a,
	0xd6, 0x23, 0x4f, 0x2c, 0xd5, 0xe1, 0xb7, 0xfb, 0xf5, 0x8a, 0x44, 0xa6, 0xbe, 0x5e, 0x25, 0x9d,
	0x4b, 0xb1, 0x1e, 0x59, 0xc0, 0xf8, 0x6a, 0x27, 0x74, 0x76, 0xdb, 0xce, 0xc4, 0xd4, 0x5e, 0xa2,
	0x33, 0x7a, 0xbf, 0xf1, 0x9f, 0xc2, 0x24, 0xc5, 0xaf, 0x42, 0x69, 0x94, 0x6d, 0x47, 0xa7, 0xfe,
	0x2b, 0x77, 0x13, 0xd9, 0x1f, 0xfb, 0xe5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xc8, 0x8a,
	0xd3, 0xe8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserRequest, error)
	// List all registered devices
	GetAllDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Devices, error)
	// Get a device by ID
	GetDeviceByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Device, error)
	// Update a device's state
	SwitchDevice(ctx context.Context, in *UpdateDevice, opts ...grpc.CallOption) (*Device, error)
	// Register a new device
	RegisterDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
}

type deviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceServiceClient(cc *grpc.ClientConn) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserRequest, error) {
	out := new(CreateUserRequest)
	err := c.cc.Invoke(ctx, "/pb.DeviceService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetAllDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Devices, error) {
	out := new(Devices)
	err := c.cc.Invoke(ctx, "/pb.DeviceService/GetAllDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetDeviceByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/pb.DeviceService/GetDeviceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) SwitchDevice(ctx context.Context, in *UpdateDevice, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/pb.DeviceService/SwitchDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) RegisterDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/pb.DeviceService/RegisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
type DeviceServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserRequest, error)
	// List all registered devices
	GetAllDevices(context.Context, *Empty) (*Devices, error)
	// Get a device by ID
	GetDeviceByID(context.Context, *ID) (*Device, error)
	// Update a device's state
	SwitchDevice(context.Context, *UpdateDevice) (*Device, error)
	// Register a new device
	RegisterDevice(context.Context, *Device) (*Device, error)
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/GetAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetAllDevices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetDeviceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetDeviceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/GetDeviceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetDeviceByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_SwitchDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDevice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).SwitchDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/SwitchDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).SwitchDevice(ctx, req.(*UpdateDevice))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeviceService/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).RegisterDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _DeviceService_CreateUser_Handler,
		},
		{
			MethodName: "GetAllDevices",
			Handler:    _DeviceService_GetAllDevices_Handler,
		},
		{
			MethodName: "GetDeviceByID",
			Handler:    _DeviceService_GetDeviceByID_Handler,
		},
		{
			MethodName: "SwitchDevice",
			Handler:    _DeviceService_SwitchDevice_Handler,
		},
		{
			MethodName: "RegisterDevice",
			Handler:    _DeviceService_RegisterDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "device.proto",
}
