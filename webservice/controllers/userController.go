package controllers

import (
	"github.com/Agile-tools-SaaS/dashboard/services"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.Engine) {

	api := router.Group("user")
	{
		api.POST("", services.CreateUser)
		api.PUT(":user", services.ChangeUserDetails)
		api.PUT(":user/changepassword", services.ChangePassword)
		api.DELETE(":user", services.DeleteUser)

		api.GET(":user", services.FindOneUser)
		api.GET(":user/spaces", services.GetSpacesByUserAndFilterWithPagination)

	}
}
