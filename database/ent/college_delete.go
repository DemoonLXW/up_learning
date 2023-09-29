// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/college"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
)

// CollegeDelete is the builder for deleting a College entity.
type CollegeDelete struct {
	config
	hooks    []Hook
	mutation *CollegeMutation
}

// Where appends a list predicates to the CollegeDelete builder.
func (cd *CollegeDelete) Where(ps ...predicate.College) *CollegeDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CollegeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CollegeDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CollegeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(college.Table, sqlgraph.NewFieldSpec(college.FieldID, field.TypeUint8))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CollegeDeleteOne is the builder for deleting a single College entity.
type CollegeDeleteOne struct {
	cd *CollegeDelete
}

// Where appends a list predicates to the CollegeDelete builder.
func (cdo *CollegeDeleteOne) Where(ps ...predicate.College) *CollegeDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CollegeDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{college.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CollegeDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
