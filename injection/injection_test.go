package injection

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideDataBase(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	client, err := ProvideDataBase()
	defer client.Close()
	assert.Nil(t, err)
}

func TestProvideService(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	service, err := ProvideService()
	assert.Nil(t, err)
	token, err := service.User.CreateAndSaveToken(1)
	assert.Nil(t, err)
	fmt.Println(token)
}
