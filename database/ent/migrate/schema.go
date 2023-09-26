// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClassColumns holds the columns for the "class" table.
	ClassColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
		{Name: "cid", Type: field.TypeUint16},
	}
	// ClassTable holds the schema information for the "class" table.
	ClassTable = &schema.Table{
		Name:       "class",
		Columns:    ClassColumns,
		PrimaryKey: []*schema.Column{ClassColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_college_classes",
				Columns:    []*schema.Column{ClassColumns[6]},
				RefColumns: []*schema.Column{CollegeColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CollegeColumns holds the columns for the "college" table.
	CollegeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint16, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
		{Name: "sid", Type: field.TypeUint16},
	}
	// CollegeTable holds the schema information for the "college" table.
	CollegeTable = &schema.Table{
		Name:       "college",
		Columns:    CollegeColumns,
		PrimaryKey: []*schema.Column{CollegeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "college_school_colleges",
				Columns:    []*schema.Column{CollegeColumns[6]},
				RefColumns: []*schema.Column{SchoolColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FileColumns holds the columns for the "file" table.
	FileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "path", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt64, Nullable: true},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
		{Name: "uid", Type: field.TypeUint32},
	}
	// FileTable holds the schema information for the "file" table.
	FileTable = &schema.Table{
		Name:       "file",
		Columns:    FileColumns,
		PrimaryKey: []*schema.Column{FileColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "file_user_files",
				Columns:    []*schema.Column{FileColumns[8]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MenuColumns holds the columns for the "menu" table.
	MenuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint8, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "json_menu", Type: field.TypeJSON, Nullable: true},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime},
		{Name: "modified_time", Type: field.TypeTime},
		{Name: "rid", Type: field.TypeUint8},
	}
	// MenuTable holds the schema information for the "menu" table.
	MenuTable = &schema.Table{
		Name:       "menu",
		Columns:    MenuColumns,
		PrimaryKey: []*schema.Column{MenuColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menu_role_menu",
				Columns:    []*schema.Column{MenuColumns[6]},
				RefColumns: []*schema.Column{RoleColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PermissionColumns holds the columns for the "permission" table.
	PermissionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint16, Increment: true},
		{Name: "action", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
	}
	// PermissionTable holds the schema information for the "permission" table.
	PermissionTable = &schema.Table{
		Name:       "permission",
		Columns:    PermissionColumns,
		PrimaryKey: []*schema.Column{PermissionColumns[0]},
	}
	// RoleColumns holds the columns for the "role" table.
	RoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint8, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "modified_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime},
	}
	// RoleTable holds the schema information for the "role" table.
	RoleTable = &schema.Table{
		Name:       "role",
		Columns:    RoleColumns,
		PrimaryKey: []*schema.Column{RoleColumns[0]},
	}
	// RolePermissionColumns holds the columns for the "role_permission" table.
	RolePermissionColumns = []*schema.Column{
		{Name: "created_time", Type: field.TypeTime},
		{Name: "rid", Type: field.TypeUint8},
		{Name: "pid", Type: field.TypeUint16},
	}
	// RolePermissionTable holds the schema information for the "role_permission" table.
	RolePermissionTable = &schema.Table{
		Name:       "role_permission",
		Columns:    RolePermissionColumns,
		PrimaryKey: []*schema.Column{RolePermissionColumns[1], RolePermissionColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permission_role_role",
				Columns:    []*schema.Column{RolePermissionColumns[1]},
				RefColumns: []*schema.Column{RoleColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "role_permission_permission_permission",
				Columns:    []*schema.Column{RolePermissionColumns[2]},
				RefColumns: []*schema.Column{PermissionColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SampleFileColumns holds the columns for the "sample_file" table.
	SampleFileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint8, Increment: true},
		{Name: "type", Type: field.TypeString, Unique: true},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
		{Name: "fid", Type: field.TypeUint32, Unique: true},
	}
	// SampleFileTable holds the schema information for the "sample_file" table.
	SampleFileTable = &schema.Table{
		Name:       "sample_file",
		Columns:    SampleFileColumns,
		PrimaryKey: []*schema.Column{SampleFileColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sample_file_file_sample",
				Columns:    []*schema.Column{SampleFileColumns[6]},
				RefColumns: []*schema.Column{FileColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SchoolColumns holds the columns for the "school" table.
	SchoolColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint16, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "location", Type: field.TypeString},
		{Name: "competent_department", Type: field.TypeString},
		{Name: "education_level", Type: field.TypeUint8},
		{Name: "remark", Type: field.TypeString},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
	}
	// SchoolTable holds the schema information for the "school" table.
	SchoolTable = &schema.Table{
		Name:       "school",
		Columns:    SchoolColumns,
		PrimaryKey: []*schema.Column{SchoolColumns[0]},
	}
	// StudentColumns holds the columns for the "student" table.
	StudentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "student_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "gender", Type: field.TypeUint8},
		{Name: "birthday", Type: field.TypeTime},
		{Name: "admission_date", Type: field.TypeTime},
		{Name: "state", Type: field.TypeUint8},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
		{Name: "cid", Type: field.TypeUint32},
		{Name: "uid", Type: field.TypeUint32},
	}
	// StudentTable holds the schema information for the "student" table.
	StudentTable = &schema.Table{
		Name:       "student",
		Columns:    StudentColumns,
		PrimaryKey: []*schema.Column{StudentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_class_students",
				Columns:    []*schema.Column{StudentColumns[11]},
				RefColumns: []*schema.Column{ClassColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "student_user_students",
				Columns:    []*schema.Column{StudentColumns[12]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "introduction", Type: field.TypeString, Nullable: true},
		{Name: "is_disabled", Type: field.TypeBool, Default: false},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime, Nullable: true},
		{Name: "modified_time", Type: field.TypeTime, Nullable: true},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// UserRoleColumns holds the columns for the "user_role" table.
	UserRoleColumns = []*schema.Column{
		{Name: "created_time", Type: field.TypeTime},
		{Name: "uid", Type: field.TypeUint32},
		{Name: "rid", Type: field.TypeUint8},
	}
	// UserRoleTable holds the schema information for the "user_role" table.
	UserRoleTable = &schema.Table{
		Name:       "user_role",
		Columns:    UserRoleColumns,
		PrimaryKey: []*schema.Column{UserRoleColumns[1], UserRoleColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_role_user_user",
				Columns:    []*schema.Column{UserRoleColumns[1]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_role_role_role",
				Columns:    []*schema.Column{UserRoleColumns[2]},
				RefColumns: []*schema.Column{RoleColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassTable,
		CollegeTable,
		FileTable,
		MenuTable,
		PermissionTable,
		RoleTable,
		RolePermissionTable,
		SampleFileTable,
		SchoolTable,
		StudentTable,
		UserTable,
		UserRoleTable,
	}
)

func init() {
	ClassTable.ForeignKeys[0].RefTable = CollegeTable
	ClassTable.Annotation = &entsql.Annotation{
		Table: "class",
	}
	CollegeTable.ForeignKeys[0].RefTable = SchoolTable
	CollegeTable.Annotation = &entsql.Annotation{
		Table: "college",
	}
	FileTable.ForeignKeys[0].RefTable = UserTable
	FileTable.Annotation = &entsql.Annotation{
		Table: "file",
	}
	MenuTable.ForeignKeys[0].RefTable = RoleTable
	MenuTable.Annotation = &entsql.Annotation{
		Table: "menu",
	}
	PermissionTable.Annotation = &entsql.Annotation{
		Table: "permission",
	}
	RoleTable.Annotation = &entsql.Annotation{
		Table: "role",
	}
	RolePermissionTable.ForeignKeys[0].RefTable = RoleTable
	RolePermissionTable.ForeignKeys[1].RefTable = PermissionTable
	RolePermissionTable.Annotation = &entsql.Annotation{
		Table: "role_permission",
	}
	SampleFileTable.ForeignKeys[0].RefTable = FileTable
	SampleFileTable.Annotation = &entsql.Annotation{
		Table: "sample_file",
	}
	SchoolTable.Annotation = &entsql.Annotation{
		Table: "school",
	}
	StudentTable.ForeignKeys[0].RefTable = ClassTable
	StudentTable.ForeignKeys[1].RefTable = UserTable
	StudentTable.Annotation = &entsql.Annotation{
		Table: "student",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	UserRoleTable.ForeignKeys[0].RefTable = UserTable
	UserRoleTable.ForeignKeys[1].RefTable = RoleTable
	UserRoleTable.Annotation = &entsql.Annotation{
		Table: "user_role",
	}
}
