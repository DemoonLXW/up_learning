// Code generated by ent, DO NOT EDIT.

package samplefile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint8) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLTE(FieldID, id))
}

// Fid applies equality check predicate on the "fid" field. It's identical to FidEQ.
func Fid(v uint32) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldFid, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldType, v))
}

// IsDisabled applies equality check predicate on the "is_disabled" field. It's identical to IsDisabledEQ.
func IsDisabled(v bool) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldIsDisabled, v))
}

// CreatedTime applies equality check predicate on the "created_time" field. It's identical to CreatedTimeEQ.
func CreatedTime(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldCreatedTime, v))
}

// DeletedTime applies equality check predicate on the "deleted_time" field. It's identical to DeletedTimeEQ.
func DeletedTime(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldDeletedTime, v))
}

// ModifiedTime applies equality check predicate on the "modified_time" field. It's identical to ModifiedTimeEQ.
func ModifiedTime(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldModifiedTime, v))
}

// FidEQ applies the EQ predicate on the "fid" field.
func FidEQ(v uint32) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldFid, v))
}

// FidNEQ applies the NEQ predicate on the "fid" field.
func FidNEQ(v uint32) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldFid, v))
}

// FidIn applies the In predicate on the "fid" field.
func FidIn(vs ...uint32) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIn(FieldFid, vs...))
}

// FidNotIn applies the NotIn predicate on the "fid" field.
func FidNotIn(vs ...uint32) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotIn(FieldFid, vs...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldContainsFold(FieldType, v))
}

// IsDisabledEQ applies the EQ predicate on the "is_disabled" field.
func IsDisabledEQ(v bool) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldIsDisabled, v))
}

// IsDisabledNEQ applies the NEQ predicate on the "is_disabled" field.
func IsDisabledNEQ(v bool) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldIsDisabled, v))
}

// CreatedTimeEQ applies the EQ predicate on the "created_time" field.
func CreatedTimeEQ(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldCreatedTime, v))
}

// CreatedTimeNEQ applies the NEQ predicate on the "created_time" field.
func CreatedTimeNEQ(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldCreatedTime, v))
}

// CreatedTimeIn applies the In predicate on the "created_time" field.
func CreatedTimeIn(vs ...time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIn(FieldCreatedTime, vs...))
}

// CreatedTimeNotIn applies the NotIn predicate on the "created_time" field.
func CreatedTimeNotIn(vs ...time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotIn(FieldCreatedTime, vs...))
}

// CreatedTimeGT applies the GT predicate on the "created_time" field.
func CreatedTimeGT(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGT(FieldCreatedTime, v))
}

// CreatedTimeGTE applies the GTE predicate on the "created_time" field.
func CreatedTimeGTE(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGTE(FieldCreatedTime, v))
}

// CreatedTimeLT applies the LT predicate on the "created_time" field.
func CreatedTimeLT(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLT(FieldCreatedTime, v))
}

// CreatedTimeLTE applies the LTE predicate on the "created_time" field.
func CreatedTimeLTE(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLTE(FieldCreatedTime, v))
}

// DeletedTimeEQ applies the EQ predicate on the "deleted_time" field.
func DeletedTimeEQ(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldDeletedTime, v))
}

// DeletedTimeNEQ applies the NEQ predicate on the "deleted_time" field.
func DeletedTimeNEQ(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldDeletedTime, v))
}

// DeletedTimeIn applies the In predicate on the "deleted_time" field.
func DeletedTimeIn(vs ...time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIn(FieldDeletedTime, vs...))
}

// DeletedTimeNotIn applies the NotIn predicate on the "deleted_time" field.
func DeletedTimeNotIn(vs ...time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotIn(FieldDeletedTime, vs...))
}

// DeletedTimeGT applies the GT predicate on the "deleted_time" field.
func DeletedTimeGT(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGT(FieldDeletedTime, v))
}

// DeletedTimeGTE applies the GTE predicate on the "deleted_time" field.
func DeletedTimeGTE(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGTE(FieldDeletedTime, v))
}

// DeletedTimeLT applies the LT predicate on the "deleted_time" field.
func DeletedTimeLT(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLT(FieldDeletedTime, v))
}

// DeletedTimeLTE applies the LTE predicate on the "deleted_time" field.
func DeletedTimeLTE(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLTE(FieldDeletedTime, v))
}

// DeletedTimeIsNil applies the IsNil predicate on the "deleted_time" field.
func DeletedTimeIsNil() predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIsNull(FieldDeletedTime))
}

// DeletedTimeNotNil applies the NotNil predicate on the "deleted_time" field.
func DeletedTimeNotNil() predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotNull(FieldDeletedTime))
}

// ModifiedTimeEQ applies the EQ predicate on the "modified_time" field.
func ModifiedTimeEQ(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldEQ(FieldModifiedTime, v))
}

// ModifiedTimeNEQ applies the NEQ predicate on the "modified_time" field.
func ModifiedTimeNEQ(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNEQ(FieldModifiedTime, v))
}

// ModifiedTimeIn applies the In predicate on the "modified_time" field.
func ModifiedTimeIn(vs ...time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIn(FieldModifiedTime, vs...))
}

// ModifiedTimeNotIn applies the NotIn predicate on the "modified_time" field.
func ModifiedTimeNotIn(vs ...time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotIn(FieldModifiedTime, vs...))
}

// ModifiedTimeGT applies the GT predicate on the "modified_time" field.
func ModifiedTimeGT(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGT(FieldModifiedTime, v))
}

// ModifiedTimeGTE applies the GTE predicate on the "modified_time" field.
func ModifiedTimeGTE(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldGTE(FieldModifiedTime, v))
}

// ModifiedTimeLT applies the LT predicate on the "modified_time" field.
func ModifiedTimeLT(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLT(FieldModifiedTime, v))
}

// ModifiedTimeLTE applies the LTE predicate on the "modified_time" field.
func ModifiedTimeLTE(v time.Time) predicate.SampleFile {
	return predicate.SampleFile(sql.FieldLTE(FieldModifiedTime, v))
}

// ModifiedTimeIsNil applies the IsNil predicate on the "modified_time" field.
func ModifiedTimeIsNil() predicate.SampleFile {
	return predicate.SampleFile(sql.FieldIsNull(FieldModifiedTime))
}

// ModifiedTimeNotNil applies the NotNil predicate on the "modified_time" field.
func ModifiedTimeNotNil() predicate.SampleFile {
	return predicate.SampleFile(sql.FieldNotNull(FieldModifiedTime))
}

// HasFile applies the HasEdge predicate on the "file" edge.
func HasFile() predicate.SampleFile {
	return predicate.SampleFile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFileWith applies the HasEdge predicate on the "file" edge with a given conditions (other predicates).
func HasFileWith(preds ...predicate.File) predicate.SampleFile {
	return predicate.SampleFile(func(s *sql.Selector) {
		step := newFileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SampleFile) predicate.SampleFile {
	return predicate.SampleFile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SampleFile) predicate.SampleFile {
	return predicate.SampleFile(func(s *sql.Selector) {
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
func Not(p predicate.SampleFile) predicate.SampleFile {
	return predicate.SampleFile(func(s *sql.Selector) {
		p(s.Not())
	})
}
