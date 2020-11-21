// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: weight_tracker.proto

package weighttracker

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

type AddRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *AddRecordRequest) Reset() {
	*x = AddRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weight_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordRequest) ProtoMessage() {}

func (x *AddRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_weight_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordRequest.ProtoReflect.Descriptor instead.
func (*AddRecordRequest) Descriptor() ([]byte, []int) {
	return file_weight_tracker_proto_rawDescGZIP(), []int{0}
}

func (x *AddRecordRequest) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type AddRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *AddRecordResponse) Reset() {
	*x = AddRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weight_tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordResponse) ProtoMessage() {}

func (x *AddRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weight_tracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordResponse.ProtoReflect.Descriptor instead.
func (*AddRecordResponse) Descriptor() ([]byte, []int) {
	return file_weight_tracker_proto_rawDescGZIP(), []int{1}
}

func (x *AddRecordResponse) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weight float32 `protobuf:"fixed32,1,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weight_tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_weight_tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_weight_tracker_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_weight_tracker_proto protoreflect.FileDescriptor

var file_weight_tracker_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x20, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x32, 0x43, 0x0a, 0x0d, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x30, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x67, 0x6f,
	0x2d, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weight_tracker_proto_rawDescOnce sync.Once
	file_weight_tracker_proto_rawDescData = file_weight_tracker_proto_rawDesc
)

func file_weight_tracker_proto_rawDescGZIP() []byte {
	file_weight_tracker_proto_rawDescOnce.Do(func() {
		file_weight_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_weight_tracker_proto_rawDescData)
	})
	return file_weight_tracker_proto_rawDescData
}

var file_weight_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_weight_tracker_proto_goTypes = []interface{}{
	(*AddRecordRequest)(nil),  // 0: AddRecordRequest
	(*AddRecordResponse)(nil), // 1: AddRecordResponse
	(*Record)(nil),            // 2: Record
}
var file_weight_tracker_proto_depIdxs = []int32{
	2, // 0: AddRecordRequest.record:type_name -> Record
	2, // 1: AddRecordResponse.record:type_name -> Record
	0, // 2: WeightTracker.AddRecord:input_type -> AddRecordRequest
	1, // 3: WeightTracker.AddRecord:output_type -> AddRecordResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_weight_tracker_proto_init() }
func file_weight_tracker_proto_init() {
	if File_weight_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weight_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordRequest); i {
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
		file_weight_tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordResponse); i {
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
		file_weight_tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
			RawDescriptor: file_weight_tracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weight_tracker_proto_goTypes,
		DependencyIndexes: file_weight_tracker_proto_depIdxs,
		MessageInfos:      file_weight_tracker_proto_msgTypes,
	}.Build()
	File_weight_tracker_proto = out.File
	file_weight_tracker_proto_rawDesc = nil
	file_weight_tracker_proto_goTypes = nil
	file_weight_tracker_proto_depIdxs = nil
}
