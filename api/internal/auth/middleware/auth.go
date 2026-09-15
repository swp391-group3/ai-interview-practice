package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/swp391-group3/ai-interview-practice/api/internal/shared/config"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/response"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/token"
)

const currentUserIDKey = "auth.middleware.currentUserID"

// RequireAuth authenticates requests using the configured access-token secret.
// Access and refresh secrets must be distinct to reject refresh tokens.
func RequireAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if cfg == nil ||
			strings.TrimSpace(cfg.JWTAccessSecret) == "" ||
			cfg.JWTAccessSecret == cfg.JWTRefreshSecret {
			response.Error(c, apperror.New(
				apperror.CodeInternal,
				"authentication is not configured correctly",
			))
			c.Abort()
			return
		}
		parts := strings.Fields(c.GetHeader("Authorization"))
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Error(c, apperror.New(apperror.CodeInvalidToken, "a bearer token is required"))
			c.Abort()
			return
		}

		id, err := token.ParseToken(parts[1], cfg.JWTAccessSecret)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		c.Set(currentUserIDKey, id)
		c.Next()
	}
}

// CurrentUserID returns the account UUID stored by RequireAuth.
func CurrentUserID(c *gin.Context) (uuid.UUID, bool) {
	if c == nil {
		return uuid.Nil, false
	}
	value, exists := c.Get(currentUserIDKey)
	if !exists {
		return uuid.Nil, false
	}
	id, ok := value.(uuid.UUID)
	return id, ok
}
