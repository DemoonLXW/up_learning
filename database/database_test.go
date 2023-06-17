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
	assert.Equal(t, "learn:123456@tcp(106.52.123.58:3306)/up_learning_lxw?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true", config["dsn"].(string))
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
	assert.Equal(t, "redis://lxw:123456@106.52.123.58:6379/0", config["url"].(string))
}

func TestProvideRedis(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	client, err := ProvideRedis(config)
	defer client.Close()
	assert.Nil(t, err)

}
