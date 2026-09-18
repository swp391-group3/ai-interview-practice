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

// Login godoc
// @Summary Log in with email and password
// @Description Returns an access token in data and sets the refresh cookie scoped to /auth/refresh. A refresh endpoint is not currently registered.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body auth.LoginRequest true "Login credentials"
// @Success 200 {object} response.Envelope{data=string} "Access token"
// @Header 200 {string} Set-Cookie "Refresh token cookie; Path=/auth/refresh"
// @Failure 400 {object} response.Envelope "Invalid request"
// @Failure 401 {object} response.Envelope "Invalid credentials"
// @Failure 403 {object} response.Envelope "Account locked"
// @Failure 404 {object} response.Envelope "Account not found"
// @Failure 500 {object} response.Envelope "Internal error"
// @Router /auth/login [post]
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
