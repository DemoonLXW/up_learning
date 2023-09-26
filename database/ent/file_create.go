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
	"github.com/DemoonLXW/up_learning/database/ent/samplefile"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (fc *FileCreate) SetUID(u uint32) *FileCreate {
	fc.mutation.SetUID(u)
	return fc
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetPath sets the "path" field.
func (fc *FileCreate) SetPath(s string) *FileCreate {
	fc.mutation.SetPath(s)
	return fc
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int64) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int64) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetIsDisabled sets the "is_disabled" field.
func (fc *FileCreate) SetIsDisabled(b bool) *FileCreate {
	fc.mutation.SetIsDisabled(b)
	return fc
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (fc *FileCreate) SetNillableIsDisabled(b *bool) *FileCreate {
	if b != nil {
		fc.SetIsDisabled(*b)
	}
	return fc
}

// SetCreatedTime sets the "created_time" field.
func (fc *FileCreate) SetCreatedTime(t time.Time) *FileCreate {
	fc.mutation.SetCreatedTime(t)
	return fc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedTime(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedTime(*t)
	}
	return fc
}

// SetDeletedTime sets the "deleted_time" field.
func (fc *FileCreate) SetDeletedTime(t time.Time) *FileCreate {
	fc.mutation.SetDeletedTime(t)
	return fc
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (fc *FileCreate) SetNillableDeletedTime(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetDeletedTime(*t)
	}
	return fc
}

// SetModifiedTime sets the "modified_time" field.
func (fc *FileCreate) SetModifiedTime(t time.Time) *FileCreate {
	fc.mutation.SetModifiedTime(t)
	return fc
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (fc *FileCreate) SetNillableModifiedTime(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetModifiedTime(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileCreate) SetID(u uint32) *FileCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (fc *FileCreate) SetCreatorID(id uint32) *FileCreate {
	fc.mutation.SetCreatorID(id)
	return fc
}

// SetCreator sets the "creator" edge to the User entity.
func (fc *FileCreate) SetCreator(u *User) *FileCreate {
	return fc.SetCreatorID(u.ID)
}

// SetSampleID sets the "sample" edge to the SampleFile entity by ID.
func (fc *FileCreate) SetSampleID(id uint8) *FileCreate {
	fc.mutation.SetSampleID(id)
	return fc
}

// SetNillableSampleID sets the "sample" edge to the SampleFile entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableSampleID(id *uint8) *FileCreate {
	if id != nil {
		fc = fc.SetSampleID(*id)
	}
	return fc
}

// SetSample sets the "sample" edge to the SampleFile entity.
func (fc *FileCreate) SetSample(s *SampleFile) *FileCreate {
	return fc.SetSampleID(s.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	fc.defaults()
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.IsDisabled(); !ok {
		v := file.DefaultIsDisabled
		fc.mutation.SetIsDisabled(v)
	}
	if _, ok := fc.mutation.CreatedTime(); !ok {
		v := file.DefaultCreatedTime()
		fc.mutation.SetCreatedTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "File.uid"`)}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := file.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "File.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "File.path"`)}
	}
	if v, ok := fc.mutation.Path(); ok {
		if err := file.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "File.path": %w`, err)}
		}
	}
	if _, ok := fc.mutation.IsDisabled(); !ok {
		return &ValidationError{Name: "is_disabled", err: errors.New(`ent: missing required field "File.is_disabled"`)}
	}
	if _, ok := fc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "File.created_time"`)}
	}
	if _, ok := fc.mutation.CreatorID(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`ent: missing required edge "File.creator"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32))
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt64, value)
		_node.Size = value
	}
	if value, ok := fc.mutation.IsDisabled(); ok {
		_spec.SetField(file.FieldIsDisabled, field.TypeBool, value)
		_node.IsDisabled = value
	}
	if value, ok := fc.mutation.CreatedTime(); ok {
		_spec.SetField(file.FieldCreatedTime, field.TypeTime, value)
		_node.CreatedTime = &value
	}
	if value, ok := fc.mutation.DeletedTime(); ok {
		_spec.SetField(file.FieldDeletedTime, field.TypeTime, value)
		_node.DeletedTime = &value
	}
	if value, ok := fc.mutation.ModifiedTime(); ok {
		_spec.SetField(file.FieldModifiedTime, field.TypeTime, value)
		_node.ModifiedTime = &value
	}
	if nodes := fc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.CreatorTable,
			Columns: []string{file.CreatorColumn},
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
	if nodes := fc.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   file.SampleTable,
			Columns: []string{file.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(samplefile.FieldID, field.TypeUint8),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
