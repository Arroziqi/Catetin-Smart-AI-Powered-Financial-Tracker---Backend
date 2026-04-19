package auth

import (
	"catetin-backend/internal/modules/auth/domain"
	"catetin-backend/internal/modules/auth/dto"
	"catetin-backend/internal/modules/auth/handler"
	"catetin-backend/internal/modules/auth/repository"
	"catetin-backend/internal/modules/auth/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Module struct {
	db *gorm.DB
}

func NewModule(db *gorm.DB) *Module {
	return &Module{db: db}
}

func (m *Module) Migrate() {
	m.db.AutoMigrate(&domain.User{})
}

func (m *Module) Register(app *fiber.App) {
	// wiring internal module
	userRepo := repository.NewUserRepository(m.db)
	authUC := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUC)

	// routes
	api := app.Group("/api/v1/auth")

	api.Post("/register", authHandler.Register).Name("auth.register")
	api.Post("/login", authHandler.Login).Name("auth.login")
}

func (m *Module) Swagger() []interface{} {
	return []interface{}{
		dto.RegisterRequest{},
		dto.LoginRequest{},
		dto.AuthResponse{},
	}
}
