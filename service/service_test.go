package service

import (
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/stretchr/testify/assert"
)

func TestCreatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	action1 := "test update actions for service2"
	description1 := "test description for service2"
	permission1 := ent.Permission{
		Action:      &action1,
		Description: &description1,
	}

	action2 := "test insert actions for new2"
	description2 := "test description for create2"
	permission2 := ent.Permission{
		Action:      &action2,
		Description: &description2,
	}
	err = serv.CreatePermission([]*ent.Permission{&permission1, &permission2})
	// err = serv.CreatePermission([]*ent.Permission{&permission1})
	assert.Nil(t, err)

}
