package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
	"github.com/swp391-group3/ai-interview-practice/api/pkg/response"
)

const refreshPath = "/auth/refresh"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"jane@example.com"`
	Password string `json:"password" binding:"required" example:"SuperSecret123"`
}

func (s *Server) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.Wrap(apperror.CodeValidation, "", err))
		return
	}

	tokenPair, err := s.svc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}
	c.SetCookieData(&http.Cookie{
		Name:  refreshTokenCookie,
		Value: tokenPair.RefreshToken,
		Path:  refreshPath,
	})

	response.OK(c, tokenPair.AccessToken)
}
