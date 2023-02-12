package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/configs"
	_ "github.com/mastervectormaster/cashflow/apps/cashflow-rest/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Cashflow
// @version         0.1.0
// @description     Internal Cash management service API in Go using Gin framework.

// @contact.name   Kenny Stacker
// @contact.url    https://www.linkedin.com/in/kenny-stacker-8705b6251/
// @contact.email  vectormaster123456789@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3001
// @BasePath  /api/v1
func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health Check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run("0.0.0.0:" + configs.RestPort())
}
