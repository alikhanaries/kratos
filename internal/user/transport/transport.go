package transport

import (
	//"kratos-practice/kratos/internal/user/service_impl"
	//v1 "kratos-practice/kratos/api/user/service/v1"

	v1 "github.com/go-kratos/kratos/v2/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterGRPC(srv *grpc.Server, s v1.UserServiceServer) {
	v1.RegisterUserServiceServer(srv, s)
}

func RegisterHTTP(srv *http.Server, s v1.UserServiceServer) {
	v1.RegisterUserServiceHTTPServer(srv, s)
}
