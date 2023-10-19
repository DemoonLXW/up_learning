package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReviewProject holds the schema definition for the ReviewProject entity.
type ReviewProject struct {
	ent.Schema
}

func (ReviewProject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "review_project"},
	}
}

// Fields of the ReviewProject.
func (ReviewProject) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Uint32("project_id"),
		field.String("workflow_id"),
		field.String("run_id"),
		field.Uint32("applicant_id"),
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

// Edges of the ReviewProject.
func (ReviewProject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applicant", User.Type).
			Ref("review_project").
			Required().
			Unique().
			Field("applicant_id"),
		edge.From("project", Project.Type).
			Ref("review_project").
			Required().
			Unique().
			Field("project_id"),
		edge.To("review_project_detail", ReviewProjectDetail.Type),
	}
}
