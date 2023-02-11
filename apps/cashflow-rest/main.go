package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/configs"
)

func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run("0.0.0.0:" + configs.RestPort())
}
