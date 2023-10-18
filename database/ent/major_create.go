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
	"github.com/DemoonLXW/up_learning/database/ent/college"
	"github.com/DemoonLXW/up_learning/database/ent/major"
)

// MajorCreate is the builder for creating a Major entity.
type MajorCreate struct {
	config
	mutation *MajorMutation
	hooks    []Hook
}

// SetCid sets the "cid" field.
func (mc *MajorCreate) SetCid(u uint8) *MajorCreate {
	mc.mutation.SetCid(u)
	return mc
}

// SetName sets the "name" field.
func (mc *MajorCreate) SetName(s string) *MajorCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetIsDisabled sets the "is_disabled" field.
func (mc *MajorCreate) SetIsDisabled(b bool) *MajorCreate {
	mc.mutation.SetIsDisabled(b)
	return mc
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (mc *MajorCreate) SetNillableIsDisabled(b *bool) *MajorCreate {
	if b != nil {
		mc.SetIsDisabled(*b)
	}
	return mc
}

// SetCreatedTime sets the "created_time" field.
func (mc *MajorCreate) SetCreatedTime(t time.Time) *MajorCreate {
	mc.mutation.SetCreatedTime(t)
	return mc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (mc *MajorCreate) SetNillableCreatedTime(t *time.Time) *MajorCreate {
	if t != nil {
		mc.SetCreatedTime(*t)
	}
	return mc
}

// SetDeletedTime sets the "deleted_time" field.
func (mc *MajorCreate) SetDeletedTime(t time.Time) *MajorCreate {
	mc.mutation.SetDeletedTime(t)
	return mc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (mc *MajorCreate) SetNillableDeletedTime(t *time.Time) *MajorCreate {
	if t != nil {
		mc.SetDeletedTime(*t)
	}
	return mc
}

// SetModifiedTime sets the "modified_time" field.
func (mc *MajorCreate) SetModifiedTime(t time.Time) *MajorCreate {
	mc.mutation.SetModifiedTime(t)
	return mc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (mc *MajorCreate) SetNillableModifiedTime(t *time.Time) *MajorCreate {
	if t != nil {
		mc.SetModifiedTime(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MajorCreate) SetID(u uint16) *MajorCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetCollegeID sets the "college" edge to the College entity by ID.
func (mc *MajorCreate) SetCollegeID(id uint8) *MajorCreate {
	mc.mutation.SetCollegeID(id)
	return mc
}

// SetCollege sets the "college" edge to the College entity.
func (mc *MajorCreate) SetCollege(c *College) *MajorCreate {
	return mc.SetCollegeID(c.ID)
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (mc *MajorCreate) AddClassIDs(ids ...uint32) *MajorCreate {
	mc.mutation.AddClassIDs(ids...)
	return mc
}

// AddClasses adds the "classes" edges to the Class entity.
func (mc *MajorCreate) AddClasses(c ...*Class) *MajorCreate {
	ids := make([]uint32, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mc.AddClassIDs(ids...)
}

// Mutation returns the MajorMutation object of the builder.
func (mc *MajorCreate) Mutation() *MajorMutation {
	return mc.mutation
}

// Save creates the Major in the database.
func (mc *MajorCreate) Save(ctx context.Context) (*Major, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MajorCreate) SaveX(ctx context.Context) *Major {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MajorCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MajorCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MajorCreate) defaults() {
	if _, ok := mc.mutation.IsDisabled(); !ok {
		v := major.DefaultIsDisabled
		mc.mutation.SetIsDisabled(v)
	}
	if _, ok := mc.mutation.CreatedTime(); !ok {
		v := major.DefaultCreatedTime()
		mc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MajorCreate) check() error {
	if _, ok := mc.mutation.Cid(); !ok {
		return &ValidationError{Name: "cid", err: errors.New(`ent: missing required field "Major.cid"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Major.name"`)}
	}
	if _, ok := mc.mutation.IsDisabled(); !ok {
		return &ValidationError{Name: "is_disabled", err: errors.New(`ent: missing required field "Major.is_disabled"`)}
	}
	if _, ok := mc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "Major.created_time"`)}
	}
	if _, ok := mc.mutation.CollegeID(); !ok {
		return &ValidationError{Name: "college", err: errors.New(`ent: missing required edge "Major.college"`)}
	}
	return nil
}

func (mc *MajorCreate) sqlSave(ctx context.Context) (*Major, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint16(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MajorCreate) createSpec() (*Major, *sqlgraph.CreateSpec) {
	var (
		_node = &Major{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(major.Table, sqlgraph.NewFieldSpec(major.FieldID, field.TypeUint16))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(major.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.IsDisabled(); ok {
		_spec.SetField(major.FieldIsDisabled, field.TypeBool, value)
		_node.IsDisabled = value
	}
	if value, ok := mc.mutation.CreatedTime(); ok {
		_spec.SetField(major.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := mc.mutation.DeletedTime(); ok {
		_spec.SetField(major.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := mc.mutation.ModifiedTime(); ok {
		_spec.SetField(major.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := mc.mutation.CollegeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   major.CollegeTable,
			Columns: []string{major.CollegeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Cid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   major.ClassesTable,
			Columns: []string{major.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MajorCreateBulk is the builder for creating many Major entities in bulk.
type MajorCreateBulk struct {
	config
	builders []*MajorCreate
}

// Save creates the Major entities in the database.
func (mcb *MajorCreateBulk) Save(ctx context.Context) ([]*Major, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Major, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MajorMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
					nodes[i].ID = uint16(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MajorCreateBulk) SaveX(ctx context.Context) []*Major {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MajorCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MajorCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}