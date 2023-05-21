package services

import (
	"strings"

	"github.com/Agile-tools-SaaS/dashboard/helpers"
	auth_helpers "github.com/Agile-tools-SaaS/dashboard/helpers/auth"
	http_response "github.com/Agile-tools-SaaS/dashboard/helpers/http"
	models "github.com/Agile-tools-SaaS/dashboard/models/user"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	db := helpers.NewContext("users")
	defer db.Close()
	user_login := new(models.LoginUser)
	if err := c.ShouldBindJSON(&user_login); err != nil {
		http_response.ServerError(c, err.Error())
		return
	}
	if strings.TrimSpace(user_login.User) == "" && strings.TrimSpace(user_login.Password) == "" {
		http_response.Bad(c, "Fill in all the required fields")
		return
	}
	username, err := user_login.UserLogin(db)
	if err != nil {
		if err.Error() == "The user does not exist" || err.Error() == "Wrong password" {
			http_response.Bad(c, err.Error())
			return
		}
		http_response.Bad(c, err.Error())
		return
	}
	token, err := auth_helpers.GenerateJWT(username)
	if err != nil {
		http_response.Bad(c, err.Error())
		return
	}
	http_response.Ok(c, gin.H{"token": token}, "")
}

func CheckUser(c *gin.Context) {

	isAuthenticated, email := auth_helpers.CheckAuthorized(c)

	if isAuthenticated {
		http_response.Ok(c, gin.H{"is_authenticated": isAuthenticated, "email": email}, "")
		return
	}
	http_response.Forbidden(c, "User not authenticated")
}
