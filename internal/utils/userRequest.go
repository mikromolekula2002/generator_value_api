package utils

import "github.com/gin-gonic/gin"

func RetrieveUserRequest(c *gin.Context) (method, userAgent, url string) {
	method = c.Request.Method
	userAgent = c.Request.UserAgent()
	url = c.Request.URL.String()

	return method, userAgent, url
}
