// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent/reviewproject"
	"github.com/DemoonLXW/up_learning/database/ent/reviewprojectdetail"
	"github.com/DemoonLXW/up_learning/entity"
)

// ReviewProjectDetail is the model entity for the ReviewProjectDetail schema.
type ReviewProjectDetail struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ReviewProjectID holds the value of the "review_project_id" field.
	ReviewProjectID uint32 `json:"review_project_id,omitempty"`
	// Order holds the value of the "order" field.
	Order uint8 `json:"order,omitempty"`
	// Reviewer holds the value of the "reviewer" field.
	Reviewer *entity.Reviewer `json:"reviewer,omitempty"`
	// Executor holds the value of the "executor" field.
	Executor *entity.Executor `json:"executor,omitempty"`
	// Typee holds the value of the "typee" field.
	Typee uint8 `json:"typee,omitempty"`
	// Status holds the value of the "status" field.
	Status uint8 `json:"status,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// DeletedTime holds the value of the "deleted_time" field.
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	// ModifiedTime holds the value of the "modified_time" field.
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewProjectDetailQuery when eager-loading is set.
	Edges        ReviewProjectDetailEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ReviewProjectDetailEdges holds the relations/edges for other nodes in the graph.
type ReviewProjectDetailEdges struct {
	// ReviewProject holds the value of the review_project edge.
	ReviewProject *ReviewProject `json:"review_project,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReviewProjectOrErr returns the ReviewProject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReviewProjectDetailEdges) ReviewProjectOrErr() (*ReviewProject, error) {
	if e.loadedTypes[0] {
		if e.ReviewProject == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: reviewproject.Label}
		}
		return e.ReviewProject, nil
	}
	return nil, &NotLoadedError{edge: "review_project"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReviewProjectDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case reviewprojectdetail.FieldReviewer, reviewprojectdetail.FieldExecutor:
			values[i] = new([]byte)
		case reviewprojectdetail.FieldID, reviewprojectdetail.FieldReviewProjectID, reviewprojectdetail.FieldOrder, reviewprojectdetail.FieldTypee, reviewprojectdetail.FieldStatus:
			values[i] = new(sql.NullInt64)
		case reviewprojectdetail.FieldCreatedTime, reviewprojectdetail.FieldDeletedTime, reviewprojectdetail.FieldModifiedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReviewProjectDetail fields.
func (rpd *ReviewProjectDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reviewprojectdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rpd.ID = int(value.Int64)
		case reviewprojectdetail.FieldReviewProjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field review_project_id", values[i])
			} else if value.Valid {
				rpd.ReviewProjectID = uint32(value.Int64)
			}
		case reviewprojectdetail.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				rpd.Order = uint8(value.Int64)
			}
		case reviewprojectdetail.FieldReviewer:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field reviewer", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rpd.Reviewer); err != nil {
					return fmt.Errorf("unmarshal field reviewer: %w", err)
				}
			}
		case reviewprojectdetail.FieldExecutor:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field executor", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rpd.Executor); err != nil {
					return fmt.Errorf("unmarshal field executor: %w", err)
				}
			}
		case reviewprojectdetail.FieldTypee:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field typee", values[i])
			} else if value.Valid {
				rpd.Typee = uint8(value.Int64)
			}
		case reviewprojectdetail.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				rpd.Status = uint8(value.Int64)
			}
		case reviewprojectdetail.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				rpd.CreatedTime = new(time.Time)
				*rpd.CreatedTime = value.Time
			}
		case reviewprojectdetail.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				rpd.DeletedTime = new(time.Time)
				*rpd.DeletedTime = value.Time
			}
		case reviewprojectdetail.FieldModifiedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_time", values[i])
			} else if value.Valid {
				rpd.ModifiedTime = new(time.Time)
				*rpd.ModifiedTime = value.Time
			}
		default:
			rpd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ReviewProjectDetail.
// This includes values selected through modifiers, order, etc.
func (rpd *ReviewProjectDetail) Value(name string) (ent.Value, error) {
	return rpd.selectValues.Get(name)
}

// QueryReviewProject queries the "review_project" edge of the ReviewProjectDetail entity.
func (rpd *ReviewProjectDetail) QueryReviewProject() *ReviewProjectQuery {
	return NewReviewProjectDetailClient(rpd.config).QueryReviewProject(rpd)
}

// Update returns a builder for updating this ReviewProjectDetail.
// Note that you need to call ReviewProjectDetail.Unwrap() before calling this method if this ReviewProjectDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (rpd *ReviewProjectDetail) Update() *ReviewProjectDetailUpdateOne {
	return NewReviewProjectDetailClient(rpd.config).UpdateOne(rpd)
}

// Unwrap unwraps the ReviewProjectDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rpd *ReviewProjectDetail) Unwrap() *ReviewProjectDetail {
	_tx, ok := rpd.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReviewProjectDetail is not a transactional entity")
	}
	rpd.config.driver = _tx.drv
	return rpd
}

// String implements the fmt.Stringer.
func (rpd *ReviewProjectDetail) String() string {
	var builder strings.Builder
	builder.WriteString("ReviewProjectDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rpd.ID))
	builder.WriteString("review_project_id=")
	builder.WriteString(fmt.Sprintf("%v", rpd.ReviewProjectID))
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", rpd.Order))
	builder.WriteString(", ")
	builder.WriteString("reviewer=")
	builder.WriteString(fmt.Sprintf("%v", rpd.Reviewer))
	builder.WriteString(", ")
	builder.WriteString("executor=")
	builder.WriteString(fmt.Sprintf("%v", rpd.Executor))
	builder.WriteString(", ")
	builder.WriteString("typee=")
	builder.WriteString(fmt.Sprintf("%v", rpd.Typee))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", rpd.Status))
	builder.WriteString(", ")
	if v := rpd.CreatedTime; v != nil {
		builder.WriteString("created_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := rpd.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := rpd.ModifiedTime; v != nil {
		builder.WriteString("modified_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ReviewProjectDetails is a parsable slice of ReviewProjectDetail.
type ReviewProjectDetails []*ReviewProjectDetail
