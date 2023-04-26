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
	assert.Equal(t, "learn:123456@tcp(106.52.123.58:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true", config["dsn"].(string))
	assert.Equal(t, float64(20), config["MaxIdleConns"].(float64))
	assert.Equal(t, 50., config["MaxOpenConns"].(float64))
	assert.Equal(t, true, config["SkipDefaultTransaction"].(bool))
	log.Println(config)
}

func TestProvideDialector(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	dialector, err := ProvideDialector(config)
	assert.Nil(t, err)
	assert.Equal(t, "mysql", dialector.Name())

}

func TestProvideOptions(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	config, err := ProvideDatabaseConfig()
	assert.Nil(t, err)
	options, err := ProvideOptions(config)
	assert.Nil(t, err)
	assert.Equal(t, options.SkipDefaultTransaction, true)
}
