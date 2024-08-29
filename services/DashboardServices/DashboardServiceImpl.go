package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

func (dashboardServicesImpl *DashboardServicesImpl) Create(app *fiber.Ctx, user *usermodel.User) error {
	if err := dashboardServicesImpl.DashboardRepository.Create(app, user); err != nil {
		panic(err.Error())
	}
	return nil
}