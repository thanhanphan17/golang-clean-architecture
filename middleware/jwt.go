package middleware

import (
	"errors"
	"fmt"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/common/requester"
	"go-clean-architecture/config"
	"go-clean-architecture/provider/tokenprovider/jwt"
	tokentype "go-clean-architecture/provider/tokenprovider/type"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

// extractTokenFromHeader extracts the token from the Authorization header.
// It expects the header value to be in the format "Bearer {token}".
// If the header is not in the correct format, it returns an error.
func extractTokenFromHeader(header string) (*string, error) {
	parts := strings.Fields(header)

	// Check if the header is in the correct format
	if len(parts) < 2 || parts[0] != "Bearer" || parts[1] == "" {
		slog.Error(fmt.Sprintf("invalid or empty Authorization header: %s", header))
		return nil, cerr.ErrWrongAuthHeader(nil)
	}

	token := parts[1]
	return &token, nil
}

// RequireJWT is a middleware that requires a JWT token for authentication.
// It extracts the token from the request header, validates it, and sets the
// requester information in the context.
func RequireAuthorization(cfg *config.AppConfig, tokenType ...string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		tokenProvider := jwt.NewJWTProvider(cfg.JWTSecretKey)

		// Extract token from header
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		// Validate token
		payload, err := tokenProvider.Validate(*token)
		if err != nil {
			panic(err)
		}

		// Set default token type if not provided
		if len(tokenType) == 0 {
			tokenType = []string{tokentype.ACCESS_TOKEN.Value()}
		}

		// Check token type
		if payload.Type != tokenType[0] {
			panic(errors.New("invalid token type"))
		}

		// Set requester information in context
		c.Set(requester.CurrentRequester, requester.JWTRequester{
			ID:   payload.UserID,
			Role: payload.Role,
		})

		c.Next()
	}
}
