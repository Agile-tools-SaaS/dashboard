package models

type ReturnUser struct {
	FirstName    string    `bson:"first_name,omitempty"`
	Surname      string    `bson:"surname,omitempty"`
	Email        string    `bson:"email,omitempty"`
	DisplayImage string    `bson:"display_image,omitempty"`
	Spaces       *[]string `bson:"spaces,omitempty"`
}
