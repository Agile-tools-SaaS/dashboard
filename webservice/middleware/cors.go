package middleware

import (
	"strings"

	"github.com/Agile-tools-SaaS/dashboard/helpers"

	"github.com/gin-gonic/gin"
)

func getAllowedOriginsList() map[string]bool {
	var allowed_origins string = helpers.GetEnvByName("ALLOWED_ORIGIN")
	allowList := make(map[string]bool)

	for _, v := range strings.Split(allowed_origins, ",") {
		allowList[v] = true
	}

	return allowList
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowList := getAllowedOriginsList()
		if origin := c.Request.Header.Get("Origin"); allowList[origin] || allowList["*"] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Next()
		} else {
			c.AbortWithStatus(403)
			return
		}
	}

}
