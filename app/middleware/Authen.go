package middleware

import (
	"ginapi-gateway/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware authenticates requests using token in header
func TokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("TokenAuthMiddleware: No Authorization header provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			return
		}

		token := authHeader[len(BearerSchema):]
		token = strings.TrimSpace(token)

		// Here, implement your logic to validate the token
		if err := services.ValidAccessToken(token); err != nil {
			log.Printf("TokenAuthMiddleware: Invalid token: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			return
		}

		log.Println("TokenAuthMiddleware: Token validation successful")
		c.Next()
	}
}
