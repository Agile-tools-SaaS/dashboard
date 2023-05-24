package services

import (
	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	http_response "github.com/Agile-tools-SaaS/dashboard/helpers/http"
	space_models "github.com/Agile-tools-SaaS/dashboard/models/space"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSpace(c *gin.Context) {
	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if !isAuthenticated {
		http_response.Forbidden(c, "You must be logged in to create a space")
		return
	}

	db := helpers.NewContext("spaces")
	defer db.Close()

	space := new(space_models.Space)

	space.Files = &[]space_models.SpaceFile{}
	space.Id = primitive.NewObjectID()
	space.AdminUsers = &[]string{email}
	space.Templates = &[]string{}
	space.Users = &[]string{email}

	if err := c.ShouldBindJSON(&space); err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	if space.Name == "" || space.AdminUsers == nil || space.Users == nil {
		http_response.Bad(c, "Fill in all the required fields")
		return
	}

	_, err := db.Collection.InsertOne(c, &space)
	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	http_response.Ok(c, space, "Successfully created space")
}

func GetSpaceById(c *gin.Context) {
	db := helpers.NewContext("spaces")
	defer db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if !isAuthenticated {
		http_response.Forbidden(c, "You must be logged in to view this space")
		return
	}

	space_id := c.Param("space_id")

	id, err := primitive.ObjectIDFromHex(space_id)
	if err != nil {
		http_response.ServerError(c, "Space does not exist")
		return
	}

	space := new(space_models.Space)

	db.Collection.FindOne(*db.Context, bson.M{"_id": id}).Decode(&space)

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	if !helpers.Contains_string(*space.Users, email) {
		http_response.Forbidden(c, "You are not allowed in this space")
		return
	}

	http_response.Ok(c, space, "")
}

func ChangeSpaceDetails(c *gin.Context) {
	db := helpers.NewContext("spaces")
	defer db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	space_id := c.Param("space_id")

	id, err := primitive.ObjectIDFromHex(space_id)
	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id+"} does not exist")
		return
	}

	space := new(space_models.Space)

	db.Collection.FindOne(*db.Context, bson.M{"_id": id}).Decode(&space)

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	if !helpers.Contains_string(*space.AdminUsers, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	changed_space := new(space_models.ChangeSpaceDetailsModel)

	c.BindJSON(&changed_space)

	result, err := db.Collection.UpdateByID(*db.Context, id, bson.M{"$set": bson.M{"name": changed_space.Name}})

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Could not find space")
		return
	}

	http_response.Ok_no_body(c, "Successfully changed space details")
}

func DeleteSpace(c *gin.Context) {
	db := helpers.NewContext("spaces")
	defer db.Close()

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	space_id := c.Param("space_id")

	id, err := primitive.ObjectIDFromHex(space_id)
	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id+"} does not exist")
		return
	}

	space := new(space_models.Space)

	db.Collection.FindOne(*db.Context, bson.M{"_id": id}).Decode(&space)

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	if !helpers.Contains_string(*space.AdminUsers, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	db.Collection.FindOneAndDelete(*db.Context, bson.M{"_id": id})
	http_response.Ok_no_body(c, "Successfully deleted space")
}
