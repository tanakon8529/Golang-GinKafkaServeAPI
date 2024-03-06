package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ginapi-gateway/services"
	"ginapi-gateway/settings"
)

// AuthHandler handles authentication requests
// @Summary Authenticate
// @Description Authenticate user and return token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  client-id  header  string  true  "client-id"
// @Param  client-secret  header  string  true  "client-secret"
// @Success 200 {object} map[string]string
// @Router /auth [post]
func AuthHandler() gin.HandlerFunc {
	// Load environment variables
	config := settings.LoadEnv(".env")

	return func(c *gin.Context) {
		username := c.Request.Header.Get("client-id")
		password := c.Request.Header.Get("client-secret")
		if username != config.UsernameGateWay || password != config.PasswordGateWay {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// generate uuidv4 and return to client
		accesstoken, err := uuid.NewRandom()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "UUID generation failed"})
			return
		}
		tokenString := accesstoken.String()

		// store accesstoken in redis
		err = services.StoreAccessToken(tokenString, 60)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to store access token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": tokenString, "token_type": "Bearer"})
	}
}
