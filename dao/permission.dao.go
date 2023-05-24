package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/DemoonLXW/up_learning/database/ent"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/entity"
)

type PermissionDao struct {
	DB *ent.Client
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func (dao *PermissionDao) CreatePermission(toCreate *entity.Permission) error {
	ctx := context.Background()

	num, err := dao.DB.Permission.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create permission count failed: %w", err)
	}

	checkRepeatAction, err := dao.DB.Permission.Query().Where(permission.Action(toCreate.Action)).Exist(ctx)
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
		SetAction(toCreate.Action).
		SetDescription(toCreate.Description).
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

func (dao *PermissionDao) UpdatePermission(toUpdate *entity.Permission) error {
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
		SetAction(toUpdate.Action).
		SetDescription(toUpdate.Description).
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

func (dao *PermissionDao) RetrievePermission(current, pageSize int, like, order string) ([]*entity.Permission, error) {
	// var permissions []entity.Permission
	// like = "%" + like + "%"
	// offset := (current - 1) * pageSize

	// queryWithoutOrder := func(db *gorm.DB) *gorm.DB {
	// 	deletedTime := time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)

	// 	return db.Where("deleted_time = ?", deletedTime).Where("action like ? or description like ?", like, like).
	// 		Limit(pageSize).Offset(offset)
	// }
	// queryWithOrder := func(db *gorm.DB) *gorm.DB {
	// 	if order != "" {
	// 		return db.Order(order)
	// 	}
	// 	return db
	// }

	// dao.DB.Scopes(queryWithoutOrder, queryWithOrder).Find(&permissions)
	// permissions_pointers := make([]*entity.Permission, len(permissions))
	// for i, v := range permissions {
	// 	permissions_pointers[i] = &v
	// }
	return make([]*entity.Permission, 3), nil

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
