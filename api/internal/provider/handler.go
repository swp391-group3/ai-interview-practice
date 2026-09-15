package provider

import (
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/auth"
	"github.com/swp391-group3/ai-interview-practice/api/internal/handler"
)

func ProvideAuthHandler(authService auth.AuthService) *handler.AuthHandler {
	return handler.NewAuthHandler(authService)
}

func ProvideHealthHandler() *handler.HealthHandler {
	return handler.NewHealthHandler()
}
