//go:build wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/service"
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

func ProvideService() (*service.Service, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider)
	return &service.Service{}, nil
}

func ProvideController() (*controller.Controller, error) {
	wire.Build(database.DataBaseProvider, service.ServiceProvider, controller.ControllerProvider)
	return &controller.Controller{}, nil
}
