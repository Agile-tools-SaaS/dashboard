package helpers

import "github.com/gin-gonic/gin"

func response[T any](status int, body T, message string) (int, gin.H) {
	response_body := gin.H{
		"status":  status,
		"body":    body,
		"message": message,
	}

	return status, response_body
}

type EmptyBody struct{}

func Ok(c *gin.Context, body interface{}, message string) {
	c.JSON(response(200, body, message))
}

func Ok_no_body(c *gin.Context, message string) {
	empty_body := EmptyBody{}
	c.JSON(response(200, empty_body, message))
}

func Forbidden(c *gin.Context, message string) {
	empty_body := EmptyBody{}
	c.JSON(response(403, empty_body, message))
}

func ServerError(c *gin.Context, message string) {
	empty_body := EmptyBody{}
	c.JSON(response(500, empty_body, message))
}

func NotFound(c *gin.Context, message string) {
	empty_body := EmptyBody{}
	c.JSON(response(404, empty_body, message))
}

func Bad(c *gin.Context, message string) {
	empty_body := EmptyBody{}
	c.JSON(response(400, empty_body, message))
}

func Response(c *gin.Context, status int, body interface{}, message string) {
	c.JSON(response(status, body, message))
}
