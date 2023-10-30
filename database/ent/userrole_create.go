// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/database/ent/userrole"
)

// UserRoleCreate is the builder for creating a UserRole entity.
type UserRoleCreate struct {
	config
	mutation *UserRoleMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (urc *UserRoleCreate) SetUID(u uint32) *UserRoleCreate {
	urc.mutation.SetUID(u)
	return urc
}

// SetRid sets the "rid" field.
func (urc *UserRoleCreate) SetRid(u uint8) *UserRoleCreate {
	urc.mutation.SetRid(u)
	return urc
}

// SetCreatedTime sets the "created_time" field.
func (urc *UserRoleCreate) SetCreatedTime(t time.Time) *UserRoleCreate {
	urc.mutation.SetCreatedTime(t)
	return urc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (urc *UserRoleCreate) SetNillableCreatedTime(t *time.Time) *UserRoleCreate {
	if t != nil {
		urc.SetCreatedTime(*t)
	}
	return urc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (urc *UserRoleCreate) SetUserID(id uint32) *UserRoleCreate {
	urc.mutation.SetUserID(id)
	return urc
}

// SetUser sets the "user" edge to the User entity.
func (urc *UserRoleCreate) SetUser(u *User) *UserRoleCreate {
	return urc.SetUserID(u.ID)
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (urc *UserRoleCreate) SetRoleID(id uint8) *UserRoleCreate {
	urc.mutation.SetRoleID(id)
	return urc
}

// SetRole sets the "role" edge to the Role entity.
func (urc *UserRoleCreate) SetRole(r *Role) *UserRoleCreate {
	return urc.SetRoleID(r.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (urc *UserRoleCreate) Mutation() *UserRoleMutation {
	return urc.mutation
}

// Save creates the UserRole in the database.
func (urc *UserRoleCreate) Save(ctx context.Context) (*UserRole, error) {
	urc.defaults()
	return withHooks[*UserRole, UserRoleMutation](ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRoleCreate) SaveX(ctx context.Context) *UserRole {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRoleCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRoleCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urc *UserRoleCreate) defaults() {
	if _, ok := urc.mutation.CreatedTime(); !ok {
		v := userrole.DefaultCreatedTime()
		urc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRoleCreate) check() error {
	if _, ok := urc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "UserRole.uid"`)}
	}
	if _, ok := urc.mutation.Rid(); !ok {
		return &ValidationError{Name: "rid", err: errors.New(`ent: missing required field "UserRole.rid"`)}
	}
	if _, ok := urc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "UserRole.created_time"`)}
	}
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserRole.user"`)}
	}
	if _, ok := urc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "UserRole.role"`)}
	}
	return nil
}

func (urc *UserRoleCreate) sqlSave(ctx context.Context) (*UserRole, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (urc *UserRoleCreate) createSpec() (*UserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRole{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(userrole.Table, nil)
	)
	if value, ok := urc.mutation.CreatedTime(); ok {
		_spec.SetField(userrole.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if nodes := urc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := urc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
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

// UserRoleCreateBulk is the builder for creating many UserRole entities in bulk.
type UserRoleCreateBulk struct {
	config
	builders []*UserRoleCreate
}

// Save creates the UserRole entities in the database.
func (urcb *UserRoleCreateBulk) Save(ctx context.Context) ([]*UserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRole, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) SaveX(ctx context.Context) []*UserRole {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
