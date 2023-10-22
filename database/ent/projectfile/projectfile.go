// Code generated by ent, DO NOT EDIT.

package projectfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the projectfile type in the database.
	Label = "project_file"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldFid holds the string denoting the fid field in the database.
	FieldFid = "fid"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// ProjectFieldID holds the string denoting the ID field of the Project.
	ProjectFieldID = "id"
	// FileFieldID holds the string denoting the ID field of the File.
	FileFieldID = "id"
	// Table holds the table name of the projectfile in the database.
	Table = "project_file"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "project_file"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "project"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "pid"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "project_file"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "file"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "fid"
)

// Columns holds all SQL columns for projectfile fields.
var Columns = []string{
	FieldPid,
	FieldFid,
	FieldCreatedTime,
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
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime func() time.Time
)

// OrderOption defines the ordering options for the ProjectFile queries.
type OrderOption func(*sql.Selector)

// ByPid orders the results by the pid field.
func ByPid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPid, opts...).ToFunc()
}

// ByFid orders the results by the fid field.
func ByFid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFid, opts...).ToFunc()
}

// ByCreatedTime orders the results by the created_time field.
func ByCreatedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedTime, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByFilesField orders the results by files field.
func ByFilesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFilesStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, ProjectColumn),
		sqlgraph.To(ProjectInverseTable, ProjectFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
	)
}
func newFilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FilesColumn),
		sqlgraph.To(FilesInverseTable, FileFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, FilesTable, FilesColumn),
	)
}
