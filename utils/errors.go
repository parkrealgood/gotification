package utils

import (
	"github.com/gin-gonic/gin"
)


type ErrorResponse struct {
    Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
    Message string `json:"message"`
    Code    string `json:"code"`
    Details string `json:"details,omitempty"`
}

// 에러 응답을 반환하는 헬퍼 함수
func RespondWithError(c *gin.Context, statusCode int, message string, code string, details string) {
    c.JSON(statusCode, ErrorResponse{
        Error: ErrorDetails{
            Message: message,
            Details: details,
        },
    })
    c.Abort()
}