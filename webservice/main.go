package main

import (
	"github.com/Agile-tools-SaaS/dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	controllers.InitRoutes(router)

	router.Run(":8080")
}
