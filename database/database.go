package database

import (
	"encoding/json"
	"os"
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
