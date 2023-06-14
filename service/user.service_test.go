package service

import (
	"os"
	"testing"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/injection"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	serv := new(ManagementService)
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
