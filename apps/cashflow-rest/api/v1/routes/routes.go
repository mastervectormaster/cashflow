package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mastervectormaster/cashflow/apps/cashflow-rest/api/v1/controllers"
)

func SetupRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HealthCheck)
	router.POST("/income", controllers.ReportIncome)
}
