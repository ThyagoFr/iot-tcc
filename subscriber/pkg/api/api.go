package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
	"github.com/thyago/tcc/subscriber/pkg/api/handler"
	"github.com/thyago/tcc/subscriber/pkg/api/service"
	"log"
)

type API struct {
	engine *fiber.App
	serv service.Servicer
}

func NewAPI(serv service.Servicer) *API {
	api := &API{serv: serv}
	api.engine = fiber.New()
	api.setupRoutes()
	api.setupMiddlewares()
	return api
}

func (api *API) setupRoutes() {
	api.engine.Get("/api/v1/home", handler.Home(api.serv))
	api.engine.Get("/api/v1/devices/:id", handler.Historical(api.serv))
}

func (api *API) setupMiddlewares() {
	api.engine.Use(requestid.New())
	api.engine.Use(recover.New())
}

func (api *API) ListenAndServe() error {
	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server running on port %s \n", port)
	return api.engine.Listen(fmt.Sprintf("localhost:%s", port))
}
