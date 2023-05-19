package controllers

import (
	"webservice/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthController(db *mongo.Client, router *gin.Engine) {

	userContext := db.Database("cluster0").Collection("users")

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
	}
}
