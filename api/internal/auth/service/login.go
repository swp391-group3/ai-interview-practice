package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/swp391-group3/ai-interview-practice/api/internal/auth/repository"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
)

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
