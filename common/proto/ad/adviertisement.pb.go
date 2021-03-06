// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.10.0
// source: ad/adviertisement.proto

package ad

import (
	proto "github.com/golang/protobuf/proto"
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

type Adviertisement struct {
	state         protoimpl.MessageState  `xorm:"-"`
	sizeCache     protoimpl.SizeCache     `xorm:"-"`
	unknownFields protoimpl.UnknownFields `xorm:"-"`

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" xorm:" 'id' not null pk autoincr "`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" xorm:" 'name' comment('时间') VARCHAR(255)" `
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" xorm:" 'title' comment('合计') VARCHAR(50)"`
	Tel   string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty" xorm:" 'tel' comment('uid') "`
	Type  uint32 `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty" xorm:" 'Type' comment('uid') "`
}

func (x *Adviertisement) Reset() {
	*x = Adviertisement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_adviertisement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adviertisement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adviertisement) ProtoMessage() {}

func (x *Adviertisement) ProtoReflect() protoreflect.Message {
	mi := &file_ad_adviertisement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adviertisement.ProtoReflect.Descriptor instead.
func (*Adviertisement) Descriptor() ([]byte, []int) {
	return file_ad_adviertisement_proto_rawDescGZIP(), []int{0}
}

func (x *Adviertisement) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Adviertisement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Adviertisement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Adviertisement) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *Adviertisement) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize uint32 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PageNum  uint32 `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	Total    uint32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_adviertisement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ad_adviertisement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_ad_adviertisement_proto_rawDescGZIP(), []int{1}
}

func (x *PageInfo) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageInfo) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageInfo) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type QueryByExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adviertisement *Adviertisement `protobuf:"bytes,1,opt,name=adviertisement,proto3" json:"adviertisement,omitempty"`
	PageInfo       *PageInfo       `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
}

func (x *QueryByExampleRequest) Reset() {
	*x = QueryByExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_adviertisement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryByExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryByExampleRequest) ProtoMessage() {}

func (x *QueryByExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ad_adviertisement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryByExampleRequest.ProtoReflect.Descriptor instead.
func (*QueryByExampleRequest) Descriptor() ([]byte, []int) {
	return file_ad_adviertisement_proto_rawDescGZIP(), []int{2}
}

func (x *QueryByExampleRequest) GetAdviertisement() *Adviertisement {
	if x != nil {
		return x.Adviertisement
	}
	return nil
}

func (x *QueryByExampleRequest) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

type QueryByExampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo       *PageInfo         `protobuf:"bytes,1,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Adviertisement []*Adviertisement `protobuf:"bytes,2,rep,name=adviertisement,proto3" json:"adviertisement,omitempty"`
}

func (x *QueryByExampleResponse) Reset() {
	*x = QueryByExampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_adviertisement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryByExampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryByExampleResponse) ProtoMessage() {}

func (x *QueryByExampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ad_adviertisement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryByExampleResponse.ProtoReflect.Descriptor instead.
func (*QueryByExampleResponse) Descriptor() ([]byte, []int) {
	return file_ad_adviertisement_proto_rawDescGZIP(), []int{3}
}

func (x *QueryByExampleResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *QueryByExampleResponse) GetAdviertisement() []*Adviertisement {
	if x != nil {
		return x.Adviertisement
	}
	return nil
}

var File_ad_adviertisement_proto protoreflect.FileDescriptor

var file_ad_adviertisement_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x64, 0x2f, 0x61, 0x64, 0x76, 0x69, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x61, 0x64, 0x22, 0x70, 0x0a,
	0x0e, 0x41, 0x64, 0x76, 0x69, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x56, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7d, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x0e, 0x61, 0x64, 0x76, 0x69, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64,
	0x76, 0x69, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x61, 0x64,
	0x76, 0x69, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x7e, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x0e, 0x61, 0x64,
	0x76, 0x69, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x64, 0x2e, 0x41, 0x64, 0x76, 0x69, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x61, 0x64, 0x76, 0x69, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x67, 0x0a, 0x15, 0x41, 0x64, 0x76, 0x69, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x64, 0x76, 0x69, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x07, 0x5a, 0x05, 0x61, 0x64, 0x3b, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ad_adviertisement_proto_rawDescOnce sync.Once
	file_ad_adviertisement_proto_rawDescData = file_ad_adviertisement_proto_rawDesc
)

func file_ad_adviertisement_proto_rawDescGZIP() []byte {
	file_ad_adviertisement_proto_rawDescOnce.Do(func() {
		file_ad_adviertisement_proto_rawDescData = protoimpl.X.CompressGZIP(file_ad_adviertisement_proto_rawDescData)
	})
	return file_ad_adviertisement_proto_rawDescData
}

var file_ad_adviertisement_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ad_adviertisement_proto_goTypes = []interface{}{
	(*Adviertisement)(nil),         // 0: ad.Adviertisement
	(*PageInfo)(nil),               // 1: ad.PageInfo
	(*QueryByExampleRequest)(nil),  // 2: ad.QueryByExampleRequest
	(*QueryByExampleResponse)(nil), // 3: ad.QueryByExampleResponse
}
var file_ad_adviertisement_proto_depIdxs = []int32{
	0, // 0: ad.QueryByExampleRequest.adviertisement:type_name -> ad.Adviertisement
	1, // 1: ad.QueryByExampleRequest.pageInfo:type_name -> ad.PageInfo
	1, // 2: ad.QueryByExampleResponse.pageInfo:type_name -> ad.PageInfo
	0, // 3: ad.QueryByExampleResponse.adviertisement:type_name -> ad.Adviertisement
	2, // 4: ad.AdviertisementService.QueryAdviertisement:input_type -> ad.QueryByExampleRequest
	3, // 5: ad.AdviertisementService.QueryAdviertisement:output_type -> ad.QueryByExampleResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ad_adviertisement_proto_init() }
func file_ad_adviertisement_proto_init() {
	if File_ad_adviertisement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ad_adviertisement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adviertisement); i {
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
		file_ad_adviertisement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_ad_adviertisement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryByExampleRequest); i {
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
		file_ad_adviertisement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryByExampleResponse); i {
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
			RawDescriptor: file_ad_adviertisement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ad_adviertisement_proto_goTypes,
		DependencyIndexes: file_ad_adviertisement_proto_depIdxs,
		MessageInfos:      file_ad_adviertisement_proto_msgTypes,
	}.Build()
	File_ad_adviertisement_proto = out.File
	file_ad_adviertisement_proto_rawDesc = nil
	file_ad_adviertisement_proto_goTypes = nil
	file_ad_adviertisement_proto_depIdxs = nil
}
