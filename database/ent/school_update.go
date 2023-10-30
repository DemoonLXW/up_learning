// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/school"
	"github.com/DemoonLXW/up_learning/database/ent/student"
)

// SchoolUpdate is the builder for updating School entities.
type SchoolUpdate struct {
	config
	hooks    []Hook
	mutation *SchoolMutation
}

// Where appends a list predicates to the SchoolUpdate builder.
func (su *SchoolUpdate) Where(ps ...predicate.School) *SchoolUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCode sets the "code" field.
func (su *SchoolUpdate) SetCode(s string) *SchoolUpdate {
	su.mutation.SetCode(s)
	return su
}

// SetName sets the "name" field.
func (su *SchoolUpdate) SetName(s string) *SchoolUpdate {
	su.mutation.SetName(s)
	return su
}

// SetLocation sets the "location" field.
func (su *SchoolUpdate) SetLocation(s string) *SchoolUpdate {
	su.mutation.SetLocation(s)
	return su
}

// SetCompetentDepartment sets the "competent_department" field.
func (su *SchoolUpdate) SetCompetentDepartment(s string) *SchoolUpdate {
	su.mutation.SetCompetentDepartment(s)
	return su
}

// SetEducationLevel sets the "education_level" field.
func (su *SchoolUpdate) SetEducationLevel(u uint8) *SchoolUpdate {
	su.mutation.ResetEducationLevel()
	su.mutation.SetEducationLevel(u)
	return su
}

// AddEducationLevel adds u to the "education_level" field.
func (su *SchoolUpdate) AddEducationLevel(u int8) *SchoolUpdate {
	su.mutation.AddEducationLevel(u)
	return su
}

// SetRemark sets the "remark" field.
func (su *SchoolUpdate) SetRemark(s string) *SchoolUpdate {
	su.mutation.SetRemark(s)
	return su
}

// SetIsDisabled sets the "is_disabled" field.
func (su *SchoolUpdate) SetIsDisabled(b bool) *SchoolUpdate {
	su.mutation.SetIsDisabled(b)
	return su
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableIsDisabled(b *bool) *SchoolUpdate {
	if b != nil {
		su.SetIsDisabled(*b)
	}
	return su
}

// SetCreatedTime sets the "created_time" field.
func (su *SchoolUpdate) SetCreatedTime(t time.Time) *SchoolUpdate {
	su.mutation.SetCreatedTime(t)
	return su
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableCreatedTime(t *time.Time) *SchoolUpdate {
	if t != nil {
		su.SetCreatedTime(*t)
	}
	return su
}

// SetDeletedTime sets the "deleted_time" field.
func (su *SchoolUpdate) SetDeletedTime(t time.Time) *SchoolUpdate {
	su.mutation.SetDeletedTime(t)
	return su
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableDeletedTime(t *time.Time) *SchoolUpdate {
	if t != nil {
		su.SetDeletedTime(*t)
	}
	return su
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (su *SchoolUpdate) ClearDeletedTime() *SchoolUpdate {
	su.mutation.ClearDeletedTime()
	return su
}

// SetModifiedTime sets the "modified_time" field.
func (su *SchoolUpdate) SetModifiedTime(t time.Time) *SchoolUpdate {
	su.mutation.SetModifiedTime(t)
	return su
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableModifiedTime(t *time.Time) *SchoolUpdate {
	if t != nil {
		su.SetModifiedTime(*t)
	}
	return su
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (su *SchoolUpdate) ClearModifiedTime() *SchoolUpdate {
	su.mutation.ClearModifiedTime()
	return su
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (su *SchoolUpdate) AddStudentIDs(ids ...uint32) *SchoolUpdate {
	su.mutation.AddStudentIDs(ids...)
	return su
}

// AddStudents adds the "students" edges to the Student entity.
func (su *SchoolUpdate) AddStudents(s ...*Student) *SchoolUpdate {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddStudentIDs(ids...)
}

// Mutation returns the SchoolMutation object of the builder.
func (su *SchoolUpdate) Mutation() *SchoolMutation {
	return su.mutation
}

// ClearStudents clears all "students" edges to the Student entity.
func (su *SchoolUpdate) ClearStudents() *SchoolUpdate {
	su.mutation.ClearStudents()
	return su
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (su *SchoolUpdate) RemoveStudentIDs(ids ...uint32) *SchoolUpdate {
	su.mutation.RemoveStudentIDs(ids...)
	return su
}

// RemoveStudents removes "students" edges to Student entities.
func (su *SchoolUpdate) RemoveStudents(s ...*Student) *SchoolUpdate {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveStudentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SchoolUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SchoolMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SchoolUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SchoolUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SchoolUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SchoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(school.Table, school.Columns, sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Code(); ok {
		_spec.SetField(school.FieldCode, field.TypeString, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(school.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Location(); ok {
		_spec.SetField(school.FieldLocation, field.TypeString, value)
	}
	if value, ok := su.mutation.CompetentDepartment(); ok {
		_spec.SetField(school.FieldCompetentDepartment, field.TypeString, value)
	}
	if value, ok := su.mutation.EducationLevel(); ok {
		_spec.SetField(school.FieldEducationLevel, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedEducationLevel(); ok {
		_spec.AddField(school.FieldEducationLevel, field.TypeUint8, value)
	}
	if value, ok := su.mutation.Remark(); ok {
		_spec.SetField(school.FieldRemark, field.TypeString, value)
	}
	if value, ok := su.mutation.IsDisabled(); ok {
		_spec.SetField(school.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := su.mutation.CreatedTime(); ok {
		_spec.SetField(school.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.DeletedTime(); ok {
		_spec.SetField(school.FieldDeletedTime, field.TypeTime, value)
	}
	if su.mutation.DeletedTimeCleared() {
		_spec.ClearField(school.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := su.mutation.ModifiedTime(); ok {
		_spec.SetField(school.FieldModifiedTime, field.TypeTime, value)
	}
	if su.mutation.ModifiedTimeCleared() {
		_spec.ClearField(school.FieldModifiedTime, field.TypeTime)
	}
	if su.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StudentsTable,
			Columns: []string{school.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !su.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StudentsTable,
			Columns: []string{school.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StudentsTable,
			Columns: []string{school.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{school.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SchoolUpdateOne is the builder for updating a single School entity.
type SchoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SchoolMutation
}

// SetCode sets the "code" field.
func (suo *SchoolUpdateOne) SetCode(s string) *SchoolUpdateOne {
	suo.mutation.SetCode(s)
	return suo
}

// SetName sets the "name" field.
func (suo *SchoolUpdateOne) SetName(s string) *SchoolUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetLocation sets the "location" field.
func (suo *SchoolUpdateOne) SetLocation(s string) *SchoolUpdateOne {
	suo.mutation.SetLocation(s)
	return suo
}

// SetCompetentDepartment sets the "competent_department" field.
func (suo *SchoolUpdateOne) SetCompetentDepartment(s string) *SchoolUpdateOne {
	suo.mutation.SetCompetentDepartment(s)
	return suo
}

// SetEducationLevel sets the "education_level" field.
func (suo *SchoolUpdateOne) SetEducationLevel(u uint8) *SchoolUpdateOne {
	suo.mutation.ResetEducationLevel()
	suo.mutation.SetEducationLevel(u)
	return suo
}

// AddEducationLevel adds u to the "education_level" field.
func (suo *SchoolUpdateOne) AddEducationLevel(u int8) *SchoolUpdateOne {
	suo.mutation.AddEducationLevel(u)
	return suo
}

// SetRemark sets the "remark" field.
func (suo *SchoolUpdateOne) SetRemark(s string) *SchoolUpdateOne {
	suo.mutation.SetRemark(s)
	return suo
}

// SetIsDisabled sets the "is_disabled" field.
func (suo *SchoolUpdateOne) SetIsDisabled(b bool) *SchoolUpdateOne {
	suo.mutation.SetIsDisabled(b)
	return suo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableIsDisabled(b *bool) *SchoolUpdateOne {
	if b != nil {
		suo.SetIsDisabled(*b)
	}
	return suo
}

// SetCreatedTime sets the "created_time" field.
func (suo *SchoolUpdateOne) SetCreatedTime(t time.Time) *SchoolUpdateOne {
	suo.mutation.SetCreatedTime(t)
	return suo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableCreatedTime(t *time.Time) *SchoolUpdateOne {
	if t != nil {
		suo.SetCreatedTime(*t)
	}
	return suo
}

// SetDeletedTime sets the "deleted_time" field.
func (suo *SchoolUpdateOne) SetDeletedTime(t time.Time) *SchoolUpdateOne {
	suo.mutation.SetDeletedTime(t)
	return suo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableDeletedTime(t *time.Time) *SchoolUpdateOne {
	if t != nil {
		suo.SetDeletedTime(*t)
	}
	return suo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (suo *SchoolUpdateOne) ClearDeletedTime() *SchoolUpdateOne {
	suo.mutation.ClearDeletedTime()
	return suo
}

// SetModifiedTime sets the "modified_time" field.
func (suo *SchoolUpdateOne) SetModifiedTime(t time.Time) *SchoolUpdateOne {
	suo.mutation.SetModifiedTime(t)
	return suo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableModifiedTime(t *time.Time) *SchoolUpdateOne {
	if t != nil {
		suo.SetModifiedTime(*t)
	}
	return suo
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (suo *SchoolUpdateOne) ClearModifiedTime() *SchoolUpdateOne {
	suo.mutation.ClearModifiedTime()
	return suo
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (suo *SchoolUpdateOne) AddStudentIDs(ids ...uint32) *SchoolUpdateOne {
	suo.mutation.AddStudentIDs(ids...)
	return suo
}

// AddStudents adds the "students" edges to the Student entity.
func (suo *SchoolUpdateOne) AddStudents(s ...*Student) *SchoolUpdateOne {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddStudentIDs(ids...)
}

// Mutation returns the SchoolMutation object of the builder.
func (suo *SchoolUpdateOne) Mutation() *SchoolMutation {
	return suo.mutation
}

// ClearStudents clears all "students" edges to the Student entity.
func (suo *SchoolUpdateOne) ClearStudents() *SchoolUpdateOne {
	suo.mutation.ClearStudents()
	return suo
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (suo *SchoolUpdateOne) RemoveStudentIDs(ids ...uint32) *SchoolUpdateOne {
	suo.mutation.RemoveStudentIDs(ids...)
	return suo
}

// RemoveStudents removes "students" edges to Student entities.
func (suo *SchoolUpdateOne) RemoveStudents(s ...*Student) *SchoolUpdateOne {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveStudentIDs(ids...)
}

// Where appends a list predicates to the SchoolUpdate builder.
func (suo *SchoolUpdateOne) Where(ps ...predicate.School) *SchoolUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SchoolUpdateOne) Select(field string, fields ...string) *SchoolUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated School entity.
func (suo *SchoolUpdateOne) Save(ctx context.Context) (*School, error) {
	return withHooks[*School, SchoolMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SchoolUpdateOne) SaveX(ctx context.Context) *School {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SchoolUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SchoolUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SchoolUpdateOne) sqlSave(ctx context.Context) (_node *School, err error) {
	_spec := sqlgraph.NewUpdateSpec(school.Table, school.Columns, sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "School.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, school.FieldID)
		for _, f := range fields {
			if !school.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != school.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Code(); ok {
		_spec.SetField(school.FieldCode, field.TypeString, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(school.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Location(); ok {
		_spec.SetField(school.FieldLocation, field.TypeString, value)
	}
	if value, ok := suo.mutation.CompetentDepartment(); ok {
		_spec.SetField(school.FieldCompetentDepartment, field.TypeString, value)
	}
	if value, ok := suo.mutation.EducationLevel(); ok {
		_spec.SetField(school.FieldEducationLevel, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedEducationLevel(); ok {
		_spec.AddField(school.FieldEducationLevel, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.Remark(); ok {
		_spec.SetField(school.FieldRemark, field.TypeString, value)
	}
	if value, ok := suo.mutation.IsDisabled(); ok {
		_spec.SetField(school.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := suo.mutation.CreatedTime(); ok {
		_spec.SetField(school.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.DeletedTime(); ok {
		_spec.SetField(school.FieldDeletedTime, field.TypeTime, value)
	}
	if suo.mutation.DeletedTimeCleared() {
		_spec.ClearField(school.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := suo.mutation.ModifiedTime(); ok {
		_spec.SetField(school.FieldModifiedTime, field.TypeTime, value)
	}
	if suo.mutation.ModifiedTimeCleared() {
		_spec.ClearField(school.FieldModifiedTime, field.TypeTime)
	}
	if suo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StudentsTable,
			Columns: []string{school.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !suo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StudentsTable,
			Columns: []string{school.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   school.StudentsTable,
			Columns: []string{school.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &School{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{school.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
