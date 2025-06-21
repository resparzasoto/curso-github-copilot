package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns 200 if the service is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "running"})
}
