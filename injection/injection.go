//go:build wireinject

package injection

import (
	"github.com/DemoonLXW/up_learning/database"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func ProvideDataBase() (*gorm.DB, error) {
	wire.Build(database.DataBaseProvider)
	return &gorm.DB{}, nil
}
