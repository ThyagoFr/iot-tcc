package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thyagofr/tcc/api/domain/service/user"
	"github.com/thyagofr/tcc/api/web/schemas"
	"github.com/thyagofr/tcc/api/web/security"
)

type Users struct {
	service user.Service
}

func (u *Users) Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(schemas.CreateUserRequest)
		err := ctx.BodyParser(request)
		if err != nil {
			return fiber.ErrUnprocessableEntity
		}
		model := request.ToModel()
		err = u.service.Register(model)
		if err != nil {
			return fiber.ErrBadRequest
		}
		return ctx.Status(fiber.StatusCreated).JSON(schemas.ToUserResponse(model))
	}
}

func (u *Users) Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(schemas.LoginRequest)
		err := ctx.BodyParser(request)
		if err != nil {
			return fiber.ErrUnprocessableEntity
		}
		model, err := u.service.Login(request.Email, request.Password)
		if err != nil {
			return fiber.ErrForbidden
		}
		token, err := security.GenerateToken(model.ID)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		ctx.Response().Header.Set(security.AuthorizationHeader, token)
		response := schemas.ToUserResponse(model)
		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
