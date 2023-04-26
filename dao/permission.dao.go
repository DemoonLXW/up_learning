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
		return err
	}

	return tx.Commit().Error

}
