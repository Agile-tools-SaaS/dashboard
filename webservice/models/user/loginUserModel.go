package models

type LoginUser struct {
	User     string `bson:"user,omitempty"`
	Password string `bson:"password,omitempty"`
}
