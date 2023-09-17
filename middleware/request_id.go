package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	xRequestIDKey = "X-Request-ID"
)

// RequestID is a middleware that injects a 'RequestID'
// into the context and header of each request.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		xRequestID := uuid.New().String()

		c.Request.Header.Set(xRequestIDKey, xRequestID)
		c.Set(xRequestIDKey, xRequestID)

		fmt.Printf("[GIN] RequestID: [%s]\n", xRequestID)

		c.Next()
	}
}

// GetRequestIDFromContext returns 'RequestID' from
// the given context if present.
func GetRequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(xRequestIDKey); ok {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}
	return ""
}

// GetRequestIDFromHeaders returns 'RequestID' from
// the headers if present.
func GetRequestIDFromHeaders(c *gin.Context) string {
	return c.Request.Header.Get(xRequestIDKey)
}
