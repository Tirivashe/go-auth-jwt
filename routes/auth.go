package routes

import (
	"github.com/Tirivashe/go-fiber-jwt/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	authRoutes(v1)
}

func authRoutes(v1 fiber.Router) {
	v1.Post("/auth/login", handlers.Login)
	v1.Post("/auth/signup", handlers.SignUp)
	v1.Get("/health", handlers.HealthCheck)
}
