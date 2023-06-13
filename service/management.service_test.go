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

func TestDeletePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	IDs := []uint16{8, 9, 10}
	err = dao.DeletePermission(IDs)
	assert.Nil(t, err)

}

func TestCreateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	// name1 := "role1"
	// description1 := "role1 description"
	// role1 := ent.Role{
	// 	Name:        &name1,
	// 	Description: &description1,
	// }

	name2 := "role3"
	description2 := "role3 description"
	role2 := ent.Role{
		Name:        &name2,
		Description: &description2,
	}
	err = serv.CreateRole([]*ent.Role{&role2})
	assert.Nil(t, err)

}

func TestUpdateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	name := "role2change"
	description := "role2change description"
	role := ent.Role{
		ID:          2,
		Name:        &name,
		Description: &description,
	}
	err = serv.UpdateRole(&role)
	assert.Nil(t, err)

}

func TestRetrieveRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	roles, err := dao.RetrieveRole(1, 10, "service", "description", "asc")
	assert.Nil(t, err)
	for _, v := range roles {

		fmt.Println(v)
	}
	assert.Len(t, roles, 1)
}