package handlers

import (
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/constants"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": constants.STATUS_OK,
	})
}
