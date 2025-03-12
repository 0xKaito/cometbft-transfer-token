// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: proto/app.proto

package proto

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

// TransferTransaction defines a token transfer request.
type TransferTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`     // Sender address
	Receiver  string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"` // Receiver address
	Amount    uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`    // Amount to transfer
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransferTransaction) Reset() {
	*x = TransferTransaction{}
	mi := &file_proto_app_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferTransaction) ProtoMessage() {}

func (x *TransferTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferTransaction.ProtoReflect.Descriptor instead.
func (*TransferTransaction) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{0}
}

func (x *TransferTransaction) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *TransferTransaction) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *TransferTransaction) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferTransaction) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Transaction wraps different types of transactions.
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TxType:
	//
	//	*Transaction_Transfer
	TxType isTransaction_TxType `protobuf_oneof:"tx_type"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_proto_app_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{1}
}

func (m *Transaction) GetTxType() isTransaction_TxType {
	if m != nil {
		return m.TxType
	}
	return nil
}

func (x *Transaction) GetTransfer() *TransferTransaction {
	if x, ok := x.GetTxType().(*Transaction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

type isTransaction_TxType interface {
	isTransaction_TxType()
}

type Transaction_Transfer struct {
	Transfer *TransferTransaction `protobuf:"bytes,1,opt,name=transfer,proto3,oneof"`
}

func (*Transaction_Transfer) isTransaction_TxType() {}

var File_proto_app_proto protoreflect.FileDescriptor

var file_proto_app_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x61, 0x70, 0x70, 0x22, 0x7f, 0x0a, 0x13, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x55, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x61, 0x70, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_app_proto_rawDescOnce sync.Once
	file_proto_app_proto_rawDescData = file_proto_app_proto_rawDesc
)

func file_proto_app_proto_rawDescGZIP() []byte {
	file_proto_app_proto_rawDescOnce.Do(func() {
		file_proto_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_app_proto_rawDescData)
	})
	return file_proto_app_proto_rawDescData
}

var file_proto_app_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_app_proto_goTypes = []any{
	(*TransferTransaction)(nil), // 0: tokenapp.TransferTransaction
	(*Transaction)(nil),         // 1: tokenapp.Transaction
}
var file_proto_app_proto_depIdxs = []int32{
	0, // 0: tokenapp.Transaction.transfer:type_name -> tokenapp.TransferTransaction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_app_proto_init() }
func file_proto_app_proto_init() {
	if File_proto_app_proto != nil {
		return
	}
	file_proto_app_proto_msgTypes[1].OneofWrappers = []any{
		(*Transaction_Transfer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_app_proto_goTypes,
		DependencyIndexes: file_proto_app_proto_depIdxs,
		MessageInfos:      file_proto_app_proto_msgTypes,
	}.Build()
	File_proto_app_proto = out.File
	file_proto_app_proto_rawDesc = nil
	file_proto_app_proto_goTypes = nil
	file_proto_app_proto_depIdxs = nil
}
