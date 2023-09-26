// Code generated by ent, DO NOT EDIT.

package college

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the college type in the database.
	Label = "college"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsDisabled holds the string denoting the is_disabled field in the database.
	FieldIsDisabled = "is_disabled"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldModifiedTime holds the string denoting the modified_time field in the database.
	FieldModifiedTime = "modified_time"
	// EdgeSchool holds the string denoting the school edge name in mutations.
	EdgeSchool = "school"
	// EdgeClasses holds the string denoting the classes edge name in mutations.
	EdgeClasses = "classes"
	// Table holds the table name of the college in the database.
	Table = "college"
	// SchoolTable is the table that holds the school relation/edge.
	SchoolTable = "college"
	// SchoolInverseTable is the table name for the School entity.
	// It exists in this package in order to avoid circular dependency with the "school" package.
	SchoolInverseTable = "school"
	// SchoolColumn is the table column denoting the school relation/edge.
	SchoolColumn = "sid"
	// ClassesTable is the table that holds the classes relation/edge.
	ClassesTable = "class"
	// ClassesInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassesInverseTable = "class"
	// ClassesColumn is the table column denoting the classes relation/edge.
	ClassesColumn = "cid"
)

// Columns holds all SQL columns for college fields.
var Columns = []string{
	FieldID,
	FieldSid,
	FieldName,
	FieldIsDisabled,
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
	// DefaultIsDisabled holds the default value on creation for the "is_disabled" field.
	DefaultIsDisabled bool
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime func() time.Time
)

// OrderOption defines the ordering options for the College queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySid orders the results by the sid field.
func BySid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSid, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIsDisabled orders the results by the is_disabled field.
func ByIsDisabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDisabled, opts...).ToFunc()
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

// BySchoolField orders the results by school field.
func BySchoolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSchoolStep(), sql.OrderByField(field, opts...))
	}
}

// ByClassesCount orders the results by classes count.
func ByClassesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClassesStep(), opts...)
	}
}

// ByClasses orders the results by classes terms.
func ByClasses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSchoolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SchoolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SchoolTable, SchoolColumn),
	)
}
func newClassesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
	)
}
