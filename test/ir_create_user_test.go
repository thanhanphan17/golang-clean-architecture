package test

import (
	"context"
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/entity"
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/internal/user/repository"
	"go-clean-architecture/provider/tokenprovider/jwt"
	"os"
	"testing"
)

func TestCreateUser(t *testing.T) {
	configPath := "../config/env"
	env := "local"
	upgrade := "false"

	_ = os.Setenv("CONFIG_PATH", configPath)
	_ = os.Setenv("CONFIG_ENV", env)
	_ = os.Setenv("CONFIG_ALLOW_MIGRATION", upgrade)

	appConfig := config.InitLoadAppConf()

	dbInstance := db.InitDatabase(true, *appConfig)
	tokenProvider := jwt.NewTokenJWTProvider(appConfig.JWTSecretKey)
	expiry := 60 * 60 * 24

	userFinder := repository.NewUserFinderImpl(*dbInstance)
	userWriter := repository.NewUserWriterImpl(*dbInstance)

	userRepository := repository.NewUserRepository(userFinder, userWriter)
	createUserUsecase := usecase.NewUserCreator(
		userRepository,
		tokenProvider,
		expiry,
	)

	_, err := createUserUsecase.Execute(context.Background(), entity.User{
		Name:     "Phan Thanh An",
		Email:    "thanhanphan17@gmail.com",
		Password: "abcd1234@Q",
	})

	if err != nil {
		t.Error(err)
	}
}
