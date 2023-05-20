package models

type ChangeAccountDetailsModel struct {
	FirstName    string `bson:"first_name,omitempty"`
	Surname      string `bson:"surname,omitempty"`
	Email        string `bson:"email,omitempty"`
	DisplayImage string `bson:"display_image,omitempty"`
}
