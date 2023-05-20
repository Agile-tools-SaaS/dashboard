package models

type ReturnUser struct {
	FirstName    string   `bson:"first_name,omitempty"`
	Surname      string   `bson:"surname,omitempty"`
	Username     string   `bson:"username,omitempty"`
	DisplayImage string   `bson:"display_image,omitempty"`
	Spaces       []string `bson:"spaces,omitempty"`
}
