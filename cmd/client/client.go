package main

import (
	"context"
	"fmt"
	"github.com/loxt/fullcycle2.0-grpc/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			log.Fatalf("Could not close gRPC Server: %v", err)
		}
	}(connection)

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Loxt",
		Email: "emannuel@hotmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}
