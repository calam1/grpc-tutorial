package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/calam1/helloworld/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "default_name"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}

	defer conn.Close()

	client := pb.NewHelloWorldServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHelloWorld(ctx, &pb.HelloWorldRequest{Name: name})
	if err != nil {
		log.Fatalf("cannot contact helloworld server: %v", err)
	}

	log.Printf("Response is: %s", response.Msg)
}
