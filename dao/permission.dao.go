package dao

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/entity"
)

type PermissionDao struct {
	DB *ent.Client
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("rollback failed %w: %v", err, rerr)
	}
	return fmt.Errorf("update permission failed: %w", err)
}

func (dao *PermissionDao) CreatePermission(toCreate *ent.Permission) error {
	ctx := context.Background()

	num, err := dao.DB.Permission.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create permission count failed: %w", err)
	}

	checkRepeatAction, err := dao.DB.Permission.Query().Where(permission.Action(*toCreate.Action)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("create permission check repeat action failed: %w", err)
	}
	if checkRepeatAction {
		return fmt.Errorf("repeat action")
	}

	tx, err := dao.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("create permission start a transaction failed: %w", err)
	}

	err = tx.Permission.Create().
		SetID(uint16(num + 1)).
		SetAction(*toCreate.Action).
		SetDescription(*toCreate.Description).
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("create permission transaction commit failed: %w", err)
	}

	return nil

}

func (dao *PermissionDao) UpdatePermission(toUpdate *ent.Permission) error {
	ctx := context.Background()

	originalPermission, err := dao.DB.Permission.Query().Where(permission.ID(toUpdate.ID)).WithRoles().First(ctx)
	if err != nil {
		return fmt.Errorf("update permission find original failed: %w", err)
	}
	if !originalPermission.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
		return fmt.Errorf("permission is already deleted")
	}
	if len(originalPermission.Edges.Roles) != 0 {
		return fmt.Errorf("permission is already used")
	}

	tx, err := dao.DB.Tx(ctx)
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
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update permission transaction commit failed: %w", err)
	}

	return nil
}

func (dao *PermissionDao) RetrievePermission(current, pageSize int, like, sort, order string) ([]*ent.Permission, error) {
	ctx := context.Background()

	offset := (current - 1) * pageSize

	permissions, err := dao.DB.Permission.Query().
		Where(
			permission.Or(
				permission.ActionContains(like),
				permission.DescriptionContains(like),
			),
		).
		Limit(pageSize).
		Offset(offset).
		Order(func(s *sql.Selector) {
			if order == "asc" {
				s.OrderBy(sql.Asc(sort))
			} else {
				s.OrderBy(sql.Desc(sort))
			}
		}).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieve permission query failed: %w", err)
	}

	return permissions, nil

}

func (dao *PermissionDao) DeletePermission(toDelete *entity.Permission) error {
	ctx := context.Background()

	originalPermission, err := dao.DB.Permission.Query().Where(permission.ID(toDelete.ID)).WithRoles().First(ctx)
	if err != nil {
		return fmt.Errorf("delete permission find original failed: %w", err)
	}
	if !originalPermission.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
		return fmt.Errorf("permission is already deleted")
	}
	if len(originalPermission.Edges.Roles) != 0 {
		return fmt.Errorf("permission is already used")
	}

	tx, err := dao.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("delete permission start a transaction failed: %w", err)
	}

	err = tx.Permission.Update().
		SetDeletedTime(time.Now()).
		Where(permission.ID(toDelete.ID)).
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("delete permission transaction commit failed: %w", err)
	}

	return nil
}
