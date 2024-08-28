package dashboardrepository

import "gorm.io/gorm"

type DashboardRepository interface {
	
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{
		DB: db,
	}
}