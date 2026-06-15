package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, data any) {
	c.JSON(code, Response{
		Code:   code,
		Status: true,
		Data:   data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Status:  false,
		Message: message,
	})
}
