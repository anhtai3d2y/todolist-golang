// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: todolist/todolist.proto

package todolist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodoListServiceClient is the client API for TodoListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoListServiceClient interface {
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	GetAllTodoList(ctx context.Context, in *GetAllTodoListRequest, opts ...grpc.CallOption) (*GetAllTodoListResponse, error)
	GetTodoList(ctx context.Context, in *GetTodoListRequest, opts ...grpc.CallOption) (*GetTodoListResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type todoListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoListServiceClient(cc grpc.ClientConnInterface) TodoListServiceClient {
	return &todoListServiceClient{cc}
}

func (c *todoListServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoListService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) GetAllTodoList(ctx context.Context, in *GetAllTodoListRequest, opts ...grpc.CallOption) (*GetAllTodoListResponse, error) {
	out := new(GetAllTodoListResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoListService/GetAllTodoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) GetTodoList(ctx context.Context, in *GetTodoListRequest, opts ...grpc.CallOption) (*GetTodoListResponse, error) {
	out := new(GetTodoListResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoListService/GetTodoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoListService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodoListService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServiceServer is the server API for TodoListService service.
// All implementations must embed UnimplementedTodoListServiceServer
// for forward compatibility
type TodoListServiceServer interface {
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	GetAllTodoList(context.Context, *GetAllTodoListRequest) (*GetAllTodoListResponse, error)
	GetTodoList(context.Context, *GetTodoListRequest) (*GetTodoListResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedTodoListServiceServer()
}

// UnimplementedTodoListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoListServiceServer struct {
}

func (UnimplementedTodoListServiceServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedTodoListServiceServer) GetAllTodoList(context.Context, *GetAllTodoListRequest) (*GetAllTodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodoList not implemented")
}
func (UnimplementedTodoListServiceServer) GetTodoList(context.Context, *GetTodoListRequest) (*GetTodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoList not implemented")
}
func (UnimplementedTodoListServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTodoListServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTodoListServiceServer) mustEmbedUnimplementedTodoListServiceServer() {}

// UnsafeTodoListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoListServiceServer will
// result in compilation errors.
type UnsafeTodoListServiceServer interface {
	mustEmbedUnimplementedTodoListServiceServer()
}

func RegisterTodoListServiceServer(s grpc.ServiceRegistrar, srv TodoListServiceServer) {
	s.RegisterService(&TodoListService_ServiceDesc, srv)
}

func _TodoListService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoListService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_GetAllTodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).GetAllTodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoListService/GetAllTodoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).GetAllTodoList(ctx, req.(*GetAllTodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_GetTodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).GetTodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoListService/GetTodoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).GetTodoList(ctx, req.(*GetTodoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoListService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoListService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodoListService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoListService_ServiceDesc is the grpc.ServiceDesc for TodoListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.TodoListService",
	HandlerType: (*TodoListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _TodoListService_Insert_Handler,
		},
		{
			MethodName: "GetAllTodoList",
			Handler:    _TodoListService_GetAllTodoList_Handler,
		},
		{
			MethodName: "GetTodoList",
			Handler:    _TodoListService_GetTodoList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoListService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoListService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todolist/todolist.proto",
}
