//go:generate proto -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
    "net"
    "log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/cw-sakamoto/grpc-example/helloworld"
	"google.golang.org/grpc/reflection"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	Port string `default:"50555"`
}

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	hostname , _ := os.Hostname()
	return &pb.HelloReply{Message: "Hello I'm " + in.Name + ". " + in.Age + " years old. At " + hostname }, nil
}

func main() {
	var config Config
	err := envconfig.Process("grpc", &config)
	lis, err := net.Listen("tcp", ":" + config.Port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}