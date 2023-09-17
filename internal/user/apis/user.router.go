package apis

import (
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/apis/handler"
	"go-clean-architecture/internal/user/repository"
	"go-clean-architecture/provider/tokenprovider/jwt"
	utils "go-clean-architecture/utils/validator"

	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine, db db.Database, appConfig config.AppConfig) {
	// create repository instance
	userFinder := repository.NewUserFinderImpl(db)
	userWritter := repository.NewUserWriterImpl(db)
	repo := *repository.NewUserRepository(userFinder, userWritter)

	// jwt init
	tokenProvider := jwt.NewTokenJWTProvider(appConfig.JWTSecretKey)
	verifyTime := appConfig.VerifyTokenExpiry
	accessTime := appConfig.AccessTokenExpiry

	// validate incoming request
	validatorRequest := utils.NewValidator()

	// handler user
	handler := handler.NewUserHandler(
		repo,
		validatorRequest,
		tokenProvider,
		verifyTime,
		accessTime,
	)

	initRouter(engine, *handler)
}

func initRouter(engine *gin.Engine, handler handler.UserHandler) {
	v1 := engine.Group("api/v1")
	{
		v1.POST("user/create", handler.HandleCreateUser)
	}
}
