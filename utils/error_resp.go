package utils

import "github.com/gin-gonic/gin"

type APIError struct {
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
}

// ErrorMessage ...
func ErrorMessage(c *gin.Context, status int, apiErr *APIError) *gin.Context {
	c.JSON(status, apiErr)
	return c
}
