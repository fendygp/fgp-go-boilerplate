package utils

import "github.com/gin-gonic/gin"

// ErrorMessage ...
func ErrorMessage(c *gin.Context, status int, code, msg string) *gin.Context {
	c.JSON(status, gin.H{
		"Code":    code,
		"Message": msg,
	})
	return c
}
