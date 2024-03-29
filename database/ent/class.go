// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent/class"
	"github.com/DemoonLXW/up_learning/database/ent/major"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Mid holds the value of the "mid" field.
	Mid uint16 `json:"mid,omitempty"`
	// Grade holds the value of the "grade" field.
	Grade string `json:"grade,omitempty"`
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
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges        ClassEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Major holds the value of the major edge.
	Major *Major `json:"major,omitempty"`
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MajorOrErr returns the Major value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) MajorOrErr() (*Major, error) {
	if e.loadedTypes[0] {
		if e.Major == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: major.Label}
		}
		return e.Major, nil
	}
	return nil, &NotLoadedError{edge: "major"}
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[1] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldIsDisabled:
			values[i] = new(sql.NullBool)
		case class.FieldID, class.FieldMid:
			values[i] = new(sql.NullInt64)
		case class.FieldGrade, class.FieldName:
			values[i] = new(sql.NullString)
		case class.FieldCreatedTime, class.FieldDeletedTime, class.FieldModifiedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint32(value.Int64)
		case class.FieldMid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mid", values[i])
			} else if value.Valid {
				c.Mid = uint16(value.Int64)
			}
		case class.FieldGrade:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field grade", values[i])
			} else if value.Valid {
				c.Grade = value.String
			}
		case class.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case class.FieldIsDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_disabled", values[i])
			} else if value.Valid {
				c.IsDisabled = value.Bool
			}
		case class.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				c.CreatedTime = new(time.Time)
				*c.CreatedTime = value.Time
			}
		case class.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				c.DeletedTime = new(time.Time)
				*c.DeletedTime = value.Time
			}
		case class.FieldModifiedTime:
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

// Value returns the ent.Value that was dynamically selected and assigned to the Class.
// This includes values selected through modifiers, order, etc.
func (c *Class) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryMajor queries the "major" edge of the Class entity.
func (c *Class) QueryMajor() *MajorQuery {
	return NewClassClient(c.config).QueryMajor(c)
}

// QueryStudents queries the "students" edge of the Class entity.
func (c *Class) QueryStudents() *StudentQuery {
	return NewClassClient(c.config).QueryStudents(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return NewClassClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("mid=")
	builder.WriteString(fmt.Sprintf("%v", c.Mid))
	builder.WriteString(", ")
	builder.WriteString("grade=")
	builder.WriteString(c.Grade)
	builder.WriteString(", ")
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

// Classes is a parsable slice of Class.
type Classes []*Class
