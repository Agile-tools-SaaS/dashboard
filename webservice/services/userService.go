package services

import (
	"strings"

	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"

	models "github.com/Agile-tools-SaaS/dashboard/models/user"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, db *helpers.Context) {

	defer db.Close()

	user := new(models.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashed_password, err := auth_helpers.HashPassword(user.Password)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = string(hashed_password)
	user.DisplayImage = ""
	user.Spaces = []string{}

	if strings.TrimSpace(user.FirstName) == "" && strings.TrimSpace(user.Surname) == "" && strings.TrimSpace(user.Email) == "" && strings.TrimSpace(user.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if checkmail.ValidateFormat(user.Email) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email format",
		})
		return
	}
	_, err = db.Collection.InsertOne(c, &user)
	if err != nil {
		if err.Error() == "User already exists" {
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
	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}

func ChangePassword(c *gin.Context, db *helpers.Context) {

}

func ChangeUserDetails(c *gin.Context, db *helpers.Context) {

}

func DeleteUser(c *gin.Context, db *helpers.Context) {

}
