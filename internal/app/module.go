package app

import "github.com/gofiber/fiber/v2"

type Module interface {
	Register(app *fiber.App)
	Migrate()
	Swagger() []interface{}
}
