package dashboardrepository

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	Create(app *fiber.App) error
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{
		DB: db,
	}
}