package app

import (
	_ "catetin-backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	"log"
)

type App struct {
	Fiber   *fiber.App
	Modules []Module
}

func New(modules ...Module) *App {
	app := fiber.New()

	app.Use(cors.New())

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
	a.Fiber.Get("/swagger/*", swagger.HandlerDefault)

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
