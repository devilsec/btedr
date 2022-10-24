// Protobuf defining tasks for agents
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative */*.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: taskpb/task.proto

package taskpb

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

// Type of task
type Type int32

const (
	Type_Ping Type = 0
	Type_Ls   Type = 1
	Type_Exec Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Ping",
		1: "Ls",
		2: "Exec",
	}
	Type_value = map[string]int32{
		"Ping": 0,
		"Ls":   1,
		"Exec": 2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_taskpb_task_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_taskpb_task_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_taskpb_task_proto_rawDescGZIP(), []int{0}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type Type   `protobuf:"varint,2,opt,name=type,proto3,enum=taskpb.Type" json:"type,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_taskpb_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_Ping
}

type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration uint32 `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_taskpb_task_proto_rawDescGZIP(), []int{1}
}

func (x *PingReq) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

var File_taskpb_task_proto protoreflect.FileDescriptor

var file_taskpb_task_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x22, 0x38, 0x0a, 0x04, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x22, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x4c, 0x73, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x10, 0x02,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x6c, 0x73, 0x65, 0x63, 0x2f, 0x62, 0x74, 0x65, 0x64, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_taskpb_task_proto_rawDescOnce sync.Once
	file_taskpb_task_proto_rawDescData = file_taskpb_task_proto_rawDesc
)

func file_taskpb_task_proto_rawDescGZIP() []byte {
	file_taskpb_task_proto_rawDescOnce.Do(func() {
		file_taskpb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_taskpb_task_proto_rawDescData)
	})
	return file_taskpb_task_proto_rawDescData
}

var file_taskpb_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_taskpb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_taskpb_task_proto_goTypes = []interface{}{
	(Type)(0),       // 0: taskpb.Type
	(*Task)(nil),    // 1: taskpb.Task
	(*PingReq)(nil), // 2: taskpb.PingReq
}
var file_taskpb_task_proto_depIdxs = []int32{
	0, // 0: taskpb.Task.type:type_name -> taskpb.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_taskpb_task_proto_init() }
func file_taskpb_task_proto_init() {
	if File_taskpb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_taskpb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_taskpb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingReq); i {
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
			RawDescriptor: file_taskpb_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_taskpb_task_proto_goTypes,
		DependencyIndexes: file_taskpb_task_proto_depIdxs,
		EnumInfos:         file_taskpb_task_proto_enumTypes,
		MessageInfos:      file_taskpb_task_proto_msgTypes,
	}.Build()
	File_taskpb_task_proto = out.File
	file_taskpb_task_proto_rawDesc = nil
	file_taskpb_task_proto_goTypes = nil
	file_taskpb_task_proto_depIdxs = nil
}
