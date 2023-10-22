//go:build wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/application"
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/DemoonLXW/up_learning/workflow"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func ProvideDataBase() (*ent.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &ent.Client{}, nil
}

func ProvideWorkflowHelper() (*workflow.WorkflowHelper, error) {
	wire.Build(workflow.WorkflowHelperProvider)
	return &workflow.WorkflowHelper{}, nil
}

func ProvideRedis() (*redis.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &redis.Client{}, nil
}

func ProvideService() (*service.Services, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, workflow.WorkflowHelperProvider)
	return &service.Services{}, nil
}

func ProvideController() (*controller.Controllers, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, controller.ControllerProvider, workflow.WorkflowHelperProvider)
	return &controller.Controllers{}, nil
}

func ProvideApplication() (*gin.Engine, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, controller.ControllerProvider, application.ApplicationProvider, workflow.WorkflowHelperProvider)
	return &gin.Engine{}, nil
}
