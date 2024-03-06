package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck responds to the health check request
// @Summary Health check
// @Description Get the health status of the API
// @Tags health
// @Accept  json
// @Produce  json
// @Param  Authorization  header  string  true  "token"
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
