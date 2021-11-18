package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestHandler -> function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler -> creates a new request handler
func NewRequestHandler() RequestHandler {
	engine := gin.New()
	engine.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "API Up and Running"})
	})
	return RequestHandler{Gin: engine}
}
