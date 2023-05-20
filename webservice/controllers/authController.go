package controllers

import (
	"github.com/Agile-tools-SaaS/dashboard/services"

	"github.com/gin-gonic/gin"
)

func AuthController(router *gin.Engine) {
	api := router.Group("auth")
	{
		api.POST("login", services.LoginUser)
		api.POST("logout", services.LogoutUser)
		api.GET("", services.CheckUser)
	}
}
