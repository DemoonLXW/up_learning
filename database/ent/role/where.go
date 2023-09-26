// Code generated by ent, DO NOT EDIT.

package role

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint8) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint8) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint8) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint8) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint8) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint8) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint8) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint8) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint8) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDescription, v))
}

// IsDisabled applies equality check predicate on the "is_disabled" field. It's identical to IsDisabledEQ.
func IsDisabled(v bool) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldIsDisabled, v))
}

// CreatedTime applies equality check predicate on the "created_time" field. It's identical to CreatedTimeEQ.
func CreatedTime(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedTime, v))
}

// ModifiedTime applies equality check predicate on the "modified_time" field. It's identical to ModifiedTimeEQ.
func ModifiedTime(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldModifiedTime, v))
}

// DeletedTime applies equality check predicate on the "deleted_time" field. It's identical to DeletedTimeEQ.
func DeletedTime(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDeletedTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Role {
	return predicate.Role(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Role {
	return predicate.Role(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldDescription, v))
}

// IsDisabledEQ applies the EQ predicate on the "is_disabled" field.
func IsDisabledEQ(v bool) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldIsDisabled, v))
}

// IsDisabledNEQ applies the NEQ predicate on the "is_disabled" field.
func IsDisabledNEQ(v bool) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldIsDisabled, v))
}

// CreatedTimeEQ applies the EQ predicate on the "created_time" field.
func CreatedTimeEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedTime, v))
}

// CreatedTimeNEQ applies the NEQ predicate on the "created_time" field.
func CreatedTimeNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldCreatedTime, v))
}

// CreatedTimeIn applies the In predicate on the "created_time" field.
func CreatedTimeIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldCreatedTime, vs...))
}

// CreatedTimeNotIn applies the NotIn predicate on the "created_time" field.
func CreatedTimeNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldCreatedTime, vs...))
}

// CreatedTimeGT applies the GT predicate on the "created_time" field.
func CreatedTimeGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldCreatedTime, v))
}

// CreatedTimeGTE applies the GTE predicate on the "created_time" field.
func CreatedTimeGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldCreatedTime, v))
}

// CreatedTimeLT applies the LT predicate on the "created_time" field.
func CreatedTimeLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldCreatedTime, v))
}

// CreatedTimeLTE applies the LTE predicate on the "created_time" field.
func CreatedTimeLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldCreatedTime, v))
}

// ModifiedTimeEQ applies the EQ predicate on the "modified_time" field.
func ModifiedTimeEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldModifiedTime, v))
}

// ModifiedTimeNEQ applies the NEQ predicate on the "modified_time" field.
func ModifiedTimeNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldModifiedTime, v))
}

// ModifiedTimeIn applies the In predicate on the "modified_time" field.
func ModifiedTimeIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldModifiedTime, vs...))
}

// ModifiedTimeNotIn applies the NotIn predicate on the "modified_time" field.
func ModifiedTimeNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldModifiedTime, vs...))
}

// ModifiedTimeGT applies the GT predicate on the "modified_time" field.
func ModifiedTimeGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldModifiedTime, v))
}

// ModifiedTimeGTE applies the GTE predicate on the "modified_time" field.
func ModifiedTimeGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldModifiedTime, v))
}

// ModifiedTimeLT applies the LT predicate on the "modified_time" field.
func ModifiedTimeLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldModifiedTime, v))
}

// ModifiedTimeLTE applies the LTE predicate on the "modified_time" field.
func ModifiedTimeLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldModifiedTime, v))
}

// ModifiedTimeIsNil applies the IsNil predicate on the "modified_time" field.
func ModifiedTimeIsNil() predicate.Role {
	return predicate.Role(sql.FieldIsNull(FieldModifiedTime))
}

// ModifiedTimeNotNil applies the NotNil predicate on the "modified_time" field.
func ModifiedTimeNotNil() predicate.Role {
	return predicate.Role(sql.FieldNotNull(FieldModifiedTime))
}

// DeletedTimeEQ applies the EQ predicate on the "deleted_time" field.
func DeletedTimeEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDeletedTime, v))
}

// DeletedTimeNEQ applies the NEQ predicate on the "deleted_time" field.
func DeletedTimeNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldDeletedTime, v))
}

// DeletedTimeIn applies the In predicate on the "deleted_time" field.
func DeletedTimeIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldDeletedTime, vs...))
}

// DeletedTimeNotIn applies the NotIn predicate on the "deleted_time" field.
func DeletedTimeNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldDeletedTime, vs...))
}

// DeletedTimeGT applies the GT predicate on the "deleted_time" field.
func DeletedTimeGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldDeletedTime, v))
}

// DeletedTimeGTE applies the GTE predicate on the "deleted_time" field.
func DeletedTimeGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldDeletedTime, v))
}

// DeletedTimeLT applies the LT predicate on the "deleted_time" field.
func DeletedTimeLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldDeletedTime, v))
}

// DeletedTimeLTE applies the LTE predicate on the "deleted_time" field.
func DeletedTimeLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldDeletedTime, v))
}

// DeletedTimeIsNil applies the IsNil predicate on the "deleted_time" field.
func DeletedTimeIsNil() predicate.Role {
	return predicate.Role(sql.FieldIsNull(FieldDeletedTime))
}

// DeletedTimeNotNil applies the NotNil predicate on the "deleted_time" field.
func DeletedTimeNotNil() predicate.Role {
	return predicate.Role(sql.FieldNotNull(FieldDeletedTime))
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMenu applies the HasEdge predicate on the "menu" edge.
func HasMenu() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MenuTable, MenuColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenuWith applies the HasEdge predicate on the "menu" edge with a given conditions (other predicates).
func HasMenuWith(preds ...predicate.Menu) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newMenuStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRolePermission applies the HasEdge predicate on the "role_permission" edge.
func HasRolePermission() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RolePermissionTable, RolePermissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolePermissionWith applies the HasEdge predicate on the "role_permission" edge with a given conditions (other predicates).
func HasRolePermissionWith(preds ...predicate.RolePermission) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newRolePermissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserRole applies the HasEdge predicate on the "user_role" edge.
func HasUserRole() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserRoleTable, UserRoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserRoleWith applies the HasEdge predicate on the "user_role" edge with a given conditions (other predicates).
func HasUserRoleWith(preds ...predicate.UserRole) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newUserRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
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
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		p(s.Not())
	})
}
