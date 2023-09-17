package app

import (
	"flag"
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/apis"
	"go-clean-architecture/middleware"
	"os"
	"strconv"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setEnv(flagArg *FlagArgument) {
	if err := os.Setenv("CONFIG_PATH", *flagArg.ConfigPath); err != nil {
		panic(err)
	}

	if err := os.Setenv("CONFIG_ENV", *flagArg.Env); err != nil {
		panic(err)
	}

	if err := os.Setenv(
		"CONFIG_ALLOW_MIGRATION",
		strconv.FormatBool(*flagArg.Upgrade),
	); err != nil {
		panic(err)
	}
}

func Run() {
	flagArg := getFlagArgument()
	flag.Parse()

	setEnv(flagArg)

	appConfig := config.InitLoadAppConf()

	envValue := os.Getenv("CONFIG_ALLOW_MIGRATION")
	allowUpgrade, err := strconv.ParseBool(envValue)
	if err != nil {
		panic(err)
	}

	dbInstance := db.InitDatabase(allowUpgrade, *appConfig)

	gin.SetMode(gin.DebugMode)
	appRouter := gin.Default()

	// Init middleware
	appRouter.Use(helmet.Default())
	appRouter.Use(middleware.Recover())
	appRouter.Use(middleware.RequestID())

	// Allow CORS
	appRouter.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowHeaders: []string{
				"Origin",
				"Content-Type",
				"Accept",
				"Authorization",
			},
		},
	))

	apis.SetupRouter(appRouter, *dbInstance, *appConfig)

	if err := appRouter.Run(
		fmt.Sprintf(":%d", appConfig.ServicePort),
	); err != nil {
		panic(err)
	}
}
