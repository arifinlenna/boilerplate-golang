package dashboardrepository

import (
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	Create(app *fiber.Ctx, user *usermodel.User) error
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{
		DB: db,
	}
}