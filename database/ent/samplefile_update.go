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
	"github.com/DemoonLXW/up_learning/database/ent/samplefile"
)

// SampleFileUpdate is the builder for updating SampleFile entities.
type SampleFileUpdate struct {
	config
	hooks    []Hook
	mutation *SampleFileMutation
}

// Where appends a list predicates to the SampleFileUpdate builder.
func (sfu *SampleFileUpdate) Where(ps ...predicate.SampleFile) *SampleFileUpdate {
	sfu.mutation.Where(ps...)
	return sfu
}

// SetFid sets the "fid" field.
func (sfu *SampleFileUpdate) SetFid(u uint32) *SampleFileUpdate {
	sfu.mutation.SetFid(u)
	return sfu
}

// SetType sets the "type" field.
func (sfu *SampleFileUpdate) SetType(s string) *SampleFileUpdate {
	sfu.mutation.SetType(s)
	return sfu
}

// SetIsDisabled sets the "is_disabled" field.
func (sfu *SampleFileUpdate) SetIsDisabled(b bool) *SampleFileUpdate {
	sfu.mutation.SetIsDisabled(b)
	return sfu
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (sfu *SampleFileUpdate) SetNillableIsDisabled(b *bool) *SampleFileUpdate {
	if b != nil {
		sfu.SetIsDisabled(*b)
	}
	return sfu
}

// SetCreatedTime sets the "created_time" field.
func (sfu *SampleFileUpdate) SetCreatedTime(t time.Time) *SampleFileUpdate {
	sfu.mutation.SetCreatedTime(t)
	return sfu
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (sfu *SampleFileUpdate) SetNillableCreatedTime(t *time.Time) *SampleFileUpdate {
	if t != nil {
		sfu.SetCreatedTime(*t)
	}
	return sfu
}

// SetDeletedTime sets the "deleted_time" field.
func (sfu *SampleFileUpdate) SetDeletedTime(t time.Time) *SampleFileUpdate {
	sfu.mutation.SetDeletedTime(t)
	return sfu
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (sfu *SampleFileUpdate) SetNillableDeletedTime(t *time.Time) *SampleFileUpdate {
	if t != nil {
		sfu.SetDeletedTime(*t)
	}
	return sfu
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (sfu *SampleFileUpdate) ClearDeletedTime() *SampleFileUpdate {
	sfu.mutation.ClearDeletedTime()
	return sfu
}

// SetModifiedTime sets the "modified_time" field.
func (sfu *SampleFileUpdate) SetModifiedTime(t time.Time) *SampleFileUpdate {
	sfu.mutation.SetModifiedTime(t)
	return sfu
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (sfu *SampleFileUpdate) SetNillableModifiedTime(t *time.Time) *SampleFileUpdate {
	if t != nil {
		sfu.SetModifiedTime(*t)
	}
	return sfu
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (sfu *SampleFileUpdate) ClearModifiedTime() *SampleFileUpdate {
	sfu.mutation.ClearModifiedTime()
	return sfu
}

// SetFileID sets the "file" edge to the File entity by ID.
func (sfu *SampleFileUpdate) SetFileID(id uint32) *SampleFileUpdate {
	sfu.mutation.SetFileID(id)
	return sfu
}

// SetFile sets the "file" edge to the File entity.
func (sfu *SampleFileUpdate) SetFile(f *File) *SampleFileUpdate {
	return sfu.SetFileID(f.ID)
}

// Mutation returns the SampleFileMutation object of the builder.
func (sfu *SampleFileUpdate) Mutation() *SampleFileMutation {
	return sfu.mutation
}

// ClearFile clears the "file" edge to the File entity.
func (sfu *SampleFileUpdate) ClearFile() *SampleFileUpdate {
	sfu.mutation.ClearFile()
	return sfu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sfu *SampleFileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SampleFileMutation](ctx, sfu.sqlSave, sfu.mutation, sfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sfu *SampleFileUpdate) SaveX(ctx context.Context) int {
	affected, err := sfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sfu *SampleFileUpdate) Exec(ctx context.Context) error {
	_, err := sfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sfu *SampleFileUpdate) ExecX(ctx context.Context) {
	if err := sfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sfu *SampleFileUpdate) check() error {
	if _, ok := sfu.mutation.FileID(); sfu.mutation.FileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SampleFile.file"`)
	}
	return nil
}

func (sfu *SampleFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sfu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(samplefile.Table, samplefile.Columns, sqlgraph.NewFieldSpec(samplefile.FieldID, field.TypeUint8))
	if ps := sfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sfu.mutation.GetType(); ok {
		_spec.SetField(samplefile.FieldType, field.TypeString, value)
	}
	if value, ok := sfu.mutation.IsDisabled(); ok {
		_spec.SetField(samplefile.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := sfu.mutation.CreatedTime(); ok {
		_spec.SetField(samplefile.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := sfu.mutation.DeletedTime(); ok {
		_spec.SetField(samplefile.FieldDeletedTime, field.TypeTime, value)
	}
	if sfu.mutation.DeletedTimeCleared() {
		_spec.ClearField(samplefile.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := sfu.mutation.ModifiedTime(); ok {
		_spec.SetField(samplefile.FieldModifiedTime, field.TypeTime, value)
	}
	if sfu.mutation.ModifiedTimeCleared() {
		_spec.ClearField(samplefile.FieldModifiedTime, field.TypeTime)
	}
	if sfu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   samplefile.FileTable,
			Columns: []string{samplefile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sfu.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   samplefile.FileTable,
			Columns: []string{samplefile.FileColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, sfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{samplefile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sfu.mutation.done = true
	return n, nil
}

// SampleFileUpdateOne is the builder for updating a single SampleFile entity.
type SampleFileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SampleFileMutation
}

// SetFid sets the "fid" field.
func (sfuo *SampleFileUpdateOne) SetFid(u uint32) *SampleFileUpdateOne {
	sfuo.mutation.SetFid(u)
	return sfuo
}

// SetType sets the "type" field.
func (sfuo *SampleFileUpdateOne) SetType(s string) *SampleFileUpdateOne {
	sfuo.mutation.SetType(s)
	return sfuo
}

// SetIsDisabled sets the "is_disabled" field.
func (sfuo *SampleFileUpdateOne) SetIsDisabled(b bool) *SampleFileUpdateOne {
	sfuo.mutation.SetIsDisabled(b)
	return sfuo
}

// SetNillableIsDisabled sets the "is_disabled" field if the given value is not nil.
func (sfuo *SampleFileUpdateOne) SetNillableIsDisabled(b *bool) *SampleFileUpdateOne {
	if b != nil {
		sfuo.SetIsDisabled(*b)
	}
	return sfuo
}

// SetCreatedTime sets the "created_time" field.
func (sfuo *SampleFileUpdateOne) SetCreatedTime(t time.Time) *SampleFileUpdateOne {
	sfuo.mutation.SetCreatedTime(t)
	return sfuo
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (sfuo *SampleFileUpdateOne) SetNillableCreatedTime(t *time.Time) *SampleFileUpdateOne {
	if t != nil {
		sfuo.SetCreatedTime(*t)
	}
	return sfuo
}

// SetDeletedTime sets the "deleted_time" field.
func (sfuo *SampleFileUpdateOne) SetDeletedTime(t time.Time) *SampleFileUpdateOne {
	sfuo.mutation.SetDeletedTime(t)
	return sfuo
}

// SetNillableDeletedTime sets the "deleted_time" field if the given value is not nil.
func (sfuo *SampleFileUpdateOne) SetNillableDeletedTime(t *time.Time) *SampleFileUpdateOne {
	if t != nil {
		sfuo.SetDeletedTime(*t)
	}
	return sfuo
}

// ClearDeletedTime clears the value of the "deleted_time" field.
func (sfuo *SampleFileUpdateOne) ClearDeletedTime() *SampleFileUpdateOne {
	sfuo.mutation.ClearDeletedTime()
	return sfuo
}

// SetModifiedTime sets the "modified_time" field.
func (sfuo *SampleFileUpdateOne) SetModifiedTime(t time.Time) *SampleFileUpdateOne {
	sfuo.mutation.SetModifiedTime(t)
	return sfuo
}

// SetNillableModifiedTime sets the "modified_time" field if the given value is not nil.
func (sfuo *SampleFileUpdateOne) SetNillableModifiedTime(t *time.Time) *SampleFileUpdateOne {
	if t != nil {
		sfuo.SetModifiedTime(*t)
	}
	return sfuo
}

// ClearModifiedTime clears the value of the "modified_time" field.
func (sfuo *SampleFileUpdateOne) ClearModifiedTime() *SampleFileUpdateOne {
	sfuo.mutation.ClearModifiedTime()
	return sfuo
}

// SetFileID sets the "file" edge to the File entity by ID.
func (sfuo *SampleFileUpdateOne) SetFileID(id uint32) *SampleFileUpdateOne {
	sfuo.mutation.SetFileID(id)
	return sfuo
}

// SetFile sets the "file" edge to the File entity.
func (sfuo *SampleFileUpdateOne) SetFile(f *File) *SampleFileUpdateOne {
	return sfuo.SetFileID(f.ID)
}

// Mutation returns the SampleFileMutation object of the builder.
func (sfuo *SampleFileUpdateOne) Mutation() *SampleFileMutation {
	return sfuo.mutation
}

// ClearFile clears the "file" edge to the File entity.
func (sfuo *SampleFileUpdateOne) ClearFile() *SampleFileUpdateOne {
	sfuo.mutation.ClearFile()
	return sfuo
}

// Where appends a list predicates to the SampleFileUpdate builder.
func (sfuo *SampleFileUpdateOne) Where(ps ...predicate.SampleFile) *SampleFileUpdateOne {
	sfuo.mutation.Where(ps...)
	return sfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sfuo *SampleFileUpdateOne) Select(field string, fields ...string) *SampleFileUpdateOne {
	sfuo.fields = append([]string{field}, fields...)
	return sfuo
}

// Save executes the query and returns the updated SampleFile entity.
func (sfuo *SampleFileUpdateOne) Save(ctx context.Context) (*SampleFile, error) {
	return withHooks[*SampleFile, SampleFileMutation](ctx, sfuo.sqlSave, sfuo.mutation, sfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sfuo *SampleFileUpdateOne) SaveX(ctx context.Context) *SampleFile {
	node, err := sfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sfuo *SampleFileUpdateOne) Exec(ctx context.Context) error {
	_, err := sfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sfuo *SampleFileUpdateOne) ExecX(ctx context.Context) {
	if err := sfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sfuo *SampleFileUpdateOne) check() error {
	if _, ok := sfuo.mutation.FileID(); sfuo.mutation.FileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SampleFile.file"`)
	}
	return nil
}

func (sfuo *SampleFileUpdateOne) sqlSave(ctx context.Context) (_node *SampleFile, err error) {
	if err := sfuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(samplefile.Table, samplefile.Columns, sqlgraph.NewFieldSpec(samplefile.FieldID, field.TypeUint8))
	id, ok := sfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SampleFile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, samplefile.FieldID)
		for _, f := range fields {
			if !samplefile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != samplefile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sfuo.mutation.GetType(); ok {
		_spec.SetField(samplefile.FieldType, field.TypeString, value)
	}
	if value, ok := sfuo.mutation.IsDisabled(); ok {
		_spec.SetField(samplefile.FieldIsDisabled, field.TypeBool, value)
	}
	if value, ok := sfuo.mutation.CreatedTime(); ok {
		_spec.SetField(samplefile.FieldCreatedTime, field.TypeTime, value)
	}
	if value, ok := sfuo.mutation.DeletedTime(); ok {
		_spec.SetField(samplefile.FieldDeletedTime, field.TypeTime, value)
	}
	if sfuo.mutation.DeletedTimeCleared() {
		_spec.ClearField(samplefile.FieldDeletedTime, field.TypeTime)
	}
	if value, ok := sfuo.mutation.ModifiedTime(); ok {
		_spec.SetField(samplefile.FieldModifiedTime, field.TypeTime, value)
	}
	if sfuo.mutation.ModifiedTimeCleared() {
		_spec.ClearField(samplefile.FieldModifiedTime, field.TypeTime)
	}
	if sfuo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   samplefile.FileTable,
			Columns: []string{samplefile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sfuo.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   samplefile.FileTable,
			Columns: []string{samplefile.FileColumn},
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
	_node = &SampleFile{config: sfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{samplefile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sfuo.mutation.done = true
	return _node, nil
}
