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
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetPhone sets the "phone" field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetIntroduction sets the "introduction" field.
func (uc *UserCreate) SetIntroduction(s string) *UserCreate {
	uc.mutation.SetIntroduction(s)
	return uc
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (uc *UserCreate) SetNillableIntroduction(s *string) *UserCreate {
	if s != nil {
		uc.SetIntroduction(*s)
	}
	return uc
}

// SetCreatedTime sets the "created_time" field.
func (uc *UserCreate) SetCreatedTime(t time.Time) *UserCreate {
	uc.mutation.SetCreatedTime(t)
	return uc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedTime(*t)
	}
	return uc
}

// SetDeletedTime sets the "deleted_time" field.
func (uc *UserCreate) SetDeletedTime(t time.Time) *UserCreate {
	uc.mutation.SetDeletedTime(t)
	return uc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetDeletedTime(*t)
	}
	return uc
}

// SetModifiedTime sets the "modified_time" field.
func (uc *UserCreate) SetModifiedTime(t time.Time) *UserCreate {
	uc.mutation.SetModifiedTime(t)
	return uc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (uc *UserCreate) SetNillableModifiedTime(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetModifiedTime(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uint32) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uc *UserCreate) AddRoleIDs(ids ...uint8) *UserCreate {
	uc.mutation.AddRoleIDs(ids...)
	return uc
}

// AddRoles adds the "roles" edges to the Role entity.
func (uc *UserCreate) AddRoles(r ...*Role) *UserCreate {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoleIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedTime(); !ok {
		v := user.DefaultCreatedTime()
		uc.mutation.SetCreatedTime(v)
	}
	if _, ok := uc.mutation.DeletedTime(); !ok {
		v := user.DefaultDeletedTime
		uc.mutation.SetDeletedTime(v)
	}
	if _, ok := uc.mutation.ModifiedTime(); !ok {
		v := user.DefaultModifiedTime
		uc.mutation.SetModifiedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if v, ok := uc.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "User.password"`)}
	}
	if v, ok := uc.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "User.created_time"`)}
	}
	if _, ok := uc.mutation.DeletedTime(); !ok {
		return &ValidationError{Name: "deleted_time", err: errors.New(`ent: missing required field "User.deleted_time"`)}
	}
	if _, ok := uc.mutation.ModifiedTime(); !ok {
		return &ValidationError{Name: "modified_time", err: errors.New(`ent: missing required field "User.modified_time"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = &value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
		_node.Password = &value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = &value
	}
	if value, ok := uc.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
		_node.Phone = &value
	}
	if value, ok := uc.mutation.Introduction(); ok {
		_spec.SetField(user.FieldIntroduction, field.TypeString, value)
		_node.Introduction = &value
	}
	if value, ok := uc.mutation.CreatedTime(); ok {
		_spec.SetField(user.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := uc.mutation.DeletedTime(); ok {
		_spec.SetField(user.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := uc.mutation.ModifiedTime(); ok {
		_spec.SetField(user.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := uc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserRoleCreate{config: uc.config, mutation: newUserRoleMutation(uc.config, OpCreate)}
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
//	client.User.Create().
//		SetUsername(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUsername sets the "username" field.
func (u *UserUpsert) SetUsername(v string) *UserUpsert {
	u.Set(user.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsert) UpdateUsername() *UserUpsert {
	u.SetExcluded(user.FieldUsername)
	return u
}

// SetPassword sets the "password" field.
func (u *UserUpsert) SetPassword(v string) *UserUpsert {
	u.Set(user.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsert) UpdatePassword() *UserUpsert {
	u.SetExcluded(user.FieldPassword)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsert) ClearEmail() *UserUpsert {
	u.SetNull(user.FieldEmail)
	return u
}

// SetPhone sets the "phone" field.
func (u *UserUpsert) SetPhone(v string) *UserUpsert {
	u.Set(user.FieldPhone, v)
	return u
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *UserUpsert) UpdatePhone() *UserUpsert {
	u.SetExcluded(user.FieldPhone)
	return u
}

// ClearPhone clears the value of the "phone" field.
func (u *UserUpsert) ClearPhone() *UserUpsert {
	u.SetNull(user.FieldPhone)
	return u
}

// SetIntroduction sets the "introduction" field.
func (u *UserUpsert) SetIntroduction(v string) *UserUpsert {
	u.Set(user.FieldIntroduction, v)
	return u
}

// UpdateIntroduction sets the "introduction" field to the value that was provided on create.
func (u *UserUpsert) UpdateIntroduction() *UserUpsert {
	u.SetExcluded(user.FieldIntroduction)
	return u
}

// ClearIntroduction clears the value of the "introduction" field.
func (u *UserUpsert) ClearIntroduction() *UserUpsert {
	u.SetNull(user.FieldIntroduction)
	return u
}

// SetCreatedTime sets the "created_time" field.
func (u *UserUpsert) SetCreatedTime(v time.Time) *UserUpsert {
	u.Set(user.FieldCreatedTime, v)
	return u
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateCreatedTime() *UserUpsert {
	u.SetExcluded(user.FieldCreatedTime)
	return u
}

// SetDeletedTime sets the "deleted_time" field.
func (u *UserUpsert) SetDeletedTime(v time.Time) *UserUpsert {
	u.Set(user.FieldDeletedTime, v)
	return u
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateDeletedTime() *UserUpsert {
	u.SetExcluded(user.FieldDeletedTime)
	return u
}

// SetModifiedTime sets the "modified_time" field.
func (u *UserUpsert) SetModifiedTime(v time.Time) *UserUpsert {
	u.Set(user.FieldModifiedTime, v)
	return u
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *UserUpsert) UpdateModifiedTime() *UserUpsert {
	u.SetExcluded(user.FieldModifiedTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *UserUpsertOne) SetUsername(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUsername() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertOne) SetPassword(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertOne) ClearEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetPhone sets the "phone" field.
func (u *UserUpsertOne) SetPhone(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *UserUpsertOne) UpdatePhone() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePhone()
	})
}

// ClearPhone clears the value of the "phone" field.
func (u *UserUpsertOne) ClearPhone() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearPhone()
	})
}

// SetIntroduction sets the "introduction" field.
func (u *UserUpsertOne) SetIntroduction(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetIntroduction(v)
	})
}

// UpdateIntroduction sets the "introduction" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateIntroduction() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateIntroduction()
	})
}

// ClearIntroduction clears the value of the "introduction" field.
func (u *UserUpsertOne) ClearIntroduction() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearIntroduction()
	})
}

// SetCreatedTime sets the "created_time" field.
func (u *UserUpsertOne) SetCreatedTime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedTime(v)
	})
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCreatedTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedTime()
	})
}

// SetDeletedTime sets the "deleted_time" field.
func (u *UserUpsertOne) SetDeletedTime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedTime(v)
	})
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDeletedTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedTime()
	})
}

// SetModifiedTime sets the "modified_time" field.
func (u *UserUpsertOne) SetModifiedTime(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetModifiedTime(v)
	})
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateModifiedTime() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateModifiedTime()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *UserUpsertBulk) SetUsername(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUsername() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// SetPassword sets the "password" field.
func (u *UserUpsertBulk) SetPassword(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePassword()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertBulk) ClearEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetPhone sets the "phone" field.
func (u *UserUpsertBulk) SetPhone(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdatePhone() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdatePhone()
	})
}

// ClearPhone clears the value of the "phone" field.
func (u *UserUpsertBulk) ClearPhone() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearPhone()
	})
}

// SetIntroduction sets the "introduction" field.
func (u *UserUpsertBulk) SetIntroduction(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetIntroduction(v)
	})
}

// UpdateIntroduction sets the "introduction" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateIntroduction() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateIntroduction()
	})
}

// ClearIntroduction clears the value of the "introduction" field.
func (u *UserUpsertBulk) ClearIntroduction() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearIntroduction()
	})
}

// SetCreatedTime sets the "created_time" field.
func (u *UserUpsertBulk) SetCreatedTime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedTime(v)
	})
}

// UpdateCreatedTime sets the "created_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCreatedTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedTime()
	})
}

// SetDeletedTime sets the "deleted_time" field.
func (u *UserUpsertBulk) SetDeletedTime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedTime(v)
	})
}

// UpdateDeletedTime sets the "deleted_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDeletedTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedTime()
	})
}

// SetModifiedTime sets the "modified_time" field.
func (u *UserUpsertBulk) SetModifiedTime(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetModifiedTime(v)
	})
}

// UpdateModifiedTime sets the "modified_time" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateModifiedTime() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateModifiedTime()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
