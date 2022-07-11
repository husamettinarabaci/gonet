// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: status.proto

package __

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

type PbStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statu         string `protobuf:"bytes,1,opt,name=statu,proto3" json:"statu,omitempty"`
	LocalAddress  string `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	LocalPort     string `protobuf:"bytes,3,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	RemoteAddress string `protobuf:"bytes,4,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	RemotePort    string `protobuf:"bytes,5,opt,name=remote_port,json=remotePort,proto3" json:"remote_port,omitempty"`
}

func (x *PbStatus) Reset() {
	*x = PbStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbStatus) ProtoMessage() {}

func (x *PbStatus) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbStatus.ProtoReflect.Descriptor instead.
func (*PbStatus) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{0}
}

func (x *PbStatus) GetStatu() string {
	if x != nil {
		return x.Statu
	}
	return ""
}

func (x *PbStatus) GetLocalAddress() string {
	if x != nil {
		return x.LocalAddress
	}
	return ""
}

func (x *PbStatus) GetLocalPort() string {
	if x != nil {
		return x.LocalPort
	}
	return ""
}

func (x *PbStatus) GetRemoteAddress() string {
	if x != nil {
		return x.RemoteAddress
	}
	return ""
}

func (x *PbStatus) GetRemotePort() string {
	if x != nil {
		return x.RemotePort
	}
	return ""
}

type PbStatuses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses []*PbStatus `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty"`
}

func (x *PbStatuses) Reset() {
	*x = PbStatuses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbStatuses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbStatuses) ProtoMessage() {}

func (x *PbStatuses) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbStatuses.ProtoReflect.Descriptor instead.
func (*PbStatuses) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{1}
}

func (x *PbStatuses) GetStatuses() []*PbStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type PbResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PbResp) Reset() {
	*x = PbResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbResp) ProtoMessage() {}

func (x *PbResp) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbResp.ProtoReflect.Descriptor instead.
func (*PbResp) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{2}
}

var File_status_proto protoreflect.FileDescriptor

var file_status_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x50, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x50, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x75, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x75, 0x12, 0x23, 0x0a, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x3c, 0x0a, 0x0a, 0x50, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x50, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x08, 0x0a, 0x06, 0x50, 0x62, 0x52, 0x65, 0x73, 0x70, 0x32,
	0x7b, 0x0a, 0x09, 0x53, 0x76, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x0a,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x2e, 0x50, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x10,
	0x2e, 0x50, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x12, 0x14, 0x2e, 0x50, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x50, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x50, 0x62, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01,
	0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_status_proto_rawDescOnce sync.Once
	file_status_proto_rawDescData = file_status_proto_rawDesc
)

func file_status_proto_rawDescGZIP() []byte {
	file_status_proto_rawDescOnce.Do(func() {
		file_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_proto_rawDescData)
	})
	return file_status_proto_rawDescData
}

var file_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_status_proto_goTypes = []interface{}{
	(*PbStatus)(nil),   // 0: PbStatus.PbStatus
	(*PbStatuses)(nil), // 1: PbStatus.PbStatuses
	(*PbResp)(nil),     // 2: PbStatus.PbResp
}
var file_status_proto_depIdxs = []int32{
	0, // 0: PbStatus.PbStatuses.statuses:type_name -> PbStatus.PbStatus
	0, // 1: PbStatus.SvcStatus.SendStatus:input_type -> PbStatus.PbStatus
	1, // 2: PbStatus.SvcStatus.SendStatuses:input_type -> PbStatus.PbStatuses
	2, // 3: PbStatus.SvcStatus.SendStatus:output_type -> PbStatus.PbResp
	2, // 4: PbStatus.SvcStatus.SendStatuses:output_type -> PbStatus.PbResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_status_proto_init() }
func file_status_proto_init() {
	if File_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbStatus); i {
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
		file_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbStatuses); i {
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
		file_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbResp); i {
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
			RawDescriptor: file_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_status_proto_goTypes,
		DependencyIndexes: file_status_proto_depIdxs,
		MessageInfos:      file_status_proto_msgTypes,
	}.Build()
	File_status_proto = out.File
	file_status_proto_rawDesc = nil
	file_status_proto_goTypes = nil
	file_status_proto_depIdxs = nil
}
