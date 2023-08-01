package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func TestCreatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	action1 := "test update actions for service2"
	description1 := "test description for service2"
	permission1 := ent.Permission{
		Action:      action1,
		Description: description1,
	}

	action2 := "test insert actions for new2"
	description2 := "test description for create2"
	permission2 := ent.Permission{
		Action:      action2,
		Description: description2,
	}
	err = serv.CreatePermission([]*ent.Permission{&permission1, &permission2})
	// err = serv.CreatePermission([]*ent.Permission{&permission1})
	assert.Nil(t, err)

}

func TestUpdatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	action := "test update permssion in service"
	description := "test update description in service"
	permission := ent.Permission{
		ID:          9,
		Action:      action,
		Description: description,
	}
	err = serv.UpdatePermission(&permission)
	assert.Nil(t, err)

}

func TestRetrievePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	order := true
	disabled := false
	permissons, err := dao.RetrievePermission(1, 10, "", "action", &order, &disabled)
	assert.Nil(t, err)
	for _, v := range permissons {

		fmt.Println(v)
	}
	assert.Len(t, permissons, 5)
}

func TestDeletePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	IDs := []uint16{8, 9, 10}
	err = dao.DeletePermission(IDs)
	assert.Nil(t, err)

}

func TestCreateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	// name1 := "role1"
	// description1 := "role1 description"
	// role1 := ent.Role{
	// 	Name:        &name1,
	// 	Description: &description1,
	// }

	name2 := "role5"
	description2 := "role5 description"
	role2 := entity.ToAddRole{
		Name:        &name2,
		Description: &description2,
	}
	err = serv.CreateRole([]*entity.ToAddRole{&role2})
	assert.Nil(t, err)

}

func TestUpdateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	name := "test change role name"
	description := "test change role service description go go"
	isDisabled := true
	role := entity.ToModifyRole{
		ID:          5,
		Name:        &name,
		Description: &description,
		IsDisabled:  &isDisabled,
	}
	err = serv.UpdateRole(&role)
	assert.Nil(t, err)

}

func TestRetrieveRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	order := false
	disabled := false
	roles, err := dao.RetrieveRole(1, 5, "", "id", &order, &disabled)
	assert.Nil(t, err)
	for _, v := range roles {
		fmt.Println(v)
	}
	assert.Len(t, roles, 5)
}

func TestGetTotalRetrievedRoles(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	// delete := false
	total, err := serv.GetTotalRetrievedRoles("role", nil)
	assert.Nil(t, err)
	assert.Equal(t, 6, total)
}

func TestDeleteRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	dao := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	dao.DB = db

	IDs := []uint8{4}
	err = dao.DeleteRole(IDs)
	assert.Nil(t, err)

}

func TestCreateUser(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	account1 := "username3"
	password1 := "password3"
	user1 := ent.User{
		Account:  &account1,
		Password: &password1,
	}

	account2 := "account3"
	password2 := "password2"
	user2 := ent.User{
		Account:  &account2,
		Password: &password2,
	}
	err = serv.CreateUser([]*ent.User{&user2, &user1})
	assert.Nil(t, err)

}

func TestFindOneRoleById(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	role, err := serv.FindOneRoleById(5)
	assert.Nil(t, err)
	fmt.Println(role)

}

func TestUpdateDeletedRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	name := "role7"
	description := "test change deleted role description"
	isDisabled := true
	role := &entity.ToModifyRole{
		ID:          7,
		Name:        &name,
		Description: &description,
		IsDisabled:  &isDisabled,
	}
	err = serv.UpdateDeletedRole(role)
	assert.Nil(t, err)
}

func TestFindADeletedRoleID(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	id, err := serv.FindADeletedRoleID()
	assert.Nil(t, err)
	fmt.Println(id)
}

func TestGetTotalRetrievedPermissions(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	delete := false
	total, err := serv.GetTotalRetrievedPermissions("created", &delete)
	assert.Nil(t, err)
	assert.Equal(t, 1, total)
}
