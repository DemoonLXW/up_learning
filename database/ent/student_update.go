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
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUID sets the "uid" field.
func (su *StudentUpdate) SetUID(u uint32) *StudentUpdate {
	su.mutation.SetUID(u)
	return su
}

// SetSid sets the "sid" field.
func (su *StudentUpdate) SetSid(u uint16) *StudentUpdate {
	su.mutation.SetSid(u)
	return su
}

// SetStudentID sets the "student_id" field.
func (su *StudentUpdate) SetStudentID(s string) *StudentUpdate {
	su.mutation.SetStudentID(s)
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetGender sets the "gender" field.
func (su *StudentUpdate) SetGender(u uint8) *StudentUpdate {
	su.mutation.ResetGender()
	su.mutation.SetGender(u)
	return su
}

// AddGender adds u to the "gender" field.
func (su *StudentUpdate) AddGender(u int8) *StudentUpdate {
	su.mutation.AddGender(u)
	return su
}

// SetIsDisabled sets the "is_disabled" field.
func (su *StudentUpdate) SetIsDisabled(b bool) *StudentUpdate {
	su.mutation.SetIsDisabled(b)
	return su
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (su *StudentUpdate) SetNillableIsDisabled(b *bool) *StudentUpdate {
	if b != nil {
		su.SetIsDisabled(*b)
	}
	return su
}

// SetCreatedTime sets the "created_time" field.
func (su *StudentUpdate) SetCreatedTime(t time.Time) *StudentUpdate {
	su.mutation.SetCreatedTime(t)
	return su
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (su *StudentUpdate) SetNillableCreatedTime(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetCreatedTime(*t)
	}
	return su
}

// SetDeletedTime sets the "deleted_time" field.
func (su *StudentUpdate) SetDeletedTime(t time.Time) *StudentUpdate {
	su.mutation.SetDeletedTime(t)
	return su
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (su *StudentUpdate) SetNillableDeletedTime(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetDeletedTime(*t)
	}
	return su
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (su *StudentUpdate) ClearDeletedTime() *StudentUpdate {
	su.mutation.ClearDeletedTime()
	return su
}

// SetModifiedTime sets the "modified_time" field.
func (su *StudentUpdate) SetModifiedTime(t time.Time) *StudentUpdate {
	su.mutation.SetModifiedTime(t)
	return su
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (su *StudentUpdate) SetNillableModifiedTime(t *time.Time) *StudentUpdate {
	if t != nil {
		su.SetModifiedTime(*t)
	}
	return su
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (su *StudentUpdate) ClearModifiedTime() *StudentUpdate {
	su.mutation.ClearModifiedTime()
	return su
}

// SetSchoolID sets the "school" edge to the School entity by ID.
func (su *StudentUpdate) SetSchoolID(id uint16) *StudentUpdate {
	su.mutation.SetSchoolID(id)
	return su
}

// SetSchool sets the "school" edge to the School entity.
func (su *StudentUpdate) SetSchool(s *School) *StudentUpdate {
	return su.SetSchoolID(s.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *StudentUpdate) SetUserID(id uint32) *StudentUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *StudentUpdate) SetUser(u *User) *StudentUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearSchool clears the "school" edge to the School entity.
func (su *StudentUpdate) ClearSchool() *StudentUpdate {
	su.mutation.ClearSchool()
	return su
}

// ClearUser clears the "user" edge to the User entity.
func (su *StudentUpdate) ClearUser() *StudentUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StudentUpdate) check() error {
	if v, ok := su.mutation.StudentID(); ok {
		if err := student.StudentIDValidator(v); err != nil {
			return &ValidationError{Name: "student_id", err: fmt.Errorf(`ent: validator failed for field "Student.student_id": %w`, err)}
		}
	}
	if v, ok := su.mutation.Name(); ok {
		if err := student.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Student.name": %w`, err)}
		}
	}
	if _, ok := su.mutation.SchoolID(); su.mutation.SchoolCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Student.school"`)
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Student.user"`)
	}
	return nil
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.StudentID(); ok {
		_spec.SetField(student.FieldStudentID, field.TypeString, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Gender(); ok {
		_spec.SetField(student.FieldGender, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedGender(); ok {
		_spec.AddField(student.FieldGender, field.TypeUint8, value)
	}
	if value, ok := su.mutation.IsDisabled(); ok {
		_spec.SetField(student.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := su.mutation.CreatedTime(); ok {
		_spec.SetField(student.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.DeletedTime(); ok {
		_spec.SetField(student.FieldDeletedTime, field.TypeTime, value)
	}
	if su.mutation.DeletedTimeCleared() {
		_spec.ClearField(student.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := su.mutation.ModifiedTime(); ok {
		_spec.SetField(student.FieldModifiedTime, field.TypeTime, value)
	}
	if su.mutation.ModifiedTimeCleared() {
		_spec.ClearField(student.FieldModifiedTime, field.TypeTime)
	}
	if su.mutation.SchoolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.SchoolTable,
			Columns: []string{student.SchoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SchoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.SchoolTable,
			Columns: []string{student.SchoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetUID sets the "uid" field.
func (suo *StudentUpdateOne) SetUID(u uint32) *StudentUpdateOne {
	suo.mutation.SetUID(u)
	return suo
}

// SetSid sets the "sid" field.
func (suo *StudentUpdateOne) SetSid(u uint16) *StudentUpdateOne {
	suo.mutation.SetSid(u)
	return suo
}

// SetStudentID sets the "student_id" field.
func (suo *StudentUpdateOne) SetStudentID(s string) *StudentUpdateOne {
	suo.mutation.SetStudentID(s)
	return suo
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetGender sets the "gender" field.
func (suo *StudentUpdateOne) SetGender(u uint8) *StudentUpdateOne {
	suo.mutation.ResetGender()
	suo.mutation.SetGender(u)
	return suo
}

// AddGender adds u to the "gender" field.
func (suo *StudentUpdateOne) AddGender(u int8) *StudentUpdateOne {
	suo.mutation.AddGender(u)
	return suo
}

// SetIsDisabled sets the "is_disabled" field.
func (suo *StudentUpdateOne) SetIsDisabled(b bool) *StudentUpdateOne {
	suo.mutation.SetIsDisabled(b)
	return suo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableIsDisabled(b *bool) *StudentUpdateOne {
	if b != nil {
		suo.SetIsDisabled(*b)
	}
	return suo
}

// SetCreatedTime sets the "created_time" field.
func (suo *StudentUpdateOne) SetCreatedTime(t time.Time) *StudentUpdateOne {
	suo.mutation.SetCreatedTime(t)
	return suo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableCreatedTime(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetCreatedTime(*t)
	}
	return suo
}

// SetDeletedTime sets the "deleted_time" field.
func (suo *StudentUpdateOne) SetDeletedTime(t time.Time) *StudentUpdateOne {
	suo.mutation.SetDeletedTime(t)
	return suo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableDeletedTime(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetDeletedTime(*t)
	}
	return suo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (suo *StudentUpdateOne) ClearDeletedTime() *StudentUpdateOne {
	suo.mutation.ClearDeletedTime()
	return suo
}

// SetModifiedTime sets the "modified_time" field.
func (suo *StudentUpdateOne) SetModifiedTime(t time.Time) *StudentUpdateOne {
	suo.mutation.SetModifiedTime(t)
	return suo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableModifiedTime(t *time.Time) *StudentUpdateOne {
	if t != nil {
		suo.SetModifiedTime(*t)
	}
	return suo
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (suo *StudentUpdateOne) ClearModifiedTime() *StudentUpdateOne {
	suo.mutation.ClearModifiedTime()
	return suo
}

// SetSchoolID sets the "school" edge to the School entity by ID.
func (suo *StudentUpdateOne) SetSchoolID(id uint16) *StudentUpdateOne {
	suo.mutation.SetSchoolID(id)
	return suo
}

// SetSchool sets the "school" edge to the School entity.
func (suo *StudentUpdateOne) SetSchool(s *School) *StudentUpdateOne {
	return suo.SetSchoolID(s.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *StudentUpdateOne) SetUserID(id uint32) *StudentUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *StudentUpdateOne) SetUser(u *User) *StudentUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearSchool clears the "school" edge to the School entity.
func (suo *StudentUpdateOne) ClearSchool() *StudentUpdateOne {
	suo.mutation.ClearSchool()
	return suo
}

// ClearUser clears the "user" edge to the User entity.
func (suo *StudentUpdateOne) ClearUser() *StudentUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StudentUpdateOne) check() error {
	if v, ok := suo.mutation.StudentID(); ok {
		if err := student.StudentIDValidator(v); err != nil {
			return &ValidationError{Name: "student_id", err: fmt.Errorf(`ent: validator failed for field "Student.student_id": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Name(); ok {
		if err := student.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Student.name": %w`, err)}
		}
	}
	if _, ok := suo.mutation.SchoolID(); suo.mutation.SchoolCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Student.school"`)
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Student.user"`)
	}
	return nil
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
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
	if value, ok := suo.mutation.StudentID(); ok {
		_spec.SetField(student.FieldStudentID, field.TypeString, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Gender(); ok {
		_spec.SetField(student.FieldGender, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedGender(); ok {
		_spec.AddField(student.FieldGender, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.IsDisabled(); ok {
		_spec.SetField(student.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := suo.mutation.CreatedTime(); ok {
		_spec.SetField(student.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.DeletedTime(); ok {
		_spec.SetField(student.FieldDeletedTime, field.TypeTime, value)
	}
	if suo.mutation.DeletedTimeCleared() {
		_spec.ClearField(student.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := suo.mutation.ModifiedTime(); ok {
		_spec.SetField(student.FieldModifiedTime, field.TypeTime, value)
	}
	if suo.mutation.ModifiedTimeCleared() {
		_spec.ClearField(student.FieldModifiedTime, field.TypeTime)
	}
	if suo.mutation.SchoolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.SchoolTable,
			Columns: []string{student.SchoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SchoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.SchoolTable,
			Columns: []string{student.SchoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
