package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "project"},
	}
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Uint32("uid").
			Optional(),
		field.String("goal"),
		field.String("principle"),
		field.String("process_and_method"),
		field.String("step"),
		field.String("result_and_conclusion"),
		field.String("requirement"),
		field.Uint8("review_status").
			Default(0),
		field.Bool("is_disabled").
			Default(false),
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

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("projects").
			Unique().
			Field("uid"),
		edge.To("attachments", File.Type),
		edge.To("review_project", ReviewProject.Type),
	}
}
