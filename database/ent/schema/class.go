package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

func (Class) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "class"},
	}
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Uint16("cid"),
		field.String("name"),
		field.Bool("is_disabled").
			Default(false),
		field.Time("created_time").Nillable().
			Default(time.Now),
		field.Time("deleted_time").
			Nillable().
			Optional(),
		field.Time("modified_time").
			Nillable().
			Optional(),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("college", College.Type).
			Ref("classes").
			Unique().
			Required().
			Field("cid"),
		// edge.To("students", Student.Type),
	}
}
