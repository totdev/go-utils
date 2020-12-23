package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler returns a handler function for API's PING
func PingHandler(apiName string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"apiName": apiName,
		})
	}
}
