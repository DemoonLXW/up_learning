package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Major holds the schema definition for the Major entity.
type Major struct {
	ent.Schema
}

func (Major) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "major"},
	}
}

// Fields of the Major.
func (Major) Fields() []ent.Field {
	return []ent.Field{
		field.Uint16("id").
			Unique(),
		field.Uint8("cid"),
		field.String("name"),
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

// Edges of the Major.
func (Major) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("college", College.Type).
			Ref("majors").
			Unique().
			Required().
			Field("cid"),
		edge.To("classes", Class.Type),
	}
}
