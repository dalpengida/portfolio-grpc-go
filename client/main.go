package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/dalpengida/portfolio-grpc-go/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn) //서버의 method를 사용할 수 있게 해줌

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	r, err = c.SayHello(ctx, &helloworld.HelloRequest{Name: "dalpengida"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SayHello: %s", r.Message)

}
