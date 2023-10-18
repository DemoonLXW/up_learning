// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (pc *ProjectCreate) SetUID(u uint32) *ProjectCreate {
	pc.mutation.SetUID(u)
	return pc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUID(u *uint32) *ProjectCreate {
	if u != nil {
		pc.SetUID(*u)
	}
	return pc
}

// SetGoal sets the "goal" field.
func (pc *ProjectCreate) SetGoal(s string) *ProjectCreate {
	pc.mutation.SetGoal(s)
	return pc
}

// SetPrinciple sets the "principle" field.
func (pc *ProjectCreate) SetPrinciple(s string) *ProjectCreate {
	pc.mutation.SetPrinciple(s)
	return pc
}

// SetProcessAndMethod sets the "process_and_method" field.
func (pc *ProjectCreate) SetProcessAndMethod(s string) *ProjectCreate {
	pc.mutation.SetProcessAndMethod(s)
	return pc
}

// SetStep sets the "step" field.
func (pc *ProjectCreate) SetStep(s string) *ProjectCreate {
	pc.mutation.SetStep(s)
	return pc
}

// SetResultAndConclusion sets the "result_and_conclusion" field.
func (pc *ProjectCreate) SetResultAndConclusion(s string) *ProjectCreate {
	pc.mutation.SetResultAndConclusion(s)
	return pc
}

// SetRequirement sets the "requirement" field.
func (pc *ProjectCreate) SetRequirement(s string) *ProjectCreate {
	pc.mutation.SetRequirement(s)
	return pc
}

// SetIsDisabled sets the "is_disabled" field.
func (pc *ProjectCreate) SetIsDisabled(b bool) *ProjectCreate {
	pc.mutation.SetIsDisabled(b)
	return pc
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableIsDisabled(b *bool) *ProjectCreate {
	if b != nil {
		pc.SetIsDisabled(*b)
	}
	return pc
}

// SetCreatedTime sets the "created_time" field.
func (pc *ProjectCreate) SetCreatedTime(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedTime(t)
	return pc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedTime(*t)
	}
	return pc
}

// SetDeletedTime sets the "deleted_time" field.
func (pc *ProjectCreate) SetDeletedTime(t time.Time) *ProjectCreate {
	pc.mutation.SetDeletedTime(t)
	return pc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDeletedTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetDeletedTime(*t)
	}
	return pc
}

// SetModifiedTime sets the "modified_time" field.
func (pc *ProjectCreate) SetModifiedTime(t time.Time) *ProjectCreate {
	pc.mutation.SetModifiedTime(t)
	return pc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableModifiedTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetModifiedTime(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProjectCreate) SetID(u uint32) *ProjectCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *ProjectCreate) SetUserID(id uint32) *ProjectCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *ProjectCreate) SetNillableUserID(id *uint32) *ProjectCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *ProjectCreate) SetUser(u *User) *ProjectCreate {
	return pc.SetUserID(u.ID)
}

// AddAttachmentIDs adds the "attachments" edge to the File entity by IDs.
func (pc *ProjectCreate) AddAttachmentIDs(ids ...uint32) *ProjectCreate {
	pc.mutation.AddAttachmentIDs(ids...)
	return pc
}

// AddAttachments adds the "attachments" edges to the File entity.
func (pc *ProjectCreate) AddAttachments(f ...*File) *ProjectCreate {
	ids := make([]uint32, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pc.AddAttachmentIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.IsDisabled(); !ok {
		v := project.DefaultIsDisabled
		pc.mutation.SetIsDisabled(v)
	}
	if _, ok := pc.mutation.CreatedTime(); !ok {
		v := project.DefaultCreatedTime()
		pc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.Goal(); !ok {
		return &ValidationError{Name: "goal", err: errors.New(`ent: missing required field "Project.goal"`)}
	}
	if _, ok := pc.mutation.Principle(); !ok {
		return &ValidationError{Name: "principle", err: errors.New(`ent: missing required field "Project.principle"`)}
	}
	if _, ok := pc.mutation.ProcessAndMethod(); !ok {
		return &ValidationError{Name: "process_and_method", err: errors.New(`ent: missing required field "Project.process_and_method"`)}
	}
	if _, ok := pc.mutation.Step(); !ok {
		return &ValidationError{Name: "step", err: errors.New(`ent: missing required field "Project.step"`)}
	}
	if _, ok := pc.mutation.ResultAndConclusion(); !ok {
		return &ValidationError{Name: "result_and_conclusion", err: errors.New(`ent: missing required field "Project.result_and_conclusion"`)}
	}
	if _, ok := pc.mutation.Requirement(); !ok {
		return &ValidationError{Name: "requirement", err: errors.New(`ent: missing required field "Project.requirement"`)}
	}
	if _, ok := pc.mutation.IsDisabled(); !ok {
		return &ValidationError{Name: "is_disabled", err: errors.New(`ent: missing required field "Project.is_disabled"`)}
	}
	if _, ok := pc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "Project.created_time"`)}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
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
		_node.ID = uint32(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUint32))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Goal(); ok {
		_spec.SetField(project.FieldGoal, field.TypeString, value)
		_node.Goal = value
	}
	if value, ok := pc.mutation.Principle(); ok {
		_spec.SetField(project.FieldPrinciple, field.TypeString, value)
		_node.Principle = value
	}
	if value, ok := pc.mutation.ProcessAndMethod(); ok {
		_spec.SetField(project.FieldProcessAndMethod, field.TypeString, value)
		_node.ProcessAndMethod = value
	}
	if value, ok := pc.mutation.Step(); ok {
		_spec.SetField(project.FieldStep, field.TypeString, value)
		_node.Step = value
	}
	if value, ok := pc.mutation.ResultAndConclusion(); ok {
		_spec.SetField(project.FieldResultAndConclusion, field.TypeString, value)
		_node.ResultAndConclusion = value
	}
	if value, ok := pc.mutation.Requirement(); ok {
		_spec.SetField(project.FieldRequirement, field.TypeString, value)
		_node.Requirement = value
	}
	if value, ok := pc.mutation.IsDisabled(); ok {
		_spec.SetField(project.FieldIsDisabled, field.TypeBool, value)
		_node.IsDisabled = value
	}
	if value, ok := pc.mutation.CreatedTime(); ok {
		_spec.SetField(project.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := pc.mutation.DeletedTime(); ok {
		_spec.SetField(project.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := pc.mutation.ModifiedTime(); ok {
		_spec.SetField(project.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.AttachmentsTable,
			Columns: []string{project.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}