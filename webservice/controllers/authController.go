package controllers

import (
	"webservice/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthController(db *mongo.Client, router *gin.Engine) {
	userContext := db.Database("AScore").Collection("users")

	api := router.Group("auth")
	{
		// log in a user
		api.POST("login", func(ctx *gin.Context) {
			services.LoginUser(ctx, userContext)
		})

		// check user is logged in
		api.GET("", func(ctx *gin.Context) {
			services.CheckUser(ctx, userContext)
		})

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
