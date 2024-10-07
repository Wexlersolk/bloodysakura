// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: data.proto

package crawler

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

type CrawlerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrawlerID  int32  `protobuf:"varint,1,opt,name=CrawlerID,proto3" json:"CrawlerID,omitempty"`
	VisitUrl   string `protobuf:"bytes,2,opt,name=VisitUrl,proto3" json:"VisitUrl,omitempty"`
	WantedText string `protobuf:"bytes,3,opt,name=WantedText,proto3" json:"WantedText,omitempty"`
	GeckoPort  int32  `protobuf:"varint,4,opt,name=GeckoPort,proto3" json:"GeckoPort,omitempty"`
	GeckoPath  string `protobuf:"bytes,5,opt,name=GeckoPath,proto3" json:"GeckoPath,omitempty"`
}

func (x *CrawlerData) Reset() {
	*x = CrawlerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerData) ProtoMessage() {}

func (x *CrawlerData) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlerData.ProtoReflect.Descriptor instead.
func (*CrawlerData) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *CrawlerData) GetCrawlerID() int32 {
	if x != nil {
		return x.CrawlerID
	}
	return 0
}

func (x *CrawlerData) GetVisitUrl() string {
	if x != nil {
		return x.VisitUrl
	}
	return ""
}

func (x *CrawlerData) GetWantedText() string {
	if x != nil {
		return x.WantedText
	}
	return ""
}

func (x *CrawlerData) GetGeckoPort() int32 {
	if x != nil {
		return x.GeckoPort
	}
	return 0
}

func (x *CrawlerData) GetGeckoPath() string {
	if x != nil {
		return x.GeckoPath
	}
	return ""
}

type CreateCrawlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisitUrl   string `protobuf:"bytes,1,opt,name=VisitUrl,proto3" json:"VisitUrl,omitempty"`
	WantedText string `protobuf:"bytes,2,opt,name=WantedText,proto3" json:"WantedText,omitempty"`
	GeckoPort  int32  `protobuf:"varint,3,opt,name=GeckoPort,proto3" json:"GeckoPort,omitempty"`
	GeckoPath  string `protobuf:"bytes,4,opt,name=GeckoPath,proto3" json:"GeckoPath,omitempty"`
}

func (x *CreateCrawlerRequest) Reset() {
	*x = CreateCrawlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCrawlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCrawlerRequest) ProtoMessage() {}

func (x *CreateCrawlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCrawlerRequest.ProtoReflect.Descriptor instead.
func (*CreateCrawlerRequest) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCrawlerRequest) GetVisitUrl() string {
	if x != nil {
		return x.VisitUrl
	}
	return ""
}

func (x *CreateCrawlerRequest) GetWantedText() string {
	if x != nil {
		return x.WantedText
	}
	return ""
}

func (x *CreateCrawlerRequest) GetGeckoPort() int32 {
	if x != nil {
		return x.GeckoPort
	}
	return 0
}

func (x *CreateCrawlerRequest) GetGeckoPath() string {
	if x != nil {
		return x.GeckoPath
	}
	return ""
}

type CreateCrawlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisitUrl string `protobuf:"bytes,1,opt,name=VisitUrl,proto3" json:"VisitUrl,omitempty"`
}

func (x *CreateCrawlerResponse) Reset() {
	*x = CreateCrawlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCrawlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCrawlerResponse) ProtoMessage() {}

func (x *CreateCrawlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCrawlerResponse.ProtoReflect.Descriptor instead.
func (*CreateCrawlerResponse) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCrawlerResponse) GetVisitUrl() string {
	if x != nil {
		return x.VisitUrl
	}
	return ""
}

type GetCrawlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrawlerID int32 `protobuf:"varint,1,opt,name=CrawlerID,proto3" json:"CrawlerID,omitempty"`
}

func (x *GetCrawlerRequest) Reset() {
	*x = GetCrawlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrawlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrawlerRequest) ProtoMessage() {}

func (x *GetCrawlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrawlerRequest.ProtoReflect.Descriptor instead.
func (*GetCrawlerRequest) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetCrawlerRequest) GetCrawlerID() int32 {
	if x != nil {
		return x.CrawlerID
	}
	return 0
}

type GetCrawlerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crawlers []*CrawlerData `protobuf:"bytes,1,rep,name=crawlers,proto3" json:"crawlers,omitempty"`
}

func (x *GetCrawlerResponse) Reset() {
	*x = GetCrawlerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCrawlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCrawlerResponse) ProtoMessage() {}

func (x *GetCrawlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCrawlerResponse.ProtoReflect.Descriptor instead.
func (*GetCrawlerResponse) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *GetCrawlerResponse) GetCrawlers() []*CrawlerData {
	if x != nil {
		return x.Crawlers
	}
	return nil
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a,
	0x0b, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x57, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x56,
	0x69, 0x73, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x56,
	0x69, 0x73, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x57, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x57, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x47, 0x65, 0x63, 0x6b,
	0x6f, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x65, 0x63, 0x6b, 0x6f, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x33, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x56, 0x69, 0x73, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x56, 0x69, 0x73, 0x69, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x73, 0x32, 0x8b, 0x01, 0x0a, 0x0e,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12,
	0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x65, 0x78, 0x6c, 0x65, 0x72, 0x73, 0x6f,
	0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_data_proto_goTypes = []any{
	(*CrawlerData)(nil),           // 0: CrawlerData
	(*CreateCrawlerRequest)(nil),  // 1: CreateCrawlerRequest
	(*CreateCrawlerResponse)(nil), // 2: CreateCrawlerResponse
	(*GetCrawlerRequest)(nil),     // 3: GetCrawlerRequest
	(*GetCrawlerResponse)(nil),    // 4: GetCrawlerResponse
}
var file_data_proto_depIdxs = []int32{
	0, // 0: GetCrawlerResponse.crawlers:type_name -> CrawlerData
	1, // 1: CrawlerService.CreateCrawler:input_type -> CreateCrawlerRequest
	3, // 2: CrawlerService.GetCrawler:input_type -> GetCrawlerRequest
	2, // 3: CrawlerService.CreateCrawler:output_type -> CreateCrawlerResponse
	4, // 4: CrawlerService.GetCrawler:output_type -> GetCrawlerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CrawlerData); i {
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
		file_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCrawlerRequest); i {
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
		file_data_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCrawlerResponse); i {
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
		file_data_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCrawlerRequest); i {
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
		file_data_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCrawlerResponse); i {
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
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}
