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
	"github.com/DemoonLXW/up_learning/database/ent/projectfile"
)

// ProjectFileCreate is the builder for creating a ProjectFile entity.
type ProjectFileCreate struct {
	config
	mutation *ProjectFileMutation
	hooks    []Hook
}

// SetPid sets the "pid" field.
func (pfc *ProjectFileCreate) SetPid(u uint32) *ProjectFileCreate {
	pfc.mutation.SetPid(u)
	return pfc
}

// SetFid sets the "fid" field.
func (pfc *ProjectFileCreate) SetFid(u uint32) *ProjectFileCreate {
	pfc.mutation.SetFid(u)
	return pfc
}

// SetCreatedTime sets the "created_time" field.
func (pfc *ProjectFileCreate) SetCreatedTime(t time.Time) *ProjectFileCreate {
	pfc.mutation.SetCreatedTime(t)
	return pfc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (pfc *ProjectFileCreate) SetNillableCreatedTime(t *time.Time) *ProjectFileCreate {
	if t != nil {
		pfc.SetCreatedTime(*t)
	}
	return pfc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (pfc *ProjectFileCreate) SetProjectID(id uint32) *ProjectFileCreate {
	pfc.mutation.SetProjectID(id)
	return pfc
}

// SetProject sets the "project" edge to the Project entity.
func (pfc *ProjectFileCreate) SetProject(p *Project) *ProjectFileCreate {
	return pfc.SetProjectID(p.ID)
}

// SetFilesID sets the "files" edge to the File entity by ID.
func (pfc *ProjectFileCreate) SetFilesID(id uint32) *ProjectFileCreate {
	pfc.mutation.SetFilesID(id)
	return pfc
}

// SetFiles sets the "files" edge to the File entity.
func (pfc *ProjectFileCreate) SetFiles(f *File) *ProjectFileCreate {
	return pfc.SetFilesID(f.ID)
}

// Mutation returns the ProjectFileMutation object of the builder.
func (pfc *ProjectFileCreate) Mutation() *ProjectFileMutation {
	return pfc.mutation
}

// Save creates the ProjectFile in the database.
func (pfc *ProjectFileCreate) Save(ctx context.Context) (*ProjectFile, error) {
	pfc.defaults()
	return withHooks(ctx, pfc.sqlSave, pfc.mutation, pfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pfc *ProjectFileCreate) SaveX(ctx context.Context) *ProjectFile {
	v, err := pfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfc *ProjectFileCreate) Exec(ctx context.Context) error {
	_, err := pfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfc *ProjectFileCreate) ExecX(ctx context.Context) {
	if err := pfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pfc *ProjectFileCreate) defaults() {
	if _, ok := pfc.mutation.CreatedTime(); !ok {
		v := projectfile.DefaultCreatedTime()
		pfc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfc *ProjectFileCreate) check() error {
	if _, ok := pfc.mutation.Pid(); !ok {
		return &ValidationError{Name: "pid", err: errors.New(`ent: missing required field "ProjectFile.pid"`)}
	}
	if _, ok := pfc.mutation.Fid(); !ok {
		return &ValidationError{Name: "fid", err: errors.New(`ent: missing required field "ProjectFile.fid"`)}
	}
	if _, ok := pfc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "ProjectFile.created_time"`)}
	}
	if _, ok := pfc.mutation.ProjectID(); !ok {
		return &ValidationError{Name: "project", err: errors.New(`ent: missing required edge "ProjectFile.project"`)}
	}
	if _, ok := pfc.mutation.FilesID(); !ok {
		return &ValidationError{Name: "files", err: errors.New(`ent: missing required edge "ProjectFile.files"`)}
	}
	return nil
}

func (pfc *ProjectFileCreate) sqlSave(ctx context.Context) (*ProjectFile, error) {
	if err := pfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pfc.mutation.id = &_node.ID
	pfc.mutation.done = true
	return _node, nil
}

func (pfc *ProjectFileCreate) createSpec() (*ProjectFile, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectFile{config: pfc.config}
		_spec = sqlgraph.NewCreateSpec(projectfile.Table, sqlgraph.NewFieldSpec(projectfile.FieldID, field.TypeInt))
	)
	if value, ok := pfc.mutation.CreatedTime(); ok {
		_spec.SetField(projectfile.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if nodes := pfc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectfile.ProjectTable,
			Columns: []string{projectfile.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Pid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectfile.FilesTable,
			Columns: []string{projectfile.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Fid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectFileCreateBulk is the builder for creating many ProjectFile entities in bulk.
type ProjectFileCreateBulk struct {
	config
	builders []*ProjectFileCreate
}

// Save creates the ProjectFile entities in the database.
func (pfcb *ProjectFileCreateBulk) Save(ctx context.Context) ([]*ProjectFile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pfcb.builders))
	nodes := make([]*ProjectFile, len(pfcb.builders))
	mutators := make([]Mutator, len(pfcb.builders))
	for i := range pfcb.builders {
		func(i int, root context.Context) {
			builder := pfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectFileMutation)
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
					_, err = mutators[i+1].Mutate(root, pfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, pfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pfcb *ProjectFileCreateBulk) SaveX(ctx context.Context) []*ProjectFile {
	v, err := pfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pfcb *ProjectFileCreateBulk) Exec(ctx context.Context) error {
	_, err := pfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfcb *ProjectFileCreateBulk) ExecX(ctx context.Context) {
	if err := pfcb.Exec(ctx); err != nil {
		panic(err)
	}
}