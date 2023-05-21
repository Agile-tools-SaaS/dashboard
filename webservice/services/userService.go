package services

import (
	"strings"
	"time"

	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	http_response "github.com/Agile-tools-SaaS/dashboard/helpers/http"
	space_models "github.com/Agile-tools-SaaS/dashboard/models/space"
	models "github.com/Agile-tools-SaaS/dashboard/models/user"
	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *gin.Context) {

	db := helpers.NewContext("users")
	defer db.Close()
	user := new(models.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	hashed_password, err := auth_helpers.HashPassword(user.Password)

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	user.Password = string(hashed_password)
	user.DisplayImage = ""
	user.Spaces = &([]string{})
	user.SpacesAdminOf = &([]string{})
	user.CreatedAt = time.Now().String()

	var find_user models.User

	db.Collection.FindOne(*db.Context, bson.M{"email": user.Email}).Decode(&find_user)

	if !find_user.CheckUserIsEmpty() {
		http_response.Bad(c, "User Already exists")
		return
	}

	if strings.TrimSpace(user.FirstName) == "" || strings.TrimSpace(user.Surname) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		http_response.Bad(c, "Fill in all the require fields")
		return
	}

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if checkmail.ValidateFormat(user.Email) != nil {
		http_response.Bad(c, "Invalid email format")
		return
	}

	_, err = db.Collection.InsertOne(c, &user)
	if err != nil {
		if err.Error() == "User already exists" {
			http_response.Bad(c, "User already exists")
			return
		}
		http_response.ServerError(c, err.Error())
		return
	}
	http_response.Ok_no_body(c, "User registered successfully")
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
			http_response.Forbidden(c, "incorrect password, cannot change password")
			return
		}
		hashed_password, err := auth_helpers.HashPassword(new_password_model.NewPassword)

		if err != nil {
			http_response.ServerError(c, err.Error())
			return
		}

		result, err := db.Collection.UpdateOne(*db.Context, bson.M{"email": user_email}, bson.M{"$set": bson.M{"password": hashed_password}})
		if err != nil {
			http_response.ServerError(c, err.Error())
			return
		}

		if result.MatchedCount == 0 {
			http_response.Bad(c, "Could not find user with email: "+email)
			return
		}

		http_response.Ok_no_body(c, "Successfully changed password")
		return
	}
	http_response.Forbidden(c, "Cannot change the password as you are not authenticated")
}

func ChangeUserDetails(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user_email := c.Param("user")

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if isAuthenticated && user_email == email {

		new_change_details_user_model := new(models.ChangeAccountDetailsModel)
		c.BindJSON(&new_change_details_user_model)

		find_user := new(models.User)

		db.Collection.FindOne(*db.Context, bson.M{"email": new_change_details_user_model.Email}).Decode(&find_user)

		if !find_user.CheckUserIsEmpty() && find_user.Email != user_email {
			http_response.Bad(c, "User Already Exists")
			return
		}

		new_change_details_user_model.Email = strings.TrimSpace(strings.ToLower(new_change_details_user_model.Email))

		if checkmail.ValidateFormat(new_change_details_user_model.Email) != nil {
			http_response.Bad(c, "Invalid email format")
			return
		}

		result, err := db.Collection.UpdateOne(*db.Context, bson.M{"email": user_email}, bson.M{"$set": new_change_details_user_model})
		if err != nil {
			http_response.ServerError(c, err.Error())
			return
		}

		if result.MatchedCount == 0 {
			http_response.Bad(c, "Could not find user with email: "+email)
			return
		}

		http_response.Ok_no_body(c, "Successfully changed account details")
		return
	}
	http_response.Forbidden(c, "Cannot change account details as you are not authenticated")
}

func DeleteUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user := c.Param("user")

	isLoggedIn, username := auth_helpers.CheckAuthorized(c)

	if isLoggedIn && username == user {

		db.Collection.FindOneAndDelete(*db.Context, bson.M{"email": user})

		http_response.Ok_no_body(c, "Successfully deleted the user")
		return
	}
	http_response.Forbidden(c, "you are not authorized to delete this user")
}

func FindOneUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()

	user := c.Param("user")

	find_user := new(models.User)

	db.Collection.FindOne(*db.Context, bson.M{"email": user}).Decode(&find_user)

	if find_user.CheckUserIsEmpty() {
		http_response.Bad(c, "User not found")
		return
	}

	http_response.Ok(c, find_user.ConvertToReturnUser(), "")
}

func GetSpacesByUserAndFilterWithPagination(c *gin.Context) {
	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if !isAuthenticated {
		http_response.Forbidden(c, "You must be logged in to view this space")
		return
	}
	users_db := helpers.NewContext("users")
	defer users_db.Close()

	user := c.Param("user")

	find_user := new(models.User)

	users_db.Collection.FindOne(*users_db.Context, bson.M{"email": user}).Decode(&find_user)

	if find_user.CheckUserIsEmpty() {
		http_response.Bad(c, "User not found")
		return
	}

	spaces_db := helpers.NewContext("spaces")
	defer spaces_db.Close()

	users_spaces := find_user.Spaces

	spaces := []space_models.Space{}

	for _, space := range *users_spaces {
		id, err := primitive.ObjectIDFromHex(space)

		if err != nil {
			http_response.ServerError(c, "Space does not exist")
			return
		}

		space_model := new(space_models.Space)

		spaces_db.Collection.FindOne(*spaces_db.Context, bson.M{"_id": id}).Decode(&space_model)

		if space_model.CheckSpaceIsEmpty() {
			http_response.Bad(c, "Space not found")
			return
		}

		if !helpers.Contains_string(*space_model.Users, email) {
			http_response.Forbidden(c, "You are not allowed in this space")
			return
		}

		spaces = append(spaces, *space_model)
	}

	http_response.Ok(c, spaces, "")
}
