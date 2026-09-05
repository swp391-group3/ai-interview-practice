package service

import (
	"github.com/google/uuid"

	"ai-interview-practice-api/pkg/token"
)

func (s *AuthServiceImpl) GenerateTokenPair(id uuid.UUID) (*TokenPair, error) {
	refresh, err := token.GenerateToken(id, s.cfg.JWTRefreshSecret, s.cfg.JWTRefreshExpiry)
	if err != nil {
		return nil, err
	}

	access, err := token.GenerateToken(id, s.cfg.JWTAccessSecret, s.cfg.JWTAccessExpiry)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		RefreshToken: refresh,
		AccessToken:  access,
	}, nil
}
