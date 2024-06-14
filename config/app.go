package config

import (
	"github.com/Tirivashe/go-fiber-jwt/routes"
	"github.com/gofiber/fiber/v2"
)

type APIServer struct {
	mux  *fiber.App
	addr string
}

func (a *APIServer) Start() error {
	return a.mux.Listen(a.addr)
}

func NewAPIServer(addr string) *APIServer {
	app := fiber.New()
	routes.RegisterRoutes(app)
	return &APIServer{
		mux:  app,
		addr: addr,
	}
}
