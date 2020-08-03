// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/test.proto

package proto

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

// Status 枚举状态
type Test_Status int32

const (
	Test_OK   Test_Status = 0
	Test_FAIL Test_Status = 1
)

// Enum value maps for Test_Status.
var (
	Test_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	Test_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x Test_Status) Enum() *Test_Status {
	p := new(Test_Status)
	*p = x
	return p
}

func (x Test_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Test_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_test_proto_enumTypes[0].Descriptor()
}

func (Test_Status) Type() protoreflect.EnumType {
	return &file_proto_test_proto_enumTypes[0]
}

func (x Test_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Test_Status.Descriptor instead.
func (Test_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_test_proto_rawDescGZIP(), []int{0, 0}
}

// Test 测试
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age    int32             `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Count  int64             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Money  float64           `protobuf:"fixed64,3,opt,name=money,proto3" json:"money,omitempty"`
	Score  float32           `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
	Name   string            `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Fat    bool              `protobuf:"varint,6,opt,name=fat,proto3" json:"fat,omitempty"`
	Char   []byte            `protobuf:"bytes,7,opt,name=char,proto3" json:"char,omitempty"`
	Status Test_Status       `protobuf:"varint,8,opt,name=status,proto3,enum=proto.Test_Status" json:"status,omitempty"`
	Child  *Test_Child       `protobuf:"bytes,9,opt,name=child,proto3" json:"child,omitempty"`
	Dict   map[string]string `protobuf:"bytes,10,rep,name=dict,proto3" json:"dict,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_proto_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Test) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Test) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *Test) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Test) GetFat() bool {
	if x != nil {
		return x.Fat
	}
	return false
}

func (x *Test) GetChar() []byte {
	if x != nil {
		return x.Char
	}
	return nil
}

func (x *Test) GetStatus() Test_Status {
	if x != nil {
		return x.Status
	}
	return Test_OK
}

func (x *Test) GetChild() *Test_Child {
	if x != nil {
		return x.Child
	}
	return nil
}

func (x *Test) GetDict() map[string]string {
	if x != nil {
		return x.Dict
	}
	return nil
}

// Child 子结构
type Test_Child struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sex string `protobuf:"bytes,1,opt,name=sex,proto3" json:"sex,omitempty"`
}

func (x *Test_Child) Reset() {
	*x = Test_Child{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test_Child) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_Child) ProtoMessage() {}

func (x *Test_Child) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_Child.ProtoReflect.Descriptor instead.
func (*Test_Child) Descriptor() ([]byte, []int) {
	return file_proto_test_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Test_Child) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

var File_proto_test_proto protoreflect.FileDescriptor

var file_proto_test_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x66, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x68, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x68, 0x61, 0x72,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x05,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x05,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x69, 0x63, 0x74, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x69, 0x63, 0x74,
	0x1a, 0x19, 0x0a, 0x05, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x1a, 0x37, 0x0a, 0x09, 0x44,
	0x69, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01,
	0x32, 0x35, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x69, 0x6c, 0x69, 0x65, 0x2f, 0x6c, 0x69, 0x63,
	0x6f, 0x5f, 0x61, 0x6c, 0x6f, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_test_proto_rawDescOnce sync.Once
	file_proto_test_proto_rawDescData = file_proto_test_proto_rawDesc
)

func file_proto_test_proto_rawDescGZIP() []byte {
	file_proto_test_proto_rawDescOnce.Do(func() {
		file_proto_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_proto_rawDescData)
	})
	return file_proto_test_proto_rawDescData
}

var file_proto_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_test_proto_goTypes = []interface{}{
	(Test_Status)(0),   // 0: proto.Test.Status
	(*Test)(nil),       // 1: proto.Test
	(*Test_Child)(nil), // 2: proto.Test.Child
	nil,                // 3: proto.Test.DictEntry
}
var file_proto_test_proto_depIdxs = []int32{
	0, // 0: proto.Test.status:type_name -> proto.Test.Status
	2, // 1: proto.Test.child:type_name -> proto.Test.Child
	3, // 2: proto.Test.dict:type_name -> proto.Test.DictEntry
	1, // 3: proto.SearchService.Search:input_type -> proto.Test
	1, // 4: proto.SearchService.Search:output_type -> proto.Test
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_test_proto_init() }
func file_proto_test_proto_init() {
	if File_proto_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_proto_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test_Child); i {
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
			RawDescriptor: file_proto_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_test_proto_goTypes,
		DependencyIndexes: file_proto_test_proto_depIdxs,
		EnumInfos:         file_proto_test_proto_enumTypes,
		MessageInfos:      file_proto_test_proto_msgTypes,
	}.Build()
	File_proto_test_proto = out.File
	file_proto_test_proto_rawDesc = nil
	file_proto_test_proto_goTypes = nil
	file_proto_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/proto.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	Search(context.Context, *Test) (*Test, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) Search(context.Context, *Test) (*Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/test.proto",
}
