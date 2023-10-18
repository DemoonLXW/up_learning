// Code generated by ent, DO NOT EDIT.

package major

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint16) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint16) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint16) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint16) predicate.Major {
	return predicate.Major(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint16) predicate.Major {
	return predicate.Major(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint16) predicate.Major {
	return predicate.Major(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint16) predicate.Major {
	return predicate.Major(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint16) predicate.Major {
	return predicate.Major(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint16) predicate.Major {
	return predicate.Major(sql.FieldLTE(FieldID, id))
}

// Cid applies equality check predicate on the "cid" field. It's identical to CidEQ.
func Cid(v uint8) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldCid, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldName, v))
}

// IsDisabled applies equality check predicate on the "is_disabled" field. It's identical to IsDisabledEQ.
func IsDisabled(v bool) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldIsDisabled, v))
}

// CreatedTime applies equality check predicate on the "created_time" field. It's identical to CreatedTimeEQ.
func CreatedTime(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldCreatedTime, v))
}

// DeletedTime applies equality check predicate on the "deleted_time" field. It's identical to DeletedTimeEQ.
func DeletedTime(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldDeletedTime, v))
}

// ModifiedTime applies equality check predicate on the "modified_time" field. It's identical to ModifiedTimeEQ.
func ModifiedTime(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldModifiedTime, v))
}

// CidEQ applies the EQ predicate on the "cid" field.
func CidEQ(v uint8) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldCid, v))
}

// CidNEQ applies the NEQ predicate on the "cid" field.
func CidNEQ(v uint8) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldCid, v))
}

// CidIn applies the In predicate on the "cid" field.
func CidIn(vs ...uint8) predicate.Major {
	return predicate.Major(sql.FieldIn(FieldCid, vs...))
}

// CidNotIn applies the NotIn predicate on the "cid" field.
func CidNotIn(vs ...uint8) predicate.Major {
	return predicate.Major(sql.FieldNotIn(FieldCid, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Major {
	return predicate.Major(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Major {
	return predicate.Major(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Major {
	return predicate.Major(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Major {
	return predicate.Major(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Major {
	return predicate.Major(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Major {
	return predicate.Major(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Major {
	return predicate.Major(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Major {
	return predicate.Major(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Major {
	return predicate.Major(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Major {
	return predicate.Major(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Major {
	return predicate.Major(sql.FieldContainsFold(FieldName, v))
}

// IsDisabledEQ applies the EQ predicate on the "is_disabled" field.
func IsDisabledEQ(v bool) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldIsDisabled, v))
}

// IsDisabledNEQ applies the NEQ predicate on the "is_disabled" field.
func IsDisabledNEQ(v bool) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldIsDisabled, v))
}

// CreatedTimeEQ applies the EQ predicate on the "created_time" field.
func CreatedTimeEQ(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldCreatedTime, v))
}

// CreatedTimeNEQ applies the NEQ predicate on the "created_time" field.
func CreatedTimeNEQ(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldCreatedTime, v))
}

// CreatedTimeIn applies the In predicate on the "created_time" field.
func CreatedTimeIn(vs ...time.Time) predicate.Major {
	return predicate.Major(sql.FieldIn(FieldCreatedTime, vs...))
}

// CreatedTimeNotIn applies the NotIn predicate on the "created_time" field.
func CreatedTimeNotIn(vs ...time.Time) predicate.Major {
	return predicate.Major(sql.FieldNotIn(FieldCreatedTime, vs...))
}

// CreatedTimeGT applies the GT predicate on the "created_time" field.
func CreatedTimeGT(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldGT(FieldCreatedTime, v))
}

// CreatedTimeGTE applies the GTE predicate on the "created_time" field.
func CreatedTimeGTE(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldGTE(FieldCreatedTime, v))
}

// CreatedTimeLT applies the LT predicate on the "created_time" field.
func CreatedTimeLT(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldLT(FieldCreatedTime, v))
}

// CreatedTimeLTE applies the LTE predicate on the "created_time" field.
func CreatedTimeLTE(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldLTE(FieldCreatedTime, v))
}

// DeletedTimeEQ applies the EQ predicate on the "deleted_time" field.
func DeletedTimeEQ(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldDeletedTime, v))
}

// DeletedTimeNEQ applies the NEQ predicate on the "deleted_time" field.
func DeletedTimeNEQ(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldDeletedTime, v))
}

// DeletedTimeIn applies the In predicate on the "deleted_time" field.
func DeletedTimeIn(vs ...time.Time) predicate.Major {
	return predicate.Major(sql.FieldIn(FieldDeletedTime, vs...))
}

// DeletedTimeNotIn applies the NotIn predicate on the "deleted_time" field.
func DeletedTimeNotIn(vs ...time.Time) predicate.Major {
	return predicate.Major(sql.FieldNotIn(FieldDeletedTime, vs...))
}

// DeletedTimeGT applies the GT predicate on the "deleted_time" field.
func DeletedTimeGT(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldGT(FieldDeletedTime, v))
}

// DeletedTimeGTE applies the GTE predicate on the "deleted_time" field.
func DeletedTimeGTE(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldGTE(FieldDeletedTime, v))
}

// DeletedTimeLT applies the LT predicate on the "deleted_time" field.
func DeletedTimeLT(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldLT(FieldDeletedTime, v))
}

// DeletedTimeLTE applies the LTE predicate on the "deleted_time" field.
func DeletedTimeLTE(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldLTE(FieldDeletedTime, v))
}

// DeletedTimeIsNil applies the IsNil predicate on the "deleted_time" field.
func DeletedTimeIsNil() predicate.Major {
	return predicate.Major(sql.FieldIsNull(FieldDeletedTime))
}

// DeletedTimeNotNil applies the NotNil predicate on the "deleted_time" field.
func DeletedTimeNotNil() predicate.Major {
	return predicate.Major(sql.FieldNotNull(FieldDeletedTime))
}

// ModifiedTimeEQ applies the EQ predicate on the "modified_time" field.
func ModifiedTimeEQ(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldEQ(FieldModifiedTime, v))
}

// ModifiedTimeNEQ applies the NEQ predicate on the "modified_time" field.
func ModifiedTimeNEQ(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldNEQ(FieldModifiedTime, v))
}

// ModifiedTimeIn applies the In predicate on the "modified_time" field.
func ModifiedTimeIn(vs ...time.Time) predicate.Major {
	return predicate.Major(sql.FieldIn(FieldModifiedTime, vs...))
}

// ModifiedTimeNotIn applies the NotIn predicate on the "modified_time" field.
func ModifiedTimeNotIn(vs ...time.Time) predicate.Major {
	return predicate.Major(sql.FieldNotIn(FieldModifiedTime, vs...))
}

// ModifiedTimeGT applies the GT predicate on the "modified_time" field.
func ModifiedTimeGT(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldGT(FieldModifiedTime, v))
}

// ModifiedTimeGTE applies the GTE predicate on the "modified_time" field.
func ModifiedTimeGTE(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldGTE(FieldModifiedTime, v))
}

// ModifiedTimeLT applies the LT predicate on the "modified_time" field.
func ModifiedTimeLT(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldLT(FieldModifiedTime, v))
}

// ModifiedTimeLTE applies the LTE predicate on the "modified_time" field.
func ModifiedTimeLTE(v time.Time) predicate.Major {
	return predicate.Major(sql.FieldLTE(FieldModifiedTime, v))
}

// ModifiedTimeIsNil applies the IsNil predicate on the "modified_time" field.
func ModifiedTimeIsNil() predicate.Major {
	return predicate.Major(sql.FieldIsNull(FieldModifiedTime))
}

// ModifiedTimeNotNil applies the NotNil predicate on the "modified_time" field.
func ModifiedTimeNotNil() predicate.Major {
	return predicate.Major(sql.FieldNotNull(FieldModifiedTime))
}

// HasCollege applies the HasEdge predicate on the "college" edge.
func HasCollege() predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CollegeTable, CollegeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCollegeWith applies the HasEdge predicate on the "college" edge with a given conditions (other predicates).
func HasCollegeWith(preds ...predicate.College) predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
		step := newCollegeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClasses applies the HasEdge predicate on the "classes" edge.
func HasClasses() predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassesWith applies the HasEdge predicate on the "classes" edge with a given conditions (other predicates).
func HasClassesWith(preds ...predicate.Class) predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
		step := newClassesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Major) predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Major) predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
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
func Not(p predicate.Major) predicate.Major {
	return predicate.Major(func(s *sql.Selector) {
		p(s.Not())
	})
}