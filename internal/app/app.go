package app

import (
	"context"
	"log"

	v1 "github.com/go-kratos/kratos/v2/api/product/service/v1"
	authservice "github.com/go-kratos/kratos/v2/internal/auth/service_impl"
	authtransport "github.com/go-kratos/kratos/v2/internal/auth/transport"
	productservice "github.com/go-kratos/kratos/v2/internal/product/service_impl"
	producttransport "github.com/go-kratos/kratos/v2/internal/product/transport"
	userservice "github.com/go-kratos/kratos/v2/internal/user/service_impl"
	usertransport "github.com/go-kratos/kratos/v2/internal/user/transport"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

// App represents the application with its dependencies
type App struct {
	grpcServer *kgrpc.Server
	httpServer *khttp.Server
	logger     *log.Logger
}

// NewApp creates a new App instance with the provided dependencies
func NewApp(grpcServer *kgrpc.Server, httpServer *khttp.Server, logger *log.Logger) *App {
	app := &App{
		grpcServer: grpcServer,
		httpServer: httpServer,
		logger:     logger,
	}

	// Register services
	app.registerServices()

	return app
}

// registerServices registers all the services with their respective transports
func (a *App) registerServices() {
	// Auth Service
	authSvc := authservice.NewAuthService(a.logger)
	authtransport.RegisterGRPC(a.grpcServer, authSvc)
	authtransport.RegisterHTTP(a.httpServer, authSvc)

	// User Service
	userSvc := userservice.NewUserServiceImpl(a.logger)
	usertransport.RegisterGRPC(a.grpcServer, userSvc)
	usertransport.RegisterHTTP(a.httpServer, userSvc)

	// Product Service
	productSvc := productservice.NewProductServiceImpl()
	v1.RegisterProductServiceServer(a.grpcServer, productSvc)

	// Register HTTP handler for product service
	productHandler := producttransport.NewProductServiceHandler(productSvc)
	a.httpServer.Handle("/products", productHandler)
}

// Start starts all the servers in the application
func (a *App) Start(ctx context.Context) error {
	// Log server information
	a.logServerInfo()

	// Start gRPC server in a goroutine
	go func() {
		if err := a.grpcServer.Start(ctx); err != nil {
			a.logger.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// Start HTTP server
	if err := a.httpServer.Start(ctx); err != nil {
		return err
	}

	return nil
}

// logServerInfo logs information about the running servers
func (a *App) logServerInfo() {
	a.logger.Println("Application starting...")
	a.logger.Println("Registered endpoints:")
	a.logger.Println("- gRPC:")
	a.logger.Println("  - Auth Service")
	a.logger.Println("  - User Service")
	a.logger.Println("  - Product Service")
	a.logger.Println("- HTTP:")
	a.logger.Println("  - /products")
	a.logger.Println("  - /auth")
	a.logger.Println("  - /user")
}

// Stop stops all the servers in the application
func (a *App) Stop(ctx context.Context) error {
	if err := a.grpcServer.Stop(ctx); err != nil {
		return err
	}
	if err := a.httpServer.Stop(ctx); err != nil {
		return err
	}
	return nil
}
