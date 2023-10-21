package app

import (
	"flag"
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/middleware"
	"os"
	"strconv"

	_ "go-clean-architecture/docs"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run is a function that initializes and runs the application.
// It sets up the environment, initializes the database, configures the router and middleware,
// and starts the application server.
func Run() {
	// Get the flag argument from command line
	flagArg := getFlagArgument()
	flag.Parse()

	// Set the environment based on the flag argument
	setEnv(flagArg)

	// Initialize the application configuration
	appConfig := config.InitLoadAppConf()

	// Initialize the database instance
	dbInstance := db.InitDatabase(parseAllowMigration(), *appConfig)

	// Create the Gin router with default middleware
	appRouter := gin.Default()

	// Set Gin to debug mode
	gin.SetMode(gin.DebugMode)

	// Initialize middleware
	initMiddleware(appRouter)

	// Allow CORS
	allowCORS(appRouter)

	// Serve Swagger documentation
	appRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Setup API routes
	initRouter(appRouter, *dbInstance, *appConfig)

	// Start the application server on the specified port
	if err := appRouter.Run(fmt.Sprintf(":%d", appConfig.ServicePort)); err != nil {
		panic(err)
	}
}

func parseAllowMigration() bool {
	envValue := os.Getenv("CONFIG_ALLOW_MIGRATION")
	allowUpgrade, err := strconv.ParseBool(envValue)
	if err != nil {
		panic(err)
	}
	return allowUpgrade
}

func initMiddleware(appRouter *gin.Engine) {
	appRouter.Use(helmet.Default())
	appRouter.Use(middleware.Recover())
	appRouter.Use(middleware.RequestID())
}

func allowCORS(appRouter *gin.Engine) {
	appRouter.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
	}))
}
