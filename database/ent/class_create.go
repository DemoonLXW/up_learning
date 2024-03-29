// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/class"
	"github.com/DemoonLXW/up_learning/database/ent/major"
	"github.com/DemoonLXW/up_learning/database/ent/student"
)

// ClassCreate is the builder for creating a Class entity.
type ClassCreate struct {
	config
	mutation *ClassMutation
	hooks    []Hook
}

// SetMid sets the "mid" field.
func (cc *ClassCreate) SetMid(u uint16) *ClassCreate {
	cc.mutation.SetMid(u)
	return cc
}

// SetGrade sets the "grade" field.
func (cc *ClassCreate) SetGrade(s string) *ClassCreate {
	cc.mutation.SetGrade(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ClassCreate) SetName(s string) *ClassCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetIsDisabled sets the "is_disabled" field.
func (cc *ClassCreate) SetIsDisabled(b bool) *ClassCreate {
	cc.mutation.SetIsDisabled(b)
	return cc
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (cc *ClassCreate) SetNillableIsDisabled(b *bool) *ClassCreate {
	if b != nil {
		cc.SetIsDisabled(*b)
	}
	return cc
}

// SetCreatedTime sets the "created_time" field.
func (cc *ClassCreate) SetCreatedTime(t time.Time) *ClassCreate {
	cc.mutation.SetCreatedTime(t)
	return cc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (cc *ClassCreate) SetNillableCreatedTime(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetCreatedTime(*t)
	}
	return cc
}

// SetDeletedTime sets the "deleted_time" field.
func (cc *ClassCreate) SetDeletedTime(t time.Time) *ClassCreate {
	cc.mutation.SetDeletedTime(t)
	return cc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (cc *ClassCreate) SetNillableDeletedTime(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetDeletedTime(*t)
	}
	return cc
}

// SetModifiedTime sets the "modified_time" field.
func (cc *ClassCreate) SetModifiedTime(t time.Time) *ClassCreate {
	cc.mutation.SetModifiedTime(t)
	return cc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (cc *ClassCreate) SetNillableModifiedTime(t *time.Time) *ClassCreate {
	if t != nil {
		cc.SetModifiedTime(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ClassCreate) SetID(u uint32) *ClassCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetMajorID sets the "major" edge to the Major entity by ID.
func (cc *ClassCreate) SetMajorID(id uint16) *ClassCreate {
	cc.mutation.SetMajorID(id)
	return cc
}

// SetMajor sets the "major" edge to the Major entity.
func (cc *ClassCreate) SetMajor(m *Major) *ClassCreate {
	return cc.SetMajorID(m.ID)
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (cc *ClassCreate) AddStudentIDs(ids ...uint32) *ClassCreate {
	cc.mutation.AddStudentIDs(ids...)
	return cc
}

// AddStudents adds the "students" edges to the Student entity.
func (cc *ClassCreate) AddStudents(s ...*Student) *ClassCreate {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStudentIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cc *ClassCreate) Mutation() *ClassMutation {
	return cc.mutation
}

// Save creates the Class in the database.
func (cc *ClassCreate) Save(ctx context.Context) (*Class, error) {
	cc.defaults()
	return withHooks[*Class, ClassMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassCreate) SaveX(ctx context.Context) *Class {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClassCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClassCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClassCreate) defaults() {
	if _, ok := cc.mutation.IsDisabled(); !ok {
		v := class.DefaultIsDisabled
		cc.mutation.SetIsDisabled(v)
	}
	if _, ok := cc.mutation.CreatedTime(); !ok {
		v := class.DefaultCreatedTime()
		cc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClassCreate) check() error {
	if _, ok := cc.mutation.Mid(); !ok {
		return &ValidationError{Name: "mid", err: errors.New(`ent: missing required field "Class.mid"`)}
	}
	if _, ok := cc.mutation.Grade(); !ok {
		return &ValidationError{Name: "grade", err: errors.New(`ent: missing required field "Class.grade"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Class.name"`)}
	}
	if _, ok := cc.mutation.IsDisabled(); !ok {
		return &ValidationError{Name: "is_disabled", err: errors.New(`ent: missing required field "Class.is_disabled"`)}
	}
	if _, ok := cc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "Class.created_time"`)}
	}
	if _, ok := cc.mutation.MajorID(); !ok {
		return &ValidationError{Name: "major", err: errors.New(`ent: missing required edge "Class.major"`)}
	}
	return nil
}

func (cc *ClassCreate) sqlSave(ctx context.Context) (*Class, error) {
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
		_node.ID = uint32(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ClassCreate) createSpec() (*Class, *sqlgraph.CreateSpec) {
	var (
		_node = &Class{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(class.Table, sqlgraph.NewFieldSpec(class.FieldID, field.TypeUint32))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Grade(); ok {
		_spec.SetField(class.FieldGrade, field.TypeString, value)
		_node.Grade = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(class.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.IsDisabled(); ok {
		_spec.SetField(class.FieldIsDisabled, field.TypeBool, value)
		_node.IsDisabled = value
	}
	if value, ok := cc.mutation.CreatedTime(); ok {
		_spec.SetField(class.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := cc.mutation.DeletedTime(); ok {
		_spec.SetField(class.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := cc.mutation.ModifiedTime(); ok {
		_spec.SetField(class.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := cc.mutation.MajorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   class.MajorTable,
			Columns: []string{class.MajorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Mid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StudentsTable,
			Columns: []string{class.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassCreateBulk is the builder for creating many Class entities in bulk.
type ClassCreateBulk struct {
	config
	builders []*ClassCreate
}

// Save creates the Class entities in the database.
func (ccb *ClassCreateBulk) Save(ctx context.Context) ([]*Class, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Class, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassMutation)
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
					nodes[i].ID = uint32(id)
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
func (ccb *ClassCreateBulk) SaveX(ctx context.Context) []*Class {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClassCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClassCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
