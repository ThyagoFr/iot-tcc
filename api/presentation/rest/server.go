package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
	"github.com/thyago/tcc/api-service/application/service"
	"log"
)

type HTTPServer struct {
	engine *fiber.App
	serv   service.Servicer
}

func NewHTTPServer(serv service.Servicer) *HTTPServer {
	api := &HTTPServer{serv: serv}
	api.engine = fiber.New()
	api.setupRoutes()
	api.setupMiddlewares()
	return api
}

func (server *HTTPServer) setupRoutes() {
	apiV1 := server.engine.Group("/api/v1")

	apiV1.Post("/register")
	apiV1.Post("/login")

	notificationResource := apiV1.Group("/notifications")
	notificationResource.Post("")
	notificationResource.Patch("/:id")
	notificationResource.Delete("/:id")
	notificationResource.Get("/:id")

	userResource := apiV1.Group("/users")
	userResource.Patch("/:id")
	userResource.Delete("/:id")

}

func (server *HTTPServer) setupMiddlewares() {
	server.engine.Use(requestid.New())
	server.engine.Use(recover.New())
}

func (server *HTTPServer) ListenAndServe() error {
	port := viper.GetString("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server running on port %s \n", port)
	return server.engine.Listen(fmt.Sprintf("localhost:%s", port))
}
