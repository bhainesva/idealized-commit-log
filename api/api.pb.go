// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: api.proto

package api

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

type StartJobRequest_SortType int32

const (
	StartJobRequest_HARDCODED  StartJobRequest_SortType = 0
	StartJobRequest_RAW        StartJobRequest_SortType = 1
	StartJobRequest_NET        StartJobRequest_SortType = 2
	StartJobRequest_IMPORTANCE StartJobRequest_SortType = 3
)

// Enum value maps for StartJobRequest_SortType.
var (
	StartJobRequest_SortType_name = map[int32]string{
		0: "HARDCODED",
		1: "RAW",
		2: "NET",
		3: "IMPORTANCE",
	}
	StartJobRequest_SortType_value = map[string]int32{
		"HARDCODED":  0,
		"RAW":        1,
		"NET":        2,
		"IMPORTANCE": 3,
	}
)

func (x StartJobRequest_SortType) Enum() *StartJobRequest_SortType {
	p := new(StartJobRequest_SortType)
	*p = x
	return p
}

func (x StartJobRequest_SortType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartJobRequest_SortType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (StartJobRequest_SortType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x StartJobRequest_SortType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartJobRequest_SortType.Descriptor instead.
func (StartJobRequest_SortType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 0}
}

type StartJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tests []string                 `protobuf:"bytes,1,rep,name=tests,proto3" json:"tests,omitempty"`
	Pkg   string                   `protobuf:"bytes,2,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Sort  StartJobRequest_SortType `protobuf:"varint,3,opt,name=sort,proto3,enum=StartJobRequest_SortType" json:"sort,omitempty"`
}

func (x *StartJobRequest) Reset() {
	*x = StartJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobRequest) ProtoMessage() {}

func (x *StartJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobRequest.ProtoReflect.Descriptor instead.
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *StartJobRequest) GetTests() []string {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *StartJobRequest) GetPkg() string {
	if x != nil {
		return x.Pkg
	}
	return ""
}

func (x *StartJobRequest) GetSort() StartJobRequest_SortType {
	if x != nil {
		return x.Sort
	}
	return StartJobRequest_HARDCODED
}

type StartJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartJobResponse) Reset() {
	*x = StartJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobResponse) ProtoMessage() {}

func (x *StartJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartJobResponse.ProtoReflect.Descriptor instead.
func (*StartJobResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *StartJobResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CheckoutFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files *FileMap `protobuf:"bytes,1,opt,name=files,proto3" json:"files,omitempty"`
}

func (x *CheckoutFilesRequest) Reset() {
	*x = CheckoutFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutFilesRequest) ProtoMessage() {}

func (x *CheckoutFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutFilesRequest.ProtoReflect.Descriptor instead.
func (*CheckoutFilesRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *CheckoutFilesRequest) GetFiles() *FileMap {
	if x != nil {
		return x.Files
	}
	return nil
}

type JobStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Complete bool        `protobuf:"varint,1,opt,name=complete,proto3" json:"complete,omitempty"`
	Details  string      `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Error    string      `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Results  *JobResults `protobuf:"bytes,4,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *JobStatusResponse) Reset() {
	*x = JobStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusResponse) ProtoMessage() {}

func (x *JobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusResponse.ProtoReflect.Descriptor instead.
func (*JobStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *JobStatusResponse) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

func (x *JobStatusResponse) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *JobStatusResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *JobStatusResponse) GetResults() *JobResults {
	if x != nil {
		return x.Results
	}
	return nil
}

type JobResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tests []string   `protobuf:"bytes,1,rep,name=tests,proto3" json:"tests,omitempty"`
	Files []*FileMap `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *JobResults) Reset() {
	*x = JobResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResults) ProtoMessage() {}

func (x *JobResults) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResults.ProtoReflect.Descriptor instead.
func (*JobResults) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *JobResults) GetTests() []string {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *JobResults) GetFiles() []*FileMap {
	if x != nil {
		return x.Files
	}
	return nil
}

type FileMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files map[string][]byte `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FileMap) Reset() {
	*x = FileMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMap) ProtoMessage() {}

func (x *FileMap) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMap.ProtoReflect.Descriptor instead.
func (*FileMap) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *FileMap) GetFiles() map[string][]byte {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0f,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6b, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x6b, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x3b, 0x0a, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x52, 0x44, 0x43, 0x4f, 0x44, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x57, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45,
	0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x4e, 0x43,
	0x45, 0x10, 0x03, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22,
	0x86, 0x01, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x25, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x07,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(StartJobRequest_SortType)(0), // 0: StartJobRequest.SortType
	(*StartJobRequest)(nil),       // 1: StartJobRequest
	(*StartJobResponse)(nil),      // 2: StartJobResponse
	(*CheckoutFilesRequest)(nil),  // 3: CheckoutFilesRequest
	(*JobStatusResponse)(nil),     // 4: JobStatusResponse
	(*JobResults)(nil),            // 5: JobResults
	(*FileMap)(nil),               // 6: FileMap
	nil,                           // 7: FileMap.FilesEntry
}
var file_api_proto_depIdxs = []int32{
	0, // 0: StartJobRequest.sort:type_name -> StartJobRequest.SortType
	6, // 1: CheckoutFilesRequest.files:type_name -> FileMap
	5, // 2: JobStatusResponse.results:type_name -> JobResults
	6, // 3: JobResults.files:type_name -> FileMap
	7, // 4: FileMap.files:type_name -> FileMap.FilesEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartJobResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutFilesRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatusResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResults); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMap); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
