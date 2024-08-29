package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/azureOneSmile.git/injector"
)

func Router(app *fiber.App)  {
	allController := injector.InitializeController()
	
	
	dashboard := app.Group("dashboard")
	dashboard.Post("/", allController.DashboardController.Create)
}