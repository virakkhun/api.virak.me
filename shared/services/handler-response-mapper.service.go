package shared_services

import "github.com/gofiber/fiber/v2"

func HandlerResponseMapper(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}
