package services

import (
	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	http_response "github.com/Agile-tools-SaaS/dashboard/helpers/http"
	space_models "github.com/Agile-tools-SaaS/dashboard/models/space"
	user_models "github.com/Agile-tools-SaaS/dashboard/models/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserToSpace(c *gin.Context) {
	space_db := helpers.NewContext("spaces")
	defer space_db.Close()

	user_db := helpers.NewContext("users")
	defer user_db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	space_id_unparsed := c.Param("space_id")

	space_id, err := primitive.ObjectIDFromHex(space_id_unparsed)

	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id_unparsed+"} does not exist")
		return
	}
	user_to_be_added_email := c.Param("user_id")

	var space space_models.Space

	space_db.Collection.FindOne(*space_db.Context, bson.M{"_id": space_id}).Decode(&space)

	if helpers.Contains_string(*space.Users, user_to_be_added_email) {
		http_response.Bad(c, "User already added")
		return
	}

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	if !helpers.Contains_string(*space.AdminUsers, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	// Add user to space
	updated_user_list := append(*space.Users, user_to_be_added_email)
	result, err := space_db.Collection.UpdateByID(*space_db.Context, space_id, bson.M{"$set": bson.M{"users": updated_user_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	// Add space to user
	user_to_be_added := new(user_models.User)
	user_db.Collection.FindOne(*user_db.Context, bson.M{"email": user_to_be_added_email}).Decode(&user_to_be_added)

	if user_to_be_added.CheckUserIsEmpty() {
		http_response.Bad(c, "User not found")
		return
	}
	updated_spaces_list := append(*user_to_be_added.Spaces, space_id_unparsed)
	result, err = user_db.Collection.UpdateOne(*user_db.Context, bson.M{"email": user_to_be_added_email}, bson.M{"$set": bson.M{"spaces": updated_spaces_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	http_response.Ok_no_body(c, "User added successfully")

}
func MakeUserAnAdminOfSpace(c *gin.Context) {
	space_db := helpers.NewContext("spaces")
	defer space_db.Close()

	user_db := helpers.NewContext("users")
	defer user_db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	space_id_unparsed := c.Param("space_id")

	space_id, err := primitive.ObjectIDFromHex(space_id_unparsed)

	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id_unparsed+"} does not exist")
		return
	}
	user_to_be_added_email := c.Param("user_id")

	var space space_models.Space

	space_db.Collection.FindOne(*space_db.Context, bson.M{"_id": space_id}).Decode(&space)

	if !helpers.Contains_string(*space.Users, user_to_be_added_email) {
		http_response.Bad(c, "User does not exist in space")
		return
	}

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	if !helpers.Contains_string(*space.AdminUsers, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	// Add user to space
	updated_user_list := append(*space.AdminUsers, user_to_be_added_email)
	result, err := space_db.Collection.UpdateByID(*space_db.Context, space_id, bson.M{"$set": bson.M{"admin_users": updated_user_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	// Add space to user
	user_to_be_added := new(user_models.User)
	user_db.Collection.FindOne(*user_db.Context, bson.M{"email": user_to_be_added_email}).Decode(&user_to_be_added)

	if user_to_be_added.CheckUserIsEmpty() {
		http_response.Bad(c, "User not found")
		return
	}
	updated_spaces_list := append(*user_to_be_added.SpacesAdminOf, space_id_unparsed)
	result, err = user_db.Collection.UpdateOne(*user_db.Context, bson.M{"email": user_to_be_added_email}, bson.M{"$set": bson.M{"spaces_admin_of": updated_spaces_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	http_response.Ok_no_body(c, "User added successfully")

}
func RemoveUserFromSpace(c *gin.Context) {
	space_db := helpers.NewContext("spaces")
	defer space_db.Close()

	user_db := helpers.NewContext("users")
	defer user_db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	space_id_unparsed := c.Param("space_id")

	space_id, err := primitive.ObjectIDFromHex(space_id_unparsed)

	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id_unparsed+"} does not exist")
		return
	}
	user_to_be_added_email := c.Param("user_id")

	var space space_models.Space

	space_db.Collection.FindOne(*space_db.Context, bson.M{"_id": space_id}).Decode(&space)

	if !helpers.Contains_string(*space.AdminUsers, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	if !helpers.Contains_string(*space.Users, user_to_be_added_email) {
		http_response.Bad(c, "User not in space")
		return
	}

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	// DELETE user from space

	var updated_user_list []string

	for _, user := range *space.Users {
		if user != user_to_be_added_email {
			updated_user_list = append(updated_user_list, user)
		}
	}

	var updated_admin_user_list []string

	for _, admin_user := range *space.AdminUsers {
		if admin_user != user_to_be_added_email {
			updated_admin_user_list = append(updated_admin_user_list, admin_user)
		}
	}

	result, err := space_db.Collection.UpdateByID(*space_db.Context, space_id, bson.M{"$set": bson.M{"users": updated_user_list, "admin_users": updated_admin_user_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	// DELETE space from user

	var user_to_be_removed_from_space user_models.User

	user_db.Collection.FindOne(*user_db.Context, bson.M{"email": user_to_be_added_email}).Decode(&user_to_be_removed_from_space)

	var user_spaces_list []string

	for _, user := range *user_to_be_removed_from_space.Spaces {
		if user != user_to_be_added_email {
			user_spaces_list = append(user_spaces_list, user)
		}
	}

	var user_admin_spaces_list []string

	for _, admin_user := range *user_to_be_removed_from_space.SpacesAdminOf {
		if admin_user != user_to_be_added_email {
			user_admin_spaces_list = append(user_admin_spaces_list, admin_user)
		}
	}

	result, err = user_db.Collection.UpdateOne(*space_db.Context, bson.M{"email": user_to_be_added_email}, bson.M{"$set": bson.M{"spaces": user_spaces_list, "spaces_admin_of": user_admin_spaces_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	http_response.Ok_no_body(c, "Successfully removed user from space")
}

func RemoveUserFromAdminOfSpace(c *gin.Context) {
	space_db := helpers.NewContext("spaces")
	defer space_db.Close()

	user_db := helpers.NewContext("users")
	defer user_db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	space_id_unparsed := c.Param("space_id")

	space_id, err := primitive.ObjectIDFromHex(space_id_unparsed)

	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id_unparsed+"} does not exist")
		return
	}
	user_to_be_added_email := c.Param("user_id")

	var space space_models.Space

	space_db.Collection.FindOne(*space_db.Context, bson.M{"_id": space_id}).Decode(&space)

	if !helpers.Contains_string(*space.AdminUsers, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	if !helpers.Contains_string(*space.Users, user_to_be_added_email) {
		http_response.Bad(c, "User not in space")
		return
	}

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	// DELETE user from space

	var updated_admin_user_list []string

	for _, admin_user := range *space.AdminUsers {
		if admin_user != user_to_be_added_email {
			updated_admin_user_list = append(updated_admin_user_list, admin_user)
		}
	}

	result, err := space_db.Collection.UpdateByID(*space_db.Context, space_id, bson.M{"$set": bson.M{"admin_users": updated_admin_user_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	// DELETE space from user

	var user_to_be_removed_from_space user_models.User

	user_db.Collection.FindOne(*user_db.Context, bson.M{"email": user_to_be_added_email}).Decode(&user_to_be_removed_from_space)

	var user_admin_spaces_list []string

	for _, admin_user := range *user_to_be_removed_from_space.SpacesAdminOf {
		if admin_user != user_to_be_added_email {
			user_admin_spaces_list = append(user_admin_spaces_list, admin_user)
		}
	}

	result, err = user_db.Collection.UpdateOne(*space_db.Context, bson.M{"email": user_to_be_added_email}, bson.M{"$set": bson.M{"spaces_admin_of": user_admin_spaces_list}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "User not admin")
		return
	}

	http_response.Ok_no_body(c, "Successfully removed admin priveledges from user")
}
