package service

import (
	"fmt"
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

func TestUpdatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	action := "test update permssion in service"
	description := "test update description in service"
	permission := ent.Permission{
		ID:          9,
		Action:      &action,
		Description: &description,
	}
	err = serv.UpdatePermission(&permission)
	assert.Nil(t, err)

}

func TestRetrievePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	permissons, err := dao.RetrievePermission(3, 3, "", "created_time", "desc")
	assert.Nil(t, err)
	for _, v := range permissons {

		fmt.Println(v)
	}
	assert.Len(t, permissons, 3)
}
