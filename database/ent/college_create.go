// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/college"
	"github.com/DemoonLXW/up_learning/database/ent/major"
	"github.com/DemoonLXW/up_learning/database/ent/teacher"
)

// CollegeCreate is the builder for creating a College entity.
type CollegeCreate struct {
	config
	mutation *CollegeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CollegeCreate) SetName(s string) *CollegeCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetIsDisabled sets the "is_disabled" field.
func (cc *CollegeCreate) SetIsDisabled(b bool) *CollegeCreate {
	cc.mutation.SetIsDisabled(b)
	return cc
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (cc *CollegeCreate) SetNillableIsDisabled(b *bool) *CollegeCreate {
	if b != nil {
		cc.SetIsDisabled(*b)
	}
	return cc
}

// SetCreatedTime sets the "created_time" field.
func (cc *CollegeCreate) SetCreatedTime(t time.Time) *CollegeCreate {
	cc.mutation.SetCreatedTime(t)
	return cc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (cc *CollegeCreate) SetNillableCreatedTime(t *time.Time) *CollegeCreate {
	if t != nil {
		cc.SetCreatedTime(*t)
	}
	return cc
}

// SetDeletedTime sets the "deleted_time" field.
func (cc *CollegeCreate) SetDeletedTime(t time.Time) *CollegeCreate {
	cc.mutation.SetDeletedTime(t)
	return cc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (cc *CollegeCreate) SetNillableDeletedTime(t *time.Time) *CollegeCreate {
	if t != nil {
		cc.SetDeletedTime(*t)
	}
	return cc
}

// SetModifiedTime sets the "modified_time" field.
func (cc *CollegeCreate) SetModifiedTime(t time.Time) *CollegeCreate {
	cc.mutation.SetModifiedTime(t)
	return cc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (cc *CollegeCreate) SetNillableModifiedTime(t *time.Time) *CollegeCreate {
	if t != nil {
		cc.SetModifiedTime(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CollegeCreate) SetID(u uint8) *CollegeCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddMajorIDs adds the "majors" edge to the Major entity by IDs.
func (cc *CollegeCreate) AddMajorIDs(ids ...uint16) *CollegeCreate {
	cc.mutation.AddMajorIDs(ids...)
	return cc
}

// AddMajors adds the "majors" edges to the Major entity.
func (cc *CollegeCreate) AddMajors(m ...*Major) *CollegeCreate {
	ids := make([]uint16, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cc.AddMajorIDs(ids...)
}

// AddTeacherIDs adds the "teachers" edge to the Teacher entity by IDs.
func (cc *CollegeCreate) AddTeacherIDs(ids ...uint32) *CollegeCreate {
	cc.mutation.AddTeacherIDs(ids...)
	return cc
}

// AddTeachers adds the "teachers" edges to the Teacher entity.
func (cc *CollegeCreate) AddTeachers(t ...*Teacher) *CollegeCreate {
	ids := make([]uint32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTeacherIDs(ids...)
}

// Mutation returns the CollegeMutation object of the builder.
func (cc *CollegeCreate) Mutation() *CollegeMutation {
	return cc.mutation
}

// Save creates the College in the database.
func (cc *CollegeCreate) Save(ctx context.Context) (*College, error) {
	cc.defaults()
	return withHooks[*College, CollegeMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CollegeCreate) SaveX(ctx context.Context) *College {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CollegeCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CollegeCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CollegeCreate) defaults() {
	if _, ok := cc.mutation.IsDisabled(); !ok {
		v := college.DefaultIsDisabled
		cc.mutation.SetIsDisabled(v)
	}
	if _, ok := cc.mutation.CreatedTime(); !ok {
		v := college.DefaultCreatedTime()
		cc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CollegeCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "College.name"`)}
	}
	if _, ok := cc.mutation.IsDisabled(); !ok {
		return &ValidationError{Name: "is_disabled", err: errors.New(`ent: missing required field "College.is_disabled"`)}
	}
	if _, ok := cc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "College.created_time"`)}
	}
	return nil
}

func (cc *CollegeCreate) sqlSave(ctx context.Context) (*College, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint8(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CollegeCreate) createSpec() (*College, *sqlgraph.CreateSpec) {
	var (
		_node = &College{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(college.Table, sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(college.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.IsDisabled(); ok {
		_spec.SetField(college.FieldIsDisabled, field.TypeBool, value)
		_node.IsDisabled = value
	}
	if value, ok := cc.mutation.CreatedTime(); ok {
		_spec.SetField(college.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := cc.mutation.DeletedTime(); ok {
		_spec.SetField(college.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := cc.mutation.ModifiedTime(); ok {
		_spec.SetField(college.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := cc.mutation.MajorsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TeachersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CollegeCreateBulk is the builder for creating many College entities in bulk.
type CollegeCreateBulk struct {
	config
	builders []*CollegeCreate
}

// Save creates the College entities in the database.
func (ccb *CollegeCreateBulk) Save(ctx context.Context) ([]*College, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*College, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CollegeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint8(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CollegeCreateBulk) SaveX(ctx context.Context) []*College {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CollegeCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CollegeCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
