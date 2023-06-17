//go:build wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func ProvideDataBase() (*ent.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &ent.Client{}, nil
}

func ProvideRedis() (*redis.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &redis.Client{}, nil
}
