package http

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Not Found",
			})
		},
	)
}
