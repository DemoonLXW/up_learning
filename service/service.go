package service

import (
	"fmt"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/google/wire"
)

var ServiceProvider = wire.NewSet(
	wire.Struct(new(Services), "*"),
	wire.Struct(new(ManagementService), "*"),
	wire.Struct(new(AuthService), "*"),
	wire.Struct(new(CommonService), "*"),
	wire.Struct(new(TeacherService), "*"),
	wire.Struct(new(WorkflowService), "*"),
)

type Services struct {
	Management *ManagementService
	Auth       *AuthService
	Common     *CommonService
	Teacher    *TeacherService
	Workflow   *WorkflowService
}

func rollback(tx *ent.Tx, description string, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%s rollback failed %w: %v", description, err, rerr)
	}
	return fmt.Errorf("%s failed: %w", description, err)
}
