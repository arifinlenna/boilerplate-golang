package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	dashboardrepository "github.com/lenna-ai/azureOneSmile.git/repositories/DashboardRepository"
)

type DashboardServices interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
}

type DashboardServicesImpl struct {
	DashboardRepository dashboardrepository.DashboardRepository
}

func NewDashboardServices(dashboardRepository dashboardrepository.DashboardRepository) *DashboardServicesImpl {
	return &DashboardServicesImpl{
		DashboardRepository: dashboardRepository,
	}
}