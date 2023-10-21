// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldTitle, v))
}

// Goal applies equality check predicate on the "goal" field. It's identical to GoalEQ.
func Goal(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldGoal, v))
}

// Principle applies equality check predicate on the "principle" field. It's identical to PrincipleEQ.
func Principle(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldPrinciple, v))
}

// ProcessAndMethod applies equality check predicate on the "process_and_method" field. It's identical to ProcessAndMethodEQ.
func ProcessAndMethod(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldProcessAndMethod, v))
}

// Step applies equality check predicate on the "step" field. It's identical to StepEQ.
func Step(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStep, v))
}

// ResultAndConclusion applies equality check predicate on the "result_and_conclusion" field. It's identical to ResultAndConclusionEQ.
func ResultAndConclusion(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldResultAndConclusion, v))
}

// Requirement applies equality check predicate on the "requirement" field. It's identical to RequirementEQ.
func Requirement(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldRequirement, v))
}

// ReviewStatus applies equality check predicate on the "review_status" field. It's identical to ReviewStatusEQ.
func ReviewStatus(v uint8) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldReviewStatus, v))
}

// IsDisabled applies equality check predicate on the "is_disabled" field. It's identical to IsDisabledEQ.
func IsDisabled(v bool) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldIsDisabled, v))
}

// CreatedTime applies equality check predicate on the "created_time" field. It's identical to CreatedTimeEQ.
func CreatedTime(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedTime, v))
}

// DeletedTime applies equality check predicate on the "deleted_time" field. It's identical to DeletedTimeEQ.
func DeletedTime(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedTime, v))
}

// ModifiedTime applies equality check predicate on the "modified_time" field. It's identical to ModifiedTimeEQ.
func ModifiedTime(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldModifiedTime, v))
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUID, v))
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v uint32) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUID, v))
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUID, vs...))
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...uint32) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUID, vs...))
}

// UIDIsNil applies the IsNil predicate on the "uid" field.
func UIDIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldUID))
}

// UIDNotNil applies the NotNil predicate on the "uid" field.
func UIDNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldUID))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldTitle, v))
}

// GoalEQ applies the EQ predicate on the "goal" field.
func GoalEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldGoal, v))
}

// GoalNEQ applies the NEQ predicate on the "goal" field.
func GoalNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldGoal, v))
}

// GoalIn applies the In predicate on the "goal" field.
func GoalIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldGoal, vs...))
}

// GoalNotIn applies the NotIn predicate on the "goal" field.
func GoalNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldGoal, vs...))
}

// GoalGT applies the GT predicate on the "goal" field.
func GoalGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldGoal, v))
}

// GoalGTE applies the GTE predicate on the "goal" field.
func GoalGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldGoal, v))
}

// GoalLT applies the LT predicate on the "goal" field.
func GoalLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldGoal, v))
}

// GoalLTE applies the LTE predicate on the "goal" field.
func GoalLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldGoal, v))
}

// GoalContains applies the Contains predicate on the "goal" field.
func GoalContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldGoal, v))
}

// GoalHasPrefix applies the HasPrefix predicate on the "goal" field.
func GoalHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldGoal, v))
}

// GoalHasSuffix applies the HasSuffix predicate on the "goal" field.
func GoalHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldGoal, v))
}

// GoalEqualFold applies the EqualFold predicate on the "goal" field.
func GoalEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldGoal, v))
}

// GoalContainsFold applies the ContainsFold predicate on the "goal" field.
func GoalContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldGoal, v))
}

// PrincipleEQ applies the EQ predicate on the "principle" field.
func PrincipleEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldPrinciple, v))
}

// PrincipleNEQ applies the NEQ predicate on the "principle" field.
func PrincipleNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldPrinciple, v))
}

// PrincipleIn applies the In predicate on the "principle" field.
func PrincipleIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldPrinciple, vs...))
}

// PrincipleNotIn applies the NotIn predicate on the "principle" field.
func PrincipleNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldPrinciple, vs...))
}

// PrincipleGT applies the GT predicate on the "principle" field.
func PrincipleGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldPrinciple, v))
}

// PrincipleGTE applies the GTE predicate on the "principle" field.
func PrincipleGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldPrinciple, v))
}

// PrincipleLT applies the LT predicate on the "principle" field.
func PrincipleLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldPrinciple, v))
}

// PrincipleLTE applies the LTE predicate on the "principle" field.
func PrincipleLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldPrinciple, v))
}

// PrincipleContains applies the Contains predicate on the "principle" field.
func PrincipleContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldPrinciple, v))
}

// PrincipleHasPrefix applies the HasPrefix predicate on the "principle" field.
func PrincipleHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldPrinciple, v))
}

// PrincipleHasSuffix applies the HasSuffix predicate on the "principle" field.
func PrincipleHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldPrinciple, v))
}

// PrincipleEqualFold applies the EqualFold predicate on the "principle" field.
func PrincipleEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldPrinciple, v))
}

// PrincipleContainsFold applies the ContainsFold predicate on the "principle" field.
func PrincipleContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldPrinciple, v))
}

// ProcessAndMethodEQ applies the EQ predicate on the "process_and_method" field.
func ProcessAndMethodEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldProcessAndMethod, v))
}

// ProcessAndMethodNEQ applies the NEQ predicate on the "process_and_method" field.
func ProcessAndMethodNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldProcessAndMethod, v))
}

// ProcessAndMethodIn applies the In predicate on the "process_and_method" field.
func ProcessAndMethodIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldProcessAndMethod, vs...))
}

// ProcessAndMethodNotIn applies the NotIn predicate on the "process_and_method" field.
func ProcessAndMethodNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldProcessAndMethod, vs...))
}

// ProcessAndMethodGT applies the GT predicate on the "process_and_method" field.
func ProcessAndMethodGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldProcessAndMethod, v))
}

// ProcessAndMethodGTE applies the GTE predicate on the "process_and_method" field.
func ProcessAndMethodGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldProcessAndMethod, v))
}

// ProcessAndMethodLT applies the LT predicate on the "process_and_method" field.
func ProcessAndMethodLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldProcessAndMethod, v))
}

// ProcessAndMethodLTE applies the LTE predicate on the "process_and_method" field.
func ProcessAndMethodLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldProcessAndMethod, v))
}

// ProcessAndMethodContains applies the Contains predicate on the "process_and_method" field.
func ProcessAndMethodContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldProcessAndMethod, v))
}

// ProcessAndMethodHasPrefix applies the HasPrefix predicate on the "process_and_method" field.
func ProcessAndMethodHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldProcessAndMethod, v))
}

// ProcessAndMethodHasSuffix applies the HasSuffix predicate on the "process_and_method" field.
func ProcessAndMethodHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldProcessAndMethod, v))
}

// ProcessAndMethodEqualFold applies the EqualFold predicate on the "process_and_method" field.
func ProcessAndMethodEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldProcessAndMethod, v))
}

// ProcessAndMethodContainsFold applies the ContainsFold predicate on the "process_and_method" field.
func ProcessAndMethodContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldProcessAndMethod, v))
}

// StepEQ applies the EQ predicate on the "step" field.
func StepEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldStep, v))
}

// StepNEQ applies the NEQ predicate on the "step" field.
func StepNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldStep, v))
}

// StepIn applies the In predicate on the "step" field.
func StepIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldStep, vs...))
}

// StepNotIn applies the NotIn predicate on the "step" field.
func StepNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldStep, vs...))
}

// StepGT applies the GT predicate on the "step" field.
func StepGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldStep, v))
}

// StepGTE applies the GTE predicate on the "step" field.
func StepGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldStep, v))
}

// StepLT applies the LT predicate on the "step" field.
func StepLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldStep, v))
}

// StepLTE applies the LTE predicate on the "step" field.
func StepLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldStep, v))
}

// StepContains applies the Contains predicate on the "step" field.
func StepContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldStep, v))
}

// StepHasPrefix applies the HasPrefix predicate on the "step" field.
func StepHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldStep, v))
}

// StepHasSuffix applies the HasSuffix predicate on the "step" field.
func StepHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldStep, v))
}

// StepEqualFold applies the EqualFold predicate on the "step" field.
func StepEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldStep, v))
}

// StepContainsFold applies the ContainsFold predicate on the "step" field.
func StepContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldStep, v))
}

// ResultAndConclusionEQ applies the EQ predicate on the "result_and_conclusion" field.
func ResultAndConclusionEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldResultAndConclusion, v))
}

// ResultAndConclusionNEQ applies the NEQ predicate on the "result_and_conclusion" field.
func ResultAndConclusionNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldResultAndConclusion, v))
}

// ResultAndConclusionIn applies the In predicate on the "result_and_conclusion" field.
func ResultAndConclusionIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldResultAndConclusion, vs...))
}

// ResultAndConclusionNotIn applies the NotIn predicate on the "result_and_conclusion" field.
func ResultAndConclusionNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldResultAndConclusion, vs...))
}

// ResultAndConclusionGT applies the GT predicate on the "result_and_conclusion" field.
func ResultAndConclusionGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldResultAndConclusion, v))
}

// ResultAndConclusionGTE applies the GTE predicate on the "result_and_conclusion" field.
func ResultAndConclusionGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldResultAndConclusion, v))
}

// ResultAndConclusionLT applies the LT predicate on the "result_and_conclusion" field.
func ResultAndConclusionLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldResultAndConclusion, v))
}

// ResultAndConclusionLTE applies the LTE predicate on the "result_and_conclusion" field.
func ResultAndConclusionLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldResultAndConclusion, v))
}

// ResultAndConclusionContains applies the Contains predicate on the "result_and_conclusion" field.
func ResultAndConclusionContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldResultAndConclusion, v))
}

// ResultAndConclusionHasPrefix applies the HasPrefix predicate on the "result_and_conclusion" field.
func ResultAndConclusionHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldResultAndConclusion, v))
}

// ResultAndConclusionHasSuffix applies the HasSuffix predicate on the "result_and_conclusion" field.
func ResultAndConclusionHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldResultAndConclusion, v))
}

// ResultAndConclusionEqualFold applies the EqualFold predicate on the "result_and_conclusion" field.
func ResultAndConclusionEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldResultAndConclusion, v))
}

// ResultAndConclusionContainsFold applies the ContainsFold predicate on the "result_and_conclusion" field.
func ResultAndConclusionContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldResultAndConclusion, v))
}

// RequirementEQ applies the EQ predicate on the "requirement" field.
func RequirementEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldRequirement, v))
}

// RequirementNEQ applies the NEQ predicate on the "requirement" field.
func RequirementNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldRequirement, v))
}

// RequirementIn applies the In predicate on the "requirement" field.
func RequirementIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldRequirement, vs...))
}

// RequirementNotIn applies the NotIn predicate on the "requirement" field.
func RequirementNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldRequirement, vs...))
}

// RequirementGT applies the GT predicate on the "requirement" field.
func RequirementGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldRequirement, v))
}

// RequirementGTE applies the GTE predicate on the "requirement" field.
func RequirementGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldRequirement, v))
}

// RequirementLT applies the LT predicate on the "requirement" field.
func RequirementLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldRequirement, v))
}

// RequirementLTE applies the LTE predicate on the "requirement" field.
func RequirementLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldRequirement, v))
}

// RequirementContains applies the Contains predicate on the "requirement" field.
func RequirementContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldRequirement, v))
}

// RequirementHasPrefix applies the HasPrefix predicate on the "requirement" field.
func RequirementHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldRequirement, v))
}

// RequirementHasSuffix applies the HasSuffix predicate on the "requirement" field.
func RequirementHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldRequirement, v))
}

// RequirementEqualFold applies the EqualFold predicate on the "requirement" field.
func RequirementEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldRequirement, v))
}

// RequirementContainsFold applies the ContainsFold predicate on the "requirement" field.
func RequirementContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldRequirement, v))
}

// ReviewStatusEQ applies the EQ predicate on the "review_status" field.
func ReviewStatusEQ(v uint8) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldReviewStatus, v))
}

// ReviewStatusNEQ applies the NEQ predicate on the "review_status" field.
func ReviewStatusNEQ(v uint8) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldReviewStatus, v))
}

// ReviewStatusIn applies the In predicate on the "review_status" field.
func ReviewStatusIn(vs ...uint8) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldReviewStatus, vs...))
}

// ReviewStatusNotIn applies the NotIn predicate on the "review_status" field.
func ReviewStatusNotIn(vs ...uint8) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldReviewStatus, vs...))
}

// ReviewStatusGT applies the GT predicate on the "review_status" field.
func ReviewStatusGT(v uint8) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldReviewStatus, v))
}

// ReviewStatusGTE applies the GTE predicate on the "review_status" field.
func ReviewStatusGTE(v uint8) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldReviewStatus, v))
}

// ReviewStatusLT applies the LT predicate on the "review_status" field.
func ReviewStatusLT(v uint8) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldReviewStatus, v))
}

// ReviewStatusLTE applies the LTE predicate on the "review_status" field.
func ReviewStatusLTE(v uint8) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldReviewStatus, v))
}

// IsDisabledEQ applies the EQ predicate on the "is_disabled" field.
func IsDisabledEQ(v bool) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldIsDisabled, v))
}

// IsDisabledNEQ applies the NEQ predicate on the "is_disabled" field.
func IsDisabledNEQ(v bool) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldIsDisabled, v))
}

// CreatedTimeEQ applies the EQ predicate on the "created_time" field.
func CreatedTimeEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedTime, v))
}

// CreatedTimeNEQ applies the NEQ predicate on the "created_time" field.
func CreatedTimeNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedTime, v))
}

// CreatedTimeIn applies the In predicate on the "created_time" field.
func CreatedTimeIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedTime, vs...))
}

// CreatedTimeNotIn applies the NotIn predicate on the "created_time" field.
func CreatedTimeNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedTime, vs...))
}

// CreatedTimeGT applies the GT predicate on the "created_time" field.
func CreatedTimeGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedTime, v))
}

// CreatedTimeGTE applies the GTE predicate on the "created_time" field.
func CreatedTimeGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedTime, v))
}

// CreatedTimeLT applies the LT predicate on the "created_time" field.
func CreatedTimeLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedTime, v))
}

// CreatedTimeLTE applies the LTE predicate on the "created_time" field.
func CreatedTimeLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedTime, v))
}

// DeletedTimeEQ applies the EQ predicate on the "deleted_time" field.
func DeletedTimeEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldDeletedTime, v))
}

// DeletedTimeNEQ applies the NEQ predicate on the "deleted_time" field.
func DeletedTimeNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldDeletedTime, v))
}

// DeletedTimeIn applies the In predicate on the "deleted_time" field.
func DeletedTimeIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldDeletedTime, vs...))
}

// DeletedTimeNotIn applies the NotIn predicate on the "deleted_time" field.
func DeletedTimeNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldDeletedTime, vs...))
}

// DeletedTimeGT applies the GT predicate on the "deleted_time" field.
func DeletedTimeGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldDeletedTime, v))
}

// DeletedTimeGTE applies the GTE predicate on the "deleted_time" field.
func DeletedTimeGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldDeletedTime, v))
}

// DeletedTimeLT applies the LT predicate on the "deleted_time" field.
func DeletedTimeLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldDeletedTime, v))
}

// DeletedTimeLTE applies the LTE predicate on the "deleted_time" field.
func DeletedTimeLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldDeletedTime, v))
}

// DeletedTimeIsNil applies the IsNil predicate on the "deleted_time" field.
func DeletedTimeIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldDeletedTime))
}

// DeletedTimeNotNil applies the NotNil predicate on the "deleted_time" field.
func DeletedTimeNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldDeletedTime))
}

// ModifiedTimeEQ applies the EQ predicate on the "modified_time" field.
func ModifiedTimeEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldModifiedTime, v))
}

// ModifiedTimeNEQ applies the NEQ predicate on the "modified_time" field.
func ModifiedTimeNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldModifiedTime, v))
}

// ModifiedTimeIn applies the In predicate on the "modified_time" field.
func ModifiedTimeIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldModifiedTime, vs...))
}

// ModifiedTimeNotIn applies the NotIn predicate on the "modified_time" field.
func ModifiedTimeNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldModifiedTime, vs...))
}

// ModifiedTimeGT applies the GT predicate on the "modified_time" field.
func ModifiedTimeGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldModifiedTime, v))
}

// ModifiedTimeGTE applies the GTE predicate on the "modified_time" field.
func ModifiedTimeGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldModifiedTime, v))
}

// ModifiedTimeLT applies the LT predicate on the "modified_time" field.
func ModifiedTimeLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldModifiedTime, v))
}

// ModifiedTimeLTE applies the LTE predicate on the "modified_time" field.
func ModifiedTimeLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldModifiedTime, v))
}

// ModifiedTimeIsNil applies the IsNil predicate on the "modified_time" field.
func ModifiedTimeIsNil() predicate.Project {
	return predicate.Project(sql.FieldIsNull(FieldModifiedTime))
}

// ModifiedTimeNotNil applies the NotNil predicate on the "modified_time" field.
func ModifiedTimeNotNil() predicate.Project {
	return predicate.Project(sql.FieldNotNull(FieldModifiedTime))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttachments applies the HasEdge predicate on the "attachments" edge.
func HasAttachments() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AttachmentsTable, AttachmentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentsWith applies the HasEdge predicate on the "attachments" edge with a given conditions (other predicates).
func HasAttachmentsWith(preds ...predicate.File) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newAttachmentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReviewProject applies the HasEdge predicate on the "review_project" edge.
func HasReviewProject() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReviewProjectTable, ReviewProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewProjectWith applies the HasEdge predicate on the "review_project" edge with a given conditions (other predicates).
func HasReviewProjectWith(preds ...predicate.ReviewProject) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newReviewProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProjectFile applies the HasEdge predicate on the "project_file" edge.
func HasProjectFile() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ProjectFileTable, ProjectFileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectFileWith applies the HasEdge predicate on the "project_file" edge with a given conditions (other predicates).
func HasProjectFileWith(preds ...predicate.ProjectFile) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := newProjectFileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
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
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		p(s.Not())
	})
}
