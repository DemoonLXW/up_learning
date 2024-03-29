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
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUID sets the "uid" field.
func (pu *ProjectUpdate) SetUID(u uint32) *ProjectUpdate {
	pu.mutation.SetUID(u)
	return pu
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableUID(u *uint32) *ProjectUpdate {
	if u != nil {
		pu.SetUID(*u)
	}
	return pu
}

// ClearUID clears the value of the "uid" field.
func (pu *ProjectUpdate) ClearUID() *ProjectUpdate {
	pu.mutation.ClearUID()
	return pu
}

// SetTitle sets the "title" field.
func (pu *ProjectUpdate) SetTitle(s string) *ProjectUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetGoal sets the "goal" field.
func (pu *ProjectUpdate) SetGoal(s string) *ProjectUpdate {
	pu.mutation.SetGoal(s)
	return pu
}

// SetPrinciple sets the "principle" field.
func (pu *ProjectUpdate) SetPrinciple(s string) *ProjectUpdate {
	pu.mutation.SetPrinciple(s)
	return pu
}

// SetProcessAndMethod sets the "process_and_method" field.
func (pu *ProjectUpdate) SetProcessAndMethod(s string) *ProjectUpdate {
	pu.mutation.SetProcessAndMethod(s)
	return pu
}

// SetStep sets the "step" field.
func (pu *ProjectUpdate) SetStep(s string) *ProjectUpdate {
	pu.mutation.SetStep(s)
	return pu
}

// SetResultAndConclusion sets the "result_and_conclusion" field.
func (pu *ProjectUpdate) SetResultAndConclusion(s string) *ProjectUpdate {
	pu.mutation.SetResultAndConclusion(s)
	return pu
}

// SetRequirement sets the "requirement" field.
func (pu *ProjectUpdate) SetRequirement(s string) *ProjectUpdate {
	pu.mutation.SetRequirement(s)
	return pu
}

// SetReviewStatus sets the "review_status" field.
func (pu *ProjectUpdate) SetReviewStatus(u uint8) *ProjectUpdate {
	pu.mutation.ResetReviewStatus()
	pu.mutation.SetReviewStatus(u)
	return pu
}

// SetNillableReviewStatus sets the "review_status" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableReviewStatus(u *uint8) *ProjectUpdate {
	if u != nil {
		pu.SetReviewStatus(*u)
	}
	return pu
}

// AddReviewStatus adds u to the "review_status" field.
func (pu *ProjectUpdate) AddReviewStatus(u int8) *ProjectUpdate {
	pu.mutation.AddReviewStatus(u)
	return pu
}

// SetIsDisabled sets the "is_disabled" field.
func (pu *ProjectUpdate) SetIsDisabled(b bool) *ProjectUpdate {
	pu.mutation.SetIsDisabled(b)
	return pu
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableIsDisabled(b *bool) *ProjectUpdate {
	if b != nil {
		pu.SetIsDisabled(*b)
	}
	return pu
}

// SetCreatedTime sets the "created_time" field.
func (pu *ProjectUpdate) SetCreatedTime(t time.Time) *ProjectUpdate {
	pu.mutation.SetCreatedTime(t)
	return pu
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCreatedTime(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetCreatedTime(*t)
	}
	return pu
}

// SetDeletedTime sets the "deleted_time" field.
func (pu *ProjectUpdate) SetDeletedTime(t time.Time) *ProjectUpdate {
	pu.mutation.SetDeletedTime(t)
	return pu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDeletedTime(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetDeletedTime(*t)
	}
	return pu
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (pu *ProjectUpdate) ClearDeletedTime() *ProjectUpdate {
	pu.mutation.ClearDeletedTime()
	return pu
}

// SetModifiedTime sets the "modified_time" field.
func (pu *ProjectUpdate) SetModifiedTime(t time.Time) *ProjectUpdate {
	pu.mutation.SetModifiedTime(t)
	return pu
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableModifiedTime(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetModifiedTime(*t)
	}
	return pu
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (pu *ProjectUpdate) ClearModifiedTime() *ProjectUpdate {
	pu.mutation.ClearModifiedTime()
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *ProjectUpdate) SetUserID(id uint32) *ProjectUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableUserID(id *uint32) *ProjectUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *ProjectUpdate) SetUser(u *User) *ProjectUpdate {
	return pu.SetUserID(u.ID)
}

// AddAttachmentIDs adds the "attachments" edge to the File entity by IDs.
func (pu *ProjectUpdate) AddAttachmentIDs(ids ...uint32) *ProjectUpdate {
	pu.mutation.AddAttachmentIDs(ids...)
	return pu
}

// AddAttachments adds the "attachments" edges to the File entity.
func (pu *ProjectUpdate) AddAttachments(f ...*File) *ProjectUpdate {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.AddAttachmentIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *ProjectUpdate) ClearUser() *ProjectUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearAttachments clears all "attachments" edges to the File entity.
func (pu *ProjectUpdate) ClearAttachments() *ProjectUpdate {
	pu.mutation.ClearAttachments()
	return pu
}

// RemoveAttachmentIDs removes the "attachments" edge to File entities by IDs.
func (pu *ProjectUpdate) RemoveAttachmentIDs(ids ...uint32) *ProjectUpdate {
	pu.mutation.RemoveAttachmentIDs(ids...)
	return pu
}

// RemoveAttachments removes "attachments" edges to File entities.
func (pu *ProjectUpdate) RemoveAttachments(f ...*File) *ProjectUpdate {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.RemoveAttachmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ProjectMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUint32))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(project.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Goal(); ok {
		_spec.SetField(project.FieldGoal, field.TypeString, value)
	}
	if value, ok := pu.mutation.Principle(); ok {
		_spec.SetField(project.FieldPrinciple, field.TypeString, value)
	}
	if value, ok := pu.mutation.ProcessAndMethod(); ok {
		_spec.SetField(project.FieldProcessAndMethod, field.TypeString, value)
	}
	if value, ok := pu.mutation.Step(); ok {
		_spec.SetField(project.FieldStep, field.TypeString, value)
	}
	if value, ok := pu.mutation.ResultAndConclusion(); ok {
		_spec.SetField(project.FieldResultAndConclusion, field.TypeString, value)
	}
	if value, ok := pu.mutation.Requirement(); ok {
		_spec.SetField(project.FieldRequirement, field.TypeString, value)
	}
	if value, ok := pu.mutation.ReviewStatus(); ok {
		_spec.SetField(project.FieldReviewStatus, field.TypeUint8, value)
	}
	if value, ok := pu.mutation.AddedReviewStatus(); ok {
		_spec.AddField(project.FieldReviewStatus, field.TypeUint8, value)
	}
	if value, ok := pu.mutation.IsDisabled(); ok {
		_spec.SetField(project.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := pu.mutation.CreatedTime(); ok {
		_spec.SetField(project.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedTime(); ok {
		_spec.SetField(project.FieldDeletedTime, field.TypeTime, value)
	}
	if pu.mutation.DeletedTimeCleared() {
		_spec.ClearField(project.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := pu.mutation.ModifiedTime(); ok {
		_spec.SetField(project.FieldModifiedTime, field.TypeTime, value)
	}
	if pu.mutation.ModifiedTimeCleared() {
		_spec.ClearField(project.FieldModifiedTime, field.TypeTime)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.UserTable,
			Columns: []string{project.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.UserTable,
			Columns: []string{project.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: project.AttachmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		createE := &ProjectFileCreate{config: pu.config, mutation: newProjectFileMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAttachmentsIDs(); len(nodes) > 0 && !pu.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: project.AttachmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProjectFileCreate{config: pu.config, mutation: newProjectFileMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: project.AttachmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProjectFileCreate{config: pu.config, mutation: newProjectFileMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetUID sets the "uid" field.
func (puo *ProjectUpdateOne) SetUID(u uint32) *ProjectUpdateOne {
	puo.mutation.SetUID(u)
	return puo
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableUID(u *uint32) *ProjectUpdateOne {
	if u != nil {
		puo.SetUID(*u)
	}
	return puo
}

// ClearUID clears the value of the "uid" field.
func (puo *ProjectUpdateOne) ClearUID() *ProjectUpdateOne {
	puo.mutation.ClearUID()
	return puo
}

// SetTitle sets the "title" field.
func (puo *ProjectUpdateOne) SetTitle(s string) *ProjectUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetGoal sets the "goal" field.
func (puo *ProjectUpdateOne) SetGoal(s string) *ProjectUpdateOne {
	puo.mutation.SetGoal(s)
	return puo
}

// SetPrinciple sets the "principle" field.
func (puo *ProjectUpdateOne) SetPrinciple(s string) *ProjectUpdateOne {
	puo.mutation.SetPrinciple(s)
	return puo
}

// SetProcessAndMethod sets the "process_and_method" field.
func (puo *ProjectUpdateOne) SetProcessAndMethod(s string) *ProjectUpdateOne {
	puo.mutation.SetProcessAndMethod(s)
	return puo
}

// SetStep sets the "step" field.
func (puo *ProjectUpdateOne) SetStep(s string) *ProjectUpdateOne {
	puo.mutation.SetStep(s)
	return puo
}

// SetResultAndConclusion sets the "result_and_conclusion" field.
func (puo *ProjectUpdateOne) SetResultAndConclusion(s string) *ProjectUpdateOne {
	puo.mutation.SetResultAndConclusion(s)
	return puo
}

// SetRequirement sets the "requirement" field.
func (puo *ProjectUpdateOne) SetRequirement(s string) *ProjectUpdateOne {
	puo.mutation.SetRequirement(s)
	return puo
}

// SetReviewStatus sets the "review_status" field.
func (puo *ProjectUpdateOne) SetReviewStatus(u uint8) *ProjectUpdateOne {
	puo.mutation.ResetReviewStatus()
	puo.mutation.SetReviewStatus(u)
	return puo
}

// SetNillableReviewStatus sets the "review_status" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableReviewStatus(u *uint8) *ProjectUpdateOne {
	if u != nil {
		puo.SetReviewStatus(*u)
	}
	return puo
}

// AddReviewStatus adds u to the "review_status" field.
func (puo *ProjectUpdateOne) AddReviewStatus(u int8) *ProjectUpdateOne {
	puo.mutation.AddReviewStatus(u)
	return puo
}

// SetIsDisabled sets the "is_disabled" field.
func (puo *ProjectUpdateOne) SetIsDisabled(b bool) *ProjectUpdateOne {
	puo.mutation.SetIsDisabled(b)
	return puo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableIsDisabled(b *bool) *ProjectUpdateOne {
	if b != nil {
		puo.SetIsDisabled(*b)
	}
	return puo
}

// SetCreatedTime sets the "created_time" field.
func (puo *ProjectUpdateOne) SetCreatedTime(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetCreatedTime(t)
	return puo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCreatedTime(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetCreatedTime(*t)
	}
	return puo
}

// SetDeletedTime sets the "deleted_time" field.
func (puo *ProjectUpdateOne) SetDeletedTime(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetDeletedTime(t)
	return puo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDeletedTime(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetDeletedTime(*t)
	}
	return puo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (puo *ProjectUpdateOne) ClearDeletedTime() *ProjectUpdateOne {
	puo.mutation.ClearDeletedTime()
	return puo
}

// SetModifiedTime sets the "modified_time" field.
func (puo *ProjectUpdateOne) SetModifiedTime(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetModifiedTime(t)
	return puo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableModifiedTime(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetModifiedTime(*t)
	}
	return puo
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (puo *ProjectUpdateOne) ClearModifiedTime() *ProjectUpdateOne {
	puo.mutation.ClearModifiedTime()
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *ProjectUpdateOne) SetUserID(id uint32) *ProjectUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableUserID(id *uint32) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *ProjectUpdateOne) SetUser(u *User) *ProjectUpdateOne {
	return puo.SetUserID(u.ID)
}

// AddAttachmentIDs adds the "attachments" edge to the File entity by IDs.
func (puo *ProjectUpdateOne) AddAttachmentIDs(ids ...uint32) *ProjectUpdateOne {
	puo.mutation.AddAttachmentIDs(ids...)
	return puo
}

// AddAttachments adds the "attachments" edges to the File entity.
func (puo *ProjectUpdateOne) AddAttachments(f ...*File) *ProjectUpdateOne {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.AddAttachmentIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *ProjectUpdateOne) ClearUser() *ProjectUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearAttachments clears all "attachments" edges to the File entity.
func (puo *ProjectUpdateOne) ClearAttachments() *ProjectUpdateOne {
	puo.mutation.ClearAttachments()
	return puo
}

// RemoveAttachmentIDs removes the "attachments" edge to File entities by IDs.
func (puo *ProjectUpdateOne) RemoveAttachmentIDs(ids ...uint32) *ProjectUpdateOne {
	puo.mutation.RemoveAttachmentIDs(ids...)
	return puo
}

// RemoveAttachments removes "attachments" edges to File entities.
func (puo *ProjectUpdateOne) RemoveAttachments(f ...*File) *ProjectUpdateOne {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.RemoveAttachmentIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks[*Project, ProjectMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUint32))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
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
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(project.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Goal(); ok {
		_spec.SetField(project.FieldGoal, field.TypeString, value)
	}
	if value, ok := puo.mutation.Principle(); ok {
		_spec.SetField(project.FieldPrinciple, field.TypeString, value)
	}
	if value, ok := puo.mutation.ProcessAndMethod(); ok {
		_spec.SetField(project.FieldProcessAndMethod, field.TypeString, value)
	}
	if value, ok := puo.mutation.Step(); ok {
		_spec.SetField(project.FieldStep, field.TypeString, value)
	}
	if value, ok := puo.mutation.ResultAndConclusion(); ok {
		_spec.SetField(project.FieldResultAndConclusion, field.TypeString, value)
	}
	if value, ok := puo.mutation.Requirement(); ok {
		_spec.SetField(project.FieldRequirement, field.TypeString, value)
	}
	if value, ok := puo.mutation.ReviewStatus(); ok {
		_spec.SetField(project.FieldReviewStatus, field.TypeUint8, value)
	}
	if value, ok := puo.mutation.AddedReviewStatus(); ok {
		_spec.AddField(project.FieldReviewStatus, field.TypeUint8, value)
	}
	if value, ok := puo.mutation.IsDisabled(); ok {
		_spec.SetField(project.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := puo.mutation.CreatedTime(); ok {
		_spec.SetField(project.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedTime(); ok {
		_spec.SetField(project.FieldDeletedTime, field.TypeTime, value)
	}
	if puo.mutation.DeletedTimeCleared() {
		_spec.ClearField(project.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := puo.mutation.ModifiedTime(); ok {
		_spec.SetField(project.FieldModifiedTime, field.TypeTime, value)
	}
	if puo.mutation.ModifiedTimeCleared() {
		_spec.ClearField(project.FieldModifiedTime, field.TypeTime)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.UserTable,
			Columns: []string{project.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.UserTable,
			Columns: []string{project.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: project.AttachmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		createE := &ProjectFileCreate{config: puo.config, mutation: newProjectFileMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAttachmentsIDs(); len(nodes) > 0 && !puo.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: project.AttachmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProjectFileCreate{config: puo.config, mutation: newProjectFileMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: project.AttachmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ProjectFileCreate{config: puo.config, mutation: newProjectFileMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
