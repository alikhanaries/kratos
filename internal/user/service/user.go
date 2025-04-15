package service

import "context"

type UserService interface {
    CreateUser(ctx context.Context, username, email, password string) error
    GetUser(ctx context.Context, id int32) (string, string, error)
}
