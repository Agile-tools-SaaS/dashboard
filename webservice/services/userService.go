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

func ChangePassword(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user_email := c.Param("user")

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if isAuthenticated && user_email == email {

		new_password_model := new(models.ChangePasswordModel)
		c.BindJSON(&new_password_model)

		user := new(models.User)

		db.Collection.FindOne(*db.Context, bson.M{"email": user_email}).Decode(&user)

		if err := auth_helpers.CheckPassword(user.Password, new_password_model.OldPassword); err != nil {
			c.JSON(403, gin.H{
				"message": "incorrect password, cannot change password",
			})
			return
		}
		hashed_password, err := auth_helpers.HashPassword(new_password_model.NewPassword)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := db.Collection.UpdateOne(*db.Context, bson.M{"email": user_email}, bson.M{"$set": bson.M{"password": hashed_password}})
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(400, gin.H{
				"error": "Could not find user with email: " + email,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Successfully changed password",
		})
		return
	}
	c.JSON(403, gin.H{
		"message": "Cannot change the password as you are not authenticated",
	})
}

func ChangeUserDetails(c *gin.Context) {}

func DeleteUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user := c.Param("user")

	isLoggedIn, username := auth_helpers.CheckAuthorized(c)

	if isLoggedIn && username == user {

		db.Collection.FindOneAndDelete(*db.Context, bson.M{"email": user})

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
		"user": find_user.ConvertToReturnUser(),
	})
}
