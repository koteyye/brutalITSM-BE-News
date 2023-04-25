package rest

import "github.com/gin-gonic/gin"

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, error{message})
}
