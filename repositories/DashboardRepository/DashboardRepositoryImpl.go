package dashboardrepository

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
)

func (dashboardRepositoryImpl *DashboardRepositoryImpl) Create(app *fiber.Ctx, user *usermodel.User) error {
	fmt.Println("user")
	fmt.Printf("%+v\n",user)
	if err := dashboardRepositoryImpl.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}