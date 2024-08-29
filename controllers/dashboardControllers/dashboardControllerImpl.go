package dashboardcontrollers

import "github.com/gofiber/fiber/v2"

func (dashboardControllerImpl *DashboardControllerImpl) Create(app *fiber.App) error {
	if err := dashboardControllerImpl.Create(app); err != nil {
		panic(err)
	}
	return nil
}