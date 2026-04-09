package handler

import (
	"catetin-backend/internal/modules/auth/dto"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login Data"
// @Success 200 {object} dto.AuthResponse
// @Failure 401 {object} map[string]string
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body dto.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	token, err := h.uc.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	return c.JSON(dto.AuthResponse{
		AccessToken: token,
	})
}
