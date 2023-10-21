// Code generated by ent, DO NOT EDIT.

package projectfile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldLTE(FieldID, id))
}

// Pid applies equality check predicate on the "pid" field. It's identical to PidEQ.
func Pid(v uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldPid, v))
}

// Fid applies equality check predicate on the "fid" field. It's identical to FidEQ.
func Fid(v uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldFid, v))
}

// CreatedTime applies equality check predicate on the "created_time" field. It's identical to CreatedTimeEQ.
func CreatedTime(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldCreatedTime, v))
}

// PidEQ applies the EQ predicate on the "pid" field.
func PidEQ(v uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldPid, v))
}

// PidNEQ applies the NEQ predicate on the "pid" field.
func PidNEQ(v uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNEQ(FieldPid, v))
}

// PidIn applies the In predicate on the "pid" field.
func PidIn(vs ...uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldIn(FieldPid, vs...))
}

// PidNotIn applies the NotIn predicate on the "pid" field.
func PidNotIn(vs ...uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNotIn(FieldPid, vs...))
}

// FidEQ applies the EQ predicate on the "fid" field.
func FidEQ(v uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldFid, v))
}

// FidNEQ applies the NEQ predicate on the "fid" field.
func FidNEQ(v uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNEQ(FieldFid, v))
}

// FidIn applies the In predicate on the "fid" field.
func FidIn(vs ...uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldIn(FieldFid, vs...))
}

// FidNotIn applies the NotIn predicate on the "fid" field.
func FidNotIn(vs ...uint32) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNotIn(FieldFid, vs...))
}

// CreatedTimeEQ applies the EQ predicate on the "created_time" field.
func CreatedTimeEQ(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldEQ(FieldCreatedTime, v))
}

// CreatedTimeNEQ applies the NEQ predicate on the "created_time" field.
func CreatedTimeNEQ(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNEQ(FieldCreatedTime, v))
}

// CreatedTimeIn applies the In predicate on the "created_time" field.
func CreatedTimeIn(vs ...time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldIn(FieldCreatedTime, vs...))
}

// CreatedTimeNotIn applies the NotIn predicate on the "created_time" field.
func CreatedTimeNotIn(vs ...time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldNotIn(FieldCreatedTime, vs...))
}

// CreatedTimeGT applies the GT predicate on the "created_time" field.
func CreatedTimeGT(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldGT(FieldCreatedTime, v))
}

// CreatedTimeGTE applies the GTE predicate on the "created_time" field.
func CreatedTimeGTE(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldGTE(FieldCreatedTime, v))
}

// CreatedTimeLT applies the LT predicate on the "created_time" field.
func CreatedTimeLT(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldLT(FieldCreatedTime, v))
}

// CreatedTimeLTE applies the LTE predicate on the "created_time" field.
func CreatedTimeLTE(v time.Time) predicate.ProjectFile {
	return predicate.ProjectFile(sql.FieldLTE(FieldCreatedTime, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProjectFile) predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProjectFile) predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProjectFile) predicate.ProjectFile {
	return predicate.ProjectFile(func(s *sql.Selector) {
		p(s.Not())
	})
}
