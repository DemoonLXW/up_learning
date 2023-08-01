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
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/role"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAction sets the "action" field.
func (pu *PermissionUpdate) SetAction(s string) *PermissionUpdate {
	pu.mutation.SetAction(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PermissionUpdate) SetDescription(s string) *PermissionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PermissionUpdate) ClearDescription() *PermissionUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetIsDisabled sets the "is_disabled" field.
func (pu *PermissionUpdate) SetIsDisabled(b bool) *PermissionUpdate {
	pu.mutation.SetIsDisabled(b)
	return pu
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableIsDisabled(b *bool) *PermissionUpdate {
	if b != nil {
		pu.SetIsDisabled(*b)
	}
	return pu
}

// SetCreatedTime sets the "created_time" field.
func (pu *PermissionUpdate) SetCreatedTime(t time.Time) *PermissionUpdate {
	pu.mutation.SetCreatedTime(t)
	return pu
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCreatedTime(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetCreatedTime(*t)
	}
	return pu
}

// SetDeletedTime sets the "deleted_time" field.
func (pu *PermissionUpdate) SetDeletedTime(t time.Time) *PermissionUpdate {
	pu.mutation.SetDeletedTime(t)
	return pu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDeletedTime(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetDeletedTime(*t)
	}
	return pu
}

// SetModifiedTime sets the "modified_time" field.
func (pu *PermissionUpdate) SetModifiedTime(t time.Time) *PermissionUpdate {
	pu.mutation.SetModifiedTime(t)
	return pu
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableModifiedTime(t *time.Time) *PermissionUpdate {
	if t != nil {
		pu.SetModifiedTime(*t)
	}
	return pu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...uint8) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRoles adds the "roles" edges to the Role entity.
func (pu *PermissionUpdate) AddRoles(r ...*Role) *PermissionUpdate {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (pu *PermissionUpdate) ClearRoles() *PermissionUpdate {
	pu.mutation.ClearRoles()
	return pu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...uint8) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRoles removes "roles" edges to Role entities.
func (pu *PermissionUpdate) RemoveRoles(r ...*Role) *PermissionUpdate {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.Action(); ok {
		if err := permission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Permission.action": %w`, err)}
		}
	}
	return nil
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUint16))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.IsDisabled(); ok {
		_spec.SetField(permission.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CreatedTime(); ok {
		_spec.SetField(permission.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedTime(); ok {
		_spec.SetField(permission.FieldDeletedTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.ModifiedTime(); ok {
		_spec.SetField(permission.FieldModifiedTime, field.TypeTime, value)
	}
	if pu.mutation.RolesCleared() {
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
		createE := &RolePermissionCreate{config: pu.config, mutation: newRolePermissionMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !pu.mutation.RolesCleared() {
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
		createE := &RolePermissionCreate{config: pu.config, mutation: newRolePermissionMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RolesIDs(); len(nodes) > 0 {
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
		createE := &RolePermissionCreate{config: pu.config, mutation: newRolePermissionMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetAction sets the "action" field.
func (puo *PermissionUpdateOne) SetAction(s string) *PermissionUpdateOne {
	puo.mutation.SetAction(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PermissionUpdateOne) SetDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PermissionUpdateOne) ClearDescription() *PermissionUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetIsDisabled sets the "is_disabled" field.
func (puo *PermissionUpdateOne) SetIsDisabled(b bool) *PermissionUpdateOne {
	puo.mutation.SetIsDisabled(b)
	return puo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableIsDisabled(b *bool) *PermissionUpdateOne {
	if b != nil {
		puo.SetIsDisabled(*b)
	}
	return puo
}

// SetCreatedTime sets the "created_time" field.
func (puo *PermissionUpdateOne) SetCreatedTime(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetCreatedTime(t)
	return puo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCreatedTime(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetCreatedTime(*t)
	}
	return puo
}

// SetDeletedTime sets the "deleted_time" field.
func (puo *PermissionUpdateOne) SetDeletedTime(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetDeletedTime(t)
	return puo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDeletedTime(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetDeletedTime(*t)
	}
	return puo
}

// SetModifiedTime sets the "modified_time" field.
func (puo *PermissionUpdateOne) SetModifiedTime(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetModifiedTime(t)
	return puo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableModifiedTime(t *time.Time) *PermissionUpdateOne {
	if t != nil {
		puo.SetModifiedTime(*t)
	}
	return puo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...uint8) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRoles adds the "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRoles() *PermissionUpdateOne {
	puo.mutation.ClearRoles()
	return puo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...uint8) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRoles removes "roles" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.Action(); ok {
		if err := permission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Permission.action": %w`, err)}
		}
	}
	return nil
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUint16))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.IsDisabled(); ok {
		_spec.SetField(permission.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CreatedTime(); ok {
		_spec.SetField(permission.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedTime(); ok {
		_spec.SetField(permission.FieldDeletedTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.ModifiedTime(); ok {
		_spec.SetField(permission.FieldModifiedTime, field.TypeTime, value)
	}
	if puo.mutation.RolesCleared() {
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
		createE := &RolePermissionCreate{config: puo.config, mutation: newRolePermissionMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !puo.mutation.RolesCleared() {
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
		createE := &RolePermissionCreate{config: puo.config, mutation: newRolePermissionMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RolesIDs(); len(nodes) > 0 {
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
		createE := &RolePermissionCreate{config: puo.config, mutation: newRolePermissionMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
