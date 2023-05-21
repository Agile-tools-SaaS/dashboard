package controllers

import (
	services "github.com/Agile-tools-SaaS/dashboard/services/space"

	"github.com/gin-gonic/gin"
)

func SpacesController(router *gin.Engine) {

	api := router.Group("space")
	{
		api.POST("", services.CreateSpace)
		api.GET(":space_id", services.GetSpaceById)
		api.PUT(":space_id", services.ChangeSpaceDetails)
		api.DELETE(":space_id", services.DeleteSpace)

		api.POST(":space_id/user/:user_id", services.AddUserToSpace)
		api.POST(":space_id/user/:user_id/makeadmin", services.MakeUserAnAdminOfSpace)
		api.PUT(":space_id/user/:user_id", services.RemoveUserFromSpace)
		api.PUT(":space_id/user/:user_id/removeadmin", services.RemoveUserFromAdminOfSpace)

		api.POST(":space_id/file", services.AddFileToSpace)
		api.GET(":space_id/file/:file_id", services.GetSpaceFileByFileNameAndSpaceName)
		api.PUT(":space_id/file/:file_id", services.EditFileInSpace)
		api.DELETE(":space_id/file/:file_id", services.DeleteFileFromSpace)
	}
}
