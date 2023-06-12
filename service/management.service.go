package service

import (
	"context"
	"fmt"
	"time"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
)

type ManagementService struct {
	DB *ent.Client
}

func rollback(tx *ent.Tx, description string, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%s rollback failed %w: %v", description, err, rerr)
	}
	return fmt.Errorf("%s failed: %w", description, err)
}

func (serv *ManagementService) CreatePermission(toCreates []*ent.Permission) error {
	ctx := context.Background()

	num, err := serv.DB.Permission.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create permissions count failed: %w", err)
	}

	actions := make([]string, len(toCreates))
	bulk := make([]*ent.PermissionCreate, len(toCreates))
	for i, v := range toCreates {
		actions[i] = *v.Action
		bulk[i] = serv.DB.Permission.Create().SetID(uint16(num + i + 1)).SetAction(*v.Action).SetDescription(*v.Description)
	}

	checkRepeatAction, err := serv.DB.Permission.Query().Where(permission.ActionIn(actions...)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("create permissions check repeat action failed: %w", err)
	}
	if checkRepeatAction {
		return fmt.Errorf("repeat action")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create permission start a transaction failed: %w", err)
	}

	err = tx.Permission.CreateBulk(bulk...).Exec(ctx)
	if err != nil {
		return rollback(tx, "create permissions", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create permissions transaction commit failed: %w", err)
	}

	return nil

}

func (serv *ManagementService) UpdatePermission(toUpdate *ent.Permission) error {
	ctx := context.Background()

	originalPermission, err := serv.DB.Permission.Query().Where(permission.ID(toUpdate.ID)).WithRoles().First(ctx)
	if err != nil {
		return fmt.Errorf("update permission find original failed: %w", err)
	}
	if !originalPermission.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
		return fmt.Errorf("permission is already deleted")
	}
	if len(originalPermission.Edges.Roles) != 0 {
		return fmt.Errorf("permission is already used")
	}

	tx, err := serv.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("update permission start a transaction failed: %w", err)
	}

	err = tx.Permission.Create().
		SetID(toUpdate.ID).
		SetAction(*toUpdate.Action).
		SetDescription(*toUpdate.Description).
		SetModifiedTime(time.Now()).
		OnConflict().
		UpdateAction().
		UpdateDescription().
		UpdateModifiedTime().
		Exec(ctx)
	if err != nil {
		return rollback(tx, "update permission", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update permission transaction commit failed: %w", err)
	}

	return nil
}
