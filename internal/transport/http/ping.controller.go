package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Pong"

	c.JSON(http.StatusOK, resp)
}
