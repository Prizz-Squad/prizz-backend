package middleware

import (
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func AuthMiddleware(roleBase fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, route := range util.Routes {
			if strings.Contains(c.Path(), route) {
				return c.Next()
			}
		}
		err := roleBase(c)
		if err != nil {
			return err
		}
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return fiber.ErrUnauthorized
		}

		token, err := services.ValidateJWTToken(tokenString)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", claims)
		return c.Next()
	}
}
