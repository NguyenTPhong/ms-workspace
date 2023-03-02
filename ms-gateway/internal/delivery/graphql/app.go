package graphql

import (
	"github.com/gofiber/fiber/v2"
	"ms-workspace/gateway/internal/delivery/middleware"
)

func InitGraphQL(app *fiber.App, authMiddleware *middleware.AuthMiddleware) {
	app.Get("/", playgroundHandler())
	app.Post("/query", authMiddleware.Authentication, graphqlHandler())
}
