package transaction

import (
	"catetin-backend/internal/middleware"
	"catetin-backend/internal/modules/transaction/domain"
	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/handler"
	"catetin-backend/internal/modules/transaction/repository"
	"catetin-backend/internal/modules/transaction/usecase"

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
	m.db.AutoMigrate(&domain.Transaction{})
}

func (m *Module) Register(app *fiber.App) {
	// Initialize Repository
	repo := repository.NewTransactionRepository(m.db)

	// Initialize Usecases
	createUc := usecase.NewCreateTransactionUsecase(repo)
	getUc := usecase.NewGetTransactionUsecase(repo)
	updateUc := usecase.NewUpdateTransactionUsecase(repo)
	deleteUc := usecase.NewDeleteTransactionUsecase(repo)

	// Initialize Handlers
	createH := handler.NewCreateTransactionHandler(createUc)
	getH := handler.NewGetTransactionHandler(getUc)
	updateH := handler.NewUpdateTransactionHandler(updateUc)
	deleteH := handler.NewDeleteTransactionHandler(deleteUc)

	// Configure Routes
	api := app.Group("/api/v1/transactions")

	api.Post("/", middleware.AuthMiddleware, createH.Handle).Name("transaction.create")
	api.Get("/", middleware.AuthMiddleware, getH.Handle).Name("transaction.get")
	api.Put("/:id", middleware.AuthMiddleware, updateH.Handle).Name("transaction.update")
	api.Delete("/:id", middleware.AuthMiddleware, deleteH.Handle).Name("transaction.delete")
}

func (m *Module) Swagger() []interface{} {
	return []interface{}{
		dto.CreateTransactionRequest{},
		dto.UpdateTransactionRequest{},
		dto.GetTransactionQuery{},
		domain.Transaction{},
	}
}
