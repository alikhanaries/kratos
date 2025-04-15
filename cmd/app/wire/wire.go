//go:build wireinject
// +build wireinject

package wire

import (
	"log"

	"github.com/go-kratos/kratos/v2/internal/app"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ProvideGRPCServerOptions provides the default gRPC server options
func ProvideGRPCServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.Address(":52937"),
	}
}

// ProvideHTTPServerOptions provides the default HTTP server options
func ProvideHTTPServerOptions() []http.ServerOption {
	return []http.ServerOption{
		http.Address(":52938"),
	}
}

// InitializeApp initializes the application with all its dependencies
func InitializeApp(logger *log.Logger) (*app.App, error) {
	wire.Build(
		ProvideGRPCServerOptions,
		ProvideHTTPServerOptions,
		grpc.NewServer,
		http.NewServer,
		app.NewApp,
	)
	return nil, nil
}
