package service

import (
	"fmt"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/google/wire"
)

var ServiceProvider = wire.NewSet(
	wire.Struct(new(Service), "*"),
	wire.Struct(new(ManagementService), "*"),
	wire.Struct(new(UserService), "*"),
)

type Service struct {
	Management *ManagementService
	User       *UserService
}

func rollback(tx *ent.Tx, description string, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%s rollback failed %w: %v", description, err, rerr)
	}
	return fmt.Errorf("%s failed: %w", description, err)
}
