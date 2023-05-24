package dao

import (
	"context"
	"fmt"

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

func (dao *PermissionDao) CreatePermission(permissionToCreate *entity.Permission) error {
	ctx := context.Background()

	num, err := dao.DB.Permission.Query().Aggregate(ent.Count()).Int(ctx)
	if err != nil {
		return fmt.Errorf("create permission count failed: %w", err)
	}

	checkRepeatAction, err := dao.DB.Permission.Query().Where(permission.Action(permissionToCreate.Action)).Exist(ctx)
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
		SetAction(permissionToCreate.Action).
		SetDescription(permissionToCreate.Description).
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

func (dao *PermissionDao) UpdatePermission(permission *entity.Permission) error {
	// var check entity.Permission
	// if err := dao.DB.Where("id = ?", permission.ID).First(&check).Error; err != nil {
	// 	return errors.Join(errors.New("permission can not be found"), err)
	// }
	// if !check.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
	// 	return errors.New("permission is deleted")
	// }

	// tx := dao.DB.Begin()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()
	// if err := tx.Error; err != nil {
	// 	return err
	// }
	// permission.ModifiedTime = time.Now()
	// if err := tx.Model(permission).Omit("id", "created_time", "deleted_time").Updates(permission).Error; err != nil {
	// 	tx.Rollback()
	// 	return errors.Join(errors.New("update permission fail"), err)
	// }

	// return tx.Commit().Error
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

func (dao *PermissionDao) DeletePermission(permission *entity.Permission) error {
	// var check entity.Permission
	// if err := dao.DB.Where("id = ?", permission.ID).First(&check).Error; err != nil {
	// 	return errors.Join(errors.New("permission can not be found"), err)
	// }
	// if !check.DeletedTime.Equal(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)) {
	// 	return errors.New("permission is already deleted")
	// }

	// var checkIsUsed entity.Permission
	// if dao.DB.Table("role_permission").Where("pid = ?", permission.ID).Limit(1).Find(&checkIsUsed).RowsAffected != 0 {
	// 	return errors.New("permission is alread used")
	// }

	// tx := dao.DB.Begin()
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		tx.Rollback()
	// 	}
	// }()
	// if err := tx.Error; err != nil {
	// 	return err
	// }
	// if err := tx.Model(permission).Update("deleted_time", time.Now()).Error; err != nil {
	// 	tx.Rollback()
	// 	return errors.Join(errors.New("delete permission fail"), err)
	// }

	// return tx.Commit().Error
	return nil
}
