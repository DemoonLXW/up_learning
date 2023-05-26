package dao

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
)

type PermissionDao struct {
	DB *ent.Client
}

func rollback(tx *ent.Tx, description string, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		return fmt.Errorf("%s rollback failed %w: %v", description, err, rerr)
	}
	return fmt.Errorf("%s failed: %w", description, err)
}

func (dao *PermissionDao) CreatePermission(toCreates []*ent.Permission) error {
	ctx := context.Background()

	num, err := dao.DB.Permission.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create permissions count failed: %w", err)
	}

	actions := make([]string, len(toCreates))
	bulk := make([]*ent.PermissionCreate, len(toCreates))
	for i, v := range toCreates {
		actions[i] = *v.Action
		bulk[i] = dao.DB.Permission.Create().SetID(uint16(num + i + 1)).SetAction(*v.Action).SetDescription(*v.Description)
	}

	checkRepeatAction, err := dao.DB.Permission.Query().Where(permission.ActionIn(actions...)).Exist(ctx)
	if err != nil {
		return fmt.Errorf("create permissions check repeat action failed: %w", err)
	}
	if checkRepeatAction {
		return fmt.Errorf("repeat action")
	}

	tx, err := dao.DB.Tx(ctx)
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
		return rollback(tx, "update permission", err)
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

func (dao *PermissionDao) DeletePermission(toDeleteIDs []uint16) error {
	ctx := context.Background()

	originalPermissions, err := dao.DB.Permission.Query().Where(permission.IDIn(toDeleteIDs...)).WithRoles().All(ctx)
	if err != nil {
		return fmt.Errorf("delete permissions find original failed: %w", err)
	}
	if len(originalPermissions) == 0 {
		return fmt.Errorf("delete permissions are not found")
	}
	for _, v := range originalPermissions {
		if !v.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
			return fmt.Errorf("permission (id:%d) is already deleted", v.ID)
		}
		if len(v.Edges.Roles) != 0 {
			return fmt.Errorf("permission (id:%d) is already used", v.ID)
		}
	}

	tx, err := dao.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("delete permissions start a transaction failed: %w", err)
	}

	err = tx.Permission.Update().
		SetDeletedTime(time.Now()).
		Where(permission.IDIn(toDeleteIDs...)).
		Exec(ctx)
	if err != nil {
		return rollback(tx, "delete permissions", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("delete permissions transaction commit failed: %w", err)
	}

	return nil
}
