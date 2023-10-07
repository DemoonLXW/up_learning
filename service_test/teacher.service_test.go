package service_test

import (
	"os"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
)

func CreateTestTeacherService() (*service.TeacherService, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.TeacherService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db

	return serv, nil
}
