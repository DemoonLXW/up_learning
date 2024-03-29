// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/projectfile"
)

// ProjectFileDelete is the builder for deleting a ProjectFile entity.
type ProjectFileDelete struct {
	config
	hooks    []Hook
	mutation *ProjectFileMutation
}

// Where appends a list predicates to the ProjectFileDelete builder.
func (pfd *ProjectFileDelete) Where(ps ...predicate.ProjectFile) *ProjectFileDelete {
	pfd.mutation.Where(ps...)
	return pfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pfd *ProjectFileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ProjectFileMutation](ctx, pfd.sqlExec, pfd.mutation, pfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pfd *ProjectFileDelete) ExecX(ctx context.Context) int {
	n, err := pfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pfd *ProjectFileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(projectfile.Table, nil)
	if ps := pfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pfd.mutation.done = true
	return affected, err
}

// ProjectFileDeleteOne is the builder for deleting a single ProjectFile entity.
type ProjectFileDeleteOne struct {
	pfd *ProjectFileDelete
}

// Where appends a list predicates to the ProjectFileDelete builder.
func (pfdo *ProjectFileDeleteOne) Where(ps ...predicate.ProjectFile) *ProjectFileDeleteOne {
	pfdo.pfd.mutation.Where(ps...)
	return pfdo
}

// Exec executes the deletion query.
func (pfdo *ProjectFileDeleteOne) Exec(ctx context.Context) error {
	n, err := pfdo.pfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projectfile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pfdo *ProjectFileDeleteOne) ExecX(ctx context.Context) {
	if err := pfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
