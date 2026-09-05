package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"ai-interview-practice-api/internal/shared/config"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string) (*TokenPair, error)
}

type AuthServiceImpl struct {
	cfg  *config.Config
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool, cfg *config.Config) AuthService {
	return &AuthServiceImpl{
		cfg:  cfg,
		pool: pool,
	}
}
