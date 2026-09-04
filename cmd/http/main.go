package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"ai-interview-practice-api/internal/config"
	"ai-interview-practice-api/internal/service"
	server "ai-interview-practice-api/internal/transport/http"
	"ai-interview-practice-api/internal/util"

	"github.com/caarlos0/env/v11"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}

func build() (*http.Server, error) {
	cfg, err := env.ParseAs[config.Config]()
	if err != nil {
		return nil, err
	}
	pool, err := util.NewDatabasePool(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	authSvc := service.NewAuthService(pool, &cfg)

	return server.New(&cfg, authSvc).Build(), nil
}

func main() {
	server, err := build()
	if err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
