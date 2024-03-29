// Code generated by ent, DO NOT EDIT.

package school

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the school type in the database.
	Label = "school"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldCompetentDepartment holds the string denoting the competent_department field in the database.
	FieldCompetentDepartment = "competent_department"
	// FieldEducationLevel holds the string denoting the education_level field in the database.
	FieldEducationLevel = "education_level"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldIsDisabled holds the string denoting the is_disabled field in the database.
	FieldIsDisabled = "is_disabled"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldModifiedTime holds the string denoting the modified_time field in the database.
	FieldModifiedTime = "modified_time"
	// Table holds the table name of the school in the database.
	Table = "school"
)

// Columns holds all SQL columns for school fields.
var Columns = []string{
	FieldID,
	FieldCode,
	FieldName,
	FieldLocation,
	FieldCompetentDepartment,
	FieldEducationLevel,
	FieldRemark,
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

// OrderOption defines the ordering options for the School queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByCompetentDepartment orders the results by the competent_department field.
func ByCompetentDepartment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompetentDepartment, opts...).ToFunc()
}

// ByEducationLevel orders the results by the education_level field.
func ByEducationLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEducationLevel, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
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
