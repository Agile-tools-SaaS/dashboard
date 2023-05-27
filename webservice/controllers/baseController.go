package controllers

import (
	"github.com/Agile-tools-SaaS/dashboard/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())

	AuthController(router)
	UserController(router)
	SpacesController(router)
	TemplateController(router)
}
