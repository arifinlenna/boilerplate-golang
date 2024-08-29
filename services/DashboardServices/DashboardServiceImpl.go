package dashboardservices

import "github.com/gofiber/fiber/v2"

func (dashboardServicesImpl *DashboardServicesImpl) Create(app *fiber.App) error {
	if err := dashboardServicesImpl.Create(app); err != nil {
		panic(err.Error())
	}
	return nil
}