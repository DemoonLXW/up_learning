// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent/college"
)

// College is the model entity for the College schema.
type College struct {
	config `json:"-"`
	// ID of the ent.
	ID uint8 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsDisabled holds the value of the "is_disabled" field.
	IsDisabled bool `json:"is_disabled,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// DeletedTime holds the value of the "deleted_time" field.
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	// ModifiedTime holds the value of the "modified_time" field.
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CollegeQuery when eager-loading is set.
	Edges        CollegeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CollegeEdges holds the relations/edges for other nodes in the graph.
type CollegeEdges struct {
	// Majors holds the value of the majors edge.
	Majors []*Major `json:"majors,omitempty"`
	// Teachers holds the value of the teachers edge.
	Teachers []*Teacher `json:"teachers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MajorsOrErr returns the Majors value or an error if the edge
// was not loaded in eager-loading.
func (e CollegeEdges) MajorsOrErr() ([]*Major, error) {
	if e.loadedTypes[0] {
		return e.Majors, nil
	}
	return nil, &NotLoadedError{edge: "majors"}
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading.
func (e CollegeEdges) TeachersOrErr() ([]*Teacher, error) {
	if e.loadedTypes[1] {
		return e.Teachers, nil
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*College) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case college.FieldIsDisabled:
			values[i] = new(sql.NullBool)
		case college.FieldID:
			values[i] = new(sql.NullInt64)
		case college.FieldName:
			values[i] = new(sql.NullString)
		case college.FieldCreatedTime, college.FieldDeletedTime, college.FieldModifiedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the College fields.
func (c *College) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case college.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint8(value.Int64)
		case college.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case college.FieldIsDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_disabled", values[i])
			} else if value.Valid {
				c.IsDisabled = value.Bool
			}
		case college.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				c.CreatedTime = new(time.Time)
				*c.CreatedTime = value.Time
			}
		case college.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				c.DeletedTime = new(time.Time)
				*c.DeletedTime = value.Time
			}
		case college.FieldModifiedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_time", values[i])
			} else if value.Valid {
				c.ModifiedTime = new(time.Time)
				*c.ModifiedTime = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the College.
// This includes values selected through modifiers, order, etc.
func (c *College) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryMajors queries the "majors" edge of the College entity.
func (c *College) QueryMajors() *MajorQuery {
	return NewCollegeClient(c.config).QueryMajors(c)
}

// QueryTeachers queries the "teachers" edge of the College entity.
func (c *College) QueryTeachers() *TeacherQuery {
	return NewCollegeClient(c.config).QueryTeachers(c)
}

// Update returns a builder for updating this College.
// Note that you need to call College.Unwrap() before calling this method if this College
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *College) Update() *CollegeUpdateOne {
	return NewCollegeClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the College entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *College) Unwrap() *College {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: College is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *College) String() string {
	var builder strings.Builder
	builder.WriteString("College(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("is_disabled=")
	builder.WriteString(fmt.Sprintf("%v", c.IsDisabled))
	builder.WriteString(", ")
	if v := c.CreatedTime; v != nil {
		builder.WriteString("created_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.ModifiedTime; v != nil {
		builder.WriteString("modified_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Colleges is a parsable slice of College.
type Colleges []*College
