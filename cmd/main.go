package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"catetin-backend/internal/config"
	"catetin-backend/internal/database"
	"catetin-backend/internal/modules/auth"
	"catetin-backend/internal/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()

	database.DB.AutoMigrate(&auth.User{})

	app := fiber.New()

	routes.Setup(app)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Fatal(app.Listen(":5000"))
}