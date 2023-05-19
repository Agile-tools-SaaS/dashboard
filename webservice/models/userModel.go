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

type LoginUser struct {
	User     string `bson:"user,omitempty"`
	Password string `bson:"password,omitempty"`
}

type ReturnUser struct {
	FirstName    string   `bson:"first_name,omitempty"`
	Surname      string   `bson:"surname,omitempty"`
	Username     string   `bson:"username,omitempty"`
	DisplayImage string   `bson:"display_image,omitempty"`
	Spaces       []string `bson:"spaces,omitempty"`
}
