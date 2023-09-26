// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/menu"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/entity"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetJSONMenu sets the "json_menu" field.
func (mc *MenuCreate) SetJSONMenu(e []*entity.Menu) *MenuCreate {
	mc.mutation.SetJSONMenu(e)
	return mc
}

// SetRid sets the "rid" field.
func (mc *MenuCreate) SetRid(u uint8) *MenuCreate {
	mc.mutation.SetRid(u)
	return mc
}

// SetCreatedTime sets the "created_time" field.
func (mc *MenuCreate) SetCreatedTime(t time.Time) *MenuCreate {
	mc.mutation.SetCreatedTime(t)
	return mc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedTime(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetCreatedTime(*t)
	}
	return mc
}

// SetDeletedTime sets the "deleted_time" field.
func (mc *MenuCreate) SetDeletedTime(t time.Time) *MenuCreate {
	mc.mutation.SetDeletedTime(t)
	return mc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDeletedTime(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetDeletedTime(*t)
	}
	return mc
}

// SetModifiedTime sets the "modified_time" field.
func (mc *MenuCreate) SetModifiedTime(t time.Time) *MenuCreate {
	mc.mutation.SetModifiedTime(t)
	return mc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (mc *MenuCreate) SetNillableModifiedTime(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetModifiedTime(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MenuCreate) SetID(u uint8) *MenuCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (mc *MenuCreate) SetRoleID(id uint8) *MenuCreate {
	mc.mutation.SetRoleID(id)
	return mc
}

// SetRole sets the "role" edge to the Role entity.
func (mc *MenuCreate) SetRole(r *Role) *MenuCreate {
	return mc.SetRoleID(r.ID)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() {
	if _, ok := mc.mutation.CreatedTime(); !ok {
		v := menu.DefaultCreatedTime()
		mc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Menu.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Menu.name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Rid(); !ok {
		return &ValidationError{Name: "rid", err: errors.New(`ent: missing required field "Menu.rid"`)}
	}
	if _, ok := mc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "Menu.created_time"`)}
	}
	if _, ok := mc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "Menu.role"`)}
	}
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
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
		_node.ID = uint8(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(menu.Table, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint8))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := mc.mutation.JSONMenu(); ok {
		_spec.SetField(menu.FieldJSONMenu, field.TypeJSON, value)
		_node.JSONMenu = value
	}
	if value, ok := mc.mutation.CreatedTime(); ok {
		_spec.SetField(menu.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := mc.mutation.DeletedTime(); ok {
		_spec.SetField(menu.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := mc.mutation.ModifiedTime(); ok {
		_spec.SetField(menu.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := mc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.RoleTable,
			Columns: []string{menu.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Rid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
