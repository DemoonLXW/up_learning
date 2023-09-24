package service_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

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
	action3 := "action13"
	description3 := "action13 description"
	permission3 := entity.ToAddPermission{
		Action:      &action3,
		Description: &description3,
	}
	action4 := "action14"
	description4 := "action14 description"
	permission4 := entity.ToAddPermission{
		Action:      &action4,
		Description: &description4,
	}
	err = serv.CreatePermission([]*entity.ToAddPermission{&permission1, &permission2, &permission3, &permission4})
	// err = serv.CreatePermission([]*ent.Permission{&permission1})
	assert.Nil(t, err)

}

func TestUpdatePermission(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.ManagementService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db

	// action := "test update permssion in service hhhhhh"
	// description := "test"
	// disabled := true
	permission := entity.ToModifyPermission{
		ID: 3,
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

	current := 2
	pageSize := 3
	order := true
	disabled := false
	permissons, err := dao.RetrievePermission(&current, &pageSize, "", "id", &order, &disabled)
	assert.Nil(t, err)
	for _, v := range permissons {

		fmt.Println(v)
	}
	assert.Len(t, permissons, 3)
}

func TestDeletePermission(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	IDs := []uint16{9, 11, 13, 14}
	err = serv.DeletePermission(IDs)
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

	current := 1
	pageSize := 5
	order := true
	// disabled := false
	roles, err := dao.RetrieveRole(&current, &pageSize, "", "id", &order, nil)
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

func TestFindPermissionsByRoleIds(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	ps, err := serv.FindPermissionsByRoleIds([]uint8{2, 99})
	assert.Nil(t, err)
	for _, v := range ps {
		fmt.Println(v)
	}
}

func TestUpdatePermissionForRole(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	rids := []uint8{2}
	pids := []uint16{99}

	err = serv.UpdatePermissionForRole(rids, pids)
	assert.Nil(t, err)
}

func TestRetrieveUser(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	current := 3
	pageSize := 2
	order := false
	// disabled := false
	users, err := serv.RetrieveUser(&current, &pageSize, "", "id", &order, nil)
	assert.Nil(t, err)
	for _, v := range users {
		fmt.Println(v)
	}
	assert.Equal(t, len(users), 2)
}

func TestGetTotalRetrievedUsers(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	// delete := false
	total, err := serv.GetTotalRetrievedUsers("", nil)
	assert.Nil(t, err)
	assert.Equal(t, 3, total)
}

func TestCreateUser(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	password := "88888888"

	account1 := "new account7"
	username1 := "new username7"
	user1 := entity.ToAddUser{
		Account:  &account1,
		Username: &username1,
		Password: &password,
	}

	account2 := "new account9"
	username2 := "new username9"
	user2 := entity.ToAddUser{
		Account:  &account2,
		Username: &username2,
		Password: &password,
	}

	// account3 := "new account3"
	// username3 := "new username3"
	// user3 := entity.ToAddUser{
	// 	Account:  &account3,
	// 	Username: &username3,
	// 	Password: &password,
	// }

	// account4 := "new account6"
	// username4 := "new username6"
	// user4 := entity.ToAddUser{
	// 	Account:  &account4,
	// 	Username: &username4,
	// 	Password: &password,
	// }

	err = serv.CreateUser([]*entity.ToAddUser{&user1, &user2}, 3)
	assert.Nil(t, err)

}

func TestSaveImportedFile(t *testing.T) {
	serv, err := CreateTestManagementService()
	assert.Nil(t, err)

	f, err := os.Open("../domain.config.json")
	assert.Nil(t, err)
	defer f.Close()

	b := bytes.Buffer{}
	writer := multipart.NewWriter(&b)
	part, err := writer.CreateFormFile("upload", filepath.Base(f.Name()))
	assert.Nil(t, err)

	_, err = io.Copy(part, f)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)

	req, err := http.NewRequest(http.MethodPost, "localhost:8080", &b)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	_, fh, err := req.FormFile("upload")
	assert.Nil(t, err)

	dir := "../temp/importSchool/"
	prefix := "uio"

	o, err := serv.SaveImportedFile(fh, dir, prefix)
	assert.Nil(t, err)
	fmt.Println(o.Name())
	os.Remove(o.Name())
}
