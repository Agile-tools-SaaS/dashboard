package controllers

import (
	"webservice/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserController(db *mongo.Client, router *gin.Engine) {
	userContext := db.Database("cluster0").Collection("users")

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
