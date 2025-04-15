package transport

import (
	// v1 "kratos-practice/kratos/api/product/service/v1"
	// "kratos-practice/kratos/internal/product/service_impl"

	v1 "github.com/go-kratos/kratos/v2/api/product/service/v1"
	"github.com/go-kratos/kratos/v2/internal/product/service_impl"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterGRPC(server *grpc.Server) {
	v1.RegisterProductServiceServer(server, service_impl.NewProductServiceImpl())
}

func RegisterHTTP(server *http.Server) {
	v1.RegisterProductServiceHTTPServer(server, service_impl.NewProductServiceImpl())
}
