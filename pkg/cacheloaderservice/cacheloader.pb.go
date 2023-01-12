// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: cacheloader.proto

package cacheloaderservice

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cacheloader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_cacheloader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_cacheloader_proto_rawDescGZIP(), []int{0}
}

func (x *Empty) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UrlShortLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short    string `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *UrlShortLink) Reset() {
	*x = UrlShortLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cacheloader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlShortLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlShortLink) ProtoMessage() {}

func (x *UrlShortLink) ProtoReflect() protoreflect.Message {
	mi := &file_cacheloader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlShortLink.ProtoReflect.Descriptor instead.
func (*UrlShortLink) Descriptor() ([]byte, []int) {
	return file_cacheloader_proto_rawDescGZIP(), []int{1}
}

func (x *UrlShortLink) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *UrlShortLink) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

var File_cacheloader_proto protoreflect.FileDescriptor

var file_cacheloader_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x63, 0x68, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x17, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x40, 0x0a, 0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x32, 0x60, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x19, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x61, 0x6b, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cacheloader_proto_rawDescOnce sync.Once
	file_cacheloader_proto_rawDescData = file_cacheloader_proto_rawDesc
)

func file_cacheloader_proto_rawDescGZIP() []byte {
	file_cacheloader_proto_rawDescOnce.Do(func() {
		file_cacheloader_proto_rawDescData = protoimpl.X.CompressGZIP(file_cacheloader_proto_rawDescData)
	})
	return file_cacheloader_proto_rawDescData
}

var file_cacheloader_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cacheloader_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: CacheLoaderService.Empty
	(*UrlShortLink)(nil), // 1: CacheLoaderService.UrlShortLink
}
var file_cacheloader_proto_depIdxs = []int32{
	0, // 0: CacheLoaderService.UrlShortlinkCache.GetItems:input_type -> CacheLoaderService.Empty
	1, // 1: CacheLoaderService.UrlShortlinkCache.GetItems:output_type -> CacheLoaderService.UrlShortLink
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cacheloader_proto_init() }
func file_cacheloader_proto_init() {
	if File_cacheloader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cacheloader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_cacheloader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlShortLink); i {
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
			RawDescriptor: file_cacheloader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cacheloader_proto_goTypes,
		DependencyIndexes: file_cacheloader_proto_depIdxs,
		MessageInfos:      file_cacheloader_proto_msgTypes,
	}.Build()
	File_cacheloader_proto = out.File
	file_cacheloader_proto_rawDesc = nil
	file_cacheloader_proto_goTypes = nil
	file_cacheloader_proto_depIdxs = nil
}