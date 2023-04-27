package dao

import (
	"errors"
	"time"

	"github.com/DemoonLXW/up_learning/entity"
	"gorm.io/gorm"
)

type PermissionDao struct {
	DB *gorm.DB
}

func (dao *PermissionDao) CreatePermission(permission *entity.Permission) error {
	var num int64
	if err := dao.DB.Table("permission").Select("count(id)").Count(&num).Error; err != nil {
		return err
	}

	var check entity.Permission
	if dao.DB.Where("action = ?", permission.Action).First(&check).RowsAffected != 0 {
		return errors.New("repeat action to insert")
	}

	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	permission.ID = uint16(num + 1)
	permission.CreateTime = time.Now()
	if err := tx.Omit("deleted_time", "modified_time").Create(permission).Error; err != nil {
		tx.Rollback()
		return errors.Join(errors.New("delete fail"), err)
	}

	return tx.Commit().Error

}

func (dao *PermissionDao) UpdatePermission(permission *entity.Permission) error {
	var check entity.Permission
	if err := dao.DB.Where("id = ?", permission.ID).First(&check).Error; err != nil {
		return errors.Join(errors.New("permission can not be found"), err)
	}
	if check.DeletedTime != time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local) {
		return errors.New("permission is deleted")
	}

	tx := dao.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	permission.ModifiedTime = time.Now()
	if err := tx.Model(permission).Omit("id", "created_time", "deleted_time").Updates(permission).Error; err != nil {
		tx.Rollback()
		return errors.Join(errors.New("update permission fail"), err)
	}

	return tx.Commit().Error
}

func (dao *PermissionDao) RetrievePermission(current, pageSize int, like, order string) ([]*entity.Permission, error) {
	var permissions []entity.Permission
	like = "%" + like + "%"
	offset := (current - 1) * pageSize

	queryWithoutOrder := func(db *gorm.DB) *gorm.DB {
		deletedTime := time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)

		return db.Where("deleted_time = ?", deletedTime).Where("action like ? or description like ?", like, like).
			Limit(pageSize).Offset(offset)
	}
	queryWithOrder := func(db *gorm.DB) *gorm.DB {
		if order != "" {
			return db.Order(order)
		}
		return db
	}

	dao.DB.Scopes(queryWithoutOrder, queryWithOrder).Find(&permissions)
	permissions_pointers := make([]*entity.Permission, len(permissions))
	for i, v := range permissions {
		permissions_pointers[i] = &v
	}
	return permissions_pointers, nil

}
