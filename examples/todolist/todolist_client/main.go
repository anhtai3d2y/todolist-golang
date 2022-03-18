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

	insertTodo(c, "Ngay mai", "Di choi", "Di choi tai noi nao do voi ny")
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
