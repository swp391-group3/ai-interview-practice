package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/swp391-group3/ai-interview-practice/api/internal/auth/service"
	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/config"
)

type Server struct {
	cfg *config.Config
	svc service.AuthService
}

func New(cfg *config.Config, pool *pgxpool.Pool) *Server {
	return &Server{
		cfg: cfg,
		svc: service.New(pool, cfg),
	}
}

func (s *Server) RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/auth/login", s.Login)
}
