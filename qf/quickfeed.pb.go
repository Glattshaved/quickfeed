// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
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
	0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x10, 0x0a, 0x10,
	0x51, 0x75, 0x69, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x71, 0x66,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x21, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x08, 0x2e,
	0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x71,
	0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x13, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x79, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x25, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x2e, 0x71,
	0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x09, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x11,
	0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x08, 0x2e,
	0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x26, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x08, 0x2e, 0x71,
	0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x71, 0x66, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x71, 0x66,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08,
	0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x71,
	0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x71, 0x66, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x71, 0x66, 0x2e,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x71, 0x66,
	0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x08, 0x2e, 0x71,
	0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x71, 0x66, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x1f, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
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
	0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x1a, 0x08, 0x2e, 0x71,
	0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x71, 0x66,
	0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x1a, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x71,
	0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x33,
	0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x71, 0x66, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x11, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x4f,
	0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x12, 0x15, 0x2e, 0x71, 0x66, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x71, 0x66, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x08, 0x2e, 0x71, 0x66, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x71, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x66, 0x65, 0x65, 0x64,
	0x2f, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x71, 0x66, 0xba, 0x02, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_qf_quickfeed_proto_goTypes = []interface{}{
	(*Void)(nil),                        // 0: qf.Void
	(*User)(nil),                        // 1: qf.User
	(*GetGroupRequest)(nil),             // 2: qf.GetGroupRequest
	(*GroupRequest)(nil),                // 3: qf.GroupRequest
	(*CourseRequest)(nil),               // 4: qf.CourseRequest
	(*Group)(nil),                       // 5: qf.Group
	(*Course)(nil),                      // 6: qf.Course
	(*Enrollment)(nil),                  // 7: qf.Enrollment
	(*EnrollmentStatusRequest)(nil),     // 8: qf.EnrollmentStatusRequest
	(*EnrollmentRequest)(nil),           // 9: qf.EnrollmentRequest
	(*Enrollments)(nil),                 // 10: qf.Enrollments
	(*SubmissionRequest)(nil),           // 11: qf.SubmissionRequest
	(*SubmissionReviewersRequest)(nil),  // 12: qf.SubmissionReviewersRequest
	(*SubmissionsForCourseRequest)(nil), // 13: qf.SubmissionsForCourseRequest
	(*UpdateSubmissionRequest)(nil),     // 14: qf.UpdateSubmissionRequest
	(*UpdateSubmissionsRequest)(nil),    // 15: qf.UpdateSubmissionsRequest
	(*RebuildRequest)(nil),              // 16: qf.RebuildRequest
	(*GradingBenchmark)(nil),            // 17: qf.GradingBenchmark
	(*GradingCriterion)(nil),            // 18: qf.GradingCriterion
	(*ReviewRequest)(nil),               // 19: qf.ReviewRequest
	(*OrgRequest)(nil),                  // 20: qf.OrgRequest
	(*URLRequest)(nil),                  // 21: qf.URLRequest
	(*RepositoryRequest)(nil),           // 22: qf.RepositoryRequest
	(*Users)(nil),                       // 23: qf.Users
	(*Groups)(nil),                      // 24: qf.Groups
	(*Courses)(nil),                     // 25: qf.Courses
	(*Assignments)(nil),                 // 26: qf.Assignments
	(*Submissions)(nil),                 // 27: qf.Submissions
	(*Submission)(nil),                  // 28: qf.Submission
	(*CourseSubmissions)(nil),           // 29: qf.CourseSubmissions
	(*Review)(nil),                      // 30: qf.Review
	(*Reviewers)(nil),                   // 31: qf.Reviewers
	(*Organization)(nil),                // 32: qf.Organization
	(*Repositories)(nil),                // 33: qf.Repositories
}
var file_qf_quickfeed_proto_depIdxs = []int32{
	0,  // 0: qf.QuickFeedService.GetUser:input_type -> qf.Void
	0,  // 1: qf.QuickFeedService.GetUsers:input_type -> qf.Void
	1,  // 2: qf.QuickFeedService.UpdateUser:input_type -> qf.User
	2,  // 3: qf.QuickFeedService.GetGroup:input_type -> qf.GetGroupRequest
	3,  // 4: qf.QuickFeedService.GetGroupByUserAndCourse:input_type -> qf.GroupRequest
	4,  // 5: qf.QuickFeedService.GetGroupsByCourse:input_type -> qf.CourseRequest
	5,  // 6: qf.QuickFeedService.CreateGroup:input_type -> qf.Group
	5,  // 7: qf.QuickFeedService.UpdateGroup:input_type -> qf.Group
	3,  // 8: qf.QuickFeedService.DeleteGroup:input_type -> qf.GroupRequest
	4,  // 9: qf.QuickFeedService.GetCourse:input_type -> qf.CourseRequest
	0,  // 10: qf.QuickFeedService.GetCourses:input_type -> qf.Void
	6,  // 11: qf.QuickFeedService.CreateCourse:input_type -> qf.Course
	6,  // 12: qf.QuickFeedService.UpdateCourse:input_type -> qf.Course
	7,  // 13: qf.QuickFeedService.UpdateCourseVisibility:input_type -> qf.Enrollment
	4,  // 14: qf.QuickFeedService.GetAssignments:input_type -> qf.CourseRequest
	4,  // 15: qf.QuickFeedService.UpdateAssignments:input_type -> qf.CourseRequest
	8,  // 16: qf.QuickFeedService.GetEnrollmentsByUser:input_type -> qf.EnrollmentStatusRequest
	9,  // 17: qf.QuickFeedService.GetEnrollmentsByCourse:input_type -> qf.EnrollmentRequest
	7,  // 18: qf.QuickFeedService.CreateEnrollment:input_type -> qf.Enrollment
	10, // 19: qf.QuickFeedService.UpdateEnrollments:input_type -> qf.Enrollments
	11, // 20: qf.QuickFeedService.GetSubmissions:input_type -> qf.SubmissionRequest
	12, // 21: qf.QuickFeedService.GetSubmission:input_type -> qf.SubmissionReviewersRequest
	13, // 22: qf.QuickFeedService.GetSubmissionsByCourse:input_type -> qf.SubmissionsForCourseRequest
	14, // 23: qf.QuickFeedService.UpdateSubmission:input_type -> qf.UpdateSubmissionRequest
	15, // 24: qf.QuickFeedService.UpdateSubmissions:input_type -> qf.UpdateSubmissionsRequest
	16, // 25: qf.QuickFeedService.RebuildSubmissions:input_type -> qf.RebuildRequest
	17, // 26: qf.QuickFeedService.CreateBenchmark:input_type -> qf.GradingBenchmark
	17, // 27: qf.QuickFeedService.UpdateBenchmark:input_type -> qf.GradingBenchmark
	17, // 28: qf.QuickFeedService.DeleteBenchmark:input_type -> qf.GradingBenchmark
	18, // 29: qf.QuickFeedService.CreateCriterion:input_type -> qf.GradingCriterion
	18, // 30: qf.QuickFeedService.UpdateCriterion:input_type -> qf.GradingCriterion
	18, // 31: qf.QuickFeedService.DeleteCriterion:input_type -> qf.GradingCriterion
	19, // 32: qf.QuickFeedService.CreateReview:input_type -> qf.ReviewRequest
	19, // 33: qf.QuickFeedService.UpdateReview:input_type -> qf.ReviewRequest
	12, // 34: qf.QuickFeedService.GetReviewers:input_type -> qf.SubmissionReviewersRequest
	20, // 35: qf.QuickFeedService.GetOrganization:input_type -> qf.OrgRequest
	21, // 36: qf.QuickFeedService.GetRepositories:input_type -> qf.URLRequest
	22, // 37: qf.QuickFeedService.IsEmptyRepo:input_type -> qf.RepositoryRequest
	0,  // 38: qf.QuickFeedService.SubmissionStream:input_type -> qf.Void
	1,  // 39: qf.QuickFeedService.GetUser:output_type -> qf.User
	23, // 40: qf.QuickFeedService.GetUsers:output_type -> qf.Users
	0,  // 41: qf.QuickFeedService.UpdateUser:output_type -> qf.Void
	5,  // 42: qf.QuickFeedService.GetGroup:output_type -> qf.Group
	5,  // 43: qf.QuickFeedService.GetGroupByUserAndCourse:output_type -> qf.Group
	24, // 44: qf.QuickFeedService.GetGroupsByCourse:output_type -> qf.Groups
	5,  // 45: qf.QuickFeedService.CreateGroup:output_type -> qf.Group
	5,  // 46: qf.QuickFeedService.UpdateGroup:output_type -> qf.Group
	0,  // 47: qf.QuickFeedService.DeleteGroup:output_type -> qf.Void
	6,  // 48: qf.QuickFeedService.GetCourse:output_type -> qf.Course
	25, // 49: qf.QuickFeedService.GetCourses:output_type -> qf.Courses
	6,  // 50: qf.QuickFeedService.CreateCourse:output_type -> qf.Course
	0,  // 51: qf.QuickFeedService.UpdateCourse:output_type -> qf.Void
	0,  // 52: qf.QuickFeedService.UpdateCourseVisibility:output_type -> qf.Void
	26, // 53: qf.QuickFeedService.GetAssignments:output_type -> qf.Assignments
	0,  // 54: qf.QuickFeedService.UpdateAssignments:output_type -> qf.Void
	10, // 55: qf.QuickFeedService.GetEnrollmentsByUser:output_type -> qf.Enrollments
	10, // 56: qf.QuickFeedService.GetEnrollmentsByCourse:output_type -> qf.Enrollments
	0,  // 57: qf.QuickFeedService.CreateEnrollment:output_type -> qf.Void
	0,  // 58: qf.QuickFeedService.UpdateEnrollments:output_type -> qf.Void
	27, // 59: qf.QuickFeedService.GetSubmissions:output_type -> qf.Submissions
	28, // 60: qf.QuickFeedService.GetSubmission:output_type -> qf.Submission
	29, // 61: qf.QuickFeedService.GetSubmissionsByCourse:output_type -> qf.CourseSubmissions
	0,  // 62: qf.QuickFeedService.UpdateSubmission:output_type -> qf.Void
	0,  // 63: qf.QuickFeedService.UpdateSubmissions:output_type -> qf.Void
	0,  // 64: qf.QuickFeedService.RebuildSubmissions:output_type -> qf.Void
	17, // 65: qf.QuickFeedService.CreateBenchmark:output_type -> qf.GradingBenchmark
	0,  // 66: qf.QuickFeedService.UpdateBenchmark:output_type -> qf.Void
	0,  // 67: qf.QuickFeedService.DeleteBenchmark:output_type -> qf.Void
	18, // 68: qf.QuickFeedService.CreateCriterion:output_type -> qf.GradingCriterion
	0,  // 69: qf.QuickFeedService.UpdateCriterion:output_type -> qf.Void
	0,  // 70: qf.QuickFeedService.DeleteCriterion:output_type -> qf.Void
	30, // 71: qf.QuickFeedService.CreateReview:output_type -> qf.Review
	30, // 72: qf.QuickFeedService.UpdateReview:output_type -> qf.Review
	31, // 73: qf.QuickFeedService.GetReviewers:output_type -> qf.Reviewers
	32, // 74: qf.QuickFeedService.GetOrganization:output_type -> qf.Organization
	33, // 75: qf.QuickFeedService.GetRepositories:output_type -> qf.Repositories
	0,  // 76: qf.QuickFeedService.IsEmptyRepo:output_type -> qf.Void
	28, // 77: qf.QuickFeedService.SubmissionStream:output_type -> qf.Submission
	39, // [39:78] is the sub-list for method output_type
	0,  // [0:39] is the sub-list for method input_type
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
