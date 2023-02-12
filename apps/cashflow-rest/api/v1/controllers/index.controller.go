package controllers

import "github.com/gin-gonic/gin"

// Health Check             godoc
// @Summary      Health Check to see if api server is running
// @Description  Responds with simple "OK"
// @Tags         Health Check
// @Produce      json
// @Success      200
// @Router       / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
