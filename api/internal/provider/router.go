package provider

import (
	"github.com/gin-gonic/gin"

	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	"github.com/swp391-group3/ai-interview-practice/api/internal/handler"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/logger"
	"github.com/swp391-group3/ai-interview-practice/api/internal/router"
	"github.com/swp391-group3/ai-interview-practice/api/internal/server"
)

func ProvideRouter(
	cfg *config.Config,
	log *logger.Logger,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
	jdHandler *handler.JDHandler,
) *gin.Engine {
	appRouter := router.NewRouter(cfg, log, authHandler, healthHandler, jdHandler)
	return appRouter.Setup()
}

func ProvideHTTPServer(
	cfg *config.Config,
	engine *gin.Engine,
	log *logger.Logger,
) *server.Server {
	return server.NewServer(cfg.Server, engine, log)
}
