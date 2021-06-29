package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajeshj3/jwt-auth/responses"
	"github.com/rajeshj3/jwt-auth/utils"
)

// JWT cookies must be sent along
func AuthMiddleware(c *fiber.Ctx) error {

	// Checking for JWT
	_, err := utils.ClaimJWT(c)

	if err != nil {
		return responses.ErrorResponse(c, err)
	}

	// User is authorized
	return c.Next()
}
