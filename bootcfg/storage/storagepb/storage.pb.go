// Code generated by protoc-gen-go.
// source: storage.proto
// DO NOT EDIT!

/*
Package storagepb is a generated protocol buffer package.

It is generated from these files:
	storage.proto

It has these top-level messages:
	Group
	Profile
	NetBoot
*/
package storagepb

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

// Group selects one or more machines and matches them to a Profile.
type Group struct {
	// machine readable Id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Profile id
	Profile string `protobuf:"bytes,3,opt,name=profile" json:"profile,omitempty"`
	// Selectors to match machines
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// JSON encoded metadata
	Metadata []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Group) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

// Profile defines the boot and provisioning behavior of a group of machines.
type Profile struct {
	// profile id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// ignition id
	IgnitionId string `protobuf:"bytes,3,opt,name=ignition_id,json=ignitionId" json:"ignition_id,omitempty"`
	// cloud config id
	CloudId string `protobuf:"bytes,4,opt,name=cloud_id,json=cloudId" json:"cloud_id,omitempty"`
	// support network boot / PXE
	Boot *NetBoot `protobuf:"bytes,5,opt,name=boot" json:"boot,omitempty"`
	// generic config id
	GenericId string `protobuf:"bytes,6,opt,name=generic_id,json=genericId" json:"generic_id,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profile) GetBoot() *NetBoot {
	if m != nil {
		return m.Boot
	}
	return nil
}

// NetBoot describes network or PXE boot settings for a machine.
type NetBoot struct {
	// the URL of the kernel image
	Kernel string `protobuf:"bytes,1,opt,name=kernel" json:"kernel,omitempty"`
	// the init RAM filesystem URLs
	Initrd []string `protobuf:"bytes,2,rep,name=initrd" json:"initrd,omitempty"`
	// kernel parameters (command line)
	Cmdline map[string]string `protobuf:"bytes,3,rep,name=cmdline" json:"cmdline,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NetBoot) Reset()                    { *m = NetBoot{} }
func (m *NetBoot) String() string            { return proto.CompactTextString(m) }
func (*NetBoot) ProtoMessage()               {}
func (*NetBoot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NetBoot) GetCmdline() map[string]string {
	if m != nil {
		return m.Cmdline
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "storagepb.Group")
	proto.RegisterType((*Profile)(nil), "storagepb.Profile")
	proto.RegisterType((*NetBoot)(nil), "storagepb.NetBoot")
}

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0x54, 0x9a, 0xf4, 0x6f, 0xdb, 0x7e, 0xfa, 0xb4, 0x42, 0x28, 0x54, 0x82, 0x56, 0x3d, 0xa0,
	0x9e, 0x72, 0x28, 0x17, 0x28, 0x37, 0x10, 0x42, 0xbd, 0x20, 0x14, 0x1e, 0x00, 0xa5, 0xf1, 0x52,
	0x59, 0x4d, 0xed, 0xca, 0x75, 0x91, 0xfa, 0x56, 0x3c, 0x0f, 0x4f, 0x83, 0xed, 0x38, 0x51, 0x11,
	0x17, 0xb8, 0xed, 0xcc, 0x8e, 0xc7, 0xb3, 0x5e, 0xc3, 0x60, 0xa7, 0xa5, 0xca, 0x56, 0x94, 0x6c,
	0x95, 0xd4, 0x12, 0xbb, 0x1e, 0x6e, 0x97, 0x93, 0xcf, 0x00, 0x9a, 0x8f, 0x4a, 0xee, 0xb7, 0xf8,
	0x0f, 0x1a, 0x9c, 0xc5, 0xc1, 0x38, 0x98, 0x76, 0x53, 0x53, 0x21, 0x42, 0x24, 0xb2, 0x0d, 0xc5,
	0x0d, 0xc7, 0xb8, 0x1a, 0x63, 0x68, 0x1b, 0x87, 0x37, 0x5e, 0x50, 0x1c, 0x3a, 0xba, 0x82, 0x38,
	0x87, 0xce, 0x8e, 0x0a, 0xca, 0x8d, 0x71, 0x1c, 0x8d, 0xc3, 0x69, 0x6f, 0x76, 0x91, 0xd4, 0xb7,
	0x24, 0xee, 0x86, 0xe4, 0xc5, 0x0b, 0x1e, 0x84, 0x56, 0x87, 0xb4, 0xd6, 0xe3, 0x10, 0x3a, 0x1b,
	0xd2, 0x19, 0xcb, 0x74, 0x16, 0x37, 0x8d, 0x6d, 0x3f, 0xad, 0xf1, 0xf0, 0x16, 0x06, 0xdf, 0x8e,
	0xe1, 0x7f, 0x08, 0xd7, 0x74, 0xf0, 0x39, 0x6d, 0x89, 0x27, 0xd0, 0x7c, 0xcf, 0x8a, 0x7d, 0x95,
	0xb4, 0x04, 0xf3, 0xc6, 0x75, 0x30, 0xf9, 0x08, 0xa0, 0xfd, 0xec, 0x03, 0xfe, 0x66, 0xbc, 0x11,
	0xf4, 0xf8, 0x4a, 0x70, 0xcd, 0xa5, 0x78, 0x35, 0xe2, 0x72, 0x44, 0xa8, 0xa8, 0x05, 0xc3, 0x33,
	0xe8, 0xe4, 0x85, 0xdc, 0x33, 0xdb, 0x8d, 0xca, 0x07, 0x70, 0xd8, 0xb4, 0x2e, 0x21, 0x5a, 0x4a,
	0xa9, 0xdd, 0x00, 0xbd, 0x19, 0x1e, 0x0d, 0xff, 0x44, 0xfa, 0xce, 0x74, 0x52, 0xd7, 0xc7, 0x73,
	0x80, 0x15, 0x09, 0x52, 0x3c, 0xb7, 0x26, 0x2d, 0x67, 0xd2, 0xf5, 0xcc, 0x82, 0xb9, 0xc8, 0xfe,
	0x00, 0x9e, 0x42, 0x6b, 0x4d, 0x4a, 0x50, 0xe1, 0x63, 0x7b, 0x64, 0x79, 0x6e, 0x22, 0x29, 0x66,
	0xc2, 0x87, 0x96, 0x2f, 0x11, 0xde, 0x40, 0x3b, 0xdf, 0xb0, 0x82, 0x0b, 0xbb, 0x1d, 0xbb, 0x82,
	0xd1, 0xcf, 0x14, 0xc9, 0x7d, 0xa9, 0x28, 0x77, 0x50, 0xe9, 0x87, 0x73, 0xe8, 0x1f, 0x37, 0xfe,
	0xf2, 0xca, 0xcb, 0x96, 0xfb, 0x54, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x0d, 0x6c,
	0x68, 0x65, 0x02, 0x00, 0x00,
}
