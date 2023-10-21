package test

import (
	"context"
	"go-clean-architecture/common"
	"go-clean-architecture/common/requester"
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

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestVerifyUser(t *testing.T) {
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
	var expiry uint = 60 * 60 * 24

	userRepository := repository.NewUserRepository(
		repository.NewUserFinderImpl(*dbInstance),
		repository.NewUserWriterImpl(*dbInstance),
	)

	verifyUserUseCase := usecase.NewUserVerifyConfirmer(
		userRepository,
		tokenProvider,
		expiry,
	)

	createUserUseCase := usecase.NewUserCreator(
		userRepository,
		tokenProvider,
		expiry,
		md5Hash,
	)

	myUser := entity.User{
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		Status:   entity.INACTIVE.Value(),
		Role:     common.STUDENT.Value(),
		OTP:      999999,
		Password: "1234000@ABC",
	}
	myUser.ID = "d54b2cec-fb7b-4411-bb46-5fac0383e336"
	_, _ = createUserUseCase.Execute(context.Background(), myUser)

	c := gin.Context{}
	user := requester.JWTRequester{
		ID:   myUser.ID,
		Role: myUser.Role,
	}

	c.Set(requester.CurrentRequester, user)
	requester := c.MustGet(requester.CurrentRequester).(requester.Requester)
	ctx := context.WithValue(context.Background(), common.VerifyTokenKey{}, requester)

	t.Run("Valid OTP", func(t *testing.T) {
		_, err = verifyUserUseCase.Execute(ctx, 999999)

		assert.NoError(t, err)
	})

	t.Run("Invalid OTP", func(t *testing.T) {
		_, err = verifyUserUseCase.Execute(ctx, 99999)

		assert.Error(t, err)
	})
}
