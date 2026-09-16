package auth

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/swp391-group3/ai-interview-practice/api/internal/config"
	"github.com/swp391-group3/ai-interview-practice/api/internal/features/auth/repository"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/token"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string) (*TokenPair, error)
	GenerateTokenPair(id uuid.UUID) (*TokenPair, error)
}

type Service struct {
	cfg     *config.Config
	queries *repository.Queries
}

func NewService(cfg *config.Config, queries *repository.Queries) AuthService {
	return &Service{
		cfg:     cfg,
		queries: queries,
	}
}

func (s *Service) Login(ctx context.Context, email string, password string) (*TokenPair, error) {
	account, err := s.queries.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, apperror.Wrap(apperror.CodeAccountNotFound, "account with given email does not exist", err)
	}

	if account.IsLocked {
		return nil, apperror.New(apperror.CodeAccountLocked, "account is locked")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(password)); err != nil {
		return nil, apperror.Wrap(apperror.CodeInvalidCredentials, "wrong password", err)
	}

	return s.GenerateTokenPair(account.ID)
}

func (s *Service) GenerateTokenPair(id uuid.UUID) (*TokenPair, error) {
	refreshExpirySec := int(s.cfg.JWT.RefreshTokenExpiry.Seconds())
	accessExpirySec := int(s.cfg.JWT.AccessTokenExpiry.Seconds())

	refresh, err := token.GenerateToken(id, s.cfg.JWT.RefreshSecret, refreshExpirySec)
	if err != nil {
		return nil, err
	}

	access, err := token.GenerateToken(id, s.cfg.JWT.AccessSecret, accessExpirySec)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		RefreshToken: refresh,
		AccessToken:  access,
	}, nil
}
