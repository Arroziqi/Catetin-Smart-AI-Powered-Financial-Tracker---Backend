package app

import (
	_ "catetin-backend/docs"

	swagger "github.com/swaggo/fiber-swagger"

	"log"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber   *fiber.App
	Modules []Module
}

func New(modules ...Module) *App {
	app := fiber.New()

	return &App{
		Fiber:   app,
		Modules: modules,
	}
}

func (a *App) Setup() {
	// 🔥 migrate semua module dulu
	for _, m := range a.Modules {
		m.Migrate()
	}

	// 🔥 register routes
	for _, m := range a.Modules {
		m.Register(a.Fiber)
	}

	// health
	a.Fiber.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Swagger
	a.Fiber.Get("/swagger/*", swagger.WrapHandler)

	// log routes
	log.Println("=== REGISTERED ROUTES ===")
	for _, r := range a.Fiber.GetRoutes() {
		log.Printf("%-6s %-30s %s\n", r.Method, r.Path, r.Name)
	}
}

func (a *App) collectSwaggerModels() []interface{} {
	var models []interface{}

	for _, m := range a.Modules {
		models = append(models, m.Swagger()...)
	}

	return models
}
