package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	AuthController(router)
	UserController(router)
}
