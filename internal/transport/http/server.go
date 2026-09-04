package http

import (
	"fmt"
	"net/http"
	"time"

	"ai-interview-practice-api/internal/config"
	"ai-interview-practice-api/internal/service"
)

type Server struct {
	cfg     *config.Config
	authSvc service.AuthService
}

func New(cfg *config.Config, authSvc service.AuthService) Server {
	return Server{
		cfg:     cfg,
		authSvc: authSvc,
	}
}

func (s Server) Build() *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
