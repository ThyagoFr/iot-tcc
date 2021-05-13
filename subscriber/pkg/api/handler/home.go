package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thyago/tcc/subscriber/pkg/api/service"
	"net/http"
)

func Home(srv service.Servicer) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response,err := srv.Home()
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, "an error occurred processing the request")
		}
		return ctx.Status(http.StatusOK).JSON(response)
	}
}