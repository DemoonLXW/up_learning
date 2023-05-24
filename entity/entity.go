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

type Role struct {
	ID           uint8     `gorm:"primaryKey;column:id;"`
	Name         string    `gorm:"column:name;"`
	Description  string    `gorm:"column:description;"`
	CreateTime   time.Time `gorm:"column:created_time;"`
	DeletedTime  time.Time `gorm:"column:deleted_time;"`
	ModifiedTime time.Time `gorm:"column:modified_time;"`
}

func (*Role) TableName() string {
	return "role"
}

type RolePermission struct {
	RoleID       uint8  `gorm:"primaryKey;"`
	PermissionID uint16 `gorm:"primaryKey;"`
}

type PermissionWithRoles struct {
	Permission Permission `gorm:"many2many:role_permission;"`
	Roles      []Role     `gorm:"many2many:role_permission;"`
}

func (*PermissionWithRoles) TableName() string {
	return "role_permission"
}
