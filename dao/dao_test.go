package dao

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
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	action1 := "test update actions for new"
	description1 := "test description for delete"
	permission1 := ent.Permission{
		Action:      &action1,
		Description: &description1,
	}

	action2 := "test insert actions for new"
	description2 := "test description for create"
	permission2 := ent.Permission{
		Action:      &action2,
		Description: &description2,
	}
	err = dao.CreatePermission([]*ent.Permission{&permission1, &permission2})
	assert.Nil(t, err)

}

func TestUpdatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	action := "test action for ent update"
	description := "test description for Nillable"
	permission := ent.Permission{
		ID:          5,
		Action:      &action,
		Description: &description,
	}
	err = dao.UpdatePermission(&permission)
	assert.Nil(t, err)

}

func TestRetrievePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	permissons, err := dao.RetrievePermission(1, 3, "", "id", "desc")
	assert.Nil(t, err)
	for _, v := range permissons {

		fmt.Println(v)
	}
	assert.Len(t, permissons, 3)
}

func TestDeletePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	IDs := []uint16{5}
	err = dao.DeletePermission(IDs)
	assert.Nil(t, err)

}
