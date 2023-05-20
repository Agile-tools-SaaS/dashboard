package controllers

import (
	"github.com/Agile-tools-SaaS/dashboard/helpers"

	"github.com/Agile-tools-SaaS/dashboard/services"

	"github.com/gin-gonic/gin"
)

func AuthController(router *gin.Engine) {
	userContext := helpers.NewContext("users")

	api := router.Group("auth")
	{
		// log in a user
		api.POST("login", func(ctx *gin.Context) {
			services.LoginUser(ctx, userContext)
		})
		// log out a user
		api.POST("logout", func(ctx *gin.Context) {
			services.LogoutUser(ctx, userContext)
		})

		// check user is logged in
		api.GET("", func(ctx *gin.Context) {
			services.CheckUser(ctx, userContext)
		})
	}
}
