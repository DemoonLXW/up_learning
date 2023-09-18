package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// College holds the schema definition for the College entity.
type College struct {
	ent.Schema
}

func (College) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "college"},
	}
}

// Fields of the College.
func (College) Fields() []ent.Field {
	return []ent.Field{
		field.Uint16("id").
			Unique(),
		field.Uint16("sid"),
		field.String("name"),
		field.Bool("is_disabled").
			Default(false),
		field.Time("created_time").
			Default(time.Now),
		field.Time("deleted_time").
			Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		field.Time("modified_time").
			Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	}
}

// Edges of the College.
func (College) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("school", School.Type).
			Ref("colleges").
			Unique().
			Required().
			Field("sid"),
		edge.To("classes", Class.Type),
	}
}
