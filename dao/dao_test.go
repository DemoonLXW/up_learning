package dao

import (
	"os"
	"testing"

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

	permission := entity.Permission{
		Action:      "test description for ent",
		Description: "test description for ent",
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
	// os.Setenv("DB_CONFIG", "../database.config.json")
	// dao := new(PermissionDao)
	// db, err := injection.ProvideDataBase()
	// assert.Nil(t, err)
	// dao.DB = db

	// permissons, err := dao.RetrievePermission(2, 10, "action", "modified_time")
	// assert.Nil(t, err)
	// assert.Len(t, permissons, 1)
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
