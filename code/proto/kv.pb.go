// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: code/proto/kv.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryTxnRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxnId         string                 `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryTxnRequest) Reset() {
	*x = QueryTxnRequest{}
	mi := &file_code_proto_kv_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryTxnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTxnRequest) ProtoMessage() {}

func (x *QueryTxnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTxnRequest.ProtoReflect.Descriptor instead.
func (*QueryTxnRequest) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{0}
}

func (x *QueryTxnRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

type QueryTxnReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // "PREPARED", "COMMITTED", or "UNKNOWN"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryTxnReply) Reset() {
	*x = QueryTxnReply{}
	mi := &file_code_proto_kv_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryTxnReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTxnReply) ProtoMessage() {}

func (x *QueryTxnReply) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTxnReply.ProtoReflect.Descriptor instead.
func (*QueryTxnReply) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{1}
}

func (x *QueryTxnReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PrepareRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxnId         string                 `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	Key           string                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Operation     string                 `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrepareRequest) Reset() {
	*x = PrepareRequest{}
	mi := &file_code_proto_kv_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrepareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRequest) ProtoMessage() {}

func (x *PrepareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRequest.ProtoReflect.Descriptor instead.
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{2}
}

func (x *PrepareRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

func (x *PrepareRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PrepareRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PrepareRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

type CommitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxnId         string                 `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommitRequest) Reset() {
	*x = CommitRequest{}
	mi := &file_code_proto_kv_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitRequest) ProtoMessage() {}

func (x *CommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitRequest.ProtoReflect.Descriptor instead.
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{3}
}

func (x *CommitRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

type AbortRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TxnId         string                 `protobuf:"bytes,1,opt,name=txn_id,json=txnId,proto3" json:"txn_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AbortRequest) Reset() {
	*x = AbortRequest{}
	mi := &file_code_proto_kv_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AbortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbortRequest) ProtoMessage() {}

func (x *AbortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbortRequest.ProtoReflect.Descriptor instead.
func (*AbortRequest) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{4}
}

func (x *AbortRequest) GetTxnId() string {
	if x != nil {
		return x.TxnId
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_code_proto_kv_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	mi := &file_code_proto_kv_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{6}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ValueReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValueReply) Reset() {
	*x = ValueReply{}
	mi := &file_code_proto_kv_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueReply) ProtoMessage() {}

func (x *ValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueReply.ProtoReflect.Descriptor instead.
func (*ValueReply) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{7}
}

func (x *ValueReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_code_proto_kv_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_code_proto_kv_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_code_proto_kv_proto_rawDescGZIP(), []int{8}
}

func (x *Ack) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_code_proto_kv_proto protoreflect.FileDescriptor

const file_code_proto_kv_proto_rawDesc = "" +
	"\n" +
	"\x13code/proto/kv.proto\x12\x04code\"(\n" +
	"\x0fQueryTxnRequest\x12\x15\n" +
	"\x06txn_id\x18\x01 \x01(\tR\x05txnId\"'\n" +
	"\rQueryTxnReply\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"m\n" +
	"\x0ePrepareRequest\x12\x15\n" +
	"\x06txn_id\x18\x01 \x01(\tR\x05txnId\x12\x10\n" +
	"\x03key\x18\x02 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x03 \x01(\tR\x05value\x12\x1c\n" +
	"\toperation\x18\x04 \x01(\tR\toperation\"&\n" +
	"\rCommitRequest\x12\x15\n" +
	"\x06txn_id\x18\x01 \x01(\tR\x05txnId\"%\n" +
	"\fAbortRequest\x12\x15\n" +
	"\x06txn_id\x18\x01 \x01(\tR\x05txnId\"\x1e\n" +
	"\n" +
	"GetRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"4\n" +
	"\n" +
	"SetRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\"\"\n" +
	"\n" +
	"ValueReply\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\"\x1f\n" +
	"\x03Ack\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\x96\x02\n" +
	"\tKVService\x12*\n" +
	"\aPrepare\x12\x14.code.PrepareRequest\x1a\t.code.Ack\x12(\n" +
	"\x06Commit\x12\x13.code.CommitRequest\x1a\t.code.Ack\x12&\n" +
	"\x05Abort\x12\x12.code.AbortRequest\x1a\t.code.Ack\x12)\n" +
	"\x03Get\x12\x10.code.GetRequest\x1a\x10.code.ValueReply\x12\"\n" +
	"\x03Set\x12\x10.code.SetRequest\x1a\t.code.Ack\x12<\n" +
	"\x0eQueryTxnStatus\x12\x15.code.QueryTxnRequest\x1a\x13.code.QueryTxnReplyB\x12Z\x10code/proto;protob\x06proto3"

var (
	file_code_proto_kv_proto_rawDescOnce sync.Once
	file_code_proto_kv_proto_rawDescData []byte
)

func file_code_proto_kv_proto_rawDescGZIP() []byte {
	file_code_proto_kv_proto_rawDescOnce.Do(func() {
		file_code_proto_kv_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_code_proto_kv_proto_rawDesc), len(file_code_proto_kv_proto_rawDesc)))
	})
	return file_code_proto_kv_proto_rawDescData
}

var file_code_proto_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_code_proto_kv_proto_goTypes = []any{
	(*QueryTxnRequest)(nil), // 0: code.QueryTxnRequest
	(*QueryTxnReply)(nil),   // 1: code.QueryTxnReply
	(*PrepareRequest)(nil),  // 2: code.PrepareRequest
	(*CommitRequest)(nil),   // 3: code.CommitRequest
	(*AbortRequest)(nil),    // 4: code.AbortRequest
	(*GetRequest)(nil),      // 5: code.GetRequest
	(*SetRequest)(nil),      // 6: code.SetRequest
	(*ValueReply)(nil),      // 7: code.ValueReply
	(*Ack)(nil),             // 8: code.Ack
}
var file_code_proto_kv_proto_depIdxs = []int32{
	2, // 0: code.KVService.Prepare:input_type -> code.PrepareRequest
	3, // 1: code.KVService.Commit:input_type -> code.CommitRequest
	4, // 2: code.KVService.Abort:input_type -> code.AbortRequest
	5, // 3: code.KVService.Get:input_type -> code.GetRequest
	6, // 4: code.KVService.Set:input_type -> code.SetRequest
	0, // 5: code.KVService.QueryTxnStatus:input_type -> code.QueryTxnRequest
	8, // 6: code.KVService.Prepare:output_type -> code.Ack
	8, // 7: code.KVService.Commit:output_type -> code.Ack
	8, // 8: code.KVService.Abort:output_type -> code.Ack
	7, // 9: code.KVService.Get:output_type -> code.ValueReply
	8, // 10: code.KVService.Set:output_type -> code.Ack
	1, // 11: code.KVService.QueryTxnStatus:output_type -> code.QueryTxnReply
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_proto_kv_proto_init() }
func file_code_proto_kv_proto_init() {
	if File_code_proto_kv_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_code_proto_kv_proto_rawDesc), len(file_code_proto_kv_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_code_proto_kv_proto_goTypes,
		DependencyIndexes: file_code_proto_kv_proto_depIdxs,
		MessageInfos:      file_code_proto_kv_proto_msgTypes,
	}.Build()
	File_code_proto_kv_proto = out.File
	file_code_proto_kv_proto_goTypes = nil
	file_code_proto_kv_proto_depIdxs = nil
}
