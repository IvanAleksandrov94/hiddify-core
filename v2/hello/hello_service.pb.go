// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: v2/hello/hello_service.proto

package hello

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

var File_v2_hello_hello_service_proto protoreflect.FileDescriptor

var file_v2_hello_hello_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x32, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x14, 0x76, 0x32, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x7f, 0x0a, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x64, 0x64, 0x69,
	0x66, 0x79, 0x2f, 0x68, 0x69, 0x64, 0x64, 0x69, 0x66, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x32, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_v2_hello_hello_service_proto_goTypes = []any{
	(*HelloRequest)(nil),  // 0: hello.HelloRequest
	(*HelloResponse)(nil), // 1: hello.HelloResponse
}
var file_v2_hello_hello_service_proto_depIdxs = []int32{
	0, // 0: hello.Hello.SayHello:input_type -> hello.HelloRequest
	0, // 1: hello.Hello.SayHelloStream:input_type -> hello.HelloRequest
	1, // 2: hello.Hello.SayHello:output_type -> hello.HelloResponse
	1, // 3: hello.Hello.SayHelloStream:output_type -> hello.HelloResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v2_hello_hello_service_proto_init() }
func file_v2_hello_hello_service_proto_init() {
	if File_v2_hello_hello_service_proto != nil {
		return
	}
	file_v2_hello_hello_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v2_hello_hello_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2_hello_hello_service_proto_goTypes,
		DependencyIndexes: file_v2_hello_hello_service_proto_depIdxs,
	}.Build()
	File_v2_hello_hello_service_proto = out.File
	file_v2_hello_hello_service_proto_rawDesc = nil
	file_v2_hello_hello_service_proto_goTypes = nil
	file_v2_hello_hello_service_proto_depIdxs = nil
}
