package services

import (
	"strings"

	helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"

	models "github.com/Agile-tools-SaaS/dashboard/models/user"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *gin.Context, db *mongo.Collection) {

	// HASHING PASSWORD NOT ADDED

	user := new(models.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashed_password, err := helpers.HashPassword(user.Password)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = string(hashed_password)
	user.DisplayImage = ""
	user.Spaces = []string{}

	if strings.TrimSpace(user.FirstName) == "" && strings.TrimSpace(user.Surname) == "" && strings.TrimSpace(user.Username) == "" && strings.TrimSpace(user.Email) == "" && strings.TrimSpace(user.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if checkmail.ValidateFormat(user.Email) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email format",
		})
		return
	}
	_, err = db.InsertOne(c, &user)
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

func ChangePassword(c *gin.Context, db *mongo.Collection) {

}

func ChangeUserDetails(c *gin.Context, db *mongo.Collection) {

}

func DeleteUser(c *gin.Context, db *mongo.Collection) {

}
