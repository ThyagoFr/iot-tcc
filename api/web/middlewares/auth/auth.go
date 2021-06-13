package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thyagofr/tcc/api/web/security"
)

func Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get(security.AuthorizationHeader)
		err := security.ValidateToken(auth)
		if err != nil {
			return fiber.ErrForbidden
		}
		return ctx.Next()
	}
}
