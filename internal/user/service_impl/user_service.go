package service_impl

import (
	"context"
	"log"

	v1 "github.com/go-kratos/kratos/v2/api/user/service/v1"
	// v1 "kratos/api/user/service/v1"
)

type UserServiceImpl struct {
	log *log.Logger
	v1.UnimplementedUserServiceServer
}

func NewUserServiceImpl(logger *log.Logger) *UserServiceImpl {
	return &UserServiceImpl{log: logger}
}

func (s *UserServiceImpl) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	s.log.Printf("Registering user: %+v\n", req)
	return &v1.UserReply{
		Id:       1,
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
	}, nil
}

func (s *UserServiceImpl) GetUser(ctx context.Context, req *v1.UserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		Id:       req.Id,
		Username: "demo_user",
		Email:    "demo@example.com",
		Age:      25,
	}, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	s.log.Printf("Updating user: %+v\n", req)
	return &v1.UserReply{
		Id:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Age:      req.Age,
	}, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *v1.UserRequest) (*v1.DeleteReply, error) {
	s.log.Printf("Deleting user ID: %d\n", req.Id)
	return &v1.DeleteReply{Message: "User deleted successfully"}, nil
}

func (s *UserServiceImpl) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersReply, error) {
	return &v1.ListUsersReply{
		Users: []*v1.UserReply{
			{Id: 1, Username: "user1", Email: "user1@example.com", Age: 30},
			{Id: 2, Username: "user2", Email: "user2@example.com", Age: 28},
		},
	}, nil
}
