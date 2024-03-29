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
	"github.com/DemoonLXW/up_learning/database/ent/major"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/teacher"
)

// CollegeUpdate is the builder for updating College entities.
type CollegeUpdate struct {
	config
	hooks    []Hook
	mutation *CollegeMutation
}

// Where appends a list predicates to the CollegeUpdate builder.
func (cu *CollegeUpdate) Where(ps ...predicate.College) *CollegeUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CollegeUpdate) SetName(s string) *CollegeUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetIsDisabled sets the "is_disabled" field.
func (cu *CollegeUpdate) SetIsDisabled(b bool) *CollegeUpdate {
	cu.mutation.SetIsDisabled(b)
	return cu
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (cu *CollegeUpdate) SetNillableIsDisabled(b *bool) *CollegeUpdate {
	if b != nil {
		cu.SetIsDisabled(*b)
	}
	return cu
}

// SetCreatedTime sets the "created_time" field.
func (cu *CollegeUpdate) SetCreatedTime(t time.Time) *CollegeUpdate {
	cu.mutation.SetCreatedTime(t)
	return cu
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (cu *CollegeUpdate) SetNillableCreatedTime(t *time.Time) *CollegeUpdate {
	if t != nil {
		cu.SetCreatedTime(*t)
	}
	return cu
}

// SetDeletedTime sets the "deleted_time" field.
func (cu *CollegeUpdate) SetDeletedTime(t time.Time) *CollegeUpdate {
	cu.mutation.SetDeletedTime(t)
	return cu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (cu *CollegeUpdate) SetNillableDeletedTime(t *time.Time) *CollegeUpdate {
	if t != nil {
		cu.SetDeletedTime(*t)
	}
	return cu
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (cu *CollegeUpdate) ClearDeletedTime() *CollegeUpdate {
	cu.mutation.ClearDeletedTime()
	return cu
}

// SetModifiedTime sets the "modified_time" field.
func (cu *CollegeUpdate) SetModifiedTime(t time.Time) *CollegeUpdate {
	cu.mutation.SetModifiedTime(t)
	return cu
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (cu *CollegeUpdate) SetNillableModifiedTime(t *time.Time) *CollegeUpdate {
	if t != nil {
		cu.SetModifiedTime(*t)
	}
	return cu
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (cu *CollegeUpdate) ClearModifiedTime() *CollegeUpdate {
	cu.mutation.ClearModifiedTime()
	return cu
}

// AddMajorIDs adds the "majors" edge to the Major entity by IDs.
func (cu *CollegeUpdate) AddMajorIDs(ids ...uint16) *CollegeUpdate {
	cu.mutation.AddMajorIDs(ids...)
	return cu
}

// AddMajors adds the "majors" edges to the Major entity.
func (cu *CollegeUpdate) AddMajors(m ...*Major) *CollegeUpdate {
	ids := make([]uint16, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMajorIDs(ids...)
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (cu *CollegeUpdate) AddTeacherIDs(ids ...uint32) *CollegeUpdate {
	cu.mutation.AddTeacherIDs(ids...)
	return cu
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (cu *CollegeUpdate) AddTeachers(t ...*Teacher) *CollegeUpdate {
	ids := make([]uint32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTeacherIDs(ids...)
}

// Mutation returns the CollegeMutation object of the builder.
func (cu *CollegeUpdate) Mutation() *CollegeMutation {
	return cu.mutation
}

// ClearMajors clears all "majors" edges to the Major entity.
func (cu *CollegeUpdate) ClearMajors() *CollegeUpdate {
	cu.mutation.ClearMajors()
	return cu
}

// RemoveMajorIDs removes the "majors" edge to Major entities by IDs.
func (cu *CollegeUpdate) RemoveMajorIDs(ids ...uint16) *CollegeUpdate {
	cu.mutation.RemoveMajorIDs(ids...)
	return cu
}

// RemoveMajors removes "majors" edges to Major entities.
func (cu *CollegeUpdate) RemoveMajors(m ...*Major) *CollegeUpdate {
	ids := make([]uint16, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMajorIDs(ids...)
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (cu *CollegeUpdate) ClearTeachers() *CollegeUpdate {
	cu.mutation.ClearTeachers()
	return cu
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (cu *CollegeUpdate) RemoveTeacherIDs(ids ...uint32) *CollegeUpdate {
	cu.mutation.RemoveTeacherIDs(ids...)
	return cu
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (cu *CollegeUpdate) RemoveTeachers(t ...*Teacher) *CollegeUpdate {
	ids := make([]uint32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTeacherIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollegeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CollegeMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollegeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollegeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollegeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CollegeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(college.Table, college.Columns, sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(college.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsDisabled(); ok {
		_spec.SetField(college.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := cu.mutation.CreatedTime(); ok {
		_spec.SetField(college.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.DeletedTime(); ok {
		_spec.SetField(college.FieldDeletedTime, field.TypeTime, value)
	}
	if cu.mutation.DeletedTimeCleared() {
		_spec.ClearField(college.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := cu.mutation.ModifiedTime(); ok {
		_spec.SetField(college.FieldModifiedTime, field.TypeTime, value)
	}
	if cu.mutation.ModifiedTimeCleared() {
		_spec.ClearField(college.FieldModifiedTime, field.TypeTime)
	}
	if cu.mutation.MajorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.MajorsTable,
			Columns: []string{college.MajorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMajorsIDs(); len(nodes) > 0 && !cu.mutation.MajorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.MajorsTable,
			Columns: []string{college.MajorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MajorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.MajorsTable,
			Columns: []string{college.MajorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.TeachersTable,
			Columns: []string{college.TeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !cu.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.TeachersTable,
			Columns: []string{college.TeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.TeachersTable,
			Columns: []string{college.TeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{college.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CollegeUpdateOne is the builder for updating a single College entity.
type CollegeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollegeMutation
}

// SetName sets the "name" field.
func (cuo *CollegeUpdateOne) SetName(s string) *CollegeUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetIsDisabled sets the "is_disabled" field.
func (cuo *CollegeUpdateOne) SetIsDisabled(b bool) *CollegeUpdateOne {
	cuo.mutation.SetIsDisabled(b)
	return cuo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (cuo *CollegeUpdateOne) SetNillableIsDisabled(b *bool) *CollegeUpdateOne {
	if b != nil {
		cuo.SetIsDisabled(*b)
	}
	return cuo
}

// SetCreatedTime sets the "created_time" field.
func (cuo *CollegeUpdateOne) SetCreatedTime(t time.Time) *CollegeUpdateOne {
	cuo.mutation.SetCreatedTime(t)
	return cuo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (cuo *CollegeUpdateOne) SetNillableCreatedTime(t *time.Time) *CollegeUpdateOne {
	if t != nil {
		cuo.SetCreatedTime(*t)
	}
	return cuo
}

// SetDeletedTime sets the "deleted_time" field.
func (cuo *CollegeUpdateOne) SetDeletedTime(t time.Time) *CollegeUpdateOne {
	cuo.mutation.SetDeletedTime(t)
	return cuo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (cuo *CollegeUpdateOne) SetNillableDeletedTime(t *time.Time) *CollegeUpdateOne {
	if t != nil {
		cuo.SetDeletedTime(*t)
	}
	return cuo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (cuo *CollegeUpdateOne) ClearDeletedTime() *CollegeUpdateOne {
	cuo.mutation.ClearDeletedTime()
	return cuo
}

// SetModifiedTime sets the "modified_time" field.
func (cuo *CollegeUpdateOne) SetModifiedTime(t time.Time) *CollegeUpdateOne {
	cuo.mutation.SetModifiedTime(t)
	return cuo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (cuo *CollegeUpdateOne) SetNillableModifiedTime(t *time.Time) *CollegeUpdateOne {
	if t != nil {
		cuo.SetModifiedTime(*t)
	}
	return cuo
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (cuo *CollegeUpdateOne) ClearModifiedTime() *CollegeUpdateOne {
	cuo.mutation.ClearModifiedTime()
	return cuo
}

// AddMajorIDs adds the "majors" edge to the Major entity by IDs.
func (cuo *CollegeUpdateOne) AddMajorIDs(ids ...uint16) *CollegeUpdateOne {
	cuo.mutation.AddMajorIDs(ids...)
	return cuo
}

// AddMajors adds the "majors" edges to the Major entity.
func (cuo *CollegeUpdateOne) AddMajors(m ...*Major) *CollegeUpdateOne {
	ids := make([]uint16, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMajorIDs(ids...)
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (cuo *CollegeUpdateOne) AddTeacherIDs(ids ...uint32) *CollegeUpdateOne {
	cuo.mutation.AddTeacherIDs(ids...)
	return cuo
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (cuo *CollegeUpdateOne) AddTeachers(t ...*Teacher) *CollegeUpdateOne {
	ids := make([]uint32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTeacherIDs(ids...)
}

// Mutation returns the CollegeMutation object of the builder.
func (cuo *CollegeUpdateOne) Mutation() *CollegeMutation {
	return cuo.mutation
}

// ClearMajors clears all "majors" edges to the Major entity.
func (cuo *CollegeUpdateOne) ClearMajors() *CollegeUpdateOne {
	cuo.mutation.ClearMajors()
	return cuo
}

// RemoveMajorIDs removes the "majors" edge to Major entities by IDs.
func (cuo *CollegeUpdateOne) RemoveMajorIDs(ids ...uint16) *CollegeUpdateOne {
	cuo.mutation.RemoveMajorIDs(ids...)
	return cuo
}

// RemoveMajors removes "majors" edges to Major entities.
func (cuo *CollegeUpdateOne) RemoveMajors(m ...*Major) *CollegeUpdateOne {
	ids := make([]uint16, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMajorIDs(ids...)
}

// ClearTeachers clears all "teachers" edges to the Teacher entity.
func (cuo *CollegeUpdateOne) ClearTeachers() *CollegeUpdateOne {
	cuo.mutation.ClearTeachers()
	return cuo
}

// RemoveTeacherIDs removes the "teachers" edge to Teacher entities by IDs.
func (cuo *CollegeUpdateOne) RemoveTeacherIDs(ids ...uint32) *CollegeUpdateOne {
	cuo.mutation.RemoveTeacherIDs(ids...)
	return cuo
}

// RemoveTeachers removes "teachers" edges to Teacher entities.
func (cuo *CollegeUpdateOne) RemoveTeachers(t ...*Teacher) *CollegeUpdateOne {
	ids := make([]uint32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTeacherIDs(ids...)
}

// Where appends a list predicates to the CollegeUpdate builder.
func (cuo *CollegeUpdateOne) Where(ps ...predicate.College) *CollegeUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollegeUpdateOne) Select(field string, fields ...string) *CollegeUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated College entity.
func (cuo *CollegeUpdateOne) Save(ctx context.Context) (*College, error) {
	return withHooks[*College, CollegeMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollegeUpdateOne) SaveX(ctx context.Context) *College {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollegeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollegeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CollegeUpdateOne) sqlSave(ctx context.Context) (_node *College, err error) {
	_spec := sqlgraph.NewUpdateSpec(college.Table, college.Columns, sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "College.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, college.FieldID)
		for _, f := range fields {
			if !college.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != college.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(college.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsDisabled(); ok {
		_spec.SetField(college.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.CreatedTime(); ok {
		_spec.SetField(college.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.DeletedTime(); ok {
		_spec.SetField(college.FieldDeletedTime, field.TypeTime, value)
	}
	if cuo.mutation.DeletedTimeCleared() {
		_spec.ClearField(college.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := cuo.mutation.ModifiedTime(); ok {
		_spec.SetField(college.FieldModifiedTime, field.TypeTime, value)
	}
	if cuo.mutation.ModifiedTimeCleared() {
		_spec.ClearField(college.FieldModifiedTime, field.TypeTime)
	}
	if cuo.mutation.MajorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.MajorsTable,
			Columns: []string{college.MajorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMajorsIDs(); len(nodes) > 0 && !cuo.mutation.MajorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.MajorsTable,
			Columns: []string{college.MajorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MajorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.MajorsTable,
			Columns: []string{college.MajorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.TeachersTable,
			Columns: []string{college.TeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTeachersIDs(); len(nodes) > 0 && !cuo.mutation.TeachersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.TeachersTable,
			Columns: []string{college.TeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TeachersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   college.TeachersTable,
			Columns: []string{college.TeachersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &College{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{college.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
