package main

import (
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/configs"
	_ "github.com/mastervectormaster/cashflow/apps/cashflow-rest/docs"
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/routes"
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

	// Setup Routes
	router := routes.SetupRouter()

	// Connect to DB
	configs.ConnectDB()

	// Run Swagger UI
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start Server
	router.Run("0.0.0.0:" + configs.RestPort())
}
