package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/swp391-group3/ai-interview-practice/api/internal/pkg/logger"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/response"
)

// RecoveryMiddleware bắt các panic trong quá trình xử lý request và log stack trace
func RecoveryMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("Panic recovered in HTTP request",
					logger.Any("error", err),
					logger.String("path", c.Request.URL.Path),
				)
				response.Error(c, apperror.New(apperror.CodeInternal, "Internal server error"))
				c.Abort()
			}
		}()
		c.Next()
	}
}
