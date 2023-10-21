package test

import (
	"context"
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/entity"
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/internal/user/repository"
	"go-clean-architecture/provider/hashprovider"
	"go-clean-architecture/provider/tokenprovider/jwt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Login user integration test
// Make sure your database is ready for connection before testing
func TestLoginUser(t *testing.T) {
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
	tokenProvider := jwt.NewJWTProvider(appConfig.JWTSecretKey)

	md5Hash := hashprovider.NewMd5Hash()
	var accessExpiry uint = 60 * 60 * 24
	var verifyExpiry uint = 60 * 60 * 24

	userRepository := repository.NewUserRepository(
		repository.NewUserFinderImpl(*dbInstance),
		repository.NewUserWriterImpl(*dbInstance),
	)

	loginUserUseCase := usecase.NewUserLoginer(
		userRepository,
		tokenProvider,
		accessExpiry,
		md5Hash,
	)

	createUserUseCase := usecase.NewUserCreator(
		userRepository,
		tokenProvider,
		verifyExpiry,
		md5Hash,
	)

	t.Run("Valid Info", func(t *testing.T) {
		userEntity := entity.User{
			Name:     "Example Name",
			Email:    "example@example.com",
			Phone:    "0123456789",
			Status:   entity.ACTIVE.Value(),
			Role:     "user",
			Password: "1234000@ABC",
		}

		_, _ = createUserUseCase.Execute(context.Background(), userEntity)

		user := entity.User{
			Email:    "example@example.com",
			Password: "1234000@ABC",
		}

		_, err = loginUserUseCase.Execute(context.Background(), user)

		assert.NoError(t, err)
	})

	t.Run("Email Not Existed", func(t *testing.T) {
		user := entity.User{
			Email:    "notexistedemail@howell.com",
			Password: "1234000@ABC",
		}

		_, err = loginUserUseCase.Execute(context.Background(), user)

		assert.Error(t, err)
	})

	t.Run("Wrong Password", func(t *testing.T) {
		user := entity.User{
			Email:    "donniegerhold@howell.com",
			Password: "1234000@AC",
		}

		_, err = loginUserUseCase.Execute(context.Background(), user)

		assert.Error(t, err)
	})
}
