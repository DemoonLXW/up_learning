package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/injection"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(UserService)
	db, err := injection.ProvideDataBase()
	assert.Nil(t, err)
	serv.DB = db
	rd, err := injection.ProvideRedis()
	assert.Nil(t, err)
	serv.Redis = rd

	// for i := 0; i < 2; i++ {
	// 	cookie, err := serv.Login("username1", "password1")
	// 	assert.Nil(t, err)
	// 	fmt.Println(cookie)
	// 	time.Sleep(3 * time.Second)
	// }
	cookie, err := serv.Login("account3", "password2")
	assert.Nil(t, err)
	fmt.Println(cookie)
}
