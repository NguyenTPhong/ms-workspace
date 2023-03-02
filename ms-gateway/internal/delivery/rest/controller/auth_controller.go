package controller

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"
	"ms-workspace/gateway/internal/entity"
	"ms-workspace/package/grpc"
	go_proto "ms-workspace/package/proto/ms-user/v1/go-proto"
)

type AuthController struct {
	userServiceHost string
}

func NewAuthController(userServiceHost string) *AuthController {
	return &AuthController{
		userServiceHost: userServiceHost,
	}
}

// RegisNewUser godoc
// @Summary register new user
// @Tags Authentication
// @ID register-new-user
// @Accept json
// @Produce json
// @Param json body entity.CreateUserRequest true "json body"
// @Success 200 {object} entity.CreateUserResponse
// @Failure 400 {object} entity.ResponseError
// @Router /api/v1/auth/register [post]
func (c *AuthController) RegisNewUser(ctx *fiber.Ctx) error {
	var req entity.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Message: err.Error(),
		})
	}

	conn, cancelFunc, err := grpc.ConnectToGrpcServer(c.userServiceHost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(entity.ResponseError{
			Message: err.Error(),
		})
	}

	defer cancelFunc()

	client := go_proto.NewUserServiceClient(conn)

	res, err := client.CreateUser(ctx.Context(), &go_proto.CreateUserRequest{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Password:    req.Password,
	})

	if err != nil {
		message := err.Error()
		if grpcErr, ok := status.FromError(err); ok {
			message = grpcErr.Message()
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Message: message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(entity.CreateUserResponse{
		Email:       res.Email,
		Id:          res.Id,
		PhoneNumber: res.PhoneNumber,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Status:      res.Status,
		CreatedAt:   res.CreatedAt.AsTime(),
	})
}

// Login godoc
// @Summary log in
// @Tags Authentication
// @ID login
// @Accept json
// @Produce json
// @Param json body entity.LoginRequest true "json body"
// @Success 200 {object} entity.LoginResponse
// @Failure 400 {object} entity.ResponseError
// @Router /api/v1/auth [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req entity.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Message: err.Error(),
		})
	}

	conn, cancelFunc, err := grpc.ConnectToGrpcServer(c.userServiceHost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(entity.ResponseError{
			Message: err.Error(),
		})
	}

	defer cancelFunc()

	client := go_proto.NewUserServiceClient(conn)

	res, err := client.Login(ctx.Context(), &go_proto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		message := err.Error()
		if grpcErr, ok := status.FromError(err); ok {
			message = grpcErr.Message()
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(entity.ResponseError{
			Message: message,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(entity.LoginResponse{
		Token:    res.Token,
		ExpireAt: res.ExpiredAt.AsTime(),
	})

}
