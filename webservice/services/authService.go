package services

import (
	"strings"

	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	models "github.com/Agile-tools-SaaS/dashboard/models/user"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()
	user_login := new(models.LoginUser)
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if strings.TrimSpace(user_login.User) == "" && strings.TrimSpace(user_login.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}
	username, err := user_login.UserLogin(db)
	if err != nil {
		if err.Error() == "The user does not exist" || err.Error() == "Wrong password" {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := auth_helpers.GenerateJWT(username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func CheckUser(c *gin.Context) {

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if isAuthenticated {
		c.JSON(200, gin.H{
			"isAuthenticated": isAuthenticated,
			"username":        email,
		})
		return
	}
	c.JSON(403, gin.H{
		"isAuthenticated": isAuthenticated,
		"message":         "User not authenticated",
	})

}

func LogoutUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()
}
