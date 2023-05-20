package models

import (
	"errors"

	"github.com/Agile-tools-SaaS/dashboard/helpers"

	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginUser struct {
	User     string `bson:"user,omitempty"`
	Password string `bson:"password,omitempty"`
}

func (u *LoginUser) UserLogin(db *helpers.Context) (string, error) {
	user := new(User)

	db.Collection.FindOne(*db.Context, bson.M{"email": u.User}).Decode(&user)

	if err := auth_helpers.CheckPassword(user.Password, u.Password); err != nil {
		err = errors.New("wrong password")
		return "", err
	}
	return user.Email, nil
}
