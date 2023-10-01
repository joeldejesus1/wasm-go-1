// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: main.proto

package subway

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

type Station int32

const (
	Station_A Station = 0
	Station_B Station = 1
	Station_C Station = 2
	Station_D Station = 3
)

// Enum value maps for Station.
var (
	Station_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
	}
	Station_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
	}
)

func (x Station) Enum() *Station {
	p := new(Station)
	*p = x
	return p
}

func (x Station) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Station) Descriptor() protoreflect.EnumDescriptor {
	return file_main_proto_enumTypes[0].Descriptor()
}

func (Station) Type() protoreflect.EnumType {
	return &file_main_proto_enumTypes[0]
}

func (x Station) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Station.Descriptor instead.
func (Station) EnumDescriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

type Segment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      Station `protobuf:"varint,1,opt,name=source,proto3,enum=subway.Station" json:"source,omitempty"`
	Destination Station `protobuf:"varint,2,opt,name=destination,proto3,enum=subway.Station" json:"destination,omitempty"`
}

func (x *Segment) Reset() {
	*x = Segment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Segment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Segment) ProtoMessage() {}

func (x *Segment) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Segment.ProtoReflect.Descriptor instead.
func (*Segment) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{0}
}

func (x *Segment) GetSource() Station {
	if x != nil {
		return x.Source
	}
	return Station_A
}

func (x *Segment) GetDestination() Station {
	if x != nil {
		return x.Destination
	}
	return Station_A
}

// [Segment; 2B?]
type DirectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segment *Segment `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
}

func (x *DirectionRequest) Reset() {
	*x = DirectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectionRequest) ProtoMessage() {}

func (x *DirectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectionRequest.ProtoReflect.Descriptor instead.
func (*DirectionRequest) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{1}
}

func (x *DirectionRequest) GetSegment() *Segment {
	if x != nil {
		return x.Segment
	}
	return nil
}

// what data you see as the client coming from the server
// <--[repeated; n segments][Segment_1; 2B][Segment_2; 2B][Segment_3; 2B]....[Segment_n; 2B]
type DirectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated means array of segments
	DirectionList []*Segment `protobuf:"bytes,1,rep,name=direction_list,json=directionList,proto3" json:"direction_list,omitempty"`
}

func (x *DirectionResponse) Reset() {
	*x = DirectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectionResponse) ProtoMessage() {}

func (x *DirectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectionResponse.ProtoReflect.Descriptor instead.
func (*DirectionResponse) Descriptor() ([]byte, []int) {
	return file_main_proto_rawDescGZIP(), []int{2}
}

func (x *DirectionResponse) GetDirectionList() []*Segment {
	if x != nil {
		return x.DirectionList
	}
	return nil
}

var File_main_proto protoreflect.FileDescriptor

var file_main_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x75,
	0x62, 0x77, 0x61, 0x79, 0x22, 0x65, 0x0a, 0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x27, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x73, 0x75, 0x62, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x73, 0x75, 0x62, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x10, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x11, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x77, 0x61, 0x79,
	0x2e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x25, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01,
	0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x44, 0x10, 0x03, 0x32, 0x52,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x77, 0x61, 0x79, 0x4d, 0x61, 0x70, 0x12, 0x45, 0x0a, 0x0c, 0x41,
	0x73, 0x6b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x73, 0x75,
	0x62, 0x77, 0x61, 0x79, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x77, 0x61, 0x79, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6f, 0x65, 0x6c, 0x64, 0x65, 0x6a, 0x65, 0x73, 0x75, 0x73, 0x31, 0x2f, 0x77, 0x61,
	0x73, 0x6d, 0x2d, 0x67, 0x6f, 0x2d, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75,
	0x62, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_proto_rawDescOnce sync.Once
	file_main_proto_rawDescData = file_main_proto_rawDesc
)

func file_main_proto_rawDescGZIP() []byte {
	file_main_proto_rawDescOnce.Do(func() {
		file_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_proto_rawDescData)
	})
	return file_main_proto_rawDescData
}

var file_main_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_main_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_main_proto_goTypes = []interface{}{
	(Station)(0),              // 0: subway.Station
	(*Segment)(nil),           // 1: subway.Segment
	(*DirectionRequest)(nil),  // 2: subway.DirectionRequest
	(*DirectionResponse)(nil), // 3: subway.DirectionResponse
}
var file_main_proto_depIdxs = []int32{
	0, // 0: subway.Segment.source:type_name -> subway.Station
	0, // 1: subway.Segment.destination:type_name -> subway.Station
	1, // 2: subway.DirectionRequest.segment:type_name -> subway.Segment
	1, // 3: subway.DirectionResponse.direction_list:type_name -> subway.Segment
	2, // 4: subway.SubwayMap.AskDirection:input_type -> subway.DirectionRequest
	3, // 5: subway.SubwayMap.AskDirection:output_type -> subway.DirectionResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_main_proto_init() }
func file_main_proto_init() {
	if File_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Segment); i {
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
		file_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectionRequest); i {
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
		file_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectionResponse); i {
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
			RawDescriptor: file_main_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_proto_goTypes,
		DependencyIndexes: file_main_proto_depIdxs,
		EnumInfos:         file_main_proto_enumTypes,
		MessageInfos:      file_main_proto_msgTypes,
	}.Build()
	File_main_proto = out.File
	file_main_proto_rawDesc = nil
	file_main_proto_goTypes = nil
	file_main_proto_depIdxs = nil
}