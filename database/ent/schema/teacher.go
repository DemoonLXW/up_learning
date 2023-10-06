package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

func (Teacher) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "teacher"},
	}
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Uint32("uid").
			Unique().
			Optional(),
		field.Uint8("cid"),
		field.String("teacher_id").
			Unique().
			NotEmpty(),
		field.String("name").
			NotEmpty(),
		field.Uint8("gender"),
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

// Edges of the Teacher.
func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("college", College.Type).
			Ref("teachers").
			Unique().
			Required().
			Field("cid"),
		edge.From("user", User.Type).
			Ref("teacher").
			Unique().
			Field("uid"),
	}
}
