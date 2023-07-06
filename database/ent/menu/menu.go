// Code generated by ent, DO NOT EDIT.

package menu

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the menu type in the database.
	Label = "menu"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldJSONMenu holds the string denoting the json_menu field in the database.
	FieldJSONMenu = "json_menu"
	// FieldRid holds the string denoting the rid field in the database.
	FieldRid = "rid"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldModifiedTime holds the string denoting the modified_time field in the database.
	FieldModifiedTime = "modified_time"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the menu in the database.
	Table = "menu"
	// RoleTable is the table that holds the role relation/edge.
	RoleTable = "menu"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "role"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "rid"
)

// Columns holds all SQL columns for menu fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldJSONMenu,
	FieldRid,
	FieldCreatedTime,
	FieldDeletedTime,
	FieldModifiedTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime func() time.Time
	// DefaultDeletedTime holds the default value on creation for the "deleted_time" field.
	DefaultDeletedTime time.Time
	// DefaultModifiedTime holds the default value on creation for the "modified_time" field.
	DefaultModifiedTime time.Time
)

// OrderOption defines the ordering options for the Menu queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByRid orders the results by the rid field.
func ByRid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRid, opts...).ToFunc()
}

// ByCreatedTime orders the results by the created_time field.
func ByCreatedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedTime, opts...).ToFunc()
}

// ByDeletedTime orders the results by the deleted_time field.
func ByDeletedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedTime, opts...).ToFunc()
}

// ByModifiedTime orders the results by the modified_time field.
func ByModifiedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedTime, opts...).ToFunc()
}

// ByRoleField orders the results by role field.
func ByRoleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleStep(), sql.OrderByField(field, opts...))
	}
}
func newRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RoleTable, RoleColumn),
	)
}
