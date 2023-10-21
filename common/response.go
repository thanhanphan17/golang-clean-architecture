package common

import (
	"github.com/gin-gonic/gin"
)

type SuccessRes struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, response SuccessRes) {
	c.JSON(response.StatusCode, response)
}
