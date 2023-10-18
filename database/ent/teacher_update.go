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
	"github.com/DemoonLXW/up_learning/database/ent/college"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/teacher"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// TeacherUpdate is the builder for updating Teacher entities.
type TeacherUpdate struct {
	config
	hooks    []Hook
	mutation *TeacherMutation
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tu *TeacherUpdate) Where(ps ...predicate.Teacher) *TeacherUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUID sets the "uid" field.
func (tu *TeacherUpdate) SetUID(u uint32) *TeacherUpdate {
	tu.mutation.SetUID(u)
	return tu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableUID(u *uint32) *TeacherUpdate {
	if u != nil {
		tu.SetUID(*u)
	}
	return tu
}

// ClearUID clears the value of the "uid" field.
func (tu *TeacherUpdate) ClearUID() *TeacherUpdate {
	tu.mutation.ClearUID()
	return tu
}

// SetCid sets the "cid" field.
func (tu *TeacherUpdate) SetCid(u uint8) *TeacherUpdate {
	tu.mutation.SetCid(u)
	return tu
}

// SetTeacherID sets the "teacher_id" field.
func (tu *TeacherUpdate) SetTeacherID(s string) *TeacherUpdate {
	tu.mutation.SetTeacherID(s)
	return tu
}

// SetName sets the "name" field.
func (tu *TeacherUpdate) SetName(s string) *TeacherUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetGender sets the "gender" field.
func (tu *TeacherUpdate) SetGender(u uint8) *TeacherUpdate {
	tu.mutation.ResetGender()
	tu.mutation.SetGender(u)
	return tu
}

// AddGender adds u to the "gender" field.
func (tu *TeacherUpdate) AddGender(u int8) *TeacherUpdate {
	tu.mutation.AddGender(u)
	return tu
}

// SetIsDisabled sets the "is_disabled" field.
func (tu *TeacherUpdate) SetIsDisabled(b bool) *TeacherUpdate {
	tu.mutation.SetIsDisabled(b)
	return tu
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableIsDisabled(b *bool) *TeacherUpdate {
	if b != nil {
		tu.SetIsDisabled(*b)
	}
	return tu
}

// SetCreatedTime sets the "created_time" field.
func (tu *TeacherUpdate) SetCreatedTime(t time.Time) *TeacherUpdate {
	tu.mutation.SetCreatedTime(t)
	return tu
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableCreatedTime(t *time.Time) *TeacherUpdate {
	if t != nil {
		tu.SetCreatedTime(*t)
	}
	return tu
}

// SetDeletedTime sets the "deleted_time" field.
func (tu *TeacherUpdate) SetDeletedTime(t time.Time) *TeacherUpdate {
	tu.mutation.SetDeletedTime(t)
	return tu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableDeletedTime(t *time.Time) *TeacherUpdate {
	if t != nil {
		tu.SetDeletedTime(*t)
	}
	return tu
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (tu *TeacherUpdate) ClearDeletedTime() *TeacherUpdate {
	tu.mutation.ClearDeletedTime()
	return tu
}

// SetModifiedTime sets the "modified_time" field.
func (tu *TeacherUpdate) SetModifiedTime(t time.Time) *TeacherUpdate {
	tu.mutation.SetModifiedTime(t)
	return tu
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableModifiedTime(t *time.Time) *TeacherUpdate {
	if t != nil {
		tu.SetModifiedTime(*t)
	}
	return tu
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (tu *TeacherUpdate) ClearModifiedTime() *TeacherUpdate {
	tu.mutation.ClearModifiedTime()
	return tu
}

// SetCollegeID sets the "college" edge to the College entity by ID.
func (tu *TeacherUpdate) SetCollegeID(id uint8) *TeacherUpdate {
	tu.mutation.SetCollegeID(id)
	return tu
}

// SetCollege sets the "college" edge to the College entity.
func (tu *TeacherUpdate) SetCollege(c *College) *TeacherUpdate {
	return tu.SetCollegeID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tu *TeacherUpdate) SetUserID(id uint32) *TeacherUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tu *TeacherUpdate) SetNillableUserID(id *uint32) *TeacherUpdate {
	if id != nil {
		tu = tu.SetUserID(*id)
	}
	return tu
}

// SetUser sets the "user" edge to the User entity.
func (tu *TeacherUpdate) SetUser(u *User) *TeacherUpdate {
	return tu.SetUserID(u.ID)
}

// Mutation returns the TeacherMutation object of the builder.
func (tu *TeacherUpdate) Mutation() *TeacherMutation {
	return tu.mutation
}

// ClearCollege clears the "college" edge to the College entity.
func (tu *TeacherUpdate) ClearCollege() *TeacherUpdate {
	tu.mutation.ClearCollege()
	return tu
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TeacherUpdate) ClearUser() *TeacherUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeacherUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeacherUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeacherUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeacherUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TeacherUpdate) check() error {
	if v, ok := tu.mutation.TeacherID(); ok {
		if err := teacher.TeacherIDValidator(v); err != nil {
			return &ValidationError{Name: "teacher_id", err: fmt.Errorf(`ent: validator failed for field "Teacher.teacher_id": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Name(); ok {
		if err := teacher.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Teacher.name": %w`, err)}
		}
	}
	if _, ok := tu.mutation.CollegeID(); tu.mutation.CollegeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Teacher.college"`)
	}
	return nil
}

func (tu *TeacherUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TeacherID(); ok {
		_spec.SetField(teacher.FieldTeacherID, field.TypeString, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Gender(); ok {
		_spec.SetField(teacher.FieldGender, field.TypeUint8, value)
	}
	if value, ok := tu.mutation.AddedGender(); ok {
		_spec.AddField(teacher.FieldGender, field.TypeUint8, value)
	}
	if value, ok := tu.mutation.IsDisabled(); ok {
		_spec.SetField(teacher.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := tu.mutation.CreatedTime(); ok {
		_spec.SetField(teacher.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.DeletedTime(); ok {
		_spec.SetField(teacher.FieldDeletedTime, field.TypeTime, value)
	}
	if tu.mutation.DeletedTimeCleared() {
		_spec.ClearField(teacher.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := tu.mutation.ModifiedTime(); ok {
		_spec.SetField(teacher.FieldModifiedTime, field.TypeTime, value)
	}
	if tu.mutation.ModifiedTimeCleared() {
		_spec.ClearField(teacher.FieldModifiedTime, field.TypeTime)
	}
	if tu.mutation.CollegeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teacher.CollegeTable,
			Columns: []string{teacher.CollegeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CollegeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teacher.CollegeTable,
			Columns: []string{teacher.CollegeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeacherUpdateOne is the builder for updating a single Teacher entity.
type TeacherUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeacherMutation
}

// SetUID sets the "uid" field.
func (tuo *TeacherUpdateOne) SetUID(u uint32) *TeacherUpdateOne {
	tuo.mutation.SetUID(u)
	return tuo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableUID(u *uint32) *TeacherUpdateOne {
	if u != nil {
		tuo.SetUID(*u)
	}
	return tuo
}

// ClearUID clears the value of the "uid" field.
func (tuo *TeacherUpdateOne) ClearUID() *TeacherUpdateOne {
	tuo.mutation.ClearUID()
	return tuo
}

// SetCid sets the "cid" field.
func (tuo *TeacherUpdateOne) SetCid(u uint8) *TeacherUpdateOne {
	tuo.mutation.SetCid(u)
	return tuo
}

// SetTeacherID sets the "teacher_id" field.
func (tuo *TeacherUpdateOne) SetTeacherID(s string) *TeacherUpdateOne {
	tuo.mutation.SetTeacherID(s)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TeacherUpdateOne) SetName(s string) *TeacherUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetGender sets the "gender" field.
func (tuo *TeacherUpdateOne) SetGender(u uint8) *TeacherUpdateOne {
	tuo.mutation.ResetGender()
	tuo.mutation.SetGender(u)
	return tuo
}

// AddGender adds u to the "gender" field.
func (tuo *TeacherUpdateOne) AddGender(u int8) *TeacherUpdateOne {
	tuo.mutation.AddGender(u)
	return tuo
}

// SetIsDisabled sets the "is_disabled" field.
func (tuo *TeacherUpdateOne) SetIsDisabled(b bool) *TeacherUpdateOne {
	tuo.mutation.SetIsDisabled(b)
	return tuo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableIsDisabled(b *bool) *TeacherUpdateOne {
	if b != nil {
		tuo.SetIsDisabled(*b)
	}
	return tuo
}

// SetCreatedTime sets the "created_time" field.
func (tuo *TeacherUpdateOne) SetCreatedTime(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetCreatedTime(t)
	return tuo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableCreatedTime(t *time.Time) *TeacherUpdateOne {
	if t != nil {
		tuo.SetCreatedTime(*t)
	}
	return tuo
}

// SetDeletedTime sets the "deleted_time" field.
func (tuo *TeacherUpdateOne) SetDeletedTime(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetDeletedTime(t)
	return tuo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableDeletedTime(t *time.Time) *TeacherUpdateOne {
	if t != nil {
		tuo.SetDeletedTime(*t)
	}
	return tuo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (tuo *TeacherUpdateOne) ClearDeletedTime() *TeacherUpdateOne {
	tuo.mutation.ClearDeletedTime()
	return tuo
}

// SetModifiedTime sets the "modified_time" field.
func (tuo *TeacherUpdateOne) SetModifiedTime(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetModifiedTime(t)
	return tuo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableModifiedTime(t *time.Time) *TeacherUpdateOne {
	if t != nil {
		tuo.SetModifiedTime(*t)
	}
	return tuo
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (tuo *TeacherUpdateOne) ClearModifiedTime() *TeacherUpdateOne {
	tuo.mutation.ClearModifiedTime()
	return tuo
}

// SetCollegeID sets the "college" edge to the College entity by ID.
func (tuo *TeacherUpdateOne) SetCollegeID(id uint8) *TeacherUpdateOne {
	tuo.mutation.SetCollegeID(id)
	return tuo
}

// SetCollege sets the "college" edge to the College entity.
func (tuo *TeacherUpdateOne) SetCollege(c *College) *TeacherUpdateOne {
	return tuo.SetCollegeID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (tuo *TeacherUpdateOne) SetUserID(id uint32) *TeacherUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableUserID(id *uint32) *TeacherUpdateOne {
	if id != nil {
		tuo = tuo.SetUserID(*id)
	}
	return tuo
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TeacherUpdateOne) SetUser(u *User) *TeacherUpdateOne {
	return tuo.SetUserID(u.ID)
}

// Mutation returns the TeacherMutation object of the builder.
func (tuo *TeacherUpdateOne) Mutation() *TeacherMutation {
	return tuo.mutation
}

// ClearCollege clears the "college" edge to the College entity.
func (tuo *TeacherUpdateOne) ClearCollege() *TeacherUpdateOne {
	tuo.mutation.ClearCollege()
	return tuo
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TeacherUpdateOne) ClearUser() *TeacherUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tuo *TeacherUpdateOne) Where(ps ...predicate.Teacher) *TeacherUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeacherUpdateOne) Select(field string, fields ...string) *TeacherUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teacher entity.
func (tuo *TeacherUpdateOne) Save(ctx context.Context) (*Teacher, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeacherUpdateOne) SaveX(ctx context.Context) *Teacher {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeacherUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeacherUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TeacherUpdateOne) check() error {
	if v, ok := tuo.mutation.TeacherID(); ok {
		if err := teacher.TeacherIDValidator(v); err != nil {
			return &ValidationError{Name: "teacher_id", err: fmt.Errorf(`ent: validator failed for field "Teacher.teacher_id": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Name(); ok {
		if err := teacher.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Teacher.name": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.CollegeID(); tuo.mutation.CollegeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Teacher.college"`)
	}
	return nil
}

func (tuo *TeacherUpdateOne) sqlSave(ctx context.Context) (_node *Teacher, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Teacher.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teacher.FieldID)
		for _, f := range fields {
			if !teacher.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TeacherID(); ok {
		_spec.SetField(teacher.FieldTeacherID, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Gender(); ok {
		_spec.SetField(teacher.FieldGender, field.TypeUint8, value)
	}
	if value, ok := tuo.mutation.AddedGender(); ok {
		_spec.AddField(teacher.FieldGender, field.TypeUint8, value)
	}
	if value, ok := tuo.mutation.IsDisabled(); ok {
		_spec.SetField(teacher.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.CreatedTime(); ok {
		_spec.SetField(teacher.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.DeletedTime(); ok {
		_spec.SetField(teacher.FieldDeletedTime, field.TypeTime, value)
	}
	if tuo.mutation.DeletedTimeCleared() {
		_spec.ClearField(teacher.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := tuo.mutation.ModifiedTime(); ok {
		_spec.SetField(teacher.FieldModifiedTime, field.TypeTime, value)
	}
	if tuo.mutation.ModifiedTimeCleared() {
		_spec.ClearField(teacher.FieldModifiedTime, field.TypeTime)
	}
	if tuo.mutation.CollegeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teacher.CollegeTable,
			Columns: []string{teacher.CollegeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CollegeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teacher.CollegeTable,
			Columns: []string{teacher.CollegeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
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
	_node = &Teacher{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}