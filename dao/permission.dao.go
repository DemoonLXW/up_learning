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
	if !check.DeletedTime.IsZero() {
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
