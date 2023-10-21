package apis

import (
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/apis/handler"
	"go-clean-architecture/internal/user/repository"
	"go-clean-architecture/middleware"
	"go-clean-architecture/provider/hashprovider"
	"go-clean-architecture/provider/tokenprovider/jwt"
	tokentype "go-clean-architecture/provider/tokenprovider/type"
	utils "go-clean-architecture/utils/validator"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router for the application.
// It initializes the necessary dependencies and handlers.
func setUpHandler(db db.Database, appConfig config.AppConfig) *handler.UserHandler {
	// Create the user repository with the database
	repo := *repository.NewUserRepository(
		repository.NewUserFinderImpl(db),
		repository.NewUserWriterImpl(db),
	)

	// Create the JWT token provider with the secret key from the app configuration
	tokenProvider := jwt.NewJWTProvider(appConfig.JWTSecretKey)

	// Create the request validator
	validatorRequest := utils.NewValidator()

	md5Hash := hashprovider.NewMd5Hash()

	// Create the user handler with the repository, request validator,
	// token provider, and token expiry settings from the app configuration
	userHandler := handler.NewUserHandler(
		repo,
		validatorRequest,
		tokenProvider,
		appConfig.VerifyTokenExpiry,
		appConfig.AccessTokenExpiry,
		md5Hash,
	)

	return userHandler
}

func InitRouter(engine *gin.Engine, db db.Database, appConfig config.AppConfig) {
	userHandler := setUpHandler(db, appConfig)

	route := engine.Group("/api/v1/user")
	{
		route.POST("/register", userHandler.HandleCreateUser)
		route.POST("/login", userHandler.HandleLoginUser)
		route.GET("/verify", userHandler.HandleVerifyUser)
		route.POST("/comfirm-verify",
			middleware.RequireAuthorization(&appConfig, tokentype.VERIFY_TOKEN.Value()),
			userHandler.HandleConfirmVerifyUser,
		)

		route.GET("/otp-resend",
			middleware.RequireAuthorization(&appConfig, tokentype.VERIFY_TOKEN.Value()),
			userHandler.HandleSendOTP,
		)
	}
}
