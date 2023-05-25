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
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAction sets the "action" field.
func (pc *PermissionCreate) SetAction(s string) *PermissionCreate {
	pc.mutation.SetAction(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PermissionCreate) SetDescription(s string) *PermissionCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDescription(s *string) *PermissionCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetCreatedTime sets the "created_time" field.
func (pc *PermissionCreate) SetCreatedTime(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedTime(t)
	return pc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedTime(*t)
	}
	return pc
}

// SetDeletedTime sets the "deleted_time" field.
func (pc *PermissionCreate) SetDeletedTime(t time.Time) *PermissionCreate {
	pc.mutation.SetDeletedTime(t)
	return pc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDeletedTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetDeletedTime(*t)
	}
	return pc
}

// SetModifiedTime sets the "modified_time" field.
func (pc *PermissionCreate) SetModifiedTime(t time.Time) *PermissionCreate {
	pc.mutation.SetModifiedTime(t)
	return pc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableModifiedTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetModifiedTime(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(u uint16) *PermissionCreate {
	pc.mutation.SetID(u)
	return pc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pc *PermissionCreate) AddRoleIDs(ids ...uint8) *PermissionCreate {
	pc.mutation.AddRoleIDs(ids...)
	return pc
}

// AddRoles adds the "roles" edges to the Role entity.
func (pc *PermissionCreate) AddRoles(r ...*Role) *PermissionCreate {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermissionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermissionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.CreatedTime(); !ok {
		v := permission.DefaultCreatedTime()
		pc.mutation.SetCreatedTime(v)
	}
	if _, ok := pc.mutation.DeletedTime(); !ok {
		v := permission.DefaultDeletedTime
		pc.mutation.SetDeletedTime(v)
	}
	if _, ok := pc.mutation.ModifiedTime(); !ok {
		v := permission.DefaultModifiedTime
		pc.mutation.SetModifiedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "Permission.action"`)}
	}
	if v, ok := pc.mutation.Action(); ok {
		if err := permission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Permission.action": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "Permission.created_time"`)}
	}
	if _, ok := pc.mutation.DeletedTime(); !ok {
		return &ValidationError{Name: "deleted_time", err: errors.New(`ent: missing required field "Permission.deleted_time"`)}
	}
	if _, ok := pc.mutation.ModifiedTime(); !ok {
		return &ValidationError{Name: "modified_time", err: errors.New(`ent: missing required field "Permission.modified_time"`)}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint16(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(permission.Table, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUint16))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
		_node.Action = &value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := pc.mutation.CreatedTime(); ok {
		_spec.SetField(permission.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := pc.mutation.DeletedTime(); ok {
		_spec.SetField(permission.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := pc.mutation.ModifiedTime(); ok {
		_spec.SetField(permission.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := pc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RolePermissionCreate{config: pc.config, mutation: newRolePermissionMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Permission.Create().
//		SetAction(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PermissionUpsert) {
//			SetAction(v+v).
//		}).
//		Exec(ctx)
func (pc *PermissionCreate) OnConflict(opts ...sql.ConflictOption) *PermissionUpsertOne {
	pc.conflict = opts
	return &PermissionUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PermissionCreate) OnConflictColumns(columns ...string) *PermissionUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PermissionUpsertOne{
		create: pc,
	}
}

type (
	// PermissionUpsertOne is the builder for "upsert"-ing
	//  one Permission node.
	PermissionUpsertOne struct {
		create *PermissionCreate
	}

	// PermissionUpsert is the "OnConflict" setter.
	PermissionUpsert struct {
		*sql.UpdateSet
	}
)

// SetAction sets the "action" field.
func (u *PermissionUpsert) SetAction(v string) *PermissionUpsert {
	u.Set(permission.FieldAction, v)
	return u
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateAction() *PermissionUpsert {
	u.SetExcluded(permission.FieldAction)
	return u
}

// SetDescription sets the "description" field.
func (u *PermissionUpsert) SetDescription(v string) *PermissionUpsert {
	u.Set(permission.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateDescription() *PermissionUpsert {
	u.SetExcluded(permission.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *PermissionUpsert) ClearDescription() *PermissionUpsert {
	u.SetNull(permission.FieldDescription)
	return u
}

// SetCreatedTime sets the "created_time" field.
func (u *PermissionUpsert) SetCreatedTime(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldCreatedTime, v)
	return u
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateCreatedTime() *PermissionUpsert {
	u.SetExcluded(permission.FieldCreatedTime)
	return u
}

// SetDeletedTime sets the "deleted_time" field.
func (u *PermissionUpsert) SetDeletedTime(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldDeletedTime, v)
	return u
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateDeletedTime() *PermissionUpsert {
	u.SetExcluded(permission.FieldDeletedTime)
	return u
}

// SetModifiedTime sets the "modified_time" field.
func (u *PermissionUpsert) SetModifiedTime(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldModifiedTime, v)
	return u
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateModifiedTime() *PermissionUpsert {
	u.SetExcluded(permission.FieldModifiedTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(permission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PermissionUpsertOne) UpdateNewValues() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(permission.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PermissionUpsertOne) Ignore() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PermissionUpsertOne) DoNothing() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PermissionCreate.OnConflict
// documentation for more info.
func (u *PermissionUpsertOne) Update(set func(*PermissionUpsert)) *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PermissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetAction sets the "action" field.
func (u *PermissionUpsertOne) SetAction(v string) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateAction() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateAction()
	})
}

// SetDescription sets the "description" field.
func (u *PermissionUpsertOne) SetDescription(v string) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateDescription() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PermissionUpsertOne) ClearDescription() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearDescription()
	})
}

// SetCreatedTime sets the "created_time" field.
func (u *PermissionUpsertOne) SetCreatedTime(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetCreatedTime(v)
	})
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateCreatedTime() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateCreatedTime()
	})
}

// SetDeletedTime sets the "deleted_time" field.
func (u *PermissionUpsertOne) SetDeletedTime(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetDeletedTime(v)
	})
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateDeletedTime() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateDeletedTime()
	})
}

// SetModifiedTime sets the "modified_time" field.
func (u *PermissionUpsertOne) SetModifiedTime(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetModifiedTime(v)
	})
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateModifiedTime() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateModifiedTime()
	})
}

// Exec executes the query.
func (u *PermissionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PermissionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PermissionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PermissionUpsertOne) ID(ctx context.Context) (id uint16, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PermissionUpsertOne) IDX(ctx context.Context) uint16 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
	conflict []sql.ConflictOption
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermissionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Permission.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PermissionUpsert) {
//			SetAction(v+v).
//		}).
//		Exec(ctx)
func (pcb *PermissionCreateBulk) OnConflict(opts ...sql.ConflictOption) *PermissionUpsertBulk {
	pcb.conflict = opts
	return &PermissionUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PermissionCreateBulk) OnConflictColumns(columns ...string) *PermissionUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PermissionUpsertBulk{
		create: pcb,
	}
}

// PermissionUpsertBulk is the builder for "upsert"-ing
// a bulk of Permission nodes.
type PermissionUpsertBulk struct {
	create *PermissionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(permission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PermissionUpsertBulk) UpdateNewValues() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(permission.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PermissionUpsertBulk) Ignore() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PermissionUpsertBulk) DoNothing() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PermissionCreateBulk.OnConflict
// documentation for more info.
func (u *PermissionUpsertBulk) Update(set func(*PermissionUpsert)) *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PermissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetAction sets the "action" field.
func (u *PermissionUpsertBulk) SetAction(v string) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateAction() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateAction()
	})
}

// SetDescription sets the "description" field.
func (u *PermissionUpsertBulk) SetDescription(v string) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateDescription() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PermissionUpsertBulk) ClearDescription() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearDescription()
	})
}

// SetCreatedTime sets the "created_time" field.
func (u *PermissionUpsertBulk) SetCreatedTime(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetCreatedTime(v)
	})
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateCreatedTime() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateCreatedTime()
	})
}

// SetDeletedTime sets the "deleted_time" field.
func (u *PermissionUpsertBulk) SetDeletedTime(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetDeletedTime(v)
	})
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateDeletedTime() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateDeletedTime()
	})
}

// SetModifiedTime sets the "modified_time" field.
func (u *PermissionUpsertBulk) SetModifiedTime(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetModifiedTime(v)
	})
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateModifiedTime() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateModifiedTime()
	})
}

// Exec executes the query.
func (u *PermissionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PermissionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PermissionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PermissionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
