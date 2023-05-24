package models

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Space struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string             `bson:"name"`
	Files      *[]SpaceFile       `bson:"files"`
	Users      *[]string          `bson:"users"`
	AdminUsers *[]string          `bson:"admin_users" json:"admin_users"`
	Templates  *[]string          `bson:"templates" json:"templates"`
}

func (s *Space) CheckSpaceIsEmpty() bool {
	empty_user := new(Space)
	empty_user.Users = nil
	empty_user.AdminUsers = nil
	empty_user.Files = nil
	empty_user.Templates = nil
	return reflect.DeepEqual(s, empty_user)
}

func (s *Space) Contains_File_By_Id(f primitive.ObjectID) bool {
	for _, a := range *s.Files {
		if a.Id == f {
			return true
		}
	}
	return false
}

func (s *Space) Get_File_By_Id(f primitive.ObjectID) (SpaceFile, error) {
	for _, a := range *s.Files {
		if a.Id == f {
			return a, nil
		}
	}

	empty_space_file := new(SpaceFile)

	return *empty_space_file, errors.New("File does not exist")
}
