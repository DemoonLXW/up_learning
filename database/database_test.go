package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideDatabaseConfig(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	config = config["mysql"].(map[string]interface{})
	assert.Nil(t, err)
	assert.Equal(t, float64(20), config["MaxIdleConns"].(float64))
	assert.Equal(t, 50., config["MaxOpenConns"].(float64))
	assert.Equal(t, true, config["SkipDefaultTransaction"].(bool))
}

func TestProvideDB(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	client, err := ProvideDB(config)
	defer client.Close()
	assert.Nil(t, err)

}

func TestProvideRedisConfig(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	config = config["redis"].(map[string]interface{})
	assert.Nil(t, err)
}

func TestProvideRedis(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	client, err := ProvideRedis(config)
	defer client.Close()
	assert.Nil(t, err)

}
