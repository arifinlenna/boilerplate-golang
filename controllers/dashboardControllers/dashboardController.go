package dashboardcontrollers

import dashboardservices "github.com/lenna-ai/azureOneSmile.git/services/DashboardServices"

type DashboardController interface {
	
}

type DashboardControllerImpl struct {
	DashboardServices dashboardservices.DashboardServices 
}

func NewDashboardController(dashboardservices dashboardservices.DashboardServices) *DashboardControllerImpl {
	return &DashboardControllerImpl{
		DashboardServices: dashboardservices,
	}
}