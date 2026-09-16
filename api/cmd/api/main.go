package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/logger"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	app, err := InitializeApplication(configPath)
	if err != nil {
		fmt.Printf("Failed to initialize application: %v\n", err)
		os.Exit(1)
	}

	app.Logger.Info("Starting AI Interview Practice API",
		logger.String("environment", app.Config.App.Environment),
		logger.String("version", app.Config.App.Version),
	)

	// Clean up resources when application shuts down
	defer func() {
		if app.Pool != nil {
			app.Logger.Info("Closing database connection pool...")
			app.Pool.Close()
		}
		if app.Tracer != nil {
			_ = app.Tracer.Shutdown(context.Background())
		}
		_ = app.Logger.Sync()
	}()

	// Run HTTP server in a goroutine
	go func() {
		if err := app.Server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatal("HTTP server failed to start", logger.Error(err))
		}
	}()

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Logger.Info("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), app.Config.Server.ShutdownTimeout)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		app.Logger.Fatal("Server forced to shutdown", logger.Error(err))
	}

	app.Logger.Info("Server exited properly")
}
