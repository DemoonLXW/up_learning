// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DemoonLXW/up_learning/database/ent/class"
	"github.com/DemoonLXW/up_learning/database/ent/college"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/menu"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/rolepermission"
	"github.com/DemoonLXW/up_learning/database/ent/samplefile"
	"github.com/DemoonLXW/up_learning/database/ent/schema"
	"github.com/DemoonLXW/up_learning/database/ent/school"
	"github.com/DemoonLXW/up_learning/database/ent/student"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/database/ent/userrole"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescIsDisabled is the schema descriptor for is_disabled field.
	classDescIsDisabled := classFields[3].Descriptor()
	// class.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	class.DefaultIsDisabled = classDescIsDisabled.Default.(bool)
	// classDescCreatedTime is the schema descriptor for created_time field.
	classDescCreatedTime := classFields[4].Descriptor()
	// class.DefaultCreatedTime holds the default value on creation for the created_time field.
	class.DefaultCreatedTime = classDescCreatedTime.Default.(func() time.Time)
	// classDescDeletedTime is the schema descriptor for deleted_time field.
	classDescDeletedTime := classFields[5].Descriptor()
	// class.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	class.DefaultDeletedTime = classDescDeletedTime.Default.(time.Time)
	// classDescModifiedTime is the schema descriptor for modified_time field.
	classDescModifiedTime := classFields[6].Descriptor()
	// class.DefaultModifiedTime holds the default value on creation for the modified_time field.
	class.DefaultModifiedTime = classDescModifiedTime.Default.(time.Time)
	collegeFields := schema.College{}.Fields()
	_ = collegeFields
	// collegeDescIsDisabled is the schema descriptor for is_disabled field.
	collegeDescIsDisabled := collegeFields[3].Descriptor()
	// college.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	college.DefaultIsDisabled = collegeDescIsDisabled.Default.(bool)
	// collegeDescCreatedTime is the schema descriptor for created_time field.
	collegeDescCreatedTime := collegeFields[4].Descriptor()
	// college.DefaultCreatedTime holds the default value on creation for the created_time field.
	college.DefaultCreatedTime = collegeDescCreatedTime.Default.(func() time.Time)
	// collegeDescDeletedTime is the schema descriptor for deleted_time field.
	collegeDescDeletedTime := collegeFields[5].Descriptor()
	// college.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	college.DefaultDeletedTime = collegeDescDeletedTime.Default.(time.Time)
	// collegeDescModifiedTime is the schema descriptor for modified_time field.
	collegeDescModifiedTime := collegeFields[6].Descriptor()
	// college.DefaultModifiedTime holds the default value on creation for the modified_time field.
	college.DefaultModifiedTime = collegeDescModifiedTime.Default.(time.Time)
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescName is the schema descriptor for name field.
	fileDescName := fileFields[2].Descriptor()
	// file.NameValidator is a validator for the "name" field. It is called by the builders before save.
	file.NameValidator = fileDescName.Validators[0].(func(string) error)
	// fileDescPath is the schema descriptor for path field.
	fileDescPath := fileFields[3].Descriptor()
	// file.PathValidator is a validator for the "path" field. It is called by the builders before save.
	file.PathValidator = fileDescPath.Validators[0].(func(string) error)
	// fileDescIsDisabled is the schema descriptor for is_disabled field.
	fileDescIsDisabled := fileFields[5].Descriptor()
	// file.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	file.DefaultIsDisabled = fileDescIsDisabled.Default.(bool)
	// fileDescCreatedTime is the schema descriptor for created_time field.
	fileDescCreatedTime := fileFields[6].Descriptor()
	// file.DefaultCreatedTime holds the default value on creation for the created_time field.
	file.DefaultCreatedTime = fileDescCreatedTime.Default.(func() time.Time)
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescName is the schema descriptor for name field.
	menuDescName := menuFields[1].Descriptor()
	// menu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	menu.NameValidator = menuDescName.Validators[0].(func(string) error)
	// menuDescCreatedTime is the schema descriptor for created_time field.
	menuDescCreatedTime := menuFields[4].Descriptor()
	// menu.DefaultCreatedTime holds the default value on creation for the created_time field.
	menu.DefaultCreatedTime = menuDescCreatedTime.Default.(func() time.Time)
	// menuDescDeletedTime is the schema descriptor for deleted_time field.
	menuDescDeletedTime := menuFields[5].Descriptor()
	// menu.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	menu.DefaultDeletedTime = menuDescDeletedTime.Default.(time.Time)
	// menuDescModifiedTime is the schema descriptor for modified_time field.
	menuDescModifiedTime := menuFields[6].Descriptor()
	// menu.DefaultModifiedTime holds the default value on creation for the modified_time field.
	menu.DefaultModifiedTime = menuDescModifiedTime.Default.(time.Time)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescAction is the schema descriptor for action field.
	permissionDescAction := permissionFields[1].Descriptor()
	// permission.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	permission.ActionValidator = permissionDescAction.Validators[0].(func(string) error)
	// permissionDescIsDisabled is the schema descriptor for is_disabled field.
	permissionDescIsDisabled := permissionFields[3].Descriptor()
	// permission.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	permission.DefaultIsDisabled = permissionDescIsDisabled.Default.(bool)
	// permissionDescCreatedTime is the schema descriptor for created_time field.
	permissionDescCreatedTime := permissionFields[4].Descriptor()
	// permission.DefaultCreatedTime holds the default value on creation for the created_time field.
	permission.DefaultCreatedTime = permissionDescCreatedTime.Default.(func() time.Time)
	// permissionDescDeletedTime is the schema descriptor for deleted_time field.
	permissionDescDeletedTime := permissionFields[5].Descriptor()
	// permission.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	permission.DefaultDeletedTime = permissionDescDeletedTime.Default.(time.Time)
	// permissionDescModifiedTime is the schema descriptor for modified_time field.
	permissionDescModifiedTime := permissionFields[6].Descriptor()
	// permission.DefaultModifiedTime holds the default value on creation for the modified_time field.
	permission.DefaultModifiedTime = permissionDescModifiedTime.Default.(time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[1].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescIsDisabled is the schema descriptor for is_disabled field.
	roleDescIsDisabled := roleFields[3].Descriptor()
	// role.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	role.DefaultIsDisabled = roleDescIsDisabled.Default.(bool)
	// roleDescCreatedTime is the schema descriptor for created_time field.
	roleDescCreatedTime := roleFields[4].Descriptor()
	// role.DefaultCreatedTime holds the default value on creation for the created_time field.
	role.DefaultCreatedTime = roleDescCreatedTime.Default.(func() time.Time)
	// roleDescModifiedTime is the schema descriptor for modified_time field.
	roleDescModifiedTime := roleFields[5].Descriptor()
	// role.DefaultModifiedTime holds the default value on creation for the modified_time field.
	role.DefaultModifiedTime = roleDescModifiedTime.Default.(time.Time)
	// roleDescDeletedTime is the schema descriptor for deleted_time field.
	roleDescDeletedTime := roleFields[6].Descriptor()
	// role.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	role.DefaultDeletedTime = roleDescDeletedTime.Default.(time.Time)
	rolepermissionFields := schema.RolePermission{}.Fields()
	_ = rolepermissionFields
	// rolepermissionDescCreatedTime is the schema descriptor for created_time field.
	rolepermissionDescCreatedTime := rolepermissionFields[2].Descriptor()
	// rolepermission.DefaultCreatedTime holds the default value on creation for the created_time field.
	rolepermission.DefaultCreatedTime = rolepermissionDescCreatedTime.Default.(func() time.Time)
	samplefileFields := schema.SampleFile{}.Fields()
	_ = samplefileFields
	// samplefileDescIsDisabled is the schema descriptor for is_disabled field.
	samplefileDescIsDisabled := samplefileFields[3].Descriptor()
	// samplefile.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	samplefile.DefaultIsDisabled = samplefileDescIsDisabled.Default.(bool)
	// samplefileDescCreatedTime is the schema descriptor for created_time field.
	samplefileDescCreatedTime := samplefileFields[4].Descriptor()
	// samplefile.DefaultCreatedTime holds the default value on creation for the created_time field.
	samplefile.DefaultCreatedTime = samplefileDescCreatedTime.Default.(func() time.Time)
	schoolFields := schema.School{}.Fields()
	_ = schoolFields
	// schoolDescIsDisabled is the schema descriptor for is_disabled field.
	schoolDescIsDisabled := schoolFields[7].Descriptor()
	// school.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	school.DefaultIsDisabled = schoolDescIsDisabled.Default.(bool)
	// schoolDescCreatedTime is the schema descriptor for created_time field.
	schoolDescCreatedTime := schoolFields[8].Descriptor()
	// school.DefaultCreatedTime holds the default value on creation for the created_time field.
	school.DefaultCreatedTime = schoolDescCreatedTime.Default.(func() time.Time)
	// schoolDescDeletedTime is the schema descriptor for deleted_time field.
	schoolDescDeletedTime := schoolFields[9].Descriptor()
	// school.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	school.DefaultDeletedTime = schoolDescDeletedTime.Default.(time.Time)
	// schoolDescModifiedTime is the schema descriptor for modified_time field.
	schoolDescModifiedTime := schoolFields[10].Descriptor()
	// school.DefaultModifiedTime holds the default value on creation for the modified_time field.
	school.DefaultModifiedTime = schoolDescModifiedTime.Default.(time.Time)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescStudentID is the schema descriptor for student_id field.
	studentDescStudentID := studentFields[3].Descriptor()
	// student.StudentIDValidator is a validator for the "student_id" field. It is called by the builders before save.
	student.StudentIDValidator = studentDescStudentID.Validators[0].(func(string) error)
	// studentDescName is the schema descriptor for name field.
	studentDescName := studentFields[4].Descriptor()
	// student.NameValidator is a validator for the "name" field. It is called by the builders before save.
	student.NameValidator = studentDescName.Validators[0].(func(string) error)
	// studentDescIsDisabled is the schema descriptor for is_disabled field.
	studentDescIsDisabled := studentFields[9].Descriptor()
	// student.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	student.DefaultIsDisabled = studentDescIsDisabled.Default.(bool)
	// studentDescCreatedTime is the schema descriptor for created_time field.
	studentDescCreatedTime := studentFields[10].Descriptor()
	// student.DefaultCreatedTime holds the default value on creation for the created_time field.
	student.DefaultCreatedTime = studentDescCreatedTime.Default.(func() time.Time)
	// studentDescDeletedTime is the schema descriptor for deleted_time field.
	studentDescDeletedTime := studentFields[11].Descriptor()
	// student.DefaultDeletedTime holds the default value on creation for the deleted_time field.
	student.DefaultDeletedTime = studentDescDeletedTime.Default.(time.Time)
	// studentDescModifiedTime is the schema descriptor for modified_time field.
	studentDescModifiedTime := studentFields[12].Descriptor()
	// student.DefaultModifiedTime holds the default value on creation for the modified_time field.
	student.DefaultModifiedTime = studentDescModifiedTime.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAccount is the schema descriptor for account field.
	userDescAccount := userFields[1].Descriptor()
	// user.AccountValidator is a validator for the "account" field. It is called by the builders before save.
	user.AccountValidator = userDescAccount.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescIsDisabled is the schema descriptor for is_disabled field.
	userDescIsDisabled := userFields[7].Descriptor()
	// user.DefaultIsDisabled holds the default value on creation for the is_disabled field.
	user.DefaultIsDisabled = userDescIsDisabled.Default.(bool)
	// userDescCreatedTime is the schema descriptor for created_time field.
	userDescCreatedTime := userFields[8].Descriptor()
	// user.DefaultCreatedTime holds the default value on creation for the created_time field.
	user.DefaultCreatedTime = userDescCreatedTime.Default.(func() time.Time)
	userroleFields := schema.UserRole{}.Fields()
	_ = userroleFields
	// userroleDescCreatedTime is the schema descriptor for created_time field.
	userroleDescCreatedTime := userroleFields[2].Descriptor()
	// userrole.DefaultCreatedTime holds the default value on creation for the created_time field.
	userrole.DefaultCreatedTime = userroleDescCreatedTime.Default.(func() time.Time)
}
