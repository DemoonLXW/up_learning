package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var DataBaseProvider = wire.NewSet(
	ProvideDatabaseConfig,
	ProvideDB,
	ProvideRedis,
)

type DataBaseConfig map[string]interface{}

func ProvideDatabaseConfig() (DataBaseConfig, error) {
	fp, ok := os.LookupEnv("DB_CONFIG")
	if !ok {
		fp = "./database.config.json"
	}
	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read config of database failed: %w", err)
	}

	var config DataBaseConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config of database failed: %w", err)
	}
	return config, nil
}

func ProvideDB(config DataBaseConfig) (*ent.Client, error) {
	config = config["mysql"].(map[string]interface{})
	dsn, ok := config["dsn"]
	if !ok {
		return nil, fmt.Errorf("read dsn in config failed")
	}
	db, err := sql.Open("mysql", dsn.(string))
	if err != nil {
		return nil, fmt.Errorf("open database failed: %w", err)
	}

	MaxIdleConns, ok := config["MaxIdleConns"]
	if ok {
		db.SetMaxIdleConns(int(MaxIdleConns.(float64)))
	}
	MaxOpenConns, ok := config["MaxOpenConns"]
	if ok {
		db.SetMaxOpenConns(int(MaxOpenConns.(float64)))
	}

	drv := entsql.OpenDB("mysql", db)

	return ent.NewClient(ent.Driver(drv)), nil
}

func ProvideRedis(config DataBaseConfig) (*redis.Client, error) {
	config = config["redis"].(map[string]interface{})
	url, ok := config["url"]
	if !ok {
		return nil, fmt.Errorf("read url in config failed")
	}
	opt, err := redis.ParseURL(url.(string))
	if err != nil {
		return nil, fmt.Errorf("open redis failed: %w", err)
	}

	rdb := redis.NewClient(opt)

	return rdb, nil
}
