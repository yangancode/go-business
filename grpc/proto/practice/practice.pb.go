// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: proto/practice/practice.proto

package practice

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

type MsgNoticeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	LatestTime string `protobuf:"bytes,2,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"` // 客户端最新时间
}

func (x *MsgNoticeRequest) Reset() {
	*x = MsgNoticeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_practice_practice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgNoticeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgNoticeRequest) ProtoMessage() {}

func (x *MsgNoticeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_practice_practice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgNoticeRequest.ProtoReflect.Descriptor instead.
func (*MsgNoticeRequest) Descriptor() ([]byte, []int) {
	return file_proto_practice_practice_proto_rawDescGZIP(), []int{0}
}

func (x *MsgNoticeRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *MsgNoticeRequest) GetLatestTime() string {
	if x != nil {
		return x.LatestTime
	}
	return ""
}

type MsgNoticeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc  string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *MsgNoticeResponse) Reset() {
	*x = MsgNoticeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_practice_practice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgNoticeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgNoticeResponse) ProtoMessage() {}

func (x *MsgNoticeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_practice_practice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgNoticeResponse.ProtoReflect.Descriptor instead.
func (*MsgNoticeResponse) Descriptor() ([]byte, []int) {
	return file_proto_practice_practice_proto_rawDescGZIP(), []int{1}
}

func (x *MsgNoticeResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MsgNoticeResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   // 姓名
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"` // 手机
	Work  string `protobuf:"bytes,3,opt,name=work,proto3" json:"work,omitempty"`   // 职位
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_practice_practice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_practice_practice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_practice_practice_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadFileRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UploadFileRequest) GetWork() string {
	if x != nil {
		return x.Work
	}
	return ""
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName    string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize    int32  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	ElapsedTime int32  `protobuf:"varint,3,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_practice_practice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_practice_practice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_practice_practice_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileResponse) GetFileSize() int32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *UploadFileResponse) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

type ChatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid uint32 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ChatMsg) Reset() {
	*x = ChatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_practice_practice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMsg) ProtoMessage() {}

func (x *ChatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_practice_practice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMsg.ProtoReflect.Descriptor instead.
func (*ChatMsg) Descriptor() ([]byte, []int) {
	return file_proto_practice_practice_proto_rawDescGZIP(), []int{4}
}

func (x *ChatMsg) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatMsg) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ChatMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_practice_practice_proto protoreflect.FileDescriptor

var file_proto_practice_practice_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x72, 0x76, 0x22, 0x45, 0x0a,
	0x10, 0x4d, 0x73, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x22, 0x51, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x71, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x07, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xf3, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x4d, 0x73, 0x67, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x4d, 0x73, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x4d, 0x73, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x40, 0x0a, 0x0a,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73,
	0x67, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x10,
	0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_practice_practice_proto_rawDescOnce sync.Once
	file_proto_practice_practice_proto_rawDescData = file_proto_practice_practice_proto_rawDesc
)

func file_proto_practice_practice_proto_rawDescGZIP() []byte {
	file_proto_practice_practice_proto_rawDescOnce.Do(func() {
		file_proto_practice_practice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_practice_practice_proto_rawDescData)
	})
	return file_proto_practice_practice_proto_rawDescData
}

var file_proto_practice_practice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_practice_practice_proto_goTypes = []interface{}{
	(*MsgNoticeRequest)(nil),   // 0: practice.srv.MsgNoticeRequest
	(*MsgNoticeResponse)(nil),  // 1: practice.srv.MsgNoticeResponse
	(*UploadFileRequest)(nil),  // 2: practice.srv.UploadFileRequest
	(*UploadFileResponse)(nil), // 3: practice.srv.UploadFileResponse
	(*ChatMsg)(nil),            // 4: practice.srv.ChatMsg
}
var file_proto_practice_practice_proto_depIdxs = []int32{
	0, // 0: practice.srv.Practice.MsgNotice:input_type -> practice.srv.MsgNoticeRequest
	2, // 1: practice.srv.Practice.UploadFile:input_type -> practice.srv.UploadFileRequest
	4, // 2: practice.srv.Practice.SimpleChat:input_type -> practice.srv.ChatMsg
	1, // 3: practice.srv.Practice.MsgNotice:output_type -> practice.srv.MsgNoticeResponse
	3, // 4: practice.srv.Practice.UploadFile:output_type -> practice.srv.UploadFileResponse
	4, // 5: practice.srv.Practice.SimpleChat:output_type -> practice.srv.ChatMsg
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_practice_practice_proto_init() }
func file_proto_practice_practice_proto_init() {
	if File_proto_practice_practice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_practice_practice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgNoticeRequest); i {
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
		file_proto_practice_practice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgNoticeResponse); i {
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
		file_proto_practice_practice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_proto_practice_practice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
		file_proto_practice_practice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMsg); i {
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
			RawDescriptor: file_proto_practice_practice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_practice_practice_proto_goTypes,
		DependencyIndexes: file_proto_practice_practice_proto_depIdxs,
		MessageInfos:      file_proto_practice_practice_proto_msgTypes,
	}.Build()
	File_proto_practice_practice_proto = out.File
	file_proto_practice_practice_proto_rawDesc = nil
	file_proto_practice_practice_proto_goTypes = nil
	file_proto_practice_practice_proto_depIdxs = nil
}
