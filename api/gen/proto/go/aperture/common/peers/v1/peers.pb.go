// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aperture/common/peers/v1/peers.proto

package peersv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Peers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerInfos []*PeerInfo `protobuf:"bytes,1,rep,name=peer_infos,json=peerInfos,proto3" json:"peer_infos,omitempty"`
}

func (x *Peers) Reset() {
	*x = Peers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peers) ProtoMessage() {}

func (x *Peers) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peers.ProtoReflect.Descriptor instead.
func (*Peers) Descriptor() ([]byte, []int) {
	return file_aperture_common_peers_v1_peers_proto_rawDescGZIP(), []int{0}
}

func (x *Peers) GetPeerInfos() []*PeerInfo {
	if x != nil {
		return x.PeerInfos
	}
	return nil
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Hostname string            `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Services map[string]string `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_aperture_common_peers_v1_peers_proto_rawDescGZIP(), []int{1}
}

func (x *PeerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PeerInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *PeerInfo) GetServices() map[string]string {
	if x != nil {
		return x.Services
	}
	return nil
}

type PeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *PeerRequest) Reset() {
	*x = PeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerRequest) ProtoMessage() {}

func (x *PeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerRequest.ProtoReflect.Descriptor instead.
func (*PeerRequest) Descriptor() ([]byte, []int) {
	return file_aperture_common_peers_v1_peers_proto_rawDescGZIP(), []int{2}
}

func (x *PeerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type PeerKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *PeerKeysResponse) Reset() {
	*x = PeerKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerKeysResponse) ProtoMessage() {}

func (x *PeerKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_common_peers_v1_peers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerKeysResponse.ProtoReflect.Descriptor instead.
func (*PeerKeysResponse) Descriptor() ([]byte, []int) {
	return file_aperture_common_peers_v1_peers_proto_rawDescGZIP(), []int{3}
}

func (x *PeerKeysResponse) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_aperture_common_peers_v1_peers_proto protoreflect.FileDescriptor

var file_aperture_common_peers_v1_peers_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x05, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x26,
	0x0a, 0x10, 0x50, 0x65, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x32, 0xe6, 0x03, 0x0a, 0x14, 0x50, 0x65, 0x65, 0x72, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x82, 0x01,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x7d, 0x12, 0x53, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50, 0x65,
	0x65, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12, 0x25,
	0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0xfe, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x6e,
	0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65, 0x65, 0x72, 0x73, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x43, 0x50, 0xaa, 0x02, 0x18, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x18, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x70,
	0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1b, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x65, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aperture_common_peers_v1_peers_proto_rawDescOnce sync.Once
	file_aperture_common_peers_v1_peers_proto_rawDescData = file_aperture_common_peers_v1_peers_proto_rawDesc
)

func file_aperture_common_peers_v1_peers_proto_rawDescGZIP() []byte {
	file_aperture_common_peers_v1_peers_proto_rawDescOnce.Do(func() {
		file_aperture_common_peers_v1_peers_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_common_peers_v1_peers_proto_rawDescData)
	})
	return file_aperture_common_peers_v1_peers_proto_rawDescData
}

var file_aperture_common_peers_v1_peers_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_aperture_common_peers_v1_peers_proto_goTypes = []interface{}{
	(*Peers)(nil),            // 0: aperture.common.peers.v1.Peers
	(*PeerInfo)(nil),         // 1: aperture.common.peers.v1.PeerInfo
	(*PeerRequest)(nil),      // 2: aperture.common.peers.v1.PeerRequest
	(*PeerKeysResponse)(nil), // 3: aperture.common.peers.v1.PeerKeysResponse
	nil,                      // 4: aperture.common.peers.v1.PeerInfo.ServicesEntry
	(*emptypb.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_aperture_common_peers_v1_peers_proto_depIdxs = []int32{
	1, // 0: aperture.common.peers.v1.Peers.peer_infos:type_name -> aperture.common.peers.v1.PeerInfo
	4, // 1: aperture.common.peers.v1.PeerInfo.services:type_name -> aperture.common.peers.v1.PeerInfo.ServicesEntry
	5, // 2: aperture.common.peers.v1.PeerDiscoveryService.GetPeers:input_type -> google.protobuf.Empty
	2, // 3: aperture.common.peers.v1.PeerDiscoveryService.GetPeer:input_type -> aperture.common.peers.v1.PeerRequest
	5, // 4: aperture.common.peers.v1.PeerDiscoveryService.GetPeerKeys:input_type -> google.protobuf.Empty
	1, // 5: aperture.common.peers.v1.PeerDiscoveryService.AddPeer:input_type -> aperture.common.peers.v1.PeerInfo
	2, // 6: aperture.common.peers.v1.PeerDiscoveryService.RemovePeer:input_type -> aperture.common.peers.v1.PeerRequest
	0, // 7: aperture.common.peers.v1.PeerDiscoveryService.GetPeers:output_type -> aperture.common.peers.v1.Peers
	1, // 8: aperture.common.peers.v1.PeerDiscoveryService.GetPeer:output_type -> aperture.common.peers.v1.PeerInfo
	3, // 9: aperture.common.peers.v1.PeerDiscoveryService.GetPeerKeys:output_type -> aperture.common.peers.v1.PeerKeysResponse
	5, // 10: aperture.common.peers.v1.PeerDiscoveryService.AddPeer:output_type -> google.protobuf.Empty
	5, // 11: aperture.common.peers.v1.PeerDiscoveryService.RemovePeer:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_aperture_common_peers_v1_peers_proto_init() }
func file_aperture_common_peers_v1_peers_proto_init() {
	if File_aperture_common_peers_v1_peers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aperture_common_peers_v1_peers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peers); i {
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
		file_aperture_common_peers_v1_peers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
		file_aperture_common_peers_v1_peers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerRequest); i {
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
		file_aperture_common_peers_v1_peers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerKeysResponse); i {
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
			RawDescriptor: file_aperture_common_peers_v1_peers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aperture_common_peers_v1_peers_proto_goTypes,
		DependencyIndexes: file_aperture_common_peers_v1_peers_proto_depIdxs,
		MessageInfos:      file_aperture_common_peers_v1_peers_proto_msgTypes,
	}.Build()
	File_aperture_common_peers_v1_peers_proto = out.File
	file_aperture_common_peers_v1_peers_proto_rawDesc = nil
	file_aperture_common_peers_v1_peers_proto_goTypes = nil
	file_aperture_common_peers_v1_peers_proto_depIdxs = nil
}
