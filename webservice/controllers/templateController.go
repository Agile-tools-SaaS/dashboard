package controllers

import (
	template_services "github.com/Agile-tools-SaaS/dashboard/services/templates"
	"github.com/gin-gonic/gin"
)

func TemplateController(router *gin.Engine) {
	api := router.Group("template")
	{
		// USED WHEN GETTING ALL IN A LIST FOR SEARHCING ON
		api.GET("", template_services.GetMultipleTemplates)
		// USED WHEN APPLYING A TEMPLATE TO A BOARD
		api.GET(":template_id", template_services.GetTemplateById)

		api.POST("", template_services.CreateNewTemplate)
		api.PUT(":template_id", template_services.EditTemplateById)
		api.DELETE(":template_id", template_services.DeleteTemplateById)

		/* BOTH BELOW ARE USED TO CREATE A STANDARD FOR COMPANIES USINg COMMUNITY TEMPLATES*/

		// USED TO ADD TO SPACES Used Templates
		api.POST(":template_id/:space_id", template_services.AddTemplateToSpacesTemplateList)

		// USED TO REMOVE FROM SPACES Used Templates
		api.PUT(":template_id/:space_id", template_services.RemoveTemplateFromSpacesTemplateList)
	}
}
