package main

import (
	"go-clean-architecture/cmd/app"
)

// @title CLEAN-EXAMPLE API
// @version 1.0
// @description More
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

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
