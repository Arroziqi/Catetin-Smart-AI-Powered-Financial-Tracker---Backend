package handler

import (
	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/usecase"

	"github.com/gofiber/fiber/v2"
)

type CreateTransactionHandler struct {
	uc usecase.CreateTransactionUsecase
}

func NewCreateTransactionHandler(uc usecase.CreateTransactionUsecase) *CreateTransactionHandler {
	return &CreateTransactionHandler{uc: uc}
}

// @Summary Create a transaction
// @Description Creates a new financial transaction (income/expense)
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateTransactionRequest true "Transaction Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/transactions [post]
func (h *CreateTransactionHandler) Handle(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	var req dto.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid payload",
		})
	}

	if err := h.uc.Execute(userID, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "transaction created successfully",
	})
}
