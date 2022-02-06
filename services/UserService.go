package services

import (
	"context"
	"fmt"
	"github.com/loxt/fullcycle2.0-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
	}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}