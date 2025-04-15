// Step 3: gRPC + HTTP transport in `internal/auth/transport/transport.go`
package transport

import (
	pb "github.com/go-kratos/kratos/v2/api/auth/service/v1"
	//"kratos/internal/auth/service_impl"

	"github.com/go-kratos/kratos/v2/internal/auth/service_impl"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func RegisterGRPC(server *grpc.Server, auth *service_impl.AuthServiceImpl) {
	pb.RegisterAuthServiceServer(server, auth)
}

func RegisterHTTP(server *http.Server, auth *service_impl.AuthServiceImpl) {
	pb.RegisterAuthServiceHTTPServer(server, auth)
}
