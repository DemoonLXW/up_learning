// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/rolepermission"
	"github.com/DemoonLXW/up_learning/database/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescAction is the schema descriptor for action field.
	permissionDescAction := permissionFields[1].Descriptor()
	// permission.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	permission.ActionValidator = permissionDescAction.Validators[0].(func(string) error)
	// permissionDescCreatedTime is the schema descriptor for created_time field.
	permissionDescCreatedTime := permissionFields[3].Descriptor()
	// permission.DefaultCreatedTime holds the default value on creation for the created_time field.
	permission.DefaultCreatedTime = permissionDescCreatedTime.Default.(func() time.Time)
	// permissionDescDeletedTime is the schema descriptor for deleted_time field.
	permissionDescDeletedTime := permissionFields[4].Descriptor()
	// permission.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	permission.DefaultDeletedTime = permissionDescDeletedTime.Default.(time.Time)
	// permissionDescModifiedTime is the schema descriptor for modified_time field.
	permissionDescModifiedTime := permissionFields[5].Descriptor()
	// permission.DefaultModifiedTime holds the default value on creation for the modified_time field.
	permission.DefaultModifiedTime = permissionDescModifiedTime.Default.(time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[1].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescCreatedTime is the schema descriptor for created_time field.
	roleDescCreatedTime := roleFields[3].Descriptor()
	// role.DefaultCreatedTime holds the default value on creation for the created_time field.
	role.DefaultCreatedTime = roleDescCreatedTime.Default.(func() time.Time)
	// roleDescDeletedTime is the schema descriptor for deleted_time field.
	roleDescDeletedTime := roleFields[4].Descriptor()
	// role.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	role.DefaultDeletedTime = roleDescDeletedTime.Default.(time.Time)
	// roleDescModifiedTime is the schema descriptor for modified_time field.
	roleDescModifiedTime := roleFields[5].Descriptor()
	// role.DefaultModifiedTime holds the default value on creation for the modified_time field.
	role.DefaultModifiedTime = roleDescModifiedTime.Default.(time.Time)
	rolepermissionFields := schema.RolePermission{}.Fields()
	_ = rolepermissionFields
	// rolepermissionDescCreatedTime is the schema descriptor for created_time field.
	rolepermissionDescCreatedTime := rolepermissionFields[2].Descriptor()
	// rolepermission.DefaultCreatedTime holds the default value on creation for the created_time field.
	rolepermission.DefaultCreatedTime = rolepermissionDescCreatedTime.Default.(func() time.Time)
}
