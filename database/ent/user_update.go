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
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/student"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetAccount sets the "account" field.
func (uu *UserUpdate) SetAccount(s string) *UserUpdate {
	uu.mutation.SetAccount(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// ClearUsername clears the value of the "username" field.
func (uu *UserUpdate) ClearUsername() *UserUpdate {
	uu.mutation.ClearUsername()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of the "phone" field.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetIntroduction sets the "introduction" field.
func (uu *UserUpdate) SetIntroduction(s string) *UserUpdate {
	uu.mutation.SetIntroduction(s)
	return uu
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIntroduction(s *string) *UserUpdate {
	if s != nil {
		uu.SetIntroduction(*s)
	}
	return uu
}

// ClearIntroduction clears the value of the "introduction" field.
func (uu *UserUpdate) ClearIntroduction() *UserUpdate {
	uu.mutation.ClearIntroduction()
	return uu
}

// SetIsDisabled sets the "is_disabled" field.
func (uu *UserUpdate) SetIsDisabled(b bool) *UserUpdate {
	uu.mutation.SetIsDisabled(b)
	return uu
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsDisabled(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsDisabled(*b)
	}
	return uu
}

// SetCreatedTime sets the "created_time" field.
func (uu *UserUpdate) SetCreatedTime(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedTime(t)
	return uu
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedTime(*t)
	}
	return uu
}

// SetDeletedTime sets the "deleted_time" field.
func (uu *UserUpdate) SetDeletedTime(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedTime(t)
	return uu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedTime(*t)
	}
	return uu
}

// SetModifiedTime sets the "modified_time" field.
func (uu *UserUpdate) SetModifiedTime(t time.Time) *UserUpdate {
	uu.mutation.SetModifiedTime(t)
	return uu
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableModifiedTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetModifiedTime(*t)
	}
	return uu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uu *UserUpdate) AddRoleIDs(ids ...uint8) *UserUpdate {
	uu.mutation.AddRoleIDs(ids...)
	return uu
}

// AddRoles adds the "roles" edges to the Role entity.
func (uu *UserUpdate) AddRoles(r ...*Role) *UserUpdate {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRoleIDs(ids...)
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (uu *UserUpdate) AddStudentIDs(ids ...uint32) *UserUpdate {
	uu.mutation.AddStudentIDs(ids...)
	return uu
}

// AddStudents adds the "students" edges to the Student entity.
func (uu *UserUpdate) AddStudents(s ...*Student) *UserUpdate {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddStudentIDs(ids...)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (uu *UserUpdate) AddFileIDs(ids ...uint32) *UserUpdate {
	uu.mutation.AddFileIDs(ids...)
	return uu
}

// AddFiles adds the "files" edges to the File entity.
func (uu *UserUpdate) AddFiles(f ...*File) *UserUpdate {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddFileIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uu *UserUpdate) ClearRoles() *UserUpdate {
	uu.mutation.ClearRoles()
	return uu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uu *UserUpdate) RemoveRoleIDs(ids ...uint8) *UserUpdate {
	uu.mutation.RemoveRoleIDs(ids...)
	return uu
}

// RemoveRoles removes "roles" edges to Role entities.
func (uu *UserUpdate) RemoveRoles(r ...*Role) *UserUpdate {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRoleIDs(ids...)
}

// ClearStudents clears all "students" edges to the Student entity.
func (uu *UserUpdate) ClearStudents() *UserUpdate {
	uu.mutation.ClearStudents()
	return uu
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (uu *UserUpdate) RemoveStudentIDs(ids ...uint32) *UserUpdate {
	uu.mutation.RemoveStudentIDs(ids...)
	return uu
}

// RemoveStudents removes "students" edges to Student entities.
func (uu *UserUpdate) RemoveStudents(s ...*Student) *UserUpdate {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveStudentIDs(ids...)
}

// ClearFiles clears all "files" edges to the File entity.
func (uu *UserUpdate) ClearFiles() *UserUpdate {
	uu.mutation.ClearFiles()
	return uu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (uu *UserUpdate) RemoveFileIDs(ids ...uint32) *UserUpdate {
	uu.mutation.RemoveFileIDs(ids...)
	return uu
}

// RemoveFiles removes "files" edges to File entities.
func (uu *UserUpdate) RemoveFiles(f ...*File) *UserUpdate {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Account(); ok {
		if err := user.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf(`ent: validator failed for field "User.account": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Account(); ok {
		_spec.SetField(user.FieldAccount, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if uu.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uu.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uu.mutation.Introduction(); ok {
		_spec.SetField(user.FieldIntroduction, field.TypeString, value)
	}
	if uu.mutation.IntroductionCleared() {
		_spec.ClearField(user.FieldIntroduction, field.TypeString)
	}
	if value, ok := uu.mutation.IsDisabled(); ok {
		_spec.SetField(user.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := uu.mutation.CreatedTime(); ok {
		_spec.SetField(user.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedTime(); ok {
		_spec.SetField(user.FieldDeletedTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.ModifiedTime(); ok {
		_spec.SetField(user.FieldModifiedTime, field.TypeTime, value)
	}
	if uu.mutation.RolesCleared() {
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
		createE := &UserRoleCreate{config: uu.config, mutation: newUserRoleMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uu.mutation.RolesCleared() {
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
		createE := &UserRoleCreate{config: uu.config, mutation: newUserRoleMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RolesIDs(); len(nodes) > 0 {
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
		createE := &UserRoleCreate{config: uu.config, mutation: newUserRoleMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StudentsTable,
			Columns: []string{user.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !uu.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StudentsTable,
			Columns: []string{user.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StudentsTable,
			Columns: []string{user.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !uu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAccount sets the "account" field.
func (uuo *UserUpdateOne) SetAccount(s string) *UserUpdateOne {
	uuo.mutation.SetAccount(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// ClearUsername clears the value of the "username" field.
func (uuo *UserUpdateOne) ClearUsername() *UserUpdateOne {
	uuo.mutation.ClearUsername()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of the "phone" field.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetIntroduction sets the "introduction" field.
func (uuo *UserUpdateOne) SetIntroduction(s string) *UserUpdateOne {
	uuo.mutation.SetIntroduction(s)
	return uuo
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIntroduction(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIntroduction(*s)
	}
	return uuo
}

// ClearIntroduction clears the value of the "introduction" field.
func (uuo *UserUpdateOne) ClearIntroduction() *UserUpdateOne {
	uuo.mutation.ClearIntroduction()
	return uuo
}

// SetIsDisabled sets the "is_disabled" field.
func (uuo *UserUpdateOne) SetIsDisabled(b bool) *UserUpdateOne {
	uuo.mutation.SetIsDisabled(b)
	return uuo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsDisabled(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsDisabled(*b)
	}
	return uuo
}

// SetCreatedTime sets the "created_time" field.
func (uuo *UserUpdateOne) SetCreatedTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedTime(t)
	return uuo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedTime(*t)
	}
	return uuo
}

// SetDeletedTime sets the "deleted_time" field.
func (uuo *UserUpdateOne) SetDeletedTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedTime(t)
	return uuo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedTime(*t)
	}
	return uuo
}

// SetModifiedTime sets the "modified_time" field.
func (uuo *UserUpdateOne) SetModifiedTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetModifiedTime(t)
	return uuo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableModifiedTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetModifiedTime(*t)
	}
	return uuo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uuo *UserUpdateOne) AddRoleIDs(ids ...uint8) *UserUpdateOne {
	uuo.mutation.AddRoleIDs(ids...)
	return uuo
}

// AddRoles adds the "roles" edges to the Role entity.
func (uuo *UserUpdateOne) AddRoles(r ...*Role) *UserUpdateOne {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRoleIDs(ids...)
}

// AddStudentIDs adds the "students" edge to the Student entity by IDs.
func (uuo *UserUpdateOne) AddStudentIDs(ids ...uint32) *UserUpdateOne {
	uuo.mutation.AddStudentIDs(ids...)
	return uuo
}

// AddStudents adds the "students" edges to the Student entity.
func (uuo *UserUpdateOne) AddStudents(s ...*Student) *UserUpdateOne {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddStudentIDs(ids...)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (uuo *UserUpdateOne) AddFileIDs(ids ...uint32) *UserUpdateOne {
	uuo.mutation.AddFileIDs(ids...)
	return uuo
}

// AddFiles adds the "files" edges to the File entity.
func (uuo *UserUpdateOne) AddFiles(f ...*File) *UserUpdateOne {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddFileIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uuo *UserUpdateOne) ClearRoles() *UserUpdateOne {
	uuo.mutation.ClearRoles()
	return uuo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uuo *UserUpdateOne) RemoveRoleIDs(ids ...uint8) *UserUpdateOne {
	uuo.mutation.RemoveRoleIDs(ids...)
	return uuo
}

// RemoveRoles removes "roles" edges to Role entities.
func (uuo *UserUpdateOne) RemoveRoles(r ...*Role) *UserUpdateOne {
	ids := make([]uint8, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRoleIDs(ids...)
}

// ClearStudents clears all "students" edges to the Student entity.
func (uuo *UserUpdateOne) ClearStudents() *UserUpdateOne {
	uuo.mutation.ClearStudents()
	return uuo
}

// RemoveStudentIDs removes the "students" edge to Student entities by IDs.
func (uuo *UserUpdateOne) RemoveStudentIDs(ids ...uint32) *UserUpdateOne {
	uuo.mutation.RemoveStudentIDs(ids...)
	return uuo
}

// RemoveStudents removes "students" edges to Student entities.
func (uuo *UserUpdateOne) RemoveStudents(s ...*Student) *UserUpdateOne {
	ids := make([]uint32, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveStudentIDs(ids...)
}

// ClearFiles clears all "files" edges to the File entity.
func (uuo *UserUpdateOne) ClearFiles() *UserUpdateOne {
	uuo.mutation.ClearFiles()
	return uuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (uuo *UserUpdateOne) RemoveFileIDs(ids ...uint32) *UserUpdateOne {
	uuo.mutation.RemoveFileIDs(ids...)
	return uuo
}

// RemoveFiles removes "files" edges to File entities.
func (uuo *UserUpdateOne) RemoveFiles(f ...*File) *UserUpdateOne {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveFileIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Account(); ok {
		if err := user.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf(`ent: validator failed for field "User.account": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Account(); ok {
		_spec.SetField(user.FieldAccount, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if uuo.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uuo.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uuo.mutation.Introduction(); ok {
		_spec.SetField(user.FieldIntroduction, field.TypeString, value)
	}
	if uuo.mutation.IntroductionCleared() {
		_spec.ClearField(user.FieldIntroduction, field.TypeString)
	}
	if value, ok := uuo.mutation.IsDisabled(); ok {
		_spec.SetField(user.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.CreatedTime(); ok {
		_spec.SetField(user.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedTime(); ok {
		_spec.SetField(user.FieldDeletedTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.ModifiedTime(); ok {
		_spec.SetField(user.FieldModifiedTime, field.TypeTime, value)
	}
	if uuo.mutation.RolesCleared() {
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
		createE := &UserRoleCreate{config: uuo.config, mutation: newUserRoleMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uuo.mutation.RolesCleared() {
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
		createE := &UserRoleCreate{config: uuo.config, mutation: newUserRoleMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RolesIDs(); len(nodes) > 0 {
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
		createE := &UserRoleCreate{config: uuo.config, mutation: newUserRoleMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StudentsTable,
			Columns: []string{user.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedStudentsIDs(); len(nodes) > 0 && !uuo.mutation.StudentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StudentsTable,
			Columns: []string{user.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.StudentsTable,
			Columns: []string{user.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !uuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
