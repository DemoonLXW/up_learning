//go:build wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/application"
	"github.com/DemoonLXW/up_learning/config"
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func ProvideDataBase() (*ent.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &ent.Client{}, nil
}

func ProvideRedis() (*redis.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &redis.Client{}, nil
}

func ProvideService() (*service.Services, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, config.ConfigProvider)
	return &service.Services{}, nil
}

func ProvideController() (*controller.Controllers, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, config.ConfigProvider, controller.ControllerProvider)
	return &controller.Controllers{}, nil
}

func ProvideApplication() (*gin.Engine, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, config.ConfigProvider, controller.ControllerProvider, application.ApplicationProvider)
	return &gin.Engine{}, nil
}
