package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(db *mongo.Client, router *gin.Engine) {
	AuthController(db, router)
}
