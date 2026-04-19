package middleware

import (
	"strings"

	"catetin-backend/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized"})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"message": "unauthorized, invalid format"})
	}
	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"message": "invalid token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	c.Locals("user_id", userID)

	return c.Next()
}
