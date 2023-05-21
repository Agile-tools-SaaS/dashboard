package services

import (
	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	http_response "github.com/Agile-tools-SaaS/dashboard/helpers/http"
	space_models "github.com/Agile-tools-SaaS/dashboard/models/space"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddFileToSpace(c *gin.Context) {
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

	if !helpers.Contains_string(*space.Users, email) || !isAuthenticated {
		http_response.Forbidden(c, "You are not allowed to make changes in this space")
		return
	}

	new_space_file := new(space_models.NewSpaceFileModel)

	c.BindJSON(&new_space_file)

	space_file := space_models.SpaceFile{
		Id:      primitive.NewObjectID(),
		Name:    new_space_file.Name,
		Content: new_space_file.Content,
	}

	new_space_files := append(*space.Files, space_file)

	space.Files = &new_space_files

	result, err := db.Collection.UpdateByID(*db.Context, id, bson.M{"$set": bson.M{"files": space.Files}})

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Could not find space")
		return
	}

	http_response.Ok_no_body(c, "Successfully added file to space")
}
func GetSpaceFileByFileNameAndSpaceName(c *gin.Context) {
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

	if space.CheckSpaceIsEmpty() {
		http_response.Bad(c, "Space not found")
		return
	}

	file_id_unsanitised := c.Param("file_id")

	println(file_id_unsanitised)
	file_id, err := primitive.ObjectIDFromHex(file_id_unsanitised)
	if err != nil {
		http_response.ServerError(c, "File does not exist")
		return
	}

	if space.Contains_File_By_Id(file_id) {
		file, err := space.Get_File_By_Id(file_id)
		if err != nil {
			http_response.Bad(c, "No file found with this id")
			return
		}
		http_response.Ok(c, file, "")
		return
	} else {
		http_response.Bad(c, "File does not exist in space")
	}
}
func EditFileInSpace(c *gin.Context) {
	db := helpers.NewContext("spaces")
	defer db.Close()

	isAuthenticated, _ := auth_helpers.CheckAuthorized(c)

	if !isAuthenticated {
		http_response.Forbidden(c, "You must be authenticated to remove this file")
	}

	space_id_unparsed := c.Param("space_id")

	space_id, err := primitive.ObjectIDFromHex(space_id_unparsed)

	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id_unparsed+"} does not exist")
		return
	}

	file_id_unsanitised := c.Param("file_id")

	file_id, err := primitive.ObjectIDFromHex(file_id_unsanitised)

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	var space space_models.Space

	db.Collection.FindOne(*db.Context, bson.M{"_id": space_id}).Decode(&space)

	var files []space_models.SpaceFile = *space.Files

	if len(files) == 0 {
		http_response.Bad(c, "No files in space")
		return
	}

	var file_to_be_updated space_models.SpaceFile

	var other_files []space_models.SpaceFile

	found := false
	for _, f := range files {
		if f.Id == file_id {
			file_to_be_updated = f
			found = true
		} else {
			other_files = append(other_files, f)
		}
	}

	if !found {
		http_response.Bad(c, "No file found")
		return
	}

	var new_file_data space_models.SpaceFile

	c.BindJSON(&new_file_data)

	file_to_be_updated.Content = new_file_data.Content
	file_to_be_updated.Name = new_file_data.Name

	other_files = append(other_files, file_to_be_updated)

	result, err := db.Collection.UpdateByID(*db.Context, space_id, bson.M{"$set": bson.M{"files": other_files}})

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	http_response.Ok_no_body(c, "Successfully updated file")

}
func DeleteFileFromSpace(c *gin.Context) {
	db := helpers.NewContext("spaces")
	defer db.Close()

	isAuthenticated, _ := auth_helpers.CheckAuthorized(c)

	if !isAuthenticated {
		http_response.Forbidden(c, "You must be authenticated to remove this file")
	}

	space_id_unparsed := c.Param("space_id")

	space_id, err := primitive.ObjectIDFromHex(space_id_unparsed)

	if err != nil {
		http_response.ServerError(c, "Space with id: {"+space_id_unparsed+"} does not exist")
		return
	}

	file_id_unsanitised := c.Param("file_id")

	file_id, err := primitive.ObjectIDFromHex(file_id_unsanitised)

	if err != nil {
		http_response.ServerError(c, err.Error())
		return
	}

	var space space_models.Space

	db.Collection.FindOne(*db.Context, bson.M{"_id": space_id}).Decode(&space)

	var files []space_models.SpaceFile = *space.Files

	if len(files) == 0 {
		http_response.Bad(c, "No files in space")
		return
	}

	file_length_before := len(files)

	var files_new []space_models.SpaceFile

	for _, f := range files {
		if f.Id != file_id {
			files_new = append(files_new, f)
		}
	}
	file_length_after := len(files_new)

	if file_length_after == file_length_before {
		http_response.Bad(c, "File does not exist")
		return
	}

	var result *mongo.UpdateResult

	if len(files_new) == 0 {
		result, err = db.Collection.UpdateByID(*db.Context, space_id, bson.M{"$set": bson.M{"files": &[]string{}}})
	} else {
		result, err = db.Collection.UpdateByID(*db.Context, space_id, bson.M{"$set": bson.M{"files": files_new}})
	}

	if err != nil {
		http_response.ServerError(c, err.Error())
	}

	if result.MatchedCount == 0 {
		http_response.Bad(c, "Couldnt find space")
		return
	}

	http_response.Ok_no_body(c, "Successfully deleted file")
}
