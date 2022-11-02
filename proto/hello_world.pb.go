// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: hello_world.proto

package grpc

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReqHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqHead) Reset() {
	*x = ReqHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqHead) ProtoMessage() {}

func (x *ReqHead) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqHead.ProtoReflect.Descriptor instead.
func (*ReqHead) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{2}
}

func (x *ReqHead) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RspHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RspHead) Reset() {
	*x = RspHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspHead) ProtoMessage() {}

func (x *RspHead) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspHead.ProtoReflect.Descriptor instead.
func (*RspHead) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{3}
}

func (x *RspHead) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RspHead) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SaveNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head  *ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Notes string   `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *SaveNotesRequest) Reset() {
	*x = SaveNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNotesRequest) ProtoMessage() {}

func (x *SaveNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNotesRequest.ProtoReflect.Descriptor instead.
func (*SaveNotesRequest) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{4}
}

func (x *SaveNotesRequest) GetHead() *ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *SaveNotesRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type SaveNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *RspHead `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *SaveNotesResponse) Reset() {
	*x = SaveNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNotesResponse) ProtoMessage() {}

func (x *SaveNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNotesResponse.ProtoReflect.Descriptor instead.
func (*SaveNotesResponse) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{5}
}

func (x *SaveNotesResponse) GetHead() *RspHead {
	if x != nil {
		return x.Head
	}
	return nil
}

type GetNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head *ReqHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *GetNotesRequest) Reset() {
	*x = GetNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesRequest) ProtoMessage() {}

func (x *GetNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesRequest.ProtoReflect.Descriptor instead.
func (*GetNotesRequest) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotesRequest) GetHead() *ReqHead {
	if x != nil {
		return x.Head
	}
	return nil
}

type GetNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head  *RspHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	Notes string   `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *GetNotesResponse) Reset() {
	*x = GetNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_world_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotesResponse) ProtoMessage() {}

func (x *GetNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_world_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotesResponse.ProtoReflect.Descriptor instead.
func (*GetNotesResponse) Descriptor() ([]byte, []int) {
	return file_hello_world_proto_rawDescGZIP(), []int{7}
}

func (x *GetNotesResponse) GetHead() *RspHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *GetNotesResponse) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

var File_hello_world_proto protoreflect.FileDescriptor

var file_hello_world_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x48, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x07, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x50, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x68, 0x65,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65,
	0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x52,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x22, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64,
	0x22, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x48, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x32, 0xdd, 0x01, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x17, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1b,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x46, 0x43, 0x48, 0x4b, 0x4b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_world_proto_rawDescOnce sync.Once
	file_hello_world_proto_rawDescData = file_hello_world_proto_rawDesc
)

func file_hello_world_proto_rawDescGZIP() []byte {
	file_hello_world_proto_rawDescOnce.Do(func() {
		file_hello_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_world_proto_rawDescData)
	})
	return file_hello_world_proto_rawDescData
}

var file_hello_world_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_hello_world_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),      // 0: http_grpc.HelloRequest
	(*HelloReply)(nil),        // 1: http_grpc.HelloReply
	(*ReqHead)(nil),           // 2: http_grpc.ReqHead
	(*RspHead)(nil),           // 3: http_grpc.RspHead
	(*SaveNotesRequest)(nil),  // 4: http_grpc.SaveNotesRequest
	(*SaveNotesResponse)(nil), // 5: http_grpc.SaveNotesResponse
	(*GetNotesRequest)(nil),   // 6: http_grpc.GetNotesRequest
	(*GetNotesResponse)(nil),  // 7: http_grpc.GetNotesResponse
}
var file_hello_world_proto_depIdxs = []int32{
	2, // 0: http_grpc.SaveNotesRequest.head:type_name -> http_grpc.ReqHead
	3, // 1: http_grpc.SaveNotesResponse.head:type_name -> http_grpc.RspHead
	2, // 2: http_grpc.GetNotesRequest.head:type_name -> http_grpc.ReqHead
	3, // 3: http_grpc.GetNotesResponse.Head:type_name -> http_grpc.RspHead
	0, // 4: http_grpc.HelloService.SayHello:input_type -> http_grpc.HelloRequest
	4, // 5: http_grpc.HelloService.SaveNotes:input_type -> http_grpc.SaveNotesRequest
	6, // 6: http_grpc.HelloService.GetNotes:input_type -> http_grpc.GetNotesRequest
	1, // 7: http_grpc.HelloService.SayHello:output_type -> http_grpc.HelloReply
	5, // 8: http_grpc.HelloService.SaveNotes:output_type -> http_grpc.SaveNotesResponse
	7, // 9: http_grpc.HelloService.GetNotes:output_type -> http_grpc.GetNotesResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_hello_world_proto_init() }
func file_hello_world_proto_init() {
	if File_hello_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_hello_world_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqHead); i {
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
		file_hello_world_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspHead); i {
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
		file_hello_world_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveNotesRequest); i {
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
		file_hello_world_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveNotesResponse); i {
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
		file_hello_world_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesRequest); i {
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
		file_hello_world_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotesResponse); i {
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
			RawDescriptor: file_hello_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_world_proto_goTypes,
		DependencyIndexes: file_hello_world_proto_depIdxs,
		MessageInfos:      file_hello_world_proto_msgTypes,
	}.Build()
	File_hello_world_proto = out.File
	file_hello_world_proto_rawDesc = nil
	file_hello_world_proto_goTypes = nil
	file_hello_world_proto_depIdxs = nil
}
