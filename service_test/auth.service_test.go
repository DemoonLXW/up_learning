package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func CreateTestAuthService() (*service.AuthService, error) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	if err != nil {
		return nil, err
	}
	serv.DB = db
	rd, err := injection.ProvideRedis()
	if err != nil {
		return nil, err
	}
	serv.Redis = rd

	return serv, nil
}

func TestLogin(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	// for i := 0; i < 3; i++ {
	// 	u, token, err := serv.Login("account3", "f8ec997eccf015b232ac2b97992ece6caf28060d95d0cbfa6da803064e941583")
	// 	assert.Nil(t, err)
	// 	fmt.Println(token)
	// 	fmt.Println(u)
	// 	time.Sleep(3 * time.Second)
	// }
	u, token, err := serv.Login("account3", "f8ec997eccf015b232ac2b97992ece6caf28060d95d0cbfa6da803064e941583")
	assert.Nil(t, err)
	fmt.Println(token)
	fmt.Println(u)
	// u, cookie, err := serv.Login("username1", "password1")
	// assert.Nil(t, err)
	// fmt.Println(cookie)
	// fmt.Println(u)
}

func TestLogout(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	token := "cb56c3b6fa1a3322fb36180b27755428"
	err = serv.Logout("1", token)
	assert.Nil(t, err)
}

func TestCheckCredential(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	token := "45fa8fc9f40cc15b0c3421e16f01d1e3"
	new_token, err := serv.CheckCredential(1, token)
	assert.Nil(t, err)
	fmt.Println(new_token)
}

func TestFindOneUserWihtRoles(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	user, err := serv.FindOneUserById(uint32(1))
	assert.Nil(t, err)
	fmt.Println(user)
	fmt.Println(user.Edges.Roles)
}

func TestFindMenusByUserId(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	m, err := serv.FindMenuByUserId(uint32(1))
	assert.Nil(t, err)
	for _, menus := range m {
		for _, v := range menus.JSONMenu {
			fmt.Println(v)
		}

	}
}

func TestFindMenusByRoleIds(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	ids := []uint8{1, 2}
	m, err := serv.FindMenuByRoleIds(ids)
	assert.Nil(t, err)
	for _, menus := range m {
		for _, v := range menus.JSONMenu {
			fmt.Println(v)
		}

	}
}

func TestFindRolesWithMenusAndPermissonsByUserId(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	var id uint32 = 1

	rs, err := serv.FindRolesWithMenusAndPermissonsByUserId(id)
	assert.Nil(t, err)
	for i := range rs {
		fmt.Print(rs[i].Edges.Permissions)
	}
}

func TestCheckPermissions(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.AuthService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	var id uint32 = 1
	var ps []string = []string{"test insert actions for new"}

	has, err := serv.CheckPermissions(id, ps)
	assert.Nil(t, err)
	assert.Equal(t, true, has)
}
