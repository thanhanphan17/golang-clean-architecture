package middleware

import (
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/common/requester"
	"go-clean-architecture/config"
	"go-clean-architecture/provider/tokenprovider/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeader(s string) (*string, error) {
	parts := strings.Split(s, " ")

	//Authorization : Bearer {token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return nil, cerr.ErrWrongAuthHeader(nil)
	}

	return &parts[1], nil
}

// middleware require jwt
func RequireJWT(cfg *config.AppConfig) func(ctx *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(cfg.JWTSecretKey)

	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(*token)

		if err != nil {
			panic(err)
		}

		requesterData := requester.JWTRequester{
			ID:   payload.UserID,
			Role: payload.Role,
		}

		if err != nil {
			panic(err)
		}

		c.Set(requester.CurrentRequester, requesterData)
		c.Next()
	}
}
