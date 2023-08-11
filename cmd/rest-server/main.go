package main

import (
	_ "loanapp/docs" // import the docs, here TODO is my directory name
	"loanapp/postgresql"
	"loanapp/rest"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Loan App Service APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/

// @contact.name Loanapp API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @schemes http
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	db := postgresql.NewDB("test")

	// Instantiate postgresql-backed services.
	userService := postgresql.NewUserService(db)

	// Create the main handler instance
	handler := rest.NewHandler(router)

	// Attach underlying services to the HTTP server/handler.
	handler.UserService = userService

	// Run
	router.Run() // listen and serve on 0.0.0.0:8080

}
