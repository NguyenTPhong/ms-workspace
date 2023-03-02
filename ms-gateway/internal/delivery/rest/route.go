package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"ms-workspace/gateway/docs"
	"ms-workspace/gateway/global/config"
)

func (h *Handler) InitV1Route() {

	docs.SwaggerInfo.Title = "API Documentations"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Domain
	docs.SwaggerInfo.Schemes = []string{"http"}

	// documentation public
	h.fiberApp.Get("/swagger/*", swagger.HandlerDefault)

	api := h.fiberApp.Group("/api")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{"message": "service health is good"})
	})

	// v1
	v1 := api.Group("/v1")

	//authentication
	authGroup := v1.Group("/auth")
	authGroup.Post("/register", h.authController.RegisNewUser)
	authGroup.Post("/", h.authController.Login)

}
