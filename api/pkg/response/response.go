package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/swp391-group3/ai-interview-practice/api/pkg/apperror"
)

type Envelope struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
}

type ErrorBody struct {
	Code    apperror.Code `json:"code"`
	Message string        `json:"message"`
}

var httpStatusByCode = map[apperror.Code]int{
	apperror.CodeInvalidCredentials: http.StatusUnauthorized,
	apperror.CodeAccountNotFound:    http.StatusNotFound,
	apperror.CodeAccountLocked:      http.StatusForbidden,
	apperror.CodeInvalidToken:       http.StatusUnauthorized,
	apperror.CodeValidation:         http.StatusBadRequest,
	apperror.CodeInternal:           http.StatusInternalServerError,
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Envelope{Success: true, Data: data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Envelope{Success: true, Data: data})
}

func Error(c *gin.Context, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		status, known := httpStatusByCode[appErr.Code]
		if !known {
			status = http.StatusInternalServerError
		}
		c.JSON(status, Envelope{
			Success: false,
			Error:   &ErrorBody{Code: appErr.Code, Message: appErr.Message},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, Envelope{
		Success: false,
		Error:   &ErrorBody{Code: apperror.CodeInternal, Message: "something wrong happen"},
	})
}
