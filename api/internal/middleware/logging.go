package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/logger"
)

// LoggingMiddleware tự động ghi log request bằng Zap logger
func LoggingMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Info("HTTP Request",
			logger.String("method", c.Request.Method),
			logger.String("path", path),
			logger.String("query", query),
			logger.Int("status", status),
			logger.Duration("latency", latency),
			logger.String("client_ip", c.ClientIP()),
		)
	}
}
