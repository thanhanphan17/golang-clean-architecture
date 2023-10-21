package test

import (
	"context"
	"go-clean-architecture/common"
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/entity"
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/internal/user/repository"
	"go-clean-architecture/provider/hashprovider"
	"go-clean-architecture/provider/tokenprovider/jwt"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

// Create user integration test
// Make sure your database is ready for connection before testing
func TestCreateUser(t *testing.T) {
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

	useCase := usecase.NewUserCreator(
		userRepository,
		tokenProvider,
		expiry,
		md5Hash,
	)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			user := entity.User{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Phone:    gofakeit.Phone(),
				Status:   entity.INACTIVE.Value(),
				Role:     common.STUDENT.Value(),
				Password: "1234000@ABC",
			}

			_, err = useCase.Execute(context.Background(), user)

			assert.NoError(t, err)
		}()

		wg.Wait()
	}
}
