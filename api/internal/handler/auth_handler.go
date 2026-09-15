package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swp391-group3/ai-interview-practice/api/internal/features/auth"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/response"
)

type AuthHandler struct {
	authService auth.AuthService
}

func NewAuthHandler(authService auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.Wrap(apperror.CodeValidation, "", err))
		return
	}

	tokenPair, err := h.authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:  auth.RefreshTokenCookie,
		Value: tokenPair.RefreshToken,
		Path:  auth.RefreshPath,
	})

	response.OK(c, tokenPair.AccessToken)
}
