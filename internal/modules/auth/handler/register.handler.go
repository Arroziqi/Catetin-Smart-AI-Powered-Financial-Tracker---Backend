package handler

import (
	"catetin-backend/internal/modules/auth/dto"

	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary Register user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.RegisterRequest true "Register Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body dto.RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	err := h.uc.Register(body.Email, body.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "success"})
}
