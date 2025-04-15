package main

import (
	"context"
	"log"
	"os"

	"github.com/go-kratos/kratos/v2/cmd/app/wire"
)

func main() {
	logger := log.New(os.Stdout, "[app] ", log.LstdFlags)

	app, err := wire.InitializeApp(logger)
	if err != nil {
		logger.Fatalf("failed to initialize application: %v", err)
	}

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		logger.Fatalf("failed to start application: %v", err)
	}
}
