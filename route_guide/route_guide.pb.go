// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: route_guide/route_guide.proto

package route_guide

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

var File_route_guide_route_guide_proto protoreflect.FileDescriptor

var file_route_guide_route_guide_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x1a, 0x17, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x44, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x36,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_route_guide_route_guide_proto_goTypes = []interface{}{
	(*Point)(nil),   // 0: route_guide.Point
	(*Feature)(nil), // 1: route_guide.Feature
}
var file_route_guide_route_guide_proto_depIdxs = []int32{
	0, // 0: route_guide.RouteGuide.GetFeature:input_type -> route_guide.Point
	1, // 1: route_guide.RouteGuide.GetFeature:output_type -> route_guide.Feature
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_route_guide_route_guide_proto_init() }
func file_route_guide_route_guide_proto_init() {
	if File_route_guide_route_guide_proto != nil {
		return
	}
	file_route_guide_point_proto_init()
	file_route_guide_feature_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_guide_route_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_guide_route_guide_proto_goTypes,
		DependencyIndexes: file_route_guide_route_guide_proto_depIdxs,
	}.Build()
	File_route_guide_route_guide_proto = out.File
	file_route_guide_route_guide_proto_rawDesc = nil
	file_route_guide_route_guide_proto_goTypes = nil
	file_route_guide_route_guide_proto_depIdxs = nil
}
