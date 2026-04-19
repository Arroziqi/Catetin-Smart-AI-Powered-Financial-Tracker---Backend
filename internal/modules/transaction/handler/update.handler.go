package handler

import (
	"strconv"

	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/usecase"

	"github.com/gofiber/fiber/v2"
)

type UpdateTransactionHandler struct {
	uc usecase.UpdateTransactionUsecase
}

func NewUpdateTransactionHandler(uc usecase.UpdateTransactionUsecase) *UpdateTransactionHandler {
	return &UpdateTransactionHandler{uc: uc}
}

// @Summary Update a transaction
// @Description Updates an existing transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Param request body dto.UpdateTransactionRequest true "Transaction Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/transactions/{id} [put]
func (h *UpdateTransactionHandler) Handle(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid transaction ID",
		})
	}

	var req dto.UpdateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid payload",
		})
	}

	if err := h.uc.Execute(uint(id), userID, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaction updated successfully",
	})
}
