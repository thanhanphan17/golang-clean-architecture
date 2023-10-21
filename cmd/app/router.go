package app

import (
	"go-clean-architecture/config"
	"go-clean-architecture/db"
	user "go-clean-architecture/internal/user/apis"

	"github.com/gin-gonic/gin"
)

func initRouter(engine *gin.Engine, db db.Database, appConfig config.AppConfig) {
	user.InitRouter(engine, db, appConfig)
}
