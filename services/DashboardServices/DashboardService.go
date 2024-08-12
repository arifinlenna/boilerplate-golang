package dashboardservices

import dashboardrepository "github.com/lenna-ai/azureOneSmile.git/repositories/DashboardRepository"

type DashboardServices interface {
	
}

type DashboardServicesImpl struct {
	DashboardRepository dashboardrepository.DashboardRepository
}

func NewDashboardServices(dashboardRepository dashboardrepository.DashboardRepository) *DashboardServicesImpl {
	return &DashboardServicesImpl{
		DashboardRepository: dashboardRepository,
	}
}