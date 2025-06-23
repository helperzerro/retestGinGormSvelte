package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewHelloWorldController untuk Gin
func EmployeeController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	}
}
