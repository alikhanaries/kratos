package main

import (
	"context"
	"log"
	"os"

	"github.com/go-kratos/kratos/v2/cmd/app/wire" // Update this path if your wire.go is elsewhere
)

func main() {
	// Set up logger
	logger := log.New(os.Stdout, "[app] ", log.LstdFlags)

	// Initialize the application with Wire
	app, err := wire.InitializeApp(logger)
	if err != nil {
		logger.Fatalf("failed to initialize application: %v", err)
	}

	// Create a base context
	ctx := context.Background()

	// Start the application
	if err := app.Start(ctx); err != nil {
		logger.Fatalf("failed to start application: %v", err)
	}
}
