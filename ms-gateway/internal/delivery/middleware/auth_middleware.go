package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"ms-workspace/gateway/global"
	"ms-workspace/gateway/internal/entity"
	"ms-workspace/package/grpc"
	"ms-workspace/package/logger"
	go_proto "ms-workspace/package/proto/ms-user/v1/go-proto"
)

const (
	Unauthorized = "unauthorized"
)

type AuthMiddleware struct {
	userEndpoint string
}

func NewAuthMiddleware(userEndpoint string) *AuthMiddleware {
	return &AuthMiddleware{userEndpoint: userEndpoint}
}

func (c *AuthMiddleware) Authentication(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	token := headers["Authorization"]

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": Unauthorized,
		})
	}

	conn, cancelFunc, err := grpc.ConnectToGrpcServer(c.userEndpoint)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(entity.ResponseError{
			Message: err.Error(),
		})
	}

	defer cancelFunc()

	client := go_proto.NewUserServiceClient(conn)
	session, err := client.Authentication(ctx.Context(), &go_proto.AuthenticationRequest{
		Token: token,
	})

	if err != nil {
		global.Logger.Error("unauthorized", zap.Error(err), zap.Any("token", token), logger.TraceID(ctx.Context()))
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": Unauthorized,
		})
	}

	ctx.Locals("session", session)
	return ctx.Next()
}
