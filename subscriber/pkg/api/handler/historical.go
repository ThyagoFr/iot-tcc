package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thyago/tcc/subscriber/pkg/api/service"
	"net/http"
	"time"
)

func Historical(srv service.Servicer) fiber.Handler{
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}
		to := ctx.Query("to")
		toParseTime,err := time.Parse("2006-02-01", to)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}
		from := ctx.Query("from")
		fromParseTime,err := time.Parse("2006-02-01", from)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}
		historical, err := srv.Historical(id,toParseTime,fromParseTime)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError,"an error occurred processing the request")
		}
		return ctx.Status(http.StatusOK).JSON(historical)
	}
}
