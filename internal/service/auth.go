package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"ai-interview-practice-api/internal/config"
	"ai-interview-practice-api/internal/repository"
	"ai-interview-practice-api/internal/util"
	"ai-interview-practice-api/pkg/apperror"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string) (*TokenPair, error)
}

type AuthServiceImpl struct {
	pool *pgxpool.Pool
	cfg  *config.Config
}

func NewAuthService(pool *pgxpool.Pool, cfg *config.Config) AuthService {
	return &AuthServiceImpl{
		pool: pool,
		cfg:  cfg,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, email string, password string) (*TokenPair, error) {
	queries := repository.New(s.pool)

	account, err := queries.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, apperror.Wrap(apperror.CodeAccountNotFound, "account with given email does not existed", err)
	}
	if account.IsLocked {
		return nil, apperror.New(apperror.CodeAccountLocked, "account is locked")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(password)); err != nil {
		return nil, apperror.Wrap(apperror.CodeInvalidCredentials, "wrong password", err)
	}

	return s.GenerateTokenPair(account.ID)
}

func (s *AuthServiceImpl) GenerateTokenPair(id uuid.UUID) (*TokenPair, error) {
	refresh, err := util.GenerateToken(id, s.cfg.JWTRefreshSecret, s.cfg.JWTRefreshExpiry)
	if err != nil {
		return nil, err
	}

	access, err := util.GenerateToken(id, s.cfg.JWTAccessSecret, s.cfg.JWTAccessExpiry)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		RefreshToken: refresh,
		AccessToken:  access,
	}, nil
}
