package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App)  {
	dashboard := app.Group("dashboard")
	dashboard.Post("/", func(c *fiber.Ctx) error {
		json := fiber.Map{
			"name":"arifin",
		}
		return c.JSON(json)
	})
}