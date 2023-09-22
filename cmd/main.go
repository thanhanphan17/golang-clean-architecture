package main

import (
	"go-clean-architecture/cmd/app"
)

// @title Golang Clean Architecture Example
// @version 1.0
// @description Simple implementation of clean architecture
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email thanhanphan17@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey jwt
// @in header
// @name Authorization

// @host localhost:8888
// @BasePath /api/v1
func main() {
	app.Run()
}
