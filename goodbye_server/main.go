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
	"time"
	"math/rand"
)

type Config struct {
	Port string `default:"50555"`
}

type server struct {}


func (s *server) SayGoodbye(ctx context.Context, in *pb.GoodbyeRequest) (*pb.GoodbyeReply, error) {
	hostname , _ := os.Hostname()

	rand.Seed(time.Now().Unix())
	// http://iyashitour.com/archives/36787/2
	chandler_words := []string{
	"There are two kinds of truth: the truth that lights the way and the truth that warms the heart. The first of these is science, and the second is art. Without art, science would be as useless as a pair of high forceps in the hands of a plumber. Without science, art would become a crude mess of folklore and emotional quackery.",
	"The more you reason the less you create.",
	"I certainly admire people who do things.",
	"If I wasn’t hard, I wouldn’t be alive. If I couldn’t ever be gentle, I wouldn’t deserve to be alive.",
	}
	n := rand.Int() % len(chandler_words)
	return &pb.GoodbyeReply{Message: "Goodbye " + in.Name + ". At " + hostname + ".\n I sent chandler's words in Collection of famous sayings: " + chandler_words[n] }, nil
}

func main() {
	var config Config
	err := envconfig.Process("grpc", &config)
	lis, err := net.Listen("tcp", ":" + config.Port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoodbyeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}