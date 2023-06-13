// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PermissionColumns holds the columns for the "permission" table.
	PermissionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint16, Increment: true},
		{Name: "action", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime},
		{Name: "modified_time", Type: field.TypeTime},
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
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime},
		{Name: "modified_time", Type: field.TypeTime},
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
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "introduction", Type: field.TypeString, Nullable: true},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "deleted_time", Type: field.TypeTime},
		{Name: "modified_time", Type: field.TypeTime},
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
		PermissionTable,
		RoleTable,
		RolePermissionTable,
		UserTable,
		UserRoleTable,
	}
)

func init() {
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
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	UserRoleTable.ForeignKeys[0].RefTable = UserTable
	UserRoleTable.ForeignKeys[1].RefTable = RoleTable
	UserRoleTable.Annotation = &entsql.Annotation{
		Table: "user_role",
	}
}
