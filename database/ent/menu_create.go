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
	"github.com/DemoonLXW/up_learning/database/ent/menu"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/entity"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
	conflict []sql.ConflictOption
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
	if _, ok := mc.mutation.DeletedTime(); !ok {
		v := menu.DefaultDeletedTime
		mc.mutation.SetDeletedTime(v)
	}
	if _, ok := mc.mutation.ModifiedTime(); !ok {
		v := menu.DefaultModifiedTime
		mc.mutation.SetModifiedTime(v)
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
	if _, ok := mc.mutation.DeletedTime(); !ok {
		return &ValidationError{Name: "deleted_time", err: errors.New(`ent: missing required field "Menu.deleted_time"`)}
	}
	if _, ok := mc.mutation.ModifiedTime(); !ok {
		return &ValidationError{Name: "modified_time", err: errors.New(`ent: missing required field "Menu.modified_time"`)}
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
	_spec.OnConflict = mc.conflict
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
			Rel:     sqlgraph.O2O,
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Menu.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MenuUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (mc *MenuCreate) OnConflict(opts ...sql.ConflictOption) *MenuUpsertOne {
	mc.conflict = opts
	return &MenuUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Menu.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MenuCreate) OnConflictColumns(columns ...string) *MenuUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MenuUpsertOne{
		create: mc,
	}
}

type (
	// MenuUpsertOne is the builder for "upsert"-ing
	//  one Menu node.
	MenuUpsertOne struct {
		create *MenuCreate
	}

	// MenuUpsert is the "OnConflict" setter.
	MenuUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *MenuUpsert) SetName(v string) *MenuUpsert {
	u.Set(menu.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MenuUpsert) UpdateName() *MenuUpsert {
	u.SetExcluded(menu.FieldName)
	return u
}

// SetJSONMenu sets the "json_menu" field.
func (u *MenuUpsert) SetJSONMenu(v []*entity.Menu) *MenuUpsert {
	u.Set(menu.FieldJSONMenu, v)
	return u
}

// UpdateJSONMenu sets the "json_menu" field to the value that was provided on create.
func (u *MenuUpsert) UpdateJSONMenu() *MenuUpsert {
	u.SetExcluded(menu.FieldJSONMenu)
	return u
}

// ClearJSONMenu clears the value of the "json_menu" field.
func (u *MenuUpsert) ClearJSONMenu() *MenuUpsert {
	u.SetNull(menu.FieldJSONMenu)
	return u
}

// SetRid sets the "rid" field.
func (u *MenuUpsert) SetRid(v uint8) *MenuUpsert {
	u.Set(menu.FieldRid, v)
	return u
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *MenuUpsert) UpdateRid() *MenuUpsert {
	u.SetExcluded(menu.FieldRid)
	return u
}

// SetCreatedTime sets the "created_time" field.
func (u *MenuUpsert) SetCreatedTime(v time.Time) *MenuUpsert {
	u.Set(menu.FieldCreatedTime, v)
	return u
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *MenuUpsert) UpdateCreatedTime() *MenuUpsert {
	u.SetExcluded(menu.FieldCreatedTime)
	return u
}

// SetDeletedTime sets the "deleted_time" field.
func (u *MenuUpsert) SetDeletedTime(v time.Time) *MenuUpsert {
	u.Set(menu.FieldDeletedTime, v)
	return u
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *MenuUpsert) UpdateDeletedTime() *MenuUpsert {
	u.SetExcluded(menu.FieldDeletedTime)
	return u
}

// SetModifiedTime sets the "modified_time" field.
func (u *MenuUpsert) SetModifiedTime(v time.Time) *MenuUpsert {
	u.Set(menu.FieldModifiedTime, v)
	return u
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *MenuUpsert) UpdateModifiedTime() *MenuUpsert {
	u.SetExcluded(menu.FieldModifiedTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Menu.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(menu.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MenuUpsertOne) UpdateNewValues() *MenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(menu.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Menu.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MenuUpsertOne) Ignore() *MenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MenuUpsertOne) DoNothing() *MenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MenuCreate.OnConflict
// documentation for more info.
func (u *MenuUpsertOne) Update(set func(*MenuUpsert)) *MenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MenuUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MenuUpsertOne) SetName(v string) *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MenuUpsertOne) UpdateName() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateName()
	})
}

// SetJSONMenu sets the "json_menu" field.
func (u *MenuUpsertOne) SetJSONMenu(v []*entity.Menu) *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.SetJSONMenu(v)
	})
}

// UpdateJSONMenu sets the "json_menu" field to the value that was provided on create.
func (u *MenuUpsertOne) UpdateJSONMenu() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateJSONMenu()
	})
}

// ClearJSONMenu clears the value of the "json_menu" field.
func (u *MenuUpsertOne) ClearJSONMenu() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.ClearJSONMenu()
	})
}

// SetRid sets the "rid" field.
func (u *MenuUpsertOne) SetRid(v uint8) *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.SetRid(v)
	})
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *MenuUpsertOne) UpdateRid() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateRid()
	})
}

// SetCreatedTime sets the "created_time" field.
func (u *MenuUpsertOne) SetCreatedTime(v time.Time) *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.SetCreatedTime(v)
	})
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *MenuUpsertOne) UpdateCreatedTime() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateCreatedTime()
	})
}

// SetDeletedTime sets the "deleted_time" field.
func (u *MenuUpsertOne) SetDeletedTime(v time.Time) *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.SetDeletedTime(v)
	})
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *MenuUpsertOne) UpdateDeletedTime() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateDeletedTime()
	})
}

// SetModifiedTime sets the "modified_time" field.
func (u *MenuUpsertOne) SetModifiedTime(v time.Time) *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.SetModifiedTime(v)
	})
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *MenuUpsertOne) UpdateModifiedTime() *MenuUpsertOne {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateModifiedTime()
	})
}

// Exec executes the query.
func (u *MenuUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MenuCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MenuUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MenuUpsertOne) ID(ctx context.Context) (id uint8, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MenuUpsertOne) IDX(ctx context.Context) uint8 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	builders []*MenuCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = mcb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Menu.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MenuUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (mcb *MenuCreateBulk) OnConflict(opts ...sql.ConflictOption) *MenuUpsertBulk {
	mcb.conflict = opts
	return &MenuUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Menu.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MenuCreateBulk) OnConflictColumns(columns ...string) *MenuUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MenuUpsertBulk{
		create: mcb,
	}
}

// MenuUpsertBulk is the builder for "upsert"-ing
// a bulk of Menu nodes.
type MenuUpsertBulk struct {
	create *MenuCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Menu.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(menu.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MenuUpsertBulk) UpdateNewValues() *MenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(menu.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Menu.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MenuUpsertBulk) Ignore() *MenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MenuUpsertBulk) DoNothing() *MenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MenuCreateBulk.OnConflict
// documentation for more info.
func (u *MenuUpsertBulk) Update(set func(*MenuUpsert)) *MenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MenuUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MenuUpsertBulk) SetName(v string) *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MenuUpsertBulk) UpdateName() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateName()
	})
}

// SetJSONMenu sets the "json_menu" field.
func (u *MenuUpsertBulk) SetJSONMenu(v []*entity.Menu) *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.SetJSONMenu(v)
	})
}

// UpdateJSONMenu sets the "json_menu" field to the value that was provided on create.
func (u *MenuUpsertBulk) UpdateJSONMenu() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateJSONMenu()
	})
}

// ClearJSONMenu clears the value of the "json_menu" field.
func (u *MenuUpsertBulk) ClearJSONMenu() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.ClearJSONMenu()
	})
}

// SetRid sets the "rid" field.
func (u *MenuUpsertBulk) SetRid(v uint8) *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.SetRid(v)
	})
}

// UpdateRid sets the "rid" field to the value that was provided on create.
func (u *MenuUpsertBulk) UpdateRid() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateRid()
	})
}

// SetCreatedTime sets the "created_time" field.
func (u *MenuUpsertBulk) SetCreatedTime(v time.Time) *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.SetCreatedTime(v)
	})
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *MenuUpsertBulk) UpdateCreatedTime() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateCreatedTime()
	})
}

// SetDeletedTime sets the "deleted_time" field.
func (u *MenuUpsertBulk) SetDeletedTime(v time.Time) *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.SetDeletedTime(v)
	})
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *MenuUpsertBulk) UpdateDeletedTime() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateDeletedTime()
	})
}

// SetModifiedTime sets the "modified_time" field.
func (u *MenuUpsertBulk) SetModifiedTime(v time.Time) *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.SetModifiedTime(v)
	})
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *MenuUpsertBulk) UpdateModifiedTime() *MenuUpsertBulk {
	return u.Update(func(s *MenuUpsert) {
		s.UpdateModifiedTime()
	})
}

// Exec executes the query.
func (u *MenuUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MenuCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MenuCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MenuUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}