package entity

import "time"

type Permission struct {
	ID           uint16    `gorm:"primaryKey;column:id;"`
	Action       string    `gorm:"column:action;"`
	Description  string    `gorm:"column:description;"`
	CreateTime   time.Time `gorm:"column:created_time;"`
	DeletedTime  time.Time `gorm:"column:deleted_time;"`
	ModifiedTime time.Time `gorm:"column:modified_time;"`
}

func (*Permission) TableName() string {
	return "permission"
}
