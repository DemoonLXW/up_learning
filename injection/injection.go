//go:build wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/database"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/google/wire"
)

func ProvideDataBase() (*ent.Client, error) {
	wire.Build(database.DataBaseProvider)
	return &ent.Client{}, nil
}
