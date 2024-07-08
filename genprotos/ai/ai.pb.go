// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: ai.proto

package ai

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AiCHat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *AiCHat) Reset() {
	*x = AiCHat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiCHat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiCHat) ProtoMessage() {}

func (x *AiCHat) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiCHat.ProtoReflect.Descriptor instead.
func (*AiCHat) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{0}
}

func (x *AiCHat) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_ai_proto protoreflect.FileDescriptor

var file_ai_proto_rawDesc = []byte{
	0x0a, 0x08, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x41, 0x49, 0x22, 0x1c,
	0x0a, 0x06, 0x41, 0x69, 0x43, 0x48, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x2b, 0x0a, 0x09,
	0x41, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x43, 0x48, 0x61,
	0x74, 0x12, 0x0a, 0x2e, 0x41, 0x49, 0x2e, 0x41, 0x69, 0x43, 0x48, 0x61, 0x74, 0x1a, 0x0a, 0x2e,
	0x41, 0x49, 0x2e, 0x41, 0x69, 0x43, 0x48, 0x61, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ai_proto_rawDescOnce sync.Once
	file_ai_proto_rawDescData = file_ai_proto_rawDesc
)

func file_ai_proto_rawDescGZIP() []byte {
	file_ai_proto_rawDescOnce.Do(func() {
		file_ai_proto_rawDescData = protoimpl.X.CompressGZIP(file_ai_proto_rawDescData)
	})
	return file_ai_proto_rawDescData
}

var file_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ai_proto_goTypes = []interface{}{
	(*AiCHat)(nil), // 0: AI.AiCHat
}
var file_ai_proto_depIdxs = []int32{
	0, // 0: AI.AiService.CHat:input_type -> AI.AiCHat
	0, // 1: AI.AiService.CHat:output_type -> AI.AiCHat
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ai_proto_init() }
func file_ai_proto_init() {
	if File_ai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiCHat); i {
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
			RawDescriptor: file_ai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ai_proto_goTypes,
		DependencyIndexes: file_ai_proto_depIdxs,
		MessageInfos:      file_ai_proto_msgTypes,
	}.Build()
	File_ai_proto = out.File
	file_ai_proto_rawDesc = nil
	file_ai_proto_goTypes = nil
	file_ai_proto_depIdxs = nil
}
