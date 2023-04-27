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
		Action:      "test insert action",
		Description: "test insert description",
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
		ID:          1,
		Action:      "test update action",
		Description: "test update description",
	}
	err = dao.UpdatePermission(&permission)
	assert.Nil(t, err)

}
