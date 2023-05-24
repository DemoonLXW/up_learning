// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/rolepermission"
)

// RolePermissionCreate is the builder for creating a RolePermission entity.
type RolePermissionCreate struct {
	config
	mutation *RolePermissionMutation
	hooks    []Hook
}

// SetRid sets the "rid" field.
func (rpc *RolePermissionCreate) SetRid(u uint8) *RolePermissionCreate {
	rpc.mutation.SetRid(u)
	return rpc
}

// SetPid sets the "pid" field.
func (rpc *RolePermissionCreate) SetPid(u uint16) *RolePermissionCreate {
	rpc.mutation.SetPid(u)
	return rpc
}

// SetCreatedTime sets the "created_time" field.
func (rpc *RolePermissionCreate) SetCreatedTime(t time.Time) *RolePermissionCreate {
	rpc.mutation.SetCreatedTime(t)
	return rpc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (rpc *RolePermissionCreate) SetNillableCreatedTime(t *time.Time) *RolePermissionCreate {
	if t != nil {
		rpc.SetCreatedTime(*t)
	}
	return rpc
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (rpc *RolePermissionCreate) SetRoleID(id uint8) *RolePermissionCreate {
	rpc.mutation.SetRoleID(id)
	return rpc
}

// SetRole sets the "role" edge to the Role entity.
func (rpc *RolePermissionCreate) SetRole(r *Role) *RolePermissionCreate {
	return rpc.SetRoleID(r.ID)
}

// SetPermissionID sets the "permission" edge to the Permission entity by ID.
func (rpc *RolePermissionCreate) SetPermissionID(id uint16) *RolePermissionCreate {
	rpc.mutation.SetPermissionID(id)
	return rpc
}

// SetPermission sets the "permission" edge to the Permission entity.
func (rpc *RolePermissionCreate) SetPermission(p *Permission) *RolePermissionCreate {
	return rpc.SetPermissionID(p.ID)
}

// Mutation returns the RolePermissionMutation object of the builder.
func (rpc *RolePermissionCreate) Mutation() *RolePermissionMutation {
	return rpc.mutation
}

// Save creates the RolePermission in the database.
func (rpc *RolePermissionCreate) Save(ctx context.Context) (*RolePermission, error) {
	rpc.defaults()
	return withHooks(ctx, rpc.sqlSave, rpc.mutation, rpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rpc *RolePermissionCreate) SaveX(ctx context.Context) *RolePermission {
	v, err := rpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpc *RolePermissionCreate) Exec(ctx context.Context) error {
	_, err := rpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpc *RolePermissionCreate) ExecX(ctx context.Context) {
	if err := rpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpc *RolePermissionCreate) defaults() {
	if _, ok := rpc.mutation.CreatedTime(); !ok {
		v := rolepermission.DefaultCreatedTime()
		rpc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpc *RolePermissionCreate) check() error {
	if _, ok := rpc.mutation.Rid(); !ok {
		return &ValidationError{Name: "rid", err: errors.New(`ent: missing required field "RolePermission.rid"`)}
	}
	if _, ok := rpc.mutation.Pid(); !ok {
		return &ValidationError{Name: "pid", err: errors.New(`ent: missing required field "RolePermission.pid"`)}
	}
	if _, ok := rpc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "RolePermission.created_time"`)}
	}
	if _, ok := rpc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "RolePermission.role"`)}
	}
	if _, ok := rpc.mutation.PermissionID(); !ok {
		return &ValidationError{Name: "permission", err: errors.New(`ent: missing required edge "RolePermission.permission"`)}
	}
	return nil
}

func (rpc *RolePermissionCreate) sqlSave(ctx context.Context) (*RolePermission, error) {
	if err := rpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (rpc *RolePermissionCreate) createSpec() (*RolePermission, *sqlgraph.CreateSpec) {
	var (
		_node = &RolePermission{config: rpc.config}
		_spec = sqlgraph.NewCreateSpec(rolepermission.Table, nil)
	)
	if value, ok := rpc.mutation.CreatedTime(); ok {
		_spec.SetField(rolepermission.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = value
	}
	if nodes := rpc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rolepermission.RoleTable,
			Columns: []string{rolepermission.RoleColumn},
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
	if nodes := rpc.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rolepermission.PermissionTable,
			Columns: []string{rolepermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUint16),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Pid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RolePermissionCreateBulk is the builder for creating many RolePermission entities in bulk.
type RolePermissionCreateBulk struct {
	config
	builders []*RolePermissionCreate
}

// Save creates the RolePermission entities in the database.
func (rpcb *RolePermissionCreateBulk) Save(ctx context.Context) ([]*RolePermission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rpcb.builders))
	nodes := make([]*RolePermission, len(rpcb.builders))
	mutators := make([]Mutator, len(rpcb.builders))
	for i := range rpcb.builders {
		func(i int, root context.Context) {
			builder := rpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RolePermissionMutation)
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
					_, err = mutators[i+1].Mutate(root, rpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, rpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rpcb *RolePermissionCreateBulk) SaveX(ctx context.Context) []*RolePermission {
	v, err := rpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpcb *RolePermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := rpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpcb *RolePermissionCreateBulk) ExecX(ctx context.Context) {
	if err := rpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
