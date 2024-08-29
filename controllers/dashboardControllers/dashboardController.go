package dashboardcontrollers

import (
	"github.com/gofiber/fiber/v2"
	dashboardservices "github.com/lenna-ai/azureOneSmile.git/services/DashboardServices"
)

type DashboardController interface {
	Create(app *fiber.Ctx) error 
}

type DashboardControllerImpl struct {
	DashboardServices dashboardservices.DashboardServices 
}

func NewDashboardController(dashboardservices dashboardservices.DashboardServices) *DashboardControllerImpl {
	return &DashboardControllerImpl{
		DashboardServices: dashboardservices,
	}
}