package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBaseConfig map[string]interface{}

func ProvideDatabaseConfig() (DataBaseConfig, error) {
	filepath, ok := os.LookupEnv("DB_CONFIG")
	if !ok {
		filepath = "./database.config.json"
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config DataBaseConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func ProvideDialector(config DataBaseConfig) (gorm.Dialector, error) {
	dsn, ok := config["dsn"]
	if !ok {
		return nil, errors.New("DSN no found")
	}

	sqlDB, err := sql.Open("mysql", dsn.(string))
	if err != nil {
		return nil, err
	}

	MaxIdleConns, ok := config["MaxIdleConns"]
	if ok {
		sqlDB.SetMaxIdleConns(int(MaxIdleConns.(float64)))
	}
	MaxOpenConns, ok := config["MaxOpenConns"]
	if ok {
		sqlDB.SetMaxOpenConns(int(MaxOpenConns.(float64)))
	}

	dialector := mysql.New(mysql.Config{
		Conn: sqlDB,
	})

	return dialector, nil
}

func ProvideOptions(config DataBaseConfig) (*gorm.Config, error) {
	var options gorm.Config
	SkipDefaultTransaction, ok := config["SkipDefaultTransaction"]
	if ok {
		options.SkipDefaultTransaction = SkipDefaultTransaction.(bool)
	}
	return &options, nil
}
