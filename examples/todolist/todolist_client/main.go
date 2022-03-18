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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/todolist/todolist"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoListServiceClient(conn)

	// insertTodo(c, "Ngay mai", "Di choi", "Di choi tai noi nao do voi ny")
	// getAllTodoList(c)
	// getTodoList(c, 6)
	// updateTodo(c, &pb.Todo{
	// 	Id:           1,
	// 	Time:         "Ngay kia",
	// 	Action:       "Hoc bai",
	// 	DetailAction: "Hoc golang",
	// })
	deleteContact(c, 6)
}

func insertTodo(c pb.TodoListServiceClient, time, action, detailAction string) {
	req := &pb.InsertRequest{
		Todo: &pb.Todo{
			Time:         time,
			Action:       action,
			DetailAction: detailAction,
		},
	}
	resp, err := c.Insert(context.Background(), req)

	if err != nil {
		log.Printf("Call insert err %v\n", err)
		return
	}

	log.Printf("Insert response %+v\n", resp)
}

func getAllTodoList(c pb.TodoListServiceClient) {
	req := &pb.GetAllTodoListRequest{}

	resp, err := c.GetAllTodoList(context.Background(), req)
	if err != nil {
		log.Printf("call read all err %v\n", err)
		return
	}

	log.Printf("read all response %+v\n", resp.GetTodo())
}

func getTodoList(c pb.TodoListServiceClient, id int64) {
	req := &pb.GetTodoListRequest{
		Id: id,
	}

	resp, err := c.GetTodoList(context.Background(), req)
	if err != nil {
		log.Printf("call read err %v\n", err)
		return
	}

	log.Printf("read response %+v\n", resp.GetTodo())
}

func updateTodo(cli pb.TodoListServiceClient, updateTodo *pb.Todo) {
	req := &pb.UpdateRequest{
		NewTodo: updateTodo,
	}

	resp, err := cli.Update(context.Background(), req)
	if err != nil {
		log.Printf("call update err %v\n", err)
		return
	}

	log.Printf("update response %+v\n", resp)
}

func deleteContact(cli pb.TodoListServiceClient, id int64) {
	req := &pb.DeleteRequest{
		Id: id,
	}
	resp, err := cli.Delete(context.Background(), req)
	if err != nil {
		log.Printf("call delete err %v\n", err)
		return
	}

	log.Printf("delete response %+v\n", resp)
}
