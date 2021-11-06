package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type successResponse struct {
	Result string `json:"result"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message, Error: fmt.Sprint(err)})
}

func NewSuccessResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, successResponse{Result: message})
}
