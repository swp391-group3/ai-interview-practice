package http

import (
	"github.com/gin-gonic/gin"

	"ai-interview-practice-api/pkg/apperror"
	"ai-interview-practice-api/pkg/response"
)

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

	response.OK(c, tokenPair)
}
