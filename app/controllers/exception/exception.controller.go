package exception

import (
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/constants"
	"github.com/gofiber/fiber/v2"
)

func ExceptionNotFound(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  constants.STATUS_FAIL,
		"message": "Not Found",
		"err":     err.Error(),
	})
}
