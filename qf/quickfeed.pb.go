// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: qf/quickfeed.proto

package qf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_qf_quickfeed_proto protoreflect.FileDescriptor

var file_qf_quickfeed_proto_rawDesc = []byte{
	0x0a, 0x12, 0x71, 0x66, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x71, 0x66, 0x1a, 0x0e, 0x71, 0x66, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x71, 0x66, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe4, 0x0e, 0x0a, 0x10,
	0x51, 0x75, 0x69, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x71, 0x66,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x21, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x08, 0x2e,
	0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x71,
	0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x71, 0x66,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00,
	0x12, 0x25, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x12, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x71, 0x66, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a,
	0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x71, 0x66, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x11,
	0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x15, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e,
	0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x08, 0x2e,
	0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f,
	0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x71,
	0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x71, 0x66,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x71,
	0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e,
	0x71, 0x66, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x12, 0x52, 0x65, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x2e,
	0x71, 0x66, 0x2e, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x24, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x09, 0x2e, 0x71,
	0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x14, 0x2e, 0x71,
	0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x08, 0x2e,
	0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x2e, 0x71,
	0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e,
	0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x71,
	0x66, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e,
	0x71, 0x66, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x2e, 0x71, 0x66, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x71, 0x66, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x0b, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x15, 0x2e,
	0x71, 0x66, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0e,
	0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x26, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b,
	0x66, 0x65, 0x65, 0x64, 0x2f, 0x71, 0x66, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_qf_quickfeed_proto_goTypes = []interface{}{
	(*Void)(nil),                     // 0: qf.Void
	(*User)(nil),                     // 1: qf.User
	(*GroupRequest)(nil),             // 2: qf.GroupRequest
	(*CourseRequest)(nil),            // 3: qf.CourseRequest
	(*Group)(nil),                    // 4: qf.Group
	(*Course)(nil),                   // 5: qf.Course
	(*Enrollment)(nil),               // 6: qf.Enrollment
	(*EnrollmentRequest)(nil),        // 7: qf.EnrollmentRequest
	(*Enrollments)(nil),              // 8: qf.Enrollments
	(*SubmissionRequest)(nil),        // 9: qf.SubmissionRequest
	(*UpdateSubmissionRequest)(nil),  // 10: qf.UpdateSubmissionRequest
	(*UpdateSubmissionsRequest)(nil), // 11: qf.UpdateSubmissionsRequest
	(*RebuildRequest)(nil),           // 12: qf.RebuildRequest
	(*Grade)(nil),                    // 13: qf.Grade
	(*GradingBenchmark)(nil),         // 14: qf.GradingBenchmark
	(*GradingCriterion)(nil),         // 15: qf.GradingCriterion
	(*ReviewRequest)(nil),            // 16: qf.ReviewRequest
	(*Organization)(nil),             // 17: qf.Organization
	(*RepositoryRequest)(nil),        // 18: qf.RepositoryRequest
	(*Users)(nil),                    // 19: qf.Users
	(*Groups)(nil),                   // 20: qf.Groups
	(*Courses)(nil),                  // 21: qf.Courses
	(*Assignments)(nil),              // 22: qf.Assignments
	(*Submission)(nil),               // 23: qf.Submission
	(*Submissions)(nil),              // 24: qf.Submissions
	(*CourseSubmissions)(nil),        // 25: qf.CourseSubmissions
	(*Review)(nil),                   // 26: qf.Review
	(*Repositories)(nil),             // 27: qf.Repositories
}
var file_qf_quickfeed_proto_depIdxs = []int32{
	0,  // 0: qf.QuickFeedService.GetUser:input_type -> qf.Void
	0,  // 1: qf.QuickFeedService.GetUsers:input_type -> qf.Void
	1,  // 2: qf.QuickFeedService.UpdateUser:input_type -> qf.User
	2,  // 3: qf.QuickFeedService.GetGroup:input_type -> qf.GroupRequest
	3,  // 4: qf.QuickFeedService.GetGroupsByCourse:input_type -> qf.CourseRequest
	4,  // 5: qf.QuickFeedService.CreateGroup:input_type -> qf.Group
	4,  // 6: qf.QuickFeedService.UpdateGroup:input_type -> qf.Group
	2,  // 7: qf.QuickFeedService.DeleteGroup:input_type -> qf.GroupRequest
	3,  // 8: qf.QuickFeedService.GetCourse:input_type -> qf.CourseRequest
	0,  // 9: qf.QuickFeedService.GetCourses:input_type -> qf.Void
	5,  // 10: qf.QuickFeedService.CreateCourse:input_type -> qf.Course
	5,  // 11: qf.QuickFeedService.UpdateCourse:input_type -> qf.Course
	6,  // 12: qf.QuickFeedService.UpdateCourseVisibility:input_type -> qf.Enrollment
	3,  // 13: qf.QuickFeedService.GetAssignments:input_type -> qf.CourseRequest
	3,  // 14: qf.QuickFeedService.UpdateAssignments:input_type -> qf.CourseRequest
	7,  // 15: qf.QuickFeedService.GetEnrollments:input_type -> qf.EnrollmentRequest
	6,  // 16: qf.QuickFeedService.CreateEnrollment:input_type -> qf.Enrollment
	8,  // 17: qf.QuickFeedService.UpdateEnrollments:input_type -> qf.Enrollments
	9,  // 18: qf.QuickFeedService.GetSubmission:input_type -> qf.SubmissionRequest
	9,  // 19: qf.QuickFeedService.GetSubmissions:input_type -> qf.SubmissionRequest
	9,  // 20: qf.QuickFeedService.GetSubmissionsByCourse:input_type -> qf.SubmissionRequest
	10, // 21: qf.QuickFeedService.UpdateSubmission:input_type -> qf.UpdateSubmissionRequest
	11, // 22: qf.QuickFeedService.UpdateSubmissions:input_type -> qf.UpdateSubmissionsRequest
	12, // 23: qf.QuickFeedService.RebuildSubmissions:input_type -> qf.RebuildRequest
	13, // 24: qf.QuickFeedService.UpdateGrade:input_type -> qf.Grade
	14, // 25: qf.QuickFeedService.CreateBenchmark:input_type -> qf.GradingBenchmark
	14, // 26: qf.QuickFeedService.UpdateBenchmark:input_type -> qf.GradingBenchmark
	14, // 27: qf.QuickFeedService.DeleteBenchmark:input_type -> qf.GradingBenchmark
	15, // 28: qf.QuickFeedService.CreateCriterion:input_type -> qf.GradingCriterion
	15, // 29: qf.QuickFeedService.UpdateCriterion:input_type -> qf.GradingCriterion
	15, // 30: qf.QuickFeedService.DeleteCriterion:input_type -> qf.GradingCriterion
	16, // 31: qf.QuickFeedService.CreateReview:input_type -> qf.ReviewRequest
	16, // 32: qf.QuickFeedService.UpdateReview:input_type -> qf.ReviewRequest
	17, // 33: qf.QuickFeedService.GetOrganization:input_type -> qf.Organization
	3,  // 34: qf.QuickFeedService.GetRepositories:input_type -> qf.CourseRequest
	18, // 35: qf.QuickFeedService.IsEmptyRepo:input_type -> qf.RepositoryRequest
	0,  // 36: qf.QuickFeedService.SubmissionStream:input_type -> qf.Void
	1,  // 37: qf.QuickFeedService.GetUser:output_type -> qf.User
	19, // 38: qf.QuickFeedService.GetUsers:output_type -> qf.Users
	0,  // 39: qf.QuickFeedService.UpdateUser:output_type -> qf.Void
	4,  // 40: qf.QuickFeedService.GetGroup:output_type -> qf.Group
	20, // 41: qf.QuickFeedService.GetGroupsByCourse:output_type -> qf.Groups
	4,  // 42: qf.QuickFeedService.CreateGroup:output_type -> qf.Group
	4,  // 43: qf.QuickFeedService.UpdateGroup:output_type -> qf.Group
	0,  // 44: qf.QuickFeedService.DeleteGroup:output_type -> qf.Void
	5,  // 45: qf.QuickFeedService.GetCourse:output_type -> qf.Course
	21, // 46: qf.QuickFeedService.GetCourses:output_type -> qf.Courses
	5,  // 47: qf.QuickFeedService.CreateCourse:output_type -> qf.Course
	0,  // 48: qf.QuickFeedService.UpdateCourse:output_type -> qf.Void
	0,  // 49: qf.QuickFeedService.UpdateCourseVisibility:output_type -> qf.Void
	22, // 50: qf.QuickFeedService.GetAssignments:output_type -> qf.Assignments
	0,  // 51: qf.QuickFeedService.UpdateAssignments:output_type -> qf.Void
	8,  // 52: qf.QuickFeedService.GetEnrollments:output_type -> qf.Enrollments
	0,  // 53: qf.QuickFeedService.CreateEnrollment:output_type -> qf.Void
	0,  // 54: qf.QuickFeedService.UpdateEnrollments:output_type -> qf.Void
	23, // 55: qf.QuickFeedService.GetSubmission:output_type -> qf.Submission
	24, // 56: qf.QuickFeedService.GetSubmissions:output_type -> qf.Submissions
	25, // 57: qf.QuickFeedService.GetSubmissionsByCourse:output_type -> qf.CourseSubmissions
	0,  // 58: qf.QuickFeedService.UpdateSubmission:output_type -> qf.Void
	0,  // 59: qf.QuickFeedService.UpdateSubmissions:output_type -> qf.Void
	0,  // 60: qf.QuickFeedService.RebuildSubmissions:output_type -> qf.Void
	0,  // 61: qf.QuickFeedService.UpdateGrade:output_type -> qf.Void
	14, // 62: qf.QuickFeedService.CreateBenchmark:output_type -> qf.GradingBenchmark
	0,  // 63: qf.QuickFeedService.UpdateBenchmark:output_type -> qf.Void
	0,  // 64: qf.QuickFeedService.DeleteBenchmark:output_type -> qf.Void
	15, // 65: qf.QuickFeedService.CreateCriterion:output_type -> qf.GradingCriterion
	0,  // 66: qf.QuickFeedService.UpdateCriterion:output_type -> qf.Void
	0,  // 67: qf.QuickFeedService.DeleteCriterion:output_type -> qf.Void
	26, // 68: qf.QuickFeedService.CreateReview:output_type -> qf.Review
	26, // 69: qf.QuickFeedService.UpdateReview:output_type -> qf.Review
	17, // 70: qf.QuickFeedService.GetOrganization:output_type -> qf.Organization
	27, // 71: qf.QuickFeedService.GetRepositories:output_type -> qf.Repositories
	0,  // 72: qf.QuickFeedService.IsEmptyRepo:output_type -> qf.Void
	23, // 73: qf.QuickFeedService.SubmissionStream:output_type -> qf.Submission
	37, // [37:74] is the sub-list for method output_type
	0,  // [0:37] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_qf_quickfeed_proto_init() }
func file_qf_quickfeed_proto_init() {
	if File_qf_quickfeed_proto != nil {
		return
	}
	file_qf_types_proto_init()
	file_qf_requests_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_qf_quickfeed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qf_quickfeed_proto_goTypes,
		DependencyIndexes: file_qf_quickfeed_proto_depIdxs,
	}.Build()
	File_qf_quickfeed_proto = out.File
	file_qf_quickfeed_proto_rawDesc = nil
	file_qf_quickfeed_proto_goTypes = nil
	file_qf_quickfeed_proto_depIdxs = nil
}
