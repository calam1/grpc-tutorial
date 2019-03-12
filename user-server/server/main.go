package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "github.com/calam1/server/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type userServer struct {
	db *sql.DB
}

// NewUserServiceServer returns implemented server
func NewUserServiceServer(db *sql.DB) pb.UserServiceServer {
	return &userServer{db: db}
}

func (u *userServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := u.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to db"+err.Error())
	}
	return c, nil
}

// Create create a user
func (u *userServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	c, err := u.connect(ctx)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	result, err := c.ExecContext(ctx, "INSERT INTO User(`FirstName`, `LastName`, `Address`) VALUES (?, ?, ?)", req.User.FirstName, req.User.LastName, req.User.Address)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into User table"+err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve last inserted record for User"+err.Error())
	}

	return &pb.CreateUserResponse{
		Id: id,
	}, nil
}

// Read get a user
func (u *userServer) Read(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	c, err := u.connect(ctx)
	if err != nil {
		return nil, err
	}

	defer c.Close()

	result, err := c.QueryContext(ctx, "SELECT `ID`, `FIRSTNAME`, `LASTNAME`, `ADDRESS` FROM User WHERE `ID` = ?", req.Id)

	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select id from User table"+err.Error())
	}

	defer result.Close()

	var usr pb.User

	if result.Next() {
		if err := result.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Address); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve values from User"+err.Error())
		}
	} else {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with id %d is not found", req.Id))
	}

	return &pb.ReadUserResponse{
		User: &usr,
	}, nil
}

func (u *userServer) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	c, err := u.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	result, err := c.ExecContext(ctx, "DELETE FROM User WHERE ID = ?", req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete User"+err.Error())
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve rows affected"+err.Error())
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("User with id %d is not found", req.Id))
	}

	return &pb.DeleteUserResponse{
		RowsDeleted: rows,
	}, nil
}

func main() {

	dsUser := "chris:password@tcp(172.17.0.2:3306)/customers?User"

	db, err := sql.Open("mysql", dsUser)
	if err != nil {
		log.Fatalf("error opening database %v", err)
	}

	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listener failed:%v", err)
	}

	userAPI := NewUserServiceServer(db)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userAPI)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
