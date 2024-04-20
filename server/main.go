package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dalpengida/portfolio-grpc-go/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "the server post")
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello dalpengida" + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{}) //client가 사용할 수 있도록 등록
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
