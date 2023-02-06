package main

import (
	"cashflow-rest/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	router.Run("localhost:" + configs.RestPort())
}
