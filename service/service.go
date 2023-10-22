package service

import (
	"context"
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
	wire.Struct(new(TeacherFacade), "*"),
	wire.Struct(new(WorkflowService), "*"),
)

type Services struct {
	Management *ManagementService
	Auth       *AuthService
	Common     *CommonService
	Teacher    *TeacherService
	TeacherFa  *TeacherFacade
	Workflow   *WorkflowService
}

func rollback(tx *ent.Tx, description string, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%s rollback failed %w: %v", description, err, rerr)
	}
	return fmt.Errorf("%s failed: %w", description, err)
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
