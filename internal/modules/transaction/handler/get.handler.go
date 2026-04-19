package handler

import (
	"catetin-backend/internal/modules/transaction/dto"
	"catetin-backend/internal/modules/transaction/usecase"

	"github.com/gofiber/fiber/v2"
)

type GetTransactionHandler struct {
	uc usecase.GetTransactionUsecase
}

func NewGetTransactionHandler(uc usecase.GetTransactionUsecase) *GetTransactionHandler {
	return &GetTransactionHandler{uc: uc}
}

// @Summary Get transactions
// @Description Retrieve a list of transactions with pagination and filtering
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param category query string false "Filter by category"
// @Param type query string false "Filter by type (income | expense)"
// @Param start_date query string false "Start date YYYY-MM-DD"
// @Param end_date query string false "End date YYYY-MM-DD"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/transactions [get]
func (h *GetTransactionHandler) Handle(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	var query dto.GetTransactionQuery
	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid query parameters",
		})
	}

	transactions, total, err := h.uc.Execute(userID, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	page := query.Page
	if page <= 0 {
		page = 1
	}
	limit := query.Limit
	if limit <= 0 {
		limit = 10
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transactions,
		"meta": fiber.Map{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}
