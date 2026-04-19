package handler

import (
	"strconv"

	"catetin-backend/internal/modules/transaction/usecase"

	"github.com/gofiber/fiber/v2"
)

type DeleteTransactionHandler struct {
	uc usecase.DeleteTransactionUsecase
}

func NewDeleteTransactionHandler(uc usecase.DeleteTransactionUsecase) *DeleteTransactionHandler {
	return &DeleteTransactionHandler{uc: uc}
}

// @Summary Delete a transaction
// @Description Deletes an existing transaction by ID
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/transactions/{id} [delete]
func (h *DeleteTransactionHandler) Handle(c *fiber.Ctx) error {
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

	if err := h.uc.Execute(uint(id), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transaction deleted successfully",
	})
}
