package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/entity"
)

// ReviewProjectDetail holds the schema definition for the ReviewProjectDetail entity.
type ReviewProjectDetail struct {
	ent.Schema
}

func (ReviewProjectDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "review_project_detail"},
	}
}

// Fields of the ReviewProjectDetail.
func (ReviewProjectDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("review_project_id"),
		field.Uint8("order"),
		field.JSON("reviewer", &entity.Reviewer{}),
		field.JSON("executor", &entity.Executor{}),
		field.Uint8("typee"),
		field.Uint8("status"),
		field.Time("created_time").
			Nillable().
			Default(time.Now),
		field.Time("deleted_time").
			Nillable().
			Optional(),
		field.Time("modified_time").
			Nillable().
			Optional(),
	}
}

// Edges of the ReviewProjectDetail.
func (ReviewProjectDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("review_project", ReviewProject.Type).
			Ref("review_project_detail").
			Required().
			Unique().
			Field("review_project_id"),
	}
}
