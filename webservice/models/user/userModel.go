package models

import "reflect"

type User struct {
	FirstName    string    `bson:"first_name,omitempty"`
	Surname      string    `bson:"surname,omitempty"`
	Email        string    `bson:"email,omitempty"`
	Password     string    `bson:"password,omitempty"`
	Spaces       *[]string `bson:"spaces,omitempty"`
	DisplayImage string    `bson:"display_image,omitempty"`
	CreatedAt    string    `bson:"created_at,omitempty"`
}

func (u *User) CheckUserIsEmpty() bool {
	empty_user := new(User)
	empty_user.Spaces = nil
	return reflect.DeepEqual(u, empty_user)
}
