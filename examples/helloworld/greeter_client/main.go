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
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/status"
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
	c := pb.NewSumServiceClient(conn)
	// callSum(c)
	callSumWithDeadline(c, 1*time.Second)
	callSumWithDeadline(c, 5*time.Second)
	// callPND(c)
	// callAverage(c)
	// callFindMax(c)
	// callSquare(c, -3)
}

func callSum(c pb.SumServiceClient) {
	log.Println("Calling sum api")
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 7,
		Num2: 6,
	})

	if err != nil {
		log.Fatalf("Call sum api err %v", err)
	}

	log.Printf("Sum api response %v \n", resp.GetResult())

}

func callSumWithDeadline(c pb.SumServiceClient, timeout time.Duration) {
	log.Println("Calling sum with deadline api")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.SumWithDeadline(ctx, &pb.SumRequest{
		Num1: 7,
		Num2: 6,
	})

	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("calling sum with deadline api with DeadlineExceeded")
			} else {
				log.Printf("Call sum with deadline api err %v", err)
			}
		} else {
			log.Fatalf("Call sum with deadline unknow err %v", err)
		}
		return
	}

	log.Printf("Sum with deadline api response %v \n", resp.GetResult())

}

func callPND(c pb.SumServiceClient) {
	log.Println("Calling PND api")
	stream, err := c.PND(context.Background(), &pb.PNDRequest{
		Num: 120,
	})

	if err != nil {
		log.Fatalf("Call PND err %v", err)
	}

	for {
		resp, recvErr := stream.Recv()
		if recvErr == io.EOF {
			log.Println("Server finish streaming")
			return
		}

		log.Printf("recv number: %v", resp.GetResult())
	}
}

func callAverage(c pb.SumServiceClient) {
	log.Println("calling average api")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("call average err %v", err)
	}

	listReq := []pb.AverageRequest{
		pb.AverageRequest{
			Num: 5,
		},
		pb.AverageRequest{
			Num: 10,
		},
		pb.AverageRequest{
			Num: 12,
		},
		pb.AverageRequest{
			Num: 3,
		},
		pb.AverageRequest{
			Num: 4.2,
		},
	}

	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("send average request err %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("receive average response err %v", err)
	}

	log.Printf("average response %+v", resp)
}

func callFindMax(c pb.SumServiceClient) {
	log.Println("calling find max api")

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("call find max err %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		// send requests
		listReq := []pb.FindMaxRequest{
			pb.FindMaxRequest{
				Num: 5,
			},
			pb.FindMaxRequest{
				Num: 10,
			},
			pb.FindMaxRequest{
				Num: 12,
			},
			pb.FindMaxRequest{
				Num: 3,
			},
			pb.FindMaxRequest{
				Num: 23,
			},
			pb.FindMaxRequest{
				Num: 12,
			},
			pb.FindMaxRequest{
				Num: 34,
			},
			pb.FindMaxRequest{
				Num: 54,
			},
		}

		for _, req := range listReq {
			err := stream.Send(&req)
			if err != nil {
				log.Fatalf("send find max request err %v", err)
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("Ending find max")
				break
			}
			if err != nil {
				log.Fatalf("recv find max request err %v", err)
				break
			}

			log.Printf("Max: %v\n", resp.GetMax())
		}
		close(waitc)
	}()

	<-waitc
}

func callSquare(c pb.SumServiceClient, num int32) {
	log.Println("calling square api")
	resp, err := c.Square(context.Background(), &pb.SquareRequest{
		Num: num,
	})
	if err != nil {
		log.Printf("call square err %v", err)
		if errStatus, ok := status.FromError(err); ok {
			log.Printf("Error message: %v", errStatus.Message())
			log.Printf("Error code: %v", errStatus.Code())
			if errStatus.Code() == codes.InvalidArgument {
				log.Printf("InvalidArgument num %v", num)
				return
			}
		}
	}

	log.Printf("Square response = %v", resp.GetResult())
}
