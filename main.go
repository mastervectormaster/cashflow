package main

import (
	"cashflow/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//run database
	configs.ConnectDB()

	router.Run("localhost:6000")
}
