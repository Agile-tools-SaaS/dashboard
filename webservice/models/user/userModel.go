package models

type User struct {
	FirstName    string   `bson:"first_name,omitempty"`
	Surname      string   `bson:"surname,omitempty"`
	Email        string   `bson:"email,omitempty"`
	Username     string   `bson:"username,omitempty"`
	Password     string   `bson:"password,omitempty"`
	Spaces       []string `bson:"spaces,omitempty"`
	DisplayImage string   `bson:"display_image,omitempty"`
	CreatedAt    string   `bson:"created_at,omitempty"`
}
