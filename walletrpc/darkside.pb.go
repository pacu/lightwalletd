// Copyright (c) 2019-2020 The Zcash developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or https://www.opensource.org/licenses/mit-license.php .

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: darkside.proto

package walletrpc

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

type DarksideMetaState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaplingActivation              int32  `protobuf:"varint,1,opt,name=saplingActivation,proto3" json:"saplingActivation,omitempty"`
	BranchID                       string `protobuf:"bytes,2,opt,name=branchID,proto3" json:"branchID,omitempty"`
	ChainName                      string `protobuf:"bytes,3,opt,name=chainName,proto3" json:"chainName,omitempty"`
	StartSaplingCommitmentTreeSize uint32 `protobuf:"varint,4,opt,name=startSaplingCommitmentTreeSize,proto3" json:"startSaplingCommitmentTreeSize,omitempty"`
	StartOrchardCommitmentTreeSize uint32 `protobuf:"varint,5,opt,name=startOrchardCommitmentTreeSize,proto3" json:"startOrchardCommitmentTreeSize,omitempty"`
}

func (x *DarksideMetaState) Reset() {
	*x = DarksideMetaState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideMetaState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideMetaState) ProtoMessage() {}

func (x *DarksideMetaState) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideMetaState.ProtoReflect.Descriptor instead.
func (*DarksideMetaState) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{0}
}

func (x *DarksideMetaState) GetSaplingActivation() int32 {
	if x != nil {
		return x.SaplingActivation
	}
	return 0
}

func (x *DarksideMetaState) GetBranchID() string {
	if x != nil {
		return x.BranchID
	}
	return ""
}

func (x *DarksideMetaState) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *DarksideMetaState) GetStartSaplingCommitmentTreeSize() uint32 {
	if x != nil {
		return x.StartSaplingCommitmentTreeSize
	}
	return 0
}

func (x *DarksideMetaState) GetStartOrchardCommitmentTreeSize() uint32 {
	if x != nil {
		return x.StartOrchardCommitmentTreeSize
	}
	return 0
}

// A block is a hex-encoded string.
type DarksideBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block string `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *DarksideBlock) Reset() {
	*x = DarksideBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideBlock) ProtoMessage() {}

func (x *DarksideBlock) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideBlock.ProtoReflect.Descriptor instead.
func (*DarksideBlock) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{1}
}

func (x *DarksideBlock) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

// DarksideBlocksURL is typically something like:
// https://raw.githubusercontent.com/zcash-hackworks/darksidewalletd-test-data/master/basic-reorg/before-reorg.txt
type DarksideBlocksURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DarksideBlocksURL) Reset() {
	*x = DarksideBlocksURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideBlocksURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideBlocksURL) ProtoMessage() {}

func (x *DarksideBlocksURL) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideBlocksURL.ProtoReflect.Descriptor instead.
func (*DarksideBlocksURL) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{2}
}

func (x *DarksideBlocksURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// DarksideTransactionsURL refers to an HTTP source that contains a list
// of hex-encoded transactions, one per line, that are to be associated
// with the given height (fake-mined into the block at that height)
type DarksideTransactionsURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int32  `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *DarksideTransactionsURL) Reset() {
	*x = DarksideTransactionsURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideTransactionsURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideTransactionsURL) ProtoMessage() {}

func (x *DarksideTransactionsURL) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideTransactionsURL.ProtoReflect.Descriptor instead.
func (*DarksideTransactionsURL) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{3}
}

func (x *DarksideTransactionsURL) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *DarksideTransactionsURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DarksideHeight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *DarksideHeight) Reset() {
	*x = DarksideHeight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideHeight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideHeight) ProtoMessage() {}

func (x *DarksideHeight) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideHeight.ProtoReflect.Descriptor instead.
func (*DarksideHeight) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{4}
}

func (x *DarksideHeight) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type DarksideEmptyBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height int32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Nonce  int32 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Count  int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *DarksideEmptyBlocks) Reset() {
	*x = DarksideEmptyBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideEmptyBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideEmptyBlocks) ProtoMessage() {}

func (x *DarksideEmptyBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideEmptyBlocks.ProtoReflect.Descriptor instead.
func (*DarksideEmptyBlocks) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{5}
}

func (x *DarksideEmptyBlocks) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *DarksideEmptyBlocks) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *DarksideEmptyBlocks) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DarksideSubTreeIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShieldedProtocol ShieldedProtocol `protobuf:"varint,1,opt,name=shieldedProtocol,proto3,enum=cash.z.wallet.sdk.rpc.ShieldedProtocol" json:"shieldedProtocol,omitempty"`
	Index            int32            `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *DarksideSubTreeIndex) Reset() {
	*x = DarksideSubTreeIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideSubTreeIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideSubTreeIndex) ProtoMessage() {}

func (x *DarksideSubTreeIndex) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideSubTreeIndex.ProtoReflect.Descriptor instead.
func (*DarksideSubTreeIndex) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{6}
}

func (x *DarksideSubTreeIndex) GetShieldedProtocol() ShieldedProtocol {
	if x != nil {
		return x.ShieldedProtocol
	}
	return ShieldedProtocol_sapling
}

func (x *DarksideSubTreeIndex) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type DarksideSubTreeRoots struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShieldedProtocol ShieldedProtocol `protobuf:"varint,1,opt,name=shieldedProtocol,proto3,enum=cash.z.wallet.sdk.rpc.ShieldedProtocol" json:"shieldedProtocol,omitempty"`
	Index            int32            `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Subtrees         []*SubtreeRoot   `protobuf:"bytes,3,rep,name=subtrees,proto3" json:"subtrees,omitempty"`
}

func (x *DarksideSubTreeRoots) Reset() {
	*x = DarksideSubTreeRoots{}
	if protoimpl.UnsafeEnabled {
		mi := &file_darkside_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DarksideSubTreeRoots) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DarksideSubTreeRoots) ProtoMessage() {}

func (x *DarksideSubTreeRoots) ProtoReflect() protoreflect.Message {
	mi := &file_darkside_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DarksideSubTreeRoots.ProtoReflect.Descriptor instead.
func (*DarksideSubTreeRoots) Descriptor() ([]byte, []int) {
	return file_darkside_proto_rawDescGZIP(), []int{7}
}

func (x *DarksideSubTreeRoots) GetShieldedProtocol() ShieldedProtocol {
	if x != nil {
		return x.ShieldedProtocol
	}
	return ShieldedProtocol_sapling
}

func (x *DarksideSubTreeRoots) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DarksideSubTreeRoots) GetSubtrees() []*SubtreeRoot {
	if x != nil {
		return x.Subtrees
	}
	return nil
}

var File_darkside_proto protoreflect.FileDescriptor

var file_darkside_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x11, 0x44, 0x61, 0x72, 0x6b, 0x73,
	0x69, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x73, 0x61, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x73, 0x61, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x1e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x61, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72,
	0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1e, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x61, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x46, 0x0a, 0x1e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x1e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x63, 0x68, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x25, 0x0a, 0x11, 0x44,
	0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x55, 0x52, 0x4c,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x43, 0x0a, 0x17, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x61, 0x72, 0x6b, 0x73,
	0x69, 0x64, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x22, 0x59, 0x0a, 0x13, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a,
	0x14, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x53, 0x75, 0x62, 0x54, 0x72, 0x65, 0x65,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x53, 0x0a, 0x10, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x10, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0xc1, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x53, 0x75, 0x62,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x10, 0x73, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68, 0x69, 0x65,
	0x6c, 0x64, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x10, 0x73, 0x68,
	0x69, 0x65, 0x6c, 0x64, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x72, 0x65, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x75, 0x62, 0x74, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74,
	0x72, 0x65, 0x65, 0x73, 0x32, 0x8b, 0x0d, 0x0a, 0x10, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x05, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73,
	0x69, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x11,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69,
	0x64, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x57, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e,
	0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x55,
	0x52, 0x4c, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x5f, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25,
	0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x63, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x55, 0x52, 0x4c, 0x1a, 0x1c, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0b,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x62, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e,
	0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x19, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55,
	0x74, 0x78, 0x6f, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x74, 0x78, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x10, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x55, 0x74, 0x78, 0x6f, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x54, 0x72, 0x65, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x72,
	0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x41, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e,
	0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x53, 0x75, 0x62, 0x54, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x2b, 0x2e,
	0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65, 0x53, 0x75,
	0x62, 0x54, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73,
	0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x14, 0x41, 0x64,
	0x64, 0x53, 0x75, 0x62, 0x54, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73,
	0x69, 0x64, 0x65, 0x53, 0x75, 0x62, 0x54, 0x72, 0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x1a,
	0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x69, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x75, 0x62, 0x54, 0x72,
	0x65, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x46, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x2b, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x72, 0x6b, 0x73, 0x69, 0x64, 0x65,
	0x53, 0x75, 0x62, 0x54, 0x72, 0x65, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x1c, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x14,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x54, 0x72, 0x65, 0x65, 0x52,
	0x6f, 0x6f, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x7a, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x1b, 0x5a, 0x16, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x64, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x72, 0x70, 0x63, 0xba, 0x02, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_darkside_proto_rawDescOnce sync.Once
	file_darkside_proto_rawDescData = file_darkside_proto_rawDesc
)

func file_darkside_proto_rawDescGZIP() []byte {
	file_darkside_proto_rawDescOnce.Do(func() {
		file_darkside_proto_rawDescData = protoimpl.X.CompressGZIP(file_darkside_proto_rawDescData)
	})
	return file_darkside_proto_rawDescData
}

var file_darkside_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_darkside_proto_goTypes = []interface{}{
	(*DarksideMetaState)(nil),       // 0: cash.z.wallet.sdk.rpc.DarksideMetaState
	(*DarksideBlock)(nil),           // 1: cash.z.wallet.sdk.rpc.DarksideBlock
	(*DarksideBlocksURL)(nil),       // 2: cash.z.wallet.sdk.rpc.DarksideBlocksURL
	(*DarksideTransactionsURL)(nil), // 3: cash.z.wallet.sdk.rpc.DarksideTransactionsURL
	(*DarksideHeight)(nil),          // 4: cash.z.wallet.sdk.rpc.DarksideHeight
	(*DarksideEmptyBlocks)(nil),     // 5: cash.z.wallet.sdk.rpc.DarksideEmptyBlocks
	(*DarksideSubTreeIndex)(nil),    // 6: cash.z.wallet.sdk.rpc.DarksideSubTreeIndex
	(*DarksideSubTreeRoots)(nil),    // 7: cash.z.wallet.sdk.rpc.DarksideSubTreeRoots
	(ShieldedProtocol)(0),           // 8: cash.z.wallet.sdk.rpc.ShieldedProtocol
	(*SubtreeRoot)(nil),             // 9: cash.z.wallet.sdk.rpc.SubtreeRoot
	(*RawTransaction)(nil),          // 10: cash.z.wallet.sdk.rpc.RawTransaction
	(*Empty)(nil),                   // 11: cash.z.wallet.sdk.rpc.Empty
	(*GetAddressUtxosReply)(nil),    // 12: cash.z.wallet.sdk.rpc.GetAddressUtxosReply
	(*TreeState)(nil),               // 13: cash.z.wallet.sdk.rpc.TreeState
	(*BlockID)(nil),                 // 14: cash.z.wallet.sdk.rpc.BlockID
}
var file_darkside_proto_depIdxs = []int32{
	8,  // 0: cash.z.wallet.sdk.rpc.DarksideSubTreeIndex.shieldedProtocol:type_name -> cash.z.wallet.sdk.rpc.ShieldedProtocol
	8,  // 1: cash.z.wallet.sdk.rpc.DarksideSubTreeRoots.shieldedProtocol:type_name -> cash.z.wallet.sdk.rpc.ShieldedProtocol
	9,  // 2: cash.z.wallet.sdk.rpc.DarksideSubTreeRoots.subtrees:type_name -> cash.z.wallet.sdk.rpc.SubtreeRoot
	0,  // 3: cash.z.wallet.sdk.rpc.DarksideStreamer.Reset:input_type -> cash.z.wallet.sdk.rpc.DarksideMetaState
	1,  // 4: cash.z.wallet.sdk.rpc.DarksideStreamer.StageBlocksStream:input_type -> cash.z.wallet.sdk.rpc.DarksideBlock
	2,  // 5: cash.z.wallet.sdk.rpc.DarksideStreamer.StageBlocks:input_type -> cash.z.wallet.sdk.rpc.DarksideBlocksURL
	5,  // 6: cash.z.wallet.sdk.rpc.DarksideStreamer.StageBlocksCreate:input_type -> cash.z.wallet.sdk.rpc.DarksideEmptyBlocks
	10, // 7: cash.z.wallet.sdk.rpc.DarksideStreamer.StageTransactionsStream:input_type -> cash.z.wallet.sdk.rpc.RawTransaction
	3,  // 8: cash.z.wallet.sdk.rpc.DarksideStreamer.StageTransactions:input_type -> cash.z.wallet.sdk.rpc.DarksideTransactionsURL
	4,  // 9: cash.z.wallet.sdk.rpc.DarksideStreamer.ApplyStaged:input_type -> cash.z.wallet.sdk.rpc.DarksideHeight
	11, // 10: cash.z.wallet.sdk.rpc.DarksideStreamer.GetIncomingTransactions:input_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 11: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearIncomingTransactions:input_type -> cash.z.wallet.sdk.rpc.Empty
	12, // 12: cash.z.wallet.sdk.rpc.DarksideStreamer.AddAddressUtxo:input_type -> cash.z.wallet.sdk.rpc.GetAddressUtxosReply
	11, // 13: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearAddressUtxo:input_type -> cash.z.wallet.sdk.rpc.Empty
	13, // 14: cash.z.wallet.sdk.rpc.DarksideStreamer.AddTreeState:input_type -> cash.z.wallet.sdk.rpc.TreeState
	14, // 15: cash.z.wallet.sdk.rpc.DarksideStreamer.RemoveTreeState:input_type -> cash.z.wallet.sdk.rpc.BlockID
	11, // 16: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearAllTreeStates:input_type -> cash.z.wallet.sdk.rpc.Empty
	7,  // 17: cash.z.wallet.sdk.rpc.DarksideStreamer.AddSubTreeRoot:input_type -> cash.z.wallet.sdk.rpc.DarksideSubTreeRoots
	7,  // 18: cash.z.wallet.sdk.rpc.DarksideStreamer.AddSubTreeRootStream:input_type -> cash.z.wallet.sdk.rpc.DarksideSubTreeRoots
	6,  // 19: cash.z.wallet.sdk.rpc.DarksideStreamer.RemoveSubTreeRootsForIndex:input_type -> cash.z.wallet.sdk.rpc.DarksideSubTreeIndex
	11, // 20: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearAllSubTreeRoots:input_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 21: cash.z.wallet.sdk.rpc.DarksideStreamer.Reset:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 22: cash.z.wallet.sdk.rpc.DarksideStreamer.StageBlocksStream:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 23: cash.z.wallet.sdk.rpc.DarksideStreamer.StageBlocks:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 24: cash.z.wallet.sdk.rpc.DarksideStreamer.StageBlocksCreate:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 25: cash.z.wallet.sdk.rpc.DarksideStreamer.StageTransactionsStream:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 26: cash.z.wallet.sdk.rpc.DarksideStreamer.StageTransactions:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 27: cash.z.wallet.sdk.rpc.DarksideStreamer.ApplyStaged:output_type -> cash.z.wallet.sdk.rpc.Empty
	10, // 28: cash.z.wallet.sdk.rpc.DarksideStreamer.GetIncomingTransactions:output_type -> cash.z.wallet.sdk.rpc.RawTransaction
	11, // 29: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearIncomingTransactions:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 30: cash.z.wallet.sdk.rpc.DarksideStreamer.AddAddressUtxo:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 31: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearAddressUtxo:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 32: cash.z.wallet.sdk.rpc.DarksideStreamer.AddTreeState:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 33: cash.z.wallet.sdk.rpc.DarksideStreamer.RemoveTreeState:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 34: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearAllTreeStates:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 35: cash.z.wallet.sdk.rpc.DarksideStreamer.AddSubTreeRoot:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 36: cash.z.wallet.sdk.rpc.DarksideStreamer.AddSubTreeRootStream:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 37: cash.z.wallet.sdk.rpc.DarksideStreamer.RemoveSubTreeRootsForIndex:output_type -> cash.z.wallet.sdk.rpc.Empty
	11, // 38: cash.z.wallet.sdk.rpc.DarksideStreamer.ClearAllSubTreeRoots:output_type -> cash.z.wallet.sdk.rpc.Empty
	21, // [21:39] is the sub-list for method output_type
	3,  // [3:21] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_darkside_proto_init() }
func file_darkside_proto_init() {
	if File_darkside_proto != nil {
		return
	}
	file_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_darkside_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideMetaState); i {
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
		file_darkside_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideBlock); i {
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
		file_darkside_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideBlocksURL); i {
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
		file_darkside_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideTransactionsURL); i {
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
		file_darkside_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideHeight); i {
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
		file_darkside_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideEmptyBlocks); i {
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
		file_darkside_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideSubTreeIndex); i {
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
		file_darkside_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DarksideSubTreeRoots); i {
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
			RawDescriptor: file_darkside_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_darkside_proto_goTypes,
		DependencyIndexes: file_darkside_proto_depIdxs,
		MessageInfos:      file_darkside_proto_msgTypes,
	}.Build()
	File_darkside_proto = out.File
	file_darkside_proto_rawDesc = nil
	file_darkside_proto_goTypes = nil
	file_darkside_proto_depIdxs = nil
}
