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

package main

import (
	"log"
	"os"
	"time"

	pb "github.com/cw-sakamoto/grpc-example/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	defaultName = "world"
)

func main() {
	// Contact the server and print out its response.
	name := defaultName
	address := "localhost:30333"

	if len(os.Args) > 1 {
		address = os.Args[1]
		name = os.Args[2]
	}
	// Set up a connection to the server.
	// hello
	conn1, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	c1 := pb.NewHelloClient(conn1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r1, err1 := c1.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err1 != nil {
		log.Fatalf("could not hello: %v", err1)
	}
	log.Printf("Hello: %s", r1.Message)

	// goodbyd
	conn2, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c2 := pb.NewGoodbyeClient(conn2)
	r2, err := c2.SayGoodbye(ctx, &pb.GoodbyeRequest{Name: name})
	if err != nil {
		log.Fatalf("could not goodbye: %v", err)
	}
	log.Printf("Goodbye: %s", r2.Message)
}
