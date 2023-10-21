package test

import (
	"context"
	"go-clean-architecture/common"
	"go-clean-architecture/common/requester"
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/internal/user/repository"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSendOTP(t *testing.T) {
	envVars := map[string]string{
		"CONFIG_PATH":            "../../config/env",
		"CONFIG_ENV":             "local",
		"CONFIG_ALLOW_MIGRATION": "true",
	}

	for key, value := range envVars {
		_ = os.Setenv(key, value)
	}

	envValue := os.Getenv("CONFIG_ALLOW_MIGRATION")
	allowUpgrade, err := strconv.ParseBool(envValue)
	assert.NoError(t, err)

	appConfig := config.InitLoadAppConf()
	dbInstance := db.InitDatabase(allowUpgrade, *appConfig)

	userRepository := repository.NewUserRepository(
		repository.NewUserFinderImpl(*dbInstance),
		repository.NewUserWriterImpl(*dbInstance),
	)

	usecase := usecase.NewOTPResender(
		userRepository,
	)

	c := gin.Context{}
	user := requester.JWTRequester{
		ID:   "03c333c0-8148-4805-8efb-8a4e6009e824",
		Role: "student",
	}

	c.Set(requester.CurrentRequester, user)
	requester := c.MustGet(requester.CurrentRequester).(requester.Requester)
	ctx := context.WithValue(context.Background(), common.VerifyTokenKey{}, requester)

	t.Run("Valid Email", func(t *testing.T) {
		err = usecase.Execute(ctx)

		assert.NoError(t, err)
	})
}
