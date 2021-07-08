package utils

import "github.com/gin-gonic/gin"

// SuccessMessage ...
func SuccessMessage(c *gin.Context, status int, msg string) *gin.Context {
	c.JSON(status, gin.H{
		"Code":    "0000",
		"Message": msg,
	})
	return c
}

// SuccessData ...
func SuccessData(c *gin.Context, status int, data interface{}) *gin.Context {
	c.JSON(status, gin.H{
		"responseCode":        "0000",
		"responseDescription": "success",
		"responseData":        data,
	})

	return c
}
