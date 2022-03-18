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
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/todolist/todolist"
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
