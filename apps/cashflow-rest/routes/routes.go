package routes

import (
	"github.com/gin-gonic/gin"
	v1Routes "github.com/mastervectormaster/cashflow/apps/cashflow-rest/api/v1/routes"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1Routes.SetupRouter(v1)
	return router
}
