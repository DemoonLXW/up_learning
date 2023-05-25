package dao

import (
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/stretchr/testify/assert"
)

func TestCreatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	action := "test description for ent Nillable"
	description := "test description for Nillable"
	permission := ent.Permission{
		Action:      &action,
		Description: &description,
	}
	err = dao.CreatePermission(&permission)
	assert.Nil(t, err)

}

func TestUpdatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	permission := entity.Permission{
		ID:          3,
		Action:      "test update action for ant again",
		Description: "don't update created time",
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

	permissons, err := dao.RetrievePermission(1, 10, "", "id", "desc")
	assert.Nil(t, err)
	for _, v := range permissons {

		fmt.Println(v)
	}
	assert.Len(t, permissons, 4)
}

func TestDeletePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(PermissionDao)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db
	permission := entity.Permission{}
	permission.ID = 3
	err = dao.DeletePermission(&permission)
	assert.Nil(t, err)

}

func TestManyToMany(t *testing.T) {
	// os.Setenv("DB_CONFIG", "../database.config.json")
	// dao := new(PermissionDao)
	// db, err := injection.ProvideDataBase()
	// assert.Nil(t, err)
	// dao.DB = db

	// var per entity.PermissionWithRoles
	// // dao.DB.SetupJoinTable(&entity.PermissionWithRoles{}, "Permission", &entity.RolePermission{})
	// // dao.DB.SetupJoinTable(&entity.PermissionWithRoles{}, "Roles", &entity.RolePermission{})
	// err = dao.DB.Where("pid = ?", 1).Preload("Permissions").Preload("Roles").Find(&per).Error

	// assert.Nil(t, err)

}
