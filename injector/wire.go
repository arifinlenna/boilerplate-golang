//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/lenna-ai/azureOneSmile.git/controllers"
	dashboardcontrollers "github.com/lenna-ai/azureOneSmile.git/controllers/dashboardControllers"
	dashboardrepository "github.com/lenna-ai/azureOneSmile.git/repositories/DashboardRepository"
	dashboardservices "github.com/lenna-ai/azureOneSmile.git/services/DashboardServices"
)


var dashboardController = wire.NewSet(
	dashboardrepository.NewDashboardRepository,
	wire.Bind(new(dashboardrepository.DashboardRepository), new(*dashboardrepository.DashboardRepositoryImpl)),
	dashboardservices.NewDashboardServices,
	wire.Bind(new(dashboardservices.DashboardServices),new(*dashboardservices.DashboardServicesImpl)),
	dashboardcontrollers.NewDashboardController,
	wire.Bind(new(dashboardcontrollers.DashboardController),new(*dashboardcontrollers.DashboardControllerImpl)),
)

var setAllControllers = wire.NewSet(
	dashboardController,
	wire.Struct(new(controllers.AllControllers),"*"),
)

func InitializeController() (*controllers.AllControllers){
	// wire.Build(setLoginController, controllers.NewAllControllers)
	wire.Build(setAllControllers)
	return &controllers.AllControllers{}
}