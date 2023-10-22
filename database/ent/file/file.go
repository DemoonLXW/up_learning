// Code generated by ent, DO NOT EDIT.

package file

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldIsDisabled holds the string denoting the is_disabled field in the database.
	FieldIsDisabled = "is_disabled"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldDeletedTime holds the string denoting the deleted_time field in the database.
	FieldDeletedTime = "deleted_time"
	// FieldModifiedTime holds the string denoting the modified_time field in the database.
	FieldModifiedTime = "modified_time"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeSample holds the string denoting the sample edge name in mutations.
	EdgeSample = "sample"
	// EdgeProjectFile holds the string denoting the project_file edge name in mutations.
	EdgeProjectFile = "project_file"
	// Table holds the table name of the file in the database.
	Table = "file"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "file"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "user"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "uid"
	// ProjectTable is the table that holds the project relation/edge. The primary key declared below.
	ProjectTable = "project_file"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "project"
	// SampleTable is the table that holds the sample relation/edge.
	SampleTable = "sample_file"
	// SampleInverseTable is the table name for the SampleFile entity.
	// It exists in this package in order to avoid circular dependency with the "samplefile" package.
	SampleInverseTable = "sample_file"
	// SampleColumn is the table column denoting the sample relation/edge.
	SampleColumn = "fid"
	// ProjectFileTable is the table that holds the project_file relation/edge.
	ProjectFileTable = "project_file"
	// ProjectFileInverseTable is the table name for the ProjectFile entity.
	// It exists in this package in order to avoid circular dependency with the "projectfile" package.
	ProjectFileInverseTable = "project_file"
	// ProjectFileColumn is the table column denoting the project_file relation/edge.
	ProjectFileColumn = "fid"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldName,
	FieldPath,
	FieldSize,
	FieldIsDisabled,
	FieldCreatedTime,
	FieldDeletedTime,
	FieldModifiedTime,
}

var (
	// ProjectPrimaryKey and ProjectColumn2 are the table columns denoting the
	// primary key for the project relation (M2M).
	ProjectPrimaryKey = []string{"pid", "fid"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// DefaultIsDisabled holds the default value on creation for the "is_disabled" field.
	DefaultIsDisabled bool
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime func() time.Time
)

// OrderOption defines the ordering options for the File queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUID orders the results by the uid field.
func ByUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
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

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectCount orders the results by project count.
func ByProjectCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProjectStep(), opts...)
	}
}

// ByProject orders the results by project terms.
func ByProject(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySampleField orders the results by sample field.
func BySampleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSampleStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectFileCount orders the results by project_file count.
func ByProjectFileCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProjectFileStep(), opts...)
	}
}

// ByProjectFile orders the results by project_file terms.
func ByProjectFile(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectFileStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
	)
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProjectTable, ProjectPrimaryKey...),
	)
}
func newSampleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SampleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SampleTable, SampleColumn),
	)
}
func newProjectFileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectFileInverseTable, ProjectFileColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, ProjectFileTable, ProjectFileColumn),
	)
}
