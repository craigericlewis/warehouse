// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: warehouse.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Language int32

const (
	Language_NONE    Language = 0
	Language_PYTHON2 Language = 1
	Language_PYTHON3 Language = 2
	Language_GO      Language = 3
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "NONE",
		1: "PYTHON2",
		2: "PYTHON3",
		3: "GO",
	}
	Language_value = map[string]int32{
		"NONE":    0,
		"PYTHON2": 1,
		"PYTHON3": 2,
		"GO":      3,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_warehouse_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_warehouse_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{0}
}

type GetPackageInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string   `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	Language    Language `protobuf:"varint,2,opt,name=language,proto3,enum=api.Language" json:"language,omitempty"`
}

func (x *GetPackageInfoRequest) Reset() {
	*x = GetPackageInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackageInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackageInfoRequest) ProtoMessage() {}

func (x *GetPackageInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackageInfoRequest.ProtoReflect.Descriptor instead.
func (*GetPackageInfoRequest) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *GetPackageInfoRequest) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *GetPackageInfoRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_NONE
}

type PackageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Language          Language       `protobuf:"varint,2,opt,name=language,proto3,enum=api.Language" json:"language,omitempty"`
	VersionInfo       []*VersionInfo `protobuf:"bytes,3,rep,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Description       string         `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Urls              *Urls          `protobuf:"bytes,5,opt,name=urls,proto3" json:"urls,omitempty"`
	DependenciesCount int64          `protobuf:"varint,6,opt,name=dependencies_count,json=dependenciesCount,proto3" json:"dependencies_count,omitempty"`
}

func (x *PackageInfo) Reset() {
	*x = PackageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageInfo) ProtoMessage() {}

func (x *PackageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageInfo.ProtoReflect.Descriptor instead.
func (*PackageInfo) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *PackageInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackageInfo) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_NONE
}

func (x *PackageInfo) GetVersionInfo() []*VersionInfo {
	if x != nil {
		return x.VersionInfo
	}
	return nil
}

func (x *PackageInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PackageInfo) GetUrls() *Urls {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *PackageInfo) GetDependenciesCount() int64 {
	if x != nil {
		return x.DependenciesCount
	}
	return 0
}

type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	PackageSize       string `protobuf:"bytes,2,opt,name=package_size,json=packageSize,proto3" json:"package_size,omitempty"`
	PackageZippedSize string `protobuf:"bytes,3,opt,name=package_zipped_size,json=packageZippedSize,proto3" json:"package_zipped_size,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *VersionInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionInfo) GetPackageSize() string {
	if x != nil {
		return x.PackageSize
	}
	return ""
}

func (x *VersionInfo) GetPackageZippedSize() string {
	if x != nil {
		return x.PackageZippedSize
	}
	return ""
}

type Urls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GithubLink string `protobuf:"bytes,1,opt,name=github_link,json=githubLink,proto3" json:"github_link,omitempty"`
	PyPiLink   string `protobuf:"bytes,2,opt,name=py_pi_link,json=pyPiLink,proto3" json:"py_pi_link,omitempty"`
}

func (x *Urls) Reset() {
	*x = Urls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Urls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Urls) ProtoMessage() {}

func (x *Urls) ProtoReflect() protoreflect.Message {
	mi := &file_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Urls.ProtoReflect.Descriptor instead.
func (*Urls) Descriptor() ([]byte, []int) {
	return file_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *Urls) GetGithubLink() string {
	if x != nil {
		return x.GithubLink
	}
	return ""
}

func (x *Urls) GetPyPiLink() string {
	if x != nil {
		return x.PyPiLink
	}
	return ""
}

var File_warehouse_proto protoreflect.FileDescriptor

var file_warehouse_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x65, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0xf1, 0x01,
	0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x0c,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x7a, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x7a, 0x69, 0x70, 0x70, 0x65, 0x64, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5a, 0x69, 0x70, 0x70, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x45, 0x0a,
	0x04, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x0a, 0x70, 0x79, 0x5f, 0x70, 0x69, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x79, 0x50, 0x69,
	0x4c, 0x69, 0x6e, 0x6b, 0x2a, 0x36, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x59,
	0x54, 0x48, 0x4f, 0x4e, 0x32, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x59, 0x54, 0x48, 0x4f,
	0x4e, 0x33, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4f, 0x10, 0x03, 0x32, 0x50, 0x0a, 0x0e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x05,
	0x5a, 0x03, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_warehouse_proto_rawDescOnce sync.Once
	file_warehouse_proto_rawDescData = file_warehouse_proto_rawDesc
)

func file_warehouse_proto_rawDescGZIP() []byte {
	file_warehouse_proto_rawDescOnce.Do(func() {
		file_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_warehouse_proto_rawDescData)
	})
	return file_warehouse_proto_rawDescData
}

var file_warehouse_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_warehouse_proto_goTypes = []interface{}{
	(Language)(0),                 // 0: api.Language
	(*GetPackageInfoRequest)(nil), // 1: api.GetPackageInfoRequest
	(*PackageInfo)(nil),           // 2: api.PackageInfo
	(*VersionInfo)(nil),           // 3: api.VersionInfo
	(*Urls)(nil),                  // 4: api.Urls
}
var file_warehouse_proto_depIdxs = []int32{
	0, // 0: api.GetPackageInfoRequest.language:type_name -> api.Language
	0, // 1: api.PackageInfo.language:type_name -> api.Language
	3, // 2: api.PackageInfo.version_info:type_name -> api.VersionInfo
	4, // 3: api.PackageInfo.urls:type_name -> api.Urls
	1, // 4: api.PackageService.GetPackageInfo:input_type -> api.GetPackageInfoRequest
	2, // 5: api.PackageService.GetPackageInfo:output_type -> api.PackageInfo
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_warehouse_proto_init() }
func file_warehouse_proto_init() {
	if File_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackageInfoRequest); i {
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
		file_warehouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageInfo); i {
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
		file_warehouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
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
		file_warehouse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Urls); i {
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
			RawDescriptor: file_warehouse_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_warehouse_proto_goTypes,
		DependencyIndexes: file_warehouse_proto_depIdxs,
		EnumInfos:         file_warehouse_proto_enumTypes,
		MessageInfos:      file_warehouse_proto_msgTypes,
	}.Build()
	File_warehouse_proto = out.File
	file_warehouse_proto_rawDesc = nil
	file_warehouse_proto_goTypes = nil
	file_warehouse_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PackageServiceClient is the client API for PackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PackageServiceClient interface {
	GetPackageInfo(ctx context.Context, in *GetPackageInfoRequest, opts ...grpc.CallOption) (*PackageInfo, error)
}

type packageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageServiceClient(cc grpc.ClientConnInterface) PackageServiceClient {
	return &packageServiceClient{cc}
}

func (c *packageServiceClient) GetPackageInfo(ctx context.Context, in *GetPackageInfoRequest, opts ...grpc.CallOption) (*PackageInfo, error) {
	out := new(PackageInfo)
	err := c.cc.Invoke(ctx, "/api.PackageService/GetPackageInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageServiceServer is the server API for PackageService service.
type PackageServiceServer interface {
	GetPackageInfo(context.Context, *GetPackageInfoRequest) (*PackageInfo, error)
}

// UnimplementedPackageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPackageServiceServer struct {
}

func (*UnimplementedPackageServiceServer) GetPackageInfo(context.Context, *GetPackageInfoRequest) (*PackageInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageInfo not implemented")
}

func RegisterPackageServiceServer(s *grpc.Server, srv PackageServiceServer) {
	s.RegisterService(&_PackageService_serviceDesc, srv)
}

func _PackageService_GetPackageInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).GetPackageInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PackageService/GetPackageInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).GetPackageInfo(ctx, req.(*GetPackageInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PackageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PackageService",
	HandlerType: (*PackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPackageInfo",
			Handler:    _PackageService_GetPackageInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse.proto",
}
