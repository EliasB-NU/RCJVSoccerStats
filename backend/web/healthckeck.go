package web

import "github.com/gofiber/fiber/v2"

func getHealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("OK")
}
