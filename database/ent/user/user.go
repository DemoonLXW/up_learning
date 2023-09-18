// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccount holds the string denoting the account field in the database.
	FieldAccount = "account"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldIntroduction holds the string denoting the introduction field in the database.
	FieldIntroduction = "introduction"
	// FieldIsDisabled holds the string denoting the is_disabled field in the database.
	FieldIsDisabled = "is_disabled"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldModifiedTime holds the string denoting the modified_time field in the database.
	FieldModifiedTime = "modified_time"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeStudents holds the string denoting the students edge name in mutations.
	EdgeStudents = "students"
	// EdgeUserRole holds the string denoting the user_role edge name in mutations.
	EdgeUserRole = "user_role"
	// Table holds the table name of the user in the database.
	Table = "user"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "user_role"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "role"
	// StudentsTable is the table that holds the students relation/edge.
	StudentsTable = "student"
	// StudentsInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentsInverseTable = "student"
	// StudentsColumn is the table column denoting the students relation/edge.
	StudentsColumn = "uid"
	// UserRoleTable is the table that holds the user_role relation/edge.
	UserRoleTable = "user_role"
	// UserRoleInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	UserRoleInverseTable = "user_role"
	// UserRoleColumn is the table column denoting the user_role relation/edge.
	UserRoleColumn = "uid"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAccount,
	FieldPassword,
	FieldUsername,
	FieldEmail,
	FieldPhone,
	FieldIntroduction,
	FieldIsDisabled,
	FieldCreatedTime,
	FieldDeletedTime,
	FieldModifiedTime,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"uid", "rid"}
)

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
	// AccountValidator is a validator for the "account" field. It is called by the builders before save.
	AccountValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultIsDisabled holds the default value on creation for the "is_disabled" field.
	DefaultIsDisabled bool
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime func() time.Time
	// DefaultDeletedTime holds the default value on creation for the "deleted_time" field.
	DefaultDeletedTime time.Time
	// DefaultModifiedTime holds the default value on creation for the "modified_time" field.
	DefaultModifiedTime time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccount orders the results by the account field.
func ByAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccount, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByIntroduction orders the results by the introduction field.
func ByIntroduction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntroduction, opts...).ToFunc()
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

// ByRolesCount orders the results by roles count.
func ByRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesStep(), opts...)
	}
}

// ByRoles orders the results by roles terms.
func ByRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStudentsCount orders the results by students count.
func ByStudentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudentsStep(), opts...)
	}
}

// ByStudents orders the results by students terms.
func ByStudents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserRoleCount orders the results by user_role count.
func ByUserRoleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserRoleStep(), opts...)
	}
}

// ByUserRole orders the results by user_role terms.
func ByUserRole(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserRoleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, RolesTable, RolesPrimaryKey...),
	)
}
func newStudentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudentsTable, StudentsColumn),
	)
}
func newUserRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserRoleInverseTable, UserRoleColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserRoleTable, UserRoleColumn),
	)
}
