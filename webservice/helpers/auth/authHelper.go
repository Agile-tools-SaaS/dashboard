package helpers

import (
	"github.com/gin-gonic/gin"
)

func CheckAuthorized(c *gin.Context) (bool, string) {
	token := c.Request.Header["Authorization"][0]

	var username string

	err := CheckJWT(token, &username)
	if err != nil {
		return false, ""
	}
	return true, username
}
