package main

import (
	"context"
	"log"
	"net"

	pb "github.com/calam1/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	log.Printf("Received: %s", req.Name)
	return &pb.HelloWorldResponse{Msg: "Hello " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listener failed:%v:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(grpcServer, &server{})
	// register reflection service on gRPC server
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
