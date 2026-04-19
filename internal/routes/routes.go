package routes

import (
	"catetin-backend/internal/modules/auth/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, authHandler *handler.AuthHandler) {
	api := app.Group("/api/v1")

	authGroup := api.Group("/auth")
	authGroup.Post("/register", authHandler.Register).Name("auth.register")
	authGroup.Post("/login", authHandler.Login).Name("auth.login")
}
