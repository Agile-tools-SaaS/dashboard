package services

import (
	"strings"
	"time"

	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	"go.mongodb.org/mongo-driver/bson"

	models "github.com/Agile-tools-SaaS/dashboard/models/user"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	db := helpers.NewContext("users")

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
	user.Spaces = &([]string{})
	user.CreatedAt = time.Now().String()

	var find_user models.User

	db.Collection.FindOne(*db.Context, bson.M{"email": user.Email}).Decode(&find_user)

	if !find_user.CheckUserIsEmpty() {
		c.JSON(400, gin.H{
			"message": "User Already Exists",
		})
		return
	}

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

func ChangePassword(c *gin.Context) {}

func ChangeUserDetails(c *gin.Context) {}

func DeleteUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user := c.Param("user")

	isLoggedIn, username := auth_helpers.CheckAuthorized(c)

	if isLoggedIn && username == user {

		result := db.Collection.FindOneAndDelete(*db.Context, bson.M{"email": user})

		println(result)

		c.JSON(200, gin.H{
			"message": "successfully deleted the user",
		})
		return
	}

	c.JSON(403, gin.H{
		"message": "you are not authorized to delete this user",
	})

}

func FindOneUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user := c.Param("user")

	find_user := new(models.User)

	db.Collection.FindOne(*db.Context, bson.M{"email": user}).Decode(&find_user)

	if find_user.CheckUserIsEmpty() {
		c.JSON(400, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": find_user,
	})
}
