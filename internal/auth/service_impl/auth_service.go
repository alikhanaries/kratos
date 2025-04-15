// Step 2: Implement the interface in `internal/auth/service_impl/auth_service.go`
package service_impl

import (
	"context"
	"errors"
	"log"

	pb "github.com/go-kratos/kratos/v2/api/auth/service/v1"
)

type AuthServiceImpl struct {
	pb.UnimplementedAuthServiceServer // for gRPC
	log                               *log.Logger
}

func NewAuthService(logger *log.Logger) *AuthServiceImpl {
	return &AuthServiceImpl{
		log: logger,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	if req.Username == "admin" && req.Password == "admin" {
		return &pb.LoginReply{Token: "fake-jwt-token"}, nil
	}
	return nil, errors.New("invalid credentials")
}
