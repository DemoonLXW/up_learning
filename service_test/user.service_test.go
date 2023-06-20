package service_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/DemoonLXW/up_learning/service"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.UserService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	for i := 0; i < 3; i++ {
		u, token, err := serv.Login("username1", "password1")
		assert.Nil(t, err)
		fmt.Println(token)
		fmt.Println(u)
		time.Sleep(3 * time.Second)
	}
	// u, cookie, err := serv.Login("username1", "password1")
	// assert.Nil(t, err)
	// fmt.Println(cookie)
	// fmt.Println(u)
}

func TestLogout(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.UserService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	token := "cb56c3b6fa1a3322fb36180b27755428"
	err = serv.Logout(1, token)
	assert.Nil(t, err)
}

func TestCheckCredential(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(service.UserService)
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
