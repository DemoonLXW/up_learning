// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent/school"
)

// School is the model entity for the School schema.
type School struct {
	config `json:"-"`
	// ID of the ent.
	ID uint16 `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// CompetentDepartment holds the value of the "competent_department" field.
	CompetentDepartment string `json:"competent_department,omitempty"`
	// EducationLevel holds the value of the "education_level" field.
	EducationLevel uint8 `json:"education_level,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// IsDisabled holds the value of the "is_disabled" field.
	IsDisabled bool `json:"is_disabled,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// DeletedTime holds the value of the "deleted_time" field.
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	// ModifiedTime holds the value of the "modified_time" field.
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*School) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case school.FieldIsDisabled:
			values[i] = new(sql.NullBool)
		case school.FieldID, school.FieldEducationLevel:
			values[i] = new(sql.NullInt64)
		case school.FieldCode, school.FieldName, school.FieldLocation, school.FieldCompetentDepartment, school.FieldRemark:
			values[i] = new(sql.NullString)
		case school.FieldCreatedTime, school.FieldDeletedTime, school.FieldModifiedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the School fields.
func (s *School) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case school.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint16(value.Int64)
		case school.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case school.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case school.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				s.Location = value.String
			}
		case school.FieldCompetentDepartment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field competent_department", values[i])
			} else if value.Valid {
				s.CompetentDepartment = value.String
			}
		case school.FieldEducationLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field education_level", values[i])
			} else if value.Valid {
				s.EducationLevel = uint8(value.Int64)
			}
		case school.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				s.Remark = value.String
			}
		case school.FieldIsDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_disabled", values[i])
			} else if value.Valid {
				s.IsDisabled = value.Bool
			}
		case school.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				s.CreatedTime = new(time.Time)
				*s.CreatedTime = value.Time
			}
		case school.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				s.DeletedTime = new(time.Time)
				*s.DeletedTime = value.Time
			}
		case school.FieldModifiedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_time", values[i])
			} else if value.Valid {
				s.ModifiedTime = new(time.Time)
				*s.ModifiedTime = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the School.
// This includes values selected through modifiers, order, etc.
func (s *School) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this School.
// Note that you need to call School.Unwrap() before calling this method if this School
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *School) Update() *SchoolUpdateOne {
	return NewSchoolClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the School entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *School) Unwrap() *School {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: School is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *School) String() string {
	var builder strings.Builder
	builder.WriteString("School(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("code=")
	builder.WriteString(s.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(s.Location)
	builder.WriteString(", ")
	builder.WriteString("competent_department=")
	builder.WriteString(s.CompetentDepartment)
	builder.WriteString(", ")
	builder.WriteString("education_level=")
	builder.WriteString(fmt.Sprintf("%v", s.EducationLevel))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(s.Remark)
	builder.WriteString(", ")
	builder.WriteString("is_disabled=")
	builder.WriteString(fmt.Sprintf("%v", s.IsDisabled))
	builder.WriteString(", ")
	if v := s.CreatedTime; v != nil {
		builder.WriteString("created_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := s.ModifiedTime; v != nil {
		builder.WriteString("modified_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Schools is a parsable slice of School.
type Schools []*School
