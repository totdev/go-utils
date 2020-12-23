package http

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AbortWithError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(
		code,
		ErrorResponse{
			Code:    code,
			Message: err.Error(),
		},
	)
}
