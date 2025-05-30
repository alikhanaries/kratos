// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v6.30.2
// source: kratos/api/user/service/v1/user.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceDeleteUser = "/user.service.v1.UserService/DeleteUser"
const OperationUserServiceGetUser = "/user.service.v1.UserService/GetUser"
const OperationUserServiceListUsers = "/user.service.v1.UserService/ListUsers"
const OperationUserServiceRegister = "/user.service.v1.UserService/Register"
const OperationUserServiceUpdateUser = "/user.service.v1.UserService/UpdateUser"

type UserServiceHTTPServer interface {
	// DeleteUser Delete user
	DeleteUser(context.Context, *UserRequest) (*DeleteReply, error)
	// GetUser Get user by ID
	GetUser(context.Context, *UserRequest) (*UserReply, error)
	// ListUsers List all users
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersReply, error)
	// Register Create a new user
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	// UpdateUser Update user
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/users/register", _UserService_Register0_HTTP_Handler(srv))
	r.GET("/v1/users/{id}", _UserService_GetUser0_HTTP_Handler(srv))
	r.PUT("/v1/users/{id}", _UserService_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("/v1/users/{id}", _UserService_DeleteUser0_HTTP_Handler(srv))
	r.GET("/v1/users", _UserService_ListUsers0_HTTP_Handler(srv))
}

func _UserService_Register0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*UserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_DeleteUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*UserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_ListUsers0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUsersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceListUsers)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUsers(ctx, req.(*ListUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUsersReply)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	DeleteUser(ctx context.Context, req *UserRequest, opts ...http.CallOption) (rsp *DeleteReply, err error)
	GetUser(ctx context.Context, req *UserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	ListUsers(ctx context.Context, req *ListUsersRequest, opts ...http.CallOption) (rsp *ListUsersReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) DeleteUser(ctx context.Context, in *UserRequest, opts ...http.CallOption) (*DeleteReply, error) {
	var out DeleteReply
	pattern := "/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetUser(ctx context.Context, in *UserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...http.CallOption) (*ListUsersReply, error) {
	var out ListUsersReply
	pattern := "/v1/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceListUsers))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/users/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/v1/users/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
