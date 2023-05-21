package models

import "reflect"

type User struct {
	FirstName     string    `bson:"first_name,omitempty"`
	Surname       string    `bson:"surname,omitempty"`
	Email         string    `bson:"email,omitempty"`
	Password      string    `bson:"password,omitempty"`
	Spaces        *[]string `bson:"spaces,omitempty"`
	SpacesAdminOf *[]string `bson:"spaces_admin_of,omitempty"`
	DisplayImage  string    `bson:"display_image,omitempty"`
	CreatedAt     string    `bson:"created_at,omitempty"`
}

func (u *User) CheckUserIsEmpty() bool {
	empty_user := new(User)
	empty_user.Spaces = nil
	return reflect.DeepEqual(u, empty_user)
}

func (u *User) ConvertToReturnUser() *ReturnUser {
	return_user := new(ReturnUser)

	return_user.FirstName = u.FirstName
	return_user.Surname = u.Surname
	return_user.DisplayImage = u.DisplayImage
	return_user.Spaces = u.Spaces
	return_user.SpacesAdminOf = u.SpacesAdminOf
	return_user.Email = u.Email

	return return_user
}
