// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: privilege.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAccessibleApplicationsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetAccessibleApplicationsByUserIdRequest) Reset() {
	*x = GetAccessibleApplicationsByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privilege_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessibleApplicationsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessibleApplicationsByUserIdRequest) ProtoMessage() {}

func (x *GetAccessibleApplicationsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_privilege_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessibleApplicationsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetAccessibleApplicationsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_privilege_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccessibleApplicationsByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetAccessibleApplicationsByUserIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications *structpb.Struct `protobuf:"bytes,1,opt,name=Applications,proto3" json:"Applications,omitempty"`
}

func (x *GetAccessibleApplicationsByUserIdReply) Reset() {
	*x = GetAccessibleApplicationsByUserIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_privilege_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessibleApplicationsByUserIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessibleApplicationsByUserIdReply) ProtoMessage() {}

func (x *GetAccessibleApplicationsByUserIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_privilege_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessibleApplicationsByUserIdReply.ProtoReflect.Descriptor instead.
func (*GetAccessibleApplicationsByUserIdReply) Descriptor() ([]byte, []int) {
	return file_privilege_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessibleApplicationsByUserIdReply) GetApplications() *structpb.Struct {
	if x != nil {
		return x.Applications
	}
	return nil
}

var File_privilege_proto protoreflect.FileDescriptor

var file_privilege_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x28, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65,
	0x0a, 0x26, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x9b, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x2e, 0x70, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x76, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x76,
	0x2d, 0x73, 0x73, 0x6f, 0x2d, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_privilege_proto_rawDescOnce sync.Once
	file_privilege_proto_rawDescData = file_privilege_proto_rawDesc
)

func file_privilege_proto_rawDescGZIP() []byte {
	file_privilege_proto_rawDescOnce.Do(func() {
		file_privilege_proto_rawDescData = protoimpl.X.CompressGZIP(file_privilege_proto_rawDescData)
	})
	return file_privilege_proto_rawDescData
}

var file_privilege_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_privilege_proto_goTypes = []interface{}{
	(*GetAccessibleApplicationsByUserIdRequest)(nil), // 0: privilege.GetAccessibleApplicationsByUserIdRequest
	(*GetAccessibleApplicationsByUserIdReply)(nil),   // 1: privilege.GetAccessibleApplicationsByUserIdReply
	(*structpb.Struct)(nil),                          // 2: google.protobuf.Struct
}
var file_privilege_proto_depIdxs = []int32{
	2, // 0: privilege.GetAccessibleApplicationsByUserIdReply.Applications:type_name -> google.protobuf.Struct
	0, // 1: privilege.Privilege.GetAccessibleApplicationsByUserId:input_type -> privilege.GetAccessibleApplicationsByUserIdRequest
	1, // 2: privilege.Privilege.GetAccessibleApplicationsByUserId:output_type -> privilege.GetAccessibleApplicationsByUserIdReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_privilege_proto_init() }
func file_privilege_proto_init() {
	if File_privilege_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_privilege_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessibleApplicationsByUserIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_privilege_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessibleApplicationsByUserIdReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_privilege_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_privilege_proto_goTypes,
		DependencyIndexes: file_privilege_proto_depIdxs,
		MessageInfos:      file_privilege_proto_msgTypes,
	}.Build()
	File_privilege_proto = out.File
	file_privilege_proto_rawDesc = nil
	file_privilege_proto_goTypes = nil
	file_privilege_proto_depIdxs = nil
}
