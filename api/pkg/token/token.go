package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/swp391-group3/ai-interview-practice/pkg/apperror"
)

func GenerateToken(
	id uuid.UUID,
	secret string,
	expiry int,
) (string, error) {
	exp := time.Duration(expiry) * time.Second
	claim := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   id.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", apperror.Wrap(apperror.CodeInternal, "something went wrong", err)
	}

	return tokenStr, nil
}

func ParseToken(
	tokenStr string,
	secret string,
) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return uuid.Nil, apperror.Wrap(apperror.CodeInvalidToken, "token is invalid, expired, or revoked", err)
	}
	if !token.Valid {
		return uuid.Nil, apperror.New(apperror.CodeInvalidToken, "token is invalid, expired, or revoked")
	}

	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, apperror.Wrap(apperror.CodeInvalidToken, "token contain invalid structure", err)
	}

	return id, nil
}
