package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swp391-group3/ai-interview-practice/api/docs"

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

	// Core middlewares
	router.Use(middleware.RecoveryMiddleware(r.logger))
	router.Use(middleware.LoggingMiddleware(r.logger))
	router.Use(middleware.CORSMiddleware(r.cfg.CORS))

	router.GET("/health", r.healthHandler.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/login", r.authHandler.Login)
	}

	return router
}
