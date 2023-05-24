package database

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideDatabaseConfig(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	assert.Equal(t, "learn:123456@tcp(106.52.123.58:3306)/up_learning_lxw?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true", config["dsn"].(string))
	assert.Equal(t, float64(20), config["MaxIdleConns"].(float64))
	assert.Equal(t, 50., config["MaxOpenConns"].(float64))
	assert.Equal(t, true, config["SkipDefaultTransaction"].(bool))
	log.Println(config)
}

func TestProvideDB(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	client, err := ProvideDB(config)
	defer client.Close()
	assert.Nil(t, err)

}
