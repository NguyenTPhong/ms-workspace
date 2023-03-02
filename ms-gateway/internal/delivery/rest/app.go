package rest

import (
	"github.com/gofiber/fiber/v2"
	"ms-workspace/gateway/global/config"
	"ms-workspace/gateway/internal/delivery/middleware"
	"ms-workspace/gateway/internal/delivery/rest/controller"
)

func InitRestApi(app *fiber.App, authMiddleware *middleware.AuthMiddleware) {
	// init controller
	handler := NewHandler(
		WithEngine(app),
	)
	
	// middleware
	handler.authMiddleware = authMiddleware

	// auth controller
	authController := controller.NewAuthController(config.UserServiceHost)
	handler.authController = authController

	// create route
	handler.InitV1Route()
}

type Handler struct {
	fiberApp       *fiber.App
	authMiddleware *middleware.AuthMiddleware
	authController *controller.AuthController
}

type HandlerOption func(*Handler)

func NewHandler(options ...HandlerOption) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithEngine(r *fiber.App) HandlerOption {
	return func(handler *Handler) {
		handler.fiberApp = r
	}
}
