package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Arg1          string                 `protobuf:"bytes,2,opt,name=arg1,proto3" json:"arg1,omitempty"`
	Arg2          string                 `protobuf:"bytes,3,opt,name=arg2,proto3" json:"arg2,omitempty"`
	Operator      string                 `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (msg *TaskRequest) ProtoReflect() protoreflect.Message {
	metaInfo := &file_calculation_proto_msgTypes[0]
	if msg != nil {
		messageState := protoimpl.X.MessageStateOf(protoimpl.Pointer(msg))
		if messageState.LoadMessageInfo() == nil {
			messageState.StoreMessageInfo(metaInfo)
		}
		return messageState
	}
	return metaInfo.MessageOf(msg)
}

func (msg *TaskRequest) Reset() {
	*msg = TaskRequest{}
	metaInfo := &file_calculation_proto_msgTypes[0]
	messageState := protoimpl.X.MessageStateOf(protoimpl.Pointer(msg))
	messageState.StoreMessageInfo(metaInfo)
}

func (msg *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(msg)
}

func (*TaskRequest) ProtoMessage() {}

func (msg *TaskRequest) Descriptor() ([]byte, []int) {
	return file_calculation_proto_rawDescGZIP(), []int{0}
}

func (msg *TaskRequest) GetId() int32 {
	if msg != nil {
		return msg.Id
	}
	return 0
}

func (msg *TaskRequest) GetArg1() string {
	if msg != nil {
		return msg.Arg1
	}
	return ""
}

func (msg *TaskRequest) GetArg2() string {
	if msg != nil {
		return msg.Arg2
	}
	return ""
}

func (msg *TaskRequest) GetOperator() string {
	if msg != nil {
		return msg.Operator
	}
	return ""
}

type AgentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Result        float32                `protobuf:"fixed32,2,opt,name=result,proto3" json:"result,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (msg *AgentResponse) ProtoReflect() protoreflect.Message {
	metaInfo := &file_calculation_proto_msgTypes[1]
	if msg != nil {
		messageState := protoimpl.X.MessageStateOf(protoimpl.Pointer(msg))
		if messageState.LoadMessageInfo() == nil {
			messageState.StoreMessageInfo(metaInfo)
		}
		return messageState
	}
	return metaInfo.MessageOf(msg)
}

func (msg *AgentResponse) Reset() {
	*msg = AgentResponse{}
	metaInfo := &file_calculation_proto_msgTypes[1]
	messageState := protoimpl.X.MessageStateOf(protoimpl.Pointer(msg))
	messageState.StoreMessageInfo(metaInfo)
}

func (msg *AgentResponse) String() string {
	return protoimpl.X.MessageStringOf(msg)
}

func (*AgentResponse) ProtoMessage() {}

func (msg *AgentResponse) Descriptor() ([]byte, []int) {
	return file_calculation_proto_rawDescGZIP(), []int{1}
}

func (msg *AgentResponse) GetId() int32 {
	if msg != nil {
		return msg.Id
	}
	return 0
}

func (msg *AgentResponse) GetResult() float32 {
	if msg != nil {
		return msg.Result
	}
	return 0
}

func (msg *AgentResponse) GetError() string {
	if msg != nil {
		return msg.Error
	}
	return ""
}

var File_calculation_proto protoreflect.FileDescriptor

const file_calculation_proto_rawDesc = "" +
	"\n" +
	"\x11calculation.proto\x12\tcalculate\"a\n" +
	"\vTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04arg1\x18\x02 \x01(\tR\x04arg1\x12\x12\n" +
	"\x04arg2\x18\x03 \x01(\tR\x04arg2\x12\x1a\n" +
	"\boperator\x18\x04 \x01(\tR\boperator\"M\n" +
	"\rAgentResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x16\n" +
	"\x06result\x18\x02 \x01(\x02R\x06result\x12\x14\n" +
	"\x05error\x18\x03 \x01(\tR\x05error2Q\n" +
	"\fOrchestrator\x12A\n" +
	"\tCalculate\x12\x18.calculate.AgentResponse\x1a\x16.calculate.TaskRequest(\x010\x01B#Z!github.com/vedsatt/calc_prl/protob\x06proto3"

var (
	file_calculation_proto_rawDescOnce sync.Once
	file_calculation_proto_rawDescData []byte
)

func file_calculation_proto_rawDescGZIP() []byte {
	file_calculation_proto_rawDescOnce.Do(func() {
		file_calculation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculation_proto_rawDesc), len(file_calculation_proto_rawDesc)))
	})
	return file_calculation_proto_rawDescData
}

var file_calculation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)

var file_calculation_proto_goTypes = []any{
	(*TaskRequest)(nil),
	(*AgentResponse)(nil),
}

var file_calculation_proto_depIdxs = []int32{
	1,
	0,
	1,
	0,
	0,
	0,
	0,
}

func init() {
	setupCalculationProto()
}

func setupCalculationProto() {
	if File_calculation_proto != nil {
		return
	}
	type dummy struct{}
	builder := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(dummy{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculation_proto_rawDesc), len(file_calculation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculation_proto_goTypes,
		DependencyIndexes: file_calculation_proto_depIdxs,
		MessageInfos:      file_calculation_proto_msgTypes,
	}.Build()
	File_calculation_proto = builder.File
	file_calculation_proto_goTypes = nil
	file_calculation_proto_depIdxs = nil
}
