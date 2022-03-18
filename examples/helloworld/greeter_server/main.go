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
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSumServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum called")
	resp := &pb.SumResponse{
		Result: in.GetNum1() + in.GetNum2(),
	}

	return resp, nil
}

func (s *server) SumWithDeadline(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum with deadline called")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("Contex.Cancel...")
			return nil, status.Errorf(codes.Canceled, "client cancel req")
		}
		time.Sleep(1 * time.Second)
	}

	resp := &pb.SumResponse{
		Result: in.GetNum1() + in.GetNum2(),
	}

	return resp, nil
}

func (s *server) PND(in *pb.PNDRequest, stream pb.SumService_PNDServer) error {
	log.Printf("PND called")

	k := int32(2)
	N := in.GetNum()
	for N > 1 {
		if N%k == 0 {
			N = N / k
			//send to client
			stream.Send(&pb.PNDResponse{
				Result: k,
			})
			time.Sleep(1000 * time.Millisecond)
		} else {
			k++
			log.Printf("k increase to: %v", k)
		}
	}
	return nil
}

func (s *server) Average(stream pb.SumService_AverageServer) error {

	log.Println("Average called...")
	var total float32
	var count int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Calculate Average and return
			resp := &pb.AverageResponse{
				Result: total / float32(count),
			}

			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatalf("err while Recv Average %v", err)
			return err
		}
		log.Printf("receive req %v", req)
		total += req.GetNum()
		count++
	}
}

func (s *server) FindMax(stream pb.SumService_FindMaxServer) error {

	log.Println("Find max called...")
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF...")
			return nil
		}
		if err != nil {
			log.Fatalf("err while Recv FindMax %v", err)
			return err
		}

		num := req.GetNum()
		log.Printf("Recv number: %v", num)
		if num > max {
			max = num
		}
		err = stream.Send(&pb.FindMaxResponse{
			Max: max,
		})
		if err != nil {
			log.Printf("Send max err %v", err)
			return err
		}
	}
}

func (s *server) Square(ctx context.Context, req *pb.SquareRequest) (*pb.SquareResponse, error) {

	log.Println("Square called...")
	num := req.GetNum()
	if num < 0 {
		log.Printf("req num < 0, num=%v, return InvalidArgument", num)
		return nil, status.Errorf(codes.InvalidArgument, "Expect num > 0, req num was %v", num)
	}

	return &pb.SquareResponse{
		Result: math.Sqrt(float64(num)),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &server{})
	log.Printf("server todo list listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
