package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"ms-workspace/gateway/global"
	"ms-workspace/gateway/global/config"
	"ms-workspace/gateway/internal/delivery/graphql"
	"ms-workspace/gateway/internal/delivery/middleware"
	"ms-workspace/gateway/internal/delivery/rest"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	global.Init()
	defer global.DeInit()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length,Authorization",
	}))

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(config.UserServiceHost)

	// init network interface
	graphql.InitGraphQL(app, authMiddleware)
	rest.InitRestApi(app, authMiddleware)

	app.Listen(":80")
}
