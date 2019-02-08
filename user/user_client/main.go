package main

import (
	"context"
	"log"
	"time"

	pb "github.com/calam1/user/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection error %v", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cxl()

	// create user
	reqCreate := pb.CreateUserRequest{
		User: &pb.User{
			FirstName: "Chris",
			LastName:  "Lam",
			Address:   "123 Main Street USA",
		},
	}

	respCreate, err := client.Create(ctx, &reqCreate)
	if err != nil {
		log.Fatalf("Create User failed %v", err)
	}
	log.Printf("Created User: %v", respCreate)

	id := respCreate.Id

	// read user
	reqRead := pb.ReadUserRequest{
		Id: id,
	}

	respRead, err := client.Read(ctx, &reqRead)
	if err != nil {
		log.Fatalf("read failed %v", err)
	}

	log.Printf("read response %v", respRead)

	// delete user
	reqDelete := pb.DeleteUserRequest{
		Id: id,
	}

	respDelete, err := client.Delete(ctx, &reqDelete)
	if err != nil {
		log.Fatalf("delete failed for %v on erro r%v", id, err)
	}

	log.Printf("delete response %v", respDelete)
}
