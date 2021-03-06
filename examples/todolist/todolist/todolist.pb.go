// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: todolist/todolist.proto

package todolist

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

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time         string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Action       string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	DetailAction string `protobuf:"bytes,4,opt,name=detailAction,proto3" json:"detailAction,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Todo) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Todo) GetDetailAction() string {
	if x != nil {
		return x.DetailAction
	}
	return ""
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{1}
}

func (x *InsertRequest) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type InsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"` //statusCode == 1 -> success
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InsertResponse) Reset() {
	*x = InsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResponse) ProtoMessage() {}

func (x *InsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResponse.ProtoReflect.Descriptor instead.
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{2}
}

func (x *InsertResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *InsertResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllTodoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllTodoListRequest) Reset() {
	*x = GetAllTodoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTodoListRequest) ProtoMessage() {}

func (x *GetAllTodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTodoListRequest.ProtoReflect.Descriptor instead.
func (*GetAllTodoListRequest) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{3}
}

type GetAllTodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetAllTodoListResponse) Reset() {
	*x = GetAllTodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTodoListResponse) ProtoMessage() {}

func (x *GetAllTodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTodoListResponse.ProtoReflect.Descriptor instead.
func (*GetAllTodoListResponse) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTodoListResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type GetTodoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTodoListRequest) Reset() {
	*x = GetTodoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoListRequest) ProtoMessage() {}

func (x *GetTodoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoListRequest.ProtoReflect.Descriptor instead.
func (*GetTodoListRequest) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{5}
}

func (x *GetTodoListRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTodoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *GetTodoListResponse) Reset() {
	*x = GetTodoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoListResponse) ProtoMessage() {}

func (x *GetTodoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoListResponse.ProtoReflect.Descriptor instead.
func (*GetTodoListResponse) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{6}
}

func (x *GetTodoListResponse) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewTodo *Todo `protobuf:"bytes,1,opt,name=newTodo,proto3" json:"newTodo,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetNewTodo() *Todo {
	if x != nil {
		return x.NewTodo
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"` //statusCode == 1 -> success
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"` //statusCode == 1 -> success
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todolist_todolist_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todolist_todolist_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_todolist_todolist_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_todolist_todolist_proto protoreflect.FileDescriptor

var file_todolist_todolist_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x0d, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f,
	0x22, 0x4a, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x17, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04,
	0x74, 0x6f, 0x64, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x64, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x64, 0x6f, 0x22,
	0x4a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf3, 0x02, 0x0a, 0x0f, 0x54, 0x6f, 0x64,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x64, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5f,
	0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x0d, 0x54, 0x6f, 0x64,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todolist_todolist_proto_rawDescOnce sync.Once
	file_todolist_todolist_proto_rawDescData = file_todolist_todolist_proto_rawDesc
)

func file_todolist_todolist_proto_rawDescGZIP() []byte {
	file_todolist_todolist_proto_rawDescOnce.Do(func() {
		file_todolist_todolist_proto_rawDescData = protoimpl.X.CompressGZIP(file_todolist_todolist_proto_rawDescData)
	})
	return file_todolist_todolist_proto_rawDescData
}

var file_todolist_todolist_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_todolist_todolist_proto_goTypes = []interface{}{
	(*Todo)(nil),                   // 0: todolist.Todo
	(*InsertRequest)(nil),          // 1: todolist.InsertRequest
	(*InsertResponse)(nil),         // 2: todolist.InsertResponse
	(*GetAllTodoListRequest)(nil),  // 3: todolist.GetAllTodoListRequest
	(*GetAllTodoListResponse)(nil), // 4: todolist.GetAllTodoListResponse
	(*GetTodoListRequest)(nil),     // 5: todolist.GetTodoListRequest
	(*GetTodoListResponse)(nil),    // 6: todolist.GetTodoListResponse
	(*UpdateRequest)(nil),          // 7: todolist.UpdateRequest
	(*UpdateResponse)(nil),         // 8: todolist.UpdateResponse
	(*DeleteRequest)(nil),          // 9: todolist.DeleteRequest
	(*DeleteResponse)(nil),         // 10: todolist.DeleteResponse
}
var file_todolist_todolist_proto_depIdxs = []int32{
	0,  // 0: todolist.InsertRequest.todo:type_name -> todolist.Todo
	0,  // 1: todolist.GetAllTodoListResponse.todo:type_name -> todolist.Todo
	0,  // 2: todolist.GetTodoListResponse.todo:type_name -> todolist.Todo
	0,  // 3: todolist.UpdateRequest.newTodo:type_name -> todolist.Todo
	1,  // 4: todolist.TodoListService.Insert:input_type -> todolist.InsertRequest
	3,  // 5: todolist.TodoListService.GetAllTodoList:input_type -> todolist.GetAllTodoListRequest
	5,  // 6: todolist.TodoListService.GetTodoList:input_type -> todolist.GetTodoListRequest
	7,  // 7: todolist.TodoListService.Update:input_type -> todolist.UpdateRequest
	9,  // 8: todolist.TodoListService.Delete:input_type -> todolist.DeleteRequest
	2,  // 9: todolist.TodoListService.Insert:output_type -> todolist.InsertResponse
	4,  // 10: todolist.TodoListService.GetAllTodoList:output_type -> todolist.GetAllTodoListResponse
	6,  // 11: todolist.TodoListService.GetTodoList:output_type -> todolist.GetTodoListResponse
	8,  // 12: todolist.TodoListService.Update:output_type -> todolist.UpdateResponse
	10, // 13: todolist.TodoListService.Delete:output_type -> todolist.DeleteResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_todolist_todolist_proto_init() }
func file_todolist_todolist_proto_init() {
	if File_todolist_todolist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todolist_todolist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_todolist_todolist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_todolist_todolist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResponse); i {
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
		file_todolist_todolist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTodoListRequest); i {
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
		file_todolist_todolist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTodoListResponse); i {
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
		file_todolist_todolist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoListRequest); i {
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
		file_todolist_todolist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodoListResponse); i {
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
		file_todolist_todolist_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_todolist_todolist_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_todolist_todolist_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_todolist_todolist_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_todolist_todolist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todolist_todolist_proto_goTypes,
		DependencyIndexes: file_todolist_todolist_proto_depIdxs,
		MessageInfos:      file_todolist_todolist_proto_msgTypes,
	}.Build()
	File_todolist_todolist_proto = out.File
	file_todolist_todolist_proto_rawDesc = nil
	file_todolist_todolist_proto_goTypes = nil
	file_todolist_todolist_proto_depIdxs = nil
}
