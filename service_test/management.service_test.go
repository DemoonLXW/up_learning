package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/entity"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestManagementService() (*service.ManagementService, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db
	return serv, nil
}

func TestCreatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	action1 := "action11"
	description1 := "action11 description"
	permission1 := entity.ToAddPermission{
		Action:      &action1,
		Description: &description1,
	}

	action2 := "action12"
	description2 := "action12 description"
	permission2 := entity.ToAddPermission{
		Action:      &action2,
		Description: &description2,
	}
	err = serv.CreatePermission([]*entity.ToAddPermission{&permission1, &permission2})
	// err = serv.CreatePermission([]*ent.Permission{&permission1})
	assert.Nil(t, err)

}

func TestUpdatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	// action := "test update permssion in service"
	// description := "come"
	// disabled := true
	permission := entity.ToModifyPermission{
		ID: 12,
		// Action:      &action,
		// Description: &description,
		// IsDisabled:  &disabled,
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

	IDs := []uint16{9, 10, 11}
	err = dao.DeletePermission(IDs)
	assert.Nil(t, err)

}

func TestCreateRole(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	name1 := "new role2"
	description1 := "new role1 description"
	role1 := entity.ToAddRole{
		Name:        &name1,
		Description: &description1,
	}

	name2 := "new role1"
	description2 := "new role2 description"
	role2 := entity.ToAddRole{
		Name:        &name2,
		Description: &description2,
	}

	name3 := "new role3"
	description3 := "new role3 description"
	role3 := entity.ToAddRole{
		Name:        &name3,
		Description: &description3,
	}

	name4 := "new role4"
	description4 := "new role4 description"
	role4 := entity.ToAddRole{
		Name:        &name4,
		Description: &description4,
	}
	err = serv.CreateRole([]*entity.ToAddRole{&role1, &role2, &role3, &role4})
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
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	IDs := []uint8{10, 12}
	err = serv.DeleteRole(IDs)
	assert.Nil(t, err)

}

func TestFindOneRoleById(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	role, err := serv.FindOneRoleById(5)
	assert.Nil(t, err)
	fmt.Println(role)

}

func TestUpdateDeletedRole(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

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
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	id, err := serv.FindADeletedRoleID()
	assert.Nil(t, err)
	fmt.Println(id)
}

func TestGetTotalRetrievedPermissions(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	delete := false
	total, err := serv.GetTotalRetrievedPermissions("created", &delete)
	assert.Nil(t, err)
	assert.Equal(t, 1, total)
}

func TestUpdateDeletedPermission(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	action := "test1 modify deleted permission"
	description := "deleted permission description"
	isDisabled := false
	p := &entity.ToModifyPermission{
		ID:          8,
		Action:      &action,
		Description: &description,
		IsDisabled:  &isDisabled,
	}
	err = serv.UpdateDeletedPermission(p)
	assert.Nil(t, err)
}

func TestFindADeletedPermissionID(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	id, err := serv.FindADeletedPermissionID()
	assert.Nil(t, err)
	fmt.Println(id)
}

func TestFindOnePermissionById(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	role, err := serv.FindOnePermissionById(3)
	assert.Nil(t, err)
	fmt.Println(role)

}

func TestFindPermissionsByRoleId(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	ps, err := serv.FindPermissionsByRoleId(2)
	assert.Nil(t, err)
	for _, v := range ps {
		fmt.Println(v)
	}
}

func TestUpdatePermissionForRole(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	rids := []uint8{2, 3}
	pids := []uint16{1, 2}
	isDeleted := true

	err = serv.UpdatePermissionForRole(rids, pids, isDeleted)
	assert.Nil(t, err)
}

func TestRetrieveUser(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	order := false
	disabled := false
	users, err := serv.RetrieveUser(1, 10, "", "id", &order, &disabled)
	assert.Nil(t, err)
	for _, v := range users {
		fmt.Println(v)
	}
	assert.Equal(t, len(users), 1)
}

func TestGetTotalRetrievedUsers(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// delete := false
	total, err := serv.GetTotalRetrievedUsers("", nil)
	assert.Nil(t, err)
	assert.Equal(t, 3, total)
}
