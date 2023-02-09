package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/configs"
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
