// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/redis/go-redis/v9"
)

// Injectors from injection.go:

func ProvideDataBase() (*ent.Client, error) {
	dataBaseConfig, err := database.ProvideDatabaseConfig()
	if err != nil {
		return nil, err
	}
	client, err := database.ProvideDB(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ProvideRedis() (*redis.Client, error) {
	dataBaseConfig, err := database.ProvideDatabaseConfig()
	if err != nil {
		return nil, err
	}
	client, err := database.ProvideRedis(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}
