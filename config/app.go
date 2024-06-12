package config

import "github.com/gofiber/fiber/v2"

type APIServer struct {
	mux  *fiber.App
	addr string
}

func (a *APIServer) Start() error {
	return a.mux.Listen(a.addr)
}

func NewAPIServer(addr string) *APIServer {
	app := fiber.New()
	return &APIServer{
		mux:  app,
		addr: addr,
	}
}
