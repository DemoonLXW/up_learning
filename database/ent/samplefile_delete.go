// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/samplefile"
)

// SampleFileDelete is the builder for deleting a SampleFile entity.
type SampleFileDelete struct {
	config
	hooks    []Hook
	mutation *SampleFileMutation
}

// Where appends a list predicates to the SampleFileDelete builder.
func (sfd *SampleFileDelete) Where(ps ...predicate.SampleFile) *SampleFileDelete {
	sfd.mutation.Where(ps...)
	return sfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sfd *SampleFileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SampleFileMutation](ctx, sfd.sqlExec, sfd.mutation, sfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sfd *SampleFileDelete) ExecX(ctx context.Context) int {
	n, err := sfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sfd *SampleFileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(samplefile.Table, sqlgraph.NewFieldSpec(samplefile.FieldID, field.TypeUint8))
	if ps := sfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sfd.mutation.done = true
	return affected, err
}

// SampleFileDeleteOne is the builder for deleting a single SampleFile entity.
type SampleFileDeleteOne struct {
	sfd *SampleFileDelete
}

// Where appends a list predicates to the SampleFileDelete builder.
func (sfdo *SampleFileDeleteOne) Where(ps ...predicate.SampleFile) *SampleFileDeleteOne {
	sfdo.sfd.mutation.Where(ps...)
	return sfdo
}

// Exec executes the deletion query.
func (sfdo *SampleFileDeleteOne) Exec(ctx context.Context) error {
	n, err := sfdo.sfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{samplefile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sfdo *SampleFileDeleteOne) ExecX(ctx context.Context) {
	if err := sfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
