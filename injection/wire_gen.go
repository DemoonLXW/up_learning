// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/application"
	"github.com/DemoonLXW/up_learning/controller"
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/gin-gonic/gin"
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

func ProvideService() (*service.Services, error) {
	dataBaseConfig, err := database.ProvideDatabaseConfig()
	if err != nil {
		return nil, err
	}
	client, err := database.ProvideDB(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	managementService := &service.ManagementService{
		DB: client,
	}
	redisClient, err := database.ProvideRedis(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	authService := &service.AuthService{
		DB:    client,
		Redis: redisClient,
	}
	commonService := &service.CommonService{
		DB: client,
	}
	services := &service.Services{
		Management: managementService,
		Auth:       authService,
		Common:     commonService,
	}
	return services, nil
}

func ProvideController() (*controller.Controllers, error) {
	dataBaseConfig, err := database.ProvideDatabaseConfig()
	if err != nil {
		return nil, err
	}
	client, err := database.ProvideDB(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	managementService := &service.ManagementService{
		DB: client,
	}
	redisClient, err := database.ProvideRedis(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	authService := &service.AuthService{
		DB:    client,
		Redis: redisClient,
	}
	commonService := &service.CommonService{
		DB: client,
	}
	services := &service.Services{
		Management: managementService,
		Auth:       authService,
		Common:     commonService,
	}
	authController := &controller.AuthController{
		Services: services,
	}
	managementController := &controller.ManagementController{
		Services: services,
	}
	commonController := &controller.CommonController{
		Services: services,
	}
	controllers := &controller.Controllers{
		Auth:       authController,
		Management: managementController,
		Common:     commonController,
	}
	return controllers, nil
}

func ProvideApplication() (*gin.Engine, error) {
	dataBaseConfig, err := database.ProvideDatabaseConfig()
	if err != nil {
		return nil, err
	}
	client, err := database.ProvideDB(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	managementService := &service.ManagementService{
		DB: client,
	}
	redisClient, err := database.ProvideRedis(dataBaseConfig)
	if err != nil {
		return nil, err
	}
	authService := &service.AuthService{
		DB:    client,
		Redis: redisClient,
	}
	commonService := &service.CommonService{
		DB: client,
	}
	services := &service.Services{
		Management: managementService,
		Auth:       authService,
		Common:     commonService,
	}
	authController := &controller.AuthController{
		Services: services,
	}
	managementController := &controller.ManagementController{
		Services: services,
	}
	commonController := &controller.CommonController{
		Services: services,
	}
	controllers := &controller.Controllers{
		Auth:       authController,
		Management: managementController,
		Common:     commonController,
	}
	engine := application.SetupApplication(controllers)
	return engine, nil
}
