/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/todolist/todolist"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	connectStr := "root:@tcp(127.0.0.1)/todolist?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", connectStr, 100, 100)
	if err != nil {
		log.Panicf("Register database err %v", err)
	}

	orm.RegisterModel(new(TodoInfo))
	err = orm.RunSyncdb("default", false, false)
	if err != nil {
		log.Panicf("Run migrate database err %v", err)
	}

	fmt.Println("Connect database successfully!")

}

// server is used to implement todolist.GreeterServer.
type server struct {
	pb.UnimplementedTodoListServiceServer
}

func (s *server) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	log.Printf("Calling insert %+v\n", req.Todo)
	ti := ConvertPbTodo2TodoInfo(req.Todo)

	err := ti.Insert()
	if err != nil {
		resp := &pb.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("Insert %+v err %v\n", ti, err),
		}
		return resp, nil
		// return nil, status.Errorf(codes.InvalidArgument, "Insert %+v err %v", ti, err)
	}

	resp := &pb.InsertResponse{
		StatusCode: 1,
		Message:    "Insert ok",
	}

	return resp, nil
}

func (s *server) GetAllTodoList(ctx context.Context, req *pb.GetAllTodoListRequest) (*pb.GetAllTodoListResponse, error) {
	log.Printf("calling read all")
	ti, err := GetAll()
	if err == orm.ErrNoRows {
		return nil, status.Errorf(codes.InvalidArgument, "Table empty")
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Get all todo list err %v", err)
	}

	return &pb.GetAllTodoListResponse{
		Todo: ConvertTodoInfo2PbTodo(ti),
	}, nil

}

func (s *server) GetTodoList(ctx context.Context, req *pb.GetTodoListRequest) (*pb.GetTodoListResponse, error) {
	log.Printf("calling read %v\n", req.GetId())
	ti, err := Get(req.GetId())
	if err == orm.ErrNoRows {
		return nil, status.Errorf(codes.InvalidArgument, "Id %s not exist", req.GetId())
	}
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Get id %s err %v", req.GetId(), err)
	}

	return &pb.GetTodoListResponse{
		Todo: ConvertTodoInfo2PbTodo(ti),
	}, nil

}

func (server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	if req.GetNewTodo() == nil {
		return nil, status.Error(codes.InvalidArgument, "update req with nil todo")
	}
	log.Printf("calling update with data %v\n", req.GetNewTodo())
	ci := ConvertPbTodo2TodoInfo(req.GetNewTodo())
	err := ci.Update()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "update %+v err %v", req.GetNewTodo(), err)
	}

	resp := &pb.UpdateResponse{
		StatusCode: 1,
		Message:    "Update successfully",
	}

	return resp, nil
}

func (server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Printf("calling delete %v\n", req.GetId())
	ci := &TodoInfo{
		Id: req.GetId(),
	}
	err := ci.Delete()
	if err != nil {
		return &pb.DeleteResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("delete todo item %v err %v", req.GetId(), err),
		}, nil
	}

	return &pb.DeleteResponse{
		StatusCode: 1,
		Message:    fmt.Sprintf("delete todo item %v successfully", req.GetId()),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodoListServiceServer(s, &server{})
	log.Printf("Server todo list listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
