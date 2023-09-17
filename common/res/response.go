package res

import (
	"github.com/gin-gonic/gin"
)

type OK struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	RequestID  string      `json:"request_id"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseOK(
	c *gin.Context,
	respone OK,
) {
	c.JSON(respone.StatusCode, respone)
}
