package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	APIKeyHeader = "api_key"
	ValidAPIKey  = "apitest"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader(APIKeyHeader)
		if apiKey != ValidAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
