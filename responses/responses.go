package responses

import "github.com/gofiber/fiber/v2"

// A standard ErrorResponse
func ErrorResponse(c *fiber.Ctx, err error) error {
	output := fiber.Map{
		"error": err.Error(),
	}
	return c.JSON(output)
}

// A standard SuccessResponse
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	output := fiber.Map{
		"output": data,
	}
	return c.JSON(output)
}
