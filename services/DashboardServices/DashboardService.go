package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardrepository "github.com/lenna-ai/azureOneSmile.git/repositories/DashboardRepository"
)

type DashboardServices interface {
	Create(app *fiber.App) error
}

type DashboardServicesImpl struct {
	DashboardRepository dashboardrepository.DashboardRepository
}

func NewDashboardServices(dashboardRepository dashboardrepository.DashboardRepository) *DashboardServicesImpl {
	return &DashboardServicesImpl{
		DashboardRepository: dashboardRepository,
	}
}