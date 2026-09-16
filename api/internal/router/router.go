package router

import (
	"github.com/gin-gonic/gin"

	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	"github.com/swp391-group3/ai-interview-practice/api/internal/handler"
	"github.com/swp391-group3/ai-interview-practice/api/internal/middleware"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/logger"
)

type Router struct {
	cfg           *config.Config
	logger        *logger.Logger
	authHandler   *handler.AuthHandler
	healthHandler *handler.HealthHandler
}

func NewRouter(
	cfg *config.Config,
	logger *logger.Logger,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
) *Router {
	return &Router{
		cfg:           cfg,
		logger:        logger,
		authHandler:   authHandler,
		healthHandler: healthHandler,
	}
}

func (r *Router) Setup() *gin.Engine {
	if r.cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Middlewares từ cinema-platform
	router.Use(middleware.RecoveryMiddleware(r.logger))
	router.Use(middleware.LoggingMiddleware(r.logger))
	router.Use(middleware.CORSMiddleware(r.cfg.CORS))

	router.GET("/health", r.healthHandler.Health)

	auth := router.Group("/auth")
	{
		auth.POST("/login", r.authHandler.Login)
	}

	return router
}
