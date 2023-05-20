package controllers

import (
	"github.com/Agile-tools-SaaS/dashboard/helpers"

	"github.com/Agile-tools-SaaS/dashboard/services"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.Engine) {
	userContext := helpers.NewContext("users")
	api := router.Group("user")
	{
		// new user
		api.POST("", func(ctx *gin.Context) {
			services.CreateUser(ctx, userContext)
		})

		// change user details
		api.PUT(":username", func(ctx *gin.Context) {
			services.ChangeUserDetails(ctx, userContext)
		})

		// change users password
		api.PUT(":username/changepassword", func(ctx *gin.Context) {
			services.ChangePassword(ctx, userContext)
		})

		// delete user
		api.DELETE(":username", func(ctx *gin.Context) {
			services.DeleteUser(ctx, userContext)
		})
	}
}
