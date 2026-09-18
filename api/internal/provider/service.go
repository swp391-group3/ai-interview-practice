package provider

import (
	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/auth"
	authRepo "github.com/swp391-group3/ai-interview-practice/api/internal/features/auth/repository"
)

func ProvideAuthService(cfg *config.Config, queries *authRepo.Queries) auth.AuthService {
	return auth.NewService(cfg, queries)
}
